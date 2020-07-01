# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetVolumeSnapshotResult:
    """
    A collection of values returned by getVolumeSnapshot.
    """
    def __init__(__self__, created_at=None, id=None, min_disk_size=None, most_recent=None, name=None, name_regex=None, region=None, regions=None, size=None, tags=None, volume_id=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        __self__.created_at = created_at
        """
        The date and time the volume snapshot was created.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if min_disk_size and not isinstance(min_disk_size, float):
            raise TypeError("Expected argument 'min_disk_size' to be a float")
        __self__.min_disk_size = min_disk_size
        """
        The minimum size in gigabytes required for a volume to be created based on this volume snapshot.
        """
        if most_recent and not isinstance(most_recent, bool):
            raise TypeError("Expected argument 'most_recent' to be a bool")
        __self__.most_recent = most_recent
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        if regions and not isinstance(regions, list):
            raise TypeError("Expected argument 'regions' to be a list")
        __self__.regions = regions
        """
        A list of DigitalOcean region "slugs" indicating where the volume snapshot is available.
        """
        if size and not isinstance(size, float):
            raise TypeError("Expected argument 'size' to be a float")
        __self__.size = size
        """
        The billable size of the volume snapshot in gigabytes.
        """
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        __self__.tags = tags
        """
        A list of the tags associated to the volume snapshot.
        """
        if volume_id and not isinstance(volume_id, str):
            raise TypeError("Expected argument 'volume_id' to be a str")
        __self__.volume_id = volume_id
        """
        The ID of the volume from which the volume snapshot originated.
        """
class AwaitableGetVolumeSnapshotResult(GetVolumeSnapshotResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVolumeSnapshotResult(
            created_at=self.created_at,
            id=self.id,
            min_disk_size=self.min_disk_size,
            most_recent=self.most_recent,
            name=self.name,
            name_regex=self.name_regex,
            region=self.region,
            regions=self.regions,
            size=self.size,
            tags=self.tags,
            volume_id=self.volume_id)

def get_volume_snapshot(most_recent=None,name=None,name_regex=None,region=None,opts=None):
    """
    Volume snapshots are saved instances of a block storage volume. Use this data
    source to retrieve the ID of a DigitalOcean volume snapshot for use in other
    resources.

    ## Example Usage

    Get the volume snapshot:

    ```python
    import pulumi
    import pulumi_digitalocean as digitalocean

    snapshot = digitalocean.get_volume_snapshot(most_recent=True,
        name_regex="^web",
        region="nyc3")
    ```

    Reuse the data about a volume snapshot to create a new volume based on it:

    ```python
    import pulumi
    import pulumi_digitalocean as digitalocean

    snapshot = digitalocean.get_volume_snapshot(name_regex="^web",
        region="nyc3",
        most_recent=True)
    foobar = digitalocean.Volume("foobar",
        region="nyc3",
        size=100,
        snapshot_id=snapshot.id)
    ```


    :param bool most_recent: If more than one result is returned, use the most recent volume snapshot.
    :param str name: The name of the volume snapshot.
    :param str name_regex: A regex string to apply to the volume snapshot list returned by DigitalOcean. This allows more advanced filtering not supported from the DigitalOcean API. This filtering is done locally on what DigitalOcean returns.
    :param str region: A "slug" representing a DigitalOcean region (e.g. `nyc1`). If set, only volume snapshots available in the region will be returned.
    """
    __args__ = dict()


    __args__['mostRecent'] = most_recent
    __args__['name'] = name
    __args__['nameRegex'] = name_regex
    __args__['region'] = region
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('digitalocean:index/getVolumeSnapshot:getVolumeSnapshot', __args__, opts=opts).value

    return AwaitableGetVolumeSnapshotResult(
        created_at=__ret__.get('createdAt'),
        id=__ret__.get('id'),
        min_disk_size=__ret__.get('minDiskSize'),
        most_recent=__ret__.get('mostRecent'),
        name=__ret__.get('name'),
        name_regex=__ret__.get('nameRegex'),
        region=__ret__.get('region'),
        regions=__ret__.get('regions'),
        size=__ret__.get('size'),
        tags=__ret__.get('tags'),
        volume_id=__ret__.get('volumeId'))
