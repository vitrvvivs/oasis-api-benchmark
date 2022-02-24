package main

import (
	"context"
	"fmt"
	"time"

	beaconAPI "github.com/oasisprotocol/oasis-core/go/beacon/api"
	oasisgrpc "github.com/oasisprotocol/oasis-core/go/common/grpc"
	consensusAPI "github.com/oasisprotocol/oasis-core/go/consensus/api"
	stakingAPI "github.com/oasisprotocol/oasis-core/go/staking/api"
	"google.golang.org/grpc"
)

func testApiCall(name string, f func() error, samples int) {
	var times []int64
	failed := false
	for i := 0; i < samples; i++ {
		starttime := time.Now()
		err := f()
		if err != nil {
			Warn(name, "error: ", err)
			failed = true
			break
		}
		times = append(times, int64(time.Since(starttime)))
	}
	if !failed {
		min := time.Duration(min(times))
		max := time.Duration(max(times))
		avg := time.Duration(sum(times) / int64(samples))
		fmt.Printf("%s:  Avg %s  Min %s  Max %s\n", name, avg, min, max)
	}
}

func getUpperLimits(grpcAddress string) (height int64, epoch int64) {

	conn, err := oasisgrpc.Dial(grpcAddress, grpc.WithInsecure())
	ctx, stop := context.WithCancel(context.Background())
	if err != nil {
		Fail(err)
	}

	cAPI := consensusAPI.NewConsensusClient(conn)
	bAPI := beaconAPI.NewBeaconClient(conn)

	// Get current block height and epoch
	bch, bsub, err := bAPI.WatchEpochs(ctx)
	if err != nil {
		Fail("WatchEpochs", err)
	}
	epoch = int64(<-bch)
	bsub.Close()

	cch, csub, err := cAPI.WatchBlocks(ctx)
	if err != nil {
		Fail("WatchEpochs", err)
	}
	height = (<-cch).Height
	csub.Close()

	conn.Close()
	stop()
	return
}

func getAddresses(grpcAddress string, currentHeight int64) (addresses []stakingAPI.Address) {

	conn, err := oasisgrpc.Dial(grpcAddress, grpc.WithInsecure())
	ctx, stop := context.WithCancel(context.Background())
	if err != nil {
		Fail(err)
	}

	sAPI := stakingAPI.NewStakingClient(conn)
	addresses, err = sAPI.Addresses(ctx, currentHeight)
	if err != nil {
		Fail("getAddresses", err)
	}

	conn.Close()
	stop()
	return
}

func getUpperLimitsV2(grpcAddress string) (height int64, epoch int64) {

	conn, err := oasisgrpc.Dial(grpcAddress, grpc.WithInsecure())
	ctx, stop := context.WithCancel(context.Background())
	if err != nil {
		Fail(err)
	}

	bAPI := beaconAPI.NewBeaconClient(conn)
	height = consensusAPI.HeightLatest
	epochtime, _ := bAPI.GetEpoch(ctx, consensusAPI.HeightLatest)
	epoch = int64(epochtime)

	conn.Close()
	stop()
	return
}

func makeTests(grpcAddress string, randomBlockId func() int64, randomEpochId func() int64, randomAddress func() stakingAPI.Address) (tests map[string](func() error), closeConn func()) {
	conn, err := oasisgrpc.Dial(grpcAddress, grpc.WithInsecure())
	ctx, stop := context.WithCancel(context.Background())
	if err != nil {
		Fail(err)
	}
	cAPI := consensusAPI.NewConsensusClient(conn)
	bAPI := beaconAPI.NewBeaconClient(conn)
	sAPI := stakingAPI.NewStakingClient(conn)
	closeConn = func() {
		conn.Close()
		stop()
	}

	tests = make(map[string](func() error))
	tests["Beacon.GetBaseEpoch"] = func() error {
		_, err := bAPI.GetBaseEpoch(ctx)
		return err
	}
	tests["Beacon.GetEpoch"] = func() error {
		_, err := bAPI.GetEpoch(ctx, randomBlockId())
		return err
	}
	tests["Beacon.GetEpochBlock"] = func() error {
		_, err := bAPI.GetEpochBlock(ctx, beaconAPI.EpochTime(randomEpochId()))
		return err
	}
	tests["Consensus.GetBlock"] = func() error {
		_, err := cAPI.GetBlock(ctx, randomBlockId())
		return err
	}
	tests["Consensus.GetTransactions"] = func() error {
		_, err := cAPI.GetTransactions(ctx, randomBlockId())
		return err
	}
	tests["Consensus.GetTransactionsWithResults"] = func() error {
		_, err := cAPI.GetTransactionsWithResults(ctx, randomBlockId())
		return err
	}
	tests["Consensus.StateToGenesis"] = func() error {
		_, err := cAPI.StateToGenesis(ctx, randomBlockId())
		return err
	}
	tests["Consensus.GetGenesisDocument"] = func() error {
		_, err := cAPI.GetGenesisDocument(ctx)
		return err
	}
	tests["Consensus.GetChainContext"] = func() error {
		_, err := cAPI.GetChainContext(ctx)
		return err
	}
	tests["Consensus.GetStatus"] = func() error {
		_, err := cAPI.GetStatus(ctx)
		return err
	}
	tests["Staking.Addresses"] = func() error {
		_, err := sAPI.Addresses(ctx, randomBlockId())
		return err
	}
	tests["Staking.Account"] = func() error {
		_, err := sAPI.Account(ctx, &stakingAPI.OwnerQuery{
			Height: randomBlockId(),
			Owner:  randomAddress(),
		})
		return err
	}
	tests["Staking.DelegationsFor"] = func() error {
		_, err := sAPI.DelegationsFor(ctx, &stakingAPI.OwnerQuery{
			Height: randomBlockId(),
			Owner:  randomAddress(),
		})
		return err
	}
	tests["Staking.DebondingDelegationsFor"] = func() error {
		_, err := sAPI.DebondingDelegationsFor(ctx, &stakingAPI.OwnerQuery{
			Height: randomBlockId(),
			Owner:  randomAddress(),
		})
		return err
	}
	return
}
