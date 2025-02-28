// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class AppSpecFunctionLogDestinationLogtail {
    /**
     * @return Logtail token.
     * 
     */
    private String token;

    private AppSpecFunctionLogDestinationLogtail() {}
    /**
     * @return Logtail token.
     * 
     */
    public String token() {
        return this.token;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AppSpecFunctionLogDestinationLogtail defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String token;
        public Builder() {}
        public Builder(AppSpecFunctionLogDestinationLogtail defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.token = defaults.token;
        }

        @CustomType.Setter
        public Builder token(String token) {
            if (token == null) {
              throw new MissingRequiredPropertyException("AppSpecFunctionLogDestinationLogtail", "token");
            }
            this.token = token;
            return this;
        }
        public AppSpecFunctionLogDestinationLogtail build() {
            final var _resultValue = new AppSpecFunctionLogDestinationLogtail();
            _resultValue.token = token;
            return _resultValue;
        }
    }
}
