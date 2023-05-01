// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Datasource representing the project, which can then be granted to different organizations or users directly, containing different applications.
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
//			projectProject, err := zitadel.LookupProject(ctx, &zitadel.LookupProjectArgs{
//				OrgId:     data.Zitadel_org.Org.Id,
//				ProjectId: "177073620768522243",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("project", projectProject)
//			return nil
//		})
//	}
//
// ```
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupProjectResult
	err := ctx.Invoke("zitadel:index/getProject:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProject.
type LookupProjectArgs struct {
	// Organization in which the project is located
	OrgId string `pulumi:"orgId"`
	// The ID of this resource.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getProject.
type LookupProjectResult struct {
	// ZITADEL checks if the org of the user has permission to this project
	HasProjectCheck bool `pulumi:"hasProjectCheck"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the project
	Name string `pulumi:"name"`
	// Organization in which the project is located
	OrgId string `pulumi:"orgId"`
	// Defines from where the private labeling should be triggered
	PrivateLabelingSetting string `pulumi:"privateLabelingSetting"`
	// The ID of this resource.
	ProjectId string `pulumi:"projectId"`
	// describes if roles of user should be added in token
	ProjectRoleAssertion bool `pulumi:"projectRoleAssertion"`
	// ZITADEL checks if the user has at least one on this project
	ProjectRoleCheck bool `pulumi:"projectRoleCheck"`
	// State of the project
	State string `pulumi:"state"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectResult, error) {
			args := v.(LookupProjectArgs)
			r, err := LookupProject(ctx, &args, opts...)
			var s LookupProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectResultOutput)
}

// A collection of arguments for invoking getProject.
type LookupProjectOutputArgs struct {
	// Organization in which the project is located
	OrgId pulumi.StringInput `pulumi:"orgId"`
	// The ID of this resource.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

// A collection of values returned by getProject.
type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

// ZITADEL checks if the org of the user has permission to this project
func (o LookupProjectResultOutput) HasProjectCheck() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.HasProjectCheck }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the project
func (o LookupProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

// Organization in which the project is located
func (o LookupProjectResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// Defines from where the private labeling should be triggered
func (o LookupProjectResultOutput) PrivateLabelingSetting() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.PrivateLabelingSetting }).(pulumi.StringOutput)
}

// The ID of this resource.
func (o LookupProjectResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// describes if roles of user should be added in token
func (o LookupProjectResultOutput) ProjectRoleAssertion() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.ProjectRoleAssertion }).(pulumi.BoolOutput)
}

// ZITADEL checks if the user has at least one on this project
func (o LookupProjectResultOutput) ProjectRoleCheck() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.ProjectRoleCheck }).(pulumi.BoolOutput)
}

// State of the project
func (o LookupProjectResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
