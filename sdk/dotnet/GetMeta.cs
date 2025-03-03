// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Buildkite
{
    public static class GetMeta
    {
        public static Task<GetMetaResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMetaResult>("buildkite:index/getMeta:getMeta", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetMetaResult
    {
        public readonly string Id;
        /// <summary>
        /// A list of strings, each one an IP address (x.x.x.x) or CIDR address (x.x.x.x/32) that Buildkite may use to send webhooks and other external requests.
        /// </summary>
        public readonly ImmutableArray<string> WebhookIps;

        [OutputConstructor]
        private GetMetaResult(
            string id,

            ImmutableArray<string> webhookIps)
        {
            Id = id;
            WebhookIps = webhookIps;
        }
    }
}
