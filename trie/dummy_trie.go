// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package trie

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/trie/trienode"
)

type EmptyTrie struct{}

// NewSecure creates a dummy trie
func NewEmptyTrie() *EmptyTrie {
	return &EmptyTrie{}
}

func (t *EmptyTrie) GetStorage(addr common.Address, key []byte) ([]byte, error) {
	return nil, nil
}

func (t *EmptyTrie) GetAccount(address common.Address) (*types.StateAccount, error) {
	return nil, nil
}

func (t *EmptyTrie) UpdateStorage(addr common.Address, key, value []byte) error {
	return nil
}

func (t *EmptyTrie) UpdateAccount(address common.Address, account *types.StateAccount) error {
	return nil
}

func (t *EmptyTrie) UpdateContractCode(address common.Address, codeHash common.Hash, code []byte) error {
	return nil
}

func (t *EmptyTrie) DeleteStorage(addr common.Address, key []byte) error {
	return nil
}

func (t *EmptyTrie) DeleteAccount(address common.Address) error {
	return nil
}

func (t *EmptyTrie) Commit(collectLeaf bool) (common.Hash, *trienode.NodeSet, error) {
	return common.Hash{}, nil, nil
}

func (t *EmptyTrie) NodeIterator(startKey []byte) (NodeIterator, error) {
	return nil, nil
}

func (t *EmptyTrie) Prove(key []byte, proofDb ethdb.KeyValueWriter) error {
	return nil
}

func (t *EmptyTrie) Get(key []byte) []byte {
	return nil
}

func (t *EmptyTrie) TryGet(key []byte) ([]byte, error) {
	return nil, nil
}

func (t *EmptyTrie) TryGetNode(path []byte) ([]byte, int, error) {
	return nil, 0, nil
}
func (t *EmptyTrie) Update(key, value []byte) {}

func (t *EmptyTrie) TryUpdate(key, value []byte) error {
	return nil
}

// Delete removes any existing value for key from the trie.
func (t *EmptyTrie) Delete(key []byte) {
	if err := t.TryDelete(key); err != nil {
		log.Error(fmt.Sprintf("Unhandled trie error: %v", err))
	}
}

func (t *EmptyTrie) TryDelete(key []byte) error {
	return nil
}

func (t *EmptyTrie) GetKey(shaKey []byte) []byte {
	return nil
}

func (t *EmptyTrie) Hash() common.Hash {
	return common.Hash{}
}

// Copy returns a copy of SecureTrie.
func (t *EmptyTrie) Copy() *EmptyTrie {
	cpy := *t
	return &cpy
}

func (t *EmptyTrie) ResetCopy() *EmptyTrie {
	cpy := *t
	return &cpy
}
