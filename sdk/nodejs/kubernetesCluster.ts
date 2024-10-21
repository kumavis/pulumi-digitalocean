// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * Before importing a Kubernetes cluster, the cluster's default node pool must be tagged with
 *
 * the `terraform:default-node-pool` tag. The provider will automatically add this tag if
 *
 * the cluster only has a single node pool. Clusters with more than one node pool, however, will require
 *
 * that you manually add the `terraform:default-node-pool` tag to the node pool that you intend to be
 *
 * the default node pool.
 *
 * Then the Kubernetes cluster and its default node pool can be imported using the cluster's `id`, e.g.
 *
 * ```sh
 * $ pulumi import digitalocean:index/kubernetesCluster:KubernetesCluster mycluster 1b8b2100-0e9f-4e8f-ad78-9eb578c2a0af
 * ```
 *
 * Additional node pools must be imported separately as `digitalocean_kubernetes_cluster`
 *
 * resources, e.g.
 *
 * ```sh
 * $ pulumi import digitalocean:index/kubernetesCluster:KubernetesCluster mynodepool 9d76f410-9284-4436-9633-4066852442c8
 * ```
 */
export class KubernetesCluster extends pulumi.CustomResource {
    /**
     * Get an existing KubernetesCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KubernetesClusterState, opts?: pulumi.CustomResourceOptions): KubernetesCluster {
        return new KubernetesCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean:index/kubernetesCluster:KubernetesCluster';

    /**
     * Returns true if the given object is an instance of KubernetesCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KubernetesCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KubernetesCluster.__pulumiType;
    }

    /**
     * A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window.
     */
    public readonly autoUpgrade!: pulumi.Output<boolean | undefined>;
    /**
     * The range of IP addresses in the overlay network of the Kubernetes cluster. For more information, see [here](https://docs.digitalocean.com/products/kubernetes/how-to/create-clusters/#create-with-vpc-native).
     */
    public readonly clusterSubnet!: pulumi.Output<string>;
    /**
     * The uniform resource name (URN) for the Kubernetes cluster.
     */
    public /*out*/ readonly clusterUrn!: pulumi.Output<string>;
    /**
     * The date and time when the node was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * **Use with caution.** When set to true, all associated DigitalOcean resources created via the Kubernetes API (load balancers, volumes, and volume snapshots) will be destroyed along with the cluster when it is destroyed.
     *
     * This resource supports customized create timeouts. The default timeout is 30 minutes.
     */
    public readonly destroyAllAssociatedResources!: pulumi.Output<boolean | undefined>;
    /**
     * The base URL of the API server on the Kubernetes master node.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * Enable/disable the high availability control plane for a cluster. Once enabled for a cluster, high availability cannot be disabled. Default: false
     */
    public readonly ha!: pulumi.Output<boolean | undefined>;
    /**
     * The public IPv4 address of the Kubernetes master node. This will not be set if high availability is configured on the cluster (v1.21+)
     */
    public /*out*/ readonly ipv4Address!: pulumi.Output<string>;
    /**
     * A representation of the Kubernetes cluster's kubeconfig with the following attributes:
     */
    public /*out*/ readonly kubeConfigs!: pulumi.Output<outputs.KubernetesClusterKubeConfig[]>;
    /**
     * A block representing the cluster's maintenance window. Updates will be applied within this window. If not specified, a default maintenance window will be chosen. `autoUpgrade` must be set to `true` for this to have an effect.
     */
    public readonly maintenancePolicy!: pulumi.Output<outputs.KubernetesClusterMaintenancePolicy>;
    /**
     * A name for the Kubernetes cluster.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A block representing the cluster's default node pool. Additional node pools may be added to the cluster using the `digitalocean.KubernetesNodePool` resource. The following arguments may be specified:
     */
    public readonly nodePool!: pulumi.Output<outputs.KubernetesClusterNodePool>;
    /**
     * The slug identifier for the region where the Kubernetes cluster will be created.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Enables or disables the DigitalOcean container registry integration for the cluster. This requires that a container registry has first been created for the account. Default: false
     */
    public readonly registryIntegration!: pulumi.Output<boolean | undefined>;
    /**
     * The range of assignable IP addresses for services running in the Kubernetes cluster. For more information, see [here](https://docs.digitalocean.com/products/kubernetes/how-to/create-clusters/#create-with-vpc-native).
     */
    public readonly serviceSubnet!: pulumi.Output<string>;
    /**
     * A string indicating the current status of the individual node.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Enable/disable surge upgrades for a cluster. Default: true
     */
    public readonly surgeUpgrade!: pulumi.Output<boolean | undefined>;
    /**
     * A list of tag names to be applied to the Kubernetes cluster.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The date and time when the node was last updated.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * The slug identifier for the version of Kubernetes used for the cluster. Use [doctl](https://github.com/digitalocean/doctl) to find the available versions `doctl kubernetes options versions`. (**Note:** A cluster may only be upgraded to newer versions in-place. If the version is decreased, a new resource will be created.)
     */
    public readonly version!: pulumi.Output<string>;
    /**
     * The ID of the VPC where the Kubernetes cluster will be located.
     */
    public readonly vpcUuid!: pulumi.Output<string>;

    /**
     * Create a KubernetesCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KubernetesClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KubernetesClusterArgs | KubernetesClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KubernetesClusterState | undefined;
            resourceInputs["autoUpgrade"] = state ? state.autoUpgrade : undefined;
            resourceInputs["clusterSubnet"] = state ? state.clusterSubnet : undefined;
            resourceInputs["clusterUrn"] = state ? state.clusterUrn : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["destroyAllAssociatedResources"] = state ? state.destroyAllAssociatedResources : undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["ha"] = state ? state.ha : undefined;
            resourceInputs["ipv4Address"] = state ? state.ipv4Address : undefined;
            resourceInputs["kubeConfigs"] = state ? state.kubeConfigs : undefined;
            resourceInputs["maintenancePolicy"] = state ? state.maintenancePolicy : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodePool"] = state ? state.nodePool : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["registryIntegration"] = state ? state.registryIntegration : undefined;
            resourceInputs["serviceSubnet"] = state ? state.serviceSubnet : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["surgeUpgrade"] = state ? state.surgeUpgrade : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["vpcUuid"] = state ? state.vpcUuid : undefined;
        } else {
            const args = argsOrState as KubernetesClusterArgs | undefined;
            if ((!args || args.nodePool === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodePool'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            resourceInputs["autoUpgrade"] = args ? args.autoUpgrade : undefined;
            resourceInputs["clusterSubnet"] = args ? args.clusterSubnet : undefined;
            resourceInputs["destroyAllAssociatedResources"] = args ? args.destroyAllAssociatedResources : undefined;
            resourceInputs["ha"] = args ? args.ha : undefined;
            resourceInputs["maintenancePolicy"] = args ? args.maintenancePolicy : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodePool"] = args ? args.nodePool : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["registryIntegration"] = args ? args.registryIntegration : undefined;
            resourceInputs["serviceSubnet"] = args ? args.serviceSubnet : undefined;
            resourceInputs["surgeUpgrade"] = args ? args.surgeUpgrade : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["vpcUuid"] = args ? args.vpcUuid : undefined;
            resourceInputs["clusterUrn"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["ipv4Address"] = undefined /*out*/;
            resourceInputs["kubeConfigs"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["kubeConfigs"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(KubernetesCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KubernetesCluster resources.
 */
export interface KubernetesClusterState {
    /**
     * A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window.
     */
    autoUpgrade?: pulumi.Input<boolean>;
    /**
     * The range of IP addresses in the overlay network of the Kubernetes cluster. For more information, see [here](https://docs.digitalocean.com/products/kubernetes/how-to/create-clusters/#create-with-vpc-native).
     */
    clusterSubnet?: pulumi.Input<string>;
    /**
     * The uniform resource name (URN) for the Kubernetes cluster.
     */
    clusterUrn?: pulumi.Input<string>;
    /**
     * The date and time when the node was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * **Use with caution.** When set to true, all associated DigitalOcean resources created via the Kubernetes API (load balancers, volumes, and volume snapshots) will be destroyed along with the cluster when it is destroyed.
     *
     * This resource supports customized create timeouts. The default timeout is 30 minutes.
     */
    destroyAllAssociatedResources?: pulumi.Input<boolean>;
    /**
     * The base URL of the API server on the Kubernetes master node.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * Enable/disable the high availability control plane for a cluster. Once enabled for a cluster, high availability cannot be disabled. Default: false
     */
    ha?: pulumi.Input<boolean>;
    /**
     * The public IPv4 address of the Kubernetes master node. This will not be set if high availability is configured on the cluster (v1.21+)
     */
    ipv4Address?: pulumi.Input<string>;
    /**
     * A representation of the Kubernetes cluster's kubeconfig with the following attributes:
     */
    kubeConfigs?: pulumi.Input<pulumi.Input<inputs.KubernetesClusterKubeConfig>[]>;
    /**
     * A block representing the cluster's maintenance window. Updates will be applied within this window. If not specified, a default maintenance window will be chosen. `autoUpgrade` must be set to `true` for this to have an effect.
     */
    maintenancePolicy?: pulumi.Input<inputs.KubernetesClusterMaintenancePolicy>;
    /**
     * A name for the Kubernetes cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * A block representing the cluster's default node pool. Additional node pools may be added to the cluster using the `digitalocean.KubernetesNodePool` resource. The following arguments may be specified:
     */
    nodePool?: pulumi.Input<inputs.KubernetesClusterNodePool>;
    /**
     * The slug identifier for the region where the Kubernetes cluster will be created.
     */
    region?: pulumi.Input<string | enums.Region>;
    /**
     * Enables or disables the DigitalOcean container registry integration for the cluster. This requires that a container registry has first been created for the account. Default: false
     */
    registryIntegration?: pulumi.Input<boolean>;
    /**
     * The range of assignable IP addresses for services running in the Kubernetes cluster. For more information, see [here](https://docs.digitalocean.com/products/kubernetes/how-to/create-clusters/#create-with-vpc-native).
     */
    serviceSubnet?: pulumi.Input<string>;
    /**
     * A string indicating the current status of the individual node.
     */
    status?: pulumi.Input<string>;
    /**
     * Enable/disable surge upgrades for a cluster. Default: true
     */
    surgeUpgrade?: pulumi.Input<boolean>;
    /**
     * A list of tag names to be applied to the Kubernetes cluster.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The date and time when the node was last updated.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * The slug identifier for the version of Kubernetes used for the cluster. Use [doctl](https://github.com/digitalocean/doctl) to find the available versions `doctl kubernetes options versions`. (**Note:** A cluster may only be upgraded to newer versions in-place. If the version is decreased, a new resource will be created.)
     */
    version?: pulumi.Input<string>;
    /**
     * The ID of the VPC where the Kubernetes cluster will be located.
     */
    vpcUuid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KubernetesCluster resource.
 */
export interface KubernetesClusterArgs {
    /**
     * A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window.
     */
    autoUpgrade?: pulumi.Input<boolean>;
    /**
     * The range of IP addresses in the overlay network of the Kubernetes cluster. For more information, see [here](https://docs.digitalocean.com/products/kubernetes/how-to/create-clusters/#create-with-vpc-native).
     */
    clusterSubnet?: pulumi.Input<string>;
    /**
     * **Use with caution.** When set to true, all associated DigitalOcean resources created via the Kubernetes API (load balancers, volumes, and volume snapshots) will be destroyed along with the cluster when it is destroyed.
     *
     * This resource supports customized create timeouts. The default timeout is 30 minutes.
     */
    destroyAllAssociatedResources?: pulumi.Input<boolean>;
    /**
     * Enable/disable the high availability control plane for a cluster. Once enabled for a cluster, high availability cannot be disabled. Default: false
     */
    ha?: pulumi.Input<boolean>;
    /**
     * A block representing the cluster's maintenance window. Updates will be applied within this window. If not specified, a default maintenance window will be chosen. `autoUpgrade` must be set to `true` for this to have an effect.
     */
    maintenancePolicy?: pulumi.Input<inputs.KubernetesClusterMaintenancePolicy>;
    /**
     * A name for the Kubernetes cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * A block representing the cluster's default node pool. Additional node pools may be added to the cluster using the `digitalocean.KubernetesNodePool` resource. The following arguments may be specified:
     */
    nodePool: pulumi.Input<inputs.KubernetesClusterNodePool>;
    /**
     * The slug identifier for the region where the Kubernetes cluster will be created.
     */
    region: pulumi.Input<string | enums.Region>;
    /**
     * Enables or disables the DigitalOcean container registry integration for the cluster. This requires that a container registry has first been created for the account. Default: false
     */
    registryIntegration?: pulumi.Input<boolean>;
    /**
     * The range of assignable IP addresses for services running in the Kubernetes cluster. For more information, see [here](https://docs.digitalocean.com/products/kubernetes/how-to/create-clusters/#create-with-vpc-native).
     */
    serviceSubnet?: pulumi.Input<string>;
    /**
     * Enable/disable surge upgrades for a cluster. Default: true
     */
    surgeUpgrade?: pulumi.Input<boolean>;
    /**
     * A list of tag names to be applied to the Kubernetes cluster.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The slug identifier for the version of Kubernetes used for the cluster. Use [doctl](https://github.com/digitalocean/doctl) to find the available versions `doctl kubernetes options versions`. (**Note:** A cluster may only be upgraded to newer versions in-place. If the version is decreased, a new resource will be created.)
     */
    version: pulumi.Input<string>;
    /**
     * The ID of the VPC where the Kubernetes cluster will be located.
     */
    vpcUuid?: pulumi.Input<string>;
}
