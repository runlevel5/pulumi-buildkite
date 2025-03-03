# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TeamArgs', 'Team']

@pulumi.input_type
class TeamArgs:
    def __init__(__self__, *,
                 default_member_role: pulumi.Input[str],
                 default_team: pulumi.Input[bool],
                 privacy: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 members_can_create_pipelines: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Team resource.
        :param pulumi.Input[str] default_member_role: Default role to assign to a team member.
        :param pulumi.Input[bool] default_team: Whether to assign this team to a user by default.
        :param pulumi.Input[str] privacy: The privacy level to set the team too.
        :param pulumi.Input[str] description: The description to assign to the team.
        :param pulumi.Input[bool] members_can_create_pipelines: Whether team members can create.
        :param pulumi.Input[str] name: The name of the team.
        """
        pulumi.set(__self__, "default_member_role", default_member_role)
        pulumi.set(__self__, "default_team", default_team)
        pulumi.set(__self__, "privacy", privacy)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if members_can_create_pipelines is not None:
            pulumi.set(__self__, "members_can_create_pipelines", members_can_create_pipelines)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="defaultMemberRole")
    def default_member_role(self) -> pulumi.Input[str]:
        """
        Default role to assign to a team member.
        """
        return pulumi.get(self, "default_member_role")

    @default_member_role.setter
    def default_member_role(self, value: pulumi.Input[str]):
        pulumi.set(self, "default_member_role", value)

    @property
    @pulumi.getter(name="defaultTeam")
    def default_team(self) -> pulumi.Input[bool]:
        """
        Whether to assign this team to a user by default.
        """
        return pulumi.get(self, "default_team")

    @default_team.setter
    def default_team(self, value: pulumi.Input[bool]):
        pulumi.set(self, "default_team", value)

    @property
    @pulumi.getter
    def privacy(self) -> pulumi.Input[str]:
        """
        The privacy level to set the team too.
        """
        return pulumi.get(self, "privacy")

    @privacy.setter
    def privacy(self, value: pulumi.Input[str]):
        pulumi.set(self, "privacy", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description to assign to the team.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="membersCanCreatePipelines")
    def members_can_create_pipelines(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether team members can create.
        """
        return pulumi.get(self, "members_can_create_pipelines")

    @members_can_create_pipelines.setter
    def members_can_create_pipelines(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "members_can_create_pipelines", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the team.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _TeamState:
    def __init__(__self__, *,
                 default_member_role: Optional[pulumi.Input[str]] = None,
                 default_team: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 members_can_create_pipelines: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 privacy: Optional[pulumi.Input[str]] = None,
                 slug: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Team resources.
        :param pulumi.Input[str] default_member_role: Default role to assign to a team member.
        :param pulumi.Input[bool] default_team: Whether to assign this team to a user by default.
        :param pulumi.Input[str] description: The description to assign to the team.
        :param pulumi.Input[bool] members_can_create_pipelines: Whether team members can create.
        :param pulumi.Input[str] name: The name of the team.
        :param pulumi.Input[str] privacy: The privacy level to set the team too.
        :param pulumi.Input[str] slug: The name of the team.
        :param pulumi.Input[str] uuid: The UUID for the team.
        """
        if default_member_role is not None:
            pulumi.set(__self__, "default_member_role", default_member_role)
        if default_team is not None:
            pulumi.set(__self__, "default_team", default_team)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if members_can_create_pipelines is not None:
            pulumi.set(__self__, "members_can_create_pipelines", members_can_create_pipelines)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if privacy is not None:
            pulumi.set(__self__, "privacy", privacy)
        if slug is not None:
            pulumi.set(__self__, "slug", slug)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter(name="defaultMemberRole")
    def default_member_role(self) -> Optional[pulumi.Input[str]]:
        """
        Default role to assign to a team member.
        """
        return pulumi.get(self, "default_member_role")

    @default_member_role.setter
    def default_member_role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_member_role", value)

    @property
    @pulumi.getter(name="defaultTeam")
    def default_team(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to assign this team to a user by default.
        """
        return pulumi.get(self, "default_team")

    @default_team.setter
    def default_team(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "default_team", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description to assign to the team.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="membersCanCreatePipelines")
    def members_can_create_pipelines(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether team members can create.
        """
        return pulumi.get(self, "members_can_create_pipelines")

    @members_can_create_pipelines.setter
    def members_can_create_pipelines(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "members_can_create_pipelines", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the team.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def privacy(self) -> Optional[pulumi.Input[str]]:
        """
        The privacy level to set the team too.
        """
        return pulumi.get(self, "privacy")

    @privacy.setter
    def privacy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "privacy", value)

    @property
    @pulumi.getter
    def slug(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the team.
        """
        return pulumi.get(self, "slug")

    @slug.setter
    def slug(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slug", value)

    @property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[str]]:
        """
        The UUID for the team.
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uuid", value)


class Team(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_member_role: Optional[pulumi.Input[str]] = None,
                 default_team: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 members_can_create_pipelines: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 privacy: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # Resource: team

        This resource allows you to create and manage teams.

        Buildkite Documentation: https://buildkite.com/docs/pipelines/permissions

        Note: You must first enable Teams on your organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_buildkite as buildkite

        team = buildkite.team.Team("team",
            default_member_role="MEMBER",
            default_team=True,
            privacy="VISIBLE")
        ```

        ## Import

        Teams can be imported using the `GraphQL ID` (not UUID), e.g.

        ```sh
         $ pulumi import buildkite:Team/team:Team fleet UGlwZWxpbmUtLS00MzVjYWQ1OC1lODFkLTQ1YWYtODYzNy1iMWNmODA3MDIzOGQ=
        ```

         To find the ID to use, you can use the GraphQL query below. Alternatively, you could use this [pre-saved query](https://buildkite.com/user/graphql/console/6e74c89c-4e91-4d1d-92ca-4fb19d0ea453), where you will need fo fill in the organization slug and search term (TEAM_SEARCH_TERM) for the team. graphql query getTeamId {

         organization(slug"ORGANIZATION_SLUG") {

         teams(first1, search"TEAM_SEARCH_TERM") {

         edges {

         node {

         id

         name

         }

         }

         }

         } }

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_member_role: Default role to assign to a team member.
        :param pulumi.Input[bool] default_team: Whether to assign this team to a user by default.
        :param pulumi.Input[str] description: The description to assign to the team.
        :param pulumi.Input[bool] members_can_create_pipelines: Whether team members can create.
        :param pulumi.Input[str] name: The name of the team.
        :param pulumi.Input[str] privacy: The privacy level to set the team too.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TeamArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Resource: team

        This resource allows you to create and manage teams.

        Buildkite Documentation: https://buildkite.com/docs/pipelines/permissions

        Note: You must first enable Teams on your organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_buildkite as buildkite

        team = buildkite.team.Team("team",
            default_member_role="MEMBER",
            default_team=True,
            privacy="VISIBLE")
        ```

        ## Import

        Teams can be imported using the `GraphQL ID` (not UUID), e.g.

        ```sh
         $ pulumi import buildkite:Team/team:Team fleet UGlwZWxpbmUtLS00MzVjYWQ1OC1lODFkLTQ1YWYtODYzNy1iMWNmODA3MDIzOGQ=
        ```

         To find the ID to use, you can use the GraphQL query below. Alternatively, you could use this [pre-saved query](https://buildkite.com/user/graphql/console/6e74c89c-4e91-4d1d-92ca-4fb19d0ea453), where you will need fo fill in the organization slug and search term (TEAM_SEARCH_TERM) for the team. graphql query getTeamId {

         organization(slug"ORGANIZATION_SLUG") {

         teams(first1, search"TEAM_SEARCH_TERM") {

         edges {

         node {

         id

         name

         }

         }

         }

         } }

        :param str resource_name: The name of the resource.
        :param TeamArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TeamArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_member_role: Optional[pulumi.Input[str]] = None,
                 default_team: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 members_can_create_pipelines: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 privacy: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TeamArgs.__new__(TeamArgs)

            if default_member_role is None and not opts.urn:
                raise TypeError("Missing required property 'default_member_role'")
            __props__.__dict__["default_member_role"] = default_member_role
            if default_team is None and not opts.urn:
                raise TypeError("Missing required property 'default_team'")
            __props__.__dict__["default_team"] = default_team
            __props__.__dict__["description"] = description
            __props__.__dict__["members_can_create_pipelines"] = members_can_create_pipelines
            __props__.__dict__["name"] = name
            if privacy is None and not opts.urn:
                raise TypeError("Missing required property 'privacy'")
            __props__.__dict__["privacy"] = privacy
            __props__.__dict__["slug"] = None
            __props__.__dict__["uuid"] = None
        super(Team, __self__).__init__(
            'buildkite:Team/team:Team',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            default_member_role: Optional[pulumi.Input[str]] = None,
            default_team: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            members_can_create_pipelines: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            privacy: Optional[pulumi.Input[str]] = None,
            slug: Optional[pulumi.Input[str]] = None,
            uuid: Optional[pulumi.Input[str]] = None) -> 'Team':
        """
        Get an existing Team resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_member_role: Default role to assign to a team member.
        :param pulumi.Input[bool] default_team: Whether to assign this team to a user by default.
        :param pulumi.Input[str] description: The description to assign to the team.
        :param pulumi.Input[bool] members_can_create_pipelines: Whether team members can create.
        :param pulumi.Input[str] name: The name of the team.
        :param pulumi.Input[str] privacy: The privacy level to set the team too.
        :param pulumi.Input[str] slug: The name of the team.
        :param pulumi.Input[str] uuid: The UUID for the team.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TeamState.__new__(_TeamState)

        __props__.__dict__["default_member_role"] = default_member_role
        __props__.__dict__["default_team"] = default_team
        __props__.__dict__["description"] = description
        __props__.__dict__["members_can_create_pipelines"] = members_can_create_pipelines
        __props__.__dict__["name"] = name
        __props__.__dict__["privacy"] = privacy
        __props__.__dict__["slug"] = slug
        __props__.__dict__["uuid"] = uuid
        return Team(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defaultMemberRole")
    def default_member_role(self) -> pulumi.Output[str]:
        """
        Default role to assign to a team member.
        """
        return pulumi.get(self, "default_member_role")

    @property
    @pulumi.getter(name="defaultTeam")
    def default_team(self) -> pulumi.Output[bool]:
        """
        Whether to assign this team to a user by default.
        """
        return pulumi.get(self, "default_team")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description to assign to the team.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="membersCanCreatePipelines")
    def members_can_create_pipelines(self) -> pulumi.Output[bool]:
        """
        Whether team members can create.
        """
        return pulumi.get(self, "members_can_create_pipelines")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the team.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def privacy(self) -> pulumi.Output[str]:
        """
        The privacy level to set the team too.
        """
        return pulumi.get(self, "privacy")

    @property
    @pulumi.getter
    def slug(self) -> pulumi.Output[str]:
        """
        The name of the team.
        """
        return pulumi.get(self, "slug")

    @property
    @pulumi.getter
    def uuid(self) -> pulumi.Output[str]:
        """
        The UUID for the team.
        """
        return pulumi.get(self, "uuid")

