// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean.Outputs
{

    [OutputType]
    public sealed class GetKubernetesClusterNodePoolResult
    {
        /// <summary>
        /// The actual number of nodes in the node pool, which is especially useful when auto-scaling is enabled.
        /// </summary>
        public readonly int ActualNodeCount;
        /// <summary>
        /// A boolean indicating whether auto-scaling is enabled on the node pool.
        /// </summary>
        public readonly bool AutoScale;
        /// <summary>
        /// A unique ID that can be used to identify and reference the node.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A map of key/value pairs applied to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/).
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to.
        /// </summary>
        public readonly int MaxNodes;
        /// <summary>
        /// If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to.
        /// </summary>
        public readonly int MinNodes;
        /// <summary>
        /// The name of Kubernetes cluster.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The number of Droplet instances in the node pool.
        /// </summary>
        public readonly int NodeCount;
        /// <summary>
        /// A list of nodes in the pool. Each node exports the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesClusterNodePoolNodeResult> Nodes;
        /// <summary>
        /// The slug identifier for the type of Droplet used as workers in the node pool.
        /// </summary>
        public readonly string Size;
        /// <summary>
        /// A list of tag names applied to the node pool.
        /// </summary>
        public readonly ImmutableArray<string> Tags;

        [OutputConstructor]
        private GetKubernetesClusterNodePoolResult(
            int actualNodeCount,

            bool autoScale,

            string id,

            ImmutableDictionary<string, string> labels,

            int maxNodes,

            int minNodes,

            string name,

            int nodeCount,

            ImmutableArray<Outputs.GetKubernetesClusterNodePoolNodeResult> nodes,

            string size,

            ImmutableArray<string> tags)
        {
            ActualNodeCount = actualNodeCount;
            AutoScale = autoScale;
            Id = id;
            Labels = labels;
            MaxNodes = maxNodes;
            MinNodes = minNodes;
            Name = name;
            NodeCount = nodeCount;
            Nodes = nodes;
            Size = size;
            Tags = tags;
        }
    }
}
