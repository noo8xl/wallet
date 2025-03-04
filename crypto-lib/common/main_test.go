package common_test

import (
	"log"
	"testing"
	"wallet/crypto-lib/common"
)

func TestInitService(t *testing.T) {
	s := common.InitService() // len is 4
	if s == nil {
		t.Errorf("got an err at initialization")
	}

	log.Printf("btc service is -> %v ", s)

}
