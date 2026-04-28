// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ctf_collateral_adapter

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

// CtfCollateralAdapterMetaData contains all meta data concerning the CtfCollateralAdapter contract.
var CtfCollateralAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_conditionalTokens\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdce\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpiredDeadline\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAsset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoHandoverRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyUnpaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"RolesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COLLATERAL_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONDITIONAL_TOKENS\",\"outputs\":[{\"internalType\":\"contractIConditionalTokens\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDCE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"completeOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"grantRoles\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"hasAllRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"hasAnyRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_conditionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mergePositions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"result\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"ownershipHandoverExpiresAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_conditionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"redeemPositions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"renounceRoles\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"revokeRoles\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"rolesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_conditionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"splitPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CtfCollateralAdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use CtfCollateralAdapterMetaData.ABI instead.
var CtfCollateralAdapterABI = CtfCollateralAdapterMetaData.ABI

// CtfCollateralAdapter is an auto generated Go binding around an Ethereum contract.
type CtfCollateralAdapter struct {
	CtfCollateralAdapterCaller     // Read-only binding to the contract
	CtfCollateralAdapterTransactor // Write-only binding to the contract
	CtfCollateralAdapterFilterer   // Log filterer for contract events
}

// CtfCollateralAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CtfCollateralAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CtfCollateralAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CtfCollateralAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CtfCollateralAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CtfCollateralAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CtfCollateralAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CtfCollateralAdapterSession struct {
	Contract     *CtfCollateralAdapter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CtfCollateralAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CtfCollateralAdapterCallerSession struct {
	Contract *CtfCollateralAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// CtfCollateralAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CtfCollateralAdapterTransactorSession struct {
	Contract     *CtfCollateralAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// CtfCollateralAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CtfCollateralAdapterRaw struct {
	Contract *CtfCollateralAdapter // Generic contract binding to access the raw methods on
}

// CtfCollateralAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CtfCollateralAdapterCallerRaw struct {
	Contract *CtfCollateralAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// CtfCollateralAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CtfCollateralAdapterTransactorRaw struct {
	Contract *CtfCollateralAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCtfCollateralAdapter creates a new instance of CtfCollateralAdapter, bound to a specific deployed contract.
func NewCtfCollateralAdapter(address common.Address, backend bind.ContractBackend) (*CtfCollateralAdapter, error) {
	contract, err := bindCtfCollateralAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CtfCollateralAdapter{CtfCollateralAdapterCaller: CtfCollateralAdapterCaller{contract: contract}, CtfCollateralAdapterTransactor: CtfCollateralAdapterTransactor{contract: contract}, CtfCollateralAdapterFilterer: CtfCollateralAdapterFilterer{contract: contract}}, nil
}

// NewCtfCollateralAdapterCaller creates a new read-only instance of CtfCollateralAdapter, bound to a specific deployed contract.
func NewCtfCollateralAdapterCaller(address common.Address, caller bind.ContractCaller) (*CtfCollateralAdapterCaller, error) {
	contract, err := bindCtfCollateralAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CtfCollateralAdapterCaller{contract: contract}, nil
}

// NewCtfCollateralAdapterTransactor creates a new write-only instance of CtfCollateralAdapter, bound to a specific deployed contract.
func NewCtfCollateralAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*CtfCollateralAdapterTransactor, error) {
	contract, err := bindCtfCollateralAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CtfCollateralAdapterTransactor{contract: contract}, nil
}

// NewCtfCollateralAdapterFilterer creates a new log filterer instance of CtfCollateralAdapter, bound to a specific deployed contract.
func NewCtfCollateralAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*CtfCollateralAdapterFilterer, error) {
	contract, err := bindCtfCollateralAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CtfCollateralAdapterFilterer{contract: contract}, nil
}

// bindCtfCollateralAdapter binds a generic wrapper to an already deployed contract.
func bindCtfCollateralAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CtfCollateralAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CtfCollateralAdapter *CtfCollateralAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CtfCollateralAdapter.Contract.CtfCollateralAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CtfCollateralAdapter *CtfCollateralAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.CtfCollateralAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CtfCollateralAdapter *CtfCollateralAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.CtfCollateralAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CtfCollateralAdapter *CtfCollateralAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CtfCollateralAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.contract.Transact(opts, method, params...)
}

// COLLATERALTOKEN is a free data retrieval call binding the contract method 0xf5f1f1a7.
//
// Solidity: function COLLATERAL_TOKEN() view returns(address)
func (_CtfCollateralAdapter *CtfCollateralAdapterCaller) COLLATERALTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CtfCollateralAdapter.contract.Call(opts, &out, "COLLATERAL_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// COLLATERALTOKEN is a free data retrieval call binding the contract method 0xf5f1f1a7.
//
// Solidity: function COLLATERAL_TOKEN() view returns(address)
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) COLLATERALTOKEN() (common.Address, error) {
	return _CtfCollateralAdapter.Contract.COLLATERALTOKEN(&_CtfCollateralAdapter.CallOpts)
}

// COLLATERALTOKEN is a free data retrieval call binding the contract method 0xf5f1f1a7.
//
// Solidity: function COLLATERAL_TOKEN() view returns(address)
func (_CtfCollateralAdapter *CtfCollateralAdapterCallerSession) COLLATERALTOKEN() (common.Address, error) {
	return _CtfCollateralAdapter.Contract.COLLATERALTOKEN(&_CtfCollateralAdapter.CallOpts)
}

// CONDITIONALTOKENS is a free data retrieval call binding the contract method 0x165d1f36.
//
// Solidity: function CONDITIONAL_TOKENS() view returns(address)
func (_CtfCollateralAdapter *CtfCollateralAdapterCaller) CONDITIONALTOKENS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CtfCollateralAdapter.contract.Call(opts, &out, "CONDITIONAL_TOKENS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CONDITIONALTOKENS is a free data retrieval call binding the contract method 0x165d1f36.
//
// Solidity: function CONDITIONAL_TOKENS() view returns(address)
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) CONDITIONALTOKENS() (common.Address, error) {
	return _CtfCollateralAdapter.Contract.CONDITIONALTOKENS(&_CtfCollateralAdapter.CallOpts)
}

// CONDITIONALTOKENS is a free data retrieval call binding the contract method 0x165d1f36.
//
// Solidity: function CONDITIONAL_TOKENS() view returns(address)
func (_CtfCollateralAdapter *CtfCollateralAdapterCallerSession) CONDITIONALTOKENS() (common.Address, error) {
	return _CtfCollateralAdapter.Contract.CONDITIONALTOKENS(&_CtfCollateralAdapter.CallOpts)
}

// USDCE is a free data retrieval call binding the contract method 0x195187e1.
//
// Solidity: function USDCE() view returns(address)
func (_CtfCollateralAdapter *CtfCollateralAdapterCaller) USDCE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CtfCollateralAdapter.contract.Call(opts, &out, "USDCE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDCE is a free data retrieval call binding the contract method 0x195187e1.
//
// Solidity: function USDCE() view returns(address)
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) USDCE() (common.Address, error) {
	return _CtfCollateralAdapter.Contract.USDCE(&_CtfCollateralAdapter.CallOpts)
}

// USDCE is a free data retrieval call binding the contract method 0x195187e1.
//
// Solidity: function USDCE() view returns(address)
func (_CtfCollateralAdapter *CtfCollateralAdapterCallerSession) USDCE() (common.Address, error) {
	return _CtfCollateralAdapter.Contract.USDCE(&_CtfCollateralAdapter.CallOpts)
}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_CtfCollateralAdapter *CtfCollateralAdapterCaller) HasAllRoles(opts *bind.CallOpts, user common.Address, roles *big.Int) (bool, error) {
	var out []interface{}
	err := _CtfCollateralAdapter.contract.Call(opts, &out, "hasAllRoles", user, roles)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) HasAllRoles(user common.Address, roles *big.Int) (bool, error) {
	return _CtfCollateralAdapter.Contract.HasAllRoles(&_CtfCollateralAdapter.CallOpts, user, roles)
}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_CtfCollateralAdapter *CtfCollateralAdapterCallerSession) HasAllRoles(user common.Address, roles *big.Int) (bool, error) {
	return _CtfCollateralAdapter.Contract.HasAllRoles(&_CtfCollateralAdapter.CallOpts, user, roles)
}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_CtfCollateralAdapter *CtfCollateralAdapterCaller) HasAnyRole(opts *bind.CallOpts, user common.Address, roles *big.Int) (bool, error) {
	var out []interface{}
	err := _CtfCollateralAdapter.contract.Call(opts, &out, "hasAnyRole", user, roles)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) HasAnyRole(user common.Address, roles *big.Int) (bool, error) {
	return _CtfCollateralAdapter.Contract.HasAnyRole(&_CtfCollateralAdapter.CallOpts, user, roles)
}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_CtfCollateralAdapter *CtfCollateralAdapterCallerSession) HasAnyRole(user common.Address, roles *big.Int) (bool, error) {
	return _CtfCollateralAdapter.Contract.HasAnyRole(&_CtfCollateralAdapter.CallOpts, user, roles)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_CtfCollateralAdapter *CtfCollateralAdapterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CtfCollateralAdapter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) Owner() (common.Address, error) {
	return _CtfCollateralAdapter.Contract.Owner(&_CtfCollateralAdapter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_CtfCollateralAdapter *CtfCollateralAdapterCallerSession) Owner() (common.Address, error) {
	return _CtfCollateralAdapter.Contract.Owner(&_CtfCollateralAdapter.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_CtfCollateralAdapter *CtfCollateralAdapterCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CtfCollateralAdapter.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _CtfCollateralAdapter.Contract.OwnershipHandoverExpiresAt(&_CtfCollateralAdapter.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_CtfCollateralAdapter *CtfCollateralAdapterCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _CtfCollateralAdapter.Contract.OwnershipHandoverExpiresAt(&_CtfCollateralAdapter.CallOpts, pendingOwner)
}

// Paused is a free data retrieval call binding the contract method 0x2e48152c.
//
// Solidity: function paused(address ) view returns(bool)
func (_CtfCollateralAdapter *CtfCollateralAdapterCaller) Paused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CtfCollateralAdapter.contract.Call(opts, &out, "paused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x2e48152c.
//
// Solidity: function paused(address ) view returns(bool)
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) Paused(arg0 common.Address) (bool, error) {
	return _CtfCollateralAdapter.Contract.Paused(&_CtfCollateralAdapter.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x2e48152c.
//
// Solidity: function paused(address ) view returns(bool)
func (_CtfCollateralAdapter *CtfCollateralAdapterCallerSession) Paused(arg0 common.Address) (bool, error) {
	return _CtfCollateralAdapter.Contract.Paused(&_CtfCollateralAdapter.CallOpts, arg0)
}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_CtfCollateralAdapter *CtfCollateralAdapterCaller) RolesOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CtfCollateralAdapter.contract.Call(opts, &out, "rolesOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) RolesOf(user common.Address) (*big.Int, error) {
	return _CtfCollateralAdapter.Contract.RolesOf(&_CtfCollateralAdapter.CallOpts, user)
}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_CtfCollateralAdapter *CtfCollateralAdapterCallerSession) RolesOf(user common.Address) (*big.Int, error) {
	return _CtfCollateralAdapter.Contract.RolesOf(&_CtfCollateralAdapter.CallOpts, user)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_CtfCollateralAdapter *CtfCollateralAdapterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CtfCollateralAdapter.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CtfCollateralAdapter.Contract.SupportsInterface(&_CtfCollateralAdapter.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_CtfCollateralAdapter *CtfCollateralAdapterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CtfCollateralAdapter.Contract.SupportsInterface(&_CtfCollateralAdapter.CallOpts, interfaceId)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactor) AddAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _CtfCollateralAdapter.contract.Transact(opts, "addAdmin", _admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) AddAdmin(_admin common.Address) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.AddAdmin(&_CtfCollateralAdapter.TransactOpts, _admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactorSession) AddAdmin(_admin common.Address) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.AddAdmin(&_CtfCollateralAdapter.TransactOpts, _admin)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CtfCollateralAdapter.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.CancelOwnershipHandover(&_CtfCollateralAdapter.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.CancelOwnershipHandover(&_CtfCollateralAdapter.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _CtfCollateralAdapter.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.CompleteOwnershipHandover(&_CtfCollateralAdapter.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.CompleteOwnershipHandover(&_CtfCollateralAdapter.TransactOpts, pendingOwner)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactor) GrantRoles(opts *bind.TransactOpts, user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _CtfCollateralAdapter.contract.Transact(opts, "grantRoles", user, roles)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) GrantRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.GrantRoles(&_CtfCollateralAdapter.TransactOpts, user, roles)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactorSession) GrantRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.GrantRoles(&_CtfCollateralAdapter.TransactOpts, user, roles)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address , bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactor) MergePositions(opts *bind.TransactOpts, arg0 common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _CtfCollateralAdapter.contract.Transact(opts, "mergePositions", arg0, arg1, _conditionId, arg3, _amount)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address , bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) MergePositions(arg0 common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.MergePositions(&_CtfCollateralAdapter.TransactOpts, arg0, arg1, _conditionId, arg3, _amount)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address , bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactorSession) MergePositions(arg0 common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.MergePositions(&_CtfCollateralAdapter.TransactOpts, arg0, arg1, _conditionId, arg3, _amount)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CtfCollateralAdapter.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.OnERC1155BatchReceived(&_CtfCollateralAdapter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.OnERC1155BatchReceived(&_CtfCollateralAdapter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CtfCollateralAdapter.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.OnERC1155Received(&_CtfCollateralAdapter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.OnERC1155Received(&_CtfCollateralAdapter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address _asset) returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactor) Pause(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _CtfCollateralAdapter.contract.Transact(opts, "pause", _asset)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address _asset) returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) Pause(_asset common.Address) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.Pause(&_CtfCollateralAdapter.TransactOpts, _asset)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address _asset) returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactorSession) Pause(_asset common.Address) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.Pause(&_CtfCollateralAdapter.TransactOpts, _asset)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0x01b7037c.
//
// Solidity: function redeemPositions(address , bytes32 , bytes32 _conditionId, uint256[] ) returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactor) RedeemPositions(opts *bind.TransactOpts, arg0 common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int) (*types.Transaction, error) {
	return _CtfCollateralAdapter.contract.Transact(opts, "redeemPositions", arg0, arg1, _conditionId, arg3)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0x01b7037c.
//
// Solidity: function redeemPositions(address , bytes32 , bytes32 _conditionId, uint256[] ) returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) RedeemPositions(arg0 common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.RedeemPositions(&_CtfCollateralAdapter.TransactOpts, arg0, arg1, _conditionId, arg3)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0x01b7037c.
//
// Solidity: function redeemPositions(address , bytes32 , bytes32 _conditionId, uint256[] ) returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactorSession) RedeemPositions(arg0 common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.RedeemPositions(&_CtfCollateralAdapter.TransactOpts, arg0, arg1, _conditionId, arg3)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactor) RemoveAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _CtfCollateralAdapter.contract.Transact(opts, "removeAdmin", _admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) RemoveAdmin(_admin common.Address) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.RemoveAdmin(&_CtfCollateralAdapter.TransactOpts, _admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactorSession) RemoveAdmin(_admin common.Address) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.RemoveAdmin(&_CtfCollateralAdapter.TransactOpts, _admin)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CtfCollateralAdapter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) RenounceOwnership() (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.RenounceOwnership(&_CtfCollateralAdapter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.RenounceOwnership(&_CtfCollateralAdapter.TransactOpts)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactor) RenounceRoles(opts *bind.TransactOpts, roles *big.Int) (*types.Transaction, error) {
	return _CtfCollateralAdapter.contract.Transact(opts, "renounceRoles", roles)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) RenounceRoles(roles *big.Int) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.RenounceRoles(&_CtfCollateralAdapter.TransactOpts, roles)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactorSession) RenounceRoles(roles *big.Int) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.RenounceRoles(&_CtfCollateralAdapter.TransactOpts, roles)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CtfCollateralAdapter.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.RequestOwnershipHandover(&_CtfCollateralAdapter.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.RequestOwnershipHandover(&_CtfCollateralAdapter.TransactOpts)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactor) RevokeRoles(opts *bind.TransactOpts, user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _CtfCollateralAdapter.contract.Transact(opts, "revokeRoles", user, roles)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) RevokeRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.RevokeRoles(&_CtfCollateralAdapter.TransactOpts, user, roles)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactorSession) RevokeRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.RevokeRoles(&_CtfCollateralAdapter.TransactOpts, user, roles)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address , bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactor) SplitPosition(opts *bind.TransactOpts, arg0 common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _CtfCollateralAdapter.contract.Transact(opts, "splitPosition", arg0, arg1, _conditionId, arg3, _amount)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address , bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) SplitPosition(arg0 common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.SplitPosition(&_CtfCollateralAdapter.TransactOpts, arg0, arg1, _conditionId, arg3, _amount)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address , bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactorSession) SplitPosition(arg0 common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.SplitPosition(&_CtfCollateralAdapter.TransactOpts, arg0, arg1, _conditionId, arg3, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CtfCollateralAdapter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.TransferOwnership(&_CtfCollateralAdapter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.TransferOwnership(&_CtfCollateralAdapter.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x57b001f9.
//
// Solidity: function unpause(address _asset) returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactor) Unpause(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _CtfCollateralAdapter.contract.Transact(opts, "unpause", _asset)
}

// Unpause is a paid mutator transaction binding the contract method 0x57b001f9.
//
// Solidity: function unpause(address _asset) returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterSession) Unpause(_asset common.Address) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.Unpause(&_CtfCollateralAdapter.TransactOpts, _asset)
}

// Unpause is a paid mutator transaction binding the contract method 0x57b001f9.
//
// Solidity: function unpause(address _asset) returns()
func (_CtfCollateralAdapter *CtfCollateralAdapterTransactorSession) Unpause(_asset common.Address) (*types.Transaction, error) {
	return _CtfCollateralAdapter.Contract.Unpause(&_CtfCollateralAdapter.TransactOpts, _asset)
}

// CtfCollateralAdapterOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the CtfCollateralAdapter contract.
type CtfCollateralAdapterOwnershipHandoverCanceledIterator struct {
	Event *CtfCollateralAdapterOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *CtfCollateralAdapterOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CtfCollateralAdapterOwnershipHandoverCanceled)
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
		it.Event = new(CtfCollateralAdapterOwnershipHandoverCanceled)
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
func (it *CtfCollateralAdapterOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CtfCollateralAdapterOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CtfCollateralAdapterOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the CtfCollateralAdapter contract.
type CtfCollateralAdapterOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_CtfCollateralAdapter *CtfCollateralAdapterFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*CtfCollateralAdapterOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _CtfCollateralAdapter.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CtfCollateralAdapterOwnershipHandoverCanceledIterator{contract: _CtfCollateralAdapter.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_CtfCollateralAdapter *CtfCollateralAdapterFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *CtfCollateralAdapterOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _CtfCollateralAdapter.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CtfCollateralAdapterOwnershipHandoverCanceled)
				if err := _CtfCollateralAdapter.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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

// ParseOwnershipHandoverCanceled is a log parse operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_CtfCollateralAdapter *CtfCollateralAdapterFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*CtfCollateralAdapterOwnershipHandoverCanceled, error) {
	event := new(CtfCollateralAdapterOwnershipHandoverCanceled)
	if err := _CtfCollateralAdapter.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CtfCollateralAdapterOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the CtfCollateralAdapter contract.
type CtfCollateralAdapterOwnershipHandoverRequestedIterator struct {
	Event *CtfCollateralAdapterOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *CtfCollateralAdapterOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CtfCollateralAdapterOwnershipHandoverRequested)
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
		it.Event = new(CtfCollateralAdapterOwnershipHandoverRequested)
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
func (it *CtfCollateralAdapterOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CtfCollateralAdapterOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CtfCollateralAdapterOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the CtfCollateralAdapter contract.
type CtfCollateralAdapterOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_CtfCollateralAdapter *CtfCollateralAdapterFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*CtfCollateralAdapterOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _CtfCollateralAdapter.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CtfCollateralAdapterOwnershipHandoverRequestedIterator{contract: _CtfCollateralAdapter.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_CtfCollateralAdapter *CtfCollateralAdapterFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *CtfCollateralAdapterOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _CtfCollateralAdapter.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CtfCollateralAdapterOwnershipHandoverRequested)
				if err := _CtfCollateralAdapter.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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

// ParseOwnershipHandoverRequested is a log parse operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_CtfCollateralAdapter *CtfCollateralAdapterFilterer) ParseOwnershipHandoverRequested(log types.Log) (*CtfCollateralAdapterOwnershipHandoverRequested, error) {
	event := new(CtfCollateralAdapterOwnershipHandoverRequested)
	if err := _CtfCollateralAdapter.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CtfCollateralAdapterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CtfCollateralAdapter contract.
type CtfCollateralAdapterOwnershipTransferredIterator struct {
	Event *CtfCollateralAdapterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CtfCollateralAdapterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CtfCollateralAdapterOwnershipTransferred)
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
		it.Event = new(CtfCollateralAdapterOwnershipTransferred)
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
func (it *CtfCollateralAdapterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CtfCollateralAdapterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CtfCollateralAdapterOwnershipTransferred represents a OwnershipTransferred event raised by the CtfCollateralAdapter contract.
type CtfCollateralAdapterOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_CtfCollateralAdapter *CtfCollateralAdapterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*CtfCollateralAdapterOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CtfCollateralAdapter.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CtfCollateralAdapterOwnershipTransferredIterator{contract: _CtfCollateralAdapter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_CtfCollateralAdapter *CtfCollateralAdapterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CtfCollateralAdapterOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CtfCollateralAdapter.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CtfCollateralAdapterOwnershipTransferred)
				if err := _CtfCollateralAdapter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_CtfCollateralAdapter *CtfCollateralAdapterFilterer) ParseOwnershipTransferred(log types.Log) (*CtfCollateralAdapterOwnershipTransferred, error) {
	event := new(CtfCollateralAdapterOwnershipTransferred)
	if err := _CtfCollateralAdapter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CtfCollateralAdapterPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CtfCollateralAdapter contract.
type CtfCollateralAdapterPausedIterator struct {
	Event *CtfCollateralAdapterPaused // Event containing the contract specifics and raw log

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
func (it *CtfCollateralAdapterPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CtfCollateralAdapterPaused)
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
		it.Event = new(CtfCollateralAdapterPaused)
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
func (it *CtfCollateralAdapterPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CtfCollateralAdapterPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CtfCollateralAdapterPaused represents a Paused event raised by the CtfCollateralAdapter contract.
type CtfCollateralAdapterPaused struct {
	Asset common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address indexed asset)
func (_CtfCollateralAdapter *CtfCollateralAdapterFilterer) FilterPaused(opts *bind.FilterOpts, asset []common.Address) (*CtfCollateralAdapterPausedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _CtfCollateralAdapter.contract.FilterLogs(opts, "Paused", assetRule)
	if err != nil {
		return nil, err
	}
	return &CtfCollateralAdapterPausedIterator{contract: _CtfCollateralAdapter.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address indexed asset)
func (_CtfCollateralAdapter *CtfCollateralAdapterFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CtfCollateralAdapterPaused, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _CtfCollateralAdapter.contract.WatchLogs(opts, "Paused", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CtfCollateralAdapterPaused)
				if err := _CtfCollateralAdapter.contract.UnpackLog(event, "Paused", log); err != nil {
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
// Solidity: event Paused(address indexed asset)
func (_CtfCollateralAdapter *CtfCollateralAdapterFilterer) ParsePaused(log types.Log) (*CtfCollateralAdapterPaused, error) {
	event := new(CtfCollateralAdapterPaused)
	if err := _CtfCollateralAdapter.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CtfCollateralAdapterRolesUpdatedIterator is returned from FilterRolesUpdated and is used to iterate over the raw logs and unpacked data for RolesUpdated events raised by the CtfCollateralAdapter contract.
type CtfCollateralAdapterRolesUpdatedIterator struct {
	Event *CtfCollateralAdapterRolesUpdated // Event containing the contract specifics and raw log

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
func (it *CtfCollateralAdapterRolesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CtfCollateralAdapterRolesUpdated)
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
		it.Event = new(CtfCollateralAdapterRolesUpdated)
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
func (it *CtfCollateralAdapterRolesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CtfCollateralAdapterRolesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CtfCollateralAdapterRolesUpdated represents a RolesUpdated event raised by the CtfCollateralAdapter contract.
type CtfCollateralAdapterRolesUpdated struct {
	User  common.Address
	Roles *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRolesUpdated is a free log retrieval operation binding the contract event 0x715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26.
//
// Solidity: event RolesUpdated(address indexed user, uint256 indexed roles)
func (_CtfCollateralAdapter *CtfCollateralAdapterFilterer) FilterRolesUpdated(opts *bind.FilterOpts, user []common.Address, roles []*big.Int) (*CtfCollateralAdapterRolesUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rolesRule []interface{}
	for _, rolesItem := range roles {
		rolesRule = append(rolesRule, rolesItem)
	}

	logs, sub, err := _CtfCollateralAdapter.contract.FilterLogs(opts, "RolesUpdated", userRule, rolesRule)
	if err != nil {
		return nil, err
	}
	return &CtfCollateralAdapterRolesUpdatedIterator{contract: _CtfCollateralAdapter.contract, event: "RolesUpdated", logs: logs, sub: sub}, nil
}

// WatchRolesUpdated is a free log subscription operation binding the contract event 0x715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26.
//
// Solidity: event RolesUpdated(address indexed user, uint256 indexed roles)
func (_CtfCollateralAdapter *CtfCollateralAdapterFilterer) WatchRolesUpdated(opts *bind.WatchOpts, sink chan<- *CtfCollateralAdapterRolesUpdated, user []common.Address, roles []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rolesRule []interface{}
	for _, rolesItem := range roles {
		rolesRule = append(rolesRule, rolesItem)
	}

	logs, sub, err := _CtfCollateralAdapter.contract.WatchLogs(opts, "RolesUpdated", userRule, rolesRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CtfCollateralAdapterRolesUpdated)
				if err := _CtfCollateralAdapter.contract.UnpackLog(event, "RolesUpdated", log); err != nil {
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

// ParseRolesUpdated is a log parse operation binding the contract event 0x715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26.
//
// Solidity: event RolesUpdated(address indexed user, uint256 indexed roles)
func (_CtfCollateralAdapter *CtfCollateralAdapterFilterer) ParseRolesUpdated(log types.Log) (*CtfCollateralAdapterRolesUpdated, error) {
	event := new(CtfCollateralAdapterRolesUpdated)
	if err := _CtfCollateralAdapter.contract.UnpackLog(event, "RolesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CtfCollateralAdapterUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CtfCollateralAdapter contract.
type CtfCollateralAdapterUnpausedIterator struct {
	Event *CtfCollateralAdapterUnpaused // Event containing the contract specifics and raw log

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
func (it *CtfCollateralAdapterUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CtfCollateralAdapterUnpaused)
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
		it.Event = new(CtfCollateralAdapterUnpaused)
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
func (it *CtfCollateralAdapterUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CtfCollateralAdapterUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CtfCollateralAdapterUnpaused represents a Unpaused event raised by the CtfCollateralAdapter contract.
type CtfCollateralAdapterUnpaused struct {
	Asset common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address indexed asset)
func (_CtfCollateralAdapter *CtfCollateralAdapterFilterer) FilterUnpaused(opts *bind.FilterOpts, asset []common.Address) (*CtfCollateralAdapterUnpausedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _CtfCollateralAdapter.contract.FilterLogs(opts, "Unpaused", assetRule)
	if err != nil {
		return nil, err
	}
	return &CtfCollateralAdapterUnpausedIterator{contract: _CtfCollateralAdapter.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address indexed asset)
func (_CtfCollateralAdapter *CtfCollateralAdapterFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CtfCollateralAdapterUnpaused, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _CtfCollateralAdapter.contract.WatchLogs(opts, "Unpaused", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CtfCollateralAdapterUnpaused)
				if err := _CtfCollateralAdapter.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
// Solidity: event Unpaused(address indexed asset)
func (_CtfCollateralAdapter *CtfCollateralAdapterFilterer) ParseUnpaused(log types.Log) (*CtfCollateralAdapterUnpaused, error) {
	event := new(CtfCollateralAdapterUnpaused)
	if err := _CtfCollateralAdapter.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
