package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/SamuelMarks/go-cidr/cidr"
	"io/ioutil"
	"net"
	"os"
	"path"
	"sync"
)

func GenerateKeyPair(dir string, wg *sync.WaitGroup) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0700)
		if err != nil {
			panic(err)
		}
	}

	k, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		_, err := fmt.Fprint(os.Stderr, err)
		if err != nil {
			panic(err)
		}
		return
	}
	pub := fmt.Sprintf("0x%X", FromECDSAPub(&k.PublicKey))
	err = ioutil.WriteFile(path.Join(dir, "pub_key.pub"), []byte(pub), 0600)
	if err != nil {
		panic(err)
	}

	b, err := x509.MarshalECPrivateKey(k)
	if err != nil {
		_, err := fmt.Fprint(os.Stderr, err)
		if err != nil {
			panic(err)
		}
		return
	}
	pemBlock := &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}
	data := pem.EncodeToMemory(pemBlock)
	err = ioutil.WriteFile(path.Join(dir, "priv_key.pem"), data, 0600)
	if err != nil {
		panic(err)
	}

	wg.Done()
}

func FromECDSAPub(pub *ecdsa.PublicKey) []byte {
	return elliptic.Marshal(elliptic.P256(), pub.X, pub.Y)
}

func visitF(seen uint64, port uint64, ip *net.IP, incPort bool) func(string, os.FileInfo, error) error {
	return func(p string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !f.IsDir() && path.Base(p) == "pub_key.pub" {
			pubKeyHex, e := ioutil.ReadFile(p)
			if e != nil {
				return e
			}
			seen--

			var endl string

			if seen > 0 {
				endl = ",\n"
			} else {
				endl = "\n"
			}

			if incPort {
				port++
			} else {
				ipInc := cidr.Inc(*ip)
				ip = &ipInc
			}
			host := ip.String() // strings.Join(ipInc[12:], ".")

			fmt.Printf("  {\n    \"NetAddr\": \"%s:%d\",\n    \"PubKeyHex\": \"%s\"\n  }%s",
				host, port, pubKeyHex, endl)
			if ip == nil {
				port++
			}
		}
		return nil
	}
}
