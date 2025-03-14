// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: api/wallet.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WalletItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CoinName      string                 `protobuf:"bytes,2,opt,name=coin_name,json=coinName,proto3" json:"coin_name,omitempty"`
	Address       string                 `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"` // double balance = 4;
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WalletItem) Reset() {
	*x = WalletItem{}
	mi := &file_api_wallet_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WalletItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletItem) ProtoMessage() {}

func (x *WalletItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_wallet_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletItem.ProtoReflect.Descriptor instead.
func (*WalletItem) Descriptor() ([]byte, []int) {
	return file_api_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *WalletItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WalletItem) GetCoinName() string {
	if x != nil {
		return x.CoinName
	}
	return ""
}

func (x *WalletItem) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type CoinBalance struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CoinName      string                 `protobuf:"bytes,1,opt,name=coin_name,json=coinName,proto3" json:"coin_name,omitempty"`
	Balance       string                 `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CoinBalance) Reset() {
	*x = CoinBalance{}
	mi := &file_api_wallet_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CoinBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinBalance) ProtoMessage() {}

func (x *CoinBalance) ProtoReflect() protoreflect.Message {
	mi := &file_api_wallet_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinBalance.ProtoReflect.Descriptor instead.
func (*CoinBalance) Descriptor() ([]byte, []int) {
	return file_api_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *CoinBalance) GetCoinName() string {
	if x != nil {
		return x.CoinName
	}
	return ""
}

func (x *CoinBalance) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

type CustomerBalance struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CoinBalance   string                 `protobuf:"bytes,1,opt,name=coin_balance,json=coinBalance,proto3" json:"coin_balance,omitempty"`
	FiatBalance   string                 `protobuf:"bytes,2,opt,name=fiat_balance,json=fiatBalance,proto3" json:"fiat_balance,omitempty"`
	CurrencyName  string                 `protobuf:"bytes,3,opt,name=currency_name,json=currencyName,proto3" json:"currency_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomerBalance) Reset() {
	*x = CustomerBalance{}
	mi := &file_api_wallet_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerBalance) ProtoMessage() {}

func (x *CustomerBalance) ProtoReflect() protoreflect.Message {
	mi := &file_api_wallet_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerBalance.ProtoReflect.Descriptor instead.
func (*CustomerBalance) Descriptor() ([]byte, []int) {
	return file_api_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *CustomerBalance) GetCoinBalance() string {
	if x != nil {
		return x.CoinBalance
	}
	return ""
}

func (x *CustomerBalance) GetFiatBalance() string {
	if x != nil {
		return x.FiatBalance
	}
	return ""
}

func (x *CustomerBalance) GetCurrencyName() string {
	if x != nil {
		return x.CurrencyName
	}
	return ""
}

type TransactionHash struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TsxHash       string                 `protobuf:"bytes,1,opt,name=tsx_hash,json=tsxHash,proto3" json:"tsx_hash,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransactionHash) Reset() {
	*x = TransactionHash{}
	mi := &file_api_wallet_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionHash) ProtoMessage() {}

func (x *TransactionHash) ProtoReflect() protoreflect.Message {
	mi := &file_api_wallet_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionHash.ProtoReflect.Descriptor instead.
func (*TransactionHash) Descriptor() ([]byte, []int) {
	return file_api_wallet_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionHash) GetTsxHash() string {
	if x != nil {
		return x.TsxHash
	}
	return ""
}

type WalletList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Wallet        []*WalletItem          `protobuf:"bytes,1,rep,name=wallet,proto3" json:"wallet,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WalletList) Reset() {
	*x = WalletList{}
	mi := &file_api_wallet_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WalletList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletList) ProtoMessage() {}

func (x *WalletList) ProtoReflect() protoreflect.Message {
	mi := &file_api_wallet_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletList.ProtoReflect.Descriptor instead.
func (*WalletList) Descriptor() ([]byte, []int) {
	return file_api_wallet_proto_rawDescGZIP(), []int{4}
}

func (x *WalletList) GetWallet() []*WalletItem {
	if x != nil {
		return x.Wallet
	}
	return nil
}

// CreateWalletRequest -> create wallet in each available blockchain (a list of)
type CreateWalletRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CustomerId    int64                  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateWalletRequest) Reset() {
	*x = CreateWalletRequest{}
	mi := &file_api_wallet_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWalletRequest) ProtoMessage() {}

func (x *CreateWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_wallet_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWalletRequest.ProtoReflect.Descriptor instead.
func (*CreateWalletRequest) Descriptor() ([]byte, []int) {
	return file_api_wallet_proto_rawDescGZIP(), []int{5}
}

func (x *CreateWalletRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

// CreateAddressRequest -> generate one-time addres in certain blockchain
type CreateAddressRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CustomerId    int64                  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	CoinName      string                 `protobuf:"bytes,2,opt,name=coin_name,json=coinName,proto3" json:"coin_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAddressRequest) Reset() {
	*x = CreateAddressRequest{}
	mi := &file_api_wallet_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAddressRequest) ProtoMessage() {}

func (x *CreateAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_wallet_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAddressRequest.ProtoReflect.Descriptor instead.
func (*CreateAddressRequest) Descriptor() ([]byte, []int) {
	return file_api_wallet_proto_rawDescGZIP(), []int{6}
}

func (x *CreateAddressRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *CreateAddressRequest) GetCoinName() string {
	if x != nil {
		return x.CoinName
	}
	return ""
}

// GetCustomerBalanceRequest -> get balance data by customer id
type GetCustomerBalanceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CustomerId    int64                  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	CurrencyName  string                 `protobuf:"bytes,2,opt,name=currency_name,json=currencyName,proto3" json:"currency_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCustomerBalanceRequest) Reset() {
	*x = GetCustomerBalanceRequest{}
	mi := &file_api_wallet_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCustomerBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerBalanceRequest) ProtoMessage() {}

func (x *GetCustomerBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_wallet_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerBalanceRequest) Descriptor() ([]byte, []int) {
	return file_api_wallet_proto_rawDescGZIP(), []int{7}
}

func (x *GetCustomerBalanceRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *GetCustomerBalanceRequest) GetCurrencyName() string {
	if x != nil {
		return x.CurrencyName
	}
	return ""
}

// GetCoinBalanceRequest -> get balance data by address and coin name
type GetCoinBalanceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CoinName      string                 `protobuf:"bytes,1,opt,name=coin_name,json=coinName,proto3" json:"coin_name,omitempty"`
	Address       string                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCoinBalanceRequest) Reset() {
	*x = GetCoinBalanceRequest{}
	mi := &file_api_wallet_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCoinBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoinBalanceRequest) ProtoMessage() {}

func (x *GetCoinBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_wallet_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCoinBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetCoinBalanceRequest) Descriptor() ([]byte, []int) {
	return file_api_wallet_proto_rawDescGZIP(), []int{8}
}

func (x *GetCoinBalanceRequest) GetCoinName() string {
	if x != nil {
		return x.CoinName
	}
	return ""
}

func (x *GetCoinBalanceRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// SendSingleTsxRequest -> send transaction direct to the certain users
type SendSingleTsxRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CoinName      string                 `protobuf:"bytes,1,opt,name=coin_name,json=coinName,proto3" json:"coin_name,omitempty"`
	Payee         *PeyeeData             `protobuf:"bytes,2,opt,name=payee,proto3" json:"payee,omitempty"`
	Beneficiar    *BeneficiarData        `protobuf:"bytes,3,opt,name=beneficiar,proto3" json:"beneficiar,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendSingleTsxRequest) Reset() {
	*x = SendSingleTsxRequest{}
	mi := &file_api_wallet_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendSingleTsxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSingleTsxRequest) ProtoMessage() {}

func (x *SendSingleTsxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_wallet_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSingleTsxRequest.ProtoReflect.Descriptor instead.
func (*SendSingleTsxRequest) Descriptor() ([]byte, []int) {
	return file_api_wallet_proto_rawDescGZIP(), []int{9}
}

func (x *SendSingleTsxRequest) GetCoinName() string {
	if x != nil {
		return x.CoinName
	}
	return ""
}

func (x *SendSingleTsxRequest) GetPayee() *PeyeeData {
	if x != nil {
		return x.Payee
	}
	return nil
}

func (x *SendSingleTsxRequest) GetBeneficiar() *BeneficiarData {
	if x != nil {
		return x.Beneficiar
	}
	return nil
}

// SendMultipleTsxRequest -> send transaction to the multiply users
type SendMultipleTsxRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	CoinName       string                 `protobuf:"bytes,1,opt,name=coin_name,json=coinName,proto3" json:"coin_name,omitempty"`
	Payee          *PeyeeData             `protobuf:"bytes,2,opt,name=payee,proto3" json:"payee,omitempty"`
	BeneficiarList []*BeneficiarData      `protobuf:"bytes,3,rep,name=beneficiar_list,json=beneficiarList,proto3" json:"beneficiar_list,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *SendMultipleTsxRequest) Reset() {
	*x = SendMultipleTsxRequest{}
	mi := &file_api_wallet_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMultipleTsxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMultipleTsxRequest) ProtoMessage() {}

func (x *SendMultipleTsxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_wallet_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMultipleTsxRequest.ProtoReflect.Descriptor instead.
func (*SendMultipleTsxRequest) Descriptor() ([]byte, []int) {
	return file_api_wallet_proto_rawDescGZIP(), []int{10}
}

func (x *SendMultipleTsxRequest) GetCoinName() string {
	if x != nil {
		return x.CoinName
	}
	return ""
}

func (x *SendMultipleTsxRequest) GetPayee() *PeyeeData {
	if x != nil {
		return x.Payee
	}
	return nil
}

func (x *SendMultipleTsxRequest) GetBeneficiarList() []*BeneficiarData {
	if x != nil {
		return x.BeneficiarList
	}
	return nil
}

type BeneficiarData struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	BeneficiarAddress string                 `protobuf:"bytes,1,opt,name=beneficiar_address,json=beneficiarAddress,proto3" json:"beneficiar_address,omitempty"`
	BeneficiarId      int64                  `protobuf:"varint,2,opt,name=beneficiar_id,json=beneficiarId,proto3" json:"beneficiar_id,omitempty"`
	Amount            string                 `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *BeneficiarData) Reset() {
	*x = BeneficiarData{}
	mi := &file_api_wallet_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BeneficiarData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeneficiarData) ProtoMessage() {}

func (x *BeneficiarData) ProtoReflect() protoreflect.Message {
	mi := &file_api_wallet_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeneficiarData.ProtoReflect.Descriptor instead.
func (*BeneficiarData) Descriptor() ([]byte, []int) {
	return file_api_wallet_proto_rawDescGZIP(), []int{11}
}

func (x *BeneficiarData) GetBeneficiarAddress() string {
	if x != nil {
		return x.BeneficiarAddress
	}
	return ""
}

func (x *BeneficiarData) GetBeneficiarId() int64 {
	if x != nil {
		return x.BeneficiarId
	}
	return 0
}

func (x *BeneficiarData) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type PeyeeData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PeyeeAddress  string                 `protobuf:"bytes,1,opt,name=peyee_address,json=peyeeAddress,proto3" json:"peyee_address,omitempty"`
	PeyeeId       int64                  `protobuf:"varint,2,opt,name=peyee_id,json=peyeeId,proto3" json:"peyee_id,omitempty"`
	Amount        string                 `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PeyeeData) Reset() {
	*x = PeyeeData{}
	mi := &file_api_wallet_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PeyeeData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeyeeData) ProtoMessage() {}

func (x *PeyeeData) ProtoReflect() protoreflect.Message {
	mi := &file_api_wallet_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeyeeData.ProtoReflect.Descriptor instead.
func (*PeyeeData) Descriptor() ([]byte, []int) {
	return file_api_wallet_proto_rawDescGZIP(), []int{12}
}

func (x *PeyeeData) GetPeyeeAddress() string {
	if x != nil {
		return x.PeyeeAddress
	}
	return ""
}

func (x *PeyeeData) GetPeyeeId() int64 {
	if x != nil {
		return x.PeyeeId
	}
	return 0
}

func (x *PeyeeData) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

var File_api_wallet_proto protoreflect.FileDescriptor

var file_api_wallet_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x53, 0x0a, 0x0a, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x44, 0x0a, 0x0b,
	0x43, 0x6f, 0x69, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6f, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x22, 0x7c, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x69,
	0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x61, 0x74,
	0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x66, 0x69, 0x61, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x2c, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x73, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x73, 0x78, 0x48, 0x61, 0x73, 0x68, 0x22, 0x35,
	0x0a, 0x0a, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x06,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x22, 0x36, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x54, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x61, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4e, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69,
	0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x8e, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x73, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x05,
	0x70, 0x61, 0x79, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x50, 0x65, 0x79, 0x65, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x70, 0x61, 0x79,
	0x65, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x65, 0x6e,
	0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x62, 0x65, 0x6e,
	0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x22, 0x99, 0x01, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x54, 0x73, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x24, 0x0a, 0x05, 0x70, 0x61, 0x79, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x65, 0x79, 0x65, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05,
	0x70, 0x61, 0x79, 0x65, 0x65, 0x12, 0x3c, 0x0a, 0x0f, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63,
	0x69, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x0e, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x7c, 0x0a, 0x0e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x12, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63,
	0x69, 0x61, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69,
	0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x62, 0x65, 0x6e,
	0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x63, 0x0a, 0x09, 0x50, 0x65, 0x79, 0x65, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x23,
	0x0a, 0x0d, 0x70, 0x65, 0x79, 0x65, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x65, 0x79, 0x65, 0x65, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x65, 0x79, 0x65, 0x65, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xf6, 0x03, 0x0a, 0x0d, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x12, 0x44, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x69, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x54, 0x73, 0x78, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x73, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x61, 0x73, 0x68, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x54, 0x73, 0x78, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x54, 0x73, 0x78,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x22, 0x00, 0x42,
	0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f,
	0x6f, 0x38, 0x78, 0x6c, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_wallet_proto_rawDescOnce sync.Once
	file_api_wallet_proto_rawDescData []byte
)

func file_api_wallet_proto_rawDescGZIP() []byte {
	file_api_wallet_proto_rawDescOnce.Do(func() {
		file_api_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_wallet_proto_rawDesc), len(file_api_wallet_proto_rawDesc)))
	})
	return file_api_wallet_proto_rawDescData
}

var file_api_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_wallet_proto_goTypes = []any{
	(*WalletItem)(nil),                // 0: api.WalletItem
	(*CoinBalance)(nil),               // 1: api.CoinBalance
	(*CustomerBalance)(nil),           // 2: api.CustomerBalance
	(*TransactionHash)(nil),           // 3: api.TransactionHash
	(*WalletList)(nil),                // 4: api.WalletList
	(*CreateWalletRequest)(nil),       // 5: api.CreateWalletRequest
	(*CreateAddressRequest)(nil),      // 6: api.CreateAddressRequest
	(*GetCustomerBalanceRequest)(nil), // 7: api.GetCustomerBalanceRequest
	(*GetCoinBalanceRequest)(nil),     // 8: api.GetCoinBalanceRequest
	(*SendSingleTsxRequest)(nil),      // 9: api.SendSingleTsxRequest
	(*SendMultipleTsxRequest)(nil),    // 10: api.SendMultipleTsxRequest
	(*BeneficiarData)(nil),            // 11: api.BeneficiarData
	(*PeyeeData)(nil),                 // 12: api.PeyeeData
}
var file_api_wallet_proto_depIdxs = []int32{
	0,  // 0: api.WalletList.wallet:type_name -> api.WalletItem
	12, // 1: api.SendSingleTsxRequest.payee:type_name -> api.PeyeeData
	11, // 2: api.SendSingleTsxRequest.beneficiar:type_name -> api.BeneficiarData
	12, // 3: api.SendMultipleTsxRequest.payee:type_name -> api.PeyeeData
	11, // 4: api.SendMultipleTsxRequest.beneficiar_list:type_name -> api.BeneficiarData
	5,  // 5: api.WalletService.CreateWallet:input_type -> api.CreateWalletRequest
	6,  // 6: api.WalletService.CreatePermanentAddress:input_type -> api.CreateAddressRequest
	6,  // 7: api.WalletService.CreateOneTimeAddress:input_type -> api.CreateAddressRequest
	8,  // 8: api.WalletService.GetCoinBalance:input_type -> api.GetCoinBalanceRequest
	7,  // 9: api.WalletService.GetCustomerBalance:input_type -> api.GetCustomerBalanceRequest
	9,  // 10: api.WalletService.SendSingleTsx:input_type -> api.SendSingleTsxRequest
	10, // 11: api.WalletService.SendMultiplyTsx:input_type -> api.SendMultipleTsxRequest
	4,  // 12: api.WalletService.CreateWallet:output_type -> api.WalletList
	0,  // 13: api.WalletService.CreatePermanentAddress:output_type -> api.WalletItem
	0,  // 14: api.WalletService.CreateOneTimeAddress:output_type -> api.WalletItem
	1,  // 15: api.WalletService.GetCoinBalance:output_type -> api.CoinBalance
	2,  // 16: api.WalletService.GetCustomerBalance:output_type -> api.CustomerBalance
	3,  // 17: api.WalletService.SendSingleTsx:output_type -> api.TransactionHash
	3,  // 18: api.WalletService.SendMultiplyTsx:output_type -> api.TransactionHash
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_wallet_proto_init() }
func file_api_wallet_proto_init() {
	if File_api_wallet_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_wallet_proto_rawDesc), len(file_api_wallet_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_wallet_proto_goTypes,
		DependencyIndexes: file_api_wallet_proto_depIdxs,
		MessageInfos:      file_api_wallet_proto_msgTypes,
	}.Build()
	File_api_wallet_proto = out.File
	file_api_wallet_proto_goTypes = nil
	file_api_wallet_proto_depIdxs = nil
}
