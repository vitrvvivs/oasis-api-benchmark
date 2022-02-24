package main

import (
	"fmt"
	"os"
)

func Fail(s ...interface{}) {
	fmt.Fprintln(os.Stderr, s...)
	os.Exit(1)
}

func Warn(s ...interface{}) {
	fmt.Fprintln(os.Stderr, s...)
}
