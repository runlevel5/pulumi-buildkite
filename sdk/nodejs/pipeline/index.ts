// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetPipelineArgs, GetPipelineResult, GetPipelineOutputArgs } from "./getPipeline";
export const getPipeline: typeof import("./getPipeline").getPipeline = null as any;
export const getPipelineOutput: typeof import("./getPipeline").getPipelineOutput = null as any;
utilities.lazyLoad(exports, ["getPipeline","getPipelineOutput"], () => require("./getPipeline"));

export { PipelineArgs, PipelineState } from "./pipeline";
export type Pipeline = import("./pipeline").Pipeline;
export const Pipeline: typeof import("./pipeline").Pipeline = null as any;
utilities.lazyLoad(exports, ["Pipeline"], () => require("./pipeline"));

export { ScheduleArgs, ScheduleState } from "./schedule";
export type Schedule = import("./schedule").Schedule;
export const Schedule: typeof import("./schedule").Schedule = null as any;
utilities.lazyLoad(exports, ["Schedule"], () => require("./schedule"));

export { TeamArgs, TeamState } from "./team";
export type Team = import("./team").Team;
export const Team: typeof import("./team").Team = null as any;
utilities.lazyLoad(exports, ["Team"], () => require("./team"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "buildkite:Pipeline/pipeline:Pipeline":
                return new Pipeline(name, <any>undefined, { urn })
            case "buildkite:Pipeline/schedule:Schedule":
                return new Schedule(name, <any>undefined, { urn })
            case "buildkite:Pipeline/team:Team":
                return new Team(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("buildkite", "Pipeline/pipeline", _module)
pulumi.runtime.registerResourceModule("buildkite", "Pipeline/schedule", _module)
pulumi.runtime.registerResourceModule("buildkite", "Pipeline/team", _module)
