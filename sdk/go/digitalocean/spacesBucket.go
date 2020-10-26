// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a bucket resource for Spaces, DigitalOcean's object storage product.
//
// The [Spaces API](https://developers.digitalocean.com/documentation/spaces/) was
// designed to be interoperable with Amazon's AWS S3 API. This allows users to
// interact with the service while using the tools they already know. Spaces
// mirrors S3's authentication framework and requests to Spaces require a key pair
// similar to Amazon's Access ID and Secret Key.
//
// The authentication requirement can be met by either setting the
// `SPACES_ACCESS_KEY_ID` and `SPACES_SECRET_ACCESS_KEY` environment variables or
// the provider's `spacesAccessId` and `spacesSecretKey` arguments to the
// access ID and secret you generate via the DigitalOcean control panel. For
// example:
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
// 		_, err := digitalocean.NewSpacesBucket(ctx, "static_assets", nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// For more information, See [An Introduction to DigitalOcean Spaces](https://www.digitalocean.com/community/tutorials/an-introduction-to-digitalocean-spaces)
//
// ## Example Usage
// ### Create a New Bucket
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
// 		_, err := digitalocean.NewSpacesBucket(ctx, "foobar", &digitalocean.SpacesBucketArgs{
// 			Region: pulumi.String("nyc3"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Create a New Bucket With CORS Rules
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
// 		_, err := digitalocean.NewSpacesBucket(ctx, "foobar", &digitalocean.SpacesBucketArgs{
// 			CorsRules: digitalocean.SpacesBucketCorsRuleArray{
// 				&digitalocean.SpacesBucketCorsRuleArgs{
// 					AllowedHeaders: pulumi.StringArray{
// 						pulumi.String("*"),
// 					},
// 					AllowedMethods: pulumi.StringArray{
// 						pulumi.String("GET"),
// 					},
// 					AllowedOrigins: pulumi.StringArray{
// 						pulumi.String("*"),
// 					},
// 					MaxAgeSeconds: pulumi.Int(3000),
// 				},
// 				&digitalocean.SpacesBucketCorsRuleArgs{
// 					AllowedHeaders: pulumi.StringArray{
// 						pulumi.String("*"),
// 					},
// 					AllowedMethods: pulumi.StringArray{
// 						pulumi.String("PUT"),
// 						pulumi.String("POST"),
// 						pulumi.String("DELETE"),
// 					},
// 					AllowedOrigins: pulumi.StringArray{
// 						pulumi.String("https://www.example.com"),
// 					},
// 					MaxAgeSeconds: pulumi.Int(3000),
// 				},
// 			},
// 			Region: pulumi.String("nyc3"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SpacesBucket struct {
	pulumi.CustomResourceState

	// Canned ACL applied on bucket creation (`private` or `public-read`)
	Acl pulumi.StringPtrOutput `pulumi:"acl"`
	// The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com)
	BucketDomainName pulumi.StringOutput `pulumi:"bucketDomainName"`
	// The uniform resource name for the bucket
	BucketUrn pulumi.StringOutput `pulumi:"bucketUrn"`
	// A rule of Cross-Origin Resource Sharing (documented below).
	CorsRules SpacesBucketCorsRuleArrayOutput `pulumi:"corsRules"`
	// Unless `true`, the bucket will only be destroyed if empty (Defaults to `false`)
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// A configuration of object lifecycle management (documented below).
	LifecycleRules SpacesBucketLifecycleRuleArrayOutput `pulumi:"lifecycleRules"`
	// The name of the bucket
	Name pulumi.StringOutput `pulumi:"name"`
	// The region where the bucket resides (Defaults to `nyc3`)
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// A state of versioning (documented below)
	Versioning SpacesBucketVersioningPtrOutput `pulumi:"versioning"`
}

// NewSpacesBucket registers a new resource with the given unique name, arguments, and options.
func NewSpacesBucket(ctx *pulumi.Context,
	name string, args *SpacesBucketArgs, opts ...pulumi.ResourceOption) (*SpacesBucket, error) {
	if args == nil {
		args = &SpacesBucketArgs{}
	}
	var resource SpacesBucket
	err := ctx.RegisterResource("digitalocean:index/spacesBucket:SpacesBucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpacesBucket gets an existing SpacesBucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpacesBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpacesBucketState, opts ...pulumi.ResourceOption) (*SpacesBucket, error) {
	var resource SpacesBucket
	err := ctx.ReadResource("digitalocean:index/spacesBucket:SpacesBucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpacesBucket resources.
type spacesBucketState struct {
	// Canned ACL applied on bucket creation (`private` or `public-read`)
	Acl *string `pulumi:"acl"`
	// The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com)
	BucketDomainName *string `pulumi:"bucketDomainName"`
	// The uniform resource name for the bucket
	BucketUrn *string `pulumi:"bucketUrn"`
	// A rule of Cross-Origin Resource Sharing (documented below).
	CorsRules []SpacesBucketCorsRule `pulumi:"corsRules"`
	// Unless `true`, the bucket will only be destroyed if empty (Defaults to `false`)
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// A configuration of object lifecycle management (documented below).
	LifecycleRules []SpacesBucketLifecycleRule `pulumi:"lifecycleRules"`
	// The name of the bucket
	Name *string `pulumi:"name"`
	// The region where the bucket resides (Defaults to `nyc3`)
	Region *string `pulumi:"region"`
	// A state of versioning (documented below)
	Versioning *SpacesBucketVersioning `pulumi:"versioning"`
}

type SpacesBucketState struct {
	// Canned ACL applied on bucket creation (`private` or `public-read`)
	Acl pulumi.StringPtrInput
	// The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com)
	BucketDomainName pulumi.StringPtrInput
	// The uniform resource name for the bucket
	BucketUrn pulumi.StringPtrInput
	// A rule of Cross-Origin Resource Sharing (documented below).
	CorsRules SpacesBucketCorsRuleArrayInput
	// Unless `true`, the bucket will only be destroyed if empty (Defaults to `false`)
	ForceDestroy pulumi.BoolPtrInput
	// A configuration of object lifecycle management (documented below).
	LifecycleRules SpacesBucketLifecycleRuleArrayInput
	// The name of the bucket
	Name pulumi.StringPtrInput
	// The region where the bucket resides (Defaults to `nyc3`)
	Region pulumi.StringPtrInput
	// A state of versioning (documented below)
	Versioning SpacesBucketVersioningPtrInput
}

func (SpacesBucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*spacesBucketState)(nil)).Elem()
}

type spacesBucketArgs struct {
	// Canned ACL applied on bucket creation (`private` or `public-read`)
	Acl *string `pulumi:"acl"`
	// A rule of Cross-Origin Resource Sharing (documented below).
	CorsRules []SpacesBucketCorsRule `pulumi:"corsRules"`
	// Unless `true`, the bucket will only be destroyed if empty (Defaults to `false`)
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// A configuration of object lifecycle management (documented below).
	LifecycleRules []SpacesBucketLifecycleRule `pulumi:"lifecycleRules"`
	// The name of the bucket
	Name *string `pulumi:"name"`
	// The region where the bucket resides (Defaults to `nyc3`)
	Region *string `pulumi:"region"`
	// A state of versioning (documented below)
	Versioning *SpacesBucketVersioning `pulumi:"versioning"`
}

// The set of arguments for constructing a SpacesBucket resource.
type SpacesBucketArgs struct {
	// Canned ACL applied on bucket creation (`private` or `public-read`)
	Acl pulumi.StringPtrInput
	// A rule of Cross-Origin Resource Sharing (documented below).
	CorsRules SpacesBucketCorsRuleArrayInput
	// Unless `true`, the bucket will only be destroyed if empty (Defaults to `false`)
	ForceDestroy pulumi.BoolPtrInput
	// A configuration of object lifecycle management (documented below).
	LifecycleRules SpacesBucketLifecycleRuleArrayInput
	// The name of the bucket
	Name pulumi.StringPtrInput
	// The region where the bucket resides (Defaults to `nyc3`)
	Region pulumi.StringPtrInput
	// A state of versioning (documented below)
	Versioning SpacesBucketVersioningPtrInput
}

func (SpacesBucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spacesBucketArgs)(nil)).Elem()
}
