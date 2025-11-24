package constants

// Authentication endpoints
const (
	EndpointAuthAPIKey       = "/auth/api-key"
	EndpointAuthDeriveAPIKey = "/auth/derive-api-key"
)

// Health and system endpoints
const (
	EndpointRoot       = "/"
	EndpointServerTime = "/time"
)

// Order management endpoints
const (
	EndpointOrder              = "/order"
	EndpointOrders             = "/orders"
	EndpointOpenOrders         = "/data/orders"
	EndpointGetOrder           = "/data/order/"
	EndpointCancelAll          = "/cancel-all"
	EndpointCancelMarketOrders = "/cancel-market-orders"
	EndpointOrderScoring       = "/orders/scoring"
	EndpointOrderByID          = "/order/%s"
	EndpointOrderScoring2      = "/order/%s/scoring"
)

// Market data endpoints
const (
	EndpointMidpoint       = "/midpoint"
	EndpointSpreads        = "/spreads"
	EndpointPrice          = "/price"
	EndpointPrices         = "/prices"
	EndpointBook           = "/book"
	EndpointBooks          = "/books"
	EndpointMarkets        = "/markets"
	EndpointMarketByID     = "/markets/%s"
	EndpointTickSize       = "/tick-size"
	EndpointPriceHistory   = "/prices-history"
	EndpointLastTradePrice = "/last-trade-price"
	EndpointNegRisk        = "/neg-risk"
)

// Trade and fill endpoints
const (
	EndpointTrades = "/data/trades"
	EndpointFills  = "/fills"
)

// Balance endpoints
const (
	EndpointBalanceAllowance = "/balance-allowance"
)

// Notification endpoints
const (
	EndpointNotifications = "/notifications"
)
