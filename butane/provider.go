package butane

import (
	"context"
	"crypto/sha256"
	"hash"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
		DataSourcesMap: map[string]*schema.Resource{
			"butane_config": dataSourceConfig(),
		},
		ResourcesMap:         map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigureContextFunc,
	}
}

type client struct {
	Hash hash.Hash
}

func providerConfigureContextFunc(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
	return client{
		Hash: sha256.New(),
	}, nil
}

type Resolver struct {
	URL string
}
