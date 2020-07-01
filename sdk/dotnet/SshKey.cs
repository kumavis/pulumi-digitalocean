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
    /// Provides a DigitalOcean SSH key resource to allow you to manage SSH
    /// keys for Droplet access. Keys created with this resource
    /// can be referenced in your Droplet configuration via their ID or
    /// fingerprint.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.IO;
    /// using Pulumi;
    /// using DigitalOcean = Pulumi.DigitalOcean;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // Create a new SSH key
    ///         var @default = new DigitalOcean.SshKey("default", new DigitalOcean.SshKeyArgs
    ///         {
    ///             PublicKey = File.ReadAllText("/Users/myuser/.ssh/id_rsa.pub"),
    ///         });
    ///         // Create a new Droplet using the SSH key
    ///         var web = new DigitalOcean.Droplet("web", new DigitalOcean.DropletArgs
    ///         {
    ///             Image = "ubuntu-18-04-x64",
    ///             Region = "nyc3",
    ///             Size = "s-1vcpu-1gb",
    ///             SshKeys = 
    ///             {
    ///                 @default.Fingerprint,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class SshKey : Pulumi.CustomResource
    {
        /// <summary>
        /// The fingerprint of the SSH key
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// The name of the SSH key for identification
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The public key. If this is a file, it
        /// can be read using the file interpolation function
        /// </summary>
        [Output("publicKey")]
        public Output<string> PublicKey { get; private set; } = null!;


        /// <summary>
        /// Create a SshKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SshKey(string name, SshKeyArgs args, CustomResourceOptions? options = null)
            : base("digitalocean:index/sshKey:SshKey", name, args ?? new SshKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SshKey(string name, Input<string> id, SshKeyState? state = null, CustomResourceOptions? options = null)
            : base("digitalocean:index/sshKey:SshKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SshKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SshKey Get(string name, Input<string> id, SshKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new SshKey(name, id, state, options);
        }
    }

    public sealed class SshKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the SSH key for identification
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The public key. If this is a file, it
        /// can be read using the file interpolation function
        /// </summary>
        [Input("publicKey", required: true)]
        public Input<string> PublicKey { get; set; } = null!;

        public SshKeyArgs()
        {
        }
    }

    public sealed class SshKeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The fingerprint of the SSH key
        /// </summary>
        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        /// <summary>
        /// The name of the SSH key for identification
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The public key. If this is a file, it
        /// can be read using the file interpolation function
        /// </summary>
        [Input("publicKey")]
        public Input<string>? PublicKey { get; set; }

        public SshKeyState()
        {
        }
    }
}
