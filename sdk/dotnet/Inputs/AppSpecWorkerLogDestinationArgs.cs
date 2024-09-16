// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean.Inputs
{

    public sealed class AppSpecWorkerLogDestinationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Datadog configuration.
        /// </summary>
        [Input("datadog")]
        public Input<Inputs.AppSpecWorkerLogDestinationDatadogArgs>? Datadog { get; set; }

        /// <summary>
        /// Logtail configuration.
        /// </summary>
        [Input("logtail")]
        public Input<Inputs.AppSpecWorkerLogDestinationLogtailArgs>? Logtail { get; set; }

        /// <summary>
        /// Name of the log destination. Minimum length: 2. Maximum length: 42.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// OpenSearch configuration.
        /// </summary>
        [Input("openSearch")]
        public Input<Inputs.AppSpecWorkerLogDestinationOpenSearchArgs>? OpenSearch { get; set; }

        /// <summary>
        /// Papertrail configuration.
        /// </summary>
        [Input("papertrail")]
        public Input<Inputs.AppSpecWorkerLogDestinationPapertrailArgs>? Papertrail { get; set; }

        public AppSpecWorkerLogDestinationArgs()
        {
        }
        public static new AppSpecWorkerLogDestinationArgs Empty => new AppSpecWorkerLogDestinationArgs();
    }
}
