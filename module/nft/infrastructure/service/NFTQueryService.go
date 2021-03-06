package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	baconContract "api-grandtheftbacon/infrastructures/smartcontracts/bacon"
	"api-grandtheftbacon/internal/constants"
	"api-grandtheftbacon/module/nft/domain/entity"
	"api-grandtheftbacon/module/nft/domain/repository"
	"api-grandtheftbacon/module/nft/infrastructure/service/types"
)

// NFTQueryService handles the nft query service logic
type NFTQueryService struct {
	repository.NFTQueryRepositoryInterface
	BaconContractInstance *baconContract.Bacon
}

func (service *NFTQueryService) GetGameStatus(ctx context.Context) types.GameStatus {
	var status types.GameStatus

	// get bacons stats
	stonedApesMinted, err := service.NFTQueryRepositoryInterface.SelectMintedTokensByMintType("bacon")
	if err == nil {
		status.BaconsMinted = len(stonedApesMinted)
	}
	stonedApesStaked, err := service.NFTQueryRepositoryInterface.SelectStakedTokensByMintType("bacon")
	if err == nil {
		status.BaconsStaked = len(stonedApesStaked)
	}
	stonedApesStolen, err := service.NFTQueryRepositoryInterface.SelectStolenTokensByMintType("bacon")
	if err == nil {
		status.BaconsStolen = len(stonedApesStolen)
	}

	// get cops stats
	fedApesMinted, err := service.NFTQueryRepositoryInterface.SelectMintedTokensByMintType("cop")
	if err == nil {
		status.CopsMinted = len(fedApesMinted)
	}
	fedApesStaked, err := service.NFTQueryRepositoryInterface.SelectStakedTokensByMintType("cop")
	if err == nil {
		status.CopsStaked = len(fedApesStaked)
	}
	fedApesStolen, err := service.NFTQueryRepositoryInterface.SelectStolenTokensByMintType("cop")
	if err == nil {
		status.CopsStolen = len(fedApesStolen)
	}

	return status
}

// GetNFTByID retrieves the nft provided by its id
func (service *NFTQueryService) GetNFTByID(ctx context.Context, ID int64) map[string]interface{} {
	placeHolderMetadata := map[string]interface{}{
		"name":        "Stoned Ape Club",
		"description": "In a metaverse nearby cannabis has been made illegal to the point where the only way to obtain some quality kush is to risk it all and grow it yourself. These brave Stoned Apes are passionate in growing and smoking cannabis, so they built secret grow houses and are growing $TOKE daily. In the hopes of having plentiful weed to smoke and cash to stash away. Thousands Stoned Apes and Feds compete on the streets of the metaverse. Dynamic PVP staking game based on wolf.game. On-chain game. Your NFT is also a ticket to our evergrowing community of cannabis & real estate investors and enthusiasts.",
		"attributes":  []string{},
	}

	// check if token id is minted
	isMinted := service.IsTokenIDMinted(ID)
	if !isMinted {
		return placeHolderMetadata
	}

	// check if bacon or cop
	metadataPath := constants.BaconNFTMetadataPath

	isCop := service.IsCop(ID)
	if isCop {
		metadataPath = constants.CopNFTMetadataPath
		log.Println("This is a cop")

	}

	var metadata map[string]interface{}

	rootPath, _ := os.Getwd()
	plan, _ := ioutil.ReadFile(fmt.Sprintf("%s/%s/%d.json", rootPath, metadataPath, ID))
	err := json.Unmarshal(plan, &metadata)
	if err != nil {
		log.Println("[Error] cannot unmarshal metadata json", err)
		return placeHolderMetadata
	}

	return metadata
}

// GetStakedTokensByOwnerAddress get staked tokens by owner address
func (service *NFTQueryService) GetStakedTokensByOwnerAddress(ctx context.Context, ownerAddress string) ([]entity.StakedToken, error) {
	res, err := service.NFTQueryRepositoryInterface.SelectStakedTokensByOwnerAddress(ownerAddress)
	if err != nil {
		return []entity.StakedToken{}, err
	}

	return res, nil
}

// IsTokenIDMinted check if token id is already minted
func (service *NFTQueryService) IsTokenIDMinted(tokenID int64) bool {
	_, err := service.BaconContractInstance.TokenURI(&bind.CallOpts{}, big.NewInt(tokenID))
	if err != nil {
		return false
	}

	return true
}

// IsCop check if token id is a cop
func (service *NFTQueryService) IsCop(tokenID int64) bool {
	isCop, _, err := service.BaconContractInstance.GetTokenTraits(&bind.CallOpts{}, big.NewInt(tokenID))
	if err != nil {
		log.Println("[error] isCop check call failed", err)
		return false
	}

	return isCop
}
