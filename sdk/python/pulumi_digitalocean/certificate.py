# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class Certificate(pulumi.CustomResource):
    certificate_chain: pulumi.Output[str]
    """
    The full PEM-formatted trust chain
    between the certificate authority's certificate and your domain's TLS
    certificate. Only valid when type is `custom`.
    """
    domains: pulumi.Output[list]
    """
    List of fully qualified domain names (FQDNs) for
    which the certificate will be issued. The domains must be managed using
    DigitalOcean's DNS. Only valid when type is `lets_encrypt`.
    """
    leaf_certificate: pulumi.Output[str]
    """
    The contents of a PEM-formatted public
    TLS certificate. Only valid when type is `custom`.
    """
    name: pulumi.Output[str]
    """
    The name of the certificate for identification.
    """
    not_after: pulumi.Output[str]
    """
    The expiration date of the certificate
    """
    private_key: pulumi.Output[str]
    """
    The contents of a PEM-formatted private-key
    corresponding to the SSL certificate. Only valid when type is `custom`.
    """
    sha1_fingerprint: pulumi.Output[str]
    """
    The SHA-1 fingerprint of the certificate
    """
    state: pulumi.Output[str]
    type: pulumi.Output[str]
    """
    The type of certificate to provision. Can be either
    `custom` or `lets_encrypt`. Defaults to `custom`.
    """
    def __init__(__self__, resource_name, opts=None, certificate_chain=None, domains=None, leaf_certificate=None, name=None, private_key=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a DigitalOcean Certificate resource that allows you to manage
        certificates for configuring TLS termination in Load Balancers.
        Certificates created with this resource can be referenced in your
        Load Balancer configuration via their ID. The certificate can either
        be a custom one provided by you or automatically generated one with
        Let's Encrypt.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_chain: The full PEM-formatted trust chain
               between the certificate authority's certificate and your domain's TLS
               certificate. Only valid when type is `custom`.
        :param pulumi.Input[list] domains: List of fully qualified domain names (FQDNs) for
               which the certificate will be issued. The domains must be managed using
               DigitalOcean's DNS. Only valid when type is `lets_encrypt`.
        :param pulumi.Input[str] leaf_certificate: The contents of a PEM-formatted public
               TLS certificate. Only valid when type is `custom`.
        :param pulumi.Input[str] name: The name of the certificate for identification.
        :param pulumi.Input[str] private_key: The contents of a PEM-formatted private-key
               corresponding to the SSL certificate. Only valid when type is `custom`.
        :param pulumi.Input[str] type: The type of certificate to provision. Can be either
               `custom` or `lets_encrypt`. Defaults to `custom`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/r/certificate.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['certificate_chain'] = certificate_chain
            __props__['domains'] = domains
            __props__['leaf_certificate'] = leaf_certificate
            __props__['name'] = name
            __props__['private_key'] = private_key
            __props__['type'] = type
            __props__['not_after'] = None
            __props__['sha1_fingerprint'] = None
            __props__['state'] = None
        super(Certificate, __self__).__init__(
            'digitalocean:index/certificate:Certificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, certificate_chain=None, domains=None, leaf_certificate=None, name=None, not_after=None, private_key=None, sha1_fingerprint=None, state=None, type=None):
        """
        Get an existing Certificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_chain: The full PEM-formatted trust chain
               between the certificate authority's certificate and your domain's TLS
               certificate. Only valid when type is `custom`.
        :param pulumi.Input[list] domains: List of fully qualified domain names (FQDNs) for
               which the certificate will be issued. The domains must be managed using
               DigitalOcean's DNS. Only valid when type is `lets_encrypt`.
        :param pulumi.Input[str] leaf_certificate: The contents of a PEM-formatted public
               TLS certificate. Only valid when type is `custom`.
        :param pulumi.Input[str] name: The name of the certificate for identification.
        :param pulumi.Input[str] not_after: The expiration date of the certificate
        :param pulumi.Input[str] private_key: The contents of a PEM-formatted private-key
               corresponding to the SSL certificate. Only valid when type is `custom`.
        :param pulumi.Input[str] sha1_fingerprint: The SHA-1 fingerprint of the certificate
        :param pulumi.Input[str] type: The type of certificate to provision. Can be either
               `custom` or `lets_encrypt`. Defaults to `custom`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/r/certificate.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["certificate_chain"] = certificate_chain
        __props__["domains"] = domains
        __props__["leaf_certificate"] = leaf_certificate
        __props__["name"] = name
        __props__["not_after"] = not_after
        __props__["private_key"] = private_key
        __props__["sha1_fingerprint"] = sha1_fingerprint
        __props__["state"] = state
        __props__["type"] = type
        return Certificate(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

