// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package thefryingpan

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

// ThefryingpanMetaData contains all meta data concerning the Thefryingpan contract.
var ThefryingpanMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"earned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"unstaked\",\"type\":\"bool\"}],\"name\":\"StonedApeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"earned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"unstaked\",\"type\":\"bool\"}],\"name\":\"FedApeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TokenStaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ALPHA_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DAILY_GREASE_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GREASE_CLAIM_TAX_PERCENTAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GreasePerAlpha\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_GLOBAL_GREASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_ALPHA\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_TO_EXIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Pan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bacon\",\"outputs\":[{\"internalType\":\"contractIBACON\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"unstake\",\"type\":\"bool\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getOwedGrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"owed\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"grease\",\"outputs\":[{\"internalType\":\"contractIGREASE\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isCop\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isCop\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastClaimTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pack\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"packIndices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"}],\"name\":\"randomCopOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"}],\"name\":\"reclaimERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"erc721Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"reclaimERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"rescue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rescueEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bacon\",\"type\":\"address\"}],\"name\":\"setBacon\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newRate\",\"type\":\"uint256\"}],\"name\":\"setDailyGreaseRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_grease\",\"type\":\"address\"}],\"name\":\"setGrease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRandomizer\",\"type\":\"address\"}],\"name\":\"setRandomizer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setRescueEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setStakeStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newRatio\",\"type\":\"uint256\"}],\"name\":\"setalphaRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAlphaStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBaconStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalGREASEEarned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unaccountedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ThefryingpanABI is the input ABI used to generate the binding from.
// Deprecated: Use ThefryingpanMetaData.ABI instead.
var ThefryingpanABI = ThefryingpanMetaData.ABI

// Thefryingpan is an auto generated Go binding around an Ethereum contract.
type Thefryingpan struct {
	ThefryingpanCaller     // Read-only binding to the contract
	ThefryingpanTransactor // Write-only binding to the contract
	ThefryingpanFilterer   // Log filterer for contract events
}

// ThefryingpanCaller is an auto generated read-only Go binding around an Ethereum contract.
type ThefryingpanCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThefryingpanTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ThefryingpanTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThefryingpanFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ThefryingpanFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThefryingpanSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ThefryingpanSession struct {
	Contract     *Thefryingpan     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ThefryingpanCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ThefryingpanCallerSession struct {
	Contract *ThefryingpanCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ThefryingpanTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ThefryingpanTransactorSession struct {
	Contract     *ThefryingpanTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ThefryingpanRaw is an auto generated low-level Go binding around an Ethereum contract.
type ThefryingpanRaw struct {
	Contract *Thefryingpan // Generic contract binding to access the raw methods on
}

// ThefryingpanCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ThefryingpanCallerRaw struct {
	Contract *ThefryingpanCaller // Generic read-only contract binding to access the raw methods on
}

// ThefryingpanTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ThefryingpanTransactorRaw struct {
	Contract *ThefryingpanTransactor // Generic write-only contract binding to access the raw methods on
}

// NewThefryingpan creates a new instance of Thefryingpan, bound to a specific deployed contract.
func NewThefryingpan(address common.Address, backend bind.ContractBackend) (*Thefryingpan, error) {
	contract, err := bindThefryingpan(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Thefryingpan{ThefryingpanCaller: ThefryingpanCaller{contract: contract}, ThefryingpanTransactor: ThefryingpanTransactor{contract: contract}, ThefryingpanFilterer: ThefryingpanFilterer{contract: contract}}, nil
}

// NewThefryingpanCaller creates a new read-only instance of Thefryingpan, bound to a specific deployed contract.
func NewThefryingpanCaller(address common.Address, caller bind.ContractCaller) (*ThefryingpanCaller, error) {
	contract, err := bindThefryingpan(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ThefryingpanCaller{contract: contract}, nil
}

// NewThefryingpanTransactor creates a new write-only instance of Thefryingpan, bound to a specific deployed contract.
func NewThefryingpanTransactor(address common.Address, transactor bind.ContractTransactor) (*ThefryingpanTransactor, error) {
	contract, err := bindThefryingpan(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ThefryingpanTransactor{contract: contract}, nil
}

// NewThefryingpanFilterer creates a new log filterer instance of Thefryingpan, bound to a specific deployed contract.
func NewThefryingpanFilterer(address common.Address, filterer bind.ContractFilterer) (*ThefryingpanFilterer, error) {
	contract, err := bindThefryingpan(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ThefryingpanFilterer{contract: contract}, nil
}

// bindThefryingpan binds a generic wrapper to an already deployed contract.
func bindThefryingpan(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ThefryingpanABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Thefryingpan *ThefryingpanRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Thefryingpan.Contract.ThefryingpanCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Thefryingpan *ThefryingpanRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Thefryingpan.Contract.ThefryingpanTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Thefryingpan *ThefryingpanRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Thefryingpan.Contract.ThefryingpanTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Thefryingpan *ThefryingpanCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Thefryingpan.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Thefryingpan *ThefryingpanTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Thefryingpan.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Thefryingpan *ThefryingpanTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Thefryingpan.Contract.contract.Transact(opts, method, params...)
}

// ALPHARATIO is a free data retrieval call binding the contract method 0x8e3d691b.
//
// Solidity: function ALPHA_RATIO() view returns(uint256)
func (_Thefryingpan *ThefryingpanCaller) ALPHARATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "ALPHA_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ALPHARATIO is a free data retrieval call binding the contract method 0x8e3d691b.
//
// Solidity: function ALPHA_RATIO() view returns(uint256)
func (_Thefryingpan *ThefryingpanSession) ALPHARATIO() (*big.Int, error) {
	return _Thefryingpan.Contract.ALPHARATIO(&_Thefryingpan.CallOpts)
}

// ALPHARATIO is a free data retrieval call binding the contract method 0x8e3d691b.
//
// Solidity: function ALPHA_RATIO() view returns(uint256)
func (_Thefryingpan *ThefryingpanCallerSession) ALPHARATIO() (*big.Int, error) {
	return _Thefryingpan.Contract.ALPHARATIO(&_Thefryingpan.CallOpts)
}

// DAILYGREASERATE is a free data retrieval call binding the contract method 0xeafe9e5e.
//
// Solidity: function DAILY_GREASE_RATE() view returns(uint256)
func (_Thefryingpan *ThefryingpanCaller) DAILYGREASERATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "DAILY_GREASE_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DAILYGREASERATE is a free data retrieval call binding the contract method 0xeafe9e5e.
//
// Solidity: function DAILY_GREASE_RATE() view returns(uint256)
func (_Thefryingpan *ThefryingpanSession) DAILYGREASERATE() (*big.Int, error) {
	return _Thefryingpan.Contract.DAILYGREASERATE(&_Thefryingpan.CallOpts)
}

// DAILYGREASERATE is a free data retrieval call binding the contract method 0xeafe9e5e.
//
// Solidity: function DAILY_GREASE_RATE() view returns(uint256)
func (_Thefryingpan *ThefryingpanCallerSession) DAILYGREASERATE() (*big.Int, error) {
	return _Thefryingpan.Contract.DAILYGREASERATE(&_Thefryingpan.CallOpts)
}

// GREASECLAIMTAXPERCENTAGE is a free data retrieval call binding the contract method 0x258d9bce.
//
// Solidity: function GREASE_CLAIM_TAX_PERCENTAGE() view returns(uint256)
func (_Thefryingpan *ThefryingpanCaller) GREASECLAIMTAXPERCENTAGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "GREASE_CLAIM_TAX_PERCENTAGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GREASECLAIMTAXPERCENTAGE is a free data retrieval call binding the contract method 0x258d9bce.
//
// Solidity: function GREASE_CLAIM_TAX_PERCENTAGE() view returns(uint256)
func (_Thefryingpan *ThefryingpanSession) GREASECLAIMTAXPERCENTAGE() (*big.Int, error) {
	return _Thefryingpan.Contract.GREASECLAIMTAXPERCENTAGE(&_Thefryingpan.CallOpts)
}

// GREASECLAIMTAXPERCENTAGE is a free data retrieval call binding the contract method 0x258d9bce.
//
// Solidity: function GREASE_CLAIM_TAX_PERCENTAGE() view returns(uint256)
func (_Thefryingpan *ThefryingpanCallerSession) GREASECLAIMTAXPERCENTAGE() (*big.Int, error) {
	return _Thefryingpan.Contract.GREASECLAIMTAXPERCENTAGE(&_Thefryingpan.CallOpts)
}

// GreasePerAlpha is a free data retrieval call binding the contract method 0xecdfa7fb.
//
// Solidity: function GreasePerAlpha() view returns(uint256)
func (_Thefryingpan *ThefryingpanCaller) GreasePerAlpha(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "GreasePerAlpha")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GreasePerAlpha is a free data retrieval call binding the contract method 0xecdfa7fb.
//
// Solidity: function GreasePerAlpha() view returns(uint256)
func (_Thefryingpan *ThefryingpanSession) GreasePerAlpha() (*big.Int, error) {
	return _Thefryingpan.Contract.GreasePerAlpha(&_Thefryingpan.CallOpts)
}

// GreasePerAlpha is a free data retrieval call binding the contract method 0xecdfa7fb.
//
// Solidity: function GreasePerAlpha() view returns(uint256)
func (_Thefryingpan *ThefryingpanCallerSession) GreasePerAlpha() (*big.Int, error) {
	return _Thefryingpan.Contract.GreasePerAlpha(&_Thefryingpan.CallOpts)
}

// MAXIMUMGLOBALGREASE is a free data retrieval call binding the contract method 0x75546e5b.
//
// Solidity: function MAXIMUM_GLOBAL_GREASE() view returns(uint256)
func (_Thefryingpan *ThefryingpanCaller) MAXIMUMGLOBALGREASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "MAXIMUM_GLOBAL_GREASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXIMUMGLOBALGREASE is a free data retrieval call binding the contract method 0x75546e5b.
//
// Solidity: function MAXIMUM_GLOBAL_GREASE() view returns(uint256)
func (_Thefryingpan *ThefryingpanSession) MAXIMUMGLOBALGREASE() (*big.Int, error) {
	return _Thefryingpan.Contract.MAXIMUMGLOBALGREASE(&_Thefryingpan.CallOpts)
}

// MAXIMUMGLOBALGREASE is a free data retrieval call binding the contract method 0x75546e5b.
//
// Solidity: function MAXIMUM_GLOBAL_GREASE() view returns(uint256)
func (_Thefryingpan *ThefryingpanCallerSession) MAXIMUMGLOBALGREASE() (*big.Int, error) {
	return _Thefryingpan.Contract.MAXIMUMGLOBALGREASE(&_Thefryingpan.CallOpts)
}

// MAXALPHA is a free data retrieval call binding the contract method 0x201ef2a4.
//
// Solidity: function MAX_ALPHA() view returns(uint8)
func (_Thefryingpan *ThefryingpanCaller) MAXALPHA(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "MAX_ALPHA")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXALPHA is a free data retrieval call binding the contract method 0x201ef2a4.
//
// Solidity: function MAX_ALPHA() view returns(uint8)
func (_Thefryingpan *ThefryingpanSession) MAXALPHA() (uint8, error) {
	return _Thefryingpan.Contract.MAXALPHA(&_Thefryingpan.CallOpts)
}

// MAXALPHA is a free data retrieval call binding the contract method 0x201ef2a4.
//
// Solidity: function MAX_ALPHA() view returns(uint8)
func (_Thefryingpan *ThefryingpanCallerSession) MAXALPHA() (uint8, error) {
	return _Thefryingpan.Contract.MAXALPHA(&_Thefryingpan.CallOpts)
}

// MINIMUMTOEXIT is a free data retrieval call binding the contract method 0x37a386b9.
//
// Solidity: function MINIMUM_TO_EXIT() view returns(uint256)
func (_Thefryingpan *ThefryingpanCaller) MINIMUMTOEXIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "MINIMUM_TO_EXIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMTOEXIT is a free data retrieval call binding the contract method 0x37a386b9.
//
// Solidity: function MINIMUM_TO_EXIT() view returns(uint256)
func (_Thefryingpan *ThefryingpanSession) MINIMUMTOEXIT() (*big.Int, error) {
	return _Thefryingpan.Contract.MINIMUMTOEXIT(&_Thefryingpan.CallOpts)
}

// MINIMUMTOEXIT is a free data retrieval call binding the contract method 0x37a386b9.
//
// Solidity: function MINIMUM_TO_EXIT() view returns(uint256)
func (_Thefryingpan *ThefryingpanCallerSession) MINIMUMTOEXIT() (*big.Int, error) {
	return _Thefryingpan.Contract.MINIMUMTOEXIT(&_Thefryingpan.CallOpts)
}

// Pan is a free data retrieval call binding the contract method 0xc0ab3d72.
//
// Solidity: function Pan(uint256 ) view returns(uint256 tokenId, uint256 value, address owner)
func (_Thefryingpan *ThefryingpanCaller) Pan(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenId *big.Int
	Value   *big.Int
	Owner   common.Address
}, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "Pan", arg0)

	outstruct := new(struct {
		TokenId *big.Int
		Value   *big.Int
		Owner   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Pan is a free data retrieval call binding the contract method 0xc0ab3d72.
//
// Solidity: function Pan(uint256 ) view returns(uint256 tokenId, uint256 value, address owner)
func (_Thefryingpan *ThefryingpanSession) Pan(arg0 *big.Int) (struct {
	TokenId *big.Int
	Value   *big.Int
	Owner   common.Address
}, error) {
	return _Thefryingpan.Contract.Pan(&_Thefryingpan.CallOpts, arg0)
}

// Pan is a free data retrieval call binding the contract method 0xc0ab3d72.
//
// Solidity: function Pan(uint256 ) view returns(uint256 tokenId, uint256 value, address owner)
func (_Thefryingpan *ThefryingpanCallerSession) Pan(arg0 *big.Int) (struct {
	TokenId *big.Int
	Value   *big.Int
	Owner   common.Address
}, error) {
	return _Thefryingpan.Contract.Pan(&_Thefryingpan.CallOpts, arg0)
}

// Bacon is a free data retrieval call binding the contract method 0xcdadcc7b.
//
// Solidity: function bacon() view returns(address)
func (_Thefryingpan *ThefryingpanCaller) Bacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "bacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bacon is a free data retrieval call binding the contract method 0xcdadcc7b.
//
// Solidity: function bacon() view returns(address)
func (_Thefryingpan *ThefryingpanSession) Bacon() (common.Address, error) {
	return _Thefryingpan.Contract.Bacon(&_Thefryingpan.CallOpts)
}

// Bacon is a free data retrieval call binding the contract method 0xcdadcc7b.
//
// Solidity: function bacon() view returns(address)
func (_Thefryingpan *ThefryingpanCallerSession) Bacon() (common.Address, error) {
	return _Thefryingpan.Contract.Bacon(&_Thefryingpan.CallOpts)
}

// GetOwedGrease is a free data retrieval call binding the contract method 0x66ef53cb.
//
// Solidity: function getOwedGrease(uint256 tokenId) view returns(uint256 owed)
func (_Thefryingpan *ThefryingpanCaller) GetOwedGrease(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "getOwedGrease", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOwedGrease is a free data retrieval call binding the contract method 0x66ef53cb.
//
// Solidity: function getOwedGrease(uint256 tokenId) view returns(uint256 owed)
func (_Thefryingpan *ThefryingpanSession) GetOwedGrease(tokenId *big.Int) (*big.Int, error) {
	return _Thefryingpan.Contract.GetOwedGrease(&_Thefryingpan.CallOpts, tokenId)
}

// GetOwedGrease is a free data retrieval call binding the contract method 0x66ef53cb.
//
// Solidity: function getOwedGrease(uint256 tokenId) view returns(uint256 owed)
func (_Thefryingpan *ThefryingpanCallerSession) GetOwedGrease(tokenId *big.Int) (*big.Int, error) {
	return _Thefryingpan.Contract.GetOwedGrease(&_Thefryingpan.CallOpts, tokenId)
}

// Grease is a free data retrieval call binding the contract method 0x701794c4.
//
// Solidity: function grease() view returns(address)
func (_Thefryingpan *ThefryingpanCaller) Grease(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "grease")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Grease is a free data retrieval call binding the contract method 0x701794c4.
//
// Solidity: function grease() view returns(address)
func (_Thefryingpan *ThefryingpanSession) Grease() (common.Address, error) {
	return _Thefryingpan.Contract.Grease(&_Thefryingpan.CallOpts)
}

// Grease is a free data retrieval call binding the contract method 0x701794c4.
//
// Solidity: function grease() view returns(address)
func (_Thefryingpan *ThefryingpanCallerSession) Grease() (common.Address, error) {
	return _Thefryingpan.Contract.Grease(&_Thefryingpan.CallOpts)
}

// IsCop is a free data retrieval call binding the contract method 0x0796895d.
//
// Solidity: function isCop(uint256 tokenId) view returns(bool _isCop)
func (_Thefryingpan *ThefryingpanCaller) IsCop(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "isCop", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCop is a free data retrieval call binding the contract method 0x0796895d.
//
// Solidity: function isCop(uint256 tokenId) view returns(bool _isCop)
func (_Thefryingpan *ThefryingpanSession) IsCop(tokenId *big.Int) (bool, error) {
	return _Thefryingpan.Contract.IsCop(&_Thefryingpan.CallOpts, tokenId)
}

// IsCop is a free data retrieval call binding the contract method 0x0796895d.
//
// Solidity: function isCop(uint256 tokenId) view returns(bool _isCop)
func (_Thefryingpan *ThefryingpanCallerSession) IsCop(tokenId *big.Int) (bool, error) {
	return _Thefryingpan.Contract.IsCop(&_Thefryingpan.CallOpts, tokenId)
}

// LastClaimTimestamp is a free data retrieval call binding the contract method 0x607af397.
//
// Solidity: function lastClaimTimestamp() view returns(uint256)
func (_Thefryingpan *ThefryingpanCaller) LastClaimTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "lastClaimTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastClaimTimestamp is a free data retrieval call binding the contract method 0x607af397.
//
// Solidity: function lastClaimTimestamp() view returns(uint256)
func (_Thefryingpan *ThefryingpanSession) LastClaimTimestamp() (*big.Int, error) {
	return _Thefryingpan.Contract.LastClaimTimestamp(&_Thefryingpan.CallOpts)
}

// LastClaimTimestamp is a free data retrieval call binding the contract method 0x607af397.
//
// Solidity: function lastClaimTimestamp() view returns(uint256)
func (_Thefryingpan *ThefryingpanCallerSession) LastClaimTimestamp() (*big.Int, error) {
	return _Thefryingpan.Contract.LastClaimTimestamp(&_Thefryingpan.CallOpts)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Thefryingpan *ThefryingpanCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Thefryingpan *ThefryingpanSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Thefryingpan.Contract.OnERC721Received(&_Thefryingpan.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Thefryingpan *ThefryingpanCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Thefryingpan.Contract.OnERC721Received(&_Thefryingpan.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Thefryingpan *ThefryingpanCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Thefryingpan *ThefryingpanSession) Owner() (common.Address, error) {
	return _Thefryingpan.Contract.Owner(&_Thefryingpan.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Thefryingpan *ThefryingpanCallerSession) Owner() (common.Address, error) {
	return _Thefryingpan.Contract.Owner(&_Thefryingpan.CallOpts)
}

// Pack is a free data retrieval call binding the contract method 0xdd55fcb3.
//
// Solidity: function pack(uint256 , uint256 ) view returns(uint256 tokenId, uint256 value, address owner)
func (_Thefryingpan *ThefryingpanCaller) Pack(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	TokenId *big.Int
	Value   *big.Int
	Owner   common.Address
}, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "pack", arg0, arg1)

	outstruct := new(struct {
		TokenId *big.Int
		Value   *big.Int
		Owner   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Pack is a free data retrieval call binding the contract method 0xdd55fcb3.
//
// Solidity: function pack(uint256 , uint256 ) view returns(uint256 tokenId, uint256 value, address owner)
func (_Thefryingpan *ThefryingpanSession) Pack(arg0 *big.Int, arg1 *big.Int) (struct {
	TokenId *big.Int
	Value   *big.Int
	Owner   common.Address
}, error) {
	return _Thefryingpan.Contract.Pack(&_Thefryingpan.CallOpts, arg0, arg1)
}

// Pack is a free data retrieval call binding the contract method 0xdd55fcb3.
//
// Solidity: function pack(uint256 , uint256 ) view returns(uint256 tokenId, uint256 value, address owner)
func (_Thefryingpan *ThefryingpanCallerSession) Pack(arg0 *big.Int, arg1 *big.Int) (struct {
	TokenId *big.Int
	Value   *big.Int
	Owner   common.Address
}, error) {
	return _Thefryingpan.Contract.Pack(&_Thefryingpan.CallOpts, arg0, arg1)
}

// PackIndices is a free data retrieval call binding the contract method 0x6f234fb5.
//
// Solidity: function packIndices(uint256 ) view returns(uint256)
func (_Thefryingpan *ThefryingpanCaller) PackIndices(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "packIndices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PackIndices is a free data retrieval call binding the contract method 0x6f234fb5.
//
// Solidity: function packIndices(uint256 ) view returns(uint256)
func (_Thefryingpan *ThefryingpanSession) PackIndices(arg0 *big.Int) (*big.Int, error) {
	return _Thefryingpan.Contract.PackIndices(&_Thefryingpan.CallOpts, arg0)
}

// PackIndices is a free data retrieval call binding the contract method 0x6f234fb5.
//
// Solidity: function packIndices(uint256 ) view returns(uint256)
func (_Thefryingpan *ThefryingpanCallerSession) PackIndices(arg0 *big.Int) (*big.Int, error) {
	return _Thefryingpan.Contract.PackIndices(&_Thefryingpan.CallOpts, arg0)
}

// RandomCopOwner is a free data retrieval call binding the contract method 0x46d4e1e6.
//
// Solidity: function randomCopOwner(uint256 seed) view returns(address)
func (_Thefryingpan *ThefryingpanCaller) RandomCopOwner(opts *bind.CallOpts, seed *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "randomCopOwner", seed)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RandomCopOwner is a free data retrieval call binding the contract method 0x46d4e1e6.
//
// Solidity: function randomCopOwner(uint256 seed) view returns(address)
func (_Thefryingpan *ThefryingpanSession) RandomCopOwner(seed *big.Int) (common.Address, error) {
	return _Thefryingpan.Contract.RandomCopOwner(&_Thefryingpan.CallOpts, seed)
}

// RandomCopOwner is a free data retrieval call binding the contract method 0x46d4e1e6.
//
// Solidity: function randomCopOwner(uint256 seed) view returns(address)
func (_Thefryingpan *ThefryingpanCallerSession) RandomCopOwner(seed *big.Int) (common.Address, error) {
	return _Thefryingpan.Contract.RandomCopOwner(&_Thefryingpan.CallOpts, seed)
}

// RescueEnabled is a free data retrieval call binding the contract method 0x39db714f.
//
// Solidity: function rescueEnabled() view returns(bool)
func (_Thefryingpan *ThefryingpanCaller) RescueEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "rescueEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RescueEnabled is a free data retrieval call binding the contract method 0x39db714f.
//
// Solidity: function rescueEnabled() view returns(bool)
func (_Thefryingpan *ThefryingpanSession) RescueEnabled() (bool, error) {
	return _Thefryingpan.Contract.RescueEnabled(&_Thefryingpan.CallOpts)
}

// RescueEnabled is a free data retrieval call binding the contract method 0x39db714f.
//
// Solidity: function rescueEnabled() view returns(bool)
func (_Thefryingpan *ThefryingpanCallerSession) RescueEnabled() (bool, error) {
	return _Thefryingpan.Contract.RescueEnabled(&_Thefryingpan.CallOpts)
}

// StakeStartTime is a free data retrieval call binding the contract method 0x7419f190.
//
// Solidity: function stakeStartTime() view returns(uint256)
func (_Thefryingpan *ThefryingpanCaller) StakeStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "stakeStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeStartTime is a free data retrieval call binding the contract method 0x7419f190.
//
// Solidity: function stakeStartTime() view returns(uint256)
func (_Thefryingpan *ThefryingpanSession) StakeStartTime() (*big.Int, error) {
	return _Thefryingpan.Contract.StakeStartTime(&_Thefryingpan.CallOpts)
}

// StakeStartTime is a free data retrieval call binding the contract method 0x7419f190.
//
// Solidity: function stakeStartTime() view returns(uint256)
func (_Thefryingpan *ThefryingpanCallerSession) StakeStartTime() (*big.Int, error) {
	return _Thefryingpan.Contract.StakeStartTime(&_Thefryingpan.CallOpts)
}

// TotalAlphaStaked is a free data retrieval call binding the contract method 0x6f257875.
//
// Solidity: function totalAlphaStaked() view returns(uint256)
func (_Thefryingpan *ThefryingpanCaller) TotalAlphaStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "totalAlphaStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAlphaStaked is a free data retrieval call binding the contract method 0x6f257875.
//
// Solidity: function totalAlphaStaked() view returns(uint256)
func (_Thefryingpan *ThefryingpanSession) TotalAlphaStaked() (*big.Int, error) {
	return _Thefryingpan.Contract.TotalAlphaStaked(&_Thefryingpan.CallOpts)
}

// TotalAlphaStaked is a free data retrieval call binding the contract method 0x6f257875.
//
// Solidity: function totalAlphaStaked() view returns(uint256)
func (_Thefryingpan *ThefryingpanCallerSession) TotalAlphaStaked() (*big.Int, error) {
	return _Thefryingpan.Contract.TotalAlphaStaked(&_Thefryingpan.CallOpts)
}

// TotalBaconStaked is a free data retrieval call binding the contract method 0x48f225d5.
//
// Solidity: function totalBaconStaked() view returns(uint256)
func (_Thefryingpan *ThefryingpanCaller) TotalBaconStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "totalBaconStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBaconStaked is a free data retrieval call binding the contract method 0x48f225d5.
//
// Solidity: function totalBaconStaked() view returns(uint256)
func (_Thefryingpan *ThefryingpanSession) TotalBaconStaked() (*big.Int, error) {
	return _Thefryingpan.Contract.TotalBaconStaked(&_Thefryingpan.CallOpts)
}

// TotalBaconStaked is a free data retrieval call binding the contract method 0x48f225d5.
//
// Solidity: function totalBaconStaked() view returns(uint256)
func (_Thefryingpan *ThefryingpanCallerSession) TotalBaconStaked() (*big.Int, error) {
	return _Thefryingpan.Contract.TotalBaconStaked(&_Thefryingpan.CallOpts)
}

// TotalGREASEEarned is a free data retrieval call binding the contract method 0x258fc4eb.
//
// Solidity: function totalGREASEEarned() view returns(uint256)
func (_Thefryingpan *ThefryingpanCaller) TotalGREASEEarned(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "totalGREASEEarned")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalGREASEEarned is a free data retrieval call binding the contract method 0x258fc4eb.
//
// Solidity: function totalGREASEEarned() view returns(uint256)
func (_Thefryingpan *ThefryingpanSession) TotalGREASEEarned() (*big.Int, error) {
	return _Thefryingpan.Contract.TotalGREASEEarned(&_Thefryingpan.CallOpts)
}

// TotalGREASEEarned is a free data retrieval call binding the contract method 0x258fc4eb.
//
// Solidity: function totalGREASEEarned() view returns(uint256)
func (_Thefryingpan *ThefryingpanCallerSession) TotalGREASEEarned() (*big.Int, error) {
	return _Thefryingpan.Contract.TotalGREASEEarned(&_Thefryingpan.CallOpts)
}

// UnaccountedRewards is a free data retrieval call binding the contract method 0xc0c472c0.
//
// Solidity: function unaccountedRewards() view returns(uint256)
func (_Thefryingpan *ThefryingpanCaller) UnaccountedRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Thefryingpan.contract.Call(opts, &out, "unaccountedRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnaccountedRewards is a free data retrieval call binding the contract method 0xc0c472c0.
//
// Solidity: function unaccountedRewards() view returns(uint256)
func (_Thefryingpan *ThefryingpanSession) UnaccountedRewards() (*big.Int, error) {
	return _Thefryingpan.Contract.UnaccountedRewards(&_Thefryingpan.CallOpts)
}

// UnaccountedRewards is a free data retrieval call binding the contract method 0xc0c472c0.
//
// Solidity: function unaccountedRewards() view returns(uint256)
func (_Thefryingpan *ThefryingpanCallerSession) UnaccountedRewards() (*big.Int, error) {
	return _Thefryingpan.Contract.UnaccountedRewards(&_Thefryingpan.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0xc50cda12.
//
// Solidity: function claim(uint256[] tokenIds, bool unstake) returns()
func (_Thefryingpan *ThefryingpanTransactor) Claim(opts *bind.TransactOpts, tokenIds []*big.Int, unstake bool) (*types.Transaction, error) {
	return _Thefryingpan.contract.Transact(opts, "claim", tokenIds, unstake)
}

// Claim is a paid mutator transaction binding the contract method 0xc50cda12.
//
// Solidity: function claim(uint256[] tokenIds, bool unstake) returns()
func (_Thefryingpan *ThefryingpanSession) Claim(tokenIds []*big.Int, unstake bool) (*types.Transaction, error) {
	return _Thefryingpan.Contract.Claim(&_Thefryingpan.TransactOpts, tokenIds, unstake)
}

// Claim is a paid mutator transaction binding the contract method 0xc50cda12.
//
// Solidity: function claim(uint256[] tokenIds, bool unstake) returns()
func (_Thefryingpan *ThefryingpanTransactorSession) Claim(tokenIds []*big.Int, unstake bool) (*types.Transaction, error) {
	return _Thefryingpan.Contract.Claim(&_Thefryingpan.TransactOpts, tokenIds, unstake)
}

// ReclaimERC20 is a paid mutator transaction binding the contract method 0x8905fd4f.
//
// Solidity: function reclaimERC20(address erc20Token) returns()
func (_Thefryingpan *ThefryingpanTransactor) ReclaimERC20(opts *bind.TransactOpts, erc20Token common.Address) (*types.Transaction, error) {
	return _Thefryingpan.contract.Transact(opts, "reclaimERC20", erc20Token)
}

// ReclaimERC20 is a paid mutator transaction binding the contract method 0x8905fd4f.
//
// Solidity: function reclaimERC20(address erc20Token) returns()
func (_Thefryingpan *ThefryingpanSession) ReclaimERC20(erc20Token common.Address) (*types.Transaction, error) {
	return _Thefryingpan.Contract.ReclaimERC20(&_Thefryingpan.TransactOpts, erc20Token)
}

// ReclaimERC20 is a paid mutator transaction binding the contract method 0x8905fd4f.
//
// Solidity: function reclaimERC20(address erc20Token) returns()
func (_Thefryingpan *ThefryingpanTransactorSession) ReclaimERC20(erc20Token common.Address) (*types.Transaction, error) {
	return _Thefryingpan.Contract.ReclaimERC20(&_Thefryingpan.TransactOpts, erc20Token)
}

// ReclaimERC721 is a paid mutator transaction binding the contract method 0x6b7d2470.
//
// Solidity: function reclaimERC721(address erc721Token, uint256 id) returns()
func (_Thefryingpan *ThefryingpanTransactor) ReclaimERC721(opts *bind.TransactOpts, erc721Token common.Address, id *big.Int) (*types.Transaction, error) {
	return _Thefryingpan.contract.Transact(opts, "reclaimERC721", erc721Token, id)
}

// ReclaimERC721 is a paid mutator transaction binding the contract method 0x6b7d2470.
//
// Solidity: function reclaimERC721(address erc721Token, uint256 id) returns()
func (_Thefryingpan *ThefryingpanSession) ReclaimERC721(erc721Token common.Address, id *big.Int) (*types.Transaction, error) {
	return _Thefryingpan.Contract.ReclaimERC721(&_Thefryingpan.TransactOpts, erc721Token, id)
}

// ReclaimERC721 is a paid mutator transaction binding the contract method 0x6b7d2470.
//
// Solidity: function reclaimERC721(address erc721Token, uint256 id) returns()
func (_Thefryingpan *ThefryingpanTransactorSession) ReclaimERC721(erc721Token common.Address, id *big.Int) (*types.Transaction, error) {
	return _Thefryingpan.Contract.ReclaimERC721(&_Thefryingpan.TransactOpts, erc721Token, id)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Thefryingpan *ThefryingpanTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Thefryingpan.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Thefryingpan *ThefryingpanSession) RenounceOwnership() (*types.Transaction, error) {
	return _Thefryingpan.Contract.RenounceOwnership(&_Thefryingpan.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Thefryingpan *ThefryingpanTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Thefryingpan.Contract.RenounceOwnership(&_Thefryingpan.TransactOpts)
}

// Rescue is a paid mutator transaction binding the contract method 0x6a3ef057.
//
// Solidity: function rescue(uint256[] tokenIds) returns()
func (_Thefryingpan *ThefryingpanTransactor) Rescue(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Thefryingpan.contract.Transact(opts, "rescue", tokenIds)
}

// Rescue is a paid mutator transaction binding the contract method 0x6a3ef057.
//
// Solidity: function rescue(uint256[] tokenIds) returns()
func (_Thefryingpan *ThefryingpanSession) Rescue(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Thefryingpan.Contract.Rescue(&_Thefryingpan.TransactOpts, tokenIds)
}

// Rescue is a paid mutator transaction binding the contract method 0x6a3ef057.
//
// Solidity: function rescue(uint256[] tokenIds) returns()
func (_Thefryingpan *ThefryingpanTransactorSession) Rescue(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Thefryingpan.Contract.Rescue(&_Thefryingpan.TransactOpts, tokenIds)
}

// SetBacon is a paid mutator transaction binding the contract method 0xf0903e78.
//
// Solidity: function setBacon(address _bacon) returns()
func (_Thefryingpan *ThefryingpanTransactor) SetBacon(opts *bind.TransactOpts, _bacon common.Address) (*types.Transaction, error) {
	return _Thefryingpan.contract.Transact(opts, "setBacon", _bacon)
}

// SetBacon is a paid mutator transaction binding the contract method 0xf0903e78.
//
// Solidity: function setBacon(address _bacon) returns()
func (_Thefryingpan *ThefryingpanSession) SetBacon(_bacon common.Address) (*types.Transaction, error) {
	return _Thefryingpan.Contract.SetBacon(&_Thefryingpan.TransactOpts, _bacon)
}

// SetBacon is a paid mutator transaction binding the contract method 0xf0903e78.
//
// Solidity: function setBacon(address _bacon) returns()
func (_Thefryingpan *ThefryingpanTransactorSession) SetBacon(_bacon common.Address) (*types.Transaction, error) {
	return _Thefryingpan.Contract.SetBacon(&_Thefryingpan.TransactOpts, _bacon)
}

// SetDailyGreaseRate is a paid mutator transaction binding the contract method 0xb450afae.
//
// Solidity: function setDailyGreaseRate(uint256 _newRate) returns()
func (_Thefryingpan *ThefryingpanTransactor) SetDailyGreaseRate(opts *bind.TransactOpts, _newRate *big.Int) (*types.Transaction, error) {
	return _Thefryingpan.contract.Transact(opts, "setDailyGreaseRate", _newRate)
}

// SetDailyGreaseRate is a paid mutator transaction binding the contract method 0xb450afae.
//
// Solidity: function setDailyGreaseRate(uint256 _newRate) returns()
func (_Thefryingpan *ThefryingpanSession) SetDailyGreaseRate(_newRate *big.Int) (*types.Transaction, error) {
	return _Thefryingpan.Contract.SetDailyGreaseRate(&_Thefryingpan.TransactOpts, _newRate)
}

// SetDailyGreaseRate is a paid mutator transaction binding the contract method 0xb450afae.
//
// Solidity: function setDailyGreaseRate(uint256 _newRate) returns()
func (_Thefryingpan *ThefryingpanTransactorSession) SetDailyGreaseRate(_newRate *big.Int) (*types.Transaction, error) {
	return _Thefryingpan.Contract.SetDailyGreaseRate(&_Thefryingpan.TransactOpts, _newRate)
}

// SetGrease is a paid mutator transaction binding the contract method 0xb11d451c.
//
// Solidity: function setGrease(address _grease) returns()
func (_Thefryingpan *ThefryingpanTransactor) SetGrease(opts *bind.TransactOpts, _grease common.Address) (*types.Transaction, error) {
	return _Thefryingpan.contract.Transact(opts, "setGrease", _grease)
}

// SetGrease is a paid mutator transaction binding the contract method 0xb11d451c.
//
// Solidity: function setGrease(address _grease) returns()
func (_Thefryingpan *ThefryingpanSession) SetGrease(_grease common.Address) (*types.Transaction, error) {
	return _Thefryingpan.Contract.SetGrease(&_Thefryingpan.TransactOpts, _grease)
}

// SetGrease is a paid mutator transaction binding the contract method 0xb11d451c.
//
// Solidity: function setGrease(address _grease) returns()
func (_Thefryingpan *ThefryingpanTransactorSession) SetGrease(_grease common.Address) (*types.Transaction, error) {
	return _Thefryingpan.Contract.SetGrease(&_Thefryingpan.TransactOpts, _grease)
}

// SetRandomizer is a paid mutator transaction binding the contract method 0x767bcab5.
//
// Solidity: function setRandomizer(address _newRandomizer) returns()
func (_Thefryingpan *ThefryingpanTransactor) SetRandomizer(opts *bind.TransactOpts, _newRandomizer common.Address) (*types.Transaction, error) {
	return _Thefryingpan.contract.Transact(opts, "setRandomizer", _newRandomizer)
}

// SetRandomizer is a paid mutator transaction binding the contract method 0x767bcab5.
//
// Solidity: function setRandomizer(address _newRandomizer) returns()
func (_Thefryingpan *ThefryingpanSession) SetRandomizer(_newRandomizer common.Address) (*types.Transaction, error) {
	return _Thefryingpan.Contract.SetRandomizer(&_Thefryingpan.TransactOpts, _newRandomizer)
}

// SetRandomizer is a paid mutator transaction binding the contract method 0x767bcab5.
//
// Solidity: function setRandomizer(address _newRandomizer) returns()
func (_Thefryingpan *ThefryingpanTransactorSession) SetRandomizer(_newRandomizer common.Address) (*types.Transaction, error) {
	return _Thefryingpan.Contract.SetRandomizer(&_Thefryingpan.TransactOpts, _newRandomizer)
}

// SetRescueEnabled is a paid mutator transaction binding the contract method 0x76531008.
//
// Solidity: function setRescueEnabled(bool _enabled) returns()
func (_Thefryingpan *ThefryingpanTransactor) SetRescueEnabled(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _Thefryingpan.contract.Transact(opts, "setRescueEnabled", _enabled)
}

// SetRescueEnabled is a paid mutator transaction binding the contract method 0x76531008.
//
// Solidity: function setRescueEnabled(bool _enabled) returns()
func (_Thefryingpan *ThefryingpanSession) SetRescueEnabled(_enabled bool) (*types.Transaction, error) {
	return _Thefryingpan.Contract.SetRescueEnabled(&_Thefryingpan.TransactOpts, _enabled)
}

// SetRescueEnabled is a paid mutator transaction binding the contract method 0x76531008.
//
// Solidity: function setRescueEnabled(bool _enabled) returns()
func (_Thefryingpan *ThefryingpanTransactorSession) SetRescueEnabled(_enabled bool) (*types.Transaction, error) {
	return _Thefryingpan.Contract.SetRescueEnabled(&_Thefryingpan.TransactOpts, _enabled)
}

// SetStakeStartTime is a paid mutator transaction binding the contract method 0xcc1b63d5.
//
// Solidity: function setStakeStartTime(uint256 newTime) returns()
func (_Thefryingpan *ThefryingpanTransactor) SetStakeStartTime(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error) {
	return _Thefryingpan.contract.Transact(opts, "setStakeStartTime", newTime)
}

// SetStakeStartTime is a paid mutator transaction binding the contract method 0xcc1b63d5.
//
// Solidity: function setStakeStartTime(uint256 newTime) returns()
func (_Thefryingpan *ThefryingpanSession) SetStakeStartTime(newTime *big.Int) (*types.Transaction, error) {
	return _Thefryingpan.Contract.SetStakeStartTime(&_Thefryingpan.TransactOpts, newTime)
}

// SetStakeStartTime is a paid mutator transaction binding the contract method 0xcc1b63d5.
//
// Solidity: function setStakeStartTime(uint256 newTime) returns()
func (_Thefryingpan *ThefryingpanTransactorSession) SetStakeStartTime(newTime *big.Int) (*types.Transaction, error) {
	return _Thefryingpan.Contract.SetStakeStartTime(&_Thefryingpan.TransactOpts, newTime)
}

// SetalphaRatio is a paid mutator transaction binding the contract method 0x4524b659.
//
// Solidity: function setalphaRatio(uint256 _newRatio) returns()
func (_Thefryingpan *ThefryingpanTransactor) SetalphaRatio(opts *bind.TransactOpts, _newRatio *big.Int) (*types.Transaction, error) {
	return _Thefryingpan.contract.Transact(opts, "setalphaRatio", _newRatio)
}

// SetalphaRatio is a paid mutator transaction binding the contract method 0x4524b659.
//
// Solidity: function setalphaRatio(uint256 _newRatio) returns()
func (_Thefryingpan *ThefryingpanSession) SetalphaRatio(_newRatio *big.Int) (*types.Transaction, error) {
	return _Thefryingpan.Contract.SetalphaRatio(&_Thefryingpan.TransactOpts, _newRatio)
}

// SetalphaRatio is a paid mutator transaction binding the contract method 0x4524b659.
//
// Solidity: function setalphaRatio(uint256 _newRatio) returns()
func (_Thefryingpan *ThefryingpanTransactorSession) SetalphaRatio(_newRatio *big.Int) (*types.Transaction, error) {
	return _Thefryingpan.Contract.SetalphaRatio(&_Thefryingpan.TransactOpts, _newRatio)
}

// Stake is a paid mutator transaction binding the contract method 0x0fbf0a93.
//
// Solidity: function stake(uint256[] tokenIds) returns()
func (_Thefryingpan *ThefryingpanTransactor) Stake(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Thefryingpan.contract.Transact(opts, "stake", tokenIds)
}

// Stake is a paid mutator transaction binding the contract method 0x0fbf0a93.
//
// Solidity: function stake(uint256[] tokenIds) returns()
func (_Thefryingpan *ThefryingpanSession) Stake(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Thefryingpan.Contract.Stake(&_Thefryingpan.TransactOpts, tokenIds)
}

// Stake is a paid mutator transaction binding the contract method 0x0fbf0a93.
//
// Solidity: function stake(uint256[] tokenIds) returns()
func (_Thefryingpan *ThefryingpanTransactorSession) Stake(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Thefryingpan.Contract.Stake(&_Thefryingpan.TransactOpts, tokenIds)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Thefryingpan *ThefryingpanTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Thefryingpan.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Thefryingpan *ThefryingpanSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Thefryingpan.Contract.TransferOwnership(&_Thefryingpan.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Thefryingpan *ThefryingpanTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Thefryingpan.Contract.TransferOwnership(&_Thefryingpan.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Thefryingpan *ThefryingpanTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Thefryingpan.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Thefryingpan *ThefryingpanSession) Withdraw() (*types.Transaction, error) {
	return _Thefryingpan.Contract.Withdraw(&_Thefryingpan.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Thefryingpan *ThefryingpanTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Thefryingpan.Contract.Withdraw(&_Thefryingpan.TransactOpts)
}

// ThefryingpanBaconClaimedIterator is returned from FilterBaconClaimed and is used to iterate over the raw logs and unpacked data for BaconClaimed events raised by the Thefryingpan contract.
type ThefryingpanBaconClaimedIterator struct {
	Event *ThefryingpanBaconClaimed // Event containing the contract specifics and raw log

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
func (it *ThefryingpanBaconClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThefryingpanBaconClaimed)
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
		it.Event = new(ThefryingpanBaconClaimed)
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
func (it *ThefryingpanBaconClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThefryingpanBaconClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThefryingpanBaconClaimed represents a BaconClaimed event raised by the Thefryingpan contract.
type ThefryingpanBaconClaimed struct {
	TokenId  *big.Int
	Earned   *big.Int
	Unstaked bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBaconClaimed is a free log retrieval operation binding the contract event 0xe19b8b68d8137c08645ab07f69b3163085c2908b8a726544319224f8fb88383c.
//
// Solidity: event BaconClaimed(uint256 indexed tokenId, uint256 earned, bool unstaked)
func (_Thefryingpan *ThefryingpanFilterer) FilterBaconClaimed(opts *bind.FilterOpts, tokenId []*big.Int) (*ThefryingpanBaconClaimedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Thefryingpan.contract.FilterLogs(opts, "StonedApeClaimed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ThefryingpanBaconClaimedIterator{contract: _Thefryingpan.contract, event: "StonedApeClaimed", logs: logs, sub: sub}, nil
}

// WatchBaconClaimed is a free log subscription operation binding the contract event 0xe19b8b68d8137c08645ab07f69b3163085c2908b8a726544319224f8fb88383c.
//
// Solidity: event BaconClaimed(uint256 indexed tokenId, uint256 earned, bool unstaked)
func (_Thefryingpan *ThefryingpanFilterer) WatchBaconClaimed(opts *bind.WatchOpts, sink chan<- *ThefryingpanBaconClaimed, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Thefryingpan.contract.WatchLogs(opts, "StonedApeClaimed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThefryingpanBaconClaimed)
				if err := _Thefryingpan.contract.UnpackLog(event, "StonedApeClaimed", log); err != nil {
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

// ParseBaconClaimed is a log parse operation binding the contract event 0xe19b8b68d8137c08645ab07f69b3163085c2908b8a726544319224f8fb88383c.
//
// Solidity: event BaconClaimed(uint256 indexed tokenId, uint256 earned, bool unstaked)
func (_Thefryingpan *ThefryingpanFilterer) ParseBaconClaimed(log types.Log) (*ThefryingpanBaconClaimed, error) {
	event := new(ThefryingpanBaconClaimed)
	if err := _Thefryingpan.contract.UnpackLog(event, "StonedApeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThefryingpanBaconCopClaimedIterator is returned from FilterBaconCopClaimed and is used to iterate over the raw logs and unpacked data for BaconCopClaimed events raised by the Thefryingpan contract.
type ThefryingpanBaconCopClaimedIterator struct {
	Event *ThefryingpanBaconCopClaimed // Event containing the contract specifics and raw log

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
func (it *ThefryingpanBaconCopClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThefryingpanBaconCopClaimed)
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
		it.Event = new(ThefryingpanBaconCopClaimed)
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
func (it *ThefryingpanBaconCopClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThefryingpanBaconCopClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThefryingpanBaconCopClaimed represents a BaconCopClaimed event raised by the Thefryingpan contract.
type ThefryingpanBaconCopClaimed struct {
	TokenId  *big.Int
	Earned   *big.Int
	Unstaked bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBaconCopClaimed is a free log retrieval operation binding the contract event 0xb2e9ca63677b55e0d6750b3a53e5dfe5013738f39b9ef62cfd62af970b779ce3.
//
// Solidity: event BaconCopClaimed(uint256 indexed tokenId, uint256 earned, bool unstaked)
func (_Thefryingpan *ThefryingpanFilterer) FilterBaconCopClaimed(opts *bind.FilterOpts, tokenId []*big.Int) (*ThefryingpanBaconCopClaimedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Thefryingpan.contract.FilterLogs(opts, "FedApeClaimed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ThefryingpanBaconCopClaimedIterator{contract: _Thefryingpan.contract, event: "FedApeClaimed", logs: logs, sub: sub}, nil
}

// WatchBaconCopClaimed is a free log subscription operation binding the contract event 0xb2e9ca63677b55e0d6750b3a53e5dfe5013738f39b9ef62cfd62af970b779ce3.
//
// Solidity: event BaconCopClaimed(uint256 indexed tokenId, uint256 earned, bool unstaked)
func (_Thefryingpan *ThefryingpanFilterer) WatchBaconCopClaimed(opts *bind.WatchOpts, sink chan<- *ThefryingpanBaconCopClaimed, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Thefryingpan.contract.WatchLogs(opts, "FedApeClaimed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThefryingpanBaconCopClaimed)
				if err := _Thefryingpan.contract.UnpackLog(event, "FedApeClaimed", log); err != nil {
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

// ParseBaconCopClaimed is a log parse operation binding the contract event 0xb2e9ca63677b55e0d6750b3a53e5dfe5013738f39b9ef62cfd62af970b779ce3.
//
// Solidity: event BaconCopClaimed(uint256 indexed tokenId, uint256 earned, bool unstaked)
func (_Thefryingpan *ThefryingpanFilterer) ParseBaconCopClaimed(log types.Log) (*ThefryingpanBaconCopClaimed, error) {
	event := new(ThefryingpanBaconCopClaimed)
	if err := _Thefryingpan.contract.UnpackLog(event, "FedApeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThefryingpanOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Thefryingpan contract.
type ThefryingpanOwnershipTransferredIterator struct {
	Event *ThefryingpanOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ThefryingpanOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThefryingpanOwnershipTransferred)
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
		it.Event = new(ThefryingpanOwnershipTransferred)
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
func (it *ThefryingpanOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThefryingpanOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThefryingpanOwnershipTransferred represents a OwnershipTransferred event raised by the Thefryingpan contract.
type ThefryingpanOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Thefryingpan *ThefryingpanFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ThefryingpanOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Thefryingpan.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ThefryingpanOwnershipTransferredIterator{contract: _Thefryingpan.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Thefryingpan *ThefryingpanFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ThefryingpanOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Thefryingpan.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThefryingpanOwnershipTransferred)
				if err := _Thefryingpan.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Thefryingpan *ThefryingpanFilterer) ParseOwnershipTransferred(log types.Log) (*ThefryingpanOwnershipTransferred, error) {
	event := new(ThefryingpanOwnershipTransferred)
	if err := _Thefryingpan.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThefryingpanTokenStakedIterator is returned from FilterTokenStaked and is used to iterate over the raw logs and unpacked data for TokenStaked events raised by the Thefryingpan contract.
type ThefryingpanTokenStakedIterator struct {
	Event *ThefryingpanTokenStaked // Event containing the contract specifics and raw log

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
func (it *ThefryingpanTokenStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThefryingpanTokenStaked)
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
		it.Event = new(ThefryingpanTokenStaked)
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
func (it *ThefryingpanTokenStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThefryingpanTokenStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThefryingpanTokenStaked represents a TokenStaked event raised by the Thefryingpan contract.
type ThefryingpanTokenStaked struct {
	Owner   common.Address
	TokenId *big.Int
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenStaked is a free log retrieval operation binding the contract event 0x6173e4d2d9dd52aae0ed37afed3adcf924a490639b759ca93d32dc43366c17d2.
//
// Solidity: event TokenStaked(address indexed owner, uint256 indexed tokenId, uint256 value)
func (_Thefryingpan *ThefryingpanFilterer) FilterTokenStaked(opts *bind.FilterOpts, owner []common.Address, tokenId []*big.Int) (*ThefryingpanTokenStakedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Thefryingpan.contract.FilterLogs(opts, "TokenStaked", ownerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ThefryingpanTokenStakedIterator{contract: _Thefryingpan.contract, event: "TokenStaked", logs: logs, sub: sub}, nil
}

// WatchTokenStaked is a free log subscription operation binding the contract event 0x6173e4d2d9dd52aae0ed37afed3adcf924a490639b759ca93d32dc43366c17d2.
//
// Solidity: event TokenStaked(address indexed owner, uint256 indexed tokenId, uint256 value)
func (_Thefryingpan *ThefryingpanFilterer) WatchTokenStaked(opts *bind.WatchOpts, sink chan<- *ThefryingpanTokenStaked, owner []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Thefryingpan.contract.WatchLogs(opts, "TokenStaked", ownerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThefryingpanTokenStaked)
				if err := _Thefryingpan.contract.UnpackLog(event, "TokenStaked", log); err != nil {
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

// ParseTokenStaked is a log parse operation binding the contract event 0x6173e4d2d9dd52aae0ed37afed3adcf924a490639b759ca93d32dc43366c17d2.
//
// Solidity: event TokenStaked(address indexed owner, uint256 indexed tokenId, uint256 value)
func (_Thefryingpan *ThefryingpanFilterer) ParseTokenStaked(log types.Log) (*ThefryingpanTokenStaked, error) {
	event := new(ThefryingpanTokenStaked)
	if err := _Thefryingpan.contract.UnpackLog(event, "TokenStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
