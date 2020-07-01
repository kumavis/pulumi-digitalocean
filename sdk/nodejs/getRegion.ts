// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Get information on a single DigitalOcean region. This is useful to find out
 * what Droplet sizes and features are supported within a region.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const sfo2 = digitalocean.getRegion({
 *     slug: "sfo2",
 * });
 * export const regionName = sfo2.then(sfo2 => sfo2.name);
 * ```
 */
export function getRegion(args: GetRegionArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("digitalocean:index/getRegion:getRegion", {
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getRegion.
 */
export interface GetRegionArgs {
    /**
     * A human-readable string that is used as a unique identifier for each region.
     */
    readonly slug: string;
}

/**
 * A collection of values returned by getRegion.
 */
export interface GetRegionResult {
    /**
     * A boolean value that represents whether new Droplets can be created in this region.
     */
    readonly available: boolean;
    /**
     * A set of features available in this region.
     */
    readonly features: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The display name of the region.
     */
    readonly name: string;
    /**
     * A set of identifying slugs for the Droplet sizes available in this region.
     */
    readonly sizes: string[];
    /**
     * A human-readable string that is used as a unique identifier for each region.
     */
    readonly slug: string;
}
