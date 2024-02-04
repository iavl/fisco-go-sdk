// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package examples

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/iavl/fisco-go-sdk/abi"
	"github.com/iavl/fisco-go-sdk/abi/bind"
	"github.com/iavl/fisco-go-sdk/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
)

// FileHubABI is the input ABI used to generate the binding from.
const FileHubABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"fileHash\",\"type\":\"bytes32\"}],\"name\":\"FileForwarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fileHash\",\"type\":\"bytes32\"}],\"name\":\"FileShared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fileHash\",\"type\":\"bytes32\"}],\"name\":\"FileUploaded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fileForwardingRecords\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fileOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fileUploadTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"fileHash\",\"type\":\"bytes32\"}],\"name\":\"forwardFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registeredUsers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"fileHash\",\"type\":\"bytes32\"}],\"name\":\"shareFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"fileHash\",\"type\":\"bytes32\"}],\"name\":\"uploadFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FileHubBin is the compiled bytecode used for deploying new contracts.
var FileHubBin = "0x608060405234801561001057600080fd5b50610da7806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80637ea8e04a1161005b5780637ea8e04a1461020757806383263f51146102755780638bd2194f146102e357806394ed4dfd1461038957610088565b80630adbc3c21461008d5780633c29fd101461011b57806368457e3f1461015d57806372997968146101b9575b600080fd5b610119600480360360808110156100a357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506103cd565b005b6101476004803603602081101561013157600080fd5b81019080803590602001909291905050506105d3565b6040518082815260200191505060405180910390f35b61019f6004803603602081101561017357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506105eb565b604051808215151515815260200191505060405180910390f35b610205600480360360408110156101cf57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061060b565b005b6102336004803603602081101561021d57600080fd5b81019080803590602001909291905050506108f1565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102e16004803603606081101561028b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610924565b005b61036f600480360360808110156102f957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610c1e565b604051808215151515815260200191505060405180910390f35b6103cb6004803603602081101561039f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c67565b005b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461046e576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f46696c654875623a2073656e646572206e6f7420617574686f72697a6564000081525060200191505060405180910390fd5b6001600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fff5afd4402f59c0d0f6836336001cf85570ae9d76ae3c9ad60318b53df6063a5846040518082815260200191505060405180910390a450505050565b60026020528060005260406000206000915090505481565b60006020528060005260406000206000915054906101000a900460ff1681565b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106ac576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f46696c654875623a2073656e646572206e6f7420617574686f72697a6564000081525060200191505060405180910390fd5b6000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661076a576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f46696c654875623a206163636f756e74206e6f7420726567697374657265640081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166001600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461083f576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f46696c654875623a2066696c6520616c7265616479206578697374730000000081525060200191505060405180910390fd5b816001600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550426002600083815260200190815260200160002081905550808273ffffffffffffffffffffffffffffffffffffffff167fa0e5f045fbeb6fe5c48c17b8072c3247639af379477463c71d9d6c2ce2006ade60405160405180910390a35050565b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109c5576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f46696c654875623a2073656e646572206e6f7420617574686f72697a6564000081525060200191505060405180910390fd5b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610a66576040517fc703cb12000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180610d056024913960400191505060405180910390fd5b6000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610b07576040517fc703cb12000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180610d296022913960400191505060405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff166001600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610bbe576040517fc703cb12000000000000000000000000000000000000000000000000000000008152600401808060200182810382526027815260200180610d4b6027913960400191505060405180910390fd5b808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167ff86c05999291216c060f8639c3c88b12fe2b126679b0bf02943e3a85d1a7256d60405160405180910390a4505050565b600360205283600052604060002060205282600052604060002060205281600052604060002060205280600052604060002060009350935050509054906101000a900460ff1681565b60016000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff167fa6419914aa99164db3f61f5297f071bc0e4d57eff1847b4c648d2240b4be49ff60405160405180910390a25056fe46696c654875623a2066726f6d206163636f756e74206e6f74207265676973746572656446696c654875623a20746f206163636f756e74206e6f74207265676973746572656446696c654875623a2066696c65206e6f74206f776e65642062792066726f6d206163636f756e74a264697066735822122092bbc8026d14f158dac83dd9672beb3ca3220b3f478405d3f5bf8915a7cd167f64736f6c634300060a0033"

// DeployFileHub deploys a new contract, binding an instance of FileHub to it.
func DeployFileHub(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FileHub, error) {
	parsed, err := abi.JSON(strings.NewReader(FileHubABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FileHubBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FileHub{FileHubCaller: FileHubCaller{contract: contract}, FileHubTransactor: FileHubTransactor{contract: contract}, FileHubFilterer: FileHubFilterer{contract: contract}}, nil
}

func AsyncDeployFileHub(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(FileHubABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(FileHubBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// FileHub is an auto generated Go binding around a Solidity contract.
type FileHub struct {
	FileHubCaller     // Read-only binding to the contract
	FileHubTransactor // Write-only binding to the contract
	FileHubFilterer   // Log filterer for contract events
}

// FileHubCaller is an auto generated read-only Go binding around a Solidity contract.
type FileHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileHubTransactor is an auto generated write-only Go binding around a Solidity contract.
type FileHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileHubFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type FileHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileHubSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type FileHubSession struct {
	Contract     *FileHub          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FileHubCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type FileHubCallerSession struct {
	Contract *FileHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FileHubTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type FileHubTransactorSession struct {
	Contract     *FileHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FileHubRaw is an auto generated low-level Go binding around a Solidity contract.
type FileHubRaw struct {
	Contract *FileHub // Generic contract binding to access the raw methods on
}

// FileHubCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type FileHubCallerRaw struct {
	Contract *FileHubCaller // Generic read-only contract binding to access the raw methods on
}

// FileHubTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type FileHubTransactorRaw struct {
	Contract *FileHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFileHub creates a new instance of FileHub, bound to a specific deployed contract.
func NewFileHub(address common.Address, backend bind.ContractBackend) (*FileHub, error) {
	contract, err := bindFileHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FileHub{FileHubCaller: FileHubCaller{contract: contract}, FileHubTransactor: FileHubTransactor{contract: contract}, FileHubFilterer: FileHubFilterer{contract: contract}}, nil
}

// NewFileHubCaller creates a new read-only instance of FileHub, bound to a specific deployed contract.
func NewFileHubCaller(address common.Address, caller bind.ContractCaller) (*FileHubCaller, error) {
	contract, err := bindFileHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FileHubCaller{contract: contract}, nil
}

// NewFileHubTransactor creates a new write-only instance of FileHub, bound to a specific deployed contract.
func NewFileHubTransactor(address common.Address, transactor bind.ContractTransactor) (*FileHubTransactor, error) {
	contract, err := bindFileHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FileHubTransactor{contract: contract}, nil
}

// NewFileHubFilterer creates a new log filterer instance of FileHub, bound to a specific deployed contract.
func NewFileHubFilterer(address common.Address, filterer bind.ContractFilterer) (*FileHubFilterer, error) {
	contract, err := bindFileHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FileHubFilterer{contract: contract}, nil
}

// bindFileHub binds a generic wrapper to an already deployed contract.
func bindFileHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FileHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FileHub *FileHubRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FileHub.Contract.FileHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FileHub *FileHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _FileHub.Contract.FileHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FileHub *FileHubRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _FileHub.Contract.FileHubTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FileHub *FileHubCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FileHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FileHub *FileHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _FileHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FileHub *FileHubTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _FileHub.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// FileForwardingRecords is a free data retrieval call binding the contract method 0x8bd2194f.
//
// Solidity: function fileForwardingRecords(address , address , address , bytes32 ) constant returns(bool)
func (_FileHub *FileHubCaller) FileForwardingRecords(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FileHub.contract.Call(opts, out, "fileForwardingRecords", arg0, arg1, arg2, arg3)
	return *ret0, err
}

// FileForwardingRecords is a free data retrieval call binding the contract method 0x8bd2194f.
//
// Solidity: function fileForwardingRecords(address , address , address , bytes32 ) constant returns(bool)
func (_FileHub *FileHubSession) FileForwardingRecords(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 [32]byte) (bool, error) {
	return _FileHub.Contract.FileForwardingRecords(&_FileHub.CallOpts, arg0, arg1, arg2, arg3)
}

// FileForwardingRecords is a free data retrieval call binding the contract method 0x8bd2194f.
//
// Solidity: function fileForwardingRecords(address , address , address , bytes32 ) constant returns(bool)
func (_FileHub *FileHubCallerSession) FileForwardingRecords(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 [32]byte) (bool, error) {
	return _FileHub.Contract.FileForwardingRecords(&_FileHub.CallOpts, arg0, arg1, arg2, arg3)
}

// FileOwners is a free data retrieval call binding the contract method 0x7ea8e04a.
//
// Solidity: function fileOwners(bytes32 ) constant returns(address)
func (_FileHub *FileHubCaller) FileOwners(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FileHub.contract.Call(opts, out, "fileOwners", arg0)
	return *ret0, err
}

// FileOwners is a free data retrieval call binding the contract method 0x7ea8e04a.
//
// Solidity: function fileOwners(bytes32 ) constant returns(address)
func (_FileHub *FileHubSession) FileOwners(arg0 [32]byte) (common.Address, error) {
	return _FileHub.Contract.FileOwners(&_FileHub.CallOpts, arg0)
}

// FileOwners is a free data retrieval call binding the contract method 0x7ea8e04a.
//
// Solidity: function fileOwners(bytes32 ) constant returns(address)
func (_FileHub *FileHubCallerSession) FileOwners(arg0 [32]byte) (common.Address, error) {
	return _FileHub.Contract.FileOwners(&_FileHub.CallOpts, arg0)
}

// FileUploadTime is a free data retrieval call binding the contract method 0x3c29fd10.
//
// Solidity: function fileUploadTime(bytes32 ) constant returns(uint256)
func (_FileHub *FileHubCaller) FileUploadTime(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FileHub.contract.Call(opts, out, "fileUploadTime", arg0)
	return *ret0, err
}

// FileUploadTime is a free data retrieval call binding the contract method 0x3c29fd10.
//
// Solidity: function fileUploadTime(bytes32 ) constant returns(uint256)
func (_FileHub *FileHubSession) FileUploadTime(arg0 [32]byte) (*big.Int, error) {
	return _FileHub.Contract.FileUploadTime(&_FileHub.CallOpts, arg0)
}

// FileUploadTime is a free data retrieval call binding the contract method 0x3c29fd10.
//
// Solidity: function fileUploadTime(bytes32 ) constant returns(uint256)
func (_FileHub *FileHubCallerSession) FileUploadTime(arg0 [32]byte) (*big.Int, error) {
	return _FileHub.Contract.FileUploadTime(&_FileHub.CallOpts, arg0)
}

// RegisteredUsers is a free data retrieval call binding the contract method 0x68457e3f.
//
// Solidity: function registeredUsers(address ) constant returns(bool)
func (_FileHub *FileHubCaller) RegisteredUsers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FileHub.contract.Call(opts, out, "registeredUsers", arg0)
	return *ret0, err
}

// RegisteredUsers is a free data retrieval call binding the contract method 0x68457e3f.
//
// Solidity: function registeredUsers(address ) constant returns(bool)
func (_FileHub *FileHubSession) RegisteredUsers(arg0 common.Address) (bool, error) {
	return _FileHub.Contract.RegisteredUsers(&_FileHub.CallOpts, arg0)
}

// RegisteredUsers is a free data retrieval call binding the contract method 0x68457e3f.
//
// Solidity: function registeredUsers(address ) constant returns(bool)
func (_FileHub *FileHubCallerSession) RegisteredUsers(arg0 common.Address) (bool, error) {
	return _FileHub.Contract.RegisteredUsers(&_FileHub.CallOpts, arg0)
}

// ForwardFile is a paid mutator transaction binding the contract method 0x0adbc3c2.
//
// Solidity: function forwardFile(address from, address to, address proxy, bytes32 fileHash) returns()
func (_FileHub *FileHubTransactor) ForwardFile(opts *bind.TransactOpts, from common.Address, to common.Address, proxy common.Address, fileHash [32]byte) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _FileHub.contract.TransactWithResult(opts, out, "forwardFile", from, to, proxy, fileHash)
	return transaction, receipt, err
}

func (_FileHub *FileHubTransactor) AsyncForwardFile(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, proxy common.Address, fileHash [32]byte) (*types.Transaction, error) {
	return _FileHub.contract.AsyncTransact(opts, handler, "forwardFile", from, to, proxy, fileHash)
}

// ForwardFile is a paid mutator transaction binding the contract method 0x0adbc3c2.
//
// Solidity: function forwardFile(address from, address to, address proxy, bytes32 fileHash) returns()
func (_FileHub *FileHubSession) ForwardFile(from common.Address, to common.Address, proxy common.Address, fileHash [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _FileHub.Contract.ForwardFile(&_FileHub.TransactOpts, from, to, proxy, fileHash)
}

func (_FileHub *FileHubSession) AsyncForwardFile(handler func(*types.Receipt, error), from common.Address, to common.Address, proxy common.Address, fileHash [32]byte) (*types.Transaction, error) {
	return _FileHub.Contract.AsyncForwardFile(handler, &_FileHub.TransactOpts, from, to, proxy, fileHash)
}

// ForwardFile is a paid mutator transaction binding the contract method 0x0adbc3c2.
//
// Solidity: function forwardFile(address from, address to, address proxy, bytes32 fileHash) returns()
func (_FileHub *FileHubTransactorSession) ForwardFile(from common.Address, to common.Address, proxy common.Address, fileHash [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _FileHub.Contract.ForwardFile(&_FileHub.TransactOpts, from, to, proxy, fileHash)
}

func (_FileHub *FileHubTransactorSession) AsyncForwardFile(handler func(*types.Receipt, error), from common.Address, to common.Address, proxy common.Address, fileHash [32]byte) (*types.Transaction, error) {
	return _FileHub.Contract.AsyncForwardFile(handler, &_FileHub.TransactOpts, from, to, proxy, fileHash)
}

// Register is a paid mutator transaction binding the contract method 0x94ed4dfd.
//
// Solidity: function register(address account) returns()
func (_FileHub *FileHubTransactor) Register(opts *bind.TransactOpts, account common.Address) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _FileHub.contract.TransactWithResult(opts, out, "register", account)
	return transaction, receipt, err
}

func (_FileHub *FileHubTransactor) AsyncRegister(handler func(*types.Receipt, error), opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _FileHub.contract.AsyncTransact(opts, handler, "register", account)
}

// Register is a paid mutator transaction binding the contract method 0x94ed4dfd.
//
// Solidity: function register(address account) returns()
func (_FileHub *FileHubSession) Register(account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _FileHub.Contract.Register(&_FileHub.TransactOpts, account)
}

func (_FileHub *FileHubSession) AsyncRegister(handler func(*types.Receipt, error), account common.Address) (*types.Transaction, error) {
	return _FileHub.Contract.AsyncRegister(handler, &_FileHub.TransactOpts, account)
}

// Register is a paid mutator transaction binding the contract method 0x94ed4dfd.
//
// Solidity: function register(address account) returns()
func (_FileHub *FileHubTransactorSession) Register(account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _FileHub.Contract.Register(&_FileHub.TransactOpts, account)
}

func (_FileHub *FileHubTransactorSession) AsyncRegister(handler func(*types.Receipt, error), account common.Address) (*types.Transaction, error) {
	return _FileHub.Contract.AsyncRegister(handler, &_FileHub.TransactOpts, account)
}

// ShareFile is a paid mutator transaction binding the contract method 0x83263f51.
//
// Solidity: function shareFile(address from, address to, bytes32 fileHash) returns()
func (_FileHub *FileHubTransactor) ShareFile(opts *bind.TransactOpts, from common.Address, to common.Address, fileHash [32]byte) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _FileHub.contract.TransactWithResult(opts, out, "shareFile", from, to, fileHash)
	return transaction, receipt, err
}

func (_FileHub *FileHubTransactor) AsyncShareFile(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, fileHash [32]byte) (*types.Transaction, error) {
	return _FileHub.contract.AsyncTransact(opts, handler, "shareFile", from, to, fileHash)
}

// ShareFile is a paid mutator transaction binding the contract method 0x83263f51.
//
// Solidity: function shareFile(address from, address to, bytes32 fileHash) returns()
func (_FileHub *FileHubSession) ShareFile(from common.Address, to common.Address, fileHash [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _FileHub.Contract.ShareFile(&_FileHub.TransactOpts, from, to, fileHash)
}

func (_FileHub *FileHubSession) AsyncShareFile(handler func(*types.Receipt, error), from common.Address, to common.Address, fileHash [32]byte) (*types.Transaction, error) {
	return _FileHub.Contract.AsyncShareFile(handler, &_FileHub.TransactOpts, from, to, fileHash)
}

// ShareFile is a paid mutator transaction binding the contract method 0x83263f51.
//
// Solidity: function shareFile(address from, address to, bytes32 fileHash) returns()
func (_FileHub *FileHubTransactorSession) ShareFile(from common.Address, to common.Address, fileHash [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _FileHub.Contract.ShareFile(&_FileHub.TransactOpts, from, to, fileHash)
}

func (_FileHub *FileHubTransactorSession) AsyncShareFile(handler func(*types.Receipt, error), from common.Address, to common.Address, fileHash [32]byte) (*types.Transaction, error) {
	return _FileHub.Contract.AsyncShareFile(handler, &_FileHub.TransactOpts, from, to, fileHash)
}

// UploadFile is a paid mutator transaction binding the contract method 0x72997968.
//
// Solidity: function uploadFile(address account, bytes32 fileHash) returns()
func (_FileHub *FileHubTransactor) UploadFile(opts *bind.TransactOpts, account common.Address, fileHash [32]byte) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _FileHub.contract.TransactWithResult(opts, out, "uploadFile", account, fileHash)
	return transaction, receipt, err
}

func (_FileHub *FileHubTransactor) AsyncUploadFile(handler func(*types.Receipt, error), opts *bind.TransactOpts, account common.Address, fileHash [32]byte) (*types.Transaction, error) {
	return _FileHub.contract.AsyncTransact(opts, handler, "uploadFile", account, fileHash)
}

// UploadFile is a paid mutator transaction binding the contract method 0x72997968.
//
// Solidity: function uploadFile(address account, bytes32 fileHash) returns()
func (_FileHub *FileHubSession) UploadFile(account common.Address, fileHash [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _FileHub.Contract.UploadFile(&_FileHub.TransactOpts, account, fileHash)
}

func (_FileHub *FileHubSession) AsyncUploadFile(handler func(*types.Receipt, error), account common.Address, fileHash [32]byte) (*types.Transaction, error) {
	return _FileHub.Contract.AsyncUploadFile(handler, &_FileHub.TransactOpts, account, fileHash)
}

// UploadFile is a paid mutator transaction binding the contract method 0x72997968.
//
// Solidity: function uploadFile(address account, bytes32 fileHash) returns()
func (_FileHub *FileHubTransactorSession) UploadFile(account common.Address, fileHash [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _FileHub.Contract.UploadFile(&_FileHub.TransactOpts, account, fileHash)
}

func (_FileHub *FileHubTransactorSession) AsyncUploadFile(handler func(*types.Receipt, error), account common.Address, fileHash [32]byte) (*types.Transaction, error) {
	return _FileHub.Contract.AsyncUploadFile(handler, &_FileHub.TransactOpts, account, fileHash)
}

// FileHubFileForwarded represents a FileForwarded event raised by the FileHub contract.
type FileHubFileForwarded struct {
	From     common.Address
	To       common.Address
	Proxy    common.Address
	FileHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// WatchFileForwarded is a free log subscription operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000ff5afd44.
//
// Solidity: event FileForwarded(address indexed from, address indexed to, address indexed proxy, bytes32 fileHash)
func (_FileHub *FileHubFilterer) WatchFileForwarded(fromBlock *uint64, handler func(int, []types.Log), from common.Address, to common.Address, proxy common.Address) error {
	return _FileHub.contract.WatchLogs(fromBlock, handler, "FileForwarded", from, to, proxy)
}

func (_FileHub *FileHubFilterer) WatchAllFileForwarded(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _FileHub.contract.WatchLogs(fromBlock, handler, "FileForwarded")
}

// ParseFileForwarded is a log parse operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000ff5afd44.
//
// Solidity: event FileForwarded(address indexed from, address indexed to, address indexed proxy, bytes32 fileHash)
func (_FileHub *FileHubFilterer) ParseFileForwarded(log types.Log) (*FileHubFileForwarded, error) {
	event := new(FileHubFileForwarded)
	if err := _FileHub.contract.UnpackLog(event, "FileForwarded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchFileForwarded is a free log subscription operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000ff5afd44.
//
// Solidity: event FileForwarded(address indexed from, address indexed to, address indexed proxy, bytes32 fileHash)
func (_FileHub *FileHubSession) WatchFileForwarded(fromBlock *uint64, handler func(int, []types.Log), from common.Address, to common.Address, proxy common.Address) error {
	return _FileHub.Contract.WatchFileForwarded(fromBlock, handler, from, to, proxy)
}

func (_FileHub *FileHubSession) WatchAllFileForwarded(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _FileHub.Contract.WatchAllFileForwarded(fromBlock, handler)
}

// ParseFileForwarded is a log parse operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000ff5afd44.
//
// Solidity: event FileForwarded(address indexed from, address indexed to, address indexed proxy, bytes32 fileHash)
func (_FileHub *FileHubSession) ParseFileForwarded(log types.Log) (*FileHubFileForwarded, error) {
	return _FileHub.Contract.ParseFileForwarded(log)
}

// FileHubFileShared represents a FileShared event raised by the FileHub contract.
type FileHubFileShared struct {
	From     common.Address
	To       common.Address
	FileHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// WatchFileShared is a free log subscription operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000f86c0599.
//
// Solidity: event FileShared(address indexed from, address indexed to, bytes32 indexed fileHash)
func (_FileHub *FileHubFilterer) WatchFileShared(fromBlock *uint64, handler func(int, []types.Log), from common.Address, to common.Address, fileHash [32]byte) error {
	return _FileHub.contract.WatchLogs(fromBlock, handler, "FileShared", from, to, fileHash)
}

func (_FileHub *FileHubFilterer) WatchAllFileShared(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _FileHub.contract.WatchLogs(fromBlock, handler, "FileShared")
}

// ParseFileShared is a log parse operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000f86c0599.
//
// Solidity: event FileShared(address indexed from, address indexed to, bytes32 indexed fileHash)
func (_FileHub *FileHubFilterer) ParseFileShared(log types.Log) (*FileHubFileShared, error) {
	event := new(FileHubFileShared)
	if err := _FileHub.contract.UnpackLog(event, "FileShared", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchFileShared is a free log subscription operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000f86c0599.
//
// Solidity: event FileShared(address indexed from, address indexed to, bytes32 indexed fileHash)
func (_FileHub *FileHubSession) WatchFileShared(fromBlock *uint64, handler func(int, []types.Log), from common.Address, to common.Address, fileHash [32]byte) error {
	return _FileHub.Contract.WatchFileShared(fromBlock, handler, from, to, fileHash)
}

func (_FileHub *FileHubSession) WatchAllFileShared(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _FileHub.Contract.WatchAllFileShared(fromBlock, handler)
}

// ParseFileShared is a log parse operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000f86c0599.
//
// Solidity: event FileShared(address indexed from, address indexed to, bytes32 indexed fileHash)
func (_FileHub *FileHubSession) ParseFileShared(log types.Log) (*FileHubFileShared, error) {
	return _FileHub.Contract.ParseFileShared(log)
}

// FileHubFileUploaded represents a FileUploaded event raised by the FileHub contract.
type FileHubFileUploaded struct {
	Account  common.Address
	FileHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// WatchFileUploaded is a free log subscription operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000a0e5f045.
//
// Solidity: event FileUploaded(address indexed account, bytes32 indexed fileHash)
func (_FileHub *FileHubFilterer) WatchFileUploaded(fromBlock *uint64, handler func(int, []types.Log), account common.Address, fileHash [32]byte) error {
	return _FileHub.contract.WatchLogs(fromBlock, handler, "FileUploaded", account, fileHash)
}

func (_FileHub *FileHubFilterer) WatchAllFileUploaded(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _FileHub.contract.WatchLogs(fromBlock, handler, "FileUploaded")
}

// ParseFileUploaded is a log parse operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000a0e5f045.
//
// Solidity: event FileUploaded(address indexed account, bytes32 indexed fileHash)
func (_FileHub *FileHubFilterer) ParseFileUploaded(log types.Log) (*FileHubFileUploaded, error) {
	event := new(FileHubFileUploaded)
	if err := _FileHub.contract.UnpackLog(event, "FileUploaded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchFileUploaded is a free log subscription operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000a0e5f045.
//
// Solidity: event FileUploaded(address indexed account, bytes32 indexed fileHash)
func (_FileHub *FileHubSession) WatchFileUploaded(fromBlock *uint64, handler func(int, []types.Log), account common.Address, fileHash [32]byte) error {
	return _FileHub.Contract.WatchFileUploaded(fromBlock, handler, account, fileHash)
}

func (_FileHub *FileHubSession) WatchAllFileUploaded(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _FileHub.Contract.WatchAllFileUploaded(fromBlock, handler)
}

// ParseFileUploaded is a log parse operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000a0e5f045.
//
// Solidity: event FileUploaded(address indexed account, bytes32 indexed fileHash)
func (_FileHub *FileHubSession) ParseFileUploaded(log types.Log) (*FileHubFileUploaded, error) {
	return _FileHub.Contract.ParseFileUploaded(log)
}

// FileHubRegistered represents a Registered event raised by the FileHub contract.
type FileHubRegistered struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// WatchRegistered is a free log subscription operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000a6419914.
//
// Solidity: event Registered(address indexed account)
func (_FileHub *FileHubFilterer) WatchRegistered(fromBlock *uint64, handler func(int, []types.Log), account common.Address) error {
	return _FileHub.contract.WatchLogs(fromBlock, handler, "Registered", account)
}

func (_FileHub *FileHubFilterer) WatchAllRegistered(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _FileHub.contract.WatchLogs(fromBlock, handler, "Registered")
}

// ParseRegistered is a log parse operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000a6419914.
//
// Solidity: event Registered(address indexed account)
func (_FileHub *FileHubFilterer) ParseRegistered(log types.Log) (*FileHubRegistered, error) {
	event := new(FileHubRegistered)
	if err := _FileHub.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000a6419914.
//
// Solidity: event Registered(address indexed account)
func (_FileHub *FileHubSession) WatchRegistered(fromBlock *uint64, handler func(int, []types.Log), account common.Address) error {
	return _FileHub.Contract.WatchRegistered(fromBlock, handler, account)
}

func (_FileHub *FileHubSession) WatchAllRegistered(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _FileHub.Contract.WatchAllRegistered(fromBlock, handler)
}

// ParseRegistered is a log parse operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000a6419914.
//
// Solidity: event Registered(address indexed account)
func (_FileHub *FileHubSession) ParseRegistered(log types.Log) (*FileHubRegistered, error) {
	return _FileHub.Contract.ParseRegistered(log)
}
