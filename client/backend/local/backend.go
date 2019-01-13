package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

type Backend struct {
	root string
}

type Metadata struct {
	Namespace string
	Name      string
	Revision  string
}

func NewBackend() (*Backend, error) {
	backend := &Backend{}

	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	backend.root = path.Join(home, ".skhema.d")

	return backend, nil
}

func (b *Backend) Write(root string, rev string, content []byte) error {
	err := ioutil.WriteFile(getFilePath(root, rev), content, 0644)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (b *Backend) Create(namespace string, name string, revision string, resource interface{}) error {
	content, err := json.Marshal(resource)
	if err != nil {
		log.Println(err)
		return err
	}

	root := getRoot(b.root, namespace, name)

	err = os.MkdirAll(root, 0700)
	if err != nil {
		log.Println(err)
		return err
	}

	err = b.Write(root, revision, content)
	if err != nil {
		log.Println(err)
		return err
	}

	err = b.Write(root, "latest", content)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (b *Backend) Read(namespace string, name string, revision string, resource interface{}) error {
	src := getFilePath(getRoot(b.root, namespace, name), revision)

	content, err := ioutil.ReadFile(src)

	if err != nil {
		log.Println(err)
		return err
	}

	err = json.Unmarshal(content, &resource)

	if err != nil {
		return err
	}

	return nil
}

func getRoot(root string, ns string, name string) string {
	return path.Join(root, ns, name)
}

func getFileName(rev string) string {
	return fmt.Sprintf("%s.json", rev)
}

func getFilePath(root string, rev string) string {
	return path.Join(root, getFileName(rev))
}
