package network

import (
	"errors"
	"math/rand"
	"time"

	"github.com/ElioenaiFerrari/8bits/lib/block"
	"github.com/ElioenaiFerrari/8bits/lib/node"
)

type Network struct {
	Chain      []*block.Block
	ChainHead  *block.Block
	Validators []*node.Node
}

func (n Network) GetWinner() (*node.Node, error) {
	var wPool []*node.Node
	tStake := 0

	for _, nd := range n.Validators {
		if nd.Stake > 0 {
			wPool = append(wPool, nd)
			tStake += nd.Stake
		}
	}

	if wPool == nil {
		return nil, errors.New("there are no nodes with stake in the network")
	}

	wNumber := rand.Intn(tStake)
	tmp := 0

	for _, nd := range n.Validators {
		tmp += nd.Stake
		if wNumber < tmp {
			return nd, nil
		}
	}

	return nil, errors.New("a winner should have been picked but wasn't")
}

func (n Network) AddNode(nd *node.Node) []*node.Node {
	return append(n.Validators, nd)
}

func (n Network) ValidateBlockCandidate(b *block.Block) error {
	if n.ChainHead.Hash != b.PrevHash {
		return errors.New("blockchain HEAD hash is not equal to new block previous hash")
	}

	if n.ChainHead.Timestamp >= b.Timestamp {
		return errors.New("blockchain HEAD timestamp is greater than or equal to new block timestamp")
	}

	if n.ChainHead.Sum() != b.Hash {
		return errors.New("new block hash of blockchain HEAD does not equal new block hash")
	}

	return nil
}

func (n Network) ValidateChain() error {
	if len(n.Chain) <= 1 {
		return nil
	}

	cbi := len(n.Chain) - 1
	pbi := len(n.Chain) - 2

	for pbi >= 0 {
		cb := n.Chain[cbi]
		pb := n.Chain[pbi]

		if cb.PrevHash != pb.Hash {
			return errors.New("blockchain has inconsistent hashes")
		}

		if cb.Timestamp <= pb.Timestamp {
			return errors.New("blockchain has inconsistent timestamps")
		}

		if pb.Sum() != cb.Hash {
			return errors.New("blockchain has inconsistent hash generation")
		}

		cbi--
		pbi--
	}

	return nil
}

func (n Network) GenerateNewBlock(v *node.Node) ([]*block.Block, *block.Block, error) {
	if err := n.ValidateChain(); err != nil {
		v.Stake -= 10
		return n.Chain, n.ChainHead, err
	}

	t := time.Now().String()

	b := &block.Block{
		Timestamp:        t,
		PrevHash:         n.ChainHead.Hash,
		Hash:             n.ChainHead.Sum(),
		ValidatorAddress: v.Address,
	}

	if err := n.ValidateBlockCandidate(b); err != nil {
		v.Stake -= 10
		return n.Chain, n.ChainHead, err
	}

	n.Chain = append(n.Chain, b)

	return n.Chain, b, nil
}
