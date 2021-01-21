package main

import (
	"github.com/turbot/steampipe-plugin-digitalocean/digitalocean"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: digitalocean.Plugin})
}
