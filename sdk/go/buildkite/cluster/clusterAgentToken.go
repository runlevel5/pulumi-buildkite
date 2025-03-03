// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cluster

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-buildkite/sdk/v2/go/buildkite/internal"
)

// ## # Resource: clusterAgentToken
//
// This resource allows you to create and manage cluster agent tokens.
//
// Buildkite Documentation: https://buildkite.com/docs/clusters/manage-clusters#set-up-clusters-connect-agents-to-a-cluster
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-buildkite/sdk/v2/go/buildkite/Cluster"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cluster.NewClusterAgentToken(ctx, "cluster-token-1", &Cluster.ClusterAgentTokenArgs{
//				ClusterId:   pulumi.String("Q2x1c3Rlci0tLTkyMmVjOTA4LWRmNWItNDhhYS1hMThjLTczMzE0YjQ1ZGYyMA=="),
//				Description: pulumi.String("agent token for cluster-1"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ClusterAgentToken struct {
	pulumi.CustomResourceState

	// The ID of the cluster that this cluster queue belongs to.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The UUID of the cluster that this cluster queue belongs to.
	ClusterUuid pulumi.StringOutput `pulumi:"clusterUuid"`
	// A description about what this cluster agent token is used for.
	Description pulumi.StringOutput `pulumi:"description"`
	Token       pulumi.StringOutput `pulumi:"token"`
	// The UUID of the created cluster queue.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
}

// NewClusterAgentToken registers a new resource with the given unique name, arguments, and options.
func NewClusterAgentToken(ctx *pulumi.Context,
	name string, args *ClusterAgentTokenArgs, opts ...pulumi.ResourceOption) (*ClusterAgentToken, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClusterAgentToken
	err := ctx.RegisterResource("buildkite:Cluster/clusterAgentToken:ClusterAgentToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterAgentToken gets an existing ClusterAgentToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterAgentToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterAgentTokenState, opts ...pulumi.ResourceOption) (*ClusterAgentToken, error) {
	var resource ClusterAgentToken
	err := ctx.ReadResource("buildkite:Cluster/clusterAgentToken:ClusterAgentToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterAgentToken resources.
type clusterAgentTokenState struct {
	// The ID of the cluster that this cluster queue belongs to.
	ClusterId *string `pulumi:"clusterId"`
	// The UUID of the cluster that this cluster queue belongs to.
	ClusterUuid *string `pulumi:"clusterUuid"`
	// A description about what this cluster agent token is used for.
	Description *string `pulumi:"description"`
	Token       *string `pulumi:"token"`
	// The UUID of the created cluster queue.
	Uuid *string `pulumi:"uuid"`
}

type ClusterAgentTokenState struct {
	// The ID of the cluster that this cluster queue belongs to.
	ClusterId pulumi.StringPtrInput
	// The UUID of the cluster that this cluster queue belongs to.
	ClusterUuid pulumi.StringPtrInput
	// A description about what this cluster agent token is used for.
	Description pulumi.StringPtrInput
	Token       pulumi.StringPtrInput
	// The UUID of the created cluster queue.
	Uuid pulumi.StringPtrInput
}

func (ClusterAgentTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterAgentTokenState)(nil)).Elem()
}

type clusterAgentTokenArgs struct {
	// The ID of the cluster that this cluster queue belongs to.
	ClusterId string `pulumi:"clusterId"`
	// A description about what this cluster agent token is used for.
	Description string `pulumi:"description"`
}

// The set of arguments for constructing a ClusterAgentToken resource.
type ClusterAgentTokenArgs struct {
	// The ID of the cluster that this cluster queue belongs to.
	ClusterId pulumi.StringInput
	// A description about what this cluster agent token is used for.
	Description pulumi.StringInput
}

func (ClusterAgentTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterAgentTokenArgs)(nil)).Elem()
}

type ClusterAgentTokenInput interface {
	pulumi.Input

	ToClusterAgentTokenOutput() ClusterAgentTokenOutput
	ToClusterAgentTokenOutputWithContext(ctx context.Context) ClusterAgentTokenOutput
}

func (*ClusterAgentToken) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterAgentToken)(nil)).Elem()
}

func (i *ClusterAgentToken) ToClusterAgentTokenOutput() ClusterAgentTokenOutput {
	return i.ToClusterAgentTokenOutputWithContext(context.Background())
}

func (i *ClusterAgentToken) ToClusterAgentTokenOutputWithContext(ctx context.Context) ClusterAgentTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterAgentTokenOutput)
}

// ClusterAgentTokenArrayInput is an input type that accepts ClusterAgentTokenArray and ClusterAgentTokenArrayOutput values.
// You can construct a concrete instance of `ClusterAgentTokenArrayInput` via:
//
//	ClusterAgentTokenArray{ ClusterAgentTokenArgs{...} }
type ClusterAgentTokenArrayInput interface {
	pulumi.Input

	ToClusterAgentTokenArrayOutput() ClusterAgentTokenArrayOutput
	ToClusterAgentTokenArrayOutputWithContext(context.Context) ClusterAgentTokenArrayOutput
}

type ClusterAgentTokenArray []ClusterAgentTokenInput

func (ClusterAgentTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterAgentToken)(nil)).Elem()
}

func (i ClusterAgentTokenArray) ToClusterAgentTokenArrayOutput() ClusterAgentTokenArrayOutput {
	return i.ToClusterAgentTokenArrayOutputWithContext(context.Background())
}

func (i ClusterAgentTokenArray) ToClusterAgentTokenArrayOutputWithContext(ctx context.Context) ClusterAgentTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterAgentTokenArrayOutput)
}

// ClusterAgentTokenMapInput is an input type that accepts ClusterAgentTokenMap and ClusterAgentTokenMapOutput values.
// You can construct a concrete instance of `ClusterAgentTokenMapInput` via:
//
//	ClusterAgentTokenMap{ "key": ClusterAgentTokenArgs{...} }
type ClusterAgentTokenMapInput interface {
	pulumi.Input

	ToClusterAgentTokenMapOutput() ClusterAgentTokenMapOutput
	ToClusterAgentTokenMapOutputWithContext(context.Context) ClusterAgentTokenMapOutput
}

type ClusterAgentTokenMap map[string]ClusterAgentTokenInput

func (ClusterAgentTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterAgentToken)(nil)).Elem()
}

func (i ClusterAgentTokenMap) ToClusterAgentTokenMapOutput() ClusterAgentTokenMapOutput {
	return i.ToClusterAgentTokenMapOutputWithContext(context.Background())
}

func (i ClusterAgentTokenMap) ToClusterAgentTokenMapOutputWithContext(ctx context.Context) ClusterAgentTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterAgentTokenMapOutput)
}

type ClusterAgentTokenOutput struct{ *pulumi.OutputState }

func (ClusterAgentTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterAgentToken)(nil)).Elem()
}

func (o ClusterAgentTokenOutput) ToClusterAgentTokenOutput() ClusterAgentTokenOutput {
	return o
}

func (o ClusterAgentTokenOutput) ToClusterAgentTokenOutputWithContext(ctx context.Context) ClusterAgentTokenOutput {
	return o
}

// The ID of the cluster that this cluster queue belongs to.
func (o ClusterAgentTokenOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterAgentToken) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The UUID of the cluster that this cluster queue belongs to.
func (o ClusterAgentTokenOutput) ClusterUuid() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterAgentToken) pulumi.StringOutput { return v.ClusterUuid }).(pulumi.StringOutput)
}

// A description about what this cluster agent token is used for.
func (o ClusterAgentTokenOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterAgentToken) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o ClusterAgentTokenOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterAgentToken) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// The UUID of the created cluster queue.
func (o ClusterAgentTokenOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterAgentToken) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

type ClusterAgentTokenArrayOutput struct{ *pulumi.OutputState }

func (ClusterAgentTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterAgentToken)(nil)).Elem()
}

func (o ClusterAgentTokenArrayOutput) ToClusterAgentTokenArrayOutput() ClusterAgentTokenArrayOutput {
	return o
}

func (o ClusterAgentTokenArrayOutput) ToClusterAgentTokenArrayOutputWithContext(ctx context.Context) ClusterAgentTokenArrayOutput {
	return o
}

func (o ClusterAgentTokenArrayOutput) Index(i pulumi.IntInput) ClusterAgentTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterAgentToken {
		return vs[0].([]*ClusterAgentToken)[vs[1].(int)]
	}).(ClusterAgentTokenOutput)
}

type ClusterAgentTokenMapOutput struct{ *pulumi.OutputState }

func (ClusterAgentTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterAgentToken)(nil)).Elem()
}

func (o ClusterAgentTokenMapOutput) ToClusterAgentTokenMapOutput() ClusterAgentTokenMapOutput {
	return o
}

func (o ClusterAgentTokenMapOutput) ToClusterAgentTokenMapOutputWithContext(ctx context.Context) ClusterAgentTokenMapOutput {
	return o
}

func (o ClusterAgentTokenMapOutput) MapIndex(k pulumi.StringInput) ClusterAgentTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterAgentToken {
		return vs[0].(map[string]*ClusterAgentToken)[vs[1].(string)]
	}).(ClusterAgentTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterAgentTokenInput)(nil)).Elem(), &ClusterAgentToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterAgentTokenArrayInput)(nil)).Elem(), ClusterAgentTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterAgentTokenMapInput)(nil)).Elem(), ClusterAgentTokenMap{})
	pulumi.RegisterOutputType(ClusterAgentTokenOutput{})
	pulumi.RegisterOutputType(ClusterAgentTokenArrayOutput{})
	pulumi.RegisterOutputType(ClusterAgentTokenMapOutput{})
}
