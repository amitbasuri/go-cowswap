package go_cowswap_test

import (
	"context"
	go_cowswap "github.com/itsahedge/go-cowswap/cmd/go-cowswap"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := go_cowswap.NewClient(util.Options)
	res, statusCode, err := client.Version(context.Background())
	if err != nil {
		t.Fatalf("Version err: %v", err)
	}
	t.Logf("status code: %v, response: %v", statusCode, res)
}
