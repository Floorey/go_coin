package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Coin struct {
	ID    string
	Owner string
	Value int
}

func NewCoin(owner string, value int) *Coin {
	coin := &Coin{
		Owner: owner,
		Value: value,
	}
	coin.ID = coin.generateID()
	return coin
}
func (c *Coin) generateID() string {
	data := fmt.Sprintf("%s:%d", c.Owner, c.Value)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
