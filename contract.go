package main

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"tkido.com/goeth/config"
	"tkido.com/goeth/storage"
)

func store(s string, number *big.Int) {
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(config.Key), config.Secret, big.NewInt(config.ChainID))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	address := common.HexToAddress(s)
	instance, err := storage.NewStorage(address, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract is loaded")
	tx, err := instance.Store(auth, number)
	if err != nil {
		log.Fatalf("Failed to update value: %v", err)
	}
	fmt.Printf("Update pending: 0x%x\n", tx.Hash())
}

func retrieve(s string) {
	address := common.HexToAddress(s)
	instance, err := storage.NewStorage(address, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract is loaded")
	big, err := instance.Retrieve(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("value is", big)
}
