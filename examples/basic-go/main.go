package main

import (
	"github.com/dagster-io/pulumi-dagsterplus/sdk/go/dagsterplus"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		team, err := dagsterplus.NewTeam(ctx, "myTeam", &dagsterplus.TeamArgs{
			Name: pulumi.String("pulumi-test-team"),
		})
		if err != nil {
			return err
		}
		ctx.Export("teamName", team.Name)
		return nil
	})
}
