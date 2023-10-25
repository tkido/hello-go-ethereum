package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"tkido.com/goeth/config"
)

func signature() {
	key, err := keystore.DecryptKey([]byte(config.Key), config.Secret)
	if err != nil {
		log.Fatalf("cannot decrypt key: %s", err)
	}

	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex())

	signature, err := crypto.Sign(hash.Bytes(), key.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hexutil.Encode(signature))

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}
	matches := bytes.Equal(sigPublicKey, crypto.FromECDSAPub(&key.PrivateKey.PublicKey))
	fmt.Println(matches)
}
