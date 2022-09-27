// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.digitalocean.outputs.AppSpecServiceLogDestinationDatadog;
import com.pulumi.digitalocean.outputs.AppSpecServiceLogDestinationLogtail;
import com.pulumi.digitalocean.outputs.AppSpecServiceLogDestinationPapertrail;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AppSpecServiceLogDestination {
    /**
     * @return Datadog configuration.
     * 
     */
    private @Nullable AppSpecServiceLogDestinationDatadog datadog;
    /**
     * @return Logtail configuration.
     * 
     */
    private @Nullable AppSpecServiceLogDestinationLogtail logtail;
    /**
     * @return The name of the component.
     * 
     */
    private String name;
    /**
     * @return Papertrail configuration.
     * 
     */
    private @Nullable AppSpecServiceLogDestinationPapertrail papertrail;

    private AppSpecServiceLogDestination() {}
    /**
     * @return Datadog configuration.
     * 
     */
    public Optional<AppSpecServiceLogDestinationDatadog> datadog() {
        return Optional.ofNullable(this.datadog);
    }
    /**
     * @return Logtail configuration.
     * 
     */
    public Optional<AppSpecServiceLogDestinationLogtail> logtail() {
        return Optional.ofNullable(this.logtail);
    }
    /**
     * @return The name of the component.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Papertrail configuration.
     * 
     */
    public Optional<AppSpecServiceLogDestinationPapertrail> papertrail() {
        return Optional.ofNullable(this.papertrail);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AppSpecServiceLogDestination defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable AppSpecServiceLogDestinationDatadog datadog;
        private @Nullable AppSpecServiceLogDestinationLogtail logtail;
        private String name;
        private @Nullable AppSpecServiceLogDestinationPapertrail papertrail;
        public Builder() {}
        public Builder(AppSpecServiceLogDestination defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.datadog = defaults.datadog;
    	      this.logtail = defaults.logtail;
    	      this.name = defaults.name;
    	      this.papertrail = defaults.papertrail;
        }

        @CustomType.Setter
        public Builder datadog(@Nullable AppSpecServiceLogDestinationDatadog datadog) {
            this.datadog = datadog;
            return this;
        }
        @CustomType.Setter
        public Builder logtail(@Nullable AppSpecServiceLogDestinationLogtail logtail) {
            this.logtail = logtail;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder papertrail(@Nullable AppSpecServiceLogDestinationPapertrail papertrail) {
            this.papertrail = papertrail;
            return this;
        }
        public AppSpecServiceLogDestination build() {
            final var o = new AppSpecServiceLogDestination();
            o.datadog = datadog;
            o.logtail = logtail;
            o.name = name;
            o.papertrail = papertrail;
            return o;
        }
    }
}
