package repository

import (
	"errors"
	"fmt"

	"api-grandtheftbacon/infrastructures/database/mysql/types"
	apiError "api-grandtheftbacon/internal/errors"
	"api-grandtheftbacon/module/nft/domain/entity"
)

// NFTQueryRepository handles the nft query repository logic
type NFTQueryRepository struct {
	types.MySQLDBHandlerInterface
}

// SelectMintedTokensByMintType select total staked by mint type
func (repository *NFTQueryRepository) SelectMintedTokensByMintType(mintType string) ([]entity.OnMintEvent, error) {
	var token entity.OnMintEvent
	var tokens []entity.OnMintEvent

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE mint_type=:mint_type", token.GetModelName())
	err := repository.Query(stmt, map[string]interface{}{
		"mint_type": mintType,
	}, &tokens)
	if err != nil {
		return []entity.OnMintEvent{}, errors.New(apiError.DatabaseError)
	} else if len(tokens) == 0 {
		return []entity.OnMintEvent{}, errors.New(apiError.MissingRecord)
	}

	return tokens, nil
}

// SelectStakedTokenByTokenID select staked token by token id
func (repository *NFTQueryRepository) SelectStakedTokenByTokenID(tokenID int64) (entity.StakedToken, error) {
	var token entity.StakedToken
	var tokens []entity.StakedToken

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE token_id=:token_id", token.GetModelName())
	err := repository.Query(stmt, map[string]interface{}{
		"token_id": tokenID,
	}, &tokens)
	if err != nil {
		return token, errors.New(apiError.DatabaseError)
	} else if len(tokens) == 0 {
		return token, errors.New(apiError.MissingRecord)
	}

	return tokens[0], nil
}

// SelectStakedTokensByOwnerAddress select staked tokens by owner address
func (repository *NFTQueryRepository) SelectStakedTokensByOwnerAddress(ownerAddress string) ([]entity.StakedToken, error) {
	var token entity.StakedToken
	var tokens []entity.StakedToken

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE owner_address=:owner_address", token.GetModelName())
	err := repository.Query(stmt, map[string]interface{}{
		"owner_address": ownerAddress,
	}, &tokens)
	if err != nil {
		return []entity.StakedToken{}, errors.New(apiError.DatabaseError)
	} else if len(tokens) == 0 {
		return []entity.StakedToken{}, errors.New(apiError.MissingRecord)
	}

	return tokens, nil
}

// SelectStakedTokensByMintType select total staked by mint type
func (repository *NFTQueryRepository) SelectStakedTokensByMintType(mintType string) ([]entity.StakedToken, error) {
	var token entity.StakedToken
	var tokens []entity.StakedToken

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE mint_type=:mint_type", token.GetModelName())
	err := repository.Query(stmt, map[string]interface{}{
		"mint_type": mintType,
	}, &tokens)
	if err != nil {
		return []entity.StakedToken{}, errors.New(apiError.DatabaseError)
	} else if len(tokens) == 0 {
		return []entity.StakedToken{}, errors.New(apiError.MissingRecord)
	}

	return tokens, nil
}

// SelectStolenTokensByMintType select total staked by mint type
func (repository *NFTQueryRepository) SelectStolenTokensByMintType(mintType string) ([]entity.OnMintEvent, error) {
	var token entity.OnMintEvent
	var tokens []entity.OnMintEvent

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE mint_type=:mint_type AND minter_address!=owner_address", token.GetModelName())
	err := repository.Query(stmt, map[string]interface{}{
		"mint_type": mintType,
	}, &tokens)
	if err != nil {
		return []entity.OnMintEvent{}, errors.New(apiError.DatabaseError)
	} else if len(tokens) == 0 {
		return []entity.OnMintEvent{}, errors.New(apiError.MissingRecord)
	}

	return tokens, nil
}
