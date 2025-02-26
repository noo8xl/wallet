package tron

import (
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"wallet/config"
	"wallet/database"
	"wallet/lib/exceptions"
	"wallet/lib/models"

	pb "wallet/api"
)

type TronService struct {
	token   string
	network string
	db      database.DatabaseService
}

func init() {
	initTrxConfig()
}

func (s *TronService) CreateWallet(userId int64) *pb.WalletItem {
	// TODO : impl db
	s.db.InsertTonWallet(&models.TonWallet{})
	return &pb.WalletItem{CoinName: "trx", Address: "address_will_be_here", Balance: 0.0}
}

// GetTrxBalance -> get balance by wallet address
func (s *TronService) GetTrxBalance(addr string) *big.Float {

	var currentBalance *big.Float
	var reqString string
	var payload *strings.Reader

	currentBalance = new(big.Float)
	currentBalance.SetString("123")

	fmt.Println(addr)
	url := "https://api.shasta.trongrid.io/wallet/getaccountbalance"
	reqString = strings.Join([]string{"{\"account_identifier\":{\"address\":", "\"", addr, "\"}"}, "")

	payload = strings.NewReader("{\"account_identifier\":{\"address\":\"TFH9ufxh8CpYxa7W9LdA8Y1iHHptYErvb7\"}")

	fmt.Println("payload =>\n", payload)
	fmt.Println("payload =>\n", reqString)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		exceptions.HandleAnException("<trx create wallet> got an err: " + err.Error())

	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

	// return response.balance / 1000000 // dividing on 1_000_000

	return currentBalance

}

// ===========================================================================================//
// ============================== function for internal usage ================================//
// ===========================================================================================//

func initTrxConfig() *TronService {
	var conf = new(TronService)
	conf.token = config.GetTronAPIKey()
	conf.network = "https://api.trongrid.io" // mainnet
	// conf.netwotk = "https://api.shasta.trongrid.io" // testnet
	//
	return &TronService{token: conf.token, network: conf.network}
}
