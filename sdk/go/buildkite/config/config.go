// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	"github.com/pulumiverse/pulumi-buildkite/sdk/v2/go/buildkite/internal"
)

var _ = internal.GetEnvOrDefault

// API token with GraphQL access and `write_pipelines, read_pipelines` and `write_suites` REST API scopes
func GetApiToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "buildkite:apiToken")
}

// Archive pipelines when destroying instead of completely deleting.
func GetArchivePipelineOnDelete(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "buildkite:archivePipelineOnDelete")
}

// Base URL for the GraphQL API to use
func GetGraphqlUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "buildkite:graphqlUrl")
}

// The Buildkite organization slug
func GetOrganization(ctx *pulumi.Context) string {
	return config.Get(ctx, "buildkite:organization")
}

// Base URL for the REST API to use
func GetRestUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "buildkite:restUrl")
}
