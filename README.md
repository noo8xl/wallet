# wallet

## REST service

In addition to the gRPC service, there is a RESTful HTTP service under the `REST` folder.

### Run REST server

From the project root:

```bash
go run ./REST
```

The REST server listens on `0.0.0.0:8080` and exposes:

- `POST /api/v1/wallets/{customerId}` – create wallets for all supported coins for a customer.
- `POST /api/v1/wallets/{coin}/{customerId}/permanent` – create a permanent wallet address.
- `POST /api/v1/wallets/{coin}/{customerId}/one-time` – create a one-time wallet address.
- `GET /api/v1/wallets/{coin}/{address}/balance` – get balance for a specific coin/address.


