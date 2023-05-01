// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Zitadel
{
    /// <summary>
    /// Resource representing the SMS provider Twilio configuration of an instance.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Zitadel = Pulumi.Zitadel;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var twilio = new Zitadel.SmsProviderTwilio("twilio", new()
    ///     {
    ///         SenderNumber = "019920892",
    ///         Sid = "sid",
    ///         Token = "token",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/smsProviderTwilio:SmsProviderTwilio")]
    public partial class SmsProviderTwilio : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sender number which is used to send the SMS.
        /// </summary>
        [Output("senderNumber")]
        public Output<string> SenderNumber { get; private set; } = null!;

        /// <summary>
        /// SID used to communicate with Twilio.
        /// </summary>
        [Output("sid")]
        public Output<string> Sid { get; private set; } = null!;

        /// <summary>
        /// Token used to communicate with Twilio.
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;


        /// <summary>
        /// Create a SmsProviderTwilio resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SmsProviderTwilio(string name, SmsProviderTwilioArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/smsProviderTwilio:SmsProviderTwilio", name, args ?? new SmsProviderTwilioArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SmsProviderTwilio(string name, Input<string> id, SmsProviderTwilioState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/smsProviderTwilio:SmsProviderTwilio", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "token",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SmsProviderTwilio resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SmsProviderTwilio Get(string name, Input<string> id, SmsProviderTwilioState? state = null, CustomResourceOptions? options = null)
        {
            return new SmsProviderTwilio(name, id, state, options);
        }
    }

    public sealed class SmsProviderTwilioArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sender number which is used to send the SMS.
        /// </summary>
        [Input("senderNumber", required: true)]
        public Input<string> SenderNumber { get; set; } = null!;

        /// <summary>
        /// SID used to communicate with Twilio.
        /// </summary>
        [Input("sid", required: true)]
        public Input<string> Sid { get; set; } = null!;

        [Input("token", required: true)]
        private Input<string>? _token;

        /// <summary>
        /// Token used to communicate with Twilio.
        /// </summary>
        public Input<string>? Token
        {
            get => _token;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _token = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public SmsProviderTwilioArgs()
        {
        }
        public static new SmsProviderTwilioArgs Empty => new SmsProviderTwilioArgs();
    }

    public sealed class SmsProviderTwilioState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sender number which is used to send the SMS.
        /// </summary>
        [Input("senderNumber")]
        public Input<string>? SenderNumber { get; set; }

        /// <summary>
        /// SID used to communicate with Twilio.
        /// </summary>
        [Input("sid")]
        public Input<string>? Sid { get; set; }

        [Input("token")]
        private Input<string>? _token;

        /// <summary>
        /// Token used to communicate with Twilio.
        /// </summary>
        public Input<string>? Token
        {
            get => _token;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _token = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public SmsProviderTwilioState()
        {
        }
        public static new SmsProviderTwilioState Empty => new SmsProviderTwilioState();
    }
}
