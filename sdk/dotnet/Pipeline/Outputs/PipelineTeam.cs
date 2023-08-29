// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Buildkite.Pipeline.Outputs
{

    [OutputType]
    public sealed class PipelineTeam
    {
        public readonly string AccessLevel;
        public readonly string? PipelineTeamId;
        /// <summary>
        /// The slug of the created pipeline.
        /// </summary>
        public readonly string Slug;
        public readonly string? TeamId;

        [OutputConstructor]
        private PipelineTeam(
            string accessLevel,

            string? pipelineTeamId,

            string slug,

            string? teamId)
        {
            AccessLevel = accessLevel;
            PipelineTeamId = pipelineTeamId;
            Slug = slug;
            TeamId = teamId;
        }
    }
}
