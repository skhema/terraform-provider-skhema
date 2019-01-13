package skhema

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/skhema/terraform-provider-skhema/client"
)

func dataSourceSkhemaApi() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSkhemaApiRead,

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
			"urn": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSkhemaApiRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Env).client

	resource := &skhema.Api{}
	filter := &skhema.ApiLookupFilter{
		Namespace: d.Get("namespace").(string),
		Name:      d.Get("name").(string),
		Revision:  d.Get("revision").(string),
	}

	err := client.Apis.Describe(filter, &resource)

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

	return nil
}
