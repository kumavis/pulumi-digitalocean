// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a DigitalOcean database user resource. When creating a new database cluster, a default admin user with name `doadmin` will be created. Then, this resource can be used to provide additional normal users inside the cluster.
//
// > **NOTE:** Any new users created will always have `normal` role, only the default user that comes with database cluster creation has `primary` role. Additional permissions must be managed manually.
//
// ## Example Usage
// ### Create a new PostgreSQL database user
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v3/go/digitalocean"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := digitalocean.NewDatabaseCluster(ctx, "postgres_example", &digitalocean.DatabaseClusterArgs{
// 			Engine:    pulumi.String("pg"),
// 			Version:   pulumi.String("11"),
// 			Size:      pulumi.String("db-s-1vcpu-1gb"),
// 			Region:    pulumi.String("nyc1"),
// 			NodeCount: pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = digitalocean.NewDatabaseUser(ctx, "user_example", &digitalocean.DatabaseUserArgs{
// 			ClusterId: postgres_example.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DatabaseUser struct {
	pulumi.CustomResourceState

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
}

// NewDatabaseUser registers a new resource with the given unique name, arguments, and options.
func NewDatabaseUser(ctx *pulumi.Context,
	name string, args *DatabaseUserArgs, opts ...pulumi.ResourceOption) (*DatabaseUser, error) {
	if args == nil || args.ClusterId == nil {
		return nil, errors.New("missing required argument 'ClusterId'")
	}
	if args == nil {
		args = &DatabaseUserArgs{}
	}
	var resource DatabaseUser
	err := ctx.RegisterResource("digitalocean:index:DatabaseUser", name, args, &resource, opts...)
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
	err := ctx.ReadResource("digitalocean:index:DatabaseUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseUser resources.
type databaseUserState struct {
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
}

type DatabaseUserState struct {
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
}

// The set of arguments for constructing a DatabaseUser resource.
type DatabaseUserArgs struct {
	// The ID of the original source database cluster.
	ClusterId pulumi.StringInput
	// The authentication method to use for connections to the MySQL user account. The valid values are `mysqlNativePassword` or `cachingSha2Password` (this is the default).
	MysqlAuthPlugin pulumi.StringPtrInput
	// The name for the database user.
	Name pulumi.StringPtrInput
}

func (DatabaseUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseUserArgs)(nil)).Elem()
}
