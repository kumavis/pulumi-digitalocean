// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AppSpecIngressRuleMatchPathArgs extends com.pulumi.resources.ResourceArgs {

    public static final AppSpecIngressRuleMatchPathArgs Empty = new AppSpecIngressRuleMatchPathArgs();

    /**
     * Prefix-based match.
     * 
     */
    @Import(name="prefix")
    private @Nullable Output<String> prefix;

    /**
     * @return Prefix-based match.
     * 
     */
    public Optional<Output<String>> prefix() {
        return Optional.ofNullable(this.prefix);
    }

    private AppSpecIngressRuleMatchPathArgs() {}

    private AppSpecIngressRuleMatchPathArgs(AppSpecIngressRuleMatchPathArgs $) {
        this.prefix = $.prefix;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppSpecIngressRuleMatchPathArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppSpecIngressRuleMatchPathArgs $;

        public Builder() {
            $ = new AppSpecIngressRuleMatchPathArgs();
        }

        public Builder(AppSpecIngressRuleMatchPathArgs defaults) {
            $ = new AppSpecIngressRuleMatchPathArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param prefix Prefix-based match.
         * 
         * @return builder
         * 
         */
        public Builder prefix(@Nullable Output<String> prefix) {
            $.prefix = prefix;
            return this;
        }

        /**
         * @param prefix Prefix-based match.
         * 
         * @return builder
         * 
         */
        public Builder prefix(String prefix) {
            return prefix(Output.of(prefix));
        }

        public AppSpecIngressRuleMatchPathArgs build() {
            return $;
        }
    }

}
