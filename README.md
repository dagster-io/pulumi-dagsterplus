# Pulumi Dagster+ Provider

A [Pulumi](https://www.pulumi.com/) provider for managing [Dagster+](https://dagster.io/) (Dagster Cloud) resources — deployments, code locations, teams, users, tokens, alert policies, and more.

This provider is built on top of the [Dagster+ Terraform provider](https://github.com/dagster-io/terraform-provider-dagsterplus) using the [Pulumi Terraform Bridge](https://github.com/pulumi/pulumi-terraform-bridge).

## Configuration

The provider requires your Dagster+ organization name and an API token. These can be set via config or environment variables.

```bash
export DAGSTER_CLOUD_ORGANIZATION="your-org"
export DAGSTER_CLOUD_API_TOKEN="your-token"
```

Or via Pulumi config:

```bash
pulumi config set dagsterplus:organization "your-org"
pulumi config set --secret dagsterplus:apiToken "your-token"
```

Generate an API token at **Dagster+ → Account Settings → API Tokens**.

## Usage

### Pulumi YAML

```yaml
name: my-dagster-infra
runtime: yaml

resources:
  dataEngineering:
    type: dagsterplus:index:Team
    properties:
      name: data-engineering

  prodDeployment:
    type: dagsterplus:index:Deployment
    properties:
      name: production
```

### TypeScript

```typescript
import * as dagsterplus from "@dagster-io/pulumi-dagsterplus";

const team = new dagsterplus.index.Team("data-engineering", {
    name: "data-engineering",
});
```

### Python

```python
import pulumi_dagsterplus as dagsterplus

team = dagsterplus.index.Team("data-engineering",
    name="data-engineering",
)
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).
