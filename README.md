# Polymarket Go Contracts

A Go library for interacting with Polymarket smart contracts on Polygon network. This library provides a convenient interface for trading on Polymarket using both EOA (Externally Owned Account) and Gnosis Safe wallets.

## Features

- 🔐 Support for both EOA and Gnosis Safe signers
- 🔄 Complete trading operations (enable trading, split, merge, redeem)
- 🏭 Safe deployment and management
- 🔌 MPC wallet integration (Cobo MPC)
- ⚡ Built on go-ethereum

## Installation

```bash
go get github.com/ivanzzeth/polymarket-go-contracts
```

## Quick Start

### EOA Signer (Direct Private Key)

EOA signer allows you to interact with Polymarket directly using a private key. This is the simplest approach for individual accounts.

```go
package main

import (
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ivanzzeth/ethsig"
    polymarketcontracts "github.com/ivanzzeth/polymarket-go-contracts"
)

func main() {
    // Connect to Polygon RPC
    client, _ := ethclient.Dial(os.Getenv("RPC_URL"))

    // Load private key
    privateKey, _ := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))

    // Create EOA signer
    eoaSigner := ethsig.NewEthPrivateKeySigner(privateKey)

    // Initialize contract interface
    config := polymarketcontracts.GetContractConfig(big.NewInt(137)) // Polygon mainnet
    polymarketInterface, _ := polymarketcontracts.NewContractInterface(
        client,
        polymarketcontracts.WithContractConfig(config),
        polymarketcontracts.WithEOASigner(eoaSigner),
    )

    // Enable trading
    ctx := context.Background()
    txHashes, _ := polymarketInterface.EnableTrading(ctx)
}
```

### Safe Signer (Gnosis Safe Single-Owner)

Safe signer uses Gnosis Safe with a single owner for fund isolation and simplified protocol interaction. Transactions are signed using EIP-712 typed data signatures and executed through the Safe proxy contract.

```go
package main

import (
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    polymarketcontracts "github.com/ivanzzeth/polymarket-go-contracts"
    "github.com/ivanzzeth/polymarket-go-contracts/signer"
)

func main() {
    // Connect to Polygon RPC
    client, _ := ethclient.Dial(os.Getenv("RPC_URL"))
    chainID, _ := client.ChainID(context.Background())

    // Load private key
    privateKey, _ := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))

    // Create Safe signer
    safeSigner, _ := signer.NewSafeTradingPrivateKeySigner(chainID, client, privateKey)

    // Initialize contract interface
    config := polymarketcontracts.GetContractConfig(big.NewInt(137))
    polymarketInterface, _ := polymarketcontracts.NewContractInterface(
        client,
        polymarketcontracts.WithContractConfig(config),
        polymarketcontracts.WithSafeSigner(safeSigner),
    )

    // Deploy Safe (first time only)
    safeProxy, txHash, _ := polymarketInterface.DeploySafe()

    // Enable trading through Safe
    ctx := context.Background()
    txHashes, _ := polymarketInterface.EnableTrading(ctx)
}
```

## EOA vs Safe: Key Differences

| Feature | EOA Signer | Safe Signer |
|---------|-----------|-------------|
| **Setup** | Direct private key usage | Requires Safe deployment |
| **Owner Model** | Single private key | Single owner (via Safe proxy) |
| **Transaction Cost** | Lower gas costs | Higher gas costs (Safe proxy) |
| **Signing Method** | Standard transaction signing | EIP-712 typed data signatures |
| **Primary Benefit** | Simplicity and low cost | Fund isolation & protocol abstraction |
| **Use Case** | Direct wallet interaction | Isolated trading accounts |
| **Recovery** | No built-in recovery | Owner can be changed if needed |

### When to Use EOA Signer

- Personal trading accounts
- Quick prototyping and testing
- Lower transaction costs are important
- Direct wallet control is preferred

### When to Use Safe Signer

- **Fund Isolation**: Separate trading funds from main wallet
- **Protocol Abstraction**: Simplified interaction with Polymarket contracts
- **MPC Integration**: Use with MPC wallets (e.g., Cobo) for institutional security
- **Account Management**: Easier to change ownership or upgrade
- **Professional Trading**: Isolated accounts for different trading strategies

## Environment Variables

```bash
# Required for all examples
RPC_URL=https://polygon-rpc.com
PRIVATE_KEY=your_private_key_hex

# Required for MPC examples
COBO_PRIVATE_KEY=your_cobo_api_private_key
COBO_WALLET_ID=your_cobo_wallet_id
COBO_ADDRESS=your_mpc_wallet_address
COBO_ENV=dev  # or 'prod'
```

## Examples

For complete, runnable examples, see the [`examples/`](./examples) directory:

### EOA Examples

- **[`enable_trading`](./examples/enable_trading)** - Enable trading for an EOA account
  - Direct USDC and CTF token approvals
  - Simplest setup for individual traders

### Safe Examples with Private Key

- **[`deploy_safe`](./examples/deploy_safe)** - Deploy a new Gnosis Safe (single-owner)
  - One-time setup for isolated trading account
  - Creates a Safe proxy with single owner for fund isolation

- **[`safe_enable_trading`](./examples/safe_enable_trading)** - Enable trading for a Safe wallet
  - Configures all necessary token approvals through Safe
  - Owner signs transactions via EIP-712

### Safe Examples with MPC (Cobo)

- **[`deploy_safe_for_mpc`](./examples/deploy_safe_for_mpc)** - Deploy Safe using Cobo MPC wallet
  - Single-owner Safe with MPC wallet as owner
  - Institutional-grade key management

- **[`enable_trading_for_mpc_safe`](./examples/enable_trading_for_mpc_safe)** - Enable trading for MPC Safe
  - Complete MPC-based Safe trading setup
  - Combines Safe fund isolation with MPC security

## Running Examples

```bash
# Navigate to an example directory
cd examples/enable_trading

# Set environment variables
export RPC_URL=https://polygon-rpc.com
export PRIVATE_KEY=your_private_key_hex

# Run the example
go run main.go
```

## Core Operations

### Enable Trading

Sets up all required token approvals for trading on Polymarket:
- USDC approvals for Exchange, NegRiskAdapter, and NegRiskExchange
- CTF token approvals for the same contracts

```go
txHashes, err := polymarketInterface.EnableTrading(ctx)
```

### Split Position

Split collateral tokens into conditional tokens:

```go
txHash, err := polymarketInterface.SplitPosition(ctx, conditionId, amount)
```

### Merge Positions

Merge conditional tokens back to collateral:

```go
txHash, err := polymarketInterface.MergePositions(ctx, conditionId, amount)
```

### Redeem Positions

Redeem conditional tokens after market resolution:

```go
txHash, err := polymarketInterface.RedeemPositions(ctx, conditionId, indexSets)
```

## Architecture

```
polymarket-go-contracts/
├── interface.go              # Main contract interface
├── config.go                 # Contract addresses and configs
├── types.go                  # Type definitions
├── signer/                   # Signing implementations
│   ├── eoa_trading_signer.go     # EOA signer interface
│   ├── safe_trading_signer.go    # Safe signer implementations
│   ├── mpc_remote_signer.go      # Cobo MPC integration
│   └── transaction_sender.go     # Transaction sending logic
├── sender/                   # Transaction sender interface
├── contracts/                # Generated contract bindings
└── examples/                 # Complete usage examples
```

## Supported Networks

- **Polygon Mainnet** (Chain ID: 137)
- **Polygon Amoy Testnet** (Chain ID: 80002)

Pre-configured contract addresses are available via `GetContractConfig()`.

## Security Considerations

- **Private Keys**: Never commit private keys to version control
- **Environment Variables**: Use `.env` files or secure secret management
- **Safe Wallets**: Recommended for institutional use and large funds
- **MPC Wallets**: Provide hardware-backed security for enterprise applications
- **Gas Estimation**: Always verify gas costs before mainnet deployment

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Support

For issues and questions:
- Open an issue on [GitHub](https://github.com/ivanzzeth/polymarket-go-contracts/issues)
- Check the [`examples/`](./examples) directory for working code

## Links

- [Polymarket](https://polymarket.com/)
- [Gnosis Safe](https://safe.global/)
- [Cobo MPC WaaS](https://www.cobo.com/)
