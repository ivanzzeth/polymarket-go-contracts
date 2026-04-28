// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package permissioned_ramp

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

// PermissionedRampMetaData contains all meta data concerning the PermissionedRamp contract.
var PermissionedRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpiredDeadline\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAsset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoHandoverRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyUnpaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"RolesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COLLATERAL_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_witness\",\"type\":\"address\"}],\"name\":\"addWitness\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"completeOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"grantRoles\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"hasAllRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"hasAnyRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"result\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"ownershipHandoverExpiresAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_witness\",\"type\":\"address\"}],\"name\":\"removeWitness\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"renounceRoles\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"revokeRoles\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"rolesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"unwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"wrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PermissionedRampABI is the input ABI used to generate the binding from.
// Deprecated: Use PermissionedRampMetaData.ABI instead.
var PermissionedRampABI = PermissionedRampMetaData.ABI

// PermissionedRamp is an auto generated Go binding around an Ethereum contract.
type PermissionedRamp struct {
	PermissionedRampCaller     // Read-only binding to the contract
	PermissionedRampTransactor // Write-only binding to the contract
	PermissionedRampFilterer   // Log filterer for contract events
}

// PermissionedRampCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermissionedRampCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionedRampTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermissionedRampTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionedRampFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermissionedRampFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionedRampSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermissionedRampSession struct {
	Contract     *PermissionedRamp // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PermissionedRampCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermissionedRampCallerSession struct {
	Contract *PermissionedRampCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PermissionedRampTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermissionedRampTransactorSession struct {
	Contract     *PermissionedRampTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PermissionedRampRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermissionedRampRaw struct {
	Contract *PermissionedRamp // Generic contract binding to access the raw methods on
}

// PermissionedRampCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermissionedRampCallerRaw struct {
	Contract *PermissionedRampCaller // Generic read-only contract binding to access the raw methods on
}

// PermissionedRampTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermissionedRampTransactorRaw struct {
	Contract *PermissionedRampTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermissionedRamp creates a new instance of PermissionedRamp, bound to a specific deployed contract.
func NewPermissionedRamp(address common.Address, backend bind.ContractBackend) (*PermissionedRamp, error) {
	contract, err := bindPermissionedRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PermissionedRamp{PermissionedRampCaller: PermissionedRampCaller{contract: contract}, PermissionedRampTransactor: PermissionedRampTransactor{contract: contract}, PermissionedRampFilterer: PermissionedRampFilterer{contract: contract}}, nil
}

// NewPermissionedRampCaller creates a new read-only instance of PermissionedRamp, bound to a specific deployed contract.
func NewPermissionedRampCaller(address common.Address, caller bind.ContractCaller) (*PermissionedRampCaller, error) {
	contract, err := bindPermissionedRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionedRampCaller{contract: contract}, nil
}

// NewPermissionedRampTransactor creates a new write-only instance of PermissionedRamp, bound to a specific deployed contract.
func NewPermissionedRampTransactor(address common.Address, transactor bind.ContractTransactor) (*PermissionedRampTransactor, error) {
	contract, err := bindPermissionedRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionedRampTransactor{contract: contract}, nil
}

// NewPermissionedRampFilterer creates a new log filterer instance of PermissionedRamp, bound to a specific deployed contract.
func NewPermissionedRampFilterer(address common.Address, filterer bind.ContractFilterer) (*PermissionedRampFilterer, error) {
	contract, err := bindPermissionedRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermissionedRampFilterer{contract: contract}, nil
}

// bindPermissionedRamp binds a generic wrapper to an already deployed contract.
func bindPermissionedRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PermissionedRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionedRamp *PermissionedRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermissionedRamp.Contract.PermissionedRampCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionedRamp *PermissionedRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.PermissionedRampTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionedRamp *PermissionedRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.PermissionedRampTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionedRamp *PermissionedRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermissionedRamp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionedRamp *PermissionedRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionedRamp *PermissionedRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.contract.Transact(opts, method, params...)
}

// COLLATERALTOKEN is a free data retrieval call binding the contract method 0xf5f1f1a7.
//
// Solidity: function COLLATERAL_TOKEN() view returns(address)
func (_PermissionedRamp *PermissionedRampCaller) COLLATERALTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PermissionedRamp.contract.Call(opts, &out, "COLLATERAL_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// COLLATERALTOKEN is a free data retrieval call binding the contract method 0xf5f1f1a7.
//
// Solidity: function COLLATERAL_TOKEN() view returns(address)
func (_PermissionedRamp *PermissionedRampSession) COLLATERALTOKEN() (common.Address, error) {
	return _PermissionedRamp.Contract.COLLATERALTOKEN(&_PermissionedRamp.CallOpts)
}

// COLLATERALTOKEN is a free data retrieval call binding the contract method 0xf5f1f1a7.
//
// Solidity: function COLLATERAL_TOKEN() view returns(address)
func (_PermissionedRamp *PermissionedRampCallerSession) COLLATERALTOKEN() (common.Address, error) {
	return _PermissionedRamp.Contract.COLLATERALTOKEN(&_PermissionedRamp.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_PermissionedRamp *PermissionedRampCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _PermissionedRamp.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_PermissionedRamp *PermissionedRampSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _PermissionedRamp.Contract.Eip712Domain(&_PermissionedRamp.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_PermissionedRamp *PermissionedRampCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _PermissionedRamp.Contract.Eip712Domain(&_PermissionedRamp.CallOpts)
}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_PermissionedRamp *PermissionedRampCaller) HasAllRoles(opts *bind.CallOpts, user common.Address, roles *big.Int) (bool, error) {
	var out []interface{}
	err := _PermissionedRamp.contract.Call(opts, &out, "hasAllRoles", user, roles)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_PermissionedRamp *PermissionedRampSession) HasAllRoles(user common.Address, roles *big.Int) (bool, error) {
	return _PermissionedRamp.Contract.HasAllRoles(&_PermissionedRamp.CallOpts, user, roles)
}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_PermissionedRamp *PermissionedRampCallerSession) HasAllRoles(user common.Address, roles *big.Int) (bool, error) {
	return _PermissionedRamp.Contract.HasAllRoles(&_PermissionedRamp.CallOpts, user, roles)
}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_PermissionedRamp *PermissionedRampCaller) HasAnyRole(opts *bind.CallOpts, user common.Address, roles *big.Int) (bool, error) {
	var out []interface{}
	err := _PermissionedRamp.contract.Call(opts, &out, "hasAnyRole", user, roles)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_PermissionedRamp *PermissionedRampSession) HasAnyRole(user common.Address, roles *big.Int) (bool, error) {
	return _PermissionedRamp.Contract.HasAnyRole(&_PermissionedRamp.CallOpts, user, roles)
}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_PermissionedRamp *PermissionedRampCallerSession) HasAnyRole(user common.Address, roles *big.Int) (bool, error) {
	return _PermissionedRamp.Contract.HasAnyRole(&_PermissionedRamp.CallOpts, user, roles)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_PermissionedRamp *PermissionedRampCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedRamp.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_PermissionedRamp *PermissionedRampSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _PermissionedRamp.Contract.Nonces(&_PermissionedRamp.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_PermissionedRamp *PermissionedRampCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _PermissionedRamp.Contract.Nonces(&_PermissionedRamp.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_PermissionedRamp *PermissionedRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PermissionedRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_PermissionedRamp *PermissionedRampSession) Owner() (common.Address, error) {
	return _PermissionedRamp.Contract.Owner(&_PermissionedRamp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_PermissionedRamp *PermissionedRampCallerSession) Owner() (common.Address, error) {
	return _PermissionedRamp.Contract.Owner(&_PermissionedRamp.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_PermissionedRamp *PermissionedRampCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedRamp.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_PermissionedRamp *PermissionedRampSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _PermissionedRamp.Contract.OwnershipHandoverExpiresAt(&_PermissionedRamp.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_PermissionedRamp *PermissionedRampCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _PermissionedRamp.Contract.OwnershipHandoverExpiresAt(&_PermissionedRamp.CallOpts, pendingOwner)
}

// Paused is a free data retrieval call binding the contract method 0x2e48152c.
//
// Solidity: function paused(address ) view returns(bool)
func (_PermissionedRamp *PermissionedRampCaller) Paused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PermissionedRamp.contract.Call(opts, &out, "paused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x2e48152c.
//
// Solidity: function paused(address ) view returns(bool)
func (_PermissionedRamp *PermissionedRampSession) Paused(arg0 common.Address) (bool, error) {
	return _PermissionedRamp.Contract.Paused(&_PermissionedRamp.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x2e48152c.
//
// Solidity: function paused(address ) view returns(bool)
func (_PermissionedRamp *PermissionedRampCallerSession) Paused(arg0 common.Address) (bool, error) {
	return _PermissionedRamp.Contract.Paused(&_PermissionedRamp.CallOpts, arg0)
}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_PermissionedRamp *PermissionedRampCaller) RolesOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedRamp.contract.Call(opts, &out, "rolesOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_PermissionedRamp *PermissionedRampSession) RolesOf(user common.Address) (*big.Int, error) {
	return _PermissionedRamp.Contract.RolesOf(&_PermissionedRamp.CallOpts, user)
}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_PermissionedRamp *PermissionedRampCallerSession) RolesOf(user common.Address) (*big.Int, error) {
	return _PermissionedRamp.Contract.RolesOf(&_PermissionedRamp.CallOpts, user)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_PermissionedRamp *PermissionedRampTransactor) AddAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.contract.Transact(opts, "addAdmin", _admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_PermissionedRamp *PermissionedRampSession) AddAdmin(_admin common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.AddAdmin(&_PermissionedRamp.TransactOpts, _admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_PermissionedRamp *PermissionedRampTransactorSession) AddAdmin(_admin common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.AddAdmin(&_PermissionedRamp.TransactOpts, _admin)
}

// AddWitness is a paid mutator transaction binding the contract method 0x59e26be1.
//
// Solidity: function addWitness(address _witness) returns()
func (_PermissionedRamp *PermissionedRampTransactor) AddWitness(opts *bind.TransactOpts, _witness common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.contract.Transact(opts, "addWitness", _witness)
}

// AddWitness is a paid mutator transaction binding the contract method 0x59e26be1.
//
// Solidity: function addWitness(address _witness) returns()
func (_PermissionedRamp *PermissionedRampSession) AddWitness(_witness common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.AddWitness(&_PermissionedRamp.TransactOpts, _witness)
}

// AddWitness is a paid mutator transaction binding the contract method 0x59e26be1.
//
// Solidity: function addWitness(address _witness) returns()
func (_PermissionedRamp *PermissionedRampTransactorSession) AddWitness(_witness common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.AddWitness(&_PermissionedRamp.TransactOpts, _witness)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_PermissionedRamp *PermissionedRampTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionedRamp.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_PermissionedRamp *PermissionedRampSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _PermissionedRamp.Contract.CancelOwnershipHandover(&_PermissionedRamp.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_PermissionedRamp *PermissionedRampTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _PermissionedRamp.Contract.CancelOwnershipHandover(&_PermissionedRamp.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_PermissionedRamp *PermissionedRampTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_PermissionedRamp *PermissionedRampSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.CompleteOwnershipHandover(&_PermissionedRamp.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_PermissionedRamp *PermissionedRampTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.CompleteOwnershipHandover(&_PermissionedRamp.TransactOpts, pendingOwner)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_PermissionedRamp *PermissionedRampTransactor) GrantRoles(opts *bind.TransactOpts, user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _PermissionedRamp.contract.Transact(opts, "grantRoles", user, roles)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_PermissionedRamp *PermissionedRampSession) GrantRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.GrantRoles(&_PermissionedRamp.TransactOpts, user, roles)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_PermissionedRamp *PermissionedRampTransactorSession) GrantRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.GrantRoles(&_PermissionedRamp.TransactOpts, user, roles)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address _asset) returns()
func (_PermissionedRamp *PermissionedRampTransactor) Pause(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.contract.Transact(opts, "pause", _asset)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address _asset) returns()
func (_PermissionedRamp *PermissionedRampSession) Pause(_asset common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.Pause(&_PermissionedRamp.TransactOpts, _asset)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address _asset) returns()
func (_PermissionedRamp *PermissionedRampTransactorSession) Pause(_asset common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.Pause(&_PermissionedRamp.TransactOpts, _asset)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_PermissionedRamp *PermissionedRampTransactor) RemoveAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.contract.Transact(opts, "removeAdmin", _admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_PermissionedRamp *PermissionedRampSession) RemoveAdmin(_admin common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.RemoveAdmin(&_PermissionedRamp.TransactOpts, _admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_PermissionedRamp *PermissionedRampTransactorSession) RemoveAdmin(_admin common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.RemoveAdmin(&_PermissionedRamp.TransactOpts, _admin)
}

// RemoveWitness is a paid mutator transaction binding the contract method 0xee2f13cd.
//
// Solidity: function removeWitness(address _witness) returns()
func (_PermissionedRamp *PermissionedRampTransactor) RemoveWitness(opts *bind.TransactOpts, _witness common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.contract.Transact(opts, "removeWitness", _witness)
}

// RemoveWitness is a paid mutator transaction binding the contract method 0xee2f13cd.
//
// Solidity: function removeWitness(address _witness) returns()
func (_PermissionedRamp *PermissionedRampSession) RemoveWitness(_witness common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.RemoveWitness(&_PermissionedRamp.TransactOpts, _witness)
}

// RemoveWitness is a paid mutator transaction binding the contract method 0xee2f13cd.
//
// Solidity: function removeWitness(address _witness) returns()
func (_PermissionedRamp *PermissionedRampTransactorSession) RemoveWitness(_witness common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.RemoveWitness(&_PermissionedRamp.TransactOpts, _witness)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_PermissionedRamp *PermissionedRampTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionedRamp.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_PermissionedRamp *PermissionedRampSession) RenounceOwnership() (*types.Transaction, error) {
	return _PermissionedRamp.Contract.RenounceOwnership(&_PermissionedRamp.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_PermissionedRamp *PermissionedRampTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PermissionedRamp.Contract.RenounceOwnership(&_PermissionedRamp.TransactOpts)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_PermissionedRamp *PermissionedRampTransactor) RenounceRoles(opts *bind.TransactOpts, roles *big.Int) (*types.Transaction, error) {
	return _PermissionedRamp.contract.Transact(opts, "renounceRoles", roles)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_PermissionedRamp *PermissionedRampSession) RenounceRoles(roles *big.Int) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.RenounceRoles(&_PermissionedRamp.TransactOpts, roles)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_PermissionedRamp *PermissionedRampTransactorSession) RenounceRoles(roles *big.Int) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.RenounceRoles(&_PermissionedRamp.TransactOpts, roles)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_PermissionedRamp *PermissionedRampTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionedRamp.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_PermissionedRamp *PermissionedRampSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _PermissionedRamp.Contract.RequestOwnershipHandover(&_PermissionedRamp.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_PermissionedRamp *PermissionedRampTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _PermissionedRamp.Contract.RequestOwnershipHandover(&_PermissionedRamp.TransactOpts)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_PermissionedRamp *PermissionedRampTransactor) RevokeRoles(opts *bind.TransactOpts, user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _PermissionedRamp.contract.Transact(opts, "revokeRoles", user, roles)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_PermissionedRamp *PermissionedRampSession) RevokeRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.RevokeRoles(&_PermissionedRamp.TransactOpts, user, roles)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_PermissionedRamp *PermissionedRampTransactorSession) RevokeRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.RevokeRoles(&_PermissionedRamp.TransactOpts, user, roles)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_PermissionedRamp *PermissionedRampTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_PermissionedRamp *PermissionedRampSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.TransferOwnership(&_PermissionedRamp.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_PermissionedRamp *PermissionedRampTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.TransferOwnership(&_PermissionedRamp.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x57b001f9.
//
// Solidity: function unpause(address _asset) returns()
func (_PermissionedRamp *PermissionedRampTransactor) Unpause(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.contract.Transact(opts, "unpause", _asset)
}

// Unpause is a paid mutator transaction binding the contract method 0x57b001f9.
//
// Solidity: function unpause(address _asset) returns()
func (_PermissionedRamp *PermissionedRampSession) Unpause(_asset common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.Unpause(&_PermissionedRamp.TransactOpts, _asset)
}

// Unpause is a paid mutator transaction binding the contract method 0x57b001f9.
//
// Solidity: function unpause(address _asset) returns()
func (_PermissionedRamp *PermissionedRampTransactorSession) Unpause(_asset common.Address) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.Unpause(&_PermissionedRamp.TransactOpts, _asset)
}

// Unwrap is a paid mutator transaction binding the contract method 0x117211b1.
//
// Solidity: function unwrap(address _asset, address _to, uint256 _amount, uint256 _nonce, uint256 _deadline, bytes _signature) returns()
func (_PermissionedRamp *PermissionedRampTransactor) Unwrap(opts *bind.TransactOpts, _asset common.Address, _to common.Address, _amount *big.Int, _nonce *big.Int, _deadline *big.Int, _signature []byte) (*types.Transaction, error) {
	return _PermissionedRamp.contract.Transact(opts, "unwrap", _asset, _to, _amount, _nonce, _deadline, _signature)
}

// Unwrap is a paid mutator transaction binding the contract method 0x117211b1.
//
// Solidity: function unwrap(address _asset, address _to, uint256 _amount, uint256 _nonce, uint256 _deadline, bytes _signature) returns()
func (_PermissionedRamp *PermissionedRampSession) Unwrap(_asset common.Address, _to common.Address, _amount *big.Int, _nonce *big.Int, _deadline *big.Int, _signature []byte) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.Unwrap(&_PermissionedRamp.TransactOpts, _asset, _to, _amount, _nonce, _deadline, _signature)
}

// Unwrap is a paid mutator transaction binding the contract method 0x117211b1.
//
// Solidity: function unwrap(address _asset, address _to, uint256 _amount, uint256 _nonce, uint256 _deadline, bytes _signature) returns()
func (_PermissionedRamp *PermissionedRampTransactorSession) Unwrap(_asset common.Address, _to common.Address, _amount *big.Int, _nonce *big.Int, _deadline *big.Int, _signature []byte) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.Unwrap(&_PermissionedRamp.TransactOpts, _asset, _to, _amount, _nonce, _deadline, _signature)
}

// Wrap is a paid mutator transaction binding the contract method 0x5f02639c.
//
// Solidity: function wrap(address _asset, address _to, uint256 _amount, uint256 _nonce, uint256 _deadline, bytes _signature) returns()
func (_PermissionedRamp *PermissionedRampTransactor) Wrap(opts *bind.TransactOpts, _asset common.Address, _to common.Address, _amount *big.Int, _nonce *big.Int, _deadline *big.Int, _signature []byte) (*types.Transaction, error) {
	return _PermissionedRamp.contract.Transact(opts, "wrap", _asset, _to, _amount, _nonce, _deadline, _signature)
}

// Wrap is a paid mutator transaction binding the contract method 0x5f02639c.
//
// Solidity: function wrap(address _asset, address _to, uint256 _amount, uint256 _nonce, uint256 _deadline, bytes _signature) returns()
func (_PermissionedRamp *PermissionedRampSession) Wrap(_asset common.Address, _to common.Address, _amount *big.Int, _nonce *big.Int, _deadline *big.Int, _signature []byte) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.Wrap(&_PermissionedRamp.TransactOpts, _asset, _to, _amount, _nonce, _deadline, _signature)
}

// Wrap is a paid mutator transaction binding the contract method 0x5f02639c.
//
// Solidity: function wrap(address _asset, address _to, uint256 _amount, uint256 _nonce, uint256 _deadline, bytes _signature) returns()
func (_PermissionedRamp *PermissionedRampTransactorSession) Wrap(_asset common.Address, _to common.Address, _amount *big.Int, _nonce *big.Int, _deadline *big.Int, _signature []byte) (*types.Transaction, error) {
	return _PermissionedRamp.Contract.Wrap(&_PermissionedRamp.TransactOpts, _asset, _to, _amount, _nonce, _deadline, _signature)
}

// PermissionedRampOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the PermissionedRamp contract.
type PermissionedRampOwnershipHandoverCanceledIterator struct {
	Event *PermissionedRampOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *PermissionedRampOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedRampOwnershipHandoverCanceled)
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
		it.Event = new(PermissionedRampOwnershipHandoverCanceled)
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
func (it *PermissionedRampOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedRampOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedRampOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the PermissionedRamp contract.
type PermissionedRampOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_PermissionedRamp *PermissionedRampFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*PermissionedRampOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _PermissionedRamp.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PermissionedRampOwnershipHandoverCanceledIterator{contract: _PermissionedRamp.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_PermissionedRamp *PermissionedRampFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *PermissionedRampOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _PermissionedRamp.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedRampOwnershipHandoverCanceled)
				if err := _PermissionedRamp.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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
func (_PermissionedRamp *PermissionedRampFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*PermissionedRampOwnershipHandoverCanceled, error) {
	event := new(PermissionedRampOwnershipHandoverCanceled)
	if err := _PermissionedRamp.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedRampOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the PermissionedRamp contract.
type PermissionedRampOwnershipHandoverRequestedIterator struct {
	Event *PermissionedRampOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *PermissionedRampOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedRampOwnershipHandoverRequested)
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
		it.Event = new(PermissionedRampOwnershipHandoverRequested)
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
func (it *PermissionedRampOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedRampOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedRampOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the PermissionedRamp contract.
type PermissionedRampOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_PermissionedRamp *PermissionedRampFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*PermissionedRampOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _PermissionedRamp.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PermissionedRampOwnershipHandoverRequestedIterator{contract: _PermissionedRamp.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_PermissionedRamp *PermissionedRampFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *PermissionedRampOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _PermissionedRamp.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedRampOwnershipHandoverRequested)
				if err := _PermissionedRamp.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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
func (_PermissionedRamp *PermissionedRampFilterer) ParseOwnershipHandoverRequested(log types.Log) (*PermissionedRampOwnershipHandoverRequested, error) {
	event := new(PermissionedRampOwnershipHandoverRequested)
	if err := _PermissionedRamp.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedRampOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PermissionedRamp contract.
type PermissionedRampOwnershipTransferredIterator struct {
	Event *PermissionedRampOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PermissionedRampOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedRampOwnershipTransferred)
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
		it.Event = new(PermissionedRampOwnershipTransferred)
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
func (it *PermissionedRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedRampOwnershipTransferred represents a OwnershipTransferred event raised by the PermissionedRamp contract.
type PermissionedRampOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_PermissionedRamp *PermissionedRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*PermissionedRampOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PermissionedRamp.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PermissionedRampOwnershipTransferredIterator{contract: _PermissionedRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_PermissionedRamp *PermissionedRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PermissionedRampOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PermissionedRamp.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedRampOwnershipTransferred)
				if err := _PermissionedRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PermissionedRamp *PermissionedRampFilterer) ParseOwnershipTransferred(log types.Log) (*PermissionedRampOwnershipTransferred, error) {
	event := new(PermissionedRampOwnershipTransferred)
	if err := _PermissionedRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedRampPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PermissionedRamp contract.
type PermissionedRampPausedIterator struct {
	Event *PermissionedRampPaused // Event containing the contract specifics and raw log

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
func (it *PermissionedRampPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedRampPaused)
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
		it.Event = new(PermissionedRampPaused)
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
func (it *PermissionedRampPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedRampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedRampPaused represents a Paused event raised by the PermissionedRamp contract.
type PermissionedRampPaused struct {
	Asset common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address indexed asset)
func (_PermissionedRamp *PermissionedRampFilterer) FilterPaused(opts *bind.FilterOpts, asset []common.Address) (*PermissionedRampPausedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _PermissionedRamp.contract.FilterLogs(opts, "Paused", assetRule)
	if err != nil {
		return nil, err
	}
	return &PermissionedRampPausedIterator{contract: _PermissionedRamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address indexed asset)
func (_PermissionedRamp *PermissionedRampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PermissionedRampPaused, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _PermissionedRamp.contract.WatchLogs(opts, "Paused", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedRampPaused)
				if err := _PermissionedRamp.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PermissionedRamp *PermissionedRampFilterer) ParsePaused(log types.Log) (*PermissionedRampPaused, error) {
	event := new(PermissionedRampPaused)
	if err := _PermissionedRamp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedRampRolesUpdatedIterator is returned from FilterRolesUpdated and is used to iterate over the raw logs and unpacked data for RolesUpdated events raised by the PermissionedRamp contract.
type PermissionedRampRolesUpdatedIterator struct {
	Event *PermissionedRampRolesUpdated // Event containing the contract specifics and raw log

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
func (it *PermissionedRampRolesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedRampRolesUpdated)
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
		it.Event = new(PermissionedRampRolesUpdated)
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
func (it *PermissionedRampRolesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedRampRolesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedRampRolesUpdated represents a RolesUpdated event raised by the PermissionedRamp contract.
type PermissionedRampRolesUpdated struct {
	User  common.Address
	Roles *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRolesUpdated is a free log retrieval operation binding the contract event 0x715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26.
//
// Solidity: event RolesUpdated(address indexed user, uint256 indexed roles)
func (_PermissionedRamp *PermissionedRampFilterer) FilterRolesUpdated(opts *bind.FilterOpts, user []common.Address, roles []*big.Int) (*PermissionedRampRolesUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rolesRule []interface{}
	for _, rolesItem := range roles {
		rolesRule = append(rolesRule, rolesItem)
	}

	logs, sub, err := _PermissionedRamp.contract.FilterLogs(opts, "RolesUpdated", userRule, rolesRule)
	if err != nil {
		return nil, err
	}
	return &PermissionedRampRolesUpdatedIterator{contract: _PermissionedRamp.contract, event: "RolesUpdated", logs: logs, sub: sub}, nil
}

// WatchRolesUpdated is a free log subscription operation binding the contract event 0x715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26.
//
// Solidity: event RolesUpdated(address indexed user, uint256 indexed roles)
func (_PermissionedRamp *PermissionedRampFilterer) WatchRolesUpdated(opts *bind.WatchOpts, sink chan<- *PermissionedRampRolesUpdated, user []common.Address, roles []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rolesRule []interface{}
	for _, rolesItem := range roles {
		rolesRule = append(rolesRule, rolesItem)
	}

	logs, sub, err := _PermissionedRamp.contract.WatchLogs(opts, "RolesUpdated", userRule, rolesRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedRampRolesUpdated)
				if err := _PermissionedRamp.contract.UnpackLog(event, "RolesUpdated", log); err != nil {
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
func (_PermissionedRamp *PermissionedRampFilterer) ParseRolesUpdated(log types.Log) (*PermissionedRampRolesUpdated, error) {
	event := new(PermissionedRampRolesUpdated)
	if err := _PermissionedRamp.contract.UnpackLog(event, "RolesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedRampUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PermissionedRamp contract.
type PermissionedRampUnpausedIterator struct {
	Event *PermissionedRampUnpaused // Event containing the contract specifics and raw log

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
func (it *PermissionedRampUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedRampUnpaused)
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
		it.Event = new(PermissionedRampUnpaused)
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
func (it *PermissionedRampUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedRampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedRampUnpaused represents a Unpaused event raised by the PermissionedRamp contract.
type PermissionedRampUnpaused struct {
	Asset common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address indexed asset)
func (_PermissionedRamp *PermissionedRampFilterer) FilterUnpaused(opts *bind.FilterOpts, asset []common.Address) (*PermissionedRampUnpausedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _PermissionedRamp.contract.FilterLogs(opts, "Unpaused", assetRule)
	if err != nil {
		return nil, err
	}
	return &PermissionedRampUnpausedIterator{contract: _PermissionedRamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address indexed asset)
func (_PermissionedRamp *PermissionedRampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PermissionedRampUnpaused, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _PermissionedRamp.contract.WatchLogs(opts, "Unpaused", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedRampUnpaused)
				if err := _PermissionedRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PermissionedRamp *PermissionedRampFilterer) ParseUnpaused(log types.Log) (*PermissionedRampUnpaused, error) {
	event := new(PermissionedRampUnpaused)
	if err := _PermissionedRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
