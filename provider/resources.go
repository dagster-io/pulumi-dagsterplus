package dagsterplus

import (
	_ "embed"

	pf "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"

	dagsterprovider "github.com/dagster-io/terraform-provider-dagsterplus/pkg/provider"
)

//go:embed cmd/pulumi-resource-dagsterplus/bridge-metadata.json
var bridgeMetadata []byte

var Version = "0.0.1"

func Provider() tfbridge.ProviderInfo {
	info := tfbridge.ProviderInfo{
		P:                 pf.ShimProvider(dagsterprovider.New("0.0.1")()),
		Name:              "dagsterplus",
		Version:           Version,
		DisplayName:       "Dagster+",
		Publisher:         "Dagster",
		Description:       "A Pulumi provider for managing Dagster+ resources.",
		License:           "Apache-2.0",
		Homepage:          "https://www.dagster.io",
		Repository:        "https://github.com/dagster-io/pulumi-dagsterplus",
		GitHubOrg:         "dagster-io",
		PluginDownloadURL: "github://api.github.com/dagster-io",
		UpstreamRepoPath:  "/Users/dennis/code/terraform-provider-dagster-plus",
		MetadataInfo:      tfbridge.NewProviderMetadata(bridgeMetadata),
	}
	info.MustComputeTokens(tokens.SingleModule(
		"dagsterplus_", "index", tokens.MakeStandard("dagsterplus"),
	))
	return info
}
