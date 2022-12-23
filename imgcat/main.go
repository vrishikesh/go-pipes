package main

import (
	"fmt"
	"imgcat/imgcat"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "missing paths of images to cat")
		os.Exit(2)
	}

	for _, path := range os.Args[1:] {
		if err := cat(path); err != nil {
			fmt.Fprintf(os.Stderr, "Could not cat %s: %v\n", path, err)
		}
	}
}

func cat(path string) error {
	f, err := os.Open(path)
	if f != nil {
		defer f.Close()
	}

	if err != nil {
		return fmt.Errorf("could now open image err: %s", err)
	}

	wc := imgcat.NewWriter(os.Stdout)
	if _, err = io.Copy(wc, f); err != nil {
		return err
	}
	return wc.Close()
}
