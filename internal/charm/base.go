// Copyright 2020 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package charm

import (
	"fmt"
	"strings"

	"github.com/juju/errors"

	"github.com/juju/juju/core/arch"
)

// Base represents an OS/Channel.
// Bases can also be converted to and from a series string.
type Base struct {
	Name          string   `json:"name,omitempty"`
	Channel       Channel  `json:"channel,omitempty"`
	Architectures []string `json:"architectures,omitempty"`
}

// Validate returns with no error when the Base is valid.
func (b Base) Validate() error {
	if b.Name == "" {
		return errors.NotValidf("base without name")
	}
	if b.Channel.Empty() {
		return errors.NotValidf("channel")
	}
	return nil
}

// String representation of the Base.
func (b Base) String() string {
	if b.Channel.Empty() {
		panic("cannot stringify invalid base. Bases should always be validated before stringifying")
	}
	str := fmt.Sprintf("%s@%s", b.Name, b.Channel)
	if len(b.Architectures) > 0 {
		str = fmt.Sprintf("%s on %s", str, strings.Join(b.Architectures, ", "))
	}
	return str
}

// ParseBase parses a base as string in the form "os@track/risk/branch"
// and an optional list of architectures
func ParseBase(s string, archs ...string) (Base, error) {
	var err error
	base := Base{}

	segments := strings.Split(s, "@")
	if len(segments) != 2 {
		return Base{}, errors.NotValidf("base string must contain exactly one @. %q", s)
	}
	base.Name = strings.ToLower(segments[0])
	channelName := segments[1]

	if channelName != "" {
		base.Channel, err = ParseChannelNormalize(channelName)
		if err != nil {
			return Base{}, errors.Annotatef(err, "malformed channel in base string %q", s)
		}
	}

	base.Architectures = make([]string, len(archs))
	for i, v := range archs {
		base.Architectures[i] = arch.NormaliseArch(v)
	}

	err = base.Validate()
	if err != nil {
		var a string
		if len(base.Architectures) > 0 {
			a = fmt.Sprintf(" with architectures %q", strings.Join(base.Architectures, ","))
		}
		return Base{}, errors.Annotatef(err, "invalid base string %q%s", s, a)
	}
	return base, nil
}
