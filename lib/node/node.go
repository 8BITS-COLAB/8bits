package node

import "github.com/google/uuid"

type Node struct {
	Stake   int    `json:"stake"`
	Address string `json:"address"`
}

func New(s int) *Node {
	return &Node{
		Stake:   s,
		Address: uuid.NewString(),
	}
}
