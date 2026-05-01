# Contributing

## Prerequisites

- [Go 1.24+](https://golang.org/dl/)
- [Pulumi CLI](https://www.pulumi.com/docs/install/)

## Building

```bash
# Build and install the provider locally
make install

# Or step by step:
make tfgen      # Build the schema generator
make schema     # Generate schema.json and bridge-metadata.json
make provider   # Build the provider binary
```

## Testing

```bash
# Fast schema validity test (no credentials needed)
make test_provider

# Full end-to-end integration test (requires credentials)
export DAGSTER_CLOUD_ORGANIZATION="your-org"
export DAGSTER_CLOUD_API_TOKEN="your-token"
make test_examples
```

## Project Structure

```
pulumi-dagsterplus/
├── .github/workflows/
│   ├── ci.yml                              # Runs on PRs and pushes to main
│   └── release.yml                         # Triggered by version tags
├── provider/
│   ├── resources.go                        # Provider configuration and token mapping
│   ├── schema_test.go                      # Schema validity test
│   └── cmd/
│       ├── pulumi-tfgen-dagsterplus/       # Schema generator binary
│       └── pulumi-resource-dagsterplus/    # Provider runtime binary
├── examples/
│   ├── basic/                              # Basic usage example
│   └── examples_test.go                   # Integration tests
├── .goreleaser.yml                         # Cross-platform release config
└── Makefile
```

## Releasing

Releases are automated via [goreleaser](https://goreleaser.com/) and GitHub Actions. Pushing a version tag triggers the release workflow, which builds cross-platform binaries and publishes them to GitHub Releases.

```bash
git tag v0.1.0
git push origin v0.1.0
```

Pulumi will then be able to auto-download the provider from GitHub Releases.

## Updating the Upstream Provider

When a new version of the Terraform provider is released:

1. Update the version in `go.mod`:
   ```
   github.com/dagster-io/terraform-provider-dagsterplus vX.Y.Z
   ```
2. Run `go mod tidy`
3. Run `make install` to regenerate the schema and rebuild
4. Run `make test_provider && make test_examples` to verify
