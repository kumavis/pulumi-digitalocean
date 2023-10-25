# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetProjectResult',
    'AwaitableGetProjectResult',
    'get_project',
    'get_project_output',
]

@pulumi.output_type
class GetProjectResult:
    """
    A collection of values returned by getProject.
    """
    def __init__(__self__, created_at=None, description=None, environment=None, id=None, is_default=None, name=None, owner_id=None, owner_uuid=None, purpose=None, resources=None, updated_at=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if environment and not isinstance(environment, str):
            raise TypeError("Expected argument 'environment' to be a str")
        pulumi.set(__self__, "environment", environment)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_default and not isinstance(is_default, bool):
            raise TypeError("Expected argument 'is_default' to be a bool")
        pulumi.set(__self__, "is_default", is_default)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if owner_id and not isinstance(owner_id, int):
            raise TypeError("Expected argument 'owner_id' to be a int")
        pulumi.set(__self__, "owner_id", owner_id)
        if owner_uuid and not isinstance(owner_uuid, str):
            raise TypeError("Expected argument 'owner_uuid' to be a str")
        pulumi.set(__self__, "owner_uuid", owner_uuid)
        if purpose and not isinstance(purpose, str):
            raise TypeError("Expected argument 'purpose' to be a str")
        pulumi.set(__self__, "purpose", purpose)
        if resources and not isinstance(resources, list):
            raise TypeError("Expected argument 'resources' to be a list")
        pulumi.set(__self__, "resources", resources)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The date and time when the project was created, (ISO8601)
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the project
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def environment(self) -> str:
        """
        The environment of the project's resources. The possible values are: `Development`, `Staging`, `Production`.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> bool:
        return pulumi.get(self, "is_default")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> int:
        """
        The ID of the project owner.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="ownerUuid")
    def owner_uuid(self) -> str:
        """
        The unique universal identifier of the project owner.
        """
        return pulumi.get(self, "owner_uuid")

    @property
    @pulumi.getter
    def purpose(self) -> str:
        """
        The purpose of the project, (Default: "Web Application")
        """
        return pulumi.get(self, "purpose")

    @property
    @pulumi.getter
    def resources(self) -> Sequence[str]:
        """
        A set of uniform resource names (URNs) for the resources associated with the project
        """
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        The date and time when the project was last updated, (ISO8601)
        """
        return pulumi.get(self, "updated_at")


class AwaitableGetProjectResult(GetProjectResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectResult(
            created_at=self.created_at,
            description=self.description,
            environment=self.environment,
            id=self.id,
            is_default=self.is_default,
            name=self.name,
            owner_id=self.owner_id,
            owner_uuid=self.owner_uuid,
            purpose=self.purpose,
            resources=self.resources,
            updated_at=self.updated_at)


def get_project(id: Optional[str] = None,
                name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectResult:
    """
    Get information on a single DigitalOcean project. If neither the `id` nor `name` attributes are provided,
    then this data source returns the default project.


    :param str id: the ID of the project to retrieve
    :param str name: the name of the project to retrieve. The data source will raise an error if more than
           one project has the provided name or if no project has that name.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean:index/getProject:getProject', __args__, opts=opts, typ=GetProjectResult).value

    return AwaitableGetProjectResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        description=pulumi.get(__ret__, 'description'),
        environment=pulumi.get(__ret__, 'environment'),
        id=pulumi.get(__ret__, 'id'),
        is_default=pulumi.get(__ret__, 'is_default'),
        name=pulumi.get(__ret__, 'name'),
        owner_id=pulumi.get(__ret__, 'owner_id'),
        owner_uuid=pulumi.get(__ret__, 'owner_uuid'),
        purpose=pulumi.get(__ret__, 'purpose'),
        resources=pulumi.get(__ret__, 'resources'),
        updated_at=pulumi.get(__ret__, 'updated_at'))


@_utilities.lift_output_func(get_project)
def get_project_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                       name: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProjectResult]:
    """
    Get information on a single DigitalOcean project. If neither the `id` nor `name` attributes are provided,
    then this data source returns the default project.


    :param str id: the ID of the project to retrieve
    :param str name: the name of the project to retrieve. The data source will raise an error if more than
           one project has the provided name or if no project has that name.
    """
    ...
