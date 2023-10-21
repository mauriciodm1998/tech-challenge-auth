package customer

import (
	"context"
	"tech-challenge-auth/internal/canonical"
	"tech-challenge-auth/internal/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	cfg = &config.Cfg
)

type customerService struct {
	client CustomerServiceClient
}

type CustomerService interface {
	Get(context.Context, canonical.Login) (canonical.Login, error)
}

func New() CustomerService {
	grpcClient := NewCustomerServiceClient(
		ConnectGrpc(cfg.Integration.Customer),
	)

	return &customerService{
		client: grpcClient,
	}
}

func ConnectGrpc(address string) *grpc.ClientConn {
	return connectGRPC(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func connectGRPC(address string, options ...grpc.DialOption) *grpc.ClientConn {
	client, err := grpc.Dial(address, options...)
	if err != nil {
		panic(err)
	}

	return client
}

func (c *customerService) Get(ctx context.Context, login canonical.Login) (canonical.Login, error) {
	customer, err := c.client.Get(ctx, translateToGRPC(login))
	if err != nil {
		return canonical.Login{}, err
	}

	return translateToCanonical(customer), nil
}
