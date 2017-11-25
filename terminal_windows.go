// +build windows

package main

import (
	"syscall"

	"golang.org/x/sys/windows"
)

const newline = "\r\n"

// TerminalDimensions returns the dimensions of the given terminal.
func TerminalDimensions() (columns, rows int, err error) {
	in, err := syscall.Open("CONOUT$", syscall.O_RDWR, 0)
	if err != nil {
		return -1, -1, err
	}

	var info windows.ConsoleScreenBufferInfo
	if err := windows.GetConsoleScreenBufferInfo(windows.Handle(in), &info); err != nil {
		return -1, -1, err
	}
	return int(info.Size.X), int(info.Size.Y), nil
}
