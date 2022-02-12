package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-chi/chi"

	"api-grandtheftbacon/interfaces/http/rest/viewmodels"
	"api-grandtheftbacon/internal/constants"
	"api-grandtheftbacon/internal/errors"
	"api-grandtheftbacon/module/nft/application"
	types "api-grandtheftbacon/module/nft/interfaces/http"
)

// NFTQueryController request controller for nft query
type NFTQueryController struct {
	application.NFTQueryServiceInterface
}

// GetGameStatus retrieves the nft by id
func (controller *NFTQueryController) GetGameStatus(w http.ResponseWriter, r *http.Request) {
	res := controller.NFTQueryServiceInterface.GetGameStatus(context.TODO())

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully fetched game status.",
		Data: &types.GameStatusResponse{
			BaconsMinted: res.BaconsMinted,
			BaconsStaked: res.BaconsStaked,
			BaconsStolen: res.BaconsStolen,
			CopsMinted:   res.CopsMinted,
			CopsStaked:   res.CopsStaked,
			CopsStolen:   res.CopsStolen,
		},
	}

	response.JSON(w)
}

// GetNFTByID retrieves the nft by id
func (controller *NFTQueryController) GetNFTByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "tokenID")

	nftID, err := strconv.Atoi(idStr)
	if err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusUnprocessableEntity,
			Success:   false,
			Message:   "Invalid nft ID",
			ErrorCode: errors.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	metadata := controller.NFTQueryServiceInterface.GetNFTByID(context.TODO(), int64(nftID))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(metadata)
}

// GetStakedTokensByOwnerAddress retrieves the nft by id
func (controller *NFTQueryController) GetStakedTokensByOwnerAddress(w http.ResponseWriter, r *http.Request) {
	ownerAddress := chi.URLParam(r, "ownerAddress")
	if len(ownerAddress) == 0 {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusUnprocessableEntity,
			Success:   false,
			Message:   "Invalid payload.",
			ErrorCode: errors.InvalidPayload,
		}

		response.JSON(w)
		return
	}

	res, err := controller.NFTQueryServiceInterface.GetStakedTokensByOwnerAddress(context.TODO(), ownerAddress)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.MissingRecord:
			httpCode = http.StatusNotFound
			errorMsg = "No records found."
		default:
			httpCode = http.StatusInternalServerError
			errorMsg = "Please contact technical support."
		}

		response := viewmodels.HTTPResponseVM{
			Status:    httpCode,
			Success:   false,
			Message:   errorMsg,
			ErrorCode: err.Error(),
		}

		response.JSON(w)
		return
	}

	var stakedTokens []types.GetStakedTokenResponse

	for _, stakedToken := range res {
		stakedTokens = append(stakedTokens, types.GetStakedTokenResponse{
			TokenID:        stakedToken.TokenID,
			MintType:       stakedToken.MintType,
			OwnerAddress:   stakedToken.MintType,
			BlockTimestamp: stakedToken.BlockTimestamp,
		})
	}

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully fetched staked tokens.",
		Data:    stakedTokens,
	}

	response.JSON(w)
}

// GetNFTImage get nft image by filename
func (controller *NFTQueryController) GetNFTImage(w http.ResponseWriter, r *http.Request) {
	fileName := chi.URLParam(r, "fileName")
	rootPath, _ := os.Getwd()

	// check if token id is already minted
	splitFileNameArr := strings.Split(fileName, ".")
	tokenId, err := strconv.Atoi(splitFileNameArr[0]) // get token id number
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	isTokenMinted := controller.NFTQueryServiceInterface.IsTokenIDMinted(int64(tokenId))
	if !isTokenMinted {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// check if token id is a bacon or cop
	imagePath := constants.BaconNFTImagePath

	isCop := controller.NFTQueryServiceInterface.IsCop(int64(tokenId))
	if isCop {
		imagePath = constants.CopNFTImagePath
	}

	// serve image
	fileBytes, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/%s", rootPath, imagePath, fileName))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
}
