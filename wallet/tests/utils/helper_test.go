package helpers

// var (
// 	strToEncrypt       = "f096feb38821ddd6035d7e2b142402483b678e582fb18f8229c65bd082cd3672"
// 	key                = config.GetAnEncryptionKey()
// 	byteSliceToEncrypt = []byte("0x25857891164D7EEB5BFF97D06FA5553055E3BC46B0D54CE8A8AF4BD328A7D37ADF27C674C3930A246BF33F32AC6C168BF4FD526C51051972F1F44B60327CBDBA")
// 	byteKey            = []byte(config.GetAnEncryptionKey())
// )

// func TestCheckNetworkConnection(t *testing.T) {
// 	err := CheckNetworkConnection()
// 	if err != nil {
// 		t.Errorf("CheckNetworkConnection got an error: %v", err.Error())
// 	}
// }

// func TestValidateArgs(t *testing.T) {
// 	err := ValidateArgs(-1, 3)
// 	if err != nil {
// 		t.Logf("tested with an error1")
// 		t.Logf("expected: err; got: %v", err.Error())
// 	}

// 	err = ValidateArgs(3, 3)
// 	if err != nil {
// 		t.Errorf("expected: nil; got: %v", err.Error())
// 	}
// 	t.Logf("expected: nil; got: %v", err)

// 	err = ValidateArgs(2, 4)
// 	if err == nil {
// 		t.Logf("tested with an error2")
// 		t.Errorf("expected: err; got: %v", err)
// 	}
// 	t.Logf("expected: err; got: %v", err.Error())
// }

// func TestBalanceFromStoreFormatter(t *testing.T) {

// 	valF := BalanceFromStoreFormatter("0.356", nil)
// 	if valF == nil {
// 		t.Logf("tested with an error1")
// 		t.Fatalf("TestBalanceFromStoreFormatter got an unexpected value: %v", valF)
// 		t.FailNow()
// 	}
// 	t.Logf("expected: *big.Float non null value; got: %v", valF)

// 	valF = BalanceFromStoreFormatter("", errors.New("unknown key"))
// 	if valF != nil {
// 		t.Logf("tested with an error2")
// 		t.Fatalf("TestBalanceFromStoreFormatter got an unexpected value: %v", valF)
// 		t.FailNow()
// 	}
// 	t.Logf("expected: nil; got: %v", valF)

// 	valF = BalanceFromStoreFormatter("0.2145", errors.New("unknown key"))
// 	if valF != nil {
// 		t.Logf("tested with an error and valid value")
// 		t.Fatalf("TestBalanceFromStoreFormatter got an empty value")
// 		t.FailNow()
// 	}
// 	t.Logf("expected: nil; got: %v", valF)

// }

// // ######################################################### encryption ########################################################

// func TestEncryptKey(t *testing.T) {

// 	hash, err := EncryptKey(key, strToEncrypt)
// 	if err != nil {
// 		t.Errorf("TestEncryptKey failed with: %v", err.Error())
// 		t.Fail()
// 	}
// 	t.Logf("encrypted val is -> %v", hash)

// }

// func TestEncryptByteArr(t *testing.T) {

// 	hash, err := EncryptByteArr(byteKey, byteSliceToEncrypt)
// 	if err != nil {
// 		t.Errorf("TestEncryptByteArr failed with: %v", err.Error())
// 		t.Fail()
// 	}
// 	t.Logf("encrypted val is -> %v", hash)

// }

// func TestDectyptByteArr(t *testing.T) {

// 	encryptedStrVal := "yzg6zI-1cVOxjGqkRxVGiDmKV-pHCzHdK-l8g8ggeLgiwv0QcrdUbjLw-WwYEiy6TgoSQW_JZABjNdp5bUYq0_gfplEnjiF5GkLM5LITaTkDlUDY8dYxsbyVymsIn5mlrW3H-ulJEZI-7aqbZeiMq4Mce1ptavvUNkP2Au9SSSRh3cXRi6OmDX3QMMAZFpo61s8="
// 	key := []byte(config.GetAnEncryptionKey())
// 	hash, err := DecryptByteArr(key, encryptedStrVal)
// 	if err != nil {
// 		t.Errorf("TestDectyptByteArr failed with: %v", err.Error())
// 		t.Fail()
// 	}
// 	if string(hash) != string(byteSliceToEncrypt) {
// 		t.Errorf("TestDectyptByteArr failed with: %v", errors.New("decoded value is different"))
// 		t.Fail()
// 	}
// 	t.Logf("decrypted val is -> %v", hash)
// }

// func TestDecryptKey(t *testing.T) {

// 	str := "SmVagK4rMWcMdoymwnXn1Fb3AYnRPO-Q6pgsWx32bii5mG_wYZlhLBFkm-NQRIYPpDPNahoGCrzUlvKaDdYUBmuBqjtB3chUchZsvotUT9k="
// 	key := config.GetAnEncryptionKey()
// 	hash, err := DecryptKey(key, str)
// 	if err != nil {
// 		t.Errorf("TestDecryptKey failed with: %v", err.Error())
// 		t.Fail()
// 	}
// 	if hash != strToEncrypt {
// 		t.Errorf("TestDecryptKey failed with: %v", errors.New("decoded value is different"))
// 		t.Fail()
// 	}
// 	t.Logf("decrypted val is -> %v", hash)
// }
