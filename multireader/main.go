package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	header := strings.NewReader("<msg>")
	body := strings.NewReader("hello")
	footer := strings.NewReader("</msg>")

	r := io.MultiReader(header, body, footer)
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		panic(err)
	}
}
