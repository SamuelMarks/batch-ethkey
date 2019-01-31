package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/SamuelMarks/go-cidr/cidr"
)

var hosts arrayHosts

func main() {
	dirPtr := flag.String("dir", ":/required",
		"parent directory containing numbered subdirectories containing keys")
	nPtr := flag.Uint64("n", 5, "number subdirectories (containing keys) to create")
	portStartPtr := flag.Uint64("port-start", 12000, "port to start counting at")
	networkPtr := flag.String("network", "",
		"network in CIDR, with start address e.g.: 192.168.0.1/16")
	incPort := flag.Bool("inc-port", false, "Increment port numbers instead of IP addresses")
	evm := flag.Bool("evm", false, "Generate evm config files: genesis.json, evml.toml and keystore")
	passwordFile := flag.String("pwd", "pwd.txt", "path to account passwords file for evm configuration. Default: pwd.txt")
	flag.Var(&hosts, "host", "hostname or IP address and base port of a node. Can be specified multiply times")
	var passwords []string

	ensureCliArgs([]string{"dir", "n", "network"})

	ip := net.ParseIP(*networkPtr)
	if ip4 := ip.To4(); ip4 == nil {
		_, err := fmt.Fprintf(os.Stderr, "non IPv4 address %s is unsupported\n", ip)
		if err != nil {
			panic(err)
		}
		os.Exit(6)
	}

	switch uint64(len(hosts)) {
	case 0:
		port := *portStartPtr
		for i := uint64(0); i < *nPtr; i++ {
			hosts = append(hosts, fmt.Sprintf("%s:%d", ip.String(), port))
			if *incPort {
				port++
			} else {
				ip = cidr.Inc(ip)
			}
		}
	case *nPtr:
		port := *portStartPtr
		for i := uint64(0); i < *nPtr; i++ {
			if !strings.ContainsAny(hosts[i], ":") {
				hosts[i] = fmt.Sprintf("%s:%d", hosts[i], port)
				if *incPort {
					port++
				}
			}
		}
	default:
		panic(fmt.Errorf("number of --host flags speified (%d) doesn't match the value of n (%d)",
			len(hosts), *nPtr))
	}

	abspath, err := filepath.Abs(*dirPtr)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Abs(%q) error: %v", *dirPtr, err)
		if err != nil {
			panic(err)
		}
		os.Exit(1)
	}

	if _, err := os.Stat(abspath); os.IsNotExist(err) {
		err := os.MkdirAll(abspath, 0700)
		if err != nil {
			panic(err)
		}
	}

	if *evm {
		passwords = makePasswordList(*passwordFile)
	}
	nPasswords := uint64(len(passwords))

	wg := new(sync.WaitGroup)
	nStr := strconv.FormatUint(*nPtr, 10)
	padding := fmt.Sprintf("%%0%dd", len(nStr))
	for i := uint64(0); i < *nPtr; i++ {
		nodeDir := path.Join(abspath, fmt.Sprintf(padding, i))
		if _, err := os.Stat(nodeDir); os.IsNotExist(err) {
			err := os.Mkdir(nodeDir, 0700)
			if err != nil {
				panic(err)
			}
		}
		wg.Add(1)
		go GenerateKeyPair(nodeDir, wg)
		wg.Add(1)
		go accountCreate(path.Join(nodeDir, "eth"), passwords[i%nPasswords], wg)
	}
	wg.Wait()

	peersFile, err := os.Create(path.Join(abspath, "peers.json"))
	if err != nil {
		panic(err)
	}
	defer peersFile.Close()

	genesisFile, err := os.Create(path.Join(abspath, "genesis.json"))
	if err != nil {
		panic(err)
	}
	defer genesisFile.Close()

	fmt.Fprintln(peersFile, "[")
	fmt.Fprintln(genesisFile, "{\n\t\"alloc\": {")
	err = filepath.Walk(*dirPtr, visitF(peersFile, genesisFile, *nPtr, hosts))
	fmt.Fprintln(peersFile, "]")
	fmt.Fprintln(genesisFile, "\t}\n}")

	if err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			panic(err)
		}
	}
}

func ensureCliArgs(required []string) {
	// From: https://stackoverflow.com/a/31795922
	flag.Parse()
	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			// or possibly use `log.Fatalf` instead of:
			_, err := fmt.Fprintf(os.Stderr, "missing required -%s argument/flag\n", req)
			if err != nil {
				panic(err)
			}
			os.Exit(2) // the same exit code flag.Parse uses
		}
	}
}
