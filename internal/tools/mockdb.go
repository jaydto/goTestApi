package tools

import (
	"time"
)

type MockDB struct {
}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {AuthToken: "1234w", Username: "alex"},
	"mike": {AuthToken: "1234a", Username: "mike"},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {Coins: 100, Username: "alex"},
	"mike": {Coins: 20000, Username: "mike"},
}

func (d *MockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData

}

func (d *MockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]

	if !ok {
		return nil
	}
	return &clientData
}

func (d *MockDB) SetupDatabases() error {
	return nil
}
