package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/node"
)

func makePasswordList(passwordFile string) []string {
	text, err := ioutil.ReadFile(passwordFile)
	if err != nil {
		utils.Fatalf("Failed to read password file: %v", err)
	}
	lines := strings.Split(string(text), "\n")
	// Sanitise DOS line endings.
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], "\r")
	}
	return lines

}

func accountCreate(dir string, password string, wg *sync.WaitGroup) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0700)
		if err != nil {
			panic(err)
		}
	}

	cfg := node.DefaultConfig
	cfg.Name = "ftm"
	cfg.HTTPModules = append(cfg.HTTPModules, "eth", "shh")
	cfg.WSModules = append(cfg.WSModules, "eth", "shh")
	cfg.IPCPath = "geth.ipc"

	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP

	address, err := keystore.StoreKey(path.Join(dir, "keystore"), password, scryptN, scryptP)

	if err != nil {
		utils.Fatalf("Failed to create account: %v", err)
	}

	addr := fmt.Sprintf("%x", address)

	err = ioutil.WriteFile(path.Join(dir, "evm-address"), []byte(addr), 0600)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(path.Join(dir, "pwd.txt"), []byte(password), 0600)
	if err != nil {
		panic(err)
	}
	evmlFile, err := os.Create(path.Join(dir, "evml.toml"))
	if err != nil {
		panic(err)
	}
	if _, err := fmt.Fprint(evmlFile, "[lachesis]\nstore = true\nheartbeat = \"50ms\"\ntimeout = \"200ms\"\n"); err != nil {
		panic(err)
	}

	wg.Done()

	if err := evmlFile.Close(); err != nil {
		panic(err)
	}
}
