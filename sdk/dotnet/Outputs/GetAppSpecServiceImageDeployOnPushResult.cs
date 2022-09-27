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
    public sealed class GetAppSpecServiceImageDeployOnPushResult
    {
        /// <summary>
        /// Whether to automatically deploy images pushed to DOCR.
        /// </summary>
        public readonly bool? Enabled;

        [OutputConstructor]
        private GetAppSpecServiceImageDeployOnPushResult(bool? enabled)
        {
            Enabled = enabled;
        }
    }
}
