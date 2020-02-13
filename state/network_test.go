// Copyright 2019 ChainSafe Systems (ON) Corp.
// This file is part of gossamer.
//
// The gossamer library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The gossamer library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the gossamer library. If not, see <http://www.gnu.org/licenses/>.

package state

import (
	"math/big"
	"testing"

	"github.com/ChainSafe/gossamer/common"
	"github.com/ChainSafe/gossamer/core/types"
	"github.com/ChainSafe/gossamer/trie"
	"github.com/stretchr/testify/require"
)

var testHealth = &common.Health{}
var testNetworkState = &common.NetworkState{}
var testPeers = &[]common.PeerInfo{}

// test state.Network
func TestNetworkState(t *testing.T) {
	state := newTestService(t)

	header := &types.Header{
		Number:    big.NewInt(0),
		StateRoot: trie.EmptyHash,
	}

	err := state.Initialize(header, trie.NewEmptyTrie(nil))
	if err != nil {
		t.Fatal(err)
	}

	err = state.Start()
	if err != nil {
		t.Fatal(err)
	}

	err = state.Network.SetHealth(testHealth)
	require.Nil(t, err)

	health, err := state.Network.GetHealth()
	require.Nil(t, err)

	require.Equal(t, health, testHealth)

	err = state.Network.SetNetworkState(testNetworkState)
	require.Nil(t, err)

	networkState, err := state.Network.GetNetworkState()
	require.Nil(t, err)

	require.Equal(t, networkState, testNetworkState)

	err = state.Network.SetPeers(testPeers)
	require.Nil(t, err)

	peers, err := state.Network.GetPeers()
	require.Nil(t, err)

	require.Equal(t, peers, testPeers)
}