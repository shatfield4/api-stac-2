package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	"api-grandtheftbacon/module/nft/domain/entity"
	"api-grandtheftbacon/module/nft/domain/repository"
)

// NFTQueryRepositoryCircuitBreaker circuit breaker for record query repository
type NFTQueryRepositoryCircuitBreaker struct {
	repository.NFTQueryRepositoryInterface
}

// SelectMintedTokensByMintType is a decorator for the select minted tokens by mint type
func (repository *NFTQueryRepositoryCircuitBreaker) SelectMintedTokensByMintType(mintType string) ([]entity.OnMintEvent, error) {
	output := make(chan []entity.OnMintEvent, 1)
	hystrix.ConfigureCommand("select_minted_tokens_by_mint_type", config.Settings())
	errors := hystrix.Go("select_minted_tokens_by_mint_type", func() error {
		tokens, err := repository.NFTQueryRepositoryInterface.SelectMintedTokensByMintType(mintType)
		if err != nil {
			return err
		}

		output <- tokens
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return []entity.OnMintEvent{}, err
	}
}

// SelectStakedTokenByTokenID is a decorator for the select staked token by token id
func (repository *NFTQueryRepositoryCircuitBreaker) SelectStakedTokenByTokenID(tokenID int64) (entity.StakedToken, error) {
	output := make(chan entity.StakedToken, 1)
	hystrix.ConfigureCommand("select_staked_token_by_token_id", config.Settings())
	errors := hystrix.Go("select_staked_token_by_token_id", func() error {
		token, err := repository.NFTQueryRepositoryInterface.SelectStakedTokenByTokenID(tokenID)
		if err != nil {
			return err
		}

		output <- token
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return entity.StakedToken{}, err
	}
}

// SelectStakedTokensByOwnerAddress is a decorator for the select staked tokens by owner address
func (repository *NFTQueryRepositoryCircuitBreaker) SelectStakedTokensByOwnerAddress(ownerAddress string) ([]entity.StakedToken, error) {
	output := make(chan []entity.StakedToken, 1)
	hystrix.ConfigureCommand("select_staked_tokens_by_owner_address", config.Settings())
	errors := hystrix.Go("select_staked_tokens_by_owner_address", func() error {
		tokens, err := repository.NFTQueryRepositoryInterface.SelectStakedTokensByOwnerAddress(ownerAddress)
		if err != nil {
			return err
		}

		output <- tokens
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return []entity.StakedToken{}, err
	}
}

// SelectStakedTokensByMintType is a decorator for the select staked tokens by mint type
func (repository *NFTQueryRepositoryCircuitBreaker) SelectStakedTokensByMintType(mintType string) ([]entity.StakedToken, error) {
	output := make(chan []entity.StakedToken, 1)
	hystrix.ConfigureCommand("select_staked_tokens_by_mint_type", config.Settings())
	errors := hystrix.Go("select_staked_tokens_by_mint_type", func() error {
		tokens, err := repository.NFTQueryRepositoryInterface.SelectStakedTokensByMintType(mintType)
		if err != nil {
			return err
		}

		output <- tokens
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return []entity.StakedToken{}, err
	}
}

// SelectStolenTokensByMintType is a decorator for the select stolen tokens by mint type
func (repository *NFTQueryRepositoryCircuitBreaker) SelectStolenTokensByMintType(mintType string) ([]entity.OnMintEvent, error) {
	output := make(chan []entity.OnMintEvent, 1)
	hystrix.ConfigureCommand("select_stolen_tokens_by_mint_type", config.Settings())
	errors := hystrix.Go("select_stolen_tokens_by_mint_type", func() error {
		tokens, err := repository.NFTQueryRepositoryInterface.SelectStolenTokensByMintType(mintType)
		if err != nil {
			return err
		}

		output <- tokens
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return []entity.OnMintEvent{}, err
	}
}
