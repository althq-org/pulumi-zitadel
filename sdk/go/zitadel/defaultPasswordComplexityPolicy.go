// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing the default password complexity policy.
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
//			_, err := zitadel.NewDefaultPasswordComplexityPolicy(ctx, "passwordComplexityPolicy", &zitadel.DefaultPasswordComplexityPolicyArgs{
//				HasLowercase: pulumi.Bool(true),
//				HasNumber:    pulumi.Bool(true),
//				HasSymbol:    pulumi.Bool(true),
//				HasUppercase: pulumi.Bool(true),
//				MinLength:    pulumi.Int(8),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DefaultPasswordComplexityPolicy struct {
	pulumi.CustomResourceState

	// defines if the password MUST contain a lower case letter
	HasLowercase pulumi.BoolOutput `pulumi:"hasLowercase"`
	// defines if the password MUST contain a number
	HasNumber pulumi.BoolOutput `pulumi:"hasNumber"`
	// defines if the password MUST contain a symbol. E.g. "$"
	HasSymbol pulumi.BoolOutput `pulumi:"hasSymbol"`
	// defines if the password MUST contain an upper case letter
	HasUppercase pulumi.BoolOutput `pulumi:"hasUppercase"`
	// Minimal length for the password
	MinLength pulumi.IntOutput `pulumi:"minLength"`
}

// NewDefaultPasswordComplexityPolicy registers a new resource with the given unique name, arguments, and options.
func NewDefaultPasswordComplexityPolicy(ctx *pulumi.Context,
	name string, args *DefaultPasswordComplexityPolicyArgs, opts ...pulumi.ResourceOption) (*DefaultPasswordComplexityPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HasLowercase == nil {
		return nil, errors.New("invalid value for required argument 'HasLowercase'")
	}
	if args.HasNumber == nil {
		return nil, errors.New("invalid value for required argument 'HasNumber'")
	}
	if args.HasSymbol == nil {
		return nil, errors.New("invalid value for required argument 'HasSymbol'")
	}
	if args.HasUppercase == nil {
		return nil, errors.New("invalid value for required argument 'HasUppercase'")
	}
	if args.MinLength == nil {
		return nil, errors.New("invalid value for required argument 'MinLength'")
	}
	var resource DefaultPasswordComplexityPolicy
	err := ctx.RegisterResource("zitadel:index/defaultPasswordComplexityPolicy:DefaultPasswordComplexityPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultPasswordComplexityPolicy gets an existing DefaultPasswordComplexityPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultPasswordComplexityPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultPasswordComplexityPolicyState, opts ...pulumi.ResourceOption) (*DefaultPasswordComplexityPolicy, error) {
	var resource DefaultPasswordComplexityPolicy
	err := ctx.ReadResource("zitadel:index/defaultPasswordComplexityPolicy:DefaultPasswordComplexityPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultPasswordComplexityPolicy resources.
type defaultPasswordComplexityPolicyState struct {
	// defines if the password MUST contain a lower case letter
	HasLowercase *bool `pulumi:"hasLowercase"`
	// defines if the password MUST contain a number
	HasNumber *bool `pulumi:"hasNumber"`
	// defines if the password MUST contain a symbol. E.g. "$"
	HasSymbol *bool `pulumi:"hasSymbol"`
	// defines if the password MUST contain an upper case letter
	HasUppercase *bool `pulumi:"hasUppercase"`
	// Minimal length for the password
	MinLength *int `pulumi:"minLength"`
}

type DefaultPasswordComplexityPolicyState struct {
	// defines if the password MUST contain a lower case letter
	HasLowercase pulumi.BoolPtrInput
	// defines if the password MUST contain a number
	HasNumber pulumi.BoolPtrInput
	// defines if the password MUST contain a symbol. E.g. "$"
	HasSymbol pulumi.BoolPtrInput
	// defines if the password MUST contain an upper case letter
	HasUppercase pulumi.BoolPtrInput
	// Minimal length for the password
	MinLength pulumi.IntPtrInput
}

func (DefaultPasswordComplexityPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultPasswordComplexityPolicyState)(nil)).Elem()
}

type defaultPasswordComplexityPolicyArgs struct {
	// defines if the password MUST contain a lower case letter
	HasLowercase bool `pulumi:"hasLowercase"`
	// defines if the password MUST contain a number
	HasNumber bool `pulumi:"hasNumber"`
	// defines if the password MUST contain a symbol. E.g. "$"
	HasSymbol bool `pulumi:"hasSymbol"`
	// defines if the password MUST contain an upper case letter
	HasUppercase bool `pulumi:"hasUppercase"`
	// Minimal length for the password
	MinLength int `pulumi:"minLength"`
}

// The set of arguments for constructing a DefaultPasswordComplexityPolicy resource.
type DefaultPasswordComplexityPolicyArgs struct {
	// defines if the password MUST contain a lower case letter
	HasLowercase pulumi.BoolInput
	// defines if the password MUST contain a number
	HasNumber pulumi.BoolInput
	// defines if the password MUST contain a symbol. E.g. "$"
	HasSymbol pulumi.BoolInput
	// defines if the password MUST contain an upper case letter
	HasUppercase pulumi.BoolInput
	// Minimal length for the password
	MinLength pulumi.IntInput
}

func (DefaultPasswordComplexityPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultPasswordComplexityPolicyArgs)(nil)).Elem()
}

type DefaultPasswordComplexityPolicyInput interface {
	pulumi.Input

	ToDefaultPasswordComplexityPolicyOutput() DefaultPasswordComplexityPolicyOutput
	ToDefaultPasswordComplexityPolicyOutputWithContext(ctx context.Context) DefaultPasswordComplexityPolicyOutput
}

func (*DefaultPasswordComplexityPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultPasswordComplexityPolicy)(nil)).Elem()
}

func (i *DefaultPasswordComplexityPolicy) ToDefaultPasswordComplexityPolicyOutput() DefaultPasswordComplexityPolicyOutput {
	return i.ToDefaultPasswordComplexityPolicyOutputWithContext(context.Background())
}

func (i *DefaultPasswordComplexityPolicy) ToDefaultPasswordComplexityPolicyOutputWithContext(ctx context.Context) DefaultPasswordComplexityPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultPasswordComplexityPolicyOutput)
}

// DefaultPasswordComplexityPolicyArrayInput is an input type that accepts DefaultPasswordComplexityPolicyArray and DefaultPasswordComplexityPolicyArrayOutput values.
// You can construct a concrete instance of `DefaultPasswordComplexityPolicyArrayInput` via:
//
//	DefaultPasswordComplexityPolicyArray{ DefaultPasswordComplexityPolicyArgs{...} }
type DefaultPasswordComplexityPolicyArrayInput interface {
	pulumi.Input

	ToDefaultPasswordComplexityPolicyArrayOutput() DefaultPasswordComplexityPolicyArrayOutput
	ToDefaultPasswordComplexityPolicyArrayOutputWithContext(context.Context) DefaultPasswordComplexityPolicyArrayOutput
}

type DefaultPasswordComplexityPolicyArray []DefaultPasswordComplexityPolicyInput

func (DefaultPasswordComplexityPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultPasswordComplexityPolicy)(nil)).Elem()
}

func (i DefaultPasswordComplexityPolicyArray) ToDefaultPasswordComplexityPolicyArrayOutput() DefaultPasswordComplexityPolicyArrayOutput {
	return i.ToDefaultPasswordComplexityPolicyArrayOutputWithContext(context.Background())
}

func (i DefaultPasswordComplexityPolicyArray) ToDefaultPasswordComplexityPolicyArrayOutputWithContext(ctx context.Context) DefaultPasswordComplexityPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultPasswordComplexityPolicyArrayOutput)
}

// DefaultPasswordComplexityPolicyMapInput is an input type that accepts DefaultPasswordComplexityPolicyMap and DefaultPasswordComplexityPolicyMapOutput values.
// You can construct a concrete instance of `DefaultPasswordComplexityPolicyMapInput` via:
//
//	DefaultPasswordComplexityPolicyMap{ "key": DefaultPasswordComplexityPolicyArgs{...} }
type DefaultPasswordComplexityPolicyMapInput interface {
	pulumi.Input

	ToDefaultPasswordComplexityPolicyMapOutput() DefaultPasswordComplexityPolicyMapOutput
	ToDefaultPasswordComplexityPolicyMapOutputWithContext(context.Context) DefaultPasswordComplexityPolicyMapOutput
}

type DefaultPasswordComplexityPolicyMap map[string]DefaultPasswordComplexityPolicyInput

func (DefaultPasswordComplexityPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultPasswordComplexityPolicy)(nil)).Elem()
}

func (i DefaultPasswordComplexityPolicyMap) ToDefaultPasswordComplexityPolicyMapOutput() DefaultPasswordComplexityPolicyMapOutput {
	return i.ToDefaultPasswordComplexityPolicyMapOutputWithContext(context.Background())
}

func (i DefaultPasswordComplexityPolicyMap) ToDefaultPasswordComplexityPolicyMapOutputWithContext(ctx context.Context) DefaultPasswordComplexityPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultPasswordComplexityPolicyMapOutput)
}

type DefaultPasswordComplexityPolicyOutput struct{ *pulumi.OutputState }

func (DefaultPasswordComplexityPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultPasswordComplexityPolicy)(nil)).Elem()
}

func (o DefaultPasswordComplexityPolicyOutput) ToDefaultPasswordComplexityPolicyOutput() DefaultPasswordComplexityPolicyOutput {
	return o
}

func (o DefaultPasswordComplexityPolicyOutput) ToDefaultPasswordComplexityPolicyOutputWithContext(ctx context.Context) DefaultPasswordComplexityPolicyOutput {
	return o
}

// defines if the password MUST contain a lower case letter
func (o DefaultPasswordComplexityPolicyOutput) HasLowercase() pulumi.BoolOutput {
	return o.ApplyT(func(v *DefaultPasswordComplexityPolicy) pulumi.BoolOutput { return v.HasLowercase }).(pulumi.BoolOutput)
}

// defines if the password MUST contain a number
func (o DefaultPasswordComplexityPolicyOutput) HasNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v *DefaultPasswordComplexityPolicy) pulumi.BoolOutput { return v.HasNumber }).(pulumi.BoolOutput)
}

// defines if the password MUST contain a symbol. E.g. "$"
func (o DefaultPasswordComplexityPolicyOutput) HasSymbol() pulumi.BoolOutput {
	return o.ApplyT(func(v *DefaultPasswordComplexityPolicy) pulumi.BoolOutput { return v.HasSymbol }).(pulumi.BoolOutput)
}

// defines if the password MUST contain an upper case letter
func (o DefaultPasswordComplexityPolicyOutput) HasUppercase() pulumi.BoolOutput {
	return o.ApplyT(func(v *DefaultPasswordComplexityPolicy) pulumi.BoolOutput { return v.HasUppercase }).(pulumi.BoolOutput)
}

// Minimal length for the password
func (o DefaultPasswordComplexityPolicyOutput) MinLength() pulumi.IntOutput {
	return o.ApplyT(func(v *DefaultPasswordComplexityPolicy) pulumi.IntOutput { return v.MinLength }).(pulumi.IntOutput)
}

type DefaultPasswordComplexityPolicyArrayOutput struct{ *pulumi.OutputState }

func (DefaultPasswordComplexityPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultPasswordComplexityPolicy)(nil)).Elem()
}

func (o DefaultPasswordComplexityPolicyArrayOutput) ToDefaultPasswordComplexityPolicyArrayOutput() DefaultPasswordComplexityPolicyArrayOutput {
	return o
}

func (o DefaultPasswordComplexityPolicyArrayOutput) ToDefaultPasswordComplexityPolicyArrayOutputWithContext(ctx context.Context) DefaultPasswordComplexityPolicyArrayOutput {
	return o
}

func (o DefaultPasswordComplexityPolicyArrayOutput) Index(i pulumi.IntInput) DefaultPasswordComplexityPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DefaultPasswordComplexityPolicy {
		return vs[0].([]*DefaultPasswordComplexityPolicy)[vs[1].(int)]
	}).(DefaultPasswordComplexityPolicyOutput)
}

type DefaultPasswordComplexityPolicyMapOutput struct{ *pulumi.OutputState }

func (DefaultPasswordComplexityPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultPasswordComplexityPolicy)(nil)).Elem()
}

func (o DefaultPasswordComplexityPolicyMapOutput) ToDefaultPasswordComplexityPolicyMapOutput() DefaultPasswordComplexityPolicyMapOutput {
	return o
}

func (o DefaultPasswordComplexityPolicyMapOutput) ToDefaultPasswordComplexityPolicyMapOutputWithContext(ctx context.Context) DefaultPasswordComplexityPolicyMapOutput {
	return o
}

func (o DefaultPasswordComplexityPolicyMapOutput) MapIndex(k pulumi.StringInput) DefaultPasswordComplexityPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DefaultPasswordComplexityPolicy {
		return vs[0].(map[string]*DefaultPasswordComplexityPolicy)[vs[1].(string)]
	}).(DefaultPasswordComplexityPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultPasswordComplexityPolicyInput)(nil)).Elem(), &DefaultPasswordComplexityPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultPasswordComplexityPolicyArrayInput)(nil)).Elem(), DefaultPasswordComplexityPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultPasswordComplexityPolicyMapInput)(nil)).Elem(), DefaultPasswordComplexityPolicyMap{})
	pulumi.RegisterOutputType(DefaultPasswordComplexityPolicyOutput{})
	pulumi.RegisterOutputType(DefaultPasswordComplexityPolicyArrayOutput{})
	pulumi.RegisterOutputType(DefaultPasswordComplexityPolicyMapOutput{})
}
