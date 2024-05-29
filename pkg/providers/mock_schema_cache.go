package providers

import "github.com/adrien-f/opentofu-opened/pkg/addrs"

func NewMockSchemaCache() *schemaCache {
	return &schemaCache{
		m: make(map[addrs.Provider]ProviderSchema),
	}
}
