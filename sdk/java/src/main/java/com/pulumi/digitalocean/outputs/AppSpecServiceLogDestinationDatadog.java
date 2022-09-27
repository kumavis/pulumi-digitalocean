// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AppSpecServiceLogDestinationDatadog {
    /**
     * @return Datadog API key.
     * 
     */
    private String apiKey;
    /**
     * @return Datadog HTTP log intake endpoint.
     * 
     */
    private @Nullable String endpoint;

    private AppSpecServiceLogDestinationDatadog() {}
    /**
     * @return Datadog API key.
     * 
     */
    public String apiKey() {
        return this.apiKey;
    }
    /**
     * @return Datadog HTTP log intake endpoint.
     * 
     */
    public Optional<String> endpoint() {
        return Optional.ofNullable(this.endpoint);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AppSpecServiceLogDestinationDatadog defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String apiKey;
        private @Nullable String endpoint;
        public Builder() {}
        public Builder(AppSpecServiceLogDestinationDatadog defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.apiKey = defaults.apiKey;
    	      this.endpoint = defaults.endpoint;
        }

        @CustomType.Setter
        public Builder apiKey(String apiKey) {
            this.apiKey = Objects.requireNonNull(apiKey);
            return this;
        }
        @CustomType.Setter
        public Builder endpoint(@Nullable String endpoint) {
            this.endpoint = endpoint;
            return this;
        }
        public AppSpecServiceLogDestinationDatadog build() {
            final var o = new AppSpecServiceLogDestinationDatadog();
            o.apiKey = apiKey;
            o.endpoint = endpoint;
            return o;
        }
    }
}
