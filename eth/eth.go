package eth

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/sirupsen/logrus"
	"math/big"
	"strings"
	ierc20 "swap-statistics/eth/contract/erc20"
	pairContract "swap-statistics/eth/contract/pair"
)

const (
	IUniswapV2PairABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"blockTimestampLast\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
)

var (
	Cli               *Client
	contractAbi, _    = abi.JSON(strings.NewReader(IUniswapV2PairABI))
	syncEventIdentity = struct {
		name string
		hash string
	}{name: "Sync", hash: "0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1"}
	swapEventIdentity = struct {
		name string
		hash string
	}{name: "Swap", hash: "0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"}
)

type (
	Client struct {
		EthClient *ethclient.Client
		RpcClient *rpc.Client
		NetworkId *big.Int
		Url       string
	}

	Swap struct {
		BlockNumber uint64 `json:"blockNumber"`
		TxHash      string `json:"txHash"`
		Index       uint   `json:"index"`
		Pair        struct {
			Address string `json:"address"`
			Dex     string `json:"-"`
			Name    string `json:"name"`
			Token0 string `json:"token0"`
			Token1 string `json:"token1"`
		} `json:"pair"`
		TokenIn struct {
			Address  string `json:"address"`
			Symbol   string `json:"symbol"`
			Decimal  int    `json:"decimal"`
			Position int    `json:"position"`
			Amount   string `json:"amount"`
		} `json:"tokenIn"`
		TokenOut struct {
			Address  string `json:"address"`
			Symbol   string `json:"symbol"`
			Decimal  int    `json:"decimal"`
			Position int    `json:"position"`
			Amount   string `json:"amount"`
		} `json:"tokenOut"`
		Sync struct {
			Reserve0   string `json:"reserve0"`
			Reserve1   string `json:"reserve1"`
			ReserveIn  string `json:"reserveIn"`
			ReserveOut string `json:"reserveOut"`
		} `json:"sync"`
	}

	SyncInfo struct {
		BlockNumber uint64
		TxHash      string
		Index       uint
		SyncEvent   SyncEvent
	}

	SwapInfo struct {
		BlockNumber uint64
		TxHash      string
		Index       uint
		SwapEvent   SwapEvent
	}

	SyncEvent struct {
		Reserve0 *big.Int
		Reserve1 *big.Int
	}

	SwapEvent struct {
		Amount0In  *big.Int
		Amount1In  *big.Int
		Amount0Out *big.Int
		Amount1Out *big.Int
	}
)

func Init(url string) {
	var err error
	Cli, err = NewClient(url)
	if err != nil {
		panic(err)
	}
}

func NewClient(url string) (*Client, error) {
	client := new(Client)
	if ethClient, err := ethclient.Dial(url); err != nil {
		panic(err)
		return nil, err
	} else {
		client.EthClient = ethClient
	}
	if networkId, err := client.EthClient.NetworkID(context.Background()); err != nil {
		panic(err)
		return nil, err
	} else {
		client.NetworkId = networkId
	}
	if rpcClient, err := rpc.Dial(url); err != nil {
		panic(err)
		return nil, err
	} else {
		client.RpcClient = rpcClient
	}
	return client, nil
}

func (client *Client) FilterSync(pairs []common.Address, fromBlock, toBlock int64) (*[]SyncInfo, error) {
	query := ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		Addresses: pairs,
		Topics:    [][]common.Hash{{common.HexToHash(syncEventIdentity.hash)}},
	}
	logs, err := client.EthClient.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}
	syncInfos := make([]SyncInfo, 0)
	for _, l := range logs {
		if l.Removed{
			continue
		}
		syncEvent := new(SyncEvent)
		err := contractAbi.UnpackIntoInterface(syncEvent, syncEventIdentity.name, l.Data)
		if err != nil {
			logrus.Error("UnpackIntoInterface: ", err, l.TxHash.String())
			continue
		}
		syncInfos = append(syncInfos, SyncInfo{
			BlockNumber: l.BlockNumber,
			TxHash:      l.TxHash.String(),
			Index:       l.TxIndex,
			SyncEvent:   *syncEvent,
		})
	}
	return &syncInfos, nil
}

func (client *Client) FilterSwap(pairs []common.Address, fromBlock, toBlock int64) (*[]SwapInfo, error) {
	query := ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		Addresses: pairs,
		Topics:    [][]common.Hash{{common.HexToHash(swapEventIdentity.hash)}},
	}
	logs, err := client.EthClient.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}
	swapInfos := make([]SwapInfo, 0)
	for _, l := range logs {
		if l.Removed{
			continue
		}
		swapEvent := new(SwapEvent)
		err := contractAbi.UnpackIntoInterface(swapEvent, swapEventIdentity.name, l.Data)
		if err != nil {
			logrus.Error("UnpackIntoInterface: ", err, l.TxHash.String())
			continue
		}
		swapInfos = append(swapInfos, SwapInfo{
			BlockNumber: l.BlockNumber,
			TxHash:      l.TxHash.String(),
			Index:       l.TxIndex,
			SwapEvent:   *swapEvent,
		})
	}
	return &swapInfos, nil
}

func (client *Client) ParseSwap(pair common.Address, fromBlock int64, toBlock int64) (*[]Swap, error) {
	pairContractObj, err := pairContract.NewIUniswapV2Pair(pair, client.EthClient)
	if err != nil {
		return nil, err
	}
	dex, err := pairContractObj.Name(nil)
	if err != nil {
		return nil, err
	}
	token0Address, err := pairContractObj.Token0(nil)
	if err != nil {
		return nil, err
	}
	token1Address, err := pairContractObj.Token1(nil)
	if err != nil {
		return nil, err
	}
	token0ContractObj, err := ierc20.NewIerc20(token0Address, client.EthClient)
	if err != nil {
		return nil, err
	}
	token1ContractObj, err := ierc20.NewIerc20(token1Address, client.EthClient)
	if err != nil {
		return nil, err
	}
	symbol0, err := token0ContractObj.Symbol(nil)
	if err != nil {
		return nil, err
	}
	symbol1, err := token1ContractObj.Symbol(nil)
	if err != nil {
		return nil, err
	}
	decimal0, err := token0ContractObj.Decimals(nil)
	if err != nil {
		return nil, err
	}
	decimal1, err := token1ContractObj.Decimals(nil)
	if err != nil {
		return nil, err
	}
	swapInfos, err := client.FilterSwap([]common.Address{pair}, fromBlock, toBlock)
	if err != nil {
		return nil, err
	}
	syncInfos, err := client.FilterSync([]common.Address{pair}, fromBlock, toBlock)
	if err != nil {
		return nil, err
	}
	swaps := make([]Swap, 0)
	for i, s := range *swapInfos {
		swap := new(Swap)
		if s.SwapEvent.Amount0In.Cmp(big.NewInt(0)) > 0 {
			swap = &Swap{
				BlockNumber: s.BlockNumber,
				TxHash:      s.TxHash,
				Index:       s.Index,
				Pair: struct {
					Address string `json:"address"`
					Dex     string `json:"-"`
					Name    string `json:"name"`
					Token0 string `json:"token0"`
					Token1 string `json:"token1"`
				}{
					Address: pair.String(),
					Dex:     dex,
					Name:    symbol0 + "-" + symbol1,
					Token0: token0Address.String(),
					Token1: token1Address.String(),
				},
				TokenIn: struct {
					Address  string `json:"address"`
					Symbol   string `json:"symbol"`
					Decimal  int    `json:"decimal"`
					Position int    `json:"position"`
					Amount   string `json:"amount"`
				}{
					Address:  token0Address.String(),
					Symbol:   symbol0,
					Decimal:  int(decimal0),
					Position: 0,
					Amount:   s.SwapEvent.Amount0In.String(),
				},
				TokenOut: struct {
					Address  string `json:"address"`
					Symbol   string `json:"symbol"`
					Decimal  int    `json:"decimal"`
					Position int    `json:"position"`
					Amount   string `json:"amount"`
				}{
					Address:  token1Address.String(),
					Symbol:   symbol1,
					Decimal:  int(decimal1),
					Position: 1,
					Amount:   s.SwapEvent.Amount1Out.String(),
				},
				Sync: struct {
					Reserve0   string `json:"reserve0"`
					Reserve1   string `json:"reserve1"`
					ReserveIn  string `json:"reserveIn"`
					ReserveOut string `json:"reserveOut"`
				}(struct {
					Reserve0   string
					Reserve1   string
					ReserveIn  string
					ReserveOut string
				}{Reserve0: (*syncInfos)[i].SyncEvent.Reserve0.String(), Reserve1: (*syncInfos)[i].SyncEvent.Reserve1.String(), ReserveIn: (*syncInfos)[i].SyncEvent.Reserve0.String(), ReserveOut: (*syncInfos)[i].SyncEvent.Reserve1.String()}),
			}
		} else {
			swap = &Swap{
				BlockNumber: s.BlockNumber,
				TxHash:      s.TxHash,
				Index:       s.Index,
				Pair: struct {
					Address string `json:"address"`
					Dex     string `json:"-"`
					Name    string `json:"name"`
					Token0 string `json:"token0"`
					Token1 string `json:"token1"`
				}{
					Address: pair.String(),
					Dex:     dex,
					Name:    symbol0 + "-" + symbol1,
					Token0: token0Address.String(),
					Token1: token1Address.String(),
				},
				TokenIn: struct {
					Address  string `json:"address"`
					Symbol   string `json:"symbol"`
					Decimal  int    `json:"decimal"`
					Position int    `json:"position"`
					Amount   string `json:"amount"`
				}{
					Address:  token1Address.String(),
					Symbol:   symbol1,
					Decimal:  int(decimal1),
					Position: 1,
					Amount:   s.SwapEvent.Amount1In.String(),
				},
				TokenOut: struct {
					Address  string `json:"address"`
					Symbol   string `json:"symbol"`
					Decimal  int    `json:"decimal"`
					Position int    `json:"position"`
					Amount   string `json:"amount"`
				}{
					Address:  token0Address.String(),
					Symbol:   symbol0,
					Decimal:  int(decimal0),
					Position: 0,
					Amount:   s.SwapEvent.Amount0Out.String(),
				},
				Sync: struct {
					Reserve0   string `json:"reserve0"`
					Reserve1   string `json:"reserve1"`
					ReserveIn  string `json:"reserveIn"`
					ReserveOut string `json:"reserveOut"`
				}(struct {
					Reserve0   string
					Reserve1   string
					ReserveIn  string
					ReserveOut string
				}{Reserve0: (*syncInfos)[i].SyncEvent.Reserve0.String(), Reserve1: (*syncInfos)[i].SyncEvent.Reserve1.String(), ReserveOut: (*syncInfos)[i].SyncEvent.Reserve0.String(), ReserveIn: (*syncInfos)[i].SyncEvent.Reserve1.String()}),
			}
		}
		swaps = append(swaps, *swap)
	}
	return &swaps, nil
}
