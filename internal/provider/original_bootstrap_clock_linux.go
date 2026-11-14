//go:build linux

package provider

import "golang.org/x/sys/unix"

func bootstrapMonotonicMillis() (int64, error) {
	var value unix.Timespec
	if err := unix.ClockGettime(unix.CLOCK_MONOTONIC, &value); err != nil {
		return 0, err
	}
	return value.Sec*1000 + value.Nsec/1000000, nil
}
