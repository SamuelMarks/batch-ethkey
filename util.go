package main

import (
	"fmt"
	cr "github.com/SamuelMarks/batch-ethkey/crypto"
	"os"
	"reflect"
)

func GeneratePemKey(dir string) {
	// fmt.Printf("GeneratePemKey::dir = %s;\n", dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0700)
	}

	// Taken straight from babble:

	// Create the PEM key
	pemKey := cr.NewPemKey(dir)
	// Try a read, should get nothing
	key, err := pemKey.ReadKey()
	if err == nil {
		fmt.Fprint(os.Stderr, "ReadKey should generate an error")
		os.Exit(1)
	}
	if key != nil {
		fmt.Fprint(os.Stderr, "key is not nil")
		os.Exit(1)
	}
	// Initialize a key
	key, _ = cr.GenerateECDSAKey()
	if err := pemKey.WriteKey(key); err != nil {
		fmt.Fprintf(os.Stderr, "err: %v", err)
		os.Exit(1)
	}
	// Try a read, should get key
	nKey, err := pemKey.ReadKey()
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v", err)
		os.Exit(1)
	}
	if !reflect.DeepEqual(*nKey, *key) {
		fmt.Fprint(os.Stderr, "Keys do not match")
		os.Exit(1)
	}
}
