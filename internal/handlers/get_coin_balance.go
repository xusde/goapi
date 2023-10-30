package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
	"github.com/xusde/goapi/api"
	"github.com/xusde/goapi/internal/tools"
)

func GetCoinBalance(w http.ResponseWriter, r*http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return 
	}
	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}
	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if (tokenDetails == nil) {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	var resp = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code: http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}