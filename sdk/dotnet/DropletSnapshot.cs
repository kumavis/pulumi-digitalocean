// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean
{
    /// <summary>
    /// Provides a resource which can be used to create a snapshot from an existing DigitalOcean Droplet.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using DigitalOcean = Pulumi.DigitalOcean;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var web = new DigitalOcean.Droplet("web", new()
    ///     {
    ///         Size = "s-1vcpu-1gb",
    ///         Image = "ubuntu-22-04-x64",
    ///         Region = "nyc3",
    ///     });
    /// 
    ///     var web_snapshot = new DigitalOcean.DropletSnapshot("web-snapshot", new()
    ///     {
    ///         DropletId = web.Id,
    ///     });
    /// 
    ///     var from_snapshot = new DigitalOcean.Droplet("from-snapshot", new()
    ///     {
    ///         Image = web_snapshot.Id,
    ///         Region = "nyc3",
    ///         Size = "s-2vcpu-4gb",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Droplet Snapshots can be imported using the `snapshot id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import digitalocean:index/dropletSnapshot:DropletSnapshot mysnapshot 123456
    /// ```
    /// </summary>
    [DigitalOceanResourceType("digitalocean:index/dropletSnapshot:DropletSnapshot")]
    public partial class DropletSnapshot : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The date and time the Droplet snapshot was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The ID of the Droplet from which the snapshot will be taken.
        /// </summary>
        [Output("dropletId")]
        public Output<string> DropletId { get; private set; } = null!;

        /// <summary>
        /// The minimum size in gigabytes required for a Droplet to be created based on this snapshot.
        /// </summary>
        [Output("minDiskSize")]
        public Output<int> MinDiskSize { get; private set; } = null!;

        /// <summary>
        /// A name for the Droplet snapshot.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of DigitalOcean region "slugs" indicating where the droplet snapshot is available.
        /// </summary>
        [Output("regions")]
        public Output<ImmutableArray<string>> Regions { get; private set; } = null!;

        /// <summary>
        /// The billable size of the Droplet snapshot in gigabytes.
        /// </summary>
        [Output("size")]
        public Output<double> Size { get; private set; } = null!;


        /// <summary>
        /// Create a DropletSnapshot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DropletSnapshot(string name, DropletSnapshotArgs args, CustomResourceOptions? options = null)
            : base("digitalocean:index/dropletSnapshot:DropletSnapshot", name, args ?? new DropletSnapshotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DropletSnapshot(string name, Input<string> id, DropletSnapshotState? state = null, CustomResourceOptions? options = null)
            : base("digitalocean:index/dropletSnapshot:DropletSnapshot", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DropletSnapshot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DropletSnapshot Get(string name, Input<string> id, DropletSnapshotState? state = null, CustomResourceOptions? options = null)
        {
            return new DropletSnapshot(name, id, state, options);
        }
    }

    public sealed class DropletSnapshotArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Droplet from which the snapshot will be taken.
        /// </summary>
        [Input("dropletId", required: true)]
        public Input<string> DropletId { get; set; } = null!;

        /// <summary>
        /// A name for the Droplet snapshot.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public DropletSnapshotArgs()
        {
        }
        public static new DropletSnapshotArgs Empty => new DropletSnapshotArgs();
    }

    public sealed class DropletSnapshotState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date and time the Droplet snapshot was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The ID of the Droplet from which the snapshot will be taken.
        /// </summary>
        [Input("dropletId")]
        public Input<string>? DropletId { get; set; }

        /// <summary>
        /// The minimum size in gigabytes required for a Droplet to be created based on this snapshot.
        /// </summary>
        [Input("minDiskSize")]
        public Input<int>? MinDiskSize { get; set; }

        /// <summary>
        /// A name for the Droplet snapshot.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("regions")]
        private InputList<string>? _regions;

        /// <summary>
        /// A list of DigitalOcean region "slugs" indicating where the droplet snapshot is available.
        /// </summary>
        public InputList<string> Regions
        {
            get => _regions ?? (_regions = new InputList<string>());
            set => _regions = value;
        }

        /// <summary>
        /// The billable size of the Droplet snapshot in gigabytes.
        /// </summary>
        [Input("size")]
        public Input<double>? Size { get; set; }

        public DropletSnapshotState()
        {
        }
        public static new DropletSnapshotState Empty => new DropletSnapshotState();
    }
}
