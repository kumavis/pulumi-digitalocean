// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean.Inputs
{

    public sealed class KubernetesClusterNodePoolTaintArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// How the node reacts to pods that it won't tolerate. Available effect values are: "NoSchedule", "PreferNoSchedule", "NoExecute".
        /// </summary>
        [Input("effect", required: true)]
        public Input<string> Effect { get; set; } = null!;

        /// <summary>
        /// An arbitrary string. The "key" and "value" fields of the "taint" object form a key-value pair.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// An arbitrary string. The "key" and "value" fields of the "taint" object form a key-value pair.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public KubernetesClusterNodePoolTaintArgs()
        {
        }
    }
}
