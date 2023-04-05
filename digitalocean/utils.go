package digitalocean

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func connect(_ context.Context, d *plugin.QueryData) (*godo.Client, error) {

	// There is no CLI order of preference that I could find, so we use the
	// terraform provider order - https://registry.terraform.io/providers/digitalocean/digitalocean/latest/docs#token
	// 1. DIGITALOCEAN_TOKEN
	// 2. DIGITALOCEAN_ACCESS_TOKEN

	digitaloceanConfig := GetConfig(d.Connection)
	if digitaloceanConfig.Token != nil {
		os.Setenv("DIGITALOCEAN_TOKEN", *digitaloceanConfig.Token)
	}

	token, ok := os.LookupEnv("DIGITALOCEAN_TOKEN")
	if !ok || token == "" {
		token, ok = os.LookupEnv("DIGITALOCEAN_ACCESS_TOKEN")
		if !ok || token == "" {
			return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
		}
	}

	client := godo.NewFromToken(token)
	client.UserAgent = "Steampipe/0.x (+https://steampipe.io)"
	return client, nil
}

func resourceInterfaceDescription(key string) string {
	switch key {
	case "akas":
		return "Array of globally unique identifier strings (also known as) for the resource."
	case "tags":
		return "A map of tags for the resource."
	case "title":
		return "Title of the resource."
	}
	return ""
}

func timestampToIsoTimestamp(_ context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}
	i := d.Value.(*godo.Timestamp)
	return i.Format(time.RFC3339), nil
}

func labelsToTagsMap(_ context.Context, d *transform.TransformData) (interface{}, error) {
	labels := d.Value.([]string)
	result := map[string]bool{}
	if labels == nil {
		return result, nil
	}
	for _, i := range labels {
		result[i] = true
	}
	return result, nil
}

func ensureStringArray(_ context.Context, d *transform.TransformData) (interface{}, error) {
	switch v := d.Value.(type) {
	case []string:
		return v, nil
	case string:
		return []string{v}, nil
	default:
		str := fmt.Sprintf("%v", d.Value)
		return []string{string(str)}, nil
	}
}
