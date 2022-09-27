// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetSpacesBucketsBucket {
    /**
     * @return The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com)
     * 
     */
    private String bucketDomainName;
    /**
     * @return The name of the Spaces bucket
     * 
     */
    private String name;
    /**
     * @return The slug of the region where the bucket is stored.
     * 
     */
    private String region;
    /**
     * @return The uniform resource name of the bucket
     * 
     */
    private String urn;

    private GetSpacesBucketsBucket() {}
    /**
     * @return The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com)
     * 
     */
    public String bucketDomainName() {
        return this.bucketDomainName;
    }
    /**
     * @return The name of the Spaces bucket
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The slug of the region where the bucket is stored.
     * 
     */
    public String region() {
        return this.region;
    }
    /**
     * @return The uniform resource name of the bucket
     * 
     */
    public String urn() {
        return this.urn;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSpacesBucketsBucket defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String bucketDomainName;
        private String name;
        private String region;
        private String urn;
        public Builder() {}
        public Builder(GetSpacesBucketsBucket defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bucketDomainName = defaults.bucketDomainName;
    	      this.name = defaults.name;
    	      this.region = defaults.region;
    	      this.urn = defaults.urn;
        }

        @CustomType.Setter
        public Builder bucketDomainName(String bucketDomainName) {
            this.bucketDomainName = Objects.requireNonNull(bucketDomainName);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        @CustomType.Setter
        public Builder urn(String urn) {
            this.urn = Objects.requireNonNull(urn);
            return this;
        }
        public GetSpacesBucketsBucket build() {
            final var o = new GetSpacesBucketsBucket();
            o.bucketDomainName = bucketDomainName;
            o.name = name;
            o.region = region;
            o.urn = urn;
            return o;
        }
    }
}
