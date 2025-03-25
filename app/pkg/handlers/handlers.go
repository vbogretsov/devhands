package handlers

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getThreadCPUMSec() (int64, error) {
	var usage syscall.Rusage
	if err := syscall.Getrusage(syscall.RUSAGE_THREAD, &usage); err != nil {
	// if err := syscall.Getrusage(syscall.RUSAGE_SELF, &usage); err != nil {
		slog.Error("getrusage syscall failed", "err", err)
		return -1, err
	}
	return (usage.Stime.Nano() + usage.Utime.Nano()) / int64(time.Millisecond), nil
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
		return c.String(http.StatusBadRequest, "Invalid msec parameter. Must be an integer.")
	}

	if msec < 0 {
		return c.String(http.StatusBadRequest, "msec must be a non-negative integer.")
	}


	cpuStarted, err := getThreadCPUMSec()
	if err != nil {
		return err
	}
	timeStarted := time.Now().UnixMilli()

	cpuFinished := int64(0)
	timeFinished := int64(0)
	
	val := int32(0)

	iters := 0
	for {
		iters += 1

		rngVal, err := getRandomValue()
		if err != nil {
			return err
		}

		val = val ^ rngVal

		now, err := getThreadCPUMSec()
		if err != nil {
			return err
		}

		if now - cpuStarted > int64(msec) {
			cpuFinished = now
			timeFinished = time.Now().UnixMilli()
			break
		}
	}

	return c.String(
		http.StatusOK,
		fmt.Sprintf(
			"value=%d cycles=%d cpuMs=%d timeMs=%d",
			val,
			iters,
			cpuFinished - cpuStarted,
			timeFinished - timeStarted,
		),
	)
}
