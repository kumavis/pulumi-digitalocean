// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a DigitalOcean SSH key resource to allow you to manage SSH
// keys for Droplet access. Keys created with this resource
// can be referenced in your Droplet configuration via their ID or
// fingerprint.
//
// ## Import
//
// SSH Keys can be imported using the `ssh key id`, e.g.
//
// ```sh
//  $ pulumi import digitalocean:index/sshKey:SshKey mykey 263654
// ```
type SshKey struct {
	pulumi.CustomResourceState

	// The fingerprint of the SSH key
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The name of the SSH key for identification
	Name pulumi.StringOutput `pulumi:"name"`
	// The public key. If this is a file, it
	// can be read using the file interpolation function
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
}

// NewSshKey registers a new resource with the given unique name, arguments, and options.
func NewSshKey(ctx *pulumi.Context,
	name string, args *SshKeyArgs, opts ...pulumi.ResourceOption) (*SshKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PublicKey == nil {
		return nil, errors.New("invalid value for required argument 'PublicKey'")
	}
	var resource SshKey
	err := ctx.RegisterResource("digitalocean:index/sshKey:SshKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSshKey gets an existing SshKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSshKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SshKeyState, opts ...pulumi.ResourceOption) (*SshKey, error) {
	var resource SshKey
	err := ctx.ReadResource("digitalocean:index/sshKey:SshKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SshKey resources.
type sshKeyState struct {
	// The fingerprint of the SSH key
	Fingerprint *string `pulumi:"fingerprint"`
	// The name of the SSH key for identification
	Name *string `pulumi:"name"`
	// The public key. If this is a file, it
	// can be read using the file interpolation function
	PublicKey *string `pulumi:"publicKey"`
}

type SshKeyState struct {
	// The fingerprint of the SSH key
	Fingerprint pulumi.StringPtrInput
	// The name of the SSH key for identification
	Name pulumi.StringPtrInput
	// The public key. If this is a file, it
	// can be read using the file interpolation function
	PublicKey pulumi.StringPtrInput
}

func (SshKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sshKeyState)(nil)).Elem()
}

type sshKeyArgs struct {
	// The name of the SSH key for identification
	Name *string `pulumi:"name"`
	// The public key. If this is a file, it
	// can be read using the file interpolation function
	PublicKey string `pulumi:"publicKey"`
}

// The set of arguments for constructing a SshKey resource.
type SshKeyArgs struct {
	// The name of the SSH key for identification
	Name pulumi.StringPtrInput
	// The public key. If this is a file, it
	// can be read using the file interpolation function
	PublicKey pulumi.StringInput
}

func (SshKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sshKeyArgs)(nil)).Elem()
}

type SshKeyInput interface {
	pulumi.Input

	ToSshKeyOutput() SshKeyOutput
	ToSshKeyOutputWithContext(ctx context.Context) SshKeyOutput
}

func (SshKey) ElementType() reflect.Type {
	return reflect.TypeOf((*SshKey)(nil)).Elem()
}

func (i SshKey) ToSshKeyOutput() SshKeyOutput {
	return i.ToSshKeyOutputWithContext(context.Background())
}

func (i SshKey) ToSshKeyOutputWithContext(ctx context.Context) SshKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshKeyOutput)
}

type SshKeyOutput struct {
	*pulumi.OutputState
}

func (SshKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshKeyOutput)(nil)).Elem()
}

func (o SshKeyOutput) ToSshKeyOutput() SshKeyOutput {
	return o
}

func (o SshKeyOutput) ToSshKeyOutputWithContext(ctx context.Context) SshKeyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SshKeyOutput{})
}
