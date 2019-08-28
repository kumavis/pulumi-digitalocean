# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class FloatingIpAssignment(pulumi.CustomResource):
    droplet_id: pulumi.Output[float]
    """
    The ID of Droplet that the Floating IP will be assigned to.
    """
    ip_address: pulumi.Output[str]
    """
    The Floating IP to assign to the Droplet.
    """
    def __init__(__self__, resource_name, opts=None, droplet_id=None, ip_address=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a resource for assigning an existing DigitalOcean Floating IP to a Droplet. This
        makes it easy to provision floating IP addresses that are not tied to the lifecycle of your
        Droplet.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] droplet_id: The ID of Droplet that the Floating IP will be assigned to.
        :param pulumi.Input[str] ip_address: The Floating IP to assign to the Droplet.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/r/floating_ip_assignment.html.markdown.
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

            if droplet_id is None:
                raise TypeError("Missing required property 'droplet_id'")
            __props__['droplet_id'] = droplet_id
            if ip_address is None:
                raise TypeError("Missing required property 'ip_address'")
            __props__['ip_address'] = ip_address
        super(FloatingIpAssignment, __self__).__init__(
            'digitalocean:index/floatingIpAssignment:FloatingIpAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, droplet_id=None, ip_address=None):
        """
        Get an existing FloatingIpAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] droplet_id: The ID of Droplet that the Floating IP will be assigned to.
        :param pulumi.Input[str] ip_address: The Floating IP to assign to the Droplet.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/r/floating_ip_assignment.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["droplet_id"] = droplet_id
        __props__["ip_address"] = ip_address
        return FloatingIpAssignment(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

