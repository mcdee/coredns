// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dnsresolvercontract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// DnsResolverContractABI is the input ABI used to generate the binding from.
const DnsResolverContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"resource\",\"type\":\"uint16\"},{\"name\":\"soaData\",\"type\":\"bytes\"}],\"name\":\"clearDnsRecord\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"records\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setText\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"contentTypes\",\"type\":\"uint256\"}],\"name\":\"ABI\",\"outputs\":[{\"name\":\"contentType\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setDnsZone\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"x\",\"type\":\"bytes32\"},{\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"setPubkey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"content\",\"outputs\":[{\"name\":\"ret\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"name\":\"ret\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"hasDnsRecords\",\"outputs\":[{\"name\":\"hasRecords\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"text\",\"outputs\":[{\"name\":\"ret\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"contentType\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setABI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"name\",\"outputs\":[{\"name\":\"ret\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"zones\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"clearDnsZone\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nameEntriesCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"resource\",\"type\":\"uint16\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"soaData\",\"type\":\"bytes\"}],\"name\":\"setDnsRecord\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"resource\",\"type\":\"uint16\"}],\"name\":\"dnsRecord\",\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"hasDnsZone\",\"outputs\":[{\"name\":\"hasRecords\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"soaRecords\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"setContent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"pubkey\",\"outputs\":[{\"name\":\"x\",\"type\":\"bytes32\"},{\"name\":\"y\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"dnsZone\",\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_registry\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddrChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ContentChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"contentType\",\"type\":\"uint256\"}],\"name\":\"ABIChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"x\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"PubkeyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"key\",\"type\":\"string\"}],\"name\":\"TextChanged\",\"type\":\"event\"}]"

// DnsResolverContract is an auto generated Go binding around an Ethereum contract.
type DnsResolverContract struct {
	DnsResolverContractCaller     // Read-only binding to the contract
	DnsResolverContractTransactor // Write-only binding to the contract
	DnsResolverContractFilterer   // Log filterer for contract events
}

// DnsResolverContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DnsResolverContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsResolverContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DnsResolverContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsResolverContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DnsResolverContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsResolverContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DnsResolverContractSession struct {
	Contract     *DnsResolverContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DnsResolverContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DnsResolverContractCallerSession struct {
	Contract *DnsResolverContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// DnsResolverContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DnsResolverContractTransactorSession struct {
	Contract     *DnsResolverContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// DnsResolverContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DnsResolverContractRaw struct {
	Contract *DnsResolverContract // Generic contract binding to access the raw methods on
}

// DnsResolverContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DnsResolverContractCallerRaw struct {
	Contract *DnsResolverContractCaller // Generic read-only contract binding to access the raw methods on
}

// DnsResolverContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DnsResolverContractTransactorRaw struct {
	Contract *DnsResolverContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDnsResolverContract creates a new instance of DnsResolverContract, bound to a specific deployed contract.
func NewDnsResolverContract(address common.Address, backend bind.ContractBackend) (*DnsResolverContract, error) {
	contract, err := bindDnsResolverContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DnsResolverContract{DnsResolverContractCaller: DnsResolverContractCaller{contract: contract}, DnsResolverContractTransactor: DnsResolverContractTransactor{contract: contract}, DnsResolverContractFilterer: DnsResolverContractFilterer{contract: contract}}, nil
}

// NewDnsResolverContractCaller creates a new read-only instance of DnsResolverContract, bound to a specific deployed contract.
func NewDnsResolverContractCaller(address common.Address, caller bind.ContractCaller) (*DnsResolverContractCaller, error) {
	contract, err := bindDnsResolverContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DnsResolverContractCaller{contract: contract}, nil
}

// NewDnsResolverContractTransactor creates a new write-only instance of DnsResolverContract, bound to a specific deployed contract.
func NewDnsResolverContractTransactor(address common.Address, transactor bind.ContractTransactor) (*DnsResolverContractTransactor, error) {
	contract, err := bindDnsResolverContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DnsResolverContractTransactor{contract: contract}, nil
}

// NewDnsResolverContractFilterer creates a new log filterer instance of DnsResolverContract, bound to a specific deployed contract.
func NewDnsResolverContractFilterer(address common.Address, filterer bind.ContractFilterer) (*DnsResolverContractFilterer, error) {
	contract, err := bindDnsResolverContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DnsResolverContractFilterer{contract: contract}, nil
}

// bindDnsResolverContract binds a generic wrapper to an already deployed contract.
func bindDnsResolverContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DnsResolverContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsResolverContract *DnsResolverContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DnsResolverContract.Contract.DnsResolverContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsResolverContract *DnsResolverContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.DnsResolverContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsResolverContract *DnsResolverContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.DnsResolverContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsResolverContract *DnsResolverContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DnsResolverContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsResolverContract *DnsResolverContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsResolverContract *DnsResolverContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.contract.Transact(opts, method, params...)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(node bytes32, contentTypes uint256) constant returns(contentType uint256, data bytes)
func (_DnsResolverContract *DnsResolverContractCaller) ABI(opts *bind.CallOpts, node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	ret := new(struct {
		ContentType *big.Int
		Data        []byte
	})
	out := ret
	err := _DnsResolverContract.contract.Call(opts, out, "ABI", node, contentTypes)
	return *ret, err
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(node bytes32, contentTypes uint256) constant returns(contentType uint256, data bytes)
func (_DnsResolverContract *DnsResolverContractSession) ABI(node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	return _DnsResolverContract.Contract.ABI(&_DnsResolverContract.CallOpts, node, contentTypes)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(node bytes32, contentTypes uint256) constant returns(contentType uint256, data bytes)
func (_DnsResolverContract *DnsResolverContractCallerSession) ABI(node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	return _DnsResolverContract.Contract.ABI(&_DnsResolverContract.CallOpts, node, contentTypes)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(ret address)
func (_DnsResolverContract *DnsResolverContractCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DnsResolverContract.contract.Call(opts, out, "addr", node)
	return *ret0, err
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(ret address)
func (_DnsResolverContract *DnsResolverContractSession) Addr(node [32]byte) (common.Address, error) {
	return _DnsResolverContract.Contract.Addr(&_DnsResolverContract.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(ret address)
func (_DnsResolverContract *DnsResolverContractCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _DnsResolverContract.Contract.Addr(&_DnsResolverContract.CallOpts, node)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(ret bytes32)
func (_DnsResolverContract *DnsResolverContractCaller) Content(opts *bind.CallOpts, node [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DnsResolverContract.contract.Call(opts, out, "content", node)
	return *ret0, err
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(ret bytes32)
func (_DnsResolverContract *DnsResolverContractSession) Content(node [32]byte) ([32]byte, error) {
	return _DnsResolverContract.Contract.Content(&_DnsResolverContract.CallOpts, node)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(ret bytes32)
func (_DnsResolverContract *DnsResolverContractCallerSession) Content(node [32]byte) ([32]byte, error) {
	return _DnsResolverContract.Contract.Content(&_DnsResolverContract.CallOpts, node)
}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Solidity: function dnsRecord(node bytes32, name bytes32, resource uint16) constant returns(data bytes)
func (_DnsResolverContract *DnsResolverContractCaller) DnsRecord(opts *bind.CallOpts, node [32]byte, name [32]byte, resource uint16) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _DnsResolverContract.contract.Call(opts, out, "dnsRecord", node, name, resource)
	return *ret0, err
}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Solidity: function dnsRecord(node bytes32, name bytes32, resource uint16) constant returns(data bytes)
func (_DnsResolverContract *DnsResolverContractSession) DnsRecord(node [32]byte, name [32]byte, resource uint16) ([]byte, error) {
	return _DnsResolverContract.Contract.DnsRecord(&_DnsResolverContract.CallOpts, node, name, resource)
}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Solidity: function dnsRecord(node bytes32, name bytes32, resource uint16) constant returns(data bytes)
func (_DnsResolverContract *DnsResolverContractCallerSession) DnsRecord(node [32]byte, name [32]byte, resource uint16) ([]byte, error) {
	return _DnsResolverContract.Contract.DnsRecord(&_DnsResolverContract.CallOpts, node, name, resource)
}

// DnsZone is a free data retrieval call binding the contract method 0xdbfc5d00.
//
// Solidity: function dnsZone(node bytes32) constant returns(data bytes)
func (_DnsResolverContract *DnsResolverContractCaller) DnsZone(opts *bind.CallOpts, node [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _DnsResolverContract.contract.Call(opts, out, "dnsZone", node)
	return *ret0, err
}

// DnsZone is a free data retrieval call binding the contract method 0xdbfc5d00.
//
// Solidity: function dnsZone(node bytes32) constant returns(data bytes)
func (_DnsResolverContract *DnsResolverContractSession) DnsZone(node [32]byte) ([]byte, error) {
	return _DnsResolverContract.Contract.DnsZone(&_DnsResolverContract.CallOpts, node)
}

// DnsZone is a free data retrieval call binding the contract method 0xdbfc5d00.
//
// Solidity: function dnsZone(node bytes32) constant returns(data bytes)
func (_DnsResolverContract *DnsResolverContractCallerSession) DnsZone(node [32]byte) ([]byte, error) {
	return _DnsResolverContract.Contract.DnsZone(&_DnsResolverContract.CallOpts, node)
}

// HasDnsRecords is a free data retrieval call binding the contract method 0x4cac6495.
//
// Solidity: function hasDnsRecords(node bytes32, name bytes32) constant returns(hasRecords bool)
func (_DnsResolverContract *DnsResolverContractCaller) HasDnsRecords(opts *bind.CallOpts, node [32]byte, name [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DnsResolverContract.contract.Call(opts, out, "hasDnsRecords", node, name)
	return *ret0, err
}

// HasDnsRecords is a free data retrieval call binding the contract method 0x4cac6495.
//
// Solidity: function hasDnsRecords(node bytes32, name bytes32) constant returns(hasRecords bool)
func (_DnsResolverContract *DnsResolverContractSession) HasDnsRecords(node [32]byte, name [32]byte) (bool, error) {
	return _DnsResolverContract.Contract.HasDnsRecords(&_DnsResolverContract.CallOpts, node, name)
}

// HasDnsRecords is a free data retrieval call binding the contract method 0x4cac6495.
//
// Solidity: function hasDnsRecords(node bytes32, name bytes32) constant returns(hasRecords bool)
func (_DnsResolverContract *DnsResolverContractCallerSession) HasDnsRecords(node [32]byte, name [32]byte) (bool, error) {
	return _DnsResolverContract.Contract.HasDnsRecords(&_DnsResolverContract.CallOpts, node, name)
}

// HasDnsZone is a free data retrieval call binding the contract method 0xa956cbbc.
//
// Solidity: function hasDnsZone(node bytes32) constant returns(hasRecords bool)
func (_DnsResolverContract *DnsResolverContractCaller) HasDnsZone(opts *bind.CallOpts, node [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DnsResolverContract.contract.Call(opts, out, "hasDnsZone", node)
	return *ret0, err
}

// HasDnsZone is a free data retrieval call binding the contract method 0xa956cbbc.
//
// Solidity: function hasDnsZone(node bytes32) constant returns(hasRecords bool)
func (_DnsResolverContract *DnsResolverContractSession) HasDnsZone(node [32]byte) (bool, error) {
	return _DnsResolverContract.Contract.HasDnsZone(&_DnsResolverContract.CallOpts, node)
}

// HasDnsZone is a free data retrieval call binding the contract method 0xa956cbbc.
//
// Solidity: function hasDnsZone(node bytes32) constant returns(hasRecords bool)
func (_DnsResolverContract *DnsResolverContractCallerSession) HasDnsZone(node [32]byte) (bool, error) {
	return _DnsResolverContract.Contract.HasDnsZone(&_DnsResolverContract.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(node bytes32) constant returns(ret string)
func (_DnsResolverContract *DnsResolverContractCaller) Name(opts *bind.CallOpts, node [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DnsResolverContract.contract.Call(opts, out, "name", node)
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(node bytes32) constant returns(ret string)
func (_DnsResolverContract *DnsResolverContractSession) Name(node [32]byte) (string, error) {
	return _DnsResolverContract.Contract.Name(&_DnsResolverContract.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(node bytes32) constant returns(ret string)
func (_DnsResolverContract *DnsResolverContractCallerSession) Name(node [32]byte) (string, error) {
	return _DnsResolverContract.Contract.Name(&_DnsResolverContract.CallOpts, node)
}

// NameEntriesCount is a free data retrieval call binding the contract method 0x932fc0e0.
//
// Solidity: function nameEntriesCount( bytes32,  bytes32) constant returns(uint16)
func (_DnsResolverContract *DnsResolverContractCaller) NameEntriesCount(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _DnsResolverContract.contract.Call(opts, out, "nameEntriesCount", arg0, arg1)
	return *ret0, err
}

// NameEntriesCount is a free data retrieval call binding the contract method 0x932fc0e0.
//
// Solidity: function nameEntriesCount( bytes32,  bytes32) constant returns(uint16)
func (_DnsResolverContract *DnsResolverContractSession) NameEntriesCount(arg0 [32]byte, arg1 [32]byte) (uint16, error) {
	return _DnsResolverContract.Contract.NameEntriesCount(&_DnsResolverContract.CallOpts, arg0, arg1)
}

// NameEntriesCount is a free data retrieval call binding the contract method 0x932fc0e0.
//
// Solidity: function nameEntriesCount( bytes32,  bytes32) constant returns(uint16)
func (_DnsResolverContract *DnsResolverContractCallerSession) NameEntriesCount(arg0 [32]byte, arg1 [32]byte) (uint16, error) {
	return _DnsResolverContract.Contract.NameEntriesCount(&_DnsResolverContract.CallOpts, arg0, arg1)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(node bytes32) constant returns(x bytes32, y bytes32)
func (_DnsResolverContract *DnsResolverContractCaller) Pubkey(opts *bind.CallOpts, node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	ret := new(struct {
		X [32]byte
		Y [32]byte
	})
	out := ret
	err := _DnsResolverContract.contract.Call(opts, out, "pubkey", node)
	return *ret, err
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(node bytes32) constant returns(x bytes32, y bytes32)
func (_DnsResolverContract *DnsResolverContractSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _DnsResolverContract.Contract.Pubkey(&_DnsResolverContract.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(node bytes32) constant returns(x bytes32, y bytes32)
func (_DnsResolverContract *DnsResolverContractCallerSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _DnsResolverContract.Contract.Pubkey(&_DnsResolverContract.CallOpts, node)
}

// Records is a free data retrieval call binding the contract method 0x107de931.
//
// Solidity: function records( bytes32,  bytes32,  uint16) constant returns(bytes)
func (_DnsResolverContract *DnsResolverContractCaller) Records(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte, arg2 uint16) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _DnsResolverContract.contract.Call(opts, out, "records", arg0, arg1, arg2)
	return *ret0, err
}

// Records is a free data retrieval call binding the contract method 0x107de931.
//
// Solidity: function records( bytes32,  bytes32,  uint16) constant returns(bytes)
func (_DnsResolverContract *DnsResolverContractSession) Records(arg0 [32]byte, arg1 [32]byte, arg2 uint16) ([]byte, error) {
	return _DnsResolverContract.Contract.Records(&_DnsResolverContract.CallOpts, arg0, arg1, arg2)
}

// Records is a free data retrieval call binding the contract method 0x107de931.
//
// Solidity: function records( bytes32,  bytes32,  uint16) constant returns(bytes)
func (_DnsResolverContract *DnsResolverContractCallerSession) Records(arg0 [32]byte, arg1 [32]byte, arg2 uint16) ([]byte, error) {
	return _DnsResolverContract.Contract.Records(&_DnsResolverContract.CallOpts, arg0, arg1, arg2)
}

// SoaRecords is a free data retrieval call binding the contract method 0xc37e3729.
//
// Solidity: function soaRecords( bytes32) constant returns(bytes)
func (_DnsResolverContract *DnsResolverContractCaller) SoaRecords(opts *bind.CallOpts, arg0 [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _DnsResolverContract.contract.Call(opts, out, "soaRecords", arg0)
	return *ret0, err
}

// SoaRecords is a free data retrieval call binding the contract method 0xc37e3729.
//
// Solidity: function soaRecords( bytes32) constant returns(bytes)
func (_DnsResolverContract *DnsResolverContractSession) SoaRecords(arg0 [32]byte) ([]byte, error) {
	return _DnsResolverContract.Contract.SoaRecords(&_DnsResolverContract.CallOpts, arg0)
}

// SoaRecords is a free data retrieval call binding the contract method 0xc37e3729.
//
// Solidity: function soaRecords( bytes32) constant returns(bytes)
func (_DnsResolverContract *DnsResolverContractCallerSession) SoaRecords(arg0 [32]byte) ([]byte, error) {
	return _DnsResolverContract.Contract.SoaRecords(&_DnsResolverContract.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceId bytes4) constant returns(bool)
func (_DnsResolverContract *DnsResolverContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DnsResolverContract.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceId bytes4) constant returns(bool)
func (_DnsResolverContract *DnsResolverContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DnsResolverContract.Contract.SupportsInterface(&_DnsResolverContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceId bytes4) constant returns(bool)
func (_DnsResolverContract *DnsResolverContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DnsResolverContract.Contract.SupportsInterface(&_DnsResolverContract.CallOpts, interfaceId)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(node bytes32, key string) constant returns(ret string)
func (_DnsResolverContract *DnsResolverContractCaller) Text(opts *bind.CallOpts, node [32]byte, key string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DnsResolverContract.contract.Call(opts, out, "text", node, key)
	return *ret0, err
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(node bytes32, key string) constant returns(ret string)
func (_DnsResolverContract *DnsResolverContractSession) Text(node [32]byte, key string) (string, error) {
	return _DnsResolverContract.Contract.Text(&_DnsResolverContract.CallOpts, node, key)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(node bytes32, key string) constant returns(ret string)
func (_DnsResolverContract *DnsResolverContractCallerSession) Text(node [32]byte, key string) (string, error) {
	return _DnsResolverContract.Contract.Text(&_DnsResolverContract.CallOpts, node, key)
}

// Zones is a free data retrieval call binding the contract method 0x6de75ff8.
//
// Solidity: function zones( bytes32) constant returns(bytes)
func (_DnsResolverContract *DnsResolverContractCaller) Zones(opts *bind.CallOpts, arg0 [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _DnsResolverContract.contract.Call(opts, out, "zones", arg0)
	return *ret0, err
}

// Zones is a free data retrieval call binding the contract method 0x6de75ff8.
//
// Solidity: function zones( bytes32) constant returns(bytes)
func (_DnsResolverContract *DnsResolverContractSession) Zones(arg0 [32]byte) ([]byte, error) {
	return _DnsResolverContract.Contract.Zones(&_DnsResolverContract.CallOpts, arg0)
}

// Zones is a free data retrieval call binding the contract method 0x6de75ff8.
//
// Solidity: function zones( bytes32) constant returns(bytes)
func (_DnsResolverContract *DnsResolverContractCallerSession) Zones(arg0 [32]byte) ([]byte, error) {
	return _DnsResolverContract.Contract.Zones(&_DnsResolverContract.CallOpts, arg0)
}

// ClearDnsRecord is a paid mutator transaction binding the contract method 0x0a9ac0d2.
//
// Solidity: function clearDnsRecord(node bytes32, name bytes32, resource uint16, soaData bytes) returns()
func (_DnsResolverContract *DnsResolverContractTransactor) ClearDnsRecord(opts *bind.TransactOpts, node [32]byte, name [32]byte, resource uint16, soaData []byte) (*types.Transaction, error) {
	return _DnsResolverContract.contract.Transact(opts, "clearDnsRecord", node, name, resource, soaData)
}

// ClearDnsRecord is a paid mutator transaction binding the contract method 0x0a9ac0d2.
//
// Solidity: function clearDnsRecord(node bytes32, name bytes32, resource uint16, soaData bytes) returns()
func (_DnsResolverContract *DnsResolverContractSession) ClearDnsRecord(node [32]byte, name [32]byte, resource uint16, soaData []byte) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.ClearDnsRecord(&_DnsResolverContract.TransactOpts, node, name, resource, soaData)
}

// ClearDnsRecord is a paid mutator transaction binding the contract method 0x0a9ac0d2.
//
// Solidity: function clearDnsRecord(node bytes32, name bytes32, resource uint16, soaData bytes) returns()
func (_DnsResolverContract *DnsResolverContractTransactorSession) ClearDnsRecord(node [32]byte, name [32]byte, resource uint16, soaData []byte) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.ClearDnsRecord(&_DnsResolverContract.TransactOpts, node, name, resource, soaData)
}

// ClearDnsZone is a paid mutator transaction binding the contract method 0x737aaf74.
//
// Solidity: function clearDnsZone(node bytes32) returns()
func (_DnsResolverContract *DnsResolverContractTransactor) ClearDnsZone(opts *bind.TransactOpts, node [32]byte) (*types.Transaction, error) {
	return _DnsResolverContract.contract.Transact(opts, "clearDnsZone", node)
}

// ClearDnsZone is a paid mutator transaction binding the contract method 0x737aaf74.
//
// Solidity: function clearDnsZone(node bytes32) returns()
func (_DnsResolverContract *DnsResolverContractSession) ClearDnsZone(node [32]byte) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.ClearDnsZone(&_DnsResolverContract.TransactOpts, node)
}

// ClearDnsZone is a paid mutator transaction binding the contract method 0x737aaf74.
//
// Solidity: function clearDnsZone(node bytes32) returns()
func (_DnsResolverContract *DnsResolverContractTransactorSession) ClearDnsZone(node [32]byte) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.ClearDnsZone(&_DnsResolverContract.TransactOpts, node)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(node bytes32, contentType uint256, data bytes) returns()
func (_DnsResolverContract *DnsResolverContractTransactor) SetABI(opts *bind.TransactOpts, node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _DnsResolverContract.contract.Transact(opts, "setABI", node, contentType, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(node bytes32, contentType uint256, data bytes) returns()
func (_DnsResolverContract *DnsResolverContractSession) SetABI(node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetABI(&_DnsResolverContract.TransactOpts, node, contentType, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(node bytes32, contentType uint256, data bytes) returns()
func (_DnsResolverContract *DnsResolverContractTransactorSession) SetABI(node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetABI(&_DnsResolverContract.TransactOpts, node, contentType, data)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(node bytes32, addr address) returns()
func (_DnsResolverContract *DnsResolverContractTransactor) SetAddr(opts *bind.TransactOpts, node [32]byte, addr common.Address) (*types.Transaction, error) {
	return _DnsResolverContract.contract.Transact(opts, "setAddr", node, addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(node bytes32, addr address) returns()
func (_DnsResolverContract *DnsResolverContractSession) SetAddr(node [32]byte, addr common.Address) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetAddr(&_DnsResolverContract.TransactOpts, node, addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(node bytes32, addr address) returns()
func (_DnsResolverContract *DnsResolverContractTransactorSession) SetAddr(node [32]byte, addr common.Address) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetAddr(&_DnsResolverContract.TransactOpts, node, addr)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(node bytes32, hash bytes32) returns()
func (_DnsResolverContract *DnsResolverContractTransactor) SetContent(opts *bind.TransactOpts, node [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _DnsResolverContract.contract.Transact(opts, "setContent", node, hash)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(node bytes32, hash bytes32) returns()
func (_DnsResolverContract *DnsResolverContractSession) SetContent(node [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetContent(&_DnsResolverContract.TransactOpts, node, hash)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(node bytes32, hash bytes32) returns()
func (_DnsResolverContract *DnsResolverContractTransactorSession) SetContent(node [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetContent(&_DnsResolverContract.TransactOpts, node, hash)
}

// SetDnsRecord is a paid mutator transaction binding the contract method 0x9a75194c.
//
// Solidity: function setDnsRecord(node bytes32, name bytes32, resource uint16, data bytes, soaData bytes) returns()
func (_DnsResolverContract *DnsResolverContractTransactor) SetDnsRecord(opts *bind.TransactOpts, node [32]byte, name [32]byte, resource uint16, data []byte, soaData []byte) (*types.Transaction, error) {
	return _DnsResolverContract.contract.Transact(opts, "setDnsRecord", node, name, resource, data, soaData)
}

// SetDnsRecord is a paid mutator transaction binding the contract method 0x9a75194c.
//
// Solidity: function setDnsRecord(node bytes32, name bytes32, resource uint16, data bytes, soaData bytes) returns()
func (_DnsResolverContract *DnsResolverContractSession) SetDnsRecord(node [32]byte, name [32]byte, resource uint16, data []byte, soaData []byte) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetDnsRecord(&_DnsResolverContract.TransactOpts, node, name, resource, data, soaData)
}

// SetDnsRecord is a paid mutator transaction binding the contract method 0x9a75194c.
//
// Solidity: function setDnsRecord(node bytes32, name bytes32, resource uint16, data bytes, soaData bytes) returns()
func (_DnsResolverContract *DnsResolverContractTransactorSession) SetDnsRecord(node [32]byte, name [32]byte, resource uint16, data []byte, soaData []byte) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetDnsRecord(&_DnsResolverContract.TransactOpts, node, name, resource, data, soaData)
}

// SetDnsZone is a paid mutator transaction binding the contract method 0x233a359c.
//
// Solidity: function setDnsZone(node bytes32, data bytes) returns()
func (_DnsResolverContract *DnsResolverContractTransactor) SetDnsZone(opts *bind.TransactOpts, node [32]byte, data []byte) (*types.Transaction, error) {
	return _DnsResolverContract.contract.Transact(opts, "setDnsZone", node, data)
}

// SetDnsZone is a paid mutator transaction binding the contract method 0x233a359c.
//
// Solidity: function setDnsZone(node bytes32, data bytes) returns()
func (_DnsResolverContract *DnsResolverContractSession) SetDnsZone(node [32]byte, data []byte) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetDnsZone(&_DnsResolverContract.TransactOpts, node, data)
}

// SetDnsZone is a paid mutator transaction binding the contract method 0x233a359c.
//
// Solidity: function setDnsZone(node bytes32, data bytes) returns()
func (_DnsResolverContract *DnsResolverContractTransactorSession) SetDnsZone(node [32]byte, data []byte) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetDnsZone(&_DnsResolverContract.TransactOpts, node, data)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(node bytes32, name string) returns()
func (_DnsResolverContract *DnsResolverContractTransactor) SetName(opts *bind.TransactOpts, node [32]byte, name string) (*types.Transaction, error) {
	return _DnsResolverContract.contract.Transact(opts, "setName", node, name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(node bytes32, name string) returns()
func (_DnsResolverContract *DnsResolverContractSession) SetName(node [32]byte, name string) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetName(&_DnsResolverContract.TransactOpts, node, name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(node bytes32, name string) returns()
func (_DnsResolverContract *DnsResolverContractTransactorSession) SetName(node [32]byte, name string) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetName(&_DnsResolverContract.TransactOpts, node, name)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(node bytes32, x bytes32, y bytes32) returns()
func (_DnsResolverContract *DnsResolverContractTransactor) SetPubkey(opts *bind.TransactOpts, node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _DnsResolverContract.contract.Transact(opts, "setPubkey", node, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(node bytes32, x bytes32, y bytes32) returns()
func (_DnsResolverContract *DnsResolverContractSession) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetPubkey(&_DnsResolverContract.TransactOpts, node, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(node bytes32, x bytes32, y bytes32) returns()
func (_DnsResolverContract *DnsResolverContractTransactorSession) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetPubkey(&_DnsResolverContract.TransactOpts, node, x, y)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(node bytes32, key string, value string) returns()
func (_DnsResolverContract *DnsResolverContractTransactor) SetText(opts *bind.TransactOpts, node [32]byte, key string, value string) (*types.Transaction, error) {
	return _DnsResolverContract.contract.Transact(opts, "setText", node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(node bytes32, key string, value string) returns()
func (_DnsResolverContract *DnsResolverContractSession) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetText(&_DnsResolverContract.TransactOpts, node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(node bytes32, key string, value string) returns()
func (_DnsResolverContract *DnsResolverContractTransactorSession) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetText(&_DnsResolverContract.TransactOpts, node, key, value)
}

// DnsResolverContractABIChangedIterator is returned from FilterABIChanged and is used to iterate over the raw logs and unpacked data for ABIChanged events raised by the DnsResolverContract contract.
type DnsResolverContractABIChangedIterator struct {
	Event *DnsResolverContractABIChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DnsResolverContractABIChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsResolverContractABIChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DnsResolverContractABIChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DnsResolverContractABIChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsResolverContractABIChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsResolverContractABIChanged represents a ABIChanged event raised by the DnsResolverContract contract.
type DnsResolverContractABIChanged struct {
	Node        [32]byte
	ContentType *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterABIChanged is a free log retrieval operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: event ABIChanged(node indexed bytes32, contentType indexed uint256)
func (_DnsResolverContract *DnsResolverContractFilterer) FilterABIChanged(opts *bind.FilterOpts, node [][32]byte, contentType []*big.Int) (*DnsResolverContractABIChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _DnsResolverContract.contract.FilterLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return &DnsResolverContractABIChangedIterator{contract: _DnsResolverContract.contract, event: "ABIChanged", logs: logs, sub: sub}, nil
}

// WatchABIChanged is a free log subscription operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: event ABIChanged(node indexed bytes32, contentType indexed uint256)
func (_DnsResolverContract *DnsResolverContractFilterer) WatchABIChanged(opts *bind.WatchOpts, sink chan<- *DnsResolverContractABIChanged, node [][32]byte, contentType []*big.Int) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _DnsResolverContract.contract.WatchLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsResolverContractABIChanged)
				if err := _DnsResolverContract.contract.UnpackLog(event, "ABIChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DnsResolverContractAddrChangedIterator is returned from FilterAddrChanged and is used to iterate over the raw logs and unpacked data for AddrChanged events raised by the DnsResolverContract contract.
type DnsResolverContractAddrChangedIterator struct {
	Event *DnsResolverContractAddrChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DnsResolverContractAddrChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsResolverContractAddrChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DnsResolverContractAddrChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DnsResolverContractAddrChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsResolverContractAddrChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsResolverContractAddrChanged represents a AddrChanged event raised by the DnsResolverContract contract.
type DnsResolverContractAddrChanged struct {
	Node [32]byte
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddrChanged is a free log retrieval operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(node indexed bytes32, a address)
func (_DnsResolverContract *DnsResolverContractFilterer) FilterAddrChanged(opts *bind.FilterOpts, node [][32]byte) (*DnsResolverContractAddrChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _DnsResolverContract.contract.FilterLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &DnsResolverContractAddrChangedIterator{contract: _DnsResolverContract.contract, event: "AddrChanged", logs: logs, sub: sub}, nil
}

// WatchAddrChanged is a free log subscription operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(node indexed bytes32, a address)
func (_DnsResolverContract *DnsResolverContractFilterer) WatchAddrChanged(opts *bind.WatchOpts, sink chan<- *DnsResolverContractAddrChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _DnsResolverContract.contract.WatchLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsResolverContractAddrChanged)
				if err := _DnsResolverContract.contract.UnpackLog(event, "AddrChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DnsResolverContractContentChangedIterator is returned from FilterContentChanged and is used to iterate over the raw logs and unpacked data for ContentChanged events raised by the DnsResolverContract contract.
type DnsResolverContractContentChangedIterator struct {
	Event *DnsResolverContractContentChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DnsResolverContractContentChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsResolverContractContentChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DnsResolverContractContentChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DnsResolverContractContentChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsResolverContractContentChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsResolverContractContentChanged represents a ContentChanged event raised by the DnsResolverContract contract.
type DnsResolverContractContentChanged struct {
	Node [32]byte
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterContentChanged is a free log retrieval operation binding the contract event 0x0424b6fe0d9c3bdbece0e7879dc241bb0c22e900be8b6c168b4ee08bd9bf83bc.
//
// Solidity: event ContentChanged(node indexed bytes32, hash bytes32)
func (_DnsResolverContract *DnsResolverContractFilterer) FilterContentChanged(opts *bind.FilterOpts, node [][32]byte) (*DnsResolverContractContentChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _DnsResolverContract.contract.FilterLogs(opts, "ContentChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &DnsResolverContractContentChangedIterator{contract: _DnsResolverContract.contract, event: "ContentChanged", logs: logs, sub: sub}, nil
}

// WatchContentChanged is a free log subscription operation binding the contract event 0x0424b6fe0d9c3bdbece0e7879dc241bb0c22e900be8b6c168b4ee08bd9bf83bc.
//
// Solidity: event ContentChanged(node indexed bytes32, hash bytes32)
func (_DnsResolverContract *DnsResolverContractFilterer) WatchContentChanged(opts *bind.WatchOpts, sink chan<- *DnsResolverContractContentChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _DnsResolverContract.contract.WatchLogs(opts, "ContentChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsResolverContractContentChanged)
				if err := _DnsResolverContract.contract.UnpackLog(event, "ContentChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DnsResolverContractNameChangedIterator is returned from FilterNameChanged and is used to iterate over the raw logs and unpacked data for NameChanged events raised by the DnsResolverContract contract.
type DnsResolverContractNameChangedIterator struct {
	Event *DnsResolverContractNameChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DnsResolverContractNameChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsResolverContractNameChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DnsResolverContractNameChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DnsResolverContractNameChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsResolverContractNameChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsResolverContractNameChanged represents a NameChanged event raised by the DnsResolverContract contract.
type DnsResolverContractNameChanged struct {
	Node [32]byte
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNameChanged is a free log retrieval operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(node indexed bytes32, name string)
func (_DnsResolverContract *DnsResolverContractFilterer) FilterNameChanged(opts *bind.FilterOpts, node [][32]byte) (*DnsResolverContractNameChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _DnsResolverContract.contract.FilterLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &DnsResolverContractNameChangedIterator{contract: _DnsResolverContract.contract, event: "NameChanged", logs: logs, sub: sub}, nil
}

// WatchNameChanged is a free log subscription operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(node indexed bytes32, name string)
func (_DnsResolverContract *DnsResolverContractFilterer) WatchNameChanged(opts *bind.WatchOpts, sink chan<- *DnsResolverContractNameChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _DnsResolverContract.contract.WatchLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsResolverContractNameChanged)
				if err := _DnsResolverContract.contract.UnpackLog(event, "NameChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DnsResolverContractPubkeyChangedIterator is returned from FilterPubkeyChanged and is used to iterate over the raw logs and unpacked data for PubkeyChanged events raised by the DnsResolverContract contract.
type DnsResolverContractPubkeyChangedIterator struct {
	Event *DnsResolverContractPubkeyChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DnsResolverContractPubkeyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsResolverContractPubkeyChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DnsResolverContractPubkeyChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DnsResolverContractPubkeyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsResolverContractPubkeyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsResolverContractPubkeyChanged represents a PubkeyChanged event raised by the DnsResolverContract contract.
type DnsResolverContractPubkeyChanged struct {
	Node [32]byte
	X    [32]byte
	Y    [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPubkeyChanged is a free log retrieval operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: event PubkeyChanged(node indexed bytes32, x bytes32, y bytes32)
func (_DnsResolverContract *DnsResolverContractFilterer) FilterPubkeyChanged(opts *bind.FilterOpts, node [][32]byte) (*DnsResolverContractPubkeyChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _DnsResolverContract.contract.FilterLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &DnsResolverContractPubkeyChangedIterator{contract: _DnsResolverContract.contract, event: "PubkeyChanged", logs: logs, sub: sub}, nil
}

// WatchPubkeyChanged is a free log subscription operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: event PubkeyChanged(node indexed bytes32, x bytes32, y bytes32)
func (_DnsResolverContract *DnsResolverContractFilterer) WatchPubkeyChanged(opts *bind.WatchOpts, sink chan<- *DnsResolverContractPubkeyChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _DnsResolverContract.contract.WatchLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsResolverContractPubkeyChanged)
				if err := _DnsResolverContract.contract.UnpackLog(event, "PubkeyChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DnsResolverContractTextChangedIterator is returned from FilterTextChanged and is used to iterate over the raw logs and unpacked data for TextChanged events raised by the DnsResolverContract contract.
type DnsResolverContractTextChangedIterator struct {
	Event *DnsResolverContractTextChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DnsResolverContractTextChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsResolverContractTextChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DnsResolverContractTextChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DnsResolverContractTextChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsResolverContractTextChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsResolverContractTextChanged represents a TextChanged event raised by the DnsResolverContract contract.
type DnsResolverContractTextChanged struct {
	Node       [32]byte
	IndexedKey common.Hash
	Key        string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTextChanged is a free log retrieval operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(node indexed bytes32, indexedKey indexed string, key string)
func (_DnsResolverContract *DnsResolverContractFilterer) FilterTextChanged(opts *bind.FilterOpts, node [][32]byte, indexedKey []string) (*DnsResolverContractTextChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _DnsResolverContract.contract.FilterLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return &DnsResolverContractTextChangedIterator{contract: _DnsResolverContract.contract, event: "TextChanged", logs: logs, sub: sub}, nil
}

// WatchTextChanged is a free log subscription operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(node indexed bytes32, indexedKey indexed string, key string)
func (_DnsResolverContract *DnsResolverContractFilterer) WatchTextChanged(opts *bind.WatchOpts, sink chan<- *DnsResolverContractTextChanged, node [][32]byte, indexedKey []string) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _DnsResolverContract.contract.WatchLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsResolverContractTextChanged)
				if err := _DnsResolverContract.contract.UnpackLog(event, "TextChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
