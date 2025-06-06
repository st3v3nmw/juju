// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	stdtesting "testing"

	"github.com/juju/tc"

	coredatabase "github.com/juju/juju/core/database"
	"github.com/juju/juju/domain"
	"github.com/juju/juju/domain/schema/testing"
	domainstorage "github.com/juju/juju/domain/storage"
	storageerrors "github.com/juju/juju/domain/storage/errors"
	"github.com/juju/juju/internal/errors"
)

type storagePoolStateSuite struct {
	testing.ModelSuite
}

func TestStoragePoolSuite(t *stdtesting.T) {
	tc.Run(t, &storagePoolStateSuite{})
}

func newStoragePoolState(factory coredatabase.TxnRunnerFactory) *State {
	return &State{
		StateBase: domain.NewStateBase(factory),
	}
}

func (s *storagePoolStateSuite) TestCreateStoragePool(c *tc.C) {
	st := newStoragePoolState(s.TxnRunnerFactory())

	sp := domainstorage.StoragePool{
		Name:     "ebs-fast",
		Provider: "ebs",
		Attrs: map[string]string{
			"foo": "foo val",
			"bar": "bar val",
		},
	}
	ctx := c.Context()
	err := st.CreateStoragePool(ctx, sp)
	c.Assert(err, tc.ErrorIsNil)

	out, err := st.GetStoragePoolByName(ctx, "ebs-fast")
	c.Assert(err, tc.ErrorIsNil)
	c.Assert(out, tc.DeepEquals, sp)
}

func (s *storagePoolStateSuite) TestCreateStoragePoolNoAttributes(c *tc.C) {
	st := newStoragePoolState(s.TxnRunnerFactory())

	sp := domainstorage.StoragePool{
		Name:     "ebs-fast",
		Provider: "ebs",
	}
	ctx := c.Context()
	err := st.CreateStoragePool(ctx, sp)
	c.Assert(err, tc.ErrorIsNil)

	out, err := st.GetStoragePoolByName(ctx, "ebs-fast")
	c.Assert(err, tc.ErrorIsNil)
	c.Assert(out, tc.DeepEquals, sp)
}

func (s *storagePoolStateSuite) TestCreateStoragePoolAlreadyExists(c *tc.C) {
	st := newStoragePoolState(s.TxnRunnerFactory())

	sp := domainstorage.StoragePool{
		Name:     "ebs-fast",
		Provider: "ebs",
		Attrs: map[string]string{
			"foo": "foo val",
			"bar": "bar val",
		},
	}
	ctx := c.Context()
	err := st.CreateStoragePool(ctx, sp)
	c.Assert(err, tc.ErrorIsNil)

	err = st.CreateStoragePool(ctx, sp)
	c.Assert(err, tc.ErrorIs, storageerrors.PoolAlreadyExists)
}

func (s *storagePoolStateSuite) TestUpdateCloudCredentialMissingName(c *tc.C) {
	st := newStoragePoolState(s.TxnRunnerFactory())

	sp := domainstorage.StoragePool{
		Provider: "ebs",
	}
	ctx := c.Context()
	err := st.CreateStoragePool(ctx, sp)
	c.Assert(errors.Is(err, storageerrors.MissingPoolNameError), tc.IsTrue)
}

func (s *storagePoolStateSuite) TestUpdateCloudCredentialMissingProvider(c *tc.C) {
	st := newStoragePoolState(s.TxnRunnerFactory())

	sp := domainstorage.StoragePool{
		Name: "ebs-fast",
	}
	ctx := c.Context()
	err := st.CreateStoragePool(ctx, sp)
	c.Assert(errors.Is(err, storageerrors.MissingPoolTypeError), tc.IsTrue)
}

func (s *storagePoolStateSuite) TestReplaceStoragePool(c *tc.C) {
	st := newStoragePoolState(s.TxnRunnerFactory())

	sp := domainstorage.StoragePool{
		Name:     "ebs-fast",
		Provider: "ebs",
		Attrs: map[string]string{
			"foo": "foo val",
			"bar": "bar val",
		},
	}
	ctx := c.Context()
	err := st.CreateStoragePool(ctx, sp)
	c.Assert(err, tc.ErrorIsNil)

	sp2 := domainstorage.StoragePool{
		Name:     "ebs-fast",
		Provider: "ebs",
		Attrs: map[string]string{
			"baz": "baz val",
		},
	}
	err = st.ReplaceStoragePool(ctx, sp2)
	c.Assert(err, tc.ErrorIsNil)

	out, err := st.GetStoragePoolByName(ctx, "ebs-fast")
	c.Assert(err, tc.ErrorIsNil)
	c.Assert(out, tc.DeepEquals, sp2)
}

func (s *storagePoolStateSuite) TestReplaceStoragePoolNoAttributes(c *tc.C) {
	st := newStoragePoolState(s.TxnRunnerFactory())

	sp := domainstorage.StoragePool{
		Name:     "ebs-fast",
		Provider: "ebs",
		Attrs: map[string]string{
			"foo": "foo val",
			"bar": "bar val",
		},
	}
	ctx := c.Context()
	err := st.CreateStoragePool(ctx, sp)
	c.Assert(err, tc.ErrorIsNil)

	sp2 := domainstorage.StoragePool{
		Name:     "ebs-fast",
		Provider: "ebs",
	}
	err = st.ReplaceStoragePool(ctx, sp2)
	c.Assert(err, tc.ErrorIsNil)

	out, err := st.GetStoragePoolByName(ctx, "ebs-fast")
	c.Assert(err, tc.ErrorIsNil)
	c.Assert(out, tc.DeepEquals, sp2)
}

func (s *storagePoolStateSuite) TestReplaceStoragePoolNotFound(c *tc.C) {
	st := newStoragePoolState(s.TxnRunnerFactory())

	sp := domainstorage.StoragePool{
		Name:     "ebs-fast",
		Provider: "ebs",
		Attrs: map[string]string{
			"baz": "baz val",
		},
	}
	ctx := c.Context()
	err := st.ReplaceStoragePool(ctx, sp)
	c.Assert(err, tc.ErrorIs, storageerrors.PoolNotFoundError)
}

func (s *storagePoolStateSuite) TestDeleteStoragePool(c *tc.C) {
	st := newStoragePoolState(s.TxnRunnerFactory())

	sp := domainstorage.StoragePool{
		Name:     "ebs-fast",
		Provider: "ebs",
		Attrs: map[string]string{
			"foo": "foo val",
			"bar": "bar val",
		},
	}
	ctx := c.Context()
	err := st.CreateStoragePool(ctx, sp)
	c.Assert(err, tc.ErrorIsNil)

	err = st.DeleteStoragePool(ctx, "ebs-fast")
	c.Assert(err, tc.ErrorIsNil)

	_, err = st.GetStoragePoolByName(ctx, "ebs-fast")
	c.Assert(err, tc.ErrorIs, storageerrors.PoolNotFoundError)
}

func (s *storagePoolStateSuite) TestDeleteStoragePoolNotFound(c *tc.C) {
	st := newStoragePoolState(s.TxnRunnerFactory())

	ctx := c.Context()
	err := st.DeleteStoragePool(ctx, "ebs-fast")
	c.Assert(err, tc.ErrorIs, storageerrors.PoolNotFoundError)
}

func (s *storagePoolStateSuite) TestListStoragePoolsWithoutBuiltins(c *tc.C) {
	c.Skip(c, "TODO: enable this test when storage pools are fully implemented in the state package")

	st := newStoragePoolState(s.TxnRunnerFactory())

	sp := domainstorage.StoragePool{
		Name:     "ebs-fast",
		Provider: "ebs",
		Attrs: map[string]string{
			"foo": "foo val",
			"bar": "bar val",
		},
	}
	sp2 := domainstorage.StoragePool{
		Name:     "ebs-faster",
		Provider: "ebs",
		Attrs: map[string]string{
			"baz": "baz val",
		},
	}
	ctx := c.Context()
	err := st.CreateStoragePool(ctx, sp)
	c.Assert(err, tc.ErrorIsNil)
	err = st.CreateStoragePool(ctx, sp2)
	c.Assert(err, tc.ErrorIsNil)

	out, err := st.ListStoragePoolsWithoutBuiltins(c.Context())
	c.Assert(err, tc.ErrorIsNil)
	c.Assert(out, tc.SameContents, []domainstorage.StoragePool{sp, sp2})
}

func (s *storagePoolStateSuite) TestGetStoragePoolByName(c *tc.C) {
	st := newStoragePoolState(s.TxnRunnerFactory())

	sp := domainstorage.StoragePool{
		Name:     "ebs-fast",
		Provider: "ebs",
		Attrs: map[string]string{
			"foo": "foo val",
			"bar": "bar val",
		},
	}
	sp2 := domainstorage.StoragePool{
		Name:     "loop",
		Provider: "loop",
	}
	ctx := c.Context()
	err := st.CreateStoragePool(ctx, sp)
	c.Assert(err, tc.ErrorIsNil)
	err = st.CreateStoragePool(ctx, sp2)
	c.Assert(err, tc.ErrorIsNil)

	out, err := st.GetStoragePoolByName(c.Context(), "ebs-fast")
	c.Assert(err, tc.ErrorIsNil)
	c.Assert(out, tc.DeepEquals, sp)
}

func (s *storagePoolStateSuite) TestListStoragePoolsByNamesAndProviders(c *tc.C) {
	c.Skip(c, "TODO: enable this test when storage pools are fully implemented in the state package")

	st := newStoragePoolState(s.TxnRunnerFactory())

	sp := domainstorage.StoragePool{
		Name:     "ebs-fast",
		Provider: "ebs",
		Attrs: map[string]string{
			"foo": "foo val",
			"bar": "bar val",
		},
	}
	sp2 := domainstorage.StoragePool{
		Name:     "loop",
		Provider: "loop",
	}
	ctx := c.Context()
	err := st.CreateStoragePool(ctx, sp)
	c.Assert(err, tc.ErrorIsNil)
	err = st.CreateStoragePool(ctx, sp2)
	c.Assert(err, tc.ErrorIsNil)

	out, err := st.ListStoragePoolsByNamesAndProviders(c.Context(), domainstorage.Names{"ebs-fast"}, domainstorage.Providers{"ebs"})
	c.Assert(err, tc.ErrorIsNil)
	c.Assert(out, tc.SameContents, []domainstorage.StoragePool{sp})
}

func (s *storagePoolStateSuite) TestListStoragePoolsByNames(c *tc.C) {
	c.Skip(c, "TODO: enable this test when storage pools are fully implemented in the state package")

	st := newStoragePoolState(s.TxnRunnerFactory())

	sp := domainstorage.StoragePool{
		Name:     "ebs-fast",
		Provider: "ebs",
		Attrs: map[string]string{
			"foo": "foo val",
			"bar": "bar val",
		},
	}
	sp2 := domainstorage.StoragePool{
		Name:     "loop",
		Provider: "loop",
	}
	ctx := c.Context()
	err := st.CreateStoragePool(ctx, sp)
	c.Assert(err, tc.ErrorIsNil)
	err = st.CreateStoragePool(ctx, sp2)
	c.Assert(err, tc.ErrorIsNil)

	out, err := st.ListStoragePoolsByNames(c.Context(), domainstorage.Names{"loop"})
	c.Assert(err, tc.ErrorIsNil)
	c.Assert(out, tc.SameContents, []domainstorage.StoragePool{sp2})
}

func (s *storagePoolStateSuite) TestListStoragePoolsByProviders(c *tc.C) {
	c.Skip(c, "TODO: enable this test when storage pools are fully implemented in the state package")

	st := newStoragePoolState(s.TxnRunnerFactory())

	sp := domainstorage.StoragePool{
		Name:     "ebs-fast",
		Provider: "ebs",
		Attrs: map[string]string{
			"foo": "foo val",
			"bar": "bar val",
		},
	}
	sp2 := domainstorage.StoragePool{
		Name:     "loop",
		Provider: "loop",
	}
	ctx := c.Context()
	err := st.CreateStoragePool(ctx, sp)
	c.Assert(err, tc.ErrorIsNil)
	err = st.CreateStoragePool(ctx, sp2)
	c.Assert(err, tc.ErrorIsNil)

	out, err := st.ListStoragePoolsByProviders(c.Context(), domainstorage.Providers{"ebs"})
	c.Assert(err, tc.ErrorIsNil)
	c.Assert(out, tc.SameContents, []domainstorage.StoragePool{sp})
}
