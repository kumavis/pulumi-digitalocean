// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a DigitalOcean Cloud Firewall resource. This can be used to create,
// modify, and delete Firewalls.
//
// ## Import
//
// Firewalls can be imported using the firewall `id`, e.g.
//
// ```sh
//  $ pulumi import digitalocean:index/firewall:Firewall myfirewall b8ecd2ab-2267-4a5e-8692-cbf1d32583e3
// ```
type Firewall struct {
	pulumi.CustomResourceState

	// A time value given in ISO8601 combined date and time format
	// that represents when the Firewall was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The list of the IDs of the Droplets assigned
	// to the Firewall.
	DropletIds pulumi.IntArrayOutput `pulumi:"dropletIds"`
	// The inbound access rule block for the Firewall.
	// The `inboundRule` block is documented below.
	InboundRules FirewallInboundRuleArrayOutput `pulumi:"inboundRules"`
	// The Firewall name
	Name pulumi.StringOutput `pulumi:"name"`
	// The outbound access rule block for the Firewall.
	// The `outboundRule` block is documented below.
	OutboundRules FirewallOutboundRuleArrayOutput `pulumi:"outboundRules"`
	// An list of object containing the fields, "dropletId",
	// "removing", and "status".  It is provided to detail exactly which Droplets
	// are having their security policies updated.  When empty, all changes
	// have been successfully applied.
	PendingChanges FirewallPendingChangeArrayOutput `pulumi:"pendingChanges"`
	// A status string indicating the current state of the Firewall.
	// This can be "waiting", "succeeded", or "failed".
	Status pulumi.StringOutput `pulumi:"status"`
	// The names of the Tags assigned to the Firewall.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewFirewall registers a new resource with the given unique name, arguments, and options.
func NewFirewall(ctx *pulumi.Context,
	name string, args *FirewallArgs, opts ...pulumi.ResourceOption) (*Firewall, error) {
	if args == nil {
		args = &FirewallArgs{}
	}

	var resource Firewall
	err := ctx.RegisterResource("digitalocean:index/firewall:Firewall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewall gets an existing Firewall resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewall(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallState, opts ...pulumi.ResourceOption) (*Firewall, error) {
	var resource Firewall
	err := ctx.ReadResource("digitalocean:index/firewall:Firewall", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Firewall resources.
type firewallState struct {
	// A time value given in ISO8601 combined date and time format
	// that represents when the Firewall was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The list of the IDs of the Droplets assigned
	// to the Firewall.
	DropletIds []int `pulumi:"dropletIds"`
	// The inbound access rule block for the Firewall.
	// The `inboundRule` block is documented below.
	InboundRules []FirewallInboundRule `pulumi:"inboundRules"`
	// The Firewall name
	Name *string `pulumi:"name"`
	// The outbound access rule block for the Firewall.
	// The `outboundRule` block is documented below.
	OutboundRules []FirewallOutboundRule `pulumi:"outboundRules"`
	// An list of object containing the fields, "dropletId",
	// "removing", and "status".  It is provided to detail exactly which Droplets
	// are having their security policies updated.  When empty, all changes
	// have been successfully applied.
	PendingChanges []FirewallPendingChange `pulumi:"pendingChanges"`
	// A status string indicating the current state of the Firewall.
	// This can be "waiting", "succeeded", or "failed".
	Status *string `pulumi:"status"`
	// The names of the Tags assigned to the Firewall.
	Tags []string `pulumi:"tags"`
}

type FirewallState struct {
	// A time value given in ISO8601 combined date and time format
	// that represents when the Firewall was created.
	CreatedAt pulumi.StringPtrInput
	// The list of the IDs of the Droplets assigned
	// to the Firewall.
	DropletIds pulumi.IntArrayInput
	// The inbound access rule block for the Firewall.
	// The `inboundRule` block is documented below.
	InboundRules FirewallInboundRuleArrayInput
	// The Firewall name
	Name pulumi.StringPtrInput
	// The outbound access rule block for the Firewall.
	// The `outboundRule` block is documented below.
	OutboundRules FirewallOutboundRuleArrayInput
	// An list of object containing the fields, "dropletId",
	// "removing", and "status".  It is provided to detail exactly which Droplets
	// are having their security policies updated.  When empty, all changes
	// have been successfully applied.
	PendingChanges FirewallPendingChangeArrayInput
	// A status string indicating the current state of the Firewall.
	// This can be "waiting", "succeeded", or "failed".
	Status pulumi.StringPtrInput
	// The names of the Tags assigned to the Firewall.
	Tags pulumi.StringArrayInput
}

func (FirewallState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallState)(nil)).Elem()
}

type firewallArgs struct {
	// The list of the IDs of the Droplets assigned
	// to the Firewall.
	DropletIds []int `pulumi:"dropletIds"`
	// The inbound access rule block for the Firewall.
	// The `inboundRule` block is documented below.
	InboundRules []FirewallInboundRule `pulumi:"inboundRules"`
	// The Firewall name
	Name *string `pulumi:"name"`
	// The outbound access rule block for the Firewall.
	// The `outboundRule` block is documented below.
	OutboundRules []FirewallOutboundRule `pulumi:"outboundRules"`
	// The names of the Tags assigned to the Firewall.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a Firewall resource.
type FirewallArgs struct {
	// The list of the IDs of the Droplets assigned
	// to the Firewall.
	DropletIds pulumi.IntArrayInput
	// The inbound access rule block for the Firewall.
	// The `inboundRule` block is documented below.
	InboundRules FirewallInboundRuleArrayInput
	// The Firewall name
	Name pulumi.StringPtrInput
	// The outbound access rule block for the Firewall.
	// The `outboundRule` block is documented below.
	OutboundRules FirewallOutboundRuleArrayInput
	// The names of the Tags assigned to the Firewall.
	Tags pulumi.StringArrayInput
}

func (FirewallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallArgs)(nil)).Elem()
}

type FirewallInput interface {
	pulumi.Input

	ToFirewallOutput() FirewallOutput
	ToFirewallOutputWithContext(ctx context.Context) FirewallOutput
}

func (*Firewall) ElementType() reflect.Type {
	return reflect.TypeOf((*Firewall)(nil))
}

func (i *Firewall) ToFirewallOutput() FirewallOutput {
	return i.ToFirewallOutputWithContext(context.Background())
}

func (i *Firewall) ToFirewallOutputWithContext(ctx context.Context) FirewallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallOutput)
}

func (i *Firewall) ToFirewallPtrOutput() FirewallPtrOutput {
	return i.ToFirewallPtrOutputWithContext(context.Background())
}

func (i *Firewall) ToFirewallPtrOutputWithContext(ctx context.Context) FirewallPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPtrOutput)
}

type FirewallPtrInput interface {
	pulumi.Input

	ToFirewallPtrOutput() FirewallPtrOutput
	ToFirewallPtrOutputWithContext(ctx context.Context) FirewallPtrOutput
}

type firewallPtrType FirewallArgs

func (*firewallPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Firewall)(nil))
}

func (i *firewallPtrType) ToFirewallPtrOutput() FirewallPtrOutput {
	return i.ToFirewallPtrOutputWithContext(context.Background())
}

func (i *firewallPtrType) ToFirewallPtrOutputWithContext(ctx context.Context) FirewallPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPtrOutput)
}

// FirewallArrayInput is an input type that accepts FirewallArray and FirewallArrayOutput values.
// You can construct a concrete instance of `FirewallArrayInput` via:
//
//          FirewallArray{ FirewallArgs{...} }
type FirewallArrayInput interface {
	pulumi.Input

	ToFirewallArrayOutput() FirewallArrayOutput
	ToFirewallArrayOutputWithContext(context.Context) FirewallArrayOutput
}

type FirewallArray []FirewallInput

func (FirewallArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Firewall)(nil))
}

func (i FirewallArray) ToFirewallArrayOutput() FirewallArrayOutput {
	return i.ToFirewallArrayOutputWithContext(context.Background())
}

func (i FirewallArray) ToFirewallArrayOutputWithContext(ctx context.Context) FirewallArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallArrayOutput)
}

// FirewallMapInput is an input type that accepts FirewallMap and FirewallMapOutput values.
// You can construct a concrete instance of `FirewallMapInput` via:
//
//          FirewallMap{ "key": FirewallArgs{...} }
type FirewallMapInput interface {
	pulumi.Input

	ToFirewallMapOutput() FirewallMapOutput
	ToFirewallMapOutputWithContext(context.Context) FirewallMapOutput
}

type FirewallMap map[string]FirewallInput

func (FirewallMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Firewall)(nil))
}

func (i FirewallMap) ToFirewallMapOutput() FirewallMapOutput {
	return i.ToFirewallMapOutputWithContext(context.Background())
}

func (i FirewallMap) ToFirewallMapOutputWithContext(ctx context.Context) FirewallMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallMapOutput)
}

type FirewallOutput struct {
	*pulumi.OutputState
}

func (FirewallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Firewall)(nil))
}

func (o FirewallOutput) ToFirewallOutput() FirewallOutput {
	return o
}

func (o FirewallOutput) ToFirewallOutputWithContext(ctx context.Context) FirewallOutput {
	return o
}

func (o FirewallOutput) ToFirewallPtrOutput() FirewallPtrOutput {
	return o.ToFirewallPtrOutputWithContext(context.Background())
}

func (o FirewallOutput) ToFirewallPtrOutputWithContext(ctx context.Context) FirewallPtrOutput {
	return o.ApplyT(func(v Firewall) *Firewall {
		return &v
	}).(FirewallPtrOutput)
}

type FirewallPtrOutput struct {
	*pulumi.OutputState
}

func (FirewallPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Firewall)(nil))
}

func (o FirewallPtrOutput) ToFirewallPtrOutput() FirewallPtrOutput {
	return o
}

func (o FirewallPtrOutput) ToFirewallPtrOutputWithContext(ctx context.Context) FirewallPtrOutput {
	return o
}

type FirewallArrayOutput struct{ *pulumi.OutputState }

func (FirewallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Firewall)(nil))
}

func (o FirewallArrayOutput) ToFirewallArrayOutput() FirewallArrayOutput {
	return o
}

func (o FirewallArrayOutput) ToFirewallArrayOutputWithContext(ctx context.Context) FirewallArrayOutput {
	return o
}

func (o FirewallArrayOutput) Index(i pulumi.IntInput) FirewallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Firewall {
		return vs[0].([]Firewall)[vs[1].(int)]
	}).(FirewallOutput)
}

type FirewallMapOutput struct{ *pulumi.OutputState }

func (FirewallMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Firewall)(nil))
}

func (o FirewallMapOutput) ToFirewallMapOutput() FirewallMapOutput {
	return o
}

func (o FirewallMapOutput) ToFirewallMapOutputWithContext(ctx context.Context) FirewallMapOutput {
	return o
}

func (o FirewallMapOutput) MapIndex(k pulumi.StringInput) FirewallOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Firewall {
		return vs[0].(map[string]Firewall)[vs[1].(string)]
	}).(FirewallOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallOutput{})
	pulumi.RegisterOutputType(FirewallPtrOutput{})
	pulumi.RegisterOutputType(FirewallArrayOutput{})
	pulumi.RegisterOutputType(FirewallMapOutput{})
}
