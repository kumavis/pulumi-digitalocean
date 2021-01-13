// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a DigitalOcean database resource. When creating a new database cluster, a default database with name `defaultdb` will be created. Then, this resource can be used to provide additional database inside the cluster.
//
// ## Example Usage
// ### Create a new PostgreSQL database
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
// 		_, err = digitalocean.NewDatabaseDb(ctx, "database_example", &digitalocean.DatabaseDbArgs{
// 			ClusterId: postgres_example.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Database can be imported using the `id` of the source database cluster and the `name` of the database joined with a comma. For example
//
// ```sh
//  $ pulumi import digitalocean:index:DatabaseDb database-example 245bcfd0-7f31-4ce6-a2bc-475a116cca97,foobar
// ```
type DatabaseDb struct {
	pulumi.CustomResourceState

	// The ID of the original source database cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The name for the database.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewDatabaseDb registers a new resource with the given unique name, arguments, and options.
func NewDatabaseDb(ctx *pulumi.Context,
	name string, args *DatabaseDbArgs, opts ...pulumi.ResourceOption) (*DatabaseDb, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	var resource DatabaseDb
	err := ctx.RegisterResource("digitalocean:index:DatabaseDb", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseDb gets an existing DatabaseDb resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseDb(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseDbState, opts ...pulumi.ResourceOption) (*DatabaseDb, error) {
	var resource DatabaseDb
	err := ctx.ReadResource("digitalocean:index:DatabaseDb", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseDb resources.
type databaseDbState struct {
	// The ID of the original source database cluster.
	ClusterId *string `pulumi:"clusterId"`
	// The name for the database.
	Name *string `pulumi:"name"`
}

type DatabaseDbState struct {
	// The ID of the original source database cluster.
	ClusterId pulumi.StringPtrInput
	// The name for the database.
	Name pulumi.StringPtrInput
}

func (DatabaseDbState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseDbState)(nil)).Elem()
}

type databaseDbArgs struct {
	// The ID of the original source database cluster.
	ClusterId string `pulumi:"clusterId"`
	// The name for the database.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a DatabaseDb resource.
type DatabaseDbArgs struct {
	// The ID of the original source database cluster.
	ClusterId pulumi.StringInput
	// The name for the database.
	Name pulumi.StringPtrInput
}

func (DatabaseDbArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseDbArgs)(nil)).Elem()
}

type DatabaseDbInput interface {
	pulumi.Input

	ToDatabaseDbOutput() DatabaseDbOutput
	ToDatabaseDbOutputWithContext(ctx context.Context) DatabaseDbOutput
}

func (DatabaseDb) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseDb)(nil)).Elem()
}

func (i DatabaseDb) ToDatabaseDbOutput() DatabaseDbOutput {
	return i.ToDatabaseDbOutputWithContext(context.Background())
}

func (i DatabaseDb) ToDatabaseDbOutputWithContext(ctx context.Context) DatabaseDbOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseDbOutput)
}

type DatabaseDbOutput struct {
	*pulumi.OutputState
}

func (DatabaseDbOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseDbOutput)(nil)).Elem()
}

func (o DatabaseDbOutput) ToDatabaseDbOutput() DatabaseDbOutput {
	return o
}

func (o DatabaseDbOutput) ToDatabaseDbOutputWithContext(ctx context.Context) DatabaseDbOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatabaseDbOutput{})
}
