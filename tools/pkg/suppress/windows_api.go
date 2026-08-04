// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build windows

package suppress

import (
	"os"
	"runtime"
	"syscall"
	"unsafe"
)

const (
	lockFileExclusiveLock   = 0x2
	moveFileReplaceExisting = 0x1
	moveFileWriteThrough    = 0x8
)

var (
	kernel32DLL      = syscall.NewLazyDLL("kernel32.dll")
	lockFileExProc   = kernel32DLL.NewProc("LockFileEx")
	unlockFileExProc = kernel32DLL.NewProc("UnlockFileEx")
	moveFileExWProc  = kernel32DLL.NewProc("MoveFileExW")
)

func lockFileWindows(file *os.File) error {
	overlapped := new(syscall.Overlapped)
	result, _, callErr := lockFileExProc.Call(
		file.Fd(),
		lockFileExclusiveLock,
		0,
		1,
		0,
		uintptr(unsafe.Pointer(overlapped)),
	)
	runtime.KeepAlive(overlapped)
	return windowsCallError(result, callErr)
}

func unlockFileWindows(file *os.File) error {
	overlapped := new(syscall.Overlapped)
	result, _, callErr := unlockFileExProc.Call(
		file.Fd(),
		0,
		1,
		0,
		uintptr(unsafe.Pointer(overlapped)),
	)
	runtime.KeepAlive(overlapped)
	return windowsCallError(result, callErr)
}

func replaceFileWindows(temporaryPath, path string) error {
	from, err := syscall.UTF16PtrFromString(temporaryPath)
	if err != nil {
		return err
	}
	to, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return err
	}
	result, _, callErr := moveFileExWProc.Call(
		uintptr(unsafe.Pointer(from)),
		uintptr(unsafe.Pointer(to)),
		moveFileReplaceExisting|moveFileWriteThrough,
	)
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
	return windowsCallError(result, callErr)
}

func windowsCallError(result uintptr, callErr error) error {
	if result != 0 {
		return nil
	}
	if callErr != nil && callErr != syscall.Errno(0) {
		return callErr
	}
	return syscall.EINVAL
}
