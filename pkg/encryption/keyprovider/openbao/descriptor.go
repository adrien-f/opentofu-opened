package openbao

import "github.com/adrien-f/opentofu-opened/pkg/encryption/keyprovider"

func New() keyprovider.Descriptor {
	return &descriptor{}
}

type descriptor struct {
}

func (f descriptor) ID() keyprovider.ID {
	return "openbao"
}

func (f descriptor) ConfigStruct() keyprovider.Config {
	return &Config{}
}
