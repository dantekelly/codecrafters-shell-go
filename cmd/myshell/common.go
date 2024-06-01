package main

import (
	"fmt"
	"os"
)

func Print(s string, args ...interface{}) {
	if len(args) == 0 {
		fmt.Fprint(os.Stdout, s)
		return
	}
	fmt.Fprintf(os.Stdout, s, args...)
}
