package gateway

import (
	"log"
	"net/http"
	"strconv"
	pb "wallet/api"
	"wallet/lib/helpers"
)

type handler struct {
	client pb.WalletServiceClient
}

func NewHandler(client pb.WalletServiceClient) *handler {
	return &handler{client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/wallet/create_wallet/{customerId}/", h.HandleCreateWallet)
	mux.HandleFunc("GET /api/wallet/create_address/{coin}/{customerId}/", h.HandleCreateOneTimeAddress)
	mux.HandleFunc("GET /api/wallet/get_balance/{coin}/{address}/{customerId}/", h.HandleGetBalance)

	mux.HandleFunc("POST /api/wallet/send_single_tsx/", h.HandleSendSingleTsx)
	mux.HandleFunc("POST /api/wallet/send_mult_tsx/", h.HandleSendMultiplyTsx)
}

func (h *handler) HandleCreateWallet(w http.ResponseWriter, r *http.Request) {
	customerId := r.PathValue("customerId")
	log.Println("customer id -> " + customerId)
}

func (h *handler) HandleCreateOneTimeAddress(w http.ResponseWriter, r *http.Request) {
	customerId, _ := strconv.Atoi(r.PathValue("customerId"))
	coinName := r.PathValue("coin")

	wallet, err := h.client.CreateAddress(r.Context(), &pb.CreateAddressRequest{
		CustomerId: int64(customerId),
		CoinName:   coinName,
	})
	if err != nil {
		helpers.WriteError(w, 500, err.Error())
	}

	helpers.WriteJSON(w, 200, wallet)
}

func (h *handler) HandleGetBalance(w http.ResponseWriter, r *http.Request) {
	customerId := r.PathValue("customer_id")
	log.Println("customer id -> " + customerId)
}

func (h *handler) HandleSendSingleTsx(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) HandleSendMultiplyTsx(w http.ResponseWriter, r *http.Request) {

}
