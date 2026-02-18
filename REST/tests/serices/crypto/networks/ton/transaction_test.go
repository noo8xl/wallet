package theopennetwork_test

import "testing"

func TestSingleTransaction(t *testing.T) {
	var hash string
	var err error
	TEST_SINGLE_TSX_DTO.Payee.PeyeeAddress = TEST_PEYEE_ADDRESS
	TEST_SINGLE_TSX_DTO.Beneficiar.BeneficiarAddress = TEST_BENEFICIAR_ADDRESS

	hash, err = TEST_SVC.DefineaTypeAndSendSingleTransaction(TEST_SINGLE_TSX_DTO)
	if err != nil {
		t.Fatalf("TestSingleTransaction expected non-empty string; got an error: %v", err)
	}

	// test one-time here
	// hash, err = TEST_SVC.DefineaTypeAndSendSingleTransaction(TEST_SINGLE_TSX_DTO, 1)
	// if err != nil {
	// 	t.Fatalf("TestSingleTransaction expected non-empty string; got an error: %v", err)
	// }

	t.Logf("TestSingleTransaction transaction hash -> %v", hash)
}

// func TestMultTransaction(t *testing.T) {

// }
