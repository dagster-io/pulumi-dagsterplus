// Command strip-schema-languages removes unsupported language entries from a
// generated Pulumi schema.
//
// The Pulumi Terraform bridge unconditionally emits a "nodejs" entry in the
// schema's language map (see genPulumiSchema in pulumi-terraform-bridge), but
// this provider does not publish a TypeScript/npm SDK. Leaving the entry in
// place would cause the Pulumi Registry to advertise an installable SDK that
// does not exist. We strip it here as a reproducible post-processing step so
// that `make schema` output stays consistent.
//
// The schema is round-tripped through schema.PackageSpec and re-marshalled with
// the same settings the bridge uses (json.MarshalIndent with a four-space
// indent), so removing the "nodejs" key is the only change to the file.
package main

import (
	"encoding/json"
	"fmt"
	"os"

	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

// languagesToStrip lists language keys that must not appear in the generated
// schema because we do not publish an SDK for them.
var languagesToStrip = []string{"nodejs"}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: strip-schema-languages <schema.json>")
		os.Exit(1)
	}
	path := os.Args[1]

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read %s: %v\n", path, err)
		os.Exit(1)
	}

	var spec pschema.PackageSpec
	if err := json.Unmarshal(data, &spec); err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse %s: %v\n", path, err)
		os.Exit(1)
	}

	for _, lang := range languagesToStrip {
		delete(spec.Language, lang)
	}

	out, err := json.MarshalIndent(spec, "", "    ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to marshal schema: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(path, out, 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write %s: %v\n", path, err)
		os.Exit(1)
	}
}
