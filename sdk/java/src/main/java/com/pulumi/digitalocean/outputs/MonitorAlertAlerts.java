// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.digitalocean.outputs.MonitorAlertAlertsSlack;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class MonitorAlertAlerts {
    private @Nullable List<String> emails;
    private @Nullable List<MonitorAlertAlertsSlack> slacks;

    private MonitorAlertAlerts() {}
    public List<String> emails() {
        return this.emails == null ? List.of() : this.emails;
    }
    public List<MonitorAlertAlertsSlack> slacks() {
        return this.slacks == null ? List.of() : this.slacks;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(MonitorAlertAlerts defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> emails;
        private @Nullable List<MonitorAlertAlertsSlack> slacks;
        public Builder() {}
        public Builder(MonitorAlertAlerts defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.emails = defaults.emails;
    	      this.slacks = defaults.slacks;
        }

        @CustomType.Setter
        public Builder emails(@Nullable List<String> emails) {
            this.emails = emails;
            return this;
        }
        public Builder emails(String... emails) {
            return emails(List.of(emails));
        }
        @CustomType.Setter
        public Builder slacks(@Nullable List<MonitorAlertAlertsSlack> slacks) {
            this.slacks = slacks;
            return this;
        }
        public Builder slacks(MonitorAlertAlertsSlack... slacks) {
            return slacks(List.of(slacks));
        }
        public MonitorAlertAlerts build() {
            final var o = new MonitorAlertAlerts();
            o.emails = emails;
            o.slacks = slacks;
            return o;
        }
    }
}
