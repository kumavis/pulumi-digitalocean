// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.digitalocean.outputs.GetAppSpecAlert;
import com.pulumi.digitalocean.outputs.GetAppSpecDatabase;
import com.pulumi.digitalocean.outputs.GetAppSpecEnv;
import com.pulumi.digitalocean.outputs.GetAppSpecFunction;
import com.pulumi.digitalocean.outputs.GetAppSpecJob;
import com.pulumi.digitalocean.outputs.GetAppSpecService;
import com.pulumi.digitalocean.outputs.GetAppSpecStaticSite;
import com.pulumi.digitalocean.outputs.GetAppSpecWorker;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetAppSpec {
    /**
     * @return Describes an alert policy for the component.
     * 
     */
    private @Nullable List<GetAppSpecAlert> alerts;
    private @Nullable List<GetAppSpecDatabase> databases;
    /**
     * @deprecated
     * This attribute has been replaced by `domain` which supports additional functionality.
     * 
     */
    @Deprecated /* This attribute has been replaced by `domain` which supports additional functionality. */
    private List<String> domains;
    /**
     * @return Describes an environment variable made available to an app competent.
     * 
     */
    private @Nullable List<GetAppSpecEnv> envs;
    private @Nullable List<GetAppSpecFunction> functions;
    private @Nullable List<GetAppSpecJob> jobs;
    /**
     * @return The name of the component.
     * 
     */
    private String name;
    private @Nullable String region;
    private @Nullable List<GetAppSpecService> services;
    private @Nullable List<GetAppSpecStaticSite> staticSites;
    private @Nullable List<GetAppSpecWorker> workers;

    private GetAppSpec() {}
    /**
     * @return Describes an alert policy for the component.
     * 
     */
    public List<GetAppSpecAlert> alerts() {
        return this.alerts == null ? List.of() : this.alerts;
    }
    public List<GetAppSpecDatabase> databases() {
        return this.databases == null ? List.of() : this.databases;
    }
    /**
     * @deprecated
     * This attribute has been replaced by `domain` which supports additional functionality.
     * 
     */
    @Deprecated /* This attribute has been replaced by `domain` which supports additional functionality. */
    public List<String> domains() {
        return this.domains;
    }
    /**
     * @return Describes an environment variable made available to an app competent.
     * 
     */
    public List<GetAppSpecEnv> envs() {
        return this.envs == null ? List.of() : this.envs;
    }
    public List<GetAppSpecFunction> functions() {
        return this.functions == null ? List.of() : this.functions;
    }
    public List<GetAppSpecJob> jobs() {
        return this.jobs == null ? List.of() : this.jobs;
    }
    /**
     * @return The name of the component.
     * 
     */
    public String name() {
        return this.name;
    }
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }
    public List<GetAppSpecService> services() {
        return this.services == null ? List.of() : this.services;
    }
    public List<GetAppSpecStaticSite> staticSites() {
        return this.staticSites == null ? List.of() : this.staticSites;
    }
    public List<GetAppSpecWorker> workers() {
        return this.workers == null ? List.of() : this.workers;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAppSpec defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<GetAppSpecAlert> alerts;
        private @Nullable List<GetAppSpecDatabase> databases;
        private List<String> domains;
        private @Nullable List<GetAppSpecEnv> envs;
        private @Nullable List<GetAppSpecFunction> functions;
        private @Nullable List<GetAppSpecJob> jobs;
        private String name;
        private @Nullable String region;
        private @Nullable List<GetAppSpecService> services;
        private @Nullable List<GetAppSpecStaticSite> staticSites;
        private @Nullable List<GetAppSpecWorker> workers;
        public Builder() {}
        public Builder(GetAppSpec defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.alerts = defaults.alerts;
    	      this.databases = defaults.databases;
    	      this.domains = defaults.domains;
    	      this.envs = defaults.envs;
    	      this.functions = defaults.functions;
    	      this.jobs = defaults.jobs;
    	      this.name = defaults.name;
    	      this.region = defaults.region;
    	      this.services = defaults.services;
    	      this.staticSites = defaults.staticSites;
    	      this.workers = defaults.workers;
        }

        @CustomType.Setter
        public Builder alerts(@Nullable List<GetAppSpecAlert> alerts) {
            this.alerts = alerts;
            return this;
        }
        public Builder alerts(GetAppSpecAlert... alerts) {
            return alerts(List.of(alerts));
        }
        @CustomType.Setter
        public Builder databases(@Nullable List<GetAppSpecDatabase> databases) {
            this.databases = databases;
            return this;
        }
        public Builder databases(GetAppSpecDatabase... databases) {
            return databases(List.of(databases));
        }
        @CustomType.Setter
        public Builder domains(List<String> domains) {
            this.domains = Objects.requireNonNull(domains);
            return this;
        }
        public Builder domains(String... domains) {
            return domains(List.of(domains));
        }
        @CustomType.Setter
        public Builder envs(@Nullable List<GetAppSpecEnv> envs) {
            this.envs = envs;
            return this;
        }
        public Builder envs(GetAppSpecEnv... envs) {
            return envs(List.of(envs));
        }
        @CustomType.Setter
        public Builder functions(@Nullable List<GetAppSpecFunction> functions) {
            this.functions = functions;
            return this;
        }
        public Builder functions(GetAppSpecFunction... functions) {
            return functions(List.of(functions));
        }
        @CustomType.Setter
        public Builder jobs(@Nullable List<GetAppSpecJob> jobs) {
            this.jobs = jobs;
            return this;
        }
        public Builder jobs(GetAppSpecJob... jobs) {
            return jobs(List.of(jobs));
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder region(@Nullable String region) {
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder services(@Nullable List<GetAppSpecService> services) {
            this.services = services;
            return this;
        }
        public Builder services(GetAppSpecService... services) {
            return services(List.of(services));
        }
        @CustomType.Setter
        public Builder staticSites(@Nullable List<GetAppSpecStaticSite> staticSites) {
            this.staticSites = staticSites;
            return this;
        }
        public Builder staticSites(GetAppSpecStaticSite... staticSites) {
            return staticSites(List.of(staticSites));
        }
        @CustomType.Setter
        public Builder workers(@Nullable List<GetAppSpecWorker> workers) {
            this.workers = workers;
            return this;
        }
        public Builder workers(GetAppSpecWorker... workers) {
            return workers(List.of(workers));
        }
        public GetAppSpec build() {
            final var o = new GetAppSpec();
            o.alerts = alerts;
            o.databases = databases;
            o.domains = domains;
            o.envs = envs;
            o.functions = functions;
            o.jobs = jobs;
            o.name = name;
            o.region = region;
            o.services = services;
            o.staticSites = staticSites;
            o.workers = workers;
            return o;
        }
    }
}
