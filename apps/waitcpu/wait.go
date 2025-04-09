package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"crypto/md5"
	"log/slog"
	"net/http"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

func getThreadCPUMSec() (int64, error) {
	var usage syscall.Rusage
	if err := syscall.Getrusage(RUsageFlag, &usage); err != nil {
		slog.Error("getrusage syscall failed", "err", err)
		return -1, err
	}
	return (usage.Stime.Nano() + usage.Utime.Nano()) / int64(time.Millisecond), nil
}

func waitCPU(msec int) (int, error) {
	buf := [16]byte{}
	if _, err := rand.Read(buf[:]); err != nil {
		return 0, err
	}

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	started, err := getThreadCPUMSec()
	if err != nil {
		return 0, err
	}

	for {
		for i := 0; i < 10; i++ {
			buf = md5.Sum(buf[:])
		}

		now, err := getThreadCPUMSec()
		if err != nil {
			return 0, err
		}

		if now - int64(started) > int64(msec) {
			break
		}
	}

	return int(binary.BigEndian.Uint16(buf[:])), nil
}

func getRandomValue() (int32, error) {
	buf := make([]byte, 4)
	_, err := rand.Read(buf)
	if err != nil {
		slog.Error("failed to get random value", "err", err)
		return -1, err
	}
	return int32(binary.BigEndian.Uint32(buf)), nil
}


func WaitCPU(c echo.Context) error {
	ms := c.Param("ms")

	msec, err := strconv.Atoi(ms)
	if err != nil {
		return c.String(http.StatusBadRequest, "expected msec to be integer")
	}

	if msec < 0 {
		return c.String(http.StatusBadRequest, "msec must be a non-negative integer")
	}

	num, err := waitCPU(msec)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, fmt.Sprintf("%d", num))
}
