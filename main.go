package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/skhema/terraform-provider-skhema/skhema"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: skhema.Provider})
}
