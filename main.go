package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Block represents each 'item' in the blockchain
type Block struct {
	Index    int
	Data     string
	PrevHash string
	Hash     string
}

// calculateHash creates a SHA-256 hash of a block
func calculateHash(block Block) string {
	// Correct way to format all parts as a string
	record := fmt.Sprintf("%d%s%s", block.Index, block.Data, block.PrevHash)
	hash := sha256.New()
	hash.Write([]byte(record))
	return hex.EncodeToString(hash.Sum(nil))
}

// createBlock creates a new block using the previous block’s hash
func createBlock(prevBlock Block, data string) Block {
	newBlock := Block{
		Index:    prevBlock.Index + 1,
		Data:     data,
		PrevHash: prevBlock.Hash,
	}
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

// printBlock prints the details of a block
func printBlock(block Block) {
	fmt.Printf("Block #%d\n", block.Index)
	fmt.Printf("Data: %s\n", block.Data)
	fmt.Printf("Previous Hash: %s\n", block.PrevHash)
	fmt.Printf("Hash: %s\n\n", block.Hash)
}

func main() {
	// Create genesis block
	genesisBlock := Block{Index: 0, Data: "Genesis Block", PrevHash: ""}
	genesisBlock.Hash = calculateHash(genesisBlock)

	// Create additional blocks
	secondBlock := createBlock(genesisBlock, "Second Block")
	thirdBlock := createBlock(secondBlock, "Third Block")

	// Print the blockchain
	printBlock(genesisBlock)
	printBlock(secondBlock)
	printBlock(thirdBlock)
}
