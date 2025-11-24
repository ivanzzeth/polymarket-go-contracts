# Polymarket Go Contracts SDK Examples

This directory contains example programs demonstrating how to use the Polymarket Go Contracts SDK.

## Prerequisites

Before running any examples, you need to:

1. Set up environment variables:
   ```bash
   export RPC_URL="https://polygon-rpc.com"  # Or your preferred Polygon RPC endpoint
   export PRIVATE_KEY="your_private_key_hex"  # Without 0x prefix
   ```

2. Ensure you have Go installed (version 1.19 or higher)

## Available Examples

### 1. Enable Trading (`enable_trading`)

Enables trading for your address by setting all required allowances for USDC and Conditional Tokens Framework (CTF) on Polymarket contracts.

**What it does:**
- Approves USDC spending for Exchange, NegRiskAdapter, and NegRiskExchange contracts
- Sets CTF approval for all three contracts
- Checks and displays current balance and allowance status

**Usage:**
```bash
cd examples/enable_trading
go run main.go
```

**Output:**
```
Address: 0x...
=== Balance and Allowance Status ===

USDC Balance: 1000.000000 USDC

Allowances:
  USDC → Exchange: ✅
  USDC → NegRiskAdapter: ✅
  USDC → NegRiskExchange: ✅
  CTF → Exchange: ✅
  CTF → NegRiskAdapter: ✅
  CTF → NegRiskExchange: ✅
```

### 2. Deploy Safe (`deploy_safe`)

Deploys a Gnosis Safe multi-signature wallet for your address on Polygon.

**What it does:**
- Creates an EIP-712 signature for Safe deployment
- Automatically estimates gas for optimal transaction costs
- Deploys a Safe proxy contract through the SafeProxyFactory
- Returns the deployed Safe address
- Verifies the deployment

**Usage:**
```bash
cd examples/deploy_safe
go run main.go
```

**Output:**
```
Deployer Address: 0x...
Safe Factory Address: 0x...
Deploying Safe contract...
✅ Safe deployed successfully!
   Transaction Hash: 0x...
   Safe Proxy Address: 0x...
   Owner Address: 0x...
   Computed Safe Address: 0x...
   ✅ Address verification successful!

You can now use this Safe address for multi-signature operations.
View on Polygonscan: https://polygonscan.com/address/0x...
```

### 3. Enable Trading for Safe (`safe_enable_trading`)

Enables trading for a Gnosis Safe wallet by setting all required allowances through Safe transactions.

**What it does:**
- Computes the Safe address for your EOA
- Verifies the Safe is already deployed
- Uses Safe's `execTransaction` to set USDC allowances for Exchange, NegRiskAdapter, and NegRiskExchange
- Uses Safe's `execTransaction` to set CTF approvals for all three contracts
- Automatically estimates gas for each Safe transaction
- Waits for each transaction confirmation before proceeding

**Prerequisites:**
- Your Safe must be deployed (use `deploy_safe` example first)
- The Safe should have USDC tokens for trading
- The Safe needs MATIC for gas fees

**Usage:**
```bash
cd examples/safe_enable_trading
go run main.go
```

**Output:**
```
EOA Address: 0x...
Safe Address: 0x...
✅ Safe is deployed

=== Enabling Trading for Safe ===
=== Safe Balance and Allowance Status ===
Safe Address: 0x...

USDC Balance: 1000.000000 USDC

Allowances:
  USDC → Exchange: ❌
  USDC → NegRiskAdapter: ❌
  USDC → NegRiskExchange: ❌
  CTF → Exchange: ❌
  CTF → NegRiskAdapter: ❌
  CTF → NegRiskExchange: ❌

=== Setting Allowances via Safe ===
Setting USDC → Exchange allowance...
  ✅ Transaction submitted: 0x...
Setting USDC → NegRiskAdapter allowance...
  ✅ Transaction submitted: 0x...
Setting USDC → NegRiskExchange allowance...
  ✅ Transaction submitted: 0x...
Setting CTF → Exchange approval...
  ✅ Transaction submitted: 0x...
Setting CTF → NegRiskAdapter approval...
  ✅ Transaction submitted: 0x...
Setting CTF → NegRiskExchange approval...
  ✅ Transaction submitted: 0x...

✅ All allowances are set! Trading is enabled for Safe.

✅ Trading successfully enabled for Safe!
Safe can now trade on Polymarket: 0x...
View on Polygonscan: https://polygonscan.com/address/0x...
```

**Notes:**
- Each allowance requires a separate Safe transaction with signature
- The process may take several minutes as each transaction needs confirmation
- Gas is automatically estimated for optimal costs
- If an allowance is already set, it will be skipped

## Network Configuration

The examples are configured for **Polygon Mainnet (Chain ID: 137)** by default.

Contract addresses used:
- Exchange: `0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E`
- NegRiskAdapter: `0xd91E80cF2E7be2e162c6513ceD06f1dD0dA35296`
- NegRiskExchange: `0xC5d563A36AE78145C45a50134d48A1215220f80a`
- Collateral (USDC): `0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174`
- ConditionalTokens: `0x4D97DCd97eC945f40cF65F87097ACe5EA0476045`
- SafeProxyFactory: `0xaacFeEa03eb1561C4e67d661e40682Bd20E3541b`

To use **Polygon Amoy Testnet (Chain ID: 80002)**, modify the code:
```go
polygonChainID := big.NewInt(80002)
```

## Building Examples

To build all examples:
```bash
go build ./examples/...
```

To build a specific example:
```bash
cd examples/enable_trading
go build
./enable_trading
```

## Security Notes

⚠️ **Never commit your private keys to version control!**

- Always use environment variables or secure key management systems
- The private key should be provided without the `0x` prefix
- For production use, consider using hardware wallets or more secure key management solutions

## Troubleshooting

### "Failed to dial ethclient"
- Verify your `RPC_URL` is correct and accessible
- Check your network connection
- Try using a different RPC endpoint

### "Failed to parse privateKey"
- Ensure the private key is in hex format without `0x` prefix
- Verify the key is exactly 64 hex characters (32 bytes)

### "Insufficient funds"
- Ensure your address has enough MATIC for gas fees
- For `enable_trading`: You also need USDC tokens for trading

### Transaction timeout
- Increase the context timeout in the code
- Check network congestion and gas prices
- Verify the transaction on Polygonscan

## Resources

- [Polymarket Documentation](https://docs.polymarket.com/)
- [Gnosis Safe Documentation](https://docs.safe.global/)
- [Polygon Network](https://polygon.technology/)
- [Go Ethereum Documentation](https://geth.ethereum.org/docs/)

## License

See the main repository LICENSE file.
