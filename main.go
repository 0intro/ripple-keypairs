package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/rubblelabs/ripple/crypto"
)

var (
	prefix     = flag.String("prefix", "", "prefix")
	seed       = flag.String("seed", "", "seed")
	passphrase = flag.String("passphrase", "", "passphrase")
	nWorkers   = flag.Int("n", 1, "number of workers")
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage: ripple-keypairs [ -n nWorkers ] [ -p prefix | -s seed ]")
	os.Exit(1)
}

func main() {
	flag.Parse()

	if flag.NArg() != 0 {
		usage()
	}

	if *seed != "" && *passphrase != "" || *seed != "" && *prefix != "" || *passphrase != "" && *prefix != "" {
		usage()
	}

	if *seed == "" && *passphrase == "" && *prefix == "" {
		err := generateKeyPairRandom()
		if err != nil {
			log.Fatal(err)
		}
	}

	if *seed != "" {
		err := generateKeyPairSeed(*seed)
		if err != nil {
			log.Fatal(err)
		}
	}

	if *passphrase != "" {
		err := generateKeyPairPassphrase(*passphrase)
		if err != nil {
			log.Fatal(err)
		}
	}

	if *prefix != "" {
		if !isValidPrefix(*prefix) {
			log.Fatalf("prefix %s is not base54", *prefix)
		}
		if *nWorkers == 1 {
			err := generateKeyPairPrefix(*prefix)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := generateKeyPairPrefixParallel(*prefix)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func generateKeyPairRandom() error {
	b := make([]byte, 16)

	_, err := rand.Read(b)
	if err != nil {
		return err
	}

	seed, err := crypto.NewFamilySeed(b)
	if err != nil {
		return err
	}

	key, err := crypto.NewECDSAKey(seed.Payload())
	if err != nil {
		return err
	}

	err = printKeys(seed, key)
	if err != nil {
		return err
	}

	return nil
}

func generateKeyPairSeed(s string) error {
	seed, err := crypto.NewRippleHash(s)
	if err != nil {
		return err
	}

	key, err := crypto.NewECDSAKey(seed.Payload())
	if err != nil {
		return err
	}

	err = printKeys(seed, key)
	if err != nil {
		return err
	}

	return nil
}

func generateKeyPairPassphrase(s string) error {
	seed, err := crypto.GenerateFamilySeed(s)
	if err != nil {
		return err
	}

	key, err := crypto.NewECDSAKey(seed.Payload())
	if err != nil {
		return err
	}

	err = printKeys(seed, key)
	if err != nil {
		return err
	}

	return nil
}

func generateKeyPairPrefix(prefix string) error {
	b := make([]byte, 16)
	var seed crypto.Hash
	var key crypto.Key

	for {
		_, err := rand.Read(b)
		if err != nil {
			return err
		}
		seed, err = crypto.NewFamilySeed(b)
		if err != nil {
			return err
		}
		key, err = crypto.NewECDSAKey(seed.Payload())
		if err != nil {
			return err
		}
		if keyHasPrefix(key, prefix) {
			break
		}
	}

	err := printKeys(seed, key)
	if err != nil {
		return err
	}

	return nil
}

func generateKeyPairPrefixParallel(prefix string) error {
	bytes := make(chan []byte, 0)
	results := make(chan crypto.Hash, 0)

	for w := 1; w <= *nWorkers; w++ {
		go worker(w, bytes, results, prefix)
	}

	go mainWorker(bytes)

	for seed := range results {
		key, err := crypto.NewECDSAKey(seed.Payload())
		if err != nil {
			return err
		}

		err = printKeys(seed, key)
		if err != nil {
			return err
		}
	}

	return nil
}

func keyHasPrefix(key crypto.Key, prefix string) bool {
	var sequenceZero uint32
	accountId, err := crypto.AccountId(key, &sequenceZero)
	if err != nil {
		return false
	}
	return strings.HasPrefix(accountId.String(), prefix)
}

func mainWorker(bytes chan<- []byte) {
	for {
		b := make([]byte, 16)
		_, err := io.ReadFull(rand.Reader, b[:])
		if err != nil {
			log.Println(err)
			continue
		}
		bytes <- b
	}
}

func worker(id int, bytes <-chan []byte, results chan<- crypto.Hash, prefix string) {
	for b := range bytes {
		seed, err := crypto.NewFamilySeed(b)
		if err != nil {
			log.Println(err)
			continue
		}
		key, err := crypto.NewECDSAKey(seed.Payload())
		if err != nil {
			log.Println(err)
			continue
		}
		if keyHasPrefix(key, prefix) {
			results <- seed
		}
	}
}

func printKeys(seed crypto.Hash, key crypto.Key) error {
	fmt.Println("Seed (secret key)", seed)

	var sequenceZero uint32

	accountId, err := crypto.AccountId(key, &sequenceZero)
	if err != nil {
		return err
	}
	fmt.Println("AccountID", accountId)

	nodePublicKey, err := crypto.NodePublicKey(key)
	if err != nil {
		return err
	}
	fmt.Println("NodePublicKey", nodePublicKey)

	nodePrivateKey, err := crypto.NodePrivateKey(key)
	if err != nil {
		return err
	}
	fmt.Println("NodePrivateKey", nodePrivateKey)

	accountPublicKey, err := crypto.AccountPublicKey(key, &sequenceZero)
	if err != nil {
		return err
	}
	fmt.Println("AccountPublicKey", accountPublicKey)

	accountPrivateKey, err := crypto.AccountPrivateKey(key, &sequenceZero)
	if err != nil {
		return err
	}
	fmt.Println("AccountPrivateKey", accountPrivateKey)

	return nil
}

const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

func isValidPrefix(prefix string) bool {
	if len(prefix) < 1 {
		return false
	}
	if prefix[0] != 'r' {
		return false
	}
	for _, r := range prefix {
		if !strings.ContainsRune(alphabet, r) {
			return false
		}
	}
	return true
}
