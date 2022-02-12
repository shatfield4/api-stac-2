package interfaces

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"api-grandtheftbacon/infrastructures/smartcontracts/bacon"
	"api-grandtheftbacon/infrastructures/smartcontracts/thefryingpan"
	"api-grandtheftbacon/internal/geth"
	serviceTypes "api-grandtheftbacon/module/nft/infrastructure/service/types"
)

var (
	BaconContractAddress        common.Address
	BaconContractABI            abi.ABI
	TheFryingPanContractAddress common.Address
	TheFryingPanContractABI     abi.ABI
)

func BaconContractEventWatcher() {
	// for nft command service
	commandService := NFTCommandServiceDI()

	// watch bacon minted event
	go func() {
		logs := make(chan *bacon.BaconBaconMinted)

		sub, err := BaconContractInstance.WatchBaconMinted(nil, logs, nil, nil, nil)
		if err != nil {
			panic(err)
		}

		for {
			select {
			case err := <-sub.Err():
				log.Println("[error] WatchBaconMinted listener suddenly stopped", err)
				panic(err)
			case t := <-logs:
				event := serviceTypes.Mint{
					TokenID:       t.TokenId.Int64(),
					MintType:      "bacon",
					MinterAddress: t.Minter.Hex(),
					OwnerAddress:  t.Owner.Hex(),
				}

				err := commandService.CreateMintEvent(context.TODO(), event)
				if err != nil {
					log.Println("[error] WatchBaconMinted cannot create bacon mint event", err)
				}
			}
		}
	}()

	// watch cop minted event
	go func() {
		logs := make(chan *bacon.BaconCopMinted)

		sub, err := BaconContractInstance.WatchCopMinted(nil, logs, nil, nil, nil)
		if err != nil {
			panic(err)
		}

		for {
			select {
			case err := <-sub.Err():
				log.Println("[error] WatchCopMinted listener suddenly stopped", err)
				panic(err)
			case t := <-logs:
				event := serviceTypes.Mint{
					TokenID:       t.TokenId.Int64(),
					MintType:      "cop",
					MinterAddress: t.Minter.Hex(),
					OwnerAddress:  t.Owner.Hex(),
				}

				err := commandService.CreateMintEvent(context.TODO(), event)
				if err != nil {
					log.Println("[error] WatchCopMinted cannot create cop mint event", err)
				}
			}
		}
	}()
}

func TheFryingPanContractEventWatcher() {
	// for nft command service
	commandService := NFTCommandServiceDI()

	// watch token staked event
	go func() {
		logs := make(chan *thefryingpan.ThefryingpanTokenStaked)

		sub, err := TheFryingPanContractInstance.WatchTokenStaked(nil, logs, nil, nil)
		if err != nil {
			panic(err)
		}

		for {
			select {
			case err := <-sub.Err():
				log.Println("[error] WatchTokenStaked listener suddenly stopped", err)
				panic(err)
			case t := <-logs:
				event := serviceTypes.StakedToken{
					TokenID:        t.TokenId.Int64(),
					OwnerAddress:   t.Owner.Hex(),
					BlockTimestamp: t.Value.Int64(),
				}

				err := commandService.StakeToken(context.TODO(), event)
				if err != nil {
					log.Println("[error] WatchTokenStaked cannot create token staked event", err)
				}
			}
		}
	}()

	// watch bacon claimed event
	go func() {
		logs := make(chan *thefryingpan.ThefryingpanBaconClaimed)

		sub, err := TheFryingPanContractInstance.WatchBaconClaimed(nil, logs, nil)
		if err != nil {
			panic(err)
		}

		for {
			select {
			case err := <-sub.Err():
				log.Println("[error] WatchBaconClaimed listener suddenly stopped", err)
				panic(err)
			case t := <-logs:
				// convert grease to float
				greaseEarnedBigInt := t.Earned
				greaseEarned, _ := geth.ToDecimal(greaseEarnedBigInt, 18).Float64()

				event := serviceTypes.Claim{
					TokenID:      t.TokenId.Int64(),
					MintType:     "bacon",
					GreaseEarned: greaseEarned,
					Unstaked:     t.Unstaked,
				}

				err := commandService.ClaimToken(context.TODO(), event)
				if err != nil {
					log.Println("[error] WatchBaconClaimed cannot create bacon claim event", err)
				}
			}
		}
	}()

	// watch cop claimed event
	go func() {
		logs := make(chan *thefryingpan.ThefryingpanBaconCopClaimed)

		sub, err := TheFryingPanContractInstance.WatchBaconCopClaimed(nil, logs, nil)
		if err != nil {
			panic(err)
		}

		for {
			select {
			case err := <-sub.Err():
				log.Println("[error] WatchBaconCopClaimed listener suddenly stopped", err)
				panic(err)
			case t := <-logs:
				// convert grease to float
				greaseEarnedBigInt := t.Earned
				greaseEarned, _ := geth.ToDecimal(greaseEarnedBigInt, 18).Float64()

				event := serviceTypes.Claim{
					TokenID:      t.TokenId.Int64(),
					MintType:     "cop",
					GreaseEarned: greaseEarned,
					Unstaked:     t.Unstaked,
				}

				err := commandService.ClaimToken(context.TODO(), event)
				if err != nil {
					log.Println("[error] WatchBaconCopClaimed cannot create cop claim event", err)
				}
			}
		}
	}()
}
