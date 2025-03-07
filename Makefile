
# ------------ export PATH="$PATH:$(go env GOPATH)/bin"
gen:
	@protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	api/wallet.proto

build-linux:
	rm -rf bin | GOFLAGS=-mod=mod GOOS=linux GOARCH=amd64 go build -o bin/wallet main.go && cp bin/wallet ../crypto_app/walletService

build-macos:
	rm -rf bin | GOFLAGS=-mod=mod GOOS=darwin GOARCH=amd64 go build -o bin/wallet main.go | cp bin/wallet ../crypto_app/walletService

# ------------- run all tests 
test:
	go test ./... -v

migrate:
	migrate create -ext=sql -dir=/database/migrations/ -seq $(migrationName)