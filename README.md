# oasis-api-benchmark

Benchmarks various read-only gRPC calls to the [oasis-node](https://github.com/oasisprotocol/oasis-core/).

## Usage
```
./oasis-api-benchmark NODE_GRPC_ADDRESS
  NODE_GRPC_ADDRESS
      accepts either unix:/path/file.sock or hostname:port
  -genesis string
      Path to genesis file (default "genesis.json")
      (only used to get the base epoch and block heights)
  -n int
      number of times to run each API call (default 10)
      (parameters are randomized, to avoid caching)
```

## Example output
```
Beacon.GetBaseEpoch:                   Avg  171.374µs      Min  63.811µs      Max  880.018µs
Beacon.GetEpoch:                       Avg  3.810537ms     Min  3.078717ms    Max  4.583733ms
Beacon.GetEpochBlock:                  Avg  3.131655521s   Min  24.877358ms   Max  21.134783859s
Consensus.GetBlock:                    Avg  4.945797ms     Min  2.062967ms    Max  11.229401ms
Consensus.GetChainContext:             Avg  32.962357ms    Min  29.552626ms   Max  40.538887ms
Consensus.GetGenesisDocument:          Avg  48.284482ms    Min  45.085165ms   Max  58.75043ms
Consensus.GetStatus:                   Avg  33.660415ms    Min  30.010797ms   Max  47.295457ms
Consensus.GetTransactions:             Avg  3.852147ms     Min  2.201224ms    Max  5.980481ms
Consensus.GetTransactionsWithResults:  Avg  6.164239ms     Min  4.157356ms    Max  12.335644ms
Consensus.StateToGenesis:              Avg  59.244568527s  Min  5.881816076s  Max  2m40.776942607s
Staking.Account:                       Avg  11.598386ms    Min  8.64362ms     Max  14.225769ms
Staking.Addresses:                     Avg  8.201841922s   Min  2.330641495s  Max  25.342797179s
Staking.DebondingDelegationsFor:       Avg  257.226942ms   Min  11.819307ms   Max  611.362491ms
Staking.DelegationsFor:                Avg  3.334258902s   Min  581.24385ms   Max  12.430740302s
```
