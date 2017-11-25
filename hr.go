package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	cols, _, err := TerminalDimensions()
	if err != nil {
		fmt.Println(err)
		return
	}

	w := bufio.NewWriter(os.Stdout)

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "#")
	}

	for _, str := range os.Args[1:] {
		var runes int
		r := NewInfiniteRuneReader(str)
		if r == nil {
			continue
		}

		for runes < cols {
			ch, _, err := r.ReadRune()
			if err != nil {
				panic(err)
			}
			if _, err := w.WriteRune(ch); err != nil {
				panic(err)
			}
			runes++
		}
		w.Write([]byte(newline))
		w.Flush()
	}

}
