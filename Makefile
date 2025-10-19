
# generate wallet protobuf files
gen-wallet:
	protoc -I=api \
	--go_out ./gen/wallet \
	--go_opt=paths=source_relative \
	--go-grpc_out ./gen/wallet \
	--go-grpc_opt=paths=source_relative \
	wallet.proto

# run tests
run-test:
	export GO_ENV=test && go test -v ./...
