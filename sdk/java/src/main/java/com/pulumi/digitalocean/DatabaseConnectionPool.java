// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.digitalocean.DatabaseConnectionPoolArgs;
import com.pulumi.digitalocean.Utilities;
import com.pulumi.digitalocean.inputs.DatabaseConnectionPoolState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a DigitalOcean database connection pool resource.
 * 
 * ## Example Usage
 * 
 * ### Create a new PostgreSQL database connection pool
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.digitalocean.DatabaseCluster;
 * import com.pulumi.digitalocean.DatabaseClusterArgs;
 * import com.pulumi.digitalocean.DatabaseConnectionPool;
 * import com.pulumi.digitalocean.DatabaseConnectionPoolArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var postgres_example = new DatabaseCluster("postgres-example", DatabaseClusterArgs.builder()
 *             .name("example-postgres-cluster")
 *             .engine("pg")
 *             .version("15")
 *             .size("db-s-1vcpu-1gb")
 *             .region("nyc1")
 *             .nodeCount(1)
 *             .build());
 * 
 *         var pool_01 = new DatabaseConnectionPool("pool-01", DatabaseConnectionPoolArgs.builder()
 *             .clusterId(postgres_example.id())
 *             .name("pool-01")
 *             .mode("transaction")
 *             .size(20)
 *             .dbName("defaultdb")
 *             .user("doadmin")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Database connection pools can be imported using the `id` of the source database cluster
 * 
 * and the `name` of the connection pool joined with a comma. For example:
 * 
 * ```sh
 * $ pulumi import digitalocean:index/databaseConnectionPool:DatabaseConnectionPool pool-01 245bcfd0-7f31-4ce6-a2bc-475a116cca97,pool-01
 * ```
 * 
 */
@ResourceType(type="digitalocean:index/databaseConnectionPool:DatabaseConnectionPool")
public class DatabaseConnectionPool extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the source database cluster. Note: This must be a PostgreSQL cluster.
     * 
     */
    @Export(name="clusterId", refs={String.class}, tree="[0]")
    private Output<String> clusterId;

    /**
     * @return The ID of the source database cluster. Note: This must be a PostgreSQL cluster.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    /**
     * The database for use with the connection pool.
     * 
     */
    @Export(name="dbName", refs={String.class}, tree="[0]")
    private Output<String> dbName;

    /**
     * @return The database for use with the connection pool.
     * 
     */
    public Output<String> dbName() {
        return this.dbName;
    }
    /**
     * The hostname used to connect to the database connection pool.
     * 
     */
    @Export(name="host", refs={String.class}, tree="[0]")
    private Output<String> host;

    /**
     * @return The hostname used to connect to the database connection pool.
     * 
     */
    public Output<String> host() {
        return this.host;
    }
    /**
     * The PGBouncer transaction mode for the connection pool. The allowed values are session, transaction, and statement.
     * 
     */
    @Export(name="mode", refs={String.class}, tree="[0]")
    private Output<String> mode;

    /**
     * @return The PGBouncer transaction mode for the connection pool. The allowed values are session, transaction, and statement.
     * 
     */
    public Output<String> mode() {
        return this.mode;
    }
    /**
     * The name for the database connection pool.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name for the database connection pool.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Password for the connection pool&#39;s user.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output<String> password;

    /**
     * @return Password for the connection pool&#39;s user.
     * 
     */
    public Output<String> password() {
        return this.password;
    }
    /**
     * Network port that the database connection pool is listening on.
     * 
     */
    @Export(name="port", refs={Integer.class}, tree="[0]")
    private Output<Integer> port;

    /**
     * @return Network port that the database connection pool is listening on.
     * 
     */
    public Output<Integer> port() {
        return this.port;
    }
    /**
     * Same as `host`, but only accessible from resources within the account and in the same region.
     * 
     */
    @Export(name="privateHost", refs={String.class}, tree="[0]")
    private Output<String> privateHost;

    /**
     * @return Same as `host`, but only accessible from resources within the account and in the same region.
     * 
     */
    public Output<String> privateHost() {
        return this.privateHost;
    }
    /**
     * Same as `uri`, but only accessible from resources within the account and in the same region.
     * 
     */
    @Export(name="privateUri", refs={String.class}, tree="[0]")
    private Output<String> privateUri;

    /**
     * @return Same as `uri`, but only accessible from resources within the account and in the same region.
     * 
     */
    public Output<String> privateUri() {
        return this.privateUri;
    }
    /**
     * The desired size of the PGBouncer connection pool.
     * 
     */
    @Export(name="size", refs={Integer.class}, tree="[0]")
    private Output<Integer> size;

    /**
     * @return The desired size of the PGBouncer connection pool.
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }
    /**
     * The full URI for connecting to the database connection pool.
     * 
     */
    @Export(name="uri", refs={String.class}, tree="[0]")
    private Output<String> uri;

    /**
     * @return The full URI for connecting to the database connection pool.
     * 
     */
    public Output<String> uri() {
        return this.uri;
    }
    /**
     * The name of the database user for use with the connection pool. When excluded, all sessions connect to the database as the inbound user.
     * 
     */
    @Export(name="user", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> user;

    /**
     * @return The name of the database user for use with the connection pool. When excluded, all sessions connect to the database as the inbound user.
     * 
     */
    public Output<Optional<String>> user() {
        return Codegen.optional(this.user);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DatabaseConnectionPool(java.lang.String name) {
        this(name, DatabaseConnectionPoolArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DatabaseConnectionPool(java.lang.String name, DatabaseConnectionPoolArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DatabaseConnectionPool(java.lang.String name, DatabaseConnectionPoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("digitalocean:index/databaseConnectionPool:DatabaseConnectionPool", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private DatabaseConnectionPool(java.lang.String name, Output<java.lang.String> id, @Nullable DatabaseConnectionPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("digitalocean:index/databaseConnectionPool:DatabaseConnectionPool", name, state, makeResourceOptions(options, id), false);
    }

    private static DatabaseConnectionPoolArgs makeArgs(DatabaseConnectionPoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DatabaseConnectionPoolArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password",
                "privateUri",
                "uri"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static DatabaseConnectionPool get(java.lang.String name, Output<java.lang.String> id, @Nullable DatabaseConnectionPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DatabaseConnectionPool(name, id, state, options);
    }
}
