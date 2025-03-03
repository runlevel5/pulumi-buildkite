// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pipeline

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-buildkite/sdk/v2/go/buildkite/internal"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "buildkite:Pipeline/pipeline:Pipeline":
		r = &Pipeline{}
	case "buildkite:Pipeline/schedule:Schedule":
		r = &Schedule{}
	case "buildkite:Pipeline/team:Team":
		r = &Team{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"buildkite",
		"Pipeline/pipeline",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"buildkite",
		"Pipeline/schedule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"buildkite",
		"Pipeline/team",
		&module{version},
	)
}
