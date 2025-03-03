// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package team

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-buildkite/sdk/v2/go/buildkite/internal"
)

// ## # Resource: teamMember
//
// This resource allows manage team membership for existing organization users.
//
// The user must already be part of the organization to which you are managing team membership. This will not invite a new user to the organization.
//
// Buildkite Documentation: https://buildkite.com/docs/pipelines/permissions
//
// Note: You must first enable Teams on your organization.
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
//			team, err := Team.NewTeam(ctx, "team", &Team.TeamArgs{
//				Privacy:           pulumi.String("VISIBLE"),
//				DefaultTeam:       pulumi.Bool(true),
//				DefaultMemberRole: pulumi.String("MEMBER"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Team.NewMember(ctx, "aSmith", &Team.MemberArgs{
//				Role:   pulumi.String("MEMBER"),
//				TeamId: team.ID(),
//				UserId: pulumi.String("VXNlci0tLWRlOTdmMjBiLWJkZmMtNGNjOC1hOTcwLTY1ODNiZTk2ZGEyYQ=="),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Team members can be imported using the GraphQL ID of the membership. Note this is different to the ID of the user.
//
// ```sh
//
//	$ pulumi import buildkite:Team/member:Member a_smith VGVhbU1lbWJlci0tLTVlZDEyMmY2LTM2NjQtNDI1MS04YzMwLTc4NjRiMDdiZDQ4Zg==
//
// ```
//
//	To find the ID to use, you can use the GraphQL query below. Alternatively, you could use this [pre-saved query](https://buildkite.com/user/graphql/console/ce4540dd-4f60-4e79-8e8f-9f4c3bc8784e), where you will need fo fill in the organization slug and search terms for teams and members. Both search terms (TEAM_SEARCH_TERM and TEAM_MEMBER_SEARCH_TERM) work on the name of the associated object. graphql query getTeamMemberId {
//
//	organization(slug"ORGANIZATION_SLUG") {
//
//	teams(first2, search"TEAM_SEARCH_TERM") {
//
//	edges {
//
//	node {
//
//	members(first2, search"TEAM_MEMBER_SEARCH_TERM") {
//
//	edges {
//
//	node {
//
//	id
//
//	}
//
//	}
//
//	}
//
//	}
//
//	}
//
//	}
//
//	} }
type Member struct {
	pulumi.CustomResourceState

	// Either MEMBER or MAINTAINER.
	Role pulumi.StringOutput `pulumi:"role"`
	// The GraphQL ID of the team to add to/remove from.
	TeamId pulumi.StringOutput `pulumi:"teamId"`
	// The GraphQL ID of the user to add/remove.
	UserId pulumi.StringOutput `pulumi:"userId"`
	// The UUID for the team membership.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
}

// NewMember registers a new resource with the given unique name, arguments, and options.
func NewMember(ctx *pulumi.Context,
	name string, args *MemberArgs, opts ...pulumi.ResourceOption) (*Member, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.TeamId == nil {
		return nil, errors.New("invalid value for required argument 'TeamId'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Member
	err := ctx.RegisterResource("buildkite:Team/member:Member", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMember gets an existing Member resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MemberState, opts ...pulumi.ResourceOption) (*Member, error) {
	var resource Member
	err := ctx.ReadResource("buildkite:Team/member:Member", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Member resources.
type memberState struct {
	// Either MEMBER or MAINTAINER.
	Role *string `pulumi:"role"`
	// The GraphQL ID of the team to add to/remove from.
	TeamId *string `pulumi:"teamId"`
	// The GraphQL ID of the user to add/remove.
	UserId *string `pulumi:"userId"`
	// The UUID for the team membership.
	Uuid *string `pulumi:"uuid"`
}

type MemberState struct {
	// Either MEMBER or MAINTAINER.
	Role pulumi.StringPtrInput
	// The GraphQL ID of the team to add to/remove from.
	TeamId pulumi.StringPtrInput
	// The GraphQL ID of the user to add/remove.
	UserId pulumi.StringPtrInput
	// The UUID for the team membership.
	Uuid pulumi.StringPtrInput
}

func (MemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*memberState)(nil)).Elem()
}

type memberArgs struct {
	// Either MEMBER or MAINTAINER.
	Role string `pulumi:"role"`
	// The GraphQL ID of the team to add to/remove from.
	TeamId string `pulumi:"teamId"`
	// The GraphQL ID of the user to add/remove.
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a Member resource.
type MemberArgs struct {
	// Either MEMBER or MAINTAINER.
	Role pulumi.StringInput
	// The GraphQL ID of the team to add to/remove from.
	TeamId pulumi.StringInput
	// The GraphQL ID of the user to add/remove.
	UserId pulumi.StringInput
}

func (MemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*memberArgs)(nil)).Elem()
}

type MemberInput interface {
	pulumi.Input

	ToMemberOutput() MemberOutput
	ToMemberOutputWithContext(ctx context.Context) MemberOutput
}

func (*Member) ElementType() reflect.Type {
	return reflect.TypeOf((**Member)(nil)).Elem()
}

func (i *Member) ToMemberOutput() MemberOutput {
	return i.ToMemberOutputWithContext(context.Background())
}

func (i *Member) ToMemberOutputWithContext(ctx context.Context) MemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberOutput)
}

// MemberArrayInput is an input type that accepts MemberArray and MemberArrayOutput values.
// You can construct a concrete instance of `MemberArrayInput` via:
//
//	MemberArray{ MemberArgs{...} }
type MemberArrayInput interface {
	pulumi.Input

	ToMemberArrayOutput() MemberArrayOutput
	ToMemberArrayOutputWithContext(context.Context) MemberArrayOutput
}

type MemberArray []MemberInput

func (MemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Member)(nil)).Elem()
}

func (i MemberArray) ToMemberArrayOutput() MemberArrayOutput {
	return i.ToMemberArrayOutputWithContext(context.Background())
}

func (i MemberArray) ToMemberArrayOutputWithContext(ctx context.Context) MemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberArrayOutput)
}

// MemberMapInput is an input type that accepts MemberMap and MemberMapOutput values.
// You can construct a concrete instance of `MemberMapInput` via:
//
//	MemberMap{ "key": MemberArgs{...} }
type MemberMapInput interface {
	pulumi.Input

	ToMemberMapOutput() MemberMapOutput
	ToMemberMapOutputWithContext(context.Context) MemberMapOutput
}

type MemberMap map[string]MemberInput

func (MemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Member)(nil)).Elem()
}

func (i MemberMap) ToMemberMapOutput() MemberMapOutput {
	return i.ToMemberMapOutputWithContext(context.Background())
}

func (i MemberMap) ToMemberMapOutputWithContext(ctx context.Context) MemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberMapOutput)
}

type MemberOutput struct{ *pulumi.OutputState }

func (MemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Member)(nil)).Elem()
}

func (o MemberOutput) ToMemberOutput() MemberOutput {
	return o
}

func (o MemberOutput) ToMemberOutputWithContext(ctx context.Context) MemberOutput {
	return o
}

// Either MEMBER or MAINTAINER.
func (o MemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *Member) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// The GraphQL ID of the team to add to/remove from.
func (o MemberOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v *Member) pulumi.StringOutput { return v.TeamId }).(pulumi.StringOutput)
}

// The GraphQL ID of the user to add/remove.
func (o MemberOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *Member) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

// The UUID for the team membership.
func (o MemberOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Member) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

type MemberArrayOutput struct{ *pulumi.OutputState }

func (MemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Member)(nil)).Elem()
}

func (o MemberArrayOutput) ToMemberArrayOutput() MemberArrayOutput {
	return o
}

func (o MemberArrayOutput) ToMemberArrayOutputWithContext(ctx context.Context) MemberArrayOutput {
	return o
}

func (o MemberArrayOutput) Index(i pulumi.IntInput) MemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Member {
		return vs[0].([]*Member)[vs[1].(int)]
	}).(MemberOutput)
}

type MemberMapOutput struct{ *pulumi.OutputState }

func (MemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Member)(nil)).Elem()
}

func (o MemberMapOutput) ToMemberMapOutput() MemberMapOutput {
	return o
}

func (o MemberMapOutput) ToMemberMapOutputWithContext(ctx context.Context) MemberMapOutput {
	return o
}

func (o MemberMapOutput) MapIndex(k pulumi.StringInput) MemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Member {
		return vs[0].(map[string]*Member)[vs[1].(string)]
	}).(MemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MemberInput)(nil)).Elem(), &Member{})
	pulumi.RegisterInputType(reflect.TypeOf((*MemberArrayInput)(nil)).Elem(), MemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MemberMapInput)(nil)).Elem(), MemberMap{})
	pulumi.RegisterOutputType(MemberOutput{})
	pulumi.RegisterOutputType(MemberArrayOutput{})
	pulumi.RegisterOutputType(MemberMapOutput{})
}
