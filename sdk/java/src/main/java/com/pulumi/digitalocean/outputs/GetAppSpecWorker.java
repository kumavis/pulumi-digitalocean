// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.digitalocean.outputs.GetAppSpecWorkerAlert;
import com.pulumi.digitalocean.outputs.GetAppSpecWorkerEnv;
import com.pulumi.digitalocean.outputs.GetAppSpecWorkerGit;
import com.pulumi.digitalocean.outputs.GetAppSpecWorkerGithub;
import com.pulumi.digitalocean.outputs.GetAppSpecWorkerGitlab;
import com.pulumi.digitalocean.outputs.GetAppSpecWorkerImage;
import com.pulumi.digitalocean.outputs.GetAppSpecWorkerLogDestination;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetAppSpecWorker {
    /**
     * @return Describes an alert policy for the component.
     * 
     */
    private @Nullable List<GetAppSpecWorkerAlert> alerts;
    /**
     * @return An optional build command to run while building this component from source.
     * 
     */
    private @Nullable String buildCommand;
    /**
     * @return The path to a Dockerfile relative to the root of the repo. If set, overrides usage of buildpacks.
     * 
     */
    private @Nullable String dockerfilePath;
    /**
     * @return An environment slug describing the type of this app.
     * 
     */
    private @Nullable String environmentSlug;
    /**
     * @return Describes an environment variable made available to an app competent.
     * 
     */
    private @Nullable List<GetAppSpecWorkerEnv> envs;
    /**
     * @return A Git repo to use as the component&#39;s source. The repository must be able to be cloned without authentication.  Only one of `git`, `github` or `gitlab`  may be set.
     * 
     */
    private @Nullable GetAppSpecWorkerGit git;
    /**
     * @return A GitHub repo to use as the component&#39;s source. DigitalOcean App Platform must have [access to the repository](https://cloud.digitalocean.com/apps/github/install). Only one of `git`, `github`, `gitlab`, or `image` may be set.
     * 
     */
    private @Nullable GetAppSpecWorkerGithub github;
    /**
     * @return A Gitlab repo to use as the component&#39;s source. DigitalOcean App Platform must have [access to the repository](https://cloud.digitalocean.com/apps/gitlab/install). Only one of `git`, `github`, `gitlab`, or `image` may be set.
     * 
     */
    private @Nullable GetAppSpecWorkerGitlab gitlab;
    /**
     * @return An image to use as the component&#39;s source. Only one of `git`, `github`, `gitlab`, or `image` may be set.
     * 
     */
    private @Nullable GetAppSpecWorkerImage image;
    /**
     * @return The amount of instances that this component should be scaled to.
     * 
     */
    private @Nullable Integer instanceCount;
    /**
     * @return The instance size to use for this component.
     * 
     */
    private @Nullable String instanceSizeSlug;
    /**
     * @return Describes a log forwarding destination.
     * 
     */
    private @Nullable List<GetAppSpecWorkerLogDestination> logDestinations;
    /**
     * @return The name of the component.
     * 
     */
    private String name;
    /**
     * @return An optional run command to override the component&#39;s default.
     * 
     */
    private @Nullable String runCommand;
    /**
     * @return An optional path to the working directory to use for the build.
     * 
     */
    private @Nullable String sourceDir;

    private GetAppSpecWorker() {}
    /**
     * @return Describes an alert policy for the component.
     * 
     */
    public List<GetAppSpecWorkerAlert> alerts() {
        return this.alerts == null ? List.of() : this.alerts;
    }
    /**
     * @return An optional build command to run while building this component from source.
     * 
     */
    public Optional<String> buildCommand() {
        return Optional.ofNullable(this.buildCommand);
    }
    /**
     * @return The path to a Dockerfile relative to the root of the repo. If set, overrides usage of buildpacks.
     * 
     */
    public Optional<String> dockerfilePath() {
        return Optional.ofNullable(this.dockerfilePath);
    }
    /**
     * @return An environment slug describing the type of this app.
     * 
     */
    public Optional<String> environmentSlug() {
        return Optional.ofNullable(this.environmentSlug);
    }
    /**
     * @return Describes an environment variable made available to an app competent.
     * 
     */
    public List<GetAppSpecWorkerEnv> envs() {
        return this.envs == null ? List.of() : this.envs;
    }
    /**
     * @return A Git repo to use as the component&#39;s source. The repository must be able to be cloned without authentication.  Only one of `git`, `github` or `gitlab`  may be set.
     * 
     */
    public Optional<GetAppSpecWorkerGit> git() {
        return Optional.ofNullable(this.git);
    }
    /**
     * @return A GitHub repo to use as the component&#39;s source. DigitalOcean App Platform must have [access to the repository](https://cloud.digitalocean.com/apps/github/install). Only one of `git`, `github`, `gitlab`, or `image` may be set.
     * 
     */
    public Optional<GetAppSpecWorkerGithub> github() {
        return Optional.ofNullable(this.github);
    }
    /**
     * @return A Gitlab repo to use as the component&#39;s source. DigitalOcean App Platform must have [access to the repository](https://cloud.digitalocean.com/apps/gitlab/install). Only one of `git`, `github`, `gitlab`, or `image` may be set.
     * 
     */
    public Optional<GetAppSpecWorkerGitlab> gitlab() {
        return Optional.ofNullable(this.gitlab);
    }
    /**
     * @return An image to use as the component&#39;s source. Only one of `git`, `github`, `gitlab`, or `image` may be set.
     * 
     */
    public Optional<GetAppSpecWorkerImage> image() {
        return Optional.ofNullable(this.image);
    }
    /**
     * @return The amount of instances that this component should be scaled to.
     * 
     */
    public Optional<Integer> instanceCount() {
        return Optional.ofNullable(this.instanceCount);
    }
    /**
     * @return The instance size to use for this component.
     * 
     */
    public Optional<String> instanceSizeSlug() {
        return Optional.ofNullable(this.instanceSizeSlug);
    }
    /**
     * @return Describes a log forwarding destination.
     * 
     */
    public List<GetAppSpecWorkerLogDestination> logDestinations() {
        return this.logDestinations == null ? List.of() : this.logDestinations;
    }
    /**
     * @return The name of the component.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return An optional run command to override the component&#39;s default.
     * 
     */
    public Optional<String> runCommand() {
        return Optional.ofNullable(this.runCommand);
    }
    /**
     * @return An optional path to the working directory to use for the build.
     * 
     */
    public Optional<String> sourceDir() {
        return Optional.ofNullable(this.sourceDir);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAppSpecWorker defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<GetAppSpecWorkerAlert> alerts;
        private @Nullable String buildCommand;
        private @Nullable String dockerfilePath;
        private @Nullable String environmentSlug;
        private @Nullable List<GetAppSpecWorkerEnv> envs;
        private @Nullable GetAppSpecWorkerGit git;
        private @Nullable GetAppSpecWorkerGithub github;
        private @Nullable GetAppSpecWorkerGitlab gitlab;
        private @Nullable GetAppSpecWorkerImage image;
        private @Nullable Integer instanceCount;
        private @Nullable String instanceSizeSlug;
        private @Nullable List<GetAppSpecWorkerLogDestination> logDestinations;
        private String name;
        private @Nullable String runCommand;
        private @Nullable String sourceDir;
        public Builder() {}
        public Builder(GetAppSpecWorker defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.alerts = defaults.alerts;
    	      this.buildCommand = defaults.buildCommand;
    	      this.dockerfilePath = defaults.dockerfilePath;
    	      this.environmentSlug = defaults.environmentSlug;
    	      this.envs = defaults.envs;
    	      this.git = defaults.git;
    	      this.github = defaults.github;
    	      this.gitlab = defaults.gitlab;
    	      this.image = defaults.image;
    	      this.instanceCount = defaults.instanceCount;
    	      this.instanceSizeSlug = defaults.instanceSizeSlug;
    	      this.logDestinations = defaults.logDestinations;
    	      this.name = defaults.name;
    	      this.runCommand = defaults.runCommand;
    	      this.sourceDir = defaults.sourceDir;
        }

        @CustomType.Setter
        public Builder alerts(@Nullable List<GetAppSpecWorkerAlert> alerts) {
            this.alerts = alerts;
            return this;
        }
        public Builder alerts(GetAppSpecWorkerAlert... alerts) {
            return alerts(List.of(alerts));
        }
        @CustomType.Setter
        public Builder buildCommand(@Nullable String buildCommand) {
            this.buildCommand = buildCommand;
            return this;
        }
        @CustomType.Setter
        public Builder dockerfilePath(@Nullable String dockerfilePath) {
            this.dockerfilePath = dockerfilePath;
            return this;
        }
        @CustomType.Setter
        public Builder environmentSlug(@Nullable String environmentSlug) {
            this.environmentSlug = environmentSlug;
            return this;
        }
        @CustomType.Setter
        public Builder envs(@Nullable List<GetAppSpecWorkerEnv> envs) {
            this.envs = envs;
            return this;
        }
        public Builder envs(GetAppSpecWorkerEnv... envs) {
            return envs(List.of(envs));
        }
        @CustomType.Setter
        public Builder git(@Nullable GetAppSpecWorkerGit git) {
            this.git = git;
            return this;
        }
        @CustomType.Setter
        public Builder github(@Nullable GetAppSpecWorkerGithub github) {
            this.github = github;
            return this;
        }
        @CustomType.Setter
        public Builder gitlab(@Nullable GetAppSpecWorkerGitlab gitlab) {
            this.gitlab = gitlab;
            return this;
        }
        @CustomType.Setter
        public Builder image(@Nullable GetAppSpecWorkerImage image) {
            this.image = image;
            return this;
        }
        @CustomType.Setter
        public Builder instanceCount(@Nullable Integer instanceCount) {
            this.instanceCount = instanceCount;
            return this;
        }
        @CustomType.Setter
        public Builder instanceSizeSlug(@Nullable String instanceSizeSlug) {
            this.instanceSizeSlug = instanceSizeSlug;
            return this;
        }
        @CustomType.Setter
        public Builder logDestinations(@Nullable List<GetAppSpecWorkerLogDestination> logDestinations) {
            this.logDestinations = logDestinations;
            return this;
        }
        public Builder logDestinations(GetAppSpecWorkerLogDestination... logDestinations) {
            return logDestinations(List.of(logDestinations));
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder runCommand(@Nullable String runCommand) {
            this.runCommand = runCommand;
            return this;
        }
        @CustomType.Setter
        public Builder sourceDir(@Nullable String sourceDir) {
            this.sourceDir = sourceDir;
            return this;
        }
        public GetAppSpecWorker build() {
            final var o = new GetAppSpecWorker();
            o.alerts = alerts;
            o.buildCommand = buildCommand;
            o.dockerfilePath = dockerfilePath;
            o.environmentSlug = environmentSlug;
            o.envs = envs;
            o.git = git;
            o.github = github;
            o.gitlab = gitlab;
            o.image = image;
            o.instanceCount = instanceCount;
            o.instanceSizeSlug = instanceSizeSlug;
            o.logDestinations = logDestinations;
            o.name = name;
            o.runCommand = runCommand;
            o.sourceDir = sourceDir;
            return o;
        }
    }
}
