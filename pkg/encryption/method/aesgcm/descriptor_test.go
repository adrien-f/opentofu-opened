package aesgcm_test

import (
	"testing"

	"github.com/adrien-f/opentofu-opened/pkg/encryption/method/aesgcm"
)

func TestDescriptor(t *testing.T) {
	if id := aesgcm.New().ID(); id != "aes_gcm" {
		t.Fatalf("Incorrect descriptor ID returned: %s", id)
	}
}
