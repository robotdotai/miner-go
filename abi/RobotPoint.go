// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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
	_ = abi.ConvertType
)

// RobotPointMetaData contains all meta data concerning the RobotPoint contract.
var RobotPointMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_powMineCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_foundationCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_miningLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limitPerMint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_rowa\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originDifficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"name\":\"DifficultyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"foundation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"FoundationMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challenge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"}],\"name\":\"changeDifficulty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"difficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"foundationCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"foundationMintedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitPerMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"mine\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"miningLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"miningTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_foundation\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"mintFoundation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"powMineCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"powMintedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rowa\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RobotPointABI is the input ABI used to generate the binding from.
// Deprecated: Use RobotPointMetaData.ABI instead.
var RobotPointABI = RobotPointMetaData.ABI

// RobotPoint is an auto generated Go binding around an Ethereum contract.
type RobotPoint struct {
	RobotPointCaller     // Read-only binding to the contract
	RobotPointTransactor // Write-only binding to the contract
	RobotPointFilterer   // Log filterer for contract events
}

// RobotPointCaller is an auto generated read-only Go binding around an Ethereum contract.
type RobotPointCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RobotPointTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RobotPointTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RobotPointFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RobotPointFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RobotPointSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RobotPointSession struct {
	Contract     *RobotPoint       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RobotPointCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RobotPointCallerSession struct {
	Contract *RobotPointCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RobotPointTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RobotPointTransactorSession struct {
	Contract     *RobotPointTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RobotPointRaw is an auto generated low-level Go binding around an Ethereum contract.
type RobotPointRaw struct {
	Contract *RobotPoint // Generic contract binding to access the raw methods on
}

// RobotPointCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RobotPointCallerRaw struct {
	Contract *RobotPointCaller // Generic read-only contract binding to access the raw methods on
}

// RobotPointTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RobotPointTransactorRaw struct {
	Contract *RobotPointTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRobotPoint creates a new instance of RobotPoint, bound to a specific deployed contract.
func NewRobotPoint(address common.Address, backend bind.ContractBackend) (*RobotPoint, error) {
	contract, err := bindRobotPoint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RobotPoint{RobotPointCaller: RobotPointCaller{contract: contract}, RobotPointTransactor: RobotPointTransactor{contract: contract}, RobotPointFilterer: RobotPointFilterer{contract: contract}}, nil
}

// NewRobotPointCaller creates a new read-only instance of RobotPoint, bound to a specific deployed contract.
func NewRobotPointCaller(address common.Address, caller bind.ContractCaller) (*RobotPointCaller, error) {
	contract, err := bindRobotPoint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RobotPointCaller{contract: contract}, nil
}

// NewRobotPointTransactor creates a new write-only instance of RobotPoint, bound to a specific deployed contract.
func NewRobotPointTransactor(address common.Address, transactor bind.ContractTransactor) (*RobotPointTransactor, error) {
	contract, err := bindRobotPoint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RobotPointTransactor{contract: contract}, nil
}

// NewRobotPointFilterer creates a new log filterer instance of RobotPoint, bound to a specific deployed contract.
func NewRobotPointFilterer(address common.Address, filterer bind.ContractFilterer) (*RobotPointFilterer, error) {
	contract, err := bindRobotPoint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RobotPointFilterer{contract: contract}, nil
}

// bindRobotPoint binds a generic wrapper to an already deployed contract.
func bindRobotPoint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RobotPointMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RobotPoint *RobotPointRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RobotPoint.Contract.RobotPointCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RobotPoint *RobotPointRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RobotPoint.Contract.RobotPointTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RobotPoint *RobotPointRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RobotPoint.Contract.RobotPointTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RobotPoint *RobotPointCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RobotPoint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RobotPoint *RobotPointTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RobotPoint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RobotPoint *RobotPointTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RobotPoint.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_RobotPoint *RobotPointCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RobotPoint.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_RobotPoint *RobotPointSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _RobotPoint.Contract.Allowance(&_RobotPoint.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_RobotPoint *RobotPointCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _RobotPoint.Contract.Allowance(&_RobotPoint.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_RobotPoint *RobotPointCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RobotPoint.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_RobotPoint *RobotPointSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _RobotPoint.Contract.BalanceOf(&_RobotPoint.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_RobotPoint *RobotPointCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _RobotPoint.Contract.BalanceOf(&_RobotPoint.CallOpts, account)
}

// Challenge is a free data retrieval call binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() view returns(uint256)
func (_RobotPoint *RobotPointCaller) Challenge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotPoint.contract.Call(opts, &out, "challenge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Challenge is a free data retrieval call binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() view returns(uint256)
func (_RobotPoint *RobotPointSession) Challenge() (*big.Int, error) {
	return _RobotPoint.Contract.Challenge(&_RobotPoint.CallOpts)
}

// Challenge is a free data retrieval call binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() view returns(uint256)
func (_RobotPoint *RobotPointCallerSession) Challenge() (*big.Int, error) {
	return _RobotPoint.Contract.Challenge(&_RobotPoint.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RobotPoint *RobotPointCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RobotPoint.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RobotPoint *RobotPointSession) Decimals() (uint8, error) {
	return _RobotPoint.Contract.Decimals(&_RobotPoint.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RobotPoint *RobotPointCallerSession) Decimals() (uint8, error) {
	return _RobotPoint.Contract.Decimals(&_RobotPoint.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_RobotPoint *RobotPointCaller) Difficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotPoint.contract.Call(opts, &out, "difficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_RobotPoint *RobotPointSession) Difficulty() (*big.Int, error) {
	return _RobotPoint.Contract.Difficulty(&_RobotPoint.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_RobotPoint *RobotPointCallerSession) Difficulty() (*big.Int, error) {
	return _RobotPoint.Contract.Difficulty(&_RobotPoint.CallOpts)
}

// FoundationCap is a free data retrieval call binding the contract method 0x0fcb497b.
//
// Solidity: function foundationCap() view returns(uint256)
func (_RobotPoint *RobotPointCaller) FoundationCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotPoint.contract.Call(opts, &out, "foundationCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FoundationCap is a free data retrieval call binding the contract method 0x0fcb497b.
//
// Solidity: function foundationCap() view returns(uint256)
func (_RobotPoint *RobotPointSession) FoundationCap() (*big.Int, error) {
	return _RobotPoint.Contract.FoundationCap(&_RobotPoint.CallOpts)
}

// FoundationCap is a free data retrieval call binding the contract method 0x0fcb497b.
//
// Solidity: function foundationCap() view returns(uint256)
func (_RobotPoint *RobotPointCallerSession) FoundationCap() (*big.Int, error) {
	return _RobotPoint.Contract.FoundationCap(&_RobotPoint.CallOpts)
}

// FoundationMintedAmount is a free data retrieval call binding the contract method 0x07fa404d.
//
// Solidity: function foundationMintedAmount() view returns(uint256)
func (_RobotPoint *RobotPointCaller) FoundationMintedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotPoint.contract.Call(opts, &out, "foundationMintedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FoundationMintedAmount is a free data retrieval call binding the contract method 0x07fa404d.
//
// Solidity: function foundationMintedAmount() view returns(uint256)
func (_RobotPoint *RobotPointSession) FoundationMintedAmount() (*big.Int, error) {
	return _RobotPoint.Contract.FoundationMintedAmount(&_RobotPoint.CallOpts)
}

// FoundationMintedAmount is a free data retrieval call binding the contract method 0x07fa404d.
//
// Solidity: function foundationMintedAmount() view returns(uint256)
func (_RobotPoint *RobotPointCallerSession) FoundationMintedAmount() (*big.Int, error) {
	return _RobotPoint.Contract.FoundationMintedAmount(&_RobotPoint.CallOpts)
}

// LimitPerMint is a free data retrieval call binding the contract method 0xe2ce9f51.
//
// Solidity: function limitPerMint() view returns(uint256)
func (_RobotPoint *RobotPointCaller) LimitPerMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotPoint.contract.Call(opts, &out, "limitPerMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitPerMint is a free data retrieval call binding the contract method 0xe2ce9f51.
//
// Solidity: function limitPerMint() view returns(uint256)
func (_RobotPoint *RobotPointSession) LimitPerMint() (*big.Int, error) {
	return _RobotPoint.Contract.LimitPerMint(&_RobotPoint.CallOpts)
}

// LimitPerMint is a free data retrieval call binding the contract method 0xe2ce9f51.
//
// Solidity: function limitPerMint() view returns(uint256)
func (_RobotPoint *RobotPointCallerSession) LimitPerMint() (*big.Int, error) {
	return _RobotPoint.Contract.LimitPerMint(&_RobotPoint.CallOpts)
}

// MiningLimit is a free data retrieval call binding the contract method 0xc2651503.
//
// Solidity: function miningLimit() view returns(uint256)
func (_RobotPoint *RobotPointCaller) MiningLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotPoint.contract.Call(opts, &out, "miningLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiningLimit is a free data retrieval call binding the contract method 0xc2651503.
//
// Solidity: function miningLimit() view returns(uint256)
func (_RobotPoint *RobotPointSession) MiningLimit() (*big.Int, error) {
	return _RobotPoint.Contract.MiningLimit(&_RobotPoint.CallOpts)
}

// MiningLimit is a free data retrieval call binding the contract method 0xc2651503.
//
// Solidity: function miningLimit() view returns(uint256)
func (_RobotPoint *RobotPointCallerSession) MiningLimit() (*big.Int, error) {
	return _RobotPoint.Contract.MiningLimit(&_RobotPoint.CallOpts)
}

// MiningTimes is a free data retrieval call binding the contract method 0x2719881e.
//
// Solidity: function miningTimes(address ) view returns(uint256)
func (_RobotPoint *RobotPointCaller) MiningTimes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RobotPoint.contract.Call(opts, &out, "miningTimes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiningTimes is a free data retrieval call binding the contract method 0x2719881e.
//
// Solidity: function miningTimes(address ) view returns(uint256)
func (_RobotPoint *RobotPointSession) MiningTimes(arg0 common.Address) (*big.Int, error) {
	return _RobotPoint.Contract.MiningTimes(&_RobotPoint.CallOpts, arg0)
}

// MiningTimes is a free data retrieval call binding the contract method 0x2719881e.
//
// Solidity: function miningTimes(address ) view returns(uint256)
func (_RobotPoint *RobotPointCallerSession) MiningTimes(arg0 common.Address) (*big.Int, error) {
	return _RobotPoint.Contract.MiningTimes(&_RobotPoint.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RobotPoint *RobotPointCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RobotPoint.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RobotPoint *RobotPointSession) Name() (string, error) {
	return _RobotPoint.Contract.Name(&_RobotPoint.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RobotPoint *RobotPointCallerSession) Name() (string, error) {
	return _RobotPoint.Contract.Name(&_RobotPoint.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RobotPoint *RobotPointCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RobotPoint.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RobotPoint *RobotPointSession) Owner() (common.Address, error) {
	return _RobotPoint.Contract.Owner(&_RobotPoint.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RobotPoint *RobotPointCallerSession) Owner() (common.Address, error) {
	return _RobotPoint.Contract.Owner(&_RobotPoint.CallOpts)
}

// PowMineCap is a free data retrieval call binding the contract method 0x5aa6552a.
//
// Solidity: function powMineCap() view returns(uint256)
func (_RobotPoint *RobotPointCaller) PowMineCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotPoint.contract.Call(opts, &out, "powMineCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PowMineCap is a free data retrieval call binding the contract method 0x5aa6552a.
//
// Solidity: function powMineCap() view returns(uint256)
func (_RobotPoint *RobotPointSession) PowMineCap() (*big.Int, error) {
	return _RobotPoint.Contract.PowMineCap(&_RobotPoint.CallOpts)
}

// PowMineCap is a free data retrieval call binding the contract method 0x5aa6552a.
//
// Solidity: function powMineCap() view returns(uint256)
func (_RobotPoint *RobotPointCallerSession) PowMineCap() (*big.Int, error) {
	return _RobotPoint.Contract.PowMineCap(&_RobotPoint.CallOpts)
}

// PowMintedAmount is a free data retrieval call binding the contract method 0xa862b9eb.
//
// Solidity: function powMintedAmount() view returns(uint256)
func (_RobotPoint *RobotPointCaller) PowMintedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotPoint.contract.Call(opts, &out, "powMintedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PowMintedAmount is a free data retrieval call binding the contract method 0xa862b9eb.
//
// Solidity: function powMintedAmount() view returns(uint256)
func (_RobotPoint *RobotPointSession) PowMintedAmount() (*big.Int, error) {
	return _RobotPoint.Contract.PowMintedAmount(&_RobotPoint.CallOpts)
}

// PowMintedAmount is a free data retrieval call binding the contract method 0xa862b9eb.
//
// Solidity: function powMintedAmount() view returns(uint256)
func (_RobotPoint *RobotPointCallerSession) PowMintedAmount() (*big.Int, error) {
	return _RobotPoint.Contract.PowMintedAmount(&_RobotPoint.CallOpts)
}

// Rowa is a free data retrieval call binding the contract method 0xb19d1580.
//
// Solidity: function rowa() view returns(address)
func (_RobotPoint *RobotPointCaller) Rowa(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RobotPoint.contract.Call(opts, &out, "rowa")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rowa is a free data retrieval call binding the contract method 0xb19d1580.
//
// Solidity: function rowa() view returns(address)
func (_RobotPoint *RobotPointSession) Rowa() (common.Address, error) {
	return _RobotPoint.Contract.Rowa(&_RobotPoint.CallOpts)
}

// Rowa is a free data retrieval call binding the contract method 0xb19d1580.
//
// Solidity: function rowa() view returns(address)
func (_RobotPoint *RobotPointCallerSession) Rowa() (common.Address, error) {
	return _RobotPoint.Contract.Rowa(&_RobotPoint.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RobotPoint *RobotPointCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RobotPoint.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RobotPoint *RobotPointSession) Symbol() (string, error) {
	return _RobotPoint.Contract.Symbol(&_RobotPoint.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RobotPoint *RobotPointCallerSession) Symbol() (string, error) {
	return _RobotPoint.Contract.Symbol(&_RobotPoint.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RobotPoint *RobotPointCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotPoint.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RobotPoint *RobotPointSession) TotalSupply() (*big.Int, error) {
	return _RobotPoint.Contract.TotalSupply(&_RobotPoint.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RobotPoint *RobotPointCallerSession) TotalSupply() (*big.Int, error) {
	return _RobotPoint.Contract.TotalSupply(&_RobotPoint.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_RobotPoint *RobotPointTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RobotPoint.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_RobotPoint *RobotPointSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RobotPoint.Contract.Approve(&_RobotPoint.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_RobotPoint *RobotPointTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RobotPoint.Contract.Approve(&_RobotPoint.TransactOpts, spender, amount)
}

// ChangeDifficulty is a paid mutator transaction binding the contract method 0x1a6204af.
//
// Solidity: function changeDifficulty(uint256 _difficulty) returns()
func (_RobotPoint *RobotPointTransactor) ChangeDifficulty(opts *bind.TransactOpts, _difficulty *big.Int) (*types.Transaction, error) {
	return _RobotPoint.contract.Transact(opts, "changeDifficulty", _difficulty)
}

// ChangeDifficulty is a paid mutator transaction binding the contract method 0x1a6204af.
//
// Solidity: function changeDifficulty(uint256 _difficulty) returns()
func (_RobotPoint *RobotPointSession) ChangeDifficulty(_difficulty *big.Int) (*types.Transaction, error) {
	return _RobotPoint.Contract.ChangeDifficulty(&_RobotPoint.TransactOpts, _difficulty)
}

// ChangeDifficulty is a paid mutator transaction binding the contract method 0x1a6204af.
//
// Solidity: function changeDifficulty(uint256 _difficulty) returns()
func (_RobotPoint *RobotPointTransactorSession) ChangeDifficulty(_difficulty *big.Int) (*types.Transaction, error) {
	return _RobotPoint.Contract.ChangeDifficulty(&_RobotPoint.TransactOpts, _difficulty)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_RobotPoint *RobotPointTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _RobotPoint.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_RobotPoint *RobotPointSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _RobotPoint.Contract.DecreaseAllowance(&_RobotPoint.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_RobotPoint *RobotPointTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _RobotPoint.Contract.DecreaseAllowance(&_RobotPoint.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_RobotPoint *RobotPointTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _RobotPoint.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_RobotPoint *RobotPointSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _RobotPoint.Contract.IncreaseAllowance(&_RobotPoint.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_RobotPoint *RobotPointTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _RobotPoint.Contract.IncreaseAllowance(&_RobotPoint.TransactOpts, spender, addedValue)
}

// Mine is a paid mutator transaction binding the contract method 0x071e9503.
//
// Solidity: function mine(uint256 _nonce, uint256 _tokenId) returns()
func (_RobotPoint *RobotPointTransactor) Mine(opts *bind.TransactOpts, _nonce *big.Int, _tokenId *big.Int) (*types.Transaction, error) {
	return _RobotPoint.contract.Transact(opts, "mine", _nonce, _tokenId)
}

// Mine is a paid mutator transaction binding the contract method 0x071e9503.
//
// Solidity: function mine(uint256 _nonce, uint256 _tokenId) returns()
func (_RobotPoint *RobotPointSession) Mine(_nonce *big.Int, _tokenId *big.Int) (*types.Transaction, error) {
	return _RobotPoint.Contract.Mine(&_RobotPoint.TransactOpts, _nonce, _tokenId)
}

// Mine is a paid mutator transaction binding the contract method 0x071e9503.
//
// Solidity: function mine(uint256 _nonce, uint256 _tokenId) returns()
func (_RobotPoint *RobotPointTransactorSession) Mine(_nonce *big.Int, _tokenId *big.Int) (*types.Transaction, error) {
	return _RobotPoint.Contract.Mine(&_RobotPoint.TransactOpts, _nonce, _tokenId)
}

// MintFoundation is a paid mutator transaction binding the contract method 0xd12bdd70.
//
// Solidity: function mintFoundation(address _foundation, uint256 _amount, string reason) returns()
func (_RobotPoint *RobotPointTransactor) MintFoundation(opts *bind.TransactOpts, _foundation common.Address, _amount *big.Int, reason string) (*types.Transaction, error) {
	return _RobotPoint.contract.Transact(opts, "mintFoundation", _foundation, _amount, reason)
}

// MintFoundation is a paid mutator transaction binding the contract method 0xd12bdd70.
//
// Solidity: function mintFoundation(address _foundation, uint256 _amount, string reason) returns()
func (_RobotPoint *RobotPointSession) MintFoundation(_foundation common.Address, _amount *big.Int, reason string) (*types.Transaction, error) {
	return _RobotPoint.Contract.MintFoundation(&_RobotPoint.TransactOpts, _foundation, _amount, reason)
}

// MintFoundation is a paid mutator transaction binding the contract method 0xd12bdd70.
//
// Solidity: function mintFoundation(address _foundation, uint256 _amount, string reason) returns()
func (_RobotPoint *RobotPointTransactorSession) MintFoundation(_foundation common.Address, _amount *big.Int, reason string) (*types.Transaction, error) {
	return _RobotPoint.Contract.MintFoundation(&_RobotPoint.TransactOpts, _foundation, _amount, reason)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RobotPoint *RobotPointTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RobotPoint.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RobotPoint *RobotPointSession) RenounceOwnership() (*types.Transaction, error) {
	return _RobotPoint.Contract.RenounceOwnership(&_RobotPoint.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RobotPoint *RobotPointTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RobotPoint.Contract.RenounceOwnership(&_RobotPoint.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_RobotPoint *RobotPointTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RobotPoint.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_RobotPoint *RobotPointSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RobotPoint.Contract.Transfer(&_RobotPoint.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_RobotPoint *RobotPointTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RobotPoint.Contract.Transfer(&_RobotPoint.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_RobotPoint *RobotPointTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RobotPoint.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_RobotPoint *RobotPointSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RobotPoint.Contract.TransferFrom(&_RobotPoint.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_RobotPoint *RobotPointTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RobotPoint.Contract.TransferFrom(&_RobotPoint.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RobotPoint *RobotPointTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RobotPoint.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RobotPoint *RobotPointSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RobotPoint.Contract.TransferOwnership(&_RobotPoint.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RobotPoint *RobotPointTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RobotPoint.Contract.TransferOwnership(&_RobotPoint.TransactOpts, newOwner)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_RobotPoint *RobotPointTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RobotPoint.contract.Transact(opts, "withdrawToken", _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_RobotPoint *RobotPointSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RobotPoint.Contract.WithdrawToken(&_RobotPoint.TransactOpts, _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_RobotPoint *RobotPointTransactorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RobotPoint.Contract.WithdrawToken(&_RobotPoint.TransactOpts, _token, _account, _amount)
}

// RobotPointApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RobotPoint contract.
type RobotPointApprovalIterator struct {
	Event *RobotPointApproval // Event containing the contract specifics and raw log

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
func (it *RobotPointApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobotPointApproval)
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
		it.Event = new(RobotPointApproval)
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
func (it *RobotPointApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobotPointApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobotPointApproval represents a Approval event raised by the RobotPoint contract.
type RobotPointApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RobotPoint *RobotPointFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*RobotPointApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RobotPoint.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &RobotPointApprovalIterator{contract: _RobotPoint.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RobotPoint *RobotPointFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RobotPointApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RobotPoint.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobotPointApproval)
				if err := _RobotPoint.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RobotPoint *RobotPointFilterer) ParseApproval(log types.Log) (*RobotPointApproval, error) {
	event := new(RobotPointApproval)
	if err := _RobotPoint.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RobotPointDifficultyChangedIterator is returned from FilterDifficultyChanged and is used to iterate over the raw logs and unpacked data for DifficultyChanged events raised by the RobotPoint contract.
type RobotPointDifficultyChangedIterator struct {
	Event *RobotPointDifficultyChanged // Event containing the contract specifics and raw log

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
func (it *RobotPointDifficultyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobotPointDifficultyChanged)
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
		it.Event = new(RobotPointDifficultyChanged)
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
func (it *RobotPointDifficultyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobotPointDifficultyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobotPointDifficultyChanged represents a DifficultyChanged event raised by the RobotPoint contract.
type RobotPointDifficultyChanged struct {
	OriginDifficulty *big.Int
	Difficulty       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDifficultyChanged is a free log retrieval operation binding the contract event 0x935ed35c9eb4d93874341d9aaea64ec36e6fa70b445537b90abbcdc6e3e9bb87.
//
// Solidity: event DifficultyChanged(uint256 originDifficulty, uint256 difficulty)
func (_RobotPoint *RobotPointFilterer) FilterDifficultyChanged(opts *bind.FilterOpts) (*RobotPointDifficultyChangedIterator, error) {

	logs, sub, err := _RobotPoint.contract.FilterLogs(opts, "DifficultyChanged")
	if err != nil {
		return nil, err
	}
	return &RobotPointDifficultyChangedIterator{contract: _RobotPoint.contract, event: "DifficultyChanged", logs: logs, sub: sub}, nil
}

// WatchDifficultyChanged is a free log subscription operation binding the contract event 0x935ed35c9eb4d93874341d9aaea64ec36e6fa70b445537b90abbcdc6e3e9bb87.
//
// Solidity: event DifficultyChanged(uint256 originDifficulty, uint256 difficulty)
func (_RobotPoint *RobotPointFilterer) WatchDifficultyChanged(opts *bind.WatchOpts, sink chan<- *RobotPointDifficultyChanged) (event.Subscription, error) {

	logs, sub, err := _RobotPoint.contract.WatchLogs(opts, "DifficultyChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobotPointDifficultyChanged)
				if err := _RobotPoint.contract.UnpackLog(event, "DifficultyChanged", log); err != nil {
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

// ParseDifficultyChanged is a log parse operation binding the contract event 0x935ed35c9eb4d93874341d9aaea64ec36e6fa70b445537b90abbcdc6e3e9bb87.
//
// Solidity: event DifficultyChanged(uint256 originDifficulty, uint256 difficulty)
func (_RobotPoint *RobotPointFilterer) ParseDifficultyChanged(log types.Log) (*RobotPointDifficultyChanged, error) {
	event := new(RobotPointDifficultyChanged)
	if err := _RobotPoint.contract.UnpackLog(event, "DifficultyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RobotPointFoundationMintedIterator is returned from FilterFoundationMinted and is used to iterate over the raw logs and unpacked data for FoundationMinted events raised by the RobotPoint contract.
type RobotPointFoundationMintedIterator struct {
	Event *RobotPointFoundationMinted // Event containing the contract specifics and raw log

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
func (it *RobotPointFoundationMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobotPointFoundationMinted)
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
		it.Event = new(RobotPointFoundationMinted)
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
func (it *RobotPointFoundationMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobotPointFoundationMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobotPointFoundationMinted represents a FoundationMinted event raised by the RobotPoint contract.
type RobotPointFoundationMinted struct {
	Foundation common.Address
	Amount     *big.Int
	Reason     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFoundationMinted is a free log retrieval operation binding the contract event 0x8ab2777cb42805ec254e8ba0437ce0e96ce42712f1e4c3c944873c2365b780a9.
//
// Solidity: event FoundationMinted(address indexed foundation, uint256 amount, string reason)
func (_RobotPoint *RobotPointFilterer) FilterFoundationMinted(opts *bind.FilterOpts, foundation []common.Address) (*RobotPointFoundationMintedIterator, error) {

	var foundationRule []interface{}
	for _, foundationItem := range foundation {
		foundationRule = append(foundationRule, foundationItem)
	}

	logs, sub, err := _RobotPoint.contract.FilterLogs(opts, "FoundationMinted", foundationRule)
	if err != nil {
		return nil, err
	}
	return &RobotPointFoundationMintedIterator{contract: _RobotPoint.contract, event: "FoundationMinted", logs: logs, sub: sub}, nil
}

// WatchFoundationMinted is a free log subscription operation binding the contract event 0x8ab2777cb42805ec254e8ba0437ce0e96ce42712f1e4c3c944873c2365b780a9.
//
// Solidity: event FoundationMinted(address indexed foundation, uint256 amount, string reason)
func (_RobotPoint *RobotPointFilterer) WatchFoundationMinted(opts *bind.WatchOpts, sink chan<- *RobotPointFoundationMinted, foundation []common.Address) (event.Subscription, error) {

	var foundationRule []interface{}
	for _, foundationItem := range foundation {
		foundationRule = append(foundationRule, foundationItem)
	}

	logs, sub, err := _RobotPoint.contract.WatchLogs(opts, "FoundationMinted", foundationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobotPointFoundationMinted)
				if err := _RobotPoint.contract.UnpackLog(event, "FoundationMinted", log); err != nil {
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

// ParseFoundationMinted is a log parse operation binding the contract event 0x8ab2777cb42805ec254e8ba0437ce0e96ce42712f1e4c3c944873c2365b780a9.
//
// Solidity: event FoundationMinted(address indexed foundation, uint256 amount, string reason)
func (_RobotPoint *RobotPointFilterer) ParseFoundationMinted(log types.Log) (*RobotPointFoundationMinted, error) {
	event := new(RobotPointFoundationMinted)
	if err := _RobotPoint.contract.UnpackLog(event, "FoundationMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RobotPointOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RobotPoint contract.
type RobotPointOwnershipTransferredIterator struct {
	Event *RobotPointOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RobotPointOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobotPointOwnershipTransferred)
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
		it.Event = new(RobotPointOwnershipTransferred)
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
func (it *RobotPointOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobotPointOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobotPointOwnershipTransferred represents a OwnershipTransferred event raised by the RobotPoint contract.
type RobotPointOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RobotPoint *RobotPointFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RobotPointOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RobotPoint.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RobotPointOwnershipTransferredIterator{contract: _RobotPoint.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RobotPoint *RobotPointFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RobotPointOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RobotPoint.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobotPointOwnershipTransferred)
				if err := _RobotPoint.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RobotPoint *RobotPointFilterer) ParseOwnershipTransferred(log types.Log) (*RobotPointOwnershipTransferred, error) {
	event := new(RobotPointOwnershipTransferred)
	if err := _RobotPoint.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RobotPointTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RobotPoint contract.
type RobotPointTransferIterator struct {
	Event *RobotPointTransfer // Event containing the contract specifics and raw log

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
func (it *RobotPointTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobotPointTransfer)
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
		it.Event = new(RobotPointTransfer)
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
func (it *RobotPointTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobotPointTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobotPointTransfer represents a Transfer event raised by the RobotPoint contract.
type RobotPointTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RobotPoint *RobotPointFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RobotPointTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RobotPoint.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RobotPointTransferIterator{contract: _RobotPoint.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RobotPoint *RobotPointFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RobotPointTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RobotPoint.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobotPointTransfer)
				if err := _RobotPoint.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RobotPoint *RobotPointFilterer) ParseTransfer(log types.Log) (*RobotPointTransfer, error) {
	event := new(RobotPointTransfer)
	if err := _RobotPoint.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
