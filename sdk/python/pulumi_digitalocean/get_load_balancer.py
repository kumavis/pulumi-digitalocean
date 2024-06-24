# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetLoadBalancerResult',
    'AwaitableGetLoadBalancerResult',
    'get_load_balancer',
    'get_load_balancer_output',
]

@pulumi.output_type
class GetLoadBalancerResult:
    """
    A collection of values returned by getLoadBalancer.
    """
    def __init__(__self__, algorithm=None, disable_lets_encrypt_dns_records=None, domains=None, droplet_ids=None, droplet_tag=None, enable_backend_keepalive=None, enable_proxy_protocol=None, firewalls=None, forwarding_rules=None, glb_settings=None, healthchecks=None, http_idle_timeout_seconds=None, id=None, ip=None, load_balancer_urn=None, name=None, project_id=None, redirect_http_to_https=None, region=None, size=None, size_unit=None, status=None, sticky_sessions=None, target_load_balancer_ids=None, type=None, vpc_uuid=None):
        if algorithm and not isinstance(algorithm, str):
            raise TypeError("Expected argument 'algorithm' to be a str")
        pulumi.set(__self__, "algorithm", algorithm)
        if disable_lets_encrypt_dns_records and not isinstance(disable_lets_encrypt_dns_records, bool):
            raise TypeError("Expected argument 'disable_lets_encrypt_dns_records' to be a bool")
        pulumi.set(__self__, "disable_lets_encrypt_dns_records", disable_lets_encrypt_dns_records)
        if domains and not isinstance(domains, list):
            raise TypeError("Expected argument 'domains' to be a list")
        pulumi.set(__self__, "domains", domains)
        if droplet_ids and not isinstance(droplet_ids, list):
            raise TypeError("Expected argument 'droplet_ids' to be a list")
        pulumi.set(__self__, "droplet_ids", droplet_ids)
        if droplet_tag and not isinstance(droplet_tag, str):
            raise TypeError("Expected argument 'droplet_tag' to be a str")
        pulumi.set(__self__, "droplet_tag", droplet_tag)
        if enable_backend_keepalive and not isinstance(enable_backend_keepalive, bool):
            raise TypeError("Expected argument 'enable_backend_keepalive' to be a bool")
        pulumi.set(__self__, "enable_backend_keepalive", enable_backend_keepalive)
        if enable_proxy_protocol and not isinstance(enable_proxy_protocol, bool):
            raise TypeError("Expected argument 'enable_proxy_protocol' to be a bool")
        pulumi.set(__self__, "enable_proxy_protocol", enable_proxy_protocol)
        if firewalls and not isinstance(firewalls, list):
            raise TypeError("Expected argument 'firewalls' to be a list")
        pulumi.set(__self__, "firewalls", firewalls)
        if forwarding_rules and not isinstance(forwarding_rules, list):
            raise TypeError("Expected argument 'forwarding_rules' to be a list")
        pulumi.set(__self__, "forwarding_rules", forwarding_rules)
        if glb_settings and not isinstance(glb_settings, list):
            raise TypeError("Expected argument 'glb_settings' to be a list")
        pulumi.set(__self__, "glb_settings", glb_settings)
        if healthchecks and not isinstance(healthchecks, list):
            raise TypeError("Expected argument 'healthchecks' to be a list")
        pulumi.set(__self__, "healthchecks", healthchecks)
        if http_idle_timeout_seconds and not isinstance(http_idle_timeout_seconds, int):
            raise TypeError("Expected argument 'http_idle_timeout_seconds' to be a int")
        pulumi.set(__self__, "http_idle_timeout_seconds", http_idle_timeout_seconds)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip and not isinstance(ip, str):
            raise TypeError("Expected argument 'ip' to be a str")
        pulumi.set(__self__, "ip", ip)
        if load_balancer_urn and not isinstance(load_balancer_urn, str):
            raise TypeError("Expected argument 'load_balancer_urn' to be a str")
        pulumi.set(__self__, "load_balancer_urn", load_balancer_urn)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if redirect_http_to_https and not isinstance(redirect_http_to_https, bool):
            raise TypeError("Expected argument 'redirect_http_to_https' to be a bool")
        pulumi.set(__self__, "redirect_http_to_https", redirect_http_to_https)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if size and not isinstance(size, str):
            raise TypeError("Expected argument 'size' to be a str")
        pulumi.set(__self__, "size", size)
        if size_unit and not isinstance(size_unit, int):
            raise TypeError("Expected argument 'size_unit' to be a int")
        pulumi.set(__self__, "size_unit", size_unit)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if sticky_sessions and not isinstance(sticky_sessions, list):
            raise TypeError("Expected argument 'sticky_sessions' to be a list")
        pulumi.set(__self__, "sticky_sessions", sticky_sessions)
        if target_load_balancer_ids and not isinstance(target_load_balancer_ids, list):
            raise TypeError("Expected argument 'target_load_balancer_ids' to be a list")
        pulumi.set(__self__, "target_load_balancer_ids", target_load_balancer_ids)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if vpc_uuid and not isinstance(vpc_uuid, str):
            raise TypeError("Expected argument 'vpc_uuid' to be a str")
        pulumi.set(__self__, "vpc_uuid", vpc_uuid)

    @property
    @pulumi.getter
    @_utilities.deprecated("""This field has been deprecated. You can no longer specify an algorithm for load balancers.""")
    def algorithm(self) -> str:
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter(name="disableLetsEncryptDnsRecords")
    def disable_lets_encrypt_dns_records(self) -> bool:
        return pulumi.get(self, "disable_lets_encrypt_dns_records")

    @property
    @pulumi.getter
    def domains(self) -> Sequence['outputs.GetLoadBalancerDomainResult']:
        return pulumi.get(self, "domains")

    @property
    @pulumi.getter(name="dropletIds")
    def droplet_ids(self) -> Sequence[int]:
        return pulumi.get(self, "droplet_ids")

    @property
    @pulumi.getter(name="dropletTag")
    def droplet_tag(self) -> str:
        return pulumi.get(self, "droplet_tag")

    @property
    @pulumi.getter(name="enableBackendKeepalive")
    def enable_backend_keepalive(self) -> bool:
        return pulumi.get(self, "enable_backend_keepalive")

    @property
    @pulumi.getter(name="enableProxyProtocol")
    def enable_proxy_protocol(self) -> bool:
        return pulumi.get(self, "enable_proxy_protocol")

    @property
    @pulumi.getter
    def firewalls(self) -> Sequence['outputs.GetLoadBalancerFirewallResult']:
        return pulumi.get(self, "firewalls")

    @property
    @pulumi.getter(name="forwardingRules")
    def forwarding_rules(self) -> Sequence['outputs.GetLoadBalancerForwardingRuleResult']:
        return pulumi.get(self, "forwarding_rules")

    @property
    @pulumi.getter(name="glbSettings")
    def glb_settings(self) -> Sequence['outputs.GetLoadBalancerGlbSettingResult']:
        return pulumi.get(self, "glb_settings")

    @property
    @pulumi.getter
    def healthchecks(self) -> Sequence['outputs.GetLoadBalancerHealthcheckResult']:
        return pulumi.get(self, "healthchecks")

    @property
    @pulumi.getter(name="httpIdleTimeoutSeconds")
    def http_idle_timeout_seconds(self) -> int:
        return pulumi.get(self, "http_idle_timeout_seconds")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip(self) -> str:
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter(name="loadBalancerUrn")
    def load_balancer_urn(self) -> str:
        return pulumi.get(self, "load_balancer_urn")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="redirectHttpToHttps")
    def redirect_http_to_https(self) -> bool:
        return pulumi.get(self, "redirect_http_to_https")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def size(self) -> str:
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="sizeUnit")
    def size_unit(self) -> int:
        return pulumi.get(self, "size_unit")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="stickySessions")
    def sticky_sessions(self) -> Sequence['outputs.GetLoadBalancerStickySessionResult']:
        return pulumi.get(self, "sticky_sessions")

    @property
    @pulumi.getter(name="targetLoadBalancerIds")
    def target_load_balancer_ids(self) -> Sequence[str]:
        return pulumi.get(self, "target_load_balancer_ids")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vpcUuid")
    def vpc_uuid(self) -> str:
        return pulumi.get(self, "vpc_uuid")


class AwaitableGetLoadBalancerResult(GetLoadBalancerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLoadBalancerResult(
            algorithm=self.algorithm,
            disable_lets_encrypt_dns_records=self.disable_lets_encrypt_dns_records,
            domains=self.domains,
            droplet_ids=self.droplet_ids,
            droplet_tag=self.droplet_tag,
            enable_backend_keepalive=self.enable_backend_keepalive,
            enable_proxy_protocol=self.enable_proxy_protocol,
            firewalls=self.firewalls,
            forwarding_rules=self.forwarding_rules,
            glb_settings=self.glb_settings,
            healthchecks=self.healthchecks,
            http_idle_timeout_seconds=self.http_idle_timeout_seconds,
            id=self.id,
            ip=self.ip,
            load_balancer_urn=self.load_balancer_urn,
            name=self.name,
            project_id=self.project_id,
            redirect_http_to_https=self.redirect_http_to_https,
            region=self.region,
            size=self.size,
            size_unit=self.size_unit,
            status=self.status,
            sticky_sessions=self.sticky_sessions,
            target_load_balancer_ids=self.target_load_balancer_ids,
            type=self.type,
            vpc_uuid=self.vpc_uuid)


def get_load_balancer(id: Optional[str] = None,
                      name: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLoadBalancerResult:
    """
    Get information on a load balancer for use in other resources. This data source
    provides all of the load balancers properties as configured on your DigitalOcean
    account. This is useful if the load balancer in question is not managed by
    the provider or you need to utilize any of the load balancers data.

    An error is triggered if the provided load balancer name does not exist.

    ## Example Usage

    Get the load balancer by name:

    ```python
    import pulumi
    import pulumi_digitalocean as digitalocean

    example = digitalocean.get_load_balancer(name="app")
    pulumi.export("lbOutput", example.ip)
    ```

    Get the load balancer by ID:

    ```python
    import pulumi
    import pulumi_digitalocean as digitalocean

    example = digitalocean.get_load_balancer(id="loadbalancer_id")
    ```


    :param str id: The ID of load balancer.
    :param str name: The name of load balancer.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean:index/getLoadBalancer:getLoadBalancer', __args__, opts=opts, typ=GetLoadBalancerResult).value

    return AwaitableGetLoadBalancerResult(
        algorithm=pulumi.get(__ret__, 'algorithm'),
        disable_lets_encrypt_dns_records=pulumi.get(__ret__, 'disable_lets_encrypt_dns_records'),
        domains=pulumi.get(__ret__, 'domains'),
        droplet_ids=pulumi.get(__ret__, 'droplet_ids'),
        droplet_tag=pulumi.get(__ret__, 'droplet_tag'),
        enable_backend_keepalive=pulumi.get(__ret__, 'enable_backend_keepalive'),
        enable_proxy_protocol=pulumi.get(__ret__, 'enable_proxy_protocol'),
        firewalls=pulumi.get(__ret__, 'firewalls'),
        forwarding_rules=pulumi.get(__ret__, 'forwarding_rules'),
        glb_settings=pulumi.get(__ret__, 'glb_settings'),
        healthchecks=pulumi.get(__ret__, 'healthchecks'),
        http_idle_timeout_seconds=pulumi.get(__ret__, 'http_idle_timeout_seconds'),
        id=pulumi.get(__ret__, 'id'),
        ip=pulumi.get(__ret__, 'ip'),
        load_balancer_urn=pulumi.get(__ret__, 'load_balancer_urn'),
        name=pulumi.get(__ret__, 'name'),
        project_id=pulumi.get(__ret__, 'project_id'),
        redirect_http_to_https=pulumi.get(__ret__, 'redirect_http_to_https'),
        region=pulumi.get(__ret__, 'region'),
        size=pulumi.get(__ret__, 'size'),
        size_unit=pulumi.get(__ret__, 'size_unit'),
        status=pulumi.get(__ret__, 'status'),
        sticky_sessions=pulumi.get(__ret__, 'sticky_sessions'),
        target_load_balancer_ids=pulumi.get(__ret__, 'target_load_balancer_ids'),
        type=pulumi.get(__ret__, 'type'),
        vpc_uuid=pulumi.get(__ret__, 'vpc_uuid'))


@_utilities.lift_output_func(get_load_balancer)
def get_load_balancer_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                             name: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLoadBalancerResult]:
    """
    Get information on a load balancer for use in other resources. This data source
    provides all of the load balancers properties as configured on your DigitalOcean
    account. This is useful if the load balancer in question is not managed by
    the provider or you need to utilize any of the load balancers data.

    An error is triggered if the provided load balancer name does not exist.

    ## Example Usage

    Get the load balancer by name:

    ```python
    import pulumi
    import pulumi_digitalocean as digitalocean

    example = digitalocean.get_load_balancer(name="app")
    pulumi.export("lbOutput", example.ip)
    ```

    Get the load balancer by ID:

    ```python
    import pulumi
    import pulumi_digitalocean as digitalocean

    example = digitalocean.get_load_balancer(id="loadbalancer_id")
    ```


    :param str id: The ID of load balancer.
    :param str name: The name of load balancer.
    """
    ...
