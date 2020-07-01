# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class ContainerRegistryDockerCredentials(pulumi.CustomResource):
    credential_expiration_time: pulumi.Output[str]
    docker_credentials: pulumi.Output[str]
    expiry_seconds: pulumi.Output[float]
    """
    The amount of time to pass before the Docker credentials expire in seconds. Defaults to 2147483647, or roughly 68 years. Must be greater than 0 and less than 2147483647.
    """
    registry_name: pulumi.Output[str]
    """
    The name of the container registry.
    """
    write: pulumi.Output[bool]
    """
    Allow for write access to the container registry. Defaults to false.
    """
    def __init__(__self__, resource_name, opts=None, expiry_seconds=None, registry_name=None, write=None, __props__=None, __name__=None, __opts__=None):
        """
        Get Docker credentials for your DigitalOcean container registry.

        An error is triggered if the provided container registry name does not exist.

        ## Example Usage
        ### Basic Example

        Get the container registry:

        ```python
        import pulumi
        import pulumi_digitalocean as digitalocean

        example = digitalocean.ContainerRegistryDockerCredentials("example", registry_name="example")
        ```
        ### Docker Provider Example

        Use the `endpoint` and `docker_credentials` with the Docker provider:

        ```python
        import pulumi
        import pulumi_digitalocean as digitalocean

        example_container_registry = digitalocean.get_container_registry(name="example")
        example_container_registry_docker_credentials = digitalocean.ContainerRegistryDockerCredentials("exampleContainerRegistryDockerCredentials", registry_name="example")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] expiry_seconds: The amount of time to pass before the Docker credentials expire in seconds. Defaults to 2147483647, or roughly 68 years. Must be greater than 0 and less than 2147483647.
        :param pulumi.Input[str] registry_name: The name of the container registry.
        :param pulumi.Input[bool] write: Allow for write access to the container registry. Defaults to false.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['expiry_seconds'] = expiry_seconds
            if registry_name is None:
                raise TypeError("Missing required property 'registry_name'")
            __props__['registry_name'] = registry_name
            __props__['write'] = write
            __props__['credential_expiration_time'] = None
            __props__['docker_credentials'] = None
        super(ContainerRegistryDockerCredentials, __self__).__init__(
            'digitalocean:index/containerRegistryDockerCredentials:ContainerRegistryDockerCredentials',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, credential_expiration_time=None, docker_credentials=None, expiry_seconds=None, registry_name=None, write=None):
        """
        Get an existing ContainerRegistryDockerCredentials resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] expiry_seconds: The amount of time to pass before the Docker credentials expire in seconds. Defaults to 2147483647, or roughly 68 years. Must be greater than 0 and less than 2147483647.
        :param pulumi.Input[str] registry_name: The name of the container registry.
        :param pulumi.Input[bool] write: Allow for write access to the container registry. Defaults to false.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["credential_expiration_time"] = credential_expiration_time
        __props__["docker_credentials"] = docker_credentials
        __props__["expiry_seconds"] = expiry_seconds
        __props__["registry_name"] = registry_name
        __props__["write"] = write
        return ContainerRegistryDockerCredentials(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
