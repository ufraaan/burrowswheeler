package main

import (
	"burrowswheeler"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	decode := flag.Bool("d", false, "")
	primaryFlag := flag.Int("primary", -1, "")
	flag.Parse()

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fail(err)
	}

	if *decode {
		if *primaryFlag < 0 {
			if flag.NArg() != 1 {
				fail(fmt.Errorf("provide the primary index with -primary or as an argument"))
			}
			*primaryFlag, err = strconv.Atoi(flag.Arg(0))
			if err != nil {
				fail(err)
			}
		}
		output, err := bwt.InverseTransform(input, *primaryFlag)
		if err != nil {
			fail(err)
		}
		if _, err := os.Stdout.Write(output); err != nil {
			fail(err)
		}
		return
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
