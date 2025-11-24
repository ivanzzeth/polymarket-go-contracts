// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gnosissafel2

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

// GnosisSafeL2MetaData contains all meta data concerning the GnosisSafeL2 contract.
var GnosisSafeL2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AddedOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"approvedHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ApproveHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"name\":\"ChangedFallbackHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"ChangedGuard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ChangedThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"DisabledModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"EnabledModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"ExecutionFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ExecutionFromModuleFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ExecutionFromModuleSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"ExecutionSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"RemovedOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initializer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fallbackHandler\",\"type\":\"address\"}],\"name\":\"SafeSetup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"SignMsg\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"addOwnerWithThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashToApprove\",\"type\":\"bytes32\"}],\"name\":\"approveHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"approvedHashes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredSignatures\",\"type\":\"uint256\"}],\"name\":\"checkNSignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"checkSignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"disableModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"enableModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"encodeTransactionData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"execTransaction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"execTransactionFromModule\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"execTransactionFromModuleReturnData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"start\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"getModulesPaginated\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"array\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"next\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getStorageAt\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getTransactionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"isModuleEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"removeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"name\":\"setFallbackHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"setGuard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"fallbackHandler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentReceiver\",\"type\":\"address\"}],\"name\":\"setup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"signedMessages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"calldataPayload\",\"type\":\"bytes\"}],\"name\":\"simulateAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"swapOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// GnosisSafeL2ABI is the input ABI used to generate the binding from.
// Deprecated: Use GnosisSafeL2MetaData.ABI instead.
var GnosisSafeL2ABI = GnosisSafeL2MetaData.ABI

// GnosisSafeL2 is an auto generated Go binding around an Ethereum contract.
type GnosisSafeL2 struct {
	GnosisSafeL2Caller     // Read-only binding to the contract
	GnosisSafeL2Transactor // Write-only binding to the contract
	GnosisSafeL2Filterer   // Log filterer for contract events
}

// GnosisSafeL2Caller is an auto generated read-only Go binding around an Ethereum contract.
type GnosisSafeL2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeL2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type GnosisSafeL2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeL2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GnosisSafeL2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeL2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GnosisSafeL2Session struct {
	Contract     *GnosisSafeL2     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GnosisSafeL2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GnosisSafeL2CallerSession struct {
	Contract *GnosisSafeL2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// GnosisSafeL2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GnosisSafeL2TransactorSession struct {
	Contract     *GnosisSafeL2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GnosisSafeL2Raw is an auto generated low-level Go binding around an Ethereum contract.
type GnosisSafeL2Raw struct {
	Contract *GnosisSafeL2 // Generic contract binding to access the raw methods on
}

// GnosisSafeL2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GnosisSafeL2CallerRaw struct {
	Contract *GnosisSafeL2Caller // Generic read-only contract binding to access the raw methods on
}

// GnosisSafeL2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GnosisSafeL2TransactorRaw struct {
	Contract *GnosisSafeL2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewGnosisSafeL2 creates a new instance of GnosisSafeL2, bound to a specific deployed contract.
func NewGnosisSafeL2(address common.Address, backend bind.ContractBackend) (*GnosisSafeL2, error) {
	contract, err := bindGnosisSafeL2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeL2{GnosisSafeL2Caller: GnosisSafeL2Caller{contract: contract}, GnosisSafeL2Transactor: GnosisSafeL2Transactor{contract: contract}, GnosisSafeL2Filterer: GnosisSafeL2Filterer{contract: contract}}, nil
}

// NewGnosisSafeL2Caller creates a new read-only instance of GnosisSafeL2, bound to a specific deployed contract.
func NewGnosisSafeL2Caller(address common.Address, caller bind.ContractCaller) (*GnosisSafeL2Caller, error) {
	contract, err := bindGnosisSafeL2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeL2Caller{contract: contract}, nil
}

// NewGnosisSafeL2Transactor creates a new write-only instance of GnosisSafeL2, bound to a specific deployed contract.
func NewGnosisSafeL2Transactor(address common.Address, transactor bind.ContractTransactor) (*GnosisSafeL2Transactor, error) {
	contract, err := bindGnosisSafeL2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeL2Transactor{contract: contract}, nil
}

// NewGnosisSafeL2Filterer creates a new log filterer instance of GnosisSafeL2, bound to a specific deployed contract.
func NewGnosisSafeL2Filterer(address common.Address, filterer bind.ContractFilterer) (*GnosisSafeL2Filterer, error) {
	contract, err := bindGnosisSafeL2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeL2Filterer{contract: contract}, nil
}

// bindGnosisSafeL2 binds a generic wrapper to an already deployed contract.
func bindGnosisSafeL2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GnosisSafeL2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GnosisSafeL2 *GnosisSafeL2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GnosisSafeL2.Contract.GnosisSafeL2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GnosisSafeL2 *GnosisSafeL2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.GnosisSafeL2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GnosisSafeL2 *GnosisSafeL2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.GnosisSafeL2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GnosisSafeL2 *GnosisSafeL2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GnosisSafeL2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GnosisSafeL2 *GnosisSafeL2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GnosisSafeL2 *GnosisSafeL2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_GnosisSafeL2 *GnosisSafeL2Caller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GnosisSafeL2.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_GnosisSafeL2 *GnosisSafeL2Session) VERSION() (string, error) {
	return _GnosisSafeL2.Contract.VERSION(&_GnosisSafeL2.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_GnosisSafeL2 *GnosisSafeL2CallerSession) VERSION() (string, error) {
	return _GnosisSafeL2.Contract.VERSION(&_GnosisSafeL2.CallOpts)
}

// ApprovedHashes is a free data retrieval call binding the contract method 0x7d832974.
//
// Solidity: function approvedHashes(address , bytes32 ) view returns(uint256)
func (_GnosisSafeL2 *GnosisSafeL2Caller) ApprovedHashes(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _GnosisSafeL2.contract.Call(opts, &out, "approvedHashes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApprovedHashes is a free data retrieval call binding the contract method 0x7d832974.
//
// Solidity: function approvedHashes(address , bytes32 ) view returns(uint256)
func (_GnosisSafeL2 *GnosisSafeL2Session) ApprovedHashes(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _GnosisSafeL2.Contract.ApprovedHashes(&_GnosisSafeL2.CallOpts, arg0, arg1)
}

// ApprovedHashes is a free data retrieval call binding the contract method 0x7d832974.
//
// Solidity: function approvedHashes(address , bytes32 ) view returns(uint256)
func (_GnosisSafeL2 *GnosisSafeL2CallerSession) ApprovedHashes(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _GnosisSafeL2.Contract.ApprovedHashes(&_GnosisSafeL2.CallOpts, arg0, arg1)
}

// CheckNSignatures is a free data retrieval call binding the contract method 0x12fb68e0.
//
// Solidity: function checkNSignatures(bytes32 dataHash, bytes data, bytes signatures, uint256 requiredSignatures) view returns()
func (_GnosisSafeL2 *GnosisSafeL2Caller) CheckNSignatures(opts *bind.CallOpts, dataHash [32]byte, data []byte, signatures []byte, requiredSignatures *big.Int) error {
	var out []interface{}
	err := _GnosisSafeL2.contract.Call(opts, &out, "checkNSignatures", dataHash, data, signatures, requiredSignatures)

	if err != nil {
		return err
	}

	return err

}

// CheckNSignatures is a free data retrieval call binding the contract method 0x12fb68e0.
//
// Solidity: function checkNSignatures(bytes32 dataHash, bytes data, bytes signatures, uint256 requiredSignatures) view returns()
func (_GnosisSafeL2 *GnosisSafeL2Session) CheckNSignatures(dataHash [32]byte, data []byte, signatures []byte, requiredSignatures *big.Int) error {
	return _GnosisSafeL2.Contract.CheckNSignatures(&_GnosisSafeL2.CallOpts, dataHash, data, signatures, requiredSignatures)
}

// CheckNSignatures is a free data retrieval call binding the contract method 0x12fb68e0.
//
// Solidity: function checkNSignatures(bytes32 dataHash, bytes data, bytes signatures, uint256 requiredSignatures) view returns()
func (_GnosisSafeL2 *GnosisSafeL2CallerSession) CheckNSignatures(dataHash [32]byte, data []byte, signatures []byte, requiredSignatures *big.Int) error {
	return _GnosisSafeL2.Contract.CheckNSignatures(&_GnosisSafeL2.CallOpts, dataHash, data, signatures, requiredSignatures)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x934f3a11.
//
// Solidity: function checkSignatures(bytes32 dataHash, bytes data, bytes signatures) view returns()
func (_GnosisSafeL2 *GnosisSafeL2Caller) CheckSignatures(opts *bind.CallOpts, dataHash [32]byte, data []byte, signatures []byte) error {
	var out []interface{}
	err := _GnosisSafeL2.contract.Call(opts, &out, "checkSignatures", dataHash, data, signatures)

	if err != nil {
		return err
	}

	return err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x934f3a11.
//
// Solidity: function checkSignatures(bytes32 dataHash, bytes data, bytes signatures) view returns()
func (_GnosisSafeL2 *GnosisSafeL2Session) CheckSignatures(dataHash [32]byte, data []byte, signatures []byte) error {
	return _GnosisSafeL2.Contract.CheckSignatures(&_GnosisSafeL2.CallOpts, dataHash, data, signatures)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x934f3a11.
//
// Solidity: function checkSignatures(bytes32 dataHash, bytes data, bytes signatures) view returns()
func (_GnosisSafeL2 *GnosisSafeL2CallerSession) CheckSignatures(dataHash [32]byte, data []byte, signatures []byte) error {
	return _GnosisSafeL2.Contract.CheckSignatures(&_GnosisSafeL2.CallOpts, dataHash, data, signatures)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_GnosisSafeL2 *GnosisSafeL2Caller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GnosisSafeL2.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_GnosisSafeL2 *GnosisSafeL2Session) DomainSeparator() ([32]byte, error) {
	return _GnosisSafeL2.Contract.DomainSeparator(&_GnosisSafeL2.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_GnosisSafeL2 *GnosisSafeL2CallerSession) DomainSeparator() ([32]byte, error) {
	return _GnosisSafeL2.Contract.DomainSeparator(&_GnosisSafeL2.CallOpts)
}

// EncodeTransactionData is a free data retrieval call binding the contract method 0xe86637db.
//
// Solidity: function encodeTransactionData(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes)
func (_GnosisSafeL2 *GnosisSafeL2Caller) EncodeTransactionData(opts *bind.CallOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([]byte, error) {
	var out []interface{}
	err := _GnosisSafeL2.contract.Call(opts, &out, "encodeTransactionData", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeTransactionData is a free data retrieval call binding the contract method 0xe86637db.
//
// Solidity: function encodeTransactionData(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes)
func (_GnosisSafeL2 *GnosisSafeL2Session) EncodeTransactionData(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([]byte, error) {
	return _GnosisSafeL2.Contract.EncodeTransactionData(&_GnosisSafeL2.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// EncodeTransactionData is a free data retrieval call binding the contract method 0xe86637db.
//
// Solidity: function encodeTransactionData(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes)
func (_GnosisSafeL2 *GnosisSafeL2CallerSession) EncodeTransactionData(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([]byte, error) {
	return _GnosisSafeL2.Contract.EncodeTransactionData(&_GnosisSafeL2.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_GnosisSafeL2 *GnosisSafeL2Caller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GnosisSafeL2.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_GnosisSafeL2 *GnosisSafeL2Session) GetChainId() (*big.Int, error) {
	return _GnosisSafeL2.Contract.GetChainId(&_GnosisSafeL2.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_GnosisSafeL2 *GnosisSafeL2CallerSession) GetChainId() (*big.Int, error) {
	return _GnosisSafeL2.Contract.GetChainId(&_GnosisSafeL2.CallOpts)
}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_GnosisSafeL2 *GnosisSafeL2Caller) GetModulesPaginated(opts *bind.CallOpts, start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	var out []interface{}
	err := _GnosisSafeL2.contract.Call(opts, &out, "getModulesPaginated", start, pageSize)

	outstruct := new(struct {
		Array []common.Address
		Next  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Array = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Next = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_GnosisSafeL2 *GnosisSafeL2Session) GetModulesPaginated(start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _GnosisSafeL2.Contract.GetModulesPaginated(&_GnosisSafeL2.CallOpts, start, pageSize)
}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_GnosisSafeL2 *GnosisSafeL2CallerSession) GetModulesPaginated(start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _GnosisSafeL2.Contract.GetModulesPaginated(&_GnosisSafeL2.CallOpts, start, pageSize)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_GnosisSafeL2 *GnosisSafeL2Caller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GnosisSafeL2.contract.Call(opts, &out, "getOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_GnosisSafeL2 *GnosisSafeL2Session) GetOwners() ([]common.Address, error) {
	return _GnosisSafeL2.Contract.GetOwners(&_GnosisSafeL2.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_GnosisSafeL2 *GnosisSafeL2CallerSession) GetOwners() ([]common.Address, error) {
	return _GnosisSafeL2.Contract.GetOwners(&_GnosisSafeL2.CallOpts)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_GnosisSafeL2 *GnosisSafeL2Caller) GetStorageAt(opts *bind.CallOpts, offset *big.Int, length *big.Int) ([]byte, error) {
	var out []interface{}
	err := _GnosisSafeL2.contract.Call(opts, &out, "getStorageAt", offset, length)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_GnosisSafeL2 *GnosisSafeL2Session) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _GnosisSafeL2.Contract.GetStorageAt(&_GnosisSafeL2.CallOpts, offset, length)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_GnosisSafeL2 *GnosisSafeL2CallerSession) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _GnosisSafeL2.Contract.GetStorageAt(&_GnosisSafeL2.CallOpts, offset, length)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_GnosisSafeL2 *GnosisSafeL2Caller) GetThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GnosisSafeL2.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_GnosisSafeL2 *GnosisSafeL2Session) GetThreshold() (*big.Int, error) {
	return _GnosisSafeL2.Contract.GetThreshold(&_GnosisSafeL2.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_GnosisSafeL2 *GnosisSafeL2CallerSession) GetThreshold() (*big.Int, error) {
	return _GnosisSafeL2.Contract.GetThreshold(&_GnosisSafeL2.CallOpts)
}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32)
func (_GnosisSafeL2 *GnosisSafeL2Caller) GetTransactionHash(opts *bind.CallOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _GnosisSafeL2.contract.Call(opts, &out, "getTransactionHash", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32)
func (_GnosisSafeL2 *GnosisSafeL2Session) GetTransactionHash(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	return _GnosisSafeL2.Contract.GetTransactionHash(&_GnosisSafeL2.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32)
func (_GnosisSafeL2 *GnosisSafeL2CallerSession) GetTransactionHash(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	return _GnosisSafeL2.Contract.GetTransactionHash(&_GnosisSafeL2.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_GnosisSafeL2 *GnosisSafeL2Caller) IsModuleEnabled(opts *bind.CallOpts, module common.Address) (bool, error) {
	var out []interface{}
	err := _GnosisSafeL2.contract.Call(opts, &out, "isModuleEnabled", module)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_GnosisSafeL2 *GnosisSafeL2Session) IsModuleEnabled(module common.Address) (bool, error) {
	return _GnosisSafeL2.Contract.IsModuleEnabled(&_GnosisSafeL2.CallOpts, module)
}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_GnosisSafeL2 *GnosisSafeL2CallerSession) IsModuleEnabled(module common.Address) (bool, error) {
	return _GnosisSafeL2.Contract.IsModuleEnabled(&_GnosisSafeL2.CallOpts, module)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_GnosisSafeL2 *GnosisSafeL2Caller) IsOwner(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _GnosisSafeL2.contract.Call(opts, &out, "isOwner", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_GnosisSafeL2 *GnosisSafeL2Session) IsOwner(owner common.Address) (bool, error) {
	return _GnosisSafeL2.Contract.IsOwner(&_GnosisSafeL2.CallOpts, owner)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_GnosisSafeL2 *GnosisSafeL2CallerSession) IsOwner(owner common.Address) (bool, error) {
	return _GnosisSafeL2.Contract.IsOwner(&_GnosisSafeL2.CallOpts, owner)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_GnosisSafeL2 *GnosisSafeL2Caller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GnosisSafeL2.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_GnosisSafeL2 *GnosisSafeL2Session) Nonce() (*big.Int, error) {
	return _GnosisSafeL2.Contract.Nonce(&_GnosisSafeL2.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_GnosisSafeL2 *GnosisSafeL2CallerSession) Nonce() (*big.Int, error) {
	return _GnosisSafeL2.Contract.Nonce(&_GnosisSafeL2.CallOpts)
}

// SignedMessages is a free data retrieval call binding the contract method 0x5ae6bd37.
//
// Solidity: function signedMessages(bytes32 ) view returns(uint256)
func (_GnosisSafeL2 *GnosisSafeL2Caller) SignedMessages(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _GnosisSafeL2.contract.Call(opts, &out, "signedMessages", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignedMessages is a free data retrieval call binding the contract method 0x5ae6bd37.
//
// Solidity: function signedMessages(bytes32 ) view returns(uint256)
func (_GnosisSafeL2 *GnosisSafeL2Session) SignedMessages(arg0 [32]byte) (*big.Int, error) {
	return _GnosisSafeL2.Contract.SignedMessages(&_GnosisSafeL2.CallOpts, arg0)
}

// SignedMessages is a free data retrieval call binding the contract method 0x5ae6bd37.
//
// Solidity: function signedMessages(bytes32 ) view returns(uint256)
func (_GnosisSafeL2 *GnosisSafeL2CallerSession) SignedMessages(arg0 [32]byte) (*big.Int, error) {
	return _GnosisSafeL2.Contract.SignedMessages(&_GnosisSafeL2.CallOpts, arg0)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_GnosisSafeL2 *GnosisSafeL2Transactor) AddOwnerWithThreshold(opts *bind.TransactOpts, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafeL2.contract.Transact(opts, "addOwnerWithThreshold", owner, _threshold)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_GnosisSafeL2 *GnosisSafeL2Session) AddOwnerWithThreshold(owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.AddOwnerWithThreshold(&_GnosisSafeL2.TransactOpts, owner, _threshold)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_GnosisSafeL2 *GnosisSafeL2TransactorSession) AddOwnerWithThreshold(owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.AddOwnerWithThreshold(&_GnosisSafeL2.TransactOpts, owner, _threshold)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(bytes32 hashToApprove) returns()
func (_GnosisSafeL2 *GnosisSafeL2Transactor) ApproveHash(opts *bind.TransactOpts, hashToApprove [32]byte) (*types.Transaction, error) {
	return _GnosisSafeL2.contract.Transact(opts, "approveHash", hashToApprove)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(bytes32 hashToApprove) returns()
func (_GnosisSafeL2 *GnosisSafeL2Session) ApproveHash(hashToApprove [32]byte) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.ApproveHash(&_GnosisSafeL2.TransactOpts, hashToApprove)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(bytes32 hashToApprove) returns()
func (_GnosisSafeL2 *GnosisSafeL2TransactorSession) ApproveHash(hashToApprove [32]byte) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.ApproveHash(&_GnosisSafeL2.TransactOpts, hashToApprove)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_GnosisSafeL2 *GnosisSafeL2Transactor) ChangeThreshold(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafeL2.contract.Transact(opts, "changeThreshold", _threshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_GnosisSafeL2 *GnosisSafeL2Session) ChangeThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.ChangeThreshold(&_GnosisSafeL2.TransactOpts, _threshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_GnosisSafeL2 *GnosisSafeL2TransactorSession) ChangeThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.ChangeThreshold(&_GnosisSafeL2.TransactOpts, _threshold)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_GnosisSafeL2 *GnosisSafeL2Transactor) DisableModule(opts *bind.TransactOpts, prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _GnosisSafeL2.contract.Transact(opts, "disableModule", prevModule, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_GnosisSafeL2 *GnosisSafeL2Session) DisableModule(prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.DisableModule(&_GnosisSafeL2.TransactOpts, prevModule, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_GnosisSafeL2 *GnosisSafeL2TransactorSession) DisableModule(prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.DisableModule(&_GnosisSafeL2.TransactOpts, prevModule, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_GnosisSafeL2 *GnosisSafeL2Transactor) EnableModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _GnosisSafeL2.contract.Transact(opts, "enableModule", module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_GnosisSafeL2 *GnosisSafeL2Session) EnableModule(module common.Address) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.EnableModule(&_GnosisSafeL2.TransactOpts, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_GnosisSafeL2 *GnosisSafeL2TransactorSession) EnableModule(module common.Address) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.EnableModule(&_GnosisSafeL2.TransactOpts, module)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool)
func (_GnosisSafeL2 *GnosisSafeL2Transactor) ExecTransaction(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _GnosisSafeL2.contract.Transact(opts, "execTransaction", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool)
func (_GnosisSafeL2 *GnosisSafeL2Session) ExecTransaction(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.ExecTransaction(&_GnosisSafeL2.TransactOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool)
func (_GnosisSafeL2 *GnosisSafeL2TransactorSession) ExecTransaction(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.ExecTransaction(&_GnosisSafeL2.TransactOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_GnosisSafeL2 *GnosisSafeL2Transactor) ExecTransactionFromModule(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafeL2.contract.Transact(opts, "execTransactionFromModule", to, value, data, operation)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_GnosisSafeL2 *GnosisSafeL2Session) ExecTransactionFromModule(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.ExecTransactionFromModule(&_GnosisSafeL2.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_GnosisSafeL2 *GnosisSafeL2TransactorSession) ExecTransactionFromModule(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.ExecTransactionFromModule(&_GnosisSafeL2.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_GnosisSafeL2 *GnosisSafeL2Transactor) ExecTransactionFromModuleReturnData(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafeL2.contract.Transact(opts, "execTransactionFromModuleReturnData", to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_GnosisSafeL2 *GnosisSafeL2Session) ExecTransactionFromModuleReturnData(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.ExecTransactionFromModuleReturnData(&_GnosisSafeL2.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_GnosisSafeL2 *GnosisSafeL2TransactorSession) ExecTransactionFromModuleReturnData(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.ExecTransactionFromModuleReturnData(&_GnosisSafeL2.TransactOpts, to, value, data, operation)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_GnosisSafeL2 *GnosisSafeL2Transactor) RemoveOwner(opts *bind.TransactOpts, prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafeL2.contract.Transact(opts, "removeOwner", prevOwner, owner, _threshold)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_GnosisSafeL2 *GnosisSafeL2Session) RemoveOwner(prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.RemoveOwner(&_GnosisSafeL2.TransactOpts, prevOwner, owner, _threshold)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_GnosisSafeL2 *GnosisSafeL2TransactorSession) RemoveOwner(prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.RemoveOwner(&_GnosisSafeL2.TransactOpts, prevOwner, owner, _threshold)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_GnosisSafeL2 *GnosisSafeL2Transactor) SetFallbackHandler(opts *bind.TransactOpts, handler common.Address) (*types.Transaction, error) {
	return _GnosisSafeL2.contract.Transact(opts, "setFallbackHandler", handler)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_GnosisSafeL2 *GnosisSafeL2Session) SetFallbackHandler(handler common.Address) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.SetFallbackHandler(&_GnosisSafeL2.TransactOpts, handler)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_GnosisSafeL2 *GnosisSafeL2TransactorSession) SetFallbackHandler(handler common.Address) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.SetFallbackHandler(&_GnosisSafeL2.TransactOpts, handler)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_GnosisSafeL2 *GnosisSafeL2Transactor) SetGuard(opts *bind.TransactOpts, guard common.Address) (*types.Transaction, error) {
	return _GnosisSafeL2.contract.Transact(opts, "setGuard", guard)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_GnosisSafeL2 *GnosisSafeL2Session) SetGuard(guard common.Address) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.SetGuard(&_GnosisSafeL2.TransactOpts, guard)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_GnosisSafeL2 *GnosisSafeL2TransactorSession) SetGuard(guard common.Address) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.SetGuard(&_GnosisSafeL2.TransactOpts, guard)
}

// Setup is a paid mutator transaction binding the contract method 0xb63e800d.
//
// Solidity: function setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address paymentReceiver) returns()
func (_GnosisSafeL2 *GnosisSafeL2Transactor) Setup(opts *bind.TransactOpts, _owners []common.Address, _threshold *big.Int, to common.Address, data []byte, fallbackHandler common.Address, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) (*types.Transaction, error) {
	return _GnosisSafeL2.contract.Transact(opts, "setup", _owners, _threshold, to, data, fallbackHandler, paymentToken, payment, paymentReceiver)
}

// Setup is a paid mutator transaction binding the contract method 0xb63e800d.
//
// Solidity: function setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address paymentReceiver) returns()
func (_GnosisSafeL2 *GnosisSafeL2Session) Setup(_owners []common.Address, _threshold *big.Int, to common.Address, data []byte, fallbackHandler common.Address, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.Setup(&_GnosisSafeL2.TransactOpts, _owners, _threshold, to, data, fallbackHandler, paymentToken, payment, paymentReceiver)
}

// Setup is a paid mutator transaction binding the contract method 0xb63e800d.
//
// Solidity: function setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address paymentReceiver) returns()
func (_GnosisSafeL2 *GnosisSafeL2TransactorSession) Setup(_owners []common.Address, _threshold *big.Int, to common.Address, data []byte, fallbackHandler common.Address, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.Setup(&_GnosisSafeL2.TransactOpts, _owners, _threshold, to, data, fallbackHandler, paymentToken, payment, paymentReceiver)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_GnosisSafeL2 *GnosisSafeL2Transactor) SimulateAndRevert(opts *bind.TransactOpts, targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _GnosisSafeL2.contract.Transact(opts, "simulateAndRevert", targetContract, calldataPayload)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_GnosisSafeL2 *GnosisSafeL2Session) SimulateAndRevert(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.SimulateAndRevert(&_GnosisSafeL2.TransactOpts, targetContract, calldataPayload)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_GnosisSafeL2 *GnosisSafeL2TransactorSession) SimulateAndRevert(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.SimulateAndRevert(&_GnosisSafeL2.TransactOpts, targetContract, calldataPayload)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_GnosisSafeL2 *GnosisSafeL2Transactor) SwapOwner(opts *bind.TransactOpts, prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _GnosisSafeL2.contract.Transact(opts, "swapOwner", prevOwner, oldOwner, newOwner)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_GnosisSafeL2 *GnosisSafeL2Session) SwapOwner(prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.SwapOwner(&_GnosisSafeL2.TransactOpts, prevOwner, oldOwner, newOwner)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_GnosisSafeL2 *GnosisSafeL2TransactorSession) SwapOwner(prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.SwapOwner(&_GnosisSafeL2.TransactOpts, prevOwner, oldOwner, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_GnosisSafeL2 *GnosisSafeL2Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _GnosisSafeL2.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_GnosisSafeL2 *GnosisSafeL2Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.Fallback(&_GnosisSafeL2.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_GnosisSafeL2 *GnosisSafeL2TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.Fallback(&_GnosisSafeL2.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GnosisSafeL2 *GnosisSafeL2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GnosisSafeL2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GnosisSafeL2 *GnosisSafeL2Session) Receive() (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.Receive(&_GnosisSafeL2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GnosisSafeL2 *GnosisSafeL2TransactorSession) Receive() (*types.Transaction, error) {
	return _GnosisSafeL2.Contract.Receive(&_GnosisSafeL2.TransactOpts)
}

// GnosisSafeL2AddedOwnerIterator is returned from FilterAddedOwner and is used to iterate over the raw logs and unpacked data for AddedOwner events raised by the GnosisSafeL2 contract.
type GnosisSafeL2AddedOwnerIterator struct {
	Event *GnosisSafeL2AddedOwner // Event containing the contract specifics and raw log

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
func (it *GnosisSafeL2AddedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeL2AddedOwner)
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
		it.Event = new(GnosisSafeL2AddedOwner)
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
func (it *GnosisSafeL2AddedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeL2AddedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeL2AddedOwner represents a AddedOwner event raised by the GnosisSafeL2 contract.
type GnosisSafeL2AddedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddedOwner is a free log retrieval operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address indexed owner)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) FilterAddedOwner(opts *bind.FilterOpts, owner []common.Address) (*GnosisSafeL2AddedOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.FilterLogs(opts, "AddedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeL2AddedOwnerIterator{contract: _GnosisSafeL2.contract, event: "AddedOwner", logs: logs, sub: sub}, nil
}

// WatchAddedOwner is a free log subscription operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address indexed owner)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) WatchAddedOwner(opts *bind.WatchOpts, sink chan<- *GnosisSafeL2AddedOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.WatchLogs(opts, "AddedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeL2AddedOwner)
				if err := _GnosisSafeL2.contract.UnpackLog(event, "AddedOwner", log); err != nil {
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

// ParseAddedOwner is a log parse operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address indexed owner)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) ParseAddedOwner(log types.Log) (*GnosisSafeL2AddedOwner, error) {
	event := new(GnosisSafeL2AddedOwner)
	if err := _GnosisSafeL2.contract.UnpackLog(event, "AddedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeL2ApproveHashIterator is returned from FilterApproveHash and is used to iterate over the raw logs and unpacked data for ApproveHash events raised by the GnosisSafeL2 contract.
type GnosisSafeL2ApproveHashIterator struct {
	Event *GnosisSafeL2ApproveHash // Event containing the contract specifics and raw log

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
func (it *GnosisSafeL2ApproveHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeL2ApproveHash)
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
		it.Event = new(GnosisSafeL2ApproveHash)
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
func (it *GnosisSafeL2ApproveHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeL2ApproveHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeL2ApproveHash represents a ApproveHash event raised by the GnosisSafeL2 contract.
type GnosisSafeL2ApproveHash struct {
	ApprovedHash [32]byte
	Owner        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterApproveHash is a free log retrieval operation binding the contract event 0xf2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c.
//
// Solidity: event ApproveHash(bytes32 indexed approvedHash, address indexed owner)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) FilterApproveHash(opts *bind.FilterOpts, approvedHash [][32]byte, owner []common.Address) (*GnosisSafeL2ApproveHashIterator, error) {

	var approvedHashRule []interface{}
	for _, approvedHashItem := range approvedHash {
		approvedHashRule = append(approvedHashRule, approvedHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.FilterLogs(opts, "ApproveHash", approvedHashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeL2ApproveHashIterator{contract: _GnosisSafeL2.contract, event: "ApproveHash", logs: logs, sub: sub}, nil
}

// WatchApproveHash is a free log subscription operation binding the contract event 0xf2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c.
//
// Solidity: event ApproveHash(bytes32 indexed approvedHash, address indexed owner)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) WatchApproveHash(opts *bind.WatchOpts, sink chan<- *GnosisSafeL2ApproveHash, approvedHash [][32]byte, owner []common.Address) (event.Subscription, error) {

	var approvedHashRule []interface{}
	for _, approvedHashItem := range approvedHash {
		approvedHashRule = append(approvedHashRule, approvedHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.WatchLogs(opts, "ApproveHash", approvedHashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeL2ApproveHash)
				if err := _GnosisSafeL2.contract.UnpackLog(event, "ApproveHash", log); err != nil {
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

// ParseApproveHash is a log parse operation binding the contract event 0xf2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c.
//
// Solidity: event ApproveHash(bytes32 indexed approvedHash, address indexed owner)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) ParseApproveHash(log types.Log) (*GnosisSafeL2ApproveHash, error) {
	event := new(GnosisSafeL2ApproveHash)
	if err := _GnosisSafeL2.contract.UnpackLog(event, "ApproveHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeL2ChangedFallbackHandlerIterator is returned from FilterChangedFallbackHandler and is used to iterate over the raw logs and unpacked data for ChangedFallbackHandler events raised by the GnosisSafeL2 contract.
type GnosisSafeL2ChangedFallbackHandlerIterator struct {
	Event *GnosisSafeL2ChangedFallbackHandler // Event containing the contract specifics and raw log

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
func (it *GnosisSafeL2ChangedFallbackHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeL2ChangedFallbackHandler)
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
		it.Event = new(GnosisSafeL2ChangedFallbackHandler)
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
func (it *GnosisSafeL2ChangedFallbackHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeL2ChangedFallbackHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeL2ChangedFallbackHandler represents a ChangedFallbackHandler event raised by the GnosisSafeL2 contract.
type GnosisSafeL2ChangedFallbackHandler struct {
	Handler common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChangedFallbackHandler is a free log retrieval operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address indexed handler)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) FilterChangedFallbackHandler(opts *bind.FilterOpts, handler []common.Address) (*GnosisSafeL2ChangedFallbackHandlerIterator, error) {

	var handlerRule []interface{}
	for _, handlerItem := range handler {
		handlerRule = append(handlerRule, handlerItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.FilterLogs(opts, "ChangedFallbackHandler", handlerRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeL2ChangedFallbackHandlerIterator{contract: _GnosisSafeL2.contract, event: "ChangedFallbackHandler", logs: logs, sub: sub}, nil
}

// WatchChangedFallbackHandler is a free log subscription operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address indexed handler)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) WatchChangedFallbackHandler(opts *bind.WatchOpts, sink chan<- *GnosisSafeL2ChangedFallbackHandler, handler []common.Address) (event.Subscription, error) {

	var handlerRule []interface{}
	for _, handlerItem := range handler {
		handlerRule = append(handlerRule, handlerItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.WatchLogs(opts, "ChangedFallbackHandler", handlerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeL2ChangedFallbackHandler)
				if err := _GnosisSafeL2.contract.UnpackLog(event, "ChangedFallbackHandler", log); err != nil {
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

// ParseChangedFallbackHandler is a log parse operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address indexed handler)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) ParseChangedFallbackHandler(log types.Log) (*GnosisSafeL2ChangedFallbackHandler, error) {
	event := new(GnosisSafeL2ChangedFallbackHandler)
	if err := _GnosisSafeL2.contract.UnpackLog(event, "ChangedFallbackHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeL2ChangedGuardIterator is returned from FilterChangedGuard and is used to iterate over the raw logs and unpacked data for ChangedGuard events raised by the GnosisSafeL2 contract.
type GnosisSafeL2ChangedGuardIterator struct {
	Event *GnosisSafeL2ChangedGuard // Event containing the contract specifics and raw log

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
func (it *GnosisSafeL2ChangedGuardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeL2ChangedGuard)
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
		it.Event = new(GnosisSafeL2ChangedGuard)
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
func (it *GnosisSafeL2ChangedGuardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeL2ChangedGuardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeL2ChangedGuard represents a ChangedGuard event raised by the GnosisSafeL2 contract.
type GnosisSafeL2ChangedGuard struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterChangedGuard is a free log retrieval operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address indexed guard)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) FilterChangedGuard(opts *bind.FilterOpts, guard []common.Address) (*GnosisSafeL2ChangedGuardIterator, error) {

	var guardRule []interface{}
	for _, guardItem := range guard {
		guardRule = append(guardRule, guardItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.FilterLogs(opts, "ChangedGuard", guardRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeL2ChangedGuardIterator{contract: _GnosisSafeL2.contract, event: "ChangedGuard", logs: logs, sub: sub}, nil
}

// WatchChangedGuard is a free log subscription operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address indexed guard)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) WatchChangedGuard(opts *bind.WatchOpts, sink chan<- *GnosisSafeL2ChangedGuard, guard []common.Address) (event.Subscription, error) {

	var guardRule []interface{}
	for _, guardItem := range guard {
		guardRule = append(guardRule, guardItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.WatchLogs(opts, "ChangedGuard", guardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeL2ChangedGuard)
				if err := _GnosisSafeL2.contract.UnpackLog(event, "ChangedGuard", log); err != nil {
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

// ParseChangedGuard is a log parse operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address indexed guard)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) ParseChangedGuard(log types.Log) (*GnosisSafeL2ChangedGuard, error) {
	event := new(GnosisSafeL2ChangedGuard)
	if err := _GnosisSafeL2.contract.UnpackLog(event, "ChangedGuard", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeL2ChangedThresholdIterator is returned from FilterChangedThreshold and is used to iterate over the raw logs and unpacked data for ChangedThreshold events raised by the GnosisSafeL2 contract.
type GnosisSafeL2ChangedThresholdIterator struct {
	Event *GnosisSafeL2ChangedThreshold // Event containing the contract specifics and raw log

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
func (it *GnosisSafeL2ChangedThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeL2ChangedThreshold)
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
		it.Event = new(GnosisSafeL2ChangedThreshold)
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
func (it *GnosisSafeL2ChangedThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeL2ChangedThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeL2ChangedThreshold represents a ChangedThreshold event raised by the GnosisSafeL2 contract.
type GnosisSafeL2ChangedThreshold struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangedThreshold is a free log retrieval operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) FilterChangedThreshold(opts *bind.FilterOpts) (*GnosisSafeL2ChangedThresholdIterator, error) {

	logs, sub, err := _GnosisSafeL2.contract.FilterLogs(opts, "ChangedThreshold")
	if err != nil {
		return nil, err
	}
	return &GnosisSafeL2ChangedThresholdIterator{contract: _GnosisSafeL2.contract, event: "ChangedThreshold", logs: logs, sub: sub}, nil
}

// WatchChangedThreshold is a free log subscription operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) WatchChangedThreshold(opts *bind.WatchOpts, sink chan<- *GnosisSafeL2ChangedThreshold) (event.Subscription, error) {

	logs, sub, err := _GnosisSafeL2.contract.WatchLogs(opts, "ChangedThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeL2ChangedThreshold)
				if err := _GnosisSafeL2.contract.UnpackLog(event, "ChangedThreshold", log); err != nil {
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

// ParseChangedThreshold is a log parse operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) ParseChangedThreshold(log types.Log) (*GnosisSafeL2ChangedThreshold, error) {
	event := new(GnosisSafeL2ChangedThreshold)
	if err := _GnosisSafeL2.contract.UnpackLog(event, "ChangedThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeL2DisabledModuleIterator is returned from FilterDisabledModule and is used to iterate over the raw logs and unpacked data for DisabledModule events raised by the GnosisSafeL2 contract.
type GnosisSafeL2DisabledModuleIterator struct {
	Event *GnosisSafeL2DisabledModule // Event containing the contract specifics and raw log

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
func (it *GnosisSafeL2DisabledModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeL2DisabledModule)
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
		it.Event = new(GnosisSafeL2DisabledModule)
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
func (it *GnosisSafeL2DisabledModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeL2DisabledModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeL2DisabledModule represents a DisabledModule event raised by the GnosisSafeL2 contract.
type GnosisSafeL2DisabledModule struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDisabledModule is a free log retrieval operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address indexed module)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) FilterDisabledModule(opts *bind.FilterOpts, module []common.Address) (*GnosisSafeL2DisabledModuleIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.FilterLogs(opts, "DisabledModule", moduleRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeL2DisabledModuleIterator{contract: _GnosisSafeL2.contract, event: "DisabledModule", logs: logs, sub: sub}, nil
}

// WatchDisabledModule is a free log subscription operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address indexed module)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) WatchDisabledModule(opts *bind.WatchOpts, sink chan<- *GnosisSafeL2DisabledModule, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.WatchLogs(opts, "DisabledModule", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeL2DisabledModule)
				if err := _GnosisSafeL2.contract.UnpackLog(event, "DisabledModule", log); err != nil {
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

// ParseDisabledModule is a log parse operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address indexed module)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) ParseDisabledModule(log types.Log) (*GnosisSafeL2DisabledModule, error) {
	event := new(GnosisSafeL2DisabledModule)
	if err := _GnosisSafeL2.contract.UnpackLog(event, "DisabledModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeL2EnabledModuleIterator is returned from FilterEnabledModule and is used to iterate over the raw logs and unpacked data for EnabledModule events raised by the GnosisSafeL2 contract.
type GnosisSafeL2EnabledModuleIterator struct {
	Event *GnosisSafeL2EnabledModule // Event containing the contract specifics and raw log

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
func (it *GnosisSafeL2EnabledModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeL2EnabledModule)
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
		it.Event = new(GnosisSafeL2EnabledModule)
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
func (it *GnosisSafeL2EnabledModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeL2EnabledModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeL2EnabledModule represents a EnabledModule event raised by the GnosisSafeL2 contract.
type GnosisSafeL2EnabledModule struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnabledModule is a free log retrieval operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address indexed module)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) FilterEnabledModule(opts *bind.FilterOpts, module []common.Address) (*GnosisSafeL2EnabledModuleIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.FilterLogs(opts, "EnabledModule", moduleRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeL2EnabledModuleIterator{contract: _GnosisSafeL2.contract, event: "EnabledModule", logs: logs, sub: sub}, nil
}

// WatchEnabledModule is a free log subscription operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address indexed module)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) WatchEnabledModule(opts *bind.WatchOpts, sink chan<- *GnosisSafeL2EnabledModule, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.WatchLogs(opts, "EnabledModule", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeL2EnabledModule)
				if err := _GnosisSafeL2.contract.UnpackLog(event, "EnabledModule", log); err != nil {
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

// ParseEnabledModule is a log parse operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address indexed module)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) ParseEnabledModule(log types.Log) (*GnosisSafeL2EnabledModule, error) {
	event := new(GnosisSafeL2EnabledModule)
	if err := _GnosisSafeL2.contract.UnpackLog(event, "EnabledModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeL2ExecutionFailureIterator is returned from FilterExecutionFailure and is used to iterate over the raw logs and unpacked data for ExecutionFailure events raised by the GnosisSafeL2 contract.
type GnosisSafeL2ExecutionFailureIterator struct {
	Event *GnosisSafeL2ExecutionFailure // Event containing the contract specifics and raw log

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
func (it *GnosisSafeL2ExecutionFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeL2ExecutionFailure)
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
		it.Event = new(GnosisSafeL2ExecutionFailure)
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
func (it *GnosisSafeL2ExecutionFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeL2ExecutionFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeL2ExecutionFailure represents a ExecutionFailure event raised by the GnosisSafeL2 contract.
type GnosisSafeL2ExecutionFailure struct {
	TxHash  [32]byte
	Payment *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailure is a free log retrieval operation binding the contract event 0x23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23.
//
// Solidity: event ExecutionFailure(bytes32 indexed txHash, uint256 payment)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) FilterExecutionFailure(opts *bind.FilterOpts, txHash [][32]byte) (*GnosisSafeL2ExecutionFailureIterator, error) {

	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.FilterLogs(opts, "ExecutionFailure", txHashRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeL2ExecutionFailureIterator{contract: _GnosisSafeL2.contract, event: "ExecutionFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFailure is a free log subscription operation binding the contract event 0x23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23.
//
// Solidity: event ExecutionFailure(bytes32 indexed txHash, uint256 payment)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) WatchExecutionFailure(opts *bind.WatchOpts, sink chan<- *GnosisSafeL2ExecutionFailure, txHash [][32]byte) (event.Subscription, error) {

	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.WatchLogs(opts, "ExecutionFailure", txHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeL2ExecutionFailure)
				if err := _GnosisSafeL2.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
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

// ParseExecutionFailure is a log parse operation binding the contract event 0x23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23.
//
// Solidity: event ExecutionFailure(bytes32 indexed txHash, uint256 payment)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) ParseExecutionFailure(log types.Log) (*GnosisSafeL2ExecutionFailure, error) {
	event := new(GnosisSafeL2ExecutionFailure)
	if err := _GnosisSafeL2.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeL2ExecutionFromModuleFailureIterator is returned from FilterExecutionFromModuleFailure and is used to iterate over the raw logs and unpacked data for ExecutionFromModuleFailure events raised by the GnosisSafeL2 contract.
type GnosisSafeL2ExecutionFromModuleFailureIterator struct {
	Event *GnosisSafeL2ExecutionFromModuleFailure // Event containing the contract specifics and raw log

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
func (it *GnosisSafeL2ExecutionFromModuleFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeL2ExecutionFromModuleFailure)
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
		it.Event = new(GnosisSafeL2ExecutionFromModuleFailure)
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
func (it *GnosisSafeL2ExecutionFromModuleFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeL2ExecutionFromModuleFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeL2ExecutionFromModuleFailure represents a ExecutionFromModuleFailure event raised by the GnosisSafeL2 contract.
type GnosisSafeL2ExecutionFromModuleFailure struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutionFromModuleFailure is a free log retrieval operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) FilterExecutionFromModuleFailure(opts *bind.FilterOpts, module []common.Address) (*GnosisSafeL2ExecutionFromModuleFailureIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.FilterLogs(opts, "ExecutionFromModuleFailure", moduleRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeL2ExecutionFromModuleFailureIterator{contract: _GnosisSafeL2.contract, event: "ExecutionFromModuleFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFromModuleFailure is a free log subscription operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) WatchExecutionFromModuleFailure(opts *bind.WatchOpts, sink chan<- *GnosisSafeL2ExecutionFromModuleFailure, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.WatchLogs(opts, "ExecutionFromModuleFailure", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeL2ExecutionFromModuleFailure)
				if err := _GnosisSafeL2.contract.UnpackLog(event, "ExecutionFromModuleFailure", log); err != nil {
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

// ParseExecutionFromModuleFailure is a log parse operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) ParseExecutionFromModuleFailure(log types.Log) (*GnosisSafeL2ExecutionFromModuleFailure, error) {
	event := new(GnosisSafeL2ExecutionFromModuleFailure)
	if err := _GnosisSafeL2.contract.UnpackLog(event, "ExecutionFromModuleFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeL2ExecutionFromModuleSuccessIterator is returned from FilterExecutionFromModuleSuccess and is used to iterate over the raw logs and unpacked data for ExecutionFromModuleSuccess events raised by the GnosisSafeL2 contract.
type GnosisSafeL2ExecutionFromModuleSuccessIterator struct {
	Event *GnosisSafeL2ExecutionFromModuleSuccess // Event containing the contract specifics and raw log

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
func (it *GnosisSafeL2ExecutionFromModuleSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeL2ExecutionFromModuleSuccess)
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
		it.Event = new(GnosisSafeL2ExecutionFromModuleSuccess)
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
func (it *GnosisSafeL2ExecutionFromModuleSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeL2ExecutionFromModuleSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeL2ExecutionFromModuleSuccess represents a ExecutionFromModuleSuccess event raised by the GnosisSafeL2 contract.
type GnosisSafeL2ExecutionFromModuleSuccess struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutionFromModuleSuccess is a free log retrieval operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) FilterExecutionFromModuleSuccess(opts *bind.FilterOpts, module []common.Address) (*GnosisSafeL2ExecutionFromModuleSuccessIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.FilterLogs(opts, "ExecutionFromModuleSuccess", moduleRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeL2ExecutionFromModuleSuccessIterator{contract: _GnosisSafeL2.contract, event: "ExecutionFromModuleSuccess", logs: logs, sub: sub}, nil
}

// WatchExecutionFromModuleSuccess is a free log subscription operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) WatchExecutionFromModuleSuccess(opts *bind.WatchOpts, sink chan<- *GnosisSafeL2ExecutionFromModuleSuccess, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.WatchLogs(opts, "ExecutionFromModuleSuccess", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeL2ExecutionFromModuleSuccess)
				if err := _GnosisSafeL2.contract.UnpackLog(event, "ExecutionFromModuleSuccess", log); err != nil {
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

// ParseExecutionFromModuleSuccess is a log parse operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) ParseExecutionFromModuleSuccess(log types.Log) (*GnosisSafeL2ExecutionFromModuleSuccess, error) {
	event := new(GnosisSafeL2ExecutionFromModuleSuccess)
	if err := _GnosisSafeL2.contract.UnpackLog(event, "ExecutionFromModuleSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeL2ExecutionSuccessIterator is returned from FilterExecutionSuccess and is used to iterate over the raw logs and unpacked data for ExecutionSuccess events raised by the GnosisSafeL2 contract.
type GnosisSafeL2ExecutionSuccessIterator struct {
	Event *GnosisSafeL2ExecutionSuccess // Event containing the contract specifics and raw log

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
func (it *GnosisSafeL2ExecutionSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeL2ExecutionSuccess)
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
		it.Event = new(GnosisSafeL2ExecutionSuccess)
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
func (it *GnosisSafeL2ExecutionSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeL2ExecutionSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeL2ExecutionSuccess represents a ExecutionSuccess event raised by the GnosisSafeL2 contract.
type GnosisSafeL2ExecutionSuccess struct {
	TxHash  [32]byte
	Payment *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecutionSuccess is a free log retrieval operation binding the contract event 0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e.
//
// Solidity: event ExecutionSuccess(bytes32 indexed txHash, uint256 payment)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) FilterExecutionSuccess(opts *bind.FilterOpts, txHash [][32]byte) (*GnosisSafeL2ExecutionSuccessIterator, error) {

	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.FilterLogs(opts, "ExecutionSuccess", txHashRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeL2ExecutionSuccessIterator{contract: _GnosisSafeL2.contract, event: "ExecutionSuccess", logs: logs, sub: sub}, nil
}

// WatchExecutionSuccess is a free log subscription operation binding the contract event 0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e.
//
// Solidity: event ExecutionSuccess(bytes32 indexed txHash, uint256 payment)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) WatchExecutionSuccess(opts *bind.WatchOpts, sink chan<- *GnosisSafeL2ExecutionSuccess, txHash [][32]byte) (event.Subscription, error) {

	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.WatchLogs(opts, "ExecutionSuccess", txHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeL2ExecutionSuccess)
				if err := _GnosisSafeL2.contract.UnpackLog(event, "ExecutionSuccess", log); err != nil {
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

// ParseExecutionSuccess is a log parse operation binding the contract event 0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e.
//
// Solidity: event ExecutionSuccess(bytes32 indexed txHash, uint256 payment)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) ParseExecutionSuccess(log types.Log) (*GnosisSafeL2ExecutionSuccess, error) {
	event := new(GnosisSafeL2ExecutionSuccess)
	if err := _GnosisSafeL2.contract.UnpackLog(event, "ExecutionSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeL2RemovedOwnerIterator is returned from FilterRemovedOwner and is used to iterate over the raw logs and unpacked data for RemovedOwner events raised by the GnosisSafeL2 contract.
type GnosisSafeL2RemovedOwnerIterator struct {
	Event *GnosisSafeL2RemovedOwner // Event containing the contract specifics and raw log

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
func (it *GnosisSafeL2RemovedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeL2RemovedOwner)
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
		it.Event = new(GnosisSafeL2RemovedOwner)
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
func (it *GnosisSafeL2RemovedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeL2RemovedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeL2RemovedOwner represents a RemovedOwner event raised by the GnosisSafeL2 contract.
type GnosisSafeL2RemovedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRemovedOwner is a free log retrieval operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address indexed owner)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) FilterRemovedOwner(opts *bind.FilterOpts, owner []common.Address) (*GnosisSafeL2RemovedOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.FilterLogs(opts, "RemovedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeL2RemovedOwnerIterator{contract: _GnosisSafeL2.contract, event: "RemovedOwner", logs: logs, sub: sub}, nil
}

// WatchRemovedOwner is a free log subscription operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address indexed owner)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) WatchRemovedOwner(opts *bind.WatchOpts, sink chan<- *GnosisSafeL2RemovedOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.WatchLogs(opts, "RemovedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeL2RemovedOwner)
				if err := _GnosisSafeL2.contract.UnpackLog(event, "RemovedOwner", log); err != nil {
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

// ParseRemovedOwner is a log parse operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address indexed owner)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) ParseRemovedOwner(log types.Log) (*GnosisSafeL2RemovedOwner, error) {
	event := new(GnosisSafeL2RemovedOwner)
	if err := _GnosisSafeL2.contract.UnpackLog(event, "RemovedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeL2SafeReceivedIterator is returned from FilterSafeReceived and is used to iterate over the raw logs and unpacked data for SafeReceived events raised by the GnosisSafeL2 contract.
type GnosisSafeL2SafeReceivedIterator struct {
	Event *GnosisSafeL2SafeReceived // Event containing the contract specifics and raw log

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
func (it *GnosisSafeL2SafeReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeL2SafeReceived)
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
		it.Event = new(GnosisSafeL2SafeReceived)
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
func (it *GnosisSafeL2SafeReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeL2SafeReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeL2SafeReceived represents a SafeReceived event raised by the GnosisSafeL2 contract.
type GnosisSafeL2SafeReceived struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSafeReceived is a free log retrieval operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) FilterSafeReceived(opts *bind.FilterOpts, sender []common.Address) (*GnosisSafeL2SafeReceivedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.FilterLogs(opts, "SafeReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeL2SafeReceivedIterator{contract: _GnosisSafeL2.contract, event: "SafeReceived", logs: logs, sub: sub}, nil
}

// WatchSafeReceived is a free log subscription operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) WatchSafeReceived(opts *bind.WatchOpts, sink chan<- *GnosisSafeL2SafeReceived, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.WatchLogs(opts, "SafeReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeL2SafeReceived)
				if err := _GnosisSafeL2.contract.UnpackLog(event, "SafeReceived", log); err != nil {
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

// ParseSafeReceived is a log parse operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) ParseSafeReceived(log types.Log) (*GnosisSafeL2SafeReceived, error) {
	event := new(GnosisSafeL2SafeReceived)
	if err := _GnosisSafeL2.contract.UnpackLog(event, "SafeReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeL2SafeSetupIterator is returned from FilterSafeSetup and is used to iterate over the raw logs and unpacked data for SafeSetup events raised by the GnosisSafeL2 contract.
type GnosisSafeL2SafeSetupIterator struct {
	Event *GnosisSafeL2SafeSetup // Event containing the contract specifics and raw log

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
func (it *GnosisSafeL2SafeSetupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeL2SafeSetup)
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
		it.Event = new(GnosisSafeL2SafeSetup)
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
func (it *GnosisSafeL2SafeSetupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeL2SafeSetupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeL2SafeSetup represents a SafeSetup event raised by the GnosisSafeL2 contract.
type GnosisSafeL2SafeSetup struct {
	Initiator       common.Address
	Owners          []common.Address
	Threshold       *big.Int
	Initializer     common.Address
	FallbackHandler common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSafeSetup is a free log retrieval operation binding the contract event 0x141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a8.
//
// Solidity: event SafeSetup(address indexed initiator, address[] owners, uint256 threshold, address initializer, address fallbackHandler)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) FilterSafeSetup(opts *bind.FilterOpts, initiator []common.Address) (*GnosisSafeL2SafeSetupIterator, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.FilterLogs(opts, "SafeSetup", initiatorRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeL2SafeSetupIterator{contract: _GnosisSafeL2.contract, event: "SafeSetup", logs: logs, sub: sub}, nil
}

// WatchSafeSetup is a free log subscription operation binding the contract event 0x141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a8.
//
// Solidity: event SafeSetup(address indexed initiator, address[] owners, uint256 threshold, address initializer, address fallbackHandler)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) WatchSafeSetup(opts *bind.WatchOpts, sink chan<- *GnosisSafeL2SafeSetup, initiator []common.Address) (event.Subscription, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.WatchLogs(opts, "SafeSetup", initiatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeL2SafeSetup)
				if err := _GnosisSafeL2.contract.UnpackLog(event, "SafeSetup", log); err != nil {
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

// ParseSafeSetup is a log parse operation binding the contract event 0x141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a8.
//
// Solidity: event SafeSetup(address indexed initiator, address[] owners, uint256 threshold, address initializer, address fallbackHandler)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) ParseSafeSetup(log types.Log) (*GnosisSafeL2SafeSetup, error) {
	event := new(GnosisSafeL2SafeSetup)
	if err := _GnosisSafeL2.contract.UnpackLog(event, "SafeSetup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeL2SignMsgIterator is returned from FilterSignMsg and is used to iterate over the raw logs and unpacked data for SignMsg events raised by the GnosisSafeL2 contract.
type GnosisSafeL2SignMsgIterator struct {
	Event *GnosisSafeL2SignMsg // Event containing the contract specifics and raw log

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
func (it *GnosisSafeL2SignMsgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeL2SignMsg)
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
		it.Event = new(GnosisSafeL2SignMsg)
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
func (it *GnosisSafeL2SignMsgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeL2SignMsgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeL2SignMsg represents a SignMsg event raised by the GnosisSafeL2 contract.
type GnosisSafeL2SignMsg struct {
	MsgHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignMsg is a free log retrieval operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) FilterSignMsg(opts *bind.FilterOpts, msgHash [][32]byte) (*GnosisSafeL2SignMsgIterator, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.FilterLogs(opts, "SignMsg", msgHashRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeL2SignMsgIterator{contract: _GnosisSafeL2.contract, event: "SignMsg", logs: logs, sub: sub}, nil
}

// WatchSignMsg is a free log subscription operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) WatchSignMsg(opts *bind.WatchOpts, sink chan<- *GnosisSafeL2SignMsg, msgHash [][32]byte) (event.Subscription, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _GnosisSafeL2.contract.WatchLogs(opts, "SignMsg", msgHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeL2SignMsg)
				if err := _GnosisSafeL2.contract.UnpackLog(event, "SignMsg", log); err != nil {
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

// ParseSignMsg is a log parse operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_GnosisSafeL2 *GnosisSafeL2Filterer) ParseSignMsg(log types.Log) (*GnosisSafeL2SignMsg, error) {
	event := new(GnosisSafeL2SignMsg)
	if err := _GnosisSafeL2.contract.UnpackLog(event, "SignMsg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
