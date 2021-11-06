package fil

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/myxtype/filecoin-client"
	"github.com/myxtype/filecoin-client/local"
	"github.com/myxtype/filecoin-client/types"
	"github.com/shopspring/decimal"
)

var Client *filecoin.Client

func main() {

	var err error

	SetHostWithToken("172.16.30.117", "")

	// 设置网络类型
	address.CurrentNetwork = address.Mainnet

	// 生产新的地址
	// 新地址有转入fil才激活，不然没法用
	ki, addr, err := local.WalletNew(types.KTSecp256k1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("key: %v\n", hex.EncodeToString(ki.PrivateKey))
	fmt.Printf("address: %v\n", addr.String())

	to, err := address.NewFromString("t1gvkap5jmv5k7gwpa42zj43i2oaai5zg74n66xra")

	// 转移0.001FIL到f1yfi4yslez2hz3ori5grvv3xdo3xkibc4v6xjusy
	msg := &types.Message{
		Version:    0,
		To:         to,
		From:       *addr,
		Nonce:      0,
		Value:      filecoin.FromFil(decimal.NewFromFloat(0.001)),
		GasLimit:   0,
		GasFeeCap:  abi.NewTokenAmount(10000),
		GasPremium: abi.NewTokenAmount(10000),
		Method:     0,
		Params:     nil,
	}

	// 最大手续费0.0001 FIL
	//maxFee := filecoin.FromFil(decimal.NewFromFloat(0.0001))

	// 估算GasLimit
	//msg, err = client.GasEstimateMessageGas(context.Background(), msg, &types.MessageSendSpec{MaxFee: maxFee}, nil)
	//if err != nil {
	//	panic(err)
	//}

	// 离线签名
	s, err := local.WalletSignMessage(types.KTSecp256k1, ki.PrivateKey, msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	println(hex.EncodeToString(s.Signature.Data))
	// 47bcbb167fd9040bd02dba02789bc7bc0463c290db1be9b07065c12a64fb84dc546bef7aedfba789d0d7ce2c4532f8fa0d2dd998985ad3ec1a8b064c26e4625a01

	// 验证签名
	if err := local.WalletVerifyMessage(s); err != nil {
		fmt.Println(err)
		return
	}

	mid, err := Client.MpoolPush(context.Background(), s)
	if err != nil {
		fmt.Println(err)
		return
	}

	println(mid.String())
}

func lotusURL(host string) string {
	return fmt.Sprintf("http://%v:1234/rpc/v0", host)
}

func SetHostWithToken(str, token string) {
	Client = filecoin.NewClient(lotusURL(str), token)
}
