package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Timestamp        string `json:"timestamp"`
	PrevHash         string `json:"prev_hash"`
	Hash             string `json:"hash"`
	ValidatorAddress string `json:"validator_address"`
}

func (b *Block) Sum() string {
	i := fmt.Sprintf("%s:%s:%s", b.Timestamp, b.PrevHash, b.ValidatorAddress)

	h := sha256.New()
	h.Write([]byte(i))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
