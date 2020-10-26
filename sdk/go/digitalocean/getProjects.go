// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Retrieve information about all DigitalOcean projects associated with an account, with
// the ability to filter and sort the results. If no filters are specified, all projects
// will be returned.
//
// Note: You can use the `Project` data source to
// obtain metadata about a single project if you already know the `id` to retrieve or the unique
// `name` of the project.
//
// ## Example Usage
//
// Use the `filter` block with a `key` string and `values` list to filter projects.
//
// For example to find all staging environment projects:
//
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
// 		_, err := digitalocean.GetProjects(ctx, &digitalocean.GetProjectsArgs{
// 			Filters: []digitalocean.GetProjectsFilter{
// 				digitalocean.GetProjectsFilter{
// 					Key: "environment",
// 					Values: []string{
// 						"Staging",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// You can filter on multiple fields and sort the results as well:
//
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
// 		_, err := digitalocean.GetProjects(ctx, &digitalocean.GetProjectsArgs{
// 			Filters: []digitalocean.GetProjectsFilter{
// 				digitalocean.GetProjectsFilter{
// 					Key: "environment",
// 					Values: []string{
// 						"Production",
// 					},
// 				},
// 				digitalocean.GetProjectsFilter{
// 					Key: "is_default",
// 					Values: []string{
// 						"false",
// 					},
// 				},
// 			},
// 			Sorts: []digitalocean.GetProjectsSort{
// 				digitalocean.GetProjectsSort{
// 					Direction: "asc",
// 					Key:       "name",
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetProjects(ctx *pulumi.Context, args *GetProjectsArgs, opts ...pulumi.InvokeOption) (*GetProjectsResult, error) {
	var rv GetProjectsResult
	err := ctx.Invoke("digitalocean:index/getProjects:getProjects", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjects.
type GetProjectsArgs struct {
	// Filter the results.
	// The `filter` block is documented below.
	Filters []GetProjectsFilter `pulumi:"filters"`
	// Sort the results.
	// The `sort` block is documented below.
	Sorts []GetProjectsSort `pulumi:"sorts"`
}

// A collection of values returned by getProjects.
type GetProjectsResult struct {
	Filters []GetProjectsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A set of projects satisfying any `filter` and `sort` criteria. Each project has
	// the following attributes:
	Projects []GetProjectsProject `pulumi:"projects"`
	Sorts    []GetProjectsSort    `pulumi:"sorts"`
}
