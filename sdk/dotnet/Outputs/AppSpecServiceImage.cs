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
    public sealed class AppSpecServiceImage
    {
        /// <summary>
        /// Configures automatically deploying images pushed to DOCR.
        /// </summary>
        public readonly ImmutableArray<Outputs.AppSpecServiceImageDeployOnPush> DeployOnPushes;
        /// <summary>
        /// The registry name. Must be left empty for the `DOCR` registry type. Required for the `DOCKER_HUB` registry type.
        /// </summary>
        public readonly string? Registry;
        /// <summary>
        /// Access credentials for third-party registries
        /// </summary>
        public readonly string RegistryCredentials;
        /// <summary>
        /// The registry type. One of `DOCR` (DigitalOcean container registry) or `DOCKER_HUB`.
        /// </summary>
        public readonly string RegistryType;
        /// <summary>
        /// The repository name.
        /// </summary>
        public readonly string Repository;
        /// <summary>
        /// The repository tag. Defaults to `latest` if not provided.
        /// </summary>
        public readonly string? Tag;

        [OutputConstructor]
        private AppSpecServiceImage(
            ImmutableArray<Outputs.AppSpecServiceImageDeployOnPush> deployOnPushes,

            string? registry,

            string registryCredentials,

            string registryType,

            string repository,

            string? tag)
        {
            DeployOnPushes = deployOnPushes;
            Registry = registry;
            RegistryCredentials = registryCredentials;
            RegistryType = registryType;
            Repository = repository;
            Tag = tag;
        }
    }
}
