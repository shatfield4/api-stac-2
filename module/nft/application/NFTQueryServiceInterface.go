package application

import (
	"context"

	"api-grandtheftbacon/module/nft/domain/entity"
	"api-grandtheftbacon/module/nft/infrastructure/service/types"
)

// NFTQueryServiceInterface holds the implementable methods for the nft query service
type NFTQueryServiceInterface interface {
	GetGameStatus(ctx context.Context) types.GameStatus
	GetNFTByID(ctx context.Context, ID int64) map[string]interface{}
	GetStakedTokensByOwnerAddress(ctx context.Context, ownerAddress string) ([]entity.StakedToken, error)
	IsTokenIDMinted(tokenID int64) bool
	IsCop(tokenID int64) bool
}
