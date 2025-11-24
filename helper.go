package polymarketcontracts

import (
	"fmt"
	"math/big"
)

// Helper functions

// formatUSDC formats USDC amount (6 decimals) to a readable string
func formatUSDC(amount *big.Int) string {
	if amount == nil {
		return "0.000000"
	}

	divisor := big.NewInt(1000000) // 10^6
	quotient := new(big.Int).Div(amount, divisor)
	remainder := new(big.Int).Mod(amount, divisor)

	return fmt.Sprintf("%s.%06d", quotient.String(), remainder.Uint64())
}

// checkmark returns a checkmark or cross based on the condition
func checkmark(condition bool) string {
	if condition {
		return "✅"
	}
	return "❌"
}
