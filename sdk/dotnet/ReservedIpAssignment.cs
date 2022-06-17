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
    /// Provides a resource for assigning an existing DigitalOcean reserved IP to a Droplet. This
    /// makes it easy to provision reserved IP addresses that are not tied to the lifecycle of your
    /// Droplet.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using DigitalOcean = Pulumi.DigitalOcean;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleReservedIp = new DigitalOcean.ReservedIp("exampleReservedIp", new DigitalOcean.ReservedIpArgs
    ///         {
    ///             Region = "nyc3",
    ///         });
    ///         var exampleDroplet = new DigitalOcean.Droplet("exampleDroplet", new DigitalOcean.DropletArgs
    ///         {
    ///             Size = "s-1vcpu-1gb",
    ///             Image = "ubuntu-22-04-x64",
    ///             Region = "nyc3",
    ///             Ipv6 = true,
    ///             PrivateNetworking = true,
    ///         });
    ///         var exampleReservedIpAssignment = new DigitalOcean.ReservedIpAssignment("exampleReservedIpAssignment", new DigitalOcean.ReservedIpAssignmentArgs
    ///         {
    ///             IpAddress = exampleReservedIp.IpAddress,
    ///             DropletId = exampleDroplet.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Reserved IP assignments can be imported using the reserved IP itself and the `id` of the Droplet joined with a comma. For example
    /// 
    /// ```sh
    ///  $ pulumi import digitalocean:index/reservedIpAssignment:ReservedIpAssignment foobar 192.0.2.1,123456
    /// ```
    /// </summary>
    [DigitalOceanResourceType("digitalocean:index/reservedIpAssignment:ReservedIpAssignment")]
    public partial class ReservedIpAssignment : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of Droplet that the reserved IP will be assigned to.
        /// </summary>
        [Output("dropletId")]
        public Output<int> DropletId { get; private set; } = null!;

        /// <summary>
        /// The reserved IP to assign to the Droplet.
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;


        /// <summary>
        /// Create a ReservedIpAssignment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReservedIpAssignment(string name, ReservedIpAssignmentArgs args, CustomResourceOptions? options = null)
            : base("digitalocean:index/reservedIpAssignment:ReservedIpAssignment", name, args ?? new ReservedIpAssignmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReservedIpAssignment(string name, Input<string> id, ReservedIpAssignmentState? state = null, CustomResourceOptions? options = null)
            : base("digitalocean:index/reservedIpAssignment:ReservedIpAssignment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReservedIpAssignment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReservedIpAssignment Get(string name, Input<string> id, ReservedIpAssignmentState? state = null, CustomResourceOptions? options = null)
        {
            return new ReservedIpAssignment(name, id, state, options);
        }
    }

    public sealed class ReservedIpAssignmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of Droplet that the reserved IP will be assigned to.
        /// </summary>
        [Input("dropletId", required: true)]
        public Input<int> DropletId { get; set; } = null!;

        /// <summary>
        /// The reserved IP to assign to the Droplet.
        /// </summary>
        [Input("ipAddress", required: true)]
        public Input<string> IpAddress { get; set; } = null!;

        public ReservedIpAssignmentArgs()
        {
        }
    }

    public sealed class ReservedIpAssignmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of Droplet that the reserved IP will be assigned to.
        /// </summary>
        [Input("dropletId")]
        public Input<int>? DropletId { get; set; }

        /// <summary>
        /// The reserved IP to assign to the Droplet.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        public ReservedIpAssignmentState()
        {
        }
    }
}
