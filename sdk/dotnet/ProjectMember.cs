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
    /// Resource representing the membership of a user on an project, defined with the given role.
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
    ///     var projectMember = new Zitadel.ProjectMember("projectMember", new()
    ///     {
    ///         OrgId = zitadel_org.Org.Id,
    ///         ProjectId = zitadel_project.Project.Id,
    ///         UserId = zitadel_human_user.Human_user.Id,
    ///         Roles = new[]
    ///         {
    ///             "PROJECT_OWNER",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/projectMember:ProjectMember")]
    public partial class ProjectMember : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the organization which owns the resource
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// ID of the project
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// List of roles granted
        /// </summary>
        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;

        /// <summary>
        /// ID of the user
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectMember(string name, ProjectMemberArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/projectMember:ProjectMember", name, args ?? new ProjectMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectMember(string name, Input<string> id, ProjectMemberState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/projectMember:ProjectMember", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ProjectMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectMember Get(string name, Input<string> id, ProjectMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectMember(name, id, state, options);
        }
    }

    public sealed class ProjectMemberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the organization which owns the resource
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        /// <summary>
        /// ID of the project
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("roles", required: true)]
        private InputList<string>? _roles;

        /// <summary>
        /// List of roles granted
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// ID of the user
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public ProjectMemberArgs()
        {
        }
        public static new ProjectMemberArgs Empty => new ProjectMemberArgs();
    }

    public sealed class ProjectMemberState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the organization which owns the resource
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// ID of the project
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// List of roles granted
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// ID of the user
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public ProjectMemberState()
        {
        }
        public static new ProjectMemberState Empty => new ProjectMemberState();
    }
}
