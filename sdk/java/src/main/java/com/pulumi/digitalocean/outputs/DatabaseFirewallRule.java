// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DatabaseFirewallRule {
    /**
     * @return The date and time when the firewall rule was created.
     * 
     */
    private @Nullable String createdAt;
    /**
     * @return The type of resource that the firewall rule allows to access the database cluster. The possible values are: `droplet`, `k8s`, `ip_addr`, `tag`, or `app`.
     * 
     */
    private String type;
    /**
     * @return A unique identifier for the firewall rule.
     * 
     */
    private @Nullable String uuid;
    /**
     * @return The ID of the specific resource, the name of a tag applied to a group of resources, or the IP address that the firewall rule allows to access the database cluster.
     * 
     */
    private String value;

    private DatabaseFirewallRule() {}
    /**
     * @return The date and time when the firewall rule was created.
     * 
     */
    public Optional<String> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }
    /**
     * @return The type of resource that the firewall rule allows to access the database cluster. The possible values are: `droplet`, `k8s`, `ip_addr`, `tag`, or `app`.
     * 
     */
    public String type() {
        return this.type;
    }
    /**
     * @return A unique identifier for the firewall rule.
     * 
     */
    public Optional<String> uuid() {
        return Optional.ofNullable(this.uuid);
    }
    /**
     * @return The ID of the specific resource, the name of a tag applied to a group of resources, or the IP address that the firewall rule allows to access the database cluster.
     * 
     */
    public String value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DatabaseFirewallRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String createdAt;
        private String type;
        private @Nullable String uuid;
        private String value;
        public Builder() {}
        public Builder(DatabaseFirewallRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.type = defaults.type;
    	      this.uuid = defaults.uuid;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder createdAt(@Nullable String createdAt) {
            this.createdAt = createdAt;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        @CustomType.Setter
        public Builder uuid(@Nullable String uuid) {
            this.uuid = uuid;
            return this;
        }
        @CustomType.Setter
        public Builder value(String value) {
            this.value = Objects.requireNonNull(value);
            return this;
        }
        public DatabaseFirewallRule build() {
            final var o = new DatabaseFirewallRule();
            o.createdAt = createdAt;
            o.type = type;
            o.uuid = uuid;
            o.value = value;
            return o;
        }
    }
}
