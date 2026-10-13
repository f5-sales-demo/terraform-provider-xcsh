// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package main

import (
	"encoding/json"
	"os"
)

func verifySuiteInventory(report *Report, names []string) {
	expected := map[string]bool{}
	for _, name := range names {
		if name == "" || expected[name] {
			report.Stream.InvalidEvents++
			continue
		}
		expected[name] = true
	}
	if len(expected) == 0 {
		report.Stream.InvalidEvents++
	}
	for _, test := range report.Tests {
		if !expected[test.Name] {
			report.Stream.InvalidEvents++
		} else {
			delete(expected, test.Name)
		}
	}
	report.Stream.IncompleteTests += len(expected)
	if len(expected) > 0 || report.Stream.InvalidEvents > 0 {
		report.Stream.Complete = false
	}
}
func loadSuiteInventory(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var names []string
	err = json.Unmarshal(data, &names)
	return names, err
}
