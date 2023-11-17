package tests

import (
	cowswap "github.com/amitbasuri/go-cowswap"
	"github.com/amitbasuri/go-cowswap/subgraph"
	"testing"
)

func TestGqlClient_NewClient(t *testing.T) {
	client, err := subgraph.NewSubgraphClient(cowswap.SUBGRAPH_MAINNET)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("client: %v", client)
}
