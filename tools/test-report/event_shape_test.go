// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package main

import (
	"strings"
	"testing"
)

func TestPackageEventsCannotMasqueradeAsTestLifecycle(t *testing.T) {
	for _, action := range []string{"run", "pause", "cont"} {
		stream := goTestEvent("start", "", "") + "\n" + goTestEvent(action, "", "") + "\n" + goTestEvent("pass", "", "")
		if r := parseTestOutputReader(strings.NewReader(stream), 10, true); r.Stream.Complete {
			t.Errorf("invalid package action %s accepted", action)
		}
	}
}
