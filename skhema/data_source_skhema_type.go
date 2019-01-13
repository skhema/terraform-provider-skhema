package skhema

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/skhema/terraform-provider-skhema/client"
)

func dataSourceSkhemaType() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSkhemaTypeRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"revision": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "latest",
			},
			"schema": {
				Type:     schema.TypeMap,
				Computed: true,
			},
			"urn": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSkhemaTypeRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Env).client

	resource := &skhema.Type{}
	filter := &skhema.TypeLookupFilter{
		Namespace: d.Get("namespace").(string),
		Name:      d.Get("name").(string),
		Revision:  d.Get("revision").(string),
	}

	err := client.Types.Describe(filter, resource)

	if err != nil {
		log.Println(err)
		d.SetId("")
		return nil
	}

	d.SetId(resource.Metadata.GetId())

	d.Set("namespace", resource.Metadata.Namespace)
	d.Set("name", resource.Metadata.Name)
	d.Set("revision", resource.Metadata.Revision)
	d.Set("urn", resource.Metadata.GetUrn())
	d.Set("schema", map[string]string{
		"namespace": resource.Metadata.Namespace,
		"name":      resource.Metadata.Name,
		"revision":  resource.Metadata.Revision,
	})

	return nil
}
