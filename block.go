package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
)

func checkBlock(x int64) {
	blockNumber := big.NewInt(x)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("block.Number()", block.Number().Uint64())              // 5671744
	fmt.Println("block.Time()", block.Time())                           // 1527211625
	fmt.Println("block.Difficulty()", block.Difficulty().Uint64())      // 3217000136609065
	fmt.Println("block.Hash()", block.Hash().Hex())                     // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println("len(block.Transactions())", len(block.Transactions())) // 144

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("TransactionCount %d\n\n", count) // 144

	for _, tx := range block.Transactions() {
		fmt.Println("tx.Hash()", tx.Hash().Hex())            // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println("tx.Value()", tx.Value().String())       // 10000000000000000
		fmt.Println("tx.Gas()", tx.Gas())                    // 105000
		fmt.Println("tx.GasPrice()", tx.GasPrice().Uint64()) // 102000000000
		fmt.Println("tx.Nonce()", tx.Nonce())                // 110644
		fmt.Println("tx.Data()", tx.Data())                  // []
		fmt.Println("tx.To()", tx.To().Hex())                // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e
	}
	fmt.Println()
}
