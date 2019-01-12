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
		},
	}
}

func resourceSkhemaTypeObject(d *schema.ResourceData) *skhema.Type {
	obj := &skhema.Type{
		Name: d.Get("name").(string),
		Type: d.Get("type").(string),
	}

	if namespace, ok := d.GetOk("namespace"); ok {
		obj.Namespace = namespace.(string)
	}

	if fields, ok := d.GetOk("field"); ok {
		obj.Fields = newFieldList(fields.([]interface{}))
	}

	return obj
}

func resourceSkhemaTypeCreate(d *schema.ResourceData, meta interface{}) error {
	id := d.Get("name").(string)
	resource := resourceSkhemaTypeObject(d)

	client := meta.(*Env).client
	err := client.CreateType(id, resource)

	if err != nil {
		log.Panic("Resource has not been created")
		return err
	}

	d.SetId(id)

	return resourceSkhemaTypeRead(d, meta)
}

func resourceSkhemaTypeRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Env).client

	resource := &skhema.Type{}
	err := client.DescribeType(d.Id(), resource)

	if err != nil {
		log.Println(err)
		d.SetId("")
		return nil
	}

	d.Set("namespace", resource.Namespace)
	d.Set("name", resource.Name)
	d.Set("type", resource.Type)
	d.Set("field", flattenFieldList(resource.Fields))

	return nil
}

func resourceSkhemaTypeUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceSkhemaTypeRead(d, meta)
}

func resourceSkhemaTypeDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
