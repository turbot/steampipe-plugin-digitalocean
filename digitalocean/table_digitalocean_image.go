package digitalocean

import (
	"context"
	"fmt"
	"strings"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableDigitalOceanImage(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_image",
		Description: "A DigitalOcean image can be used to create a Droplet and may come in a number of flavors. Currently, there are five types of images: snapshots, backups, applications, distributions, and custom images.",
		List: &plugin.ListConfig{
			Hydrate: listImage,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id", "slug"}),
			Hydrate:    getImage,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "A unique number that can be used to identify and reference a specific image."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The display name that has been given to an image. This is what is shown in the control panel and is generally a descriptive title for the image in question."},
			// Other columns
			{Name: "created_at", Type: proto.ColumnType_DATETIME, Description: "Time when the image was created."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "An optional free-form text field to describe an image."},
			{Name: "distribution", Type: proto.ColumnType_STRING, Description: "This attribute describes the base distribution used for this image. For custom images, this is user defined."},
			{Name: "error_message", Type: proto.ColumnType_STRING, Description: "A string containing information about errors that may occur when importing a custom image."},
			{Name: "min_disk_size", Type: proto.ColumnType_INT, Description: "The minimum disk size in GB required for a Droplet to use this image."},
			{Name: "public", Type: proto.ColumnType_BOOL, Description: "This is a boolean value that indicates whether the image in question is public or not. An image that is public is available to all accounts. A non-public image is only accessible from your account."},
			{Name: "regions", Type: proto.ColumnType_JSON, Description: "Array of region slugs where the image is available."},
			{Name: "size_gigabytes", Type: proto.ColumnType_INT, Description: "The size of the image in gigabytes."},
			{Name: "slug", Type: proto.ColumnType_STRING, Description: "A uniquely identifying string that is associated with each of the DigitalOcean-provided public images. These can be used to reference a public image as an alternative to the numeric id."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "A status string indicating the state of a custom image. This may be \"NEW\", \"available\", \"pending\", or \"deleted\"."},
			{Name: "tags_src", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags"), Description: "An array containing the names of the tags the image has been tagged with."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Describes the kind of image. It may be one of \"snapshot\", \"backup\", or \"custom\"."},
			{Name: "urn", Type: proto.ColumnType_STRING, Transform: transform.FromValue().Transform(imageToURN), Description: "The uniform resource name (URN) for the volume."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(imageToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags").Transform(labelsToTagsMap), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listImage(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_image.listImage", "connection_error", err)
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		images, resp, err := conn.Images.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("digitalocean_image.listImage", "query_error", err, "opts", opts)
			return nil, err
		}
		for _, t := range images {
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

func getImage(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_image.getImage", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	id := int(quals["id"].GetInt64Value())
	slug := quals["slug"].GetStringValue()
	var result *godo.Image
	var resp *godo.Response
	if id != 0 {
		result, resp, err = conn.Images.GetByID(ctx, id)
	} else if len(slug) > 0 {
		result, resp, err = conn.Images.GetBySlug(ctx, slug)
	} else {
		plugin.Logger(ctx).Warn("digitalocean_image.getImage", "invalid_quals", "id and slug both empty", "quals", quals)
		return nil, nil
	}
	if err != nil {
		if strings.Contains(err.Error(), ": 404") {
			plugin.Logger(ctx).Warn("digitalocean_image.getImage", "not_found_error", err, "quals", quals, "resp", resp)
			return nil, nil
		}
		plugin.Logger(ctx).Error("digitalocean_image.getImage", "query_error", err, "quals", quals, "resp", resp)
		return nil, err
	}
	return *result, nil
}

func imageToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.Value.(godo.Image)
	return fmt.Sprintf("do:image:%d", i.ID), nil
}
