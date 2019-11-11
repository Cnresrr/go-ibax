/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package daemons

import (
	"encoding/hex"

	"github.com/IBAX-io/go-ibax/packages/conf"
	"github.com/IBAX-io/go-ibax/packages/consts"
	"github.com/IBAX-io/go-ibax/packages/model"
	"github.com/IBAX-io/go-ibax/packages/smart"
	"github.com/IBAX-io/go-ibax/packages/utils/tx"

	log "github.com/sirupsen/logrus"
)

const (
	callDelayedContract = "CallDelayedContract"
	firstEcosystemID    = 1
)

// DelayedTx represents struct which works with delayed contracts
type DelayedTx struct {
	logger     *log.Entry
	privateKey string
	publicKey  string
	time       int64
}

// RunForDelayBlockID creates the transactions that need to be run for blockID
func (dtx *DelayedTx) RunForDelayBlockID(blockID int64) ([]*model.Transaction, error) {
	}

	return txList, nil
}

func (dtx *DelayedTx) createDelayTx(keyID, highRate int64, params map[string]interface{}) (*model.Transaction, error) {
	vm := smart.GetVM()
	contract := smart.VMGetContract(vm, callDelayedContract, uint32(firstEcosystemID))
	info := contract.Info()

	smartTx := tx.SmartContract{
		Header: tx.Header{
			ID:          int(info.ID),
			Time:        dtx.time,
			EcosystemID: firstEcosystemID,
			KeyID:       keyID,
			NetworkID:   conf.Config.NetworkID,
		},
		SignedBy: smart.PubToID(dtx.publicKey),
		Params:   params,
	}

	privateKey, err := hex.DecodeString(dtx.privateKey)
	if err != nil {
		return nil, err
	}

	txData, txHash, err := tx.NewInternalTransaction(smartTx, privateKey)
	if err != nil {
		return nil, err
	}
	return tx.CreateDelayTransactionHighRate(txData, txHash, keyID, highRate), nil
}