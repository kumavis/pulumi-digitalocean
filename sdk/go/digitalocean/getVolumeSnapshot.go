// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Volume snapshots are saved instances of a block storage volume. Use this data
// source to retrieve the ID of a DigitalOcean volume snapshot for use in other
// resources.
//
// ## Example Usage
//
// Get the volume snapshot:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := digitalocean.LookupVolumeSnapshot(ctx, &digitalocean.LookupVolumeSnapshotArgs{
//				MostRecent: pulumi.BoolRef(true),
//				NameRegex:  pulumi.StringRef("^web"),
//				Region:     pulumi.StringRef("nyc3"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Reuse the data about a volume snapshot to create a new volume based on it:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			snapshot, err := digitalocean.LookupVolumeSnapshot(ctx, &digitalocean.LookupVolumeSnapshotArgs{
//				NameRegex:  pulumi.StringRef("^web"),
//				Region:     pulumi.StringRef("nyc3"),
//				MostRecent: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = digitalocean.NewVolume(ctx, "foobar", &digitalocean.VolumeArgs{
//				Region:     pulumi.String("nyc3"),
//				Size:       pulumi.Int(100),
//				SnapshotId: *pulumi.String(snapshot.Id),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupVolumeSnapshot(ctx *pulumi.Context, args *LookupVolumeSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupVolumeSnapshotResult, error) {
	var rv LookupVolumeSnapshotResult
	err := ctx.Invoke("digitalocean:index/getVolumeSnapshot:getVolumeSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVolumeSnapshot.
type LookupVolumeSnapshotArgs struct {
	// If more than one result is returned, use the most recent volume snapshot.
	//
	// > **NOTE:** If more or less than a single match is returned by the search,
	// the provider will fail. Ensure that your search is specific enough to return
	// a single volume snapshot ID only, or use `mostRecent` to choose the most recent one.
	MostRecent *bool `pulumi:"mostRecent"`
	// The name of the volume snapshot.
	Name *string `pulumi:"name"`
	// A regex string to apply to the volume snapshot list returned by DigitalOcean. This allows more advanced filtering not supported from the DigitalOcean API. This filtering is done locally on what DigitalOcean returns.
	NameRegex *string `pulumi:"nameRegex"`
	// A "slug" representing a DigitalOcean region (e.g. `nyc1`). If set, only volume snapshots available in the region will be returned.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getVolumeSnapshot.
type LookupVolumeSnapshotResult struct {
	// The date and time the volume snapshot was created.
	CreatedAt string `pulumi:"createdAt"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The minimum size in gigabytes required for a volume to be created based on this volume snapshot.
	MinDiskSize int     `pulumi:"minDiskSize"`
	MostRecent  *bool   `pulumi:"mostRecent"`
	Name        *string `pulumi:"name"`
	NameRegex   *string `pulumi:"nameRegex"`
	Region      *string `pulumi:"region"`
	// A list of DigitalOcean region "slugs" indicating where the volume snapshot is available.
	Regions []string `pulumi:"regions"`
	// The billable size of the volume snapshot in gigabytes.
	Size float64 `pulumi:"size"`
	// A list of the tags associated to the volume snapshot.
	Tags []string `pulumi:"tags"`
	// The ID of the volume from which the volume snapshot originated.
	VolumeId string `pulumi:"volumeId"`
}

func LookupVolumeSnapshotOutput(ctx *pulumi.Context, args LookupVolumeSnapshotOutputArgs, opts ...pulumi.InvokeOption) LookupVolumeSnapshotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVolumeSnapshotResult, error) {
			args := v.(LookupVolumeSnapshotArgs)
			r, err := LookupVolumeSnapshot(ctx, &args, opts...)
			var s LookupVolumeSnapshotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVolumeSnapshotResultOutput)
}

// A collection of arguments for invoking getVolumeSnapshot.
type LookupVolumeSnapshotOutputArgs struct {
	// If more than one result is returned, use the most recent volume snapshot.
	//
	// > **NOTE:** If more or less than a single match is returned by the search,
	// the provider will fail. Ensure that your search is specific enough to return
	// a single volume snapshot ID only, or use `mostRecent` to choose the most recent one.
	MostRecent pulumi.BoolPtrInput `pulumi:"mostRecent"`
	// The name of the volume snapshot.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// A regex string to apply to the volume snapshot list returned by DigitalOcean. This allows more advanced filtering not supported from the DigitalOcean API. This filtering is done locally on what DigitalOcean returns.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// A "slug" representing a DigitalOcean region (e.g. `nyc1`). If set, only volume snapshots available in the region will be returned.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupVolumeSnapshotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeSnapshotArgs)(nil)).Elem()
}

// A collection of values returned by getVolumeSnapshot.
type LookupVolumeSnapshotResultOutput struct{ *pulumi.OutputState }

func (LookupVolumeSnapshotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeSnapshotResult)(nil)).Elem()
}

func (o LookupVolumeSnapshotResultOutput) ToLookupVolumeSnapshotResultOutput() LookupVolumeSnapshotResultOutput {
	return o
}

func (o LookupVolumeSnapshotResultOutput) ToLookupVolumeSnapshotResultOutputWithContext(ctx context.Context) LookupVolumeSnapshotResultOutput {
	return o
}

// The date and time the volume snapshot was created.
func (o LookupVolumeSnapshotResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeSnapshotResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVolumeSnapshotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeSnapshotResult) string { return v.Id }).(pulumi.StringOutput)
}

// The minimum size in gigabytes required for a volume to be created based on this volume snapshot.
func (o LookupVolumeSnapshotResultOutput) MinDiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVolumeSnapshotResult) int { return v.MinDiskSize }).(pulumi.IntOutput)
}

func (o LookupVolumeSnapshotResultOutput) MostRecent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVolumeSnapshotResult) *bool { return v.MostRecent }).(pulumi.BoolPtrOutput)
}

func (o LookupVolumeSnapshotResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeSnapshotResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeSnapshotResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeSnapshotResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeSnapshotResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeSnapshotResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// A list of DigitalOcean region "slugs" indicating where the volume snapshot is available.
func (o LookupVolumeSnapshotResultOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVolumeSnapshotResult) []string { return v.Regions }).(pulumi.StringArrayOutput)
}

// The billable size of the volume snapshot in gigabytes.
func (o LookupVolumeSnapshotResultOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVolumeSnapshotResult) float64 { return v.Size }).(pulumi.Float64Output)
}

// A list of the tags associated to the volume snapshot.
func (o LookupVolumeSnapshotResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVolumeSnapshotResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The ID of the volume from which the volume snapshot originated.
func (o LookupVolumeSnapshotResultOutput) VolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeSnapshotResult) string { return v.VolumeId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeSnapshotResultOutput{})
}
