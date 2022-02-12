package http

// CreateRecordRequest request struct for create record
type CreateRecordRequest struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

// GameStatusResponse response struct for game status
type GameStatusResponse struct {
	BaconsMinted int `json:"baconsMinted"`
	BaconsStaked int `json:"baconsStaked"`
	BaconsStolen int `json:"baconsStolen"`
	CopsMinted   int `json:"copsMinted"`
	CopsStaked   int `json:"copsStaked"`
	CopsStolen   int `json:"copsStolen"`
}

// GetStakedTokenResponse response struct for staked tokens
type GetStakedTokenResponse struct {
	TokenID        int64  `json:"tokenId"`
	MintType       string `json:"mintType"`
	OwnerAddress   string `json:"ownerAddress"`
	BlockTimestamp int64  `json:"blockTimestamp"`
}
