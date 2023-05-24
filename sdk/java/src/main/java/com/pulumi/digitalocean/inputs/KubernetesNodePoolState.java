// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.inputs;

import com.pulumi.core.Either;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.digitalocean.enums.DropletSlug;
import com.pulumi.digitalocean.inputs.KubernetesNodePoolNodeArgs;
import com.pulumi.digitalocean.inputs.KubernetesNodePoolTaintArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KubernetesNodePoolState extends com.pulumi.resources.ResourceArgs {

    public static final KubernetesNodePoolState Empty = new KubernetesNodePoolState();

    /**
     * A computed field representing the actual number of nodes in the node pool, which is especially useful when auto-scaling is enabled.
     * 
     */
    @Import(name="actualNodeCount")
    private @Nullable Output<Integer> actualNodeCount;

    /**
     * @return A computed field representing the actual number of nodes in the node pool, which is especially useful when auto-scaling is enabled.
     * 
     */
    public Optional<Output<Integer>> actualNodeCount() {
        return Optional.ofNullable(this.actualNodeCount);
    }

    /**
     * Enable auto-scaling of the number of nodes in the node pool within the given min/max range.
     * 
     */
    @Import(name="autoScale")
    private @Nullable Output<Boolean> autoScale;

    /**
     * @return Enable auto-scaling of the number of nodes in the node pool within the given min/max range.
     * 
     */
    public Optional<Output<Boolean>> autoScale() {
        return Optional.ofNullable(this.autoScale);
    }

    /**
     * The ID of the Kubernetes cluster to which the node pool is associated.
     * 
     */
    @Import(name="clusterId")
    private @Nullable Output<String> clusterId;

    /**
     * @return The ID of the Kubernetes cluster to which the node pool is associated.
     * 
     */
    public Optional<Output<String>> clusterId() {
        return Optional.ofNullable(this.clusterId);
    }

    /**
     * A map of key/value pairs to apply to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/).
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return A map of key/value pairs to apply to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/).
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to.
     * 
     */
    @Import(name="maxNodes")
    private @Nullable Output<Integer> maxNodes;

    /**
     * @return If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to.
     * 
     */
    public Optional<Output<Integer>> maxNodes() {
        return Optional.ofNullable(this.maxNodes);
    }

    /**
     * If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to.
     * 
     */
    @Import(name="minNodes")
    private @Nullable Output<Integer> minNodes;

    /**
     * @return If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to.
     * 
     */
    public Optional<Output<Integer>> minNodes() {
        return Optional.ofNullable(this.minNodes);
    }

    /**
     * A name for the node pool.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A name for the node pool.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The number of Droplet instances in the node pool. If auto-scaling is enabled, this should only be set if the desired result is to explicitly reset the number of nodes to this value. If auto-scaling is enabled, and the node count is outside of the given min/max range, it will use the min nodes value.
     * 
     */
    @Import(name="nodeCount")
    private @Nullable Output<Integer> nodeCount;

    /**
     * @return The number of Droplet instances in the node pool. If auto-scaling is enabled, this should only be set if the desired result is to explicitly reset the number of nodes to this value. If auto-scaling is enabled, and the node count is outside of the given min/max range, it will use the min nodes value.
     * 
     */
    public Optional<Output<Integer>> nodeCount() {
        return Optional.ofNullable(this.nodeCount);
    }

    /**
     * A list of nodes in the pool. Each node exports the following attributes:
     * 
     */
    @Import(name="nodes")
    private @Nullable Output<List<KubernetesNodePoolNodeArgs>> nodes;

    /**
     * @return A list of nodes in the pool. Each node exports the following attributes:
     * 
     */
    public Optional<Output<List<KubernetesNodePoolNodeArgs>>> nodes() {
        return Optional.ofNullable(this.nodes);
    }

    /**
     * The slug identifier for the type of Droplet to be used as workers in the node pool.
     * 
     */
    @Import(name="size")
    private @Nullable Output<Either<String,DropletSlug>> size;

    /**
     * @return The slug identifier for the type of Droplet to be used as workers in the node pool.
     * 
     */
    public Optional<Output<Either<String,DropletSlug>>> size() {
        return Optional.ofNullable(this.size);
    }

    /**
     * A list of tag names to be applied to the Kubernetes cluster.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return A list of tag names to be applied to the Kubernetes cluster.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A list of taints applied to all nodes in the pool.
     * 
     * This resource supports customized create timeouts. The default timeout is 30 minutes.
     * 
     */
    @Import(name="taints")
    private @Nullable Output<List<KubernetesNodePoolTaintArgs>> taints;

    /**
     * @return A list of taints applied to all nodes in the pool.
     * 
     * This resource supports customized create timeouts. The default timeout is 30 minutes.
     * 
     */
    public Optional<Output<List<KubernetesNodePoolTaintArgs>>> taints() {
        return Optional.ofNullable(this.taints);
    }

    private KubernetesNodePoolState() {}

    private KubernetesNodePoolState(KubernetesNodePoolState $) {
        this.actualNodeCount = $.actualNodeCount;
        this.autoScale = $.autoScale;
        this.clusterId = $.clusterId;
        this.labels = $.labels;
        this.maxNodes = $.maxNodes;
        this.minNodes = $.minNodes;
        this.name = $.name;
        this.nodeCount = $.nodeCount;
        this.nodes = $.nodes;
        this.size = $.size;
        this.tags = $.tags;
        this.taints = $.taints;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KubernetesNodePoolState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KubernetesNodePoolState $;

        public Builder() {
            $ = new KubernetesNodePoolState();
        }

        public Builder(KubernetesNodePoolState defaults) {
            $ = new KubernetesNodePoolState(Objects.requireNonNull(defaults));
        }

        /**
         * @param actualNodeCount A computed field representing the actual number of nodes in the node pool, which is especially useful when auto-scaling is enabled.
         * 
         * @return builder
         * 
         */
        public Builder actualNodeCount(@Nullable Output<Integer> actualNodeCount) {
            $.actualNodeCount = actualNodeCount;
            return this;
        }

        /**
         * @param actualNodeCount A computed field representing the actual number of nodes in the node pool, which is especially useful when auto-scaling is enabled.
         * 
         * @return builder
         * 
         */
        public Builder actualNodeCount(Integer actualNodeCount) {
            return actualNodeCount(Output.of(actualNodeCount));
        }

        /**
         * @param autoScale Enable auto-scaling of the number of nodes in the node pool within the given min/max range.
         * 
         * @return builder
         * 
         */
        public Builder autoScale(@Nullable Output<Boolean> autoScale) {
            $.autoScale = autoScale;
            return this;
        }

        /**
         * @param autoScale Enable auto-scaling of the number of nodes in the node pool within the given min/max range.
         * 
         * @return builder
         * 
         */
        public Builder autoScale(Boolean autoScale) {
            return autoScale(Output.of(autoScale));
        }

        /**
         * @param clusterId The ID of the Kubernetes cluster to which the node pool is associated.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(@Nullable Output<String> clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param clusterId The ID of the Kubernetes cluster to which the node pool is associated.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            return clusterId(Output.of(clusterId));
        }

        /**
         * @param labels A map of key/value pairs to apply to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/).
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels A map of key/value pairs to apply to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/).
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param maxNodes If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to.
         * 
         * @return builder
         * 
         */
        public Builder maxNodes(@Nullable Output<Integer> maxNodes) {
            $.maxNodes = maxNodes;
            return this;
        }

        /**
         * @param maxNodes If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to.
         * 
         * @return builder
         * 
         */
        public Builder maxNodes(Integer maxNodes) {
            return maxNodes(Output.of(maxNodes));
        }

        /**
         * @param minNodes If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to.
         * 
         * @return builder
         * 
         */
        public Builder minNodes(@Nullable Output<Integer> minNodes) {
            $.minNodes = minNodes;
            return this;
        }

        /**
         * @param minNodes If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to.
         * 
         * @return builder
         * 
         */
        public Builder minNodes(Integer minNodes) {
            return minNodes(Output.of(minNodes));
        }

        /**
         * @param name A name for the node pool.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A name for the node pool.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param nodeCount The number of Droplet instances in the node pool. If auto-scaling is enabled, this should only be set if the desired result is to explicitly reset the number of nodes to this value. If auto-scaling is enabled, and the node count is outside of the given min/max range, it will use the min nodes value.
         * 
         * @return builder
         * 
         */
        public Builder nodeCount(@Nullable Output<Integer> nodeCount) {
            $.nodeCount = nodeCount;
            return this;
        }

        /**
         * @param nodeCount The number of Droplet instances in the node pool. If auto-scaling is enabled, this should only be set if the desired result is to explicitly reset the number of nodes to this value. If auto-scaling is enabled, and the node count is outside of the given min/max range, it will use the min nodes value.
         * 
         * @return builder
         * 
         */
        public Builder nodeCount(Integer nodeCount) {
            return nodeCount(Output.of(nodeCount));
        }

        /**
         * @param nodes A list of nodes in the pool. Each node exports the following attributes:
         * 
         * @return builder
         * 
         */
        public Builder nodes(@Nullable Output<List<KubernetesNodePoolNodeArgs>> nodes) {
            $.nodes = nodes;
            return this;
        }

        /**
         * @param nodes A list of nodes in the pool. Each node exports the following attributes:
         * 
         * @return builder
         * 
         */
        public Builder nodes(List<KubernetesNodePoolNodeArgs> nodes) {
            return nodes(Output.of(nodes));
        }

        /**
         * @param nodes A list of nodes in the pool. Each node exports the following attributes:
         * 
         * @return builder
         * 
         */
        public Builder nodes(KubernetesNodePoolNodeArgs... nodes) {
            return nodes(List.of(nodes));
        }

        /**
         * @param size The slug identifier for the type of Droplet to be used as workers in the node pool.
         * 
         * @return builder
         * 
         */
        public Builder size(@Nullable Output<Either<String,DropletSlug>> size) {
            $.size = size;
            return this;
        }

        /**
         * @param size The slug identifier for the type of Droplet to be used as workers in the node pool.
         * 
         * @return builder
         * 
         */
        public Builder size(Either<String,DropletSlug> size) {
            return size(Output.of(size));
        }

        /**
         * @param size The slug identifier for the type of Droplet to be used as workers in the node pool.
         * 
         * @return builder
         * 
         */
        public Builder size(String size) {
            return size(Either.ofLeft(size));
        }

        /**
         * @param size The slug identifier for the type of Droplet to be used as workers in the node pool.
         * 
         * @return builder
         * 
         */
        public Builder size(DropletSlug size) {
            return size(Either.ofRight(size));
        }

        /**
         * @param tags A list of tag names to be applied to the Kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A list of tag names to be applied to the Kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags A list of tag names to be applied to the Kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param taints A list of taints applied to all nodes in the pool.
         * 
         * This resource supports customized create timeouts. The default timeout is 30 minutes.
         * 
         * @return builder
         * 
         */
        public Builder taints(@Nullable Output<List<KubernetesNodePoolTaintArgs>> taints) {
            $.taints = taints;
            return this;
        }

        /**
         * @param taints A list of taints applied to all nodes in the pool.
         * 
         * This resource supports customized create timeouts. The default timeout is 30 minutes.
         * 
         * @return builder
         * 
         */
        public Builder taints(List<KubernetesNodePoolTaintArgs> taints) {
            return taints(Output.of(taints));
        }

        /**
         * @param taints A list of taints applied to all nodes in the pool.
         * 
         * This resource supports customized create timeouts. The default timeout is 30 minutes.
         * 
         * @return builder
         * 
         */
        public Builder taints(KubernetesNodePoolTaintArgs... taints) {
            return taints(List.of(taints));
        }

        public KubernetesNodePoolState build() {
            return $;
        }
    }

}
