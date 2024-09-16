// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.digitalocean.outputs.AppSpecServiceLogDestinationOpenSearchBasicAuth;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AppSpecServiceLogDestinationOpenSearch {
    /**
     * @return Basic authentication details.
     * 
     */
    private AppSpecServiceLogDestinationOpenSearchBasicAuth basicAuth;
    /**
     * @return The name of the underlying DigitalOcean DBaaS cluster. This is required for production databases. For dev databases, if `cluster_name` is not set, a new cluster will be provisioned.
     * 
     */
    private @Nullable String clusterName;
    /**
     * @return OpenSearch endpoint.
     * 
     */
    private @Nullable String endpoint;
    /**
     * @return OpenSearch index name.
     * 
     */
    private @Nullable String indexName;

    private AppSpecServiceLogDestinationOpenSearch() {}
    /**
     * @return Basic authentication details.
     * 
     */
    public AppSpecServiceLogDestinationOpenSearchBasicAuth basicAuth() {
        return this.basicAuth;
    }
    /**
     * @return The name of the underlying DigitalOcean DBaaS cluster. This is required for production databases. For dev databases, if `cluster_name` is not set, a new cluster will be provisioned.
     * 
     */
    public Optional<String> clusterName() {
        return Optional.ofNullable(this.clusterName);
    }
    /**
     * @return OpenSearch endpoint.
     * 
     */
    public Optional<String> endpoint() {
        return Optional.ofNullable(this.endpoint);
    }
    /**
     * @return OpenSearch index name.
     * 
     */
    public Optional<String> indexName() {
        return Optional.ofNullable(this.indexName);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AppSpecServiceLogDestinationOpenSearch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private AppSpecServiceLogDestinationOpenSearchBasicAuth basicAuth;
        private @Nullable String clusterName;
        private @Nullable String endpoint;
        private @Nullable String indexName;
        public Builder() {}
        public Builder(AppSpecServiceLogDestinationOpenSearch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.basicAuth = defaults.basicAuth;
    	      this.clusterName = defaults.clusterName;
    	      this.endpoint = defaults.endpoint;
    	      this.indexName = defaults.indexName;
        }

        @CustomType.Setter
        public Builder basicAuth(AppSpecServiceLogDestinationOpenSearchBasicAuth basicAuth) {
            if (basicAuth == null) {
              throw new MissingRequiredPropertyException("AppSpecServiceLogDestinationOpenSearch", "basicAuth");
            }
            this.basicAuth = basicAuth;
            return this;
        }
        @CustomType.Setter
        public Builder clusterName(@Nullable String clusterName) {

            this.clusterName = clusterName;
            return this;
        }
        @CustomType.Setter
        public Builder endpoint(@Nullable String endpoint) {

            this.endpoint = endpoint;
            return this;
        }
        @CustomType.Setter
        public Builder indexName(@Nullable String indexName) {

            this.indexName = indexName;
            return this;
        }
        public AppSpecServiceLogDestinationOpenSearch build() {
            final var _resultValue = new AppSpecServiceLogDestinationOpenSearch();
            _resultValue.basicAuth = basicAuth;
            _resultValue.clusterName = clusterName;
            _resultValue.endpoint = endpoint;
            _resultValue.indexName = indexName;
            return _resultValue;
        }
    }
}
