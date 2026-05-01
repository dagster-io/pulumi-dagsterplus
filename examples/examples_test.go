package examples_test

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getEnvOrSkip(t *testing.T, key string) string {
	t.Helper()
	val := os.Getenv(key)
	if val == "" {
		t.Skipf("skipping: %s not set", key)
	}
	return val
}

func TestBasicExample(t *testing.T) {
	apiToken := getEnvOrSkip(t, "DAGSTER_CLOUD_API_TOKEN")
	org := getEnvOrSkip(t, "DAGSTER_CLOUD_ORGANIZATION")

	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir: "basic",
		LocalProviders: []integration.LocalDependency{
			{Package: "dagsterplus", Path: ".."},
		},
		Config: map[string]string{
			"dagsterplus:organization": org,
		},
		Secrets: map[string]string{
			"dagsterplus:apiToken": apiToken,
		},
		Quick:            true,
		DestroyOnCleanup: true,
	})
}
