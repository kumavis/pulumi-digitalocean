// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.digitalocean.inputs.SpacesBucketLifecycleRuleExpirationArgs;
import com.pulumi.digitalocean.inputs.SpacesBucketLifecycleRuleNoncurrentVersionExpirationArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SpacesBucketLifecycleRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final SpacesBucketLifecycleRuleArgs Empty = new SpacesBucketLifecycleRuleArgs();

    /**
     * Specifies the number of days after initiating a multipart
     * upload when the multipart upload must be completed or else Spaces will abort the upload.
     * 
     */
    @Import(name="abortIncompleteMultipartUploadDays")
    private @Nullable Output<Integer> abortIncompleteMultipartUploadDays;

    /**
     * @return Specifies the number of days after initiating a multipart
     * upload when the multipart upload must be completed or else Spaces will abort the upload.
     * 
     */
    public Optional<Output<Integer>> abortIncompleteMultipartUploadDays() {
        return Optional.ofNullable(this.abortIncompleteMultipartUploadDays);
    }

    /**
     * Specifies lifecycle rule status.
     * 
     */
    @Import(name="enabled", required=true)
    private Output<Boolean> enabled;

    /**
     * @return Specifies lifecycle rule status.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }

    /**
     * Specifies a time period after which applicable objects expire (documented below).
     * 
     */
    @Import(name="expiration")
    private @Nullable Output<SpacesBucketLifecycleRuleExpirationArgs> expiration;

    /**
     * @return Specifies a time period after which applicable objects expire (documented below).
     * 
     */
    public Optional<Output<SpacesBucketLifecycleRuleExpirationArgs>> expiration() {
        return Optional.ofNullable(this.expiration);
    }

    /**
     * Unique identifier for the rule.
     * 
     */
    @Import(name="id")
    private @Nullable Output<String> id;

    /**
     * @return Unique identifier for the rule.
     * 
     */
    public Optional<Output<String>> id() {
        return Optional.ofNullable(this.id);
    }

    /**
     * Specifies when non-current object versions expire (documented below).
     * 
     * At least one of `expiration` or `noncurrent_version_expiration` must be specified.
     * 
     */
    @Import(name="noncurrentVersionExpiration")
    private @Nullable Output<SpacesBucketLifecycleRuleNoncurrentVersionExpirationArgs> noncurrentVersionExpiration;

    /**
     * @return Specifies when non-current object versions expire (documented below).
     * 
     * At least one of `expiration` or `noncurrent_version_expiration` must be specified.
     * 
     */
    public Optional<Output<SpacesBucketLifecycleRuleNoncurrentVersionExpirationArgs>> noncurrentVersionExpiration() {
        return Optional.ofNullable(this.noncurrentVersionExpiration);
    }

    /**
     * Object key prefix identifying one or more objects to which the rule applies.
     * 
     */
    @Import(name="prefix")
    private @Nullable Output<String> prefix;

    /**
     * @return Object key prefix identifying one or more objects to which the rule applies.
     * 
     */
    public Optional<Output<String>> prefix() {
        return Optional.ofNullable(this.prefix);
    }

    private SpacesBucketLifecycleRuleArgs() {}

    private SpacesBucketLifecycleRuleArgs(SpacesBucketLifecycleRuleArgs $) {
        this.abortIncompleteMultipartUploadDays = $.abortIncompleteMultipartUploadDays;
        this.enabled = $.enabled;
        this.expiration = $.expiration;
        this.id = $.id;
        this.noncurrentVersionExpiration = $.noncurrentVersionExpiration;
        this.prefix = $.prefix;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SpacesBucketLifecycleRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SpacesBucketLifecycleRuleArgs $;

        public Builder() {
            $ = new SpacesBucketLifecycleRuleArgs();
        }

        public Builder(SpacesBucketLifecycleRuleArgs defaults) {
            $ = new SpacesBucketLifecycleRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param abortIncompleteMultipartUploadDays Specifies the number of days after initiating a multipart
         * upload when the multipart upload must be completed or else Spaces will abort the upload.
         * 
         * @return builder
         * 
         */
        public Builder abortIncompleteMultipartUploadDays(@Nullable Output<Integer> abortIncompleteMultipartUploadDays) {
            $.abortIncompleteMultipartUploadDays = abortIncompleteMultipartUploadDays;
            return this;
        }

        /**
         * @param abortIncompleteMultipartUploadDays Specifies the number of days after initiating a multipart
         * upload when the multipart upload must be completed or else Spaces will abort the upload.
         * 
         * @return builder
         * 
         */
        public Builder abortIncompleteMultipartUploadDays(Integer abortIncompleteMultipartUploadDays) {
            return abortIncompleteMultipartUploadDays(Output.of(abortIncompleteMultipartUploadDays));
        }

        /**
         * @param enabled Specifies lifecycle rule status.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Specifies lifecycle rule status.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param expiration Specifies a time period after which applicable objects expire (documented below).
         * 
         * @return builder
         * 
         */
        public Builder expiration(@Nullable Output<SpacesBucketLifecycleRuleExpirationArgs> expiration) {
            $.expiration = expiration;
            return this;
        }

        /**
         * @param expiration Specifies a time period after which applicable objects expire (documented below).
         * 
         * @return builder
         * 
         */
        public Builder expiration(SpacesBucketLifecycleRuleExpirationArgs expiration) {
            return expiration(Output.of(expiration));
        }

        /**
         * @param id Unique identifier for the rule.
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable Output<String> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id Unique identifier for the rule.
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            return id(Output.of(id));
        }

        /**
         * @param noncurrentVersionExpiration Specifies when non-current object versions expire (documented below).
         * 
         * At least one of `expiration` or `noncurrent_version_expiration` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder noncurrentVersionExpiration(@Nullable Output<SpacesBucketLifecycleRuleNoncurrentVersionExpirationArgs> noncurrentVersionExpiration) {
            $.noncurrentVersionExpiration = noncurrentVersionExpiration;
            return this;
        }

        /**
         * @param noncurrentVersionExpiration Specifies when non-current object versions expire (documented below).
         * 
         * At least one of `expiration` or `noncurrent_version_expiration` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder noncurrentVersionExpiration(SpacesBucketLifecycleRuleNoncurrentVersionExpirationArgs noncurrentVersionExpiration) {
            return noncurrentVersionExpiration(Output.of(noncurrentVersionExpiration));
        }

        /**
         * @param prefix Object key prefix identifying one or more objects to which the rule applies.
         * 
         * @return builder
         * 
         */
        public Builder prefix(@Nullable Output<String> prefix) {
            $.prefix = prefix;
            return this;
        }

        /**
         * @param prefix Object key prefix identifying one or more objects to which the rule applies.
         * 
         * @return builder
         * 
         */
        public Builder prefix(String prefix) {
            return prefix(Output.of(prefix));
        }

        public SpacesBucketLifecycleRuleArgs build() {
            $.enabled = Objects.requireNonNull($.enabled, "expected parameter 'enabled' to be non-null");
            return $;
        }
    }

}
