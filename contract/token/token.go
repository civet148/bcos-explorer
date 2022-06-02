// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

import (
	"math/big"
	"strings"

	"git.impool.com/storeros/bcos-sdk/abi"
	"git.impool.com/storeros/bcos-sdk/abi/bind"
	"git.impool.com/storeros/bcos-sdk/core/types"
	"git.impool.com/storeros/bcos-sdk/event"
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
	_ = event.NewSubscription
)

// LoggerABI is the input ABI used to generate the binding from.
const LoggerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"LogUint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"\",\"type\":\"int256\"}],\"name\":\"LogInt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"LogBytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"LogBytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"\",\"type\":\"address\"}],\"name\":\"LogAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"\",\"type\":\"bool\"}],\"name\":\"LogBool\",\"type\":\"event\"}]"

// LoggerBin is the compiled bytecode used for deploying new contracts.
var LoggerBin = "0x6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a72305820fbfb1210b7eee90a86e80606730c59252a59ccd8e7fad4b446b478e6dea8f5750029"

// DeployLogger deploys a new contract, binding an instance of Logger to it.
func DeployLogger(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Logger, error) {
	parsed, err := abi.JSON(strings.NewReader(LoggerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LoggerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Logger{LoggerCaller: LoggerCaller{contract: contract}, LoggerTransactor: LoggerTransactor{contract: contract}, LoggerFilterer: LoggerFilterer{contract: contract}}, nil
}

func AsyncDeployLogger(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(LoggerABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(LoggerBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Logger is an auto generated Go binding around a Solidity contract.
type Logger struct {
	LoggerCaller     // Read-only binding to the contract
	LoggerTransactor // Write-only binding to the contract
	LoggerFilterer   // Log filterer for contract events
}

// LoggerCaller is an auto generated read-only Go binding around a Solidity contract.
type LoggerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoggerTransactor is an auto generated write-only Go binding around a Solidity contract.
type LoggerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoggerFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type LoggerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoggerSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type LoggerSession struct {
	Contract     *Logger           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoggerCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type LoggerCallerSession struct {
	Contract *LoggerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LoggerTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type LoggerTransactorSession struct {
	Contract     *LoggerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoggerRaw is an auto generated low-level Go binding around a Solidity contract.
type LoggerRaw struct {
	Contract *Logger // Generic contract binding to access the raw methods on
}

// LoggerCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type LoggerCallerRaw struct {
	Contract *LoggerCaller // Generic read-only contract binding to access the raw methods on
}

// LoggerTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type LoggerTransactorRaw struct {
	Contract *LoggerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLogger creates a new instance of Logger, bound to a specific deployed contract.
func NewLogger(address common.Address, backend bind.ContractBackend) (*Logger, error) {
	contract, err := bindLogger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Logger{LoggerCaller: LoggerCaller{contract: contract}, LoggerTransactor: LoggerTransactor{contract: contract}, LoggerFilterer: LoggerFilterer{contract: contract}}, nil
}

// NewLoggerCaller creates a new read-only instance of Logger, bound to a specific deployed contract.
func NewLoggerCaller(address common.Address, caller bind.ContractCaller) (*LoggerCaller, error) {
	contract, err := bindLogger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoggerCaller{contract: contract}, nil
}

// NewLoggerTransactor creates a new write-only instance of Logger, bound to a specific deployed contract.
func NewLoggerTransactor(address common.Address, transactor bind.ContractTransactor) (*LoggerTransactor, error) {
	contract, err := bindLogger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoggerTransactor{contract: contract}, nil
}

// NewLoggerFilterer creates a new log filterer instance of Logger, bound to a specific deployed contract.
func NewLoggerFilterer(address common.Address, filterer bind.ContractFilterer) (*LoggerFilterer, error) {
	contract, err := bindLogger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoggerFilterer{contract: contract}, nil
}

// bindLogger binds a generic wrapper to an already deployed contract.
func bindLogger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LoggerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Logger *LoggerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Logger.Contract.LoggerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Logger *LoggerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Logger.Contract.LoggerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Logger *LoggerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Logger.Contract.LoggerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Logger *LoggerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Logger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Logger *LoggerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Logger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Logger *LoggerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Logger.Contract.contract.Transact(opts, method, params...)
}

// LoggerLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the Logger contract.
type LoggerLogAddressIterator struct {
	Event *LoggerLogAddress // Event containing the contract specifics and raw log

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
func (it *LoggerLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoggerLogAddress)
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
		it.Event = new(LoggerLogAddress)
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
func (it *LoggerLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoggerLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoggerLogAddress represents a LogAddress event raised by the Logger contract.
type LoggerLogAddress struct {
	Arg0 string
	Arg1 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x62ddffe5b5108385f7a590f100e1ee414ad9551a31f089e64e82998440785e1e.
//
// Solidity: event LogAddress(string , address )
func (_Logger *LoggerFilterer) FilterLogAddress(opts *bind.FilterOpts) (*LoggerLogAddressIterator, error) {

	logs, sub, err := _Logger.contract.FilterLogs(opts, "LogAddress")
	if err != nil {
		return nil, err
	}
	return &LoggerLogAddressIterator{contract: _Logger.contract, event: "LogAddress", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x62ddffe5b5108385f7a590f100e1ee414ad9551a31f089e64e82998440785e1e.
//
// Solidity: event LogAddress(string , address )
func (_Logger *LoggerFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *LoggerLogAddress) (event.Subscription, error) {

	logs, sub, err := _Logger.contract.WatchLogs(opts, "LogAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoggerLogAddress)
				if err := _Logger.contract.UnpackLog(event, "LogAddress", log); err != nil {
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

// ParseLogAddress is a log parse operation binding the contract event 0x62ddffe5b5108385f7a590f100e1ee414ad9551a31f089e64e82998440785e1e.
//
// Solidity: event LogAddress(string , address )
func (_Logger *LoggerFilterer) ParseLogAddress(log types.Log) (*LoggerLogAddress, error) {
	event := new(LoggerLogAddress)
	if err := _Logger.contract.UnpackLog(event, "LogAddress", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoggerLogBoolIterator is returned from FilterLogBool and is used to iterate over the raw logs and unpacked data for LogBool events raised by the Logger contract.
type LoggerLogBoolIterator struct {
	Event *LoggerLogBool // Event containing the contract specifics and raw log

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
func (it *LoggerLogBoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoggerLogBool)
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
		it.Event = new(LoggerLogBool)
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
func (it *LoggerLogBoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoggerLogBoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoggerLogBool represents a LogBool event raised by the Logger contract.
type LoggerLogBool struct {
	Arg0 string
	Arg1 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBool is a free log retrieval operation binding the contract event 0x4c34c2f9a78632f29fa59aaed5514cb742fd9fbcfd7ccc2c03c85f2bbc621c47.
//
// Solidity: event LogBool(string , bool )
func (_Logger *LoggerFilterer) FilterLogBool(opts *bind.FilterOpts) (*LoggerLogBoolIterator, error) {

	logs, sub, err := _Logger.contract.FilterLogs(opts, "LogBool")
	if err != nil {
		return nil, err
	}
	return &LoggerLogBoolIterator{contract: _Logger.contract, event: "LogBool", logs: logs, sub: sub}, nil
}

// WatchLogBool is a free log subscription operation binding the contract event 0x4c34c2f9a78632f29fa59aaed5514cb742fd9fbcfd7ccc2c03c85f2bbc621c47.
//
// Solidity: event LogBool(string , bool )
func (_Logger *LoggerFilterer) WatchLogBool(opts *bind.WatchOpts, sink chan<- *LoggerLogBool) (event.Subscription, error) {

	logs, sub, err := _Logger.contract.WatchLogs(opts, "LogBool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoggerLogBool)
				if err := _Logger.contract.UnpackLog(event, "LogBool", log); err != nil {
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

// ParseLogBool is a log parse operation binding the contract event 0x4c34c2f9a78632f29fa59aaed5514cb742fd9fbcfd7ccc2c03c85f2bbc621c47.
//
// Solidity: event LogBool(string , bool )
func (_Logger *LoggerFilterer) ParseLogBool(log types.Log) (*LoggerLogBool, error) {
	event := new(LoggerLogBool)
	if err := _Logger.contract.UnpackLog(event, "LogBool", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoggerLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the Logger contract.
type LoggerLogBytesIterator struct {
	Event *LoggerLogBytes // Event containing the contract specifics and raw log

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
func (it *LoggerLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoggerLogBytes)
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
		it.Event = new(LoggerLogBytes)
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
func (it *LoggerLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoggerLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoggerLogBytes represents a LogBytes event raised by the Logger contract.
type LoggerLogBytes struct {
	Arg0 string
	Arg1 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0xe8407a0209fa99ec3a7228aff140c3d3e68bd279391739c7e0b65cd406cc93b5.
//
// Solidity: event LogBytes(string , bytes )
func (_Logger *LoggerFilterer) FilterLogBytes(opts *bind.FilterOpts) (*LoggerLogBytesIterator, error) {

	logs, sub, err := _Logger.contract.FilterLogs(opts, "LogBytes")
	if err != nil {
		return nil, err
	}
	return &LoggerLogBytesIterator{contract: _Logger.contract, event: "LogBytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0xe8407a0209fa99ec3a7228aff140c3d3e68bd279391739c7e0b65cd406cc93b5.
//
// Solidity: event LogBytes(string , bytes )
func (_Logger *LoggerFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *LoggerLogBytes) (event.Subscription, error) {

	logs, sub, err := _Logger.contract.WatchLogs(opts, "LogBytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoggerLogBytes)
				if err := _Logger.contract.UnpackLog(event, "LogBytes", log); err != nil {
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

// ParseLogBytes is a log parse operation binding the contract event 0xe8407a0209fa99ec3a7228aff140c3d3e68bd279391739c7e0b65cd406cc93b5.
//
// Solidity: event LogBytes(string , bytes )
func (_Logger *LoggerFilterer) ParseLogBytes(log types.Log) (*LoggerLogBytes, error) {
	event := new(LoggerLogBytes)
	if err := _Logger.contract.UnpackLog(event, "LogBytes", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoggerLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the Logger contract.
type LoggerLogBytes32Iterator struct {
	Event *LoggerLogBytes32 // Event containing the contract specifics and raw log

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
func (it *LoggerLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoggerLogBytes32)
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
		it.Event = new(LoggerLogBytes32)
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
func (it *LoggerLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoggerLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoggerLogBytes32 represents a LogBytes32 event raised by the Logger contract.
type LoggerLogBytes32 struct {
	Arg0 string
	Arg1 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0x02d93529bba9d141e5e06733c52c7e6fbcb1149586adb5c24064b522ab26f1d7.
//
// Solidity: event LogBytes32(string , bytes32 )
func (_Logger *LoggerFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*LoggerLogBytes32Iterator, error) {

	logs, sub, err := _Logger.contract.FilterLogs(opts, "LogBytes32")
	if err != nil {
		return nil, err
	}
	return &LoggerLogBytes32Iterator{contract: _Logger.contract, event: "LogBytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0x02d93529bba9d141e5e06733c52c7e6fbcb1149586adb5c24064b522ab26f1d7.
//
// Solidity: event LogBytes32(string , bytes32 )
func (_Logger *LoggerFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *LoggerLogBytes32) (event.Subscription, error) {

	logs, sub, err := _Logger.contract.WatchLogs(opts, "LogBytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoggerLogBytes32)
				if err := _Logger.contract.UnpackLog(event, "LogBytes32", log); err != nil {
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

// ParseLogBytes32 is a log parse operation binding the contract event 0x02d93529bba9d141e5e06733c52c7e6fbcb1149586adb5c24064b522ab26f1d7.
//
// Solidity: event LogBytes32(string , bytes32 )
func (_Logger *LoggerFilterer) ParseLogBytes32(log types.Log) (*LoggerLogBytes32, error) {
	event := new(LoggerLogBytes32)
	if err := _Logger.contract.UnpackLog(event, "LogBytes32", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoggerLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the Logger contract.
type LoggerLogIntIterator struct {
	Event *LoggerLogInt // Event containing the contract specifics and raw log

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
func (it *LoggerLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoggerLogInt)
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
		it.Event = new(LoggerLogInt)
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
func (it *LoggerLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoggerLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoggerLogInt represents a LogInt event raised by the Logger contract.
type LoggerLogInt struct {
	Arg0 string
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x6a837ff3973aa4181e7b9f07756f8b7ee366dd85a36655d2cb42cd47f10b9638.
//
// Solidity: event LogInt(string , int256 )
func (_Logger *LoggerFilterer) FilterLogInt(opts *bind.FilterOpts) (*LoggerLogIntIterator, error) {

	logs, sub, err := _Logger.contract.FilterLogs(opts, "LogInt")
	if err != nil {
		return nil, err
	}
	return &LoggerLogIntIterator{contract: _Logger.contract, event: "LogInt", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x6a837ff3973aa4181e7b9f07756f8b7ee366dd85a36655d2cb42cd47f10b9638.
//
// Solidity: event LogInt(string , int256 )
func (_Logger *LoggerFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *LoggerLogInt) (event.Subscription, error) {

	logs, sub, err := _Logger.contract.WatchLogs(opts, "LogInt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoggerLogInt)
				if err := _Logger.contract.UnpackLog(event, "LogInt", log); err != nil {
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

// ParseLogInt is a log parse operation binding the contract event 0x6a837ff3973aa4181e7b9f07756f8b7ee366dd85a36655d2cb42cd47f10b9638.
//
// Solidity: event LogInt(string , int256 )
func (_Logger *LoggerFilterer) ParseLogInt(log types.Log) (*LoggerLogInt, error) {
	event := new(LoggerLogInt)
	if err := _Logger.contract.UnpackLog(event, "LogInt", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoggerLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the Logger contract.
type LoggerLogUintIterator struct {
	Event *LoggerLogUint // Event containing the contract specifics and raw log

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
func (it *LoggerLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoggerLogUint)
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
		it.Event = new(LoggerLogUint)
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
func (it *LoggerLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoggerLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoggerLogUint represents a LogUint event raised by the Logger contract.
type LoggerLogUint struct {
	Arg0 string
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x941296a39ea107bde685522318a4b6c2b544904a5dd82a512748ca2cf839bef7.
//
// Solidity: event LogUint(string , uint256 )
func (_Logger *LoggerFilterer) FilterLogUint(opts *bind.FilterOpts) (*LoggerLogUintIterator, error) {

	logs, sub, err := _Logger.contract.FilterLogs(opts, "LogUint")
	if err != nil {
		return nil, err
	}
	return &LoggerLogUintIterator{contract: _Logger.contract, event: "LogUint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x941296a39ea107bde685522318a4b6c2b544904a5dd82a512748ca2cf839bef7.
//
// Solidity: event LogUint(string , uint256 )
func (_Logger *LoggerFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *LoggerLogUint) (event.Subscription, error) {

	logs, sub, err := _Logger.contract.WatchLogs(opts, "LogUint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoggerLogUint)
				if err := _Logger.contract.UnpackLog(event, "LogUint", log); err != nil {
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

// ParseLogUint is a log parse operation binding the contract event 0x941296a39ea107bde685522318a4b6c2b544904a5dd82a512748ca2cf839bef7.
//
// Solidity: event LogUint(string , uint256 )
func (_Logger *LoggerFilterer) ParseLogUint(log types.Log) (*LoggerLogUint, error) {
	event := new(LoggerLogUint)
	if err := _Logger.contract.UnpackLog(event, "LogUint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"amount\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"LogUint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"\",\"type\":\"int256\"}],\"name\":\"LogInt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"LogBytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"LogBytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"\",\"type\":\"address\"}],\"name\":\"LogAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"\",\"type\":\"bool\"}],\"name\":\"LogBool\",\"type\":\"event\"}]"

// TokenFuncSigs maps the 4-byte function signature to its string representation.
var TokenFuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"893d20e8": "getOwner()",
	"2ea0dfe1": "transferFrom(address,address,uint64)",
}

// TokenBin is the compiled bytecode used for deploying new contracts.
var TokenBin = "0x608060405234801561001057600080fd5b50604051602080610346833981018060405261002f9190810190610086565b60008054600160a060020a031916339081178255815260016020526040902080546001604060020a0390921667ffffffffffffffff199092169190911790556100b8565b600061007f82516100ac565b9392505050565b60006020828403121561009857600080fd5b60006100a48484610073565b949350505050565b6001604060020a031690565b61027f806100c76000396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632ea0dfe1811461005b57806370a0823114610091578063893d20e8146100be575b600080fd5b34801561006757600080fd5b5061007b61007636600461017c565b6100e0565b60405161008891906101fe565b60405180910390f35b34801561009d57600080fd5b506100b16100ac366004610156565b6100e9565b604051610088919061020c565b3480156100ca57600080fd5b506100d361011b565b60405161008891906101ea565b60019392505050565b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604090205467ffffffffffffffff1690565b60005473ffffffffffffffffffffffffffffffffffffffff1690565b6000610143823561021a565b9392505050565b60006101438235610238565b60006020828403121561016857600080fd5b60006101748484610137565b949350505050565b60008060006060848603121561019157600080fd5b600061019d8686610137565b93505060206101ae86828701610137565b92505060406101bf8682870161014a565b9150509250925092565b6101d28161021a565b82525050565b6101d281610233565b6101d281610238565b602081016101f882846101c9565b92915050565b602081016101f882846101d8565b602081016101f882846101e1565b73ffffffffffffffffffffffffffffffffffffffff1690565b151590565b67ffffffffffffffff16905600a265627a7a72305820f5787fc07bd1bfb9402f61a1920ff3975673caa5eab3423559194fe4cb67ad726c6578706572696d656e74616cf50037"

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

// TokenLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the Token contract.
type TokenLogAddressIterator struct {
	Event *TokenLogAddress // Event containing the contract specifics and raw log

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
func (it *TokenLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLogAddress)
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
		it.Event = new(TokenLogAddress)
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
func (it *TokenLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLogAddress represents a LogAddress event raised by the Token contract.
type TokenLogAddress struct {
	Arg0 string
	Arg1 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x62ddffe5b5108385f7a590f100e1ee414ad9551a31f089e64e82998440785e1e.
//
// Solidity: event LogAddress(string , address )
func (_Token *TokenFilterer) FilterLogAddress(opts *bind.FilterOpts) (*TokenLogAddressIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "LogAddress")
	if err != nil {
		return nil, err
	}
	return &TokenLogAddressIterator{contract: _Token.contract, event: "LogAddress", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x62ddffe5b5108385f7a590f100e1ee414ad9551a31f089e64e82998440785e1e.
//
// Solidity: event LogAddress(string , address )
func (_Token *TokenFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *TokenLogAddress) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "LogAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLogAddress)
				if err := _Token.contract.UnpackLog(event, "LogAddress", log); err != nil {
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

// ParseLogAddress is a log parse operation binding the contract event 0x62ddffe5b5108385f7a590f100e1ee414ad9551a31f089e64e82998440785e1e.
//
// Solidity: event LogAddress(string , address )
func (_Token *TokenFilterer) ParseLogAddress(log types.Log) (*TokenLogAddress, error) {
	event := new(TokenLogAddress)
	if err := _Token.contract.UnpackLog(event, "LogAddress", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenLogBoolIterator is returned from FilterLogBool and is used to iterate over the raw logs and unpacked data for LogBool events raised by the Token contract.
type TokenLogBoolIterator struct {
	Event *TokenLogBool // Event containing the contract specifics and raw log

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
func (it *TokenLogBoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLogBool)
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
		it.Event = new(TokenLogBool)
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
func (it *TokenLogBoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLogBoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLogBool represents a LogBool event raised by the Token contract.
type TokenLogBool struct {
	Arg0 string
	Arg1 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBool is a free log retrieval operation binding the contract event 0x4c34c2f9a78632f29fa59aaed5514cb742fd9fbcfd7ccc2c03c85f2bbc621c47.
//
// Solidity: event LogBool(string , bool )
func (_Token *TokenFilterer) FilterLogBool(opts *bind.FilterOpts) (*TokenLogBoolIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "LogBool")
	if err != nil {
		return nil, err
	}
	return &TokenLogBoolIterator{contract: _Token.contract, event: "LogBool", logs: logs, sub: sub}, nil
}

// WatchLogBool is a free log subscription operation binding the contract event 0x4c34c2f9a78632f29fa59aaed5514cb742fd9fbcfd7ccc2c03c85f2bbc621c47.
//
// Solidity: event LogBool(string , bool )
func (_Token *TokenFilterer) WatchLogBool(opts *bind.WatchOpts, sink chan<- *TokenLogBool) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "LogBool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLogBool)
				if err := _Token.contract.UnpackLog(event, "LogBool", log); err != nil {
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

// ParseLogBool is a log parse operation binding the contract event 0x4c34c2f9a78632f29fa59aaed5514cb742fd9fbcfd7ccc2c03c85f2bbc621c47.
//
// Solidity: event LogBool(string , bool )
func (_Token *TokenFilterer) ParseLogBool(log types.Log) (*TokenLogBool, error) {
	event := new(TokenLogBool)
	if err := _Token.contract.UnpackLog(event, "LogBool", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the Token contract.
type TokenLogBytesIterator struct {
	Event *TokenLogBytes // Event containing the contract specifics and raw log

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
func (it *TokenLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLogBytes)
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
		it.Event = new(TokenLogBytes)
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
func (it *TokenLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLogBytes represents a LogBytes event raised by the Token contract.
type TokenLogBytes struct {
	Arg0 string
	Arg1 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0xe8407a0209fa99ec3a7228aff140c3d3e68bd279391739c7e0b65cd406cc93b5.
//
// Solidity: event LogBytes(string , bytes )
func (_Token *TokenFilterer) FilterLogBytes(opts *bind.FilterOpts) (*TokenLogBytesIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "LogBytes")
	if err != nil {
		return nil, err
	}
	return &TokenLogBytesIterator{contract: _Token.contract, event: "LogBytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0xe8407a0209fa99ec3a7228aff140c3d3e68bd279391739c7e0b65cd406cc93b5.
//
// Solidity: event LogBytes(string , bytes )
func (_Token *TokenFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *TokenLogBytes) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "LogBytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLogBytes)
				if err := _Token.contract.UnpackLog(event, "LogBytes", log); err != nil {
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

// ParseLogBytes is a log parse operation binding the contract event 0xe8407a0209fa99ec3a7228aff140c3d3e68bd279391739c7e0b65cd406cc93b5.
//
// Solidity: event LogBytes(string , bytes )
func (_Token *TokenFilterer) ParseLogBytes(log types.Log) (*TokenLogBytes, error) {
	event := new(TokenLogBytes)
	if err := _Token.contract.UnpackLog(event, "LogBytes", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the Token contract.
type TokenLogBytes32Iterator struct {
	Event *TokenLogBytes32 // Event containing the contract specifics and raw log

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
func (it *TokenLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLogBytes32)
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
		it.Event = new(TokenLogBytes32)
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
func (it *TokenLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLogBytes32 represents a LogBytes32 event raised by the Token contract.
type TokenLogBytes32 struct {
	Arg0 string
	Arg1 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0x02d93529bba9d141e5e06733c52c7e6fbcb1149586adb5c24064b522ab26f1d7.
//
// Solidity: event LogBytes32(string , bytes32 )
func (_Token *TokenFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*TokenLogBytes32Iterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "LogBytes32")
	if err != nil {
		return nil, err
	}
	return &TokenLogBytes32Iterator{contract: _Token.contract, event: "LogBytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0x02d93529bba9d141e5e06733c52c7e6fbcb1149586adb5c24064b522ab26f1d7.
//
// Solidity: event LogBytes32(string , bytes32 )
func (_Token *TokenFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *TokenLogBytes32) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "LogBytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLogBytes32)
				if err := _Token.contract.UnpackLog(event, "LogBytes32", log); err != nil {
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

// ParseLogBytes32 is a log parse operation binding the contract event 0x02d93529bba9d141e5e06733c52c7e6fbcb1149586adb5c24064b522ab26f1d7.
//
// Solidity: event LogBytes32(string , bytes32 )
func (_Token *TokenFilterer) ParseLogBytes32(log types.Log) (*TokenLogBytes32, error) {
	event := new(TokenLogBytes32)
	if err := _Token.contract.UnpackLog(event, "LogBytes32", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the Token contract.
type TokenLogIntIterator struct {
	Event *TokenLogInt // Event containing the contract specifics and raw log

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
func (it *TokenLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLogInt)
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
		it.Event = new(TokenLogInt)
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
func (it *TokenLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLogInt represents a LogInt event raised by the Token contract.
type TokenLogInt struct {
	Arg0 string
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x6a837ff3973aa4181e7b9f07756f8b7ee366dd85a36655d2cb42cd47f10b9638.
//
// Solidity: event LogInt(string , int256 )
func (_Token *TokenFilterer) FilterLogInt(opts *bind.FilterOpts) (*TokenLogIntIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "LogInt")
	if err != nil {
		return nil, err
	}
	return &TokenLogIntIterator{contract: _Token.contract, event: "LogInt", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x6a837ff3973aa4181e7b9f07756f8b7ee366dd85a36655d2cb42cd47f10b9638.
//
// Solidity: event LogInt(string , int256 )
func (_Token *TokenFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *TokenLogInt) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "LogInt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLogInt)
				if err := _Token.contract.UnpackLog(event, "LogInt", log); err != nil {
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

// ParseLogInt is a log parse operation binding the contract event 0x6a837ff3973aa4181e7b9f07756f8b7ee366dd85a36655d2cb42cd47f10b9638.
//
// Solidity: event LogInt(string , int256 )
func (_Token *TokenFilterer) ParseLogInt(log types.Log) (*TokenLogInt, error) {
	event := new(TokenLogInt)
	if err := _Token.contract.UnpackLog(event, "LogInt", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the Token contract.
type TokenLogUintIterator struct {
	Event *TokenLogUint // Event containing the contract specifics and raw log

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
func (it *TokenLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLogUint)
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
		it.Event = new(TokenLogUint)
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
func (it *TokenLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLogUint represents a LogUint event raised by the Token contract.
type TokenLogUint struct {
	Arg0 string
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x941296a39ea107bde685522318a4b6c2b544904a5dd82a512748ca2cf839bef7.
//
// Solidity: event LogUint(string , uint256 )
func (_Token *TokenFilterer) FilterLogUint(opts *bind.FilterOpts) (*TokenLogUintIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "LogUint")
	if err != nil {
		return nil, err
	}
	return &TokenLogUintIterator{contract: _Token.contract, event: "LogUint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x941296a39ea107bde685522318a4b6c2b544904a5dd82a512748ca2cf839bef7.
//
// Solidity: event LogUint(string , uint256 )
func (_Token *TokenFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *TokenLogUint) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "LogUint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLogUint)
				if err := _Token.contract.UnpackLog(event, "LogUint", log); err != nil {
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

// ParseLogUint is a log parse operation binding the contract event 0x941296a39ea107bde685522318a4b6c2b544904a5dd82a512748ca2cf839bef7.
//
// Solidity: event LogUint(string , uint256 )
func (_Token *TokenFilterer) ParseLogUint(log types.Log) (*TokenLogUint, error) {
	event := new(TokenLogUint)
	if err := _Token.contract.UnpackLog(event, "LogUint", log); err != nil {
		return nil, err
	}
	return event, nil
}
