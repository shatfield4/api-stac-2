package rest

import (
	"api-grandtheftbacon/module/nft/application"
)

// NFTCommandController request controller for nft command
type NFTCommandController struct {
	application.NFTCommandServiceInterface
}
