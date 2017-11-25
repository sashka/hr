// +build !windows,!plan9,!solaris

package main

import (
	"os"

	"golang.org/x/sys/unix"
)

const newline = "\n"

// TerminalDimensions returns the dimensions of the given terminal.
func TerminalDimensions() (columns, rows int, err error) {
	ws, err := unix.IoctlGetWinsize(int(os.Stdin.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return -1, -1, err
	}
	return int(ws.Col), int(ws.Row), nil
}
