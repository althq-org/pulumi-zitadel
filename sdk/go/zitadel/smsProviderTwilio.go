// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing the SMS provider Twilio configuration of an instance.
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
//			_, err := zitadel.NewSmsProviderTwilio(ctx, "twilio", &zitadel.SmsProviderTwilioArgs{
//				SenderNumber: pulumi.String("019920892"),
//				Sid:          pulumi.String("sid"),
//				Token:        pulumi.String("token"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SmsProviderTwilio struct {
	pulumi.CustomResourceState

	// Sender number which is used to send the SMS.
	SenderNumber pulumi.StringOutput `pulumi:"senderNumber"`
	// SID used to communicate with Twilio.
	Sid pulumi.StringOutput `pulumi:"sid"`
	// Token used to communicate with Twilio.
	Token pulumi.StringOutput `pulumi:"token"`
}

// NewSmsProviderTwilio registers a new resource with the given unique name, arguments, and options.
func NewSmsProviderTwilio(ctx *pulumi.Context,
	name string, args *SmsProviderTwilioArgs, opts ...pulumi.ResourceOption) (*SmsProviderTwilio, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SenderNumber == nil {
		return nil, errors.New("invalid value for required argument 'SenderNumber'")
	}
	if args.Sid == nil {
		return nil, errors.New("invalid value for required argument 'Sid'")
	}
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	if args.Token != nil {
		args.Token = pulumi.ToSecret(args.Token).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource SmsProviderTwilio
	err := ctx.RegisterResource("zitadel:index/smsProviderTwilio:SmsProviderTwilio", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSmsProviderTwilio gets an existing SmsProviderTwilio resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSmsProviderTwilio(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SmsProviderTwilioState, opts ...pulumi.ResourceOption) (*SmsProviderTwilio, error) {
	var resource SmsProviderTwilio
	err := ctx.ReadResource("zitadel:index/smsProviderTwilio:SmsProviderTwilio", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SmsProviderTwilio resources.
type smsProviderTwilioState struct {
	// Sender number which is used to send the SMS.
	SenderNumber *string `pulumi:"senderNumber"`
	// SID used to communicate with Twilio.
	Sid *string `pulumi:"sid"`
	// Token used to communicate with Twilio.
	Token *string `pulumi:"token"`
}

type SmsProviderTwilioState struct {
	// Sender number which is used to send the SMS.
	SenderNumber pulumi.StringPtrInput
	// SID used to communicate with Twilio.
	Sid pulumi.StringPtrInput
	// Token used to communicate with Twilio.
	Token pulumi.StringPtrInput
}

func (SmsProviderTwilioState) ElementType() reflect.Type {
	return reflect.TypeOf((*smsProviderTwilioState)(nil)).Elem()
}

type smsProviderTwilioArgs struct {
	// Sender number which is used to send the SMS.
	SenderNumber string `pulumi:"senderNumber"`
	// SID used to communicate with Twilio.
	Sid string `pulumi:"sid"`
	// Token used to communicate with Twilio.
	Token string `pulumi:"token"`
}

// The set of arguments for constructing a SmsProviderTwilio resource.
type SmsProviderTwilioArgs struct {
	// Sender number which is used to send the SMS.
	SenderNumber pulumi.StringInput
	// SID used to communicate with Twilio.
	Sid pulumi.StringInput
	// Token used to communicate with Twilio.
	Token pulumi.StringInput
}

func (SmsProviderTwilioArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*smsProviderTwilioArgs)(nil)).Elem()
}

type SmsProviderTwilioInput interface {
	pulumi.Input

	ToSmsProviderTwilioOutput() SmsProviderTwilioOutput
	ToSmsProviderTwilioOutputWithContext(ctx context.Context) SmsProviderTwilioOutput
}

func (*SmsProviderTwilio) ElementType() reflect.Type {
	return reflect.TypeOf((**SmsProviderTwilio)(nil)).Elem()
}

func (i *SmsProviderTwilio) ToSmsProviderTwilioOutput() SmsProviderTwilioOutput {
	return i.ToSmsProviderTwilioOutputWithContext(context.Background())
}

func (i *SmsProviderTwilio) ToSmsProviderTwilioOutputWithContext(ctx context.Context) SmsProviderTwilioOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmsProviderTwilioOutput)
}

// SmsProviderTwilioArrayInput is an input type that accepts SmsProviderTwilioArray and SmsProviderTwilioArrayOutput values.
// You can construct a concrete instance of `SmsProviderTwilioArrayInput` via:
//
//	SmsProviderTwilioArray{ SmsProviderTwilioArgs{...} }
type SmsProviderTwilioArrayInput interface {
	pulumi.Input

	ToSmsProviderTwilioArrayOutput() SmsProviderTwilioArrayOutput
	ToSmsProviderTwilioArrayOutputWithContext(context.Context) SmsProviderTwilioArrayOutput
}

type SmsProviderTwilioArray []SmsProviderTwilioInput

func (SmsProviderTwilioArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SmsProviderTwilio)(nil)).Elem()
}

func (i SmsProviderTwilioArray) ToSmsProviderTwilioArrayOutput() SmsProviderTwilioArrayOutput {
	return i.ToSmsProviderTwilioArrayOutputWithContext(context.Background())
}

func (i SmsProviderTwilioArray) ToSmsProviderTwilioArrayOutputWithContext(ctx context.Context) SmsProviderTwilioArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmsProviderTwilioArrayOutput)
}

// SmsProviderTwilioMapInput is an input type that accepts SmsProviderTwilioMap and SmsProviderTwilioMapOutput values.
// You can construct a concrete instance of `SmsProviderTwilioMapInput` via:
//
//	SmsProviderTwilioMap{ "key": SmsProviderTwilioArgs{...} }
type SmsProviderTwilioMapInput interface {
	pulumi.Input

	ToSmsProviderTwilioMapOutput() SmsProviderTwilioMapOutput
	ToSmsProviderTwilioMapOutputWithContext(context.Context) SmsProviderTwilioMapOutput
}

type SmsProviderTwilioMap map[string]SmsProviderTwilioInput

func (SmsProviderTwilioMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SmsProviderTwilio)(nil)).Elem()
}

func (i SmsProviderTwilioMap) ToSmsProviderTwilioMapOutput() SmsProviderTwilioMapOutput {
	return i.ToSmsProviderTwilioMapOutputWithContext(context.Background())
}

func (i SmsProviderTwilioMap) ToSmsProviderTwilioMapOutputWithContext(ctx context.Context) SmsProviderTwilioMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmsProviderTwilioMapOutput)
}

type SmsProviderTwilioOutput struct{ *pulumi.OutputState }

func (SmsProviderTwilioOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SmsProviderTwilio)(nil)).Elem()
}

func (o SmsProviderTwilioOutput) ToSmsProviderTwilioOutput() SmsProviderTwilioOutput {
	return o
}

func (o SmsProviderTwilioOutput) ToSmsProviderTwilioOutputWithContext(ctx context.Context) SmsProviderTwilioOutput {
	return o
}

// Sender number which is used to send the SMS.
func (o SmsProviderTwilioOutput) SenderNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *SmsProviderTwilio) pulumi.StringOutput { return v.SenderNumber }).(pulumi.StringOutput)
}

// SID used to communicate with Twilio.
func (o SmsProviderTwilioOutput) Sid() pulumi.StringOutput {
	return o.ApplyT(func(v *SmsProviderTwilio) pulumi.StringOutput { return v.Sid }).(pulumi.StringOutput)
}

// Token used to communicate with Twilio.
func (o SmsProviderTwilioOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *SmsProviderTwilio) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

type SmsProviderTwilioArrayOutput struct{ *pulumi.OutputState }

func (SmsProviderTwilioArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SmsProviderTwilio)(nil)).Elem()
}

func (o SmsProviderTwilioArrayOutput) ToSmsProviderTwilioArrayOutput() SmsProviderTwilioArrayOutput {
	return o
}

func (o SmsProviderTwilioArrayOutput) ToSmsProviderTwilioArrayOutputWithContext(ctx context.Context) SmsProviderTwilioArrayOutput {
	return o
}

func (o SmsProviderTwilioArrayOutput) Index(i pulumi.IntInput) SmsProviderTwilioOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SmsProviderTwilio {
		return vs[0].([]*SmsProviderTwilio)[vs[1].(int)]
	}).(SmsProviderTwilioOutput)
}

type SmsProviderTwilioMapOutput struct{ *pulumi.OutputState }

func (SmsProviderTwilioMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SmsProviderTwilio)(nil)).Elem()
}

func (o SmsProviderTwilioMapOutput) ToSmsProviderTwilioMapOutput() SmsProviderTwilioMapOutput {
	return o
}

func (o SmsProviderTwilioMapOutput) ToSmsProviderTwilioMapOutputWithContext(ctx context.Context) SmsProviderTwilioMapOutput {
	return o
}

func (o SmsProviderTwilioMapOutput) MapIndex(k pulumi.StringInput) SmsProviderTwilioOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SmsProviderTwilio {
		return vs[0].(map[string]*SmsProviderTwilio)[vs[1].(string)]
	}).(SmsProviderTwilioOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SmsProviderTwilioInput)(nil)).Elem(), &SmsProviderTwilio{})
	pulumi.RegisterInputType(reflect.TypeOf((*SmsProviderTwilioArrayInput)(nil)).Elem(), SmsProviderTwilioArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SmsProviderTwilioMapInput)(nil)).Elem(), SmsProviderTwilioMap{})
	pulumi.RegisterOutputType(SmsProviderTwilioOutput{})
	pulumi.RegisterOutputType(SmsProviderTwilioArrayOutput{})
	pulumi.RegisterOutputType(SmsProviderTwilioMapOutput{})
}
