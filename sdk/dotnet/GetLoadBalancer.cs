// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean
{
    public static class GetLoadBalancer
    {
        public static Task<GetLoadBalancerResult> InvokeAsync(GetLoadBalancerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLoadBalancerResult>("digitalocean:index/getLoadBalancer:getLoadBalancer", args ?? new GetLoadBalancerArgs(), options.WithVersion());
    }


    public sealed class GetLoadBalancerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of load balancer.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetLoadBalancerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLoadBalancerResult
    {
        public readonly string Algorithm;
        public readonly ImmutableArray<int> DropletIds;
        public readonly string DropletTag;
        public readonly bool EnableBackendKeepalive;
        public readonly bool EnableProxyProtocol;
        public readonly ImmutableArray<Outputs.GetLoadBalancerForwardingRuleResult> ForwardingRules;
        public readonly ImmutableArray<Outputs.GetLoadBalancerHealthcheckResult> Healthchecks;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Ip;
        public readonly string LoadBalancerUrn;
        public readonly string Name;
        public readonly bool RedirectHttpToHttps;
        public readonly string Region;
        public readonly string Size;
        public readonly string Status;
        public readonly ImmutableArray<Outputs.GetLoadBalancerStickySessionResult> StickySessions;
        public readonly string VpcUuid;

        [OutputConstructor]
        private GetLoadBalancerResult(
            string algorithm,

            ImmutableArray<int> dropletIds,

            string dropletTag,

            bool enableBackendKeepalive,

            bool enableProxyProtocol,

            ImmutableArray<Outputs.GetLoadBalancerForwardingRuleResult> forwardingRules,

            ImmutableArray<Outputs.GetLoadBalancerHealthcheckResult> healthchecks,

            string id,

            string ip,

            string loadBalancerUrn,

            string name,

            bool redirectHttpToHttps,

            string region,

            string size,

            string status,

            ImmutableArray<Outputs.GetLoadBalancerStickySessionResult> stickySessions,

            string vpcUuid)
        {
            Algorithm = algorithm;
            DropletIds = dropletIds;
            DropletTag = dropletTag;
            EnableBackendKeepalive = enableBackendKeepalive;
            EnableProxyProtocol = enableProxyProtocol;
            ForwardingRules = forwardingRules;
            Healthchecks = healthchecks;
            Id = id;
            Ip = ip;
            LoadBalancerUrn = loadBalancerUrn;
            Name = name;
            RedirectHttpToHttps = redirectHttpToHttps;
            Region = region;
            Size = size;
            Status = status;
            StickySessions = stickySessions;
            VpcUuid = vpcUuid;
        }
    }
}
