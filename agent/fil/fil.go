package fil

import (
	"github.com/cyvadra/filecoin-client/local"
	"github.com/cyvadra/filecoin-client/types"
	"github.com/filecoin-project/go-address"
)

func SignScript(ki *types.KeyInfo, msg *types.Message) (s *types.SignedMessage, err error) {
	// 离线签名
	s, err = local.WalletSignMessage(types.KTBLS, ki.PrivateKey, msg)
	if err != nil {
		return
	}
	err = local.WalletVerifyMessage(s)
	return
}

func CreateAccount() (ki *types.KeyInfo, addr *address.Address, err error) {
	ki, addr, err = local.WalletNew(types.KTBLS)
	return
}
