// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"github.com/adrien-f/opentofu-opened/pkg/grpcwrap"
	"github.com/adrien-f/opentofu-opened/pkg/plugin"
	simple "github.com/adrien-f/opentofu-opened/pkg/provider-simple"
	"github.com/adrien-f/opentofu-opened/pkg/tfplugin5"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin5.ProviderServer {
			return grpcwrap.Provider(simple.Provider())
		},
	})
}
