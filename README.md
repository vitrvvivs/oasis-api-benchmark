# oasis-api-benchmark

Benchmarks various read-only gRPC calls to the [oasis-node](https://github.com/oasisprotocol/oasis-core/).

## Usage
```
./oasis-api-benchmark NODE_GRPC_ADDRESS
  -genesis string
    	Path to genesis file (default "genesis.json")
      (only used to get the base epoch and block heights)
  -n int
    	number of times to run each API call (default 10)
      (parameters are randomized, to avoid caching)
