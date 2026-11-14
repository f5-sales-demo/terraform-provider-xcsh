//go:build linux

package provider

import "golang.org/x/sys/unix"

func bootstrapMonotonicMillis() (int64, error) {
	var value unix.Timespec
	if err := unix.ClockGettime(unix.CLOCK_MONOTONIC, &value); err != nil {
		return 0, err
	}
	// Timespec fields are int32 on Linux 386 and ARM; widen before arithmetic.
	return int64(value.Sec)*1000 + int64(value.Nsec)/1000000, nil
}
