package dagsterplus_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSchemaIsValid(t *testing.T) {
	data, err := os.ReadFile("cmd/pulumi-resource-dagsterplus/schema.json")
	require.NoError(t, err, "schema.json must exist — run 'make schema' to generate it")

	var spec schema.PackageSpec
	err = json.Unmarshal(data, &spec)
	require.NoError(t, err, "schema.json must be valid JSON conforming to Pulumi PackageSpec")

	pkg, diags, err := schema.BindSpec(spec, nil, schema.ValidationOptions{})
	require.NoError(t, err)
	require.False(t, diags.HasErrors(), diags.Error())

	assert.Equal(t, "dagsterplus", pkg.Name)
	assert.Len(t, pkg.Resources, 33, "expected 33 resources")
	assert.Len(t, pkg.Functions, 21, "expected 21 functions (data sources)")
}
