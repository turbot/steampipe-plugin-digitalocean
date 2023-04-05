package digitalocean

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type digitaloceanConfig struct {
	Token *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &digitaloceanConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) digitaloceanConfig {
	if connection == nil || connection.Config == nil {
		return digitaloceanConfig{}
	}
	config, _ := connection.Config.(digitaloceanConfig)
	return config
}
