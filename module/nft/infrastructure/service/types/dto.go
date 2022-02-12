package types

type Claim struct {
	TokenID      int64
	MintType     string
	GreaseEarned float64
	Unstaked     bool
}

type GameStatus struct {
	BaconsMinted int
	BaconsStaked int
	BaconsStolen int
	CopsMinted   int
	CopsStaked   int
	CopsStolen   int
}

type Mint struct {
	TokenID       int64
	MintType      string
	MinterAddress string
	OwnerAddress  string
}

type StakedToken struct {
	TokenID        int64
	OwnerAddress   string
	BlockTimestamp int64
}
