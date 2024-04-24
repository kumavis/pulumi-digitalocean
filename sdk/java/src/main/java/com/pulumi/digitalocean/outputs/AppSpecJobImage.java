// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.digitalocean.outputs.AppSpecJobImageDeployOnPush;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AppSpecJobImage {
    /**
     * @return Configures automatically deploying images pushed to DOCR.
     * 
     */
    private @Nullable List<AppSpecJobImageDeployOnPush> deployOnPushes;
    /**
     * @return The registry name. Must be left empty for the `DOCR` registry type. Required for the `DOCKER_HUB` registry type.
     * 
     */
    private @Nullable String registry;
    /**
     * @return Access credentials for third-party registries
     * 
     */
    private String registryCredentials;
    /**
     * @return The registry type. One of `DOCR` (DigitalOcean container registry) or `DOCKER_HUB`.
     * 
     */
    private String registryType;
    /**
     * @return The repository name.
     * 
     */
    private String repository;
    /**
     * @return The repository tag. Defaults to `latest` if not provided.
     * 
     */
    private @Nullable String tag;

    private AppSpecJobImage() {}
    /**
     * @return Configures automatically deploying images pushed to DOCR.
     * 
     */
    public List<AppSpecJobImageDeployOnPush> deployOnPushes() {
        return this.deployOnPushes == null ? List.of() : this.deployOnPushes;
    }
    /**
     * @return The registry name. Must be left empty for the `DOCR` registry type. Required for the `DOCKER_HUB` registry type.
     * 
     */
    public Optional<String> registry() {
        return Optional.ofNullable(this.registry);
    }
    /**
     * @return Access credentials for third-party registries
     * 
     */
    public String registryCredentials() {
        return this.registryCredentials;
    }
    /**
     * @return The registry type. One of `DOCR` (DigitalOcean container registry) or `DOCKER_HUB`.
     * 
     */
    public String registryType() {
        return this.registryType;
    }
    /**
     * @return The repository name.
     * 
     */
    public String repository() {
        return this.repository;
    }
    /**
     * @return The repository tag. Defaults to `latest` if not provided.
     * 
     */
    public Optional<String> tag() {
        return Optional.ofNullable(this.tag);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AppSpecJobImage defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<AppSpecJobImageDeployOnPush> deployOnPushes;
        private @Nullable String registry;
        private String registryCredentials;
        private String registryType;
        private String repository;
        private @Nullable String tag;
        public Builder() {}
        public Builder(AppSpecJobImage defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.deployOnPushes = defaults.deployOnPushes;
    	      this.registry = defaults.registry;
    	      this.registryCredentials = defaults.registryCredentials;
    	      this.registryType = defaults.registryType;
    	      this.repository = defaults.repository;
    	      this.tag = defaults.tag;
        }

        @CustomType.Setter
        public Builder deployOnPushes(@Nullable List<AppSpecJobImageDeployOnPush> deployOnPushes) {

            this.deployOnPushes = deployOnPushes;
            return this;
        }
        public Builder deployOnPushes(AppSpecJobImageDeployOnPush... deployOnPushes) {
            return deployOnPushes(List.of(deployOnPushes));
        }
        @CustomType.Setter
        public Builder registry(@Nullable String registry) {

            this.registry = registry;
            return this;
        }
        @CustomType.Setter
        public Builder registryCredentials(String registryCredentials) {
            if (registryCredentials == null) {
              throw new MissingRequiredPropertyException("AppSpecJobImage", "registryCredentials");
            }
            this.registryCredentials = registryCredentials;
            return this;
        }
        @CustomType.Setter
        public Builder registryType(String registryType) {
            if (registryType == null) {
              throw new MissingRequiredPropertyException("AppSpecJobImage", "registryType");
            }
            this.registryType = registryType;
            return this;
        }
        @CustomType.Setter
        public Builder repository(String repository) {
            if (repository == null) {
              throw new MissingRequiredPropertyException("AppSpecJobImage", "repository");
            }
            this.repository = repository;
            return this;
        }
        @CustomType.Setter
        public Builder tag(@Nullable String tag) {

            this.tag = tag;
            return this;
        }
        public AppSpecJobImage build() {
            final var _resultValue = new AppSpecJobImage();
            _resultValue.deployOnPushes = deployOnPushes;
            _resultValue.registry = registry;
            _resultValue.registryCredentials = registryCredentials;
            _resultValue.registryType = registryType;
            _resultValue.repository = repository;
            _resultValue.tag = tag;
            return _resultValue;
        }
    }
}
