package attack

import (
	"context"

	attack_grpc "github.com/RioRizkyRainey/pokedex/internal/gateway/delivery/attack/grpc"
	"google.golang.org/grpc"
)

func AttackClientGrpc(grpcClient *grpc.ClientConn) *Client {
	client := &Client{
		rpcClient: attack_grpc.NewAttackHandlerClient(grpcClient),
	}

	return client
}

type Client struct {
	rpcClient attack_grpc.AttackHandlerClient
}

func (c *Client) GetDamage(ctx context.Context, attackName string, attackMove string, defendName string) (int32, error) {
	params := &attack_grpc.Params{
		AttackName: attackName,
		AttackMove: attackMove,
		DefendName: defendName,
	}

	data, err := c.rpcClient.GetDamage(ctx, params)

	if err != nil {
		return 0, err
	}

	return data.Data, nil
}
