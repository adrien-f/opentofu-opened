// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	localexec "github.com/adrien-f/opentofu-opened/pkg/builtin/provisioners/local-exec"
	"github.com/adrien-f/opentofu-opened/pkg/grpcwrap"
	"github.com/adrien-f/opentofu-opened/pkg/plugin"
	"github.com/adrien-f/opentofu-opened/pkg/tfplugin5"
)

func main() {
	// Provide a binary version of the internal terraform provider for testing
	plugin.Serve(&plugin.ServeOpts{
		GRPCProvisionerFunc: func() tfplugin5.ProvisionerServer {
			return grpcwrap.Provisioner(localexec.New())
		},
	})
}
