// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bscvalidatorset

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

// BscvalidatorsetABI is the input ABI used to generate the binding from.
const BscvalidatorsetABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"batchTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"batchTransferFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"batchTransferLowerFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deprecatedDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"addresspayable\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"directTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"addresspayable\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"directTransferFail\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"failReasonWithStr\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"paramChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"systemTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"unexpectedPackage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"validatorDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"validatorEmptyJailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"validatorFelony\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"validatorJailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"validatorMisdemeanor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"validatorSetUpdated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"BIND_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CODE_OK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DUSTY_INCOMING\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERROR_FAIL_CHECK_VALIDATORS\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERROR_FAIL_DECODE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERROR_LEN_OF_VAL_MISMATCH\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERROR_RELAYFEE_TOO_LARGE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERROR_UNKNOWN_PACKAGE_TYPE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EXPIRE_TIME_SECOND_GAP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GOV_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GOV_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INCENTIVIZE_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INIT_VALIDATORSET_BYTES\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"JAIL_MESSAGE_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"LIGHT_CLIENT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RELAYERHUB_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SLASH_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SLASH_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STAKING_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SYSTEM_REWARD_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKEN_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKEN_MANAGER_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VALIDATORS_UPDATE_MESSAGE_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bscChainID\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentValidatorSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"feeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BBCFeeAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"votingPower\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"incoming\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"currentValidatorSetMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"expireTimeSecondGap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"felony\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getIncoming\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleAckPackage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleFailAckPackage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleSynPackage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"responsePayload\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"misdemeanor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalInComing\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"updateParam\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Bscvalidatorset is an auto generated Go binding around an Ethereum contract.
type Bscvalidatorset struct {
	BscvalidatorsetCaller     // Read-only binding to the contract
	BscvalidatorsetTransactor // Write-only binding to the contract
	BscvalidatorsetFilterer   // Log filterer for contract events
}

// BscvalidatorsetCaller is an auto generated read-only Go binding around an Ethereum contract.
type BscvalidatorsetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscvalidatorsetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BscvalidatorsetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscvalidatorsetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BscvalidatorsetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscvalidatorsetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BscvalidatorsetSession struct {
	Contract     *Bscvalidatorset  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BscvalidatorsetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BscvalidatorsetCallerSession struct {
	Contract *BscvalidatorsetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BscvalidatorsetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BscvalidatorsetTransactorSession struct {
	Contract     *BscvalidatorsetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BscvalidatorsetRaw is an auto generated low-level Go binding around an Ethereum contract.
type BscvalidatorsetRaw struct {
	Contract *Bscvalidatorset // Generic contract binding to access the raw methods on
}

// BscvalidatorsetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BscvalidatorsetCallerRaw struct {
	Contract *BscvalidatorsetCaller // Generic read-only contract binding to access the raw methods on
}

// BscvalidatorsetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BscvalidatorsetTransactorRaw struct {
	Contract *BscvalidatorsetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBscvalidatorset creates a new instance of Bscvalidatorset, bound to a specific deployed contract.
func NewBscvalidatorset(address common.Address, backend bind.ContractBackend) (*Bscvalidatorset, error) {
	contract, err := bindBscvalidatorset(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bscvalidatorset{BscvalidatorsetCaller: BscvalidatorsetCaller{contract: contract}, BscvalidatorsetTransactor: BscvalidatorsetTransactor{contract: contract}, BscvalidatorsetFilterer: BscvalidatorsetFilterer{contract: contract}}, nil
}

// NewBscvalidatorsetCaller creates a new read-only instance of Bscvalidatorset, bound to a specific deployed contract.
func NewBscvalidatorsetCaller(address common.Address, caller bind.ContractCaller) (*BscvalidatorsetCaller, error) {
	contract, err := bindBscvalidatorset(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BscvalidatorsetCaller{contract: contract}, nil
}

// NewBscvalidatorsetTransactor creates a new write-only instance of Bscvalidatorset, bound to a specific deployed contract.
func NewBscvalidatorsetTransactor(address common.Address, transactor bind.ContractTransactor) (*BscvalidatorsetTransactor, error) {
	contract, err := bindBscvalidatorset(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BscvalidatorsetTransactor{contract: contract}, nil
}

// NewBscvalidatorsetFilterer creates a new log filterer instance of Bscvalidatorset, bound to a specific deployed contract.
func NewBscvalidatorsetFilterer(address common.Address, filterer bind.ContractFilterer) (*BscvalidatorsetFilterer, error) {
	contract, err := bindBscvalidatorset(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BscvalidatorsetFilterer{contract: contract}, nil
}

// bindBscvalidatorset binds a generic wrapper to an already deployed contract.
func bindBscvalidatorset(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BscvalidatorsetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bscvalidatorset *BscvalidatorsetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bscvalidatorset.Contract.BscvalidatorsetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bscvalidatorset *BscvalidatorsetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bscvalidatorset.Contract.BscvalidatorsetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bscvalidatorset *BscvalidatorsetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bscvalidatorset.Contract.BscvalidatorsetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bscvalidatorset *BscvalidatorsetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bscvalidatorset.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bscvalidatorset *BscvalidatorsetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bscvalidatorset.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bscvalidatorset *BscvalidatorsetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bscvalidatorset.Contract.contract.Transact(opts, method, params...)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetCaller) BINDCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "BIND_CHANNELID")
	return *ret0, err
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetSession) BINDCHANNELID() (uint8, error) {
	return _Bscvalidatorset.Contract.BINDCHANNELID(&_Bscvalidatorset.CallOpts)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) BINDCHANNELID() (uint8, error) {
	return _Bscvalidatorset.Contract.BINDCHANNELID(&_Bscvalidatorset.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() constant returns(uint32)
func (_Bscvalidatorset *BscvalidatorsetCaller) CODEOK(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "CODE_OK")
	return *ret0, err
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() constant returns(uint32)
func (_Bscvalidatorset *BscvalidatorsetSession) CODEOK() (uint32, error) {
	return _Bscvalidatorset.Contract.CODEOK(&_Bscvalidatorset.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() constant returns(uint32)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) CODEOK() (uint32, error) {
	return _Bscvalidatorset.Contract.CODEOK(&_Bscvalidatorset.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetCaller) CROSSCHAINCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "CROSS_CHAIN_CONTRACT_ADDR")
	return *ret0, err
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Bscvalidatorset.Contract.CROSSCHAINCONTRACTADDR(&_Bscvalidatorset.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Bscvalidatorset.Contract.CROSSCHAINCONTRACTADDR(&_Bscvalidatorset.CallOpts)
}

// DUSTYINCOMING is a free data retrieval call binding the contract method 0xd86222d5.
//
// Solidity: function DUSTY_INCOMING() constant returns(uint256)
func (_Bscvalidatorset *BscvalidatorsetCaller) DUSTYINCOMING(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "DUSTY_INCOMING")
	return *ret0, err
}

// DUSTYINCOMING is a free data retrieval call binding the contract method 0xd86222d5.
//
// Solidity: function DUSTY_INCOMING() constant returns(uint256)
func (_Bscvalidatorset *BscvalidatorsetSession) DUSTYINCOMING() (*big.Int, error) {
	return _Bscvalidatorset.Contract.DUSTYINCOMING(&_Bscvalidatorset.CallOpts)
}

// DUSTYINCOMING is a free data retrieval call binding the contract method 0xd86222d5.
//
// Solidity: function DUSTY_INCOMING() constant returns(uint256)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) DUSTYINCOMING() (*big.Int, error) {
	return _Bscvalidatorset.Contract.DUSTYINCOMING(&_Bscvalidatorset.CallOpts)
}

// ERRORFAILCHECKVALIDATORS is a free data retrieval call binding the contract method 0x81650b62.
//
// Solidity: function ERROR_FAIL_CHECK_VALIDATORS() constant returns(uint32)
func (_Bscvalidatorset *BscvalidatorsetCaller) ERRORFAILCHECKVALIDATORS(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "ERROR_FAIL_CHECK_VALIDATORS")
	return *ret0, err
}

// ERRORFAILCHECKVALIDATORS is a free data retrieval call binding the contract method 0x81650b62.
//
// Solidity: function ERROR_FAIL_CHECK_VALIDATORS() constant returns(uint32)
func (_Bscvalidatorset *BscvalidatorsetSession) ERRORFAILCHECKVALIDATORS() (uint32, error) {
	return _Bscvalidatorset.Contract.ERRORFAILCHECKVALIDATORS(&_Bscvalidatorset.CallOpts)
}

// ERRORFAILCHECKVALIDATORS is a free data retrieval call binding the contract method 0x81650b62.
//
// Solidity: function ERROR_FAIL_CHECK_VALIDATORS() constant returns(uint32)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) ERRORFAILCHECKVALIDATORS() (uint32, error) {
	return _Bscvalidatorset.Contract.ERRORFAILCHECKVALIDATORS(&_Bscvalidatorset.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() constant returns(uint32)
func (_Bscvalidatorset *BscvalidatorsetCaller) ERRORFAILDECODE(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "ERROR_FAIL_DECODE")
	return *ret0, err
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() constant returns(uint32)
func (_Bscvalidatorset *BscvalidatorsetSession) ERRORFAILDECODE() (uint32, error) {
	return _Bscvalidatorset.Contract.ERRORFAILDECODE(&_Bscvalidatorset.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() constant returns(uint32)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) ERRORFAILDECODE() (uint32, error) {
	return _Bscvalidatorset.Contract.ERRORFAILDECODE(&_Bscvalidatorset.CallOpts)
}

// ERRORLENOFVALMISMATCH is a free data retrieval call binding the contract method 0x5d77156c.
//
// Solidity: function ERROR_LEN_OF_VAL_MISMATCH() constant returns(uint32)
func (_Bscvalidatorset *BscvalidatorsetCaller) ERRORLENOFVALMISMATCH(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "ERROR_LEN_OF_VAL_MISMATCH")
	return *ret0, err
}

// ERRORLENOFVALMISMATCH is a free data retrieval call binding the contract method 0x5d77156c.
//
// Solidity: function ERROR_LEN_OF_VAL_MISMATCH() constant returns(uint32)
func (_Bscvalidatorset *BscvalidatorsetSession) ERRORLENOFVALMISMATCH() (uint32, error) {
	return _Bscvalidatorset.Contract.ERRORLENOFVALMISMATCH(&_Bscvalidatorset.CallOpts)
}

// ERRORLENOFVALMISMATCH is a free data retrieval call binding the contract method 0x5d77156c.
//
// Solidity: function ERROR_LEN_OF_VAL_MISMATCH() constant returns(uint32)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) ERRORLENOFVALMISMATCH() (uint32, error) {
	return _Bscvalidatorset.Contract.ERRORLENOFVALMISMATCH(&_Bscvalidatorset.CallOpts)
}

// ERRORRELAYFEETOOLARGE is a free data retrieval call binding the contract method 0x219f22d5.
//
// Solidity: function ERROR_RELAYFEE_TOO_LARGE() constant returns(uint32)
func (_Bscvalidatorset *BscvalidatorsetCaller) ERRORRELAYFEETOOLARGE(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "ERROR_RELAYFEE_TOO_LARGE")
	return *ret0, err
}

// ERRORRELAYFEETOOLARGE is a free data retrieval call binding the contract method 0x219f22d5.
//
// Solidity: function ERROR_RELAYFEE_TOO_LARGE() constant returns(uint32)
func (_Bscvalidatorset *BscvalidatorsetSession) ERRORRELAYFEETOOLARGE() (uint32, error) {
	return _Bscvalidatorset.Contract.ERRORRELAYFEETOOLARGE(&_Bscvalidatorset.CallOpts)
}

// ERRORRELAYFEETOOLARGE is a free data retrieval call binding the contract method 0x219f22d5.
//
// Solidity: function ERROR_RELAYFEE_TOO_LARGE() constant returns(uint32)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) ERRORRELAYFEETOOLARGE() (uint32, error) {
	return _Bscvalidatorset.Contract.ERRORRELAYFEETOOLARGE(&_Bscvalidatorset.CallOpts)
}

// ERRORUNKNOWNPACKAGETYPE is a free data retrieval call binding the contract method 0xeda5868c.
//
// Solidity: function ERROR_UNKNOWN_PACKAGE_TYPE() constant returns(uint32)
func (_Bscvalidatorset *BscvalidatorsetCaller) ERRORUNKNOWNPACKAGETYPE(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "ERROR_UNKNOWN_PACKAGE_TYPE")
	return *ret0, err
}

// ERRORUNKNOWNPACKAGETYPE is a free data retrieval call binding the contract method 0xeda5868c.
//
// Solidity: function ERROR_UNKNOWN_PACKAGE_TYPE() constant returns(uint32)
func (_Bscvalidatorset *BscvalidatorsetSession) ERRORUNKNOWNPACKAGETYPE() (uint32, error) {
	return _Bscvalidatorset.Contract.ERRORUNKNOWNPACKAGETYPE(&_Bscvalidatorset.CallOpts)
}

// ERRORUNKNOWNPACKAGETYPE is a free data retrieval call binding the contract method 0xeda5868c.
//
// Solidity: function ERROR_UNKNOWN_PACKAGE_TYPE() constant returns(uint32)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) ERRORUNKNOWNPACKAGETYPE() (uint32, error) {
	return _Bscvalidatorset.Contract.ERRORUNKNOWNPACKAGETYPE(&_Bscvalidatorset.CallOpts)
}

// EXPIRETIMESECONDGAP is a free data retrieval call binding the contract method 0x853230aa.
//
// Solidity: function EXPIRE_TIME_SECOND_GAP() constant returns(uint256)
func (_Bscvalidatorset *BscvalidatorsetCaller) EXPIRETIMESECONDGAP(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "EXPIRE_TIME_SECOND_GAP")
	return *ret0, err
}

// EXPIRETIMESECONDGAP is a free data retrieval call binding the contract method 0x853230aa.
//
// Solidity: function EXPIRE_TIME_SECOND_GAP() constant returns(uint256)
func (_Bscvalidatorset *BscvalidatorsetSession) EXPIRETIMESECONDGAP() (*big.Int, error) {
	return _Bscvalidatorset.Contract.EXPIRETIMESECONDGAP(&_Bscvalidatorset.CallOpts)
}

// EXPIRETIMESECONDGAP is a free data retrieval call binding the contract method 0x853230aa.
//
// Solidity: function EXPIRE_TIME_SECOND_GAP() constant returns(uint256)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) EXPIRETIMESECONDGAP() (*big.Int, error) {
	return _Bscvalidatorset.Contract.EXPIRETIMESECONDGAP(&_Bscvalidatorset.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "GOV_CHANNELID")
	return *ret0, err
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetSession) GOVCHANNELID() (uint8, error) {
	return _Bscvalidatorset.Contract.GOVCHANNELID(&_Bscvalidatorset.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) GOVCHANNELID() (uint8, error) {
	return _Bscvalidatorset.Contract.GOVCHANNELID(&_Bscvalidatorset.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetCaller) GOVHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "GOV_HUB_ADDR")
	return *ret0, err
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetSession) GOVHUBADDR() (common.Address, error) {
	return _Bscvalidatorset.Contract.GOVHUBADDR(&_Bscvalidatorset.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) GOVHUBADDR() (common.Address, error) {
	return _Bscvalidatorset.Contract.GOVHUBADDR(&_Bscvalidatorset.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetCaller) INCENTIVIZEADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "INCENTIVIZE_ADDR")
	return *ret0, err
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Bscvalidatorset.Contract.INCENTIVIZEADDR(&_Bscvalidatorset.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Bscvalidatorset.Contract.INCENTIVIZEADDR(&_Bscvalidatorset.CallOpts)
}

// INITVALIDATORSETBYTES is a free data retrieval call binding the contract method 0xa5422d5c.
//
// Solidity: function INIT_VALIDATORSET_BYTES() constant returns(bytes)
func (_Bscvalidatorset *BscvalidatorsetCaller) INITVALIDATORSETBYTES(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "INIT_VALIDATORSET_BYTES")
	return *ret0, err
}

// INITVALIDATORSETBYTES is a free data retrieval call binding the contract method 0xa5422d5c.
//
// Solidity: function INIT_VALIDATORSET_BYTES() constant returns(bytes)
func (_Bscvalidatorset *BscvalidatorsetSession) INITVALIDATORSETBYTES() ([]byte, error) {
	return _Bscvalidatorset.Contract.INITVALIDATORSETBYTES(&_Bscvalidatorset.CallOpts)
}

// INITVALIDATORSETBYTES is a free data retrieval call binding the contract method 0xa5422d5c.
//
// Solidity: function INIT_VALIDATORSET_BYTES() constant returns(bytes)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) INITVALIDATORSETBYTES() ([]byte, error) {
	return _Bscvalidatorset.Contract.INITVALIDATORSETBYTES(&_Bscvalidatorset.CallOpts)
}

// JAILMESSAGETYPE is a free data retrieval call binding the contract method 0xbf9f4995.
//
// Solidity: function JAIL_MESSAGE_TYPE() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetCaller) JAILMESSAGETYPE(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "JAIL_MESSAGE_TYPE")
	return *ret0, err
}

// JAILMESSAGETYPE is a free data retrieval call binding the contract method 0xbf9f4995.
//
// Solidity: function JAIL_MESSAGE_TYPE() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetSession) JAILMESSAGETYPE() (uint8, error) {
	return _Bscvalidatorset.Contract.JAILMESSAGETYPE(&_Bscvalidatorset.CallOpts)
}

// JAILMESSAGETYPE is a free data retrieval call binding the contract method 0xbf9f4995.
//
// Solidity: function JAIL_MESSAGE_TYPE() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) JAILMESSAGETYPE() (uint8, error) {
	return _Bscvalidatorset.Contract.JAILMESSAGETYPE(&_Bscvalidatorset.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetCaller) LIGHTCLIENTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "LIGHT_CLIENT_ADDR")
	return *ret0, err
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Bscvalidatorset.Contract.LIGHTCLIENTADDR(&_Bscvalidatorset.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Bscvalidatorset.Contract.LIGHTCLIENTADDR(&_Bscvalidatorset.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() constant returns(uint256)
func (_Bscvalidatorset *BscvalidatorsetCaller) PRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "PRECISION")
	return *ret0, err
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() constant returns(uint256)
func (_Bscvalidatorset *BscvalidatorsetSession) PRECISION() (*big.Int, error) {
	return _Bscvalidatorset.Contract.PRECISION(&_Bscvalidatorset.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() constant returns(uint256)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) PRECISION() (*big.Int, error) {
	return _Bscvalidatorset.Contract.PRECISION(&_Bscvalidatorset.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetCaller) RELAYERHUBCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "RELAYERHUB_CONTRACT_ADDR")
	return *ret0, err
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Bscvalidatorset.Contract.RELAYERHUBCONTRACTADDR(&_Bscvalidatorset.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Bscvalidatorset.Contract.RELAYERHUBCONTRACTADDR(&_Bscvalidatorset.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetCaller) SLASHCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "SLASH_CHANNELID")
	return *ret0, err
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetSession) SLASHCHANNELID() (uint8, error) {
	return _Bscvalidatorset.Contract.SLASHCHANNELID(&_Bscvalidatorset.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) SLASHCHANNELID() (uint8, error) {
	return _Bscvalidatorset.Contract.SLASHCHANNELID(&_Bscvalidatorset.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetCaller) SLASHCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "SLASH_CONTRACT_ADDR")
	return *ret0, err
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Bscvalidatorset.Contract.SLASHCONTRACTADDR(&_Bscvalidatorset.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Bscvalidatorset.Contract.SLASHCONTRACTADDR(&_Bscvalidatorset.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetCaller) STAKINGCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "STAKING_CHANNELID")
	return *ret0, err
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetSession) STAKINGCHANNELID() (uint8, error) {
	return _Bscvalidatorset.Contract.STAKINGCHANNELID(&_Bscvalidatorset.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) STAKINGCHANNELID() (uint8, error) {
	return _Bscvalidatorset.Contract.STAKINGCHANNELID(&_Bscvalidatorset.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetCaller) SYSTEMREWARDADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "SYSTEM_REWARD_ADDR")
	return *ret0, err
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Bscvalidatorset.Contract.SYSTEMREWARDADDR(&_Bscvalidatorset.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Bscvalidatorset.Contract.SYSTEMREWARDADDR(&_Bscvalidatorset.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetCaller) TOKENHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "TOKEN_HUB_ADDR")
	return *ret0, err
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetSession) TOKENHUBADDR() (common.Address, error) {
	return _Bscvalidatorset.Contract.TOKENHUBADDR(&_Bscvalidatorset.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) TOKENHUBADDR() (common.Address, error) {
	return _Bscvalidatorset.Contract.TOKENHUBADDR(&_Bscvalidatorset.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetCaller) TOKENMANAGERADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "TOKEN_MANAGER_ADDR")
	return *ret0, err
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetSession) TOKENMANAGERADDR() (common.Address, error) {
	return _Bscvalidatorset.Contract.TOKENMANAGERADDR(&_Bscvalidatorset.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) TOKENMANAGERADDR() (common.Address, error) {
	return _Bscvalidatorset.Contract.TOKENMANAGERADDR(&_Bscvalidatorset.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "TRANSFER_IN_CHANNELID")
	return *ret0, err
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Bscvalidatorset.Contract.TRANSFERINCHANNELID(&_Bscvalidatorset.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Bscvalidatorset.Contract.TRANSFERINCHANNELID(&_Bscvalidatorset.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "TRANSFER_OUT_CHANNELID")
	return *ret0, err
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Bscvalidatorset.Contract.TRANSFEROUTCHANNELID(&_Bscvalidatorset.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Bscvalidatorset.Contract.TRANSFEROUTCHANNELID(&_Bscvalidatorset.CallOpts)
}

// VALIDATORSUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x5667515a.
//
// Solidity: function VALIDATORS_UPDATE_MESSAGE_TYPE() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetCaller) VALIDATORSUPDATEMESSAGETYPE(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "VALIDATORS_UPDATE_MESSAGE_TYPE")
	return *ret0, err
}

// VALIDATORSUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x5667515a.
//
// Solidity: function VALIDATORS_UPDATE_MESSAGE_TYPE() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetSession) VALIDATORSUPDATEMESSAGETYPE() (uint8, error) {
	return _Bscvalidatorset.Contract.VALIDATORSUPDATEMESSAGETYPE(&_Bscvalidatorset.CallOpts)
}

// VALIDATORSUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x5667515a.
//
// Solidity: function VALIDATORS_UPDATE_MESSAGE_TYPE() constant returns(uint8)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) VALIDATORSUPDATEMESSAGETYPE() (uint8, error) {
	return _Bscvalidatorset.Contract.VALIDATORSUPDATEMESSAGETYPE(&_Bscvalidatorset.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetCaller) VALIDATORCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "VALIDATOR_CONTRACT_ADDR")
	return *ret0, err
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Bscvalidatorset.Contract.VALIDATORCONTRACTADDR(&_Bscvalidatorset.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Bscvalidatorset.Contract.VALIDATORCONTRACTADDR(&_Bscvalidatorset.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Bscvalidatorset *BscvalidatorsetCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "alreadyInit")
	return *ret0, err
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Bscvalidatorset *BscvalidatorsetSession) AlreadyInit() (bool, error) {
	return _Bscvalidatorset.Contract.AlreadyInit(&_Bscvalidatorset.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) AlreadyInit() (bool, error) {
	return _Bscvalidatorset.Contract.AlreadyInit(&_Bscvalidatorset.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() constant returns(uint16)
func (_Bscvalidatorset *BscvalidatorsetCaller) BscChainID(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "bscChainID")
	return *ret0, err
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() constant returns(uint16)
func (_Bscvalidatorset *BscvalidatorsetSession) BscChainID() (uint16, error) {
	return _Bscvalidatorset.Contract.BscChainID(&_Bscvalidatorset.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() constant returns(uint16)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) BscChainID() (uint16, error) {
	return _Bscvalidatorset.Contract.BscChainID(&_Bscvalidatorset.CallOpts)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) constant returns(address consensusAddress, address feeAddress, address BBCFeeAddress, uint64 votingPower, bool jailed, uint256 incoming)
func (_Bscvalidatorset *BscvalidatorsetCaller) CurrentValidatorSet(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	FeeAddress       common.Address
	BBCFeeAddress    common.Address
	VotingPower      uint64
	Jailed           bool
	Incoming         *big.Int
}, error) {
	ret := new(struct {
		ConsensusAddress common.Address
		FeeAddress       common.Address
		BBCFeeAddress    common.Address
		VotingPower      uint64
		Jailed           bool
		Incoming         *big.Int
	})
	out := ret
	err := _Bscvalidatorset.contract.Call(opts, out, "currentValidatorSet", arg0)
	return *ret, err
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) constant returns(address consensusAddress, address feeAddress, address BBCFeeAddress, uint64 votingPower, bool jailed, uint256 incoming)
func (_Bscvalidatorset *BscvalidatorsetSession) CurrentValidatorSet(arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	FeeAddress       common.Address
	BBCFeeAddress    common.Address
	VotingPower      uint64
	Jailed           bool
	Incoming         *big.Int
}, error) {
	return _Bscvalidatorset.Contract.CurrentValidatorSet(&_Bscvalidatorset.CallOpts, arg0)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) constant returns(address consensusAddress, address feeAddress, address BBCFeeAddress, uint64 votingPower, bool jailed, uint256 incoming)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) CurrentValidatorSet(arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	FeeAddress       common.Address
	BBCFeeAddress    common.Address
	VotingPower      uint64
	Jailed           bool
	Incoming         *big.Int
}, error) {
	return _Bscvalidatorset.Contract.CurrentValidatorSet(&_Bscvalidatorset.CallOpts, arg0)
}

// CurrentValidatorSetMap is a free data retrieval call binding the contract method 0xad3c9da6.
//
// Solidity: function currentValidatorSetMap(address ) constant returns(uint256)
func (_Bscvalidatorset *BscvalidatorsetCaller) CurrentValidatorSetMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "currentValidatorSetMap", arg0)
	return *ret0, err
}

// CurrentValidatorSetMap is a free data retrieval call binding the contract method 0xad3c9da6.
//
// Solidity: function currentValidatorSetMap(address ) constant returns(uint256)
func (_Bscvalidatorset *BscvalidatorsetSession) CurrentValidatorSetMap(arg0 common.Address) (*big.Int, error) {
	return _Bscvalidatorset.Contract.CurrentValidatorSetMap(&_Bscvalidatorset.CallOpts, arg0)
}

// CurrentValidatorSetMap is a free data retrieval call binding the contract method 0xad3c9da6.
//
// Solidity: function currentValidatorSetMap(address ) constant returns(uint256)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) CurrentValidatorSetMap(arg0 common.Address) (*big.Int, error) {
	return _Bscvalidatorset.Contract.CurrentValidatorSetMap(&_Bscvalidatorset.CallOpts, arg0)
}

// ExpireTimeSecondGap is a free data retrieval call binding the contract method 0x86249882.
//
// Solidity: function expireTimeSecondGap() constant returns(uint256)
func (_Bscvalidatorset *BscvalidatorsetCaller) ExpireTimeSecondGap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "expireTimeSecondGap")
	return *ret0, err
}

// ExpireTimeSecondGap is a free data retrieval call binding the contract method 0x86249882.
//
// Solidity: function expireTimeSecondGap() constant returns(uint256)
func (_Bscvalidatorset *BscvalidatorsetSession) ExpireTimeSecondGap() (*big.Int, error) {
	return _Bscvalidatorset.Contract.ExpireTimeSecondGap(&_Bscvalidatorset.CallOpts)
}

// ExpireTimeSecondGap is a free data retrieval call binding the contract method 0x86249882.
//
// Solidity: function expireTimeSecondGap() constant returns(uint256)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) ExpireTimeSecondGap() (*big.Int, error) {
	return _Bscvalidatorset.Contract.ExpireTimeSecondGap(&_Bscvalidatorset.CallOpts)
}

// GetIncoming is a free data retrieval call binding the contract method 0x565c56b3.
//
// Solidity: function getIncoming(address validator) constant returns(uint256)
func (_Bscvalidatorset *BscvalidatorsetCaller) GetIncoming(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "getIncoming", validator)
	return *ret0, err
}

// GetIncoming is a free data retrieval call binding the contract method 0x565c56b3.
//
// Solidity: function getIncoming(address validator) constant returns(uint256)
func (_Bscvalidatorset *BscvalidatorsetSession) GetIncoming(validator common.Address) (*big.Int, error) {
	return _Bscvalidatorset.Contract.GetIncoming(&_Bscvalidatorset.CallOpts, validator)
}

// GetIncoming is a free data retrieval call binding the contract method 0x565c56b3.
//
// Solidity: function getIncoming(address validator) constant returns(uint256)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) GetIncoming(validator common.Address) (*big.Int, error) {
	return _Bscvalidatorset.Contract.GetIncoming(&_Bscvalidatorset.CallOpts, validator)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() constant returns(address[])
func (_Bscvalidatorset *BscvalidatorsetCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "getValidators")
	return *ret0, err
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() constant returns(address[])
func (_Bscvalidatorset *BscvalidatorsetSession) GetValidators() ([]common.Address, error) {
	return _Bscvalidatorset.Contract.GetValidators(&_Bscvalidatorset.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() constant returns(address[])
func (_Bscvalidatorset *BscvalidatorsetCallerSession) GetValidators() ([]common.Address, error) {
	return _Bscvalidatorset.Contract.GetValidators(&_Bscvalidatorset.CallOpts)
}

// TotalInComing is a free data retrieval call binding the contract method 0x1ff18069.
//
// Solidity: function totalInComing() constant returns(uint256)
func (_Bscvalidatorset *BscvalidatorsetCaller) TotalInComing(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bscvalidatorset.contract.Call(opts, out, "totalInComing")
	return *ret0, err
}

// TotalInComing is a free data retrieval call binding the contract method 0x1ff18069.
//
// Solidity: function totalInComing() constant returns(uint256)
func (_Bscvalidatorset *BscvalidatorsetSession) TotalInComing() (*big.Int, error) {
	return _Bscvalidatorset.Contract.TotalInComing(&_Bscvalidatorset.CallOpts)
}

// TotalInComing is a free data retrieval call binding the contract method 0x1ff18069.
//
// Solidity: function totalInComing() constant returns(uint256)
func (_Bscvalidatorset *BscvalidatorsetCallerSession) TotalInComing() (*big.Int, error) {
	return _Bscvalidatorset.Contract.TotalInComing(&_Bscvalidatorset.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address valAddr) returns()
func (_Bscvalidatorset *BscvalidatorsetTransactor) Deposit(opts *bind.TransactOpts, valAddr common.Address) (*types.Transaction, error) {
	return _Bscvalidatorset.contract.Transact(opts, "deposit", valAddr)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address valAddr) returns()
func (_Bscvalidatorset *BscvalidatorsetSession) Deposit(valAddr common.Address) (*types.Transaction, error) {
	return _Bscvalidatorset.Contract.Deposit(&_Bscvalidatorset.TransactOpts, valAddr)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address valAddr) returns()
func (_Bscvalidatorset *BscvalidatorsetTransactorSession) Deposit(valAddr common.Address) (*types.Transaction, error) {
	return _Bscvalidatorset.Contract.Deposit(&_Bscvalidatorset.TransactOpts, valAddr)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address validator) returns()
func (_Bscvalidatorset *BscvalidatorsetTransactor) Felony(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Bscvalidatorset.contract.Transact(opts, "felony", validator)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address validator) returns()
func (_Bscvalidatorset *BscvalidatorsetSession) Felony(validator common.Address) (*types.Transaction, error) {
	return _Bscvalidatorset.Contract.Felony(&_Bscvalidatorset.TransactOpts, validator)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address validator) returns()
func (_Bscvalidatorset *BscvalidatorsetTransactorSession) Felony(validator common.Address) (*types.Transaction, error) {
	return _Bscvalidatorset.Contract.Felony(&_Bscvalidatorset.TransactOpts, validator)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Bscvalidatorset *BscvalidatorsetTransactor) HandleAckPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Bscvalidatorset.contract.Transact(opts, "handleAckPackage", channelId, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Bscvalidatorset *BscvalidatorsetSession) HandleAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Bscvalidatorset.Contract.HandleAckPackage(&_Bscvalidatorset.TransactOpts, channelId, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Bscvalidatorset *BscvalidatorsetTransactorSession) HandleAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Bscvalidatorset.Contract.HandleAckPackage(&_Bscvalidatorset.TransactOpts, channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Bscvalidatorset *BscvalidatorsetTransactor) HandleFailAckPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Bscvalidatorset.contract.Transact(opts, "handleFailAckPackage", channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Bscvalidatorset *BscvalidatorsetSession) HandleFailAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Bscvalidatorset.Contract.HandleFailAckPackage(&_Bscvalidatorset.TransactOpts, channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Bscvalidatorset *BscvalidatorsetTransactorSession) HandleFailAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Bscvalidatorset.Contract.HandleFailAckPackage(&_Bscvalidatorset.TransactOpts, channelId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes responsePayload)
func (_Bscvalidatorset *BscvalidatorsetTransactor) HandleSynPackage(opts *bind.TransactOpts, arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Bscvalidatorset.contract.Transact(opts, "handleSynPackage", arg0, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes responsePayload)
func (_Bscvalidatorset *BscvalidatorsetSession) HandleSynPackage(arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Bscvalidatorset.Contract.HandleSynPackage(&_Bscvalidatorset.TransactOpts, arg0, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes responsePayload)
func (_Bscvalidatorset *BscvalidatorsetTransactorSession) HandleSynPackage(arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Bscvalidatorset.Contract.HandleSynPackage(&_Bscvalidatorset.TransactOpts, arg0, msgBytes)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Bscvalidatorset *BscvalidatorsetTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bscvalidatorset.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Bscvalidatorset *BscvalidatorsetSession) Init() (*types.Transaction, error) {
	return _Bscvalidatorset.Contract.Init(&_Bscvalidatorset.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Bscvalidatorset *BscvalidatorsetTransactorSession) Init() (*types.Transaction, error) {
	return _Bscvalidatorset.Contract.Init(&_Bscvalidatorset.TransactOpts)
}

// Misdemeanor is a paid mutator transaction binding the contract method 0xeb57e202.
//
// Solidity: function misdemeanor(address validator) returns()
func (_Bscvalidatorset *BscvalidatorsetTransactor) Misdemeanor(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Bscvalidatorset.contract.Transact(opts, "misdemeanor", validator)
}

// Misdemeanor is a paid mutator transaction binding the contract method 0xeb57e202.
//
// Solidity: function misdemeanor(address validator) returns()
func (_Bscvalidatorset *BscvalidatorsetSession) Misdemeanor(validator common.Address) (*types.Transaction, error) {
	return _Bscvalidatorset.Contract.Misdemeanor(&_Bscvalidatorset.TransactOpts, validator)
}

// Misdemeanor is a paid mutator transaction binding the contract method 0xeb57e202.
//
// Solidity: function misdemeanor(address validator) returns()
func (_Bscvalidatorset *BscvalidatorsetTransactorSession) Misdemeanor(validator common.Address) (*types.Transaction, error) {
	return _Bscvalidatorset.Contract.Misdemeanor(&_Bscvalidatorset.TransactOpts, validator)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Bscvalidatorset *BscvalidatorsetTransactor) UpdateParam(opts *bind.TransactOpts, key string, value []byte) (*types.Transaction, error) {
	return _Bscvalidatorset.contract.Transact(opts, "updateParam", key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Bscvalidatorset *BscvalidatorsetSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Bscvalidatorset.Contract.UpdateParam(&_Bscvalidatorset.TransactOpts, key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Bscvalidatorset *BscvalidatorsetTransactorSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Bscvalidatorset.Contract.UpdateParam(&_Bscvalidatorset.TransactOpts, key, value)
}

// BscvalidatorsetBatchTransferIterator is returned from FilterBatchTransfer and is used to iterate over the raw logs and unpacked data for BatchTransfer events raised by the Bscvalidatorset contract.
type BscvalidatorsetBatchTransferIterator struct {
	Event *BscvalidatorsetBatchTransfer // Event containing the contract specifics and raw log

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
func (it *BscvalidatorsetBatchTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscvalidatorsetBatchTransfer)
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
		it.Event = new(BscvalidatorsetBatchTransfer)
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
func (it *BscvalidatorsetBatchTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscvalidatorsetBatchTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscvalidatorsetBatchTransfer represents a BatchTransfer event raised by the Bscvalidatorset contract.
type BscvalidatorsetBatchTransfer struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBatchTransfer is a free log retrieval operation binding the contract event 0xa217d08e65f80c73121cd9db834d81652d544bfbf452f6d04922b16c90a37b70.
//
// Solidity: event batchTransfer(uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) FilterBatchTransfer(opts *bind.FilterOpts) (*BscvalidatorsetBatchTransferIterator, error) {

	logs, sub, err := _Bscvalidatorset.contract.FilterLogs(opts, "batchTransfer")
	if err != nil {
		return nil, err
	}
	return &BscvalidatorsetBatchTransferIterator{contract: _Bscvalidatorset.contract, event: "batchTransfer", logs: logs, sub: sub}, nil
}

// WatchBatchTransfer is a free log subscription operation binding the contract event 0xa217d08e65f80c73121cd9db834d81652d544bfbf452f6d04922b16c90a37b70.
//
// Solidity: event batchTransfer(uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) WatchBatchTransfer(opts *bind.WatchOpts, sink chan<- *BscvalidatorsetBatchTransfer) (event.Subscription, error) {

	logs, sub, err := _Bscvalidatorset.contract.WatchLogs(opts, "batchTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscvalidatorsetBatchTransfer)
				if err := _Bscvalidatorset.contract.UnpackLog(event, "batchTransfer", log); err != nil {
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

// ParseBatchTransfer is a log parse operation binding the contract event 0xa217d08e65f80c73121cd9db834d81652d544bfbf452f6d04922b16c90a37b70.
//
// Solidity: event batchTransfer(uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) ParseBatchTransfer(log types.Log) (*BscvalidatorsetBatchTransfer, error) {
	event := new(BscvalidatorsetBatchTransfer)
	if err := _Bscvalidatorset.contract.UnpackLog(event, "batchTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BscvalidatorsetBatchTransferFailedIterator is returned from FilterBatchTransferFailed and is used to iterate over the raw logs and unpacked data for BatchTransferFailed events raised by the Bscvalidatorset contract.
type BscvalidatorsetBatchTransferFailedIterator struct {
	Event *BscvalidatorsetBatchTransferFailed // Event containing the contract specifics and raw log

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
func (it *BscvalidatorsetBatchTransferFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscvalidatorsetBatchTransferFailed)
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
		it.Event = new(BscvalidatorsetBatchTransferFailed)
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
func (it *BscvalidatorsetBatchTransferFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscvalidatorsetBatchTransferFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscvalidatorsetBatchTransferFailed represents a BatchTransferFailed event raised by the Bscvalidatorset contract.
type BscvalidatorsetBatchTransferFailed struct {
	Amount *big.Int
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBatchTransferFailed is a free log retrieval operation binding the contract event 0xa7cdeed7d0db45e3219a6e5d60838824c16f1d39991fcfe3f963029c844bf280.
//
// Solidity: event batchTransferFailed(uint256 indexed amount, string reason)
func (_Bscvalidatorset *BscvalidatorsetFilterer) FilterBatchTransferFailed(opts *bind.FilterOpts, amount []*big.Int) (*BscvalidatorsetBatchTransferFailedIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Bscvalidatorset.contract.FilterLogs(opts, "batchTransferFailed", amountRule)
	if err != nil {
		return nil, err
	}
	return &BscvalidatorsetBatchTransferFailedIterator{contract: _Bscvalidatorset.contract, event: "batchTransferFailed", logs: logs, sub: sub}, nil
}

// WatchBatchTransferFailed is a free log subscription operation binding the contract event 0xa7cdeed7d0db45e3219a6e5d60838824c16f1d39991fcfe3f963029c844bf280.
//
// Solidity: event batchTransferFailed(uint256 indexed amount, string reason)
func (_Bscvalidatorset *BscvalidatorsetFilterer) WatchBatchTransferFailed(opts *bind.WatchOpts, sink chan<- *BscvalidatorsetBatchTransferFailed, amount []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Bscvalidatorset.contract.WatchLogs(opts, "batchTransferFailed", amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscvalidatorsetBatchTransferFailed)
				if err := _Bscvalidatorset.contract.UnpackLog(event, "batchTransferFailed", log); err != nil {
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

// ParseBatchTransferFailed is a log parse operation binding the contract event 0xa7cdeed7d0db45e3219a6e5d60838824c16f1d39991fcfe3f963029c844bf280.
//
// Solidity: event batchTransferFailed(uint256 indexed amount, string reason)
func (_Bscvalidatorset *BscvalidatorsetFilterer) ParseBatchTransferFailed(log types.Log) (*BscvalidatorsetBatchTransferFailed, error) {
	event := new(BscvalidatorsetBatchTransferFailed)
	if err := _Bscvalidatorset.contract.UnpackLog(event, "batchTransferFailed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BscvalidatorsetBatchTransferLowerFailedIterator is returned from FilterBatchTransferLowerFailed and is used to iterate over the raw logs and unpacked data for BatchTransferLowerFailed events raised by the Bscvalidatorset contract.
type BscvalidatorsetBatchTransferLowerFailedIterator struct {
	Event *BscvalidatorsetBatchTransferLowerFailed // Event containing the contract specifics and raw log

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
func (it *BscvalidatorsetBatchTransferLowerFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscvalidatorsetBatchTransferLowerFailed)
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
		it.Event = new(BscvalidatorsetBatchTransferLowerFailed)
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
func (it *BscvalidatorsetBatchTransferLowerFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscvalidatorsetBatchTransferLowerFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscvalidatorsetBatchTransferLowerFailed represents a BatchTransferLowerFailed event raised by the Bscvalidatorset contract.
type BscvalidatorsetBatchTransferLowerFailed struct {
	Amount *big.Int
	Reason []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBatchTransferLowerFailed is a free log retrieval operation binding the contract event 0xbfa884552dd8921b6ce90bfe906952ae5b3b29be0cc1a951d4f62697635a3a45.
//
// Solidity: event batchTransferLowerFailed(uint256 indexed amount, bytes reason)
func (_Bscvalidatorset *BscvalidatorsetFilterer) FilterBatchTransferLowerFailed(opts *bind.FilterOpts, amount []*big.Int) (*BscvalidatorsetBatchTransferLowerFailedIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Bscvalidatorset.contract.FilterLogs(opts, "batchTransferLowerFailed", amountRule)
	if err != nil {
		return nil, err
	}
	return &BscvalidatorsetBatchTransferLowerFailedIterator{contract: _Bscvalidatorset.contract, event: "batchTransferLowerFailed", logs: logs, sub: sub}, nil
}

// WatchBatchTransferLowerFailed is a free log subscription operation binding the contract event 0xbfa884552dd8921b6ce90bfe906952ae5b3b29be0cc1a951d4f62697635a3a45.
//
// Solidity: event batchTransferLowerFailed(uint256 indexed amount, bytes reason)
func (_Bscvalidatorset *BscvalidatorsetFilterer) WatchBatchTransferLowerFailed(opts *bind.WatchOpts, sink chan<- *BscvalidatorsetBatchTransferLowerFailed, amount []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Bscvalidatorset.contract.WatchLogs(opts, "batchTransferLowerFailed", amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscvalidatorsetBatchTransferLowerFailed)
				if err := _Bscvalidatorset.contract.UnpackLog(event, "batchTransferLowerFailed", log); err != nil {
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

// ParseBatchTransferLowerFailed is a log parse operation binding the contract event 0xbfa884552dd8921b6ce90bfe906952ae5b3b29be0cc1a951d4f62697635a3a45.
//
// Solidity: event batchTransferLowerFailed(uint256 indexed amount, bytes reason)
func (_Bscvalidatorset *BscvalidatorsetFilterer) ParseBatchTransferLowerFailed(log types.Log) (*BscvalidatorsetBatchTransferLowerFailed, error) {
	event := new(BscvalidatorsetBatchTransferLowerFailed)
	if err := _Bscvalidatorset.contract.UnpackLog(event, "batchTransferLowerFailed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BscvalidatorsetDeprecatedDepositIterator is returned from FilterDeprecatedDeposit and is used to iterate over the raw logs and unpacked data for DeprecatedDeposit events raised by the Bscvalidatorset contract.
type BscvalidatorsetDeprecatedDepositIterator struct {
	Event *BscvalidatorsetDeprecatedDeposit // Event containing the contract specifics and raw log

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
func (it *BscvalidatorsetDeprecatedDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscvalidatorsetDeprecatedDeposit)
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
		it.Event = new(BscvalidatorsetDeprecatedDeposit)
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
func (it *BscvalidatorsetDeprecatedDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscvalidatorsetDeprecatedDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscvalidatorsetDeprecatedDeposit represents a DeprecatedDeposit event raised by the Bscvalidatorset contract.
type BscvalidatorsetDeprecatedDeposit struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeprecatedDeposit is a free log retrieval operation binding the contract event 0xf177e5d6c5764d79c32883ed824111d9b13f5668cf6ab1cc12dd36791dd955b4.
//
// Solidity: event deprecatedDeposit(address indexed validator, uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) FilterDeprecatedDeposit(opts *bind.FilterOpts, validator []common.Address) (*BscvalidatorsetDeprecatedDepositIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Bscvalidatorset.contract.FilterLogs(opts, "deprecatedDeposit", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BscvalidatorsetDeprecatedDepositIterator{contract: _Bscvalidatorset.contract, event: "deprecatedDeposit", logs: logs, sub: sub}, nil
}

// WatchDeprecatedDeposit is a free log subscription operation binding the contract event 0xf177e5d6c5764d79c32883ed824111d9b13f5668cf6ab1cc12dd36791dd955b4.
//
// Solidity: event deprecatedDeposit(address indexed validator, uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) WatchDeprecatedDeposit(opts *bind.WatchOpts, sink chan<- *BscvalidatorsetDeprecatedDeposit, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Bscvalidatorset.contract.WatchLogs(opts, "deprecatedDeposit", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscvalidatorsetDeprecatedDeposit)
				if err := _Bscvalidatorset.contract.UnpackLog(event, "deprecatedDeposit", log); err != nil {
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

// ParseDeprecatedDeposit is a log parse operation binding the contract event 0xf177e5d6c5764d79c32883ed824111d9b13f5668cf6ab1cc12dd36791dd955b4.
//
// Solidity: event deprecatedDeposit(address indexed validator, uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) ParseDeprecatedDeposit(log types.Log) (*BscvalidatorsetDeprecatedDeposit, error) {
	event := new(BscvalidatorsetDeprecatedDeposit)
	if err := _Bscvalidatorset.contract.UnpackLog(event, "deprecatedDeposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BscvalidatorsetDirectTransferIterator is returned from FilterDirectTransfer and is used to iterate over the raw logs and unpacked data for DirectTransfer events raised by the Bscvalidatorset contract.
type BscvalidatorsetDirectTransferIterator struct {
	Event *BscvalidatorsetDirectTransfer // Event containing the contract specifics and raw log

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
func (it *BscvalidatorsetDirectTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscvalidatorsetDirectTransfer)
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
		it.Event = new(BscvalidatorsetDirectTransfer)
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
func (it *BscvalidatorsetDirectTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscvalidatorsetDirectTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscvalidatorsetDirectTransfer represents a DirectTransfer event raised by the Bscvalidatorset contract.
type BscvalidatorsetDirectTransfer struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDirectTransfer is a free log retrieval operation binding the contract event 0x6c61d60f69a7beb3e1c80db7f39f37b208537cbb19da3174511b477812b2fc7d.
//
// Solidity: event directTransfer(address indexed validator, uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) FilterDirectTransfer(opts *bind.FilterOpts, validator []common.Address) (*BscvalidatorsetDirectTransferIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Bscvalidatorset.contract.FilterLogs(opts, "directTransfer", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BscvalidatorsetDirectTransferIterator{contract: _Bscvalidatorset.contract, event: "directTransfer", logs: logs, sub: sub}, nil
}

// WatchDirectTransfer is a free log subscription operation binding the contract event 0x6c61d60f69a7beb3e1c80db7f39f37b208537cbb19da3174511b477812b2fc7d.
//
// Solidity: event directTransfer(address indexed validator, uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) WatchDirectTransfer(opts *bind.WatchOpts, sink chan<- *BscvalidatorsetDirectTransfer, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Bscvalidatorset.contract.WatchLogs(opts, "directTransfer", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscvalidatorsetDirectTransfer)
				if err := _Bscvalidatorset.contract.UnpackLog(event, "directTransfer", log); err != nil {
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

// ParseDirectTransfer is a log parse operation binding the contract event 0x6c61d60f69a7beb3e1c80db7f39f37b208537cbb19da3174511b477812b2fc7d.
//
// Solidity: event directTransfer(address indexed validator, uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) ParseDirectTransfer(log types.Log) (*BscvalidatorsetDirectTransfer, error) {
	event := new(BscvalidatorsetDirectTransfer)
	if err := _Bscvalidatorset.contract.UnpackLog(event, "directTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BscvalidatorsetDirectTransferFailIterator is returned from FilterDirectTransferFail and is used to iterate over the raw logs and unpacked data for DirectTransferFail events raised by the Bscvalidatorset contract.
type BscvalidatorsetDirectTransferFailIterator struct {
	Event *BscvalidatorsetDirectTransferFail // Event containing the contract specifics and raw log

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
func (it *BscvalidatorsetDirectTransferFailIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscvalidatorsetDirectTransferFail)
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
		it.Event = new(BscvalidatorsetDirectTransferFail)
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
func (it *BscvalidatorsetDirectTransferFailIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscvalidatorsetDirectTransferFailIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscvalidatorsetDirectTransferFail represents a DirectTransferFail event raised by the Bscvalidatorset contract.
type BscvalidatorsetDirectTransferFail struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDirectTransferFail is a free log retrieval operation binding the contract event 0x25d0ce7d2f0cec669a8c17efe49d195c13455bb8872b65fa610ac7f53fe4ca7d.
//
// Solidity: event directTransferFail(address indexed validator, uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) FilterDirectTransferFail(opts *bind.FilterOpts, validator []common.Address) (*BscvalidatorsetDirectTransferFailIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Bscvalidatorset.contract.FilterLogs(opts, "directTransferFail", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BscvalidatorsetDirectTransferFailIterator{contract: _Bscvalidatorset.contract, event: "directTransferFail", logs: logs, sub: sub}, nil
}

// WatchDirectTransferFail is a free log subscription operation binding the contract event 0x25d0ce7d2f0cec669a8c17efe49d195c13455bb8872b65fa610ac7f53fe4ca7d.
//
// Solidity: event directTransferFail(address indexed validator, uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) WatchDirectTransferFail(opts *bind.WatchOpts, sink chan<- *BscvalidatorsetDirectTransferFail, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Bscvalidatorset.contract.WatchLogs(opts, "directTransferFail", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscvalidatorsetDirectTransferFail)
				if err := _Bscvalidatorset.contract.UnpackLog(event, "directTransferFail", log); err != nil {
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

// ParseDirectTransferFail is a log parse operation binding the contract event 0x25d0ce7d2f0cec669a8c17efe49d195c13455bb8872b65fa610ac7f53fe4ca7d.
//
// Solidity: event directTransferFail(address indexed validator, uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) ParseDirectTransferFail(log types.Log) (*BscvalidatorsetDirectTransferFail, error) {
	event := new(BscvalidatorsetDirectTransferFail)
	if err := _Bscvalidatorset.contract.UnpackLog(event, "directTransferFail", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BscvalidatorsetFailReasonWithStrIterator is returned from FilterFailReasonWithStr and is used to iterate over the raw logs and unpacked data for FailReasonWithStr events raised by the Bscvalidatorset contract.
type BscvalidatorsetFailReasonWithStrIterator struct {
	Event *BscvalidatorsetFailReasonWithStr // Event containing the contract specifics and raw log

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
func (it *BscvalidatorsetFailReasonWithStrIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscvalidatorsetFailReasonWithStr)
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
		it.Event = new(BscvalidatorsetFailReasonWithStr)
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
func (it *BscvalidatorsetFailReasonWithStrIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscvalidatorsetFailReasonWithStrIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscvalidatorsetFailReasonWithStr represents a FailReasonWithStr event raised by the Bscvalidatorset contract.
type BscvalidatorsetFailReasonWithStr struct {
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFailReasonWithStr is a free log retrieval operation binding the contract event 0x70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb2.
//
// Solidity: event failReasonWithStr(string message)
func (_Bscvalidatorset *BscvalidatorsetFilterer) FilterFailReasonWithStr(opts *bind.FilterOpts) (*BscvalidatorsetFailReasonWithStrIterator, error) {

	logs, sub, err := _Bscvalidatorset.contract.FilterLogs(opts, "failReasonWithStr")
	if err != nil {
		return nil, err
	}
	return &BscvalidatorsetFailReasonWithStrIterator{contract: _Bscvalidatorset.contract, event: "failReasonWithStr", logs: logs, sub: sub}, nil
}

// WatchFailReasonWithStr is a free log subscription operation binding the contract event 0x70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb2.
//
// Solidity: event failReasonWithStr(string message)
func (_Bscvalidatorset *BscvalidatorsetFilterer) WatchFailReasonWithStr(opts *bind.WatchOpts, sink chan<- *BscvalidatorsetFailReasonWithStr) (event.Subscription, error) {

	logs, sub, err := _Bscvalidatorset.contract.WatchLogs(opts, "failReasonWithStr")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscvalidatorsetFailReasonWithStr)
				if err := _Bscvalidatorset.contract.UnpackLog(event, "failReasonWithStr", log); err != nil {
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

// ParseFailReasonWithStr is a log parse operation binding the contract event 0x70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb2.
//
// Solidity: event failReasonWithStr(string message)
func (_Bscvalidatorset *BscvalidatorsetFilterer) ParseFailReasonWithStr(log types.Log) (*BscvalidatorsetFailReasonWithStr, error) {
	event := new(BscvalidatorsetFailReasonWithStr)
	if err := _Bscvalidatorset.contract.UnpackLog(event, "failReasonWithStr", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BscvalidatorsetParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the Bscvalidatorset contract.
type BscvalidatorsetParamChangeIterator struct {
	Event *BscvalidatorsetParamChange // Event containing the contract specifics and raw log

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
func (it *BscvalidatorsetParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscvalidatorsetParamChange)
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
		it.Event = new(BscvalidatorsetParamChange)
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
func (it *BscvalidatorsetParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscvalidatorsetParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscvalidatorsetParamChange represents a ParamChange event raised by the Bscvalidatorset contract.
type BscvalidatorsetParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Bscvalidatorset *BscvalidatorsetFilterer) FilterParamChange(opts *bind.FilterOpts) (*BscvalidatorsetParamChangeIterator, error) {

	logs, sub, err := _Bscvalidatorset.contract.FilterLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return &BscvalidatorsetParamChangeIterator{contract: _Bscvalidatorset.contract, event: "paramChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Bscvalidatorset *BscvalidatorsetFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *BscvalidatorsetParamChange) (event.Subscription, error) {

	logs, sub, err := _Bscvalidatorset.contract.WatchLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscvalidatorsetParamChange)
				if err := _Bscvalidatorset.contract.UnpackLog(event, "paramChange", log); err != nil {
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

// ParseParamChange is a log parse operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Bscvalidatorset *BscvalidatorsetFilterer) ParseParamChange(log types.Log) (*BscvalidatorsetParamChange, error) {
	event := new(BscvalidatorsetParamChange)
	if err := _Bscvalidatorset.contract.UnpackLog(event, "paramChange", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BscvalidatorsetSystemTransferIterator is returned from FilterSystemTransfer and is used to iterate over the raw logs and unpacked data for SystemTransfer events raised by the Bscvalidatorset contract.
type BscvalidatorsetSystemTransferIterator struct {
	Event *BscvalidatorsetSystemTransfer // Event containing the contract specifics and raw log

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
func (it *BscvalidatorsetSystemTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscvalidatorsetSystemTransfer)
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
		it.Event = new(BscvalidatorsetSystemTransfer)
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
func (it *BscvalidatorsetSystemTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscvalidatorsetSystemTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscvalidatorsetSystemTransfer represents a SystemTransfer event raised by the Bscvalidatorset contract.
type BscvalidatorsetSystemTransfer struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSystemTransfer is a free log retrieval operation binding the contract event 0x6ecc855f9440a9282c90913bbc91619fd44f5ec0b462af28d127b116f130aa4d.
//
// Solidity: event systemTransfer(uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) FilterSystemTransfer(opts *bind.FilterOpts) (*BscvalidatorsetSystemTransferIterator, error) {

	logs, sub, err := _Bscvalidatorset.contract.FilterLogs(opts, "systemTransfer")
	if err != nil {
		return nil, err
	}
	return &BscvalidatorsetSystemTransferIterator{contract: _Bscvalidatorset.contract, event: "systemTransfer", logs: logs, sub: sub}, nil
}

// WatchSystemTransfer is a free log subscription operation binding the contract event 0x6ecc855f9440a9282c90913bbc91619fd44f5ec0b462af28d127b116f130aa4d.
//
// Solidity: event systemTransfer(uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) WatchSystemTransfer(opts *bind.WatchOpts, sink chan<- *BscvalidatorsetSystemTransfer) (event.Subscription, error) {

	logs, sub, err := _Bscvalidatorset.contract.WatchLogs(opts, "systemTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscvalidatorsetSystemTransfer)
				if err := _Bscvalidatorset.contract.UnpackLog(event, "systemTransfer", log); err != nil {
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

// ParseSystemTransfer is a log parse operation binding the contract event 0x6ecc855f9440a9282c90913bbc91619fd44f5ec0b462af28d127b116f130aa4d.
//
// Solidity: event systemTransfer(uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) ParseSystemTransfer(log types.Log) (*BscvalidatorsetSystemTransfer, error) {
	event := new(BscvalidatorsetSystemTransfer)
	if err := _Bscvalidatorset.contract.UnpackLog(event, "systemTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BscvalidatorsetUnexpectedPackageIterator is returned from FilterUnexpectedPackage and is used to iterate over the raw logs and unpacked data for UnexpectedPackage events raised by the Bscvalidatorset contract.
type BscvalidatorsetUnexpectedPackageIterator struct {
	Event *BscvalidatorsetUnexpectedPackage // Event containing the contract specifics and raw log

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
func (it *BscvalidatorsetUnexpectedPackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscvalidatorsetUnexpectedPackage)
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
		it.Event = new(BscvalidatorsetUnexpectedPackage)
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
func (it *BscvalidatorsetUnexpectedPackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscvalidatorsetUnexpectedPackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscvalidatorsetUnexpectedPackage represents a UnexpectedPackage event raised by the Bscvalidatorset contract.
type BscvalidatorsetUnexpectedPackage struct {
	ChannelId uint8
	MsgBytes  []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnexpectedPackage is a free log retrieval operation binding the contract event 0x41ce201247b6ceb957dcdb217d0b8acb50b9ea0e12af9af4f5e7f38902101605.
//
// Solidity: event unexpectedPackage(uint8 channelId, bytes msgBytes)
func (_Bscvalidatorset *BscvalidatorsetFilterer) FilterUnexpectedPackage(opts *bind.FilterOpts) (*BscvalidatorsetUnexpectedPackageIterator, error) {

	logs, sub, err := _Bscvalidatorset.contract.FilterLogs(opts, "unexpectedPackage")
	if err != nil {
		return nil, err
	}
	return &BscvalidatorsetUnexpectedPackageIterator{contract: _Bscvalidatorset.contract, event: "unexpectedPackage", logs: logs, sub: sub}, nil
}

// WatchUnexpectedPackage is a free log subscription operation binding the contract event 0x41ce201247b6ceb957dcdb217d0b8acb50b9ea0e12af9af4f5e7f38902101605.
//
// Solidity: event unexpectedPackage(uint8 channelId, bytes msgBytes)
func (_Bscvalidatorset *BscvalidatorsetFilterer) WatchUnexpectedPackage(opts *bind.WatchOpts, sink chan<- *BscvalidatorsetUnexpectedPackage) (event.Subscription, error) {

	logs, sub, err := _Bscvalidatorset.contract.WatchLogs(opts, "unexpectedPackage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscvalidatorsetUnexpectedPackage)
				if err := _Bscvalidatorset.contract.UnpackLog(event, "unexpectedPackage", log); err != nil {
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

// ParseUnexpectedPackage is a log parse operation binding the contract event 0x41ce201247b6ceb957dcdb217d0b8acb50b9ea0e12af9af4f5e7f38902101605.
//
// Solidity: event unexpectedPackage(uint8 channelId, bytes msgBytes)
func (_Bscvalidatorset *BscvalidatorsetFilterer) ParseUnexpectedPackage(log types.Log) (*BscvalidatorsetUnexpectedPackage, error) {
	event := new(BscvalidatorsetUnexpectedPackage)
	if err := _Bscvalidatorset.contract.UnpackLog(event, "unexpectedPackage", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BscvalidatorsetValidatorDepositIterator is returned from FilterValidatorDeposit and is used to iterate over the raw logs and unpacked data for ValidatorDeposit events raised by the Bscvalidatorset contract.
type BscvalidatorsetValidatorDepositIterator struct {
	Event *BscvalidatorsetValidatorDeposit // Event containing the contract specifics and raw log

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
func (it *BscvalidatorsetValidatorDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscvalidatorsetValidatorDeposit)
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
		it.Event = new(BscvalidatorsetValidatorDeposit)
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
func (it *BscvalidatorsetValidatorDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscvalidatorsetValidatorDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscvalidatorsetValidatorDeposit represents a ValidatorDeposit event raised by the Bscvalidatorset contract.
type BscvalidatorsetValidatorDeposit struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorDeposit is a free log retrieval operation binding the contract event 0x93a090ecc682c002995fad3c85b30c5651d7fd29b0be5da9d784a3302aedc055.
//
// Solidity: event validatorDeposit(address indexed validator, uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) FilterValidatorDeposit(opts *bind.FilterOpts, validator []common.Address) (*BscvalidatorsetValidatorDepositIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Bscvalidatorset.contract.FilterLogs(opts, "validatorDeposit", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BscvalidatorsetValidatorDepositIterator{contract: _Bscvalidatorset.contract, event: "validatorDeposit", logs: logs, sub: sub}, nil
}

// WatchValidatorDeposit is a free log subscription operation binding the contract event 0x93a090ecc682c002995fad3c85b30c5651d7fd29b0be5da9d784a3302aedc055.
//
// Solidity: event validatorDeposit(address indexed validator, uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) WatchValidatorDeposit(opts *bind.WatchOpts, sink chan<- *BscvalidatorsetValidatorDeposit, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Bscvalidatorset.contract.WatchLogs(opts, "validatorDeposit", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscvalidatorsetValidatorDeposit)
				if err := _Bscvalidatorset.contract.UnpackLog(event, "validatorDeposit", log); err != nil {
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

// ParseValidatorDeposit is a log parse operation binding the contract event 0x93a090ecc682c002995fad3c85b30c5651d7fd29b0be5da9d784a3302aedc055.
//
// Solidity: event validatorDeposit(address indexed validator, uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) ParseValidatorDeposit(log types.Log) (*BscvalidatorsetValidatorDeposit, error) {
	event := new(BscvalidatorsetValidatorDeposit)
	if err := _Bscvalidatorset.contract.UnpackLog(event, "validatorDeposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BscvalidatorsetValidatorEmptyJailedIterator is returned from FilterValidatorEmptyJailed and is used to iterate over the raw logs and unpacked data for ValidatorEmptyJailed events raised by the Bscvalidatorset contract.
type BscvalidatorsetValidatorEmptyJailedIterator struct {
	Event *BscvalidatorsetValidatorEmptyJailed // Event containing the contract specifics and raw log

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
func (it *BscvalidatorsetValidatorEmptyJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscvalidatorsetValidatorEmptyJailed)
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
		it.Event = new(BscvalidatorsetValidatorEmptyJailed)
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
func (it *BscvalidatorsetValidatorEmptyJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscvalidatorsetValidatorEmptyJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscvalidatorsetValidatorEmptyJailed represents a ValidatorEmptyJailed event raised by the Bscvalidatorset contract.
type BscvalidatorsetValidatorEmptyJailed struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorEmptyJailed is a free log retrieval operation binding the contract event 0xe209c46bebf57cf265d5d9009a00870e256d9150f3ed5281ab9d9eb3cec6e4be.
//
// Solidity: event validatorEmptyJailed(address indexed validator)
func (_Bscvalidatorset *BscvalidatorsetFilterer) FilterValidatorEmptyJailed(opts *bind.FilterOpts, validator []common.Address) (*BscvalidatorsetValidatorEmptyJailedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Bscvalidatorset.contract.FilterLogs(opts, "validatorEmptyJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BscvalidatorsetValidatorEmptyJailedIterator{contract: _Bscvalidatorset.contract, event: "validatorEmptyJailed", logs: logs, sub: sub}, nil
}

// WatchValidatorEmptyJailed is a free log subscription operation binding the contract event 0xe209c46bebf57cf265d5d9009a00870e256d9150f3ed5281ab9d9eb3cec6e4be.
//
// Solidity: event validatorEmptyJailed(address indexed validator)
func (_Bscvalidatorset *BscvalidatorsetFilterer) WatchValidatorEmptyJailed(opts *bind.WatchOpts, sink chan<- *BscvalidatorsetValidatorEmptyJailed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Bscvalidatorset.contract.WatchLogs(opts, "validatorEmptyJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscvalidatorsetValidatorEmptyJailed)
				if err := _Bscvalidatorset.contract.UnpackLog(event, "validatorEmptyJailed", log); err != nil {
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

// ParseValidatorEmptyJailed is a log parse operation binding the contract event 0xe209c46bebf57cf265d5d9009a00870e256d9150f3ed5281ab9d9eb3cec6e4be.
//
// Solidity: event validatorEmptyJailed(address indexed validator)
func (_Bscvalidatorset *BscvalidatorsetFilterer) ParseValidatorEmptyJailed(log types.Log) (*BscvalidatorsetValidatorEmptyJailed, error) {
	event := new(BscvalidatorsetValidatorEmptyJailed)
	if err := _Bscvalidatorset.contract.UnpackLog(event, "validatorEmptyJailed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BscvalidatorsetValidatorFelonyIterator is returned from FilterValidatorFelony and is used to iterate over the raw logs and unpacked data for ValidatorFelony events raised by the Bscvalidatorset contract.
type BscvalidatorsetValidatorFelonyIterator struct {
	Event *BscvalidatorsetValidatorFelony // Event containing the contract specifics and raw log

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
func (it *BscvalidatorsetValidatorFelonyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscvalidatorsetValidatorFelony)
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
		it.Event = new(BscvalidatorsetValidatorFelony)
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
func (it *BscvalidatorsetValidatorFelonyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscvalidatorsetValidatorFelonyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscvalidatorsetValidatorFelony represents a ValidatorFelony event raised by the Bscvalidatorset contract.
type BscvalidatorsetValidatorFelony struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorFelony is a free log retrieval operation binding the contract event 0x3b6f9ef90462b512a1293ecec018670bf7b7f1876fb727590a8a6d7643130a70.
//
// Solidity: event validatorFelony(address indexed validator, uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) FilterValidatorFelony(opts *bind.FilterOpts, validator []common.Address) (*BscvalidatorsetValidatorFelonyIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Bscvalidatorset.contract.FilterLogs(opts, "validatorFelony", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BscvalidatorsetValidatorFelonyIterator{contract: _Bscvalidatorset.contract, event: "validatorFelony", logs: logs, sub: sub}, nil
}

// WatchValidatorFelony is a free log subscription operation binding the contract event 0x3b6f9ef90462b512a1293ecec018670bf7b7f1876fb727590a8a6d7643130a70.
//
// Solidity: event validatorFelony(address indexed validator, uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) WatchValidatorFelony(opts *bind.WatchOpts, sink chan<- *BscvalidatorsetValidatorFelony, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Bscvalidatorset.contract.WatchLogs(opts, "validatorFelony", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscvalidatorsetValidatorFelony)
				if err := _Bscvalidatorset.contract.UnpackLog(event, "validatorFelony", log); err != nil {
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

// ParseValidatorFelony is a log parse operation binding the contract event 0x3b6f9ef90462b512a1293ecec018670bf7b7f1876fb727590a8a6d7643130a70.
//
// Solidity: event validatorFelony(address indexed validator, uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) ParseValidatorFelony(log types.Log) (*BscvalidatorsetValidatorFelony, error) {
	event := new(BscvalidatorsetValidatorFelony)
	if err := _Bscvalidatorset.contract.UnpackLog(event, "validatorFelony", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BscvalidatorsetValidatorJailedIterator is returned from FilterValidatorJailed and is used to iterate over the raw logs and unpacked data for ValidatorJailed events raised by the Bscvalidatorset contract.
type BscvalidatorsetValidatorJailedIterator struct {
	Event *BscvalidatorsetValidatorJailed // Event containing the contract specifics and raw log

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
func (it *BscvalidatorsetValidatorJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscvalidatorsetValidatorJailed)
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
		it.Event = new(BscvalidatorsetValidatorJailed)
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
func (it *BscvalidatorsetValidatorJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscvalidatorsetValidatorJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscvalidatorsetValidatorJailed represents a ValidatorJailed event raised by the Bscvalidatorset contract.
type BscvalidatorsetValidatorJailed struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorJailed is a free log retrieval operation binding the contract event 0xf226e7d8f547ff903d9d419cf5f54e0d7d07efa9584135a53a057c5f1f27f49a.
//
// Solidity: event validatorJailed(address indexed validator)
func (_Bscvalidatorset *BscvalidatorsetFilterer) FilterValidatorJailed(opts *bind.FilterOpts, validator []common.Address) (*BscvalidatorsetValidatorJailedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Bscvalidatorset.contract.FilterLogs(opts, "validatorJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BscvalidatorsetValidatorJailedIterator{contract: _Bscvalidatorset.contract, event: "validatorJailed", logs: logs, sub: sub}, nil
}

// WatchValidatorJailed is a free log subscription operation binding the contract event 0xf226e7d8f547ff903d9d419cf5f54e0d7d07efa9584135a53a057c5f1f27f49a.
//
// Solidity: event validatorJailed(address indexed validator)
func (_Bscvalidatorset *BscvalidatorsetFilterer) WatchValidatorJailed(opts *bind.WatchOpts, sink chan<- *BscvalidatorsetValidatorJailed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Bscvalidatorset.contract.WatchLogs(opts, "validatorJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscvalidatorsetValidatorJailed)
				if err := _Bscvalidatorset.contract.UnpackLog(event, "validatorJailed", log); err != nil {
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

// ParseValidatorJailed is a log parse operation binding the contract event 0xf226e7d8f547ff903d9d419cf5f54e0d7d07efa9584135a53a057c5f1f27f49a.
//
// Solidity: event validatorJailed(address indexed validator)
func (_Bscvalidatorset *BscvalidatorsetFilterer) ParseValidatorJailed(log types.Log) (*BscvalidatorsetValidatorJailed, error) {
	event := new(BscvalidatorsetValidatorJailed)
	if err := _Bscvalidatorset.contract.UnpackLog(event, "validatorJailed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BscvalidatorsetValidatorMisdemeanorIterator is returned from FilterValidatorMisdemeanor and is used to iterate over the raw logs and unpacked data for ValidatorMisdemeanor events raised by the Bscvalidatorset contract.
type BscvalidatorsetValidatorMisdemeanorIterator struct {
	Event *BscvalidatorsetValidatorMisdemeanor // Event containing the contract specifics and raw log

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
func (it *BscvalidatorsetValidatorMisdemeanorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscvalidatorsetValidatorMisdemeanor)
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
		it.Event = new(BscvalidatorsetValidatorMisdemeanor)
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
func (it *BscvalidatorsetValidatorMisdemeanorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscvalidatorsetValidatorMisdemeanorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscvalidatorsetValidatorMisdemeanor represents a ValidatorMisdemeanor event raised by the Bscvalidatorset contract.
type BscvalidatorsetValidatorMisdemeanor struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorMisdemeanor is a free log retrieval operation binding the contract event 0x8cd4e147d8af98a9e3b6724021b8bf6aed2e5dac71c38f2dce8161b82585b25d.
//
// Solidity: event validatorMisdemeanor(address indexed validator, uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) FilterValidatorMisdemeanor(opts *bind.FilterOpts, validator []common.Address) (*BscvalidatorsetValidatorMisdemeanorIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Bscvalidatorset.contract.FilterLogs(opts, "validatorMisdemeanor", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BscvalidatorsetValidatorMisdemeanorIterator{contract: _Bscvalidatorset.contract, event: "validatorMisdemeanor", logs: logs, sub: sub}, nil
}

// WatchValidatorMisdemeanor is a free log subscription operation binding the contract event 0x8cd4e147d8af98a9e3b6724021b8bf6aed2e5dac71c38f2dce8161b82585b25d.
//
// Solidity: event validatorMisdemeanor(address indexed validator, uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) WatchValidatorMisdemeanor(opts *bind.WatchOpts, sink chan<- *BscvalidatorsetValidatorMisdemeanor, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Bscvalidatorset.contract.WatchLogs(opts, "validatorMisdemeanor", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscvalidatorsetValidatorMisdemeanor)
				if err := _Bscvalidatorset.contract.UnpackLog(event, "validatorMisdemeanor", log); err != nil {
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

// ParseValidatorMisdemeanor is a log parse operation binding the contract event 0x8cd4e147d8af98a9e3b6724021b8bf6aed2e5dac71c38f2dce8161b82585b25d.
//
// Solidity: event validatorMisdemeanor(address indexed validator, uint256 amount)
func (_Bscvalidatorset *BscvalidatorsetFilterer) ParseValidatorMisdemeanor(log types.Log) (*BscvalidatorsetValidatorMisdemeanor, error) {
	event := new(BscvalidatorsetValidatorMisdemeanor)
	if err := _Bscvalidatorset.contract.UnpackLog(event, "validatorMisdemeanor", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BscvalidatorsetValidatorSetUpdatedIterator is returned from FilterValidatorSetUpdated and is used to iterate over the raw logs and unpacked data for ValidatorSetUpdated events raised by the Bscvalidatorset contract.
type BscvalidatorsetValidatorSetUpdatedIterator struct {
	Event *BscvalidatorsetValidatorSetUpdated // Event containing the contract specifics and raw log

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
func (it *BscvalidatorsetValidatorSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscvalidatorsetValidatorSetUpdated)
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
		it.Event = new(BscvalidatorsetValidatorSetUpdated)
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
func (it *BscvalidatorsetValidatorSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscvalidatorsetValidatorSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscvalidatorsetValidatorSetUpdated represents a ValidatorSetUpdated event raised by the Bscvalidatorset contract.
type BscvalidatorsetValidatorSetUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetUpdated is a free log retrieval operation binding the contract event 0xedd8d7296956dd970ab4de3f2fc03be2b0ffc615d20cd4c72c6e44f928630ebf.
//
// Solidity: event validatorSetUpdated()
func (_Bscvalidatorset *BscvalidatorsetFilterer) FilterValidatorSetUpdated(opts *bind.FilterOpts) (*BscvalidatorsetValidatorSetUpdatedIterator, error) {

	logs, sub, err := _Bscvalidatorset.contract.FilterLogs(opts, "validatorSetUpdated")
	if err != nil {
		return nil, err
	}
	return &BscvalidatorsetValidatorSetUpdatedIterator{contract: _Bscvalidatorset.contract, event: "validatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorSetUpdated is a free log subscription operation binding the contract event 0xedd8d7296956dd970ab4de3f2fc03be2b0ffc615d20cd4c72c6e44f928630ebf.
//
// Solidity: event validatorSetUpdated()
func (_Bscvalidatorset *BscvalidatorsetFilterer) WatchValidatorSetUpdated(opts *bind.WatchOpts, sink chan<- *BscvalidatorsetValidatorSetUpdated) (event.Subscription, error) {

	logs, sub, err := _Bscvalidatorset.contract.WatchLogs(opts, "validatorSetUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscvalidatorsetValidatorSetUpdated)
				if err := _Bscvalidatorset.contract.UnpackLog(event, "validatorSetUpdated", log); err != nil {
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

// ParseValidatorSetUpdated is a log parse operation binding the contract event 0xedd8d7296956dd970ab4de3f2fc03be2b0ffc615d20cd4c72c6e44f928630ebf.
//
// Solidity: event validatorSetUpdated()
func (_Bscvalidatorset *BscvalidatorsetFilterer) ParseValidatorSetUpdated(log types.Log) (*BscvalidatorsetValidatorSetUpdated, error) {
	event := new(BscvalidatorsetValidatorSetUpdated)
	if err := _Bscvalidatorset.contract.UnpackLog(event, "validatorSetUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
