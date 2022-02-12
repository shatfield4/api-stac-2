package repository

import (
	"api-grandtheftbacon/module/nft/domain/entity"
)

// NFTQueryRepositoryInterface holds the implementable method for nft query repository
type NFTQueryRepositoryInterface interface {
	SelectMintedTokensByMintType(mintType string) ([]entity.OnMintEvent, error)
	SelectStakedTokenByTokenID(tokenID int64) (entity.StakedToken, error)
	SelectStakedTokensByOwnerAddress(ownerAddress string) ([]entity.StakedToken, error)
	SelectStakedTokensByMintType(mintType string) ([]entity.StakedToken, error)
	SelectStolenTokensByMintType(mintType string) ([]entity.OnMintEvent, error)
}
