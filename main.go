package main

import "fmt"

func main() {
	bc := NewBlockchain()

	coin := NewCoin("Bob", 100)
	fmt.Printf("New coin created: %+v\n", coin)

	tx := Transaction{
		Sender:   "Bob",
		Receiver: "Luke",
		Value:    50,
	}
	if err := bc.AddTransaction(tx); err != nil {
		fmt.Println("Error with transaction:", err)

	} else {
		fmt.Println("Transaction successfully added!")
	}
	fmt.Printf("Momentum Blockchain: %fv\n", bc.Blocks)
}
