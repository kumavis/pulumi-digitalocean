# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['ContainerRegistry']


class ContainerRegistry(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 subscription_tier_slug: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a DigitalOcean Container Registry resource. A Container Registry is
        a secure, private location to store your containers for rapid deployment.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_digitalocean as digitalocean

        # Create a new container registry
        foobar = digitalocean.ContainerRegistry("foobar", subscription_tier_slug="starter")
        ```

        ## Import

        Container Registries can be imported using the `name`, e.g.

        ```sh
         $ pulumi import digitalocean:index/containerRegistry:ContainerRegistry myregistry registryname
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the container_registry
        :param pulumi.Input[str] subscription_tier_slug: The slug identifier for the subscription tier to use (`starter`, `basic`, or `professional`)
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['name'] = name
            if subscription_tier_slug is None and not opts.urn:
                raise TypeError("Missing required property 'subscription_tier_slug'")
            __props__['subscription_tier_slug'] = subscription_tier_slug
            __props__['endpoint'] = None
            __props__['server_url'] = None
        super(ContainerRegistry, __self__).__init__(
            'digitalocean:index/containerRegistry:ContainerRegistry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            server_url: Optional[pulumi.Input[str]] = None,
            subscription_tier_slug: Optional[pulumi.Input[str]] = None) -> 'ContainerRegistry':
        """
        Get an existing ContainerRegistry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the container_registry
        :param pulumi.Input[str] subscription_tier_slug: The slug identifier for the subscription tier to use (`starter`, `basic`, or `professional`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["endpoint"] = endpoint
        __props__["name"] = name
        __props__["server_url"] = server_url
        __props__["subscription_tier_slug"] = subscription_tier_slug
        return ContainerRegistry(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the container_registry
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serverUrl")
    def server_url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "server_url")

    @property
    @pulumi.getter(name="subscriptionTierSlug")
    def subscription_tier_slug(self) -> pulumi.Output[str]:
        """
        The slug identifier for the subscription tier to use (`starter`, `basic`, or `professional`)
        """
        return pulumi.get(self, "subscription_tier_slug")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

