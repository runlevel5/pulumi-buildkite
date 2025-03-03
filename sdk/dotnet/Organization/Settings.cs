// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Buildkite.Organization
{
    /// <summary>
    /// ## # Resource: organization_settings
    /// 
    /// **Note**: This resource has been deprecated. In a future minor release, we will remove this resource in favour of the newer `buildkite.Organization.Organization` resource that aligns with the datasource of the same name.
    /// 
    /// This resource allows you to manage the settings for an organization.
    /// 
    /// You must be an organization administrator to manage organization settings.
    /// 
    /// Note: The "Allowed API IP Addresses" feature must be enabled on your organization in order to manage the `allowed_api_ip_addresses` attribute.
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
    ///     var testSettings = new Buildkite.Organization.Settings("testSettings", new()
    ///     {
    ///         AllowedApiIpAddresses = new[]
    ///         {
    ///             "1.1.1.1/32",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Organization settings can be imported by passing the organization slug to the import command, along with the identifier of the resource.
    /// 
    /// ```sh
    ///  $ pulumi import buildkite:Organization/settings:Settings test_settings test_org
    /// ```
    /// 
    ///  Your organization's slug can be found in your organisation's [settings](https://buildkite.com/organizations/~/settings) page.
    /// </summary>
    [BuildkiteResourceType("buildkite:Organization/settings:Settings")]
    public partial class Settings : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of IP addresses in CIDR format that are allowed to access the Buildkite API. If not set, all IP addresses are allowed (the same as setting 0.0.0.0/0).
        /// </summary>
        [Output("allowedApiIpAddresses")]
        public Output<ImmutableArray<string>> AllowedApiIpAddresses { get; private set; } = null!;

        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;


        /// <summary>
        /// Create a Settings resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Settings(string name, SettingsArgs? args = null, CustomResourceOptions? options = null)
            : base("buildkite:Organization/settings:Settings", name, args ?? new SettingsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Settings(string name, Input<string> id, SettingsState? state = null, CustomResourceOptions? options = null)
            : base("buildkite:Organization/settings:Settings", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Settings resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Settings Get(string name, Input<string> id, SettingsState? state = null, CustomResourceOptions? options = null)
        {
            return new Settings(name, id, state, options);
        }
    }

    public sealed class SettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedApiIpAddresses")]
        private InputList<string>? _allowedApiIpAddresses;

        /// <summary>
        /// A list of IP addresses in CIDR format that are allowed to access the Buildkite API. If not set, all IP addresses are allowed (the same as setting 0.0.0.0/0).
        /// </summary>
        public InputList<string> AllowedApiIpAddresses
        {
            get => _allowedApiIpAddresses ?? (_allowedApiIpAddresses = new InputList<string>());
            set => _allowedApiIpAddresses = value;
        }

        public SettingsArgs()
        {
        }
        public static new SettingsArgs Empty => new SettingsArgs();
    }

    public sealed class SettingsState : global::Pulumi.ResourceArgs
    {
        [Input("allowedApiIpAddresses")]
        private InputList<string>? _allowedApiIpAddresses;

        /// <summary>
        /// A list of IP addresses in CIDR format that are allowed to access the Buildkite API. If not set, all IP addresses are allowed (the same as setting 0.0.0.0/0).
        /// </summary>
        public InputList<string> AllowedApiIpAddresses
        {
            get => _allowedApiIpAddresses ?? (_allowedApiIpAddresses = new InputList<string>());
            set => _allowedApiIpAddresses = value;
        }

        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        public SettingsState()
        {
        }
        public static new SettingsState Empty => new SettingsState();
    }
}
