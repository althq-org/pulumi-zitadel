// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing the SMTP configuration of an instance.
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
//			_, err := zitadel.NewSmtpConfig(ctx, "smtp", &zitadel.SmtpConfigArgs{
//				Host:          pulumi.String("localhost:25"),
//				Password:      pulumi.String("password"),
//				SenderAddress: pulumi.String("address"),
//				SenderName:    pulumi.String("no-reply"),
//				Tls:           pulumi.Bool(true),
//				User:          pulumi.String("user"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SmtpConfig struct {
	pulumi.CustomResourceState

	// Host and port address to your SMTP server.
	Host pulumi.StringOutput `pulumi:"host"`
	// Password used to communicate with your SMTP server.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Address used to send emails.
	SenderAddress pulumi.StringOutput `pulumi:"senderAddress"`
	// Sender name used to send emails.
	SenderName pulumi.StringOutput `pulumi:"senderName"`
	// TLS used to communicate with your SMTP server.
	Tls pulumi.BoolPtrOutput `pulumi:"tls"`
	// User used to communicate with your SMTP server.
	User pulumi.StringPtrOutput `pulumi:"user"`
}

// NewSmtpConfig registers a new resource with the given unique name, arguments, and options.
func NewSmtpConfig(ctx *pulumi.Context,
	name string, args *SmtpConfigArgs, opts ...pulumi.ResourceOption) (*SmtpConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Host == nil {
		return nil, errors.New("invalid value for required argument 'Host'")
	}
	if args.SenderAddress == nil {
		return nil, errors.New("invalid value for required argument 'SenderAddress'")
	}
	if args.SenderName == nil {
		return nil, errors.New("invalid value for required argument 'SenderName'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource SmtpConfig
	err := ctx.RegisterResource("zitadel:index/smtpConfig:SmtpConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSmtpConfig gets an existing SmtpConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSmtpConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SmtpConfigState, opts ...pulumi.ResourceOption) (*SmtpConfig, error) {
	var resource SmtpConfig
	err := ctx.ReadResource("zitadel:index/smtpConfig:SmtpConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SmtpConfig resources.
type smtpConfigState struct {
	// Host and port address to your SMTP server.
	Host *string `pulumi:"host"`
	// Password used to communicate with your SMTP server.
	Password *string `pulumi:"password"`
	// Address used to send emails.
	SenderAddress *string `pulumi:"senderAddress"`
	// Sender name used to send emails.
	SenderName *string `pulumi:"senderName"`
	// TLS used to communicate with your SMTP server.
	Tls *bool `pulumi:"tls"`
	// User used to communicate with your SMTP server.
	User *string `pulumi:"user"`
}

type SmtpConfigState struct {
	// Host and port address to your SMTP server.
	Host pulumi.StringPtrInput
	// Password used to communicate with your SMTP server.
	Password pulumi.StringPtrInput
	// Address used to send emails.
	SenderAddress pulumi.StringPtrInput
	// Sender name used to send emails.
	SenderName pulumi.StringPtrInput
	// TLS used to communicate with your SMTP server.
	Tls pulumi.BoolPtrInput
	// User used to communicate with your SMTP server.
	User pulumi.StringPtrInput
}

func (SmtpConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*smtpConfigState)(nil)).Elem()
}

type smtpConfigArgs struct {
	// Host and port address to your SMTP server.
	Host string `pulumi:"host"`
	// Password used to communicate with your SMTP server.
	Password *string `pulumi:"password"`
	// Address used to send emails.
	SenderAddress string `pulumi:"senderAddress"`
	// Sender name used to send emails.
	SenderName string `pulumi:"senderName"`
	// TLS used to communicate with your SMTP server.
	Tls *bool `pulumi:"tls"`
	// User used to communicate with your SMTP server.
	User *string `pulumi:"user"`
}

// The set of arguments for constructing a SmtpConfig resource.
type SmtpConfigArgs struct {
	// Host and port address to your SMTP server.
	Host pulumi.StringInput
	// Password used to communicate with your SMTP server.
	Password pulumi.StringPtrInput
	// Address used to send emails.
	SenderAddress pulumi.StringInput
	// Sender name used to send emails.
	SenderName pulumi.StringInput
	// TLS used to communicate with your SMTP server.
	Tls pulumi.BoolPtrInput
	// User used to communicate with your SMTP server.
	User pulumi.StringPtrInput
}

func (SmtpConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*smtpConfigArgs)(nil)).Elem()
}

type SmtpConfigInput interface {
	pulumi.Input

	ToSmtpConfigOutput() SmtpConfigOutput
	ToSmtpConfigOutputWithContext(ctx context.Context) SmtpConfigOutput
}

func (*SmtpConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**SmtpConfig)(nil)).Elem()
}

func (i *SmtpConfig) ToSmtpConfigOutput() SmtpConfigOutput {
	return i.ToSmtpConfigOutputWithContext(context.Background())
}

func (i *SmtpConfig) ToSmtpConfigOutputWithContext(ctx context.Context) SmtpConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmtpConfigOutput)
}

// SmtpConfigArrayInput is an input type that accepts SmtpConfigArray and SmtpConfigArrayOutput values.
// You can construct a concrete instance of `SmtpConfigArrayInput` via:
//
//	SmtpConfigArray{ SmtpConfigArgs{...} }
type SmtpConfigArrayInput interface {
	pulumi.Input

	ToSmtpConfigArrayOutput() SmtpConfigArrayOutput
	ToSmtpConfigArrayOutputWithContext(context.Context) SmtpConfigArrayOutput
}

type SmtpConfigArray []SmtpConfigInput

func (SmtpConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SmtpConfig)(nil)).Elem()
}

func (i SmtpConfigArray) ToSmtpConfigArrayOutput() SmtpConfigArrayOutput {
	return i.ToSmtpConfigArrayOutputWithContext(context.Background())
}

func (i SmtpConfigArray) ToSmtpConfigArrayOutputWithContext(ctx context.Context) SmtpConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmtpConfigArrayOutput)
}

// SmtpConfigMapInput is an input type that accepts SmtpConfigMap and SmtpConfigMapOutput values.
// You can construct a concrete instance of `SmtpConfigMapInput` via:
//
//	SmtpConfigMap{ "key": SmtpConfigArgs{...} }
type SmtpConfigMapInput interface {
	pulumi.Input

	ToSmtpConfigMapOutput() SmtpConfigMapOutput
	ToSmtpConfigMapOutputWithContext(context.Context) SmtpConfigMapOutput
}

type SmtpConfigMap map[string]SmtpConfigInput

func (SmtpConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SmtpConfig)(nil)).Elem()
}

func (i SmtpConfigMap) ToSmtpConfigMapOutput() SmtpConfigMapOutput {
	return i.ToSmtpConfigMapOutputWithContext(context.Background())
}

func (i SmtpConfigMap) ToSmtpConfigMapOutputWithContext(ctx context.Context) SmtpConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmtpConfigMapOutput)
}

type SmtpConfigOutput struct{ *pulumi.OutputState }

func (SmtpConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SmtpConfig)(nil)).Elem()
}

func (o SmtpConfigOutput) ToSmtpConfigOutput() SmtpConfigOutput {
	return o
}

func (o SmtpConfigOutput) ToSmtpConfigOutputWithContext(ctx context.Context) SmtpConfigOutput {
	return o
}

// Host and port address to your SMTP server.
func (o SmtpConfigOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *SmtpConfig) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// Password used to communicate with your SMTP server.
func (o SmtpConfigOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SmtpConfig) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Address used to send emails.
func (o SmtpConfigOutput) SenderAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *SmtpConfig) pulumi.StringOutput { return v.SenderAddress }).(pulumi.StringOutput)
}

// Sender name used to send emails.
func (o SmtpConfigOutput) SenderName() pulumi.StringOutput {
	return o.ApplyT(func(v *SmtpConfig) pulumi.StringOutput { return v.SenderName }).(pulumi.StringOutput)
}

// TLS used to communicate with your SMTP server.
func (o SmtpConfigOutput) Tls() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SmtpConfig) pulumi.BoolPtrOutput { return v.Tls }).(pulumi.BoolPtrOutput)
}

// User used to communicate with your SMTP server.
func (o SmtpConfigOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SmtpConfig) pulumi.StringPtrOutput { return v.User }).(pulumi.StringPtrOutput)
}

type SmtpConfigArrayOutput struct{ *pulumi.OutputState }

func (SmtpConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SmtpConfig)(nil)).Elem()
}

func (o SmtpConfigArrayOutput) ToSmtpConfigArrayOutput() SmtpConfigArrayOutput {
	return o
}

func (o SmtpConfigArrayOutput) ToSmtpConfigArrayOutputWithContext(ctx context.Context) SmtpConfigArrayOutput {
	return o
}

func (o SmtpConfigArrayOutput) Index(i pulumi.IntInput) SmtpConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SmtpConfig {
		return vs[0].([]*SmtpConfig)[vs[1].(int)]
	}).(SmtpConfigOutput)
}

type SmtpConfigMapOutput struct{ *pulumi.OutputState }

func (SmtpConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SmtpConfig)(nil)).Elem()
}

func (o SmtpConfigMapOutput) ToSmtpConfigMapOutput() SmtpConfigMapOutput {
	return o
}

func (o SmtpConfigMapOutput) ToSmtpConfigMapOutputWithContext(ctx context.Context) SmtpConfigMapOutput {
	return o
}

func (o SmtpConfigMapOutput) MapIndex(k pulumi.StringInput) SmtpConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SmtpConfig {
		return vs[0].(map[string]*SmtpConfig)[vs[1].(string)]
	}).(SmtpConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SmtpConfigInput)(nil)).Elem(), &SmtpConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*SmtpConfigArrayInput)(nil)).Elem(), SmtpConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SmtpConfigMapInput)(nil)).Elem(), SmtpConfigMap{})
	pulumi.RegisterOutputType(SmtpConfigOutput{})
	pulumi.RegisterOutputType(SmtpConfigArrayOutput{})
	pulumi.RegisterOutputType(SmtpConfigMapOutput{})
}
