package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"tkido.com/goeth/config"
)

var client *ethclient.Client

func init() {
	fmt.Println("init")
	var err error
	client, err = ethclient.Dial(config.RpcURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to", config.RpcURL)

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Header is %s\n\n", header.Number.String())
}
