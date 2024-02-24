package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	filehub "github.com/iavl/fisco-go-sdk/examples"
	"math/big"
	"strings"

	"github.com/iavl/fisco-go-sdk/abi"
	"github.com/iavl/fisco-go-sdk/client"
	"github.com/iavl/fisco-go-sdk/conf"
	"github.com/iavl/fisco-go-sdk/core/types"
	"github.com/sirupsen/logrus"
)

func main() {
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		logrus.Fatal(err)
	}
	config := &configs[0]
	client, err := client.Dial(config)
	if err != nil {
		logrus.Fatal(err)
	}

	// deploy contract
	fmt.Println("-------------------starting deploy contract-----------------------")
	address, tx, _, err := filehub.DeployFileHub(client.GetTransactOpts(), client)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Println("contract address: ", address.Hex()) // the address should be saved
	fmt.Println("transaction hash: ", tx.Hash().Hex())

	// setup FileHub contract instance
	fmt.Println("\n-------------------setup fileHub contract instance-----------------------")
	//fileHubAddress := common.HexToAddress("0xd84ba29772d30b73320310f39e235d024ead7615")
	fileHubAddress := address
	instance, err := filehub.NewFileHub(fileHubAddress, client)
	if err != nil {
		logrus.Fatal(err)
	}

	aliceAddress := common.HexToAddress("0x0000000000000000000000000000000000000001")
	bobAddress := common.HexToAddress("0x0000000000000000000000000000000000000002")

	// register
	// call `register` to register user
	fmt.Println("\n-------------------starting call register-----------------------")
	filehubSession := &filehub.FileHubSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
	tx, _, err = filehubSession.Register(aliceAddress) // call `register` API
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

	// query registered user
	// invoke `registeredUsers` to query info
	fmt.Println("\n-------------------starting call registeredUsers to query-----------------------")
	isRegisteredUser, err := filehubSession.RegisteredUsers(aliceAddress) // call `registeredUsers` API
	if err != nil {
		logrus.Fatal(err)
	}
	if !isRegisteredUser {
		logrus.Fatalf("%v is not found \n", isRegisteredUser)
	}
	fmt.Printf("isRegisteredUser: %v \n", isRegisteredUser)

	// upload file
	// call `uploadFile`
	fmt.Println("\n-------------------starting call uploadFile-----------------------")
	fileHash := sha256.Sum256([]byte("hello world"))
	tx, _, err = filehubSession.UploadFile(aliceAddress, fileHash) // call `uploadFile` API
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

	// share file
	// call `shareFile`
	fmt.Println("\n-------------------starting call shareFile-----------------------")
	fromAddress := aliceAddress
	toAdddress := bobAddress
	tx, _, err = filehubSession.ShareFile(fromAddress, toAdddress, fileHash) // call `shareFile` API
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

	// forward file
	// call `forwardFile`
	fmt.Println("\n-------------------starting call forwardFile-----------------------")
	fromAddress = aliceAddress
	toAdddress = bobAddress
	proxyAddress := common.HexToAddress("0x1d29e94559B887e32020074c190522D5354F121b")
	tx, _, err = filehubSession.ForwardFile(fromAddress, toAdddress, proxyAddress, fileHash) // call `forwardFile` API
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

	// query forward record
	fmt.Println("\n-------------------starting call fileForwardingRecords to query-----------------------")
	hasForwarded, err := filehubSession.FileForwardingRecords(fromAddress, toAdddress, proxyAddress, fileHash) // call `registeredUsers` API
	if err != nil {
		logrus.Fatal(err)
	}
	if !hasForwarded {
		logrus.Fatalf("%v is not found \n", hasForwarded)
	}
	fmt.Printf("hasForwarded: %v \n", hasForwarded)

}

func parseOutput(abiStr, name string, receipt *types.Receipt) (*big.Int, error) {
	//fmt.Printf("receipt: %v\n", receipt)
	parsed, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		fmt.Printf("parse ABI failed, err: %v", err)
	}
	var ret *big.Int
	b, err := hex.DecodeString(receipt.Output[2:])
	if err != nil {
		return nil, fmt.Errorf("decode receipt.Output[2:] failed, err: %v", err)
	}
	err = parsed.Unpack(&ret, name, b)
	if err != nil {
		return nil, fmt.Errorf("unpack %v failed, err: %v", name, err)
	}
	return ret, nil
}
