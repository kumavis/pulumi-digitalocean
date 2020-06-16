// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean.Inputs
{

    public sealed class GetTagsSortArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The sort direction. This may be either `asc` or `desc`.
        /// </summary>
        [Input("direction")]
        public string? Direction { get; set; }

        /// <summary>
        /// Sort the tags by this key. This may be one of `name`, `total_resource_count`,  `droplets_count`, `images_count`, `volumes_count`, `volume_snapshots_count`, or `databases_count`.
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        public GetTagsSortArgs()
        {
        }
    }
}
