# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['PipelineScheduleArgs', 'PipelineSchedule']

@pulumi.input_type
class PipelineScheduleArgs:
    def __init__(__self__, *,
                 branch: pulumi.Input[str],
                 cronline: pulumi.Input[str],
                 label: pulumi.Input[str],
                 pipeline_id: pulumi.Input[str],
                 commit: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 env: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 message: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PipelineSchedule resource.
        :param pulumi.Input[str] branch: The branch to use for the build.
        :param pulumi.Input[str] cronline: Schedule interval (see [docs](https://buildkite.com/docs/pipelines/scheduled-builds#schedule-intervals)).
        :param pulumi.Input[str] label: Schedule label.
        :param pulumi.Input[str] commit: The commit ref to use for the build.
        :param pulumi.Input[bool] enabled: Whether the schedule should run.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] env: A map of environment variables to use for the build.
        :param pulumi.Input[str] message: The message to use for the build.
        """
        pulumi.set(__self__, "branch", branch)
        pulumi.set(__self__, "cronline", cronline)
        pulumi.set(__self__, "label", label)
        pulumi.set(__self__, "pipeline_id", pipeline_id)
        if commit is not None:
            pulumi.set(__self__, "commit", commit)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if env is not None:
            pulumi.set(__self__, "env", env)
        if message is not None:
            pulumi.set(__self__, "message", message)

    @property
    @pulumi.getter
    def branch(self) -> pulumi.Input[str]:
        """
        The branch to use for the build.
        """
        return pulumi.get(self, "branch")

    @branch.setter
    def branch(self, value: pulumi.Input[str]):
        pulumi.set(self, "branch", value)

    @property
    @pulumi.getter
    def cronline(self) -> pulumi.Input[str]:
        """
        Schedule interval (see [docs](https://buildkite.com/docs/pipelines/scheduled-builds#schedule-intervals)).
        """
        return pulumi.get(self, "cronline")

    @cronline.setter
    def cronline(self, value: pulumi.Input[str]):
        pulumi.set(self, "cronline", value)

    @property
    @pulumi.getter
    def label(self) -> pulumi.Input[str]:
        """
        Schedule label.
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: pulumi.Input[str]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter(name="pipelineId")
    def pipeline_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "pipeline_id")

    @pipeline_id.setter
    def pipeline_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "pipeline_id", value)

    @property
    @pulumi.getter
    def commit(self) -> Optional[pulumi.Input[str]]:
        """
        The commit ref to use for the build.
        """
        return pulumi.get(self, "commit")

    @commit.setter
    def commit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "commit", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the schedule should run.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def env(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of environment variables to use for the build.
        """
        return pulumi.get(self, "env")

    @env.setter
    def env(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "env", value)

    @property
    @pulumi.getter
    def message(self) -> Optional[pulumi.Input[str]]:
        """
        The message to use for the build.
        """
        return pulumi.get(self, "message")

    @message.setter
    def message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "message", value)


@pulumi.input_type
class _PipelineScheduleState:
    def __init__(__self__, *,
                 branch: Optional[pulumi.Input[str]] = None,
                 commit: Optional[pulumi.Input[str]] = None,
                 cronline: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 env: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 message: Optional[pulumi.Input[str]] = None,
                 pipeline_id: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PipelineSchedule resources.
        :param pulumi.Input[str] branch: The branch to use for the build.
        :param pulumi.Input[str] commit: The commit ref to use for the build.
        :param pulumi.Input[str] cronline: Schedule interval (see [docs](https://buildkite.com/docs/pipelines/scheduled-builds#schedule-intervals)).
        :param pulumi.Input[bool] enabled: Whether the schedule should run.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] env: A map of environment variables to use for the build.
        :param pulumi.Input[str] label: Schedule label.
        :param pulumi.Input[str] message: The message to use for the build.
        :param pulumi.Input[str] uuid: The UUID of the pipeline schedule
        """
        if branch is not None:
            pulumi.set(__self__, "branch", branch)
        if commit is not None:
            pulumi.set(__self__, "commit", commit)
        if cronline is not None:
            pulumi.set(__self__, "cronline", cronline)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if env is not None:
            pulumi.set(__self__, "env", env)
        if label is not None:
            pulumi.set(__self__, "label", label)
        if message is not None:
            pulumi.set(__self__, "message", message)
        if pipeline_id is not None:
            pulumi.set(__self__, "pipeline_id", pipeline_id)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter
    def branch(self) -> Optional[pulumi.Input[str]]:
        """
        The branch to use for the build.
        """
        return pulumi.get(self, "branch")

    @branch.setter
    def branch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "branch", value)

    @property
    @pulumi.getter
    def commit(self) -> Optional[pulumi.Input[str]]:
        """
        The commit ref to use for the build.
        """
        return pulumi.get(self, "commit")

    @commit.setter
    def commit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "commit", value)

    @property
    @pulumi.getter
    def cronline(self) -> Optional[pulumi.Input[str]]:
        """
        Schedule interval (see [docs](https://buildkite.com/docs/pipelines/scheduled-builds#schedule-intervals)).
        """
        return pulumi.get(self, "cronline")

    @cronline.setter
    def cronline(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cronline", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the schedule should run.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def env(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of environment variables to use for the build.
        """
        return pulumi.get(self, "env")

    @env.setter
    def env(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "env", value)

    @property
    @pulumi.getter
    def label(self) -> Optional[pulumi.Input[str]]:
        """
        Schedule label.
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def message(self) -> Optional[pulumi.Input[str]]:
        """
        The message to use for the build.
        """
        return pulumi.get(self, "message")

    @message.setter
    def message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "message", value)

    @property
    @pulumi.getter(name="pipelineId")
    def pipeline_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "pipeline_id")

    @pipeline_id.setter
    def pipeline_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pipeline_id", value)

    @property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[str]]:
        """
        The UUID of the pipeline schedule
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uuid", value)


class PipelineSchedule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 commit: Optional[pulumi.Input[str]] = None,
                 cronline: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 env: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 message: Optional[pulumi.Input[str]] = None,
                 pipeline_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # Resource: pipeline_schedule

        This resource allows you to create and manage pipeline schedules.

        Buildkite Documentation: https://buildkite.com/docs/pipelines/scheduled-builds

        ## Example Usage

        ```python
        import pulumi
        import pulumi_buildkite as buildkite

        repo2_nightly = buildkite.PipelineSchedule("repo2Nightly",
            pipeline_id=buildkite_pipeline["repo2"]["id"],
            label="Nightly build",
            cronline="@midnight",
            branch=buildkite_pipeline["repo2"]["default_branch"])
        ```

        ## Import

        Pipeline schedules can be imported using a slug (which consists of `$BUILDKITE_ORGANIZATION_SLUG/$BUILDKITE_PIPELINE_SLUG/$PIPELINE_SCHEDULE_UUID`), e.g.

        ```sh
         $ pulumi import buildkite:index/pipelineSchedule:PipelineSchedule test myorg/test/1be3e7c7-1e03-4011-accf-b2d8eec90222
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] branch: The branch to use for the build.
        :param pulumi.Input[str] commit: The commit ref to use for the build.
        :param pulumi.Input[str] cronline: Schedule interval (see [docs](https://buildkite.com/docs/pipelines/scheduled-builds#schedule-intervals)).
        :param pulumi.Input[bool] enabled: Whether the schedule should run.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] env: A map of environment variables to use for the build.
        :param pulumi.Input[str] label: Schedule label.
        :param pulumi.Input[str] message: The message to use for the build.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PipelineScheduleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Resource: pipeline_schedule

        This resource allows you to create and manage pipeline schedules.

        Buildkite Documentation: https://buildkite.com/docs/pipelines/scheduled-builds

        ## Example Usage

        ```python
        import pulumi
        import pulumi_buildkite as buildkite

        repo2_nightly = buildkite.PipelineSchedule("repo2Nightly",
            pipeline_id=buildkite_pipeline["repo2"]["id"],
            label="Nightly build",
            cronline="@midnight",
            branch=buildkite_pipeline["repo2"]["default_branch"])
        ```

        ## Import

        Pipeline schedules can be imported using a slug (which consists of `$BUILDKITE_ORGANIZATION_SLUG/$BUILDKITE_PIPELINE_SLUG/$PIPELINE_SCHEDULE_UUID`), e.g.

        ```sh
         $ pulumi import buildkite:index/pipelineSchedule:PipelineSchedule test myorg/test/1be3e7c7-1e03-4011-accf-b2d8eec90222
        ```

        :param str resource_name: The name of the resource.
        :param PipelineScheduleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PipelineScheduleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 commit: Optional[pulumi.Input[str]] = None,
                 cronline: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 env: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 message: Optional[pulumi.Input[str]] = None,
                 pipeline_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PipelineScheduleArgs.__new__(PipelineScheduleArgs)

            if branch is None and not opts.urn:
                raise TypeError("Missing required property 'branch'")
            __props__.__dict__["branch"] = branch
            __props__.__dict__["commit"] = commit
            if cronline is None and not opts.urn:
                raise TypeError("Missing required property 'cronline'")
            __props__.__dict__["cronline"] = cronline
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["env"] = env
            if label is None and not opts.urn:
                raise TypeError("Missing required property 'label'")
            __props__.__dict__["label"] = label
            __props__.__dict__["message"] = message
            if pipeline_id is None and not opts.urn:
                raise TypeError("Missing required property 'pipeline_id'")
            __props__.__dict__["pipeline_id"] = pipeline_id
            __props__.__dict__["uuid"] = None
        super(PipelineSchedule, __self__).__init__(
            'buildkite:index/pipelineSchedule:PipelineSchedule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            branch: Optional[pulumi.Input[str]] = None,
            commit: Optional[pulumi.Input[str]] = None,
            cronline: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            env: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            label: Optional[pulumi.Input[str]] = None,
            message: Optional[pulumi.Input[str]] = None,
            pipeline_id: Optional[pulumi.Input[str]] = None,
            uuid: Optional[pulumi.Input[str]] = None) -> 'PipelineSchedule':
        """
        Get an existing PipelineSchedule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] branch: The branch to use for the build.
        :param pulumi.Input[str] commit: The commit ref to use for the build.
        :param pulumi.Input[str] cronline: Schedule interval (see [docs](https://buildkite.com/docs/pipelines/scheduled-builds#schedule-intervals)).
        :param pulumi.Input[bool] enabled: Whether the schedule should run.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] env: A map of environment variables to use for the build.
        :param pulumi.Input[str] label: Schedule label.
        :param pulumi.Input[str] message: The message to use for the build.
        :param pulumi.Input[str] uuid: The UUID of the pipeline schedule
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PipelineScheduleState.__new__(_PipelineScheduleState)

        __props__.__dict__["branch"] = branch
        __props__.__dict__["commit"] = commit
        __props__.__dict__["cronline"] = cronline
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["env"] = env
        __props__.__dict__["label"] = label
        __props__.__dict__["message"] = message
        __props__.__dict__["pipeline_id"] = pipeline_id
        __props__.__dict__["uuid"] = uuid
        return PipelineSchedule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def branch(self) -> pulumi.Output[str]:
        """
        The branch to use for the build.
        """
        return pulumi.get(self, "branch")

    @property
    @pulumi.getter
    def commit(self) -> pulumi.Output[Optional[str]]:
        """
        The commit ref to use for the build.
        """
        return pulumi.get(self, "commit")

    @property
    @pulumi.getter
    def cronline(self) -> pulumi.Output[str]:
        """
        Schedule interval (see [docs](https://buildkite.com/docs/pipelines/scheduled-builds#schedule-intervals)).
        """
        return pulumi.get(self, "cronline")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the schedule should run.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def env(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of environment variables to use for the build.
        """
        return pulumi.get(self, "env")

    @property
    @pulumi.getter
    def label(self) -> pulumi.Output[str]:
        """
        Schedule label.
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter
    def message(self) -> pulumi.Output[str]:
        """
        The message to use for the build.
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter(name="pipelineId")
    def pipeline_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "pipeline_id")

    @property
    @pulumi.getter
    def uuid(self) -> pulumi.Output[str]:
        """
        The UUID of the pipeline schedule
        """
        return pulumi.get(self, "uuid")

