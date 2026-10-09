package blindfold

import (
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

//go:embed contract/*
var contract embed.FS

func ValidateContract() error {
	bundle, err := contract.ReadFile("contract/blindfold-contract-v1.txt")
	if err != nil {
		return err
	}
	lock, err := contract.ReadFile("contract/lock.json")
	if err != nil {
		return err
	}
	var pin struct {
		SHA256  string `json:"sha256"`
		Version int    `json:"version"`
		Commit  string `json:"commit"`
	}
	if err := json.Unmarshal(lock, &pin); err != nil {
		return err
	}
	sum := sha256.Sum256(bundle)
	if pin.Version != 1 || pin.Commit == "" || pin.SHA256 != hex.EncodeToString(sum[:]) {
		return fmt.Errorf("embedded Blindfold contract integrity failure")
	}
	return nil
}
