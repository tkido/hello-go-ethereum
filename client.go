package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

const rpcURL = "http://localhost:8545"

var client *ethclient.Client

func init() {
	fmt.Println("init")
	var err error
	client, err = ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection to:", rpcURL)
}
