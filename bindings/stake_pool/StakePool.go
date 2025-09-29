// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stake_pool

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// StakePoolMetaData contains all meta data concerning the StakePool contract.
var StakePoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegate\",\"inputs\":[{\"name\":\"_depositToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDelegated\",\"inputs\":[{\"name\":\"_stablecoin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_stakeManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_govInstantManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_govOracleAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ondoInstantManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOndoInstantManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ondoOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOndoOracle\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payMissingUnbondingFee\",\"inputs\":[{\"name\":\"_stablecoin\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakeManagerAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalMissingUnbondingFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"undelegate\",\"inputs\":[{\"name\":\"_receivingToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_undelegateAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawForStaker\",\"inputs\":[{\"name\":\"_receivingToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Delegate\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"stablecoin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MissingUnbondingFee\",\"inputs\":[{\"name\":\"stablecoin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PayMissingUnbondingFee\",\"inputs\":[{\"name\":\"payer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"stablecoin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Undelegate\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"rwaToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawForStaker\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"receivingToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CommissionRateInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedToCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedToWithdrawForStaker\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LengthNotMatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAuthorizedLsdToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ValidatorsEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorsLenExceedLimit\",\"inputs\":[]}]",
}

// StakePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use StakePoolMetaData.ABI instead.
var StakePoolABI = StakePoolMetaData.ABI

// StakePool is an auto generated Go binding around an Ethereum contract.
type StakePool struct {
	StakePoolCaller     // Read-only binding to the contract
	StakePoolTransactor // Write-only binding to the contract
	StakePoolFilterer   // Log filterer for contract events
}

// StakePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakePoolSession struct {
	Contract     *StakePool        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakePoolCallerSession struct {
	Contract *StakePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// StakePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakePoolTransactorSession struct {
	Contract     *StakePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StakePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakePoolRaw struct {
	Contract *StakePool // Generic contract binding to access the raw methods on
}

// StakePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakePoolCallerRaw struct {
	Contract *StakePoolCaller // Generic read-only contract binding to access the raw methods on
}

// StakePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakePoolTransactorRaw struct {
	Contract *StakePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakePool creates a new instance of StakePool, bound to a specific deployed contract.
func NewStakePool(address common.Address, backend bind.ContractBackend) (*StakePool, error) {
	contract, err := bindStakePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakePool{StakePoolCaller: StakePoolCaller{contract: contract}, StakePoolTransactor: StakePoolTransactor{contract: contract}, StakePoolFilterer: StakePoolFilterer{contract: contract}}, nil
}

// NewStakePoolCaller creates a new read-only instance of StakePool, bound to a specific deployed contract.
func NewStakePoolCaller(address common.Address, caller bind.ContractCaller) (*StakePoolCaller, error) {
	contract, err := bindStakePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakePoolCaller{contract: contract}, nil
}

// NewStakePoolTransactor creates a new write-only instance of StakePool, bound to a specific deployed contract.
func NewStakePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*StakePoolTransactor, error) {
	contract, err := bindStakePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakePoolTransactor{contract: contract}, nil
}

// NewStakePoolFilterer creates a new log filterer instance of StakePool, bound to a specific deployed contract.
func NewStakePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*StakePoolFilterer, error) {
	contract, err := bindStakePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakePoolFilterer{contract: contract}, nil
}

// bindStakePool binds a generic wrapper to an already deployed contract.
func bindStakePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakePoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakePool *StakePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakePool.Contract.StakePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakePool *StakePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakePool.Contract.StakePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakePool *StakePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakePool.Contract.StakePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakePool *StakePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakePool *StakePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakePool *StakePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakePool.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_StakePool *StakePoolCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakePool.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_StakePool *StakePoolSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _StakePool.Contract.UPGRADEINTERFACEVERSION(&_StakePool.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_StakePool *StakePoolCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _StakePool.Contract.UPGRADEINTERFACEVERSION(&_StakePool.CallOpts)
}

// GetDelegated is a free data retrieval call binding the contract method 0x56f0b57e.
//
// Solidity: function getDelegated(address _stablecoin) view returns(uint256)
func (_StakePool *StakePoolCaller) GetDelegated(opts *bind.CallOpts, _stablecoin common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakePool.contract.Call(opts, &out, "getDelegated", _stablecoin)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDelegated is a free data retrieval call binding the contract method 0x56f0b57e.
//
// Solidity: function getDelegated(address _stablecoin) view returns(uint256)
func (_StakePool *StakePoolSession) GetDelegated(_stablecoin common.Address) (*big.Int, error) {
	return _StakePool.Contract.GetDelegated(&_StakePool.CallOpts, _stablecoin)
}

// GetDelegated is a free data retrieval call binding the contract method 0x56f0b57e.
//
// Solidity: function getDelegated(address _stablecoin) view returns(uint256)
func (_StakePool *StakePoolCallerSession) GetDelegated(_stablecoin common.Address) (*big.Int, error) {
	return _StakePool.Contract.GetDelegated(&_StakePool.CallOpts, _stablecoin)
}

// OndoInstantManager is a free data retrieval call binding the contract method 0x7a764a2e.
//
// Solidity: function ondoInstantManager() view returns(address)
func (_StakePool *StakePoolCaller) OndoInstantManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakePool.contract.Call(opts, &out, "ondoInstantManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OndoInstantManager is a free data retrieval call binding the contract method 0x7a764a2e.
//
// Solidity: function ondoInstantManager() view returns(address)
func (_StakePool *StakePoolSession) OndoInstantManager() (common.Address, error) {
	return _StakePool.Contract.OndoInstantManager(&_StakePool.CallOpts)
}

// OndoInstantManager is a free data retrieval call binding the contract method 0x7a764a2e.
//
// Solidity: function ondoInstantManager() view returns(address)
func (_StakePool *StakePoolCallerSession) OndoInstantManager() (common.Address, error) {
	return _StakePool.Contract.OndoInstantManager(&_StakePool.CallOpts)
}

// OndoOracle is a free data retrieval call binding the contract method 0x21e0e0cf.
//
// Solidity: function ondoOracle() view returns(address)
func (_StakePool *StakePoolCaller) OndoOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakePool.contract.Call(opts, &out, "ondoOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OndoOracle is a free data retrieval call binding the contract method 0x21e0e0cf.
//
// Solidity: function ondoOracle() view returns(address)
func (_StakePool *StakePoolSession) OndoOracle() (common.Address, error) {
	return _StakePool.Contract.OndoOracle(&_StakePool.CallOpts)
}

// OndoOracle is a free data retrieval call binding the contract method 0x21e0e0cf.
//
// Solidity: function ondoOracle() view returns(address)
func (_StakePool *StakePoolCallerSession) OndoOracle() (common.Address, error) {
	return _StakePool.Contract.OndoOracle(&_StakePool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakePool *StakePoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakePool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakePool *StakePoolSession) Owner() (common.Address, error) {
	return _StakePool.Contract.Owner(&_StakePool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakePool *StakePoolCallerSession) Owner() (common.Address, error) {
	return _StakePool.Contract.Owner(&_StakePool.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StakePool *StakePoolCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakePool.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StakePool *StakePoolSession) ProxiableUUID() ([32]byte, error) {
	return _StakePool.Contract.ProxiableUUID(&_StakePool.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StakePool *StakePoolCallerSession) ProxiableUUID() ([32]byte, error) {
	return _StakePool.Contract.ProxiableUUID(&_StakePool.CallOpts)
}

// StakeManagerAddress is a free data retrieval call binding the contract method 0x15f69471.
//
// Solidity: function stakeManagerAddress() view returns(address)
func (_StakePool *StakePoolCaller) StakeManagerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakePool.contract.Call(opts, &out, "stakeManagerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeManagerAddress is a free data retrieval call binding the contract method 0x15f69471.
//
// Solidity: function stakeManagerAddress() view returns(address)
func (_StakePool *StakePoolSession) StakeManagerAddress() (common.Address, error) {
	return _StakePool.Contract.StakeManagerAddress(&_StakePool.CallOpts)
}

// StakeManagerAddress is a free data retrieval call binding the contract method 0x15f69471.
//
// Solidity: function stakeManagerAddress() view returns(address)
func (_StakePool *StakePoolCallerSession) StakeManagerAddress() (common.Address, error) {
	return _StakePool.Contract.StakeManagerAddress(&_StakePool.CallOpts)
}

// TotalMissingUnbondingFee is a free data retrieval call binding the contract method 0xab9466d2.
//
// Solidity: function totalMissingUnbondingFee(address ) view returns(uint256)
func (_StakePool *StakePoolCaller) TotalMissingUnbondingFee(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakePool.contract.Call(opts, &out, "totalMissingUnbondingFee", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMissingUnbondingFee is a free data retrieval call binding the contract method 0xab9466d2.
//
// Solidity: function totalMissingUnbondingFee(address ) view returns(uint256)
func (_StakePool *StakePoolSession) TotalMissingUnbondingFee(arg0 common.Address) (*big.Int, error) {
	return _StakePool.Contract.TotalMissingUnbondingFee(&_StakePool.CallOpts, arg0)
}

// TotalMissingUnbondingFee is a free data retrieval call binding the contract method 0xab9466d2.
//
// Solidity: function totalMissingUnbondingFee(address ) view returns(uint256)
func (_StakePool *StakePoolCallerSession) TotalMissingUnbondingFee(arg0 common.Address) (*big.Int, error) {
	return _StakePool.Contract.TotalMissingUnbondingFee(&_StakePool.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint64)
func (_StakePool *StakePoolCaller) Version(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _StakePool.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint64)
func (_StakePool *StakePoolSession) Version() (uint64, error) {
	return _StakePool.Contract.Version(&_StakePool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint64)
func (_StakePool *StakePoolCallerSession) Version() (uint64, error) {
	return _StakePool.Contract.Version(&_StakePool.CallOpts)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address _depositToken, uint256 _amount) returns(uint256)
func (_StakePool *StakePoolTransactor) Delegate(opts *bind.TransactOpts, _depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakePool.contract.Transact(opts, "delegate", _depositToken, _amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address _depositToken, uint256 _amount) returns(uint256)
func (_StakePool *StakePoolSession) Delegate(_depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakePool.Contract.Delegate(&_StakePool.TransactOpts, _depositToken, _amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address _depositToken, uint256 _amount) returns(uint256)
func (_StakePool *StakePoolTransactorSession) Delegate(_depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakePool.Contract.Delegate(&_StakePool.TransactOpts, _depositToken, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _stakeManagerAddress, address _govInstantManagerAddress, address _govOracleAddress, address _owner) returns()
func (_StakePool *StakePoolTransactor) Initialize(opts *bind.TransactOpts, _stakeManagerAddress common.Address, _govInstantManagerAddress common.Address, _govOracleAddress common.Address, _owner common.Address) (*types.Transaction, error) {
	return _StakePool.contract.Transact(opts, "initialize", _stakeManagerAddress, _govInstantManagerAddress, _govOracleAddress, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _stakeManagerAddress, address _govInstantManagerAddress, address _govOracleAddress, address _owner) returns()
func (_StakePool *StakePoolSession) Initialize(_stakeManagerAddress common.Address, _govInstantManagerAddress common.Address, _govOracleAddress common.Address, _owner common.Address) (*types.Transaction, error) {
	return _StakePool.Contract.Initialize(&_StakePool.TransactOpts, _stakeManagerAddress, _govInstantManagerAddress, _govOracleAddress, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _stakeManagerAddress, address _govInstantManagerAddress, address _govOracleAddress, address _owner) returns()
func (_StakePool *StakePoolTransactorSession) Initialize(_stakeManagerAddress common.Address, _govInstantManagerAddress common.Address, _govOracleAddress common.Address, _owner common.Address) (*types.Transaction, error) {
	return _StakePool.Contract.Initialize(&_StakePool.TransactOpts, _stakeManagerAddress, _govInstantManagerAddress, _govOracleAddress, _owner)
}

// PayMissingUnbondingFee is a paid mutator transaction binding the contract method 0xb355e98a.
//
// Solidity: function payMissingUnbondingFee(address[] _stablecoin, uint256[] _amount) returns()
func (_StakePool *StakePoolTransactor) PayMissingUnbondingFee(opts *bind.TransactOpts, _stablecoin []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _StakePool.contract.Transact(opts, "payMissingUnbondingFee", _stablecoin, _amount)
}

// PayMissingUnbondingFee is a paid mutator transaction binding the contract method 0xb355e98a.
//
// Solidity: function payMissingUnbondingFee(address[] _stablecoin, uint256[] _amount) returns()
func (_StakePool *StakePoolSession) PayMissingUnbondingFee(_stablecoin []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _StakePool.Contract.PayMissingUnbondingFee(&_StakePool.TransactOpts, _stablecoin, _amount)
}

// PayMissingUnbondingFee is a paid mutator transaction binding the contract method 0xb355e98a.
//
// Solidity: function payMissingUnbondingFee(address[] _stablecoin, uint256[] _amount) returns()
func (_StakePool *StakePoolTransactorSession) PayMissingUnbondingFee(_stablecoin []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _StakePool.Contract.PayMissingUnbondingFee(&_StakePool.TransactOpts, _stablecoin, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_StakePool *StakePoolTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _StakePool.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_StakePool *StakePoolSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _StakePool.Contract.TransferOwnership(&_StakePool.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_StakePool *StakePoolTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _StakePool.Contract.TransferOwnership(&_StakePool.TransactOpts, _newOwner)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _receivingToken, uint256 _undelegateAmount) returns(uint256)
func (_StakePool *StakePoolTransactor) Undelegate(opts *bind.TransactOpts, _receivingToken common.Address, _undelegateAmount *big.Int) (*types.Transaction, error) {
	return _StakePool.contract.Transact(opts, "undelegate", _receivingToken, _undelegateAmount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _receivingToken, uint256 _undelegateAmount) returns(uint256)
func (_StakePool *StakePoolSession) Undelegate(_receivingToken common.Address, _undelegateAmount *big.Int) (*types.Transaction, error) {
	return _StakePool.Contract.Undelegate(&_StakePool.TransactOpts, _receivingToken, _undelegateAmount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _receivingToken, uint256 _undelegateAmount) returns(uint256)
func (_StakePool *StakePoolTransactorSession) Undelegate(_receivingToken common.Address, _undelegateAmount *big.Int) (*types.Transaction, error) {
	return _StakePool.Contract.Undelegate(&_StakePool.TransactOpts, _receivingToken, _undelegateAmount)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakePool *StakePoolTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakePool.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakePool *StakePoolSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakePool.Contract.UpgradeToAndCall(&_StakePool.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakePool *StakePoolTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakePool.Contract.UpgradeToAndCall(&_StakePool.TransactOpts, newImplementation, data)
}

// WithdrawForStaker is a paid mutator transaction binding the contract method 0x6d008623.
//
// Solidity: function withdrawForStaker(address _receivingToken, address _staker, uint256 _amount) returns()
func (_StakePool *StakePoolTransactor) WithdrawForStaker(opts *bind.TransactOpts, _receivingToken common.Address, _staker common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakePool.contract.Transact(opts, "withdrawForStaker", _receivingToken, _staker, _amount)
}

// WithdrawForStaker is a paid mutator transaction binding the contract method 0x6d008623.
//
// Solidity: function withdrawForStaker(address _receivingToken, address _staker, uint256 _amount) returns()
func (_StakePool *StakePoolSession) WithdrawForStaker(_receivingToken common.Address, _staker common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakePool.Contract.WithdrawForStaker(&_StakePool.TransactOpts, _receivingToken, _staker, _amount)
}

// WithdrawForStaker is a paid mutator transaction binding the contract method 0x6d008623.
//
// Solidity: function withdrawForStaker(address _receivingToken, address _staker, uint256 _amount) returns()
func (_StakePool *StakePoolTransactorSession) WithdrawForStaker(_receivingToken common.Address, _staker common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakePool.Contract.WithdrawForStaker(&_StakePool.TransactOpts, _receivingToken, _staker, _amount)
}

// StakePoolDelegateIterator is returned from FilterDelegate and is used to iterate over the raw logs and unpacked data for Delegate events raised by the StakePool contract.
type StakePoolDelegateIterator struct {
	Event *StakePoolDelegate // Event containing the contract specifics and raw log

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
func (it *StakePoolDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolDelegate)
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
		it.Event = new(StakePoolDelegate)
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
func (it *StakePoolDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolDelegate represents a Delegate event raised by the StakePool contract.
type StakePoolDelegate struct {
	Pool       common.Address
	Stablecoin common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDelegate is a free log retrieval operation binding the contract event 0x510b11bb3f3c799b11307c01ab7db0d335683ef5b2da98f7697de744f465eacc.
//
// Solidity: event Delegate(address pool, address stablecoin, uint256 amount)
func (_StakePool *StakePoolFilterer) FilterDelegate(opts *bind.FilterOpts) (*StakePoolDelegateIterator, error) {

	logs, sub, err := _StakePool.contract.FilterLogs(opts, "Delegate")
	if err != nil {
		return nil, err
	}
	return &StakePoolDelegateIterator{contract: _StakePool.contract, event: "Delegate", logs: logs, sub: sub}, nil
}

// WatchDelegate is a free log subscription operation binding the contract event 0x510b11bb3f3c799b11307c01ab7db0d335683ef5b2da98f7697de744f465eacc.
//
// Solidity: event Delegate(address pool, address stablecoin, uint256 amount)
func (_StakePool *StakePoolFilterer) WatchDelegate(opts *bind.WatchOpts, sink chan<- *StakePoolDelegate) (event.Subscription, error) {

	logs, sub, err := _StakePool.contract.WatchLogs(opts, "Delegate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolDelegate)
				if err := _StakePool.contract.UnpackLog(event, "Delegate", log); err != nil {
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

// ParseDelegate is a log parse operation binding the contract event 0x510b11bb3f3c799b11307c01ab7db0d335683ef5b2da98f7697de744f465eacc.
//
// Solidity: event Delegate(address pool, address stablecoin, uint256 amount)
func (_StakePool *StakePoolFilterer) ParseDelegate(log types.Log) (*StakePoolDelegate, error) {
	event := new(StakePoolDelegate)
	if err := _StakePool.contract.UnpackLog(event, "Delegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakePoolInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StakePool contract.
type StakePoolInitializedIterator struct {
	Event *StakePoolInitialized // Event containing the contract specifics and raw log

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
func (it *StakePoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolInitialized)
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
		it.Event = new(StakePoolInitialized)
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
func (it *StakePoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolInitialized represents a Initialized event raised by the StakePool contract.
type StakePoolInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StakePool *StakePoolFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakePoolInitializedIterator, error) {

	logs, sub, err := _StakePool.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakePoolInitializedIterator{contract: _StakePool.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StakePool *StakePoolFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakePoolInitialized) (event.Subscription, error) {

	logs, sub, err := _StakePool.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolInitialized)
				if err := _StakePool.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StakePool *StakePoolFilterer) ParseInitialized(log types.Log) (*StakePoolInitialized, error) {
	event := new(StakePoolInitialized)
	if err := _StakePool.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakePoolMissingUnbondingFeeIterator is returned from FilterMissingUnbondingFee and is used to iterate over the raw logs and unpacked data for MissingUnbondingFee events raised by the StakePool contract.
type StakePoolMissingUnbondingFeeIterator struct {
	Event *StakePoolMissingUnbondingFee // Event containing the contract specifics and raw log

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
func (it *StakePoolMissingUnbondingFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolMissingUnbondingFee)
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
		it.Event = new(StakePoolMissingUnbondingFee)
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
func (it *StakePoolMissingUnbondingFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolMissingUnbondingFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolMissingUnbondingFee represents a MissingUnbondingFee event raised by the StakePool contract.
type StakePoolMissingUnbondingFee struct {
	Stablecoin common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMissingUnbondingFee is a free log retrieval operation binding the contract event 0xbea589036b1df2135726bf5c0a4dc245ef6bf8cbd6af638dfc2ccf3321f38ed1.
//
// Solidity: event MissingUnbondingFee(address stablecoin, uint256 amount)
func (_StakePool *StakePoolFilterer) FilterMissingUnbondingFee(opts *bind.FilterOpts) (*StakePoolMissingUnbondingFeeIterator, error) {

	logs, sub, err := _StakePool.contract.FilterLogs(opts, "MissingUnbondingFee")
	if err != nil {
		return nil, err
	}
	return &StakePoolMissingUnbondingFeeIterator{contract: _StakePool.contract, event: "MissingUnbondingFee", logs: logs, sub: sub}, nil
}

// WatchMissingUnbondingFee is a free log subscription operation binding the contract event 0xbea589036b1df2135726bf5c0a4dc245ef6bf8cbd6af638dfc2ccf3321f38ed1.
//
// Solidity: event MissingUnbondingFee(address stablecoin, uint256 amount)
func (_StakePool *StakePoolFilterer) WatchMissingUnbondingFee(opts *bind.WatchOpts, sink chan<- *StakePoolMissingUnbondingFee) (event.Subscription, error) {

	logs, sub, err := _StakePool.contract.WatchLogs(opts, "MissingUnbondingFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolMissingUnbondingFee)
				if err := _StakePool.contract.UnpackLog(event, "MissingUnbondingFee", log); err != nil {
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

// ParseMissingUnbondingFee is a log parse operation binding the contract event 0xbea589036b1df2135726bf5c0a4dc245ef6bf8cbd6af638dfc2ccf3321f38ed1.
//
// Solidity: event MissingUnbondingFee(address stablecoin, uint256 amount)
func (_StakePool *StakePoolFilterer) ParseMissingUnbondingFee(log types.Log) (*StakePoolMissingUnbondingFee, error) {
	event := new(StakePoolMissingUnbondingFee)
	if err := _StakePool.contract.UnpackLog(event, "MissingUnbondingFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakePoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakePool contract.
type StakePoolOwnershipTransferredIterator struct {
	Event *StakePoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakePoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolOwnershipTransferred)
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
		it.Event = new(StakePoolOwnershipTransferred)
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
func (it *StakePoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolOwnershipTransferred represents a OwnershipTransferred event raised by the StakePool contract.
type StakePoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakePool *StakePoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakePoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakePool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakePoolOwnershipTransferredIterator{contract: _StakePool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakePool *StakePoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakePoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakePool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolOwnershipTransferred)
				if err := _StakePool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakePool *StakePoolFilterer) ParseOwnershipTransferred(log types.Log) (*StakePoolOwnershipTransferred, error) {
	event := new(StakePoolOwnershipTransferred)
	if err := _StakePool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakePoolPayMissingUnbondingFeeIterator is returned from FilterPayMissingUnbondingFee and is used to iterate over the raw logs and unpacked data for PayMissingUnbondingFee events raised by the StakePool contract.
type StakePoolPayMissingUnbondingFeeIterator struct {
	Event *StakePoolPayMissingUnbondingFee // Event containing the contract specifics and raw log

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
func (it *StakePoolPayMissingUnbondingFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolPayMissingUnbondingFee)
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
		it.Event = new(StakePoolPayMissingUnbondingFee)
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
func (it *StakePoolPayMissingUnbondingFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolPayMissingUnbondingFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolPayMissingUnbondingFee represents a PayMissingUnbondingFee event raised by the StakePool contract.
type StakePoolPayMissingUnbondingFee struct {
	Payer      common.Address
	Stablecoin common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPayMissingUnbondingFee is a free log retrieval operation binding the contract event 0x5b5eacdccf2a7feab7735a38c3766084802e999d6849f94928b86af877439b38.
//
// Solidity: event PayMissingUnbondingFee(address payer, address stablecoin, uint256 amount)
func (_StakePool *StakePoolFilterer) FilterPayMissingUnbondingFee(opts *bind.FilterOpts) (*StakePoolPayMissingUnbondingFeeIterator, error) {

	logs, sub, err := _StakePool.contract.FilterLogs(opts, "PayMissingUnbondingFee")
	if err != nil {
		return nil, err
	}
	return &StakePoolPayMissingUnbondingFeeIterator{contract: _StakePool.contract, event: "PayMissingUnbondingFee", logs: logs, sub: sub}, nil
}

// WatchPayMissingUnbondingFee is a free log subscription operation binding the contract event 0x5b5eacdccf2a7feab7735a38c3766084802e999d6849f94928b86af877439b38.
//
// Solidity: event PayMissingUnbondingFee(address payer, address stablecoin, uint256 amount)
func (_StakePool *StakePoolFilterer) WatchPayMissingUnbondingFee(opts *bind.WatchOpts, sink chan<- *StakePoolPayMissingUnbondingFee) (event.Subscription, error) {

	logs, sub, err := _StakePool.contract.WatchLogs(opts, "PayMissingUnbondingFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolPayMissingUnbondingFee)
				if err := _StakePool.contract.UnpackLog(event, "PayMissingUnbondingFee", log); err != nil {
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

// ParsePayMissingUnbondingFee is a log parse operation binding the contract event 0x5b5eacdccf2a7feab7735a38c3766084802e999d6849f94928b86af877439b38.
//
// Solidity: event PayMissingUnbondingFee(address payer, address stablecoin, uint256 amount)
func (_StakePool *StakePoolFilterer) ParsePayMissingUnbondingFee(log types.Log) (*StakePoolPayMissingUnbondingFee, error) {
	event := new(StakePoolPayMissingUnbondingFee)
	if err := _StakePool.contract.UnpackLog(event, "PayMissingUnbondingFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakePoolUndelegateIterator is returned from FilterUndelegate and is used to iterate over the raw logs and unpacked data for Undelegate events raised by the StakePool contract.
type StakePoolUndelegateIterator struct {
	Event *StakePoolUndelegate // Event containing the contract specifics and raw log

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
func (it *StakePoolUndelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolUndelegate)
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
		it.Event = new(StakePoolUndelegate)
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
func (it *StakePoolUndelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolUndelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolUndelegate represents a Undelegate event raised by the StakePool contract.
type StakePoolUndelegate struct {
	Pool     common.Address
	RwaToken common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUndelegate is a free log retrieval operation binding the contract event 0xbda8c0e95802a0e6788c3e9027292382d5a41b86556015f846b03a9874b2b827.
//
// Solidity: event Undelegate(address pool, address rwaToken, uint256 amount)
func (_StakePool *StakePoolFilterer) FilterUndelegate(opts *bind.FilterOpts) (*StakePoolUndelegateIterator, error) {

	logs, sub, err := _StakePool.contract.FilterLogs(opts, "Undelegate")
	if err != nil {
		return nil, err
	}
	return &StakePoolUndelegateIterator{contract: _StakePool.contract, event: "Undelegate", logs: logs, sub: sub}, nil
}

// WatchUndelegate is a free log subscription operation binding the contract event 0xbda8c0e95802a0e6788c3e9027292382d5a41b86556015f846b03a9874b2b827.
//
// Solidity: event Undelegate(address pool, address rwaToken, uint256 amount)
func (_StakePool *StakePoolFilterer) WatchUndelegate(opts *bind.WatchOpts, sink chan<- *StakePoolUndelegate) (event.Subscription, error) {

	logs, sub, err := _StakePool.contract.WatchLogs(opts, "Undelegate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolUndelegate)
				if err := _StakePool.contract.UnpackLog(event, "Undelegate", log); err != nil {
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

// ParseUndelegate is a log parse operation binding the contract event 0xbda8c0e95802a0e6788c3e9027292382d5a41b86556015f846b03a9874b2b827.
//
// Solidity: event Undelegate(address pool, address rwaToken, uint256 amount)
func (_StakePool *StakePoolFilterer) ParseUndelegate(log types.Log) (*StakePoolUndelegate, error) {
	event := new(StakePoolUndelegate)
	if err := _StakePool.contract.UnpackLog(event, "Undelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakePoolUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the StakePool contract.
type StakePoolUpgradedIterator struct {
	Event *StakePoolUpgraded // Event containing the contract specifics and raw log

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
func (it *StakePoolUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolUpgraded)
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
		it.Event = new(StakePoolUpgraded)
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
func (it *StakePoolUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolUpgraded represents a Upgraded event raised by the StakePool contract.
type StakePoolUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StakePool *StakePoolFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StakePoolUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StakePool.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StakePoolUpgradedIterator{contract: _StakePool.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StakePool *StakePoolFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StakePoolUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StakePool.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolUpgraded)
				if err := _StakePool.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StakePool *StakePoolFilterer) ParseUpgraded(log types.Log) (*StakePoolUpgraded, error) {
	event := new(StakePoolUpgraded)
	if err := _StakePool.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakePoolWithdrawForStakerIterator is returned from FilterWithdrawForStaker and is used to iterate over the raw logs and unpacked data for WithdrawForStaker events raised by the StakePool contract.
type StakePoolWithdrawForStakerIterator struct {
	Event *StakePoolWithdrawForStaker // Event containing the contract specifics and raw log

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
func (it *StakePoolWithdrawForStakerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolWithdrawForStaker)
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
		it.Event = new(StakePoolWithdrawForStaker)
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
func (it *StakePoolWithdrawForStakerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolWithdrawForStakerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolWithdrawForStaker represents a WithdrawForStaker event raised by the StakePool contract.
type StakePoolWithdrawForStaker struct {
	Staker         common.Address
	ReceivingToken common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawForStaker is a free log retrieval operation binding the contract event 0xa3c271a114cd6bf8f9fe089a01c4e676ba9dfa0e89a31056fa4a6282041de599.
//
// Solidity: event WithdrawForStaker(address staker, address receivingToken, uint256 amount)
func (_StakePool *StakePoolFilterer) FilterWithdrawForStaker(opts *bind.FilterOpts) (*StakePoolWithdrawForStakerIterator, error) {

	logs, sub, err := _StakePool.contract.FilterLogs(opts, "WithdrawForStaker")
	if err != nil {
		return nil, err
	}
	return &StakePoolWithdrawForStakerIterator{contract: _StakePool.contract, event: "WithdrawForStaker", logs: logs, sub: sub}, nil
}

// WatchWithdrawForStaker is a free log subscription operation binding the contract event 0xa3c271a114cd6bf8f9fe089a01c4e676ba9dfa0e89a31056fa4a6282041de599.
//
// Solidity: event WithdrawForStaker(address staker, address receivingToken, uint256 amount)
func (_StakePool *StakePoolFilterer) WatchWithdrawForStaker(opts *bind.WatchOpts, sink chan<- *StakePoolWithdrawForStaker) (event.Subscription, error) {

	logs, sub, err := _StakePool.contract.WatchLogs(opts, "WithdrawForStaker")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolWithdrawForStaker)
				if err := _StakePool.contract.UnpackLog(event, "WithdrawForStaker", log); err != nil {
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

// ParseWithdrawForStaker is a log parse operation binding the contract event 0xa3c271a114cd6bf8f9fe089a01c4e676ba9dfa0e89a31056fa4a6282041de599.
//
// Solidity: event WithdrawForStaker(address staker, address receivingToken, uint256 amount)
func (_StakePool *StakePoolFilterer) ParseWithdrawForStaker(log types.Log) (*StakePoolWithdrawForStaker, error) {
	event := new(StakePoolWithdrawForStaker)
	if err := _StakePool.contract.UnpackLog(event, "WithdrawForStaker", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
