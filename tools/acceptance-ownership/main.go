// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
	"os"
	"time"
)

func main() {
	mode := flag.String("mode", "cleanup", "init or cleanup")
	journal := flag.String("journal", "", "Private journal path")
	run := flag.String("run-id", "", "Run identity")
	receipt := flag.String("receipt", "", "Sanitized cleanup receipt path")
	flag.Parse()
	var err error
	switch *mode {
	case "reference-probe":
		err = acctest.ProbeReferenceResolution(context.Background())
	case "provision":
		err = acctest.ProvisionDisposablePrerequisites(context.Background())
	case "schema":
		err = acctest.ExportAcceptanceSchema(*receipt)
	case "init":
		err = acctest.InitializeOwnershipRun(*journal, *run)
	case "cleanup":
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
		defer cancel()
		err = acctest.CleanupOwnershipRun(ctx, *journal, *run, *receipt)
	default:
		err = fmt.Errorf("unknown ownership operation")
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ownership operation failed; inspect private evidence.")
		os.Exit(1)
	}
}
