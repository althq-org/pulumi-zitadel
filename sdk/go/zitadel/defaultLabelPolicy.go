// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing the default label policy.
type DefaultLabelPolicy struct {
	pulumi.CustomResourceState

	// hex value for background color
	BackgroundColor pulumi.StringOutput `pulumi:"backgroundColor"`
	// hex value for background color dark theme
	BackgroundColorDark pulumi.StringOutput `pulumi:"backgroundColorDark"`
	// disable watermark
	DisableWatermark pulumi.BoolOutput `pulumi:"disableWatermark"`
	// hex value for font color
	FontColor pulumi.StringOutput `pulumi:"fontColor"`
	// hex value for font color dark theme
	FontColorDark pulumi.StringOutput    `pulumi:"fontColorDark"`
	FontHash      pulumi.StringPtrOutput `pulumi:"fontHash"`
	FontPath      pulumi.StringPtrOutput `pulumi:"fontPath"`
	FontUrl       pulumi.StringOutput    `pulumi:"fontUrl"`
	// hides the org suffix on the login form if the scope "urn:zitadel:iam:org:domain:primary:{domainname}" is set. Details about this scope in https://docs.zitadel.ch/concepts#Reserved_Scopes
	HideLoginNameSuffix pulumi.BoolOutput      `pulumi:"hideLoginNameSuffix"`
	IconDarkHash        pulumi.StringPtrOutput `pulumi:"iconDarkHash"`
	IconDarkPath        pulumi.StringPtrOutput `pulumi:"iconDarkPath"`
	IconHash            pulumi.StringPtrOutput `pulumi:"iconHash"`
	IconPath            pulumi.StringPtrOutput `pulumi:"iconPath"`
	IconUrl             pulumi.StringOutput    `pulumi:"iconUrl"`
	IconUrlDark         pulumi.StringOutput    `pulumi:"iconUrlDark"`
	LogoDarkHash        pulumi.StringPtrOutput `pulumi:"logoDarkHash"`
	LogoDarkPath        pulumi.StringPtrOutput `pulumi:"logoDarkPath"`
	LogoHash            pulumi.StringPtrOutput `pulumi:"logoHash"`
	LogoPath            pulumi.StringPtrOutput `pulumi:"logoPath"`
	LogoUrl             pulumi.StringOutput    `pulumi:"logoUrl"`
	LogoUrlDark         pulumi.StringOutput    `pulumi:"logoUrlDark"`
	// hex value for primary color
	PrimaryColor pulumi.StringOutput `pulumi:"primaryColor"`
	// hex value for primary color dark theme
	PrimaryColorDark pulumi.StringOutput `pulumi:"primaryColorDark"`
	// set the label policy active after creating/updating
	SetActive pulumi.BoolPtrOutput `pulumi:"setActive"`
	// hex value for warn color
	WarnColor pulumi.StringOutput `pulumi:"warnColor"`
	// hex value for warn color dark theme
	WarnColorDark pulumi.StringOutput `pulumi:"warnColorDark"`
}

// NewDefaultLabelPolicy registers a new resource with the given unique name, arguments, and options.
func NewDefaultLabelPolicy(ctx *pulumi.Context,
	name string, args *DefaultLabelPolicyArgs, opts ...pulumi.ResourceOption) (*DefaultLabelPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackgroundColor == nil {
		return nil, errors.New("invalid value for required argument 'BackgroundColor'")
	}
	if args.BackgroundColorDark == nil {
		return nil, errors.New("invalid value for required argument 'BackgroundColorDark'")
	}
	if args.DisableWatermark == nil {
		return nil, errors.New("invalid value for required argument 'DisableWatermark'")
	}
	if args.FontColor == nil {
		return nil, errors.New("invalid value for required argument 'FontColor'")
	}
	if args.FontColorDark == nil {
		return nil, errors.New("invalid value for required argument 'FontColorDark'")
	}
	if args.HideLoginNameSuffix == nil {
		return nil, errors.New("invalid value for required argument 'HideLoginNameSuffix'")
	}
	if args.PrimaryColor == nil {
		return nil, errors.New("invalid value for required argument 'PrimaryColor'")
	}
	if args.PrimaryColorDark == nil {
		return nil, errors.New("invalid value for required argument 'PrimaryColorDark'")
	}
	if args.WarnColor == nil {
		return nil, errors.New("invalid value for required argument 'WarnColor'")
	}
	if args.WarnColorDark == nil {
		return nil, errors.New("invalid value for required argument 'WarnColorDark'")
	}
	var resource DefaultLabelPolicy
	err := ctx.RegisterResource("zitadel:index/defaultLabelPolicy:DefaultLabelPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultLabelPolicy gets an existing DefaultLabelPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultLabelPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultLabelPolicyState, opts ...pulumi.ResourceOption) (*DefaultLabelPolicy, error) {
	var resource DefaultLabelPolicy
	err := ctx.ReadResource("zitadel:index/defaultLabelPolicy:DefaultLabelPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultLabelPolicy resources.
type defaultLabelPolicyState struct {
	// hex value for background color
	BackgroundColor *string `pulumi:"backgroundColor"`
	// hex value for background color dark theme
	BackgroundColorDark *string `pulumi:"backgroundColorDark"`
	// disable watermark
	DisableWatermark *bool `pulumi:"disableWatermark"`
	// hex value for font color
	FontColor *string `pulumi:"fontColor"`
	// hex value for font color dark theme
	FontColorDark *string `pulumi:"fontColorDark"`
	FontHash      *string `pulumi:"fontHash"`
	FontPath      *string `pulumi:"fontPath"`
	FontUrl       *string `pulumi:"fontUrl"`
	// hides the org suffix on the login form if the scope "urn:zitadel:iam:org:domain:primary:{domainname}" is set. Details about this scope in https://docs.zitadel.ch/concepts#Reserved_Scopes
	HideLoginNameSuffix *bool   `pulumi:"hideLoginNameSuffix"`
	IconDarkHash        *string `pulumi:"iconDarkHash"`
	IconDarkPath        *string `pulumi:"iconDarkPath"`
	IconHash            *string `pulumi:"iconHash"`
	IconPath            *string `pulumi:"iconPath"`
	IconUrl             *string `pulumi:"iconUrl"`
	IconUrlDark         *string `pulumi:"iconUrlDark"`
	LogoDarkHash        *string `pulumi:"logoDarkHash"`
	LogoDarkPath        *string `pulumi:"logoDarkPath"`
	LogoHash            *string `pulumi:"logoHash"`
	LogoPath            *string `pulumi:"logoPath"`
	LogoUrl             *string `pulumi:"logoUrl"`
	LogoUrlDark         *string `pulumi:"logoUrlDark"`
	// hex value for primary color
	PrimaryColor *string `pulumi:"primaryColor"`
	// hex value for primary color dark theme
	PrimaryColorDark *string `pulumi:"primaryColorDark"`
	// set the label policy active after creating/updating
	SetActive *bool `pulumi:"setActive"`
	// hex value for warn color
	WarnColor *string `pulumi:"warnColor"`
	// hex value for warn color dark theme
	WarnColorDark *string `pulumi:"warnColorDark"`
}

type DefaultLabelPolicyState struct {
	// hex value for background color
	BackgroundColor pulumi.StringPtrInput
	// hex value for background color dark theme
	BackgroundColorDark pulumi.StringPtrInput
	// disable watermark
	DisableWatermark pulumi.BoolPtrInput
	// hex value for font color
	FontColor pulumi.StringPtrInput
	// hex value for font color dark theme
	FontColorDark pulumi.StringPtrInput
	FontHash      pulumi.StringPtrInput
	FontPath      pulumi.StringPtrInput
	FontUrl       pulumi.StringPtrInput
	// hides the org suffix on the login form if the scope "urn:zitadel:iam:org:domain:primary:{domainname}" is set. Details about this scope in https://docs.zitadel.ch/concepts#Reserved_Scopes
	HideLoginNameSuffix pulumi.BoolPtrInput
	IconDarkHash        pulumi.StringPtrInput
	IconDarkPath        pulumi.StringPtrInput
	IconHash            pulumi.StringPtrInput
	IconPath            pulumi.StringPtrInput
	IconUrl             pulumi.StringPtrInput
	IconUrlDark         pulumi.StringPtrInput
	LogoDarkHash        pulumi.StringPtrInput
	LogoDarkPath        pulumi.StringPtrInput
	LogoHash            pulumi.StringPtrInput
	LogoPath            pulumi.StringPtrInput
	LogoUrl             pulumi.StringPtrInput
	LogoUrlDark         pulumi.StringPtrInput
	// hex value for primary color
	PrimaryColor pulumi.StringPtrInput
	// hex value for primary color dark theme
	PrimaryColorDark pulumi.StringPtrInput
	// set the label policy active after creating/updating
	SetActive pulumi.BoolPtrInput
	// hex value for warn color
	WarnColor pulumi.StringPtrInput
	// hex value for warn color dark theme
	WarnColorDark pulumi.StringPtrInput
}

func (DefaultLabelPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultLabelPolicyState)(nil)).Elem()
}

type defaultLabelPolicyArgs struct {
	// hex value for background color
	BackgroundColor string `pulumi:"backgroundColor"`
	// hex value for background color dark theme
	BackgroundColorDark string `pulumi:"backgroundColorDark"`
	// disable watermark
	DisableWatermark bool `pulumi:"disableWatermark"`
	// hex value for font color
	FontColor string `pulumi:"fontColor"`
	// hex value for font color dark theme
	FontColorDark string  `pulumi:"fontColorDark"`
	FontHash      *string `pulumi:"fontHash"`
	FontPath      *string `pulumi:"fontPath"`
	// hides the org suffix on the login form if the scope "urn:zitadel:iam:org:domain:primary:{domainname}" is set. Details about this scope in https://docs.zitadel.ch/concepts#Reserved_Scopes
	HideLoginNameSuffix bool    `pulumi:"hideLoginNameSuffix"`
	IconDarkHash        *string `pulumi:"iconDarkHash"`
	IconDarkPath        *string `pulumi:"iconDarkPath"`
	IconHash            *string `pulumi:"iconHash"`
	IconPath            *string `pulumi:"iconPath"`
	LogoDarkHash        *string `pulumi:"logoDarkHash"`
	LogoDarkPath        *string `pulumi:"logoDarkPath"`
	LogoHash            *string `pulumi:"logoHash"`
	LogoPath            *string `pulumi:"logoPath"`
	// hex value for primary color
	PrimaryColor string `pulumi:"primaryColor"`
	// hex value for primary color dark theme
	PrimaryColorDark string `pulumi:"primaryColorDark"`
	// set the label policy active after creating/updating
	SetActive *bool `pulumi:"setActive"`
	// hex value for warn color
	WarnColor string `pulumi:"warnColor"`
	// hex value for warn color dark theme
	WarnColorDark string `pulumi:"warnColorDark"`
}

// The set of arguments for constructing a DefaultLabelPolicy resource.
type DefaultLabelPolicyArgs struct {
	// hex value for background color
	BackgroundColor pulumi.StringInput
	// hex value for background color dark theme
	BackgroundColorDark pulumi.StringInput
	// disable watermark
	DisableWatermark pulumi.BoolInput
	// hex value for font color
	FontColor pulumi.StringInput
	// hex value for font color dark theme
	FontColorDark pulumi.StringInput
	FontHash      pulumi.StringPtrInput
	FontPath      pulumi.StringPtrInput
	// hides the org suffix on the login form if the scope "urn:zitadel:iam:org:domain:primary:{domainname}" is set. Details about this scope in https://docs.zitadel.ch/concepts#Reserved_Scopes
	HideLoginNameSuffix pulumi.BoolInput
	IconDarkHash        pulumi.StringPtrInput
	IconDarkPath        pulumi.StringPtrInput
	IconHash            pulumi.StringPtrInput
	IconPath            pulumi.StringPtrInput
	LogoDarkHash        pulumi.StringPtrInput
	LogoDarkPath        pulumi.StringPtrInput
	LogoHash            pulumi.StringPtrInput
	LogoPath            pulumi.StringPtrInput
	// hex value for primary color
	PrimaryColor pulumi.StringInput
	// hex value for primary color dark theme
	PrimaryColorDark pulumi.StringInput
	// set the label policy active after creating/updating
	SetActive pulumi.BoolPtrInput
	// hex value for warn color
	WarnColor pulumi.StringInput
	// hex value for warn color dark theme
	WarnColorDark pulumi.StringInput
}

func (DefaultLabelPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultLabelPolicyArgs)(nil)).Elem()
}

type DefaultLabelPolicyInput interface {
	pulumi.Input

	ToDefaultLabelPolicyOutput() DefaultLabelPolicyOutput
	ToDefaultLabelPolicyOutputWithContext(ctx context.Context) DefaultLabelPolicyOutput
}

func (*DefaultLabelPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultLabelPolicy)(nil)).Elem()
}

func (i *DefaultLabelPolicy) ToDefaultLabelPolicyOutput() DefaultLabelPolicyOutput {
	return i.ToDefaultLabelPolicyOutputWithContext(context.Background())
}

func (i *DefaultLabelPolicy) ToDefaultLabelPolicyOutputWithContext(ctx context.Context) DefaultLabelPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultLabelPolicyOutput)
}

// DefaultLabelPolicyArrayInput is an input type that accepts DefaultLabelPolicyArray and DefaultLabelPolicyArrayOutput values.
// You can construct a concrete instance of `DefaultLabelPolicyArrayInput` via:
//
//	DefaultLabelPolicyArray{ DefaultLabelPolicyArgs{...} }
type DefaultLabelPolicyArrayInput interface {
	pulumi.Input

	ToDefaultLabelPolicyArrayOutput() DefaultLabelPolicyArrayOutput
	ToDefaultLabelPolicyArrayOutputWithContext(context.Context) DefaultLabelPolicyArrayOutput
}

type DefaultLabelPolicyArray []DefaultLabelPolicyInput

func (DefaultLabelPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultLabelPolicy)(nil)).Elem()
}

func (i DefaultLabelPolicyArray) ToDefaultLabelPolicyArrayOutput() DefaultLabelPolicyArrayOutput {
	return i.ToDefaultLabelPolicyArrayOutputWithContext(context.Background())
}

func (i DefaultLabelPolicyArray) ToDefaultLabelPolicyArrayOutputWithContext(ctx context.Context) DefaultLabelPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultLabelPolicyArrayOutput)
}

// DefaultLabelPolicyMapInput is an input type that accepts DefaultLabelPolicyMap and DefaultLabelPolicyMapOutput values.
// You can construct a concrete instance of `DefaultLabelPolicyMapInput` via:
//
//	DefaultLabelPolicyMap{ "key": DefaultLabelPolicyArgs{...} }
type DefaultLabelPolicyMapInput interface {
	pulumi.Input

	ToDefaultLabelPolicyMapOutput() DefaultLabelPolicyMapOutput
	ToDefaultLabelPolicyMapOutputWithContext(context.Context) DefaultLabelPolicyMapOutput
}

type DefaultLabelPolicyMap map[string]DefaultLabelPolicyInput

func (DefaultLabelPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultLabelPolicy)(nil)).Elem()
}

func (i DefaultLabelPolicyMap) ToDefaultLabelPolicyMapOutput() DefaultLabelPolicyMapOutput {
	return i.ToDefaultLabelPolicyMapOutputWithContext(context.Background())
}

func (i DefaultLabelPolicyMap) ToDefaultLabelPolicyMapOutputWithContext(ctx context.Context) DefaultLabelPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultLabelPolicyMapOutput)
}

type DefaultLabelPolicyOutput struct{ *pulumi.OutputState }

func (DefaultLabelPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultLabelPolicy)(nil)).Elem()
}

func (o DefaultLabelPolicyOutput) ToDefaultLabelPolicyOutput() DefaultLabelPolicyOutput {
	return o
}

func (o DefaultLabelPolicyOutput) ToDefaultLabelPolicyOutputWithContext(ctx context.Context) DefaultLabelPolicyOutput {
	return o
}

// hex value for background color
func (o DefaultLabelPolicyOutput) BackgroundColor() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringOutput { return v.BackgroundColor }).(pulumi.StringOutput)
}

// hex value for background color dark theme
func (o DefaultLabelPolicyOutput) BackgroundColorDark() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringOutput { return v.BackgroundColorDark }).(pulumi.StringOutput)
}

// disable watermark
func (o DefaultLabelPolicyOutput) DisableWatermark() pulumi.BoolOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.BoolOutput { return v.DisableWatermark }).(pulumi.BoolOutput)
}

// hex value for font color
func (o DefaultLabelPolicyOutput) FontColor() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringOutput { return v.FontColor }).(pulumi.StringOutput)
}

// hex value for font color dark theme
func (o DefaultLabelPolicyOutput) FontColorDark() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringOutput { return v.FontColorDark }).(pulumi.StringOutput)
}

func (o DefaultLabelPolicyOutput) FontHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringPtrOutput { return v.FontHash }).(pulumi.StringPtrOutput)
}

func (o DefaultLabelPolicyOutput) FontPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringPtrOutput { return v.FontPath }).(pulumi.StringPtrOutput)
}

func (o DefaultLabelPolicyOutput) FontUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringOutput { return v.FontUrl }).(pulumi.StringOutput)
}

// hides the org suffix on the login form if the scope "urn:zitadel:iam:org:domain:primary:{domainname}" is set. Details about this scope in https://docs.zitadel.ch/concepts#Reserved_Scopes
func (o DefaultLabelPolicyOutput) HideLoginNameSuffix() pulumi.BoolOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.BoolOutput { return v.HideLoginNameSuffix }).(pulumi.BoolOutput)
}

func (o DefaultLabelPolicyOutput) IconDarkHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringPtrOutput { return v.IconDarkHash }).(pulumi.StringPtrOutput)
}

func (o DefaultLabelPolicyOutput) IconDarkPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringPtrOutput { return v.IconDarkPath }).(pulumi.StringPtrOutput)
}

func (o DefaultLabelPolicyOutput) IconHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringPtrOutput { return v.IconHash }).(pulumi.StringPtrOutput)
}

func (o DefaultLabelPolicyOutput) IconPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringPtrOutput { return v.IconPath }).(pulumi.StringPtrOutput)
}

func (o DefaultLabelPolicyOutput) IconUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringOutput { return v.IconUrl }).(pulumi.StringOutput)
}

func (o DefaultLabelPolicyOutput) IconUrlDark() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringOutput { return v.IconUrlDark }).(pulumi.StringOutput)
}

func (o DefaultLabelPolicyOutput) LogoDarkHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringPtrOutput { return v.LogoDarkHash }).(pulumi.StringPtrOutput)
}

func (o DefaultLabelPolicyOutput) LogoDarkPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringPtrOutput { return v.LogoDarkPath }).(pulumi.StringPtrOutput)
}

func (o DefaultLabelPolicyOutput) LogoHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringPtrOutput { return v.LogoHash }).(pulumi.StringPtrOutput)
}

func (o DefaultLabelPolicyOutput) LogoPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringPtrOutput { return v.LogoPath }).(pulumi.StringPtrOutput)
}

func (o DefaultLabelPolicyOutput) LogoUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringOutput { return v.LogoUrl }).(pulumi.StringOutput)
}

func (o DefaultLabelPolicyOutput) LogoUrlDark() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringOutput { return v.LogoUrlDark }).(pulumi.StringOutput)
}

// hex value for primary color
func (o DefaultLabelPolicyOutput) PrimaryColor() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringOutput { return v.PrimaryColor }).(pulumi.StringOutput)
}

// hex value for primary color dark theme
func (o DefaultLabelPolicyOutput) PrimaryColorDark() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringOutput { return v.PrimaryColorDark }).(pulumi.StringOutput)
}

// set the label policy active after creating/updating
func (o DefaultLabelPolicyOutput) SetActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.BoolPtrOutput { return v.SetActive }).(pulumi.BoolPtrOutput)
}

// hex value for warn color
func (o DefaultLabelPolicyOutput) WarnColor() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringOutput { return v.WarnColor }).(pulumi.StringOutput)
}

// hex value for warn color dark theme
func (o DefaultLabelPolicyOutput) WarnColorDark() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultLabelPolicy) pulumi.StringOutput { return v.WarnColorDark }).(pulumi.StringOutput)
}

type DefaultLabelPolicyArrayOutput struct{ *pulumi.OutputState }

func (DefaultLabelPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultLabelPolicy)(nil)).Elem()
}

func (o DefaultLabelPolicyArrayOutput) ToDefaultLabelPolicyArrayOutput() DefaultLabelPolicyArrayOutput {
	return o
}

func (o DefaultLabelPolicyArrayOutput) ToDefaultLabelPolicyArrayOutputWithContext(ctx context.Context) DefaultLabelPolicyArrayOutput {
	return o
}

func (o DefaultLabelPolicyArrayOutput) Index(i pulumi.IntInput) DefaultLabelPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DefaultLabelPolicy {
		return vs[0].([]*DefaultLabelPolicy)[vs[1].(int)]
	}).(DefaultLabelPolicyOutput)
}

type DefaultLabelPolicyMapOutput struct{ *pulumi.OutputState }

func (DefaultLabelPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultLabelPolicy)(nil)).Elem()
}

func (o DefaultLabelPolicyMapOutput) ToDefaultLabelPolicyMapOutput() DefaultLabelPolicyMapOutput {
	return o
}

func (o DefaultLabelPolicyMapOutput) ToDefaultLabelPolicyMapOutputWithContext(ctx context.Context) DefaultLabelPolicyMapOutput {
	return o
}

func (o DefaultLabelPolicyMapOutput) MapIndex(k pulumi.StringInput) DefaultLabelPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DefaultLabelPolicy {
		return vs[0].(map[string]*DefaultLabelPolicy)[vs[1].(string)]
	}).(DefaultLabelPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultLabelPolicyInput)(nil)).Elem(), &DefaultLabelPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultLabelPolicyArrayInput)(nil)).Elem(), DefaultLabelPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultLabelPolicyMapInput)(nil)).Elem(), DefaultLabelPolicyMap{})
	pulumi.RegisterOutputType(DefaultLabelPolicyOutput{})
	pulumi.RegisterOutputType(DefaultLabelPolicyArrayOutput{})
	pulumi.RegisterOutputType(DefaultLabelPolicyMapOutput{})
}
