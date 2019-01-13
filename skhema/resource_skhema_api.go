package skhema

import (
	//"context"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/skhema/terraform-provider-skhema/client"
)

func resourceSkhemaApiOperation() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"method": {
				Type:     schema.TypeString,
				Required: true,
			},
			"param": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"segment": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"consume": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"format": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"schema": {
							Type:     schema.TypeMap,
							Required: true,
						},
					},
				},
			},
			"produce": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Type:     schema.TypeString,
							Required: true,
						},
						"format": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"schema": {
							Type:     schema.TypeMap,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceSkhemaApi() *schema.Resource {
	return &schema.Resource{
		Create: resourceSkhemaApiCreate,
		Read:   resourceSkhemaApiRead,
		Update: resourceSkhemaApiUpdate,
		Delete: resourceSkhemaApiDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"operation": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceSkhemaApiOperation(),
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

func resourceSkhemaApiObject(d *schema.ResourceData) *skhema.Api {
	obj := &skhema.Api{
		Metadata: &skhema.Metadata{
			Name:     d.Get("name").(string),
			Revision: "0",
		},
	}

	if operations, ok := d.GetOk("operation"); ok {
		obj.Operations = newApiOperationsList(operations.([]interface{}))
	}

	if revision, ok := d.GetOk("revision"); ok {
		obj.Metadata.Revision = revision.(string)
	}

	return obj
}

func resourceSkhemaApiCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Env).client
	resource := resourceSkhemaApiObject(d)

	err := client.Apis.Create(resource)
	if err != nil {
		log.Println(err)
		return err
	}

	d.SetId(resource.Metadata.GetId())

	return resourceSkhemaApiRead(d, meta)
}

func resourceSkhemaApiRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Env).client
	resource := &skhema.Api{}

	err := client.Apis.DescribeById(d.Id(), resource)
	if err != nil {
		log.Println(err)
		d.SetId("")
		return nil
	}

	d.Set("namespace", resource.Metadata.Namespace)
	d.Set("name", resource.Metadata.Name)
	d.Set("urn", resource.Metadata.GetUrn())
	d.Set("operation", flattenApiOperations(resource.Operations))

	return nil
}

func resourceSkhemaApiUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Env).client
	resource := resourceSkhemaApiObject(d)

	err := client.Apis.Update(resource)
	if err != nil {
		log.Println(err)
		return err
	}

	d.SetId(resource.Metadata.GetId())

	return resourceSkhemaApiRead(d, meta)
}

func resourceSkhemaApiDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
