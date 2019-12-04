// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Digitalocean
{
    /// <summary>
    /// Provides a DigitalOcean Kubernetes cluster resource. This can be used to create, delete, and modify clusters. For more information see the [official documentation](https://www.digitalocean.com/docs/kubernetes/).
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/r/kubernetes_cluster.html.markdown.
    /// </summary>
    public partial class KubernetesCluster : Pulumi.CustomResource
    {
        /// <summary>
        /// The range of IP addresses in the overlay network of the Kubernetes cluster.
        /// </summary>
        [Output("clusterSubnet")]
        public Output<string> ClusterSubnet { get; private set; } = null!;

        /// <summary>
        /// The date and time when the Kubernetes cluster was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The base URL of the API server on the Kubernetes master node.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// The public IPv4 address of the Kubernetes master node.
        /// </summary>
        [Output("ipv4Address")]
        public Output<string> Ipv4Address { get; private set; } = null!;

        [Output("kubeConfigs")]
        public Output<ImmutableArray<Outputs.KubernetesClusterKubeConfigs>> KubeConfigs { get; private set; } = null!;

        /// <summary>
        /// A name for the Kubernetes cluster.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A block representing the cluster's default node pool. Additional node pools may be added to the cluster using the `digitalocean..KubernetesNodePool` resource. The following arguments may be specified:
        /// - `name` - (Required) A name for the node pool.
        /// - `size` - (Required) The slug identifier for the type of Droplet to be used as workers in the node pool.
        /// - `node_count` - (Optional) The number of Droplet instances in the node pool. If auto-scaling is enabled, this should only be set if the desired result is to explicitly reset the number of nodes to this value. If auto-scaling is enabled, and the node count is outside of the given min/max range, it will use the min nodes value.
        /// - `auto_scale` - (Optional) Enable auto-scaling of the number of nodes in the node pool within the given min/max range.
        /// - `min_nodes` - (Optional) If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to.
        /// - `max_nodes` - (Optional) If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to.
        /// - `tags` - (Optional) A list of tag names to be applied to the Kubernetes cluster.
        /// </summary>
        [Output("nodePool")]
        public Output<Outputs.KubernetesClusterNodePool> NodePool { get; private set; } = null!;

        /// <summary>
        /// The slug identifier for the region where the Kubernetes cluster will be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The range of assignable IP addresses for services running in the Kubernetes cluster.
        /// </summary>
        [Output("serviceSubnet")]
        public Output<string> ServiceSubnet { get; private set; } = null!;

        /// <summary>
        /// A string indicating the current status of the cluster. Potential values include running, provisioning, and errored.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A list of tag names to be applied to the Kubernetes cluster.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The date and time when the Kubernetes cluster was last updated.
        /// * `kube_config.0` - A representation of the Kubernetes cluster's kubeconfig with the following attributes:
        /// - `raw_config` - The full contents of the Kubernetes cluster's kubeconfig file.
        /// - `host` - The URL of the API server on the Kubernetes master node.
        /// - `cluster_ca_certificate` - The base64 encoded public certificate for the cluster's certificate authority.
        /// - `token` - The DigitalOcean API access token used by clients to access the cluster.
        /// - `client_key` - The base64 encoded private key used by clients to access the cluster. Only available if token authentication is not supported on your cluster.
        /// - `client_certificate` - The base64 encoded public certificate used by clients to access the cluster. Only available if token authentication is not supported on your cluster.
        /// - `expires_at` - The date and time when the credentials will expire and need to be regenerated.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The slug identifier for the version of Kubernetes used for the cluster. Use [doctl](https://github.com/digitalocean/doctl) to find the available versions `doctl kubernetes options versions`. (**Note:** A cluster may only be upgraded to newer versions in-place. If the version is decreased, a new resource will be created.)
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a KubernetesCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KubernetesCluster(string name, KubernetesClusterArgs args, CustomResourceOptions? options = null)
            : base("digitalocean:index/kubernetesCluster:KubernetesCluster", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private KubernetesCluster(string name, Input<string> id, KubernetesClusterState? state = null, CustomResourceOptions? options = null)
            : base("digitalocean:index/kubernetesCluster:KubernetesCluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing KubernetesCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KubernetesCluster Get(string name, Input<string> id, KubernetesClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new KubernetesCluster(name, id, state, options);
        }
    }

    public sealed class KubernetesClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A name for the Kubernetes cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A block representing the cluster's default node pool. Additional node pools may be added to the cluster using the `digitalocean..KubernetesNodePool` resource. The following arguments may be specified:
        /// - `name` - (Required) A name for the node pool.
        /// - `size` - (Required) The slug identifier for the type of Droplet to be used as workers in the node pool.
        /// - `node_count` - (Optional) The number of Droplet instances in the node pool. If auto-scaling is enabled, this should only be set if the desired result is to explicitly reset the number of nodes to this value. If auto-scaling is enabled, and the node count is outside of the given min/max range, it will use the min nodes value.
        /// - `auto_scale` - (Optional) Enable auto-scaling of the number of nodes in the node pool within the given min/max range.
        /// - `min_nodes` - (Optional) If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to.
        /// - `max_nodes` - (Optional) If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to.
        /// - `tags` - (Optional) A list of tag names to be applied to the Kubernetes cluster.
        /// </summary>
        [Input("nodePool", required: true)]
        public Input<Inputs.KubernetesClusterNodePoolArgs> NodePool { get; set; } = null!;

        /// <summary>
        /// The slug identifier for the region where the Kubernetes cluster will be created.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tag names to be applied to the Kubernetes cluster.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The slug identifier for the version of Kubernetes used for the cluster. Use [doctl](https://github.com/digitalocean/doctl) to find the available versions `doctl kubernetes options versions`. (**Note:** A cluster may only be upgraded to newer versions in-place. If the version is decreased, a new resource will be created.)
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public KubernetesClusterArgs()
        {
        }
    }

    public sealed class KubernetesClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The range of IP addresses in the overlay network of the Kubernetes cluster.
        /// </summary>
        [Input("clusterSubnet")]
        public Input<string>? ClusterSubnet { get; set; }

        /// <summary>
        /// The date and time when the Kubernetes cluster was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The base URL of the API server on the Kubernetes master node.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// The public IPv4 address of the Kubernetes master node.
        /// </summary>
        [Input("ipv4Address")]
        public Input<string>? Ipv4Address { get; set; }

        [Input("kubeConfigs")]
        private InputList<Inputs.KubernetesClusterKubeConfigsGetArgs>? _kubeConfigs;
        public InputList<Inputs.KubernetesClusterKubeConfigsGetArgs> KubeConfigs
        {
            get => _kubeConfigs ?? (_kubeConfigs = new InputList<Inputs.KubernetesClusterKubeConfigsGetArgs>());
            set => _kubeConfigs = value;
        }

        /// <summary>
        /// A name for the Kubernetes cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A block representing the cluster's default node pool. Additional node pools may be added to the cluster using the `digitalocean..KubernetesNodePool` resource. The following arguments may be specified:
        /// - `name` - (Required) A name for the node pool.
        /// - `size` - (Required) The slug identifier for the type of Droplet to be used as workers in the node pool.
        /// - `node_count` - (Optional) The number of Droplet instances in the node pool. If auto-scaling is enabled, this should only be set if the desired result is to explicitly reset the number of nodes to this value. If auto-scaling is enabled, and the node count is outside of the given min/max range, it will use the min nodes value.
        /// - `auto_scale` - (Optional) Enable auto-scaling of the number of nodes in the node pool within the given min/max range.
        /// - `min_nodes` - (Optional) If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to.
        /// - `max_nodes` - (Optional) If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to.
        /// - `tags` - (Optional) A list of tag names to be applied to the Kubernetes cluster.
        /// </summary>
        [Input("nodePool")]
        public Input<Inputs.KubernetesClusterNodePoolGetArgs>? NodePool { get; set; }

        /// <summary>
        /// The slug identifier for the region where the Kubernetes cluster will be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The range of assignable IP addresses for services running in the Kubernetes cluster.
        /// </summary>
        [Input("serviceSubnet")]
        public Input<string>? ServiceSubnet { get; set; }

        /// <summary>
        /// A string indicating the current status of the cluster. Potential values include running, provisioning, and errored.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tag names to be applied to the Kubernetes cluster.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The date and time when the Kubernetes cluster was last updated.
        /// * `kube_config.0` - A representation of the Kubernetes cluster's kubeconfig with the following attributes:
        /// - `raw_config` - The full contents of the Kubernetes cluster's kubeconfig file.
        /// - `host` - The URL of the API server on the Kubernetes master node.
        /// - `cluster_ca_certificate` - The base64 encoded public certificate for the cluster's certificate authority.
        /// - `token` - The DigitalOcean API access token used by clients to access the cluster.
        /// - `client_key` - The base64 encoded private key used by clients to access the cluster. Only available if token authentication is not supported on your cluster.
        /// - `client_certificate` - The base64 encoded public certificate used by clients to access the cluster. Only available if token authentication is not supported on your cluster.
        /// - `expires_at` - The date and time when the credentials will expire and need to be regenerated.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The slug identifier for the version of Kubernetes used for the cluster. Use [doctl](https://github.com/digitalocean/doctl) to find the available versions `doctl kubernetes options versions`. (**Note:** A cluster may only be upgraded to newer versions in-place. If the version is decreased, a new resource will be created.)
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public KubernetesClusterState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class KubernetesClusterKubeConfigsGetArgs : Pulumi.ResourceArgs
    {
        [Input("clientCertificate")]
        public Input<string>? ClientCertificate { get; set; }

        [Input("clientKey")]
        public Input<string>? ClientKey { get; set; }

        [Input("clusterCaCertificate")]
        public Input<string>? ClusterCaCertificate { get; set; }

        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        [Input("host")]
        public Input<string>? Host { get; set; }

        [Input("rawConfig")]
        public Input<string>? RawConfig { get; set; }

        [Input("token")]
        public Input<string>? Token { get; set; }

        public KubernetesClusterKubeConfigsGetArgs()
        {
        }
    }

    public sealed class KubernetesClusterNodePoolArgs : Pulumi.ResourceArgs
    {
        [Input("actualNodeCount")]
        public Input<int>? ActualNodeCount { get; set; }

        [Input("autoScale")]
        public Input<bool>? AutoScale { get; set; }

        /// <summary>
        /// A unique ID that can be used to identify and reference a Kubernetes cluster.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("maxNodes")]
        public Input<int>? MaxNodes { get; set; }

        [Input("minNodes")]
        public Input<int>? MinNodes { get; set; }

        /// <summary>
        /// A name for the Kubernetes cluster.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("nodeCount")]
        public Input<int>? NodeCount { get; set; }

        [Input("nodes")]
        private InputList<KubernetesClusterNodePoolNodesArgs>? _nodes;
        public InputList<KubernetesClusterNodePoolNodesArgs> Nodes
        {
            get => _nodes ?? (_nodes = new InputList<KubernetesClusterNodePoolNodesArgs>());
            set => _nodes = value;
        }

        [Input("size", required: true)]
        public Input<string> Size { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tag names to be applied to the Kubernetes cluster.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public KubernetesClusterNodePoolArgs()
        {
        }
    }

    public sealed class KubernetesClusterNodePoolGetArgs : Pulumi.ResourceArgs
    {
        [Input("actualNodeCount")]
        public Input<int>? ActualNodeCount { get; set; }

        [Input("autoScale")]
        public Input<bool>? AutoScale { get; set; }

        /// <summary>
        /// A unique ID that can be used to identify and reference a Kubernetes cluster.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("maxNodes")]
        public Input<int>? MaxNodes { get; set; }

        [Input("minNodes")]
        public Input<int>? MinNodes { get; set; }

        /// <summary>
        /// A name for the Kubernetes cluster.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("nodeCount")]
        public Input<int>? NodeCount { get; set; }

        [Input("nodes")]
        private InputList<KubernetesClusterNodePoolNodesGetArgs>? _nodes;
        public InputList<KubernetesClusterNodePoolNodesGetArgs> Nodes
        {
            get => _nodes ?? (_nodes = new InputList<KubernetesClusterNodePoolNodesGetArgs>());
            set => _nodes = value;
        }

        [Input("size", required: true)]
        public Input<string> Size { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tag names to be applied to the Kubernetes cluster.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public KubernetesClusterNodePoolGetArgs()
        {
        }
    }

    public sealed class KubernetesClusterNodePoolNodesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date and time when the Kubernetes cluster was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// A unique ID that can be used to identify and reference a Kubernetes cluster.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// A name for the Kubernetes cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A string indicating the current status of the cluster. Potential values include running, provisioning, and errored.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The date and time when the Kubernetes cluster was last updated.
        /// * `kube_config.0` - A representation of the Kubernetes cluster's kubeconfig with the following attributes:
        /// - `raw_config` - The full contents of the Kubernetes cluster's kubeconfig file.
        /// - `host` - The URL of the API server on the Kubernetes master node.
        /// - `cluster_ca_certificate` - The base64 encoded public certificate for the cluster's certificate authority.
        /// - `token` - The DigitalOcean API access token used by clients to access the cluster.
        /// - `client_key` - The base64 encoded private key used by clients to access the cluster. Only available if token authentication is not supported on your cluster.
        /// - `client_certificate` - The base64 encoded public certificate used by clients to access the cluster. Only available if token authentication is not supported on your cluster.
        /// - `expires_at` - The date and time when the credentials will expire and need to be regenerated.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public KubernetesClusterNodePoolNodesArgs()
        {
        }
    }

    public sealed class KubernetesClusterNodePoolNodesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date and time when the Kubernetes cluster was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// A unique ID that can be used to identify and reference a Kubernetes cluster.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// A name for the Kubernetes cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A string indicating the current status of the cluster. Potential values include running, provisioning, and errored.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The date and time when the Kubernetes cluster was last updated.
        /// * `kube_config.0` - A representation of the Kubernetes cluster's kubeconfig with the following attributes:
        /// - `raw_config` - The full contents of the Kubernetes cluster's kubeconfig file.
        /// - `host` - The URL of the API server on the Kubernetes master node.
        /// - `cluster_ca_certificate` - The base64 encoded public certificate for the cluster's certificate authority.
        /// - `token` - The DigitalOcean API access token used by clients to access the cluster.
        /// - `client_key` - The base64 encoded private key used by clients to access the cluster. Only available if token authentication is not supported on your cluster.
        /// - `client_certificate` - The base64 encoded public certificate used by clients to access the cluster. Only available if token authentication is not supported on your cluster.
        /// - `expires_at` - The date and time when the credentials will expire and need to be regenerated.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public KubernetesClusterNodePoolNodesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class KubernetesClusterKubeConfigs
    {
        public readonly string ClientCertificate;
        public readonly string ClientKey;
        public readonly string ClusterCaCertificate;
        public readonly string ExpiresAt;
        public readonly string Host;
        public readonly string RawConfig;
        public readonly string Token;

        [OutputConstructor]
        private KubernetesClusterKubeConfigs(
            string clientCertificate,
            string clientKey,
            string clusterCaCertificate,
            string expiresAt,
            string host,
            string rawConfig,
            string token)
        {
            ClientCertificate = clientCertificate;
            ClientKey = clientKey;
            ClusterCaCertificate = clusterCaCertificate;
            ExpiresAt = expiresAt;
            Host = host;
            RawConfig = rawConfig;
            Token = token;
        }
    }

    [OutputType]
    public sealed class KubernetesClusterNodePool
    {
        public readonly int ActualNodeCount;
        public readonly bool? AutoScale;
        /// <summary>
        /// A unique ID that can be used to identify and reference a Kubernetes cluster.
        /// </summary>
        public readonly string Id;
        public readonly int? MaxNodes;
        public readonly int? MinNodes;
        /// <summary>
        /// A name for the Kubernetes cluster.
        /// </summary>
        public readonly string Name;
        public readonly int? NodeCount;
        public readonly ImmutableArray<KubernetesClusterNodePoolNodes> Nodes;
        public readonly string Size;
        /// <summary>
        /// A list of tag names to be applied to the Kubernetes cluster.
        /// </summary>
        public readonly ImmutableArray<string> Tags;

        [OutputConstructor]
        private KubernetesClusterNodePool(
            int actualNodeCount,
            bool? autoScale,
            string id,
            int? maxNodes,
            int? minNodes,
            string name,
            int? nodeCount,
            ImmutableArray<KubernetesClusterNodePoolNodes> nodes,
            string size,
            ImmutableArray<string> tags)
        {
            ActualNodeCount = actualNodeCount;
            AutoScale = autoScale;
            Id = id;
            MaxNodes = maxNodes;
            MinNodes = minNodes;
            Name = name;
            NodeCount = nodeCount;
            Nodes = nodes;
            Size = size;
            Tags = tags;
        }
    }

    [OutputType]
    public sealed class KubernetesClusterNodePoolNodes
    {
        /// <summary>
        /// The date and time when the Kubernetes cluster was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// A unique ID that can be used to identify and reference a Kubernetes cluster.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A name for the Kubernetes cluster.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A string indicating the current status of the cluster. Potential values include running, provisioning, and errored.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The date and time when the Kubernetes cluster was last updated.
        /// * `kube_config.0` - A representation of the Kubernetes cluster's kubeconfig with the following attributes:
        /// - `raw_config` - The full contents of the Kubernetes cluster's kubeconfig file.
        /// - `host` - The URL of the API server on the Kubernetes master node.
        /// - `cluster_ca_certificate` - The base64 encoded public certificate for the cluster's certificate authority.
        /// - `token` - The DigitalOcean API access token used by clients to access the cluster.
        /// - `client_key` - The base64 encoded private key used by clients to access the cluster. Only available if token authentication is not supported on your cluster.
        /// - `client_certificate` - The base64 encoded public certificate used by clients to access the cluster. Only available if token authentication is not supported on your cluster.
        /// - `expires_at` - The date and time when the credentials will expire and need to be regenerated.
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private KubernetesClusterNodePoolNodes(
            string createdAt,
            string id,
            string name,
            string status,
            string updatedAt)
        {
            CreatedAt = createdAt;
            Id = id;
            Name = name;
            Status = status;
            UpdatedAt = updatedAt;
        }
    }
    }
}
