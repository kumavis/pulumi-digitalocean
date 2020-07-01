# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class DatabaseConnectionPool(pulumi.CustomResource):
    cluster_id: pulumi.Output[str]
    """
    The ID of the source database cluster. Note: This must be a PostgreSQL cluster.
    """
    db_name: pulumi.Output[str]
    """
    The database for use with the connection pool.
    """
    host: pulumi.Output[str]
    """
    The hostname used to connect to the database connection pool.
    """
    mode: pulumi.Output[str]
    """
    The PGBouncer transaction mode for the connection pool. The allowed values are session, transaction, and statement.
    """
    name: pulumi.Output[str]
    """
    The name for the database connection pool.
    """
    password: pulumi.Output[str]
    """
    Password for the connection pool's user.
    """
    port: pulumi.Output[float]
    """
    Network port that the database connection pool is listening on.
    """
    private_host: pulumi.Output[str]
    """
    Same as `host`, but only accessible from resources within the account and in the same region.
    """
    private_uri: pulumi.Output[str]
    """
    Same as `uri`, but only accessible from resources within the account and in the same region.
    """
    size: pulumi.Output[float]
    """
    The desired size of the PGBouncer connection pool.
    """
    uri: pulumi.Output[str]
    """
    The full URI for connecting to the database connection pool.
    """
    user: pulumi.Output[str]
    """
    The name of the database user for use with the connection pool.
    """
    def __init__(__self__, resource_name, opts=None, cluster_id=None, db_name=None, mode=None, name=None, size=None, user=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a DigitalOcean database connection pool resource.

        ## Example Usage
        ### Create a new PostgreSQL database connection pool
        ```python
        import pulumi
        import pulumi_digitalocean as digitalocean

        postgres_example = digitalocean.DatabaseCluster("postgres-example",
            engine="pg",
            version="11",
            size="db-s-1vcpu-1gb",
            region="nyc1",
            node_count=1)
        pool_01 = digitalocean.DatabaseConnectionPool("pool-01",
            cluster_id=postgres_example.id,
            mode="transaction",
            size=20,
            db_name="defaultdb",
            user="doadmin")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: The ID of the source database cluster. Note: This must be a PostgreSQL cluster.
        :param pulumi.Input[str] db_name: The database for use with the connection pool.
        :param pulumi.Input[str] mode: The PGBouncer transaction mode for the connection pool. The allowed values are session, transaction, and statement.
        :param pulumi.Input[str] name: The name for the database connection pool.
        :param pulumi.Input[float] size: The desired size of the PGBouncer connection pool.
        :param pulumi.Input[str] user: The name of the database user for use with the connection pool.
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

            if cluster_id is None:
                raise TypeError("Missing required property 'cluster_id'")
            __props__['cluster_id'] = cluster_id
            if db_name is None:
                raise TypeError("Missing required property 'db_name'")
            __props__['db_name'] = db_name
            if mode is None:
                raise TypeError("Missing required property 'mode'")
            __props__['mode'] = mode
            __props__['name'] = name
            if size is None:
                raise TypeError("Missing required property 'size'")
            __props__['size'] = size
            if user is None:
                raise TypeError("Missing required property 'user'")
            __props__['user'] = user
            __props__['host'] = None
            __props__['password'] = None
            __props__['port'] = None
            __props__['private_host'] = None
            __props__['private_uri'] = None
            __props__['uri'] = None
        super(DatabaseConnectionPool, __self__).__init__(
            'digitalocean:index/databaseConnectionPool:DatabaseConnectionPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, cluster_id=None, db_name=None, host=None, mode=None, name=None, password=None, port=None, private_host=None, private_uri=None, size=None, uri=None, user=None):
        """
        Get an existing DatabaseConnectionPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: The ID of the source database cluster. Note: This must be a PostgreSQL cluster.
        :param pulumi.Input[str] db_name: The database for use with the connection pool.
        :param pulumi.Input[str] host: The hostname used to connect to the database connection pool.
        :param pulumi.Input[str] mode: The PGBouncer transaction mode for the connection pool. The allowed values are session, transaction, and statement.
        :param pulumi.Input[str] name: The name for the database connection pool.
        :param pulumi.Input[str] password: Password for the connection pool's user.
        :param pulumi.Input[float] port: Network port that the database connection pool is listening on.
        :param pulumi.Input[str] private_host: Same as `host`, but only accessible from resources within the account and in the same region.
        :param pulumi.Input[str] private_uri: Same as `uri`, but only accessible from resources within the account and in the same region.
        :param pulumi.Input[float] size: The desired size of the PGBouncer connection pool.
        :param pulumi.Input[str] uri: The full URI for connecting to the database connection pool.
        :param pulumi.Input[str] user: The name of the database user for use with the connection pool.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cluster_id"] = cluster_id
        __props__["db_name"] = db_name
        __props__["host"] = host
        __props__["mode"] = mode
        __props__["name"] = name
        __props__["password"] = password
        __props__["port"] = port
        __props__["private_host"] = private_host
        __props__["private_uri"] = private_uri
        __props__["size"] = size
        __props__["uri"] = uri
        __props__["user"] = user
        return DatabaseConnectionPool(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
