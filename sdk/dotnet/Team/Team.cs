// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Buildkite.Team
{
    /// <summary>
    /// ## # Resource: team
    /// 
    /// This resource allows you to create and manage teams.
    /// 
    /// Buildkite Documentation: https://buildkite.com/docs/pipelines/permissions
    /// 
    /// Note: You must first enable Teams on your organization.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Buildkite = Pulumiverse.Buildkite;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var team = new Buildkite.Team.Team("team", new()
    ///     {
    ///         DefaultMemberRole = "MEMBER",
    ///         DefaultTeam = true,
    ///         Privacy = "VISIBLE",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Teams can be imported using the `GraphQL ID` (not UUID), e.g.
    /// 
    /// ```sh
    ///  $ pulumi import buildkite:Team/team:Team fleet UGlwZWxpbmUtLS00MzVjYWQ1OC1lODFkLTQ1YWYtODYzNy1iMWNmODA3MDIzOGQ=
    /// ```
    /// 
    ///  To find the ID to use, you can use the GraphQL query below. Alternatively, you could use this [pre-saved query](https://buildkite.com/user/graphql/console/6e74c89c-4e91-4d1d-92ca-4fb19d0ea453), where you will need fo fill in the organization slug and search term (TEAM_SEARCH_TERM) for the team. graphql query getTeamId {
    /// 
    ///  organization(slug"ORGANIZATION_SLUG") {
    /// 
    ///  teams(first1, search"TEAM_SEARCH_TERM") {
    /// 
    ///  edges {
    /// 
    ///  node {
    /// 
    ///  id
    /// 
    ///  name
    /// 
    ///  }
    /// 
    ///  }
    /// 
    ///  }
    /// 
    ///  } }
    /// </summary>
    [BuildkiteResourceType("buildkite:Team/team:Team")]
    public partial class Team : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Default role to assign to a team member.
        /// </summary>
        [Output("defaultMemberRole")]
        public Output<string> DefaultMemberRole { get; private set; } = null!;

        /// <summary>
        /// Whether to assign this team to a user by default.
        /// </summary>
        [Output("defaultTeam")]
        public Output<bool> DefaultTeam { get; private set; } = null!;

        /// <summary>
        /// The description to assign to the team.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether team members can create.
        /// </summary>
        [Output("membersCanCreatePipelines")]
        public Output<bool> MembersCanCreatePipelines { get; private set; } = null!;

        /// <summary>
        /// The name of the team.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The privacy level to set the team too.
        /// </summary>
        [Output("privacy")]
        public Output<string> Privacy { get; private set; } = null!;

        /// <summary>
        /// The name of the team.
        /// </summary>
        [Output("slug")]
        public Output<string> Slug { get; private set; } = null!;

        /// <summary>
        /// The UUID for the team.
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;


        /// <summary>
        /// Create a Team resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Team(string name, TeamArgs args, CustomResourceOptions? options = null)
            : base("buildkite:Team/team:Team", name, args ?? new TeamArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Team(string name, Input<string> id, TeamState? state = null, CustomResourceOptions? options = null)
            : base("buildkite:Team/team:Team", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-buildkite",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Team resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Team Get(string name, Input<string> id, TeamState? state = null, CustomResourceOptions? options = null)
        {
            return new Team(name, id, state, options);
        }
    }

    public sealed class TeamArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Default role to assign to a team member.
        /// </summary>
        [Input("defaultMemberRole", required: true)]
        public Input<string> DefaultMemberRole { get; set; } = null!;

        /// <summary>
        /// Whether to assign this team to a user by default.
        /// </summary>
        [Input("defaultTeam", required: true)]
        public Input<bool> DefaultTeam { get; set; } = null!;

        /// <summary>
        /// The description to assign to the team.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether team members can create.
        /// </summary>
        [Input("membersCanCreatePipelines")]
        public Input<bool>? MembersCanCreatePipelines { get; set; }

        /// <summary>
        /// The name of the team.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The privacy level to set the team too.
        /// </summary>
        [Input("privacy", required: true)]
        public Input<string> Privacy { get; set; } = null!;

        public TeamArgs()
        {
        }
        public static new TeamArgs Empty => new TeamArgs();
    }

    public sealed class TeamState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Default role to assign to a team member.
        /// </summary>
        [Input("defaultMemberRole")]
        public Input<string>? DefaultMemberRole { get; set; }

        /// <summary>
        /// Whether to assign this team to a user by default.
        /// </summary>
        [Input("defaultTeam")]
        public Input<bool>? DefaultTeam { get; set; }

        /// <summary>
        /// The description to assign to the team.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether team members can create.
        /// </summary>
        [Input("membersCanCreatePipelines")]
        public Input<bool>? MembersCanCreatePipelines { get; set; }

        /// <summary>
        /// The name of the team.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The privacy level to set the team too.
        /// </summary>
        [Input("privacy")]
        public Input<string>? Privacy { get; set; }

        /// <summary>
        /// The name of the team.
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        /// <summary>
        /// The UUID for the team.
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        public TeamState()
        {
        }
        public static new TeamState Empty => new TeamState();
    }
}
