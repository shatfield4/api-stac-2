package repository

import (
	"errors"
	"fmt"
	"strings"

	"api-grandtheftbacon/infrastructures/database/mysql/types"
	apiError "api-grandtheftbacon/internal/errors"
	"api-grandtheftbacon/module/nft/domain/entity"
	repositoryTypes "api-grandtheftbacon/module/nft/infrastructure/repository/types"
)

// NFTCommandRepository handles the nft command repository logic
type NFTCommandRepository struct {
	types.MySQLDBHandlerInterface
}

// DeleteStakedToken unstake token
func (repository *NFTCommandRepository) DeleteStakedToken(tokenID int64) error {
	stakedToken := entity.StakedToken{
		TokenID: tokenID,
	}

	stmt := fmt.Sprintf("DELETE FROM %s WHERE token_id=:token_id", stakedToken.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, stakedToken)
	if err != nil {
		return err
	}

	return nil
}

// InsertClaimEvent creates a new claim event record
func (repository *NFTCommandRepository) InsertClaimEvent(data repositoryTypes.Claim) error {
	claim := entity.OnClaimEvent{
		WalletAddress: data.WalletAddress,
		TokenID:       data.TokenID,
		MintType:      data.MintType,
		GreaseEarned:  data.GreaseEarned,
	}

	stmt := fmt.Sprintf("INSERT INTO %s (wallet_address,token_id,mint_type,grease_earned) VALUES (:wallet_address,:token_id,:mint_type,:grease_earned)", claim.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, claim)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New(apiError.DuplicateRecord)
		}
		return errors.New(apiError.DatabaseError)
	}

	return nil
}

// InsertMintEvent creates a new mint event record
func (repository *NFTCommandRepository) InsertMintEvent(data repositoryTypes.Mint) error {
	nft := entity.OnMintEvent{
		TokenID:       data.TokenID,
		MintType:      data.MintType,
		MinterAddress: data.MinterAddress,
		OwnerAddress:  data.OwnerAddress,
	}

	stmt := fmt.Sprintf("INSERT INTO %s (token_id,mint_type,minter_address,owner_address) VALUES (:token_id,:mint_type,:minter_address,:owner_address)", nft.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, nft)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New(apiError.DuplicateRecord)
		}
		return errors.New(apiError.DatabaseError)
	}

	return nil
}

// InsertStakedToken creates a record for new staked token
func (repository *NFTCommandRepository) InsertStakedToken(data repositoryTypes.StakedToken) error {
	stakedToken := entity.StakedToken{
		TokenID:        data.TokenID,
		MintType:       data.MintType,
		OwnerAddress:   data.OwnerAddress,
		BlockTimestamp: data.BlockTimestamp,
	}

	stmt := fmt.Sprintf("INSERT INTO %s (token_id,mint_type,owner_address,block_timestamp) VALUES (:token_id,:mint_type,:owner_address,:block_timestamp)", stakedToken.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, stakedToken)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New(apiError.DuplicateRecord)
		}
		return errors.New(apiError.DatabaseError)
	}

	return nil
}
