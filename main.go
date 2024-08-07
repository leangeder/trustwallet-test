package main

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	_u    = "https://polygon-rpc.com/"
	_port = 9090
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	if os.Getenv("RPC_URL") != "" {
		_u = os.Getenv("RPC_URL")
	}

	client, err := ethclient.Dial(_u)
	if err != nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n", err)
		os.Exit(1)
	}
	_chan := make(chan *big.Int)

	go func(ctx context.Context, client *ethclient.Client, _chan chan *big.Int) {
		for {
			time.Sleep(5 * time.Second)

			block, err := client.BlockNumber(ctx)
			if err != nil {
				fmt.Errorf("Error in BlockNumber(): %v\n", err)
			}
			if block > 0 {
				_chan <- new(big.Int).SetUint64(block)
				fmt.Printf("Block Number via BlockNumber function: %d\n", block)
			}
		}
	}(ctx, client, _chan)

	go func(ctx context.Context, client *ethclient.Client, _chan chan *big.Int) {
		for block := range _chan {
			_block, err := client.BlockByNumber(ctx, block)
			if err != nil {
				fmt.Errorf("Error in BlockNumber(): %v\n", err)
			}
			fmt.Printf("Block Number via BlockByNumber function: %d\n", _block.NumberU64())
		}
	}(ctx, client, _chan)

	// s := http.Server{
	// 	Addr: fmt.Sprintf("0.0.0.0:%d", _port),
	// 	// Handler:      handlers.CORS(headers, methods, origins)(router),
	// 	IdleTimeout:  120 * time.Second,
	// 	ReadTimeout:  1 * time.Second,
	// 	WriteTimeout: 1 * time.Second,
	// }

	// fmt.Println("Listening on", _port)
	// err = s.ListenAndServe()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	fmt.Println("Listening on", _port)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", _port), mux)

}
