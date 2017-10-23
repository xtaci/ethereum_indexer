package main

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/Jeffail/gabs"
)

func main() {
	fmt.Println("vim-go")
}

const (
	transfer = "0xa9059cbb"
)

const (
	methodGetBlockByNumber       = "eth_getBlockByNumber"
	methodGetBlockNumber         = "eth_blockNumber"
	methodeGetTransactionReceipt = "eth_getTransactionReceipt"
)

// indexing prefix
// eg : FR_0xaabcd, TO_0xdefg
const (
	idxBlockNumber = "BL_"
	idxFrom        = "FR_"
	idxTo          = "TO_"
	idxHash        = "HA_"
)

func indexTransactions(bts []byte) error {
	jsonParsed, err := gabs.ParseJSON(bts)
	if value, ok := jsonParsed.Path("error.code").Data().(float64); ok && value > 0 {
		msg, _ := jsonParsed.Path("error.message").Data().(string)
		return errors.New(msg)
	}

	count, err := jsonParsed.ArrayCount("result", "transactions")
	if err != nil {
		return err
	}

	if count == 0 {
		return nil
	}

	var timestamp big.Int
	var blocknumber big.Int
	t, _ := jsonParsed.Path("result.timestamp").Data().(string)
	r, _ := jsonParsed.Path("result.number").Data().(string)

	timestamp.UnmarshalJSON([]byte(t))
	blocknumber.UnmarshalJSON([]byte(r))
	//children, _ := jsonParsed.S("result", "transactions").Children()
	return nil
}
