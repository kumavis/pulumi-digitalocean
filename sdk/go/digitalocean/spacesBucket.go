// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
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
// 
// For more information, See [An Introduction to DigitalOcean Spaces](https://www.digitalocean.com/community/tutorials/an-introduction-to-digitalocean-spaces)
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/r/spaces_bucket.html.markdown.
type SpacesBucket struct {
	s *pulumi.ResourceState
}

// NewSpacesBucket registers a new resource with the given unique name, arguments, and options.
func NewSpacesBucket(ctx *pulumi.Context,
	name string, args *SpacesBucketArgs, opts ...pulumi.ResourceOpt) (*SpacesBucket, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["acl"] = nil
		inputs["corsRules"] = nil
		inputs["forceDestroy"] = nil
		inputs["name"] = nil
		inputs["region"] = nil
	} else {
		inputs["acl"] = args.Acl
		inputs["corsRules"] = args.CorsRules
		inputs["forceDestroy"] = args.ForceDestroy
		inputs["name"] = args.Name
		inputs["region"] = args.Region
	}
	inputs["bucketDomainName"] = nil
	inputs["urn"] = nil
	s, err := ctx.RegisterResource("digitalocean:index/spacesBucket:SpacesBucket", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SpacesBucket{s: s}, nil
}

// GetSpacesBucket gets an existing SpacesBucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpacesBucket(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SpacesBucketState, opts ...pulumi.ResourceOpt) (*SpacesBucket, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["acl"] = state.Acl
		inputs["bucketDomainName"] = state.BucketDomainName
		inputs["corsRules"] = state.CorsRules
		inputs["forceDestroy"] = state.ForceDestroy
		inputs["name"] = state.Name
		inputs["region"] = state.Region
		inputs["urn"] = state.Urn
	}
	s, err := ctx.ReadResource("digitalocean:index/spacesBucket:SpacesBucket", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SpacesBucket{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SpacesBucket) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SpacesBucket) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Canned ACL applied on bucket creation (`private` or `public-read`)
func (r *SpacesBucket) Acl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["acl"])
}

// The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com)
func (r *SpacesBucket) BucketDomainName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["bucketDomainName"])
}

// A container holding a list of elements describing allowed methods for a specific origin.
func (r *SpacesBucket) CorsRules() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["corsRules"])
}

// Unless `true`, the bucket will only be destroyed if empty (Defaults to `false`)
func (r *SpacesBucket) ForceDestroy() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["forceDestroy"])
}

// The name of the bucket
func (r *SpacesBucket) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The region where the bucket resides (Defaults to `nyc3`)
func (r *SpacesBucket) Region() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["region"])
}

// The uniform resource name for the bucket
func (r *SpacesBucket) Urn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["urn"])
}

// Input properties used for looking up and filtering SpacesBucket resources.
type SpacesBucketState struct {
	// Canned ACL applied on bucket creation (`private` or `public-read`)
	Acl interface{}
	// The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com)
	BucketDomainName interface{}
	// A container holding a list of elements describing allowed methods for a specific origin.
	CorsRules interface{}
	// Unless `true`, the bucket will only be destroyed if empty (Defaults to `false`)
	ForceDestroy interface{}
	// The name of the bucket
	Name interface{}
	// The region where the bucket resides (Defaults to `nyc3`)
	Region interface{}
	// The uniform resource name for the bucket
	Urn interface{}
}

// The set of arguments for constructing a SpacesBucket resource.
type SpacesBucketArgs struct {
	// Canned ACL applied on bucket creation (`private` or `public-read`)
	Acl interface{}
	// A container holding a list of elements describing allowed methods for a specific origin.
	CorsRules interface{}
	// Unless `true`, the bucket will only be destroyed if empty (Defaults to `false`)
	ForceDestroy interface{}
	// The name of the bucket
	Name interface{}
	// The region where the bucket resides (Defaults to `nyc3`)
	Region interface{}
}
