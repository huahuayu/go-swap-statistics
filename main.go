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
	logrus.Info("serve at port: " + global.Port + " using node: " + global.Node)
	logrus.Info("try visit: http://localhost:" + global.Port+"/swaps?pair=0x1b96b92314c44b159149f7e0303511fb2fc4774f&from=12643365&to=12643377")
	srv.ListenAndServe()
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
		fmt.Println("e2")
		fmt.Fprintf(w,err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(bs)
}
