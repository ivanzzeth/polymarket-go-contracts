package polymarket

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	client         ethclient.EthClientInterface
	contractConfig *ContractConfig
	signatureType  SignatureType
	txSender       TransactionSender

	collateralContract        *erc20.Erc20
	conditionalTokensContract *conditional_tokens.ConditionalTokens
	exchangeContract          *exchange.Exchange
	exchangeFeesContract      *exchangefees.ExchangeFees
	negRiskAdapterContract    *negriskadapter.NegRiskAdapter
	negRiskContract           *negrisk.NegRisk
	negRiskFeesContract       *negriskfees.NegRiskFees
	safeProxyFactoryContract  *safeproxyfactory.SafeProxyFactory
}

type ContractInterfaceConfig struct {
	SignatureType SignatureType
	TxSender      TransactionSender

	ContractConfig *ContractConfig
}

type ContractInterfaceOption func(c *ContractInterfaceConfig)

func NewContractInterface(
	client ethclient.EthClientInterface,
	options ...ContractInterfaceOption,
) (*ContractInterface, error) {
	defaultOptions := &ContractInterfaceConfig{
		SignatureType:  SignatureTypeEOA,
		ContractConfig: MATIC_CONTRACTS,
	}

	for _, opFn := range options {
		opFn(defaultOptions)
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
		client:         client,
		contractConfig: defaultOptions.ContractConfig,
		signatureType:  defaultOptions.SignatureType,
		txSender:       defaultOptions.TxSender,

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

// CheckBalanceAndAllowance checks the USDC balance and all allowances for the given address
func (b *ContractInterface) CheckBalanceAndAllowance(ctx context.Context, address common.Address) (*BalanceAllowanceInfo, error) {
	info := &BalanceAllowanceInfo{}

	// Check USDC balance
	balance, err := b.collateralContract.BalanceOf(&bind.CallOpts{Context: ctx}, address)
	if err != nil {
		return nil, fmt.Errorf("failed to get USDC balance: %w", err)
	}
	info.Balance = balance

	// Check USDC allowances
	allowanceExchange, err := b.collateralContract.Allowance(&bind.CallOpts{Context: ctx}, address, b.contractConfig.Exchange)
	if err != nil {
		return nil, fmt.Errorf("failed to get USDC allowance for Exchange: %w", err)
	}
	info.AllowanceExchange = allowanceExchange

	allowanceNegRiskAdapter, err := b.collateralContract.Allowance(&bind.CallOpts{Context: ctx}, address, b.contractConfig.NegRiskAdapter)
	if err != nil {
		return nil, fmt.Errorf("failed to get USDC allowance for NegRiskAdapter: %w", err)
	}
	info.AllowanceNegRiskAdapter = allowanceNegRiskAdapter

	allowanceNegRiskExchange, err := b.collateralContract.Allowance(&bind.CallOpts{Context: ctx}, address, b.contractConfig.NegRiskExchange)
	if err != nil {
		return nil, fmt.Errorf("failed to get USDC allowance for NegRiskExchange: %w", err)
	}
	info.AllowanceNegRiskExchange = allowanceNegRiskExchange

	// Check CTF approvals
	ctfApprovedExchange, err := b.conditionalTokensContract.IsApprovedForAll(&bind.CallOpts{Context: ctx}, address, b.contractConfig.Exchange)
	if err != nil {
		return nil, fmt.Errorf("failed to check CTF approval for Exchange: %w", err)
	}
	info.CTFApprovedExchange = ctfApprovedExchange

	ctfApprovedNegRiskAdapter, err := b.conditionalTokensContract.IsApprovedForAll(&bind.CallOpts{Context: ctx}, address, b.contractConfig.NegRiskAdapter)
	if err != nil {
		return nil, fmt.Errorf("failed to check CTF approval for NegRiskAdapter: %w", err)
	}
	info.CTFApprovedNegRiskAdapter = ctfApprovedNegRiskAdapter

	ctfApprovedNegRiskExchange, err := b.conditionalTokensContract.IsApprovedForAll(&bind.CallOpts{Context: ctx}, address, b.contractConfig.NegRiskExchange)
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
	return b.GetSafeProxyFactory().ComputeProxyAddress(nil, eoa)
}

func (b *ContractInterface) DeploySafeBySender(txSender TransactionSender, signer ethsig.TypedDataSigner) (safeProxy common.Address, txHash common.Hash, err error) {
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

func (b *ContractInterface) DeploySafeWithSig(txSender TransactionSender, chainID *big.Int, safeFactory, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address, createSig safeproxyfactory.SafeProxyFactorySig) (safeProxy common.Address, txHash common.Hash, err error) {
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
	fmt.Printf("DeploySafeWithSig for address %v\n", addr.Hex())

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

	safeProxy, err = b.GetSafeAddress(addr)
	return
}

type TransactionSigner interface {
	ethsig.TransactionSigner
	ethsig.AddressGetter
}

func (b *ContractInterface) GetTransactionSenderrByTransactionSigner(client ethclient.EthClientInterface, txSigner TransactionSigner) (TransactionSender, error) {
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	return &TransactionSenderrByTransactionSigner{
		chainId:  chainId,
		client:   client,
		txSigner: txSigner,
	}, nil
}

type TransactionSenderrByTransactionSigner struct {
	chainId  *big.Int
	client   ethclient.EthClientInterface
	txSigner TransactionSigner
}

func (b *TransactionSenderrByTransactionSigner) SendEthereumTransaction(to common.Address, data []byte, value *big.Int) (common.Hash, error) {
	ctx := context.Background()

	// Get nonce
	nonce, err := b.client.NonceAt(ctx, b.txSigner.GetAddress(), nil)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get nonce: %w", err)
	}

	// Get gas price
	gasPrice, err := b.client.SuggestGasPrice(ctx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get gas price: %w", err)
	}

	// Estimate gas limit
	msg := ethereum.CallMsg{
		From:     b.txSigner.GetAddress(),
		To:       &to,
		Value:    value,
		Data:     data,
		GasPrice: gasPrice,
	}
	gasLimit, err := b.client.EstimateGas(ctx, msg)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to estimate gas: %w", err)
	}

	// Create and sign the transaction using BoundContract
	// This is a workaround since we need to create a raw transaction
	// We'll use the ethereum types directly
	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, data)

	// Sign the transaction
	signedTx, err := b.txSigner.SignTransactionWithChainID(tx, b.chainId)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Send the signed transaction
	err = b.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to send transaction: %w", err)
	}

	return signedTx.Hash(), nil
}

func (b *ContractInterface) GetTransactionSenderrByCoboMpcTransactionSender(mpcSigner *signer.CoboMpcSigner) (TransactionSender, error) {
	return &CoboMpcTransactionSender{signer: mpcSigner}, nil
}

type CoboMpcTransactionSender struct {
	signer *signer.CoboMpcSigner
}

func (s *CoboMpcTransactionSender) SendEthereumTransaction(to common.Address, data []byte, value *big.Int) (common.Hash, error) {
	txResp, err := s.signer.CallContract(to.Hex(), fmt.Sprintf("0x%x", data), value.String(), nil)
	if err != nil {
		return common.Hash{}, err
	}

	txDetail, err := s.signer.WaitTransactionStatus(txResp.TransactionId, cobo_waas2.TRANSACTIONSTATUS_CONFIRMING, 100)
	if err != nil {
		return common.Hash{}, err
	}

	txHash := common.HexToHash(*txDetail.TransactionHash)
	return txHash, nil
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

func (b *ContractInterface) ExecuteTransactionBySafe(txSender TransactionSender, safeAddr common.Address, to common.Address, value *big.Int, data []byte, operation SafeOperation, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (common.Hash, error) {
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

// SafeTradingSigner defines the interface requirements for Safe trading operations
type SafeTradingSigner interface {
	ethsig.AddressGetter
	ethsig.TypedDataSigner
	TransactionSender
}

// EnableTradingForSafe enables trading for a Safe wallet by setting all required allowances
// This function uses ExecuteTransactionBySafeAndSingleSigner to execute all approval operations
func (b *ContractInterface) EnableTradingForSafe(
	ctx context.Context,
	safeSigner SafeTradingSigner,
	// txSigner signer.TypedDataSigner,
	chainID *big.Int,
	safeAddr common.Address,
) error {
	// Check current status
	info, err := b.CheckBalanceAndAllowance(ctx, safeAddr)
	if err != nil {
		return fmt.Errorf("failed to check balance and allowance: %w", err)
	}

	// Print current status
	fmt.Println("=== Safe Balance and Allowance Status ===")
	fmt.Printf("Safe Address: %s\n", safeAddr.Hex())
	if err := b.PrintBalanceAndAllowance(ctx, safeAddr); err != nil {
		return err
	}

	// Maximum allowance for ERC20 approvals
	maxAllowance := new(big.Int)
	maxAllowance.SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)

	// Parse ERC20 ABI for approve function
	erc20ABI, err := abi.JSON(strings.NewReader(erc20.Erc20MetaData.ABI))
	if err != nil {
		return fmt.Errorf("failed to parse ERC20 ABI: %w", err)
	}

	// Parse ConditionalTokens ABI for setApprovalForAll function
	ctfABI, err := abi.JSON(strings.NewReader(conditional_tokens.ConditionalTokensMetaData.ABI))
	if err != nil {
		return fmt.Errorf("failed to parse ConditionalTokens ABI: %w", err)
	}

	fmt.Println("\n=== Setting Allowances via Safe ===")

	// Approve USDC for all contracts if needed
	if info.AllowanceExchange.Cmp(big.NewInt(0)) == 0 {
		fmt.Printf("Setting USDC → Exchange allowance...\n")
		approveData, err := erc20ABI.Pack("approve", b.contractConfig.Exchange, maxAllowance)
		if err != nil {
			return fmt.Errorf("failed to pack approve data for Exchange: %w", err)
		}

		txHash, err := b.ExecuteTransactionBySafeAndSingleSigner(
			safeSigner, chainID, safeAddr,
			b.contractConfig.Collateral, big.NewInt(0), approveData,
			SafeOperationCall, big.NewInt(0),
		)
		if err != nil {
			return fmt.Errorf("failed to execute Safe transaction for USDC → Exchange approval: %w", err)
		}
		fmt.Printf("  ✅ Transaction submitted: %s\n", txHash.Hex())
	}

	if info.AllowanceNegRiskAdapter.Cmp(big.NewInt(0)) == 0 {
		fmt.Printf("Setting USDC → NegRiskAdapter allowance...\n")
		approveData, err := erc20ABI.Pack("approve", b.contractConfig.NegRiskAdapter, maxAllowance)
		if err != nil {
			return fmt.Errorf("failed to pack approve data for NegRiskAdapter: %w", err)
		}

		txHash, err := b.ExecuteTransactionBySafeAndSingleSigner(
			safeSigner, chainID, safeAddr,
			b.contractConfig.Collateral, big.NewInt(0), approveData,
			SafeOperationCall, big.NewInt(0),
		)
		if err != nil {
			return fmt.Errorf("failed to execute Safe transaction for USDC → NegRiskAdapter approval: %w", err)
		}
		fmt.Printf("  ✅ Transaction submitted: %s\n", txHash.Hex())
	}

	if info.AllowanceNegRiskExchange.Cmp(big.NewInt(0)) == 0 {
		fmt.Printf("Setting USDC → NegRiskExchange allowance...\n")
		approveData, err := erc20ABI.Pack("approve", b.contractConfig.NegRiskExchange, maxAllowance)
		if err != nil {
			return fmt.Errorf("failed to pack approve data for NegRiskExchange: %w", err)
		}

		txHash, err := b.ExecuteTransactionBySafeAndSingleSigner(
			safeSigner, chainID, safeAddr,
			b.contractConfig.Collateral, big.NewInt(0), approveData,
			SafeOperationCall, big.NewInt(0),
		)
		if err != nil {
			return fmt.Errorf("failed to execute Safe transaction for USDC → NegRiskExchange approval: %w", err)
		}
		fmt.Printf("  ✅ Transaction submitted: %s\n", txHash.Hex())
	}

	// Approve CTF for all contracts if needed
	if !info.CTFApprovedExchange {
		fmt.Printf("Setting CTF → Exchange approval...\n")
		setApprovalData, err := ctfABI.Pack("setApprovalForAll", b.contractConfig.Exchange, true)
		if err != nil {
			return fmt.Errorf("failed to pack setApprovalForAll data for Exchange: %w", err)
		}

		txHash, err := b.ExecuteTransactionBySafeAndSingleSigner(
			safeSigner, chainID, safeAddr,
			b.contractConfig.ConditionalTokens, big.NewInt(0), setApprovalData,
			SafeOperationCall, big.NewInt(0),
		)
		if err != nil {
			return fmt.Errorf("failed to execute Safe transaction for CTF → Exchange approval: %w", err)
		}
		fmt.Printf("  ✅ Transaction submitted: %s\n", txHash.Hex())
	}

	if !info.CTFApprovedNegRiskAdapter {
		fmt.Printf("Setting CTF → NegRiskAdapter approval...\n")
		setApprovalData, err := ctfABI.Pack("setApprovalForAll", b.contractConfig.NegRiskAdapter, true)
		if err != nil {
			return fmt.Errorf("failed to pack setApprovalForAll data for NegRiskAdapter: %w", err)
		}

		txHash, err := b.ExecuteTransactionBySafeAndSingleSigner(
			safeSigner, chainID, safeAddr,
			b.contractConfig.ConditionalTokens, big.NewInt(0), setApprovalData,
			SafeOperationCall, big.NewInt(0),
		)
		if err != nil {
			return fmt.Errorf("failed to execute Safe transaction for CTF → NegRiskAdapter approval: %w", err)
		}
		fmt.Printf("  ✅ Transaction submitted: %s\n", txHash.Hex())
	}

	if !info.CTFApprovedNegRiskExchange {
		fmt.Printf("Setting CTF → NegRiskExchange approval...\n")
		setApprovalData, err := ctfABI.Pack("setApprovalForAll", b.contractConfig.NegRiskExchange, true)
		if err != nil {
			return fmt.Errorf("failed to pack setApprovalForAll data for NegRiskExchange: %w", err)
		}

		txHash, err := b.ExecuteTransactionBySafeAndSingleSigner(
			safeSigner, chainID, safeAddr,
			b.contractConfig.ConditionalTokens, big.NewInt(0), setApprovalData,
			SafeOperationCall, big.NewInt(0),
		)
		if err != nil {
			return fmt.Errorf("failed to execute Safe transaction for CTF → NegRiskExchange approval: %w", err)
		}
		fmt.Printf("  ✅ Transaction submitted: %s\n", txHash.Hex())
	}

	fmt.Println("\n✅ All allowances are set! Trading is enabled for Safe.")
	return nil
}

// RedeemPositionsForSafe redeems conditional tokens for a resolved market using TransactionSender
func (b *ContractInterface) RedeemPositionsForSafe(
	ctx context.Context,
	safeSigner signer.SafeTradingSigner,
	chainID *big.Int,
	safeAddr common.Address,
	conditionId [32]byte,
	indexSets []*big.Int,
) error {
	// Check USDC balance before redeem
	balanceBefore, err := b.collateralContract.BalanceOf(&bind.CallOpts{Context: ctx}, safeAddr)
	if err != nil {
		return fmt.Errorf("failed to get balance before redeem: %w", err)
	}

	// Prepare calldata for redeemPositions
	parsedABI, err := conditional_tokens.ConditionalTokensMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to parse ConditionalTokens ABI: %w", err)
	}
	parentCollectionId := [32]byte{} // empty for root collection
	calldata, err := parsedABI.Pack("redeemPositions", b.contractConfig.Collateral, parentCollectionId, conditionId, indexSets)
	if err != nil {
		return fmt.Errorf("failed to pack redeemPositions calldata: %w", err)
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
		return fmt.Errorf("failed to execute Safe transaction: %w", err)
	}

	fmt.Printf("✅ Redeem transaction submitted: %s\n", txHash.Hex())

	// Check USDC balance after redeem
	balanceAfter, err := b.collateralContract.BalanceOf(&bind.CallOpts{Context: ctx}, safeAddr)
	if err != nil {
		return fmt.Errorf("failed to get balance after redeem: %w", err)
	}

	// Calculate the increase
	balanceIncrease := new(big.Int).Sub(balanceAfter, balanceBefore)
	fmt.Printf("✅ Redeemed %s USDC\n", formatUSDC(balanceIncrease))

	return nil
}

// RedeemPositionsNegRiskForSafe redeems NegRisk market positions using TransactionSender
func (b *ContractInterface) RedeemPositionsNegRiskForSafe(
	ctx context.Context,
	safeSigner signer.SafeTradingSigner,
	chainID *big.Int,
	safeAddr common.Address,
	conditionId [32]byte,
	amounts []*big.Int,
) error {
	// Check USDC balance before redeem
	balanceBefore, err := b.collateralContract.BalanceOf(&bind.CallOpts{Context: ctx}, safeAddr)
	if err != nil {
		return fmt.Errorf("failed to get balance before redeem: %w", err)
	}

	// Prepare calldata for redeemPositions (NegRisk)
	parsedABI, err := negriskadapter.NegRiskAdapterMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to parse NegRiskAdapter ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("redeemPositions", conditionId, amounts)
	if err != nil {
		return fmt.Errorf("failed to pack redeemPositions calldata: %w", err)
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
		return fmt.Errorf("failed to execute Safe transaction: %w", err)
	}

	fmt.Printf("✅ NegRisk redeem transaction submitted: %s\n", txHash.Hex())

	// Check USDC balance after redeem
	balanceAfter, err := b.collateralContract.BalanceOf(&bind.CallOpts{Context: ctx}, safeAddr)
	if err != nil {
		return fmt.Errorf("failed to get balance after redeem: %w", err)
	}

	// Calculate the increase
	balanceIncrease := new(big.Int).Sub(balanceAfter, balanceBefore)
	fmt.Printf("✅ Redeemed %s USDC from NegRisk market\n", formatUSDC(balanceIncrease))

	return nil
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
			Name: constants.EIP712_GNOSIS_SAFE_FACTORY_DOMAIN_NAME,
			// NO VErSION
			// Version: constants.EIP712_GNOSIS_SAFE_FACTORY_DOMAIN_VERSION,
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
