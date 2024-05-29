// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package encryption

import (
	"github.com/adrien-f/opentofu-opened/pkg/encryption/keyprovider/aws_kms"
	"github.com/adrien-f/opentofu-opened/pkg/encryption/keyprovider/gcp_kms"
	"github.com/adrien-f/opentofu-opened/pkg/encryption/keyprovider/openbao"
	"github.com/adrien-f/opentofu-opened/pkg/encryption/keyprovider/pbkdf2"
	"github.com/adrien-f/opentofu-opened/pkg/encryption/method/aesgcm"
	"github.com/adrien-f/opentofu-opened/pkg/encryption/method/unencrypted"
	"github.com/adrien-f/opentofu-opened/pkg/encryption/registry/lockingencryptionregistry"
)

var DefaultRegistry = lockingencryptionregistry.New()

func init() {
	if err := DefaultRegistry.RegisterKeyProvider(pbkdf2.New()); err != nil {
		panic(err)
	}
	if err := DefaultRegistry.RegisterKeyProvider(aws_kms.New()); err != nil {
		panic(err)
	}
	if err := DefaultRegistry.RegisterKeyProvider(gcp_kms.New()); err != nil {
		panic(err)
	}
	if err := DefaultRegistry.RegisterKeyProvider(openbao.New()); err != nil {
		panic(err)
	}
	if err := DefaultRegistry.RegisterMethod(aesgcm.New()); err != nil {
		panic(err)
	}
	if err := DefaultRegistry.RegisterMethod(unencrypted.New()); err != nil {
		panic(err)
	}
}
