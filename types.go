package polymarketcontracts

type SafeOperation uint8

const (
	SafeOperationCall         SafeOperation = 0
	SafeOperationDelegateCall SafeOperation = 1
)

// SignatureType represents the type of signature used for signing orders
type SignatureType int

// Signature type constants
const (
	SignatureTypeEOA            SignatureType = 0 // Externally Owned Account
	SignatureTypePolyProxy      SignatureType = 1 // Polymarket Proxy
	SignatureTypePolyGnosisSafe SignatureType = 2 // Polymarket Gnosis Safe
)

const COLLATERAL_TOKEN_DECIMALS = 6
const CONDITIONAL_TOKEN_DECIMALS = 6
