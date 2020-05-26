package client

import (
	"context"
	"google.golang.org/grpc"

	"github.com/figment-networks/oasis-rpc-proxy/grpc/state/statepb"
)

var (
	_ StateClient = (*stateClient)(nil)
)

type StateClient interface {
	GetByHeight(int64) (*statepb.GetByHeightResponse, error)
}

func NewStateClient(conn *grpc.ClientConn) StateClient {
	return &stateClient{
		client: statepb.NewStateServiceClient(conn),
	}
}

type stateClient struct {
	client statepb.StateServiceClient
}

func (r *stateClient) GetByHeight(h int64) (*statepb.GetByHeightResponse, error) {
	ctx := context.Background()

	return r.client.GetByHeight(ctx, &statepb.GetByHeightRequest{Height: h})
}
