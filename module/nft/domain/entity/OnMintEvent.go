package entity

import (
	"time"
)

// OnMintEvent holds the on mint event entity fields
type OnMintEvent struct {
	TokenID       int64     `db:"token_id"`
	MintType      string    `db:"mint_type"`
	MinterAddress string    `db:"minter_address"`
	OwnerAddress  string    `db:"owner_address"`
	CreatedAt     time.Time `db:"created_at"`
}

// GetModelName returns the model name of on mint event entity that can be used for naming schemas
func (entity *OnMintEvent) GetModelName() string {
	return "on_mint_events"
}
