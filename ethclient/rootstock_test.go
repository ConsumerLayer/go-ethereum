package ethclient

import (
	"context"
	"github.com/hoaleee/go-ethereum/common"
	"math/big"
	"testing"
)

func Test_Rootstock(t *testing.T) {
	EthClient, err := Dial(`https://public-node.testnet.rsk.co`)
	if err != nil {
		t.Fatalf("error connecting %s", err)
	}
	block, err := EthClient.BlockByNumber(context.Background(), big.NewInt(1))
	if err != nil {
		t.Fatalf("crawl failed %s", err)
	}
	block1, err := EthClient.BlockByHash(context.Background(), block.Hash())
	if err != nil {
		t.Fatalf("crawl failed %s", err)
	}
	if block.Hash() != block1.Hash() || block.Hash() == common.BigToHash(big.NewInt(0)) {
		t.Fatalf(`Test failed`)
	}
}
