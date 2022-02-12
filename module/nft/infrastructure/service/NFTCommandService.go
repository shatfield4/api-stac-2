package service

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	baconContract "api-grandtheftbacon/infrastructures/smartcontracts/bacon"
	"api-grandtheftbacon/module/nft/domain/repository"
	repositoryTypes "api-grandtheftbacon/module/nft/infrastructure/repository/types"
	"api-grandtheftbacon/module/nft/infrastructure/service/types"
)

// NFTCommandService handles the nft query service logic
type NFTCommandService struct {
	repository.NFTCommandRepositoryInterface
	repository.NFTQueryRepositoryInterface
	BaconContractInstance *baconContract.Bacon
}

func (service *NFTCommandService) ClaimToken(ctx context.Context, data types.Claim) error {
	// get token owner wallet address
	res, err := service.NFTQueryRepositoryInterface.SelectStakedTokenByTokenID(data.TokenID)
	if err != nil {
		return err
	}

	// add claim record
	err = service.NFTCommandRepositoryInterface.InsertClaimEvent(repositoryTypes.Claim{
		WalletAddress: res.OwnerAddress,
		TokenID:       data.TokenID,
		MintType:      data.MintType,
		GreaseEarned:  data.GreaseEarned,
	})

	// check if token is unstaked
	if data.Unstaked {
		err := service.NFTCommandRepositoryInterface.DeleteStakedToken(data.TokenID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (service *NFTCommandService) CreateMintEvent(ctx context.Context, data types.Mint) error {
	err := service.NFTCommandRepositoryInterface.InsertMintEvent(repositoryTypes.Mint{
		TokenID:       data.TokenID,
		MintType:      data.MintType,
		MinterAddress: data.MinterAddress,
		OwnerAddress:  data.OwnerAddress,
	})
	if err != nil {
		return err
	}

	return nil
}

func (service *NFTCommandService) StakeToken(ctx context.Context, data types.StakedToken) error {
	token := repositoryTypes.StakedToken{
		TokenID:        data.TokenID,
		MintType:       "bacon", // default
		OwnerAddress:   data.OwnerAddress,
		BlockTimestamp: data.BlockTimestamp,
	}

	// check if token is bacon or cop
	isCop, _, err := service.BaconContractInstance.GetTokenTraits(&bind.CallOpts{}, big.NewInt(data.TokenID))
	if err == nil && isCop {
		token.MintType = "cop"
	}

	err = service.NFTCommandRepositoryInterface.InsertStakedToken(token)
	if err != nil {
		return err
	}

	return nil
}
