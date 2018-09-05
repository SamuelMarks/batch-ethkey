package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"
)

func GeneratePemKey(dir string, wg *sync.WaitGroup) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0700)
	}

	k, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	pub := fmt.Sprintf("0x%X", FromECDSAPub(&k.PublicKey))
	ioutil.WriteFile(path.Join(dir, "pub_key.pub"), []byte(pub), 0600)

	b, err := x509.MarshalECPrivateKey(k)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	pemBlock := &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}
	data := pem.EncodeToMemory(pemBlock)
	ioutil.WriteFile(path.Join(dir, "priv_key.pem"), data, 0600)

	wg.Done()
}
func FromECDSAPub(pub *ecdsa.PublicKey) []byte {
	return elliptic.Marshal(elliptic.P256(), pub.X, pub.Y)
}
