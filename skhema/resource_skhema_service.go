package skhema

import (
	//"context"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/skhema/terraform-provider-skhema/client"
)

func resourceSkhemaService() *schema.Resource {
	return &schema.Resource{
		Create: resourceSkhemaServiceCreate,
		Read:   resourceSkhemaServiceRead,
		Update: resourceSkhemaServiceUpdate,
		Delete: resourceSkhemaServiceDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"api": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"namespace": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"revision": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"urn": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSkhemaServiceObject(d *schema.ResourceData) *skhema.Service {
	obj := &skhema.Service{
		Metadata: &skhema.Metadata{
			Name:     d.Get("name").(string),
			Revision: "0",
		},
	}

	if api, ok := d.GetOk("api"); ok {
		seq := api.([]interface{})
		obj.Api = make([]string, len(seq))

		for i, v := range seq {
			obj.Api[i] = v.(string)
		}
	}

	if revision, ok := d.GetOk("revision"); ok {
		obj.Metadata.Revision = revision.(string)
	}

	return obj
}

func resourceSkhemaServiceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Env).client
	resource := resourceSkhemaServiceObject(d)

	err := client.Services.Create(resource)
	if err != nil {
		log.Println(err)
		return err
	}

	d.SetId(resource.Metadata.GetId())

	return resourceSkhemaServiceRead(d, meta)
}

func resourceSkhemaServiceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Env).client
	resource := &skhema.Service{}

	err := client.Services.DescribeById(d.Id(), resource)
	if err != nil {
		log.Println(err)
		d.SetId("")
		return nil
	}

	d.Set("namespace", resource.Metadata.Namespace)
	d.Set("name", resource.Metadata.Name)
	d.Set("urn", resource.Metadata.GetUrn())
	d.Set("api", resource.Api)

	return nil
}

func resourceSkhemaServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Env).client
	resource := resourceSkhemaServiceObject(d)

	err := client.Services.Update(resource)
	if err != nil {
		log.Println(err)
		return err
	}

	d.SetId(resource.Metadata.GetId())

	return resourceSkhemaServiceRead(d, meta)
}

func resourceSkhemaServiceDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
