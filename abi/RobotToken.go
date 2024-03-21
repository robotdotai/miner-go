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

// RobotTokenMetaData contains all meta data concerning the RobotToken contract.
var RobotTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_powMineCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockMineCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_presaleCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_foundationCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_miningLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limitPerMint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_presaleDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_presaleUserLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originDifficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"name\":\"DifficultyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"foundation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"FoundationMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ParticipatedPresale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PowMineStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PresaleStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challenge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"}],\"name\":\"changeDifficulty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"claimPresale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"difficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyEnableClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"foundationCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"foundationMintedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitPerMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockMineCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_referral\",\"type\":\"address\"}],\"name\":\"mine\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"minedNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"miningLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"miningTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_foundation\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"mintFoundation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"powEnable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"powMineCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"powMintedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"presaleAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"presaleCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"presaleClaimEnable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"presaleDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"presaleMinted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"presaleShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"presaleStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"presaleUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"presaleUserLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"referralPerMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"}],\"name\":\"startLockMine\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startPowMine\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startPresale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RobotTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use RobotTokenMetaData.ABI instead.
var RobotTokenABI = RobotTokenMetaData.ABI

// RobotToken is an auto generated Go binding around an Ethereum contract.
type RobotToken struct {
	RobotTokenCaller     // Read-only binding to the contract
	RobotTokenTransactor // Write-only binding to the contract
	RobotTokenFilterer   // Log filterer for contract events
}

// RobotTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type RobotTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RobotTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RobotTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RobotTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RobotTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RobotTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RobotTokenSession struct {
	Contract     *RobotToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RobotTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RobotTokenCallerSession struct {
	Contract *RobotTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RobotTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RobotTokenTransactorSession struct {
	Contract     *RobotTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RobotTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type RobotTokenRaw struct {
	Contract *RobotToken // Generic contract binding to access the raw methods on
}

// RobotTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RobotTokenCallerRaw struct {
	Contract *RobotTokenCaller // Generic read-only contract binding to access the raw methods on
}

// RobotTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RobotTokenTransactorRaw struct {
	Contract *RobotTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRobotToken creates a new instance of RobotToken, bound to a specific deployed contract.
func NewRobotToken(address common.Address, backend bind.ContractBackend) (*RobotToken, error) {
	contract, err := bindRobotToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RobotToken{RobotTokenCaller: RobotTokenCaller{contract: contract}, RobotTokenTransactor: RobotTokenTransactor{contract: contract}, RobotTokenFilterer: RobotTokenFilterer{contract: contract}}, nil
}

// NewRobotTokenCaller creates a new read-only instance of RobotToken, bound to a specific deployed contract.
func NewRobotTokenCaller(address common.Address, caller bind.ContractCaller) (*RobotTokenCaller, error) {
	contract, err := bindRobotToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RobotTokenCaller{contract: contract}, nil
}

// NewRobotTokenTransactor creates a new write-only instance of RobotToken, bound to a specific deployed contract.
func NewRobotTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*RobotTokenTransactor, error) {
	contract, err := bindRobotToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RobotTokenTransactor{contract: contract}, nil
}

// NewRobotTokenFilterer creates a new log filterer instance of RobotToken, bound to a specific deployed contract.
func NewRobotTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*RobotTokenFilterer, error) {
	contract, err := bindRobotToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RobotTokenFilterer{contract: contract}, nil
}

// bindRobotToken binds a generic wrapper to an already deployed contract.
func bindRobotToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RobotTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RobotToken *RobotTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RobotToken.Contract.RobotTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RobotToken *RobotTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RobotToken.Contract.RobotTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RobotToken *RobotTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RobotToken.Contract.RobotTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RobotToken *RobotTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RobotToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RobotToken *RobotTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RobotToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RobotToken *RobotTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RobotToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_RobotToken *RobotTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_RobotToken *RobotTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _RobotToken.Contract.Allowance(&_RobotToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_RobotToken *RobotTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _RobotToken.Contract.Allowance(&_RobotToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_RobotToken *RobotTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_RobotToken *RobotTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _RobotToken.Contract.BalanceOf(&_RobotToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_RobotToken *RobotTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _RobotToken.Contract.BalanceOf(&_RobotToken.CallOpts, account)
}

// Challenge is a free data retrieval call binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() view returns(uint256)
func (_RobotToken *RobotTokenCaller) Challenge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "challenge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Challenge is a free data retrieval call binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() view returns(uint256)
func (_RobotToken *RobotTokenSession) Challenge() (*big.Int, error) {
	return _RobotToken.Contract.Challenge(&_RobotToken.CallOpts)
}

// Challenge is a free data retrieval call binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() view returns(uint256)
func (_RobotToken *RobotTokenCallerSession) Challenge() (*big.Int, error) {
	return _RobotToken.Contract.Challenge(&_RobotToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RobotToken *RobotTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RobotToken *RobotTokenSession) Decimals() (uint8, error) {
	return _RobotToken.Contract.Decimals(&_RobotToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RobotToken *RobotTokenCallerSession) Decimals() (uint8, error) {
	return _RobotToken.Contract.Decimals(&_RobotToken.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_RobotToken *RobotTokenCaller) Difficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "difficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_RobotToken *RobotTokenSession) Difficulty() (*big.Int, error) {
	return _RobotToken.Contract.Difficulty(&_RobotToken.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_RobotToken *RobotTokenCallerSession) Difficulty() (*big.Int, error) {
	return _RobotToken.Contract.Difficulty(&_RobotToken.CallOpts)
}

// FoundationCap is a free data retrieval call binding the contract method 0x0fcb497b.
//
// Solidity: function foundationCap() view returns(uint256)
func (_RobotToken *RobotTokenCaller) FoundationCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "foundationCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FoundationCap is a free data retrieval call binding the contract method 0x0fcb497b.
//
// Solidity: function foundationCap() view returns(uint256)
func (_RobotToken *RobotTokenSession) FoundationCap() (*big.Int, error) {
	return _RobotToken.Contract.FoundationCap(&_RobotToken.CallOpts)
}

// FoundationCap is a free data retrieval call binding the contract method 0x0fcb497b.
//
// Solidity: function foundationCap() view returns(uint256)
func (_RobotToken *RobotTokenCallerSession) FoundationCap() (*big.Int, error) {
	return _RobotToken.Contract.FoundationCap(&_RobotToken.CallOpts)
}

// FoundationMintedAmount is a free data retrieval call binding the contract method 0x07fa404d.
//
// Solidity: function foundationMintedAmount() view returns(uint256)
func (_RobotToken *RobotTokenCaller) FoundationMintedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "foundationMintedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FoundationMintedAmount is a free data retrieval call binding the contract method 0x07fa404d.
//
// Solidity: function foundationMintedAmount() view returns(uint256)
func (_RobotToken *RobotTokenSession) FoundationMintedAmount() (*big.Int, error) {
	return _RobotToken.Contract.FoundationMintedAmount(&_RobotToken.CallOpts)
}

// FoundationMintedAmount is a free data retrieval call binding the contract method 0x07fa404d.
//
// Solidity: function foundationMintedAmount() view returns(uint256)
func (_RobotToken *RobotTokenCallerSession) FoundationMintedAmount() (*big.Int, error) {
	return _RobotToken.Contract.FoundationMintedAmount(&_RobotToken.CallOpts)
}

// LimitPerMint is a free data retrieval call binding the contract method 0xe2ce9f51.
//
// Solidity: function limitPerMint() view returns(uint256)
func (_RobotToken *RobotTokenCaller) LimitPerMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "limitPerMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitPerMint is a free data retrieval call binding the contract method 0xe2ce9f51.
//
// Solidity: function limitPerMint() view returns(uint256)
func (_RobotToken *RobotTokenSession) LimitPerMint() (*big.Int, error) {
	return _RobotToken.Contract.LimitPerMint(&_RobotToken.CallOpts)
}

// LimitPerMint is a free data retrieval call binding the contract method 0xe2ce9f51.
//
// Solidity: function limitPerMint() view returns(uint256)
func (_RobotToken *RobotTokenCallerSession) LimitPerMint() (*big.Int, error) {
	return _RobotToken.Contract.LimitPerMint(&_RobotToken.CallOpts)
}

// LockMineCap is a free data retrieval call binding the contract method 0x85e50001.
//
// Solidity: function lockMineCap() view returns(uint256)
func (_RobotToken *RobotTokenCaller) LockMineCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "lockMineCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockMineCap is a free data retrieval call binding the contract method 0x85e50001.
//
// Solidity: function lockMineCap() view returns(uint256)
func (_RobotToken *RobotTokenSession) LockMineCap() (*big.Int, error) {
	return _RobotToken.Contract.LockMineCap(&_RobotToken.CallOpts)
}

// LockMineCap is a free data retrieval call binding the contract method 0x85e50001.
//
// Solidity: function lockMineCap() view returns(uint256)
func (_RobotToken *RobotTokenCallerSession) LockMineCap() (*big.Int, error) {
	return _RobotToken.Contract.LockMineCap(&_RobotToken.CallOpts)
}

// MinedNonces is a free data retrieval call binding the contract method 0x342a252a.
//
// Solidity: function minedNonces(address , uint256 ) view returns(bool)
func (_RobotToken *RobotTokenCaller) MinedNonces(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "minedNonces", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MinedNonces is a free data retrieval call binding the contract method 0x342a252a.
//
// Solidity: function minedNonces(address , uint256 ) view returns(bool)
func (_RobotToken *RobotTokenSession) MinedNonces(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _RobotToken.Contract.MinedNonces(&_RobotToken.CallOpts, arg0, arg1)
}

// MinedNonces is a free data retrieval call binding the contract method 0x342a252a.
//
// Solidity: function minedNonces(address , uint256 ) view returns(bool)
func (_RobotToken *RobotTokenCallerSession) MinedNonces(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _RobotToken.Contract.MinedNonces(&_RobotToken.CallOpts, arg0, arg1)
}

// MiningLimit is a free data retrieval call binding the contract method 0xc2651503.
//
// Solidity: function miningLimit() view returns(uint256)
func (_RobotToken *RobotTokenCaller) MiningLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "miningLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiningLimit is a free data retrieval call binding the contract method 0xc2651503.
//
// Solidity: function miningLimit() view returns(uint256)
func (_RobotToken *RobotTokenSession) MiningLimit() (*big.Int, error) {
	return _RobotToken.Contract.MiningLimit(&_RobotToken.CallOpts)
}

// MiningLimit is a free data retrieval call binding the contract method 0xc2651503.
//
// Solidity: function miningLimit() view returns(uint256)
func (_RobotToken *RobotTokenCallerSession) MiningLimit() (*big.Int, error) {
	return _RobotToken.Contract.MiningLimit(&_RobotToken.CallOpts)
}

// MiningTimes is a free data retrieval call binding the contract method 0x2719881e.
//
// Solidity: function miningTimes(address ) view returns(uint256)
func (_RobotToken *RobotTokenCaller) MiningTimes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "miningTimes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiningTimes is a free data retrieval call binding the contract method 0x2719881e.
//
// Solidity: function miningTimes(address ) view returns(uint256)
func (_RobotToken *RobotTokenSession) MiningTimes(arg0 common.Address) (*big.Int, error) {
	return _RobotToken.Contract.MiningTimes(&_RobotToken.CallOpts, arg0)
}

// MiningTimes is a free data retrieval call binding the contract method 0x2719881e.
//
// Solidity: function miningTimes(address ) view returns(uint256)
func (_RobotToken *RobotTokenCallerSession) MiningTimes(arg0 common.Address) (*big.Int, error) {
	return _RobotToken.Contract.MiningTimes(&_RobotToken.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RobotToken *RobotTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RobotToken *RobotTokenSession) Name() (string, error) {
	return _RobotToken.Contract.Name(&_RobotToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RobotToken *RobotTokenCallerSession) Name() (string, error) {
	return _RobotToken.Contract.Name(&_RobotToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RobotToken *RobotTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RobotToken *RobotTokenSession) Owner() (common.Address, error) {
	return _RobotToken.Contract.Owner(&_RobotToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RobotToken *RobotTokenCallerSession) Owner() (common.Address, error) {
	return _RobotToken.Contract.Owner(&_RobotToken.CallOpts)
}

// PowEnable is a free data retrieval call binding the contract method 0xe25830b2.
//
// Solidity: function powEnable() view returns(bool)
func (_RobotToken *RobotTokenCaller) PowEnable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "powEnable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PowEnable is a free data retrieval call binding the contract method 0xe25830b2.
//
// Solidity: function powEnable() view returns(bool)
func (_RobotToken *RobotTokenSession) PowEnable() (bool, error) {
	return _RobotToken.Contract.PowEnable(&_RobotToken.CallOpts)
}

// PowEnable is a free data retrieval call binding the contract method 0xe25830b2.
//
// Solidity: function powEnable() view returns(bool)
func (_RobotToken *RobotTokenCallerSession) PowEnable() (bool, error) {
	return _RobotToken.Contract.PowEnable(&_RobotToken.CallOpts)
}

// PowMineCap is a free data retrieval call binding the contract method 0x5aa6552a.
//
// Solidity: function powMineCap() view returns(uint256)
func (_RobotToken *RobotTokenCaller) PowMineCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "powMineCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PowMineCap is a free data retrieval call binding the contract method 0x5aa6552a.
//
// Solidity: function powMineCap() view returns(uint256)
func (_RobotToken *RobotTokenSession) PowMineCap() (*big.Int, error) {
	return _RobotToken.Contract.PowMineCap(&_RobotToken.CallOpts)
}

// PowMineCap is a free data retrieval call binding the contract method 0x5aa6552a.
//
// Solidity: function powMineCap() view returns(uint256)
func (_RobotToken *RobotTokenCallerSession) PowMineCap() (*big.Int, error) {
	return _RobotToken.Contract.PowMineCap(&_RobotToken.CallOpts)
}

// PowMintedAmount is a free data retrieval call binding the contract method 0xa862b9eb.
//
// Solidity: function powMintedAmount() view returns(uint256)
func (_RobotToken *RobotTokenCaller) PowMintedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "powMintedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PowMintedAmount is a free data retrieval call binding the contract method 0xa862b9eb.
//
// Solidity: function powMintedAmount() view returns(uint256)
func (_RobotToken *RobotTokenSession) PowMintedAmount() (*big.Int, error) {
	return _RobotToken.Contract.PowMintedAmount(&_RobotToken.CallOpts)
}

// PowMintedAmount is a free data retrieval call binding the contract method 0xa862b9eb.
//
// Solidity: function powMintedAmount() view returns(uint256)
func (_RobotToken *RobotTokenCallerSession) PowMintedAmount() (*big.Int, error) {
	return _RobotToken.Contract.PowMintedAmount(&_RobotToken.CallOpts)
}

// PresaleAmount is a free data retrieval call binding the contract method 0xb138d500.
//
// Solidity: function presaleAmount() view returns(uint256)
func (_RobotToken *RobotTokenCaller) PresaleAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "presaleAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PresaleAmount is a free data retrieval call binding the contract method 0xb138d500.
//
// Solidity: function presaleAmount() view returns(uint256)
func (_RobotToken *RobotTokenSession) PresaleAmount() (*big.Int, error) {
	return _RobotToken.Contract.PresaleAmount(&_RobotToken.CallOpts)
}

// PresaleAmount is a free data retrieval call binding the contract method 0xb138d500.
//
// Solidity: function presaleAmount() view returns(uint256)
func (_RobotToken *RobotTokenCallerSession) PresaleAmount() (*big.Int, error) {
	return _RobotToken.Contract.PresaleAmount(&_RobotToken.CallOpts)
}

// PresaleCap is a free data retrieval call binding the contract method 0x63d5502f.
//
// Solidity: function presaleCap() view returns(uint256)
func (_RobotToken *RobotTokenCaller) PresaleCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "presaleCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PresaleCap is a free data retrieval call binding the contract method 0x63d5502f.
//
// Solidity: function presaleCap() view returns(uint256)
func (_RobotToken *RobotTokenSession) PresaleCap() (*big.Int, error) {
	return _RobotToken.Contract.PresaleCap(&_RobotToken.CallOpts)
}

// PresaleCap is a free data retrieval call binding the contract method 0x63d5502f.
//
// Solidity: function presaleCap() view returns(uint256)
func (_RobotToken *RobotTokenCallerSession) PresaleCap() (*big.Int, error) {
	return _RobotToken.Contract.PresaleCap(&_RobotToken.CallOpts)
}

// PresaleClaimEnable is a free data retrieval call binding the contract method 0xccb843f6.
//
// Solidity: function presaleClaimEnable() view returns(bool)
func (_RobotToken *RobotTokenCaller) PresaleClaimEnable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "presaleClaimEnable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PresaleClaimEnable is a free data retrieval call binding the contract method 0xccb843f6.
//
// Solidity: function presaleClaimEnable() view returns(bool)
func (_RobotToken *RobotTokenSession) PresaleClaimEnable() (bool, error) {
	return _RobotToken.Contract.PresaleClaimEnable(&_RobotToken.CallOpts)
}

// PresaleClaimEnable is a free data retrieval call binding the contract method 0xccb843f6.
//
// Solidity: function presaleClaimEnable() view returns(bool)
func (_RobotToken *RobotTokenCallerSession) PresaleClaimEnable() (bool, error) {
	return _RobotToken.Contract.PresaleClaimEnable(&_RobotToken.CallOpts)
}

// PresaleDuration is a free data retrieval call binding the contract method 0x5868c32a.
//
// Solidity: function presaleDuration() view returns(uint256)
func (_RobotToken *RobotTokenCaller) PresaleDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "presaleDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PresaleDuration is a free data retrieval call binding the contract method 0x5868c32a.
//
// Solidity: function presaleDuration() view returns(uint256)
func (_RobotToken *RobotTokenSession) PresaleDuration() (*big.Int, error) {
	return _RobotToken.Contract.PresaleDuration(&_RobotToken.CallOpts)
}

// PresaleDuration is a free data retrieval call binding the contract method 0x5868c32a.
//
// Solidity: function presaleDuration() view returns(uint256)
func (_RobotToken *RobotTokenCallerSession) PresaleDuration() (*big.Int, error) {
	return _RobotToken.Contract.PresaleDuration(&_RobotToken.CallOpts)
}

// PresaleMinted is a free data retrieval call binding the contract method 0xbc660cac.
//
// Solidity: function presaleMinted(address ) view returns(bool)
func (_RobotToken *RobotTokenCaller) PresaleMinted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "presaleMinted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PresaleMinted is a free data retrieval call binding the contract method 0xbc660cac.
//
// Solidity: function presaleMinted(address ) view returns(bool)
func (_RobotToken *RobotTokenSession) PresaleMinted(arg0 common.Address) (bool, error) {
	return _RobotToken.Contract.PresaleMinted(&_RobotToken.CallOpts, arg0)
}

// PresaleMinted is a free data retrieval call binding the contract method 0xbc660cac.
//
// Solidity: function presaleMinted(address ) view returns(bool)
func (_RobotToken *RobotTokenCallerSession) PresaleMinted(arg0 common.Address) (bool, error) {
	return _RobotToken.Contract.PresaleMinted(&_RobotToken.CallOpts, arg0)
}

// PresaleShare is a free data retrieval call binding the contract method 0x171796fd.
//
// Solidity: function presaleShare(address _user) view returns(uint256)
func (_RobotToken *RobotTokenCaller) PresaleShare(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "presaleShare", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PresaleShare is a free data retrieval call binding the contract method 0x171796fd.
//
// Solidity: function presaleShare(address _user) view returns(uint256)
func (_RobotToken *RobotTokenSession) PresaleShare(_user common.Address) (*big.Int, error) {
	return _RobotToken.Contract.PresaleShare(&_RobotToken.CallOpts, _user)
}

// PresaleShare is a free data retrieval call binding the contract method 0x171796fd.
//
// Solidity: function presaleShare(address _user) view returns(uint256)
func (_RobotToken *RobotTokenCallerSession) PresaleShare(_user common.Address) (*big.Int, error) {
	return _RobotToken.Contract.PresaleShare(&_RobotToken.CallOpts, _user)
}

// PresaleStart is a free data retrieval call binding the contract method 0xde8801e5.
//
// Solidity: function presaleStart() view returns(uint256)
func (_RobotToken *RobotTokenCaller) PresaleStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "presaleStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PresaleStart is a free data retrieval call binding the contract method 0xde8801e5.
//
// Solidity: function presaleStart() view returns(uint256)
func (_RobotToken *RobotTokenSession) PresaleStart() (*big.Int, error) {
	return _RobotToken.Contract.PresaleStart(&_RobotToken.CallOpts)
}

// PresaleStart is a free data retrieval call binding the contract method 0xde8801e5.
//
// Solidity: function presaleStart() view returns(uint256)
func (_RobotToken *RobotTokenCallerSession) PresaleStart() (*big.Int, error) {
	return _RobotToken.Contract.PresaleStart(&_RobotToken.CallOpts)
}

// PresaleUser is a free data retrieval call binding the contract method 0x502a4e87.
//
// Solidity: function presaleUser(address ) view returns(uint256)
func (_RobotToken *RobotTokenCaller) PresaleUser(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "presaleUser", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PresaleUser is a free data retrieval call binding the contract method 0x502a4e87.
//
// Solidity: function presaleUser(address ) view returns(uint256)
func (_RobotToken *RobotTokenSession) PresaleUser(arg0 common.Address) (*big.Int, error) {
	return _RobotToken.Contract.PresaleUser(&_RobotToken.CallOpts, arg0)
}

// PresaleUser is a free data retrieval call binding the contract method 0x502a4e87.
//
// Solidity: function presaleUser(address ) view returns(uint256)
func (_RobotToken *RobotTokenCallerSession) PresaleUser(arg0 common.Address) (*big.Int, error) {
	return _RobotToken.Contract.PresaleUser(&_RobotToken.CallOpts, arg0)
}

// PresaleUserLimit is a free data retrieval call binding the contract method 0x70aa395d.
//
// Solidity: function presaleUserLimit() view returns(uint256)
func (_RobotToken *RobotTokenCaller) PresaleUserLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "presaleUserLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PresaleUserLimit is a free data retrieval call binding the contract method 0x70aa395d.
//
// Solidity: function presaleUserLimit() view returns(uint256)
func (_RobotToken *RobotTokenSession) PresaleUserLimit() (*big.Int, error) {
	return _RobotToken.Contract.PresaleUserLimit(&_RobotToken.CallOpts)
}

// PresaleUserLimit is a free data retrieval call binding the contract method 0x70aa395d.
//
// Solidity: function presaleUserLimit() view returns(uint256)
func (_RobotToken *RobotTokenCallerSession) PresaleUserLimit() (*big.Int, error) {
	return _RobotToken.Contract.PresaleUserLimit(&_RobotToken.CallOpts)
}

// ReferralPerMint is a free data retrieval call binding the contract method 0xc1da3a9f.
//
// Solidity: function referralPerMint() view returns(uint256)
func (_RobotToken *RobotTokenCaller) ReferralPerMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "referralPerMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReferralPerMint is a free data retrieval call binding the contract method 0xc1da3a9f.
//
// Solidity: function referralPerMint() view returns(uint256)
func (_RobotToken *RobotTokenSession) ReferralPerMint() (*big.Int, error) {
	return _RobotToken.Contract.ReferralPerMint(&_RobotToken.CallOpts)
}

// ReferralPerMint is a free data retrieval call binding the contract method 0xc1da3a9f.
//
// Solidity: function referralPerMint() view returns(uint256)
func (_RobotToken *RobotTokenCallerSession) ReferralPerMint() (*big.Int, error) {
	return _RobotToken.Contract.ReferralPerMint(&_RobotToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RobotToken *RobotTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RobotToken *RobotTokenSession) Symbol() (string, error) {
	return _RobotToken.Contract.Symbol(&_RobotToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RobotToken *RobotTokenCallerSession) Symbol() (string, error) {
	return _RobotToken.Contract.Symbol(&_RobotToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RobotToken *RobotTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RobotToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RobotToken *RobotTokenSession) TotalSupply() (*big.Int, error) {
	return _RobotToken.Contract.TotalSupply(&_RobotToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RobotToken *RobotTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _RobotToken.Contract.TotalSupply(&_RobotToken.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe3412e3d.
//
// Solidity: function addLiquidity(address _router) returns()
func (_RobotToken *RobotTokenTransactor) AddLiquidity(opts *bind.TransactOpts, _router common.Address) (*types.Transaction, error) {
	return _RobotToken.contract.Transact(opts, "addLiquidity", _router)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe3412e3d.
//
// Solidity: function addLiquidity(address _router) returns()
func (_RobotToken *RobotTokenSession) AddLiquidity(_router common.Address) (*types.Transaction, error) {
	return _RobotToken.Contract.AddLiquidity(&_RobotToken.TransactOpts, _router)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe3412e3d.
//
// Solidity: function addLiquidity(address _router) returns()
func (_RobotToken *RobotTokenTransactorSession) AddLiquidity(_router common.Address) (*types.Transaction, error) {
	return _RobotToken.Contract.AddLiquidity(&_RobotToken.TransactOpts, _router)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_RobotToken *RobotTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RobotToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_RobotToken *RobotTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RobotToken.Contract.Approve(&_RobotToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_RobotToken *RobotTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RobotToken.Contract.Approve(&_RobotToken.TransactOpts, spender, amount)
}

// ChangeDifficulty is a paid mutator transaction binding the contract method 0x1a6204af.
//
// Solidity: function changeDifficulty(uint256 _difficulty) returns()
func (_RobotToken *RobotTokenTransactor) ChangeDifficulty(opts *bind.TransactOpts, _difficulty *big.Int) (*types.Transaction, error) {
	return _RobotToken.contract.Transact(opts, "changeDifficulty", _difficulty)
}

// ChangeDifficulty is a paid mutator transaction binding the contract method 0x1a6204af.
//
// Solidity: function changeDifficulty(uint256 _difficulty) returns()
func (_RobotToken *RobotTokenSession) ChangeDifficulty(_difficulty *big.Int) (*types.Transaction, error) {
	return _RobotToken.Contract.ChangeDifficulty(&_RobotToken.TransactOpts, _difficulty)
}

// ChangeDifficulty is a paid mutator transaction binding the contract method 0x1a6204af.
//
// Solidity: function changeDifficulty(uint256 _difficulty) returns()
func (_RobotToken *RobotTokenTransactorSession) ChangeDifficulty(_difficulty *big.Int) (*types.Transaction, error) {
	return _RobotToken.Contract.ChangeDifficulty(&_RobotToken.TransactOpts, _difficulty)
}

// ClaimPresale is a paid mutator transaction binding the contract method 0x25bcb9fb.
//
// Solidity: function claimPresale(address _user) returns()
func (_RobotToken *RobotTokenTransactor) ClaimPresale(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _RobotToken.contract.Transact(opts, "claimPresale", _user)
}

// ClaimPresale is a paid mutator transaction binding the contract method 0x25bcb9fb.
//
// Solidity: function claimPresale(address _user) returns()
func (_RobotToken *RobotTokenSession) ClaimPresale(_user common.Address) (*types.Transaction, error) {
	return _RobotToken.Contract.ClaimPresale(&_RobotToken.TransactOpts, _user)
}

// ClaimPresale is a paid mutator transaction binding the contract method 0x25bcb9fb.
//
// Solidity: function claimPresale(address _user) returns()
func (_RobotToken *RobotTokenTransactorSession) ClaimPresale(_user common.Address) (*types.Transaction, error) {
	return _RobotToken.Contract.ClaimPresale(&_RobotToken.TransactOpts, _user)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_RobotToken *RobotTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _RobotToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_RobotToken *RobotTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _RobotToken.Contract.DecreaseAllowance(&_RobotToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_RobotToken *RobotTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _RobotToken.Contract.DecreaseAllowance(&_RobotToken.TransactOpts, spender, subtractedValue)
}

// EmergencyEnableClaim is a paid mutator transaction binding the contract method 0xae2f5ab7.
//
// Solidity: function emergencyEnableClaim() returns()
func (_RobotToken *RobotTokenTransactor) EmergencyEnableClaim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RobotToken.contract.Transact(opts, "emergencyEnableClaim")
}

// EmergencyEnableClaim is a paid mutator transaction binding the contract method 0xae2f5ab7.
//
// Solidity: function emergencyEnableClaim() returns()
func (_RobotToken *RobotTokenSession) EmergencyEnableClaim() (*types.Transaction, error) {
	return _RobotToken.Contract.EmergencyEnableClaim(&_RobotToken.TransactOpts)
}

// EmergencyEnableClaim is a paid mutator transaction binding the contract method 0xae2f5ab7.
//
// Solidity: function emergencyEnableClaim() returns()
func (_RobotToken *RobotTokenTransactorSession) EmergencyEnableClaim() (*types.Transaction, error) {
	return _RobotToken.Contract.EmergencyEnableClaim(&_RobotToken.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_RobotToken *RobotTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _RobotToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_RobotToken *RobotTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _RobotToken.Contract.IncreaseAllowance(&_RobotToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_RobotToken *RobotTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _RobotToken.Contract.IncreaseAllowance(&_RobotToken.TransactOpts, spender, addedValue)
}

// Mine is a paid mutator transaction binding the contract method 0x414da005.
//
// Solidity: function mine(uint256 _nonce, address _referral) returns()
func (_RobotToken *RobotTokenTransactor) Mine(opts *bind.TransactOpts, _nonce *big.Int, _referral common.Address) (*types.Transaction, error) {
	return _RobotToken.contract.Transact(opts, "mine", _nonce, _referral)
}

// Mine is a paid mutator transaction binding the contract method 0x414da005.
//
// Solidity: function mine(uint256 _nonce, address _referral) returns()
func (_RobotToken *RobotTokenSession) Mine(_nonce *big.Int, _referral common.Address) (*types.Transaction, error) {
	return _RobotToken.Contract.Mine(&_RobotToken.TransactOpts, _nonce, _referral)
}

// Mine is a paid mutator transaction binding the contract method 0x414da005.
//
// Solidity: function mine(uint256 _nonce, address _referral) returns()
func (_RobotToken *RobotTokenTransactorSession) Mine(_nonce *big.Int, _referral common.Address) (*types.Transaction, error) {
	return _RobotToken.Contract.Mine(&_RobotToken.TransactOpts, _nonce, _referral)
}

// MintFoundation is a paid mutator transaction binding the contract method 0xd12bdd70.
//
// Solidity: function mintFoundation(address _foundation, uint256 _amount, string reason) returns()
func (_RobotToken *RobotTokenTransactor) MintFoundation(opts *bind.TransactOpts, _foundation common.Address, _amount *big.Int, reason string) (*types.Transaction, error) {
	return _RobotToken.contract.Transact(opts, "mintFoundation", _foundation, _amount, reason)
}

// MintFoundation is a paid mutator transaction binding the contract method 0xd12bdd70.
//
// Solidity: function mintFoundation(address _foundation, uint256 _amount, string reason) returns()
func (_RobotToken *RobotTokenSession) MintFoundation(_foundation common.Address, _amount *big.Int, reason string) (*types.Transaction, error) {
	return _RobotToken.Contract.MintFoundation(&_RobotToken.TransactOpts, _foundation, _amount, reason)
}

// MintFoundation is a paid mutator transaction binding the contract method 0xd12bdd70.
//
// Solidity: function mintFoundation(address _foundation, uint256 _amount, string reason) returns()
func (_RobotToken *RobotTokenTransactorSession) MintFoundation(_foundation common.Address, _amount *big.Int, reason string) (*types.Transaction, error) {
	return _RobotToken.Contract.MintFoundation(&_RobotToken.TransactOpts, _foundation, _amount, reason)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RobotToken *RobotTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RobotToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RobotToken *RobotTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _RobotToken.Contract.RenounceOwnership(&_RobotToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RobotToken *RobotTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RobotToken.Contract.RenounceOwnership(&_RobotToken.TransactOpts)
}

// StartLockMine is a paid mutator transaction binding the contract method 0x32f51e93.
//
// Solidity: function startLockMine(address _minter, uint256 _difficulty) returns()
func (_RobotToken *RobotTokenTransactor) StartLockMine(opts *bind.TransactOpts, _minter common.Address, _difficulty *big.Int) (*types.Transaction, error) {
	return _RobotToken.contract.Transact(opts, "startLockMine", _minter, _difficulty)
}

// StartLockMine is a paid mutator transaction binding the contract method 0x32f51e93.
//
// Solidity: function startLockMine(address _minter, uint256 _difficulty) returns()
func (_RobotToken *RobotTokenSession) StartLockMine(_minter common.Address, _difficulty *big.Int) (*types.Transaction, error) {
	return _RobotToken.Contract.StartLockMine(&_RobotToken.TransactOpts, _minter, _difficulty)
}

// StartLockMine is a paid mutator transaction binding the contract method 0x32f51e93.
//
// Solidity: function startLockMine(address _minter, uint256 _difficulty) returns()
func (_RobotToken *RobotTokenTransactorSession) StartLockMine(_minter common.Address, _difficulty *big.Int) (*types.Transaction, error) {
	return _RobotToken.Contract.StartLockMine(&_RobotToken.TransactOpts, _minter, _difficulty)
}

// StartPowMine is a paid mutator transaction binding the contract method 0xd22271ab.
//
// Solidity: function startPowMine() returns()
func (_RobotToken *RobotTokenTransactor) StartPowMine(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RobotToken.contract.Transact(opts, "startPowMine")
}

// StartPowMine is a paid mutator transaction binding the contract method 0xd22271ab.
//
// Solidity: function startPowMine() returns()
func (_RobotToken *RobotTokenSession) StartPowMine() (*types.Transaction, error) {
	return _RobotToken.Contract.StartPowMine(&_RobotToken.TransactOpts)
}

// StartPowMine is a paid mutator transaction binding the contract method 0xd22271ab.
//
// Solidity: function startPowMine() returns()
func (_RobotToken *RobotTokenTransactorSession) StartPowMine() (*types.Transaction, error) {
	return _RobotToken.Contract.StartPowMine(&_RobotToken.TransactOpts)
}

// StartPresale is a paid mutator transaction binding the contract method 0x04c98b2b.
//
// Solidity: function startPresale() returns()
func (_RobotToken *RobotTokenTransactor) StartPresale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RobotToken.contract.Transact(opts, "startPresale")
}

// StartPresale is a paid mutator transaction binding the contract method 0x04c98b2b.
//
// Solidity: function startPresale() returns()
func (_RobotToken *RobotTokenSession) StartPresale() (*types.Transaction, error) {
	return _RobotToken.Contract.StartPresale(&_RobotToken.TransactOpts)
}

// StartPresale is a paid mutator transaction binding the contract method 0x04c98b2b.
//
// Solidity: function startPresale() returns()
func (_RobotToken *RobotTokenTransactorSession) StartPresale() (*types.Transaction, error) {
	return _RobotToken.Contract.StartPresale(&_RobotToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_RobotToken *RobotTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RobotToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_RobotToken *RobotTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RobotToken.Contract.Transfer(&_RobotToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_RobotToken *RobotTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RobotToken.Contract.Transfer(&_RobotToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_RobotToken *RobotTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RobotToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_RobotToken *RobotTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RobotToken.Contract.TransferFrom(&_RobotToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_RobotToken *RobotTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RobotToken.Contract.TransferFrom(&_RobotToken.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RobotToken *RobotTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RobotToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RobotToken *RobotTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RobotToken.Contract.TransferOwnership(&_RobotToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RobotToken *RobotTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RobotToken.Contract.TransferOwnership(&_RobotToken.TransactOpts, newOwner)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_RobotToken *RobotTokenTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RobotToken.contract.Transact(opts, "withdrawToken", _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_RobotToken *RobotTokenSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RobotToken.Contract.WithdrawToken(&_RobotToken.TransactOpts, _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_RobotToken *RobotTokenTransactorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RobotToken.Contract.WithdrawToken(&_RobotToken.TransactOpts, _token, _account, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RobotToken *RobotTokenTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RobotToken.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RobotToken *RobotTokenSession) Receive() (*types.Transaction, error) {
	return _RobotToken.Contract.Receive(&_RobotToken.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RobotToken *RobotTokenTransactorSession) Receive() (*types.Transaction, error) {
	return _RobotToken.Contract.Receive(&_RobotToken.TransactOpts)
}

// RobotTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RobotToken contract.
type RobotTokenApprovalIterator struct {
	Event *RobotTokenApproval // Event containing the contract specifics and raw log

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
func (it *RobotTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobotTokenApproval)
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
		it.Event = new(RobotTokenApproval)
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
func (it *RobotTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobotTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobotTokenApproval represents a Approval event raised by the RobotToken contract.
type RobotTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RobotToken *RobotTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*RobotTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RobotToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &RobotTokenApprovalIterator{contract: _RobotToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RobotToken *RobotTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RobotTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RobotToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobotTokenApproval)
				if err := _RobotToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_RobotToken *RobotTokenFilterer) ParseApproval(log types.Log) (*RobotTokenApproval, error) {
	event := new(RobotTokenApproval)
	if err := _RobotToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RobotTokenDifficultyChangedIterator is returned from FilterDifficultyChanged and is used to iterate over the raw logs and unpacked data for DifficultyChanged events raised by the RobotToken contract.
type RobotTokenDifficultyChangedIterator struct {
	Event *RobotTokenDifficultyChanged // Event containing the contract specifics and raw log

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
func (it *RobotTokenDifficultyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobotTokenDifficultyChanged)
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
		it.Event = new(RobotTokenDifficultyChanged)
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
func (it *RobotTokenDifficultyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobotTokenDifficultyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobotTokenDifficultyChanged represents a DifficultyChanged event raised by the RobotToken contract.
type RobotTokenDifficultyChanged struct {
	OriginDifficulty *big.Int
	Difficulty       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDifficultyChanged is a free log retrieval operation binding the contract event 0x935ed35c9eb4d93874341d9aaea64ec36e6fa70b445537b90abbcdc6e3e9bb87.
//
// Solidity: event DifficultyChanged(uint256 originDifficulty, uint256 difficulty)
func (_RobotToken *RobotTokenFilterer) FilterDifficultyChanged(opts *bind.FilterOpts) (*RobotTokenDifficultyChangedIterator, error) {

	logs, sub, err := _RobotToken.contract.FilterLogs(opts, "DifficultyChanged")
	if err != nil {
		return nil, err
	}
	return &RobotTokenDifficultyChangedIterator{contract: _RobotToken.contract, event: "DifficultyChanged", logs: logs, sub: sub}, nil
}

// WatchDifficultyChanged is a free log subscription operation binding the contract event 0x935ed35c9eb4d93874341d9aaea64ec36e6fa70b445537b90abbcdc6e3e9bb87.
//
// Solidity: event DifficultyChanged(uint256 originDifficulty, uint256 difficulty)
func (_RobotToken *RobotTokenFilterer) WatchDifficultyChanged(opts *bind.WatchOpts, sink chan<- *RobotTokenDifficultyChanged) (event.Subscription, error) {

	logs, sub, err := _RobotToken.contract.WatchLogs(opts, "DifficultyChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobotTokenDifficultyChanged)
				if err := _RobotToken.contract.UnpackLog(event, "DifficultyChanged", log); err != nil {
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
func (_RobotToken *RobotTokenFilterer) ParseDifficultyChanged(log types.Log) (*RobotTokenDifficultyChanged, error) {
	event := new(RobotTokenDifficultyChanged)
	if err := _RobotToken.contract.UnpackLog(event, "DifficultyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RobotTokenFoundationMintedIterator is returned from FilterFoundationMinted and is used to iterate over the raw logs and unpacked data for FoundationMinted events raised by the RobotToken contract.
type RobotTokenFoundationMintedIterator struct {
	Event *RobotTokenFoundationMinted // Event containing the contract specifics and raw log

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
func (it *RobotTokenFoundationMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobotTokenFoundationMinted)
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
		it.Event = new(RobotTokenFoundationMinted)
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
func (it *RobotTokenFoundationMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobotTokenFoundationMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobotTokenFoundationMinted represents a FoundationMinted event raised by the RobotToken contract.
type RobotTokenFoundationMinted struct {
	Foundation common.Address
	Amount     *big.Int
	Reason     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFoundationMinted is a free log retrieval operation binding the contract event 0x8ab2777cb42805ec254e8ba0437ce0e96ce42712f1e4c3c944873c2365b780a9.
//
// Solidity: event FoundationMinted(address indexed foundation, uint256 amount, string reason)
func (_RobotToken *RobotTokenFilterer) FilterFoundationMinted(opts *bind.FilterOpts, foundation []common.Address) (*RobotTokenFoundationMintedIterator, error) {

	var foundationRule []interface{}
	for _, foundationItem := range foundation {
		foundationRule = append(foundationRule, foundationItem)
	}

	logs, sub, err := _RobotToken.contract.FilterLogs(opts, "FoundationMinted", foundationRule)
	if err != nil {
		return nil, err
	}
	return &RobotTokenFoundationMintedIterator{contract: _RobotToken.contract, event: "FoundationMinted", logs: logs, sub: sub}, nil
}

// WatchFoundationMinted is a free log subscription operation binding the contract event 0x8ab2777cb42805ec254e8ba0437ce0e96ce42712f1e4c3c944873c2365b780a9.
//
// Solidity: event FoundationMinted(address indexed foundation, uint256 amount, string reason)
func (_RobotToken *RobotTokenFilterer) WatchFoundationMinted(opts *bind.WatchOpts, sink chan<- *RobotTokenFoundationMinted, foundation []common.Address) (event.Subscription, error) {

	var foundationRule []interface{}
	for _, foundationItem := range foundation {
		foundationRule = append(foundationRule, foundationItem)
	}

	logs, sub, err := _RobotToken.contract.WatchLogs(opts, "FoundationMinted", foundationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobotTokenFoundationMinted)
				if err := _RobotToken.contract.UnpackLog(event, "FoundationMinted", log); err != nil {
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
func (_RobotToken *RobotTokenFilterer) ParseFoundationMinted(log types.Log) (*RobotTokenFoundationMinted, error) {
	event := new(RobotTokenFoundationMinted)
	if err := _RobotToken.contract.UnpackLog(event, "FoundationMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RobotTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RobotToken contract.
type RobotTokenOwnershipTransferredIterator struct {
	Event *RobotTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RobotTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobotTokenOwnershipTransferred)
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
		it.Event = new(RobotTokenOwnershipTransferred)
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
func (it *RobotTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobotTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobotTokenOwnershipTransferred represents a OwnershipTransferred event raised by the RobotToken contract.
type RobotTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RobotToken *RobotTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RobotTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RobotToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RobotTokenOwnershipTransferredIterator{contract: _RobotToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RobotToken *RobotTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RobotTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RobotToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobotTokenOwnershipTransferred)
				if err := _RobotToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RobotToken *RobotTokenFilterer) ParseOwnershipTransferred(log types.Log) (*RobotTokenOwnershipTransferred, error) {
	event := new(RobotTokenOwnershipTransferred)
	if err := _RobotToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RobotTokenParticipatedPresaleIterator is returned from FilterParticipatedPresale and is used to iterate over the raw logs and unpacked data for ParticipatedPresale events raised by the RobotToken contract.
type RobotTokenParticipatedPresaleIterator struct {
	Event *RobotTokenParticipatedPresale // Event containing the contract specifics and raw log

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
func (it *RobotTokenParticipatedPresaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobotTokenParticipatedPresale)
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
		it.Event = new(RobotTokenParticipatedPresale)
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
func (it *RobotTokenParticipatedPresaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobotTokenParticipatedPresaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobotTokenParticipatedPresale represents a ParticipatedPresale event raised by the RobotToken contract.
type RobotTokenParticipatedPresale struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterParticipatedPresale is a free log retrieval operation binding the contract event 0xbb286a68c133c09db67567e13e258df5ce81208a09b603627cf2ab8c7c044264.
//
// Solidity: event ParticipatedPresale(address indexed user, uint256 amount)
func (_RobotToken *RobotTokenFilterer) FilterParticipatedPresale(opts *bind.FilterOpts, user []common.Address) (*RobotTokenParticipatedPresaleIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _RobotToken.contract.FilterLogs(opts, "ParticipatedPresale", userRule)
	if err != nil {
		return nil, err
	}
	return &RobotTokenParticipatedPresaleIterator{contract: _RobotToken.contract, event: "ParticipatedPresale", logs: logs, sub: sub}, nil
}

// WatchParticipatedPresale is a free log subscription operation binding the contract event 0xbb286a68c133c09db67567e13e258df5ce81208a09b603627cf2ab8c7c044264.
//
// Solidity: event ParticipatedPresale(address indexed user, uint256 amount)
func (_RobotToken *RobotTokenFilterer) WatchParticipatedPresale(opts *bind.WatchOpts, sink chan<- *RobotTokenParticipatedPresale, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _RobotToken.contract.WatchLogs(opts, "ParticipatedPresale", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobotTokenParticipatedPresale)
				if err := _RobotToken.contract.UnpackLog(event, "ParticipatedPresale", log); err != nil {
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

// ParseParticipatedPresale is a log parse operation binding the contract event 0xbb286a68c133c09db67567e13e258df5ce81208a09b603627cf2ab8c7c044264.
//
// Solidity: event ParticipatedPresale(address indexed user, uint256 amount)
func (_RobotToken *RobotTokenFilterer) ParseParticipatedPresale(log types.Log) (*RobotTokenParticipatedPresale, error) {
	event := new(RobotTokenParticipatedPresale)
	if err := _RobotToken.contract.UnpackLog(event, "ParticipatedPresale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RobotTokenPowMineStartedIterator is returned from FilterPowMineStarted and is used to iterate over the raw logs and unpacked data for PowMineStarted events raised by the RobotToken contract.
type RobotTokenPowMineStartedIterator struct {
	Event *RobotTokenPowMineStarted // Event containing the contract specifics and raw log

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
func (it *RobotTokenPowMineStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobotTokenPowMineStarted)
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
		it.Event = new(RobotTokenPowMineStarted)
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
func (it *RobotTokenPowMineStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobotTokenPowMineStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobotTokenPowMineStarted represents a PowMineStarted event raised by the RobotToken contract.
type RobotTokenPowMineStarted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPowMineStarted is a free log retrieval operation binding the contract event 0xfa0763bffdd0f8a347379599f86a13fe74cb6d8ca1455f23e1352c0c33f563f1.
//
// Solidity: event PowMineStarted()
func (_RobotToken *RobotTokenFilterer) FilterPowMineStarted(opts *bind.FilterOpts) (*RobotTokenPowMineStartedIterator, error) {

	logs, sub, err := _RobotToken.contract.FilterLogs(opts, "PowMineStarted")
	if err != nil {
		return nil, err
	}
	return &RobotTokenPowMineStartedIterator{contract: _RobotToken.contract, event: "PowMineStarted", logs: logs, sub: sub}, nil
}

// WatchPowMineStarted is a free log subscription operation binding the contract event 0xfa0763bffdd0f8a347379599f86a13fe74cb6d8ca1455f23e1352c0c33f563f1.
//
// Solidity: event PowMineStarted()
func (_RobotToken *RobotTokenFilterer) WatchPowMineStarted(opts *bind.WatchOpts, sink chan<- *RobotTokenPowMineStarted) (event.Subscription, error) {

	logs, sub, err := _RobotToken.contract.WatchLogs(opts, "PowMineStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobotTokenPowMineStarted)
				if err := _RobotToken.contract.UnpackLog(event, "PowMineStarted", log); err != nil {
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

// ParsePowMineStarted is a log parse operation binding the contract event 0xfa0763bffdd0f8a347379599f86a13fe74cb6d8ca1455f23e1352c0c33f563f1.
//
// Solidity: event PowMineStarted()
func (_RobotToken *RobotTokenFilterer) ParsePowMineStarted(log types.Log) (*RobotTokenPowMineStarted, error) {
	event := new(RobotTokenPowMineStarted)
	if err := _RobotToken.contract.UnpackLog(event, "PowMineStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RobotTokenPresaleStartedIterator is returned from FilterPresaleStarted and is used to iterate over the raw logs and unpacked data for PresaleStarted events raised by the RobotToken contract.
type RobotTokenPresaleStartedIterator struct {
	Event *RobotTokenPresaleStarted // Event containing the contract specifics and raw log

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
func (it *RobotTokenPresaleStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobotTokenPresaleStarted)
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
		it.Event = new(RobotTokenPresaleStarted)
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
func (it *RobotTokenPresaleStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobotTokenPresaleStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobotTokenPresaleStarted represents a PresaleStarted event raised by the RobotToken contract.
type RobotTokenPresaleStarted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPresaleStarted is a free log retrieval operation binding the contract event 0x17c3338141363aab2512c08f8a7764328ca95979f7057663eb93f7e250139b4c.
//
// Solidity: event PresaleStarted()
func (_RobotToken *RobotTokenFilterer) FilterPresaleStarted(opts *bind.FilterOpts) (*RobotTokenPresaleStartedIterator, error) {

	logs, sub, err := _RobotToken.contract.FilterLogs(opts, "PresaleStarted")
	if err != nil {
		return nil, err
	}
	return &RobotTokenPresaleStartedIterator{contract: _RobotToken.contract, event: "PresaleStarted", logs: logs, sub: sub}, nil
}

// WatchPresaleStarted is a free log subscription operation binding the contract event 0x17c3338141363aab2512c08f8a7764328ca95979f7057663eb93f7e250139b4c.
//
// Solidity: event PresaleStarted()
func (_RobotToken *RobotTokenFilterer) WatchPresaleStarted(opts *bind.WatchOpts, sink chan<- *RobotTokenPresaleStarted) (event.Subscription, error) {

	logs, sub, err := _RobotToken.contract.WatchLogs(opts, "PresaleStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobotTokenPresaleStarted)
				if err := _RobotToken.contract.UnpackLog(event, "PresaleStarted", log); err != nil {
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

// ParsePresaleStarted is a log parse operation binding the contract event 0x17c3338141363aab2512c08f8a7764328ca95979f7057663eb93f7e250139b4c.
//
// Solidity: event PresaleStarted()
func (_RobotToken *RobotTokenFilterer) ParsePresaleStarted(log types.Log) (*RobotTokenPresaleStarted, error) {
	event := new(RobotTokenPresaleStarted)
	if err := _RobotToken.contract.UnpackLog(event, "PresaleStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RobotTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RobotToken contract.
type RobotTokenTransferIterator struct {
	Event *RobotTokenTransfer // Event containing the contract specifics and raw log

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
func (it *RobotTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobotTokenTransfer)
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
		it.Event = new(RobotTokenTransfer)
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
func (it *RobotTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobotTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobotTokenTransfer represents a Transfer event raised by the RobotToken contract.
type RobotTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RobotToken *RobotTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RobotTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RobotToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RobotTokenTransferIterator{contract: _RobotToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RobotToken *RobotTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RobotTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RobotToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobotTokenTransfer)
				if err := _RobotToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_RobotToken *RobotTokenFilterer) ParseTransfer(log types.Log) (*RobotTokenTransfer, error) {
	event := new(RobotTokenTransfer)
	if err := _RobotToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
