package tests

import (
	"context"
	"encoding/json"
	"github.com/itsahedge/go-cowswap/subgraph"
	"github.com/itsahedge/go-cowswap/util"
	"testing"
)

func Test_GetOrders(t *testing.T) {
	gql_client, err := subgraph.NewSubgraphClient(util.SUBGRAPH_MAINNET)
	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]interface{}{
		"orderBy":        "id",
		"orderDirection": "desc",
		"first":          5,
	}

	resp, err := gql_client.GetOrders(context.Background(), vars)
	if err != nil {
		t.Fatal(err)
	}

	r, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("%v", string(r))
}
