// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
)

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"amount\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TokenFuncSigs maps the 4-byte function signature to its string representation.
var TokenFuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"893d20e8": "getOwner()",
	"2ea0dfe1": "transferFrom(address,address,uint64)",
}

// TokenBin is the compiled bytecode used for deploying new contracts.
var TokenBin = "0x608060405234801561001057600080fd5b506040516020806103e6833981016040908152905160008054600160a060020a0319163390811782558152600160205291909120805467ffffffffffffffff90921667ffffffffffffffff19909216919091179055610372806100746000396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632ea0dfe1811461005b57806370a08231146100a3578063893d20e8146100e1575b600080fd5b34801561006757600080fd5b5061008f600160a060020a036004358116906024351667ffffffffffffffff60443516610112565b604080519115158252519081900360200190f35b3480156100af57600080fd5b506100c4600160a060020a0360043516610312565b6040805167ffffffffffffffff9092168252519081900360200190f35b3480156100ed57600080fd5b506100f6610337565b60408051600160a060020a039092168252519081900360200190f35b600033600160a060020a03851614610174576040805160e560020a62461bcd02815260206004820152601e60248201527f66726f6d2061646472657373206973206e6f74206d73672073656e6465720000604482015290519081900360640190fd5b600160a060020a03831615156101d4576040805160e560020a62461bcd02815260206004820152601160248201527f746f206164647265737320697320307830000000000000000000000000000000604482015290519081900360640190fd5b67ffffffffffffffff82161515610235576040805160e560020a62461bcd02815260206004820152601460248201527f7472616e7366657220616d6f756e742069732030000000000000000000000000604482015290519081900360640190fd5b600160a060020a03841660009081526001602052604090205467ffffffffffffffff808416911610156102b2576040805160e560020a62461bcd02815260206004820152601e60248201527f66726f6d2061646472657373206e6f20656e6f7567682062616c616e63650000604482015290519081900360640190fd5b50600160a060020a039283166000908152600160208190526040808320805467ffffffffffffffff808216879003811667ffffffffffffffff19928316179092559590961683529091208054808616909301909416919092161790915590565b600160a060020a031660009081526001602052604090205467ffffffffffffffff1690565b600054600160a060020a0316905600a165627a7a723058200958f6c4c1578ceb0207e67e8834313d86c5fbb64782d7d60b004403a7344fc00029"

// DeployToken deploys a new contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend, amount uint64) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend, amount)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

func AsyncDeployToken(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend, amount uint64) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(TokenBin), backend, amount)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Token is an auto generated Go binding around a Solidity contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around a Solidity contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around a Solidity contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around a Solidity contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) constant returns(uint64)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, addr common.Address) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", addr)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) constant returns(uint64)
func (_Token *TokenSession) BalanceOf(addr common.Address) (uint64, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) constant returns(uint64)
func (_Token *TokenCallerSession) BalanceOf(addr common.Address) (uint64, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, addr)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Token *TokenCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Token *TokenSession) GetOwner() (common.Address, error) {
	return _Token.Contract.GetOwner(&_Token.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Token *TokenCallerSession) GetOwner() (common.Address, error) {
	return _Token.Contract.GetOwner(&_Token.CallOpts)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x2ea0dfe1.
//
// Solidity: function transferFrom(address from, address to, uint64 amount) returns(bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount uint64) (*types.Transaction, *types.Receipt, error) {
	return _Token.contract.Transact(opts, "transferFrom", from, to, amount)
}

func (_Token *TokenTransactor) AsyncTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, amount uint64) (*types.Transaction, error) {
	return _Token.contract.AsyncTransact(opts, handler, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x2ea0dfe1.
//
// Solidity: function transferFrom(address from, address to, uint64 amount) returns(bool)
func (_Token *TokenSession) TransferFrom(from common.Address, to common.Address, amount uint64) (*types.Transaction, *types.Receipt, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, from, to, amount)
}

func (_Token *TokenSession) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, amount uint64) (*types.Transaction, error) {
	return _Token.Contract.AsyncTransferFrom(handler, &_Token.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x2ea0dfe1.
//
// Solidity: function transferFrom(address from, address to, uint64 amount) returns(bool)
func (_Token *TokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount uint64) (*types.Transaction, *types.Receipt, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, from, to, amount)
}

func (_Token *TokenTransactorSession) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, amount uint64) (*types.Transaction, error) {
	return _Token.Contract.AsyncTransferFrom(handler, &_Token.TransactOpts, from, to, amount)
}
