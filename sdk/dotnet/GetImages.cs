// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean
{
    public static class GetImages
    {
        /// <summary>
        /// Get information on images for use in other resources (e.g. creating a Droplet
        /// based on a snapshot), with the ability to filter and sort the results. If no filters are specified,
        /// all images will be returned.
        /// 
        /// This data source is useful if the image in question is not managed by this provider or you need to utilize any
        /// of the image's data.
        /// 
        /// Note: You can use the `digitalocean.getImage` data source to obtain metadata
        /// about a single image if you already know the `slug`, unique `name`, or `id` to retrieve.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Use the `filter` block with a `key` string and `values` list to filter images.
        /// 
        /// For example to find all Ubuntu images:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var ubuntu = Output.Create(DigitalOcean.GetImages.InvokeAsync(new DigitalOcean.GetImagesArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new DigitalOcean.Inputs.GetImagesFilterArgs
        ///                 {
        ///                     Key = "distribution",
        ///                     Values = 
        ///                     {
        ///                         "Ubuntu",
        ///                     },
        ///                 },
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// You can filter on multiple fields and sort the results as well:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var available = Output.Create(DigitalOcean.GetImages.InvokeAsync(new DigitalOcean.GetImagesArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new DigitalOcean.Inputs.GetImagesFilterArgs
        ///                 {
        ///                     Key = "distribution",
        ///                     Values = 
        ///                     {
        ///                         "Ubuntu",
        ///                     },
        ///                 },
        ///                 new DigitalOcean.Inputs.GetImagesFilterArgs
        ///                 {
        ///                     Key = "regions",
        ///                     Values = 
        ///                     {
        ///                         "nyc3",
        ///                     },
        ///                 },
        ///             },
        ///             Sorts = 
        ///             {
        ///                 new DigitalOcean.Inputs.GetImagesSortArgs
        ///                 {
        ///                     Direction = "desc",
        ///                     Key = "created",
        ///                 },
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetImagesResult> InvokeAsync(GetImagesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetImagesResult>("digitalocean:index/getImages:getImages", args ?? new GetImagesArgs(), options.WithVersion());
    }


    public sealed class GetImagesArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetImagesFilterArgs>? _filters;

        /// <summary>
        /// Filter the results.
        /// The `filter` block is documented below.
        /// </summary>
        public List<Inputs.GetImagesFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetImagesFilterArgs>());
            set => _filters = value;
        }

        [Input("sorts")]
        private List<Inputs.GetImagesSortArgs>? _sorts;

        /// <summary>
        /// Sort the results.
        /// The `sort` block is documented below.
        /// </summary>
        public List<Inputs.GetImagesSortArgs> Sorts
        {
            get => _sorts ?? (_sorts = new List<Inputs.GetImagesSortArgs>());
            set => _sorts = value;
        }

        public GetImagesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetImagesResult
    {
        public readonly ImmutableArray<Outputs.GetImagesFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A set of images satisfying any `filter` and `sort` criteria. Each image has the following attributes:  
        /// - `slug`: Unique text identifier of the image.
        /// - `id`: The ID of the image.
        /// - `name`: The name of the image.
        /// - `type`: Type of the image.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetImagesImageResult> Images;
        public readonly ImmutableArray<Outputs.GetImagesSortResult> Sorts;

        [OutputConstructor]
        private GetImagesResult(
            ImmutableArray<Outputs.GetImagesFilterResult> filters,

            string id,

            ImmutableArray<Outputs.GetImagesImageResult> images,

            ImmutableArray<Outputs.GetImagesSortResult> sorts)
        {
            Filters = filters;
            Id = id;
            Images = images;
            Sorts = sorts;
        }
    }
}
