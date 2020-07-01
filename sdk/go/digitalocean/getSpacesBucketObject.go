// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The Spaces object data source allows access to the metadata and
// _optionally_ (see below) content of an object stored inside a Spaces bucket.
//
// > **Note:** The content of an object (`body` field) is available only for objects which have a human-readable
// `Content-Type` (`text/*` and `application/json`). This is to prevent printing unsafe characters and potentially
// downloading large amount of data which would be thrown away in favor of metadata.
//
// ## Example Usage
//
// The following example retrieves a text object (which must have a `Content-Type`
// value starting with `text/`) and uses it as the `userData` for a Droplet:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v2/go/digitalocean"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bootstrapScript, err := digitalocean.LookupSpacesBucketObject(ctx, &digitalocean.LookupSpacesBucketObjectArgs{
// 			Bucket: "ourcorp-deploy-config",
// 			Region: "nyc3",
// 			Key:    "droplet-bootstrap-script.sh",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = digitalocean.NewDroplet(ctx, "web", &digitalocean.DropletArgs{
// 			Image:    pulumi.String("ubuntu-18-04-x64"),
// 			Region:   pulumi.String("nyc2"),
// 			Size:     pulumi.String("s-1vcpu-1gb"),
// 			UserData: pulumi.String(bootstrapScript.Body),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupSpacesBucketObject(ctx *pulumi.Context, args *LookupSpacesBucketObjectArgs, opts ...pulumi.InvokeOption) (*LookupSpacesBucketObjectResult, error) {
	var rv LookupSpacesBucketObjectResult
	err := ctx.Invoke("digitalocean:index/getSpacesBucketObject:getSpacesBucketObject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSpacesBucketObject.
type LookupSpacesBucketObjectArgs struct {
	// The name of the bucket to read the object from.
	Bucket string `pulumi:"bucket"`
	// The full path to the object inside the bucket
	Key   string  `pulumi:"key"`
	Range *string `pulumi:"range"`
	// The slug of the region where the bucket is stored.
	Region string `pulumi:"region"`
	// Specific version ID of the object returned (defaults to latest version)
	VersionId *string `pulumi:"versionId"`
}

// A collection of values returned by getSpacesBucketObject.
type LookupSpacesBucketObjectResult struct {
	// Object data (see **limitations above** to understand cases in which this field is actually available)
	Body   string `pulumi:"body"`
	Bucket string `pulumi:"bucket"`
	// Specifies caching behavior along the request/reply chain.
	CacheControl string `pulumi:"cacheControl"`
	// Specifies presentational information for the object.
	ContentDisposition string `pulumi:"contentDisposition"`
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
	ContentEncoding string `pulumi:"contentEncoding"`
	// The language the content is in.
	ContentLanguage string `pulumi:"contentLanguage"`
	// Size of the body in bytes.
	ContentLength int `pulumi:"contentLength"`
	// A standard MIME type describing the format of the object data.
	ContentType string `pulumi:"contentType"`
	// [ETag](https://en.wikipedia.org/wiki/HTTP_ETag) generated for the object (an MD5 sum of the object content in case it's not encrypted)
	Etag string `pulumi:"etag"`
	// If the object expiration is configured (see [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html)), the field includes this header. It includes the expiry-date and rule-id key value pairs providing object expiration information. The value of the rule-id is URL encoded.
	Expiration string `pulumi:"expiration"`
	// The date and time at which the object is no longer cacheable.
	Expires string `pulumi:"expires"`
	// The provider-assigned unique ID for this managed resource.
	Id  string `pulumi:"id"`
	Key string `pulumi:"key"`
	// Last modified date of the object in RFC1123 format (e.g. `Mon, 02 Jan 2006 15:04:05 MST`)
	LastModified string `pulumi:"lastModified"`
	// A map of metadata stored with the object in Spaces
	Metadata map[string]interface{} `pulumi:"metadata"`
	Range    *string                `pulumi:"range"`
	Region   string                 `pulumi:"region"`
	// The latest version ID of the object returned.
	VersionId string `pulumi:"versionId"`
	// If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. Spaces stores the value of this header in the object metadata.
	WebsiteRedirectLocation string `pulumi:"websiteRedirectLocation"`
}
