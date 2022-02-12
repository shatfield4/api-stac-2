package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	hystrix_config "api-grandtheftbacon/configs/hystrix"
	"api-grandtheftbacon/module/nft/domain/repository"
	repositoryTypes "api-grandtheftbacon/module/nft/infrastructure/repository/types"
)

// NFTCommandRepositoryCircuitBreaker circuit breaker for record command repository
type NFTCommandRepositoryCircuitBreaker struct {
	repository.NFTCommandRepositoryInterface
}

var config = hystrix_config.Config{}

// DeleteStakedToken decorator pattern to delete staked token
func (repository *NFTCommandRepositoryCircuitBreaker) DeleteStakedToken(tokenID int64) error {
	hystrix.ConfigureCommand("delete_staked_token", config.Settings())
	errors := hystrix.Go("delete_staked_token", func() error {
		err := repository.NFTCommandRepositoryInterface.DeleteStakedToken(tokenID)
		if err != nil {
			return err
		}

		return nil
	}, nil)

	select {
	case err := <-errors:
		return err
	default:
		return nil
	}
}

// InsertClaimEvent decorator pattern to insert claim event
func (repository *NFTCommandRepositoryCircuitBreaker) InsertClaimEvent(data repositoryTypes.Claim) error {
	hystrix.ConfigureCommand("insert_claim_event", config.Settings())
	errors := hystrix.Go("insert_claim_event", func() error {
		err := repository.NFTCommandRepositoryInterface.InsertClaimEvent(data)
		if err != nil {
			return err
		}

		return nil
	}, nil)

	select {
	case err := <-errors:
		return err
	default:
		return nil
	}
}

// InsertMintEvent decorator pattern to insert mint event
func (repository *NFTCommandRepositoryCircuitBreaker) InsertMintEvent(data repositoryTypes.Mint) error {
	hystrix.ConfigureCommand("insert_mint_event", config.Settings())
	errors := hystrix.Go("insert_mint_event", func() error {
		err := repository.NFTCommandRepositoryInterface.InsertMintEvent(data)
		if err != nil {
			return err
		}

		return nil
	}, nil)

	select {
	case err := <-errors:
		return err
	default:
		return nil
	}
}

// InsertStakedToken decorator pattern to insert staked token
func (repository *NFTCommandRepositoryCircuitBreaker) InsertStakedToken(data repositoryTypes.StakedToken) error {
	hystrix.ConfigureCommand("insert_staked_token", config.Settings())
	errors := hystrix.Go("insert_staked_token", func() error {
		err := repository.NFTCommandRepositoryInterface.InsertStakedToken(data)
		if err != nil {
			return err
		}

		return nil
	}, nil)

	select {
	case err := <-errors:
		return err
	default:
		return nil
	}
}
