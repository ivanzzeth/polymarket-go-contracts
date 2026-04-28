// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package neg_risk_v2

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

// ExchangeInitParams is an auto generated low-level Go binding around an user-defined struct.
type ExchangeInitParams struct {
	Admin               common.Address
	Collateral          common.Address
	Ctf                 common.Address
	CtfCollateral       common.Address
	OutcomeTokenFactory common.Address
	ProxyFactory        common.Address
	SafeFactory         common.Address
	FeeReceiver         common.Address
}

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Salt          *big.Int
	Maker         common.Address
	Signer        common.Address
	TokenId       *big.Int
	MakerAmount   *big.Int
	TakerAmount   *big.Int
	Side          uint8
	SignatureType uint8
	Timestamp     *big.Int
	Metadata      [32]byte
	Builder       [32]byte
	Signature     []byte
}

// OrderStatus is an auto generated low-level Go binding around an user-defined struct.
type OrderStatus struct {
	Filled    bool
	Remaining *big.Int
}

// NegRiskV2MetaData contains all meta data concerning the NegRiskV2 contract.
var NegRiskV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ctf\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ctfCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"outcomeTokenFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proxyFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"safeFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"}],\"internalType\":\"structExchangeInitParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ComplementaryFillExceedsTakerFill\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedsMaxPauseInterval\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeExceedsMaxRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeExceedsProceeds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LastAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MakingGtRemaining\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxFeeRateExceedsCeiling\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MismatchedArrayLengths\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MismatchedTokenIds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMakerOrders\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCrossing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderAlreadyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooLittleTokensReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserAlreadyPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserIsPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroMakerAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"}],\"name\":\"FeeReceiverUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxFeeRate\",\"type\":\"uint256\"}],\"name\":\"MaxFeeRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdminAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOperatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"NewOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAmountFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAmountFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"builder\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"metadata\",\"type\":\"bytes32\"}],\"name\":\"OrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderPreapprovalInvalidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderPreapproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"takerOrderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"takerOrderMaker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAmountFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAmountFilled\",\"type\":\"uint256\"}],\"name\":\"OrdersMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"RemovedAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedOperator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"RemovedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"TradingPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"TradingUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldInterval\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newInterval\",\"type\":\"uint256\"}],\"name\":\"UserPauseBlockIntervalUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"effectivePauseBlock\",\"type\":\"uint256\"}],\"name\":\"UserPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"UserUnpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PARENT_COLLECTION_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollateral\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCtf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCtfCollateral\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"getOrderStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"filled\",\"type\":\"bool\"},{\"internalType\":\"uint248\",\"name\":\"remaining\",\"type\":\"uint248\"}],\"internalType\":\"structOrderStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOutcomeTokenFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxyFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxyImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getProxyWalletAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSafeFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSafeImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getSafeWalletAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"metadata\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"builder\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"hashOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"invalidatePreapprovedOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usr\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isUserPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"metadata\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"builder\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"takerOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"metadata\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"builder\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"makerOrders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"takerFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"makerFillAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"takerFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"makerFeeAmounts\",\"type\":\"uint256[]\"}],\"name\":\"matchOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"filled\",\"type\":\"bool\"},{\"internalType\":\"uint248\",\"name\":\"remaining\",\"type\":\"uint248\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseTrading\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"metadata\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"builder\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"preapproveOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOperatorRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"setFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"setMaxFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_interval\",\"type\":\"uint256\"}],\"name\":\"setUserPauseBlockInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseTrading\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userPauseBlockInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userPausedBlockAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cashValue\",\"type\":\"uint256\"}],\"name\":\"validateFee\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"metadata\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"builder\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"validateOrder\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"metadata\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"builder\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"validateOrderSignature\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NegRiskV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use NegRiskV2MetaData.ABI instead.
var NegRiskV2ABI = NegRiskV2MetaData.ABI

// NegRiskV2 is an auto generated Go binding around an Ethereum contract.
type NegRiskV2 struct {
	NegRiskV2Caller     // Read-only binding to the contract
	NegRiskV2Transactor // Write-only binding to the contract
	NegRiskV2Filterer   // Log filterer for contract events
}

// NegRiskV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type NegRiskV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NegRiskV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type NegRiskV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NegRiskV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NegRiskV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NegRiskV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NegRiskV2Session struct {
	Contract     *NegRiskV2        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NegRiskV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NegRiskV2CallerSession struct {
	Contract *NegRiskV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// NegRiskV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NegRiskV2TransactorSession struct {
	Contract     *NegRiskV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// NegRiskV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type NegRiskV2Raw struct {
	Contract *NegRiskV2 // Generic contract binding to access the raw methods on
}

// NegRiskV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NegRiskV2CallerRaw struct {
	Contract *NegRiskV2Caller // Generic read-only contract binding to access the raw methods on
}

// NegRiskV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NegRiskV2TransactorRaw struct {
	Contract *NegRiskV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewNegRiskV2 creates a new instance of NegRiskV2, bound to a specific deployed contract.
func NewNegRiskV2(address common.Address, backend bind.ContractBackend) (*NegRiskV2, error) {
	contract, err := bindNegRiskV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NegRiskV2{NegRiskV2Caller: NegRiskV2Caller{contract: contract}, NegRiskV2Transactor: NegRiskV2Transactor{contract: contract}, NegRiskV2Filterer: NegRiskV2Filterer{contract: contract}}, nil
}

// NewNegRiskV2Caller creates a new read-only instance of NegRiskV2, bound to a specific deployed contract.
func NewNegRiskV2Caller(address common.Address, caller bind.ContractCaller) (*NegRiskV2Caller, error) {
	contract, err := bindNegRiskV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NegRiskV2Caller{contract: contract}, nil
}

// NewNegRiskV2Transactor creates a new write-only instance of NegRiskV2, bound to a specific deployed contract.
func NewNegRiskV2Transactor(address common.Address, transactor bind.ContractTransactor) (*NegRiskV2Transactor, error) {
	contract, err := bindNegRiskV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NegRiskV2Transactor{contract: contract}, nil
}

// NewNegRiskV2Filterer creates a new log filterer instance of NegRiskV2, bound to a specific deployed contract.
func NewNegRiskV2Filterer(address common.Address, filterer bind.ContractFilterer) (*NegRiskV2Filterer, error) {
	contract, err := bindNegRiskV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NegRiskV2Filterer{contract: contract}, nil
}

// bindNegRiskV2 binds a generic wrapper to an already deployed contract.
func bindNegRiskV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NegRiskV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NegRiskV2 *NegRiskV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NegRiskV2.Contract.NegRiskV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NegRiskV2 *NegRiskV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskV2.Contract.NegRiskV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NegRiskV2 *NegRiskV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NegRiskV2.Contract.NegRiskV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NegRiskV2 *NegRiskV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NegRiskV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NegRiskV2 *NegRiskV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NegRiskV2 *NegRiskV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NegRiskV2.Contract.contract.Transact(opts, method, params...)
}

// PARENTCOLLECTIONID is a free data retrieval call binding the contract method 0x7afda8b8.
//
// Solidity: function PARENT_COLLECTION_ID() view returns(bytes32)
func (_NegRiskV2 *NegRiskV2Caller) PARENTCOLLECTIONID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "PARENT_COLLECTION_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PARENTCOLLECTIONID is a free data retrieval call binding the contract method 0x7afda8b8.
//
// Solidity: function PARENT_COLLECTION_ID() view returns(bytes32)
func (_NegRiskV2 *NegRiskV2Session) PARENTCOLLECTIONID() ([32]byte, error) {
	return _NegRiskV2.Contract.PARENTCOLLECTIONID(&_NegRiskV2.CallOpts)
}

// PARENTCOLLECTIONID is a free data retrieval call binding the contract method 0x7afda8b8.
//
// Solidity: function PARENT_COLLECTION_ID() view returns(bytes32)
func (_NegRiskV2 *NegRiskV2CallerSession) PARENTCOLLECTIONID() ([32]byte, error) {
	return _NegRiskV2.Contract.PARENTCOLLECTIONID(&_NegRiskV2.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_NegRiskV2 *NegRiskV2Caller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "eip712Domain")

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
func (_NegRiskV2 *NegRiskV2Session) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _NegRiskV2.Contract.Eip712Domain(&_NegRiskV2.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_NegRiskV2 *NegRiskV2CallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _NegRiskV2.Contract.Eip712Domain(&_NegRiskV2.CallOpts)
}

// GetCollateral is a free data retrieval call binding the contract method 0x5c1548fb.
//
// Solidity: function getCollateral() view returns(address)
func (_NegRiskV2 *NegRiskV2Caller) GetCollateral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "getCollateral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCollateral is a free data retrieval call binding the contract method 0x5c1548fb.
//
// Solidity: function getCollateral() view returns(address)
func (_NegRiskV2 *NegRiskV2Session) GetCollateral() (common.Address, error) {
	return _NegRiskV2.Contract.GetCollateral(&_NegRiskV2.CallOpts)
}

// GetCollateral is a free data retrieval call binding the contract method 0x5c1548fb.
//
// Solidity: function getCollateral() view returns(address)
func (_NegRiskV2 *NegRiskV2CallerSession) GetCollateral() (common.Address, error) {
	return _NegRiskV2.Contract.GetCollateral(&_NegRiskV2.CallOpts)
}

// GetCtf is a free data retrieval call binding the contract method 0x3b521d78.
//
// Solidity: function getCtf() view returns(address)
func (_NegRiskV2 *NegRiskV2Caller) GetCtf(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "getCtf")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCtf is a free data retrieval call binding the contract method 0x3b521d78.
//
// Solidity: function getCtf() view returns(address)
func (_NegRiskV2 *NegRiskV2Session) GetCtf() (common.Address, error) {
	return _NegRiskV2.Contract.GetCtf(&_NegRiskV2.CallOpts)
}

// GetCtf is a free data retrieval call binding the contract method 0x3b521d78.
//
// Solidity: function getCtf() view returns(address)
func (_NegRiskV2 *NegRiskV2CallerSession) GetCtf() (common.Address, error) {
	return _NegRiskV2.Contract.GetCtf(&_NegRiskV2.CallOpts)
}

// GetCtfCollateral is a free data retrieval call binding the contract method 0x03cee3df.
//
// Solidity: function getCtfCollateral() view returns(address)
func (_NegRiskV2 *NegRiskV2Caller) GetCtfCollateral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "getCtfCollateral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCtfCollateral is a free data retrieval call binding the contract method 0x03cee3df.
//
// Solidity: function getCtfCollateral() view returns(address)
func (_NegRiskV2 *NegRiskV2Session) GetCtfCollateral() (common.Address, error) {
	return _NegRiskV2.Contract.GetCtfCollateral(&_NegRiskV2.CallOpts)
}

// GetCtfCollateral is a free data retrieval call binding the contract method 0x03cee3df.
//
// Solidity: function getCtfCollateral() view returns(address)
func (_NegRiskV2 *NegRiskV2CallerSession) GetCtfCollateral() (common.Address, error) {
	return _NegRiskV2.Contract.GetCtfCollateral(&_NegRiskV2.CallOpts)
}

// GetFeeReceiver is a free data retrieval call binding the contract method 0xe8a35392.
//
// Solidity: function getFeeReceiver() view returns(address)
func (_NegRiskV2 *NegRiskV2Caller) GetFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "getFeeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeReceiver is a free data retrieval call binding the contract method 0xe8a35392.
//
// Solidity: function getFeeReceiver() view returns(address)
func (_NegRiskV2 *NegRiskV2Session) GetFeeReceiver() (common.Address, error) {
	return _NegRiskV2.Contract.GetFeeReceiver(&_NegRiskV2.CallOpts)
}

// GetFeeReceiver is a free data retrieval call binding the contract method 0xe8a35392.
//
// Solidity: function getFeeReceiver() view returns(address)
func (_NegRiskV2 *NegRiskV2CallerSession) GetFeeReceiver() (common.Address, error) {
	return _NegRiskV2.Contract.GetFeeReceiver(&_NegRiskV2.CallOpts)
}

// GetMaxFeeRate is a free data retrieval call binding the contract method 0x4a2a11f5.
//
// Solidity: function getMaxFeeRate() view returns(uint256)
func (_NegRiskV2 *NegRiskV2Caller) GetMaxFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "getMaxFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxFeeRate is a free data retrieval call binding the contract method 0x4a2a11f5.
//
// Solidity: function getMaxFeeRate() view returns(uint256)
func (_NegRiskV2 *NegRiskV2Session) GetMaxFeeRate() (*big.Int, error) {
	return _NegRiskV2.Contract.GetMaxFeeRate(&_NegRiskV2.CallOpts)
}

// GetMaxFeeRate is a free data retrieval call binding the contract method 0x4a2a11f5.
//
// Solidity: function getMaxFeeRate() view returns(uint256)
func (_NegRiskV2 *NegRiskV2CallerSession) GetMaxFeeRate() (*big.Int, error) {
	return _NegRiskV2.Contract.GetMaxFeeRate(&_NegRiskV2.CallOpts)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns((bool,uint248))
func (_NegRiskV2 *NegRiskV2Caller) GetOrderStatus(opts *bind.CallOpts, orderHash [32]byte) (OrderStatus, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "getOrderStatus", orderHash)

	if err != nil {
		return *new(OrderStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(OrderStatus)).(*OrderStatus)

	return out0, err

}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns((bool,uint248))
func (_NegRiskV2 *NegRiskV2Session) GetOrderStatus(orderHash [32]byte) (OrderStatus, error) {
	return _NegRiskV2.Contract.GetOrderStatus(&_NegRiskV2.CallOpts, orderHash)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns((bool,uint248))
func (_NegRiskV2 *NegRiskV2CallerSession) GetOrderStatus(orderHash [32]byte) (OrderStatus, error) {
	return _NegRiskV2.Contract.GetOrderStatus(&_NegRiskV2.CallOpts, orderHash)
}

// GetOutcomeTokenFactory is a free data retrieval call binding the contract method 0x29cf67f2.
//
// Solidity: function getOutcomeTokenFactory() view returns(address)
func (_NegRiskV2 *NegRiskV2Caller) GetOutcomeTokenFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "getOutcomeTokenFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOutcomeTokenFactory is a free data retrieval call binding the contract method 0x29cf67f2.
//
// Solidity: function getOutcomeTokenFactory() view returns(address)
func (_NegRiskV2 *NegRiskV2Session) GetOutcomeTokenFactory() (common.Address, error) {
	return _NegRiskV2.Contract.GetOutcomeTokenFactory(&_NegRiskV2.CallOpts)
}

// GetOutcomeTokenFactory is a free data retrieval call binding the contract method 0x29cf67f2.
//
// Solidity: function getOutcomeTokenFactory() view returns(address)
func (_NegRiskV2 *NegRiskV2CallerSession) GetOutcomeTokenFactory() (common.Address, error) {
	return _NegRiskV2.Contract.GetOutcomeTokenFactory(&_NegRiskV2.CallOpts)
}

// GetProxyFactory is a free data retrieval call binding the contract method 0xb28c51c0.
//
// Solidity: function getProxyFactory() view returns(address)
func (_NegRiskV2 *NegRiskV2Caller) GetProxyFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "getProxyFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyFactory is a free data retrieval call binding the contract method 0xb28c51c0.
//
// Solidity: function getProxyFactory() view returns(address)
func (_NegRiskV2 *NegRiskV2Session) GetProxyFactory() (common.Address, error) {
	return _NegRiskV2.Contract.GetProxyFactory(&_NegRiskV2.CallOpts)
}

// GetProxyFactory is a free data retrieval call binding the contract method 0xb28c51c0.
//
// Solidity: function getProxyFactory() view returns(address)
func (_NegRiskV2 *NegRiskV2CallerSession) GetProxyFactory() (common.Address, error) {
	return _NegRiskV2.Contract.GetProxyFactory(&_NegRiskV2.CallOpts)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x90e4b720.
//
// Solidity: function getProxyImplementation() view returns(address)
func (_NegRiskV2 *NegRiskV2Caller) GetProxyImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "getProxyImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x90e4b720.
//
// Solidity: function getProxyImplementation() view returns(address)
func (_NegRiskV2 *NegRiskV2Session) GetProxyImplementation() (common.Address, error) {
	return _NegRiskV2.Contract.GetProxyImplementation(&_NegRiskV2.CallOpts)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x90e4b720.
//
// Solidity: function getProxyImplementation() view returns(address)
func (_NegRiskV2 *NegRiskV2CallerSession) GetProxyImplementation() (common.Address, error) {
	return _NegRiskV2.Contract.GetProxyImplementation(&_NegRiskV2.CallOpts)
}

// GetProxyWalletAddress is a free data retrieval call binding the contract method 0x58d8b6bb.
//
// Solidity: function getProxyWalletAddress(address _addr) view returns(address)
func (_NegRiskV2 *NegRiskV2Caller) GetProxyWalletAddress(opts *bind.CallOpts, _addr common.Address) (common.Address, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "getProxyWalletAddress", _addr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyWalletAddress is a free data retrieval call binding the contract method 0x58d8b6bb.
//
// Solidity: function getProxyWalletAddress(address _addr) view returns(address)
func (_NegRiskV2 *NegRiskV2Session) GetProxyWalletAddress(_addr common.Address) (common.Address, error) {
	return _NegRiskV2.Contract.GetProxyWalletAddress(&_NegRiskV2.CallOpts, _addr)
}

// GetProxyWalletAddress is a free data retrieval call binding the contract method 0x58d8b6bb.
//
// Solidity: function getProxyWalletAddress(address _addr) view returns(address)
func (_NegRiskV2 *NegRiskV2CallerSession) GetProxyWalletAddress(_addr common.Address) (common.Address, error) {
	return _NegRiskV2.Contract.GetProxyWalletAddress(&_NegRiskV2.CallOpts, _addr)
}

// GetSafeFactory is a free data retrieval call binding the contract method 0x75d7370a.
//
// Solidity: function getSafeFactory() view returns(address)
func (_NegRiskV2 *NegRiskV2Caller) GetSafeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "getSafeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSafeFactory is a free data retrieval call binding the contract method 0x75d7370a.
//
// Solidity: function getSafeFactory() view returns(address)
func (_NegRiskV2 *NegRiskV2Session) GetSafeFactory() (common.Address, error) {
	return _NegRiskV2.Contract.GetSafeFactory(&_NegRiskV2.CallOpts)
}

// GetSafeFactory is a free data retrieval call binding the contract method 0x75d7370a.
//
// Solidity: function getSafeFactory() view returns(address)
func (_NegRiskV2 *NegRiskV2CallerSession) GetSafeFactory() (common.Address, error) {
	return _NegRiskV2.Contract.GetSafeFactory(&_NegRiskV2.CallOpts)
}

// GetSafeImplementation is a free data retrieval call binding the contract method 0xe3b59000.
//
// Solidity: function getSafeImplementation() view returns(address)
func (_NegRiskV2 *NegRiskV2Caller) GetSafeImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "getSafeImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSafeImplementation is a free data retrieval call binding the contract method 0xe3b59000.
//
// Solidity: function getSafeImplementation() view returns(address)
func (_NegRiskV2 *NegRiskV2Session) GetSafeImplementation() (common.Address, error) {
	return _NegRiskV2.Contract.GetSafeImplementation(&_NegRiskV2.CallOpts)
}

// GetSafeImplementation is a free data retrieval call binding the contract method 0xe3b59000.
//
// Solidity: function getSafeImplementation() view returns(address)
func (_NegRiskV2 *NegRiskV2CallerSession) GetSafeImplementation() (common.Address, error) {
	return _NegRiskV2.Contract.GetSafeImplementation(&_NegRiskV2.CallOpts)
}

// GetSafeWalletAddress is a free data retrieval call binding the contract method 0x70bf48e5.
//
// Solidity: function getSafeWalletAddress(address _addr) view returns(address)
func (_NegRiskV2 *NegRiskV2Caller) GetSafeWalletAddress(opts *bind.CallOpts, _addr common.Address) (common.Address, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "getSafeWalletAddress", _addr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSafeWalletAddress is a free data retrieval call binding the contract method 0x70bf48e5.
//
// Solidity: function getSafeWalletAddress(address _addr) view returns(address)
func (_NegRiskV2 *NegRiskV2Session) GetSafeWalletAddress(_addr common.Address) (common.Address, error) {
	return _NegRiskV2.Contract.GetSafeWalletAddress(&_NegRiskV2.CallOpts, _addr)
}

// GetSafeWalletAddress is a free data retrieval call binding the contract method 0x70bf48e5.
//
// Solidity: function getSafeWalletAddress(address _addr) view returns(address)
func (_NegRiskV2 *NegRiskV2CallerSession) GetSafeWalletAddress(_addr common.Address) (common.Address, error) {
	return _NegRiskV2.Contract.GetSafeWalletAddress(&_NegRiskV2.CallOpts, _addr)
}

// HashOrder is a free data retrieval call binding the contract method 0x3d861a4d.
//
// Solidity: function hashOrder((uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) view returns(bytes32)
func (_NegRiskV2 *NegRiskV2Caller) HashOrder(opts *bind.CallOpts, order Order) ([32]byte, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "hashOrder", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOrder is a free data retrieval call binding the contract method 0x3d861a4d.
//
// Solidity: function hashOrder((uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) view returns(bytes32)
func (_NegRiskV2 *NegRiskV2Session) HashOrder(order Order) ([32]byte, error) {
	return _NegRiskV2.Contract.HashOrder(&_NegRiskV2.CallOpts, order)
}

// HashOrder is a free data retrieval call binding the contract method 0x3d861a4d.
//
// Solidity: function hashOrder((uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) view returns(bytes32)
func (_NegRiskV2 *NegRiskV2CallerSession) HashOrder(order Order) ([32]byte, error) {
	return _NegRiskV2.Contract.HashOrder(&_NegRiskV2.CallOpts, order)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _usr) view returns(bool)
func (_NegRiskV2 *NegRiskV2Caller) IsAdmin(opts *bind.CallOpts, _usr common.Address) (bool, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "isAdmin", _usr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _usr) view returns(bool)
func (_NegRiskV2 *NegRiskV2Session) IsAdmin(_usr common.Address) (bool, error) {
	return _NegRiskV2.Contract.IsAdmin(&_NegRiskV2.CallOpts, _usr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _usr) view returns(bool)
func (_NegRiskV2 *NegRiskV2CallerSession) IsAdmin(_usr common.Address) (bool, error) {
	return _NegRiskV2.Contract.IsAdmin(&_NegRiskV2.CallOpts, _usr)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address _usr) view returns(bool)
func (_NegRiskV2 *NegRiskV2Caller) IsOperator(opts *bind.CallOpts, _usr common.Address) (bool, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "isOperator", _usr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address _usr) view returns(bool)
func (_NegRiskV2 *NegRiskV2Session) IsOperator(_usr common.Address) (bool, error) {
	return _NegRiskV2.Contract.IsOperator(&_NegRiskV2.CallOpts, _usr)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address _usr) view returns(bool)
func (_NegRiskV2 *NegRiskV2CallerSession) IsOperator(_usr common.Address) (bool, error) {
	return _NegRiskV2.Contract.IsOperator(&_NegRiskV2.CallOpts, _usr)
}

// IsUserPaused is a free data retrieval call binding the contract method 0x28872101.
//
// Solidity: function isUserPaused(address user) view returns(bool)
func (_NegRiskV2 *NegRiskV2Caller) IsUserPaused(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "isUserPaused", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUserPaused is a free data retrieval call binding the contract method 0x28872101.
//
// Solidity: function isUserPaused(address user) view returns(bool)
func (_NegRiskV2 *NegRiskV2Session) IsUserPaused(user common.Address) (bool, error) {
	return _NegRiskV2.Contract.IsUserPaused(&_NegRiskV2.CallOpts, user)
}

// IsUserPaused is a free data retrieval call binding the contract method 0x28872101.
//
// Solidity: function isUserPaused(address user) view returns(bool)
func (_NegRiskV2 *NegRiskV2CallerSession) IsUserPaused(user common.Address) (bool, error) {
	return _NegRiskV2.Contract.IsUserPaused(&_NegRiskV2.CallOpts, user)
}

// OrderStatus is a free data retrieval call binding the contract method 0x2dff692d.
//
// Solidity: function orderStatus(bytes32 ) view returns(bool filled, uint248 remaining)
func (_NegRiskV2 *NegRiskV2Caller) OrderStatus(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Filled    bool
	Remaining *big.Int
}, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "orderStatus", arg0)

	outstruct := new(struct {
		Filled    bool
		Remaining *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Filled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Remaining = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OrderStatus is a free data retrieval call binding the contract method 0x2dff692d.
//
// Solidity: function orderStatus(bytes32 ) view returns(bool filled, uint248 remaining)
func (_NegRiskV2 *NegRiskV2Session) OrderStatus(arg0 [32]byte) (struct {
	Filled    bool
	Remaining *big.Int
}, error) {
	return _NegRiskV2.Contract.OrderStatus(&_NegRiskV2.CallOpts, arg0)
}

// OrderStatus is a free data retrieval call binding the contract method 0x2dff692d.
//
// Solidity: function orderStatus(bytes32 ) view returns(bool filled, uint248 remaining)
func (_NegRiskV2 *NegRiskV2CallerSession) OrderStatus(arg0 [32]byte) (struct {
	Filled    bool
	Remaining *big.Int
}, error) {
	return _NegRiskV2.Contract.OrderStatus(&_NegRiskV2.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NegRiskV2 *NegRiskV2Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NegRiskV2 *NegRiskV2Session) Paused() (bool, error) {
	return _NegRiskV2.Contract.Paused(&_NegRiskV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NegRiskV2 *NegRiskV2CallerSession) Paused() (bool, error) {
	return _NegRiskV2.Contract.Paused(&_NegRiskV2.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_NegRiskV2 *NegRiskV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_NegRiskV2 *NegRiskV2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NegRiskV2.Contract.SupportsInterface(&_NegRiskV2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_NegRiskV2 *NegRiskV2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NegRiskV2.Contract.SupportsInterface(&_NegRiskV2.CallOpts, interfaceId)
}

// UserPauseBlockInterval is a free data retrieval call binding the contract method 0xe3bf917a.
//
// Solidity: function userPauseBlockInterval() view returns(uint256)
func (_NegRiskV2 *NegRiskV2Caller) UserPauseBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "userPauseBlockInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserPauseBlockInterval is a free data retrieval call binding the contract method 0xe3bf917a.
//
// Solidity: function userPauseBlockInterval() view returns(uint256)
func (_NegRiskV2 *NegRiskV2Session) UserPauseBlockInterval() (*big.Int, error) {
	return _NegRiskV2.Contract.UserPauseBlockInterval(&_NegRiskV2.CallOpts)
}

// UserPauseBlockInterval is a free data retrieval call binding the contract method 0xe3bf917a.
//
// Solidity: function userPauseBlockInterval() view returns(uint256)
func (_NegRiskV2 *NegRiskV2CallerSession) UserPauseBlockInterval() (*big.Int, error) {
	return _NegRiskV2.Contract.UserPauseBlockInterval(&_NegRiskV2.CallOpts)
}

// UserPausedBlockAt is a free data retrieval call binding the contract method 0x234d81b9.
//
// Solidity: function userPausedBlockAt(address ) view returns(uint256)
func (_NegRiskV2 *NegRiskV2Caller) UserPausedBlockAt(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "userPausedBlockAt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserPausedBlockAt is a free data retrieval call binding the contract method 0x234d81b9.
//
// Solidity: function userPausedBlockAt(address ) view returns(uint256)
func (_NegRiskV2 *NegRiskV2Session) UserPausedBlockAt(arg0 common.Address) (*big.Int, error) {
	return _NegRiskV2.Contract.UserPausedBlockAt(&_NegRiskV2.CallOpts, arg0)
}

// UserPausedBlockAt is a free data retrieval call binding the contract method 0x234d81b9.
//
// Solidity: function userPausedBlockAt(address ) view returns(uint256)
func (_NegRiskV2 *NegRiskV2CallerSession) UserPausedBlockAt(arg0 common.Address) (*big.Int, error) {
	return _NegRiskV2.Contract.UserPausedBlockAt(&_NegRiskV2.CallOpts, arg0)
}

// ValidateFee is a free data retrieval call binding the contract method 0x0ffea65d.
//
// Solidity: function validateFee(uint256 fee, uint256 cashValue) view returns()
func (_NegRiskV2 *NegRiskV2Caller) ValidateFee(opts *bind.CallOpts, fee *big.Int, cashValue *big.Int) error {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "validateFee", fee, cashValue)

	if err != nil {
		return err
	}

	return err

}

// ValidateFee is a free data retrieval call binding the contract method 0x0ffea65d.
//
// Solidity: function validateFee(uint256 fee, uint256 cashValue) view returns()
func (_NegRiskV2 *NegRiskV2Session) ValidateFee(fee *big.Int, cashValue *big.Int) error {
	return _NegRiskV2.Contract.ValidateFee(&_NegRiskV2.CallOpts, fee, cashValue)
}

// ValidateFee is a free data retrieval call binding the contract method 0x0ffea65d.
//
// Solidity: function validateFee(uint256 fee, uint256 cashValue) view returns()
func (_NegRiskV2 *NegRiskV2CallerSession) ValidateFee(fee *big.Int, cashValue *big.Int) error {
	return _NegRiskV2.Contract.ValidateFee(&_NegRiskV2.CallOpts, fee, cashValue)
}

// ValidateOrder is a free data retrieval call binding the contract method 0x088170cb.
//
// Solidity: function validateOrder((uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) view returns()
func (_NegRiskV2 *NegRiskV2Caller) ValidateOrder(opts *bind.CallOpts, order Order) error {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "validateOrder", order)

	if err != nil {
		return err
	}

	return err

}

// ValidateOrder is a free data retrieval call binding the contract method 0x088170cb.
//
// Solidity: function validateOrder((uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) view returns()
func (_NegRiskV2 *NegRiskV2Session) ValidateOrder(order Order) error {
	return _NegRiskV2.Contract.ValidateOrder(&_NegRiskV2.CallOpts, order)
}

// ValidateOrder is a free data retrieval call binding the contract method 0x088170cb.
//
// Solidity: function validateOrder((uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) view returns()
func (_NegRiskV2 *NegRiskV2CallerSession) ValidateOrder(order Order) error {
	return _NegRiskV2.Contract.ValidateOrder(&_NegRiskV2.CallOpts, order)
}

// ValidateOrderSignature is a free data retrieval call binding the contract method 0xf5db3e4b.
//
// Solidity: function validateOrderSignature(bytes32 orderHash, (uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) view returns()
func (_NegRiskV2 *NegRiskV2Caller) ValidateOrderSignature(opts *bind.CallOpts, orderHash [32]byte, order Order) error {
	var out []interface{}
	err := _NegRiskV2.contract.Call(opts, &out, "validateOrderSignature", orderHash, order)

	if err != nil {
		return err
	}

	return err

}

// ValidateOrderSignature is a free data retrieval call binding the contract method 0xf5db3e4b.
//
// Solidity: function validateOrderSignature(bytes32 orderHash, (uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) view returns()
func (_NegRiskV2 *NegRiskV2Session) ValidateOrderSignature(orderHash [32]byte, order Order) error {
	return _NegRiskV2.Contract.ValidateOrderSignature(&_NegRiskV2.CallOpts, orderHash, order)
}

// ValidateOrderSignature is a free data retrieval call binding the contract method 0xf5db3e4b.
//
// Solidity: function validateOrderSignature(bytes32 orderHash, (uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) view returns()
func (_NegRiskV2 *NegRiskV2CallerSession) ValidateOrderSignature(orderHash [32]byte, order Order) error {
	return _NegRiskV2.Contract.ValidateOrderSignature(&_NegRiskV2.CallOpts, orderHash, order)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_NegRiskV2 *NegRiskV2Transactor) AddAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _NegRiskV2.contract.Transact(opts, "addAdmin", _admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_NegRiskV2 *NegRiskV2Session) AddAdmin(_admin common.Address) (*types.Transaction, error) {
	return _NegRiskV2.Contract.AddAdmin(&_NegRiskV2.TransactOpts, _admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_NegRiskV2 *NegRiskV2TransactorSession) AddAdmin(_admin common.Address) (*types.Transaction, error) {
	return _NegRiskV2.Contract.AddAdmin(&_NegRiskV2.TransactOpts, _admin)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address _operator) returns()
func (_NegRiskV2 *NegRiskV2Transactor) AddOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _NegRiskV2.contract.Transact(opts, "addOperator", _operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address _operator) returns()
func (_NegRiskV2 *NegRiskV2Session) AddOperator(_operator common.Address) (*types.Transaction, error) {
	return _NegRiskV2.Contract.AddOperator(&_NegRiskV2.TransactOpts, _operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address _operator) returns()
func (_NegRiskV2 *NegRiskV2TransactorSession) AddOperator(_operator common.Address) (*types.Transaction, error) {
	return _NegRiskV2.Contract.AddOperator(&_NegRiskV2.TransactOpts, _operator)
}

// InvalidatePreapprovedOrder is a paid mutator transaction binding the contract method 0x437f1994.
//
// Solidity: function invalidatePreapprovedOrder(bytes32 orderHash) returns()
func (_NegRiskV2 *NegRiskV2Transactor) InvalidatePreapprovedOrder(opts *bind.TransactOpts, orderHash [32]byte) (*types.Transaction, error) {
	return _NegRiskV2.contract.Transact(opts, "invalidatePreapprovedOrder", orderHash)
}

// InvalidatePreapprovedOrder is a paid mutator transaction binding the contract method 0x437f1994.
//
// Solidity: function invalidatePreapprovedOrder(bytes32 orderHash) returns()
func (_NegRiskV2 *NegRiskV2Session) InvalidatePreapprovedOrder(orderHash [32]byte) (*types.Transaction, error) {
	return _NegRiskV2.Contract.InvalidatePreapprovedOrder(&_NegRiskV2.TransactOpts, orderHash)
}

// InvalidatePreapprovedOrder is a paid mutator transaction binding the contract method 0x437f1994.
//
// Solidity: function invalidatePreapprovedOrder(bytes32 orderHash) returns()
func (_NegRiskV2 *NegRiskV2TransactorSession) InvalidatePreapprovedOrder(orderHash [32]byte) (*types.Transaction, error) {
	return _NegRiskV2.Contract.InvalidatePreapprovedOrder(&_NegRiskV2.TransactOpts, orderHash)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x3c2b4399.
//
// Solidity: function matchOrders(bytes32 conditionId, (uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) takerOrder, (uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes)[] makerOrders, uint256 takerFillAmount, uint256[] makerFillAmounts, uint256 takerFeeAmount, uint256[] makerFeeAmounts) returns()
func (_NegRiskV2 *NegRiskV2Transactor) MatchOrders(opts *bind.TransactOpts, conditionId [32]byte, takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, makerFillAmounts []*big.Int, takerFeeAmount *big.Int, makerFeeAmounts []*big.Int) (*types.Transaction, error) {
	return _NegRiskV2.contract.Transact(opts, "matchOrders", conditionId, takerOrder, makerOrders, takerFillAmount, makerFillAmounts, takerFeeAmount, makerFeeAmounts)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x3c2b4399.
//
// Solidity: function matchOrders(bytes32 conditionId, (uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) takerOrder, (uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes)[] makerOrders, uint256 takerFillAmount, uint256[] makerFillAmounts, uint256 takerFeeAmount, uint256[] makerFeeAmounts) returns()
func (_NegRiskV2 *NegRiskV2Session) MatchOrders(conditionId [32]byte, takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, makerFillAmounts []*big.Int, takerFeeAmount *big.Int, makerFeeAmounts []*big.Int) (*types.Transaction, error) {
	return _NegRiskV2.Contract.MatchOrders(&_NegRiskV2.TransactOpts, conditionId, takerOrder, makerOrders, takerFillAmount, makerFillAmounts, takerFeeAmount, makerFeeAmounts)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x3c2b4399.
//
// Solidity: function matchOrders(bytes32 conditionId, (uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) takerOrder, (uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes)[] makerOrders, uint256 takerFillAmount, uint256[] makerFillAmounts, uint256 takerFeeAmount, uint256[] makerFeeAmounts) returns()
func (_NegRiskV2 *NegRiskV2TransactorSession) MatchOrders(conditionId [32]byte, takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, makerFillAmounts []*big.Int, takerFeeAmount *big.Int, makerFeeAmounts []*big.Int) (*types.Transaction, error) {
	return _NegRiskV2.Contract.MatchOrders(&_NegRiskV2.TransactOpts, conditionId, takerOrder, makerOrders, takerFillAmount, makerFillAmounts, takerFeeAmount, makerFeeAmounts)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_NegRiskV2 *NegRiskV2Transactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskV2.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_NegRiskV2 *NegRiskV2Session) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskV2.Contract.OnERC1155BatchReceived(&_NegRiskV2.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_NegRiskV2 *NegRiskV2TransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskV2.Contract.OnERC1155BatchReceived(&_NegRiskV2.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_NegRiskV2 *NegRiskV2Transactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskV2.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_NegRiskV2 *NegRiskV2Session) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskV2.Contract.OnERC1155Received(&_NegRiskV2.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_NegRiskV2 *NegRiskV2TransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskV2.Contract.OnERC1155Received(&_NegRiskV2.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// PauseTrading is a paid mutator transaction binding the contract method 0x1031e36e.
//
// Solidity: function pauseTrading() returns()
func (_NegRiskV2 *NegRiskV2Transactor) PauseTrading(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskV2.contract.Transact(opts, "pauseTrading")
}

// PauseTrading is a paid mutator transaction binding the contract method 0x1031e36e.
//
// Solidity: function pauseTrading() returns()
func (_NegRiskV2 *NegRiskV2Session) PauseTrading() (*types.Transaction, error) {
	return _NegRiskV2.Contract.PauseTrading(&_NegRiskV2.TransactOpts)
}

// PauseTrading is a paid mutator transaction binding the contract method 0x1031e36e.
//
// Solidity: function pauseTrading() returns()
func (_NegRiskV2 *NegRiskV2TransactorSession) PauseTrading() (*types.Transaction, error) {
	return _NegRiskV2.Contract.PauseTrading(&_NegRiskV2.TransactOpts)
}

// PauseUser is a paid mutator transaction binding the contract method 0xc0c3132c.
//
// Solidity: function pauseUser() returns()
func (_NegRiskV2 *NegRiskV2Transactor) PauseUser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskV2.contract.Transact(opts, "pauseUser")
}

// PauseUser is a paid mutator transaction binding the contract method 0xc0c3132c.
//
// Solidity: function pauseUser() returns()
func (_NegRiskV2 *NegRiskV2Session) PauseUser() (*types.Transaction, error) {
	return _NegRiskV2.Contract.PauseUser(&_NegRiskV2.TransactOpts)
}

// PauseUser is a paid mutator transaction binding the contract method 0xc0c3132c.
//
// Solidity: function pauseUser() returns()
func (_NegRiskV2 *NegRiskV2TransactorSession) PauseUser() (*types.Transaction, error) {
	return _NegRiskV2.Contract.PauseUser(&_NegRiskV2.TransactOpts)
}

// PreapproveOrder is a paid mutator transaction binding the contract method 0xe3a5ced5.
//
// Solidity: function preapproveOrder((uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) returns()
func (_NegRiskV2 *NegRiskV2Transactor) PreapproveOrder(opts *bind.TransactOpts, order Order) (*types.Transaction, error) {
	return _NegRiskV2.contract.Transact(opts, "preapproveOrder", order)
}

// PreapproveOrder is a paid mutator transaction binding the contract method 0xe3a5ced5.
//
// Solidity: function preapproveOrder((uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) returns()
func (_NegRiskV2 *NegRiskV2Session) PreapproveOrder(order Order) (*types.Transaction, error) {
	return _NegRiskV2.Contract.PreapproveOrder(&_NegRiskV2.TransactOpts, order)
}

// PreapproveOrder is a paid mutator transaction binding the contract method 0xe3a5ced5.
//
// Solidity: function preapproveOrder((uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) returns()
func (_NegRiskV2 *NegRiskV2TransactorSession) PreapproveOrder(order Order) (*types.Transaction, error) {
	return _NegRiskV2.Contract.PreapproveOrder(&_NegRiskV2.TransactOpts, order)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_NegRiskV2 *NegRiskV2Transactor) RemoveAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _NegRiskV2.contract.Transact(opts, "removeAdmin", _admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_NegRiskV2 *NegRiskV2Session) RemoveAdmin(_admin common.Address) (*types.Transaction, error) {
	return _NegRiskV2.Contract.RemoveAdmin(&_NegRiskV2.TransactOpts, _admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_NegRiskV2 *NegRiskV2TransactorSession) RemoveAdmin(_admin common.Address) (*types.Transaction, error) {
	return _NegRiskV2.Contract.RemoveAdmin(&_NegRiskV2.TransactOpts, _admin)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address _operator) returns()
func (_NegRiskV2 *NegRiskV2Transactor) RemoveOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _NegRiskV2.contract.Transact(opts, "removeOperator", _operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address _operator) returns()
func (_NegRiskV2 *NegRiskV2Session) RemoveOperator(_operator common.Address) (*types.Transaction, error) {
	return _NegRiskV2.Contract.RemoveOperator(&_NegRiskV2.TransactOpts, _operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address _operator) returns()
func (_NegRiskV2 *NegRiskV2TransactorSession) RemoveOperator(_operator common.Address) (*types.Transaction, error) {
	return _NegRiskV2.Contract.RemoveOperator(&_NegRiskV2.TransactOpts, _operator)
}

// RenounceOperatorRole is a paid mutator transaction binding the contract method 0x3d6d3598.
//
// Solidity: function renounceOperatorRole() returns()
func (_NegRiskV2 *NegRiskV2Transactor) RenounceOperatorRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskV2.contract.Transact(opts, "renounceOperatorRole")
}

// RenounceOperatorRole is a paid mutator transaction binding the contract method 0x3d6d3598.
//
// Solidity: function renounceOperatorRole() returns()
func (_NegRiskV2 *NegRiskV2Session) RenounceOperatorRole() (*types.Transaction, error) {
	return _NegRiskV2.Contract.RenounceOperatorRole(&_NegRiskV2.TransactOpts)
}

// RenounceOperatorRole is a paid mutator transaction binding the contract method 0x3d6d3598.
//
// Solidity: function renounceOperatorRole() returns()
func (_NegRiskV2 *NegRiskV2TransactorSession) RenounceOperatorRole() (*types.Transaction, error) {
	return _NegRiskV2.Contract.RenounceOperatorRole(&_NegRiskV2.TransactOpts)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address receiver) returns()
func (_NegRiskV2 *NegRiskV2Transactor) SetFeeReceiver(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _NegRiskV2.contract.Transact(opts, "setFeeReceiver", receiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address receiver) returns()
func (_NegRiskV2 *NegRiskV2Session) SetFeeReceiver(receiver common.Address) (*types.Transaction, error) {
	return _NegRiskV2.Contract.SetFeeReceiver(&_NegRiskV2.TransactOpts, receiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address receiver) returns()
func (_NegRiskV2 *NegRiskV2TransactorSession) SetFeeReceiver(receiver common.Address) (*types.Transaction, error) {
	return _NegRiskV2.Contract.SetFeeReceiver(&_NegRiskV2.TransactOpts, receiver)
}

// SetMaxFeeRate is a paid mutator transaction binding the contract method 0x8cda96de.
//
// Solidity: function setMaxFeeRate(uint256 rate) returns()
func (_NegRiskV2 *NegRiskV2Transactor) SetMaxFeeRate(opts *bind.TransactOpts, rate *big.Int) (*types.Transaction, error) {
	return _NegRiskV2.contract.Transact(opts, "setMaxFeeRate", rate)
}

// SetMaxFeeRate is a paid mutator transaction binding the contract method 0x8cda96de.
//
// Solidity: function setMaxFeeRate(uint256 rate) returns()
func (_NegRiskV2 *NegRiskV2Session) SetMaxFeeRate(rate *big.Int) (*types.Transaction, error) {
	return _NegRiskV2.Contract.SetMaxFeeRate(&_NegRiskV2.TransactOpts, rate)
}

// SetMaxFeeRate is a paid mutator transaction binding the contract method 0x8cda96de.
//
// Solidity: function setMaxFeeRate(uint256 rate) returns()
func (_NegRiskV2 *NegRiskV2TransactorSession) SetMaxFeeRate(rate *big.Int) (*types.Transaction, error) {
	return _NegRiskV2.Contract.SetMaxFeeRate(&_NegRiskV2.TransactOpts, rate)
}

// SetUserPauseBlockInterval is a paid mutator transaction binding the contract method 0xcd4d8cb4.
//
// Solidity: function setUserPauseBlockInterval(uint256 _interval) returns()
func (_NegRiskV2 *NegRiskV2Transactor) SetUserPauseBlockInterval(opts *bind.TransactOpts, _interval *big.Int) (*types.Transaction, error) {
	return _NegRiskV2.contract.Transact(opts, "setUserPauseBlockInterval", _interval)
}

// SetUserPauseBlockInterval is a paid mutator transaction binding the contract method 0xcd4d8cb4.
//
// Solidity: function setUserPauseBlockInterval(uint256 _interval) returns()
func (_NegRiskV2 *NegRiskV2Session) SetUserPauseBlockInterval(_interval *big.Int) (*types.Transaction, error) {
	return _NegRiskV2.Contract.SetUserPauseBlockInterval(&_NegRiskV2.TransactOpts, _interval)
}

// SetUserPauseBlockInterval is a paid mutator transaction binding the contract method 0xcd4d8cb4.
//
// Solidity: function setUserPauseBlockInterval(uint256 _interval) returns()
func (_NegRiskV2 *NegRiskV2TransactorSession) SetUserPauseBlockInterval(_interval *big.Int) (*types.Transaction, error) {
	return _NegRiskV2.Contract.SetUserPauseBlockInterval(&_NegRiskV2.TransactOpts, _interval)
}

// UnpauseTrading is a paid mutator transaction binding the contract method 0x456068d2.
//
// Solidity: function unpauseTrading() returns()
func (_NegRiskV2 *NegRiskV2Transactor) UnpauseTrading(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskV2.contract.Transact(opts, "unpauseTrading")
}

// UnpauseTrading is a paid mutator transaction binding the contract method 0x456068d2.
//
// Solidity: function unpauseTrading() returns()
func (_NegRiskV2 *NegRiskV2Session) UnpauseTrading() (*types.Transaction, error) {
	return _NegRiskV2.Contract.UnpauseTrading(&_NegRiskV2.TransactOpts)
}

// UnpauseTrading is a paid mutator transaction binding the contract method 0x456068d2.
//
// Solidity: function unpauseTrading() returns()
func (_NegRiskV2 *NegRiskV2TransactorSession) UnpauseTrading() (*types.Transaction, error) {
	return _NegRiskV2.Contract.UnpauseTrading(&_NegRiskV2.TransactOpts)
}

// UnpauseUser is a paid mutator transaction binding the contract method 0x4cce30c9.
//
// Solidity: function unpauseUser() returns()
func (_NegRiskV2 *NegRiskV2Transactor) UnpauseUser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskV2.contract.Transact(opts, "unpauseUser")
}

// UnpauseUser is a paid mutator transaction binding the contract method 0x4cce30c9.
//
// Solidity: function unpauseUser() returns()
func (_NegRiskV2 *NegRiskV2Session) UnpauseUser() (*types.Transaction, error) {
	return _NegRiskV2.Contract.UnpauseUser(&_NegRiskV2.TransactOpts)
}

// UnpauseUser is a paid mutator transaction binding the contract method 0x4cce30c9.
//
// Solidity: function unpauseUser() returns()
func (_NegRiskV2 *NegRiskV2TransactorSession) UnpauseUser() (*types.Transaction, error) {
	return _NegRiskV2.Contract.UnpauseUser(&_NegRiskV2.TransactOpts)
}

// NegRiskV2FeeChargedIterator is returned from FilterFeeCharged and is used to iterate over the raw logs and unpacked data for FeeCharged events raised by the NegRiskV2 contract.
type NegRiskV2FeeChargedIterator struct {
	Event *NegRiskV2FeeCharged // Event containing the contract specifics and raw log

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
func (it *NegRiskV2FeeChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskV2FeeCharged)
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
		it.Event = new(NegRiskV2FeeCharged)
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
func (it *NegRiskV2FeeChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskV2FeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskV2FeeCharged represents a FeeCharged event raised by the NegRiskV2 contract.
type NegRiskV2FeeCharged struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeeCharged is a free log retrieval operation binding the contract event 0x55bb3cade9d43b798a4fe5ffdd05024b2d7870df53920673bfc7e68047cd0ab1.
//
// Solidity: event FeeCharged(address indexed receiver, uint256 amount)
func (_NegRiskV2 *NegRiskV2Filterer) FilterFeeCharged(opts *bind.FilterOpts, receiver []common.Address) (*NegRiskV2FeeChargedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _NegRiskV2.contract.FilterLogs(opts, "FeeCharged", receiverRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskV2FeeChargedIterator{contract: _NegRiskV2.contract, event: "FeeCharged", logs: logs, sub: sub}, nil
}

// WatchFeeCharged is a free log subscription operation binding the contract event 0x55bb3cade9d43b798a4fe5ffdd05024b2d7870df53920673bfc7e68047cd0ab1.
//
// Solidity: event FeeCharged(address indexed receiver, uint256 amount)
func (_NegRiskV2 *NegRiskV2Filterer) WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *NegRiskV2FeeCharged, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _NegRiskV2.contract.WatchLogs(opts, "FeeCharged", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskV2FeeCharged)
				if err := _NegRiskV2.contract.UnpackLog(event, "FeeCharged", log); err != nil {
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

// ParseFeeCharged is a log parse operation binding the contract event 0x55bb3cade9d43b798a4fe5ffdd05024b2d7870df53920673bfc7e68047cd0ab1.
//
// Solidity: event FeeCharged(address indexed receiver, uint256 amount)
func (_NegRiskV2 *NegRiskV2Filterer) ParseFeeCharged(log types.Log) (*NegRiskV2FeeCharged, error) {
	event := new(NegRiskV2FeeCharged)
	if err := _NegRiskV2.contract.UnpackLog(event, "FeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskV2FeeReceiverUpdatedIterator is returned from FilterFeeReceiverUpdated and is used to iterate over the raw logs and unpacked data for FeeReceiverUpdated events raised by the NegRiskV2 contract.
type NegRiskV2FeeReceiverUpdatedIterator struct {
	Event *NegRiskV2FeeReceiverUpdated // Event containing the contract specifics and raw log

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
func (it *NegRiskV2FeeReceiverUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskV2FeeReceiverUpdated)
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
		it.Event = new(NegRiskV2FeeReceiverUpdated)
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
func (it *NegRiskV2FeeReceiverUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskV2FeeReceiverUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskV2FeeReceiverUpdated represents a FeeReceiverUpdated event raised by the NegRiskV2 contract.
type NegRiskV2FeeReceiverUpdated struct {
	FeeReceiver common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeReceiverUpdated is a free log retrieval operation binding the contract event 0x27aae5db36d94179909d019ae0b1ac7c16d96d953148f63c0f6a0a9c8ead79ee.
//
// Solidity: event FeeReceiverUpdated(address indexed feeReceiver)
func (_NegRiskV2 *NegRiskV2Filterer) FilterFeeReceiverUpdated(opts *bind.FilterOpts, feeReceiver []common.Address) (*NegRiskV2FeeReceiverUpdatedIterator, error) {

	var feeReceiverRule []interface{}
	for _, feeReceiverItem := range feeReceiver {
		feeReceiverRule = append(feeReceiverRule, feeReceiverItem)
	}

	logs, sub, err := _NegRiskV2.contract.FilterLogs(opts, "FeeReceiverUpdated", feeReceiverRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskV2FeeReceiverUpdatedIterator{contract: _NegRiskV2.contract, event: "FeeReceiverUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeReceiverUpdated is a free log subscription operation binding the contract event 0x27aae5db36d94179909d019ae0b1ac7c16d96d953148f63c0f6a0a9c8ead79ee.
//
// Solidity: event FeeReceiverUpdated(address indexed feeReceiver)
func (_NegRiskV2 *NegRiskV2Filterer) WatchFeeReceiverUpdated(opts *bind.WatchOpts, sink chan<- *NegRiskV2FeeReceiverUpdated, feeReceiver []common.Address) (event.Subscription, error) {

	var feeReceiverRule []interface{}
	for _, feeReceiverItem := range feeReceiver {
		feeReceiverRule = append(feeReceiverRule, feeReceiverItem)
	}

	logs, sub, err := _NegRiskV2.contract.WatchLogs(opts, "FeeReceiverUpdated", feeReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskV2FeeReceiverUpdated)
				if err := _NegRiskV2.contract.UnpackLog(event, "FeeReceiverUpdated", log); err != nil {
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

// ParseFeeReceiverUpdated is a log parse operation binding the contract event 0x27aae5db36d94179909d019ae0b1ac7c16d96d953148f63c0f6a0a9c8ead79ee.
//
// Solidity: event FeeReceiverUpdated(address indexed feeReceiver)
func (_NegRiskV2 *NegRiskV2Filterer) ParseFeeReceiverUpdated(log types.Log) (*NegRiskV2FeeReceiverUpdated, error) {
	event := new(NegRiskV2FeeReceiverUpdated)
	if err := _NegRiskV2.contract.UnpackLog(event, "FeeReceiverUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskV2MaxFeeRateUpdatedIterator is returned from FilterMaxFeeRateUpdated and is used to iterate over the raw logs and unpacked data for MaxFeeRateUpdated events raised by the NegRiskV2 contract.
type NegRiskV2MaxFeeRateUpdatedIterator struct {
	Event *NegRiskV2MaxFeeRateUpdated // Event containing the contract specifics and raw log

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
func (it *NegRiskV2MaxFeeRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskV2MaxFeeRateUpdated)
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
		it.Event = new(NegRiskV2MaxFeeRateUpdated)
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
func (it *NegRiskV2MaxFeeRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskV2MaxFeeRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskV2MaxFeeRateUpdated represents a MaxFeeRateUpdated event raised by the NegRiskV2 contract.
type NegRiskV2MaxFeeRateUpdated struct {
	MaxFeeRate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMaxFeeRateUpdated is a free log retrieval operation binding the contract event 0xe380d7c3967dd06cc7c01db8b17332a1d806fd18f63206dcbd12aaef455c7ff2.
//
// Solidity: event MaxFeeRateUpdated(uint256 maxFeeRate)
func (_NegRiskV2 *NegRiskV2Filterer) FilterMaxFeeRateUpdated(opts *bind.FilterOpts) (*NegRiskV2MaxFeeRateUpdatedIterator, error) {

	logs, sub, err := _NegRiskV2.contract.FilterLogs(opts, "MaxFeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return &NegRiskV2MaxFeeRateUpdatedIterator{contract: _NegRiskV2.contract, event: "MaxFeeRateUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxFeeRateUpdated is a free log subscription operation binding the contract event 0xe380d7c3967dd06cc7c01db8b17332a1d806fd18f63206dcbd12aaef455c7ff2.
//
// Solidity: event MaxFeeRateUpdated(uint256 maxFeeRate)
func (_NegRiskV2 *NegRiskV2Filterer) WatchMaxFeeRateUpdated(opts *bind.WatchOpts, sink chan<- *NegRiskV2MaxFeeRateUpdated) (event.Subscription, error) {

	logs, sub, err := _NegRiskV2.contract.WatchLogs(opts, "MaxFeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskV2MaxFeeRateUpdated)
				if err := _NegRiskV2.contract.UnpackLog(event, "MaxFeeRateUpdated", log); err != nil {
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

// ParseMaxFeeRateUpdated is a log parse operation binding the contract event 0xe380d7c3967dd06cc7c01db8b17332a1d806fd18f63206dcbd12aaef455c7ff2.
//
// Solidity: event MaxFeeRateUpdated(uint256 maxFeeRate)
func (_NegRiskV2 *NegRiskV2Filterer) ParseMaxFeeRateUpdated(log types.Log) (*NegRiskV2MaxFeeRateUpdated, error) {
	event := new(NegRiskV2MaxFeeRateUpdated)
	if err := _NegRiskV2.contract.UnpackLog(event, "MaxFeeRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskV2NewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the NegRiskV2 contract.
type NegRiskV2NewAdminIterator struct {
	Event *NegRiskV2NewAdmin // Event containing the contract specifics and raw log

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
func (it *NegRiskV2NewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskV2NewAdmin)
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
		it.Event = new(NegRiskV2NewAdmin)
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
func (it *NegRiskV2NewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskV2NewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskV2NewAdmin represents a NewAdmin event raised by the NegRiskV2 contract.
type NegRiskV2NewAdmin struct {
	NewAdminAddress common.Address
	Admin           common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed newAdminAddress, address indexed admin)
func (_NegRiskV2 *NegRiskV2Filterer) FilterNewAdmin(opts *bind.FilterOpts, newAdminAddress []common.Address, admin []common.Address) (*NegRiskV2NewAdminIterator, error) {

	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskV2.contract.FilterLogs(opts, "NewAdmin", newAdminAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskV2NewAdminIterator{contract: _NegRiskV2.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed newAdminAddress, address indexed admin)
func (_NegRiskV2 *NegRiskV2Filterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *NegRiskV2NewAdmin, newAdminAddress []common.Address, admin []common.Address) (event.Subscription, error) {

	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskV2.contract.WatchLogs(opts, "NewAdmin", newAdminAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskV2NewAdmin)
				if err := _NegRiskV2.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed newAdminAddress, address indexed admin)
func (_NegRiskV2 *NegRiskV2Filterer) ParseNewAdmin(log types.Log) (*NegRiskV2NewAdmin, error) {
	event := new(NegRiskV2NewAdmin)
	if err := _NegRiskV2.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskV2NewOperatorIterator is returned from FilterNewOperator and is used to iterate over the raw logs and unpacked data for NewOperator events raised by the NegRiskV2 contract.
type NegRiskV2NewOperatorIterator struct {
	Event *NegRiskV2NewOperator // Event containing the contract specifics and raw log

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
func (it *NegRiskV2NewOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskV2NewOperator)
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
		it.Event = new(NegRiskV2NewOperator)
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
func (it *NegRiskV2NewOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskV2NewOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskV2NewOperator represents a NewOperator event raised by the NegRiskV2 contract.
type NegRiskV2NewOperator struct {
	NewOperatorAddress common.Address
	Admin              common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewOperator is a free log retrieval operation binding the contract event 0xf1e04d73c4304b5ff164f9d10c7473e2a1593b740674a6107975e2a7001c1e5c.
//
// Solidity: event NewOperator(address indexed newOperatorAddress, address indexed admin)
func (_NegRiskV2 *NegRiskV2Filterer) FilterNewOperator(opts *bind.FilterOpts, newOperatorAddress []common.Address, admin []common.Address) (*NegRiskV2NewOperatorIterator, error) {

	var newOperatorAddressRule []interface{}
	for _, newOperatorAddressItem := range newOperatorAddress {
		newOperatorAddressRule = append(newOperatorAddressRule, newOperatorAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskV2.contract.FilterLogs(opts, "NewOperator", newOperatorAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskV2NewOperatorIterator{contract: _NegRiskV2.contract, event: "NewOperator", logs: logs, sub: sub}, nil
}

// WatchNewOperator is a free log subscription operation binding the contract event 0xf1e04d73c4304b5ff164f9d10c7473e2a1593b740674a6107975e2a7001c1e5c.
//
// Solidity: event NewOperator(address indexed newOperatorAddress, address indexed admin)
func (_NegRiskV2 *NegRiskV2Filterer) WatchNewOperator(opts *bind.WatchOpts, sink chan<- *NegRiskV2NewOperator, newOperatorAddress []common.Address, admin []common.Address) (event.Subscription, error) {

	var newOperatorAddressRule []interface{}
	for _, newOperatorAddressItem := range newOperatorAddress {
		newOperatorAddressRule = append(newOperatorAddressRule, newOperatorAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskV2.contract.WatchLogs(opts, "NewOperator", newOperatorAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskV2NewOperator)
				if err := _NegRiskV2.contract.UnpackLog(event, "NewOperator", log); err != nil {
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

// ParseNewOperator is a log parse operation binding the contract event 0xf1e04d73c4304b5ff164f9d10c7473e2a1593b740674a6107975e2a7001c1e5c.
//
// Solidity: event NewOperator(address indexed newOperatorAddress, address indexed admin)
func (_NegRiskV2 *NegRiskV2Filterer) ParseNewOperator(log types.Log) (*NegRiskV2NewOperator, error) {
	event := new(NegRiskV2NewOperator)
	if err := _NegRiskV2.contract.UnpackLog(event, "NewOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskV2OrderFilledIterator is returned from FilterOrderFilled and is used to iterate over the raw logs and unpacked data for OrderFilled events raised by the NegRiskV2 contract.
type NegRiskV2OrderFilledIterator struct {
	Event *NegRiskV2OrderFilled // Event containing the contract specifics and raw log

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
func (it *NegRiskV2OrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskV2OrderFilled)
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
		it.Event = new(NegRiskV2OrderFilled)
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
func (it *NegRiskV2OrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskV2OrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskV2OrderFilled represents a OrderFilled event raised by the NegRiskV2 contract.
type NegRiskV2OrderFilled struct {
	OrderHash         [32]byte
	Maker             common.Address
	Taker             common.Address
	Side              uint8
	TokenId           *big.Int
	MakerAmountFilled *big.Int
	TakerAmountFilled *big.Int
	Fee               *big.Int
	Builder           [32]byte
	Metadata          [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOrderFilled is a free log retrieval operation binding the contract event 0xd543adfd945773f1a62f74f0ee55a5e3b9b1a28262980ba90b1a89f2ea84d8ee.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, address indexed maker, address indexed taker, uint8 side, uint256 tokenId, uint256 makerAmountFilled, uint256 takerAmountFilled, uint256 fee, bytes32 builder, bytes32 metadata)
func (_NegRiskV2 *NegRiskV2Filterer) FilterOrderFilled(opts *bind.FilterOpts, orderHash [][32]byte, maker []common.Address, taker []common.Address) (*NegRiskV2OrderFilledIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _NegRiskV2.contract.FilterLogs(opts, "OrderFilled", orderHashRule, makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskV2OrderFilledIterator{contract: _NegRiskV2.contract, event: "OrderFilled", logs: logs, sub: sub}, nil
}

// WatchOrderFilled is a free log subscription operation binding the contract event 0xd543adfd945773f1a62f74f0ee55a5e3b9b1a28262980ba90b1a89f2ea84d8ee.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, address indexed maker, address indexed taker, uint8 side, uint256 tokenId, uint256 makerAmountFilled, uint256 takerAmountFilled, uint256 fee, bytes32 builder, bytes32 metadata)
func (_NegRiskV2 *NegRiskV2Filterer) WatchOrderFilled(opts *bind.WatchOpts, sink chan<- *NegRiskV2OrderFilled, orderHash [][32]byte, maker []common.Address, taker []common.Address) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _NegRiskV2.contract.WatchLogs(opts, "OrderFilled", orderHashRule, makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskV2OrderFilled)
				if err := _NegRiskV2.contract.UnpackLog(event, "OrderFilled", log); err != nil {
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

// ParseOrderFilled is a log parse operation binding the contract event 0xd543adfd945773f1a62f74f0ee55a5e3b9b1a28262980ba90b1a89f2ea84d8ee.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, address indexed maker, address indexed taker, uint8 side, uint256 tokenId, uint256 makerAmountFilled, uint256 takerAmountFilled, uint256 fee, bytes32 builder, bytes32 metadata)
func (_NegRiskV2 *NegRiskV2Filterer) ParseOrderFilled(log types.Log) (*NegRiskV2OrderFilled, error) {
	event := new(NegRiskV2OrderFilled)
	if err := _NegRiskV2.contract.UnpackLog(event, "OrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskV2OrderPreapprovalInvalidatedIterator is returned from FilterOrderPreapprovalInvalidated and is used to iterate over the raw logs and unpacked data for OrderPreapprovalInvalidated events raised by the NegRiskV2 contract.
type NegRiskV2OrderPreapprovalInvalidatedIterator struct {
	Event *NegRiskV2OrderPreapprovalInvalidated // Event containing the contract specifics and raw log

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
func (it *NegRiskV2OrderPreapprovalInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskV2OrderPreapprovalInvalidated)
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
		it.Event = new(NegRiskV2OrderPreapprovalInvalidated)
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
func (it *NegRiskV2OrderPreapprovalInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskV2OrderPreapprovalInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskV2OrderPreapprovalInvalidated represents a OrderPreapprovalInvalidated event raised by the NegRiskV2 contract.
type NegRiskV2OrderPreapprovalInvalidated struct {
	OrderHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderPreapprovalInvalidated is a free log retrieval operation binding the contract event 0xb766aa470f20b094f26a9a14ea5bf63a60af51703c15776e2e739b6a0428adf6.
//
// Solidity: event OrderPreapprovalInvalidated(bytes32 indexed orderHash)
func (_NegRiskV2 *NegRiskV2Filterer) FilterOrderPreapprovalInvalidated(opts *bind.FilterOpts, orderHash [][32]byte) (*NegRiskV2OrderPreapprovalInvalidatedIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _NegRiskV2.contract.FilterLogs(opts, "OrderPreapprovalInvalidated", orderHashRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskV2OrderPreapprovalInvalidatedIterator{contract: _NegRiskV2.contract, event: "OrderPreapprovalInvalidated", logs: logs, sub: sub}, nil
}

// WatchOrderPreapprovalInvalidated is a free log subscription operation binding the contract event 0xb766aa470f20b094f26a9a14ea5bf63a60af51703c15776e2e739b6a0428adf6.
//
// Solidity: event OrderPreapprovalInvalidated(bytes32 indexed orderHash)
func (_NegRiskV2 *NegRiskV2Filterer) WatchOrderPreapprovalInvalidated(opts *bind.WatchOpts, sink chan<- *NegRiskV2OrderPreapprovalInvalidated, orderHash [][32]byte) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _NegRiskV2.contract.WatchLogs(opts, "OrderPreapprovalInvalidated", orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskV2OrderPreapprovalInvalidated)
				if err := _NegRiskV2.contract.UnpackLog(event, "OrderPreapprovalInvalidated", log); err != nil {
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

// ParseOrderPreapprovalInvalidated is a log parse operation binding the contract event 0xb766aa470f20b094f26a9a14ea5bf63a60af51703c15776e2e739b6a0428adf6.
//
// Solidity: event OrderPreapprovalInvalidated(bytes32 indexed orderHash)
func (_NegRiskV2 *NegRiskV2Filterer) ParseOrderPreapprovalInvalidated(log types.Log) (*NegRiskV2OrderPreapprovalInvalidated, error) {
	event := new(NegRiskV2OrderPreapprovalInvalidated)
	if err := _NegRiskV2.contract.UnpackLog(event, "OrderPreapprovalInvalidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskV2OrderPreapprovedIterator is returned from FilterOrderPreapproved and is used to iterate over the raw logs and unpacked data for OrderPreapproved events raised by the NegRiskV2 contract.
type NegRiskV2OrderPreapprovedIterator struct {
	Event *NegRiskV2OrderPreapproved // Event containing the contract specifics and raw log

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
func (it *NegRiskV2OrderPreapprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskV2OrderPreapproved)
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
		it.Event = new(NegRiskV2OrderPreapproved)
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
func (it *NegRiskV2OrderPreapprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskV2OrderPreapprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskV2OrderPreapproved represents a OrderPreapproved event raised by the NegRiskV2 contract.
type NegRiskV2OrderPreapproved struct {
	OrderHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderPreapproved is a free log retrieval operation binding the contract event 0xe92c22722d9c284034b6c9f5aaec018edb3e593c0e084900b6b9d390a1182a0b.
//
// Solidity: event OrderPreapproved(bytes32 indexed orderHash)
func (_NegRiskV2 *NegRiskV2Filterer) FilterOrderPreapproved(opts *bind.FilterOpts, orderHash [][32]byte) (*NegRiskV2OrderPreapprovedIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _NegRiskV2.contract.FilterLogs(opts, "OrderPreapproved", orderHashRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskV2OrderPreapprovedIterator{contract: _NegRiskV2.contract, event: "OrderPreapproved", logs: logs, sub: sub}, nil
}

// WatchOrderPreapproved is a free log subscription operation binding the contract event 0xe92c22722d9c284034b6c9f5aaec018edb3e593c0e084900b6b9d390a1182a0b.
//
// Solidity: event OrderPreapproved(bytes32 indexed orderHash)
func (_NegRiskV2 *NegRiskV2Filterer) WatchOrderPreapproved(opts *bind.WatchOpts, sink chan<- *NegRiskV2OrderPreapproved, orderHash [][32]byte) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _NegRiskV2.contract.WatchLogs(opts, "OrderPreapproved", orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskV2OrderPreapproved)
				if err := _NegRiskV2.contract.UnpackLog(event, "OrderPreapproved", log); err != nil {
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

// ParseOrderPreapproved is a log parse operation binding the contract event 0xe92c22722d9c284034b6c9f5aaec018edb3e593c0e084900b6b9d390a1182a0b.
//
// Solidity: event OrderPreapproved(bytes32 indexed orderHash)
func (_NegRiskV2 *NegRiskV2Filterer) ParseOrderPreapproved(log types.Log) (*NegRiskV2OrderPreapproved, error) {
	event := new(NegRiskV2OrderPreapproved)
	if err := _NegRiskV2.contract.UnpackLog(event, "OrderPreapproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskV2OrdersMatchedIterator is returned from FilterOrdersMatched and is used to iterate over the raw logs and unpacked data for OrdersMatched events raised by the NegRiskV2 contract.
type NegRiskV2OrdersMatchedIterator struct {
	Event *NegRiskV2OrdersMatched // Event containing the contract specifics and raw log

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
func (it *NegRiskV2OrdersMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskV2OrdersMatched)
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
		it.Event = new(NegRiskV2OrdersMatched)
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
func (it *NegRiskV2OrdersMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskV2OrdersMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskV2OrdersMatched represents a OrdersMatched event raised by the NegRiskV2 contract.
type NegRiskV2OrdersMatched struct {
	TakerOrderHash    [32]byte
	TakerOrderMaker   common.Address
	Side              uint8
	TokenId           *big.Int
	MakerAmountFilled *big.Int
	TakerAmountFilled *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOrdersMatched is a free log retrieval operation binding the contract event 0x174b3811690657c217184f89418266767c87e4805d09680c39fc9c031c0cab7c.
//
// Solidity: event OrdersMatched(bytes32 indexed takerOrderHash, address indexed takerOrderMaker, uint8 side, uint256 tokenId, uint256 makerAmountFilled, uint256 takerAmountFilled)
func (_NegRiskV2 *NegRiskV2Filterer) FilterOrdersMatched(opts *bind.FilterOpts, takerOrderHash [][32]byte, takerOrderMaker []common.Address) (*NegRiskV2OrdersMatchedIterator, error) {

	var takerOrderHashRule []interface{}
	for _, takerOrderHashItem := range takerOrderHash {
		takerOrderHashRule = append(takerOrderHashRule, takerOrderHashItem)
	}
	var takerOrderMakerRule []interface{}
	for _, takerOrderMakerItem := range takerOrderMaker {
		takerOrderMakerRule = append(takerOrderMakerRule, takerOrderMakerItem)
	}

	logs, sub, err := _NegRiskV2.contract.FilterLogs(opts, "OrdersMatched", takerOrderHashRule, takerOrderMakerRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskV2OrdersMatchedIterator{contract: _NegRiskV2.contract, event: "OrdersMatched", logs: logs, sub: sub}, nil
}

// WatchOrdersMatched is a free log subscription operation binding the contract event 0x174b3811690657c217184f89418266767c87e4805d09680c39fc9c031c0cab7c.
//
// Solidity: event OrdersMatched(bytes32 indexed takerOrderHash, address indexed takerOrderMaker, uint8 side, uint256 tokenId, uint256 makerAmountFilled, uint256 takerAmountFilled)
func (_NegRiskV2 *NegRiskV2Filterer) WatchOrdersMatched(opts *bind.WatchOpts, sink chan<- *NegRiskV2OrdersMatched, takerOrderHash [][32]byte, takerOrderMaker []common.Address) (event.Subscription, error) {

	var takerOrderHashRule []interface{}
	for _, takerOrderHashItem := range takerOrderHash {
		takerOrderHashRule = append(takerOrderHashRule, takerOrderHashItem)
	}
	var takerOrderMakerRule []interface{}
	for _, takerOrderMakerItem := range takerOrderMaker {
		takerOrderMakerRule = append(takerOrderMakerRule, takerOrderMakerItem)
	}

	logs, sub, err := _NegRiskV2.contract.WatchLogs(opts, "OrdersMatched", takerOrderHashRule, takerOrderMakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskV2OrdersMatched)
				if err := _NegRiskV2.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
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

// ParseOrdersMatched is a log parse operation binding the contract event 0x174b3811690657c217184f89418266767c87e4805d09680c39fc9c031c0cab7c.
//
// Solidity: event OrdersMatched(bytes32 indexed takerOrderHash, address indexed takerOrderMaker, uint8 side, uint256 tokenId, uint256 makerAmountFilled, uint256 takerAmountFilled)
func (_NegRiskV2 *NegRiskV2Filterer) ParseOrdersMatched(log types.Log) (*NegRiskV2OrdersMatched, error) {
	event := new(NegRiskV2OrdersMatched)
	if err := _NegRiskV2.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskV2RemovedAdminIterator is returned from FilterRemovedAdmin and is used to iterate over the raw logs and unpacked data for RemovedAdmin events raised by the NegRiskV2 contract.
type NegRiskV2RemovedAdminIterator struct {
	Event *NegRiskV2RemovedAdmin // Event containing the contract specifics and raw log

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
func (it *NegRiskV2RemovedAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskV2RemovedAdmin)
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
		it.Event = new(NegRiskV2RemovedAdmin)
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
func (it *NegRiskV2RemovedAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskV2RemovedAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskV2RemovedAdmin represents a RemovedAdmin event raised by the NegRiskV2 contract.
type NegRiskV2RemovedAdmin struct {
	RemovedAdmin common.Address
	Admin        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemovedAdmin is a free log retrieval operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed removedAdmin, address indexed admin)
func (_NegRiskV2 *NegRiskV2Filterer) FilterRemovedAdmin(opts *bind.FilterOpts, removedAdmin []common.Address, admin []common.Address) (*NegRiskV2RemovedAdminIterator, error) {

	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskV2.contract.FilterLogs(opts, "RemovedAdmin", removedAdminRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskV2RemovedAdminIterator{contract: _NegRiskV2.contract, event: "RemovedAdmin", logs: logs, sub: sub}, nil
}

// WatchRemovedAdmin is a free log subscription operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed removedAdmin, address indexed admin)
func (_NegRiskV2 *NegRiskV2Filterer) WatchRemovedAdmin(opts *bind.WatchOpts, sink chan<- *NegRiskV2RemovedAdmin, removedAdmin []common.Address, admin []common.Address) (event.Subscription, error) {

	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskV2.contract.WatchLogs(opts, "RemovedAdmin", removedAdminRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskV2RemovedAdmin)
				if err := _NegRiskV2.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
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

// ParseRemovedAdmin is a log parse operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed removedAdmin, address indexed admin)
func (_NegRiskV2 *NegRiskV2Filterer) ParseRemovedAdmin(log types.Log) (*NegRiskV2RemovedAdmin, error) {
	event := new(NegRiskV2RemovedAdmin)
	if err := _NegRiskV2.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskV2RemovedOperatorIterator is returned from FilterRemovedOperator and is used to iterate over the raw logs and unpacked data for RemovedOperator events raised by the NegRiskV2 contract.
type NegRiskV2RemovedOperatorIterator struct {
	Event *NegRiskV2RemovedOperator // Event containing the contract specifics and raw log

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
func (it *NegRiskV2RemovedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskV2RemovedOperator)
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
		it.Event = new(NegRiskV2RemovedOperator)
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
func (it *NegRiskV2RemovedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskV2RemovedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskV2RemovedOperator represents a RemovedOperator event raised by the NegRiskV2 contract.
type NegRiskV2RemovedOperator struct {
	RemovedOperator common.Address
	Admin           common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRemovedOperator is a free log retrieval operation binding the contract event 0xf7262ed0443cc211121ceb1a80d69004f319245615a7488f951f1437fd91642c.
//
// Solidity: event RemovedOperator(address indexed removedOperator, address indexed admin)
func (_NegRiskV2 *NegRiskV2Filterer) FilterRemovedOperator(opts *bind.FilterOpts, removedOperator []common.Address, admin []common.Address) (*NegRiskV2RemovedOperatorIterator, error) {

	var removedOperatorRule []interface{}
	for _, removedOperatorItem := range removedOperator {
		removedOperatorRule = append(removedOperatorRule, removedOperatorItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskV2.contract.FilterLogs(opts, "RemovedOperator", removedOperatorRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskV2RemovedOperatorIterator{contract: _NegRiskV2.contract, event: "RemovedOperator", logs: logs, sub: sub}, nil
}

// WatchRemovedOperator is a free log subscription operation binding the contract event 0xf7262ed0443cc211121ceb1a80d69004f319245615a7488f951f1437fd91642c.
//
// Solidity: event RemovedOperator(address indexed removedOperator, address indexed admin)
func (_NegRiskV2 *NegRiskV2Filterer) WatchRemovedOperator(opts *bind.WatchOpts, sink chan<- *NegRiskV2RemovedOperator, removedOperator []common.Address, admin []common.Address) (event.Subscription, error) {

	var removedOperatorRule []interface{}
	for _, removedOperatorItem := range removedOperator {
		removedOperatorRule = append(removedOperatorRule, removedOperatorItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskV2.contract.WatchLogs(opts, "RemovedOperator", removedOperatorRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskV2RemovedOperator)
				if err := _NegRiskV2.contract.UnpackLog(event, "RemovedOperator", log); err != nil {
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

// ParseRemovedOperator is a log parse operation binding the contract event 0xf7262ed0443cc211121ceb1a80d69004f319245615a7488f951f1437fd91642c.
//
// Solidity: event RemovedOperator(address indexed removedOperator, address indexed admin)
func (_NegRiskV2 *NegRiskV2Filterer) ParseRemovedOperator(log types.Log) (*NegRiskV2RemovedOperator, error) {
	event := new(NegRiskV2RemovedOperator)
	if err := _NegRiskV2.contract.UnpackLog(event, "RemovedOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskV2TradingPausedIterator is returned from FilterTradingPaused and is used to iterate over the raw logs and unpacked data for TradingPaused events raised by the NegRiskV2 contract.
type NegRiskV2TradingPausedIterator struct {
	Event *NegRiskV2TradingPaused // Event containing the contract specifics and raw log

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
func (it *NegRiskV2TradingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskV2TradingPaused)
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
		it.Event = new(NegRiskV2TradingPaused)
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
func (it *NegRiskV2TradingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskV2TradingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskV2TradingPaused represents a TradingPaused event raised by the NegRiskV2 contract.
type NegRiskV2TradingPaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTradingPaused is a free log retrieval operation binding the contract event 0x203c4bd3e526634f661575359ff30de3b0edaba6c2cb1eac60f730b6d2d9d536.
//
// Solidity: event TradingPaused(address indexed pauser)
func (_NegRiskV2 *NegRiskV2Filterer) FilterTradingPaused(opts *bind.FilterOpts, pauser []common.Address) (*NegRiskV2TradingPausedIterator, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _NegRiskV2.contract.FilterLogs(opts, "TradingPaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskV2TradingPausedIterator{contract: _NegRiskV2.contract, event: "TradingPaused", logs: logs, sub: sub}, nil
}

// WatchTradingPaused is a free log subscription operation binding the contract event 0x203c4bd3e526634f661575359ff30de3b0edaba6c2cb1eac60f730b6d2d9d536.
//
// Solidity: event TradingPaused(address indexed pauser)
func (_NegRiskV2 *NegRiskV2Filterer) WatchTradingPaused(opts *bind.WatchOpts, sink chan<- *NegRiskV2TradingPaused, pauser []common.Address) (event.Subscription, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _NegRiskV2.contract.WatchLogs(opts, "TradingPaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskV2TradingPaused)
				if err := _NegRiskV2.contract.UnpackLog(event, "TradingPaused", log); err != nil {
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

// ParseTradingPaused is a log parse operation binding the contract event 0x203c4bd3e526634f661575359ff30de3b0edaba6c2cb1eac60f730b6d2d9d536.
//
// Solidity: event TradingPaused(address indexed pauser)
func (_NegRiskV2 *NegRiskV2Filterer) ParseTradingPaused(log types.Log) (*NegRiskV2TradingPaused, error) {
	event := new(NegRiskV2TradingPaused)
	if err := _NegRiskV2.contract.UnpackLog(event, "TradingPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskV2TradingUnpausedIterator is returned from FilterTradingUnpaused and is used to iterate over the raw logs and unpacked data for TradingUnpaused events raised by the NegRiskV2 contract.
type NegRiskV2TradingUnpausedIterator struct {
	Event *NegRiskV2TradingUnpaused // Event containing the contract specifics and raw log

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
func (it *NegRiskV2TradingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskV2TradingUnpaused)
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
		it.Event = new(NegRiskV2TradingUnpaused)
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
func (it *NegRiskV2TradingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskV2TradingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskV2TradingUnpaused represents a TradingUnpaused event raised by the NegRiskV2 contract.
type NegRiskV2TradingUnpaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTradingUnpaused is a free log retrieval operation binding the contract event 0xa1e8a54850dbd7f520bcc09f47bff152294b77b2081da545a7adf531b7ea283b.
//
// Solidity: event TradingUnpaused(address indexed pauser)
func (_NegRiskV2 *NegRiskV2Filterer) FilterTradingUnpaused(opts *bind.FilterOpts, pauser []common.Address) (*NegRiskV2TradingUnpausedIterator, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _NegRiskV2.contract.FilterLogs(opts, "TradingUnpaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskV2TradingUnpausedIterator{contract: _NegRiskV2.contract, event: "TradingUnpaused", logs: logs, sub: sub}, nil
}

// WatchTradingUnpaused is a free log subscription operation binding the contract event 0xa1e8a54850dbd7f520bcc09f47bff152294b77b2081da545a7adf531b7ea283b.
//
// Solidity: event TradingUnpaused(address indexed pauser)
func (_NegRiskV2 *NegRiskV2Filterer) WatchTradingUnpaused(opts *bind.WatchOpts, sink chan<- *NegRiskV2TradingUnpaused, pauser []common.Address) (event.Subscription, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _NegRiskV2.contract.WatchLogs(opts, "TradingUnpaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskV2TradingUnpaused)
				if err := _NegRiskV2.contract.UnpackLog(event, "TradingUnpaused", log); err != nil {
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

// ParseTradingUnpaused is a log parse operation binding the contract event 0xa1e8a54850dbd7f520bcc09f47bff152294b77b2081da545a7adf531b7ea283b.
//
// Solidity: event TradingUnpaused(address indexed pauser)
func (_NegRiskV2 *NegRiskV2Filterer) ParseTradingUnpaused(log types.Log) (*NegRiskV2TradingUnpaused, error) {
	event := new(NegRiskV2TradingUnpaused)
	if err := _NegRiskV2.contract.UnpackLog(event, "TradingUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskV2UserPauseBlockIntervalUpdatedIterator is returned from FilterUserPauseBlockIntervalUpdated and is used to iterate over the raw logs and unpacked data for UserPauseBlockIntervalUpdated events raised by the NegRiskV2 contract.
type NegRiskV2UserPauseBlockIntervalUpdatedIterator struct {
	Event *NegRiskV2UserPauseBlockIntervalUpdated // Event containing the contract specifics and raw log

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
func (it *NegRiskV2UserPauseBlockIntervalUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskV2UserPauseBlockIntervalUpdated)
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
		it.Event = new(NegRiskV2UserPauseBlockIntervalUpdated)
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
func (it *NegRiskV2UserPauseBlockIntervalUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskV2UserPauseBlockIntervalUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskV2UserPauseBlockIntervalUpdated represents a UserPauseBlockIntervalUpdated event raised by the NegRiskV2 contract.
type NegRiskV2UserPauseBlockIntervalUpdated struct {
	OldInterval *big.Int
	NewInterval *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUserPauseBlockIntervalUpdated is a free log retrieval operation binding the contract event 0x8c8acf678b7cd311e3b5768c92794d63943684862fdea390856e14d9e2a9ef88.
//
// Solidity: event UserPauseBlockIntervalUpdated(uint256 oldInterval, uint256 newInterval)
func (_NegRiskV2 *NegRiskV2Filterer) FilterUserPauseBlockIntervalUpdated(opts *bind.FilterOpts) (*NegRiskV2UserPauseBlockIntervalUpdatedIterator, error) {

	logs, sub, err := _NegRiskV2.contract.FilterLogs(opts, "UserPauseBlockIntervalUpdated")
	if err != nil {
		return nil, err
	}
	return &NegRiskV2UserPauseBlockIntervalUpdatedIterator{contract: _NegRiskV2.contract, event: "UserPauseBlockIntervalUpdated", logs: logs, sub: sub}, nil
}

// WatchUserPauseBlockIntervalUpdated is a free log subscription operation binding the contract event 0x8c8acf678b7cd311e3b5768c92794d63943684862fdea390856e14d9e2a9ef88.
//
// Solidity: event UserPauseBlockIntervalUpdated(uint256 oldInterval, uint256 newInterval)
func (_NegRiskV2 *NegRiskV2Filterer) WatchUserPauseBlockIntervalUpdated(opts *bind.WatchOpts, sink chan<- *NegRiskV2UserPauseBlockIntervalUpdated) (event.Subscription, error) {

	logs, sub, err := _NegRiskV2.contract.WatchLogs(opts, "UserPauseBlockIntervalUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskV2UserPauseBlockIntervalUpdated)
				if err := _NegRiskV2.contract.UnpackLog(event, "UserPauseBlockIntervalUpdated", log); err != nil {
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

// ParseUserPauseBlockIntervalUpdated is a log parse operation binding the contract event 0x8c8acf678b7cd311e3b5768c92794d63943684862fdea390856e14d9e2a9ef88.
//
// Solidity: event UserPauseBlockIntervalUpdated(uint256 oldInterval, uint256 newInterval)
func (_NegRiskV2 *NegRiskV2Filterer) ParseUserPauseBlockIntervalUpdated(log types.Log) (*NegRiskV2UserPauseBlockIntervalUpdated, error) {
	event := new(NegRiskV2UserPauseBlockIntervalUpdated)
	if err := _NegRiskV2.contract.UnpackLog(event, "UserPauseBlockIntervalUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskV2UserPausedIterator is returned from FilterUserPaused and is used to iterate over the raw logs and unpacked data for UserPaused events raised by the NegRiskV2 contract.
type NegRiskV2UserPausedIterator struct {
	Event *NegRiskV2UserPaused // Event containing the contract specifics and raw log

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
func (it *NegRiskV2UserPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskV2UserPaused)
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
		it.Event = new(NegRiskV2UserPaused)
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
func (it *NegRiskV2UserPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskV2UserPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskV2UserPaused represents a UserPaused event raised by the NegRiskV2 contract.
type NegRiskV2UserPaused struct {
	User                common.Address
	EffectivePauseBlock *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUserPaused is a free log retrieval operation binding the contract event 0xa3e76126f19eb25001b29726d2a9502b6377938633d2d6a955107dd442e7a14a.
//
// Solidity: event UserPaused(address indexed user, uint256 effectivePauseBlock)
func (_NegRiskV2 *NegRiskV2Filterer) FilterUserPaused(opts *bind.FilterOpts, user []common.Address) (*NegRiskV2UserPausedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _NegRiskV2.contract.FilterLogs(opts, "UserPaused", userRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskV2UserPausedIterator{contract: _NegRiskV2.contract, event: "UserPaused", logs: logs, sub: sub}, nil
}

// WatchUserPaused is a free log subscription operation binding the contract event 0xa3e76126f19eb25001b29726d2a9502b6377938633d2d6a955107dd442e7a14a.
//
// Solidity: event UserPaused(address indexed user, uint256 effectivePauseBlock)
func (_NegRiskV2 *NegRiskV2Filterer) WatchUserPaused(opts *bind.WatchOpts, sink chan<- *NegRiskV2UserPaused, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _NegRiskV2.contract.WatchLogs(opts, "UserPaused", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskV2UserPaused)
				if err := _NegRiskV2.contract.UnpackLog(event, "UserPaused", log); err != nil {
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

// ParseUserPaused is a log parse operation binding the contract event 0xa3e76126f19eb25001b29726d2a9502b6377938633d2d6a955107dd442e7a14a.
//
// Solidity: event UserPaused(address indexed user, uint256 effectivePauseBlock)
func (_NegRiskV2 *NegRiskV2Filterer) ParseUserPaused(log types.Log) (*NegRiskV2UserPaused, error) {
	event := new(NegRiskV2UserPaused)
	if err := _NegRiskV2.contract.UnpackLog(event, "UserPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskV2UserUnpausedIterator is returned from FilterUserUnpaused and is used to iterate over the raw logs and unpacked data for UserUnpaused events raised by the NegRiskV2 contract.
type NegRiskV2UserUnpausedIterator struct {
	Event *NegRiskV2UserUnpaused // Event containing the contract specifics and raw log

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
func (it *NegRiskV2UserUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskV2UserUnpaused)
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
		it.Event = new(NegRiskV2UserUnpaused)
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
func (it *NegRiskV2UserUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskV2UserUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskV2UserUnpaused represents a UserUnpaused event raised by the NegRiskV2 contract.
type NegRiskV2UserUnpaused struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUserUnpaused is a free log retrieval operation binding the contract event 0x1419d4111b5c8636aecff843bf618525f4f8e1aa6898a14357021d68dde8af12.
//
// Solidity: event UserUnpaused(address indexed user)
func (_NegRiskV2 *NegRiskV2Filterer) FilterUserUnpaused(opts *bind.FilterOpts, user []common.Address) (*NegRiskV2UserUnpausedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _NegRiskV2.contract.FilterLogs(opts, "UserUnpaused", userRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskV2UserUnpausedIterator{contract: _NegRiskV2.contract, event: "UserUnpaused", logs: logs, sub: sub}, nil
}

// WatchUserUnpaused is a free log subscription operation binding the contract event 0x1419d4111b5c8636aecff843bf618525f4f8e1aa6898a14357021d68dde8af12.
//
// Solidity: event UserUnpaused(address indexed user)
func (_NegRiskV2 *NegRiskV2Filterer) WatchUserUnpaused(opts *bind.WatchOpts, sink chan<- *NegRiskV2UserUnpaused, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _NegRiskV2.contract.WatchLogs(opts, "UserUnpaused", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskV2UserUnpaused)
				if err := _NegRiskV2.contract.UnpackLog(event, "UserUnpaused", log); err != nil {
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

// ParseUserUnpaused is a log parse operation binding the contract event 0x1419d4111b5c8636aecff843bf618525f4f8e1aa6898a14357021d68dde8af12.
//
// Solidity: event UserUnpaused(address indexed user)
func (_NegRiskV2 *NegRiskV2Filterer) ParseUserUnpaused(log types.Log) (*NegRiskV2UserUnpaused, error) {
	event := new(NegRiskV2UserUnpaused)
	if err := _NegRiskV2.contract.UnpackLog(event, "UserUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
