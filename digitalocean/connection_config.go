package digitalocean

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type digitaloceanConfig struct {
	Token        *string `cty:"token"`
	SpacesKey    *string `cty:"spaces_key"`
	SpacesSecret *string `cty:"spaces_secret"`
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
