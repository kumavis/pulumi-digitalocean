// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";

export interface DatabaseClusterMaintenanceWindow {
    /**
     * The day of the week on which to apply maintenance updates.
     */
    day: string;
    /**
     * The hour in UTC at which maintenance updates will be applied in 24 hour format.
     */
    hour: string;
}

export interface FirewallInboundRule {
    /**
     * The ports on which traffic will be allowed
     * specified as a string containing a single port, a range (e.g. "8000-9000"),
     * or "1-65535" to open all ports for a protocol. Required for when protocol is
     * `tcp` or `udp`.
     */
    portRange?: string;
    /**
     * The type of traffic to be allowed.
     * This may be one of "tcp", "udp", or "icmp".
     */
    protocol: string;
    /**
     * An array of strings containing the IPv4
     * addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs from which the
     * inbound traffic will be accepted.
     */
    sourceAddresses?: string[];
    /**
     * An array containing the IDs of
     * the Droplets from which the inbound traffic will be accepted.
     */
    sourceDropletIds?: number[];
    /**
     * An array containing the IDs
     * of the Load Balancers from which the inbound traffic will be accepted.
     */
    sourceLoadBalancerUids?: string[];
    /**
     * An array containing the names of Tags
     * corresponding to groups of Droplets from which the inbound traffic
     * will be accepted.
     */
    sourceTags?: string[];
}

export interface FirewallOutboundRule {
    /**
     * An array of strings containing the IPv4
     * addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs to which the
     * outbound traffic will be allowed.
     */
    destinationAddresses?: string[];
    /**
     * An array containing the IDs of
     * the Droplets to which the outbound traffic will be allowed.
     */
    destinationDropletIds?: number[];
    /**
     * An array containing the IDs
     * of the Load Balancers to which the outbound traffic will be allowed.
     */
    destinationLoadBalancerUids?: string[];
    /**
     * An array containing the names of Tags
     * corresponding to groups of Droplets to which the outbound traffic will
     * be allowed.
     * traffic.
     */
    destinationTags?: string[];
    /**
     * The ports on which traffic will be allowed
     * specified as a string containing a single port, a range (e.g. "8000-9000"),
     * or "1-65535" to open all ports for a protocol. Required for when protocol is
     * `tcp` or `udp`.
     */
    portRange?: string;
    /**
     * The type of traffic to be allowed.
     * This may be one of "tcp", "udp", or "icmp".
     */
    protocol: string;
}

export interface FirewallPendingChange {
    dropletId?: number;
    removing?: boolean;
    /**
     * A status string indicating the current state of the Firewall.
     * This can be "waiting", "succeeded", or "failed".
     */
    status?: string;
}

export interface GetDatabaseClusterMaintenanceWindow {
    /**
     * The day of the week on which to apply maintenance updates.
     */
    day: string;
    /**
     * The hour in UTC at which maintenance updates will be applied in 24 hour format.
     */
    hour: string;
}

export interface GetKubernetesClusterKubeConfig {
    clientCertificate: string;
    clientKey: string;
    clusterCaCertificate: string;
    expiresAt: string;
    host: string;
    rawConfig: string;
    token: string;
}

export interface GetKubernetesClusterNodePool {
    actualNodeCount: number;
    autoScale: boolean;
    /**
     * The unique ID that can be used to identify and reference a Kubernetes cluster.
     */
    id: string;
    maxNodes: number;
    minNodes: number;
    /**
     * The name of Kubernetes cluster.
     */
    name: string;
    nodeCount: number;
    nodes: outputs.GetKubernetesClusterNodePoolNode[];
    size: string;
    /**
     * A list of tag names to be applied to the Kubernetes cluster.
     */
    tags?: string[];
}

export interface GetKubernetesClusterNodePoolNode {
    /**
     * The date and time when the Kubernetes cluster was created.
     */
    createdAt: string;
    /**
     * The unique ID that can be used to identify and reference a Kubernetes cluster.
     */
    id: string;
    /**
     * The name of Kubernetes cluster.
     */
    name: string;
    /**
     * A string indicating the current status of the cluster. Potential values include running, provisioning, and errored.
     */
    status: string;
    /**
     * The date and time when the Kubernetes cluster was last updated.
     * * `kube_config.0` - A representation of the Kubernetes cluster's kubeconfig with the following attributes:
     * - `rawConfig` - The full contents of the Kubernetes cluster's kubeconfig file.
     * - `host` - The URL of the API server on the Kubernetes master node.
     * - `clusterCaCertificate` - The base64 encoded public certificate for the cluster's certificate authority.
     * - `token` - The DigitalOcean API access token used by clients to access the cluster.
     * - `clientKey` - The base64 encoded private key used by clients to access the cluster. Only available if token authentication is not supported on your cluster.
     * - `clientCertificate` - The base64 encoded public certificate used by clients to access the cluster. Only available if token authentication is not supported on your cluster.
     * - `expiresAt` - The date and time when the credentials will expire and need to be regenerated.
     */
    updatedAt: string;
}

export interface GetLoadBalancerForwardingRule {
    certificateId: string;
    entryPort: number;
    entryProtocol: string;
    targetPort: number;
    targetProtocol: string;
    tlsPassthrough: boolean;
}

export interface GetLoadBalancerHealthcheck {
    checkIntervalSeconds: number;
    healthyThreshold: number;
    path: string;
    port: number;
    protocol: string;
    responseTimeoutSeconds: number;
    unhealthyThreshold: number;
}

export interface GetLoadBalancerStickySessions {
    cookieName: string;
    cookieTtlSeconds: number;
    type: string;
}

export interface GetSizesFilter {
    /**
     * Sort the sizes by this key. This may be one of `slug`,
     * `memory`, `vcpus`, `disk`, `transfer`, `priceMonthly`, or `priceHourly`.
     */
    key: string;
    /**
     * Only retrieves images which keys has value that matches
     * one of the values provided here.
     */
    values: string[];
}

export interface GetSizesSize {
    /**
     * This represents whether new Droplets can be created with this size.
     */
    available: boolean;
    /**
     * The amount of disk space set aside for Droplets of this size. The value is measured in gigabytes.
     */
    disk: number;
    /**
     * The amount of RAM allocated to Droplets created of this size. The value is measured in megabytes.
     */
    memory: number;
    /**
     * The hourly cost of Droplets created in this size as measured hourly. The value is measured in US dollars.
     */
    priceHourly: number;
    /**
     * The monthly cost of Droplets created in this size if they are kept for an entire month. The value is measured in US dollars.
     */
    priceMonthly: number;
    /**
     * List of region slugs where Droplets can be created in this size.
     */
    regions: string[];
    /**
     * A human-readable string that is used to uniquely identify each size.
     */
    slug: string;
    /**
     * The amount of transfer bandwidth that is available for Droplets created in this size. This only counts traffic on the public interface. The value is given in terabytes.
     */
    transfer: number;
    /**
     * The number of CPUs allocated to Droplets of this size.
     */
    vcpus: number;
}

export interface GetSizesSort {
    /**
     * The sort direction. This may be either `asc` or `desc`.
     */
    direction?: string;
    /**
     * Sort the sizes by this key. This may be one of `slug`,
     * `memory`, `vcpus`, `disk`, `transfer`, `priceMonthly`, or `priceHourly`.
     */
    key: string;
}

export interface KubernetesClusterKubeConfig {
    clientCertificate: string;
    clientKey: string;
    clusterCaCertificate: string;
    expiresAt: string;
    host: string;
    rawConfig: string;
    token: string;
}

export interface KubernetesClusterNodePool {
    actualNodeCount: number;
    autoScale?: boolean;
    /**
     * A unique ID that can be used to identify and reference a Kubernetes cluster.
     */
    id: string;
    maxNodes?: number;
    minNodes?: number;
    /**
     * A name for the Kubernetes cluster.
     */
    name: string;
    nodeCount?: number;
    nodes: outputs.KubernetesClusterNodePoolNode[];
    size: string;
    /**
     * A list of tag names to be applied to the Kubernetes cluster.
     */
    tags?: string[];
}

export interface KubernetesClusterNodePoolNode {
    /**
     * The date and time when the Kubernetes cluster was created.
     */
    createdAt: string;
    /**
     * A unique ID that can be used to identify and reference a Kubernetes cluster.
     */
    id: string;
    /**
     * A name for the Kubernetes cluster.
     */
    name: string;
    /**
     * A string indicating the current status of the cluster. Potential values include running, provisioning, and errored.
     */
    status: string;
    /**
     * The date and time when the Kubernetes cluster was last updated.
     * * `kube_config.0` - A representation of the Kubernetes cluster's kubeconfig with the following attributes:
     * - `rawConfig` - The full contents of the Kubernetes cluster's kubeconfig file.
     * - `host` - The URL of the API server on the Kubernetes master node.
     * - `clusterCaCertificate` - The base64 encoded public certificate for the cluster's certificate authority.
     * - `token` - The DigitalOcean API access token used by clients to access the cluster.
     * - `clientKey` - The base64 encoded private key used by clients to access the cluster. Only available if token authentication is not supported on your cluster.
     * - `clientCertificate` - The base64 encoded public certificate used by clients to access the cluster. Only available if token authentication is not supported on your cluster.
     * - `expiresAt` - The date and time when the credentials will expire and need to be regenerated.
     */
    updatedAt: string;
}

export interface KubernetesNodePoolNode {
    createdAt: string;
    /**
     * A unique ID that can be used to identify and reference the node pool.
     */
    id: string;
    /**
     * A name for the node pool.
     */
    name: string;
    status: string;
    updatedAt: string;
}

export interface LoadBalancerForwardingRule {
    /**
     * The ID of the TLS certificate to be used for SSL termination.
     */
    certificateId?: string;
    /**
     * An integer representing the port on which the Load Balancer instance will listen.
     */
    entryPort: number;
    /**
     * The protocol used for traffic to the Load Balancer. The possible values are: `http`, `https`, `http2` or `tcp`.
     */
    entryProtocol: string;
    /**
     * An integer representing the port on the backend Droplets to which the Load Balancer will send traffic.
     */
    targetPort: number;
    /**
     * The protocol used for traffic from the Load Balancer to the backend Droplets. The possible values are: `http`, `https`, `http2` or `tcp`.
     */
    targetProtocol: string;
    /**
     * A boolean value indicating whether SSL encrypted traffic will be passed through to the backend Droplets. The default value is `false`.
     */
    tlsPassthrough?: boolean;
}

export interface LoadBalancerHealthcheck {
    /**
     * The number of seconds between between two consecutive health checks. If not specified, the default value is `10`.
     */
    checkIntervalSeconds?: number;
    /**
     * The number of times a health check must pass for a backend Droplet to be marked "healthy" and be re-added to the pool. If not specified, the default value is `5`.
     */
    healthyThreshold?: number;
    /**
     * The path on the backend Droplets to which the Load Balancer instance will send a request.
     */
    path?: string;
    /**
     * An integer representing the port on the backend Droplets on which the health check will attempt a connection.
     */
    port: number;
    /**
     * The protocol used for health checks sent to the backend Droplets. The possible values are `http` or `tcp`.
     */
    protocol: string;
    /**
     * The number of seconds the Load Balancer instance will wait for a response until marking a health check as failed. If not specified, the default value is `5`.
     */
    responseTimeoutSeconds?: number;
    /**
     * The number of times a health check must fail for a backend Droplet to be marked "unhealthy" and be removed from the pool. If not specified, the default value is `3`.
     */
    unhealthyThreshold?: number;
}

export interface LoadBalancerStickySessions {
    /**
     * The name to be used for the cookie sent to the client. This attribute is required when using `cookies` for the sticky sessions type.
     */
    cookieName?: string;
    /**
     * The number of seconds until the cookie set by the Load Balancer expires. This attribute is required when using `cookies` for the sticky sessions type.
     */
    cookieTtlSeconds?: number;
    /**
     * An attribute indicating how and if requests from a client will be persistently served by the same backend Droplet. The possible values are `cookies` or `none`. If not specified, the default value is `none`.
     */
    type?: string;
}

export interface SpacesBucketCorsRule {
    /**
     * A list of headers that will be included in the CORS preflight request's `Access-Control-Request-Headers`. A header may contain one wildcard (e.g. `x-amz-*`).
     */
    allowedHeaders?: string[];
    /**
     * A list of HTTP methods (e.g. `GET`) which are allowed from the specified origin.
     */
    allowedMethods: string[];
    /**
     * A list of hosts from which requests using the specified methods are allowed. A host may contain one wildcard (e.g. http://*.example.com).
     */
    allowedOrigins: string[];
    /**
     * The time in seconds that browser can cache the response for a preflight request.
     */
    maxAgeSeconds?: number;
}
