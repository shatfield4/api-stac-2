package entity

import (
	"time"
)

// OnClaimEvent holds the on claim event entity fields
type OnClaimEvent struct {
	ID            int64
	WalletAddress string    `db:"wallet_address"`
	TokenID       int64     `db:"token_id"`
	MintType      string    `db:"mint_type"`
	GreaseEarned  float64   `db:"grease_earned"`
	CreatedAt     time.Time `db:"created_at"`
}

// GetModelName returns the model name of on claim event entity that can be used for naming schemas
func (entity *OnClaimEvent) GetModelName() string {
	return "on_claim_events"
}
