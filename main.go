package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"sync"
)

func main() {
	dirPtr := flag.String("dir", ":/required",
		"parent directory containing numbered subdirectories containing keys")
	nPtr := flag.Uint64("n", 5, "number subdirectories (containing keys) to create")

	ensureCliArgs([]string{"dir", "n"})

	abspath, err := filepath.Abs(*dirPtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Abs(%q) error: %v", *dirPtr, err)
		os.Exit(1)
	}

	if _, err := os.Stat(abspath); os.IsNotExist(err) {
		os.MkdirAll(abspath, 0700)
	}

	wg := new(sync.WaitGroup)
	for i := uint64(0); i < *nPtr; i++ {
		wg.Add(1)
		go GeneratePemKey(path.Join(abspath, strconv.FormatUint(i, 10)), wg)
	}
	wg.Wait()
}

func ensureCliArgs(required []string) {
	// From: https://stackoverflow.com/a/31795922
	flag.Parse()
	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			// or possibly use `log.Fatalf` instead of:
			fmt.Fprintf(os.Stderr, "missing required -%s argument/flag\n", req)
			os.Exit(2) // the same exit code flag.Parse uses
		}
	}
}
