package skhema

import (
	"fmt"
	local "github.com/skhema/terraform-provider-skhema/client/backend/local"
	"strconv"
	"strings"
)

type Client struct {
	config  *Config
	service service // reuse a single struct instead of allocating new ones for each service on the heap.
	backend *local.Backend

	Types    *TypesService
	Apis     *ApisService
	Services *ServicesService
}

type Metadata struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Revision  string `json:"revision"`
}

func (m *Metadata) GetUrn() string {
	return fmt.Sprintf("%s/%s@%s", m.Namespace, m.Name, m.Revision)
}

func (m *Metadata) GetId() string {
	return fmt.Sprintf("%s/%s@%s", m.Namespace, m.Name, m.Revision)
}

func getMetadataFromId(id string) *Metadata {
	metadata := &Metadata{}

	tokens := strings.Split(id, "/")
	metadata.Namespace = tokens[0]

	tokens = strings.Split(tokens[1], "@")
	metadata.Name = tokens[0]
	metadata.Revision = tokens[1]

	return metadata
}

type service struct {
	client *Client
}

func NewClient(config *Config) (*Client, error) {
	backend, err := local.NewBackend()

	if err != nil {
		return nil, err
	}

	client := &Client{
		config:  config,
		backend: backend,
	}
	client.service.client = client
	client.Types = (*TypesService)(&client.service)
	client.Apis = (*ApisService)(&client.service)
	client.Services = (*ServicesService)(&client.service)

	return client, nil
}

func getNextRev(curr string) (string, error) {
	i, err := strconv.Atoi(curr)

	if err != nil {
		return curr, err
	}

	return strconv.Itoa(i + 1), nil
}
