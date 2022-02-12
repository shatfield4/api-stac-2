package application

import (
	"api-grandtheftbacon/module/nft/infrastructure/service/types"
	"context"
)

// NFTCommandServiceInterface holds the implementable methods for the nft command service
type NFTCommandServiceInterface interface {
	ClaimToken(ctx context.Context, data types.Claim) error
	CreateMintEvent(ctx context.Context, data types.Mint) error
	StakeToken(ctx context.Context, data types.StakedToken) error
}
