package skhema

import (
	"log"
)

type ApiOperationSchema struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Revision  string `json:"revision"`
}

type ApiOperationParam struct {
	Name    string `json:"name"`
	Segment string `json:"segment"`
	Type    string `json:"type"`
}

type ApiOperationConsumable struct {
	Format string              `json:"format"`
	Type   string              `json:"type"`
	Schema *ApiOperationSchema `json:"schema"`
}

type ApiOperationProducible struct {
	Status string              `json:"status"`
	Format string              `json:"format"`
	Type   string              `json:"type"`
	Schema *ApiOperationSchema `json:"schema"`
}

type ApiOperation struct {
	Name        string                    `json:"name"`
	Path        string                    `json:"path"`
	Method      string                    `json:"method"`
	Params      []*ApiOperationParam      `json:"params,omitempty"`
	Consumables []*ApiOperationConsumable `json:"consumables,omitempty"`
	Producibles []*ApiOperationProducible `json:"producibles,omitempty"`
}

type Api struct {
	Metadata   *Metadata       `json:"metadata"`
	Operations []*ApiOperation `json:"operations,omitempty"`
}

type ApiLookupFilter struct {
	Id        *string
	Namespace string
	Name      string
	Revision  string
}

type ApisService service

func (a *ApisService) Create(resource *Api) error {
	resource.Metadata.Namespace = a.client.config.Namespace

	err := a.client.backend.Create(
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

func (a *ApisService) Update(resource *Api) error {
	next, err := getNextRev(resource.Metadata.Revision)
	if err != nil {
		return err
	}

	resource.Metadata.Revision = next

	return a.Create(resource)
}

func (a *ApisService) DescribeById(id string, resource interface{}) error {
	metadata := getMetadataFromId(id)

	return a.Describe(&ApiLookupFilter{
		Namespace: metadata.Namespace,
		Name:      metadata.Name,
		Revision:  metadata.Revision,
	}, &resource)
}

func (a *ApisService) Describe(lookup *ApiLookupFilter, resource interface{}) error {
	err := a.client.backend.Read(
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
