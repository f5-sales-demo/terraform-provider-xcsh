// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package discovery

import "time"

// DeleteWithRetry calls del up to attempts times, sleeping backoff (doubled after each
// failure) between tries, and returns the error from the last attempt (nil as soon as one
// succeeds). attempts < 1 is treated as 1.
//
// Discovery creates real `tf-discover-*` probe objects in a SHARED tenant namespace
// (securemesh_site_v2 lands in "system", right next to live demo sites), so a DELETE that
// fails once — a transient 5xx, or the brief referential BAD_REQUEST the F5 XC API returns
// while an object is still referenced — must be retried and, if it still fails, reported
// to the caller. Discarding the error strands the probe object forever.
func DeleteWithRetry(del func() error, attempts int, backoff time.Duration) error {
	if attempts < 1 {
		attempts = 1
	}
	var err error
	for i := 0; i < attempts; i++ {
		if i > 0 {
			time.Sleep(backoff)
			backoff *= 2
		}
		if err = del(); err == nil {
			return nil
		}
	}
	return err
}
