// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"os"
	"time"
)

func InitializeOwnershipRun(path, run string) error {
	_, err := CreateOwnershipJournal(path, run)
	return err
}
func CleanupOwnershipRun(ctx context.Context, path, run, receiptPath string) error {
	receipt := CleanupReceipt{Version: 1, Failures: 1}
	j, err := LoadOwnershipJournal(path, run)
	if err == nil {
		c := client.NewClient(os.Getenv(EnvXCSHURL), os.Getenv(EnvXCSHToken))
		if os.Getenv(EnvXCSHURL) == "" || os.Getenv(EnvXCSHToken) == "" {
			err = errors.New("cleanup credentials unavailable")
		} else {
			receipt, err = j.Cleanup(ctx, apiOwnershipBackend{c: c, run: j.RunID}, 12, 5*time.Second)
		}
	}
	data, marshalErr := json.MarshalIndent(receipt, "", "  ")
	if marshalErr != nil {
		return marshalErr
	}
	if writeErr := os.WriteFile(receiptPath, append(data, '\n'), 0600); writeErr != nil {
		return writeErr
	}
	return err
}
