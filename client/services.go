package skhema

import (
	"log"
)

type Service struct {
	Metadata *Metadata `json:"metadata"`
	Api      []string  `json:"api,omitempty"`
}

type ServiceLookupFilter struct {
	Id        *string
	Namespace string
	Name      string
	Revision  string
}

type ServicesService service

func (t *ServicesService) Create(resource *Service) error {
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

func (t *ServicesService) Update(resource *Service) error {
	next, err := getNextRev(resource.Metadata.Revision)
	if err != nil {
		return err
	}

	resource.Metadata.Revision = next

	return t.Create(resource)
}

func (t *ServicesService) DescribeById(id string, resource interface{}) error {
	metadata := getMetadataFromId(id)

	return t.Describe(&ServiceLookupFilter{
		Namespace: metadata.Namespace,
		Name:      metadata.Name,
		Revision:  metadata.Revision,
	}, &resource)
}

func (t *ServicesService) Describe(lookup *ServiceLookupFilter, resource interface{}) error {
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
