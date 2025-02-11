package tests

import (
	"context"
	"encoding/json"
	cowswap "github.com/amitbasuri/go-cowswap"
	"github.com/amitbasuri/go-cowswap/subgraph"
	"testing"
)

func Test_GetUniswapPools(t *testing.T) {
	gql_client, err := subgraph.NewSubgraphClient(cowswap.SUBGRAPH_MAINNET)
	if err != nil {
		t.Fatal(err)
	}

	//vars := map[string]interface{}{
	//	"orderBy":        "id",
	//	"orderDirection": "desc",
	//	"first":          5,
	//}

	resp, err := gql_client.GetUniswapPools(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	r, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("%v", string(r))
}
