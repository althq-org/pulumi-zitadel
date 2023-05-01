// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing an API application belonging to a project, with all configuration possibilities.
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
//			_, err := zitadel.NewApplicationApi(ctx, "applicationApi", &zitadel.ApplicationApiArgs{
//				OrgId:          pulumi.Any(zitadel_org.Org.Id),
//				ProjectId:      pulumi.Any(zitadel_project.Project.Id),
//				AuthMethodType: pulumi.String("API_AUTH_METHOD_TYPE_PRIVATE_KEY_JWT"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ApplicationApi struct {
	pulumi.CustomResourceState

	// Auth method type, supported values: API*AUTH*METHOD*TYPE*BASIC, API*AUTH*METHOD*TYPE*PRIVATE*KEY*JWT
	AuthMethodType pulumi.StringPtrOutput `pulumi:"authMethodType"`
	// generated ID for this config
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// generated secret for this config
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// Name of the application
	Name pulumi.StringOutput `pulumi:"name"`
	// orgID of the application
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// ID of the project
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
}

// NewApplicationApi registers a new resource with the given unique name, arguments, and options.
func NewApplicationApi(ctx *pulumi.Context,
	name string, args *ApplicationApiArgs, opts ...pulumi.ResourceOption) (*ApplicationApi, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"clientId",
		"clientSecret",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource ApplicationApi
	err := ctx.RegisterResource("zitadel:index/applicationApi:ApplicationApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationApi gets an existing ApplicationApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationApiState, opts ...pulumi.ResourceOption) (*ApplicationApi, error) {
	var resource ApplicationApi
	err := ctx.ReadResource("zitadel:index/applicationApi:ApplicationApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationApi resources.
type applicationApiState struct {
	// Auth method type, supported values: API*AUTH*METHOD*TYPE*BASIC, API*AUTH*METHOD*TYPE*PRIVATE*KEY*JWT
	AuthMethodType *string `pulumi:"authMethodType"`
	// generated ID for this config
	ClientId *string `pulumi:"clientId"`
	// generated secret for this config
	ClientSecret *string `pulumi:"clientSecret"`
	// Name of the application
	Name *string `pulumi:"name"`
	// orgID of the application
	OrgId *string `pulumi:"orgId"`
	// ID of the project
	ProjectId *string `pulumi:"projectId"`
}

type ApplicationApiState struct {
	// Auth method type, supported values: API*AUTH*METHOD*TYPE*BASIC, API*AUTH*METHOD*TYPE*PRIVATE*KEY*JWT
	AuthMethodType pulumi.StringPtrInput
	// generated ID for this config
	ClientId pulumi.StringPtrInput
	// generated secret for this config
	ClientSecret pulumi.StringPtrInput
	// Name of the application
	Name pulumi.StringPtrInput
	// orgID of the application
	OrgId pulumi.StringPtrInput
	// ID of the project
	ProjectId pulumi.StringPtrInput
}

func (ApplicationApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationApiState)(nil)).Elem()
}

type applicationApiArgs struct {
	// Auth method type, supported values: API*AUTH*METHOD*TYPE*BASIC, API*AUTH*METHOD*TYPE*PRIVATE*KEY*JWT
	AuthMethodType *string `pulumi:"authMethodType"`
	// Name of the application
	Name *string `pulumi:"name"`
	// orgID of the application
	OrgId string `pulumi:"orgId"`
	// ID of the project
	ProjectId string `pulumi:"projectId"`
}

// The set of arguments for constructing a ApplicationApi resource.
type ApplicationApiArgs struct {
	// Auth method type, supported values: API*AUTH*METHOD*TYPE*BASIC, API*AUTH*METHOD*TYPE*PRIVATE*KEY*JWT
	AuthMethodType pulumi.StringPtrInput
	// Name of the application
	Name pulumi.StringPtrInput
	// orgID of the application
	OrgId pulumi.StringInput
	// ID of the project
	ProjectId pulumi.StringInput
}

func (ApplicationApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationApiArgs)(nil)).Elem()
}

type ApplicationApiInput interface {
	pulumi.Input

	ToApplicationApiOutput() ApplicationApiOutput
	ToApplicationApiOutputWithContext(ctx context.Context) ApplicationApiOutput
}

func (*ApplicationApi) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationApi)(nil)).Elem()
}

func (i *ApplicationApi) ToApplicationApiOutput() ApplicationApiOutput {
	return i.ToApplicationApiOutputWithContext(context.Background())
}

func (i *ApplicationApi) ToApplicationApiOutputWithContext(ctx context.Context) ApplicationApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationApiOutput)
}

// ApplicationApiArrayInput is an input type that accepts ApplicationApiArray and ApplicationApiArrayOutput values.
// You can construct a concrete instance of `ApplicationApiArrayInput` via:
//
//	ApplicationApiArray{ ApplicationApiArgs{...} }
type ApplicationApiArrayInput interface {
	pulumi.Input

	ToApplicationApiArrayOutput() ApplicationApiArrayOutput
	ToApplicationApiArrayOutputWithContext(context.Context) ApplicationApiArrayOutput
}

type ApplicationApiArray []ApplicationApiInput

func (ApplicationApiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationApi)(nil)).Elem()
}

func (i ApplicationApiArray) ToApplicationApiArrayOutput() ApplicationApiArrayOutput {
	return i.ToApplicationApiArrayOutputWithContext(context.Background())
}

func (i ApplicationApiArray) ToApplicationApiArrayOutputWithContext(ctx context.Context) ApplicationApiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationApiArrayOutput)
}

// ApplicationApiMapInput is an input type that accepts ApplicationApiMap and ApplicationApiMapOutput values.
// You can construct a concrete instance of `ApplicationApiMapInput` via:
//
//	ApplicationApiMap{ "key": ApplicationApiArgs{...} }
type ApplicationApiMapInput interface {
	pulumi.Input

	ToApplicationApiMapOutput() ApplicationApiMapOutput
	ToApplicationApiMapOutputWithContext(context.Context) ApplicationApiMapOutput
}

type ApplicationApiMap map[string]ApplicationApiInput

func (ApplicationApiMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationApi)(nil)).Elem()
}

func (i ApplicationApiMap) ToApplicationApiMapOutput() ApplicationApiMapOutput {
	return i.ToApplicationApiMapOutputWithContext(context.Background())
}

func (i ApplicationApiMap) ToApplicationApiMapOutputWithContext(ctx context.Context) ApplicationApiMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationApiMapOutput)
}

type ApplicationApiOutput struct{ *pulumi.OutputState }

func (ApplicationApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationApi)(nil)).Elem()
}

func (o ApplicationApiOutput) ToApplicationApiOutput() ApplicationApiOutput {
	return o
}

func (o ApplicationApiOutput) ToApplicationApiOutputWithContext(ctx context.Context) ApplicationApiOutput {
	return o
}

// Auth method type, supported values: API*AUTH*METHOD*TYPE*BASIC, API*AUTH*METHOD*TYPE*PRIVATE*KEY*JWT
func (o ApplicationApiOutput) AuthMethodType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationApi) pulumi.StringPtrOutput { return v.AuthMethodType }).(pulumi.StringPtrOutput)
}

// generated ID for this config
func (o ApplicationApiOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationApi) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// generated secret for this config
func (o ApplicationApiOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationApi) pulumi.StringOutput { return v.ClientSecret }).(pulumi.StringOutput)
}

// Name of the application
func (o ApplicationApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationApi) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// orgID of the application
func (o ApplicationApiOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationApi) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// ID of the project
func (o ApplicationApiOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationApi) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

type ApplicationApiArrayOutput struct{ *pulumi.OutputState }

func (ApplicationApiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationApi)(nil)).Elem()
}

func (o ApplicationApiArrayOutput) ToApplicationApiArrayOutput() ApplicationApiArrayOutput {
	return o
}

func (o ApplicationApiArrayOutput) ToApplicationApiArrayOutputWithContext(ctx context.Context) ApplicationApiArrayOutput {
	return o
}

func (o ApplicationApiArrayOutput) Index(i pulumi.IntInput) ApplicationApiOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationApi {
		return vs[0].([]*ApplicationApi)[vs[1].(int)]
	}).(ApplicationApiOutput)
}

type ApplicationApiMapOutput struct{ *pulumi.OutputState }

func (ApplicationApiMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationApi)(nil)).Elem()
}

func (o ApplicationApiMapOutput) ToApplicationApiMapOutput() ApplicationApiMapOutput {
	return o
}

func (o ApplicationApiMapOutput) ToApplicationApiMapOutputWithContext(ctx context.Context) ApplicationApiMapOutput {
	return o
}

func (o ApplicationApiMapOutput) MapIndex(k pulumi.StringInput) ApplicationApiOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationApi {
		return vs[0].(map[string]*ApplicationApi)[vs[1].(string)]
	}).(ApplicationApiOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationApiInput)(nil)).Elem(), &ApplicationApi{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationApiArrayInput)(nil)).Elem(), ApplicationApiArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationApiMapInput)(nil)).Elem(), ApplicationApiMap{})
	pulumi.RegisterOutputType(ApplicationApiOutput{})
	pulumi.RegisterOutputType(ApplicationApiArrayOutput{})
	pulumi.RegisterOutputType(ApplicationApiMapOutput{})
}
