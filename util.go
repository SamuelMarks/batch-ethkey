package main

import (
	"fmt"
	cr "github.com/SamuelMarks/batch-ethkey/crypto"
	"os"
	"sync"
)

func GeneratePemKey(dir string, wg *sync.WaitGroup) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0700)
	}

	cr.NewPemKey(dir)

	pemKey := cr.NewPemKey(dir)
	key, _ := cr.GenerateECDSAKey()
	if err := pemKey.WriteKey(key); err != nil {
		fmt.Fprintf(os.Stderr, "err: %v", err)
	}

	wg.Done()
}
