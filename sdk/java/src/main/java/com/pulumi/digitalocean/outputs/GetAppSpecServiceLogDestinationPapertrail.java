// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetAppSpecServiceLogDestinationPapertrail {
    /**
     * @return Datadog HTTP log intake endpoint.
     * 
     */
    private String endpoint;

    private GetAppSpecServiceLogDestinationPapertrail() {}
    /**
     * @return Datadog HTTP log intake endpoint.
     * 
     */
    public String endpoint() {
        return this.endpoint;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAppSpecServiceLogDestinationPapertrail defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String endpoint;
        public Builder() {}
        public Builder(GetAppSpecServiceLogDestinationPapertrail defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.endpoint = defaults.endpoint;
        }

        @CustomType.Setter
        public Builder endpoint(String endpoint) {
            this.endpoint = Objects.requireNonNull(endpoint);
            return this;
        }
        public GetAppSpecServiceLogDestinationPapertrail build() {
            final var o = new GetAppSpecServiceLogDestinationPapertrail();
            o.endpoint = endpoint;
            return o;
        }
    }
}
