package digitalocean

import (
	"context"
	"strings"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableDigitalOceanProject(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_project",
		Description: "Projects allow you to organize your resources into groups that fit the way you work. You can group resources (like Droplets, Spaces, load balancers, domains, and floating IPs) in ways that align with the applications you host on DigitalOcean.",
		List: &plugin.ListConfig{
			Hydrate: listProject,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getProject,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique universal identifier of this project."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The globally unique human-readable name for the project."},
			// Other columns
			{Name: "created_at", Type: proto.ColumnType_DATETIME, Description: "Time when the project was created."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of the project."},
			{Name: "environment", Type: proto.ColumnType_STRING, Description: "The environment of the project's resources."},
			{Name: "is_default", Type: proto.ColumnType_BOOL, Description: "If true, all resources will be added to this project if no project is specified."},
			{Name: "owner_id", Type: proto.ColumnType_INT, Description: "The integer id of the project owner."},
			{Name: "owner_uuid", Type: proto.ColumnType_STRING, Description: "The unique universal identifier of the project owner."},
			{Name: "purpose", Type: proto.ColumnType_STRING, Description: "The purpose of the project."},
			{Name: "updated_at", Type: proto.ColumnType_DATETIME, Description: "Time when the project was updated."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(projectToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromConstant(map[string]bool{}), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listProject(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_project.listProject", "connection_error", err)
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		projects, resp, err := conn.Projects.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("digitalocean_project.listProject", "query_error", err, "opts", opts)
			return nil, err
		}
		for _, t := range projects {
			d.StreamListItem(ctx, t)
		}
		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		// set the page we want for the next request
		opts.Page = page + 1
	}
	return nil, nil
}

func getProject(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_project.getProject", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	id := quals["id"].GetStringValue()
	result, resp, err := conn.Projects.Get(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), ": 404") {
			plugin.Logger(ctx).Warn("digitalocean_project.getProject", "not_found_error", err, "quals", quals, "resp", resp)
			return nil, nil
		}
		plugin.Logger(ctx).Error("digitalocean_project.getProject", "query_error", err, "quals", quals, "resp", resp)
		return nil, err
	}
	return result, nil
}

func projectToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	var i godo.Project
	switch d.Value.(type) {
	case *godo.Project:
		i = *d.Value.(*godo.Project)
	case godo.Project:
		i = d.Value.(godo.Project)
	}
	return "do:project:" + i.ID, nil
}
