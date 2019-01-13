package skhema

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	var p *schema.Provider
	p = &schema.Provider{
		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["namespace"],
			},
			"bucket": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SKHEMA_BUCKET", "TODO"),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"skhema_type":    resourceSkhemaType(),
			"skhema_api":     resourceSkhemaApi(),
			"skhema_service": resourceSkhemaService(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"skhema_type": dataSourceSkhemaType(),
			"skhema_api":  dataSourceSkhemaApi(),
		},
	}

	p.ConfigureFunc = providerConfigure(p)

	return p
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"namespace": "Schema namespace",
	}
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		config := Config{
			Namespace: d.Get("namespace").(string),
			Bucket:    d.Get("bucket").(string),
		}

		meta, err := config.NewEnv()
		if err != nil {
			return nil, err
		}

		meta.(*Env).StopContext = p.StopContext()

		return meta, nil
	}
}
