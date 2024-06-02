package main

import "errors"

type Transaction struct {
	Sender   string
	Receiver string
	Value    int
}
type Block struct {
	Transactions []Transaction
}
type Blockchain struct {
	Blocks []Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]Block{}}
}
func (bc *Blockchain) AddBlock(transactions []Transaction) {
	block := Block{transactions}
	bc.Blocks = append(bc.Blocks, block)
}

func (bc *Blockchain) ValidateTransaction(tx Transaction) error {
	if tx.Value <= 0 {
		return errors.New("invalid transaction!")
	}
	return nil
}
func (bc *Blockchain) AddTransaction(tx Transaction) error {
	if err := bc.ValidateTransaction(tx); err != nil {
		return err
	}
	bc.AddBlock([]Transaction{tx})
	return nil
}
