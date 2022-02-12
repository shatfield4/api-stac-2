package types

type Claim struct {
	WalletAddress string
	TokenID       int64
	MintType      string
	GreaseEarned  float64
}

type Mint struct {
	TokenID       int64
	MintType      string
	MinterAddress string
	OwnerAddress  string
}

type StakedToken struct {
	TokenID        int64
	MintType       string
	OwnerAddress   string
	BlockTimestamp int64
}
