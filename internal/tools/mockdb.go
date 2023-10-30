package tools

import (
	"time"
)

type mockDB struct {}
var mockLoginDetails = map[string]LoginDetails {
	"alex": {
		AuthToken: "123ABC",
		Username: "alex",
	},
	"jason": {
		AuthToken: "456def",
		Username: "jason",
	},
	"marie": {
		AuthToken: "789ghi",
		Username: "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails {
	"alex": {
		Coins: 100.0,
		Username: "alex",
	}, 
	"jason": {
		Coins: 200.0,
		Username: "jason",
	},
	"marie": {
		Coins: 300.0,
		Username: "marie",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// simulate db latency
	time.Sleep(1 * time.Second)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// simulate db latency
	time.Sleep(1 * time.Second)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}
func (d *mockDB) SetupDatabase() error {
	return nil

}