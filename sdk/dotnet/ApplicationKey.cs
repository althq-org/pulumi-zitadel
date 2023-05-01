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
    /// Resource representing a app key
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
    ///     var appKey = new Zitadel.ApplicationKey("appKey", new()
    ///     {
    ///         OrgId = zitadel_org.Org.Id,
    ///         ProjectId = zitadel_project.Project.Id,
    ///         AppId = zitadel_application_api.Application_api.Id,
    ///         KeyType = "KEY_TYPE_JSON",
    ///         ExpirationDate = "2519-04-01T08:45:00Z",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/applicationKey:ApplicationKey")]
    public partial class ApplicationKey : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the application
        /// </summary>
        [Output("appId")]
        public Output<string> AppId { get; private set; } = null!;

        /// <summary>
        /// Expiration date of the app key in the RFC3339 format
        /// </summary>
        [Output("expirationDate")]
        public Output<string> ExpirationDate { get; private set; } = null!;

        /// <summary>
        /// Value of the app key
        /// </summary>
        [Output("keyDetails")]
        public Output<string> KeyDetails { get; private set; } = null!;

        /// <summary>
        /// Type of the app key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
        /// </summary>
        [Output("keyType")]
        public Output<string> KeyType { get; private set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// ID of the project
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationKey(string name, ApplicationKeyArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/applicationKey:ApplicationKey", name, args ?? new ApplicationKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationKey(string name, Input<string> id, ApplicationKeyState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/applicationKey:ApplicationKey", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "keyDetails",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ApplicationKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationKey Get(string name, Input<string> id, ApplicationKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationKey(name, id, state, options);
        }
    }

    public sealed class ApplicationKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the application
        /// </summary>
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        /// <summary>
        /// Expiration date of the app key in the RFC3339 format
        /// </summary>
        [Input("expirationDate", required: true)]
        public Input<string> ExpirationDate { get; set; } = null!;

        /// <summary>
        /// Type of the app key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
        /// </summary>
        [Input("keyType", required: true)]
        public Input<string> KeyType { get; set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        /// <summary>
        /// ID of the project
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public ApplicationKeyArgs()
        {
        }
        public static new ApplicationKeyArgs Empty => new ApplicationKeyArgs();
    }

    public sealed class ApplicationKeyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the application
        /// </summary>
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        /// <summary>
        /// Expiration date of the app key in the RFC3339 format
        /// </summary>
        [Input("expirationDate")]
        public Input<string>? ExpirationDate { get; set; }

        [Input("keyDetails")]
        private Input<string>? _keyDetails;

        /// <summary>
        /// Value of the app key
        /// </summary>
        public Input<string>? KeyDetails
        {
            get => _keyDetails;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _keyDetails = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Type of the app key, supported values: KEY*TYPE*UNSPECIFIED, KEY*TYPE*JSON
        /// </summary>
        [Input("keyType")]
        public Input<string>? KeyType { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// ID of the project
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public ApplicationKeyState()
        {
        }
        public static new ApplicationKeyState Empty => new ApplicationKeyState();
    }
}
