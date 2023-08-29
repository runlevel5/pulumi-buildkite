// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package team

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-buildkite/sdk/v2/go/buildkite/internal"
)

// ## # Data Source: team
//
// Use this data source to look up properties of a team. This can be used to
// validate that a team exists before setting the team slug on a pipeline.
//
// Buildkite documentation: https://buildkite.com/docs/pipelines/permissions
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-buildkite/sdk/v2/go/buildkite/Team"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Team.GetTeam(ctx, &team.GetTeamArgs{
//				Id: pulumi.StringRef("<team id>"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupTeam(ctx *pulumi.Context, args *LookupTeamArgs, opts ...pulumi.InvokeOption) (*LookupTeamResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTeamResult
	err := ctx.Invoke("buildkite:Team/getTeam:getTeam", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTeam.
type LookupTeamArgs struct {
	// The GraphQL ID of the team, available in the Settings page for the team.
	Id *string `pulumi:"id"`
	// The slug of the team. Available in the URL of the team on buildkite.com; in the format
	// "<organizaton/team-name>"
	//
	// The `team` data-source supports **either** the use of `id` or `slug` for lookup of a team.
	Slug *string `pulumi:"slug"`
}

// A collection of values returned by getTeam.
type LookupTeamResult struct {
	// Default role to assign to a team member
	DefaultMemberRole string `pulumi:"defaultMemberRole"`
	// Whether new org members will be automatically added to this team
	DefaultTeam bool `pulumi:"defaultTeam"`
	// A description of the team
	Description string `pulumi:"description"`
	// The GraphQL ID of the team
	Id string `pulumi:"id"`
	// Whether team members can create new pipelines and add them to the team
	MembersCanCreatePipelines bool `pulumi:"membersCanCreatePipelines"`
	// The name of the team
	Name string `pulumi:"name"`
	// Whether the team is visible to org members outside this team
	Privacy string `pulumi:"privacy"`
	Slug    string `pulumi:"slug"`
	// The UUID of the team
	Uuid string `pulumi:"uuid"`
}

func LookupTeamOutput(ctx *pulumi.Context, args LookupTeamOutputArgs, opts ...pulumi.InvokeOption) LookupTeamResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTeamResult, error) {
			args := v.(LookupTeamArgs)
			r, err := LookupTeam(ctx, &args, opts...)
			var s LookupTeamResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTeamResultOutput)
}

// A collection of arguments for invoking getTeam.
type LookupTeamOutputArgs struct {
	// The GraphQL ID of the team, available in the Settings page for the team.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The slug of the team. Available in the URL of the team on buildkite.com; in the format
	// "<organizaton/team-name>"
	//
	// The `team` data-source supports **either** the use of `id` or `slug` for lookup of a team.
	Slug pulumi.StringPtrInput `pulumi:"slug"`
}

func (LookupTeamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTeamArgs)(nil)).Elem()
}

// A collection of values returned by getTeam.
type LookupTeamResultOutput struct{ *pulumi.OutputState }

func (LookupTeamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTeamResult)(nil)).Elem()
}

func (o LookupTeamResultOutput) ToLookupTeamResultOutput() LookupTeamResultOutput {
	return o
}

func (o LookupTeamResultOutput) ToLookupTeamResultOutputWithContext(ctx context.Context) LookupTeamResultOutput {
	return o
}

// Default role to assign to a team member
func (o LookupTeamResultOutput) DefaultMemberRole() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.DefaultMemberRole }).(pulumi.StringOutput)
}

// Whether new org members will be automatically added to this team
func (o LookupTeamResultOutput) DefaultTeam() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupTeamResult) bool { return v.DefaultTeam }).(pulumi.BoolOutput)
}

// A description of the team
func (o LookupTeamResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Description }).(pulumi.StringOutput)
}

// The GraphQL ID of the team
func (o LookupTeamResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Id }).(pulumi.StringOutput)
}

// Whether team members can create new pipelines and add them to the team
func (o LookupTeamResultOutput) MembersCanCreatePipelines() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupTeamResult) bool { return v.MembersCanCreatePipelines }).(pulumi.BoolOutput)
}

// The name of the team
func (o LookupTeamResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Name }).(pulumi.StringOutput)
}

// Whether the team is visible to org members outside this team
func (o LookupTeamResultOutput) Privacy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Privacy }).(pulumi.StringOutput)
}

func (o LookupTeamResultOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Slug }).(pulumi.StringOutput)
}

// The UUID of the team
func (o LookupTeamResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTeamResultOutput{})
}
