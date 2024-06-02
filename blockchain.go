package main

import (
	"errors"
	"fmt"
)

// Transaction repräsentiert eine Transaktion mit Coins.
type Transaction struct {
	Sender   string
	Receiver string
	Value    int
}

// Block repräsentiert einen Block in der Blockchain.
type Block struct {
	Transactions []Transaction
}

// Blockchain repräsentiert eine einfache Blockchain.
type Blockchain struct {
	Blocks []Block
}

// NewBlockchain erstellt eine neue Blockchain und lädt die Blöcke aus der Datenbank.
func NewBlockchain() *Blockchain {
	blocks, err := LoadBlocks()
	if err != nil {
		fmt.Println("Fehler beim Laden der Blöcke:", err)
		return &Blockchain{[]Block{}}
	}
	return &Blockchain{blocks}
}

// AddBlock fügt einen neuen Block zur Blockchain hinzu und speichert ihn in der Datenbank.
func (bc *Blockchain) AddBlock(transactions []Transaction) {
	block := Block{transactions}
	bc.Blocks = append(bc.Blocks, block)
	if err := SaveBlock(block); err != nil {
		fmt.Println("Fehler beim Speichern des Blocks:", err)
	}
}

// ValidateTransaction validiert eine Transaktion.
func (bc *Blockchain) ValidateTransaction(tx Transaction) error {
	if tx.Value <= 0 {
		return errors.New("ungültiger Transaktionswert")
	}
	return nil
}

// AddTransaction fügt eine Transaktion zur Blockchain hinzu.
func (bc *Blockchain) AddTransaction(tx Transaction) error {
	if err := bc.ValidateTransaction(tx); err != nil {
		return err
	}
	bc.AddBlock([]Transaction{tx})
	return nil
}
