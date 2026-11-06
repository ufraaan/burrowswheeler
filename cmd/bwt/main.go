package main

import (
	"burrowswheeler"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Parse()

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fail(err)
	}

	output, primary := bwt.Transform(input)
	if _, err := os.Stdout.Write(output); err != nil {
		fail(err)
	}
	if _, err := fmt.Fprintf(os.Stderr, "\nprimary index: %d\n", primary); err != nil {
		fail(err)
	}
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
