package butane

import (
	"context"

	butane "github.com/coreos/butane/config"
	"github.com/coreos/butane/config/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.in/yaml.v3"
)

func dataSourceConfig() *schema.Resource {
	return &schema.Resource{
		Description: "Validate and transpile Butane config to Ignition config.",
		Schema: map[string]*schema.Schema{
			"content": {
				Description: "Butane configuration file.",
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: func(val any, key string) (warns []string, errs []error) {
					content := val.(string)

					var t any
					if err := yaml.Unmarshal([]byte(content), &t); err != nil {
						errs = append(errs, err)
					}

					return
				},
			},
			"files_dir": {
				Description: "Directory to embed the local files. Maps to `--files-dir` option on Butane CLI.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ignition": {
				Description: "Result Ignition configuration.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			// Property follows the official Butane package.
			// Ref: https://github.com/coreos/butane/blob/55aa746eb0b43099040268ba0c70ae3ac2a19567/internal/main.go#L50
			"pretty": {
				Description: "Output formatted results. Maps to `--pretty` option on Butane CLI.",
				Default:     false,
				Optional:    true,
				Type:        schema.TypeBool,
			},
			// Property follows the official Butane package.
			// Ref: https://github.com/coreos/butane/blob/55aa746eb0b43099040268ba0c70ae3ac2a19567/internal/main.go#L49
			"strict": {
				Description: "Strictly check the format. Any warning will make transpile fail. Maps to `--strict` option on Butane CLI.",
				Default:     false,
				Optional:    true,
				Type:        schema.TypeBool,
			},
		},
		ReadContext: dataSourceButaneReadContext,
	}
}

func dataSourceButaneReadContext(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*client)
	var diags diag.Diagnostics
	opts := common.TranslateBytesOptions{}

	content := d.Get("content").(string)
	strict := d.Get("strict").(bool)

	if ok, v := d.Get("content").(bool); ok {
		opts.Pretty = v
	}

	if v, ok := d.GetOk("files_dir"); ok {
		opts.FilesDir = v.(string)
	}

	b, r, err := butane.TranslateBytes([]byte(content), opts)
	if err != nil {
		return diag.FromErr(err)
	}

	// Inspired by official Butane.
	// https://github.com/coreos/butane/blob/55aa746eb0b43099040268ba0c70ae3ac2a19567/internal/main.go#L104-L106
	if strict && len(r.Entries) > 0 {
		return diag.Errorf("parsing error: %s, strict: %t", r.String(), strict)
	}

	d.SetId(string(c.Hash.Sum(b)))
	d.Set("ignition", string(b))
	return diags
}
