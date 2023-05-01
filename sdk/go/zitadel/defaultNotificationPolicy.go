// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing the default notification policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-zitadel/sdk/go/zitadel"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := zitadel.NewDefaultNotificationPolicy(ctx, "notificationPolicy", &zitadel.DefaultNotificationPolicyArgs{
//				PasswordChange: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DefaultNotificationPolicy struct {
	pulumi.CustomResourceState

	// Send notification if a user changes his password
	PasswordChange pulumi.BoolOutput `pulumi:"passwordChange"`
}

// NewDefaultNotificationPolicy registers a new resource with the given unique name, arguments, and options.
func NewDefaultNotificationPolicy(ctx *pulumi.Context,
	name string, args *DefaultNotificationPolicyArgs, opts ...pulumi.ResourceOption) (*DefaultNotificationPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PasswordChange == nil {
		return nil, errors.New("invalid value for required argument 'PasswordChange'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DefaultNotificationPolicy
	err := ctx.RegisterResource("zitadel:index/defaultNotificationPolicy:DefaultNotificationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultNotificationPolicy gets an existing DefaultNotificationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultNotificationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultNotificationPolicyState, opts ...pulumi.ResourceOption) (*DefaultNotificationPolicy, error) {
	var resource DefaultNotificationPolicy
	err := ctx.ReadResource("zitadel:index/defaultNotificationPolicy:DefaultNotificationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultNotificationPolicy resources.
type defaultNotificationPolicyState struct {
	// Send notification if a user changes his password
	PasswordChange *bool `pulumi:"passwordChange"`
}

type DefaultNotificationPolicyState struct {
	// Send notification if a user changes his password
	PasswordChange pulumi.BoolPtrInput
}

func (DefaultNotificationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultNotificationPolicyState)(nil)).Elem()
}

type defaultNotificationPolicyArgs struct {
	// Send notification if a user changes his password
	PasswordChange bool `pulumi:"passwordChange"`
}

// The set of arguments for constructing a DefaultNotificationPolicy resource.
type DefaultNotificationPolicyArgs struct {
	// Send notification if a user changes his password
	PasswordChange pulumi.BoolInput
}

func (DefaultNotificationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultNotificationPolicyArgs)(nil)).Elem()
}

type DefaultNotificationPolicyInput interface {
	pulumi.Input

	ToDefaultNotificationPolicyOutput() DefaultNotificationPolicyOutput
	ToDefaultNotificationPolicyOutputWithContext(ctx context.Context) DefaultNotificationPolicyOutput
}

func (*DefaultNotificationPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultNotificationPolicy)(nil)).Elem()
}

func (i *DefaultNotificationPolicy) ToDefaultNotificationPolicyOutput() DefaultNotificationPolicyOutput {
	return i.ToDefaultNotificationPolicyOutputWithContext(context.Background())
}

func (i *DefaultNotificationPolicy) ToDefaultNotificationPolicyOutputWithContext(ctx context.Context) DefaultNotificationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultNotificationPolicyOutput)
}

// DefaultNotificationPolicyArrayInput is an input type that accepts DefaultNotificationPolicyArray and DefaultNotificationPolicyArrayOutput values.
// You can construct a concrete instance of `DefaultNotificationPolicyArrayInput` via:
//
//	DefaultNotificationPolicyArray{ DefaultNotificationPolicyArgs{...} }
type DefaultNotificationPolicyArrayInput interface {
	pulumi.Input

	ToDefaultNotificationPolicyArrayOutput() DefaultNotificationPolicyArrayOutput
	ToDefaultNotificationPolicyArrayOutputWithContext(context.Context) DefaultNotificationPolicyArrayOutput
}

type DefaultNotificationPolicyArray []DefaultNotificationPolicyInput

func (DefaultNotificationPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultNotificationPolicy)(nil)).Elem()
}

func (i DefaultNotificationPolicyArray) ToDefaultNotificationPolicyArrayOutput() DefaultNotificationPolicyArrayOutput {
	return i.ToDefaultNotificationPolicyArrayOutputWithContext(context.Background())
}

func (i DefaultNotificationPolicyArray) ToDefaultNotificationPolicyArrayOutputWithContext(ctx context.Context) DefaultNotificationPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultNotificationPolicyArrayOutput)
}

// DefaultNotificationPolicyMapInput is an input type that accepts DefaultNotificationPolicyMap and DefaultNotificationPolicyMapOutput values.
// You can construct a concrete instance of `DefaultNotificationPolicyMapInput` via:
//
//	DefaultNotificationPolicyMap{ "key": DefaultNotificationPolicyArgs{...} }
type DefaultNotificationPolicyMapInput interface {
	pulumi.Input

	ToDefaultNotificationPolicyMapOutput() DefaultNotificationPolicyMapOutput
	ToDefaultNotificationPolicyMapOutputWithContext(context.Context) DefaultNotificationPolicyMapOutput
}

type DefaultNotificationPolicyMap map[string]DefaultNotificationPolicyInput

func (DefaultNotificationPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultNotificationPolicy)(nil)).Elem()
}

func (i DefaultNotificationPolicyMap) ToDefaultNotificationPolicyMapOutput() DefaultNotificationPolicyMapOutput {
	return i.ToDefaultNotificationPolicyMapOutputWithContext(context.Background())
}

func (i DefaultNotificationPolicyMap) ToDefaultNotificationPolicyMapOutputWithContext(ctx context.Context) DefaultNotificationPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultNotificationPolicyMapOutput)
}

type DefaultNotificationPolicyOutput struct{ *pulumi.OutputState }

func (DefaultNotificationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultNotificationPolicy)(nil)).Elem()
}

func (o DefaultNotificationPolicyOutput) ToDefaultNotificationPolicyOutput() DefaultNotificationPolicyOutput {
	return o
}

func (o DefaultNotificationPolicyOutput) ToDefaultNotificationPolicyOutputWithContext(ctx context.Context) DefaultNotificationPolicyOutput {
	return o
}

// Send notification if a user changes his password
func (o DefaultNotificationPolicyOutput) PasswordChange() pulumi.BoolOutput {
	return o.ApplyT(func(v *DefaultNotificationPolicy) pulumi.BoolOutput { return v.PasswordChange }).(pulumi.BoolOutput)
}

type DefaultNotificationPolicyArrayOutput struct{ *pulumi.OutputState }

func (DefaultNotificationPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultNotificationPolicy)(nil)).Elem()
}

func (o DefaultNotificationPolicyArrayOutput) ToDefaultNotificationPolicyArrayOutput() DefaultNotificationPolicyArrayOutput {
	return o
}

func (o DefaultNotificationPolicyArrayOutput) ToDefaultNotificationPolicyArrayOutputWithContext(ctx context.Context) DefaultNotificationPolicyArrayOutput {
	return o
}

func (o DefaultNotificationPolicyArrayOutput) Index(i pulumi.IntInput) DefaultNotificationPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DefaultNotificationPolicy {
		return vs[0].([]*DefaultNotificationPolicy)[vs[1].(int)]
	}).(DefaultNotificationPolicyOutput)
}

type DefaultNotificationPolicyMapOutput struct{ *pulumi.OutputState }

func (DefaultNotificationPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultNotificationPolicy)(nil)).Elem()
}

func (o DefaultNotificationPolicyMapOutput) ToDefaultNotificationPolicyMapOutput() DefaultNotificationPolicyMapOutput {
	return o
}

func (o DefaultNotificationPolicyMapOutput) ToDefaultNotificationPolicyMapOutputWithContext(ctx context.Context) DefaultNotificationPolicyMapOutput {
	return o
}

func (o DefaultNotificationPolicyMapOutput) MapIndex(k pulumi.StringInput) DefaultNotificationPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DefaultNotificationPolicy {
		return vs[0].(map[string]*DefaultNotificationPolicy)[vs[1].(string)]
	}).(DefaultNotificationPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultNotificationPolicyInput)(nil)).Elem(), &DefaultNotificationPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultNotificationPolicyArrayInput)(nil)).Elem(), DefaultNotificationPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultNotificationPolicyMapInput)(nil)).Elem(), DefaultNotificationPolicyMap{})
	pulumi.RegisterOutputType(DefaultNotificationPolicyOutput{})
	pulumi.RegisterOutputType(DefaultNotificationPolicyArrayOutput{})
	pulumi.RegisterOutputType(DefaultNotificationPolicyMapOutput{})
}
