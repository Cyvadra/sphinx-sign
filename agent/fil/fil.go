package fil

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/myxtype/filecoin-client"
	"github.com/myxtype/filecoin-client/local"
	"github.com/myxtype/filecoin-client/types"
	"github.com/shopspring/decimal"
)

var Client *filecoin.Client

func main() {
	var err error
	to_addr := "t1gvkap5jmv5k7gwpa42zj43i2oaai5zg74n66xra"
	// 设置主机
	SetHostWithToken("172.16.30.117", "")
	// 设置🔑
	pk := []byte("7b2254797065223a22736563703235366b31222c22507269766174654b6579223a223772503034624643507854562b356a6f6954644b76366d2f61763064335a716c304a757879577856346e493d227d")
	ki := types.KeyInfo{
		Type:       "secp256k1",
		PrivateKey: pk,
	}
	addr, err := local.WalletPrivateToAddress(crypto.SigTypeSecp256k1, pk)
	if err != nil {
		fmt.Println(err)
	}

	// 设置网络类型
	address.CurrentNetwork = address.Mainnet

	fmt.Printf("key: %v\n", hex.EncodeToString(ki.PrivateKey))
	fmt.Printf("address: %v\n", addr.String())

	to, err := address.NewFromString(to_addr)

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
	} else {
		fmt.Println("signed message: ")
		fmt.Println(s)
	}

	println(hex.EncodeToString(s.Signature.Data))

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
