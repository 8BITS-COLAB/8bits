package main

import (
	"log"
	"time"

	"github.com/ElioenaiFerrari/8bits/lib/block"
	"github.com/ElioenaiFerrari/8bits/lib/network"
	"github.com/ElioenaiFerrari/8bits/lib/node"
)

func main() {

	g := &block.Block{
		Timestamp:        time.Now().String(),
		PrevHash:         "",
		ValidatorAddress: "",
	}

	g.Hash = g.Sum()

	n := &network.Network{
		Chain: []*block.Block{
			g,
		},
	}

	n.ChainHead = n.Chain[0]
	n.Validators = n.AddNode(node.New(60))
	n.Validators = n.AddNode(node.New(40))

	for i := 0; i < 5; i++ {
		w, err := n.GetWinner()

		if err != nil {
			log.Fatal(err)
		}

		w.Stake += 10
		n.Chain, n.ChainHead, err = n.GenerateNewBlock(w)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Round ", i)
		log.Println("\tAddress:", n.Validators[0].Address, "-Stake:", n.Validators[0].Stake)
		log.Println("\tAddress:", n.Validators[1].Address, "-Stake:", n.Validators[1].Stake)
	}

	log.Printf("%+v", n.Chain)
}
