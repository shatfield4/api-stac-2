// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bacon

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BaconMetaData contains all meta data concerning the Bacon contract.
var BaconMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"StonedApeBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"StonedApeMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"FedApeBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"FedApeMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20PaymentReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"PayeeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaymentReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaymentReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_MINT_QTY\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAID_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_contractBaseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"changePricePerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPaidTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenTraits\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenWriteBlock\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"greaseERC20\",\"outputs\":[{\"internalType\":\"contractIGREASE\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"isApprovedOrOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"}],\"name\":\"isTokenValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"qty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_seed\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mintCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"payee\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicSaleStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"}],\"name\":\"reclaimERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"erc721Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"reclaimERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"released\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"released\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newBaseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newuri\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fryingPanAddress\",\"type\":\"address\"}],\"name\":\"setFryingPan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGreaseAddress\",\"type\":\"address\"}],\"name\":\"setGrease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"setMaxAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newQty\",\"type\":\"uint16\"}],\"name\":\"setMaxMintQty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"setMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"setPaidTokensAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_setPaused\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setPublicSaleStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRandomizer\",\"type\":\"address\"}],\"name\":\"setRandomizer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setWhitelistStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"shares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"theFryingPan\",\"outputs\":[{\"internalType\":\"contractIFRYINGPAN\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"totalReleased\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalReleased\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"qty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_seed\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"whitelistMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BaconABI is the input ABI used to generate the binding from.
// Deprecated: Use BaconMetaData.ABI instead.
var BaconABI = BaconMetaData.ABI

// Bacon is an auto generated Go binding around an Ethereum contract.
type Bacon struct {
	BaconCaller     // Read-only binding to the contract
	BaconTransactor // Write-only binding to the contract
	BaconFilterer   // Log filterer for contract events
}

// BaconCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaconCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaconTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaconTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaconFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaconFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaconSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaconSession struct {
	Contract     *Bacon            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaconCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaconCallerSession struct {
	Contract *BaconCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BaconTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaconTransactorSession struct {
	Contract     *BaconTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaconRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaconRaw struct {
	Contract *Bacon // Generic contract binding to access the raw methods on
}

// BaconCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaconCallerRaw struct {
	Contract *BaconCaller // Generic read-only contract binding to access the raw methods on
}

// BaconTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaconTransactorRaw struct {
	Contract *BaconTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBacon creates a new instance of Bacon, bound to a specific deployed contract.
func NewBacon(address common.Address, backend bind.ContractBackend) (*Bacon, error) {
	contract, err := bindBacon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bacon{BaconCaller: BaconCaller{contract: contract}, BaconTransactor: BaconTransactor{contract: contract}, BaconFilterer: BaconFilterer{contract: contract}}, nil
}

// NewBaconCaller creates a new read-only instance of Bacon, bound to a specific deployed contract.
func NewBaconCaller(address common.Address, caller bind.ContractCaller) (*BaconCaller, error) {
	contract, err := bindBacon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaconCaller{contract: contract}, nil
}

// NewBaconTransactor creates a new write-only instance of Bacon, bound to a specific deployed contract.
func NewBaconTransactor(address common.Address, transactor bind.ContractTransactor) (*BaconTransactor, error) {
	contract, err := bindBacon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaconTransactor{contract: contract}, nil
}

// NewBaconFilterer creates a new log filterer instance of Bacon, bound to a specific deployed contract.
func NewBaconFilterer(address common.Address, filterer bind.ContractFilterer) (*BaconFilterer, error) {
	contract, err := bindBacon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaconFilterer{contract: contract}, nil
}

// bindBacon binds a generic wrapper to an already deployed contract.
func bindBacon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaconABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bacon *BaconRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bacon.Contract.BaconCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bacon *BaconRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bacon.Contract.BaconTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bacon *BaconRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bacon.Contract.BaconTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bacon *BaconCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bacon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bacon *BaconTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bacon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bacon *BaconTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bacon.Contract.contract.Transact(opts, method, params...)
}

// MAXMINTQTY is a free data retrieval call binding the contract method 0x16ffecb7.
//
// Solidity: function MAX_MINT_QTY() view returns(uint16)
func (_Bacon *BaconCaller) MAXMINTQTY(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "MAX_MINT_QTY")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXMINTQTY is a free data retrieval call binding the contract method 0x16ffecb7.
//
// Solidity: function MAX_MINT_QTY() view returns(uint16)
func (_Bacon *BaconSession) MAXMINTQTY() (uint16, error) {
	return _Bacon.Contract.MAXMINTQTY(&_Bacon.CallOpts)
}

// MAXMINTQTY is a free data retrieval call binding the contract method 0x16ffecb7.
//
// Solidity: function MAX_MINT_QTY() view returns(uint16)
func (_Bacon *BaconCallerSession) MAXMINTQTY() (uint16, error) {
	return _Bacon.Contract.MAXMINTQTY(&_Bacon.CallOpts)
}

// MAXTOKENS is a free data retrieval call binding the contract method 0xf47c84c5.
//
// Solidity: function MAX_TOKENS() view returns(uint256)
func (_Bacon *BaconCaller) MAXTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "MAX_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOKENS is a free data retrieval call binding the contract method 0xf47c84c5.
//
// Solidity: function MAX_TOKENS() view returns(uint256)
func (_Bacon *BaconSession) MAXTOKENS() (*big.Int, error) {
	return _Bacon.Contract.MAXTOKENS(&_Bacon.CallOpts)
}

// MAXTOKENS is a free data retrieval call binding the contract method 0xf47c84c5.
//
// Solidity: function MAX_TOKENS() view returns(uint256)
func (_Bacon *BaconCallerSession) MAXTOKENS() (*big.Int, error) {
	return _Bacon.Contract.MAXTOKENS(&_Bacon.CallOpts)
}

// MINTPRICE is a free data retrieval call binding the contract method 0xc002d23d.
//
// Solidity: function MINT_PRICE() view returns(uint256)
func (_Bacon *BaconCaller) MINTPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "MINT_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTPRICE is a free data retrieval call binding the contract method 0xc002d23d.
//
// Solidity: function MINT_PRICE() view returns(uint256)
func (_Bacon *BaconSession) MINTPRICE() (*big.Int, error) {
	return _Bacon.Contract.MINTPRICE(&_Bacon.CallOpts)
}

// MINTPRICE is a free data retrieval call binding the contract method 0xc002d23d.
//
// Solidity: function MINT_PRICE() view returns(uint256)
func (_Bacon *BaconCallerSession) MINTPRICE() (*big.Int, error) {
	return _Bacon.Contract.MINTPRICE(&_Bacon.CallOpts)
}

// PAIDTOKENS is a free data retrieval call binding the contract method 0xc084f540.
//
// Solidity: function PAID_TOKENS() view returns(uint256)
func (_Bacon *BaconCaller) PAIDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "PAID_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PAIDTOKENS is a free data retrieval call binding the contract method 0xc084f540.
//
// Solidity: function PAID_TOKENS() view returns(uint256)
func (_Bacon *BaconSession) PAIDTOKENS() (*big.Int, error) {
	return _Bacon.Contract.PAIDTOKENS(&_Bacon.CallOpts)
}

// PAIDTOKENS is a free data retrieval call binding the contract method 0xc084f540.
//
// Solidity: function PAID_TOKENS() view returns(uint256)
func (_Bacon *BaconCallerSession) PAIDTOKENS() (*big.Int, error) {
	return _Bacon.Contract.PAIDTOKENS(&_Bacon.CallOpts)
}

// ContractBaseURI is a free data retrieval call binding the contract method 0x7101ebca.
//
// Solidity: function _contractBaseURI() view returns(string)
func (_Bacon *BaconCaller) ContractBaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "_contractBaseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractBaseURI is a free data retrieval call binding the contract method 0x7101ebca.
//
// Solidity: function _contractBaseURI() view returns(string)
func (_Bacon *BaconSession) ContractBaseURI() (string, error) {
	return _Bacon.Contract.ContractBaseURI(&_Bacon.CallOpts)
}

// ContractBaseURI is a free data retrieval call binding the contract method 0x7101ebca.
//
// Solidity: function _contractBaseURI() view returns(string)
func (_Bacon *BaconCallerSession) ContractBaseURI() (string, error) {
	return _Bacon.Contract.ContractBaseURI(&_Bacon.CallOpts)
}

// ContractURI1 is a free data retrieval call binding the contract method 0xc0e72740.
//
// Solidity: function _contractURI() view returns(string)
func (_Bacon *BaconCaller) ContractURI1(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "_contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI1 is a free data retrieval call binding the contract method 0xc0e72740.
//
// Solidity: function _contractURI() view returns(string)
func (_Bacon *BaconSession) ContractURI1() (string, error) {
	return _Bacon.Contract.ContractURI1(&_Bacon.CallOpts)
}

// ContractURI1 is a free data retrieval call binding the contract method 0xc0e72740.
//
// Solidity: function _contractURI() view returns(string)
func (_Bacon *BaconCallerSession) ContractURI1() (string, error) {
	return _Bacon.Contract.ContractURI1(&_Bacon.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Bacon *BaconCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Bacon *BaconSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Bacon.Contract.BalanceOf(&_Bacon.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Bacon *BaconCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Bacon.Contract.BalanceOf(&_Bacon.CallOpts, owner)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Bacon *BaconCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Bacon *BaconSession) ContractURI() (string, error) {
	return _Bacon.Contract.ContractURI(&_Bacon.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Bacon *BaconCallerSession) ContractURI() (string, error) {
	return _Bacon.Contract.ContractURI(&_Bacon.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Bacon *BaconCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Bacon *BaconSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Bacon.Contract.GetApproved(&_Bacon.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Bacon *BaconCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Bacon.Contract.GetApproved(&_Bacon.CallOpts, tokenId)
}

// GetPaidTokens is a free data retrieval call binding the contract method 0x4018b1f8.
//
// Solidity: function getPaidTokens() view returns(uint256)
func (_Bacon *BaconCaller) GetPaidTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "getPaidTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPaidTokens is a free data retrieval call binding the contract method 0x4018b1f8.
//
// Solidity: function getPaidTokens() view returns(uint256)
func (_Bacon *BaconSession) GetPaidTokens() (*big.Int, error) {
	return _Bacon.Contract.GetPaidTokens(&_Bacon.CallOpts)
}

// GetPaidTokens is a free data retrieval call binding the contract method 0x4018b1f8.
//
// Solidity: function getPaidTokens() view returns(uint256)
func (_Bacon *BaconCallerSession) GetPaidTokens() (*big.Int, error) {
	return _Bacon.Contract.GetPaidTokens(&_Bacon.CallOpts)
}

// GetTokenTraits is a free data retrieval call binding the contract method 0x94e56847.
//
// Solidity: function getTokenTraits(uint256 tokenId) view returns(bool, uint256)
func (_Bacon *BaconCaller) GetTokenTraits(opts *bind.CallOpts, tokenId *big.Int) (bool, *big.Int, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "getTokenTraits", tokenId)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetTokenTraits is a free data retrieval call binding the contract method 0x94e56847.
//
// Solidity: function getTokenTraits(uint256 tokenId) view returns(bool, uint256)
func (_Bacon *BaconSession) GetTokenTraits(tokenId *big.Int) (bool, *big.Int, error) {
	return _Bacon.Contract.GetTokenTraits(&_Bacon.CallOpts, tokenId)
}

// GetTokenTraits is a free data retrieval call binding the contract method 0x94e56847.
//
// Solidity: function getTokenTraits(uint256 tokenId) view returns(bool, uint256)
func (_Bacon *BaconCallerSession) GetTokenTraits(tokenId *big.Int) (bool, *big.Int, error) {
	return _Bacon.Contract.GetTokenTraits(&_Bacon.CallOpts, tokenId)
}

// GetTokenWriteBlock is a free data retrieval call binding the contract method 0xebd17368.
//
// Solidity: function getTokenWriteBlock(uint256 tokenId) view returns(uint64)
func (_Bacon *BaconCaller) GetTokenWriteBlock(opts *bind.CallOpts, tokenId *big.Int) (uint64, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "getTokenWriteBlock", tokenId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetTokenWriteBlock is a free data retrieval call binding the contract method 0xebd17368.
//
// Solidity: function getTokenWriteBlock(uint256 tokenId) view returns(uint64)
func (_Bacon *BaconSession) GetTokenWriteBlock(tokenId *big.Int) (uint64, error) {
	return _Bacon.Contract.GetTokenWriteBlock(&_Bacon.CallOpts, tokenId)
}

// GetTokenWriteBlock is a free data retrieval call binding the contract method 0xebd17368.
//
// Solidity: function getTokenWriteBlock(uint256 tokenId) view returns(uint64)
func (_Bacon *BaconCallerSession) GetTokenWriteBlock(tokenId *big.Int) (uint64, error) {
	return _Bacon.Contract.GetTokenWriteBlock(&_Bacon.CallOpts, tokenId)
}

// GreaseERC20 is a free data retrieval call binding the contract method 0x964312e0.
//
// Solidity: function greaseERC20() view returns(address)
func (_Bacon *BaconCaller) GreaseERC20(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "greaseERC20")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GreaseERC20 is a free data retrieval call binding the contract method 0x964312e0.
//
// Solidity: function greaseERC20() view returns(address)
func (_Bacon *BaconSession) GreaseERC20() (common.Address, error) {
	return _Bacon.Contract.GreaseERC20(&_Bacon.CallOpts)
}

// GreaseERC20 is a free data retrieval call binding the contract method 0x964312e0.
//
// Solidity: function greaseERC20() view returns(address)
func (_Bacon *BaconCallerSession) GreaseERC20() (common.Address, error) {
	return _Bacon.Contract.GreaseERC20(&_Bacon.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Bacon *BaconCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Bacon *BaconSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Bacon.Contract.IsApprovedForAll(&_Bacon.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Bacon *BaconCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Bacon.Contract.IsApprovedForAll(&_Bacon.CallOpts, owner, operator)
}

// IsApprovedOrOwner is a free data retrieval call binding the contract method 0x430c2081.
//
// Solidity: function isApprovedOrOwner(address _spender, uint256 _tokenId) view returns(bool)
func (_Bacon *BaconCaller) IsApprovedOrOwner(opts *bind.CallOpts, _spender common.Address, _tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "isApprovedOrOwner", _spender, _tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedOrOwner is a free data retrieval call binding the contract method 0x430c2081.
//
// Solidity: function isApprovedOrOwner(address _spender, uint256 _tokenId) view returns(bool)
func (_Bacon *BaconSession) IsApprovedOrOwner(_spender common.Address, _tokenId *big.Int) (bool, error) {
	return _Bacon.Contract.IsApprovedOrOwner(&_Bacon.CallOpts, _spender, _tokenId)
}

// IsApprovedOrOwner is a free data retrieval call binding the contract method 0x430c2081.
//
// Solidity: function isApprovedOrOwner(address _spender, uint256 _tokenId) view returns(bool)
func (_Bacon *BaconCallerSession) IsApprovedOrOwner(_spender common.Address, _tokenId *big.Int) (bool, error) {
	return _Bacon.Contract.IsApprovedOrOwner(&_Bacon.CallOpts, _spender, _tokenId)
}

// IsTokenValid is a free data retrieval call binding the contract method 0x275efe1a.
//
// Solidity: function isTokenValid(address _to, uint256 _tokenId, bytes32[] _proof) view returns(bool)
func (_Bacon *BaconCaller) IsTokenValid(opts *bind.CallOpts, _to common.Address, _tokenId *big.Int, _proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "isTokenValid", _to, _tokenId, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenValid is a free data retrieval call binding the contract method 0x275efe1a.
//
// Solidity: function isTokenValid(address _to, uint256 _tokenId, bytes32[] _proof) view returns(bool)
func (_Bacon *BaconSession) IsTokenValid(_to common.Address, _tokenId *big.Int, _proof [][32]byte) (bool, error) {
	return _Bacon.Contract.IsTokenValid(&_Bacon.CallOpts, _to, _tokenId, _proof)
}

// IsTokenValid is a free data retrieval call binding the contract method 0x275efe1a.
//
// Solidity: function isTokenValid(address _to, uint256 _tokenId, bytes32[] _proof) view returns(bool)
func (_Bacon *BaconCallerSession) IsTokenValid(_to common.Address, _tokenId *big.Int, _proof [][32]byte) (bool, error) {
	return _Bacon.Contract.IsTokenValid(&_Bacon.CallOpts, _to, _tokenId, _proof)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_Bacon *BaconCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "locked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_Bacon *BaconSession) Locked() (bool, error) {
	return _Bacon.Contract.Locked(&_Bacon.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_Bacon *BaconCallerSession) Locked() (bool, error) {
	return _Bacon.Contract.Locked(&_Bacon.CallOpts)
}

// MintCost is a free data retrieval call binding the contract method 0x27de8f27.
//
// Solidity: function mintCost(uint256 tokenId) view returns(uint256)
func (_Bacon *BaconCaller) MintCost(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "mintCost", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintCost is a free data retrieval call binding the contract method 0x27de8f27.
//
// Solidity: function mintCost(uint256 tokenId) view returns(uint256)
func (_Bacon *BaconSession) MintCost(tokenId *big.Int) (*big.Int, error) {
	return _Bacon.Contract.MintCost(&_Bacon.CallOpts, tokenId)
}

// MintCost is a free data retrieval call binding the contract method 0x27de8f27.
//
// Solidity: function mintCost(uint256 tokenId) view returns(uint256)
func (_Bacon *BaconCallerSession) MintCost(tokenId *big.Int) (*big.Int, error) {
	return _Bacon.Contract.MintCost(&_Bacon.CallOpts, tokenId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bacon *BaconCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bacon *BaconSession) Name() (string, error) {
	return _Bacon.Contract.Name(&_Bacon.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bacon *BaconCallerSession) Name() (string, error) {
	return _Bacon.Contract.Name(&_Bacon.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bacon *BaconCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bacon *BaconSession) Owner() (common.Address, error) {
	return _Bacon.Contract.Owner(&_Bacon.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bacon *BaconCallerSession) Owner() (common.Address, error) {
	return _Bacon.Contract.Owner(&_Bacon.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Bacon *BaconCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Bacon *BaconSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Bacon.Contract.OwnerOf(&_Bacon.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Bacon *BaconCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Bacon.Contract.OwnerOf(&_Bacon.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bacon *BaconCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bacon *BaconSession) Paused() (bool, error) {
	return _Bacon.Contract.Paused(&_Bacon.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bacon *BaconCallerSession) Paused() (bool, error) {
	return _Bacon.Contract.Paused(&_Bacon.CallOpts)
}

// Payee is a free data retrieval call binding the contract method 0x8b83209b.
//
// Solidity: function payee(uint256 index) view returns(address)
func (_Bacon *BaconCaller) Payee(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "payee", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Payee is a free data retrieval call binding the contract method 0x8b83209b.
//
// Solidity: function payee(uint256 index) view returns(address)
func (_Bacon *BaconSession) Payee(index *big.Int) (common.Address, error) {
	return _Bacon.Contract.Payee(&_Bacon.CallOpts, index)
}

// Payee is a free data retrieval call binding the contract method 0x8b83209b.
//
// Solidity: function payee(uint256 index) view returns(address)
func (_Bacon *BaconCallerSession) Payee(index *big.Int) (common.Address, error) {
	return _Bacon.Contract.Payee(&_Bacon.CallOpts, index)
}

// PublicSaleStartTime is a free data retrieval call binding the contract method 0x6bb7b1d9.
//
// Solidity: function publicSaleStartTime() view returns(uint256)
func (_Bacon *BaconCaller) PublicSaleStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "publicSaleStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicSaleStartTime is a free data retrieval call binding the contract method 0x6bb7b1d9.
//
// Solidity: function publicSaleStartTime() view returns(uint256)
func (_Bacon *BaconSession) PublicSaleStartTime() (*big.Int, error) {
	return _Bacon.Contract.PublicSaleStartTime(&_Bacon.CallOpts)
}

// PublicSaleStartTime is a free data retrieval call binding the contract method 0x6bb7b1d9.
//
// Solidity: function publicSaleStartTime() view returns(uint256)
func (_Bacon *BaconCallerSession) PublicSaleStartTime() (*big.Int, error) {
	return _Bacon.Contract.PublicSaleStartTime(&_Bacon.CallOpts)
}

// Released is a free data retrieval call binding the contract method 0x406072a9.
//
// Solidity: function released(address token, address account) view returns(uint256)
func (_Bacon *BaconCaller) Released(opts *bind.CallOpts, token common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "released", token, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Released is a free data retrieval call binding the contract method 0x406072a9.
//
// Solidity: function released(address token, address account) view returns(uint256)
func (_Bacon *BaconSession) Released(token common.Address, account common.Address) (*big.Int, error) {
	return _Bacon.Contract.Released(&_Bacon.CallOpts, token, account)
}

// Released is a free data retrieval call binding the contract method 0x406072a9.
//
// Solidity: function released(address token, address account) view returns(uint256)
func (_Bacon *BaconCallerSession) Released(token common.Address, account common.Address) (*big.Int, error) {
	return _Bacon.Contract.Released(&_Bacon.CallOpts, token, account)
}

// Released0 is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released(address account) view returns(uint256)
func (_Bacon *BaconCaller) Released0(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "released0", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Released0 is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released(address account) view returns(uint256)
func (_Bacon *BaconSession) Released0(account common.Address) (*big.Int, error) {
	return _Bacon.Contract.Released0(&_Bacon.CallOpts, account)
}

// Released0 is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released(address account) view returns(uint256)
func (_Bacon *BaconCallerSession) Released0(account common.Address) (*big.Int, error) {
	return _Bacon.Contract.Released0(&_Bacon.CallOpts, account)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Bacon *BaconCaller) Root(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "root")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Bacon *BaconSession) Root() ([32]byte, error) {
	return _Bacon.Contract.Root(&_Bacon.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Bacon *BaconCallerSession) Root() ([32]byte, error) {
	return _Bacon.Contract.Root(&_Bacon.CallOpts)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address account) view returns(uint256)
func (_Bacon *BaconCaller) Shares(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "shares", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address account) view returns(uint256)
func (_Bacon *BaconSession) Shares(account common.Address) (*big.Int, error) {
	return _Bacon.Contract.Shares(&_Bacon.CallOpts, account)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address account) view returns(uint256)
func (_Bacon *BaconCallerSession) Shares(account common.Address) (*big.Int, error) {
	return _Bacon.Contract.Shares(&_Bacon.CallOpts, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bacon *BaconCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bacon *BaconSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bacon.Contract.SupportsInterface(&_Bacon.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bacon *BaconCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bacon.Contract.SupportsInterface(&_Bacon.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bacon *BaconCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bacon *BaconSession) Symbol() (string, error) {
	return _Bacon.Contract.Symbol(&_Bacon.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bacon *BaconCallerSession) Symbol() (string, error) {
	return _Bacon.Contract.Symbol(&_Bacon.CallOpts)
}

// TheFryingPan is a free data retrieval call binding the contract method 0xbc121866.
//
// Solidity: function theFryingPan() view returns(address)
func (_Bacon *BaconCaller) TheFryingPan(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "theFryingPan")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TheFryingPan is a free data retrieval call binding the contract method 0xbc121866.
//
// Solidity: function theFryingPan() view returns(address)
func (_Bacon *BaconSession) TheFryingPan() (common.Address, error) {
	return _Bacon.Contract.TheFryingPan(&_Bacon.CallOpts)
}

// TheFryingPan is a free data retrieval call binding the contract method 0xbc121866.
//
// Solidity: function theFryingPan() view returns(address)
func (_Bacon *BaconCallerSession) TheFryingPan() (common.Address, error) {
	return _Bacon.Contract.TheFryingPan(&_Bacon.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Bacon *BaconCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Bacon *BaconSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Bacon.Contract.TokenByIndex(&_Bacon.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Bacon *BaconCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Bacon.Contract.TokenByIndex(&_Bacon.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Bacon *BaconCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Bacon *BaconSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Bacon.Contract.TokenOfOwnerByIndex(&_Bacon.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Bacon *BaconCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Bacon.Contract.TokenOfOwnerByIndex(&_Bacon.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Bacon *BaconCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Bacon *BaconSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Bacon.Contract.TokenURI(&_Bacon.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Bacon *BaconCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Bacon.Contract.TokenURI(&_Bacon.CallOpts, _tokenId)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[])
func (_Bacon *BaconCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "tokensOfOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[])
func (_Bacon *BaconSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _Bacon.Contract.TokensOfOwner(&_Bacon.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[])
func (_Bacon *BaconCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _Bacon.Contract.TokensOfOwner(&_Bacon.CallOpts, _owner)
}

// TotalReleased is a free data retrieval call binding the contract method 0xd79779b2.
//
// Solidity: function totalReleased(address token) view returns(uint256)
func (_Bacon *BaconCaller) TotalReleased(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "totalReleased", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReleased is a free data retrieval call binding the contract method 0xd79779b2.
//
// Solidity: function totalReleased(address token) view returns(uint256)
func (_Bacon *BaconSession) TotalReleased(token common.Address) (*big.Int, error) {
	return _Bacon.Contract.TotalReleased(&_Bacon.CallOpts, token)
}

// TotalReleased is a free data retrieval call binding the contract method 0xd79779b2.
//
// Solidity: function totalReleased(address token) view returns(uint256)
func (_Bacon *BaconCallerSession) TotalReleased(token common.Address) (*big.Int, error) {
	return _Bacon.Contract.TotalReleased(&_Bacon.CallOpts, token)
}

// TotalReleased0 is a free data retrieval call binding the contract method 0xe33b7de3.
//
// Solidity: function totalReleased() view returns(uint256)
func (_Bacon *BaconCaller) TotalReleased0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "totalReleased0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReleased0 is a free data retrieval call binding the contract method 0xe33b7de3.
//
// Solidity: function totalReleased() view returns(uint256)
func (_Bacon *BaconSession) TotalReleased0() (*big.Int, error) {
	return _Bacon.Contract.TotalReleased0(&_Bacon.CallOpts)
}

// TotalReleased0 is a free data retrieval call binding the contract method 0xe33b7de3.
//
// Solidity: function totalReleased() view returns(uint256)
func (_Bacon *BaconCallerSession) TotalReleased0() (*big.Int, error) {
	return _Bacon.Contract.TotalReleased0(&_Bacon.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Bacon *BaconCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Bacon *BaconSession) TotalShares() (*big.Int, error) {
	return _Bacon.Contract.TotalShares(&_Bacon.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Bacon *BaconCallerSession) TotalShares() (*big.Int, error) {
	return _Bacon.Contract.TotalShares(&_Bacon.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bacon *BaconCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bacon *BaconSession) TotalSupply() (*big.Int, error) {
	return _Bacon.Contract.TotalSupply(&_Bacon.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bacon *BaconCallerSession) TotalSupply() (*big.Int, error) {
	return _Bacon.Contract.TotalSupply(&_Bacon.CallOpts)
}

// WhitelistStartTime is a free data retrieval call binding the contract method 0x9292caaf.
//
// Solidity: function whitelistStartTime() view returns(uint256)
func (_Bacon *BaconCaller) WhitelistStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bacon.contract.Call(opts, &out, "whitelistStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WhitelistStartTime is a free data retrieval call binding the contract method 0x9292caaf.
//
// Solidity: function whitelistStartTime() view returns(uint256)
func (_Bacon *BaconSession) WhitelistStartTime() (*big.Int, error) {
	return _Bacon.Contract.WhitelistStartTime(&_Bacon.CallOpts)
}

// WhitelistStartTime is a free data retrieval call binding the contract method 0x9292caaf.
//
// Solidity: function whitelistStartTime() view returns(uint256)
func (_Bacon *BaconCallerSession) WhitelistStartTime() (*big.Int, error) {
	return _Bacon.Contract.WhitelistStartTime(&_Bacon.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Bacon *BaconTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Bacon *BaconSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bacon.Contract.Approve(&_Bacon.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Bacon *BaconTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bacon.Contract.Approve(&_Bacon.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Bacon *BaconTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Bacon *BaconSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Bacon.Contract.Burn(&_Bacon.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Bacon *BaconTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Bacon.Contract.Burn(&_Bacon.TransactOpts, tokenId)
}

// ChangePricePerToken is a paid mutator transaction binding the contract method 0x2e143909.
//
// Solidity: function changePricePerToken(uint256 newPrice) returns()
func (_Bacon *BaconTransactor) ChangePricePerToken(opts *bind.TransactOpts, newPrice *big.Int) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "changePricePerToken", newPrice)
}

// ChangePricePerToken is a paid mutator transaction binding the contract method 0x2e143909.
//
// Solidity: function changePricePerToken(uint256 newPrice) returns()
func (_Bacon *BaconSession) ChangePricePerToken(newPrice *big.Int) (*types.Transaction, error) {
	return _Bacon.Contract.ChangePricePerToken(&_Bacon.TransactOpts, newPrice)
}

// ChangePricePerToken is a paid mutator transaction binding the contract method 0x2e143909.
//
// Solidity: function changePricePerToken(uint256 newPrice) returns()
func (_Bacon *BaconTransactorSession) ChangePricePerToken(newPrice *big.Int) (*types.Transaction, error) {
	return _Bacon.Contract.ChangePricePerToken(&_Bacon.TransactOpts, newPrice)
}

// LockMetadata is a paid mutator transaction binding the contract method 0x989bdbb6.
//
// Solidity: function lockMetadata() returns()
func (_Bacon *BaconTransactor) LockMetadata(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "lockMetadata")
}

// LockMetadata is a paid mutator transaction binding the contract method 0x989bdbb6.
//
// Solidity: function lockMetadata() returns()
func (_Bacon *BaconSession) LockMetadata() (*types.Transaction, error) {
	return _Bacon.Contract.LockMetadata(&_Bacon.TransactOpts)
}

// LockMetadata is a paid mutator transaction binding the contract method 0x989bdbb6.
//
// Solidity: function lockMetadata() returns()
func (_Bacon *BaconTransactorSession) LockMetadata() (*types.Transaction, error) {
	return _Bacon.Contract.LockMetadata(&_Bacon.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 qty, uint256 _seed) payable returns()
func (_Bacon *BaconTransactor) Mint(opts *bind.TransactOpts, qty *big.Int, _seed *big.Int) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "mint", qty, _seed)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 qty, uint256 _seed) payable returns()
func (_Bacon *BaconSession) Mint(qty *big.Int, _seed *big.Int) (*types.Transaction, error) {
	return _Bacon.Contract.Mint(&_Bacon.TransactOpts, qty, _seed)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 qty, uint256 _seed) payable returns()
func (_Bacon *BaconTransactorSession) Mint(qty *big.Int, _seed *big.Int) (*types.Transaction, error) {
	return _Bacon.Contract.Mint(&_Bacon.TransactOpts, qty, _seed)
}

// ReclaimERC20 is a paid mutator transaction binding the contract method 0x8905fd4f.
//
// Solidity: function reclaimERC20(address erc20Token) returns()
func (_Bacon *BaconTransactor) ReclaimERC20(opts *bind.TransactOpts, erc20Token common.Address) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "reclaimERC20", erc20Token)
}

// ReclaimERC20 is a paid mutator transaction binding the contract method 0x8905fd4f.
//
// Solidity: function reclaimERC20(address erc20Token) returns()
func (_Bacon *BaconSession) ReclaimERC20(erc20Token common.Address) (*types.Transaction, error) {
	return _Bacon.Contract.ReclaimERC20(&_Bacon.TransactOpts, erc20Token)
}

// ReclaimERC20 is a paid mutator transaction binding the contract method 0x8905fd4f.
//
// Solidity: function reclaimERC20(address erc20Token) returns()
func (_Bacon *BaconTransactorSession) ReclaimERC20(erc20Token common.Address) (*types.Transaction, error) {
	return _Bacon.Contract.ReclaimERC20(&_Bacon.TransactOpts, erc20Token)
}

// ReclaimERC721 is a paid mutator transaction binding the contract method 0x6b7d2470.
//
// Solidity: function reclaimERC721(address erc721Token, uint256 id) returns()
func (_Bacon *BaconTransactor) ReclaimERC721(opts *bind.TransactOpts, erc721Token common.Address, id *big.Int) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "reclaimERC721", erc721Token, id)
}

// ReclaimERC721 is a paid mutator transaction binding the contract method 0x6b7d2470.
//
// Solidity: function reclaimERC721(address erc721Token, uint256 id) returns()
func (_Bacon *BaconSession) ReclaimERC721(erc721Token common.Address, id *big.Int) (*types.Transaction, error) {
	return _Bacon.Contract.ReclaimERC721(&_Bacon.TransactOpts, erc721Token, id)
}

// ReclaimERC721 is a paid mutator transaction binding the contract method 0x6b7d2470.
//
// Solidity: function reclaimERC721(address erc721Token, uint256 id) returns()
func (_Bacon *BaconTransactorSession) ReclaimERC721(erc721Token common.Address, id *big.Int) (*types.Transaction, error) {
	return _Bacon.Contract.ReclaimERC721(&_Bacon.TransactOpts, erc721Token, id)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address account) returns()
func (_Bacon *BaconTransactor) Release(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "release", account)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address account) returns()
func (_Bacon *BaconSession) Release(account common.Address) (*types.Transaction, error) {
	return _Bacon.Contract.Release(&_Bacon.TransactOpts, account)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address account) returns()
func (_Bacon *BaconTransactorSession) Release(account common.Address) (*types.Transaction, error) {
	return _Bacon.Contract.Release(&_Bacon.TransactOpts, account)
}

// Release0 is a paid mutator transaction binding the contract method 0x48b75044.
//
// Solidity: function release(address token, address account) returns()
func (_Bacon *BaconTransactor) Release0(opts *bind.TransactOpts, token common.Address, account common.Address) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "release0", token, account)
}

// Release0 is a paid mutator transaction binding the contract method 0x48b75044.
//
// Solidity: function release(address token, address account) returns()
func (_Bacon *BaconSession) Release0(token common.Address, account common.Address) (*types.Transaction, error) {
	return _Bacon.Contract.Release0(&_Bacon.TransactOpts, token, account)
}

// Release0 is a paid mutator transaction binding the contract method 0x48b75044.
//
// Solidity: function release(address token, address account) returns()
func (_Bacon *BaconTransactorSession) Release0(token common.Address, account common.Address) (*types.Transaction, error) {
	return _Bacon.Contract.Release0(&_Bacon.TransactOpts, token, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bacon *BaconTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bacon *BaconSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bacon.Contract.RenounceOwnership(&_Bacon.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bacon *BaconTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bacon.Contract.RenounceOwnership(&_Bacon.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Bacon *BaconTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Bacon *BaconSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bacon.Contract.SafeTransferFrom(&_Bacon.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Bacon *BaconTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bacon.Contract.SafeTransferFrom(&_Bacon.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Bacon *BaconTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Bacon *BaconSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Bacon.Contract.SafeTransferFrom0(&_Bacon.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Bacon *BaconTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Bacon.Contract.SafeTransferFrom0(&_Bacon.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Bacon *BaconTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Bacon *BaconSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Bacon.Contract.SetApprovalForAll(&_Bacon.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Bacon *BaconTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Bacon.Contract.SetApprovalForAll(&_Bacon.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Bacon *BaconTransactor) SetBaseURI(opts *bind.TransactOpts, newBaseURI string) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "setBaseURI", newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Bacon *BaconSession) SetBaseURI(newBaseURI string) (*types.Transaction, error) {
	return _Bacon.Contract.SetBaseURI(&_Bacon.TransactOpts, newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Bacon *BaconTransactorSession) SetBaseURI(newBaseURI string) (*types.Transaction, error) {
	return _Bacon.Contract.SetBaseURI(&_Bacon.TransactOpts, newBaseURI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string newuri) returns()
func (_Bacon *BaconTransactor) SetContractURI(opts *bind.TransactOpts, newuri string) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "setContractURI", newuri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string newuri) returns()
func (_Bacon *BaconSession) SetContractURI(newuri string) (*types.Transaction, error) {
	return _Bacon.Contract.SetContractURI(&_Bacon.TransactOpts, newuri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string newuri) returns()
func (_Bacon *BaconTransactorSession) SetContractURI(newuri string) (*types.Transaction, error) {
	return _Bacon.Contract.SetContractURI(&_Bacon.TransactOpts, newuri)
}

// SetFryingPan is a paid mutator transaction binding the contract method 0x07bd46b8.
//
// Solidity: function setFryingPan(address _fryingPanAddress) returns()
func (_Bacon *BaconTransactor) SetFryingPan(opts *bind.TransactOpts, _fryingPanAddress common.Address) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "setFryingPan", _fryingPanAddress)
}

// SetFryingPan is a paid mutator transaction binding the contract method 0x07bd46b8.
//
// Solidity: function setFryingPan(address _fryingPanAddress) returns()
func (_Bacon *BaconSession) SetFryingPan(_fryingPanAddress common.Address) (*types.Transaction, error) {
	return _Bacon.Contract.SetFryingPan(&_Bacon.TransactOpts, _fryingPanAddress)
}

// SetFryingPan is a paid mutator transaction binding the contract method 0x07bd46b8.
//
// Solidity: function setFryingPan(address _fryingPanAddress) returns()
func (_Bacon *BaconTransactorSession) SetFryingPan(_fryingPanAddress common.Address) (*types.Transaction, error) {
	return _Bacon.Contract.SetFryingPan(&_Bacon.TransactOpts, _fryingPanAddress)
}

// SetGrease is a paid mutator transaction binding the contract method 0xb11d451c.
//
// Solidity: function setGrease(address _newGreaseAddress) returns()
func (_Bacon *BaconTransactor) SetGrease(opts *bind.TransactOpts, _newGreaseAddress common.Address) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "setGrease", _newGreaseAddress)
}

// SetGrease is a paid mutator transaction binding the contract method 0xb11d451c.
//
// Solidity: function setGrease(address _newGreaseAddress) returns()
func (_Bacon *BaconSession) SetGrease(_newGreaseAddress common.Address) (*types.Transaction, error) {
	return _Bacon.Contract.SetGrease(&_Bacon.TransactOpts, _newGreaseAddress)
}

// SetGrease is a paid mutator transaction binding the contract method 0xb11d451c.
//
// Solidity: function setGrease(address _newGreaseAddress) returns()
func (_Bacon *BaconTransactorSession) SetGrease(_newGreaseAddress common.Address) (*types.Transaction, error) {
	return _Bacon.Contract.SetGrease(&_Bacon.TransactOpts, _newGreaseAddress)
}

// SetMaxAmount is a paid mutator transaction binding the contract method 0x4fe47f70.
//
// Solidity: function setMaxAmount(uint256 newAmount) returns()
func (_Bacon *BaconTransactor) SetMaxAmount(opts *bind.TransactOpts, newAmount *big.Int) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "setMaxAmount", newAmount)
}

// SetMaxAmount is a paid mutator transaction binding the contract method 0x4fe47f70.
//
// Solidity: function setMaxAmount(uint256 newAmount) returns()
func (_Bacon *BaconSession) SetMaxAmount(newAmount *big.Int) (*types.Transaction, error) {
	return _Bacon.Contract.SetMaxAmount(&_Bacon.TransactOpts, newAmount)
}

// SetMaxAmount is a paid mutator transaction binding the contract method 0x4fe47f70.
//
// Solidity: function setMaxAmount(uint256 newAmount) returns()
func (_Bacon *BaconTransactorSession) SetMaxAmount(newAmount *big.Int) (*types.Transaction, error) {
	return _Bacon.Contract.SetMaxAmount(&_Bacon.TransactOpts, newAmount)
}

// SetMaxMintQty is a paid mutator transaction binding the contract method 0x10ec7374.
//
// Solidity: function setMaxMintQty(uint16 newQty) returns()
func (_Bacon *BaconTransactor) SetMaxMintQty(opts *bind.TransactOpts, newQty uint16) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "setMaxMintQty", newQty)
}

// SetMaxMintQty is a paid mutator transaction binding the contract method 0x10ec7374.
//
// Solidity: function setMaxMintQty(uint16 newQty) returns()
func (_Bacon *BaconSession) SetMaxMintQty(newQty uint16) (*types.Transaction, error) {
	return _Bacon.Contract.SetMaxMintQty(&_Bacon.TransactOpts, newQty)
}

// SetMaxMintQty is a paid mutator transaction binding the contract method 0x10ec7374.
//
// Solidity: function setMaxMintQty(uint16 newQty) returns()
func (_Bacon *BaconTransactorSession) SetMaxMintQty(newQty uint16) (*types.Transaction, error) {
	return _Bacon.Contract.SetMaxMintQty(&_Bacon.TransactOpts, newQty)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _root) returns()
func (_Bacon *BaconTransactor) SetMerkleRoot(opts *bind.TransactOpts, _root [32]byte) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "setMerkleRoot", _root)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _root) returns()
func (_Bacon *BaconSession) SetMerkleRoot(_root [32]byte) (*types.Transaction, error) {
	return _Bacon.Contract.SetMerkleRoot(&_Bacon.TransactOpts, _root)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _root) returns()
func (_Bacon *BaconTransactorSession) SetMerkleRoot(_root [32]byte) (*types.Transaction, error) {
	return _Bacon.Contract.SetMerkleRoot(&_Bacon.TransactOpts, _root)
}

// SetPaidTokensAmount is a paid mutator transaction binding the contract method 0x3dc94d9c.
//
// Solidity: function setPaidTokensAmount(uint256 newAmount) returns()
func (_Bacon *BaconTransactor) SetPaidTokensAmount(opts *bind.TransactOpts, newAmount *big.Int) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "setPaidTokensAmount", newAmount)
}

// SetPaidTokensAmount is a paid mutator transaction binding the contract method 0x3dc94d9c.
//
// Solidity: function setPaidTokensAmount(uint256 newAmount) returns()
func (_Bacon *BaconSession) SetPaidTokensAmount(newAmount *big.Int) (*types.Transaction, error) {
	return _Bacon.Contract.SetPaidTokensAmount(&_Bacon.TransactOpts, newAmount)
}

// SetPaidTokensAmount is a paid mutator transaction binding the contract method 0x3dc94d9c.
//
// Solidity: function setPaidTokensAmount(uint256 newAmount) returns()
func (_Bacon *BaconTransactorSession) SetPaidTokensAmount(newAmount *big.Int) (*types.Transaction, error) {
	return _Bacon.Contract.SetPaidTokensAmount(&_Bacon.TransactOpts, newAmount)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _setPaused) returns()
func (_Bacon *BaconTransactor) SetPaused(opts *bind.TransactOpts, _setPaused bool) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "setPaused", _setPaused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _setPaused) returns()
func (_Bacon *BaconSession) SetPaused(_setPaused bool) (*types.Transaction, error) {
	return _Bacon.Contract.SetPaused(&_Bacon.TransactOpts, _setPaused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _setPaused) returns()
func (_Bacon *BaconTransactorSession) SetPaused(_setPaused bool) (*types.Transaction, error) {
	return _Bacon.Contract.SetPaused(&_Bacon.TransactOpts, _setPaused)
}

// SetPublicSaleStartTime is a paid mutator transaction binding the contract method 0x6d5d40c6.
//
// Solidity: function setPublicSaleStartTime(uint256 newTime) returns()
func (_Bacon *BaconTransactor) SetPublicSaleStartTime(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "setPublicSaleStartTime", newTime)
}

// SetPublicSaleStartTime is a paid mutator transaction binding the contract method 0x6d5d40c6.
//
// Solidity: function setPublicSaleStartTime(uint256 newTime) returns()
func (_Bacon *BaconSession) SetPublicSaleStartTime(newTime *big.Int) (*types.Transaction, error) {
	return _Bacon.Contract.SetPublicSaleStartTime(&_Bacon.TransactOpts, newTime)
}

// SetPublicSaleStartTime is a paid mutator transaction binding the contract method 0x6d5d40c6.
//
// Solidity: function setPublicSaleStartTime(uint256 newTime) returns()
func (_Bacon *BaconTransactorSession) SetPublicSaleStartTime(newTime *big.Int) (*types.Transaction, error) {
	return _Bacon.Contract.SetPublicSaleStartTime(&_Bacon.TransactOpts, newTime)
}

// SetRandomizer is a paid mutator transaction binding the contract method 0x767bcab5.
//
// Solidity: function setRandomizer(address _newRandomizer) returns()
func (_Bacon *BaconTransactor) SetRandomizer(opts *bind.TransactOpts, _newRandomizer common.Address) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "setRandomizer", _newRandomizer)
}

// SetRandomizer is a paid mutator transaction binding the contract method 0x767bcab5.
//
// Solidity: function setRandomizer(address _newRandomizer) returns()
func (_Bacon *BaconSession) SetRandomizer(_newRandomizer common.Address) (*types.Transaction, error) {
	return _Bacon.Contract.SetRandomizer(&_Bacon.TransactOpts, _newRandomizer)
}

// SetRandomizer is a paid mutator transaction binding the contract method 0x767bcab5.
//
// Solidity: function setRandomizer(address _newRandomizer) returns()
func (_Bacon *BaconTransactorSession) SetRandomizer(_newRandomizer common.Address) (*types.Transaction, error) {
	return _Bacon.Contract.SetRandomizer(&_Bacon.TransactOpts, _newRandomizer)
}

// SetWhitelistStartTime is a paid mutator transaction binding the contract method 0x1c0ce3d3.
//
// Solidity: function setWhitelistStartTime(uint256 newTime) returns()
func (_Bacon *BaconTransactor) SetWhitelistStartTime(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "setWhitelistStartTime", newTime)
}

// SetWhitelistStartTime is a paid mutator transaction binding the contract method 0x1c0ce3d3.
//
// Solidity: function setWhitelistStartTime(uint256 newTime) returns()
func (_Bacon *BaconSession) SetWhitelistStartTime(newTime *big.Int) (*types.Transaction, error) {
	return _Bacon.Contract.SetWhitelistStartTime(&_Bacon.TransactOpts, newTime)
}

// SetWhitelistStartTime is a paid mutator transaction binding the contract method 0x1c0ce3d3.
//
// Solidity: function setWhitelistStartTime(uint256 newTime) returns()
func (_Bacon *BaconTransactorSession) SetWhitelistStartTime(newTime *big.Int) (*types.Transaction, error) {
	return _Bacon.Contract.SetWhitelistStartTime(&_Bacon.TransactOpts, newTime)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Bacon *BaconTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Bacon *BaconSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bacon.Contract.TransferFrom(&_Bacon.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Bacon *BaconTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bacon.Contract.TransferFrom(&_Bacon.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bacon *BaconTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bacon *BaconSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bacon.Contract.TransferOwnership(&_Bacon.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bacon *BaconTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bacon.Contract.TransferOwnership(&_Bacon.TransactOpts, newOwner)
}

// WhitelistMint is a paid mutator transaction binding the contract method 0xdf725396.
//
// Solidity: function whitelistMint(uint256 qty, uint256 tokenId, uint256 _seed, bytes32[] proof) payable returns()
func (_Bacon *BaconTransactor) WhitelistMint(opts *bind.TransactOpts, qty *big.Int, tokenId *big.Int, _seed *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Bacon.contract.Transact(opts, "whitelistMint", qty, tokenId, _seed, proof)
}

// WhitelistMint is a paid mutator transaction binding the contract method 0xdf725396.
//
// Solidity: function whitelistMint(uint256 qty, uint256 tokenId, uint256 _seed, bytes32[] proof) payable returns()
func (_Bacon *BaconSession) WhitelistMint(qty *big.Int, tokenId *big.Int, _seed *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Bacon.Contract.WhitelistMint(&_Bacon.TransactOpts, qty, tokenId, _seed, proof)
}

// WhitelistMint is a paid mutator transaction binding the contract method 0xdf725396.
//
// Solidity: function whitelistMint(uint256 qty, uint256 tokenId, uint256 _seed, bytes32[] proof) payable returns()
func (_Bacon *BaconTransactorSession) WhitelistMint(qty *big.Int, tokenId *big.Int, _seed *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Bacon.Contract.WhitelistMint(&_Bacon.TransactOpts, qty, tokenId, _seed, proof)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bacon *BaconTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bacon.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bacon *BaconSession) Receive() (*types.Transaction, error) {
	return _Bacon.Contract.Receive(&_Bacon.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bacon *BaconTransactorSession) Receive() (*types.Transaction, error) {
	return _Bacon.Contract.Receive(&_Bacon.TransactOpts)
}

// BaconApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Bacon contract.
type BaconApprovalIterator struct {
	Event *BaconApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaconApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaconApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BaconApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BaconApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaconApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaconApproval represents a Approval event raised by the Bacon contract.
type BaconApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Bacon *BaconFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BaconApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bacon.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BaconApprovalIterator{contract: _Bacon.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Bacon *BaconFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BaconApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bacon.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaconApproval)
				if err := _Bacon.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Bacon *BaconFilterer) ParseApproval(log types.Log) (*BaconApproval, error) {
	event := new(BaconApproval)
	if err := _Bacon.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaconApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Bacon contract.
type BaconApprovalForAllIterator struct {
	Event *BaconApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaconApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaconApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BaconApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BaconApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaconApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaconApprovalForAll represents a ApprovalForAll event raised by the Bacon contract.
type BaconApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Bacon *BaconFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BaconApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Bacon.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BaconApprovalForAllIterator{contract: _Bacon.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Bacon *BaconFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BaconApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Bacon.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaconApprovalForAll)
				if err := _Bacon.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Bacon *BaconFilterer) ParseApprovalForAll(log types.Log) (*BaconApprovalForAll, error) {
	event := new(BaconApprovalForAll)
	if err := _Bacon.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaconBaconBurnedIterator is returned from FilterBaconBurned and is used to iterate over the raw logs and unpacked data for BaconBurned events raised by the Bacon contract.
type BaconBaconBurnedIterator struct {
	Event *BaconBaconBurned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaconBaconBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaconBaconBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BaconBaconBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BaconBaconBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaconBaconBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaconBaconBurned represents a BaconBurned event raised by the Bacon contract.
type BaconBaconBurned struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBaconBurned is a free log retrieval operation binding the contract event 0x7d1aa707e55ca3b5e94639805bad011741adb6ad933fc9fc487ef5a6df89cea1.
//
// Solidity: event BaconBurned(uint256 indexed tokenId)
func (_Bacon *BaconFilterer) FilterBaconBurned(opts *bind.FilterOpts, tokenId []*big.Int) (*BaconBaconBurnedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bacon.contract.FilterLogs(opts, "StonedApeBurned", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BaconBaconBurnedIterator{contract: _Bacon.contract, event: "StonedApeBurned", logs: logs, sub: sub}, nil
}

// WatchBaconBurned is a free log subscription operation binding the contract event 0x7d1aa707e55ca3b5e94639805bad011741adb6ad933fc9fc487ef5a6df89cea1.
//
// Solidity: event BaconBurned(uint256 indexed tokenId)
func (_Bacon *BaconFilterer) WatchBaconBurned(opts *bind.WatchOpts, sink chan<- *BaconBaconBurned, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bacon.contract.WatchLogs(opts, "StonedApeBurned", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaconBaconBurned)
				if err := _Bacon.contract.UnpackLog(event, "StonedApeBurned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBaconBurned is a log parse operation binding the contract event 0x7d1aa707e55ca3b5e94639805bad011741adb6ad933fc9fc487ef5a6df89cea1.
//
// Solidity: event BaconBurned(uint256 indexed tokenId)
func (_Bacon *BaconFilterer) ParseBaconBurned(log types.Log) (*BaconBaconBurned, error) {
	event := new(BaconBaconBurned)
	if err := _Bacon.contract.UnpackLog(event, "StonedApeBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaconBaconMintedIterator is returned from FilterBaconMinted and is used to iterate over the raw logs and unpacked data for BaconMinted events raised by the Bacon contract.
type BaconBaconMintedIterator struct {
	Event *BaconBaconMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaconBaconMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaconBaconMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BaconBaconMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BaconBaconMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaconBaconMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaconBaconMinted represents a BaconMinted event raised by the Bacon contract.
type BaconBaconMinted struct {
	TokenId *big.Int
	Minter  common.Address
	Owner   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBaconMinted is a free log retrieval operation binding the contract event 0x015eb83cb5a996a49b6fd3ca1cb8d4506550dab2c65cb737d3200d5228bab0dc.
//
// Solidity: event BaconMinted(uint256 indexed tokenId, address indexed minter, address indexed owner)
func (_Bacon *BaconFilterer) FilterBaconMinted(opts *bind.FilterOpts, tokenId []*big.Int, minter []common.Address, owner []common.Address) (*BaconBaconMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Bacon.contract.FilterLogs(opts, "StonedApeMinted", tokenIdRule, minterRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &BaconBaconMintedIterator{contract: _Bacon.contract, event: "StonedApeMinted", logs: logs, sub: sub}, nil
}

// WatchBaconMinted is a free log subscription operation binding the contract event 0x015eb83cb5a996a49b6fd3ca1cb8d4506550dab2c65cb737d3200d5228bab0dc.
//
// Solidity: event BaconMinted(uint256 indexed tokenId, address indexed minter, address indexed owner)
func (_Bacon *BaconFilterer) WatchBaconMinted(opts *bind.WatchOpts, sink chan<- *BaconBaconMinted, tokenId []*big.Int, minter []common.Address, owner []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Bacon.contract.WatchLogs(opts, "StonedApeMinted", tokenIdRule, minterRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaconBaconMinted)
				if err := _Bacon.contract.UnpackLog(event, "StonedApeMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBaconMinted is a log parse operation binding the contract event 0x015eb83cb5a996a49b6fd3ca1cb8d4506550dab2c65cb737d3200d5228bab0dc.
//
// Solidity: event BaconMinted(uint256 indexed tokenId, address indexed minter, address indexed owner)
func (_Bacon *BaconFilterer) ParseBaconMinted(log types.Log) (*BaconBaconMinted, error) {
	event := new(BaconBaconMinted)
	if err := _Bacon.contract.UnpackLog(event, "StonedApeMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaconCopBurnedIterator is returned from FilterCopBurned and is used to iterate over the raw logs and unpacked data for CopBurned events raised by the Bacon contract.
type BaconCopBurnedIterator struct {
	Event *BaconCopBurned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaconCopBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaconCopBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BaconCopBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BaconCopBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaconCopBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaconCopBurned represents a CopBurned event raised by the Bacon contract.
type BaconCopBurned struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCopBurned is a free log retrieval operation binding the contract event 0x82c0e805aac3366c2ae2b33a8710e4c1a91f98b567ab7e83433e4e1c3485a381.
//
// Solidity: event CopBurned(uint256 indexed tokenId)
func (_Bacon *BaconFilterer) FilterCopBurned(opts *bind.FilterOpts, tokenId []*big.Int) (*BaconCopBurnedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bacon.contract.FilterLogs(opts, "FedApeBurned", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BaconCopBurnedIterator{contract: _Bacon.contract, event: "FedApeBurned", logs: logs, sub: sub}, nil
}

// WatchCopBurned is a free log subscription operation binding the contract event 0x82c0e805aac3366c2ae2b33a8710e4c1a91f98b567ab7e83433e4e1c3485a381.
//
// Solidity: event CopBurned(uint256 indexed tokenId)
func (_Bacon *BaconFilterer) WatchCopBurned(opts *bind.WatchOpts, sink chan<- *BaconCopBurned, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bacon.contract.WatchLogs(opts, "FedApeBurned", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaconCopBurned)
				if err := _Bacon.contract.UnpackLog(event, "FedApeBurned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCopBurned is a log parse operation binding the contract event 0x82c0e805aac3366c2ae2b33a8710e4c1a91f98b567ab7e83433e4e1c3485a381.
//
// Solidity: event CopBurned(uint256 indexed tokenId)
func (_Bacon *BaconFilterer) ParseCopBurned(log types.Log) (*BaconCopBurned, error) {
	event := new(BaconCopBurned)
	if err := _Bacon.contract.UnpackLog(event, "FedApeBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaconCopMintedIterator is returned from FilterCopMinted and is used to iterate over the raw logs and unpacked data for CopMinted events raised by the Bacon contract.
type BaconCopMintedIterator struct {
	Event *BaconCopMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaconCopMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaconCopMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BaconCopMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BaconCopMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaconCopMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaconCopMinted represents a CopMinted event raised by the Bacon contract.
type BaconCopMinted struct {
	TokenId *big.Int
	Minter  common.Address
	Owner   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCopMinted is a free log retrieval operation binding the contract event 0xd21c09fbb86df21368ed1d801f185395d31ffbd7e85fa70ad051b366ef3f27e1.
//
// Solidity: event CopMinted(uint256 indexed tokenId, address indexed minter, address indexed owner)
func (_Bacon *BaconFilterer) FilterCopMinted(opts *bind.FilterOpts, tokenId []*big.Int, minter []common.Address, owner []common.Address) (*BaconCopMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Bacon.contract.FilterLogs(opts, "FedApeMinted", tokenIdRule, minterRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &BaconCopMintedIterator{contract: _Bacon.contract, event: "FedApeMinted", logs: logs, sub: sub}, nil
}

// WatchCopMinted is a free log subscription operation binding the contract event 0xd21c09fbb86df21368ed1d801f185395d31ffbd7e85fa70ad051b366ef3f27e1.
//
// Solidity: event CopMinted(uint256 indexed tokenId, address indexed minter, address indexed owner)
func (_Bacon *BaconFilterer) WatchCopMinted(opts *bind.WatchOpts, sink chan<- *BaconCopMinted, tokenId []*big.Int, minter []common.Address, owner []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Bacon.contract.WatchLogs(opts, "FedApeMinted", tokenIdRule, minterRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaconCopMinted)
				if err := _Bacon.contract.UnpackLog(event, "FedApeMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCopMinted is a log parse operation binding the contract event 0xd21c09fbb86df21368ed1d801f185395d31ffbd7e85fa70ad051b366ef3f27e1.
//
// Solidity: event CopMinted(uint256 indexed tokenId, address indexed minter, address indexed owner)
func (_Bacon *BaconFilterer) ParseCopMinted(log types.Log) (*BaconCopMinted, error) {
	event := new(BaconCopMinted)
	if err := _Bacon.contract.UnpackLog(event, "FedApeMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaconERC20PaymentReleasedIterator is returned from FilterERC20PaymentReleased and is used to iterate over the raw logs and unpacked data for ERC20PaymentReleased events raised by the Bacon contract.
type BaconERC20PaymentReleasedIterator struct {
	Event *BaconERC20PaymentReleased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaconERC20PaymentReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaconERC20PaymentReleased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BaconERC20PaymentReleased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BaconERC20PaymentReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaconERC20PaymentReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaconERC20PaymentReleased represents a ERC20PaymentReleased event raised by the Bacon contract.
type BaconERC20PaymentReleased struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterERC20PaymentReleased is a free log retrieval operation binding the contract event 0x3be5b7a71e84ed12875d241991c70855ac5817d847039e17a9d895c1ceb0f18a.
//
// Solidity: event ERC20PaymentReleased(address indexed token, address to, uint256 amount)
func (_Bacon *BaconFilterer) FilterERC20PaymentReleased(opts *bind.FilterOpts, token []common.Address) (*BaconERC20PaymentReleasedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bacon.contract.FilterLogs(opts, "ERC20PaymentReleased", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BaconERC20PaymentReleasedIterator{contract: _Bacon.contract, event: "ERC20PaymentReleased", logs: logs, sub: sub}, nil
}

// WatchERC20PaymentReleased is a free log subscription operation binding the contract event 0x3be5b7a71e84ed12875d241991c70855ac5817d847039e17a9d895c1ceb0f18a.
//
// Solidity: event ERC20PaymentReleased(address indexed token, address to, uint256 amount)
func (_Bacon *BaconFilterer) WatchERC20PaymentReleased(opts *bind.WatchOpts, sink chan<- *BaconERC20PaymentReleased, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bacon.contract.WatchLogs(opts, "ERC20PaymentReleased", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaconERC20PaymentReleased)
				if err := _Bacon.contract.UnpackLog(event, "ERC20PaymentReleased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseERC20PaymentReleased is a log parse operation binding the contract event 0x3be5b7a71e84ed12875d241991c70855ac5817d847039e17a9d895c1ceb0f18a.
//
// Solidity: event ERC20PaymentReleased(address indexed token, address to, uint256 amount)
func (_Bacon *BaconFilterer) ParseERC20PaymentReleased(log types.Log) (*BaconERC20PaymentReleased, error) {
	event := new(BaconERC20PaymentReleased)
	if err := _Bacon.contract.UnpackLog(event, "ERC20PaymentReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaconOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bacon contract.
type BaconOwnershipTransferredIterator struct {
	Event *BaconOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaconOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaconOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BaconOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BaconOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaconOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaconOwnershipTransferred represents a OwnershipTransferred event raised by the Bacon contract.
type BaconOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bacon *BaconFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BaconOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bacon.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BaconOwnershipTransferredIterator{contract: _Bacon.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bacon *BaconFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BaconOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bacon.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaconOwnershipTransferred)
				if err := _Bacon.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bacon *BaconFilterer) ParseOwnershipTransferred(log types.Log) (*BaconOwnershipTransferred, error) {
	event := new(BaconOwnershipTransferred)
	if err := _Bacon.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaconPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bacon contract.
type BaconPausedIterator struct {
	Event *BaconPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaconPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaconPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BaconPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BaconPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaconPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaconPaused represents a Paused event raised by the Bacon contract.
type BaconPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bacon *BaconFilterer) FilterPaused(opts *bind.FilterOpts) (*BaconPausedIterator, error) {

	logs, sub, err := _Bacon.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BaconPausedIterator{contract: _Bacon.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bacon *BaconFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BaconPaused) (event.Subscription, error) {

	logs, sub, err := _Bacon.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaconPaused)
				if err := _Bacon.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bacon *BaconFilterer) ParsePaused(log types.Log) (*BaconPaused, error) {
	event := new(BaconPaused)
	if err := _Bacon.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaconPayeeAddedIterator is returned from FilterPayeeAdded and is used to iterate over the raw logs and unpacked data for PayeeAdded events raised by the Bacon contract.
type BaconPayeeAddedIterator struct {
	Event *BaconPayeeAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaconPayeeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaconPayeeAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BaconPayeeAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BaconPayeeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaconPayeeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaconPayeeAdded represents a PayeeAdded event raised by the Bacon contract.
type BaconPayeeAdded struct {
	Account common.Address
	Shares  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPayeeAdded is a free log retrieval operation binding the contract event 0x40c340f65e17194d14ddddb073d3c9f888e3cb52b5aae0c6c7706b4fbc905fac.
//
// Solidity: event PayeeAdded(address account, uint256 shares)
func (_Bacon *BaconFilterer) FilterPayeeAdded(opts *bind.FilterOpts) (*BaconPayeeAddedIterator, error) {

	logs, sub, err := _Bacon.contract.FilterLogs(opts, "PayeeAdded")
	if err != nil {
		return nil, err
	}
	return &BaconPayeeAddedIterator{contract: _Bacon.contract, event: "PayeeAdded", logs: logs, sub: sub}, nil
}

// WatchPayeeAdded is a free log subscription operation binding the contract event 0x40c340f65e17194d14ddddb073d3c9f888e3cb52b5aae0c6c7706b4fbc905fac.
//
// Solidity: event PayeeAdded(address account, uint256 shares)
func (_Bacon *BaconFilterer) WatchPayeeAdded(opts *bind.WatchOpts, sink chan<- *BaconPayeeAdded) (event.Subscription, error) {

	logs, sub, err := _Bacon.contract.WatchLogs(opts, "PayeeAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaconPayeeAdded)
				if err := _Bacon.contract.UnpackLog(event, "PayeeAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePayeeAdded is a log parse operation binding the contract event 0x40c340f65e17194d14ddddb073d3c9f888e3cb52b5aae0c6c7706b4fbc905fac.
//
// Solidity: event PayeeAdded(address account, uint256 shares)
func (_Bacon *BaconFilterer) ParsePayeeAdded(log types.Log) (*BaconPayeeAdded, error) {
	event := new(BaconPayeeAdded)
	if err := _Bacon.contract.UnpackLog(event, "PayeeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaconPaymentReceivedIterator is returned from FilterPaymentReceived and is used to iterate over the raw logs and unpacked data for PaymentReceived events raised by the Bacon contract.
type BaconPaymentReceivedIterator struct {
	Event *BaconPaymentReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaconPaymentReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaconPaymentReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BaconPaymentReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BaconPaymentReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaconPaymentReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaconPaymentReceived represents a PaymentReceived event raised by the Bacon contract.
type BaconPaymentReceived struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaymentReceived is a free log retrieval operation binding the contract event 0x6ef95f06320e7a25a04a175ca677b7052bdd97131872c2192525a629f51be770.
//
// Solidity: event PaymentReceived(address from, uint256 amount)
func (_Bacon *BaconFilterer) FilterPaymentReceived(opts *bind.FilterOpts) (*BaconPaymentReceivedIterator, error) {

	logs, sub, err := _Bacon.contract.FilterLogs(opts, "PaymentReceived")
	if err != nil {
		return nil, err
	}
	return &BaconPaymentReceivedIterator{contract: _Bacon.contract, event: "PaymentReceived", logs: logs, sub: sub}, nil
}

// WatchPaymentReceived is a free log subscription operation binding the contract event 0x6ef95f06320e7a25a04a175ca677b7052bdd97131872c2192525a629f51be770.
//
// Solidity: event PaymentReceived(address from, uint256 amount)
func (_Bacon *BaconFilterer) WatchPaymentReceived(opts *bind.WatchOpts, sink chan<- *BaconPaymentReceived) (event.Subscription, error) {

	logs, sub, err := _Bacon.contract.WatchLogs(opts, "PaymentReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaconPaymentReceived)
				if err := _Bacon.contract.UnpackLog(event, "PaymentReceived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaymentReceived is a log parse operation binding the contract event 0x6ef95f06320e7a25a04a175ca677b7052bdd97131872c2192525a629f51be770.
//
// Solidity: event PaymentReceived(address from, uint256 amount)
func (_Bacon *BaconFilterer) ParsePaymentReceived(log types.Log) (*BaconPaymentReceived, error) {
	event := new(BaconPaymentReceived)
	if err := _Bacon.contract.UnpackLog(event, "PaymentReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaconPaymentReleasedIterator is returned from FilterPaymentReleased and is used to iterate over the raw logs and unpacked data for PaymentReleased events raised by the Bacon contract.
type BaconPaymentReleasedIterator struct {
	Event *BaconPaymentReleased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaconPaymentReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaconPaymentReleased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BaconPaymentReleased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BaconPaymentReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaconPaymentReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaconPaymentReleased represents a PaymentReleased event raised by the Bacon contract.
type BaconPaymentReleased struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaymentReleased is a free log retrieval operation binding the contract event 0xdf20fd1e76bc69d672e4814fafb2c449bba3a5369d8359adf9e05e6fde87b056.
//
// Solidity: event PaymentReleased(address to, uint256 amount)
func (_Bacon *BaconFilterer) FilterPaymentReleased(opts *bind.FilterOpts) (*BaconPaymentReleasedIterator, error) {

	logs, sub, err := _Bacon.contract.FilterLogs(opts, "PaymentReleased")
	if err != nil {
		return nil, err
	}
	return &BaconPaymentReleasedIterator{contract: _Bacon.contract, event: "PaymentReleased", logs: logs, sub: sub}, nil
}

// WatchPaymentReleased is a free log subscription operation binding the contract event 0xdf20fd1e76bc69d672e4814fafb2c449bba3a5369d8359adf9e05e6fde87b056.
//
// Solidity: event PaymentReleased(address to, uint256 amount)
func (_Bacon *BaconFilterer) WatchPaymentReleased(opts *bind.WatchOpts, sink chan<- *BaconPaymentReleased) (event.Subscription, error) {

	logs, sub, err := _Bacon.contract.WatchLogs(opts, "PaymentReleased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaconPaymentReleased)
				if err := _Bacon.contract.UnpackLog(event, "PaymentReleased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaymentReleased is a log parse operation binding the contract event 0xdf20fd1e76bc69d672e4814fafb2c449bba3a5369d8359adf9e05e6fde87b056.
//
// Solidity: event PaymentReleased(address to, uint256 amount)
func (_Bacon *BaconFilterer) ParsePaymentReleased(log types.Log) (*BaconPaymentReleased, error) {
	event := new(BaconPaymentReleased)
	if err := _Bacon.contract.UnpackLog(event, "PaymentReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaconTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Bacon contract.
type BaconTransferIterator struct {
	Event *BaconTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaconTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaconTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BaconTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BaconTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaconTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaconTransfer represents a Transfer event raised by the Bacon contract.
type BaconTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Bacon *BaconFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BaconTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bacon.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BaconTransferIterator{contract: _Bacon.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Bacon *BaconFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BaconTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bacon.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaconTransfer)
				if err := _Bacon.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Bacon *BaconFilterer) ParseTransfer(log types.Log) (*BaconTransfer, error) {
	event := new(BaconTransfer)
	if err := _Bacon.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaconUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Bacon contract.
type BaconUnpausedIterator struct {
	Event *BaconUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaconUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaconUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BaconUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BaconUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaconUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaconUnpaused represents a Unpaused event raised by the Bacon contract.
type BaconUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bacon *BaconFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BaconUnpausedIterator, error) {

	logs, sub, err := _Bacon.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BaconUnpausedIterator{contract: _Bacon.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bacon *BaconFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BaconUnpaused) (event.Subscription, error) {

	logs, sub, err := _Bacon.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaconUnpaused)
				if err := _Bacon.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bacon *BaconFilterer) ParseUnpaused(log types.Log) (*BaconUnpaused, error) {
	event := new(BaconUnpaused)
	if err := _Bacon.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
