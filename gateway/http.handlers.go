package main

import (
	"log"
	"net/http"
	"strconv"
	pb "wallet/gen/wallet"
	"wallet/pkg/utils"
)

type handler struct {
	client pb.WalletServiceClient
}

func NewHandler(client pb.WalletServiceClient) *handler {
	return &handler{client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /wallet/create_wallet/{customerId}/", h.HandleCreateWallet)
	mux.HandleFunc("POST /wallet/create_one_time_address/{coin}/{customerId}/", h.HandleCreateOneTimeAddress)
	mux.HandleFunc("POST /wallet/create_permanent_address/{coin}/{customerId}/", h.HandleCreatePermanentAddress)

	mux.HandleFunc("GET /wallet/get_balance/{coin}/{address}/{customerId}/", h.HandleGetBalance)

	mux.HandleFunc("POST /wallet/send_single_tsx/", h.HandleSendSingleTsx)
	mux.HandleFunc("POST /wallet/send_mult_tsx/", h.HandleSendMultiplyTsx)
}

func (h *handler) HandleCreateWallet(w http.ResponseWriter, r *http.Request) {
	customerId := r.PathValue("customerId")
	log.Println("customer id -> " + customerId)
}

func (h *handler) HandleCreatePermanentAddress(w http.ResponseWriter, r *http.Request) {
	customerId, _ := strconv.Atoi(r.PathValue("customerId"))
	coinName := r.PathValue("coin")

	log.Println("req -> ", customerId, coinName)

	wallet, err := h.client.CreatePermanentAddress(r.Context(), &pb.CreateAddressRequest{
		CustomerId: int64(customerId),
		CoinName:   coinName,
	})
	if err != nil {
		if err.Error() == "already exists" {
			utils.WriteError(w, 404, "Bad request. Permanent wallet already exists")
		}
		utils.WriteError(w, 500, err.Error())
	}

	utils.WriteJSON(w, 200, wallet)
}

func (h *handler) HandleCreateOneTimeAddress(w http.ResponseWriter, r *http.Request) {
	customerId, _ := strconv.Atoi(r.PathValue("customerId"))
	coinName := r.PathValue("coin")

	log.Println("req -> ", customerId, coinName)

	wallet, err := h.client.CreateOneTimeAddress(r.Context(), &pb.CreateAddressRequest{
		CustomerId: int64(customerId),
		CoinName:   coinName,
	})
	if err != nil {
		utils.WriteError(w, 500, err.Error())
	}

	utils.WriteJSON(w, 200, wallet)
}

func (h *handler) HandleGetBalance(w http.ResponseWriter, r *http.Request) {
	customerId := r.PathValue("customer_id")
	log.Println("customer id -> " + customerId)
}

func (h *handler) HandleSendSingleTsx(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) HandleSendMultiplyTsx(w http.ResponseWriter, r *http.Request) {

}
