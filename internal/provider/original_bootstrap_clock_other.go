//go:build !linux

package provider

func bootstrapMonotonicMillis() (int64, error) { return 0, bootstrapError() }
