# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ClusterArgs', 'Cluster']

@pulumi.input_type
class ClusterArgs:
    def __init__(__self__, *,
                 color: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 emoji: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Cluster resource.
        :param pulumi.Input[str] color: A color to associate with the Cluster. Perhaps a team related color, or one related to an environment. This is set using hex value, such as `#BADA55`.
        :param pulumi.Input[str] description: This is a description for the cluster, this may describe the usage for it, the region, or something else which would help identify the Cluster's purpose.
        :param pulumi.Input[str] emoji: An emoji to use with the Cluster, this can either be set using `:buildkite:` notation, or with the emoji itself, such as 😎.
        :param pulumi.Input[str] name: This is the name of the cluster.
        """
        if color is not None:
            pulumi.set(__self__, "color", color)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if emoji is not None:
            pulumi.set(__self__, "emoji", emoji)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[str]]:
        """
        A color to associate with the Cluster. Perhaps a team related color, or one related to an environment. This is set using hex value, such as `#BADA55`.
        """
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        This is a description for the cluster, this may describe the usage for it, the region, or something else which would help identify the Cluster's purpose.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def emoji(self) -> Optional[pulumi.Input[str]]:
        """
        An emoji to use with the Cluster, this can either be set using `:buildkite:` notation, or with the emoji itself, such as 😎.
        """
        return pulumi.get(self, "emoji")

    @emoji.setter
    def emoji(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "emoji", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        This is the name of the cluster.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ClusterState:
    def __init__(__self__, *,
                 color: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 emoji: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Cluster resources.
        :param pulumi.Input[str] color: A color to associate with the Cluster. Perhaps a team related color, or one related to an environment. This is set using hex value, such as `#BADA55`.
        :param pulumi.Input[str] description: This is a description for the cluster, this may describe the usage for it, the region, or something else which would help identify the Cluster's purpose.
        :param pulumi.Input[str] emoji: An emoji to use with the Cluster, this can either be set using `:buildkite:` notation, or with the emoji itself, such as 😎.
        :param pulumi.Input[str] name: This is the name of the cluster.
        :param pulumi.Input[str] uuid: The UUID created with the Cluster.
        """
        if color is not None:
            pulumi.set(__self__, "color", color)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if emoji is not None:
            pulumi.set(__self__, "emoji", emoji)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[str]]:
        """
        A color to associate with the Cluster. Perhaps a team related color, or one related to an environment. This is set using hex value, such as `#BADA55`.
        """
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        This is a description for the cluster, this may describe the usage for it, the region, or something else which would help identify the Cluster's purpose.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def emoji(self) -> Optional[pulumi.Input[str]]:
        """
        An emoji to use with the Cluster, this can either be set using `:buildkite:` notation, or with the emoji itself, such as 😎.
        """
        return pulumi.get(self, "emoji")

    @emoji.setter
    def emoji(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "emoji", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        This is the name of the cluster.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[str]]:
        """
        The UUID created with the Cluster.
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uuid", value)


class Cluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 color: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 emoji: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # Resource: cluster

        This resource allows you to create, manage and import Clusters.

        Buildkite documentation: https://buildkite.com/docs/clusters/overview

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_buildkite as buildkite

        linux = buildkite.cluster.Cluster("linux")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] color: A color to associate with the Cluster. Perhaps a team related color, or one related to an environment. This is set using hex value, such as `#BADA55`.
        :param pulumi.Input[str] description: This is a description for the cluster, this may describe the usage for it, the region, or something else which would help identify the Cluster's purpose.
        :param pulumi.Input[str] emoji: An emoji to use with the Cluster, this can either be set using `:buildkite:` notation, or with the emoji itself, such as 😎.
        :param pulumi.Input[str] name: This is the name of the cluster.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ClusterArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Resource: cluster

        This resource allows you to create, manage and import Clusters.

        Buildkite documentation: https://buildkite.com/docs/clusters/overview

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_buildkite as buildkite

        linux = buildkite.cluster.Cluster("linux")
        ```

        :param str resource_name: The name of the resource.
        :param ClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 color: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 emoji: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClusterArgs.__new__(ClusterArgs)

            __props__.__dict__["color"] = color
            __props__.__dict__["description"] = description
            __props__.__dict__["emoji"] = emoji
            __props__.__dict__["name"] = name
            __props__.__dict__["uuid"] = None
        super(Cluster, __self__).__init__(
            'buildkite:Cluster/cluster:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            color: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            emoji: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            uuid: Optional[pulumi.Input[str]] = None) -> 'Cluster':
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] color: A color to associate with the Cluster. Perhaps a team related color, or one related to an environment. This is set using hex value, such as `#BADA55`.
        :param pulumi.Input[str] description: This is a description for the cluster, this may describe the usage for it, the region, or something else which would help identify the Cluster's purpose.
        :param pulumi.Input[str] emoji: An emoji to use with the Cluster, this can either be set using `:buildkite:` notation, or with the emoji itself, such as 😎.
        :param pulumi.Input[str] name: This is the name of the cluster.
        :param pulumi.Input[str] uuid: The UUID created with the Cluster.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClusterState.__new__(_ClusterState)

        __props__.__dict__["color"] = color
        __props__.__dict__["description"] = description
        __props__.__dict__["emoji"] = emoji
        __props__.__dict__["name"] = name
        __props__.__dict__["uuid"] = uuid
        return Cluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def color(self) -> pulumi.Output[Optional[str]]:
        """
        A color to associate with the Cluster. Perhaps a team related color, or one related to an environment. This is set using hex value, such as `#BADA55`.
        """
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        This is a description for the cluster, this may describe the usage for it, the region, or something else which would help identify the Cluster's purpose.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def emoji(self) -> pulumi.Output[Optional[str]]:
        """
        An emoji to use with the Cluster, this can either be set using `:buildkite:` notation, or with the emoji itself, such as 😎.
        """
        return pulumi.get(self, "emoji")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        This is the name of the cluster.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def uuid(self) -> pulumi.Output[str]:
        """
        The UUID created with the Cluster.
        """
        return pulumi.get(self, "uuid")

