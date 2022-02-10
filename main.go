package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	blocks []*Block
}

func (c *BlockChain) AddBlock(data string) {
	prevBlock := c.blocks[len(c.blocks)-1]
	b := CreateBlock(data, prevBlock.Hash)
	c.blocks = append(c.blocks, b)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{blocks: []*Block{Genesis()}}
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	b := &Block{
		Data:     []byte(data),
		PrevHash: prevHash,
	}
	b.DeriveHash()
	return b
}
func main() {
	chain := InitBlockChain()
	chain.AddBlock("First book")
	chain.AddBlock("Second book")
	chain.AddBlock("Third book")

	for _, b := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", b.PrevHash)
		fmt.Printf("Data in Block: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Hash)
	}
}
