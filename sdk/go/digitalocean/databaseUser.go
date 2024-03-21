// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DigitalOcean database user resource. When creating a new database cluster, a default admin user with name `doadmin` will be created. Then, this resource can be used to provide additional normal users inside the cluster.
//
// > **NOTE:** Any new users created will always have `normal` role, only the default user that comes with database cluster creation has `primary` role. Additional permissions must be managed manually.
//
// ## Example Usage
//
// ### Create a new PostgreSQL database user
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := digitalocean.NewDatabaseCluster(ctx, "postgres-example", &digitalocean.DatabaseClusterArgs{
//				Engine:    pulumi.String("pg"),
//				Version:   pulumi.String("11"),
//				Size:      pulumi.String(digitalocean.DatabaseSlug_DB_1VPCU1GB),
//				Region:    pulumi.String(digitalocean.RegionNYC1),
//				NodeCount: pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = digitalocean.NewDatabaseUser(ctx, "user-example", &digitalocean.DatabaseUserArgs{
//				ClusterId: postgres_example.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### Create a new user for a PostgreSQL database replica
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := digitalocean.NewDatabaseCluster(ctx, "postgres-example", &digitalocean.DatabaseClusterArgs{
//				Engine:    pulumi.String("pg"),
//				Version:   pulumi.String("11"),
//				Size:      pulumi.String(digitalocean.DatabaseSlug_DB_1VPCU1GB),
//				Region:    pulumi.String(digitalocean.RegionNYC1),
//				NodeCount: pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = digitalocean.NewDatabaseReplica(ctx, "replica-example", &digitalocean.DatabaseReplicaArgs{
//				ClusterId: postgres_example.ID(),
//				Size:      pulumi.String(digitalocean.DatabaseSlug_DB_1VPCU1GB),
//				Region:    pulumi.String(digitalocean.RegionNYC1),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = digitalocean.NewDatabaseUser(ctx, "user-example", &digitalocean.DatabaseUserArgs{
//				ClusterId: replica_example.Uuid,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### Create a new user for a Kafka database cluster
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := digitalocean.NewDatabaseCluster(ctx, "kafka-example", &digitalocean.DatabaseClusterArgs{
//				Engine:    pulumi.String("kafka"),
//				Version:   pulumi.String("3.5"),
//				Size:      pulumi.String("db-s-2vcpu-2gb"),
//				Region:    pulumi.String(digitalocean.RegionNYC1),
//				NodeCount: pulumi.Int(3),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = digitalocean.NewDatabaseKafkaTopic(ctx, "foobarTopic", &digitalocean.DatabaseKafkaTopicArgs{
//				ClusterId: pulumi.Any(digitalocean_database_cluster.Foobar.Id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = digitalocean.NewDatabaseUser(ctx, "foobarUser", &digitalocean.DatabaseUserArgs{
//				ClusterId: pulumi.Any(digitalocean_database_cluster.Foobar.Id),
//				Settings: digitalocean.DatabaseUserSettingArray{
//					&digitalocean.DatabaseUserSettingArgs{
//						Acls: digitalocean.DatabaseUserSettingAclArray{
//							&digitalocean.DatabaseUserSettingAclArgs{
//								Topic:      pulumi.String("topic-1"),
//								Permission: pulumi.String("produce"),
//							},
//							&digitalocean.DatabaseUserSettingAclArgs{
//								Topic:      pulumi.String("topic-2"),
//								Permission: pulumi.String("produceconsume"),
//							},
//							&digitalocean.DatabaseUserSettingAclArgs{
//								Topic:      pulumi.String("topic-*"),
//								Permission: pulumi.String("consume"),
//							},
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// # Database user can be imported using the `id` of the source database cluster
//
// and the `name` of the user joined with a comma. For example:
//
// ```sh
// $ pulumi import digitalocean:index/databaseUser:DatabaseUser user-example 245bcfd0-7f31-4ce6-a2bc-475a116cca97,foobar
// ```
type DatabaseUser struct {
	pulumi.CustomResourceState

	// Access certificate for TLS client authentication. (Kafka only)
	AccessCert pulumi.StringOutput `pulumi:"accessCert"`
	// Access key for TLS client authentication. (Kafka only)
	AccessKey pulumi.StringOutput `pulumi:"accessKey"`
	// The ID of the original source database cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The authentication method to use for connections to the MySQL user account. The valid values are `mysqlNativePassword` or `cachingSha2Password` (this is the default).
	MysqlAuthPlugin pulumi.StringPtrOutput `pulumi:"mysqlAuthPlugin"`
	// The name for the database user.
	Name pulumi.StringOutput `pulumi:"name"`
	// Password for the database user.
	Password pulumi.StringOutput `pulumi:"password"`
	// Role for the database user. The value will be either "primary" or "normal".
	Role pulumi.StringOutput `pulumi:"role"`
	// Contains optional settings for the user.
	// The `settings` block is documented below.
	Settings DatabaseUserSettingArrayOutput `pulumi:"settings"`
}

// NewDatabaseUser registers a new resource with the given unique name, arguments, and options.
func NewDatabaseUser(ctx *pulumi.Context,
	name string, args *DatabaseUserArgs, opts ...pulumi.ResourceOption) (*DatabaseUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accessCert",
		"accessKey",
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DatabaseUser
	err := ctx.RegisterResource("digitalocean:index/databaseUser:DatabaseUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseUser gets an existing DatabaseUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseUserState, opts ...pulumi.ResourceOption) (*DatabaseUser, error) {
	var resource DatabaseUser
	err := ctx.ReadResource("digitalocean:index/databaseUser:DatabaseUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseUser resources.
type databaseUserState struct {
	// Access certificate for TLS client authentication. (Kafka only)
	AccessCert *string `pulumi:"accessCert"`
	// Access key for TLS client authentication. (Kafka only)
	AccessKey *string `pulumi:"accessKey"`
	// The ID of the original source database cluster.
	ClusterId *string `pulumi:"clusterId"`
	// The authentication method to use for connections to the MySQL user account. The valid values are `mysqlNativePassword` or `cachingSha2Password` (this is the default).
	MysqlAuthPlugin *string `pulumi:"mysqlAuthPlugin"`
	// The name for the database user.
	Name *string `pulumi:"name"`
	// Password for the database user.
	Password *string `pulumi:"password"`
	// Role for the database user. The value will be either "primary" or "normal".
	Role *string `pulumi:"role"`
	// Contains optional settings for the user.
	// The `settings` block is documented below.
	Settings []DatabaseUserSetting `pulumi:"settings"`
}

type DatabaseUserState struct {
	// Access certificate for TLS client authentication. (Kafka only)
	AccessCert pulumi.StringPtrInput
	// Access key for TLS client authentication. (Kafka only)
	AccessKey pulumi.StringPtrInput
	// The ID of the original source database cluster.
	ClusterId pulumi.StringPtrInput
	// The authentication method to use for connections to the MySQL user account. The valid values are `mysqlNativePassword` or `cachingSha2Password` (this is the default).
	MysqlAuthPlugin pulumi.StringPtrInput
	// The name for the database user.
	Name pulumi.StringPtrInput
	// Password for the database user.
	Password pulumi.StringPtrInput
	// Role for the database user. The value will be either "primary" or "normal".
	Role pulumi.StringPtrInput
	// Contains optional settings for the user.
	// The `settings` block is documented below.
	Settings DatabaseUserSettingArrayInput
}

func (DatabaseUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseUserState)(nil)).Elem()
}

type databaseUserArgs struct {
	// The ID of the original source database cluster.
	ClusterId string `pulumi:"clusterId"`
	// The authentication method to use for connections to the MySQL user account. The valid values are `mysqlNativePassword` or `cachingSha2Password` (this is the default).
	MysqlAuthPlugin *string `pulumi:"mysqlAuthPlugin"`
	// The name for the database user.
	Name *string `pulumi:"name"`
	// Contains optional settings for the user.
	// The `settings` block is documented below.
	Settings []DatabaseUserSetting `pulumi:"settings"`
}

// The set of arguments for constructing a DatabaseUser resource.
type DatabaseUserArgs struct {
	// The ID of the original source database cluster.
	ClusterId pulumi.StringInput
	// The authentication method to use for connections to the MySQL user account. The valid values are `mysqlNativePassword` or `cachingSha2Password` (this is the default).
	MysqlAuthPlugin pulumi.StringPtrInput
	// The name for the database user.
	Name pulumi.StringPtrInput
	// Contains optional settings for the user.
	// The `settings` block is documented below.
	Settings DatabaseUserSettingArrayInput
}

func (DatabaseUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseUserArgs)(nil)).Elem()
}

type DatabaseUserInput interface {
	pulumi.Input

	ToDatabaseUserOutput() DatabaseUserOutput
	ToDatabaseUserOutputWithContext(ctx context.Context) DatabaseUserOutput
}

func (*DatabaseUser) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseUser)(nil)).Elem()
}

func (i *DatabaseUser) ToDatabaseUserOutput() DatabaseUserOutput {
	return i.ToDatabaseUserOutputWithContext(context.Background())
}

func (i *DatabaseUser) ToDatabaseUserOutputWithContext(ctx context.Context) DatabaseUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseUserOutput)
}

// DatabaseUserArrayInput is an input type that accepts DatabaseUserArray and DatabaseUserArrayOutput values.
// You can construct a concrete instance of `DatabaseUserArrayInput` via:
//
//	DatabaseUserArray{ DatabaseUserArgs{...} }
type DatabaseUserArrayInput interface {
	pulumi.Input

	ToDatabaseUserArrayOutput() DatabaseUserArrayOutput
	ToDatabaseUserArrayOutputWithContext(context.Context) DatabaseUserArrayOutput
}

type DatabaseUserArray []DatabaseUserInput

func (DatabaseUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseUser)(nil)).Elem()
}

func (i DatabaseUserArray) ToDatabaseUserArrayOutput() DatabaseUserArrayOutput {
	return i.ToDatabaseUserArrayOutputWithContext(context.Background())
}

func (i DatabaseUserArray) ToDatabaseUserArrayOutputWithContext(ctx context.Context) DatabaseUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseUserArrayOutput)
}

// DatabaseUserMapInput is an input type that accepts DatabaseUserMap and DatabaseUserMapOutput values.
// You can construct a concrete instance of `DatabaseUserMapInput` via:
//
//	DatabaseUserMap{ "key": DatabaseUserArgs{...} }
type DatabaseUserMapInput interface {
	pulumi.Input

	ToDatabaseUserMapOutput() DatabaseUserMapOutput
	ToDatabaseUserMapOutputWithContext(context.Context) DatabaseUserMapOutput
}

type DatabaseUserMap map[string]DatabaseUserInput

func (DatabaseUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseUser)(nil)).Elem()
}

func (i DatabaseUserMap) ToDatabaseUserMapOutput() DatabaseUserMapOutput {
	return i.ToDatabaseUserMapOutputWithContext(context.Background())
}

func (i DatabaseUserMap) ToDatabaseUserMapOutputWithContext(ctx context.Context) DatabaseUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseUserMapOutput)
}

type DatabaseUserOutput struct{ *pulumi.OutputState }

func (DatabaseUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseUser)(nil)).Elem()
}

func (o DatabaseUserOutput) ToDatabaseUserOutput() DatabaseUserOutput {
	return o
}

func (o DatabaseUserOutput) ToDatabaseUserOutputWithContext(ctx context.Context) DatabaseUserOutput {
	return o
}

// Access certificate for TLS client authentication. (Kafka only)
func (o DatabaseUserOutput) AccessCert() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseUser) pulumi.StringOutput { return v.AccessCert }).(pulumi.StringOutput)
}

// Access key for TLS client authentication. (Kafka only)
func (o DatabaseUserOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseUser) pulumi.StringOutput { return v.AccessKey }).(pulumi.StringOutput)
}

// The ID of the original source database cluster.
func (o DatabaseUserOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseUser) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The authentication method to use for connections to the MySQL user account. The valid values are `mysqlNativePassword` or `cachingSha2Password` (this is the default).
func (o DatabaseUserOutput) MysqlAuthPlugin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseUser) pulumi.StringPtrOutput { return v.MysqlAuthPlugin }).(pulumi.StringPtrOutput)
}

// The name for the database user.
func (o DatabaseUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Password for the database user.
func (o DatabaseUserOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseUser) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// Role for the database user. The value will be either "primary" or "normal".
func (o DatabaseUserOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseUser) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// Contains optional settings for the user.
// The `settings` block is documented below.
func (o DatabaseUserOutput) Settings() DatabaseUserSettingArrayOutput {
	return o.ApplyT(func(v *DatabaseUser) DatabaseUserSettingArrayOutput { return v.Settings }).(DatabaseUserSettingArrayOutput)
}

type DatabaseUserArrayOutput struct{ *pulumi.OutputState }

func (DatabaseUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseUser)(nil)).Elem()
}

func (o DatabaseUserArrayOutput) ToDatabaseUserArrayOutput() DatabaseUserArrayOutput {
	return o
}

func (o DatabaseUserArrayOutput) ToDatabaseUserArrayOutputWithContext(ctx context.Context) DatabaseUserArrayOutput {
	return o
}

func (o DatabaseUserArrayOutput) Index(i pulumi.IntInput) DatabaseUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatabaseUser {
		return vs[0].([]*DatabaseUser)[vs[1].(int)]
	}).(DatabaseUserOutput)
}

type DatabaseUserMapOutput struct{ *pulumi.OutputState }

func (DatabaseUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseUser)(nil)).Elem()
}

func (o DatabaseUserMapOutput) ToDatabaseUserMapOutput() DatabaseUserMapOutput {
	return o
}

func (o DatabaseUserMapOutput) ToDatabaseUserMapOutputWithContext(ctx context.Context) DatabaseUserMapOutput {
	return o
}

func (o DatabaseUserMapOutput) MapIndex(k pulumi.StringInput) DatabaseUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatabaseUser {
		return vs[0].(map[string]*DatabaseUser)[vs[1].(string)]
	}).(DatabaseUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseUserInput)(nil)).Elem(), &DatabaseUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseUserArrayInput)(nil)).Elem(), DatabaseUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseUserMapInput)(nil)).Elem(), DatabaseUserMap{})
	pulumi.RegisterOutputType(DatabaseUserOutput{})
	pulumi.RegisterOutputType(DatabaseUserArrayOutput{})
	pulumi.RegisterOutputType(DatabaseUserMapOutput{})
}
