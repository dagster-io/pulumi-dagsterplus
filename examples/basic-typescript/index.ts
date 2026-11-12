import * as pulumi from "@pulumi/pulumi";
import * as dagsterplus from "@pulumi/dagsterplus";

const team = new dagsterplus.Team("pulumi-test-team", {
    name: "pulumi-test-team",
});

export const teamName = team.name;
