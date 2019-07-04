package calc

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/processout/grpc-go-pool"
	"google.golang.org/grpc"

	calcSource "github.com/instabledesign/go-skeleton/internal/grpc/calc/pb"
)

type Client struct {
	url     string
	timeout time.Duration
	pool    *grpcpool.Pool
}

func (c *Client) getClientAndConnection(ctx context.Context) (calcSource.CalcClient, *grpcpool.ClientConn, error) {
	conn, err := c.pool.Get(ctx)
	if err != nil {
		return nil, conn, err
	}
	client := calcSource.NewCalcClient(conn.ClientConn)
	return client, conn, nil
}

func (c *Client) Add() (int64, error) {
	client, conn, err := c.getClientAndConnection(context.Background())
	defer conn.Close()
	if err != nil {
		return 0, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	res, err := client.Operation(ctx, &calcSource.MyRequest{Type: calcSource.MyRequest_ADDITION, OperandeA: 12, OperandeB: 3})
	if err != nil {
		return 0, err
	}
	return res.GetResult(), nil
}

func NewClient(url string, timeout time.Duration) *Client {
	factory := func() (*grpc.ClientConn, error) {
		return grpc.Dial(url, grpc.WithInsecure())
	}

	p, err := grpcpool.New(factory, 10, 50, 5*time.Minute)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to create grpc pool"))
	}

	return &Client{
		url:     url,
		timeout: timeout,
		pool:    p,
	}
}
