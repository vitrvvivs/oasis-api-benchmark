package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type GenesisDocument struct {
	GenesisHeight int64  `json:"height"`
	HaltEpoch     int64  `json:"halt_epoch"`
	Beacon        Beacon `json:"beacon"`
}

type Beacon struct {
	Base int64 `json:"base"` //base epoch
}

const DefaultGenesisFileName = "genesis.json"

func loadGenesis(fileName string) (gen GenesisDocument, err error) {

	//Use root folder as default
	file, err := os.Open(fmt.Sprint("./", fileName))
	if err != nil {
		return gen, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&gen)
	if err != nil {
		return gen, err
	}

	return gen, nil
}

func getLowerLimits(genesisFilename string) (height int64, epoch int64) {

	gen, err := loadGenesis(DefaultGenesisFileName)
	if err != nil {
		Fail("ReadGenesisFile", err)
	}
	return gen.GenesisHeight, gen.Beacon.Base
}
