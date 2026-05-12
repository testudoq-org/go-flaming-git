package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/stephenhstewart/fgit/internal/coverage"
)

func main() {
	inPath := flag.String("in", "coverage.out", "input Go coverage profile")
	outPath := flag.String("out", "coverage.lcov", "output LCOV file")
	flag.Parse()

	if err := run(*inPath, *outPath); err != nil {
		exitf("%v", err)
	}
}

func run(inPath, outPath string) error {
	inFile, err := os.Open(inPath)
	if err != nil {
		return fmt.Errorf("open input %q: %w", inPath, err)
	}
	defer inFile.Close()

	outFile, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("create output %q: %w", outPath, err)
	}
	defer outFile.Close()

	if err := coverage.ConvertProfileToLCOV(inFile, outFile); err != nil {
		return fmt.Errorf("convert to lcov: %w", err)
	}

	return nil
}

func exitf(format string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
