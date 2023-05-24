// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * Provides a DigitalOcean database cluster resource.
 *
 * ## Example Usage
 * ### Create a new PostgreSQL database cluster
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const postgres_example = new digitalocean.DatabaseCluster("postgres-example", {
 *     engine: "pg",
 *     nodeCount: 1,
 *     region: "nyc1",
 *     size: "db-s-1vcpu-1gb",
 *     version: "11",
 * });
 * ```
 * ### Create a new MySQL database cluster
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const mysql_example = new digitalocean.DatabaseCluster("mysql-example", {
 *     engine: "mysql",
 *     nodeCount: 1,
 *     region: "nyc1",
 *     size: "db-s-1vcpu-1gb",
 *     version: "8",
 * });
 * ```
 * ### Create a new Redis database cluster
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const redis_example = new digitalocean.DatabaseCluster("redis-example", {
 *     engine: "redis",
 *     nodeCount: 1,
 *     region: "nyc1",
 *     size: "db-s-1vcpu-1gb",
 *     version: "6",
 * });
 * ```
 * ### Create a new MongoDB database cluster
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const mongodb_example = new digitalocean.DatabaseCluster("mongodb-example", {
 *     engine: "mongodb",
 *     nodeCount: 1,
 *     region: "nyc3",
 *     size: "db-s-1vcpu-1gb",
 *     version: "4",
 * });
 * ```
 * ## Create a new database cluster based on a backup of an existing cluster.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const doby = new digitalocean.DatabaseCluster("doby", {
 *     engine: "pg",
 *     version: "11",
 *     size: "db-s-1vcpu-2gb",
 *     region: "nyc1",
 *     nodeCount: 1,
 *     tags: ["production"],
 * });
 * const dobyBackup = new digitalocean.DatabaseCluster("dobyBackup", {
 *     engine: "pg",
 *     version: "11",
 *     size: "db-s-1vcpu-2gb",
 *     region: "nyc1",
 *     nodeCount: 1,
 *     tags: ["production"],
 *     backupRestore: {
 *         databaseName: "dobydb",
 *     },
 * }, {
 *     dependsOn: [doby],
 * });
 * ```
 *
 * ## Import
 *
 * Database clusters can be imported using the `id` returned from DigitalOcean, e.g.
 *
 * ```sh
 *  $ pulumi import digitalocean:index/databaseCluster:DatabaseCluster mycluster 245bcfd0-7f31-4ce6-a2bc-475a116cca97
 * ```
 */
export class DatabaseCluster extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseClusterState, opts?: pulumi.CustomResourceOptions): DatabaseCluster {
        return new DatabaseCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean:index/databaseCluster:DatabaseCluster';

    /**
     * Returns true if the given object is an instance of DatabaseCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseCluster.__pulumiType;
    }

    /**
     * Create a new database cluster based on a backup of an existing cluster.
     */
    public readonly backupRestore!: pulumi.Output<outputs.DatabaseClusterBackupRestore | undefined>;
    /**
     * The uniform resource name of the database cluster.
     */
    public /*out*/ readonly clusterUrn!: pulumi.Output<string>;
    /**
     * Name of the cluster's default database.
     */
    public /*out*/ readonly database!: pulumi.Output<string>;
    /**
     * Database engine used by the cluster (ex. `pg` for PostreSQL, `mysql` for MySQL, `redis` for Redis, or `mongodb` for MongoDB).
     */
    public readonly engine!: pulumi.Output<string>;
    /**
     * A string specifying the eviction policy for a Redis cluster. Valid values are: `noeviction`, `allkeysLru`, `allkeysRandom`, `volatileLru`, `volatileRandom`, or `volatileTtl`.
     */
    public readonly evictionPolicy!: pulumi.Output<string | undefined>;
    /**
     * Database cluster's hostname.
     */
    public /*out*/ readonly host!: pulumi.Output<string>;
    /**
     * Defines when the automatic maintenance should be performed for the database cluster.
     */
    public readonly maintenanceWindows!: pulumi.Output<outputs.DatabaseClusterMaintenanceWindow[] | undefined>;
    /**
     * The name of the database cluster.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Number of nodes that will be included in the cluster.
     */
    public readonly nodeCount!: pulumi.Output<number>;
    /**
     * Password for the cluster's default user.
     */
    public /*out*/ readonly password!: pulumi.Output<string>;
    /**
     * Network port that the database cluster is listening on.
     */
    public /*out*/ readonly port!: pulumi.Output<number>;
    /**
     * Same as `host`, but only accessible from resources within the account and in the same region.
     */
    public /*out*/ readonly privateHost!: pulumi.Output<string>;
    /**
     * The ID of the VPC where the database cluster will be located.
     */
    public readonly privateNetworkUuid!: pulumi.Output<string>;
    /**
     * Same as `uri`, but only accessible from resources within the account and in the same region.
     */
    public /*out*/ readonly privateUri!: pulumi.Output<string>;
    /**
     * The ID of the project that the database cluster is assigned to. If excluded when creating a new database cluster, it will be assigned to your default project.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * DigitalOcean region where the cluster will reside.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Database Droplet size associated with the cluster (ex. `db-s-1vcpu-1gb`). See here for a [list of valid size slugs](https://docs.digitalocean.com/reference/api/api-reference/#tag/Databases).
     */
    public readonly size!: pulumi.Output<string>;
    /**
     * A comma separated string specifying the  SQL modes for a MySQL cluster.
     */
    public readonly sqlMode!: pulumi.Output<string | undefined>;
    /**
     * A list of tag names to be applied to the database cluster.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The full URI for connecting to the database cluster.
     */
    public /*out*/ readonly uri!: pulumi.Output<string>;
    /**
     * Username for the cluster's default user.
     */
    public /*out*/ readonly user!: pulumi.Output<string>;
    /**
     * Engine version used by the cluster (ex. `14` for PostgreSQL 14).
     * When this value is changed, a call to the [Upgrade major Version for a Database](https://docs.digitalocean.com/reference/api/api-reference/#operation/databases_update_major_version) API operation is made with the new version.
     */
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a DatabaseCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseClusterArgs | DatabaseClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatabaseClusterState | undefined;
            resourceInputs["backupRestore"] = state ? state.backupRestore : undefined;
            resourceInputs["clusterUrn"] = state ? state.clusterUrn : undefined;
            resourceInputs["database"] = state ? state.database : undefined;
            resourceInputs["engine"] = state ? state.engine : undefined;
            resourceInputs["evictionPolicy"] = state ? state.evictionPolicy : undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["maintenanceWindows"] = state ? state.maintenanceWindows : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodeCount"] = state ? state.nodeCount : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["privateHost"] = state ? state.privateHost : undefined;
            resourceInputs["privateNetworkUuid"] = state ? state.privateNetworkUuid : undefined;
            resourceInputs["privateUri"] = state ? state.privateUri : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["sqlMode"] = state ? state.sqlMode : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["uri"] = state ? state.uri : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as DatabaseClusterArgs | undefined;
            if ((!args || args.engine === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engine'");
            }
            if ((!args || args.nodeCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeCount'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            if ((!args || args.size === undefined) && !opts.urn) {
                throw new Error("Missing required property 'size'");
            }
            resourceInputs["backupRestore"] = args ? args.backupRestore : undefined;
            resourceInputs["engine"] = args ? args.engine : undefined;
            resourceInputs["evictionPolicy"] = args ? args.evictionPolicy : undefined;
            resourceInputs["maintenanceWindows"] = args ? args.maintenanceWindows : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodeCount"] = args ? args.nodeCount : undefined;
            resourceInputs["privateNetworkUuid"] = args ? args.privateNetworkUuid : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["sqlMode"] = args ? args.sqlMode : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["clusterUrn"] = undefined /*out*/;
            resourceInputs["database"] = undefined /*out*/;
            resourceInputs["host"] = undefined /*out*/;
            resourceInputs["password"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["privateHost"] = undefined /*out*/;
            resourceInputs["privateUri"] = undefined /*out*/;
            resourceInputs["uri"] = undefined /*out*/;
            resourceInputs["user"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password", "privateUri", "uri"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(DatabaseCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatabaseCluster resources.
 */
export interface DatabaseClusterState {
    /**
     * Create a new database cluster based on a backup of an existing cluster.
     */
    backupRestore?: pulumi.Input<inputs.DatabaseClusterBackupRestore>;
    /**
     * The uniform resource name of the database cluster.
     */
    clusterUrn?: pulumi.Input<string>;
    /**
     * Name of the cluster's default database.
     */
    database?: pulumi.Input<string>;
    /**
     * Database engine used by the cluster (ex. `pg` for PostreSQL, `mysql` for MySQL, `redis` for Redis, or `mongodb` for MongoDB).
     */
    engine?: pulumi.Input<string>;
    /**
     * A string specifying the eviction policy for a Redis cluster. Valid values are: `noeviction`, `allkeysLru`, `allkeysRandom`, `volatileLru`, `volatileRandom`, or `volatileTtl`.
     */
    evictionPolicy?: pulumi.Input<string>;
    /**
     * Database cluster's hostname.
     */
    host?: pulumi.Input<string>;
    /**
     * Defines when the automatic maintenance should be performed for the database cluster.
     */
    maintenanceWindows?: pulumi.Input<pulumi.Input<inputs.DatabaseClusterMaintenanceWindow>[]>;
    /**
     * The name of the database cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * Number of nodes that will be included in the cluster.
     */
    nodeCount?: pulumi.Input<number>;
    /**
     * Password for the cluster's default user.
     */
    password?: pulumi.Input<string>;
    /**
     * Network port that the database cluster is listening on.
     */
    port?: pulumi.Input<number>;
    /**
     * Same as `host`, but only accessible from resources within the account and in the same region.
     */
    privateHost?: pulumi.Input<string>;
    /**
     * The ID of the VPC where the database cluster will be located.
     */
    privateNetworkUuid?: pulumi.Input<string>;
    /**
     * Same as `uri`, but only accessible from resources within the account and in the same region.
     */
    privateUri?: pulumi.Input<string>;
    /**
     * The ID of the project that the database cluster is assigned to. If excluded when creating a new database cluster, it will be assigned to your default project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * DigitalOcean region where the cluster will reside.
     */
    region?: pulumi.Input<string | enums.Region>;
    /**
     * Database Droplet size associated with the cluster (ex. `db-s-1vcpu-1gb`). See here for a [list of valid size slugs](https://docs.digitalocean.com/reference/api/api-reference/#tag/Databases).
     */
    size?: pulumi.Input<string | enums.DatabaseSlug>;
    /**
     * A comma separated string specifying the  SQL modes for a MySQL cluster.
     */
    sqlMode?: pulumi.Input<string>;
    /**
     * A list of tag names to be applied to the database cluster.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The full URI for connecting to the database cluster.
     */
    uri?: pulumi.Input<string>;
    /**
     * Username for the cluster's default user.
     */
    user?: pulumi.Input<string>;
    /**
     * Engine version used by the cluster (ex. `14` for PostgreSQL 14).
     * When this value is changed, a call to the [Upgrade major Version for a Database](https://docs.digitalocean.com/reference/api/api-reference/#operation/databases_update_major_version) API operation is made with the new version.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatabaseCluster resource.
 */
export interface DatabaseClusterArgs {
    /**
     * Create a new database cluster based on a backup of an existing cluster.
     */
    backupRestore?: pulumi.Input<inputs.DatabaseClusterBackupRestore>;
    /**
     * Database engine used by the cluster (ex. `pg` for PostreSQL, `mysql` for MySQL, `redis` for Redis, or `mongodb` for MongoDB).
     */
    engine: pulumi.Input<string>;
    /**
     * A string specifying the eviction policy for a Redis cluster. Valid values are: `noeviction`, `allkeysLru`, `allkeysRandom`, `volatileLru`, `volatileRandom`, or `volatileTtl`.
     */
    evictionPolicy?: pulumi.Input<string>;
    /**
     * Defines when the automatic maintenance should be performed for the database cluster.
     */
    maintenanceWindows?: pulumi.Input<pulumi.Input<inputs.DatabaseClusterMaintenanceWindow>[]>;
    /**
     * The name of the database cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * Number of nodes that will be included in the cluster.
     */
    nodeCount: pulumi.Input<number>;
    /**
     * The ID of the VPC where the database cluster will be located.
     */
    privateNetworkUuid?: pulumi.Input<string>;
    /**
     * The ID of the project that the database cluster is assigned to. If excluded when creating a new database cluster, it will be assigned to your default project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * DigitalOcean region where the cluster will reside.
     */
    region: pulumi.Input<string | enums.Region>;
    /**
     * Database Droplet size associated with the cluster (ex. `db-s-1vcpu-1gb`). See here for a [list of valid size slugs](https://docs.digitalocean.com/reference/api/api-reference/#tag/Databases).
     */
    size: pulumi.Input<string | enums.DatabaseSlug>;
    /**
     * A comma separated string specifying the  SQL modes for a MySQL cluster.
     */
    sqlMode?: pulumi.Input<string>;
    /**
     * A list of tag names to be applied to the database cluster.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Engine version used by the cluster (ex. `14` for PostgreSQL 14).
     * When this value is changed, a call to the [Upgrade major Version for a Database](https://docs.digitalocean.com/reference/api/api-reference/#operation/databases_update_major_version) API operation is made with the new version.
     */
    version?: pulumi.Input<string>;
}
