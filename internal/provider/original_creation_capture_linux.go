//go:build linux

package provider

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func captureOwned(info os.FileInfo) bool {
	stat, ok := info.Sys().(*syscall.Stat_t)
	return ok && stat.Uid == uint32(os.Getuid()) && (info.IsDir() || stat.Nlink == 1)
}

func originalCreationProcess() (OriginalCreationProcess, error) {
	var p OriginalCreationProcess
	p.PID = os.Getpid()
	p.ParentPID = os.Getppid()
	stat, err := os.ReadFile("/proc/self/stat")
	if err != nil {
		return p, captureError()
	}
	end := strings.LastIndex(string(stat), ") ")
	if end < 0 {
		return p, captureError()
	}
	fields := strings.Fields(string(stat)[end+2:])
	if len(fields) < 20 {
		return p, captureError()
	}
	ticks, err := strconv.ParseUint(fields[19], 10, 64)
	if err != nil || ticks == 0 {
		return p, captureError()
	}
	p.StartTicks = fields[19]
	boot, err := os.ReadFile("/proc/sys/kernel/random/boot_id")
	if err != nil || len(strings.TrimSpace(string(boot))) != 36 {
		return p, captureError()
	}
	bootHash := sha256.Sum256(boot)
	p.BootIDSHA256 = hex.EncodeToString(bootHash[:])
	exe, err := os.Open("/proc/self/exe")
	if err != nil {
		return p, captureError()
	}
	defer func() { _ = exe.Close() }()
	info, err := exe.Stat()
	if err != nil || !info.Mode().IsRegular() {
		return p, captureError()
	}
	identity, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return p, captureError()
	}
	p.ExecutableDevice = uint64(identity.Dev)
	p.ExecutableInode = identity.Ino
	digest := sha256.New()
	if _, err = io.Copy(digest, exe); err != nil {
		return p, captureError()
	}
	p.ExecutableSHA256 = hex.EncodeToString(digest.Sum(nil))
	after, err := exe.Stat()
	if err != nil || !os.SameFile(info, after) || info.Size() != after.Size() || !info.ModTime().Equal(after.ModTime()) {
		return p, captureError()
	}
	if err = exe.Close(); err != nil {
		return p, captureError()
	}
	return p, nil
}
