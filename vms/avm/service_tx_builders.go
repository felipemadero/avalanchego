// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"context"
	"fmt"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/avm/txs"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/components/verify"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	"github.com/ava-labs/avalanchego/wallet/chain/x/backends"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary/common"
)

func buildCreateAssetTx(
	backend *Backend,
	name, symbol string,
	denomination byte,
	initialStates map[uint32][]verify.State,
	kc *secp256k1fx.Keychain,
	changeAddr ids.ShortID,
) (*txs.Tx, ids.ShortID, error) {
	pBuilder, pSigner := builders(backend, kc)

	utx, err := pBuilder.NewCreateAssetTx(
		name,
		symbol,
		denomination,
		initialStates,
		options(changeAddr, nil /*memo*/)...,
	)
	if err != nil {
		return nil, ids.ShortEmpty, fmt.Errorf("failed building base tx: %w", err)
	}

	tx, err := backends.SignUnsigned(context.Background(), pSigner, utx)
	if err != nil {
		return nil, ids.ShortEmpty, err
	}

	return tx, changeAddr, nil
}

func buildBaseTx(
	backend *Backend,
	outs []*avax.TransferableOutput,
	memo []byte,
	kc *secp256k1fx.Keychain,
	changeAddr ids.ShortID,
) (*txs.Tx, ids.ShortID, error) {
	pBuilder, pSigner := builders(backend, kc)

	utx, err := pBuilder.NewBaseTx(
		outs,
		options(changeAddr, memo)...,
	)
	if err != nil {
		return nil, ids.ShortEmpty, fmt.Errorf("failed building base tx: %w", err)
	}

	tx, err := backends.SignUnsigned(context.Background(), pSigner, utx)
	if err != nil {
		return nil, ids.ShortEmpty, err
	}

	return tx, changeAddr, nil
}

func buildOperation(
	backend *Backend,
	ops []*txs.Operation,
	kc *secp256k1fx.Keychain,
	changeAddr ids.ShortID,
) (*txs.Tx, error) {
	pBuilder, pSigner := builders(backend, kc)

	utx, err := pBuilder.NewOperationTx(
		ops,
		options(changeAddr, nil /*memo*/)...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed building import tx: %w", err)
	}

	return backends.SignUnsigned(context.Background(), pSigner, utx)
}

func buildImportTx(
	backend *Backend,
	sourceChain ids.ID,
	to ids.ShortID,
	kc *secp256k1fx.Keychain,
) (*txs.Tx, error) {
	pBuilder, pSigner := builders(backend, kc)

	outOwner := &secp256k1fx.OutputOwners{
		Locktime:  0,
		Threshold: 1,
		Addrs:     []ids.ShortID{to},
	}

	utx, err := pBuilder.NewImportTx(
		sourceChain,
		outOwner,
	)
	if err != nil {
		return nil, fmt.Errorf("failed building import tx: %w", err)
	}

	return backends.SignUnsigned(context.Background(), pSigner, utx)
}

func buildExportTx(
	backend *Backend,
	destinationChain ids.ID,
	to ids.ShortID,
	exportedAssetID ids.ID,
	exportedAmt uint64,
	kc *secp256k1fx.Keychain,
	changeAddr ids.ShortID,
) (*txs.Tx, ids.ShortID, error) {
	pBuilder, pSigner := builders(backend, kc)

	outputs := []*avax.TransferableOutput{{
		Asset: avax.Asset{ID: exportedAssetID},
		Out: &secp256k1fx.TransferOutput{
			Amt: exportedAmt,
			OutputOwners: secp256k1fx.OutputOwners{
				Locktime:  0,
				Threshold: 1,
				Addrs:     []ids.ShortID{to},
			},
		},
	}}

	utx, err := pBuilder.NewExportTx(
		destinationChain,
		outputs,
		options(changeAddr, nil /*memo*/)...,
	)
	if err != nil {
		return nil, ids.ShortEmpty, fmt.Errorf("failed building export tx: %w", err)
	}

	tx, err := backends.SignUnsigned(context.Background(), pSigner, utx)
	if err != nil {
		return nil, ids.ShortEmpty, err
	}
	return tx, changeAddr, nil
}

func builders(backend *Backend, kc *secp256k1fx.Keychain) (backends.Builder, backends.Signer) {
	var (
		addrs   = kc.Addresses()
		builder = backends.NewBuilder(addrs, backend)
		signer  = backends.NewSigner(kc, backend)
	)
	backend.ResetAddresses(addrs)

	return builder, signer
}

func options(changeAddr ids.ShortID, memo []byte) []common.Option {
	return common.UnionOptions(
		[]common.Option{common.WithChangeOwner(&secp256k1fx.OutputOwners{
			Threshold: 1,
			Addrs:     []ids.ShortID{changeAddr},
		})},
		[]common.Option{common.WithMemo(memo)},
	)
}
