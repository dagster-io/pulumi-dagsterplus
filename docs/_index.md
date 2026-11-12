---
title: Dagster+
meta_desc: Provides an overview of the Dagster+ provider for Pulumi.
layout: package
---

## Overview

The Dagster+ provider for Pulumi can be used to provision and manage resources in [Dagster+](https://dagster.io/cloud), the managed cloud platform for Dagster. Use it to manage deployments, code locations, agents, teams, secrets, and more — all as code.

## Installation

The Dagster+ provider is available as a package in the following languages:

* JavaScript/TypeScript: [`@pulumi/dagsterplus`](https://www.npmjs.com/package/@pulumi/dagsterplus)
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

{{< chooser language "typescript,python,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as dagsterplus from "@pulumi/dagsterplus";

const provider = new dagsterplus.Provider("dagsterplus", {
    apiToken: "your-user-token",
    organization: "your-org-name",
});
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

{{< chooser language "typescript,python,yaml" >}}

{{% choosable language typescript %}}

```typescript
import * as dagsterplus from "@pulumi/dagsterplus";

const team = new dagsterplus.Team("my-team", {
    name: "platform-engineers",
});
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
