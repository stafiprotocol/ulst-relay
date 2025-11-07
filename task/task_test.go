package task

import (
	"testing"
	"ulst-relay/pkg/eth1"

	"github.com/ethereum/go-ethereum/common"
)

func Test_GetTransactionByHash(t *testing.T) {
	client, err := eth1.NewClient("https://1rpc.io/sepolia", nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	txHash := common.HexToHash("0xba5d40cc36a2914d42df560afe2399fa65369a8bc10d731c9f3d85ebdb21b6a8")
	tx, pending, err := client.TransactionByHash(txHash)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tx, pending)

	receipt, err := client.TransactionReceipt(txHash)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(receipt)
}
