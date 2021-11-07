package main

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/cyvadra/filecoin-client"
	"github.com/cyvadra/filecoin-client/local"
	"github.com/cyvadra/filecoin-client/types"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/shopspring/decimal"
)

var Client *filecoin.Client

func main() {
	var err error
	toAddr := "t1gvkap5jmv5k7gwpa42zj43i2oaai5zg74n66xra"
	pkStr := "c3pS5JcZEM1C5Yukor63mQ8DvADh1qQN/GrUsRA20XE="
	var pk []byte
	_, err = base64.StdEncoding.Decode([]byte(pkStr), pk)
	if err != nil {
		fmt.Println("pk解码失败")
	}
	// 设置主机
	address.CurrentNetwork = address.Mainnet
	SetHostWithToken("172.16.30.117", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.ppK_nggwygh6kCPDlktdBtkGaqQXxoXM99iNx3-tZ9E")
	// 设置key
	ki := types.KeyInfo{
		Type:       types.KTBLS,
		PrivateKey: pk,
	}

	// 测试
	fmt.Println("===============")
	fmt.Println(crypto.SigTypeSecp256k1)
	fmt.Println(crypto.SigTypeBLS)
	// 由key生成并确认地址
	fromAddr, err := local.WalletPrivateToAddress(crypto.SigTypeBLS, pk)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Debug output
	fmt.Println("key: ", hex.EncodeToString(ki.PrivateKey))
	fmt.Printf("address: %v\n", fromAddr.String())
	fmt.Println("expected address: t3vciz4skae3yyfnsm54rjri3ubhbk4vklxmdvykwosnqhikehocf26rfiegekwa743ypem4oiv3uq7wuxwcca")
	to, err := address.NewFromString(toAddr)
	// 转移0.001FIL到目标地址
	msg := &types.Message{
		Version:    0,
		To:         to,
		From:       *fromAddr,
		Nonce:      0,
		Value:      filecoin.FromFil(decimal.NewFromFloat(0.001)),
		GasLimit:   0,
		GasFeeCap:  abi.NewTokenAmount(10000),
		GasPremium: abi.NewTokenAmount(10000),
		Method:     0,
		Params:     nil,
	}

	// 最大手续费0.0001 FIL
	// maxFee := filecoin.FromFil(decimal.NewFromFloat(0.0001))

	// 估算GasLimit
	//msg, err = client.GasEstimateMessageGas(context.Background(), msg, &types.MessageSendSpec{MaxFee: maxFee}, nil)
	//if err != nil {
	//	panic(err)
	//}

	// 离线签名
	s, err := local.WalletSignMessage(types.KTSecp256k1, ki.PrivateKey, msg)
	if err != nil {
		fmt.Println("离线签名失败")
		fmt.Println(s)
		fmt.Println(err)
		return
	} else {
		fmt.Println("signed message: ")
		fmt.Println(s)
	}

	println(hex.EncodeToString(s.Signature.Data))

	// 验证签名
	if err := local.WalletVerifyMessage(s); err != nil {
		fmt.Println("验证签名失败")
		fmt.Println(err)
		return
	}

	mid, err := Client.MpoolPush(context.Background(), s)
	if err != nil {
		fmt.Println("消息广播失败")
		fmt.Println(err)
		return
	}
	println("消息发送成功，message id:")
	println(mid.String())
}

func lotusURL(host string) string {
	return fmt.Sprintf("http://%v:1234/rpc/v0", host)
}

func SetHostWithToken(str, token string) {
	Client = filecoin.NewClient(lotusURL(str), token)
}
