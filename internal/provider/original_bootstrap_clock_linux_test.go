//go:build linux

package provider

import (
	"testing"

	"golang.org/x/sys/unix"
)

func TestBootstrapMonotonicMillisClockBounds(t *testing.T) {
	var before, after unix.Timespec
	if err := unix.ClockGettime(unix.CLOCK_MONOTONIC, &before); err != nil {
		t.Fatal(err)
	}
	got, err := bootstrapMonotonicMillis()
	if err != nil {
		t.Fatal(err)
	}
	if err := unix.ClockGettime(unix.CLOCK_MONOTONIC, &after); err != nil {
		t.Fatal(err)
	}
	minimum := int64(before.Sec)*1000 + int64(before.Nsec)/1000000
	maximum := int64(after.Sec)*1000 + int64(after.Nsec)/1000000
	if got < minimum || got > maximum {
		t.Fatalf("monotonic milliseconds %d outside sampled bounds [%d, %d]", got, minimum, maximum)
	}
}
