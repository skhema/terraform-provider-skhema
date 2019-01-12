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
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceSkhemaTypeRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Env).client

	resource := &skhema.Type{}
	err := client.DescribeTypeByNamespace(d.Get("namespace").(string), d.Get("name").(string), resource)

	if err != nil {
		log.Println(err)
		d.SetId("")
		return nil
	}

	d.SetId(resource.Name)

	d.Set("namespace", resource.Namespace)
	d.Set("name", resource.Name)
	d.Set("type", resource.Type)

	return nil
}
