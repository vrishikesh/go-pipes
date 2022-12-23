package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	buf := new(bytes.Buffer)
	mw := io.MultiWriter(os.Stdout, os.Stderr, buf)

	fmt.Fprintf(mw, "hello\n")
	fmt.Println("from buffer:", buf)
}
