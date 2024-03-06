package blockchain

import (
    "fmt"
    "time"
)

// Define the structure of a block
type Block struct {
    Index     int
    Timestamp time.Time
    Data      string
    PrevHash  string
    Hash      string
}

// Blockchain is a slice of blocks
type Blockchain []Block

// DisplayAllBlocks prints all blocks in the blockchain
func DisplayAllBlocks(chain Blockchain) {
    for _, block := range chain {
        fmt.Printf("Index: %d\n", block.Index)
        fmt.Printf("Timestamp: %s\n", block.Timestamp.Format(time.RFC3339))
        fmt.Printf("Data: %s\n", block.Data)
        fmt.Printf("PrevHash: %s\n", block.PrevHash)
        fmt.Printf("Hash: %s\n\n", block.Hash)
    }
}

// NewBlock creates a new block and adds it to the blockchain
func NewBlock(prevBlock Block, data string) Block {
    newIndex := prevBlock.Index + 1
    newTimestamp := time.Now()
    newHash := "calculate_hash_function" // Implement hash calculation
    return Block{newIndex, newTimestamp, data, prevBlock.Hash, newHash}
}

// ModifyBlock updates the data of a specific block in the blockchain
func ModifyBlock(chain Blockchain, index int, newData string) {
    if index >= 0 && index < len(chain) {
        chain[index].Data = newData
        // Recalculate hash if necessary
        chain[index].Hash = "recalculate_hash_function"
    } else {
        fmt.Println("Invalid block index")
    }
}
