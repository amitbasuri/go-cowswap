package go_cowswap_test

import (
	"context"
	go_cowswap "github.com/itsahedge/go-cowswap/cmd/go-cowswap"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	if err != nil {
		t.Fatal(err)
	}
	if client.TransactionSigner != nil {
		t.Logf("initialized client with a transaction signer: %v", client)
	} else {
		t.Logf("initialized client without a transaction signer: %v", client)
	}

	block, err := client.EthClient.BlockNumber(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("current block: %v", block)
	addressList := util.TOKEN_ADDRESSES[util.Options.Network]
	for s, s2 := range addressList {
		t.Logf("%v, %v \n", s, s2)
	}
	resp, err := client.Version()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("version resp: %v", resp)
}
