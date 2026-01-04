package polymarketcontracts

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ivanzzeth/ethclient"
	"github.com/ivanzzeth/ethsig"
	"github.com/ivanzzeth/ethsig/eip712"
	"github.com/ivanzzeth/polymarket-go-contracts/constants"
	conditional_tokens "github.com/ivanzzeth/polymarket-go-contracts/contracts/conditional-tokens"
	"github.com/ivanzzeth/polymarket-go-contracts/contracts/erc20"
	"github.com/ivanzzeth/polymarket-go-contracts/contracts/exchange"
	exchangefees "github.com/ivanzzeth/polymarket-go-contracts/contracts/exchange-fees"
	gnosissafel2 "github.com/ivanzzeth/polymarket-go-contracts/contracts/gnosis-safe-l2"
	negrisk "github.com/ivanzzeth/polymarket-go-contracts/contracts/neg-risk"
	negriskadapter "github.com/ivanzzeth/polymarket-go-contracts/contracts/neg-risk-adapter"
	negriskfees "github.com/ivanzzeth/polymarket-go-contracts/contracts/neg-risk-fees"
	safeproxyfactory "github.com/ivanzzeth/polymarket-go-contracts/contracts/safe-proxy-factory"
	"github.com/ivanzzeth/polymarket-go-contracts/sender"
	"github.com/ivanzzeth/polymarket-go-contracts/signer"
)

// BalanceAllowanceInfo holds balance and allowance information
type BalanceAllowanceInfo struct {
	Balance                    *big.Int
	AllowanceExchange          *big.Int
	AllowanceNegRiskAdapter    *big.Int
	AllowanceNegRiskExchange   *big.Int
	CTFApprovedExchange        bool
	CTFApprovedNegRiskAdapter  bool
	CTFApprovedNegRiskExchange bool
}

type ContractInterface struct {
	chainID           *big.Int
	client            ethclient.EthClientInterface
	contractConfig    *ContractConfig
	signatureType     SignatureType
	eoaTradingSigner  signer.EOATradingSigner
	safeTradingSigner signer.SafeTradingSigner
	txSender          sender.TransactionSender

	collateralContract        *erc20.Erc20
	conditionalTokensContract *conditional_tokens.ConditionalTokens
	exchangeContract          *exchange.Exchange
	exchangeFeesContract      *exchangefees.ExchangeFees
	negRiskAdapterContract    *negriskadapter.NegRiskAdapter
	negRiskContract           *negrisk.NegRisk
	negRiskFeesContract       *negriskfees.NegRiskFees
	safeProxyFactoryContract  *safeproxyfactory.SafeProxyFactory

	// Cache for Safe addresses (key: EOA address string, value: Safe address)
	safeAddressCache sync.Map
}

type ContractInterfaceConfig struct {
	SignatureType     SignatureType
	TxSender          sender.TransactionSender
	EOATradingSigner  signer.EOATradingSigner
	SafeTradingSigner signer.SafeTradingSigner

	ContractConfig *ContractConfig
}

type ContractInterfaceOption func(c *ContractInterfaceConfig)

func WithContractConfig(cc *ContractConfig) ContractInterfaceOption {
	return func(c *ContractInterfaceConfig) {
		c.ContractConfig = cc
	}
}

func WithSafeSigner(safeSigner signer.SafeTradingSigner) ContractInterfaceOption {
	return func(c *ContractInterfaceConfig) {
		c.SignatureType = SignatureTypePolyGnosisSafe
		c.SafeTradingSigner = safeSigner
	}
}

func WithEOASigner(eoaSigner signer.EOATradingSigner) ContractInterfaceOption {
	return func(c *ContractInterfaceConfig) {
		c.SignatureType = SignatureTypeEOA
		c.EOATradingSigner = eoaSigner
	}
}

func WithTransactionSender(txSender sender.TransactionSender) ContractInterfaceOption {
	return func(c *ContractInterfaceConfig) {
		c.TxSender = txSender
	}
}

func NewContractInterface(
	client ethclient.EthClientInterface,
	options ...ContractInterfaceOption,
) (*ContractInterface, error) {
	defaultOptions := &ContractInterfaceConfig{
		SignatureType:  SignatureTypeEOA,
		ContractConfig: MATIC_CONTRACTS,
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chainID: %v", err)
	}

	for _, opFn := range options {
		opFn(defaultOptions)
	}

	if defaultOptions.TxSender == nil {
		var txSender sender.TransactionSender
		switch defaultOptions.SignatureType {
		case SignatureTypePolyGnosisSafe:
			txSender = defaultOptions.SafeTradingSigner
		case SignatureTypeEOA:
			txSender, err = signer.GetTransactionSenderBySigner(chainID, client, defaultOptions.EOATradingSigner)
			if err != nil {
				return nil, fmt.Errorf("failed to get txSender for EOA signer: %v", err)
			}
		case SignatureTypePolyProxy:
			return nil, fmt.Errorf("not support polyProxy")
		}

		defaultOptions.TxSender = txSender
	}

	// Initialize Collateral (USDC) contract
	usdcContract, err := erc20.NewErc20(defaultOptions.ContractConfig.Collateral, client)
	if err != nil {
		return nil, fmt.Errorf("failed to setup USDC: %v", err)
	}

	// Initialize Conditional Tokens Framework contract
	ctfContract, err := conditional_tokens.NewConditionalTokens(defaultOptions.ContractConfig.ConditionalTokens, client)
	if err != nil {
		return nil, fmt.Errorf("failed to setup ConditionalTokens: %v", err)
	}

	// Initialize Exchange contract
	exchangeContract, err := exchange.NewExchange(defaultOptions.ContractConfig.Exchange, client)
	if err != nil {
		return nil, fmt.Errorf("failed to setup Exchange: %v", err)
	}

	// Initialize Exchange Fees contract
	exchangeFeesContract, err := exchangefees.NewExchangeFees(defaultOptions.ContractConfig.Exchange, client)
	if err != nil {
		return nil, fmt.Errorf("failed to setup ExchangeFees: %v", err)
	}

	// Initialize NegRisk Adapter contract
	negRiskAdapterContract, err := negriskadapter.NewNegRiskAdapter(defaultOptions.ContractConfig.NegRiskAdapter, client)
	if err != nil {
		return nil, fmt.Errorf("failed to setup NegRiskAdapter: %v", err)
	}

	// Initialize NegRisk Exchange contract
	negRiskContract, err := negrisk.NewNegRisk(defaultOptions.ContractConfig.NegRiskExchange, client)
	if err != nil {
		return nil, fmt.Errorf("failed to setup NegRisk: %v", err)
	}

	// Initialize NegRisk Fees contract
	negRiskFeesContract, err := negriskfees.NewNegRiskFees(defaultOptions.ContractConfig.NegRiskExchange, client)
	if err != nil {
		return nil, fmt.Errorf("failed to setup NegRiskFees: %v", err)
	}

	// Initialize SafeProxyFactory contract (optional, may be zero address)
	var safeProxyFactoryContract *safeproxyfactory.SafeProxyFactory
	if defaultOptions.ContractConfig.SafeProxyFactory != (common.Address{}) {
		safeProxyFactoryContract, err = safeproxyfactory.NewSafeProxyFactory(defaultOptions.ContractConfig.SafeProxyFactory, client)
		if err != nil {
			return nil, fmt.Errorf("failed to setup SafeProxyFactory: %v", err)
		}
	}

	return &ContractInterface{
		chainID:           chainID,
		client:            client,
		contractConfig:    defaultOptions.ContractConfig,
		signatureType:     defaultOptions.SignatureType,
		txSender:          defaultOptions.TxSender,
		safeTradingSigner: defaultOptions.SafeTradingSigner,
		eoaTradingSigner:  defaultOptions.EOATradingSigner,

		collateralContract:        usdcContract,
		conditionalTokensContract: ctfContract,
		exchangeContract:          exchangeContract,
		exchangeFeesContract:      exchangeFeesContract,
		negRiskAdapterContract:    negRiskAdapterContract,
		negRiskContract:           negRiskContract,
		negRiskFeesContract:       negRiskFeesContract,
		safeProxyFactoryContract:  safeProxyFactoryContract,
	}, nil
}

// GetCollateral returns the Collateral (USDC) contract instance
func (b *ContractInterface) GetCollateral() *erc20.Erc20 {
	return b.collateralContract
}

// GetConditionalTokens returns the Conditional Tokens Framework contract instance
func (b *ContractInterface) GetConditionalTokens() *conditional_tokens.ConditionalTokens {
	return b.conditionalTokensContract
}

// GetExchange returns the Exchange contract instance
func (b *ContractInterface) GetExchange() *exchange.Exchange {
	return b.exchangeContract
}

// GetExchangeFees returns the Exchange Fees contract instance
func (b *ContractInterface) GetExchangeFees() *exchangefees.ExchangeFees {
	return b.exchangeFeesContract
}

// GetNegRiskAdapter returns the NegRisk Adapter contract instance
func (b *ContractInterface) GetNegRiskAdapter() *negriskadapter.NegRiskAdapter {
	return b.negRiskAdapterContract
}

// GetNegRisk returns the NegRisk Exchange contract instance
func (b *ContractInterface) GetNegRisk() *negrisk.NegRisk {
	return b.negRiskContract
}

// GetNegRiskFees returns the NegRisk Fees contract instance
func (b *ContractInterface) GetNegRiskFees() *negriskfees.NegRiskFees {
	return b.negRiskFeesContract
}

// GetSafeProxyFactory returns the Safe Proxy Factory contract instance (may be nil)
func (b *ContractInterface) GetSafeProxyFactory() *safeproxyfactory.SafeProxyFactory {
	return b.safeProxyFactoryContract
}

// GetGnosisSafeL2 creates and returns a GnosisSafeL2 contract instance for the given Safe address
func (b *ContractInterface) GetGnosisSafeL2(safeAddress common.Address) (*gnosissafel2.GnosisSafeL2, error) {
	return gnosissafel2.NewGnosisSafeL2(safeAddress, b.client)
}

// GetConfig returns the contract configuration
func (b *ContractInterface) GetConfig() *ContractConfig {
	return b.contractConfig
}

// GetClient returns the Ethereum client
func (b *ContractInterface) GetClient() ethclient.EthClientInterface {
	return b.client
}

// GetSignatureType returns the signature type used by this interface
func (b *ContractInterface) GetSignatureType() SignatureType {
	return b.signatureType
}

// GetTxSender returns the transaction sender
func (b *ContractInterface) GetTxSender() sender.TransactionSender {
	return b.getTxSender()
}

// GetSafeTradingSigner returns the Safe trading signer
func (b *ContractInterface) GetSafeTradingSigner() signer.SafeTradingSigner {
	return b.safeTradingSigner
}

// GetEOATradingSigner returns the EOA trading signer
func (b *ContractInterface) GetEOATradingSigner() signer.EOATradingSigner {
	return b.eoaTradingSigner
}

func (b *ContractInterface) getTxSender() sender.TransactionSender {
	if b.txSender != nil {
		return b.txSender
	}

	if b.safeTradingSigner != nil {
		return b.safeTradingSigner
	}

	panic("no any transaction sender provided")
}

func (b *ContractInterface) getSafeTradingSigner() signer.SafeTradingSigner {
	if b.safeTradingSigner == nil {
		panic("no safe trading signer provided")
	}

	return b.safeTradingSigner
}

func (b *ContractInterface) getEOATradingSigner() signer.EOATradingSigner {
	if b.eoaTradingSigner == nil {
		panic("no EOA trading signer provided")
	}

	return b.eoaTradingSigner
}

// CheckBalanceAndAllowance checks the USDC balance and all allowances for the given address
func (b *ContractInterface) CheckBalanceAndAllowance(ctx context.Context, address common.Address) (*BalanceAllowanceInfo, error) {
	return b.CheckBalanceAndAllowanceAtBlock(ctx, address, nil)
}

// CheckBalanceAndAllowanceAtBlock checks the USDC balance and all allowances for the given address at a specific block
// If blockNumber is nil, it uses the latest block
func (b *ContractInterface) CheckBalanceAndAllowanceAtBlock(ctx context.Context, address common.Address, blockNumber *big.Int) (*BalanceAllowanceInfo, error) {
	info := &BalanceAllowanceInfo{}
	callOpts := &bind.CallOpts{Context: ctx, BlockNumber: blockNumber}

	// Check USDC balance
	balance, err := b.collateralContract.BalanceOf(callOpts, address)
	if err != nil {
		return nil, fmt.Errorf("failed to get USDC balance: %w", err)
	}
	info.Balance = balance

	// Check USDC allowances
	allowanceExchange, err := b.collateralContract.Allowance(callOpts, address, b.contractConfig.Exchange)
	if err != nil {
		return nil, fmt.Errorf("failed to get USDC allowance for Exchange: %w", err)
	}
	info.AllowanceExchange = allowanceExchange

	allowanceNegRiskAdapter, err := b.collateralContract.Allowance(callOpts, address, b.contractConfig.NegRiskAdapter)
	if err != nil {
		return nil, fmt.Errorf("failed to get USDC allowance for NegRiskAdapter: %w", err)
	}
	info.AllowanceNegRiskAdapter = allowanceNegRiskAdapter

	allowanceNegRiskExchange, err := b.collateralContract.Allowance(callOpts, address, b.contractConfig.NegRiskExchange)
	if err != nil {
		return nil, fmt.Errorf("failed to get USDC allowance for NegRiskExchange: %w", err)
	}
	info.AllowanceNegRiskExchange = allowanceNegRiskExchange

	// Check CTF approvals
	ctfApprovedExchange, err := b.conditionalTokensContract.IsApprovedForAll(callOpts, address, b.contractConfig.Exchange)
	if err != nil {
		return nil, fmt.Errorf("failed to check CTF approval for Exchange: %w", err)
	}
	info.CTFApprovedExchange = ctfApprovedExchange

	ctfApprovedNegRiskAdapter, err := b.conditionalTokensContract.IsApprovedForAll(callOpts, address, b.contractConfig.NegRiskAdapter)
	if err != nil {
		return nil, fmt.Errorf("failed to check CTF approval for NegRiskAdapter: %w", err)
	}
	info.CTFApprovedNegRiskAdapter = ctfApprovedNegRiskAdapter

	ctfApprovedNegRiskExchange, err := b.conditionalTokensContract.IsApprovedForAll(callOpts, address, b.contractConfig.NegRiskExchange)
	if err != nil {
		return nil, fmt.Errorf("failed to check CTF approval for NegRiskExchange: %w", err)
	}
	info.CTFApprovedNegRiskExchange = ctfApprovedNegRiskExchange

	return info, nil
}

// PrintBalanceAndAllowance prints the balance and allowance information in a user-friendly format
func (b *ContractInterface) PrintBalanceAndAllowance(ctx context.Context, address common.Address) error {
	info, err := b.CheckBalanceAndAllowance(ctx, address)
	if err != nil {
		return err
	}

	fmt.Println("=== Balance and Allowance Status ===")

	// Print USDC Balance (assuming 6 decimals)
	fmt.Printf("USDC Balance: %s USDC\n\n", formatUSDC(info.Balance))

	// Print Allowances
	fmt.Println("Allowances:")
	fmt.Printf("  USDC → Exchange: %s\n", checkmark(info.AllowanceExchange.Cmp(big.NewInt(0)) > 0))
	fmt.Printf("  USDC → NegRiskAdapter: %s\n", checkmark(info.AllowanceNegRiskAdapter.Cmp(big.NewInt(0)) > 0))
	fmt.Printf("  USDC → NegRiskExchange: %s\n", checkmark(info.AllowanceNegRiskExchange.Cmp(big.NewInt(0)) > 0))
	fmt.Printf("  CTF → Exchange: %s\n", checkmark(info.CTFApprovedExchange))
	fmt.Printf("  CTF → NegRiskAdapter: %s\n", checkmark(info.CTFApprovedNegRiskAdapter))
	fmt.Printf("  CTF → NegRiskExchange: %s\n\n", checkmark(info.CTFApprovedNegRiskExchange))

	return nil
}

// EstimateSafeTxGas estimates the gas required for a Safe transaction execution
// This uses simulateAndRevert to accurately estimate gas without requiring valid signatures
func (b *ContractInterface) EstimateSafeTxGas(safeAddr, to common.Address, value *big.Int, data []byte, operation SafeOperation) (*big.Int, error) {
	// Parse Safe ABI to encode the simulateAndRevert call
	parsedABI, err := abi.JSON(strings.NewReader(gnosissafel2.GnosisSafeL2MetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse GnosisSafeL2 ABI: %w", err)
	}

	// Encode the target transaction call data (the transaction we want to simulate)
	// We create the call data for the internal transaction
	var targetCallData []byte
	if len(data) > 0 {
		targetCallData = data
	} else {
		targetCallData = []byte{}
	}

	// Use simulateAndRevert to estimate gas
	// simulateAndRevert(address targetContract, bytes calldataPayload)
	simulateCallData, err := parsedABI.Pack("simulateAndRevert", to, targetCallData)
	if err != nil {
		return nil, fmt.Errorf("failed to pack simulateAndRevert call data: %w", err)
	}

	// Call simulateAndRevert - it will always revert with the execution result
	msg := ethereum.CallMsg{
		To:   &safeAddr,
		Data: simulateCallData,
	}

	// EstimateGas will fail because simulateAndRevert always reverts,
	// but we can extract the gas estimation from the error
	gasLimit, err := b.client.EstimateGas(context.Background(), msg)
	if err != nil {
		// simulateAndRevert always reverts, so we expect an error
		// However, EstimateGas still provides a gas estimate
		// If it's a revert with execution result, that's expected
		// For now, we'll use a fallback approach if estimation fails

		// Try a direct estimation on the target call
		directMsg := ethereum.CallMsg{
			From:  safeAddr, // Simulate call from Safe
			To:    &to,
			Value: value,
			Data:  targetCallData,
		}

		directGas, directErr := b.client.EstimateGas(context.Background(), directMsg)
		if directErr != nil {
			return nil, fmt.Errorf("failed to estimate gas: %w", directErr)
		}

		// Add Safe overhead: approximately 15000 gas for Safe execution logic
		// This includes signature verification, storage operations, and event emissions
		safeTxGas := big.NewInt(int64(directGas + 15000))
		// Add 50% buffer for safety to account for GS010 check requirements
		// GS010 requires: gasleft() >= ((safeTxGas * 64) / 63).max(safeTxGas + 2500) + 500
		safeTxGas = new(big.Int).Mul(safeTxGas, big.NewInt(150))
		safeTxGas = new(big.Int).Div(safeTxGas, big.NewInt(100))
		return safeTxGas, nil
	}

	// If EstimateGas succeeded (shouldn't happen with simulateAndRevert, but handle it)
	safeTxGas := big.NewInt(int64(gasLimit))
	// Add 50% buffer for safety to account for GS010 check requirements
	safeTxGas = new(big.Int).Mul(safeTxGas, big.NewInt(150))
	safeTxGas = new(big.Int).Div(safeTxGas, big.NewInt(100))
	return safeTxGas, nil
}

func (b *ContractInterface) GetSafeAddress(eoa common.Address) (common.Address, error) {
	// Check cache first
	if cached, ok := b.safeAddressCache.Load(eoa.Hex()); ok {
		return cached.(common.Address), nil
	}

	// Compute safe address
	safeAddr, err := b.GetSafeProxyFactory().ComputeProxyAddress(nil, eoa)
	if err != nil {
		return common.Address{}, err
	}

	// Store in cache
	b.safeAddressCache.Store(eoa.Hex(), safeAddr)

	return safeAddr, nil
}

func (b *ContractInterface) DeploySafe() (safeProxy common.Address, txHash common.Hash, err error) {
	return b.DeploySafeBySender(b.getTxSender(), b.getSafeTradingSigner())
}

func (b *ContractInterface) EnableTrading(ctx context.Context) ([]common.Hash, error) {
	switch b.signatureType {
	case SignatureTypePolyGnosisSafe:
		return b.EnableTradingForSafe(ctx, b.getSafeTradingSigner(), b.chainID)
	case SignatureTypeEOA:
		return b.EnableTradingForEOA(ctx, b.getEOATradingSigner())
	default:
		return nil, fmt.Errorf("unsupported signature type: %v", b.signatureType)
	}
}

func (b *ContractInterface) Redeem(ctx context.Context, conditionId [32]byte) (common.Hash, error) {
	// For Polymarket binary markets, indexSets is always [1, 2] representing Yes/No outcomes
	indexSets := []*big.Int{big.NewInt(1), big.NewInt(2)}

	switch b.signatureType {
	case SignatureTypePolyGnosisSafe:
		return b.RedeemPositionsForSafe(ctx, b.getSafeTradingSigner(), b.chainID, conditionId, indexSets)
	case SignatureTypeEOA:
		return b.RedeemPositionsForEOA(ctx, b.getEOATradingSigner(), conditionId, indexSets)
	default:
		return common.Hash{}, fmt.Errorf("unsupported signature type: %v", b.signatureType)
	}
}

func (b *ContractInterface) RedeemNegRisk(ctx context.Context, conditionId [32]byte, amounts []*big.Int) (common.Hash, error) {
	switch b.signatureType {
	case SignatureTypePolyGnosisSafe:
		return b.RedeemPositionsNegRiskForSafe(ctx, b.getSafeTradingSigner(), b.chainID, conditionId, amounts)
	case SignatureTypeEOA:
		return b.RedeemPositionsNegRiskForEOA(ctx, b.getEOATradingSigner(), conditionId, amounts)
	default:
		return common.Hash{}, fmt.Errorf("unsupported signature type: %v", b.signatureType)
	}
}

func (b *ContractInterface) Split(ctx context.Context, conditionId [32]byte, amount *big.Int) (common.Hash, error) {
	// For Polymarket binary markets, partition is always [1, 2] representing Yes/No outcomes
	partition := []*big.Int{big.NewInt(1), big.NewInt(2)}

	switch b.signatureType {
	case SignatureTypePolyGnosisSafe:
		return b.SplitPositionForSafe(ctx, b.getSafeTradingSigner(), b.chainID, conditionId, partition, amount)
	case SignatureTypeEOA:
		return b.SplitPositionForEOA(ctx, b.getEOATradingSigner(), conditionId, partition, amount)
	default:
		return common.Hash{}, fmt.Errorf("unsupported signature type: %v", b.signatureType)
	}
}

func (b *ContractInterface) Merge(ctx context.Context, conditionId [32]byte, amount *big.Int) (common.Hash, error) {
	// For Polymarket binary markets, partition is always [1, 2] representing Yes/No outcomes
	partition := []*big.Int{big.NewInt(1), big.NewInt(2)}

	switch b.signatureType {
	case SignatureTypePolyGnosisSafe:
		return b.MergePositionsForSafe(ctx, b.getSafeTradingSigner(), b.chainID, conditionId, partition, amount)
	case SignatureTypeEOA:
		return b.MergePositionsForEOA(ctx, b.getEOATradingSigner(), conditionId, partition, amount)
	default:
		return common.Hash{}, fmt.Errorf("unsupported signature type: %v", b.signatureType)
	}
}

func (b *ContractInterface) SplitNegRisk(ctx context.Context, conditionId [32]byte, amount *big.Int) (common.Hash, error) {
	switch b.signatureType {
	case SignatureTypePolyGnosisSafe:
		return b.SplitPositionNegRiskForSafe(ctx, b.getSafeTradingSigner(), b.chainID, conditionId, amount)
	case SignatureTypeEOA:
		return b.SplitPositionNegRiskForEOA(ctx, b.getEOATradingSigner(), conditionId, amount)
	default:
		return common.Hash{}, fmt.Errorf("unsupported signature type: %v", b.signatureType)
	}
}

func (b *ContractInterface) MergeNegRisk(ctx context.Context, conditionId [32]byte, amount *big.Int) (common.Hash, error) {
	switch b.signatureType {
	case SignatureTypePolyGnosisSafe:
		return b.MergePositionsNegRiskForSafe(ctx, b.getSafeTradingSigner(), b.chainID, conditionId, amount)
	case SignatureTypeEOA:
		return b.MergePositionsNegRiskForEOA(ctx, b.getEOATradingSigner(), conditionId, amount)
	default:
		return common.Hash{}, fmt.Errorf("unsupported signature type: %v", b.signatureType)
	}
}

func (b *ContractInterface) DeploySafeBySender(txSender sender.TransactionSender, signer ethsig.TypedDataSigner) (safeProxy common.Address, txHash common.Hash, err error) {
	zeroAddr := common.Address{}
	paymentToken := zeroAddr
	payment := big.NewInt(0)
	paymentReceiver := zeroAddr

	chainID, err := b.client.ChainID(context.Background())
	if err != nil {
		return common.Address{}, common.Hash{}, err
	}

	typedData := BuildCreateProxyTypedData(chainID, b.contractConfig.SafeProxyFactory, paymentToken, payment, paymentReceiver)
	signatureBytes, err := signer.SignTypedData(typedData)
	if err != nil {
		return common.Address{}, common.Hash{}, err
	}

	r, s, v, err := ethsig.ConvertSigBytes2RSV(signatureBytes)
	if err != nil {
		return common.Address{}, common.Hash{}, err
	}

	createSig := safeproxyfactory.SafeProxyFactorySig{
		R: r,
		S: s,
		V: v,
	}

	return b.DeploySafeWithSig(txSender, chainID, b.contractConfig.SafeProxyFactory, paymentToken, payment, paymentReceiver, createSig)
}

func (b *ContractInterface) DeploySafeWithSig(txSender sender.TransactionSender, chainID *big.Int, safeFactory, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address, createSig safeproxyfactory.SafeProxyFactorySig) (safeProxy common.Address, txHash common.Hash, err error) {
	typedData := BuildCreateProxyTypedData(chainID, safeFactory, paymentToken, payment, paymentReceiver)
	typedDataHash, _, err := eip712.TypedDataAndHash(typedData)
	if err != nil {
		return
	}

	// Convert to signature bytes with V normalized to 0/1 for crypto.SigToPub
	vNormalized := ethsig.DenormalizeV(createSig.V)
	sigBytes := ethsig.ConvertRSV2SigBytes(createSig.R, createSig.S, vNormalized)

	pubkey, err := crypto.SigToPub(typedDataHash, sigBytes)
	if err != nil {
		return
	}

	addr := crypto.PubkeyToAddress(*pubkey)
	safeProxy, err = b.GetSafeAddress(addr)
	if err != nil {
		return
	}

	code, err := b.client.CodeAt(context.Background(), safeProxy, nil)
	if err != nil {
		return
	}

	if len(code) != 0 {
		err = fmt.Errorf("already deployed")
		return
	}

	zeroAddr := common.Address{}
	factoryAbi, err := safeproxyfactory.SafeProxyFactoryMetaData.GetAbi()
	if err != nil {
		return
	}

	createProxyData, err := factoryAbi.Pack("createProxy", zeroAddr, big.NewInt(0), zeroAddr, createSig)
	if err != nil {
		return
	}

	txHash, err = txSender.SendEthereumTransaction(b.contractConfig.SafeProxyFactory, createProxyData, big.NewInt(0))
	if err != nil {
		return
	}

	return
}

func (b *ContractInterface) ExecuteTransactionBySafeAndSingleSigner(safeSigner signer.SafeTradingSigner, chainID *big.Int, safeAddr common.Address, to common.Address, value *big.Int, data []byte, operation SafeOperation, safeTxGas *big.Int) (common.Hash, error) {
	baseGas := big.NewInt(0)
	gasPrice := big.NewInt(0)
	gasToken := common.Address{}
	refundReceiver := common.Address{}
	safeContract, err := b.GetGnosisSafeL2(safeAddr)
	if err != nil {
		return common.Hash{}, err
	}

	nonce, err := safeContract.Nonce(nil)
	if err != nil {
		return common.Hash{}, err
	}

	// Estimate safeTxGas if not explicitly set or set to 0
	// IMPORTANT: Must do this BEFORE signing, as safeTxGas is part of the signed data
	if safeTxGas == nil || safeTxGas.Cmp(big.NewInt(0)) == 0 {
		estimatedGas, err := b.EstimateSafeTxGas(safeAddr, to, value, data, operation)
		if err != nil {
			return common.Hash{}, fmt.Errorf("failed to estimate safeTxGas: %w", err)
		}
		safeTxGas = estimatedGas
	}

	// Build the Safe transaction typed data with the correct safeTxGas
	typedData := BuildSafeTransactionTypedData(chainID, safeAddr, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, nonce)

	// Sign the typed data
	signature, err := safeSigner.SignTypedData(typedData)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to sign Safe transaction: %w", err)
	}

	// Execute the transaction with the signature
	return b.ExecuteTransactionBySafe(safeSigner, safeAddr, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signature)
}

func (b *ContractInterface) ExecuteTransactionBySafe(txSender sender.TransactionSender, safeAddr common.Address, to common.Address, value *big.Int, data []byte, operation SafeOperation, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (common.Hash, error) {
	safeAbi, err := gnosissafel2.GnosisSafeL2MetaData.GetAbi()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get Safe ABI: %w", err)
	}

	execTransactionData, err := safeAbi.Pack(
		"execTransaction",
		to,
		value,
		data,
		uint8(operation),
		safeTxGas,
		baseGas,
		gasPrice,
		gasToken,
		refundReceiver,
		signatures,
	)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack execTransaction: %w", err)
	}

	return txSender.SendEthereumTransaction(safeAddr, execTransactionData, big.NewInt(0))
}

// EnableTradingForSafe enables trading for a Safe wallet by setting all required allowances
// This function uses ExecuteTransactionBySafeAndSingleSigner to execute all approval operations
func (b *ContractInterface) EnableTradingForSafe(
	ctx context.Context,
	safeSigner signer.SafeTradingSigner,
	// txSigner signer.TypedDataSigner,
	chainID *big.Int,
) ([]common.Hash, error) {
	safeAddr, err := b.GetSafeAddress(safeSigner.GetAddress())
	if err != nil {
		return nil, fmt.Errorf("failed to get Safe address: %w", err)
	}
	var txHashes []common.Hash

	// Check current status
	info, err := b.CheckBalanceAndAllowance(ctx, safeAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to check balance and allowance: %w", err)
	}

	// Maximum allowance for ERC20 approvals
	maxAllowance := new(big.Int)
	maxAllowance.SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)

	// Parse ERC20 ABI for approve function
	erc20ABI, err := abi.JSON(strings.NewReader(erc20.Erc20MetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ERC20 ABI: %w", err)
	}

	// Parse ConditionalTokens ABI for setApprovalForAll function
	ctfABI, err := abi.JSON(strings.NewReader(conditional_tokens.ConditionalTokensMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ConditionalTokens ABI: %w", err)
	}

	// Approve USDC for all contracts if needed
	if info.AllowanceExchange.Cmp(big.NewInt(0)) == 0 {
		approveData, err := erc20ABI.Pack("approve", b.contractConfig.Exchange, maxAllowance)
		if err != nil {
			return nil, fmt.Errorf("failed to pack approve data for Exchange: %w", err)
		}

		txHash, err := b.ExecuteTransactionBySafeAndSingleSigner(
			safeSigner, chainID, safeAddr,
			b.contractConfig.Collateral, big.NewInt(0), approveData,
			SafeOperationCall, big.NewInt(0),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to execute Safe transaction for USDC → Exchange approval: %w", err)
		}
		txHashes = append(txHashes, txHash)
	}

	if info.AllowanceNegRiskAdapter.Cmp(big.NewInt(0)) == 0 {
		approveData, err := erc20ABI.Pack("approve", b.contractConfig.NegRiskAdapter, maxAllowance)
		if err != nil {
			return nil, fmt.Errorf("failed to pack approve data for NegRiskAdapter: %w", err)
		}

		txHash, err := b.ExecuteTransactionBySafeAndSingleSigner(
			safeSigner, chainID, safeAddr,
			b.contractConfig.Collateral, big.NewInt(0), approveData,
			SafeOperationCall, big.NewInt(0),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to execute Safe transaction for USDC → NegRiskAdapter approval: %w", err)
		}
		txHashes = append(txHashes, txHash)
	}

	if info.AllowanceNegRiskExchange.Cmp(big.NewInt(0)) == 0 {
		approveData, err := erc20ABI.Pack("approve", b.contractConfig.NegRiskExchange, maxAllowance)
		if err != nil {
			return nil, fmt.Errorf("failed to pack approve data for NegRiskExchange: %w", err)
		}

		txHash, err := b.ExecuteTransactionBySafeAndSingleSigner(
			safeSigner, chainID, safeAddr,
			b.contractConfig.Collateral, big.NewInt(0), approveData,
			SafeOperationCall, big.NewInt(0),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to execute Safe transaction for USDC → NegRiskExchange approval: %w", err)
		}
		txHashes = append(txHashes, txHash)
	}

	// Approve CTF for all contracts if needed
	if !info.CTFApprovedExchange {
		setApprovalData, err := ctfABI.Pack("setApprovalForAll", b.contractConfig.Exchange, true)
		if err != nil {
			return nil, fmt.Errorf("failed to pack setApprovalForAll data for Exchange: %w", err)
		}

		txHash, err := b.ExecuteTransactionBySafeAndSingleSigner(
			safeSigner, chainID, safeAddr,
			b.contractConfig.ConditionalTokens, big.NewInt(0), setApprovalData,
			SafeOperationCall, big.NewInt(0),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to execute Safe transaction for CTF → Exchange approval: %w", err)
		}
		txHashes = append(txHashes, txHash)
	}

	if !info.CTFApprovedNegRiskAdapter {
		setApprovalData, err := ctfABI.Pack("setApprovalForAll", b.contractConfig.NegRiskAdapter, true)
		if err != nil {
			return nil, fmt.Errorf("failed to pack setApprovalForAll data for NegRiskAdapter: %w", err)
		}

		txHash, err := b.ExecuteTransactionBySafeAndSingleSigner(
			safeSigner, chainID, safeAddr,
			b.contractConfig.ConditionalTokens, big.NewInt(0), setApprovalData,
			SafeOperationCall, big.NewInt(0),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to execute Safe transaction for CTF → NegRiskAdapter approval: %w", err)
		}
		txHashes = append(txHashes, txHash)
	}

	if !info.CTFApprovedNegRiskExchange {
		setApprovalData, err := ctfABI.Pack("setApprovalForAll", b.contractConfig.NegRiskExchange, true)
		if err != nil {
			return nil, fmt.Errorf("failed to pack setApprovalForAll data for NegRiskExchange: %w", err)
		}

		txHash, err := b.ExecuteTransactionBySafeAndSingleSigner(
			safeSigner, chainID, safeAddr,
			b.contractConfig.ConditionalTokens, big.NewInt(0), setApprovalData,
			SafeOperationCall, big.NewInt(0),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to execute Safe transaction for CTF → NegRiskExchange approval: %w", err)
		}
		txHashes = append(txHashes, txHash)
	}

	return txHashes, nil
}

// EnableTradingForEOA enables trading for an EOA wallet by setting all required allowances
// This function directly calls contract methods without going through Safe
func (b *ContractInterface) EnableTradingForEOA(
	ctx context.Context,
	eoaSigner signer.EOATradingSigner,
) ([]common.Hash, error) {
	var txHashes []common.Hash

	eoaAddr := eoaSigner.GetAddress()
	txSender := b.getTxSender()

	// Check current status
	info, err := b.CheckBalanceAndAllowance(ctx, eoaAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to check balance and allowance: %w", err)
	}

	// Maximum allowance for ERC20 approvals
	maxAllowance := new(big.Int)
	maxAllowance.SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)

	// Parse ERC20 ABI for approve function
	erc20ABI, err := abi.JSON(strings.NewReader(erc20.Erc20MetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ERC20 ABI: %w", err)
	}

	// Parse ConditionalTokens ABI for setApprovalForAll function
	ctfABI, err := abi.JSON(strings.NewReader(conditional_tokens.ConditionalTokensMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ConditionalTokens ABI: %w", err)
	}

	// Approve USDC for all contracts if needed
	if info.AllowanceExchange.Cmp(big.NewInt(0)) == 0 {
		approveData, err := erc20ABI.Pack("approve", b.contractConfig.Exchange, maxAllowance)
		if err != nil {
			return nil, fmt.Errorf("failed to pack approve data for Exchange: %w", err)
		}
		txHash, err := txSender.SendEthereumTransaction(b.contractConfig.Collateral, approveData, big.NewInt(0))
		if err != nil {
			return nil, fmt.Errorf("failed to send USDC → Exchange approval transaction: %w", err)
		}
		txHashes = append(txHashes, txHash)
	}

	if info.AllowanceNegRiskAdapter.Cmp(big.NewInt(0)) == 0 {
		approveData, err := erc20ABI.Pack("approve", b.contractConfig.NegRiskAdapter, maxAllowance)
		if err != nil {
			return nil, fmt.Errorf("failed to pack approve data for NegRiskAdapter: %w", err)
		}
		txHash, err := txSender.SendEthereumTransaction(b.contractConfig.Collateral, approveData, big.NewInt(0))
		if err != nil {
			return nil, fmt.Errorf("failed to send USDC → NegRiskAdapter approval transaction: %w", err)
		}
		txHashes = append(txHashes, txHash)
	}

	if info.AllowanceNegRiskExchange.Cmp(big.NewInt(0)) == 0 {
		approveData, err := erc20ABI.Pack("approve", b.contractConfig.NegRiskExchange, maxAllowance)
		if err != nil {
			return nil, fmt.Errorf("failed to pack approve data for NegRiskExchange: %w", err)
		}
		txHash, err := txSender.SendEthereumTransaction(b.contractConfig.Collateral, approveData, big.NewInt(0))
		if err != nil {
			return nil, fmt.Errorf("failed to send USDC → NegRiskExchange approval transaction: %w", err)
		}
		txHashes = append(txHashes, txHash)
	}

	// Approve CTF for all contracts if needed
	if !info.CTFApprovedExchange {
		setApprovalData, err := ctfABI.Pack("setApprovalForAll", b.contractConfig.Exchange, true)
		if err != nil {
			return nil, fmt.Errorf("failed to pack setApprovalForAll data for Exchange: %w", err)
		}
		txHash, err := txSender.SendEthereumTransaction(b.contractConfig.ConditionalTokens, setApprovalData, big.NewInt(0))
		if err != nil {
			return nil, fmt.Errorf("failed to send CTF → Exchange approval transaction: %w", err)
		}
		txHashes = append(txHashes, txHash)
	}

	if !info.CTFApprovedNegRiskAdapter {
		setApprovalData, err := ctfABI.Pack("setApprovalForAll", b.contractConfig.NegRiskAdapter, true)
		if err != nil {
			return nil, fmt.Errorf("failed to pack setApprovalForAll data for NegRiskAdapter: %w", err)
		}
		txHash, err := txSender.SendEthereumTransaction(b.contractConfig.ConditionalTokens, setApprovalData, big.NewInt(0))
		if err != nil {
			return nil, fmt.Errorf("failed to send CTF → NegRiskAdapter approval transaction: %w", err)
		}
		txHashes = append(txHashes, txHash)
	}

	if !info.CTFApprovedNegRiskExchange {
		setApprovalData, err := ctfABI.Pack("setApprovalForAll", b.contractConfig.NegRiskExchange, true)
		if err != nil {
			return nil, fmt.Errorf("failed to pack setApprovalForAll data for NegRiskExchange: %w", err)
		}
		txHash, err := txSender.SendEthereumTransaction(b.contractConfig.ConditionalTokens, setApprovalData, big.NewInt(0))
		if err != nil {
			return nil, fmt.Errorf("failed to send CTF → NegRiskExchange approval transaction: %w", err)
		}
		txHashes = append(txHashes, txHash)
	}

	return txHashes, nil
}

// RedeemPositionsForEOA redeems conditional tokens for a resolved market using EOA
func (b *ContractInterface) RedeemPositionsForEOA(
	ctx context.Context,
	eoaSigner signer.EOATradingSigner,
	conditionId [32]byte,
	indexSets []*big.Int,
) (common.Hash, error) {
	txSender := b.getTxSender()

	// Prepare calldata for redeemPositions
	parsedABI, err := conditional_tokens.ConditionalTokensMetaData.GetAbi()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to parse ConditionalTokens ABI: %w", err)
	}
	parentCollectionId := [32]byte{} // empty for root collection
	calldata, err := parsedABI.Pack("redeemPositions", b.contractConfig.Collateral, parentCollectionId, conditionId, indexSets)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack redeemPositions calldata: %w", err)
	}

	// Send transaction
	txHash, err := txSender.SendEthereumTransaction(b.contractConfig.ConditionalTokens, calldata, big.NewInt(0))
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to send redeem transaction: %w", err)
	}

	return txHash, nil
}

// RedeemPositionsNegRiskForEOA redeems NegRisk market positions using EOA
func (b *ContractInterface) RedeemPositionsNegRiskForEOA(
	ctx context.Context,
	eoaSigner signer.EOATradingSigner,
	conditionId [32]byte,
	amounts []*big.Int,
) (common.Hash, error) {
	txSender := b.getTxSender()

	// Prepare calldata for redeemPositions (NegRisk)
	parsedABI, err := negriskadapter.NegRiskAdapterMetaData.GetAbi()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to parse NegRiskAdapter ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("redeemPositions", conditionId, amounts)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack redeemPositions calldata: %w", err)
	}

	// Send transaction
	txHash, err := txSender.SendEthereumTransaction(b.contractConfig.NegRiskAdapter, calldata, big.NewInt(0))
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to send NegRisk redeem transaction: %w", err)
	}

	return txHash, nil
}

// SplitPositionForEOA splits collateral into conditional tokens for an EOA wallet
func (b *ContractInterface) SplitPositionForEOA(
	ctx context.Context,
	eoaSigner signer.EOATradingSigner,
	conditionId [32]byte,
	partition []*big.Int,
	amount *big.Int,
) (common.Hash, error) {
	txSender := b.getTxSender()

	// Prepare calldata for splitPosition
	parsedABI, err := conditional_tokens.ConditionalTokensMetaData.GetAbi()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to parse ConditionalTokens ABI: %w", err)
	}
	parentCollectionId := [32]byte{} // empty for root collection
	calldata, err := parsedABI.Pack("splitPosition", b.contractConfig.Collateral, parentCollectionId, conditionId, partition, amount)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack splitPosition calldata: %w", err)
	}

	// Send transaction
	txHash, err := txSender.SendEthereumTransaction(b.contractConfig.ConditionalTokens, calldata, big.NewInt(0))
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to send split transaction: %w", err)
	}

	return txHash, nil
}

// MergePositionsForEOA merges conditional tokens back into collateral for an EOA wallet
func (b *ContractInterface) MergePositionsForEOA(
	ctx context.Context,
	eoaSigner signer.EOATradingSigner,
	conditionId [32]byte,
	partition []*big.Int,
	amount *big.Int,
) (common.Hash, error) {
	txSender := b.getTxSender()

	// Prepare calldata for mergePositions
	parsedABI, err := conditional_tokens.ConditionalTokensMetaData.GetAbi()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to parse ConditionalTokens ABI: %w", err)
	}
	parentCollectionId := [32]byte{} // empty for root collection
	calldata, err := parsedABI.Pack("mergePositions", b.contractConfig.Collateral, parentCollectionId, conditionId, partition, amount)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack mergePositions calldata: %w", err)
	}

	// Send transaction
	txHash, err := txSender.SendEthereumTransaction(b.contractConfig.ConditionalTokens, calldata, big.NewInt(0))
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to send merge transaction: %w", err)
	}

	return txHash, nil
}

// RedeemPositionsForSafe redeems conditional tokens for a resolved market using TransactionSender
func (b *ContractInterface) RedeemPositionsForSafe(
	ctx context.Context,
	safeSigner signer.SafeTradingSigner,
	chainID *big.Int,
	conditionId [32]byte,
	indexSets []*big.Int,
) (common.Hash, error) {
	safeAddr, err := b.GetSafeAddress(safeSigner.GetAddress())
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get Safe address: %w", err)
	}

	// Prepare calldata for redeemPositions
	parsedABI, err := conditional_tokens.ConditionalTokensMetaData.GetAbi()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to parse ConditionalTokens ABI: %w", err)
	}
	parentCollectionId := [32]byte{} // empty for root collection
	calldata, err := parsedABI.Pack("redeemPositions", b.contractConfig.Collateral, parentCollectionId, conditionId, indexSets)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack redeemPositions calldata: %w", err)
	}

	// Execute Safe transaction
	txHash, err := b.ExecuteTransactionBySafeAndSingleSigner(
		safeSigner, chainID, safeAddr,
		b.contractConfig.ConditionalTokens,
		big.NewInt(0), // value
		calldata,
		SafeOperationCall,
		big.NewInt(0), // safeTxGas will be auto-estimated
	)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to execute Safe transaction: %w", err)
	}

	return txHash, nil
}

// RedeemPositionsNegRiskForSafe redeems NegRisk market positions using TransactionSender
func (b *ContractInterface) RedeemPositionsNegRiskForSafe(
	ctx context.Context,
	safeSigner signer.SafeTradingSigner,
	chainID *big.Int,
	conditionId [32]byte,
	amounts []*big.Int,
) (common.Hash, error) {
	safeAddr, err := b.GetSafeAddress(safeSigner.GetAddress())
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get Safe address: %w", err)
	}

	// Prepare calldata for redeemPositions (NegRisk)
	parsedABI, err := negriskadapter.NegRiskAdapterMetaData.GetAbi()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to parse NegRiskAdapter ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("redeemPositions", conditionId, amounts)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack redeemPositions calldata: %w", err)
	}

	// Execute Safe transaction
	txHash, err := b.ExecuteTransactionBySafeAndSingleSigner(
		safeSigner, chainID, safeAddr,
		b.contractConfig.NegRiskAdapter,
		big.NewInt(0), // value
		calldata,
		SafeOperationCall,
		big.NewInt(0), // safeTxGas will be auto-estimated
	)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to execute Safe transaction: %w", err)
	}

	return txHash, nil
}

// SplitPositionForSafe splits collateral into conditional tokens for a Safe wallet
func (b *ContractInterface) SplitPositionForSafe(
	ctx context.Context,
	safeSigner signer.SafeTradingSigner,
	chainID *big.Int,
	conditionId [32]byte,
	partition []*big.Int,
	amount *big.Int,
) (common.Hash, error) {
	safeAddr, err := b.GetSafeAddress(safeSigner.GetAddress())
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get Safe address: %w", err)
	}

	// Prepare calldata for splitPosition
	parsedABI, err := conditional_tokens.ConditionalTokensMetaData.GetAbi()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to parse ConditionalTokens ABI: %w", err)
	}
	parentCollectionId := [32]byte{} // empty for root collection
	calldata, err := parsedABI.Pack("splitPosition", b.contractConfig.Collateral, parentCollectionId, conditionId, partition, amount)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack splitPosition calldata: %w", err)
	}

	// Execute Safe transaction
	txHash, err := b.ExecuteTransactionBySafeAndSingleSigner(
		safeSigner, chainID, safeAddr,
		b.contractConfig.ConditionalTokens,
		big.NewInt(0), // value
		calldata,
		SafeOperationCall,
		big.NewInt(0), // safeTxGas will be auto-estimated
	)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to execute Safe transaction: %w", err)
	}

	return txHash, nil
}

// MergePositionsForSafe merges conditional tokens back into collateral for a Safe wallet
func (b *ContractInterface) MergePositionsForSafe(
	ctx context.Context,
	safeSigner signer.SafeTradingSigner,
	chainID *big.Int,
	conditionId [32]byte,
	partition []*big.Int,
	amount *big.Int,
) (common.Hash, error) {
	safeAddr, err := b.GetSafeAddress(safeSigner.GetAddress())
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get Safe address: %w", err)
	}

	// Prepare calldata for mergePositions
	parsedABI, err := conditional_tokens.ConditionalTokensMetaData.GetAbi()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to parse ConditionalTokens ABI: %w", err)
	}
	parentCollectionId := [32]byte{} // empty for root collection
	calldata, err := parsedABI.Pack("mergePositions", b.contractConfig.Collateral, parentCollectionId, conditionId, partition, amount)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack mergePositions calldata: %w", err)
	}

	// Execute Safe transaction
	txHash, err := b.ExecuteTransactionBySafeAndSingleSigner(
		safeSigner, chainID, safeAddr,
		b.contractConfig.ConditionalTokens,
		big.NewInt(0), // value
		calldata,
		SafeOperationCall,
		big.NewInt(0), // safeTxGas will be auto-estimated
	)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to execute Safe transaction: %w", err)
	}

	return txHash, nil
}

// SplitPositionNegRiskForEOA splits NegRisk market positions using EOA
func (b *ContractInterface) SplitPositionNegRiskForEOA(
	ctx context.Context,
	eoaSigner signer.EOATradingSigner,
	conditionId [32]byte,
	amount *big.Int,
) (common.Hash, error) {
	txSender := b.getTxSender()

	// Prepare calldata for splitPosition0 (NegRisk)
	parsedABI, err := negriskadapter.NegRiskAdapterMetaData.GetAbi()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to parse NegRiskAdapter ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("splitPosition0", conditionId, amount)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack splitPosition0 calldata: %w", err)
	}

	// Send transaction
	txHash, err := txSender.SendEthereumTransaction(b.contractConfig.NegRiskAdapter, calldata, big.NewInt(0))
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to send NegRisk split transaction: %w", err)
	}

	return txHash, nil
}

// MergePositionsNegRiskForEOA merges NegRisk market positions using EOA
func (b *ContractInterface) MergePositionsNegRiskForEOA(
	ctx context.Context,
	eoaSigner signer.EOATradingSigner,
	conditionId [32]byte,
	amount *big.Int,
) (common.Hash, error) {
	txSender := b.getTxSender()

	// Prepare calldata for mergePositions0 (NegRisk)
	parsedABI, err := negriskadapter.NegRiskAdapterMetaData.GetAbi()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to parse NegRiskAdapter ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("mergePositions0", conditionId, amount)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack mergePositions0 calldata: %w", err)
	}

	// Send transaction
	txHash, err := txSender.SendEthereumTransaction(b.contractConfig.NegRiskAdapter, calldata, big.NewInt(0))
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to send NegRisk merge transaction: %w", err)
	}

	return txHash, nil
}

// SplitPositionNegRiskForSafe splits NegRisk market positions using Safe
func (b *ContractInterface) SplitPositionNegRiskForSafe(
	ctx context.Context,
	safeSigner signer.SafeTradingSigner,
	chainID *big.Int,
	conditionId [32]byte,
	amount *big.Int,
) (common.Hash, error) {
	safeAddr, err := b.GetSafeAddress(safeSigner.GetAddress())
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get Safe address: %w", err)
	}

	// Prepare calldata for splitPosition0 (NegRisk)
	parsedABI, err := negriskadapter.NegRiskAdapterMetaData.GetAbi()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to parse NegRiskAdapter ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("splitPosition0", conditionId, amount)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack splitPosition0 calldata: %w", err)
	}

	// Execute Safe transaction
	txHash, err := b.ExecuteTransactionBySafeAndSingleSigner(
		safeSigner, chainID, safeAddr,
		b.contractConfig.NegRiskAdapter,
		big.NewInt(0), // value
		calldata,
		SafeOperationCall,
		big.NewInt(0), // safeTxGas will be auto-estimated
	)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to execute Safe transaction: %w", err)
	}

	return txHash, nil
}

// MergePositionsNegRiskForSafe merges NegRisk market positions using Safe
func (b *ContractInterface) MergePositionsNegRiskForSafe(
	ctx context.Context,
	safeSigner signer.SafeTradingSigner,
	chainID *big.Int,
	conditionId [32]byte,
	amount *big.Int,
) (common.Hash, error) {
	safeAddr, err := b.GetSafeAddress(safeSigner.GetAddress())
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get Safe address: %w", err)
	}

	// Prepare calldata for mergePositions0 (NegRisk)
	parsedABI, err := negriskadapter.NegRiskAdapterMetaData.GetAbi()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to parse NegRiskAdapter ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("mergePositions0", conditionId, amount)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack mergePositions0 calldata: %w", err)
	}

	// Execute Safe transaction
	txHash, err := b.ExecuteTransactionBySafeAndSingleSigner(
		safeSigner, chainID, safeAddr,
		b.contractConfig.NegRiskAdapter,
		big.NewInt(0), // value
		calldata,
		SafeOperationCall,
		big.NewInt(0), // safeTxGas will be auto-estimated
	)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to execute Safe transaction: %w", err)
	}

	return txHash, nil
}

// BuildSafeTransactionTypedData builds the typed data for Gnosis Safe transaction
// This follows the standard Gnosis Safe EIP-712 SafeTx structure
func BuildSafeTransactionTypedData(chainID *big.Int, safeAddr common.Address, to common.Address, value *big.Int, data []byte, operation SafeOperation, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, nonce *big.Int) eip712.TypedData {
	return eip712.TypedData{
		Types: eip712.Types{
			"EIP712Domain": []eip712.Type{
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"SafeTx": []eip712.Type{
				{Name: "to", Type: "address"},
				{Name: "value", Type: "uint256"},
				{Name: "data", Type: "bytes"},
				{Name: "operation", Type: "uint8"},
				{Name: "safeTxGas", Type: "uint256"},
				{Name: "baseGas", Type: "uint256"},
				{Name: "gasPrice", Type: "uint256"},
				{Name: "gasToken", Type: "address"},
				{Name: "refundReceiver", Type: "address"},
				{Name: "nonce", Type: "uint256"},
			},
		},
		PrimaryType: "SafeTx",
		Domain: eip712.TypedDataDomain{
			ChainId:           chainID.String(),
			VerifyingContract: strings.ToLower(safeAddr.Hex()),
		},
		Message: eip712.TypedDataMessage{
			"to":             to.Hex(),
			"value":          value.String(),
			"data":           fmt.Sprintf("0x%x", data),
			"operation":      fmt.Sprintf("%d", operation),
			"safeTxGas":      safeTxGas.String(),
			"baseGas":        baseGas.String(),
			"gasPrice":       gasPrice.String(),
			"gasToken":       gasToken.Hex(),
			"refundReceiver": refundReceiver.Hex(),
			"nonce":          nonce.String(),
		},
	}
}

// BuildCreateProxyTypedData builds the typed data for creating proxy of Safe
func BuildCreateProxyTypedData(chainID *big.Int, safeFactory, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) eip712.TypedData {
	return eip712.TypedData{
		Types: eip712.Types{
			"EIP712Domain": []eip712.Type{
				{Name: "name", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"CreateProxy": []eip712.Type{
				{Name: "paymentToken", Type: "address"},
				{Name: "payment", Type: "uint256"},
				{Name: "paymentReceiver", Type: "address"},
			},
		},
		PrimaryType: "CreateProxy",
		Domain: eip712.TypedDataDomain{
			Name:              constants.EIP712_GNOSIS_SAFE_FACTORY_DOMAIN_NAME,
			ChainId:           chainID.String(),
			VerifyingContract: strings.ToLower(safeFactory.Hex()),
		},
		Message: eip712.TypedDataMessage{
			"paymentToken":    paymentToken.Hex(),
			"payment":         payment.String(),
			"paymentReceiver": paymentReceiver.Hex(),
		},
	}
}

// BuildClobAuthTypedData builds the typed data for CLOB authentication
func BuildClobAuthTypedData(signer common.Address, chainID *big.Int, timestamp int64, nonce int) eip712.TypedData {
	return eip712.TypedData{
		Types: eip712.Types{
			"EIP712Domain": []eip712.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
			},
			"ClobAuth": []eip712.Type{
				{Name: "address", Type: "address"},
				{Name: "timestamp", Type: "string"},
				{Name: "nonce", Type: "uint256"},
				{Name: "message", Type: "string"},
			},
		},
		PrimaryType: "ClobAuth",
		Domain: eip712.TypedDataDomain{
			Name:    constants.EIP712_CLOB_AUTH_DOMAIN_NAME,
			Version: constants.EIP712_CLOB_AUTH_DOMAIN_VERSION,
			ChainId: chainID.String(),
		},
		Message: eip712.TypedDataMessage{
			"address":   signer.Hex(),
			"timestamp": fmt.Sprintf("%d", timestamp),
			"nonce":     fmt.Sprintf("%d", nonce),
			"message":   "This message attests that I control the given wallet",
		},
	}
}
