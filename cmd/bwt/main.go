package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/ufraaan/burrowswheeler"
)

func main() {
	decode := flag.Bool("d", false, "")
	primaryFlag := flag.Int("primary", -1, "")
	blockMode := flag.Bool("block", false, "")
	blockSize := flag.Int("block-size", 65536, "")
	flag.Parse()

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fail(err)
	}

	if *decode {
		if *blockMode {
			output, err := bwt.DecodeBlocks(input)
			if err != nil {
				fail(err)
			}
			if _, err := os.Stdout.Write(output); err != nil {
				fail(err)
			}
			return
		}
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
	if *blockMode {
		output, err := bwt.EncodeBlocks(input, *blockSize)
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
