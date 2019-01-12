package skhema

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/mitchellh/go-homedir"
)

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Type struct {
	Namespace string   `json:"namespace,omitempty"`
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Fields    []*Field `json:"fields,omitempty"`
}

type ResourceNotFound struct{}

func (e *ResourceNotFound) Error() string {
	return "Unable to find resource"
}

func NewResourceNotFoundError() error {
	return &ResourceNotFound{}
}

type Client struct {
	namespace string
	session   *session.Session
	service   *s3.S3
	root      string
}

func NewClient(namespace string) (*Client, error) {
	var client Client

	session, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		return nil, err
	}

	client.namespace = namespace
	client.session = session
	client.service = s3.New(session)

	home, err := homedir.Dir()
	if err == nil {
		client.root = path.Join(home, ".skhema.d")
	}

	return &client, nil
}

func (c *Client) CreateType(id string, resource interface{}) error {
	r := resource.(*Type)
	r.Namespace = c.namespace

	content, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
		return err
	}

	err = os.MkdirAll(c.getRoot(c.namespace), 0700)

	if err != nil {
		log.Println(err)
		return err
	}

	dest := c.getPath(c.namespace, id)
	err = ioutil.WriteFile(dest, content, 0644)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c *Client) DescribeType(id string, resource interface{}) error {
	return c.DescribeTypeByNamespace(c.namespace, id, resource)
}

func (c *Client) DescribeTypeByNamespace(ns string, id string, resource interface{}) error {
	src := c.getPath(ns, id)
	content, err := ioutil.ReadFile(src)

	if err != nil {
		log.Println(err)
		return NewResourceNotFoundError()
	}

	err = json.Unmarshal(content, &resource)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c *Client) getRoot(namespace string) string {
	return path.Join(c.root, namespace)
}

func (c *Client) getFilename(name string) string {
	return fmt.Sprintf("%s.json", name)
}

func (c *Client) getPath(ns string, id string) string {
	return path.Join(c.getRoot(ns), c.getFilename(id))
}
