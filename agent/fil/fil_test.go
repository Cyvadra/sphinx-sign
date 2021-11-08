package fil

import (
	"testing"

	"github.com/cyvadra/filecoin-client/local"
	"github.com/cyvadra/filecoin-client/types"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/stretchr/testify/assert"
)

func TestSignScript(t *testing.T) {
	// (s *types.SignedMessage, err error)
	ki, addr, err := local.WalletNew(types.KTBLS)
	assert.Nil(t, err)
	msg := &types.Message{
		Version:    0,
		To:         *addr,
		From:       *addr,
		Nonce:      0,
		Value:      big.Int{123456789},
		GasLimit:   0,
		GasFeeCap:  big.Int{10000},
		GasPremium: big.Int{10000},
		Method:     0,
		Params:     []byte{},
	}
	s, err := SignScript(ki, msg)
	err = local.WalletVerifyMessage(s)
	assert.Nil(t, err)
	return
}

func TestCreateAccount(t *testing.T) {
	// (ki *types.KeyInfo, addr *address.Address, err error)
	ki, addr, err = local.WalletNew(types.KTBLS)
	assert.Nil(t, err)
	return
}
