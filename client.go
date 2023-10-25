package main

import (
	"context"
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
	fmt.Println("connected to", rpcURL)

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Header is %s\n\n", header.Number.String())
}
