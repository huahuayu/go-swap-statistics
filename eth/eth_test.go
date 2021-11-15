package eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

var url = "https://bsc-dataseed1.binance.org"

func TestSyncInfo(t *testing.T) {
	Init(url)
	info, err := Cli.FilterSync([]common.Address{common.HexToAddress("0x1b96b92314c44b159149f7e0303511fb2fc4774f")}, 12643483, 12643484)
	if err != nil {
		return
	}
	for _, s := range *info {
		fmt.Println(s.BlockNumber, s.TxHash, s.Index, s.SyncEvent.Reserve0, s.SyncEvent.Reserve1)
	}
}

func TestSwapInfo(t *testing.T) {
	Init(url)
	info, err := Cli.FilterSwap([]common.Address{common.HexToAddress("0x1b96b92314c44b159149f7e0303511fb2fc4774f")}, 12643483, 12643484)
	if err != nil {
		return
	}
	for _, s := range *info {
		fmt.Println(s.BlockNumber, s.TxHash, s.Index, s.SwapEvent.Amount0Out, s.SwapEvent.Amount1Out)
	}
}

func TestClient_ParseSwap(t *testing.T) {
	Init(url)
	swaps, err := Cli.ParseSwap(common.HexToAddress("0x1b96b92314c44b159149f7e0303511fb2fc4774f"), 12643365, 12643566)
	if err != nil {
		t.Error(err)
	}
	for _, s := range *swaps {
		fmt.Printf("%+v\n", s)
	}
}
