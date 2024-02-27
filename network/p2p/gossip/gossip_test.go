// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gossip

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"
	"google.golang.org/protobuf/proto"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/proto/pb/sdk"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/utils/units"
)

func TestGossiperShutdown(*testing.T) {
	gossiper := NewPullGossiper[*testTx](
		logging.NoLog{},
		nil,
		nil,
		nil,
		Metrics{},
		0,
	)
	ctx, cancel := context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		Every(ctx, logging.NoLog{}, gossiper, time.Second)
		wg.Done()
	}()

	cancel()
	wg.Wait()
}

func TestGossiperGossip(t *testing.T) {
	tests := []struct {
		name                   string
		targetResponseSize     int
		requester              []*testTx // what we have
		responder              []*testTx // what the peer we're requesting gossip from has
		expectedPossibleValues []*testTx // possible values we can have
		expectedLen            int
	}{
		{
			name: "no gossip - no one knows anything",
		},
		{
			name:                   "no gossip - requester knows more than responder",
			targetResponseSize:     1024,
			requester:              []*testTx{{id: ids.ID{0}}},
			expectedPossibleValues: []*testTx{{id: ids.ID{0}}},
			expectedLen:            1,
		},
		{
			name:                   "no gossip - requester knows everything responder knows",
			targetResponseSize:     1024,
			requester:              []*testTx{{id: ids.ID{0}}},
			responder:              []*testTx{{id: ids.ID{0}}},
			expectedPossibleValues: []*testTx{{id: ids.ID{0}}},
			expectedLen:            1,
		},
		{
			name:                   "gossip - requester knows nothing",
			targetResponseSize:     1024,
			responder:              []*testTx{{id: ids.ID{0}}},
			expectedPossibleValues: []*testTx{{id: ids.ID{0}}},
			expectedLen:            1,
		},
		{
			name:                   "gossip - requester knows less than responder",
			targetResponseSize:     1024,
			requester:              []*testTx{{id: ids.ID{0}}},
			responder:              []*testTx{{id: ids.ID{0}}, {id: ids.ID{1}}},
			expectedPossibleValues: []*testTx{{id: ids.ID{0}}, {id: ids.ID{1}}},
			expectedLen:            2,
		},
		{
			name:                   "gossip - target response size exceeded",
			targetResponseSize:     32,
			responder:              []*testTx{{id: ids.ID{0}}, {id: ids.ID{1}}, {id: ids.ID{2}}},
			expectedPossibleValues: []*testTx{{id: ids.ID{0}}, {id: ids.ID{1}}, {id: ids.ID{2}}},
			expectedLen:            2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()

			responseSender := &common.FakeSender{
				SentAppResponse: make(chan []byte, 1),
			}
			responseNetwork, err := p2p.NewNetwork(logging.NoLog{}, responseSender, prometheus.NewRegistry(), "")
			require.NoError(err)

			responseBloom, err := NewBloomFilter(prometheus.NewRegistry(), "", 1000, 0.01, 0.05)
			require.NoError(err)
			responseSet := &testSet{
				txs:   make(map[ids.ID]*testTx),
				bloom: responseBloom,
			}
			for _, item := range tt.responder {
				require.NoError(responseSet.Add(item))
			}

			metrics, err := NewMetrics(prometheus.NewRegistry(), "")
			require.NoError(err)
			marshaller := testMarshaller{}
			handler := NewHandler[*testTx](
				logging.NoLog{},
				marshaller,
				responseSet,
				metrics,
				tt.targetResponseSize,
			)
			require.NoError(err)
			require.NoError(responseNetwork.AddHandler(0x0, handler))

			requestSender := &common.FakeSender{
				SentAppRequest: make(chan []byte, 1),
			}

			requestNetwork, err := p2p.NewNetwork(logging.NoLog{}, requestSender, prometheus.NewRegistry(), "")
			require.NoError(err)
			require.NoError(requestNetwork.Connected(context.Background(), ids.EmptyNodeID, nil))

			bloom, err := NewBloomFilter(prometheus.NewRegistry(), "", 1000, 0.01, 0.05)
			require.NoError(err)
			requestSet := &testSet{
				txs:   make(map[ids.ID]*testTx),
				bloom: bloom,
			}
			for _, item := range tt.requester {
				require.NoError(requestSet.Add(item))
			}

			requestClient := requestNetwork.NewClient(0x0)

			require.NoError(err)
			gossiper := NewPullGossiper[*testTx](
				logging.NoLog{},
				marshaller,
				requestSet,
				requestClient,
				metrics,
				1,
			)
			require.NoError(err)
			received := set.Set[*testTx]{}
			requestSet.onAdd = func(tx *testTx) {
				received.Add(tx)
			}

			require.NoError(gossiper.Gossip(ctx))
			require.NoError(responseNetwork.AppRequest(ctx, ids.EmptyNodeID, 1, time.Time{}, <-requestSender.SentAppRequest))
			require.NoError(requestNetwork.AppResponse(ctx, ids.EmptyNodeID, 1, <-responseSender.SentAppResponse))

			require.Len(requestSet.txs, tt.expectedLen)
			require.Subset(tt.expectedPossibleValues, maps.Values(requestSet.txs))

			// we should not receive anything that we already had before we
			// requested the gossip
			for _, tx := range tt.requester {
				require.NotContains(received, tx)
			}
		})
	}
}

func TestEvery(*testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	calls := 0
	gossiper := &TestGossiper{
		GossipF: func(context.Context) error {
			if calls >= 10 {
				cancel()
				return nil
			}

			calls++
			return nil
		},
	}

	go Every(ctx, logging.NoLog{}, gossiper, time.Millisecond)
	<-ctx.Done()
}

func TestValidatorGossiper(t *testing.T) {
	require := require.New(t)

	nodeID := ids.GenerateTestNodeID()

	validators := testValidatorSet{
		validators: set.Of(nodeID),
	}

	calls := 0
	gossiper := ValidatorGossiper{
		Gossiper: &TestGossiper{
			GossipF: func(context.Context) error {
				calls++
				return nil
			},
		},
		NodeID:     nodeID,
		Validators: validators,
	}

	// we are a validator, so we should request gossip
	require.NoError(gossiper.Gossip(context.Background()))
	require.Equal(1, calls)

	// we are not a validator, so we should not request gossip
	validators.validators = set.Set[ids.NodeID]{}
	require.NoError(gossiper.Gossip(context.Background()))
	require.Equal(2, calls)
}

// Tests that the outgoing gossip is equivalent to what was accumulated
func TestPushGossiper(t *testing.T) {
	type cycle struct {
		toAdd    []*testTx
		expected []*testTx
	}
	tests := []struct {
		name   string
		cycles []cycle
	}{
		{
			name: "single cycle",
			cycles: []cycle{
				{
					toAdd: []*testTx{
						{
							id: ids.ID{0},
						},
						{
							id: ids.ID{1},
						},
						{
							id: ids.ID{2},
						},
					},
					expected: []*testTx{
						{
							id: ids.ID{0},
						},
						{
							id: ids.ID{1},
						},
						{
							id: ids.ID{2},
						},
					},
				},
			},
		},
		{
			name: "multiple cycles",
			cycles: []cycle{
				{
					toAdd: []*testTx{
						{
							id: ids.ID{0},
						},
					},
					expected: []*testTx{
						{
							id: ids.ID{0},
						},
					},
				},
				{
					toAdd: []*testTx{
						{
							id: ids.ID{1},
						},
					},
					expected: []*testTx{
						{
							id: ids.ID{1},
						},
						{
							id: ids.ID{0},
						},
					},
				},
				{
					toAdd: []*testTx{
						{
							id: ids.ID{2},
						},
					},
					expected: []*testTx{
						{
							id: ids.ID{2},
						},
						{
							id: ids.ID{1},
						},
						{
							id: ids.ID{0},
						},
					},
				},
			},
		},
	}

	const regossipTime = time.Nanosecond
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()

			sender := &common.FakeSender{
				SentAppGossip: make(chan []byte, 1),
			}
			network, err := p2p.NewNetwork(
				logging.NoLog{},
				sender,
				prometheus.NewRegistry(),
				"",
			)
			require.NoError(err)
			client := network.NewClient(0)
			metrics, err := NewMetrics(prometheus.NewRegistry(), "")
			require.NoError(err)
			marshaller := testMarshaller{}
			gossiper := NewPushGossiper[*testTx](
				marshaller,
				FullSet[*testTx]{},
				client,
				metrics,
				0, // the discarded cache size doesn't matter for this test
				units.MiB,
				regossipTime,
			)

			for _, cycle := range tt.cycles {
				gossiper.Add(cycle.toAdd...)
				require.NoError(gossiper.Gossip(ctx))

				want := &sdk.PushGossip{
					Gossip: make([][]byte, 0, len(cycle.expected)),
				}

				for _, gossipable := range cycle.expected {
					bytes, err := marshaller.MarshalGossip(gossipable)
					require.NoError(err)

					want.Gossip = append(want.Gossip, bytes)
				}

				// remove the handler prefix
				sentMsg := <-sender.SentAppGossip
				got := &sdk.PushGossip{}
				require.NoError(proto.Unmarshal(sentMsg[1:], got))

				require.Equal(want.Gossip, got.Gossip)

				// Ensure that subsequent calls to `time.Now()` are sufficient
				// for regossip.
				time.Sleep(regossipTime)
			}
		})
	}
}

type testValidatorSet struct {
	validators set.Set[ids.NodeID]
}

func (t testValidatorSet) Has(_ context.Context, nodeID ids.NodeID) bool {
	return t.validators.Contains(nodeID)
}
