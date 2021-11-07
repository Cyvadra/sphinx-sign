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

type SignInfoFIL struct {
	Nonce      int
	Version    int
	GasLimit   int // filecoin.FromFil(decimal.NewFromFloat(0.0001))
	GasFeeCap  int
	GasPremium int
	Method     int
	Value      int
}

func OfflineSign(ki *types.KeyInfo, msg *types.Message) (s *types.SignedMessage, err error) {
	// 离线签名
	s, err = local.WalletSignMessage(types.KTBLS, ki.PrivateKey, msg)
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
	if err = local.WalletVerifyMessage(s); err != nil {
		fmt.Println("验证签名失败", err)
	}
	return
}

func BroadcastMessage(s *types.SignedMessage) (err error) {
	// 消息广播
	mid, err := Client.MpoolPush(context.Background(), s)
	if err != nil {
		fmt.Println("消息广播失败")
		fmt.Println(err)
	} else {
		fmt.Println("消息发送成功，message id:")
		fmt.Println(mid.String())
	}
	return
}

func lotusURL(host string) string {
	return fmt.Sprintf("http://%v:1234/rpc/v0", host)
}

func SetHostWithToken(str, token string) {
	Client = filecoin.NewClient(lotusURL(str), token)
}

func main() {
	var err error
	// 静态设置
	toAddr, err := address.NewFromString("t1gvkap5jmv5k7gwpa42zj43i2oaai5zg74n66xra")
	pkStr := "c3pS5JcZEM1C5Yukor63mQ8DvADh1qQN/GrUsRA20XE="
	if err != nil {
		fmt.Println("收款地址错误", err)
		return
	}
	var fromAddr *address.Address
	var pk []byte
	pk, err = base64.StdEncoding.DecodeString(pkStr)
	if err != nil {
		fmt.Println("pk解码失败", err)
		return
	}
	// 设置key
	ki := &types.KeyInfo{
		Type:       types.KTBLS,
		PrivateKey: pk,
	}
	// 由key生成并确认地址
	fromAddr, err = local.WalletPrivateToAddress(crypto.SigTypeBLS, ki.PrivateKey)
	if err != nil {
		fmt.Println("生成地址失败", err)
		return
	}
	// 设置主机
	address.CurrentNetwork = address.Mainnet
	SetHostWithToken("172.16.30.117", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.ppK_nggwygh6kCPDlktdBtkGaqQXxoXM99iNx3-tZ9E")
	// 在此编辑message
	// 需要获取Nonce
	msg := &types.Message{
		Version:    0,
		To:         toAddr,
		From:       *fromAddr,
		Nonce:      14,
		Value:      filecoin.FromFil(decimal.NewFromFloat(123.456)),
		GasLimit:   0,
		GasFeeCap:  abi.NewTokenAmount(100),
		GasPremium: abi.NewTokenAmount(100),
		Method:     0,
		Params:     nil,
	}
	// 设置最大手续费
	maxFee := filecoin.FromFil(decimal.NewFromFloat(0.0001))
	// 估算GasLimit
	msg, err = Client.GasEstimateMessageGas(context.Background(), msg, &types.MessageSendSpec{MaxFee: maxFee}, nil)
	if err != nil {
		fmt.Println("GasEstimateMessageGas错误", err)
		return
	}
	signedMsg, err := OfflineSign(ki, msg)
	if err != nil {
		fmt.Println("签名失败", err)
		return
	}
	err = BroadcastMessage(signedMsg)
	if err != nil {
		fmt.Println("广播失败", err)
		return
	}
	return
}
