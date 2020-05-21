// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides a DigitalOcean Floating IP to represent a publicly-accessible static IP addresses that can be mapped to one of your Droplets.
 *
 * > **NOTE:** Floating IPs can be assigned to a Droplet either directly on the `digitalocean..FloatingIp` resource by setting a `dropletId` or using the `digitalocean..FloatingIpAssignment` resource, but the two cannot be used together.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const foobarDroplet = new digitalocean.Droplet("foobarDroplet", {
 *     size: "s-1vcpu-1gb",
 *     image: "ubuntu-18-04-x64",
 *     region: "sgp1",
 *     ipv6: true,
 *     privateNetworking: true,
 * });
 * const foobarFloatingIp = new digitalocean.FloatingIp("foobarFloatingIp", {
 *     dropletId: foobarDroplet.id,
 *     region: foobarDroplet.region,
 * });
 * ```
 */
export class FloatingIp extends pulumi.CustomResource {
    /**
     * Get an existing FloatingIp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FloatingIpState, opts?: pulumi.CustomResourceOptions): FloatingIp {
        return new FloatingIp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean:index/floatingIp:FloatingIp';

    /**
     * Returns true if the given object is an instance of FloatingIp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FloatingIp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FloatingIp.__pulumiType;
    }

    /**
     * The ID of Droplet that the Floating IP will be assigned to.
     */
    public readonly dropletId!: pulumi.Output<number | undefined>;
    /**
     * The IP Address of the resource
     */
    public readonly ipAddress!: pulumi.Output<string>;
    /**
     * The region that the Floating IP is reserved to.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The uniform resource name of the floating ip
     */
    public /*out*/ readonly urn!: pulumi.Output<string>;

    /**
     * Create a FloatingIp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FloatingIpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FloatingIpArgs | FloatingIpState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FloatingIpState | undefined;
            inputs["dropletId"] = state ? state.dropletId : undefined;
            inputs["ipAddress"] = state ? state.ipAddress : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["urn"] = state ? state.urn : undefined;
        } else {
            const args = argsOrState as FloatingIpArgs | undefined;
            if (!args || args.region === undefined) {
                throw new Error("Missing required property 'region'");
            }
            inputs["dropletId"] = args ? args.dropletId : undefined;
            inputs["ipAddress"] = args ? args.ipAddress : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["urn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(FloatingIp.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FloatingIp resources.
 */
export interface FloatingIpState {
    /**
     * The ID of Droplet that the Floating IP will be assigned to.
     */
    readonly dropletId?: pulumi.Input<number>;
    /**
     * The IP Address of the resource
     */
    readonly ipAddress?: pulumi.Input<string>;
    /**
     * The region that the Floating IP is reserved to.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The uniform resource name of the floating ip
     */
    readonly urn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FloatingIp resource.
 */
export interface FloatingIpArgs {
    /**
     * The ID of Droplet that the Floating IP will be assigned to.
     */
    readonly dropletId?: pulumi.Input<number>;
    /**
     * The IP Address of the resource
     */
    readonly ipAddress?: pulumi.Input<string>;
    /**
     * The region that the Floating IP is reserved to.
     */
    readonly region: pulumi.Input<string>;
}
