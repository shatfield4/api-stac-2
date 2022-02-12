package entity

// StakedToken holds the staked token entity fields
type StakedToken struct {
	TokenID        int64  `db:"token_id"`
	MintType       string `db:"mint_type"`
	OwnerAddress   string `db:"owner_address"`
	BlockTimestamp int64  `db:"block_timestamp"`
}

// GetModelName returns the model name of staked token entity that can be used for naming schemas
func (entity *StakedToken) GetModelName() string {
	return "staked_tokens"
}
