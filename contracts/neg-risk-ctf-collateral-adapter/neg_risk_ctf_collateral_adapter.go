// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package neg_risk_ctf_collateral_adapter

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

// NegRiskCtfCollateralAdapterMetaData contains all meta data concerning the NegRiskCtfCollateralAdapter contract.
var NegRiskCtfCollateralAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_conditionalTokens\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdce\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_negRiskAdapter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpiredDeadline\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAsset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoHandoverRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyUnpaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"RolesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COLLATERAL_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONDITIONAL_TOKENS\",\"outputs\":[{\"internalType\":\"contractIConditionalTokens\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NEG_RISK_ADAPTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDCE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WRAPPED_COLLATERAL\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"completeOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_marketId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_indexSet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"convertPositions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"grantRoles\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"hasAllRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"hasAnyRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_conditionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mergePositions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"result\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"ownershipHandoverExpiresAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_conditionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"redeemPositions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"renounceRoles\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"name\":\"revokeRoles\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"rolesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"roles\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_conditionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"splitPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NegRiskCtfCollateralAdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use NegRiskCtfCollateralAdapterMetaData.ABI instead.
var NegRiskCtfCollateralAdapterABI = NegRiskCtfCollateralAdapterMetaData.ABI

// NegRiskCtfCollateralAdapter is an auto generated Go binding around an Ethereum contract.
type NegRiskCtfCollateralAdapter struct {
	NegRiskCtfCollateralAdapterCaller     // Read-only binding to the contract
	NegRiskCtfCollateralAdapterTransactor // Write-only binding to the contract
	NegRiskCtfCollateralAdapterFilterer   // Log filterer for contract events
}

// NegRiskCtfCollateralAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type NegRiskCtfCollateralAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NegRiskCtfCollateralAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NegRiskCtfCollateralAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NegRiskCtfCollateralAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NegRiskCtfCollateralAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NegRiskCtfCollateralAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NegRiskCtfCollateralAdapterSession struct {
	Contract     *NegRiskCtfCollateralAdapter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// NegRiskCtfCollateralAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NegRiskCtfCollateralAdapterCallerSession struct {
	Contract *NegRiskCtfCollateralAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// NegRiskCtfCollateralAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NegRiskCtfCollateralAdapterTransactorSession struct {
	Contract     *NegRiskCtfCollateralAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// NegRiskCtfCollateralAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type NegRiskCtfCollateralAdapterRaw struct {
	Contract *NegRiskCtfCollateralAdapter // Generic contract binding to access the raw methods on
}

// NegRiskCtfCollateralAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NegRiskCtfCollateralAdapterCallerRaw struct {
	Contract *NegRiskCtfCollateralAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// NegRiskCtfCollateralAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NegRiskCtfCollateralAdapterTransactorRaw struct {
	Contract *NegRiskCtfCollateralAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNegRiskCtfCollateralAdapter creates a new instance of NegRiskCtfCollateralAdapter, bound to a specific deployed contract.
func NewNegRiskCtfCollateralAdapter(address common.Address, backend bind.ContractBackend) (*NegRiskCtfCollateralAdapter, error) {
	contract, err := bindNegRiskCtfCollateralAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfCollateralAdapter{NegRiskCtfCollateralAdapterCaller: NegRiskCtfCollateralAdapterCaller{contract: contract}, NegRiskCtfCollateralAdapterTransactor: NegRiskCtfCollateralAdapterTransactor{contract: contract}, NegRiskCtfCollateralAdapterFilterer: NegRiskCtfCollateralAdapterFilterer{contract: contract}}, nil
}

// NewNegRiskCtfCollateralAdapterCaller creates a new read-only instance of NegRiskCtfCollateralAdapter, bound to a specific deployed contract.
func NewNegRiskCtfCollateralAdapterCaller(address common.Address, caller bind.ContractCaller) (*NegRiskCtfCollateralAdapterCaller, error) {
	contract, err := bindNegRiskCtfCollateralAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfCollateralAdapterCaller{contract: contract}, nil
}

// NewNegRiskCtfCollateralAdapterTransactor creates a new write-only instance of NegRiskCtfCollateralAdapter, bound to a specific deployed contract.
func NewNegRiskCtfCollateralAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*NegRiskCtfCollateralAdapterTransactor, error) {
	contract, err := bindNegRiskCtfCollateralAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfCollateralAdapterTransactor{contract: contract}, nil
}

// NewNegRiskCtfCollateralAdapterFilterer creates a new log filterer instance of NegRiskCtfCollateralAdapter, bound to a specific deployed contract.
func NewNegRiskCtfCollateralAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*NegRiskCtfCollateralAdapterFilterer, error) {
	contract, err := bindNegRiskCtfCollateralAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfCollateralAdapterFilterer{contract: contract}, nil
}

// bindNegRiskCtfCollateralAdapter binds a generic wrapper to an already deployed contract.
func bindNegRiskCtfCollateralAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NegRiskCtfCollateralAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NegRiskCtfCollateralAdapter.Contract.NegRiskCtfCollateralAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.NegRiskCtfCollateralAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.NegRiskCtfCollateralAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NegRiskCtfCollateralAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.contract.Transact(opts, method, params...)
}

// COLLATERALTOKEN is a free data retrieval call binding the contract method 0xf5f1f1a7.
//
// Solidity: function COLLATERAL_TOKEN() view returns(address)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCaller) COLLATERALTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCtfCollateralAdapter.contract.Call(opts, &out, "COLLATERAL_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// COLLATERALTOKEN is a free data retrieval call binding the contract method 0xf5f1f1a7.
//
// Solidity: function COLLATERAL_TOKEN() view returns(address)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) COLLATERALTOKEN() (common.Address, error) {
	return _NegRiskCtfCollateralAdapter.Contract.COLLATERALTOKEN(&_NegRiskCtfCollateralAdapter.CallOpts)
}

// COLLATERALTOKEN is a free data retrieval call binding the contract method 0xf5f1f1a7.
//
// Solidity: function COLLATERAL_TOKEN() view returns(address)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCallerSession) COLLATERALTOKEN() (common.Address, error) {
	return _NegRiskCtfCollateralAdapter.Contract.COLLATERALTOKEN(&_NegRiskCtfCollateralAdapter.CallOpts)
}

// CONDITIONALTOKENS is a free data retrieval call binding the contract method 0x165d1f36.
//
// Solidity: function CONDITIONAL_TOKENS() view returns(address)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCaller) CONDITIONALTOKENS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCtfCollateralAdapter.contract.Call(opts, &out, "CONDITIONAL_TOKENS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CONDITIONALTOKENS is a free data retrieval call binding the contract method 0x165d1f36.
//
// Solidity: function CONDITIONAL_TOKENS() view returns(address)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) CONDITIONALTOKENS() (common.Address, error) {
	return _NegRiskCtfCollateralAdapter.Contract.CONDITIONALTOKENS(&_NegRiskCtfCollateralAdapter.CallOpts)
}

// CONDITIONALTOKENS is a free data retrieval call binding the contract method 0x165d1f36.
//
// Solidity: function CONDITIONAL_TOKENS() view returns(address)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCallerSession) CONDITIONALTOKENS() (common.Address, error) {
	return _NegRiskCtfCollateralAdapter.Contract.CONDITIONALTOKENS(&_NegRiskCtfCollateralAdapter.CallOpts)
}

// NEGRISKADAPTER is a free data retrieval call binding the contract method 0xf6f88a8d.
//
// Solidity: function NEG_RISK_ADAPTER() view returns(address)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCaller) NEGRISKADAPTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCtfCollateralAdapter.contract.Call(opts, &out, "NEG_RISK_ADAPTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NEGRISKADAPTER is a free data retrieval call binding the contract method 0xf6f88a8d.
//
// Solidity: function NEG_RISK_ADAPTER() view returns(address)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) NEGRISKADAPTER() (common.Address, error) {
	return _NegRiskCtfCollateralAdapter.Contract.NEGRISKADAPTER(&_NegRiskCtfCollateralAdapter.CallOpts)
}

// NEGRISKADAPTER is a free data retrieval call binding the contract method 0xf6f88a8d.
//
// Solidity: function NEG_RISK_ADAPTER() view returns(address)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCallerSession) NEGRISKADAPTER() (common.Address, error) {
	return _NegRiskCtfCollateralAdapter.Contract.NEGRISKADAPTER(&_NegRiskCtfCollateralAdapter.CallOpts)
}

// USDCE is a free data retrieval call binding the contract method 0x195187e1.
//
// Solidity: function USDCE() view returns(address)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCaller) USDCE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCtfCollateralAdapter.contract.Call(opts, &out, "USDCE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDCE is a free data retrieval call binding the contract method 0x195187e1.
//
// Solidity: function USDCE() view returns(address)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) USDCE() (common.Address, error) {
	return _NegRiskCtfCollateralAdapter.Contract.USDCE(&_NegRiskCtfCollateralAdapter.CallOpts)
}

// USDCE is a free data retrieval call binding the contract method 0x195187e1.
//
// Solidity: function USDCE() view returns(address)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCallerSession) USDCE() (common.Address, error) {
	return _NegRiskCtfCollateralAdapter.Contract.USDCE(&_NegRiskCtfCollateralAdapter.CallOpts)
}

// WRAPPEDCOLLATERAL is a free data retrieval call binding the contract method 0x2d277260.
//
// Solidity: function WRAPPED_COLLATERAL() view returns(address)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCaller) WRAPPEDCOLLATERAL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCtfCollateralAdapter.contract.Call(opts, &out, "WRAPPED_COLLATERAL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WRAPPEDCOLLATERAL is a free data retrieval call binding the contract method 0x2d277260.
//
// Solidity: function WRAPPED_COLLATERAL() view returns(address)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) WRAPPEDCOLLATERAL() (common.Address, error) {
	return _NegRiskCtfCollateralAdapter.Contract.WRAPPEDCOLLATERAL(&_NegRiskCtfCollateralAdapter.CallOpts)
}

// WRAPPEDCOLLATERAL is a free data retrieval call binding the contract method 0x2d277260.
//
// Solidity: function WRAPPED_COLLATERAL() view returns(address)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCallerSession) WRAPPEDCOLLATERAL() (common.Address, error) {
	return _NegRiskCtfCollateralAdapter.Contract.WRAPPEDCOLLATERAL(&_NegRiskCtfCollateralAdapter.CallOpts)
}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCaller) HasAllRoles(opts *bind.CallOpts, user common.Address, roles *big.Int) (bool, error) {
	var out []interface{}
	err := _NegRiskCtfCollateralAdapter.contract.Call(opts, &out, "hasAllRoles", user, roles)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) HasAllRoles(user common.Address, roles *big.Int) (bool, error) {
	return _NegRiskCtfCollateralAdapter.Contract.HasAllRoles(&_NegRiskCtfCollateralAdapter.CallOpts, user, roles)
}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCallerSession) HasAllRoles(user common.Address, roles *big.Int) (bool, error) {
	return _NegRiskCtfCollateralAdapter.Contract.HasAllRoles(&_NegRiskCtfCollateralAdapter.CallOpts, user, roles)
}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCaller) HasAnyRole(opts *bind.CallOpts, user common.Address, roles *big.Int) (bool, error) {
	var out []interface{}
	err := _NegRiskCtfCollateralAdapter.contract.Call(opts, &out, "hasAnyRole", user, roles)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) HasAnyRole(user common.Address, roles *big.Int) (bool, error) {
	return _NegRiskCtfCollateralAdapter.Contract.HasAnyRole(&_NegRiskCtfCollateralAdapter.CallOpts, user, roles)
}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCallerSession) HasAnyRole(user common.Address, roles *big.Int) (bool, error) {
	return _NegRiskCtfCollateralAdapter.Contract.HasAnyRole(&_NegRiskCtfCollateralAdapter.CallOpts, user, roles)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCtfCollateralAdapter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) Owner() (common.Address, error) {
	return _NegRiskCtfCollateralAdapter.Contract.Owner(&_NegRiskCtfCollateralAdapter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCallerSession) Owner() (common.Address, error) {
	return _NegRiskCtfCollateralAdapter.Contract.Owner(&_NegRiskCtfCollateralAdapter.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NegRiskCtfCollateralAdapter.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _NegRiskCtfCollateralAdapter.Contract.OwnershipHandoverExpiresAt(&_NegRiskCtfCollateralAdapter.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _NegRiskCtfCollateralAdapter.Contract.OwnershipHandoverExpiresAt(&_NegRiskCtfCollateralAdapter.CallOpts, pendingOwner)
}

// Paused is a free data retrieval call binding the contract method 0x2e48152c.
//
// Solidity: function paused(address ) view returns(bool)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCaller) Paused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _NegRiskCtfCollateralAdapter.contract.Call(opts, &out, "paused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x2e48152c.
//
// Solidity: function paused(address ) view returns(bool)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) Paused(arg0 common.Address) (bool, error) {
	return _NegRiskCtfCollateralAdapter.Contract.Paused(&_NegRiskCtfCollateralAdapter.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x2e48152c.
//
// Solidity: function paused(address ) view returns(bool)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCallerSession) Paused(arg0 common.Address) (bool, error) {
	return _NegRiskCtfCollateralAdapter.Contract.Paused(&_NegRiskCtfCollateralAdapter.CallOpts, arg0)
}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCaller) RolesOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NegRiskCtfCollateralAdapter.contract.Call(opts, &out, "rolesOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) RolesOf(user common.Address) (*big.Int, error) {
	return _NegRiskCtfCollateralAdapter.Contract.RolesOf(&_NegRiskCtfCollateralAdapter.CallOpts, user)
}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCallerSession) RolesOf(user common.Address) (*big.Int, error) {
	return _NegRiskCtfCollateralAdapter.Contract.RolesOf(&_NegRiskCtfCollateralAdapter.CallOpts, user)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NegRiskCtfCollateralAdapter.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NegRiskCtfCollateralAdapter.Contract.SupportsInterface(&_NegRiskCtfCollateralAdapter.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NegRiskCtfCollateralAdapter.Contract.SupportsInterface(&_NegRiskCtfCollateralAdapter.CallOpts, interfaceId)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactor) AddAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.contract.Transact(opts, "addAdmin", _admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) AddAdmin(_admin common.Address) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.AddAdmin(&_NegRiskCtfCollateralAdapter.TransactOpts, _admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactorSession) AddAdmin(_admin common.Address) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.AddAdmin(&_NegRiskCtfCollateralAdapter.TransactOpts, _admin)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.CancelOwnershipHandover(&_NegRiskCtfCollateralAdapter.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.CancelOwnershipHandover(&_NegRiskCtfCollateralAdapter.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.CompleteOwnershipHandover(&_NegRiskCtfCollateralAdapter.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.CompleteOwnershipHandover(&_NegRiskCtfCollateralAdapter.TransactOpts, pendingOwner)
}

// ConvertPositions is a paid mutator transaction binding the contract method 0xc64748c4.
//
// Solidity: function convertPositions(bytes32 _marketId, uint256 _indexSet, uint256 _amount) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactor) ConvertPositions(opts *bind.TransactOpts, _marketId [32]byte, _indexSet *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.contract.Transact(opts, "convertPositions", _marketId, _indexSet, _amount)
}

// ConvertPositions is a paid mutator transaction binding the contract method 0xc64748c4.
//
// Solidity: function convertPositions(bytes32 _marketId, uint256 _indexSet, uint256 _amount) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) ConvertPositions(_marketId [32]byte, _indexSet *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.ConvertPositions(&_NegRiskCtfCollateralAdapter.TransactOpts, _marketId, _indexSet, _amount)
}

// ConvertPositions is a paid mutator transaction binding the contract method 0xc64748c4.
//
// Solidity: function convertPositions(bytes32 _marketId, uint256 _indexSet, uint256 _amount) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactorSession) ConvertPositions(_marketId [32]byte, _indexSet *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.ConvertPositions(&_NegRiskCtfCollateralAdapter.TransactOpts, _marketId, _indexSet, _amount)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactor) GrantRoles(opts *bind.TransactOpts, user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.contract.Transact(opts, "grantRoles", user, roles)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) GrantRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.GrantRoles(&_NegRiskCtfCollateralAdapter.TransactOpts, user, roles)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactorSession) GrantRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.GrantRoles(&_NegRiskCtfCollateralAdapter.TransactOpts, user, roles)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address , bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactor) MergePositions(opts *bind.TransactOpts, arg0 common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.contract.Transact(opts, "mergePositions", arg0, arg1, _conditionId, arg3, _amount)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address , bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) MergePositions(arg0 common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.MergePositions(&_NegRiskCtfCollateralAdapter.TransactOpts, arg0, arg1, _conditionId, arg3, _amount)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address , bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactorSession) MergePositions(arg0 common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.MergePositions(&_NegRiskCtfCollateralAdapter.TransactOpts, arg0, arg1, _conditionId, arg3, _amount)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.OnERC1155BatchReceived(&_NegRiskCtfCollateralAdapter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.OnERC1155BatchReceived(&_NegRiskCtfCollateralAdapter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.OnERC1155Received(&_NegRiskCtfCollateralAdapter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.OnERC1155Received(&_NegRiskCtfCollateralAdapter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address _asset) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactor) Pause(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.contract.Transact(opts, "pause", _asset)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address _asset) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) Pause(_asset common.Address) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.Pause(&_NegRiskCtfCollateralAdapter.TransactOpts, _asset)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address _asset) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactorSession) Pause(_asset common.Address) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.Pause(&_NegRiskCtfCollateralAdapter.TransactOpts, _asset)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0x01b7037c.
//
// Solidity: function redeemPositions(address , bytes32 , bytes32 _conditionId, uint256[] ) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactor) RedeemPositions(opts *bind.TransactOpts, arg0 common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.contract.Transact(opts, "redeemPositions", arg0, arg1, _conditionId, arg3)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0x01b7037c.
//
// Solidity: function redeemPositions(address , bytes32 , bytes32 _conditionId, uint256[] ) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) RedeemPositions(arg0 common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.RedeemPositions(&_NegRiskCtfCollateralAdapter.TransactOpts, arg0, arg1, _conditionId, arg3)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0x01b7037c.
//
// Solidity: function redeemPositions(address , bytes32 , bytes32 _conditionId, uint256[] ) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactorSession) RedeemPositions(arg0 common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.RedeemPositions(&_NegRiskCtfCollateralAdapter.TransactOpts, arg0, arg1, _conditionId, arg3)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactor) RemoveAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.contract.Transact(opts, "removeAdmin", _admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) RemoveAdmin(_admin common.Address) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.RemoveAdmin(&_NegRiskCtfCollateralAdapter.TransactOpts, _admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactorSession) RemoveAdmin(_admin common.Address) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.RemoveAdmin(&_NegRiskCtfCollateralAdapter.TransactOpts, _admin)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) RenounceOwnership() (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.RenounceOwnership(&_NegRiskCtfCollateralAdapter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.RenounceOwnership(&_NegRiskCtfCollateralAdapter.TransactOpts)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactor) RenounceRoles(opts *bind.TransactOpts, roles *big.Int) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.contract.Transact(opts, "renounceRoles", roles)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) RenounceRoles(roles *big.Int) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.RenounceRoles(&_NegRiskCtfCollateralAdapter.TransactOpts, roles)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactorSession) RenounceRoles(roles *big.Int) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.RenounceRoles(&_NegRiskCtfCollateralAdapter.TransactOpts, roles)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.RequestOwnershipHandover(&_NegRiskCtfCollateralAdapter.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.RequestOwnershipHandover(&_NegRiskCtfCollateralAdapter.TransactOpts)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactor) RevokeRoles(opts *bind.TransactOpts, user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.contract.Transact(opts, "revokeRoles", user, roles)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) RevokeRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.RevokeRoles(&_NegRiskCtfCollateralAdapter.TransactOpts, user, roles)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactorSession) RevokeRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.RevokeRoles(&_NegRiskCtfCollateralAdapter.TransactOpts, user, roles)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address , bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactor) SplitPosition(opts *bind.TransactOpts, arg0 common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.contract.Transact(opts, "splitPosition", arg0, arg1, _conditionId, arg3, _amount)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address , bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) SplitPosition(arg0 common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.SplitPosition(&_NegRiskCtfCollateralAdapter.TransactOpts, arg0, arg1, _conditionId, arg3, _amount)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address , bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactorSession) SplitPosition(arg0 common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.SplitPosition(&_NegRiskCtfCollateralAdapter.TransactOpts, arg0, arg1, _conditionId, arg3, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.TransferOwnership(&_NegRiskCtfCollateralAdapter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.TransferOwnership(&_NegRiskCtfCollateralAdapter.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x57b001f9.
//
// Solidity: function unpause(address _asset) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactor) Unpause(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.contract.Transact(opts, "unpause", _asset)
}

// Unpause is a paid mutator transaction binding the contract method 0x57b001f9.
//
// Solidity: function unpause(address _asset) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterSession) Unpause(_asset common.Address) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.Unpause(&_NegRiskCtfCollateralAdapter.TransactOpts, _asset)
}

// Unpause is a paid mutator transaction binding the contract method 0x57b001f9.
//
// Solidity: function unpause(address _asset) returns()
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterTransactorSession) Unpause(_asset common.Address) (*types.Transaction, error) {
	return _NegRiskCtfCollateralAdapter.Contract.Unpause(&_NegRiskCtfCollateralAdapter.TransactOpts, _asset)
}

// NegRiskCtfCollateralAdapterOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the NegRiskCtfCollateralAdapter contract.
type NegRiskCtfCollateralAdapterOwnershipHandoverCanceledIterator struct {
	Event *NegRiskCtfCollateralAdapterOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *NegRiskCtfCollateralAdapterOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCtfCollateralAdapterOwnershipHandoverCanceled)
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
		it.Event = new(NegRiskCtfCollateralAdapterOwnershipHandoverCanceled)
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
func (it *NegRiskCtfCollateralAdapterOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCtfCollateralAdapterOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCtfCollateralAdapterOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the NegRiskCtfCollateralAdapter contract.
type NegRiskCtfCollateralAdapterOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*NegRiskCtfCollateralAdapterOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _NegRiskCtfCollateralAdapter.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfCollateralAdapterOwnershipHandoverCanceledIterator{contract: _NegRiskCtfCollateralAdapter.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *NegRiskCtfCollateralAdapterOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _NegRiskCtfCollateralAdapter.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCtfCollateralAdapterOwnershipHandoverCanceled)
				if err := _NegRiskCtfCollateralAdapter.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*NegRiskCtfCollateralAdapterOwnershipHandoverCanceled, error) {
	event := new(NegRiskCtfCollateralAdapterOwnershipHandoverCanceled)
	if err := _NegRiskCtfCollateralAdapter.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCtfCollateralAdapterOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the NegRiskCtfCollateralAdapter contract.
type NegRiskCtfCollateralAdapterOwnershipHandoverRequestedIterator struct {
	Event *NegRiskCtfCollateralAdapterOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *NegRiskCtfCollateralAdapterOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCtfCollateralAdapterOwnershipHandoverRequested)
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
		it.Event = new(NegRiskCtfCollateralAdapterOwnershipHandoverRequested)
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
func (it *NegRiskCtfCollateralAdapterOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCtfCollateralAdapterOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCtfCollateralAdapterOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the NegRiskCtfCollateralAdapter contract.
type NegRiskCtfCollateralAdapterOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*NegRiskCtfCollateralAdapterOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _NegRiskCtfCollateralAdapter.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfCollateralAdapterOwnershipHandoverRequestedIterator{contract: _NegRiskCtfCollateralAdapter.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *NegRiskCtfCollateralAdapterOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _NegRiskCtfCollateralAdapter.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCtfCollateralAdapterOwnershipHandoverRequested)
				if err := _NegRiskCtfCollateralAdapter.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterFilterer) ParseOwnershipHandoverRequested(log types.Log) (*NegRiskCtfCollateralAdapterOwnershipHandoverRequested, error) {
	event := new(NegRiskCtfCollateralAdapterOwnershipHandoverRequested)
	if err := _NegRiskCtfCollateralAdapter.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCtfCollateralAdapterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NegRiskCtfCollateralAdapter contract.
type NegRiskCtfCollateralAdapterOwnershipTransferredIterator struct {
	Event *NegRiskCtfCollateralAdapterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NegRiskCtfCollateralAdapterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCtfCollateralAdapterOwnershipTransferred)
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
		it.Event = new(NegRiskCtfCollateralAdapterOwnershipTransferred)
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
func (it *NegRiskCtfCollateralAdapterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCtfCollateralAdapterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCtfCollateralAdapterOwnershipTransferred represents a OwnershipTransferred event raised by the NegRiskCtfCollateralAdapter contract.
type NegRiskCtfCollateralAdapterOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*NegRiskCtfCollateralAdapterOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NegRiskCtfCollateralAdapter.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfCollateralAdapterOwnershipTransferredIterator{contract: _NegRiskCtfCollateralAdapter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NegRiskCtfCollateralAdapterOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NegRiskCtfCollateralAdapter.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCtfCollateralAdapterOwnershipTransferred)
				if err := _NegRiskCtfCollateralAdapter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterFilterer) ParseOwnershipTransferred(log types.Log) (*NegRiskCtfCollateralAdapterOwnershipTransferred, error) {
	event := new(NegRiskCtfCollateralAdapterOwnershipTransferred)
	if err := _NegRiskCtfCollateralAdapter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCtfCollateralAdapterPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the NegRiskCtfCollateralAdapter contract.
type NegRiskCtfCollateralAdapterPausedIterator struct {
	Event *NegRiskCtfCollateralAdapterPaused // Event containing the contract specifics and raw log

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
func (it *NegRiskCtfCollateralAdapterPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCtfCollateralAdapterPaused)
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
		it.Event = new(NegRiskCtfCollateralAdapterPaused)
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
func (it *NegRiskCtfCollateralAdapterPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCtfCollateralAdapterPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCtfCollateralAdapterPaused represents a Paused event raised by the NegRiskCtfCollateralAdapter contract.
type NegRiskCtfCollateralAdapterPaused struct {
	Asset common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address indexed asset)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterFilterer) FilterPaused(opts *bind.FilterOpts, asset []common.Address) (*NegRiskCtfCollateralAdapterPausedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _NegRiskCtfCollateralAdapter.contract.FilterLogs(opts, "Paused", assetRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfCollateralAdapterPausedIterator{contract: _NegRiskCtfCollateralAdapter.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address indexed asset)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *NegRiskCtfCollateralAdapterPaused, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _NegRiskCtfCollateralAdapter.contract.WatchLogs(opts, "Paused", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCtfCollateralAdapterPaused)
				if err := _NegRiskCtfCollateralAdapter.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterFilterer) ParsePaused(log types.Log) (*NegRiskCtfCollateralAdapterPaused, error) {
	event := new(NegRiskCtfCollateralAdapterPaused)
	if err := _NegRiskCtfCollateralAdapter.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCtfCollateralAdapterRolesUpdatedIterator is returned from FilterRolesUpdated and is used to iterate over the raw logs and unpacked data for RolesUpdated events raised by the NegRiskCtfCollateralAdapter contract.
type NegRiskCtfCollateralAdapterRolesUpdatedIterator struct {
	Event *NegRiskCtfCollateralAdapterRolesUpdated // Event containing the contract specifics and raw log

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
func (it *NegRiskCtfCollateralAdapterRolesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCtfCollateralAdapterRolesUpdated)
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
		it.Event = new(NegRiskCtfCollateralAdapterRolesUpdated)
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
func (it *NegRiskCtfCollateralAdapterRolesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCtfCollateralAdapterRolesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCtfCollateralAdapterRolesUpdated represents a RolesUpdated event raised by the NegRiskCtfCollateralAdapter contract.
type NegRiskCtfCollateralAdapterRolesUpdated struct {
	User  common.Address
	Roles *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRolesUpdated is a free log retrieval operation binding the contract event 0x715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26.
//
// Solidity: event RolesUpdated(address indexed user, uint256 indexed roles)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterFilterer) FilterRolesUpdated(opts *bind.FilterOpts, user []common.Address, roles []*big.Int) (*NegRiskCtfCollateralAdapterRolesUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rolesRule []interface{}
	for _, rolesItem := range roles {
		rolesRule = append(rolesRule, rolesItem)
	}

	logs, sub, err := _NegRiskCtfCollateralAdapter.contract.FilterLogs(opts, "RolesUpdated", userRule, rolesRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfCollateralAdapterRolesUpdatedIterator{contract: _NegRiskCtfCollateralAdapter.contract, event: "RolesUpdated", logs: logs, sub: sub}, nil
}

// WatchRolesUpdated is a free log subscription operation binding the contract event 0x715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26.
//
// Solidity: event RolesUpdated(address indexed user, uint256 indexed roles)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterFilterer) WatchRolesUpdated(opts *bind.WatchOpts, sink chan<- *NegRiskCtfCollateralAdapterRolesUpdated, user []common.Address, roles []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rolesRule []interface{}
	for _, rolesItem := range roles {
		rolesRule = append(rolesRule, rolesItem)
	}

	logs, sub, err := _NegRiskCtfCollateralAdapter.contract.WatchLogs(opts, "RolesUpdated", userRule, rolesRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCtfCollateralAdapterRolesUpdated)
				if err := _NegRiskCtfCollateralAdapter.contract.UnpackLog(event, "RolesUpdated", log); err != nil {
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
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterFilterer) ParseRolesUpdated(log types.Log) (*NegRiskCtfCollateralAdapterRolesUpdated, error) {
	event := new(NegRiskCtfCollateralAdapterRolesUpdated)
	if err := _NegRiskCtfCollateralAdapter.contract.UnpackLog(event, "RolesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCtfCollateralAdapterUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the NegRiskCtfCollateralAdapter contract.
type NegRiskCtfCollateralAdapterUnpausedIterator struct {
	Event *NegRiskCtfCollateralAdapterUnpaused // Event containing the contract specifics and raw log

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
func (it *NegRiskCtfCollateralAdapterUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCtfCollateralAdapterUnpaused)
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
		it.Event = new(NegRiskCtfCollateralAdapterUnpaused)
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
func (it *NegRiskCtfCollateralAdapterUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCtfCollateralAdapterUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCtfCollateralAdapterUnpaused represents a Unpaused event raised by the NegRiskCtfCollateralAdapter contract.
type NegRiskCtfCollateralAdapterUnpaused struct {
	Asset common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address indexed asset)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterFilterer) FilterUnpaused(opts *bind.FilterOpts, asset []common.Address) (*NegRiskCtfCollateralAdapterUnpausedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _NegRiskCtfCollateralAdapter.contract.FilterLogs(opts, "Unpaused", assetRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCtfCollateralAdapterUnpausedIterator{contract: _NegRiskCtfCollateralAdapter.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address indexed asset)
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *NegRiskCtfCollateralAdapterUnpaused, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _NegRiskCtfCollateralAdapter.contract.WatchLogs(opts, "Unpaused", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCtfCollateralAdapterUnpaused)
				if err := _NegRiskCtfCollateralAdapter.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_NegRiskCtfCollateralAdapter *NegRiskCtfCollateralAdapterFilterer) ParseUnpaused(log types.Log) (*NegRiskCtfCollateralAdapterUnpaused, error) {
	event := new(NegRiskCtfCollateralAdapterUnpaused)
	if err := _NegRiskCtfCollateralAdapter.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
