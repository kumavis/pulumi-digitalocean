# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['DropletSnapshot']


class DropletSnapshot(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 droplet_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a resource which can be used to create a snapshot from an existing DigitalOcean Droplet.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_digitalocean as digitalocean

        web = digitalocean.Droplet("web",
            size="s-1vcpu-1gb",
            image="centos-7-x64",
            region="nyc3")
        web_snapshot = digitalocean.DropletSnapshot("web-snapshot", droplet_id=web.id)
        ```

        ## Import

        Droplet Snapshots can be imported using the `snapshot id`, e.g.

        ```sh
         $ pulumi import digitalocean:index/dropletSnapshot:DropletSnapshot mysnapshot 123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] droplet_id: The ID of the Droplet from which the snapshot will be taken.
        :param pulumi.Input[str] name: A name for the Droplet snapshot.
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

            if droplet_id is None and not opts.urn:
                raise TypeError("Missing required property 'droplet_id'")
            __props__['droplet_id'] = droplet_id
            __props__['name'] = name
            __props__['created_at'] = None
            __props__['min_disk_size'] = None
            __props__['regions'] = None
            __props__['size'] = None
        super(DropletSnapshot, __self__).__init__(
            'digitalocean:index/dropletSnapshot:DropletSnapshot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            droplet_id: Optional[pulumi.Input[str]] = None,
            min_disk_size: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            size: Optional[pulumi.Input[float]] = None) -> 'DropletSnapshot':
        """
        Get an existing DropletSnapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The date and time the Droplet snapshot was created.
        :param pulumi.Input[str] droplet_id: The ID of the Droplet from which the snapshot will be taken.
        :param pulumi.Input[int] min_disk_size: The minimum size in gigabytes required for a Droplet to be created based on this snapshot.
        :param pulumi.Input[str] name: A name for the Droplet snapshot.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] regions: A list of DigitalOcean region "slugs" indicating where the droplet snapshot is available.
        :param pulumi.Input[float] size: The billable size of the Droplet snapshot in gigabytes.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["created_at"] = created_at
        __props__["droplet_id"] = droplet_id
        __props__["min_disk_size"] = min_disk_size
        __props__["name"] = name
        __props__["regions"] = regions
        __props__["size"] = size
        return DropletSnapshot(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The date and time the Droplet snapshot was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dropletId")
    def droplet_id(self) -> pulumi.Output[str]:
        """
        The ID of the Droplet from which the snapshot will be taken.
        """
        return pulumi.get(self, "droplet_id")

    @property
    @pulumi.getter(name="minDiskSize")
    def min_disk_size(self) -> pulumi.Output[int]:
        """
        The minimum size in gigabytes required for a Droplet to be created based on this snapshot.
        """
        return pulumi.get(self, "min_disk_size")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A name for the Droplet snapshot.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def regions(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of DigitalOcean region "slugs" indicating where the droplet snapshot is available.
        """
        return pulumi.get(self, "regions")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[float]:
        """
        The billable size of the Droplet snapshot in gigabytes.
        """
        return pulumi.get(self, "size")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

