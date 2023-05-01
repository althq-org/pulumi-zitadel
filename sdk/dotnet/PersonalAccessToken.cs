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
    /// Resource representing a personal access token of a user
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
    ///     var pat = new Zitadel.PersonalAccessToken("pat", new()
    ///     {
    ///         OrgId = zitadel_org.Org.Id,
    ///         UserId = zitadel_machine_user.Machine_user.Id,
    ///         ExpirationDate = "2519-04-01T08:45:00Z",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/personalAccessToken:PersonalAccessToken")]
    public partial class PersonalAccessToken : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Expiration date of the token in the RFC3339 format
        /// </summary>
        [Output("expirationDate")]
        public Output<string?> ExpirationDate { get; private set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// Value of the token
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        /// <summary>
        /// ID of the user
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a PersonalAccessToken resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PersonalAccessToken(string name, PersonalAccessTokenArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/personalAccessToken:PersonalAccessToken", name, args ?? new PersonalAccessTokenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PersonalAccessToken(string name, Input<string> id, PersonalAccessTokenState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/personalAccessToken:PersonalAccessToken", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PersonalAccessToken resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PersonalAccessToken Get(string name, Input<string> id, PersonalAccessTokenState? state = null, CustomResourceOptions? options = null)
        {
            return new PersonalAccessToken(name, id, state, options);
        }
    }

    public sealed class PersonalAccessTokenArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Expiration date of the token in the RFC3339 format
        /// </summary>
        [Input("expirationDate")]
        public Input<string>? ExpirationDate { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        /// <summary>
        /// ID of the user
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public PersonalAccessTokenArgs()
        {
        }
        public static new PersonalAccessTokenArgs Empty => new PersonalAccessTokenArgs();
    }

    public sealed class PersonalAccessTokenState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Expiration date of the token in the RFC3339 format
        /// </summary>
        [Input("expirationDate")]
        public Input<string>? ExpirationDate { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("token")]
        private Input<string>? _token;

        /// <summary>
        /// Value of the token
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

        /// <summary>
        /// ID of the user
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public PersonalAccessTokenState()
        {
        }
        public static new PersonalAccessTokenState Empty => new PersonalAccessTokenState();
    }
}
