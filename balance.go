package go_cowswap

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// GetBalance - Retrieves the balance of the given address
func (c *Client) GetBalance(ctx context.Context, address string) (*big.Int, error) {
	if address == "" {
		return nil, errors.New("must provide an address")
	}

	// Convert the address from a hex string to an Ethereum address
	addr := common.HexToAddress(address)

	// Query the balance
	balance, err := c.EthClient.BalanceAt(ctx, addr, nil)
	if err != nil {
		return nil, err
	}

	return balance, nil
}
