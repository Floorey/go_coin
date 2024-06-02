package main

import (
	"fmt"
)

func main() {
	InitDB("blockchain.db")

	bc := NewBlockchain()

	coin := NewCoin("Alice", 100)
	fmt.Printf("Neuer Coin erstellt: %+v\n", coin)

	tx := Transaction{
		Sender:   "Alice",
		Receiver: "Bob",
		Value:    50,
	}

	if err := bc.AddTransaction(tx); err != nil {
		fmt.Println("Fehler bei der Transaktion:", err)
	} else {
		fmt.Println("Transaktion erfolgreich hinzugefügt!")
	}

	fmt.Printf("Aktuelle Blockchain: %+v\n", bc.Blocks)
}
