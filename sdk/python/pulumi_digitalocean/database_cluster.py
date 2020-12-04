# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['DatabaseCluster']


class DatabaseCluster(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 eviction_policy: Optional[pulumi.Input[str]] = None,
                 maintenance_windows: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatabaseClusterMaintenanceWindowArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_count: Optional[pulumi.Input[int]] = None,
                 private_network_uuid: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[str]] = None,
                 sql_mode: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a DigitalOcean database cluster resource.

        ## Example Usage
        ### Create a new PostgreSQL database cluster
        ```python
        import pulumi
        import pulumi_digitalocean as digitalocean

        postgres_example = digitalocean.DatabaseCluster("postgres-example",
            engine="pg",
            node_count=1,
            region="nyc1",
            size="db-s-1vcpu-1gb",
            version="11")
        ```
        ### Create a new MySQL database cluster
        ```python
        import pulumi
        import pulumi_digitalocean as digitalocean

        mysql_example = digitalocean.DatabaseCluster("mysql-example",
            engine="mysql",
            node_count=1,
            region="nyc1",
            size="db-s-1vcpu-1gb",
            version="8")
        ```
        ### Create a new Redis database cluster
        ```python
        import pulumi
        import pulumi_digitalocean as digitalocean

        redis_example = digitalocean.DatabaseCluster("redis-example",
            engine="redis",
            node_count=1,
            region="nyc1",
            size="db-s-1vcpu-1gb",
            version="6")
        ```

        ## Import

        Database clusters can be imported using the `id` returned from DigitalOcean, e.g.

        ```sh
         $ pulumi import digitalocean:index/databaseCluster:DatabaseCluster mycluster 245bcfd0-7f31-4ce6-a2bc-475a116cca97
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] engine: Database engine used by the cluster (ex. `pg` for PostreSQL, `mysql` for MySQL, or `redis` for Redis).
        :param pulumi.Input[str] eviction_policy: A string specifying the eviction policy for a Redis cluster. Valid values are: `noeviction`, `allkeys_lru`, `allkeys_random`, `volatile_lru`, `volatile_random`, or `volatile_ttl`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatabaseClusterMaintenanceWindowArgs']]]] maintenance_windows: Defines when the automatic maintenance should be performed for the database cluster.
        :param pulumi.Input[str] name: The name of the database cluster.
        :param pulumi.Input[int] node_count: Number of nodes that will be included in the cluster.
        :param pulumi.Input[str] private_network_uuid: The ID of the VPC where the database cluster will be located.
        :param pulumi.Input[str] region: DigitalOcean region where the cluster will reside.
        :param pulumi.Input[str] size: Database Droplet size associated with the cluster (ex. `db-s-1vcpu-1gb`). See here for a [list of valid size slugs](https://developers.digitalocean.com/documentation/v2/#databases).
        :param pulumi.Input[str] sql_mode: A comma separated string specifying the  SQL modes for a MySQL cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tag names to be applied to the database cluster.
        :param pulumi.Input[str] version: Engine version used by the cluster (ex. `11` for PostgreSQL 11).
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

            if engine is None:
                raise TypeError("Missing required property 'engine'")
            __props__['engine'] = engine
            __props__['eviction_policy'] = eviction_policy
            __props__['maintenance_windows'] = maintenance_windows
            __props__['name'] = name
            if node_count is None:
                raise TypeError("Missing required property 'node_count'")
            __props__['node_count'] = node_count
            __props__['private_network_uuid'] = private_network_uuid
            if region is None:
                raise TypeError("Missing required property 'region'")
            __props__['region'] = region
            if size is None:
                raise TypeError("Missing required property 'size'")
            __props__['size'] = size
            __props__['sql_mode'] = sql_mode
            __props__['tags'] = tags
            __props__['version'] = version
            __props__['cluster_urn'] = None
            __props__['database'] = None
            __props__['host'] = None
            __props__['password'] = None
            __props__['port'] = None
            __props__['private_host'] = None
            __props__['private_uri'] = None
            __props__['uri'] = None
            __props__['user'] = None
        super(DatabaseCluster, __self__).__init__(
            'digitalocean:index/databaseCluster:DatabaseCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_urn: Optional[pulumi.Input[str]] = None,
            database: Optional[pulumi.Input[str]] = None,
            engine: Optional[pulumi.Input[str]] = None,
            eviction_policy: Optional[pulumi.Input[str]] = None,
            host: Optional[pulumi.Input[str]] = None,
            maintenance_windows: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatabaseClusterMaintenanceWindowArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            node_count: Optional[pulumi.Input[int]] = None,
            password: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            private_host: Optional[pulumi.Input[str]] = None,
            private_network_uuid: Optional[pulumi.Input[str]] = None,
            private_uri: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[str]] = None,
            sql_mode: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            uri: Optional[pulumi.Input[str]] = None,
            user: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'DatabaseCluster':
        """
        Get an existing DatabaseCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_urn: The uniform resource name of the database cluster.
        :param pulumi.Input[str] database: Name of the cluster's default database.
        :param pulumi.Input[str] engine: Database engine used by the cluster (ex. `pg` for PostreSQL, `mysql` for MySQL, or `redis` for Redis).
        :param pulumi.Input[str] eviction_policy: A string specifying the eviction policy for a Redis cluster. Valid values are: `noeviction`, `allkeys_lru`, `allkeys_random`, `volatile_lru`, `volatile_random`, or `volatile_ttl`.
        :param pulumi.Input[str] host: Database cluster's hostname.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatabaseClusterMaintenanceWindowArgs']]]] maintenance_windows: Defines when the automatic maintenance should be performed for the database cluster.
        :param pulumi.Input[str] name: The name of the database cluster.
        :param pulumi.Input[int] node_count: Number of nodes that will be included in the cluster.
        :param pulumi.Input[str] password: Password for the cluster's default user.
        :param pulumi.Input[int] port: Network port that the database cluster is listening on.
        :param pulumi.Input[str] private_host: Same as `host`, but only accessible from resources within the account and in the same region.
        :param pulumi.Input[str] private_network_uuid: The ID of the VPC where the database cluster will be located.
        :param pulumi.Input[str] private_uri: Same as `uri`, but only accessible from resources within the account and in the same region.
        :param pulumi.Input[str] region: DigitalOcean region where the cluster will reside.
        :param pulumi.Input[str] size: Database Droplet size associated with the cluster (ex. `db-s-1vcpu-1gb`). See here for a [list of valid size slugs](https://developers.digitalocean.com/documentation/v2/#databases).
        :param pulumi.Input[str] sql_mode: A comma separated string specifying the  SQL modes for a MySQL cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tag names to be applied to the database cluster.
        :param pulumi.Input[str] uri: The full URI for connecting to the database cluster.
        :param pulumi.Input[str] user: Username for the cluster's default user.
        :param pulumi.Input[str] version: Engine version used by the cluster (ex. `11` for PostgreSQL 11).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cluster_urn"] = cluster_urn
        __props__["database"] = database
        __props__["engine"] = engine
        __props__["eviction_policy"] = eviction_policy
        __props__["host"] = host
        __props__["maintenance_windows"] = maintenance_windows
        __props__["name"] = name
        __props__["node_count"] = node_count
        __props__["password"] = password
        __props__["port"] = port
        __props__["private_host"] = private_host
        __props__["private_network_uuid"] = private_network_uuid
        __props__["private_uri"] = private_uri
        __props__["region"] = region
        __props__["size"] = size
        __props__["sql_mode"] = sql_mode
        __props__["tags"] = tags
        __props__["uri"] = uri
        __props__["user"] = user
        __props__["version"] = version
        return DatabaseCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterUrn")
    def cluster_urn(self) -> pulumi.Output[str]:
        """
        The uniform resource name of the database cluster.
        """
        return pulumi.get(self, "cluster_urn")

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output[str]:
        """
        Name of the cluster's default database.
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Output[str]:
        """
        Database engine used by the cluster (ex. `pg` for PostreSQL, `mysql` for MySQL, or `redis` for Redis).
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter(name="evictionPolicy")
    def eviction_policy(self) -> pulumi.Output[Optional[str]]:
        """
        A string specifying the eviction policy for a Redis cluster. Valid values are: `noeviction`, `allkeys_lru`, `allkeys_random`, `volatile_lru`, `volatile_random`, or `volatile_ttl`.
        """
        return pulumi.get(self, "eviction_policy")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[str]:
        """
        Database cluster's hostname.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="maintenanceWindows")
    def maintenance_windows(self) -> pulumi.Output[Optional[Sequence['outputs.DatabaseClusterMaintenanceWindow']]]:
        """
        Defines when the automatic maintenance should be performed for the database cluster.
        """
        return pulumi.get(self, "maintenance_windows")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the database cluster.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeCount")
    def node_count(self) -> pulumi.Output[int]:
        """
        Number of nodes that will be included in the cluster.
        """
        return pulumi.get(self, "node_count")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        Password for the cluster's default user.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        Network port that the database cluster is listening on.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="privateHost")
    def private_host(self) -> pulumi.Output[str]:
        """
        Same as `host`, but only accessible from resources within the account and in the same region.
        """
        return pulumi.get(self, "private_host")

    @property
    @pulumi.getter(name="privateNetworkUuid")
    def private_network_uuid(self) -> pulumi.Output[str]:
        """
        The ID of the VPC where the database cluster will be located.
        """
        return pulumi.get(self, "private_network_uuid")

    @property
    @pulumi.getter(name="privateUri")
    def private_uri(self) -> pulumi.Output[str]:
        """
        Same as `uri`, but only accessible from resources within the account and in the same region.
        """
        return pulumi.get(self, "private_uri")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        DigitalOcean region where the cluster will reside.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[str]:
        """
        Database Droplet size associated with the cluster (ex. `db-s-1vcpu-1gb`). See here for a [list of valid size slugs](https://developers.digitalocean.com/documentation/v2/#databases).
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="sqlMode")
    def sql_mode(self) -> pulumi.Output[Optional[str]]:
        """
        A comma separated string specifying the  SQL modes for a MySQL cluster.
        """
        return pulumi.get(self, "sql_mode")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of tag names to be applied to the database cluster.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def uri(self) -> pulumi.Output[str]:
        """
        The full URI for connecting to the database cluster.
        """
        return pulumi.get(self, "uri")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[str]:
        """
        Username for the cluster's default user.
        """
        return pulumi.get(self, "user")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[Optional[str]]:
        """
        Engine version used by the cluster (ex. `11` for PostgreSQL 11).
        """
        return pulumi.get(self, "version")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

