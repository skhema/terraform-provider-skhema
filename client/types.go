package skhema

import (
	"log"
)

type TypeField struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Type struct {
	Metadata *Metadata    `json:"metadata"`
	Type     string       `json:"type"`
	Items    string       `json:"items,omitempty"`
	Fields   []*TypeField `json:"fields,omitempty"`
}

type TypeLookupFilter struct {
	Id        *string
	Namespace string
	Name      string
	Revision  string
}

type TypesService service

func (t *TypesService) Create(resource *Type) error {
	resource.Metadata.Namespace = t.client.config.Namespace

	err := t.client.backend.Create(
		resource.Metadata.Namespace,
		resource.Metadata.Name,
		resource.Metadata.Revision,
		&resource,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (t *TypesService) Update(resource *Type) error {
	next, err := getNextRev(resource.Metadata.Revision)
	if err != nil {
		return err
	}

	resource.Metadata.Revision = next

	return t.Create(resource)
}

func (t *TypesService) DescribeById(id string, resource interface{}) error {
	metadata := getMetadataFromId(id)

	return t.Describe(&TypeLookupFilter{
		Namespace: metadata.Namespace,
		Name:      metadata.Name,
		Revision:  metadata.Revision,
	}, &resource)
}

func (t *TypesService) Describe(lookup *TypeLookupFilter, resource interface{}) error {
	err := t.client.backend.Read(
		lookup.Namespace,
		lookup.Name,
		lookup.Revision,
		&resource,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
