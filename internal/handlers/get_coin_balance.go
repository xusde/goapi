package handlers

import (
	"net/http"

	"github.com/gorilla/schema"
	"github.com/xusde/goapi/api"
)

func GetCoinBalance(w http.ResponseWriter, r*http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
}