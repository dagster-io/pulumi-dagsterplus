# Pulumi Dagster+ Provider

The Dagster+ resource provider for Pulumi lets you manage [Dagster+](https://dagster.io/) (Dagster Cloud) resources — deployments, code locations, teams, users, tokens, alert policies, and more.

This provider is built on top of the [Dagster+ Terraform provider](https://github.com/dagster-io/terraform-provider-dagsterplus) using the [Pulumi Terraform Bridge](https://github.com/pulumi/pulumi-terraform-bridge).

## Installing

This package is available in the following languages and packaging formats.

### TypeScript / JavaScript

To use from TypeScript or JavaScript, install using `npm`:

```bash
npm install @pulumi/dagsterplus
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_dagsterplus
```

### Pulumi YAML

No installation required. Reference the provider directly in your `Pulumi.yaml`:

```yaml
plugins:
  providers:
    - name: dagsterplus
```

## Configuration

The provider requires your Dagster+ organization name and an API token.

```bash
pulumi config set dagsterplus:organization "your-org"
pulumi config set --secret dagsterplus:apiToken "your-token"
```

Or via environment variables:

```bash
export DAGSTER_CLOUD_ORGANIZATION="your-org"
export DAGSTER_CLOUD_API_TOKEN="your-token"
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).
