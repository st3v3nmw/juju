// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package service

//go:generate go run go.uber.org/mock/mockgen -typed -package service -destination mock_test.go github.com/juju/juju/domain/agentprovisioner/service State,Provider
