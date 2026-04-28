// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package collateral_token

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

// CollateralTokenMetaData contains all meta data concerning the CollateralToken contract.
var CollateralTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdce\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AllowanceOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllowanceUnderflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpiredDeadline\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAsset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPermit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoHandoverRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyUnpaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Permit2AllowanceIsFixedAtInfinity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermitExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TotalSupplyOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpgradeFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"RolesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unwrapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Wrapped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDCE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VAULT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wrapper\",\"type\":\"address\"}],\"name\":\"addWrapper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"completeOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"grantRoles\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"hasAllRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"hasAnyRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"result\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"ownershipHandoverExpiresAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"removeMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wrapper\",\"type\":\"address\"}],\"name\":\"removeWrapper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"renounceRoles\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"revokeRoles\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"rolesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_callbackReceiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"unwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_callbackReceiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"wrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CollateralTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use CollateralTokenMetaData.ABI instead.
var CollateralTokenABI = CollateralTokenMetaData.ABI

// CollateralToken is an auto generated Go binding around an Ethereum contract.
type CollateralToken struct {
	CollateralTokenCaller     // Read-only binding to the contract
	CollateralTokenTransactor // Write-only binding to the contract
	CollateralTokenFilterer   // Log filterer for contract events
}

// CollateralTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type CollateralTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollateralTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CollateralTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollateralTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CollateralTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollateralTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CollateralTokenSession struct {
	Contract     *CollateralToken  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CollateralTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CollateralTokenCallerSession struct {
	Contract *CollateralTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CollateralTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CollateralTokenTransactorSession struct {
	Contract     *CollateralTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CollateralTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type CollateralTokenRaw struct {
	Contract *CollateralToken // Generic contract binding to access the raw methods on
}

// CollateralTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CollateralTokenCallerRaw struct {
	Contract *CollateralTokenCaller // Generic read-only contract binding to access the raw methods on
}

// CollateralTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CollateralTokenTransactorRaw struct {
	Contract *CollateralTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCollateralToken creates a new instance of CollateralToken, bound to a specific deployed contract.
func NewCollateralToken(address common.Address, backend bind.ContractBackend) (*CollateralToken, error) {
	contract, err := bindCollateralToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CollateralToken{CollateralTokenCaller: CollateralTokenCaller{contract: contract}, CollateralTokenTransactor: CollateralTokenTransactor{contract: contract}, CollateralTokenFilterer: CollateralTokenFilterer{contract: contract}}, nil
}

// NewCollateralTokenCaller creates a new read-only instance of CollateralToken, bound to a specific deployed contract.
func NewCollateralTokenCaller(address common.Address, caller bind.ContractCaller) (*CollateralTokenCaller, error) {
	contract, err := bindCollateralToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CollateralTokenCaller{contract: contract}, nil
}

// NewCollateralTokenTransactor creates a new write-only instance of CollateralToken, bound to a specific deployed contract.
func NewCollateralTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CollateralTokenTransactor, error) {
	contract, err := bindCollateralToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CollateralTokenTransactor{contract: contract}, nil
}

// NewCollateralTokenFilterer creates a new log filterer instance of CollateralToken, bound to a specific deployed contract.
func NewCollateralTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*CollateralTokenFilterer, error) {
	contract, err := bindCollateralToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CollateralTokenFilterer{contract: contract}, nil
}

// bindCollateralToken binds a generic wrapper to an already deployed contract.
func bindCollateralToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CollateralTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CollateralToken *CollateralTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CollateralToken.Contract.CollateralTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CollateralToken *CollateralTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollateralToken.Contract.CollateralTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CollateralToken *CollateralTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CollateralToken.Contract.CollateralTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CollateralToken *CollateralTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CollateralToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CollateralToken *CollateralTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollateralToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CollateralToken *CollateralTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CollateralToken.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 result)
func (_CollateralToken *CollateralTokenCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CollateralToken.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 result)
func (_CollateralToken *CollateralTokenSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _CollateralToken.Contract.DOMAINSEPARATOR(&_CollateralToken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 result)
func (_CollateralToken *CollateralTokenCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _CollateralToken.Contract.DOMAINSEPARATOR(&_CollateralToken.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_CollateralToken *CollateralTokenCaller) USDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollateralToken.contract.Call(opts, &out, "USDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_CollateralToken *CollateralTokenSession) USDC() (common.Address, error) {
	return _CollateralToken.Contract.USDC(&_CollateralToken.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_CollateralToken *CollateralTokenCallerSession) USDC() (common.Address, error) {
	return _CollateralToken.Contract.USDC(&_CollateralToken.CallOpts)
}

// USDCE is a free data retrieval call binding the contract method 0x195187e1.
//
// Solidity: function USDCE() view returns(address)
func (_CollateralToken *CollateralTokenCaller) USDCE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollateralToken.contract.Call(opts, &out, "USDCE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDCE is a free data retrieval call binding the contract method 0x195187e1.
//
// Solidity: function USDCE() view returns(address)
func (_CollateralToken *CollateralTokenSession) USDCE() (common.Address, error) {
	return _CollateralToken.Contract.USDCE(&_CollateralToken.CallOpts)
}

// USDCE is a free data retrieval call binding the contract method 0x195187e1.
//
// Solidity: function USDCE() view returns(address)
func (_CollateralToken *CollateralTokenCallerSession) USDCE() (common.Address, error) {
	return _CollateralToken.Contract.USDCE(&_CollateralToken.CallOpts)
}

// VAULT is a free data retrieval call binding the contract method 0x411557d1.
//
// Solidity: function VAULT() view returns(address)
func (_CollateralToken *CollateralTokenCaller) VAULT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollateralToken.contract.Call(opts, &out, "VAULT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VAULT is a free data retrieval call binding the contract method 0x411557d1.
//
// Solidity: function VAULT() view returns(address)
func (_CollateralToken *CollateralTokenSession) VAULT() (common.Address, error) {
	return _CollateralToken.Contract.VAULT(&_CollateralToken.CallOpts)
}

// VAULT is a free data retrieval call binding the contract method 0x411557d1.
//
// Solidity: function VAULT() view returns(address)
func (_CollateralToken *CollateralTokenCallerSession) VAULT() (common.Address, error) {
	return _CollateralToken.Contract.VAULT(&_CollateralToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 result)
func (_CollateralToken *CollateralTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CollateralToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 result)
func (_CollateralToken *CollateralTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CollateralToken.Contract.Allowance(&_CollateralToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 result)
func (_CollateralToken *CollateralTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CollateralToken.Contract.Allowance(&_CollateralToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 result)
func (_CollateralToken *CollateralTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CollateralToken.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 result)
func (_CollateralToken *CollateralTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CollateralToken.Contract.BalanceOf(&_CollateralToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 result)
func (_CollateralToken *CollateralTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CollateralToken.Contract.BalanceOf(&_CollateralToken.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_CollateralToken *CollateralTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CollateralToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_CollateralToken *CollateralTokenSession) Decimals() (uint8, error) {
	return _CollateralToken.Contract.Decimals(&_CollateralToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_CollateralToken *CollateralTokenCallerSession) Decimals() (uint8, error) {
	return _CollateralToken.Contract.Decimals(&_CollateralToken.CallOpts)
}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_CollateralToken *CollateralTokenCaller) HasAllRoles(opts *bind.CallOpts, user common.Address, roles *big.Int) (bool, error) {
	var out []interface{}
	err := _CollateralToken.contract.Call(opts, &out, "hasAllRoles", user, roles)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_CollateralToken *CollateralTokenSession) HasAllRoles(user common.Address, roles *big.Int) (bool, error) {
	return _CollateralToken.Contract.HasAllRoles(&_CollateralToken.CallOpts, user, roles)
}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_CollateralToken *CollateralTokenCallerSession) HasAllRoles(user common.Address, roles *big.Int) (bool, error) {
	return _CollateralToken.Contract.HasAllRoles(&_CollateralToken.CallOpts, user, roles)
}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_CollateralToken *CollateralTokenCaller) HasAnyRole(opts *bind.CallOpts, user common.Address, roles *big.Int) (bool, error) {
	var out []interface{}
	err := _CollateralToken.contract.Call(opts, &out, "hasAnyRole", user, roles)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_CollateralToken *CollateralTokenSession) HasAnyRole(user common.Address, roles *big.Int) (bool, error) {
	return _CollateralToken.Contract.HasAnyRole(&_CollateralToken.CallOpts, user, roles)
}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_CollateralToken *CollateralTokenCallerSession) HasAnyRole(user common.Address, roles *big.Int) (bool, error) {
	return _CollateralToken.Contract.HasAnyRole(&_CollateralToken.CallOpts, user, roles)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_CollateralToken *CollateralTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CollateralToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_CollateralToken *CollateralTokenSession) Name() (string, error) {
	return _CollateralToken.Contract.Name(&_CollateralToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_CollateralToken *CollateralTokenCallerSession) Name() (string, error) {
	return _CollateralToken.Contract.Name(&_CollateralToken.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256 result)
func (_CollateralToken *CollateralTokenCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CollateralToken.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256 result)
func (_CollateralToken *CollateralTokenSession) Nonces(owner common.Address) (*big.Int, error) {
	return _CollateralToken.Contract.Nonces(&_CollateralToken.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256 result)
func (_CollateralToken *CollateralTokenCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _CollateralToken.Contract.Nonces(&_CollateralToken.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_CollateralToken *CollateralTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollateralToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_CollateralToken *CollateralTokenSession) Owner() (common.Address, error) {
	return _CollateralToken.Contract.Owner(&_CollateralToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_CollateralToken *CollateralTokenCallerSession) Owner() (common.Address, error) {
	return _CollateralToken.Contract.Owner(&_CollateralToken.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_CollateralToken *CollateralTokenCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CollateralToken.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_CollateralToken *CollateralTokenSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _CollateralToken.Contract.OwnershipHandoverExpiresAt(&_CollateralToken.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_CollateralToken *CollateralTokenCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _CollateralToken.Contract.OwnershipHandoverExpiresAt(&_CollateralToken.CallOpts, pendingOwner)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CollateralToken *CollateralTokenCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CollateralToken.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CollateralToken *CollateralTokenSession) ProxiableUUID() ([32]byte, error) {
	return _CollateralToken.Contract.ProxiableUUID(&_CollateralToken.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CollateralToken *CollateralTokenCallerSession) ProxiableUUID() ([32]byte, error) {
	return _CollateralToken.Contract.ProxiableUUID(&_CollateralToken.CallOpts)
}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_CollateralToken *CollateralTokenCaller) RolesOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CollateralToken.contract.Call(opts, &out, "rolesOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_CollateralToken *CollateralTokenSession) RolesOf(user common.Address) (*big.Int, error) {
	return _CollateralToken.Contract.RolesOf(&_CollateralToken.CallOpts, user)
}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_CollateralToken *CollateralTokenCallerSession) RolesOf(user common.Address) (*big.Int, error) {
	return _CollateralToken.Contract.RolesOf(&_CollateralToken.CallOpts, user)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_CollateralToken *CollateralTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CollateralToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_CollateralToken *CollateralTokenSession) Symbol() (string, error) {
	return _CollateralToken.Contract.Symbol(&_CollateralToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_CollateralToken *CollateralTokenCallerSession) Symbol() (string, error) {
	return _CollateralToken.Contract.Symbol(&_CollateralToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_CollateralToken *CollateralTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CollateralToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_CollateralToken *CollateralTokenSession) TotalSupply() (*big.Int, error) {
	return _CollateralToken.Contract.TotalSupply(&_CollateralToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_CollateralToken *CollateralTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _CollateralToken.Contract.TotalSupply(&_CollateralToken.CallOpts)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address _minter) returns()
func (_CollateralToken *CollateralTokenTransactor) AddMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _CollateralToken.contract.Transact(opts, "addMinter", _minter)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address _minter) returns()
func (_CollateralToken *CollateralTokenSession) AddMinter(_minter common.Address) (*types.Transaction, error) {
	return _CollateralToken.Contract.AddMinter(&_CollateralToken.TransactOpts, _minter)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address _minter) returns()
func (_CollateralToken *CollateralTokenTransactorSession) AddMinter(_minter common.Address) (*types.Transaction, error) {
	return _CollateralToken.Contract.AddMinter(&_CollateralToken.TransactOpts, _minter)
}

// AddWrapper is a paid mutator transaction binding the contract method 0x26b09c29.
//
// Solidity: function addWrapper(address _wrapper) returns()
func (_CollateralToken *CollateralTokenTransactor) AddWrapper(opts *bind.TransactOpts, _wrapper common.Address) (*types.Transaction, error) {
	return _CollateralToken.contract.Transact(opts, "addWrapper", _wrapper)
}

// AddWrapper is a paid mutator transaction binding the contract method 0x26b09c29.
//
// Solidity: function addWrapper(address _wrapper) returns()
func (_CollateralToken *CollateralTokenSession) AddWrapper(_wrapper common.Address) (*types.Transaction, error) {
	return _CollateralToken.Contract.AddWrapper(&_CollateralToken.TransactOpts, _wrapper)
}

// AddWrapper is a paid mutator transaction binding the contract method 0x26b09c29.
//
// Solidity: function addWrapper(address _wrapper) returns()
func (_CollateralToken *CollateralTokenTransactorSession) AddWrapper(_wrapper common.Address) (*types.Transaction, error) {
	return _CollateralToken.Contract.AddWrapper(&_CollateralToken.TransactOpts, _wrapper)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CollateralToken *CollateralTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CollateralToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CollateralToken *CollateralTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CollateralToken.Contract.Approve(&_CollateralToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CollateralToken *CollateralTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CollateralToken.Contract.Approve(&_CollateralToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns()
func (_CollateralToken *CollateralTokenTransactor) Burn(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _CollateralToken.contract.Transact(opts, "burn", _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns()
func (_CollateralToken *CollateralTokenSession) Burn(_amount *big.Int) (*types.Transaction, error) {
	return _CollateralToken.Contract.Burn(&_CollateralToken.TransactOpts, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns()
func (_CollateralToken *CollateralTokenTransactorSession) Burn(_amount *big.Int) (*types.Transaction, error) {
	return _CollateralToken.Contract.Burn(&_CollateralToken.TransactOpts, _amount)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_CollateralToken *CollateralTokenTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollateralToken.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_CollateralToken *CollateralTokenSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _CollateralToken.Contract.CancelOwnershipHandover(&_CollateralToken.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_CollateralToken *CollateralTokenTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _CollateralToken.Contract.CancelOwnershipHandover(&_CollateralToken.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_CollateralToken *CollateralTokenTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _CollateralToken.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_CollateralToken *CollateralTokenSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _CollateralToken.Contract.CompleteOwnershipHandover(&_CollateralToken.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_CollateralToken *CollateralTokenTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _CollateralToken.Contract.CompleteOwnershipHandover(&_CollateralToken.TransactOpts, pendingOwner)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_CollateralToken *CollateralTokenTransactor) GrantRoles(opts *bind.TransactOpts, user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _CollateralToken.contract.Transact(opts, "grantRoles", user, roles)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_CollateralToken *CollateralTokenSession) GrantRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _CollateralToken.Contract.GrantRoles(&_CollateralToken.TransactOpts, user, roles)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_CollateralToken *CollateralTokenTransactorSession) GrantRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _CollateralToken.Contract.GrantRoles(&_CollateralToken.TransactOpts, user, roles)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_CollateralToken *CollateralTokenTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _CollateralToken.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_CollateralToken *CollateralTokenSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _CollateralToken.Contract.Initialize(&_CollateralToken.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_CollateralToken *CollateralTokenTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _CollateralToken.Contract.Initialize(&_CollateralToken.TransactOpts, _owner)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_CollateralToken *CollateralTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CollateralToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_CollateralToken *CollateralTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CollateralToken.Contract.Mint(&_CollateralToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_CollateralToken *CollateralTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CollateralToken.Contract.Mint(&_CollateralToken.TransactOpts, _to, _amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CollateralToken *CollateralTokenTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CollateralToken.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CollateralToken *CollateralTokenSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CollateralToken.Contract.Permit(&_CollateralToken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CollateralToken *CollateralTokenTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CollateralToken.Contract.Permit(&_CollateralToken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address _minter) returns()
func (_CollateralToken *CollateralTokenTransactor) RemoveMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _CollateralToken.contract.Transact(opts, "removeMinter", _minter)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address _minter) returns()
func (_CollateralToken *CollateralTokenSession) RemoveMinter(_minter common.Address) (*types.Transaction, error) {
	return _CollateralToken.Contract.RemoveMinter(&_CollateralToken.TransactOpts, _minter)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address _minter) returns()
func (_CollateralToken *CollateralTokenTransactorSession) RemoveMinter(_minter common.Address) (*types.Transaction, error) {
	return _CollateralToken.Contract.RemoveMinter(&_CollateralToken.TransactOpts, _minter)
}

// RemoveWrapper is a paid mutator transaction binding the contract method 0xe914d494.
//
// Solidity: function removeWrapper(address _wrapper) returns()
func (_CollateralToken *CollateralTokenTransactor) RemoveWrapper(opts *bind.TransactOpts, _wrapper common.Address) (*types.Transaction, error) {
	return _CollateralToken.contract.Transact(opts, "removeWrapper", _wrapper)
}

// RemoveWrapper is a paid mutator transaction binding the contract method 0xe914d494.
//
// Solidity: function removeWrapper(address _wrapper) returns()
func (_CollateralToken *CollateralTokenSession) RemoveWrapper(_wrapper common.Address) (*types.Transaction, error) {
	return _CollateralToken.Contract.RemoveWrapper(&_CollateralToken.TransactOpts, _wrapper)
}

// RemoveWrapper is a paid mutator transaction binding the contract method 0xe914d494.
//
// Solidity: function removeWrapper(address _wrapper) returns()
func (_CollateralToken *CollateralTokenTransactorSession) RemoveWrapper(_wrapper common.Address) (*types.Transaction, error) {
	return _CollateralToken.Contract.RemoveWrapper(&_CollateralToken.TransactOpts, _wrapper)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_CollateralToken *CollateralTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollateralToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_CollateralToken *CollateralTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _CollateralToken.Contract.RenounceOwnership(&_CollateralToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_CollateralToken *CollateralTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CollateralToken.Contract.RenounceOwnership(&_CollateralToken.TransactOpts)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_CollateralToken *CollateralTokenTransactor) RenounceRoles(opts *bind.TransactOpts, roles *big.Int) (*types.Transaction, error) {
	return _CollateralToken.contract.Transact(opts, "renounceRoles", roles)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_CollateralToken *CollateralTokenSession) RenounceRoles(roles *big.Int) (*types.Transaction, error) {
	return _CollateralToken.Contract.RenounceRoles(&_CollateralToken.TransactOpts, roles)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_CollateralToken *CollateralTokenTransactorSession) RenounceRoles(roles *big.Int) (*types.Transaction, error) {
	return _CollateralToken.Contract.RenounceRoles(&_CollateralToken.TransactOpts, roles)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_CollateralToken *CollateralTokenTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollateralToken.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_CollateralToken *CollateralTokenSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _CollateralToken.Contract.RequestOwnershipHandover(&_CollateralToken.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_CollateralToken *CollateralTokenTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _CollateralToken.Contract.RequestOwnershipHandover(&_CollateralToken.TransactOpts)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_CollateralToken *CollateralTokenTransactor) RevokeRoles(opts *bind.TransactOpts, user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _CollateralToken.contract.Transact(opts, "revokeRoles", user, roles)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_CollateralToken *CollateralTokenSession) RevokeRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _CollateralToken.Contract.RevokeRoles(&_CollateralToken.TransactOpts, user, roles)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_CollateralToken *CollateralTokenTransactorSession) RevokeRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _CollateralToken.Contract.RevokeRoles(&_CollateralToken.TransactOpts, user, roles)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_CollateralToken *CollateralTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CollateralToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_CollateralToken *CollateralTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CollateralToken.Contract.Transfer(&_CollateralToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_CollateralToken *CollateralTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CollateralToken.Contract.Transfer(&_CollateralToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_CollateralToken *CollateralTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CollateralToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_CollateralToken *CollateralTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CollateralToken.Contract.TransferFrom(&_CollateralToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_CollateralToken *CollateralTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CollateralToken.Contract.TransferFrom(&_CollateralToken.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_CollateralToken *CollateralTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CollateralToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_CollateralToken *CollateralTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CollateralToken.Contract.TransferOwnership(&_CollateralToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_CollateralToken *CollateralTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CollateralToken.Contract.TransferOwnership(&_CollateralToken.TransactOpts, newOwner)
}

// Unwrap is a paid mutator transaction binding the contract method 0xd600875d.
//
// Solidity: function unwrap(address _asset, address _to, uint256 _amount, address _callbackReceiver, bytes _data) returns()
func (_CollateralToken *CollateralTokenTransactor) Unwrap(opts *bind.TransactOpts, _asset common.Address, _to common.Address, _amount *big.Int, _callbackReceiver common.Address, _data []byte) (*types.Transaction, error) {
	return _CollateralToken.contract.Transact(opts, "unwrap", _asset, _to, _amount, _callbackReceiver, _data)
}

// Unwrap is a paid mutator transaction binding the contract method 0xd600875d.
//
// Solidity: function unwrap(address _asset, address _to, uint256 _amount, address _callbackReceiver, bytes _data) returns()
func (_CollateralToken *CollateralTokenSession) Unwrap(_asset common.Address, _to common.Address, _amount *big.Int, _callbackReceiver common.Address, _data []byte) (*types.Transaction, error) {
	return _CollateralToken.Contract.Unwrap(&_CollateralToken.TransactOpts, _asset, _to, _amount, _callbackReceiver, _data)
}

// Unwrap is a paid mutator transaction binding the contract method 0xd600875d.
//
// Solidity: function unwrap(address _asset, address _to, uint256 _amount, address _callbackReceiver, bytes _data) returns()
func (_CollateralToken *CollateralTokenTransactorSession) Unwrap(_asset common.Address, _to common.Address, _amount *big.Int, _callbackReceiver common.Address, _data []byte) (*types.Transaction, error) {
	return _CollateralToken.Contract.Unwrap(&_CollateralToken.TransactOpts, _asset, _to, _amount, _callbackReceiver, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CollateralToken *CollateralTokenTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CollateralToken.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CollateralToken *CollateralTokenSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CollateralToken.Contract.UpgradeToAndCall(&_CollateralToken.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CollateralToken *CollateralTokenTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CollateralToken.Contract.UpgradeToAndCall(&_CollateralToken.TransactOpts, newImplementation, data)
}

// Wrap is a paid mutator transaction binding the contract method 0xb97b57c7.
//
// Solidity: function wrap(address _asset, address _to, uint256 _amount, address _callbackReceiver, bytes _data) returns()
func (_CollateralToken *CollateralTokenTransactor) Wrap(opts *bind.TransactOpts, _asset common.Address, _to common.Address, _amount *big.Int, _callbackReceiver common.Address, _data []byte) (*types.Transaction, error) {
	return _CollateralToken.contract.Transact(opts, "wrap", _asset, _to, _amount, _callbackReceiver, _data)
}

// Wrap is a paid mutator transaction binding the contract method 0xb97b57c7.
//
// Solidity: function wrap(address _asset, address _to, uint256 _amount, address _callbackReceiver, bytes _data) returns()
func (_CollateralToken *CollateralTokenSession) Wrap(_asset common.Address, _to common.Address, _amount *big.Int, _callbackReceiver common.Address, _data []byte) (*types.Transaction, error) {
	return _CollateralToken.Contract.Wrap(&_CollateralToken.TransactOpts, _asset, _to, _amount, _callbackReceiver, _data)
}

// Wrap is a paid mutator transaction binding the contract method 0xb97b57c7.
//
// Solidity: function wrap(address _asset, address _to, uint256 _amount, address _callbackReceiver, bytes _data) returns()
func (_CollateralToken *CollateralTokenTransactorSession) Wrap(_asset common.Address, _to common.Address, _amount *big.Int, _callbackReceiver common.Address, _data []byte) (*types.Transaction, error) {
	return _CollateralToken.Contract.Wrap(&_CollateralToken.TransactOpts, _asset, _to, _amount, _callbackReceiver, _data)
}

// CollateralTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CollateralToken contract.
type CollateralTokenApprovalIterator struct {
	Event *CollateralTokenApproval // Event containing the contract specifics and raw log

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
func (it *CollateralTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralTokenApproval)
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
		it.Event = new(CollateralTokenApproval)
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
func (it *CollateralTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralTokenApproval represents a Approval event raised by the CollateralToken contract.
type CollateralTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_CollateralToken *CollateralTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CollateralTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CollateralToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CollateralTokenApprovalIterator{contract: _CollateralToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_CollateralToken *CollateralTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CollateralTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CollateralToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralTokenApproval)
				if err := _CollateralToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_CollateralToken *CollateralTokenFilterer) ParseApproval(log types.Log) (*CollateralTokenApproval, error) {
	event := new(CollateralTokenApproval)
	if err := _CollateralToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralTokenInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CollateralToken contract.
type CollateralTokenInitializedIterator struct {
	Event *CollateralTokenInitialized // Event containing the contract specifics and raw log

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
func (it *CollateralTokenInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralTokenInitialized)
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
		it.Event = new(CollateralTokenInitialized)
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
func (it *CollateralTokenInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralTokenInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralTokenInitialized represents a Initialized event raised by the CollateralToken contract.
type CollateralTokenInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CollateralToken *CollateralTokenFilterer) FilterInitialized(opts *bind.FilterOpts) (*CollateralTokenInitializedIterator, error) {

	logs, sub, err := _CollateralToken.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CollateralTokenInitializedIterator{contract: _CollateralToken.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CollateralToken *CollateralTokenFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CollateralTokenInitialized) (event.Subscription, error) {

	logs, sub, err := _CollateralToken.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralTokenInitialized)
				if err := _CollateralToken.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CollateralToken *CollateralTokenFilterer) ParseInitialized(log types.Log) (*CollateralTokenInitialized, error) {
	event := new(CollateralTokenInitialized)
	if err := _CollateralToken.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralTokenOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the CollateralToken contract.
type CollateralTokenOwnershipHandoverCanceledIterator struct {
	Event *CollateralTokenOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *CollateralTokenOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralTokenOwnershipHandoverCanceled)
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
		it.Event = new(CollateralTokenOwnershipHandoverCanceled)
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
func (it *CollateralTokenOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralTokenOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralTokenOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the CollateralToken contract.
type CollateralTokenOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_CollateralToken *CollateralTokenFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*CollateralTokenOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _CollateralToken.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CollateralTokenOwnershipHandoverCanceledIterator{contract: _CollateralToken.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_CollateralToken *CollateralTokenFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *CollateralTokenOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _CollateralToken.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralTokenOwnershipHandoverCanceled)
				if err := _CollateralToken.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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
func (_CollateralToken *CollateralTokenFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*CollateralTokenOwnershipHandoverCanceled, error) {
	event := new(CollateralTokenOwnershipHandoverCanceled)
	if err := _CollateralToken.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralTokenOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the CollateralToken contract.
type CollateralTokenOwnershipHandoverRequestedIterator struct {
	Event *CollateralTokenOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *CollateralTokenOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralTokenOwnershipHandoverRequested)
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
		it.Event = new(CollateralTokenOwnershipHandoverRequested)
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
func (it *CollateralTokenOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralTokenOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralTokenOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the CollateralToken contract.
type CollateralTokenOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_CollateralToken *CollateralTokenFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*CollateralTokenOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _CollateralToken.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CollateralTokenOwnershipHandoverRequestedIterator{contract: _CollateralToken.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_CollateralToken *CollateralTokenFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *CollateralTokenOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _CollateralToken.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralTokenOwnershipHandoverRequested)
				if err := _CollateralToken.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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
func (_CollateralToken *CollateralTokenFilterer) ParseOwnershipHandoverRequested(log types.Log) (*CollateralTokenOwnershipHandoverRequested, error) {
	event := new(CollateralTokenOwnershipHandoverRequested)
	if err := _CollateralToken.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CollateralToken contract.
type CollateralTokenOwnershipTransferredIterator struct {
	Event *CollateralTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CollateralTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralTokenOwnershipTransferred)
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
		it.Event = new(CollateralTokenOwnershipTransferred)
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
func (it *CollateralTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralTokenOwnershipTransferred represents a OwnershipTransferred event raised by the CollateralToken contract.
type CollateralTokenOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_CollateralToken *CollateralTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*CollateralTokenOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CollateralToken.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CollateralTokenOwnershipTransferredIterator{contract: _CollateralToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_CollateralToken *CollateralTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CollateralTokenOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CollateralToken.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralTokenOwnershipTransferred)
				if err := _CollateralToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CollateralToken *CollateralTokenFilterer) ParseOwnershipTransferred(log types.Log) (*CollateralTokenOwnershipTransferred, error) {
	event := new(CollateralTokenOwnershipTransferred)
	if err := _CollateralToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralTokenRolesUpdatedIterator is returned from FilterRolesUpdated and is used to iterate over the raw logs and unpacked data for RolesUpdated events raised by the CollateralToken contract.
type CollateralTokenRolesUpdatedIterator struct {
	Event *CollateralTokenRolesUpdated // Event containing the contract specifics and raw log

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
func (it *CollateralTokenRolesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralTokenRolesUpdated)
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
		it.Event = new(CollateralTokenRolesUpdated)
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
func (it *CollateralTokenRolesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralTokenRolesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralTokenRolesUpdated represents a RolesUpdated event raised by the CollateralToken contract.
type CollateralTokenRolesUpdated struct {
	User  common.Address
	Roles *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRolesUpdated is a free log retrieval operation binding the contract event 0x715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26.
//
// Solidity: event RolesUpdated(address indexed user, uint256 indexed roles)
func (_CollateralToken *CollateralTokenFilterer) FilterRolesUpdated(opts *bind.FilterOpts, user []common.Address, roles []*big.Int) (*CollateralTokenRolesUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rolesRule []interface{}
	for _, rolesItem := range roles {
		rolesRule = append(rolesRule, rolesItem)
	}

	logs, sub, err := _CollateralToken.contract.FilterLogs(opts, "RolesUpdated", userRule, rolesRule)
	if err != nil {
		return nil, err
	}
	return &CollateralTokenRolesUpdatedIterator{contract: _CollateralToken.contract, event: "RolesUpdated", logs: logs, sub: sub}, nil
}

// WatchRolesUpdated is a free log subscription operation binding the contract event 0x715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26.
//
// Solidity: event RolesUpdated(address indexed user, uint256 indexed roles)
func (_CollateralToken *CollateralTokenFilterer) WatchRolesUpdated(opts *bind.WatchOpts, sink chan<- *CollateralTokenRolesUpdated, user []common.Address, roles []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rolesRule []interface{}
	for _, rolesItem := range roles {
		rolesRule = append(rolesRule, rolesItem)
	}

	logs, sub, err := _CollateralToken.contract.WatchLogs(opts, "RolesUpdated", userRule, rolesRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralTokenRolesUpdated)
				if err := _CollateralToken.contract.UnpackLog(event, "RolesUpdated", log); err != nil {
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
func (_CollateralToken *CollateralTokenFilterer) ParseRolesUpdated(log types.Log) (*CollateralTokenRolesUpdated, error) {
	event := new(CollateralTokenRolesUpdated)
	if err := _CollateralToken.contract.UnpackLog(event, "RolesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CollateralToken contract.
type CollateralTokenTransferIterator struct {
	Event *CollateralTokenTransfer // Event containing the contract specifics and raw log

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
func (it *CollateralTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralTokenTransfer)
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
		it.Event = new(CollateralTokenTransfer)
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
func (it *CollateralTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralTokenTransfer represents a Transfer event raised by the CollateralToken contract.
type CollateralTokenTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_CollateralToken *CollateralTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CollateralTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CollateralToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CollateralTokenTransferIterator{contract: _CollateralToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_CollateralToken *CollateralTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CollateralTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CollateralToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralTokenTransfer)
				if err := _CollateralToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_CollateralToken *CollateralTokenFilterer) ParseTransfer(log types.Log) (*CollateralTokenTransfer, error) {
	event := new(CollateralTokenTransfer)
	if err := _CollateralToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralTokenUnwrappedIterator is returned from FilterUnwrapped and is used to iterate over the raw logs and unpacked data for Unwrapped events raised by the CollateralToken contract.
type CollateralTokenUnwrappedIterator struct {
	Event *CollateralTokenUnwrapped // Event containing the contract specifics and raw log

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
func (it *CollateralTokenUnwrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralTokenUnwrapped)
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
		it.Event = new(CollateralTokenUnwrapped)
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
func (it *CollateralTokenUnwrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralTokenUnwrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralTokenUnwrapped represents a Unwrapped event raised by the CollateralToken contract.
type CollateralTokenUnwrapped struct {
	Caller common.Address
	Asset  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnwrapped is a free log retrieval operation binding the contract event 0x18b42b684d0b621cc609f4d888916e5ed9e934a476259ec1c11ec116f2b9aa7f.
//
// Solidity: event Unwrapped(address indexed caller, address indexed asset, address indexed to, uint256 amount)
func (_CollateralToken *CollateralTokenFilterer) FilterUnwrapped(opts *bind.FilterOpts, caller []common.Address, asset []common.Address, to []common.Address) (*CollateralTokenUnwrappedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CollateralToken.contract.FilterLogs(opts, "Unwrapped", callerRule, assetRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CollateralTokenUnwrappedIterator{contract: _CollateralToken.contract, event: "Unwrapped", logs: logs, sub: sub}, nil
}

// WatchUnwrapped is a free log subscription operation binding the contract event 0x18b42b684d0b621cc609f4d888916e5ed9e934a476259ec1c11ec116f2b9aa7f.
//
// Solidity: event Unwrapped(address indexed caller, address indexed asset, address indexed to, uint256 amount)
func (_CollateralToken *CollateralTokenFilterer) WatchUnwrapped(opts *bind.WatchOpts, sink chan<- *CollateralTokenUnwrapped, caller []common.Address, asset []common.Address, to []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CollateralToken.contract.WatchLogs(opts, "Unwrapped", callerRule, assetRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralTokenUnwrapped)
				if err := _CollateralToken.contract.UnpackLog(event, "Unwrapped", log); err != nil {
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

// ParseUnwrapped is a log parse operation binding the contract event 0x18b42b684d0b621cc609f4d888916e5ed9e934a476259ec1c11ec116f2b9aa7f.
//
// Solidity: event Unwrapped(address indexed caller, address indexed asset, address indexed to, uint256 amount)
func (_CollateralToken *CollateralTokenFilterer) ParseUnwrapped(log types.Log) (*CollateralTokenUnwrapped, error) {
	event := new(CollateralTokenUnwrapped)
	if err := _CollateralToken.contract.UnpackLog(event, "Unwrapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralTokenUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the CollateralToken contract.
type CollateralTokenUpgradedIterator struct {
	Event *CollateralTokenUpgraded // Event containing the contract specifics and raw log

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
func (it *CollateralTokenUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralTokenUpgraded)
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
		it.Event = new(CollateralTokenUpgraded)
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
func (it *CollateralTokenUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralTokenUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralTokenUpgraded represents a Upgraded event raised by the CollateralToken contract.
type CollateralTokenUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CollateralToken *CollateralTokenFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CollateralTokenUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CollateralToken.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CollateralTokenUpgradedIterator{contract: _CollateralToken.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CollateralToken *CollateralTokenFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CollateralTokenUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CollateralToken.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralTokenUpgraded)
				if err := _CollateralToken.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CollateralToken *CollateralTokenFilterer) ParseUpgraded(log types.Log) (*CollateralTokenUpgraded, error) {
	event := new(CollateralTokenUpgraded)
	if err := _CollateralToken.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralTokenWrappedIterator is returned from FilterWrapped and is used to iterate over the raw logs and unpacked data for Wrapped events raised by the CollateralToken contract.
type CollateralTokenWrappedIterator struct {
	Event *CollateralTokenWrapped // Event containing the contract specifics and raw log

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
func (it *CollateralTokenWrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralTokenWrapped)
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
		it.Event = new(CollateralTokenWrapped)
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
func (it *CollateralTokenWrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralTokenWrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralTokenWrapped represents a Wrapped event raised by the CollateralToken contract.
type CollateralTokenWrapped struct {
	Caller common.Address
	Asset  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWrapped is a free log retrieval operation binding the contract event 0xc00a5c84859ae82a7f5e6a2773283fb525335d5b3195f61174aa1ecc7e15dd84.
//
// Solidity: event Wrapped(address indexed caller, address indexed asset, address indexed to, uint256 amount)
func (_CollateralToken *CollateralTokenFilterer) FilterWrapped(opts *bind.FilterOpts, caller []common.Address, asset []common.Address, to []common.Address) (*CollateralTokenWrappedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CollateralToken.contract.FilterLogs(opts, "Wrapped", callerRule, assetRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CollateralTokenWrappedIterator{contract: _CollateralToken.contract, event: "Wrapped", logs: logs, sub: sub}, nil
}

// WatchWrapped is a free log subscription operation binding the contract event 0xc00a5c84859ae82a7f5e6a2773283fb525335d5b3195f61174aa1ecc7e15dd84.
//
// Solidity: event Wrapped(address indexed caller, address indexed asset, address indexed to, uint256 amount)
func (_CollateralToken *CollateralTokenFilterer) WatchWrapped(opts *bind.WatchOpts, sink chan<- *CollateralTokenWrapped, caller []common.Address, asset []common.Address, to []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CollateralToken.contract.WatchLogs(opts, "Wrapped", callerRule, assetRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralTokenWrapped)
				if err := _CollateralToken.contract.UnpackLog(event, "Wrapped", log); err != nil {
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

// ParseWrapped is a log parse operation binding the contract event 0xc00a5c84859ae82a7f5e6a2773283fb525335d5b3195f61174aa1ecc7e15dd84.
//
// Solidity: event Wrapped(address indexed caller, address indexed asset, address indexed to, uint256 amount)
func (_CollateralToken *CollateralTokenFilterer) ParseWrapped(log types.Log) (*CollateralTokenWrapped, error) {
	event := new(CollateralTokenWrapped)
	if err := _CollateralToken.contract.UnpackLog(event, "Wrapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
