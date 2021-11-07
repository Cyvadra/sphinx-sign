package fil

import (
	"fmt"

	"github.com/cyvadra/filecoin-client"
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

func SetHostWithToken(str, token string) {
	Client = filecoin.NewClient(fmt.Sprintf("http://%v:1234/rpc/v0", str), token)
}
