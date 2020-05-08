// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package relayerincentivize

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RelayerincentivizeABI is the input ABI used to generate the binding from.
const RelayerincentivizeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"roundRewardForHeaderRelayer\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"roundRewardForTransferRelayer\",\"type\":\"uint256\"}],\"name\":\"LogDistributeCollectedReward\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"GOV_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INCENTIVIZE_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"LIGHT_CLIENT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAXIMUM_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RELAYERHUB_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RELAYER_REWARD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUND_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SLASH_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SYSTEM_REWARD_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKEN_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_collectedRewardForHeaderRelayer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_collectedRewardForTransferRelayer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_countInRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_headerRelayerAddressRecord\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_headerRelayersSubmitCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_roundSequence\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_transferRelayerAddressRecord\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_transferRelayersSubmitCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"headerRelayerAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"addReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"calculateHeaderRelayerWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"calculateTransferRelayerWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"denominaroeCallerCompensation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"denominaroeHeaderRelayer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"moleculeCallerCompensation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"moleculeHeaderRelayer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Relayerincentivize is an auto generated Go binding around an Ethereum contract.
type Relayerincentivize struct {
	RelayerincentivizeCaller     // Read-only binding to the contract
	RelayerincentivizeTransactor // Write-only binding to the contract
	RelayerincentivizeFilterer   // Log filterer for contract events
}

// RelayerincentivizeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayerincentivizeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerincentivizeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayerincentivizeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerincentivizeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayerincentivizeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerincentivizeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelayerincentivizeSession struct {
	Contract     *Relayerincentivize // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RelayerincentivizeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayerincentivizeCallerSession struct {
	Contract *RelayerincentivizeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// RelayerincentivizeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayerincentivizeTransactorSession struct {
	Contract     *RelayerincentivizeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// RelayerincentivizeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayerincentivizeRaw struct {
	Contract *Relayerincentivize // Generic contract binding to access the raw methods on
}

// RelayerincentivizeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayerincentivizeCallerRaw struct {
	Contract *RelayerincentivizeCaller // Generic read-only contract binding to access the raw methods on
}

// RelayerincentivizeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayerincentivizeTransactorRaw struct {
	Contract *RelayerincentivizeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelayerincentivize creates a new instance of Relayerincentivize, bound to a specific deployed contract.
func NewRelayerincentivize(address common.Address, backend bind.ContractBackend) (*Relayerincentivize, error) {
	contract, err := bindRelayerincentivize(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Relayerincentivize{RelayerincentivizeCaller: RelayerincentivizeCaller{contract: contract}, RelayerincentivizeTransactor: RelayerincentivizeTransactor{contract: contract}, RelayerincentivizeFilterer: RelayerincentivizeFilterer{contract: contract}}, nil
}

// NewRelayerincentivizeCaller creates a new read-only instance of Relayerincentivize, bound to a specific deployed contract.
func NewRelayerincentivizeCaller(address common.Address, caller bind.ContractCaller) (*RelayerincentivizeCaller, error) {
	contract, err := bindRelayerincentivize(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerincentivizeCaller{contract: contract}, nil
}

// NewRelayerincentivizeTransactor creates a new write-only instance of Relayerincentivize, bound to a specific deployed contract.
func NewRelayerincentivizeTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayerincentivizeTransactor, error) {
	contract, err := bindRelayerincentivize(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerincentivizeTransactor{contract: contract}, nil
}

// NewRelayerincentivizeFilterer creates a new log filterer instance of Relayerincentivize, bound to a specific deployed contract.
func NewRelayerincentivizeFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayerincentivizeFilterer, error) {
	contract, err := bindRelayerincentivize(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayerincentivizeFilterer{contract: contract}, nil
}

// bindRelayerincentivize binds a generic wrapper to an already deployed contract.
func bindRelayerincentivize(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayerincentivizeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relayerincentivize *RelayerincentivizeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Relayerincentivize.Contract.RelayerincentivizeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relayerincentivize *RelayerincentivizeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayerincentivize.Contract.RelayerincentivizeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relayerincentivize *RelayerincentivizeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relayerincentivize.Contract.RelayerincentivizeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relayerincentivize *RelayerincentivizeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Relayerincentivize.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relayerincentivize *RelayerincentivizeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayerincentivize.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relayerincentivize *RelayerincentivizeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relayerincentivize.Contract.contract.Transact(opts, method, params...)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCaller) GOVHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "GOV_HUB_ADDR")
	return *ret0, err
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeSession) GOVHUBADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.GOVHUBADDR(&_Relayerincentivize.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCallerSession) GOVHUBADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.GOVHUBADDR(&_Relayerincentivize.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCaller) INCENTIVIZEADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "INCENTIVIZE_ADDR")
	return *ret0, err
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.INCENTIVIZEADDR(&_Relayerincentivize.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCallerSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.INCENTIVIZEADDR(&_Relayerincentivize.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCaller) LIGHTCLIENTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "LIGHT_CLIENT_ADDR")
	return *ret0, err
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.LIGHTCLIENTADDR(&_Relayerincentivize.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCallerSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.LIGHTCLIENTADDR(&_Relayerincentivize.CallOpts)
}

// MAXIMUMWEIGHT is a free data retrieval call binding the contract method 0x08f2ec06.
//
// Solidity: function MAXIMUM_WEIGHT() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) MAXIMUMWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "MAXIMUM_WEIGHT")
	return *ret0, err
}

// MAXIMUMWEIGHT is a free data retrieval call binding the contract method 0x08f2ec06.
//
// Solidity: function MAXIMUM_WEIGHT() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) MAXIMUMWEIGHT() (*big.Int, error) {
	return _Relayerincentivize.Contract.MAXIMUMWEIGHT(&_Relayerincentivize.CallOpts)
}

// MAXIMUMWEIGHT is a free data retrieval call binding the contract method 0x08f2ec06.
//
// Solidity: function MAXIMUM_WEIGHT() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) MAXIMUMWEIGHT() (*big.Int, error) {
	return _Relayerincentivize.Contract.MAXIMUMWEIGHT(&_Relayerincentivize.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCaller) RELAYERHUBCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "RELAYERHUB_CONTRACT_ADDR")
	return *ret0, err
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.RELAYERHUBCONTRACTADDR(&_Relayerincentivize.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCallerSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.RELAYERHUBCONTRACTADDR(&_Relayerincentivize.CallOpts)
}

// RELAYERREWARD is a free data retrieval call binding the contract method 0x75405d0d.
//
// Solidity: function RELAYER_REWARD() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) RELAYERREWARD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "RELAYER_REWARD")
	return *ret0, err
}

// RELAYERREWARD is a free data retrieval call binding the contract method 0x75405d0d.
//
// Solidity: function RELAYER_REWARD() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) RELAYERREWARD() (*big.Int, error) {
	return _Relayerincentivize.Contract.RELAYERREWARD(&_Relayerincentivize.CallOpts)
}

// RELAYERREWARD is a free data retrieval call binding the contract method 0x75405d0d.
//
// Solidity: function RELAYER_REWARD() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) RELAYERREWARD() (*big.Int, error) {
	return _Relayerincentivize.Contract.RELAYERREWARD(&_Relayerincentivize.CallOpts)
}

// ROUNDSIZE is a free data retrieval call binding the contract method 0x54133307.
//
// Solidity: function ROUND_SIZE() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) ROUNDSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "ROUND_SIZE")
	return *ret0, err
}

// ROUNDSIZE is a free data retrieval call binding the contract method 0x54133307.
//
// Solidity: function ROUND_SIZE() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) ROUNDSIZE() (*big.Int, error) {
	return _Relayerincentivize.Contract.ROUNDSIZE(&_Relayerincentivize.CallOpts)
}

// ROUNDSIZE is a free data retrieval call binding the contract method 0x54133307.
//
// Solidity: function ROUND_SIZE() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) ROUNDSIZE() (*big.Int, error) {
	return _Relayerincentivize.Contract.ROUNDSIZE(&_Relayerincentivize.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCaller) SLASHCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "SLASH_CONTRACT_ADDR")
	return *ret0, err
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.SLASHCONTRACTADDR(&_Relayerincentivize.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCallerSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.SLASHCONTRACTADDR(&_Relayerincentivize.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCaller) SYSTEMREWARDADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "SYSTEM_REWARD_ADDR")
	return *ret0, err
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.SYSTEMREWARDADDR(&_Relayerincentivize.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCallerSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.SYSTEMREWARDADDR(&_Relayerincentivize.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCaller) TOKENHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "TOKEN_HUB_ADDR")
	return *ret0, err
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeSession) TOKENHUBADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.TOKENHUBADDR(&_Relayerincentivize.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCallerSession) TOKENHUBADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.TOKENHUBADDR(&_Relayerincentivize.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCaller) VALIDATORCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "VALIDATOR_CONTRACT_ADDR")
	return *ret0, err
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.VALIDATORCONTRACTADDR(&_Relayerincentivize.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCallerSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.VALIDATORCONTRACTADDR(&_Relayerincentivize.CallOpts)
}

// CollectedRewardForHeaderRelayer is a free data retrieval call binding the contract method 0x1f4c0dab.
//
// Solidity: function _collectedRewardForHeaderRelayer() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) CollectedRewardForHeaderRelayer(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "_collectedRewardForHeaderRelayer")
	return *ret0, err
}

// CollectedRewardForHeaderRelayer is a free data retrieval call binding the contract method 0x1f4c0dab.
//
// Solidity: function _collectedRewardForHeaderRelayer() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) CollectedRewardForHeaderRelayer() (*big.Int, error) {
	return _Relayerincentivize.Contract.CollectedRewardForHeaderRelayer(&_Relayerincentivize.CallOpts)
}

// CollectedRewardForHeaderRelayer is a free data retrieval call binding the contract method 0x1f4c0dab.
//
// Solidity: function _collectedRewardForHeaderRelayer() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) CollectedRewardForHeaderRelayer() (*big.Int, error) {
	return _Relayerincentivize.Contract.CollectedRewardForHeaderRelayer(&_Relayerincentivize.CallOpts)
}

// CollectedRewardForTransferRelayer is a free data retrieval call binding the contract method 0x7342a926.
//
// Solidity: function _collectedRewardForTransferRelayer() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) CollectedRewardForTransferRelayer(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "_collectedRewardForTransferRelayer")
	return *ret0, err
}

// CollectedRewardForTransferRelayer is a free data retrieval call binding the contract method 0x7342a926.
//
// Solidity: function _collectedRewardForTransferRelayer() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) CollectedRewardForTransferRelayer() (*big.Int, error) {
	return _Relayerincentivize.Contract.CollectedRewardForTransferRelayer(&_Relayerincentivize.CallOpts)
}

// CollectedRewardForTransferRelayer is a free data retrieval call binding the contract method 0x7342a926.
//
// Solidity: function _collectedRewardForTransferRelayer() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) CollectedRewardForTransferRelayer() (*big.Int, error) {
	return _Relayerincentivize.Contract.CollectedRewardForTransferRelayer(&_Relayerincentivize.CallOpts)
}

// CountInRound is a free data retrieval call binding the contract method 0xc780ba13.
//
// Solidity: function _countInRound() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) CountInRound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "_countInRound")
	return *ret0, err
}

// CountInRound is a free data retrieval call binding the contract method 0xc780ba13.
//
// Solidity: function _countInRound() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) CountInRound() (*big.Int, error) {
	return _Relayerincentivize.Contract.CountInRound(&_Relayerincentivize.CallOpts)
}

// CountInRound is a free data retrieval call binding the contract method 0xc780ba13.
//
// Solidity: function _countInRound() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) CountInRound() (*big.Int, error) {
	return _Relayerincentivize.Contract.CountInRound(&_Relayerincentivize.CallOpts)
}

// HeaderRelayerAddressRecord is a free data retrieval call binding the contract method 0x9b589faf.
//
// Solidity: function _headerRelayerAddressRecord(uint256 ) constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCaller) HeaderRelayerAddressRecord(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "_headerRelayerAddressRecord", arg0)
	return *ret0, err
}

// HeaderRelayerAddressRecord is a free data retrieval call binding the contract method 0x9b589faf.
//
// Solidity: function _headerRelayerAddressRecord(uint256 ) constant returns(address)
func (_Relayerincentivize *RelayerincentivizeSession) HeaderRelayerAddressRecord(arg0 *big.Int) (common.Address, error) {
	return _Relayerincentivize.Contract.HeaderRelayerAddressRecord(&_Relayerincentivize.CallOpts, arg0)
}

// HeaderRelayerAddressRecord is a free data retrieval call binding the contract method 0x9b589faf.
//
// Solidity: function _headerRelayerAddressRecord(uint256 ) constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCallerSession) HeaderRelayerAddressRecord(arg0 *big.Int) (common.Address, error) {
	return _Relayerincentivize.Contract.HeaderRelayerAddressRecord(&_Relayerincentivize.CallOpts, arg0)
}

// HeaderRelayersSubmitCount is a free data retrieval call binding the contract method 0x0dae2f09.
//
// Solidity: function _headerRelayersSubmitCount(address ) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) HeaderRelayersSubmitCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "_headerRelayersSubmitCount", arg0)
	return *ret0, err
}

// HeaderRelayersSubmitCount is a free data retrieval call binding the contract method 0x0dae2f09.
//
// Solidity: function _headerRelayersSubmitCount(address ) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) HeaderRelayersSubmitCount(arg0 common.Address) (*big.Int, error) {
	return _Relayerincentivize.Contract.HeaderRelayersSubmitCount(&_Relayerincentivize.CallOpts, arg0)
}

// HeaderRelayersSubmitCount is a free data retrieval call binding the contract method 0x0dae2f09.
//
// Solidity: function _headerRelayersSubmitCount(address ) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) HeaderRelayersSubmitCount(arg0 common.Address) (*big.Int, error) {
	return _Relayerincentivize.Contract.HeaderRelayersSubmitCount(&_Relayerincentivize.CallOpts, arg0)
}

// RoundSequence is a free data retrieval call binding the contract method 0x2d804f04.
//
// Solidity: function _roundSequence() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) RoundSequence(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "_roundSequence")
	return *ret0, err
}

// RoundSequence is a free data retrieval call binding the contract method 0x2d804f04.
//
// Solidity: function _roundSequence() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) RoundSequence() (*big.Int, error) {
	return _Relayerincentivize.Contract.RoundSequence(&_Relayerincentivize.CallOpts)
}

// RoundSequence is a free data retrieval call binding the contract method 0x2d804f04.
//
// Solidity: function _roundSequence() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) RoundSequence() (*big.Int, error) {
	return _Relayerincentivize.Contract.RoundSequence(&_Relayerincentivize.CallOpts)
}

// TransferRelayerAddressRecord is a free data retrieval call binding the contract method 0x89518970.
//
// Solidity: function _transferRelayerAddressRecord(uint256 ) constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCaller) TransferRelayerAddressRecord(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "_transferRelayerAddressRecord", arg0)
	return *ret0, err
}

// TransferRelayerAddressRecord is a free data retrieval call binding the contract method 0x89518970.
//
// Solidity: function _transferRelayerAddressRecord(uint256 ) constant returns(address)
func (_Relayerincentivize *RelayerincentivizeSession) TransferRelayerAddressRecord(arg0 *big.Int) (common.Address, error) {
	return _Relayerincentivize.Contract.TransferRelayerAddressRecord(&_Relayerincentivize.CallOpts, arg0)
}

// TransferRelayerAddressRecord is a free data retrieval call binding the contract method 0x89518970.
//
// Solidity: function _transferRelayerAddressRecord(uint256 ) constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCallerSession) TransferRelayerAddressRecord(arg0 *big.Int) (common.Address, error) {
	return _Relayerincentivize.Contract.TransferRelayerAddressRecord(&_Relayerincentivize.CallOpts, arg0)
}

// TransferRelayersSubmitCount is a free data retrieval call binding the contract method 0x6984bc2b.
//
// Solidity: function _transferRelayersSubmitCount(address ) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) TransferRelayersSubmitCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "_transferRelayersSubmitCount", arg0)
	return *ret0, err
}

// TransferRelayersSubmitCount is a free data retrieval call binding the contract method 0x6984bc2b.
//
// Solidity: function _transferRelayersSubmitCount(address ) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) TransferRelayersSubmitCount(arg0 common.Address) (*big.Int, error) {
	return _Relayerincentivize.Contract.TransferRelayersSubmitCount(&_Relayerincentivize.CallOpts, arg0)
}

// TransferRelayersSubmitCount is a free data retrieval call binding the contract method 0x6984bc2b.
//
// Solidity: function _transferRelayersSubmitCount(address ) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) TransferRelayersSubmitCount(arg0 common.Address) (*big.Int, error) {
	return _Relayerincentivize.Contract.TransferRelayersSubmitCount(&_Relayerincentivize.CallOpts, arg0)
}

// CalculateHeaderRelayerWeight is a free data retrieval call binding the contract method 0xbd4cc830.
//
// Solidity: function calculateHeaderRelayerWeight(uint256 count) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) CalculateHeaderRelayerWeight(opts *bind.CallOpts, count *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "calculateHeaderRelayerWeight", count)
	return *ret0, err
}

// CalculateHeaderRelayerWeight is a free data retrieval call binding the contract method 0xbd4cc830.
//
// Solidity: function calculateHeaderRelayerWeight(uint256 count) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) CalculateHeaderRelayerWeight(count *big.Int) (*big.Int, error) {
	return _Relayerincentivize.Contract.CalculateHeaderRelayerWeight(&_Relayerincentivize.CallOpts, count)
}

// CalculateHeaderRelayerWeight is a free data retrieval call binding the contract method 0xbd4cc830.
//
// Solidity: function calculateHeaderRelayerWeight(uint256 count) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) CalculateHeaderRelayerWeight(count *big.Int) (*big.Int, error) {
	return _Relayerincentivize.Contract.CalculateHeaderRelayerWeight(&_Relayerincentivize.CallOpts, count)
}

// CalculateTransferRelayerWeight is a free data retrieval call binding the contract method 0xaf400681.
//
// Solidity: function calculateTransferRelayerWeight(uint256 count) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) CalculateTransferRelayerWeight(opts *bind.CallOpts, count *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "calculateTransferRelayerWeight", count)
	return *ret0, err
}

// CalculateTransferRelayerWeight is a free data retrieval call binding the contract method 0xaf400681.
//
// Solidity: function calculateTransferRelayerWeight(uint256 count) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) CalculateTransferRelayerWeight(count *big.Int) (*big.Int, error) {
	return _Relayerincentivize.Contract.CalculateTransferRelayerWeight(&_Relayerincentivize.CallOpts, count)
}

// CalculateTransferRelayerWeight is a free data retrieval call binding the contract method 0xaf400681.
//
// Solidity: function calculateTransferRelayerWeight(uint256 count) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) CalculateTransferRelayerWeight(count *big.Int) (*big.Int, error) {
	return _Relayerincentivize.Contract.CalculateTransferRelayerWeight(&_Relayerincentivize.CallOpts, count)
}

// DenominaroeCallerCompensation is a free data retrieval call binding the contract method 0x3d8e1a41.
//
// Solidity: function denominaroeCallerCompensation() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) DenominaroeCallerCompensation(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "denominaroeCallerCompensation")
	return *ret0, err
}

// DenominaroeCallerCompensation is a free data retrieval call binding the contract method 0x3d8e1a41.
//
// Solidity: function denominaroeCallerCompensation() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) DenominaroeCallerCompensation() (*big.Int, error) {
	return _Relayerincentivize.Contract.DenominaroeCallerCompensation(&_Relayerincentivize.CallOpts)
}

// DenominaroeCallerCompensation is a free data retrieval call binding the contract method 0x3d8e1a41.
//
// Solidity: function denominaroeCallerCompensation() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) DenominaroeCallerCompensation() (*big.Int, error) {
	return _Relayerincentivize.Contract.DenominaroeCallerCompensation(&_Relayerincentivize.CallOpts)
}

// DenominaroeHeaderRelayer is a free data retrieval call binding the contract method 0xe12c146b.
//
// Solidity: function denominaroeHeaderRelayer() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) DenominaroeHeaderRelayer(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "denominaroeHeaderRelayer")
	return *ret0, err
}

// DenominaroeHeaderRelayer is a free data retrieval call binding the contract method 0xe12c146b.
//
// Solidity: function denominaroeHeaderRelayer() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) DenominaroeHeaderRelayer() (*big.Int, error) {
	return _Relayerincentivize.Contract.DenominaroeHeaderRelayer(&_Relayerincentivize.CallOpts)
}

// DenominaroeHeaderRelayer is a free data retrieval call binding the contract method 0xe12c146b.
//
// Solidity: function denominaroeHeaderRelayer() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) DenominaroeHeaderRelayer() (*big.Int, error) {
	return _Relayerincentivize.Contract.DenominaroeHeaderRelayer(&_Relayerincentivize.CallOpts)
}

// MoleculeCallerCompensation is a free data retrieval call binding the contract method 0x0c732068.
//
// Solidity: function moleculeCallerCompensation() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) MoleculeCallerCompensation(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "moleculeCallerCompensation")
	return *ret0, err
}

// MoleculeCallerCompensation is a free data retrieval call binding the contract method 0x0c732068.
//
// Solidity: function moleculeCallerCompensation() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) MoleculeCallerCompensation() (*big.Int, error) {
	return _Relayerincentivize.Contract.MoleculeCallerCompensation(&_Relayerincentivize.CallOpts)
}

// MoleculeCallerCompensation is a free data retrieval call binding the contract method 0x0c732068.
//
// Solidity: function moleculeCallerCompensation() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) MoleculeCallerCompensation() (*big.Int, error) {
	return _Relayerincentivize.Contract.MoleculeCallerCompensation(&_Relayerincentivize.CallOpts)
}

// MoleculeHeaderRelayer is a free data retrieval call binding the contract method 0x87c1830a.
//
// Solidity: function moleculeHeaderRelayer() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) MoleculeHeaderRelayer(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "moleculeHeaderRelayer")
	return *ret0, err
}

// MoleculeHeaderRelayer is a free data retrieval call binding the contract method 0x87c1830a.
//
// Solidity: function moleculeHeaderRelayer() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) MoleculeHeaderRelayer() (*big.Int, error) {
	return _Relayerincentivize.Contract.MoleculeHeaderRelayer(&_Relayerincentivize.CallOpts)
}

// MoleculeHeaderRelayer is a free data retrieval call binding the contract method 0x87c1830a.
//
// Solidity: function moleculeHeaderRelayer() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) MoleculeHeaderRelayer() (*big.Int, error) {
	return _Relayerincentivize.Contract.MoleculeHeaderRelayer(&_Relayerincentivize.CallOpts)
}

// AddReward is a paid mutator transaction binding the contract method 0x40b47e1a.
//
// Solidity: function addReward(address headerRelayerAddr, address caller) returns(bool)
func (_Relayerincentivize *RelayerincentivizeTransactor) AddReward(opts *bind.TransactOpts, headerRelayerAddr common.Address, caller common.Address) (*types.Transaction, error) {
	return _Relayerincentivize.contract.Transact(opts, "addReward", headerRelayerAddr, caller)
}

// AddReward is a paid mutator transaction binding the contract method 0x40b47e1a.
//
// Solidity: function addReward(address headerRelayerAddr, address caller) returns(bool)
func (_Relayerincentivize *RelayerincentivizeSession) AddReward(headerRelayerAddr common.Address, caller common.Address) (*types.Transaction, error) {
	return _Relayerincentivize.Contract.AddReward(&_Relayerincentivize.TransactOpts, headerRelayerAddr, caller)
}

// AddReward is a paid mutator transaction binding the contract method 0x40b47e1a.
//
// Solidity: function addReward(address headerRelayerAddr, address caller) returns(bool)
func (_Relayerincentivize *RelayerincentivizeTransactorSession) AddReward(headerRelayerAddr common.Address, caller common.Address) (*types.Transaction, error) {
	return _Relayerincentivize.Contract.AddReward(&_Relayerincentivize.TransactOpts, headerRelayerAddr, caller)
}

// RelayerincentivizeLogDistributeCollectedRewardIterator is returned from FilterLogDistributeCollectedReward and is used to iterate over the raw logs and unpacked data for LogDistributeCollectedReward events raised by the Relayerincentivize contract.
type RelayerincentivizeLogDistributeCollectedRewardIterator struct {
	Event *RelayerincentivizeLogDistributeCollectedReward // Event containing the contract specifics and raw log

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
func (it *RelayerincentivizeLogDistributeCollectedRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerincentivizeLogDistributeCollectedReward)
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
		it.Event = new(RelayerincentivizeLogDistributeCollectedReward)
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
func (it *RelayerincentivizeLogDistributeCollectedRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerincentivizeLogDistributeCollectedRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerincentivizeLogDistributeCollectedReward represents a LogDistributeCollectedReward event raised by the Relayerincentivize contract.
type RelayerincentivizeLogDistributeCollectedReward struct {
	Sequence                      *big.Int
	RoundRewardForHeaderRelayer   *big.Int
	RoundRewardForTransferRelayer *big.Int
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterLogDistributeCollectedReward is a free log retrieval operation binding the contract event 0x127fd0fd1fcf84c71c5c649625ef186be105a71ecc80c7cb3f96cd99ccae1e0f.
//
// Solidity: event LogDistributeCollectedReward(uint256 sequence, uint256 roundRewardForHeaderRelayer, uint256 roundRewardForTransferRelayer)
func (_Relayerincentivize *RelayerincentivizeFilterer) FilterLogDistributeCollectedReward(opts *bind.FilterOpts) (*RelayerincentivizeLogDistributeCollectedRewardIterator, error) {

	logs, sub, err := _Relayerincentivize.contract.FilterLogs(opts, "LogDistributeCollectedReward")
	if err != nil {
		return nil, err
	}
	return &RelayerincentivizeLogDistributeCollectedRewardIterator{contract: _Relayerincentivize.contract, event: "LogDistributeCollectedReward", logs: logs, sub: sub}, nil
}

// WatchLogDistributeCollectedReward is a free log subscription operation binding the contract event 0x127fd0fd1fcf84c71c5c649625ef186be105a71ecc80c7cb3f96cd99ccae1e0f.
//
// Solidity: event LogDistributeCollectedReward(uint256 sequence, uint256 roundRewardForHeaderRelayer, uint256 roundRewardForTransferRelayer)
func (_Relayerincentivize *RelayerincentivizeFilterer) WatchLogDistributeCollectedReward(opts *bind.WatchOpts, sink chan<- *RelayerincentivizeLogDistributeCollectedReward) (event.Subscription, error) {

	logs, sub, err := _Relayerincentivize.contract.WatchLogs(opts, "LogDistributeCollectedReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerincentivizeLogDistributeCollectedReward)
				if err := _Relayerincentivize.contract.UnpackLog(event, "LogDistributeCollectedReward", log); err != nil {
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

// ParseLogDistributeCollectedReward is a log parse operation binding the contract event 0x127fd0fd1fcf84c71c5c649625ef186be105a71ecc80c7cb3f96cd99ccae1e0f.
//
// Solidity: event LogDistributeCollectedReward(uint256 sequence, uint256 roundRewardForHeaderRelayer, uint256 roundRewardForTransferRelayer)
func (_Relayerincentivize *RelayerincentivizeFilterer) ParseLogDistributeCollectedReward(log types.Log) (*RelayerincentivizeLogDistributeCollectedReward, error) {
	event := new(RelayerincentivizeLogDistributeCollectedReward)
	if err := _Relayerincentivize.contract.UnpackLog(event, "LogDistributeCollectedReward", log); err != nil {
		return nil, err
	}
	return event, nil
}
