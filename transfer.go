package main

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"tkido.com/goeth/config"
)

func transfer(senderAddress, toAddress string, value *big.Int) {
	key, err := keystore.DecryptKey([]byte(config.Key), config.Secret)
	if err != nil {
		log.Fatalf("cannot decrypt key: %s", err)
	}
	var (
		to       = common.HexToAddress(toAddress)
		sender   = common.HexToAddress(senderAddress)
		gasLimit = uint64(21000)
	)
	// Retrieve the pending nonce
	nonce, err := client.PendingNonceAt(context.Background(), sender)
	if err != nil {
		log.Fatal(err)
	}
	// Get suggested gas price
	tipCap, _ := client.SuggestGasTipCap(context.Background())
	feeCap, _ := client.SuggestGasPrice(context.Background())
	// Create a new transaction
	tx := types.NewTx(
		&types.DynamicFeeTx{
			ChainID:   big.NewInt(config.ChainID),
			Nonce:     nonce,
			GasTipCap: tipCap,
			GasFeeCap: feeCap,
			Gas:       gasLimit,
			To:        &to,
			Value:     value,
			Data:      nil,
		})
	// Sign the transaction using our keys

	signedTx, _ := types.SignTx(tx, types.NewLondonSigner(big.NewInt(config.ChainID)), key.PrivateKey)
	// Send the transaction to our node
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("cannot send transaction: %s", err)
	}
}
