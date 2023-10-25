package config

import (
	"io"
	"log"
	"os"
)

const (
	ChainID       = 11155111
	RpcURL        = "https://eth-sepolia.public.blastapi.io"
	keyStorePath  = "./wallets/UTC--2023-10-25T00-01-06.160532000Z--fe57fabdbe012243b9bc7f1d9f76d44fc320cb4f"
	keySecretPath = "./.secret.txt"
)

var (
	Key    string
	Secret string
)

func init() {
	initKey()
	initSecret()
}

func initKey() {
	file, err := os.Open(keyStorePath)
	if err != nil {
		log.Fatalf("cannot open key store file %s", err)
	}
	defer file.Close()
	bs, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("cannot read key store file %s", err)
	}
	Key = string(bs)
}

func initSecret() {
	file, err := os.Open(keySecretPath)
	if err != nil {
		log.Fatalf("cannot open key secret file %s", err)
	}
	defer file.Close()
	bs, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("cannot read key secret file %s", err)
	}
	Secret = string(bs)
}
