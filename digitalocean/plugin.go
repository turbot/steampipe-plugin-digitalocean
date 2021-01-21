package digitalocean

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-digitalocean",
		DefaultTransform: transform.FromJSONTag().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"digitalocean_account":     tableDigitalOceanAccount(ctx),
			"digitalocean_droplet":     tableDigitalOceanDroplet(ctx),
			"digitalocean_floating_ip": tableDigitalOceanFloatingIP(ctx),
			"digitalocean_key":         tableDigitalOceanKey(ctx),
			/*
				"digitalocean_action":        tableDigitalOceanAction(ctx),
				"digitalocean_app":           tableDigitalOceanApp(ctx),
				"digitalocean_balance":       tableDigitalOceanBalance(ctx),
				"digitalocean_bill":          tableDigitalOceanBill(ctx),
			*/
			"digitalocean_database":      tableDigitalOceanDatabase(ctx),
			"digitalocean_image":         tableDigitalOceanImage(ctx),
			"digitalocean_load_balancer": tableDigitalOceanLoadBalancer(ctx),
			"digitalocean_project":       tableDigitalOceanProject(ctx),
			"digitalocean_region":        tableDigitalOceanRegion(ctx),
			"digitalocean_size":          tableDigitalOceanSize(ctx),
			"digitalocean_snapshot":      tableDigitalOceanSnapshot(ctx),
			"digitalocean_tag":           tableDigitalOceanTag(ctx),
			"digitalocean_volume":        tableDigitalOceanVolume(ctx),
			"digitalocean_vpc":           tableDigitalOceanVPC(ctx),
		},
	}
	return p
}
