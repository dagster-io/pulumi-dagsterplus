---
title: Dagster+
meta_desc: Provides an overview of the Dagster+ provider for Pulumi.
layout: package
---

## Overview

The Dagster+ provider for Pulumi can be used to provision and manage resources in [Dagster+](https://dagster.io/cloud), the managed cloud platform for Dagster. Use it to manage deployments, code locations, agents, teams, secrets, and more — all as code.

## Installation

The Dagster+ provider is available as a package in the following languages:

* Go: [`github.com/dagster-io/pulumi-dagsterplus/sdk/go`](https://pkg.go.dev/github.com/dagster-io/pulumi-dagsterplus/sdk/go)
* Python: [`pulumi-dagsterplus`](https://pypi.org/project/pulumi-dagsterplus/)
* YAML: Install the [Pulumi CLI](https://www.pulumi.com/docs/install/) and reference the provider directly in your YAML programs.

## Authentication

The Dagster+ provider requires a Dagster+ user token and your organization name to authenticate. You can generate a user token from your [Dagster+ account settings](https://dagster.cloud).

Configure the provider via environment variables:

```bash
export DAGSTER_CLOUD_API_TOKEN=your-user-token
export DAGSTER_CLOUD_ORGANIZATION=your-org-name
```

Or pass them explicitly in your Pulumi program:

{{< chooser language "go,python,yaml" >}}

{{% choosable language go %}}

```go
import (
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    dagsterplus "github.com/dagster-io/pulumi-dagsterplus/sdk/go/dagsterplus"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        _, err := dagsterplus.NewProvider(ctx, "dagsterplus", &dagsterplus.ProviderArgs{
            ApiToken:     pulumi.String("your-user-token"),
            Organization: pulumi.String("your-org-name"),
        })
        return err
    })
}
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi_dagsterplus as dagsterplus

provider = dagsterplus.Provider("dagsterplus",
    api_token="your-user-token",
    organization="your-org-name",
)
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
resources:
  dagsterplus:
    type: pulumi:providers:dagsterplus
    properties:
      apiToken: your-user-token
      organization: your-org-name
```

{{% /choosable %}}

{{< /chooser >}}

## Example Usage

Create a Dagster+ team:

{{< chooser language "go,python,yaml" >}}

{{% choosable language go %}}

```go
import (
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    dagsterplus "github.com/dagster-io/pulumi-dagsterplus/sdk/go/dagsterplus"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        _, err := dagsterplus.NewTeam(ctx, "my-team", &dagsterplus.TeamArgs{
            Name: pulumi.String("platform-engineers"),
        })
        return err
    })
}
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi_dagsterplus as dagsterplus

team = dagsterplus.Team("my-team",
    name="platform-engineers",
)
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
resources:
  my-team:
    type: dagsterplus:Team
    properties:
      name: platform-engineers
```

{{% /choosable %}}

{{< /chooser >}}
