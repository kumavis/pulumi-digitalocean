// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean
{
    public static class GetApp
    {
        /// <summary>
        /// Get information on a DigitalOcean App.
        /// </summary>
        public static Task<GetAppResult> InvokeAsync(GetAppArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAppResult>("digitalocean:index/getApp:getApp", args ?? new GetAppArgs(), options.WithDefaults());

        /// <summary>
        /// Get information on a DigitalOcean App.
        /// </summary>
        public static Output<GetAppResult> Invoke(GetAppInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAppResult>("digitalocean:index/getApp:getApp", args ?? new GetAppInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the app to retrieve information about.
        /// </summary>
        [Input("appId", required: true)]
        public string AppId { get; set; } = null!;

        public GetAppArgs()
        {
        }
        public static new GetAppArgs Empty => new GetAppArgs();
    }

    public sealed class GetAppInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the app to retrieve information about.
        /// </summary>
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        public GetAppInvokeArgs()
        {
        }
        public static new GetAppInvokeArgs Empty => new GetAppInvokeArgs();
    }


    [OutputType]
    public sealed class GetAppResult
    {
        /// <summary>
        /// The ID the app's currently active deployment.
        /// </summary>
        public readonly string ActiveDeploymentId;
        public readonly string AppId;
        /// <summary>
        /// The date and time of when the app was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The default URL to access the app.
        /// </summary>
        public readonly string DefaultIngress;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The live URL of the app.
        /// </summary>
        public readonly string LiveUrl;
        /// <summary>
        /// A DigitalOcean App spec describing the app.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAppSpecResult> Specs;
        /// <summary>
        /// The date and time of when the app was last updated.
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// The uniform resource identifier for the app.
        /// </summary>
        public readonly string Urn;

        [OutputConstructor]
        private GetAppResult(
            string activeDeploymentId,

            string appId,

            string createdAt,

            string defaultIngress,

            string id,

            string liveUrl,

            ImmutableArray<Outputs.GetAppSpecResult> specs,

            string updatedAt,

            string urn)
        {
            ActiveDeploymentId = activeDeploymentId;
            AppId = appId;
            CreatedAt = createdAt;
            DefaultIngress = defaultIngress;
            Id = id;
            LiveUrl = liveUrl;
            Specs = specs;
            UpdatedAt = updatedAt;
            Urn = urn;
        }
    }
}
