// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Get information on your DigitalOcean account.
 * 
 * ## Example Usage
 * 
 * Get the account:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 * 
 * const example = digitalocean.getAccount();
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/d/account.html.markdown.
 */
export function getAccount(opts?: pulumi.InvokeOptions): Promise<GetAccountResult> & GetAccountResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetAccountResult> = pulumi.runtime.invoke("digitalocean:index/getAccount:getAccount", {
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of values returned by getAccount.
 */
export interface GetAccountResult {
    readonly dropletLimit: number;
    readonly email: string;
    readonly emailVerified: boolean;
    readonly floatingIpLimit: number;
    readonly status: string;
    readonly statusMessage: string;
    readonly uuid: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
