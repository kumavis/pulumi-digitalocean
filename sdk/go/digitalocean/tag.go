// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a DigitalOcean Tag resource. A Tag is a label that can be applied to a
// Droplet resource in order to better organize or facilitate the lookups and
// actions on it. Tags created with this resource can be referenced in your Droplet
// configuration via their ID or name.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v2/go/digitalocean"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		foobar, err := digitalocean.NewTag(ctx, "foobar", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = digitalocean.NewDroplet(ctx, "web", &digitalocean.DropletArgs{
// 			Image:  pulumi.String("ubuntu-18-04-x64"),
// 			Region: pulumi.String("nyc3"),
// 			Size:   pulumi.String("s-1vcpu-1gb"),
// 			Tags: pulumi.StringArray{
// 				foobar.ID(),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Tag struct {
	pulumi.CustomResourceState

	// A count of the database clusters that the tag is applied to.
	DatabasesCount pulumi.IntOutput `pulumi:"databasesCount"`
	// A count of the Droplets the tag is applied to.
	DropletsCount pulumi.IntOutput `pulumi:"dropletsCount"`
	// A count of the images that the tag is applied to.
	ImagesCount pulumi.IntOutput `pulumi:"imagesCount"`
	// The name of the tag
	Name pulumi.StringOutput `pulumi:"name"`
	// A count of the total number of resources that the tag is applied to.
	TotalResourceCount pulumi.IntOutput `pulumi:"totalResourceCount"`
	// A count of the volume snapshots that the tag is applied to.
	VolumeSnapshotsCount pulumi.IntOutput `pulumi:"volumeSnapshotsCount"`
	// A count of the volumes that the tag is applied to.
	VolumesCount pulumi.IntOutput `pulumi:"volumesCount"`
}

// NewTag registers a new resource with the given unique name, arguments, and options.
func NewTag(ctx *pulumi.Context,
	name string, args *TagArgs, opts ...pulumi.ResourceOption) (*Tag, error) {
	if args == nil {
		args = &TagArgs{}
	}
	var resource Tag
	err := ctx.RegisterResource("digitalocean:index/tag:Tag", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTag gets an existing Tag resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTag(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagState, opts ...pulumi.ResourceOption) (*Tag, error) {
	var resource Tag
	err := ctx.ReadResource("digitalocean:index/tag:Tag", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Tag resources.
type tagState struct {
	// A count of the database clusters that the tag is applied to.
	DatabasesCount *int `pulumi:"databasesCount"`
	// A count of the Droplets the tag is applied to.
	DropletsCount *int `pulumi:"dropletsCount"`
	// A count of the images that the tag is applied to.
	ImagesCount *int `pulumi:"imagesCount"`
	// The name of the tag
	Name *string `pulumi:"name"`
	// A count of the total number of resources that the tag is applied to.
	TotalResourceCount *int `pulumi:"totalResourceCount"`
	// A count of the volume snapshots that the tag is applied to.
	VolumeSnapshotsCount *int `pulumi:"volumeSnapshotsCount"`
	// A count of the volumes that the tag is applied to.
	VolumesCount *int `pulumi:"volumesCount"`
}

type TagState struct {
	// A count of the database clusters that the tag is applied to.
	DatabasesCount pulumi.IntPtrInput
	// A count of the Droplets the tag is applied to.
	DropletsCount pulumi.IntPtrInput
	// A count of the images that the tag is applied to.
	ImagesCount pulumi.IntPtrInput
	// The name of the tag
	Name pulumi.StringPtrInput
	// A count of the total number of resources that the tag is applied to.
	TotalResourceCount pulumi.IntPtrInput
	// A count of the volume snapshots that the tag is applied to.
	VolumeSnapshotsCount pulumi.IntPtrInput
	// A count of the volumes that the tag is applied to.
	VolumesCount pulumi.IntPtrInput
}

func (TagState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagState)(nil)).Elem()
}

type tagArgs struct {
	// The name of the tag
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Tag resource.
type TagArgs struct {
	// The name of the tag
	Name pulumi.StringPtrInput
}

func (TagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagArgs)(nil)).Elem()
}
