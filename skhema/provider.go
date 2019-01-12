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
		},

		ResourcesMap: map[string]*schema.Resource{
			"skhema_type": resourceSkhemaType(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"skhema_type": dataSourceSkhemaType(),
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
		}

		meta, err := config.NewEnv()
		if err != nil {
			return nil, err
		}

		meta.(*Env).StopContext = p.StopContext()

		return meta, nil
	}
}
