package digitalocean

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type digitaloceanConfig struct {
	Token *string `hcl:"token"`
}

func ConfigInstance() interface{} {
	return &digitaloceanConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) digitaloceanConfig {
	if connection == nil || connection.GetConfig() == nil {
		return digitaloceanConfig{}
	}
	config, _ := connection.GetConfig().(digitaloceanConfig)
	return config
}
