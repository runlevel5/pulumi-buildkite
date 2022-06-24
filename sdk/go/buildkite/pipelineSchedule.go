// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package buildkite

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Resource: pipelineSchedule
//
// This resource allows you to create and manage pipeline schedules.
//
// Buildkite Documentation: https://buildkite.com/docs/pipelines/scheduled-builds
//
// ## Import
//
// Pipeline schedules can be imported using a slug (which consists of `$BUILDKITE_ORGANIZATION_SLUG/$BUILDKITE_PIPELINE_SLUG/$PIPELINE_SCHEDULE_UUID`), e.g.
//
// ```sh
//  $ pulumi import buildkite:index/pipelineSchedule:PipelineSchedule test myorg/test/1be3e7c7-1e03-4011-accf-b2d8eec90222
// ```
type PipelineSchedule struct {
	pulumi.CustomResourceState

	// The branch to use for the build.
	Branch pulumi.StringOutput `pulumi:"branch"`
	// The commit ref to use for the build.
	Commit pulumi.StringPtrOutput `pulumi:"commit"`
	// Schedule interval (see [docs](https://buildkite.com/docs/pipelines/scheduled-builds#schedule-intervals)).
	Cronline pulumi.StringOutput `pulumi:"cronline"`
	// Whether the schedule should run.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// A map of environment variables to use for the build.
	Env pulumi.StringMapOutput `pulumi:"env"`
	// Schedule label.
	Label pulumi.StringOutput `pulumi:"label"`
	// The message to use for the build.
	Message    pulumi.StringOutput `pulumi:"message"`
	PipelineId pulumi.StringOutput `pulumi:"pipelineId"`
	// The UUID of the pipeline schedule
	Uuid pulumi.StringOutput `pulumi:"uuid"`
}

// NewPipelineSchedule registers a new resource with the given unique name, arguments, and options.
func NewPipelineSchedule(ctx *pulumi.Context,
	name string, args *PipelineScheduleArgs, opts ...pulumi.ResourceOption) (*PipelineSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Branch == nil {
		return nil, errors.New("invalid value for required argument 'Branch'")
	}
	if args.Cronline == nil {
		return nil, errors.New("invalid value for required argument 'Cronline'")
	}
	if args.Label == nil {
		return nil, errors.New("invalid value for required argument 'Label'")
	}
	if args.PipelineId == nil {
		return nil, errors.New("invalid value for required argument 'PipelineId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource PipelineSchedule
	err := ctx.RegisterResource("buildkite:index/pipelineSchedule:PipelineSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipelineSchedule gets an existing PipelineSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipelineSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineScheduleState, opts ...pulumi.ResourceOption) (*PipelineSchedule, error) {
	var resource PipelineSchedule
	err := ctx.ReadResource("buildkite:index/pipelineSchedule:PipelineSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PipelineSchedule resources.
type pipelineScheduleState struct {
	// The branch to use for the build.
	Branch *string `pulumi:"branch"`
	// The commit ref to use for the build.
	Commit *string `pulumi:"commit"`
	// Schedule interval (see [docs](https://buildkite.com/docs/pipelines/scheduled-builds#schedule-intervals)).
	Cronline *string `pulumi:"cronline"`
	// Whether the schedule should run.
	Enabled *bool `pulumi:"enabled"`
	// A map of environment variables to use for the build.
	Env map[string]string `pulumi:"env"`
	// Schedule label.
	Label *string `pulumi:"label"`
	// The message to use for the build.
	Message    *string `pulumi:"message"`
	PipelineId *string `pulumi:"pipelineId"`
	// The UUID of the pipeline schedule
	Uuid *string `pulumi:"uuid"`
}

type PipelineScheduleState struct {
	// The branch to use for the build.
	Branch pulumi.StringPtrInput
	// The commit ref to use for the build.
	Commit pulumi.StringPtrInput
	// Schedule interval (see [docs](https://buildkite.com/docs/pipelines/scheduled-builds#schedule-intervals)).
	Cronline pulumi.StringPtrInput
	// Whether the schedule should run.
	Enabled pulumi.BoolPtrInput
	// A map of environment variables to use for the build.
	Env pulumi.StringMapInput
	// Schedule label.
	Label pulumi.StringPtrInput
	// The message to use for the build.
	Message    pulumi.StringPtrInput
	PipelineId pulumi.StringPtrInput
	// The UUID of the pipeline schedule
	Uuid pulumi.StringPtrInput
}

func (PipelineScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineScheduleState)(nil)).Elem()
}

type pipelineScheduleArgs struct {
	// The branch to use for the build.
	Branch string `pulumi:"branch"`
	// The commit ref to use for the build.
	Commit *string `pulumi:"commit"`
	// Schedule interval (see [docs](https://buildkite.com/docs/pipelines/scheduled-builds#schedule-intervals)).
	Cronline string `pulumi:"cronline"`
	// Whether the schedule should run.
	Enabled *bool `pulumi:"enabled"`
	// A map of environment variables to use for the build.
	Env map[string]string `pulumi:"env"`
	// Schedule label.
	Label string `pulumi:"label"`
	// The message to use for the build.
	Message    *string `pulumi:"message"`
	PipelineId string  `pulumi:"pipelineId"`
}

// The set of arguments for constructing a PipelineSchedule resource.
type PipelineScheduleArgs struct {
	// The branch to use for the build.
	Branch pulumi.StringInput
	// The commit ref to use for the build.
	Commit pulumi.StringPtrInput
	// Schedule interval (see [docs](https://buildkite.com/docs/pipelines/scheduled-builds#schedule-intervals)).
	Cronline pulumi.StringInput
	// Whether the schedule should run.
	Enabled pulumi.BoolPtrInput
	// A map of environment variables to use for the build.
	Env pulumi.StringMapInput
	// Schedule label.
	Label pulumi.StringInput
	// The message to use for the build.
	Message    pulumi.StringPtrInput
	PipelineId pulumi.StringInput
}

func (PipelineScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineScheduleArgs)(nil)).Elem()
}

type PipelineScheduleInput interface {
	pulumi.Input

	ToPipelineScheduleOutput() PipelineScheduleOutput
	ToPipelineScheduleOutputWithContext(ctx context.Context) PipelineScheduleOutput
}

func (*PipelineSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineSchedule)(nil)).Elem()
}

func (i *PipelineSchedule) ToPipelineScheduleOutput() PipelineScheduleOutput {
	return i.ToPipelineScheduleOutputWithContext(context.Background())
}

func (i *PipelineSchedule) ToPipelineScheduleOutputWithContext(ctx context.Context) PipelineScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineScheduleOutput)
}

// PipelineScheduleArrayInput is an input type that accepts PipelineScheduleArray and PipelineScheduleArrayOutput values.
// You can construct a concrete instance of `PipelineScheduleArrayInput` via:
//
//          PipelineScheduleArray{ PipelineScheduleArgs{...} }
type PipelineScheduleArrayInput interface {
	pulumi.Input

	ToPipelineScheduleArrayOutput() PipelineScheduleArrayOutput
	ToPipelineScheduleArrayOutputWithContext(context.Context) PipelineScheduleArrayOutput
}

type PipelineScheduleArray []PipelineScheduleInput

func (PipelineScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PipelineSchedule)(nil)).Elem()
}

func (i PipelineScheduleArray) ToPipelineScheduleArrayOutput() PipelineScheduleArrayOutput {
	return i.ToPipelineScheduleArrayOutputWithContext(context.Background())
}

func (i PipelineScheduleArray) ToPipelineScheduleArrayOutputWithContext(ctx context.Context) PipelineScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineScheduleArrayOutput)
}

// PipelineScheduleMapInput is an input type that accepts PipelineScheduleMap and PipelineScheduleMapOutput values.
// You can construct a concrete instance of `PipelineScheduleMapInput` via:
//
//          PipelineScheduleMap{ "key": PipelineScheduleArgs{...} }
type PipelineScheduleMapInput interface {
	pulumi.Input

	ToPipelineScheduleMapOutput() PipelineScheduleMapOutput
	ToPipelineScheduleMapOutputWithContext(context.Context) PipelineScheduleMapOutput
}

type PipelineScheduleMap map[string]PipelineScheduleInput

func (PipelineScheduleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PipelineSchedule)(nil)).Elem()
}

func (i PipelineScheduleMap) ToPipelineScheduleMapOutput() PipelineScheduleMapOutput {
	return i.ToPipelineScheduleMapOutputWithContext(context.Background())
}

func (i PipelineScheduleMap) ToPipelineScheduleMapOutputWithContext(ctx context.Context) PipelineScheduleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineScheduleMapOutput)
}

type PipelineScheduleOutput struct{ *pulumi.OutputState }

func (PipelineScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineSchedule)(nil)).Elem()
}

func (o PipelineScheduleOutput) ToPipelineScheduleOutput() PipelineScheduleOutput {
	return o
}

func (o PipelineScheduleOutput) ToPipelineScheduleOutputWithContext(ctx context.Context) PipelineScheduleOutput {
	return o
}

// The branch to use for the build.
func (o PipelineScheduleOutput) Branch() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineSchedule) pulumi.StringOutput { return v.Branch }).(pulumi.StringOutput)
}

// The commit ref to use for the build.
func (o PipelineScheduleOutput) Commit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineSchedule) pulumi.StringPtrOutput { return v.Commit }).(pulumi.StringPtrOutput)
}

// Schedule interval (see [docs](https://buildkite.com/docs/pipelines/scheduled-builds#schedule-intervals)).
func (o PipelineScheduleOutput) Cronline() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineSchedule) pulumi.StringOutput { return v.Cronline }).(pulumi.StringOutput)
}

// Whether the schedule should run.
func (o PipelineScheduleOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PipelineSchedule) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// A map of environment variables to use for the build.
func (o PipelineScheduleOutput) Env() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PipelineSchedule) pulumi.StringMapOutput { return v.Env }).(pulumi.StringMapOutput)
}

// Schedule label.
func (o PipelineScheduleOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineSchedule) pulumi.StringOutput { return v.Label }).(pulumi.StringOutput)
}

// The message to use for the build.
func (o PipelineScheduleOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineSchedule) pulumi.StringOutput { return v.Message }).(pulumi.StringOutput)
}

func (o PipelineScheduleOutput) PipelineId() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineSchedule) pulumi.StringOutput { return v.PipelineId }).(pulumi.StringOutput)
}

// The UUID of the pipeline schedule
func (o PipelineScheduleOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineSchedule) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

type PipelineScheduleArrayOutput struct{ *pulumi.OutputState }

func (PipelineScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PipelineSchedule)(nil)).Elem()
}

func (o PipelineScheduleArrayOutput) ToPipelineScheduleArrayOutput() PipelineScheduleArrayOutput {
	return o
}

func (o PipelineScheduleArrayOutput) ToPipelineScheduleArrayOutputWithContext(ctx context.Context) PipelineScheduleArrayOutput {
	return o
}

func (o PipelineScheduleArrayOutput) Index(i pulumi.IntInput) PipelineScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PipelineSchedule {
		return vs[0].([]*PipelineSchedule)[vs[1].(int)]
	}).(PipelineScheduleOutput)
}

type PipelineScheduleMapOutput struct{ *pulumi.OutputState }

func (PipelineScheduleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PipelineSchedule)(nil)).Elem()
}

func (o PipelineScheduleMapOutput) ToPipelineScheduleMapOutput() PipelineScheduleMapOutput {
	return o
}

func (o PipelineScheduleMapOutput) ToPipelineScheduleMapOutputWithContext(ctx context.Context) PipelineScheduleMapOutput {
	return o
}

func (o PipelineScheduleMapOutput) MapIndex(k pulumi.StringInput) PipelineScheduleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PipelineSchedule {
		return vs[0].(map[string]*PipelineSchedule)[vs[1].(string)]
	}).(PipelineScheduleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineScheduleInput)(nil)).Elem(), &PipelineSchedule{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineScheduleArrayInput)(nil)).Elem(), PipelineScheduleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineScheduleMapInput)(nil)).Elem(), PipelineScheduleMap{})
	pulumi.RegisterOutputType(PipelineScheduleOutput{})
	pulumi.RegisterOutputType(PipelineScheduleArrayOutput{})
	pulumi.RegisterOutputType(PipelineScheduleMapOutput{})
}
