// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetAppSpecJobLogDestinationDatadog {
    /**
     * @return Datadog API key.
     * 
     */
    private String apiKey;
    /**
     * @return OpenSearch API Endpoint. Only HTTPS is supported. Format: https://&lt;host&gt;:&lt;port&gt;.
     * 
     */
    private @Nullable String endpoint;

    private GetAppSpecJobLogDestinationDatadog() {}
    /**
     * @return Datadog API key.
     * 
     */
    public String apiKey() {
        return this.apiKey;
    }
    /**
     * @return OpenSearch API Endpoint. Only HTTPS is supported. Format: https://&lt;host&gt;:&lt;port&gt;.
     * 
     */
    public Optional<String> endpoint() {
        return Optional.ofNullable(this.endpoint);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAppSpecJobLogDestinationDatadog defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String apiKey;
        private @Nullable String endpoint;
        public Builder() {}
        public Builder(GetAppSpecJobLogDestinationDatadog defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.apiKey = defaults.apiKey;
    	      this.endpoint = defaults.endpoint;
        }

        @CustomType.Setter
        public Builder apiKey(String apiKey) {
            if (apiKey == null) {
              throw new MissingRequiredPropertyException("GetAppSpecJobLogDestinationDatadog", "apiKey");
            }
            this.apiKey = apiKey;
            return this;
        }
        @CustomType.Setter
        public Builder endpoint(@Nullable String endpoint) {

            this.endpoint = endpoint;
            return this;
        }
        public GetAppSpecJobLogDestinationDatadog build() {
            final var _resultValue = new GetAppSpecJobLogDestinationDatadog();
            _resultValue.apiKey = apiKey;
            _resultValue.endpoint = endpoint;
            return _resultValue;
        }
    }
}
