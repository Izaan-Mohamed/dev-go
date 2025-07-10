package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)


type Block struct {
	Index    int
	Data     string
	PrevHash string
	Hash     string
}


func calculateHash(block Block) string {
	
	record := fmt.Sprintf("%d%s%s", block.Index, block.Data, block.PrevHash)
	hash := sha256.New()
	hash.Write([]byte(record))
	return hex.EncodeToString(hash.Sum(nil))
}


func createBlock(prevBlock Block, data string) Block {
	newBlock := Block{
		Index:    prevBlock.Index + 1,
		Data:     data,
		PrevHash: prevBlock.Hash,
	}
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}


func printBlock(block Block) {
	fmt.Printf("Block #%d\n", block.Index)
	fmt.Printf("Data: %s\n", block.Data)
	fmt.Printf("Previous Hash: %s\n", block.PrevHash)
	fmt.Printf("Hash: %s\n\n", block.Hash)
}

func main() {

	genesisBlock := Block{Index: 0, Data: "Genesis Block", PrevHash: ""}
	genesisBlock.Hash = calculateHash(genesisBlock)


	secondBlock := createBlock(genesisBlock, "Second Block")
	thirdBlock := createBlock(secondBlock, "Third Block")

	
	printBlock(genesisBlock)
	printBlock(secondBlock)
	printBlock(thirdBlock)
}
