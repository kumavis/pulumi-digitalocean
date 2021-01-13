// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a DigitalOcean Floating IP to represent a publicly-accessible static IP addresses that can be mapped to one of your Droplets.
//
// > **NOTE:** Floating IPs can be assigned to a Droplet either directly on the `FloatingIp` resource by setting a `dropletId` or using the `FloatingIpAssignment` resource, but the two cannot be used together.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v3/go/digitalocean"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		foobarDroplet, err := digitalocean.NewDroplet(ctx, "foobarDroplet", &digitalocean.DropletArgs{
// 			Size:              pulumi.String("s-1vcpu-1gb"),
// 			Image:             pulumi.String("ubuntu-18-04-x64"),
// 			Region:            pulumi.String("sgp1"),
// 			Ipv6:              pulumi.Bool(true),
// 			PrivateNetworking: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = digitalocean.NewFloatingIp(ctx, "foobarFloatingIp", &digitalocean.FloatingIpArgs{
// 			DropletId: foobarDroplet.ID(),
// 			Region:    foobarDroplet.Region,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Floating IPs can be imported using the `ip`, e.g.
//
// ```sh
//  $ pulumi import digitalocean:index/floatingIp:FloatingIp myip 192.168.0.1
// ```
type FloatingIp struct {
	pulumi.CustomResourceState

	// The ID of Droplet that the Floating IP will be assigned to.
	DropletId pulumi.IntPtrOutput `pulumi:"dropletId"`
	// The uniform resource name of the floating ip
	FloatingIpUrn pulumi.StringOutput `pulumi:"floatingIpUrn"`
	// The IP Address of the resource
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The region that the Floating IP is reserved to.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewFloatingIp registers a new resource with the given unique name, arguments, and options.
func NewFloatingIp(ctx *pulumi.Context,
	name string, args *FloatingIpArgs, opts ...pulumi.ResourceOption) (*FloatingIp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource FloatingIp
	err := ctx.RegisterResource("digitalocean:index/floatingIp:FloatingIp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFloatingIp gets an existing FloatingIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFloatingIp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FloatingIpState, opts ...pulumi.ResourceOption) (*FloatingIp, error) {
	var resource FloatingIp
	err := ctx.ReadResource("digitalocean:index/floatingIp:FloatingIp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FloatingIp resources.
type floatingIpState struct {
	// The ID of Droplet that the Floating IP will be assigned to.
	DropletId *int `pulumi:"dropletId"`
	// The uniform resource name of the floating ip
	FloatingIpUrn *string `pulumi:"floatingIpUrn"`
	// The IP Address of the resource
	IpAddress *string `pulumi:"ipAddress"`
	// The region that the Floating IP is reserved to.
	Region *string `pulumi:"region"`
}

type FloatingIpState struct {
	// The ID of Droplet that the Floating IP will be assigned to.
	DropletId pulumi.IntPtrInput
	// The uniform resource name of the floating ip
	FloatingIpUrn pulumi.StringPtrInput
	// The IP Address of the resource
	IpAddress pulumi.StringPtrInput
	// The region that the Floating IP is reserved to.
	Region pulumi.StringPtrInput
}

func (FloatingIpState) ElementType() reflect.Type {
	return reflect.TypeOf((*floatingIpState)(nil)).Elem()
}

type floatingIpArgs struct {
	// The ID of Droplet that the Floating IP will be assigned to.
	DropletId *int `pulumi:"dropletId"`
	// The IP Address of the resource
	IpAddress *string `pulumi:"ipAddress"`
	// The region that the Floating IP is reserved to.
	Region string `pulumi:"region"`
}

// The set of arguments for constructing a FloatingIp resource.
type FloatingIpArgs struct {
	// The ID of Droplet that the Floating IP will be assigned to.
	DropletId pulumi.IntPtrInput
	// The IP Address of the resource
	IpAddress pulumi.StringPtrInput
	// The region that the Floating IP is reserved to.
	Region pulumi.StringInput
}

func (FloatingIpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*floatingIpArgs)(nil)).Elem()
}

type FloatingIpInput interface {
	pulumi.Input

	ToFloatingIpOutput() FloatingIpOutput
	ToFloatingIpOutputWithContext(ctx context.Context) FloatingIpOutput
}

func (FloatingIp) ElementType() reflect.Type {
	return reflect.TypeOf((*FloatingIp)(nil)).Elem()
}

func (i FloatingIp) ToFloatingIpOutput() FloatingIpOutput {
	return i.ToFloatingIpOutputWithContext(context.Background())
}

func (i FloatingIp) ToFloatingIpOutputWithContext(ctx context.Context) FloatingIpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FloatingIpOutput)
}

type FloatingIpOutput struct {
	*pulumi.OutputState
}

func (FloatingIpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FloatingIpOutput)(nil)).Elem()
}

func (o FloatingIpOutput) ToFloatingIpOutput() FloatingIpOutput {
	return o
}

func (o FloatingIpOutput) ToFloatingIpOutputWithContext(ctx context.Context) FloatingIpOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FloatingIpOutput{})
}
