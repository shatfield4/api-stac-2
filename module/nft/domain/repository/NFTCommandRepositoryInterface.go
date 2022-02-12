package repository

import (
	"api-grandtheftbacon/module/nft/infrastructure/repository/types"
)

// NFTCommandRepositoryInterface holds the implementable methods for nft command repository
type NFTCommandRepositoryInterface interface {
	DeleteStakedToken(tokenID int64) error
	InsertClaimEvent(data types.Claim) error
	InsertMintEvent(data types.Mint) error
	InsertStakedToken(data types.StakedToken) error
}
