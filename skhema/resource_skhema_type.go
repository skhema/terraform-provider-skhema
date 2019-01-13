package skhema

import (
	//"context"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/skhema/terraform-provider-skhema/client"
)

func resourceSkhemaType() *schema.Resource {
	return &schema.Resource{
		Create: resourceSkhemaTypeCreate,
		Read:   resourceSkhemaTypeRead,
		Update: resourceSkhemaTypeUpdate,
		Delete: resourceSkhemaTypeDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"items": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"field": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
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

func resourceSkhemaTypeObject(d *schema.ResourceData) *skhema.Type {
	obj := &skhema.Type{
		Metadata: &skhema.Metadata{
			Name:     d.Get("name").(string),
			Revision: "0",
		},
		Type: d.Get("type").(string),
	}

	// TODO: check whether struct (field) or array (items)
	if items, ok := d.GetOk("items"); ok {
		obj.Items = items.(string)
	}

	if fields, ok := d.GetOk("field"); ok {
		obj.Fields = newTypeFieldsList(fields.([]interface{}))
	}

	if revision, ok := d.GetOk("revision"); ok {
		obj.Metadata.Revision = revision.(string)
	}

	return obj
}

func resourceSkhemaTypeCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Env).client
	resource := resourceSkhemaTypeObject(d)

	err := client.Types.Create(resource)
	if err != nil {
		log.Println(err)
		return err
	}

	d.SetId(resource.Metadata.GetId())

	return resourceSkhemaTypeRead(d, meta)
}

func resourceSkhemaTypeRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Env).client
	resource := &skhema.Type{}

	err := client.Types.DescribeById(d.Id(), &resource)

	if err != nil {
		log.Println(err)
		d.SetId("")
		return nil
	}

	d.Set("namespace", resource.Metadata.Namespace)
	d.Set("name", resource.Metadata.Name)
	d.Set("revision", resource.Metadata.Revision)
	d.Set("urn", resource.Metadata.GetUrn())
	d.Set("type", resource.Type)
	d.Set("field", flattenTypeFieldsList(resource.Fields))

	return nil
}

func resourceSkhemaTypeUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Env).client
	resource := resourceSkhemaTypeObject(d)

	err := client.Types.Update(resource)
	if err != nil {
		log.Println(err)
		return err
	}

	d.SetId(resource.Metadata.GetId())

	return resourceSkhemaTypeRead(d, meta)
}

func resourceSkhemaTypeDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
