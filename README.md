## go-swap-statistics

get all the swap details within the block range

## get started

```bash
git clone https://github.com/huahuayu/go-swap-statistics.git
cd go-swap-statistics
go run main.go
```

## example

after the service up: 

```bash
curl http://localhost:8080/swaps?pair=0x1b96b92314c44b159149f7e0303511fb2fc4774f&from=12643365&to=12643377 
```

response: 

```json
[
  {
    "blockNumber": 12643366,
    "txHash": "0x72435e79898b97e9537f408cb2a0aac6ac3de4c744203ed64fc51a7778a4dd26",
    "index": 1,
    "pair": {
      "address": "0x1B96B92314C44b159149f7E0303511fB2Fc4774f",
      "name": "WBNB-BUSD",
      "token0": "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c",
      "token1": "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"
    },
    "tokenIn": {
      "address": "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c",
      "symbol": "WBNB",
      "decimal": 18,
      "position": 0,
      "amount": "864658749306541950"
    },
    "tokenOut": {
      "address": "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56",
      "symbol": "BUSD",
      "decimal": 18,
      "position": 1,
      "amount": "564278537813833653492"
    },
    "sync": {
      "reserve0": "2182311902354743252393",
      "reserve1": "1426471059586802372815975",
      "reserveIn": "2182311902354743252393",
      "reserveOut": "1426471059586802372815975"
    },
    "gas": {
      "gasPrice": "20000000000",
      "gasLimit": "334587",
      "gasUsed": "256381",
      "txFee": "0.00512762"
    },
    "price": {
      "token0Price": "652.602588323296528",
      "token1Price": "0.0015323261321553"
    }
  },
  {
    "blockNumber": 12643374,
    "txHash": "0x8b826208644779f74bf0ecdc123c2990411ef3668b2589f3b9462340a8f043fb",
    "index": 212,
    "pair": {
      "address": "0x1B96B92314C44b159149f7E0303511fB2Fc4774f",
      "name": "WBNB-BUSD",
      "token0": "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c",
      "token1": "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"
    },
    "tokenIn": {
      "address": "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56",
      "symbol": "BUSD",
      "decimal": 18,
      "position": 1,
      "amount": "240000000000000000000"
    },
    "tokenOut": {
      "address": "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c",
      "symbol": "WBNB",
      "decimal": 18,
      "position": 0,
      "amount": "366372377362893251"
    },
    "sync": {
      "reserve0": "2181945529977380359142",
      "reserve1": "1426711059586802372815975",
      "reserveIn": "1426711059586802372815975",
      "reserveOut": "2181945529977380359142"
    },
    "gas": {
      "gasPrice": "5000000000",
      "gasLimit": "330256",
      "gasUsed": "176907",
      "txFee": "0.000884535"
    },
    "price": {
      "token0Price": "655.0712194175028527",
      "token1Price": "0.0015265515723454"
    }
  },
  {
    "blockNumber": 12643376,
    "txHash": "0x381c68a0282ca636321c8759b556c9def2daa55f93c9b076a3e2d33f38da3e6d",
    "index": 181,
    "pair": {
      "address": "0x1B96B92314C44b159149f7E0303511fB2Fc4774f",
      "name": "WBNB-BUSD",
      "token0": "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c",
      "token1": "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"
    },
    "tokenIn": {
      "address": "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56",
      "symbol": "BUSD",
      "decimal": 18,
      "position": 1,
      "amount": "28734187627437500000"
    },
    "tokenOut": {
      "address": "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c",
      "symbol": "WBNB",
      "decimal": 18,
      "position": 0,
      "amount": "43812016040996259"
    },
    "sync": {
      "reserve0": "2181901717961339362883",
      "reserve1": "1426739793774429810315975",
      "reserveIn": "1426739793774429810315975",
      "reserveOut": "2181901717961339362883"
    },
    "gas": {
      "gasPrice": "5000000000",
      "gasLimit": "288919",
      "gasUsed": "174542",
      "txFee": "0.00087271"
    },
    "price": {
      "token0Price": "655.8517553848704787",
      "token1Price": "0.0015247348075072"
    }
  }
]
```