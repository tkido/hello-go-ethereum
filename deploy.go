package main

import (
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"tkido.com/goeth/config"
	"tkido.com/goeth/storage"
)

func deploy() {
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(config.Key), config.Secret, big.NewInt(config.ChainID))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	// Deploy the contract passing the newly created `auth` and `client` vars
	address, tx, instance, err := storage.DeployStorage(auth, client)
	if err != nil {
		log.Fatalf("Failed to deploy new storage contract: %v", err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	time.Sleep(250 * time.Millisecond) // Allow it to be processed by the local node :P
	_ = instance
}
