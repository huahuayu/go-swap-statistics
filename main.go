package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"swap-statistics/eth"
	"swap-statistics/global"
	"time"
)

func init() {
	flag.StringVar(&global.Port, "port", "8080", "server port")
	flag.StringVar(&global.Node, "node", "https://bsc-dataseed1.binance.org", "node url")
	flag.Parse()
}

func main() {
	eth.Init(global.Node)
	srv := http.Server{
		Addr:    ":" + global.Port,
		WriteTimeout: 1 * time.Minute,
	}
	http.HandleFunc("/swaps", swapsHandler)
	http.HandleFunc("/tx", txHandler)
	logrus.Info("serve at port: " + global.Port + " using node: " + global.Node)
	logrus.Info("get tx block: http://localhost:" + global.Port+"/tx?tx_hash=0xd6c03e92ec778951f4a3d13b75e2a71454d13daae83e1048bcfef8ae7ddff80f")
	logrus.Info("get swaps in block range: http://localhost:" + global.Port+"/swaps?pair=0x1b96b92314c44b159149f7e0303511fb2fc4774f&from=12643365&to=12643377")
	srv.ListenAndServe()
}

func txHandler(w http.ResponseWriter, r *http.Request){
	logrus.Info("request: ", r.URL.String())
	txHash, ok := r.URL.Query()["tx_hash"]
	if !ok || len(txHash[0]) < 1{
		w.WriteHeader(400)
		fmt.Fprintf(w, "tx_hash is missing\n")
		return
	}
	blockNumber, err := eth.Cli.TxBlock(common.HexToHash(txHash[0]))
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	bs, err := json.Marshal(struct {
		TxHash string `json:"txHash"`
		BlockNumber uint64 `json:"blockNumber"`
	}{
		TxHash:      txHash[0],
		BlockNumber: blockNumber.Uint64(),
	})
	if err != nil {
		fmt.Fprintf(w,err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(bs)
	return
}

func swapsHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Info("request: ", r.URL.String())
	pair, ok := r.URL.Query()["pair"]
	if !ok || len(pair[0]) < 1{
		w.WriteHeader(400)
		fmt.Fprintf(w, "pair is missing\n")
		return
	}
	from, ok := r.URL.Query()["from"]
	if !ok || len(from[0]) < 1{
		w.WriteHeader(400)
		fmt.Fprintf(w, "from is missing\n")
		return
	}
	to, ok := r.URL.Query()["to"]
	if !ok || len(to[0]) < 1{
		w.WriteHeader(400)
		fmt.Fprintf(w, "to is missing\n")
		return
	}
	fromBlock,_ := strconv.ParseInt(from[0],10,64)
	toBlock,_ := strconv.ParseInt(to[0],10,64)
	if fromBlock > toBlock{
		w.WriteHeader(400)
		fmt.Fprintf(w, "to block should great or equal than from block\n")
		return
	}
	swap, err := eth.Cli.ParseSwap(common.HexToAddress(pair[0]), fromBlock, toBlock)
	if err != nil {
		fmt.Fprintf(w,err.Error())
		return
	}
	bs, err := json.Marshal(swap)
	if err != nil {
		fmt.Fprintf(w,err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(bs)
}
