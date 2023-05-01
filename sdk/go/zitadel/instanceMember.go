// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing the membership of a user on an instance, defined with the given role.
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
//			_, err := zitadel.NewInstanceMember(ctx, "instanceMember", &zitadel.InstanceMemberArgs{
//				UserId: pulumi.Any(zitadel_human_user.Human_user.Id),
//				Roles: pulumi.StringArray{
//					pulumi.String("IAM_OWNER"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type InstanceMember struct {
	pulumi.CustomResourceState

	// List of roles granted, full list available here: https://zitadel.com/docs/guides/manage/console/managers#roles
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// ID of the user
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewInstanceMember registers a new resource with the given unique name, arguments, and options.
func NewInstanceMember(ctx *pulumi.Context,
	name string, args *InstanceMemberArgs, opts ...pulumi.ResourceOption) (*InstanceMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource InstanceMember
	err := ctx.RegisterResource("zitadel:index/instanceMember:InstanceMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceMember gets an existing InstanceMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceMemberState, opts ...pulumi.ResourceOption) (*InstanceMember, error) {
	var resource InstanceMember
	err := ctx.ReadResource("zitadel:index/instanceMember:InstanceMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceMember resources.
type instanceMemberState struct {
	// List of roles granted, full list available here: https://zitadel.com/docs/guides/manage/console/managers#roles
	Roles []string `pulumi:"roles"`
	// ID of the user
	UserId *string `pulumi:"userId"`
}

type InstanceMemberState struct {
	// List of roles granted, full list available here: https://zitadel.com/docs/guides/manage/console/managers#roles
	Roles pulumi.StringArrayInput
	// ID of the user
	UserId pulumi.StringPtrInput
}

func (InstanceMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceMemberState)(nil)).Elem()
}

type instanceMemberArgs struct {
	// List of roles granted, full list available here: https://zitadel.com/docs/guides/manage/console/managers#roles
	Roles []string `pulumi:"roles"`
	// ID of the user
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a InstanceMember resource.
type InstanceMemberArgs struct {
	// List of roles granted, full list available here: https://zitadel.com/docs/guides/manage/console/managers#roles
	Roles pulumi.StringArrayInput
	// ID of the user
	UserId pulumi.StringInput
}

func (InstanceMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceMemberArgs)(nil)).Elem()
}

type InstanceMemberInput interface {
	pulumi.Input

	ToInstanceMemberOutput() InstanceMemberOutput
	ToInstanceMemberOutputWithContext(ctx context.Context) InstanceMemberOutput
}

func (*InstanceMember) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceMember)(nil)).Elem()
}

func (i *InstanceMember) ToInstanceMemberOutput() InstanceMemberOutput {
	return i.ToInstanceMemberOutputWithContext(context.Background())
}

func (i *InstanceMember) ToInstanceMemberOutputWithContext(ctx context.Context) InstanceMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMemberOutput)
}

// InstanceMemberArrayInput is an input type that accepts InstanceMemberArray and InstanceMemberArrayOutput values.
// You can construct a concrete instance of `InstanceMemberArrayInput` via:
//
//	InstanceMemberArray{ InstanceMemberArgs{...} }
type InstanceMemberArrayInput interface {
	pulumi.Input

	ToInstanceMemberArrayOutput() InstanceMemberArrayOutput
	ToInstanceMemberArrayOutputWithContext(context.Context) InstanceMemberArrayOutput
}

type InstanceMemberArray []InstanceMemberInput

func (InstanceMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceMember)(nil)).Elem()
}

func (i InstanceMemberArray) ToInstanceMemberArrayOutput() InstanceMemberArrayOutput {
	return i.ToInstanceMemberArrayOutputWithContext(context.Background())
}

func (i InstanceMemberArray) ToInstanceMemberArrayOutputWithContext(ctx context.Context) InstanceMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMemberArrayOutput)
}

// InstanceMemberMapInput is an input type that accepts InstanceMemberMap and InstanceMemberMapOutput values.
// You can construct a concrete instance of `InstanceMemberMapInput` via:
//
//	InstanceMemberMap{ "key": InstanceMemberArgs{...} }
type InstanceMemberMapInput interface {
	pulumi.Input

	ToInstanceMemberMapOutput() InstanceMemberMapOutput
	ToInstanceMemberMapOutputWithContext(context.Context) InstanceMemberMapOutput
}

type InstanceMemberMap map[string]InstanceMemberInput

func (InstanceMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceMember)(nil)).Elem()
}

func (i InstanceMemberMap) ToInstanceMemberMapOutput() InstanceMemberMapOutput {
	return i.ToInstanceMemberMapOutputWithContext(context.Background())
}

func (i InstanceMemberMap) ToInstanceMemberMapOutputWithContext(ctx context.Context) InstanceMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMemberMapOutput)
}

type InstanceMemberOutput struct{ *pulumi.OutputState }

func (InstanceMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceMember)(nil)).Elem()
}

func (o InstanceMemberOutput) ToInstanceMemberOutput() InstanceMemberOutput {
	return o
}

func (o InstanceMemberOutput) ToInstanceMemberOutputWithContext(ctx context.Context) InstanceMemberOutput {
	return o
}

// List of roles granted, full list available here: https://zitadel.com/docs/guides/manage/console/managers#roles
func (o InstanceMemberOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstanceMember) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

// ID of the user
func (o InstanceMemberOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceMember) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

type InstanceMemberArrayOutput struct{ *pulumi.OutputState }

func (InstanceMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceMember)(nil)).Elem()
}

func (o InstanceMemberArrayOutput) ToInstanceMemberArrayOutput() InstanceMemberArrayOutput {
	return o
}

func (o InstanceMemberArrayOutput) ToInstanceMemberArrayOutputWithContext(ctx context.Context) InstanceMemberArrayOutput {
	return o
}

func (o InstanceMemberArrayOutput) Index(i pulumi.IntInput) InstanceMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceMember {
		return vs[0].([]*InstanceMember)[vs[1].(int)]
	}).(InstanceMemberOutput)
}

type InstanceMemberMapOutput struct{ *pulumi.OutputState }

func (InstanceMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceMember)(nil)).Elem()
}

func (o InstanceMemberMapOutput) ToInstanceMemberMapOutput() InstanceMemberMapOutput {
	return o
}

func (o InstanceMemberMapOutput) ToInstanceMemberMapOutputWithContext(ctx context.Context) InstanceMemberMapOutput {
	return o
}

func (o InstanceMemberMapOutput) MapIndex(k pulumi.StringInput) InstanceMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceMember {
		return vs[0].(map[string]*InstanceMember)[vs[1].(string)]
	}).(InstanceMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMemberInput)(nil)).Elem(), &InstanceMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMemberArrayInput)(nil)).Elem(), InstanceMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMemberMapInput)(nil)).Elem(), InstanceMemberMap{})
	pulumi.RegisterOutputType(InstanceMemberOutput{})
	pulumi.RegisterOutputType(InstanceMemberArrayOutput{})
	pulumi.RegisterOutputType(InstanceMemberMapOutput{})
}
