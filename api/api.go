package api

// coin balance params
type CoinBalanceParams struct {
	Username string
}

// coin balance response
type CoinBalanceResponse struct {
	Code int
	Balance float64
}

// error response if error occurs
type ErrorResponse struct {
	Code int
	Message string
}