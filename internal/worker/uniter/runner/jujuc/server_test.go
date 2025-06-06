// Copyright 2012, 2013 Canonical Ltd.
// Copyright 2014 Cloudbase Solutions SRL
// Licensed under the AGPLv3, see LICENCE file for details.

package jujuc_test

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	stdtesting "testing"
	"time"

	"github.com/juju/gnuflag"
	"github.com/juju/tc"
	"github.com/juju/utils/v4/exec"

	jujucmd "github.com/juju/juju/cmd"
	"github.com/juju/juju/internal/cmd"
	"github.com/juju/juju/internal/testing"
	"github.com/juju/juju/internal/worker/uniter/runner/jujuc"
	"github.com/juju/juju/juju/sockets"
)

type RpcCommand struct {
	cmd.CommandBase
	Value string
	Slow  bool
	Echo  bool
}

func (c *RpcCommand) Info() *cmd.Info {
	return jujucmd.Info(&cmd.Info{
		Name:    "remote",
		Purpose: "act at a distance",
		Doc:     "blah doc",
	})
}

func (c *RpcCommand) SetFlags(f *gnuflag.FlagSet) {
	f.StringVar(&c.Value, "value", "", "doc")
	f.BoolVar(&c.Slow, "slow", false, "doc")
	f.BoolVar(&c.Echo, "echo", false, "doc")
}

func (c *RpcCommand) Init(args []string) error {
	return cmd.CheckEmpty(args)
}

func (c *RpcCommand) Run(ctx *cmd.Context) error {
	if c.Value == "error" {
		return errors.New("blam")
	}
	if c.Slow {
		time.Sleep(testing.ShortWait)
		return nil
	}
	if c.Echo {
		if _, err := io.Copy(ctx.Stdout, ctx.Stdin); err != nil {
			return err
		}
	}
	ctx.Stdout.Write([]byte("eye of newt\n"))
	ctx.Stderr.Write([]byte("toe of frog\n"))
	return os.WriteFile(ctx.AbsPath("local"), []byte(c.Value), 0644)
}

func factory(contextId, cmdName string) (cmd.Command, error) {
	if contextId != "validCtx" {
		return nil, fmt.Errorf("unknown context %q", contextId)
	}
	if cmdName != "remote" {
		return nil, fmt.Errorf("unknown command %q", cmdName)
	}
	return &RpcCommand{}, nil
}

type ServerSuite struct {
	testing.BaseSuite
	server *jujuc.Server
	socket sockets.Socket
	err    chan error
}

func TestServerSuite(t *stdtesting.T) {
	tc.Run(t, &ServerSuite{})
}

func (s *ServerSuite) osDependentSockPath(c *tc.C) sockets.Socket {
	pipeRoot := c.MkDir()
	sock := filepath.Join(pipeRoot, "test.sock")
	return sockets.Socket{Network: "unix", Address: sock}
}

func (s *ServerSuite) SetUpTest(c *tc.C) {
	s.BaseSuite.SetUpTest(c)
	s.socket = s.osDependentSockPath(c)
	srv, err := jujuc.NewServer(factory, s.socket)
	c.Assert(err, tc.ErrorIsNil)
	c.Assert(srv, tc.NotNil)
	s.server = srv
	s.err = make(chan error)
	go func() { s.err <- s.server.Run() }()
}

func (s *ServerSuite) TearDownTest(c *tc.C) {
	s.server.Close()
	c.Assert(<-s.err, tc.IsNil)
	_, err := os.Open(s.socket.Address)
	c.Assert(err, tc.Satisfies, os.IsNotExist)
	s.BaseSuite.TearDownTest(c)
}

func (s *ServerSuite) Call(c *tc.C, req jujuc.Request) (resp exec.ExecResponse, err error) {
	client, err := sockets.Dial(s.socket)
	c.Assert(err, tc.ErrorIsNil)
	defer client.Close()
	err = client.Call("Jujuc.Main", req, &resp)
	return resp, err
}

func (s *ServerSuite) TestHappyPath(c *tc.C) {
	dir := c.MkDir()
	resp, err := s.Call(c, jujuc.Request{
		ContextId:   "validCtx",
		Dir:         dir,
		CommandName: "remote",
		Args:        []string{"--value", "something", "--echo"},
		StdinSet:    true,
		Stdin:       []byte("wool of bat\n"),
	})
	c.Assert(err, tc.ErrorIsNil)
	c.Assert(resp.Code, tc.Equals, 0)
	c.Assert(string(resp.Stdout), tc.Equals, "wool of bat\neye of newt\n")
	c.Assert(string(resp.Stderr), tc.Equals, "toe of frog\n")
	content, err := os.ReadFile(filepath.Join(dir, "local"))
	c.Assert(err, tc.ErrorIsNil)
	c.Assert(string(content), tc.Equals, "something")
}

func (s *ServerSuite) TestNoStdin(c *tc.C) {
	dir := c.MkDir()
	_, err := s.Call(c, jujuc.Request{
		ContextId:   "validCtx",
		Dir:         dir,
		CommandName: "remote",
		Args:        []string{"--echo"},
	})
	c.Assert(err, tc.ErrorMatches, jujuc.ErrNoStdin.Error())
}

func (s *ServerSuite) TestLocks(c *tc.C) {
	var wg sync.WaitGroup
	t0 := time.Now()
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			dir := c.MkDir()
			resp, err := s.Call(c, jujuc.Request{
				ContextId:   "validCtx",
				Dir:         dir,
				CommandName: "remote",
				Args:        []string{"--slow"},
			})
			c.Assert(err, tc.ErrorIsNil)
			c.Assert(resp.Code, tc.Equals, 0)
			wg.Done()
		}()
	}
	wg.Wait()
	t1 := time.Now()
	c.Assert(t0.Add(4*testing.ShortWait).Before(t1), tc.IsTrue)
}

func (s *ServerSuite) TestBadCommandName(c *tc.C) {
	dir := c.MkDir()
	_, err := s.Call(c, jujuc.Request{
		ContextId: "validCtx",
		Dir:       dir,
	})
	c.Assert(err, tc.ErrorMatches, "bad request: command not specified")
	_, err = s.Call(c, jujuc.Request{
		ContextId:   "validCtx",
		Dir:         dir,
		CommandName: "witchcraft",
	})
	c.Assert(err, tc.ErrorMatches, `bad request: unknown command "witchcraft"`)
}

func (s *ServerSuite) TestBadDir(c *tc.C) {
	for _, req := range []jujuc.Request{{
		ContextId:   "validCtx",
		CommandName: "anything",
	}, {
		ContextId:   "validCtx",
		Dir:         "foo/bar",
		CommandName: "anything",
	}} {
		_, err := s.Call(c, req)
		c.Assert(err, tc.ErrorMatches, "bad request: Dir is not absolute")
	}
}

func (s *ServerSuite) TestBadContextId(c *tc.C) {
	_, err := s.Call(c, jujuc.Request{
		ContextId:   "whatever",
		Dir:         c.MkDir(),
		CommandName: "remote",
	})
	c.Assert(err, tc.ErrorMatches, `bad request: unknown context "whatever"`)
}

func (s *ServerSuite) AssertBadCommand(c *tc.C, args []string, code int) exec.ExecResponse {
	resp, err := s.Call(c, jujuc.Request{
		ContextId:   "validCtx",
		Dir:         c.MkDir(),
		CommandName: args[0],
		Args:        args[1:],
	})
	c.Assert(err, tc.ErrorIsNil)
	c.Assert(resp.Code, tc.Equals, code)
	return resp
}

func (s *ServerSuite) TestParseError(c *tc.C) {
	resp := s.AssertBadCommand(c, []string{"remote", "--cheese"}, 2)
	c.Assert(string(resp.Stdout), tc.Equals, "")
	c.Assert(string(resp.Stderr), tc.Equals, "ERROR option provided but not defined: --cheese\n")
}

func (s *ServerSuite) TestBrokenCommand(c *tc.C) {
	resp := s.AssertBadCommand(c, []string{"remote", "--value", "error"}, 1)
	c.Assert(string(resp.Stdout), tc.Equals, "")
	c.Assert(string(resp.Stderr), tc.Equals, "ERROR blam\n")
}

type NewCommandSuite struct {
	relationSuite
}

func TestNewCommandSuite(t *stdtesting.T) {
	tc.Run(t, &NewCommandSuite{})
}

var newCommandTests = []struct {
	name string
	err  string
}{
	{"close-port", ""},
	{"config-get", ""},
	{"juju-log", ""},
	{"open-port", ""},
	{"opened-ports", ""},
	{"relation-get", ""},
	{"relation-ids", ""},
	{"relation-list", ""},
	{"relation-model-get", ""},
	{"relation-set", ""},
	{"unit-get", ""},
	{"storage-add", ""},
	{"storage-get", ""},
	{"status-get", ""},
	{"status-set", ""},
	{"random", "unknown command: random"},
}

func (s *NewCommandSuite) TestNewCommand(c *tc.C) {
	ctx, _ := s.newHookContext(0, "", "")
	for _, t := range newCommandTests {
		com, err := jujuc.NewCommand(ctx, t.name)
		if t.err == "" {
			// At this level, just check basic sanity; commands are tested in
			// more detail elsewhere.
			c.Assert(err, tc.ErrorIsNil)
			c.Assert(com.Info().Name, tc.Equals, t.name)
		} else {
			c.Assert(com, tc.IsNil)
			c.Assert(err, tc.ErrorMatches, t.err)
		}
	}
}
