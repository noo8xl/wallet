// Package tron -> is tron network interacting
package tron

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"wallet/config"
	"wallet/lib/exceptions"
	"wallet/lib/helpers"

	pb "wallet/api"
)

// https://developers.tron.network/reference/background#note
// -> the doc is here <-
//
// ValidateTrxAddress -> validate address in tron network
func (s *Service) ValidateTrxAddress(addr string) bool {

	// testNet + path + payload
	var response struct {
		Result  bool   `json:"result"`
		Message string `json:"message"`
	}
	path := "/wallet/validateaddress"
	payload := strings.NewReader(strings.Join([]string{"{\"address\":", "\"", addr, "\",", "\"visible\":true}"}, ""))
	url := strings.Join([]string{s.network, path}, "")

	// fmt.Println("payload ->", payload)
	// fmt.Println(" url -> ", url)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		exceptions.HandleAnException("<trx validate address> got an http err: " + err.Error())

	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	// fmt.Println(" res is => ", string(body))
	// fmt.Println("response is -> ", response.result)
	return response.Result
}

func (s *Service) SendSingleTransaction(dto *pb.SendSingleTsxRequest) string {

	key := config.GetAnEncryptionKey()
	encryptedPk := s.db.SelectTrxPrivate(dto.Payee.PeyeeAddress)
	privateKey, err := helpers.DecryptKey(key, encryptedPk)
	if err != nil {
		exceptions.HandleAnException("SendSingleTrxTransaction got an error: " + err.Error())
	}

	log.Printf("trx private key is -> %s", privateKey)
	// save a tsx details to db
	return "trx tsx hash"
}
