// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Buildkite.Pipeline.Inputs
{

    public sealed class PipelineTeamGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessLevel", required: true)]
        public Input<string> AccessLevel { get; set; } = null!;

        [Input("pipelineTeamId")]
        public Input<string>? PipelineTeamId { get; set; }

        /// <summary>
        /// The slug of the created pipeline.
        /// </summary>
        [Input("slug", required: true)]
        public Input<string> Slug { get; set; } = null!;

        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        public PipelineTeamGetArgs()
        {
        }
        public static new PipelineTeamGetArgs Empty => new PipelineTeamGetArgs();
    }
}
