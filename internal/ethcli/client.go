package ethcli

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	url    string
	client *ethclient.Client
}

func NewClient(url string) (*Client, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	return &Client{
		url:    url,
		client: client,
	}, nil
}

func (c *Client) ChainID() (*big.Int, error) {
	return c.client.NetworkID(context.Background())
}
