package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"

	stakingAPI "github.com/oasisprotocol/oasis-core/go/staking/api"
)

const DefaultNumSamples = 10

func main() {
	var genesisFile string
	var samples int
	flag.StringVar(&genesisFile, "genesis", DefaultGenesisFileName, "Path to genesis file")
	flag.IntVar(&samples, "n", DefaultNumSamples, "number of times to run each API call")
	flag.Parse()
	if flag.NArg() != 1 {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "Usage: %s GRPC_ADDRESS\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(64)
	}
	grpcAddress := flag.Arg(0)

	genesisHeight, genesisEpoch := getLowerLimits(genesisFile)
	currentHeight, currentEpoch := getUpperLimits(grpcAddress)
	addresses := getAddresses(grpcAddress, currentHeight)
	randomBlockId := func() int64 {
		return rand.Int63n(currentHeight-genesisHeight) + genesisHeight
	}
	randomEpochId := func() int64 {
		return rand.Int63n(currentEpoch-genesisEpoch) + genesisEpoch
	}
	randomAddress := func() stakingAPI.Address {
		return addresses[rand.Intn(len(addresses))]
	}

	fmt.Printf("Epoch: %d-%d\n", genesisEpoch, currentEpoch)
	fmt.Printf("Height: %d-%d\n", genesisHeight, currentHeight)
	fmt.Printf("Addresses: %d\n", len(addresses))
	fmt.Println()

	tests, closeConn := makeTests(grpcAddress, randomBlockId, randomEpochId, randomAddress)

	for name, f := range tests {
		testApiCall(name, f, samples)
	}

	closeConn()
}
