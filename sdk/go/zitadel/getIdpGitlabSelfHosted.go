// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Datasource representing a GitLab Self Hosted IDP on the instance.
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
//			_, err := zitadel.LookupIdpGitlabSelfHosted(ctx, &zitadel.LookupIdpGitlabSelfHostedArgs{
//				Id: "177073614158299139",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupIdpGitlabSelfHosted(ctx *pulumi.Context, args *LookupIdpGitlabSelfHostedArgs, opts ...pulumi.InvokeOption) (*LookupIdpGitlabSelfHostedResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupIdpGitlabSelfHostedResult
	err := ctx.Invoke("zitadel:index/getIdpGitlabSelfHosted:getIdpGitlabSelfHosted", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIdpGitlabSelfHosted.
type LookupIdpGitlabSelfHostedArgs struct {
	// The ID of this resource.
	Id string `pulumi:"id"`
}

// A collection of values returned by getIdpGitlabSelfHosted.
type LookupIdpGitlabSelfHostedResult struct {
	// client id generated by the identity provider
	ClientId string `pulumi:"clientId"`
	// client secret generated by the identity provider
	ClientSecret string `pulumi:"clientSecret"`
	// The ID of this resource.
	Id string `pulumi:"id"`
	// enabled if a new account in ZITADEL are created automatically on login with an external account
	IsAutoCreation bool `pulumi:"isAutoCreation"`
	// enabled if a the ZITADEL account fields are updated automatically on each login
	IsAutoUpdate bool `pulumi:"isAutoUpdate"`
	// enabled if users are able to create a new account in ZITADEL when using an external account
	IsCreationAllowed bool `pulumi:"isCreationAllowed"`
	// enabled if users are able to link an existing ZITADEL user with an external account
	IsLinkingAllowed bool `pulumi:"isLinkingAllowed"`
	// the providers issuer
	Issuer string `pulumi:"issuer"`
	// Name of the IDP
	Name string `pulumi:"name"`
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes []string `pulumi:"scopes"`
}

func LookupIdpGitlabSelfHostedOutput(ctx *pulumi.Context, args LookupIdpGitlabSelfHostedOutputArgs, opts ...pulumi.InvokeOption) LookupIdpGitlabSelfHostedResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIdpGitlabSelfHostedResult, error) {
			args := v.(LookupIdpGitlabSelfHostedArgs)
			r, err := LookupIdpGitlabSelfHosted(ctx, &args, opts...)
			var s LookupIdpGitlabSelfHostedResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIdpGitlabSelfHostedResultOutput)
}

// A collection of arguments for invoking getIdpGitlabSelfHosted.
type LookupIdpGitlabSelfHostedOutputArgs struct {
	// The ID of this resource.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupIdpGitlabSelfHostedOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdpGitlabSelfHostedArgs)(nil)).Elem()
}

// A collection of values returned by getIdpGitlabSelfHosted.
type LookupIdpGitlabSelfHostedResultOutput struct{ *pulumi.OutputState }

func (LookupIdpGitlabSelfHostedResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdpGitlabSelfHostedResult)(nil)).Elem()
}

func (o LookupIdpGitlabSelfHostedResultOutput) ToLookupIdpGitlabSelfHostedResultOutput() LookupIdpGitlabSelfHostedResultOutput {
	return o
}

func (o LookupIdpGitlabSelfHostedResultOutput) ToLookupIdpGitlabSelfHostedResultOutputWithContext(ctx context.Context) LookupIdpGitlabSelfHostedResultOutput {
	return o
}

// client id generated by the identity provider
func (o LookupIdpGitlabSelfHostedResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpGitlabSelfHostedResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// client secret generated by the identity provider
func (o LookupIdpGitlabSelfHostedResultOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpGitlabSelfHostedResult) string { return v.ClientSecret }).(pulumi.StringOutput)
}

// The ID of this resource.
func (o LookupIdpGitlabSelfHostedResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpGitlabSelfHostedResult) string { return v.Id }).(pulumi.StringOutput)
}

// enabled if a new account in ZITADEL are created automatically on login with an external account
func (o LookupIdpGitlabSelfHostedResultOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIdpGitlabSelfHostedResult) bool { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enabled if a the ZITADEL account fields are updated automatically on each login
func (o LookupIdpGitlabSelfHostedResultOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIdpGitlabSelfHostedResult) bool { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enabled if users are able to create a new account in ZITADEL when using an external account
func (o LookupIdpGitlabSelfHostedResultOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIdpGitlabSelfHostedResult) bool { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enabled if users are able to link an existing ZITADEL user with an external account
func (o LookupIdpGitlabSelfHostedResultOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIdpGitlabSelfHostedResult) bool { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// the providers issuer
func (o LookupIdpGitlabSelfHostedResultOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpGitlabSelfHostedResult) string { return v.Issuer }).(pulumi.StringOutput)
}

// Name of the IDP
func (o LookupIdpGitlabSelfHostedResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpGitlabSelfHostedResult) string { return v.Name }).(pulumi.StringOutput)
}

// the scopes requested by ZITADEL during the request on the identity provider
func (o LookupIdpGitlabSelfHostedResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIdpGitlabSelfHostedResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIdpGitlabSelfHostedResultOutput{})
}
