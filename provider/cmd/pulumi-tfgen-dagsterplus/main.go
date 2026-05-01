package main

import (
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfgen"

	dagsterplus "github.com/dagster-io/pulumi-dagsterplus/provider"
)

func main() {
	tfgen.Main("dagsterplus", dagsterplus.Provider())
}
