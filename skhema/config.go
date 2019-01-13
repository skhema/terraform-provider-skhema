package skhema

import (
	"context"
	"log"

	"github.com/skhema/terraform-provider-skhema/client"
)

type Config struct {
	Bucket    string
	Namespace string
}

type Env struct {
	namespace   string
	client      *skhema.Client
	StopContext context.Context
}

func (c *Config) NewEnv() (interface{}, error) {
	var env Env

	env.namespace = c.Namespace

	client, err := skhema.NewClient(&skhema.Config{
		Bucket:    c.Bucket,
		Namespace: c.Namespace,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	env.client = client

	return &env, nil
}
