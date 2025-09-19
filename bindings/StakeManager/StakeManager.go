// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stake_manager

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

// StakeManagerMetaData contains all meta data concerning the StakeManager contract.
var StakeManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ERA_SECONDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_ERA_SECONDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_PROTOCOL_FEE_COMMISSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_RATE_CHANGE_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_UNBONDING_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_VALIDATORS_LEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_ERA_SECONDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UNSTAKE_TIMES_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addStablecoin\",\"inputs\":[{\"name\":\"_stablecoin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addStakePool\",\"inputs\":[{\"name\":\"_poolAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"currentEra\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eraOffset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eraRate\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eraSeconds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"factoryAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"factoryFeeCommission\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBondedPools\",\"inputs\":[],\"outputs\":[{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolInfoByStablecoin\",\"inputs\":[{\"name\":\"_poolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_stablecoin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"fullfilledWithdrawalAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bond\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unbond\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStablecoins\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnstakeIndexListOf\",\"inputs\":[{\"name\":\"_staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"unstakeIndexList\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_lsdToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_stablecoins\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_factoryAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestEra\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lsdToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minStakeAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"newEra\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nextUnstakeIndex\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolInfoOf\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"era\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"active\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"protocolFeeCommission\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rateChangeLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rmStablecoin\",\"inputs\":[{\"name\":\"_stablecoin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEraParams\",\"inputs\":[{\"name\":\"_eraSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_eraOffset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFactoryFeeCommission\",\"inputs\":[{\"name\":\"_factoryFeeCommission\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinStakeAmount\",\"inputs\":[{\"name\":\"_minStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProtocolFeeCommission\",\"inputs\":[{\"name\":\"_protocolFeeCommission\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateChangeLimit\",\"inputs\":[{\"name\":\"_rateChangeLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUnbondingDuration\",\"inputs\":[{\"name\":\"_unbondingDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"_stablecoin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeWithPool\",\"inputs\":[{\"name\":\"_poolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_stablecoin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalProtocolFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unbondingDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unstake\",\"inputs\":[{\"name\":\"_stablecoin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_lsdTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstakeAtIndex\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"era\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stablecoin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetFullfilledAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unstakeWithPool\",\"inputs\":[{\"name\":\"_poolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_stablecoin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_lsdTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawProtocolFee\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawWithPool\",\"inputs\":[{\"name\":\"_poolAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ExecuteNewEra\",\"inputs\":[{\"name\":\"era\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewReward\",\"inputs\":[{\"name\":\"poolAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newReward\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetUnbondingDuration\",\"inputs\":[{\"name\":\"unbondingDuration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Stake\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"poolAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"stablecoin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"lsdTokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unstake\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"poolAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"stablecoin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"lsdTokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"unstakeIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"poolAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"unstakeIndexList\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyWithdrawed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CommissionRateInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DelegateNotEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EraNotMatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedToCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GreaterThanMaxEraSeconds\",\"inputs\":[{\"name\":\"eraSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GreaterThanMaxProtocolFeeCommission\",\"inputs\":[{\"name\":\"protocolFeeCommission\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GreaterThanMaxRateChangeLimit\",\"inputs\":[{\"name\":\"rateChangeLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GreaterThanMaxUnbondingDuration\",\"inputs\":[{\"name\":\"unbondingDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LessThanMinEraSeconds\",\"inputs\":[{\"name\":\"eraSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotAuthorizedLsdToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughAmountToUndelegate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughStakeAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolExist\",\"inputs\":[{\"name\":\"poolAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"PoolNotEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolNotExist\",\"inputs\":[{\"name\":\"poolAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RateChangeExceedLimit\",\"inputs\":[{\"name\":\"oldRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StablecoinDuplicated\",\"inputs\":[{\"name\":\"stablecoin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StablecoinNotExist\",\"inputs\":[{\"name\":\"stablecoin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"UnstakeTimesExceedLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorDuplicated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorsEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorsLenExceedLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongEraParameters\",\"inputs\":[{\"name\":\"eraSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"eraOffset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroRedelegateAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroUnbondingDuration\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroUnstakeAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroWithdrawAmount\",\"inputs\":[]}]",
}

// StakeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeManagerMetaData.ABI instead.
var StakeManagerABI = StakeManagerMetaData.ABI

// StakeManager is an auto generated Go binding around an Ethereum contract.
type StakeManager struct {
	StakeManagerCaller     // Read-only binding to the contract
	StakeManagerTransactor // Write-only binding to the contract
	StakeManagerFilterer   // Log filterer for contract events
}

// StakeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeManagerSession struct {
	Contract     *StakeManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeManagerCallerSession struct {
	Contract *StakeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StakeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeManagerTransactorSession struct {
	Contract     *StakeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StakeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeManagerRaw struct {
	Contract *StakeManager // Generic contract binding to access the raw methods on
}

// StakeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeManagerCallerRaw struct {
	Contract *StakeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeManagerTransactorRaw struct {
	Contract *StakeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakeManager creates a new instance of StakeManager, bound to a specific deployed contract.
func NewStakeManager(address common.Address, backend bind.ContractBackend) (*StakeManager, error) {
	contract, err := bindStakeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakeManager{StakeManagerCaller: StakeManagerCaller{contract: contract}, StakeManagerTransactor: StakeManagerTransactor{contract: contract}, StakeManagerFilterer: StakeManagerFilterer{contract: contract}}, nil
}

// NewStakeManagerCaller creates a new read-only instance of StakeManager, bound to a specific deployed contract.
func NewStakeManagerCaller(address common.Address, caller bind.ContractCaller) (*StakeManagerCaller, error) {
	contract, err := bindStakeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeManagerCaller{contract: contract}, nil
}

// NewStakeManagerTransactor creates a new write-only instance of StakeManager, bound to a specific deployed contract.
func NewStakeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeManagerTransactor, error) {
	contract, err := bindStakeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeManagerTransactor{contract: contract}, nil
}

// NewStakeManagerFilterer creates a new log filterer instance of StakeManager, bound to a specific deployed contract.
func NewStakeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeManagerFilterer, error) {
	contract, err := bindStakeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeManagerFilterer{contract: contract}, nil
}

// bindStakeManager binds a generic wrapper to an already deployed contract.
func bindStakeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakeManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeManager *StakeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeManager.Contract.StakeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeManager *StakeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeManager.Contract.StakeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeManager *StakeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeManager.Contract.StakeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeManager *StakeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeManager *StakeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeManager *StakeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTERASECONDS is a free data retrieval call binding the contract method 0xf4bf0811.
//
// Solidity: function DEFAULT_ERA_SECONDS() view returns(uint256)
func (_StakeManager *StakeManagerCaller) DEFAULTERASECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "DEFAULT_ERA_SECONDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTERASECONDS is a free data retrieval call binding the contract method 0xf4bf0811.
//
// Solidity: function DEFAULT_ERA_SECONDS() view returns(uint256)
func (_StakeManager *StakeManagerSession) DEFAULTERASECONDS() (*big.Int, error) {
	return _StakeManager.Contract.DEFAULTERASECONDS(&_StakeManager.CallOpts)
}

// DEFAULTERASECONDS is a free data retrieval call binding the contract method 0xf4bf0811.
//
// Solidity: function DEFAULT_ERA_SECONDS() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) DEFAULTERASECONDS() (*big.Int, error) {
	return _StakeManager.Contract.DEFAULTERASECONDS(&_StakeManager.CallOpts)
}

// MAXERASECONDS is a free data retrieval call binding the contract method 0xf8d9867a.
//
// Solidity: function MAX_ERA_SECONDS() view returns(uint256)
func (_StakeManager *StakeManagerCaller) MAXERASECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "MAX_ERA_SECONDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXERASECONDS is a free data retrieval call binding the contract method 0xf8d9867a.
//
// Solidity: function MAX_ERA_SECONDS() view returns(uint256)
func (_StakeManager *StakeManagerSession) MAXERASECONDS() (*big.Int, error) {
	return _StakeManager.Contract.MAXERASECONDS(&_StakeManager.CallOpts)
}

// MAXERASECONDS is a free data retrieval call binding the contract method 0xf8d9867a.
//
// Solidity: function MAX_ERA_SECONDS() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) MAXERASECONDS() (*big.Int, error) {
	return _StakeManager.Contract.MAXERASECONDS(&_StakeManager.CallOpts)
}

// MAXPROTOCOLFEECOMMISSION is a free data retrieval call binding the contract method 0xa73a7277.
//
// Solidity: function MAX_PROTOCOL_FEE_COMMISSION() view returns(uint256)
func (_StakeManager *StakeManagerCaller) MAXPROTOCOLFEECOMMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "MAX_PROTOCOL_FEE_COMMISSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPROTOCOLFEECOMMISSION is a free data retrieval call binding the contract method 0xa73a7277.
//
// Solidity: function MAX_PROTOCOL_FEE_COMMISSION() view returns(uint256)
func (_StakeManager *StakeManagerSession) MAXPROTOCOLFEECOMMISSION() (*big.Int, error) {
	return _StakeManager.Contract.MAXPROTOCOLFEECOMMISSION(&_StakeManager.CallOpts)
}

// MAXPROTOCOLFEECOMMISSION is a free data retrieval call binding the contract method 0xa73a7277.
//
// Solidity: function MAX_PROTOCOL_FEE_COMMISSION() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) MAXPROTOCOLFEECOMMISSION() (*big.Int, error) {
	return _StakeManager.Contract.MAXPROTOCOLFEECOMMISSION(&_StakeManager.CallOpts)
}

// MAXRATECHANGELIMIT is a free data retrieval call binding the contract method 0x0d27b7eb.
//
// Solidity: function MAX_RATE_CHANGE_LIMIT() view returns(uint256)
func (_StakeManager *StakeManagerCaller) MAXRATECHANGELIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "MAX_RATE_CHANGE_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXRATECHANGELIMIT is a free data retrieval call binding the contract method 0x0d27b7eb.
//
// Solidity: function MAX_RATE_CHANGE_LIMIT() view returns(uint256)
func (_StakeManager *StakeManagerSession) MAXRATECHANGELIMIT() (*big.Int, error) {
	return _StakeManager.Contract.MAXRATECHANGELIMIT(&_StakeManager.CallOpts)
}

// MAXRATECHANGELIMIT is a free data retrieval call binding the contract method 0x0d27b7eb.
//
// Solidity: function MAX_RATE_CHANGE_LIMIT() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) MAXRATECHANGELIMIT() (*big.Int, error) {
	return _StakeManager.Contract.MAXRATECHANGELIMIT(&_StakeManager.CallOpts)
}

// MAXUNBONDINGDURATION is a free data retrieval call binding the contract method 0x9a4dac92.
//
// Solidity: function MAX_UNBONDING_DURATION() view returns(uint256)
func (_StakeManager *StakeManagerCaller) MAXUNBONDINGDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "MAX_UNBONDING_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXUNBONDINGDURATION is a free data retrieval call binding the contract method 0x9a4dac92.
//
// Solidity: function MAX_UNBONDING_DURATION() view returns(uint256)
func (_StakeManager *StakeManagerSession) MAXUNBONDINGDURATION() (*big.Int, error) {
	return _StakeManager.Contract.MAXUNBONDINGDURATION(&_StakeManager.CallOpts)
}

// MAXUNBONDINGDURATION is a free data retrieval call binding the contract method 0x9a4dac92.
//
// Solidity: function MAX_UNBONDING_DURATION() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) MAXUNBONDINGDURATION() (*big.Int, error) {
	return _StakeManager.Contract.MAXUNBONDINGDURATION(&_StakeManager.CallOpts)
}

// MAXVALIDATORSLEN is a free data retrieval call binding the contract method 0x0ea6f80a.
//
// Solidity: function MAX_VALIDATORS_LEN() view returns(uint256)
func (_StakeManager *StakeManagerCaller) MAXVALIDATORSLEN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "MAX_VALIDATORS_LEN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVALIDATORSLEN is a free data retrieval call binding the contract method 0x0ea6f80a.
//
// Solidity: function MAX_VALIDATORS_LEN() view returns(uint256)
func (_StakeManager *StakeManagerSession) MAXVALIDATORSLEN() (*big.Int, error) {
	return _StakeManager.Contract.MAXVALIDATORSLEN(&_StakeManager.CallOpts)
}

// MAXVALIDATORSLEN is a free data retrieval call binding the contract method 0x0ea6f80a.
//
// Solidity: function MAX_VALIDATORS_LEN() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) MAXVALIDATORSLEN() (*big.Int, error) {
	return _StakeManager.Contract.MAXVALIDATORSLEN(&_StakeManager.CallOpts)
}

// MINERASECONDS is a free data retrieval call binding the contract method 0xbd4feb87.
//
// Solidity: function MIN_ERA_SECONDS() view returns(uint256)
func (_StakeManager *StakeManagerCaller) MINERASECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "MIN_ERA_SECONDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINERASECONDS is a free data retrieval call binding the contract method 0xbd4feb87.
//
// Solidity: function MIN_ERA_SECONDS() view returns(uint256)
func (_StakeManager *StakeManagerSession) MINERASECONDS() (*big.Int, error) {
	return _StakeManager.Contract.MINERASECONDS(&_StakeManager.CallOpts)
}

// MINERASECONDS is a free data retrieval call binding the contract method 0xbd4feb87.
//
// Solidity: function MIN_ERA_SECONDS() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) MINERASECONDS() (*big.Int, error) {
	return _StakeManager.Contract.MINERASECONDS(&_StakeManager.CallOpts)
}

// UNSTAKETIMESLIMIT is a free data retrieval call binding the contract method 0xe34719d9.
//
// Solidity: function UNSTAKE_TIMES_LIMIT() view returns(uint256)
func (_StakeManager *StakeManagerCaller) UNSTAKETIMESLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "UNSTAKE_TIMES_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNSTAKETIMESLIMIT is a free data retrieval call binding the contract method 0xe34719d9.
//
// Solidity: function UNSTAKE_TIMES_LIMIT() view returns(uint256)
func (_StakeManager *StakeManagerSession) UNSTAKETIMESLIMIT() (*big.Int, error) {
	return _StakeManager.Contract.UNSTAKETIMESLIMIT(&_StakeManager.CallOpts)
}

// UNSTAKETIMESLIMIT is a free data retrieval call binding the contract method 0xe34719d9.
//
// Solidity: function UNSTAKE_TIMES_LIMIT() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) UNSTAKETIMESLIMIT() (*big.Int, error) {
	return _StakeManager.Contract.UNSTAKETIMESLIMIT(&_StakeManager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_StakeManager *StakeManagerCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_StakeManager *StakeManagerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _StakeManager.Contract.UPGRADEINTERFACEVERSION(&_StakeManager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_StakeManager *StakeManagerCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _StakeManager.Contract.UPGRADEINTERFACEVERSION(&_StakeManager.CallOpts)
}

// CurrentEra is a free data retrieval call binding the contract method 0x973628f6.
//
// Solidity: function currentEra() view returns(uint256)
func (_StakeManager *StakeManagerCaller) CurrentEra(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "currentEra")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEra is a free data retrieval call binding the contract method 0x973628f6.
//
// Solidity: function currentEra() view returns(uint256)
func (_StakeManager *StakeManagerSession) CurrentEra() (*big.Int, error) {
	return _StakeManager.Contract.CurrentEra(&_StakeManager.CallOpts)
}

// CurrentEra is a free data retrieval call binding the contract method 0x973628f6.
//
// Solidity: function currentEra() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) CurrentEra() (*big.Int, error) {
	return _StakeManager.Contract.CurrentEra(&_StakeManager.CallOpts)
}

// EraOffset is a free data retrieval call binding the contract method 0xc8c20263.
//
// Solidity: function eraOffset() view returns(uint256)
func (_StakeManager *StakeManagerCaller) EraOffset(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "eraOffset")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EraOffset is a free data retrieval call binding the contract method 0xc8c20263.
//
// Solidity: function eraOffset() view returns(uint256)
func (_StakeManager *StakeManagerSession) EraOffset() (*big.Int, error) {
	return _StakeManager.Contract.EraOffset(&_StakeManager.CallOpts)
}

// EraOffset is a free data retrieval call binding the contract method 0xc8c20263.
//
// Solidity: function eraOffset() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) EraOffset() (*big.Int, error) {
	return _StakeManager.Contract.EraOffset(&_StakeManager.CallOpts)
}

// EraRate is a free data retrieval call binding the contract method 0xeb8ad76e.
//
// Solidity: function eraRate(uint256 ) view returns(uint256)
func (_StakeManager *StakeManagerCaller) EraRate(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "eraRate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EraRate is a free data retrieval call binding the contract method 0xeb8ad76e.
//
// Solidity: function eraRate(uint256 ) view returns(uint256)
func (_StakeManager *StakeManagerSession) EraRate(arg0 *big.Int) (*big.Int, error) {
	return _StakeManager.Contract.EraRate(&_StakeManager.CallOpts, arg0)
}

// EraRate is a free data retrieval call binding the contract method 0xeb8ad76e.
//
// Solidity: function eraRate(uint256 ) view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) EraRate(arg0 *big.Int) (*big.Int, error) {
	return _StakeManager.Contract.EraRate(&_StakeManager.CallOpts, arg0)
}

// EraSeconds is a free data retrieval call binding the contract method 0xe81f1553.
//
// Solidity: function eraSeconds() view returns(uint256)
func (_StakeManager *StakeManagerCaller) EraSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "eraSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EraSeconds is a free data retrieval call binding the contract method 0xe81f1553.
//
// Solidity: function eraSeconds() view returns(uint256)
func (_StakeManager *StakeManagerSession) EraSeconds() (*big.Int, error) {
	return _StakeManager.Contract.EraSeconds(&_StakeManager.CallOpts)
}

// EraSeconds is a free data retrieval call binding the contract method 0xe81f1553.
//
// Solidity: function eraSeconds() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) EraSeconds() (*big.Int, error) {
	return _StakeManager.Contract.EraSeconds(&_StakeManager.CallOpts)
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() view returns(address)
func (_StakeManager *StakeManagerCaller) FactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "factoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() view returns(address)
func (_StakeManager *StakeManagerSession) FactoryAddress() (common.Address, error) {
	return _StakeManager.Contract.FactoryAddress(&_StakeManager.CallOpts)
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() view returns(address)
func (_StakeManager *StakeManagerCallerSession) FactoryAddress() (common.Address, error) {
	return _StakeManager.Contract.FactoryAddress(&_StakeManager.CallOpts)
}

// FactoryFeeCommission is a free data retrieval call binding the contract method 0x0e088d37.
//
// Solidity: function factoryFeeCommission() view returns(uint256)
func (_StakeManager *StakeManagerCaller) FactoryFeeCommission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "factoryFeeCommission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FactoryFeeCommission is a free data retrieval call binding the contract method 0x0e088d37.
//
// Solidity: function factoryFeeCommission() view returns(uint256)
func (_StakeManager *StakeManagerSession) FactoryFeeCommission() (*big.Int, error) {
	return _StakeManager.Contract.FactoryFeeCommission(&_StakeManager.CallOpts)
}

// FactoryFeeCommission is a free data retrieval call binding the contract method 0x0e088d37.
//
// Solidity: function factoryFeeCommission() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) FactoryFeeCommission() (*big.Int, error) {
	return _StakeManager.Contract.FactoryFeeCommission(&_StakeManager.CallOpts)
}

// GetBondedPools is a free data retrieval call binding the contract method 0x58af8bf0.
//
// Solidity: function getBondedPools() view returns(address[] pools)
func (_StakeManager *StakeManagerCaller) GetBondedPools(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "getBondedPools")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBondedPools is a free data retrieval call binding the contract method 0x58af8bf0.
//
// Solidity: function getBondedPools() view returns(address[] pools)
func (_StakeManager *StakeManagerSession) GetBondedPools() ([]common.Address, error) {
	return _StakeManager.Contract.GetBondedPools(&_StakeManager.CallOpts)
}

// GetBondedPools is a free data retrieval call binding the contract method 0x58af8bf0.
//
// Solidity: function getBondedPools() view returns(address[] pools)
func (_StakeManager *StakeManagerCallerSession) GetBondedPools() ([]common.Address, error) {
	return _StakeManager.Contract.GetBondedPools(&_StakeManager.CallOpts)
}

// GetPoolInfoByStablecoin is a free data retrieval call binding the contract method 0x57e6fa61.
//
// Solidity: function getPoolInfoByStablecoin(address _poolAddress, address _stablecoin) view returns(uint256 fullfilledWithdrawalAmount, uint256 bond, uint256 unbond)
func (_StakeManager *StakeManagerCaller) GetPoolInfoByStablecoin(opts *bind.CallOpts, _poolAddress common.Address, _stablecoin common.Address) (struct {
	FullfilledWithdrawalAmount *big.Int
	Bond                       *big.Int
	Unbond                     *big.Int
}, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "getPoolInfoByStablecoin", _poolAddress, _stablecoin)

	outstruct := new(struct {
		FullfilledWithdrawalAmount *big.Int
		Bond                       *big.Int
		Unbond                     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FullfilledWithdrawalAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Bond = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Unbond = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPoolInfoByStablecoin is a free data retrieval call binding the contract method 0x57e6fa61.
//
// Solidity: function getPoolInfoByStablecoin(address _poolAddress, address _stablecoin) view returns(uint256 fullfilledWithdrawalAmount, uint256 bond, uint256 unbond)
func (_StakeManager *StakeManagerSession) GetPoolInfoByStablecoin(_poolAddress common.Address, _stablecoin common.Address) (struct {
	FullfilledWithdrawalAmount *big.Int
	Bond                       *big.Int
	Unbond                     *big.Int
}, error) {
	return _StakeManager.Contract.GetPoolInfoByStablecoin(&_StakeManager.CallOpts, _poolAddress, _stablecoin)
}

// GetPoolInfoByStablecoin is a free data retrieval call binding the contract method 0x57e6fa61.
//
// Solidity: function getPoolInfoByStablecoin(address _poolAddress, address _stablecoin) view returns(uint256 fullfilledWithdrawalAmount, uint256 bond, uint256 unbond)
func (_StakeManager *StakeManagerCallerSession) GetPoolInfoByStablecoin(_poolAddress common.Address, _stablecoin common.Address) (struct {
	FullfilledWithdrawalAmount *big.Int
	Bond                       *big.Int
	Unbond                     *big.Int
}, error) {
	return _StakeManager.Contract.GetPoolInfoByStablecoin(&_StakeManager.CallOpts, _poolAddress, _stablecoin)
}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_StakeManager *StakeManagerCaller) GetRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "getRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_StakeManager *StakeManagerSession) GetRate() (*big.Int, error) {
	return _StakeManager.Contract.GetRate(&_StakeManager.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) GetRate() (*big.Int, error) {
	return _StakeManager.Contract.GetRate(&_StakeManager.CallOpts)
}

// GetStablecoins is a free data retrieval call binding the contract method 0xb6020edb.
//
// Solidity: function getStablecoins() view returns(address[])
func (_StakeManager *StakeManagerCaller) GetStablecoins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "getStablecoins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStablecoins is a free data retrieval call binding the contract method 0xb6020edb.
//
// Solidity: function getStablecoins() view returns(address[])
func (_StakeManager *StakeManagerSession) GetStablecoins() ([]common.Address, error) {
	return _StakeManager.Contract.GetStablecoins(&_StakeManager.CallOpts)
}

// GetStablecoins is a free data retrieval call binding the contract method 0xb6020edb.
//
// Solidity: function getStablecoins() view returns(address[])
func (_StakeManager *StakeManagerCallerSession) GetStablecoins() ([]common.Address, error) {
	return _StakeManager.Contract.GetStablecoins(&_StakeManager.CallOpts)
}

// GetUnstakeIndexListOf is a free data retrieval call binding the contract method 0x615acc36.
//
// Solidity: function getUnstakeIndexListOf(address _staker) view returns(uint256[] unstakeIndexList)
func (_StakeManager *StakeManagerCaller) GetUnstakeIndexListOf(opts *bind.CallOpts, _staker common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "getUnstakeIndexListOf", _staker)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUnstakeIndexListOf is a free data retrieval call binding the contract method 0x615acc36.
//
// Solidity: function getUnstakeIndexListOf(address _staker) view returns(uint256[] unstakeIndexList)
func (_StakeManager *StakeManagerSession) GetUnstakeIndexListOf(_staker common.Address) ([]*big.Int, error) {
	return _StakeManager.Contract.GetUnstakeIndexListOf(&_StakeManager.CallOpts, _staker)
}

// GetUnstakeIndexListOf is a free data retrieval call binding the contract method 0x615acc36.
//
// Solidity: function getUnstakeIndexListOf(address _staker) view returns(uint256[] unstakeIndexList)
func (_StakeManager *StakeManagerCallerSession) GetUnstakeIndexListOf(_staker common.Address) ([]*big.Int, error) {
	return _StakeManager.Contract.GetUnstakeIndexListOf(&_StakeManager.CallOpts, _staker)
}

// LatestEra is a free data retrieval call binding the contract method 0x3f6f5f32.
//
// Solidity: function latestEra() view returns(uint256)
func (_StakeManager *StakeManagerCaller) LatestEra(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "latestEra")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestEra is a free data retrieval call binding the contract method 0x3f6f5f32.
//
// Solidity: function latestEra() view returns(uint256)
func (_StakeManager *StakeManagerSession) LatestEra() (*big.Int, error) {
	return _StakeManager.Contract.LatestEra(&_StakeManager.CallOpts)
}

// LatestEra is a free data retrieval call binding the contract method 0x3f6f5f32.
//
// Solidity: function latestEra() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) LatestEra() (*big.Int, error) {
	return _StakeManager.Contract.LatestEra(&_StakeManager.CallOpts)
}

// LsdToken is a free data retrieval call binding the contract method 0xa6645fdd.
//
// Solidity: function lsdToken() view returns(address)
func (_StakeManager *StakeManagerCaller) LsdToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "lsdToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LsdToken is a free data retrieval call binding the contract method 0xa6645fdd.
//
// Solidity: function lsdToken() view returns(address)
func (_StakeManager *StakeManagerSession) LsdToken() (common.Address, error) {
	return _StakeManager.Contract.LsdToken(&_StakeManager.CallOpts)
}

// LsdToken is a free data retrieval call binding the contract method 0xa6645fdd.
//
// Solidity: function lsdToken() view returns(address)
func (_StakeManager *StakeManagerCallerSession) LsdToken() (common.Address, error) {
	return _StakeManager.Contract.LsdToken(&_StakeManager.CallOpts)
}

// MinStakeAmount is a free data retrieval call binding the contract method 0xf1887684.
//
// Solidity: function minStakeAmount() view returns(uint256)
func (_StakeManager *StakeManagerCaller) MinStakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "minStakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakeAmount is a free data retrieval call binding the contract method 0xf1887684.
//
// Solidity: function minStakeAmount() view returns(uint256)
func (_StakeManager *StakeManagerSession) MinStakeAmount() (*big.Int, error) {
	return _StakeManager.Contract.MinStakeAmount(&_StakeManager.CallOpts)
}

// MinStakeAmount is a free data retrieval call binding the contract method 0xf1887684.
//
// Solidity: function minStakeAmount() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) MinStakeAmount() (*big.Int, error) {
	return _StakeManager.Contract.MinStakeAmount(&_StakeManager.CallOpts)
}

// NextUnstakeIndex is a free data retrieval call binding the contract method 0x3bea9ee3.
//
// Solidity: function nextUnstakeIndex() view returns(uint256)
func (_StakeManager *StakeManagerCaller) NextUnstakeIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "nextUnstakeIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextUnstakeIndex is a free data retrieval call binding the contract method 0x3bea9ee3.
//
// Solidity: function nextUnstakeIndex() view returns(uint256)
func (_StakeManager *StakeManagerSession) NextUnstakeIndex() (*big.Int, error) {
	return _StakeManager.Contract.NextUnstakeIndex(&_StakeManager.CallOpts)
}

// NextUnstakeIndex is a free data retrieval call binding the contract method 0x3bea9ee3.
//
// Solidity: function nextUnstakeIndex() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) NextUnstakeIndex() (*big.Int, error) {
	return _StakeManager.Contract.NextUnstakeIndex(&_StakeManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakeManager *StakeManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakeManager *StakeManagerSession) Owner() (common.Address, error) {
	return _StakeManager.Contract.Owner(&_StakeManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakeManager *StakeManagerCallerSession) Owner() (common.Address, error) {
	return _StakeManager.Contract.Owner(&_StakeManager.CallOpts)
}

// PoolInfoOf is a free data retrieval call binding the contract method 0x008b5dd2.
//
// Solidity: function poolInfoOf(address ) view returns(uint256 era, uint256 active)
func (_StakeManager *StakeManagerCaller) PoolInfoOf(opts *bind.CallOpts, arg0 common.Address) (struct {
	Era    *big.Int
	Active *big.Int
}, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "poolInfoOf", arg0)

	outstruct := new(struct {
		Era    *big.Int
		Active *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Era = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Active = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfoOf is a free data retrieval call binding the contract method 0x008b5dd2.
//
// Solidity: function poolInfoOf(address ) view returns(uint256 era, uint256 active)
func (_StakeManager *StakeManagerSession) PoolInfoOf(arg0 common.Address) (struct {
	Era    *big.Int
	Active *big.Int
}, error) {
	return _StakeManager.Contract.PoolInfoOf(&_StakeManager.CallOpts, arg0)
}

// PoolInfoOf is a free data retrieval call binding the contract method 0x008b5dd2.
//
// Solidity: function poolInfoOf(address ) view returns(uint256 era, uint256 active)
func (_StakeManager *StakeManagerCallerSession) PoolInfoOf(arg0 common.Address) (struct {
	Era    *big.Int
	Active *big.Int
}, error) {
	return _StakeManager.Contract.PoolInfoOf(&_StakeManager.CallOpts, arg0)
}

// ProtocolFeeCommission is a free data retrieval call binding the contract method 0x19301c26.
//
// Solidity: function protocolFeeCommission() view returns(uint256)
func (_StakeManager *StakeManagerCaller) ProtocolFeeCommission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "protocolFeeCommission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeeCommission is a free data retrieval call binding the contract method 0x19301c26.
//
// Solidity: function protocolFeeCommission() view returns(uint256)
func (_StakeManager *StakeManagerSession) ProtocolFeeCommission() (*big.Int, error) {
	return _StakeManager.Contract.ProtocolFeeCommission(&_StakeManager.CallOpts)
}

// ProtocolFeeCommission is a free data retrieval call binding the contract method 0x19301c26.
//
// Solidity: function protocolFeeCommission() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) ProtocolFeeCommission() (*big.Int, error) {
	return _StakeManager.Contract.ProtocolFeeCommission(&_StakeManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StakeManager *StakeManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StakeManager *StakeManagerSession) ProxiableUUID() ([32]byte, error) {
	return _StakeManager.Contract.ProxiableUUID(&_StakeManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StakeManager *StakeManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _StakeManager.Contract.ProxiableUUID(&_StakeManager.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_StakeManager *StakeManagerCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_StakeManager *StakeManagerSession) Rate() (*big.Int, error) {
	return _StakeManager.Contract.Rate(&_StakeManager.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) Rate() (*big.Int, error) {
	return _StakeManager.Contract.Rate(&_StakeManager.CallOpts)
}

// RateChangeLimit is a free data retrieval call binding the contract method 0xc0152b71.
//
// Solidity: function rateChangeLimit() view returns(uint256)
func (_StakeManager *StakeManagerCaller) RateChangeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "rateChangeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RateChangeLimit is a free data retrieval call binding the contract method 0xc0152b71.
//
// Solidity: function rateChangeLimit() view returns(uint256)
func (_StakeManager *StakeManagerSession) RateChangeLimit() (*big.Int, error) {
	return _StakeManager.Contract.RateChangeLimit(&_StakeManager.CallOpts)
}

// RateChangeLimit is a free data retrieval call binding the contract method 0xc0152b71.
//
// Solidity: function rateChangeLimit() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) RateChangeLimit() (*big.Int, error) {
	return _StakeManager.Contract.RateChangeLimit(&_StakeManager.CallOpts)
}

// TotalProtocolFee is a free data retrieval call binding the contract method 0x88611f35.
//
// Solidity: function totalProtocolFee() view returns(uint256)
func (_StakeManager *StakeManagerCaller) TotalProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "totalProtocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalProtocolFee is a free data retrieval call binding the contract method 0x88611f35.
//
// Solidity: function totalProtocolFee() view returns(uint256)
func (_StakeManager *StakeManagerSession) TotalProtocolFee() (*big.Int, error) {
	return _StakeManager.Contract.TotalProtocolFee(&_StakeManager.CallOpts)
}

// TotalProtocolFee is a free data retrieval call binding the contract method 0x88611f35.
//
// Solidity: function totalProtocolFee() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) TotalProtocolFee() (*big.Int, error) {
	return _StakeManager.Contract.TotalProtocolFee(&_StakeManager.CallOpts)
}

// UnbondingDuration is a free data retrieval call binding the contract method 0xccf6802a.
//
// Solidity: function unbondingDuration() view returns(uint256)
func (_StakeManager *StakeManagerCaller) UnbondingDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "unbondingDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnbondingDuration is a free data retrieval call binding the contract method 0xccf6802a.
//
// Solidity: function unbondingDuration() view returns(uint256)
func (_StakeManager *StakeManagerSession) UnbondingDuration() (*big.Int, error) {
	return _StakeManager.Contract.UnbondingDuration(&_StakeManager.CallOpts)
}

// UnbondingDuration is a free data retrieval call binding the contract method 0xccf6802a.
//
// Solidity: function unbondingDuration() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) UnbondingDuration() (*big.Int, error) {
	return _StakeManager.Contract.UnbondingDuration(&_StakeManager.CallOpts)
}

// UnstakeAtIndex is a free data retrieval call binding the contract method 0x6e436c6e.
//
// Solidity: function unstakeAtIndex(uint256 ) view returns(uint256 era, address pool, address stablecoin, address receiver, uint256 amount, uint256 targetFullfilledAmount)
func (_StakeManager *StakeManagerCaller) UnstakeAtIndex(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Era                    *big.Int
	Pool                   common.Address
	Stablecoin             common.Address
	Receiver               common.Address
	Amount                 *big.Int
	TargetFullfilledAmount *big.Int
}, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "unstakeAtIndex", arg0)

	outstruct := new(struct {
		Era                    *big.Int
		Pool                   common.Address
		Stablecoin             common.Address
		Receiver               common.Address
		Amount                 *big.Int
		TargetFullfilledAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Era = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Pool = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Stablecoin = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Receiver = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TargetFullfilledAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UnstakeAtIndex is a free data retrieval call binding the contract method 0x6e436c6e.
//
// Solidity: function unstakeAtIndex(uint256 ) view returns(uint256 era, address pool, address stablecoin, address receiver, uint256 amount, uint256 targetFullfilledAmount)
func (_StakeManager *StakeManagerSession) UnstakeAtIndex(arg0 *big.Int) (struct {
	Era                    *big.Int
	Pool                   common.Address
	Stablecoin             common.Address
	Receiver               common.Address
	Amount                 *big.Int
	TargetFullfilledAmount *big.Int
}, error) {
	return _StakeManager.Contract.UnstakeAtIndex(&_StakeManager.CallOpts, arg0)
}

// UnstakeAtIndex is a free data retrieval call binding the contract method 0x6e436c6e.
//
// Solidity: function unstakeAtIndex(uint256 ) view returns(uint256 era, address pool, address stablecoin, address receiver, uint256 amount, uint256 targetFullfilledAmount)
func (_StakeManager *StakeManagerCallerSession) UnstakeAtIndex(arg0 *big.Int) (struct {
	Era                    *big.Int
	Pool                   common.Address
	Stablecoin             common.Address
	Receiver               common.Address
	Amount                 *big.Int
	TargetFullfilledAmount *big.Int
}, error) {
	return _StakeManager.Contract.UnstakeAtIndex(&_StakeManager.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint64)
func (_StakeManager *StakeManagerCaller) Version(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint64)
func (_StakeManager *StakeManagerSession) Version() (uint64, error) {
	return _StakeManager.Contract.Version(&_StakeManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint64)
func (_StakeManager *StakeManagerCallerSession) Version() (uint64, error) {
	return _StakeManager.Contract.Version(&_StakeManager.CallOpts)
}

// AddStablecoin is a paid mutator transaction binding the contract method 0x0418945a.
//
// Solidity: function addStablecoin(address _stablecoin) returns()
func (_StakeManager *StakeManagerTransactor) AddStablecoin(opts *bind.TransactOpts, _stablecoin common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "addStablecoin", _stablecoin)
}

// AddStablecoin is a paid mutator transaction binding the contract method 0x0418945a.
//
// Solidity: function addStablecoin(address _stablecoin) returns()
func (_StakeManager *StakeManagerSession) AddStablecoin(_stablecoin common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.AddStablecoin(&_StakeManager.TransactOpts, _stablecoin)
}

// AddStablecoin is a paid mutator transaction binding the contract method 0x0418945a.
//
// Solidity: function addStablecoin(address _stablecoin) returns()
func (_StakeManager *StakeManagerTransactorSession) AddStablecoin(_stablecoin common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.AddStablecoin(&_StakeManager.TransactOpts, _stablecoin)
}

// AddStakePool is a paid mutator transaction binding the contract method 0x0f772a1d.
//
// Solidity: function addStakePool(address _poolAddress) returns()
func (_StakeManager *StakeManagerTransactor) AddStakePool(opts *bind.TransactOpts, _poolAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "addStakePool", _poolAddress)
}

// AddStakePool is a paid mutator transaction binding the contract method 0x0f772a1d.
//
// Solidity: function addStakePool(address _poolAddress) returns()
func (_StakeManager *StakeManagerSession) AddStakePool(_poolAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.AddStakePool(&_StakeManager.TransactOpts, _poolAddress)
}

// AddStakePool is a paid mutator transaction binding the contract method 0x0f772a1d.
//
// Solidity: function addStakePool(address _poolAddress) returns()
func (_StakeManager *StakeManagerTransactorSession) AddStakePool(_poolAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.AddStakePool(&_StakeManager.TransactOpts, _poolAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xe3b1a2b3.
//
// Solidity: function initialize(address _lsdToken, address _poolAddress, address _owner, address[] _stablecoins, address _factoryAddress) returns()
func (_StakeManager *StakeManagerTransactor) Initialize(opts *bind.TransactOpts, _lsdToken common.Address, _poolAddress common.Address, _owner common.Address, _stablecoins []common.Address, _factoryAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "initialize", _lsdToken, _poolAddress, _owner, _stablecoins, _factoryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xe3b1a2b3.
//
// Solidity: function initialize(address _lsdToken, address _poolAddress, address _owner, address[] _stablecoins, address _factoryAddress) returns()
func (_StakeManager *StakeManagerSession) Initialize(_lsdToken common.Address, _poolAddress common.Address, _owner common.Address, _stablecoins []common.Address, _factoryAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.Initialize(&_StakeManager.TransactOpts, _lsdToken, _poolAddress, _owner, _stablecoins, _factoryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xe3b1a2b3.
//
// Solidity: function initialize(address _lsdToken, address _poolAddress, address _owner, address[] _stablecoins, address _factoryAddress) returns()
func (_StakeManager *StakeManagerTransactorSession) Initialize(_lsdToken common.Address, _poolAddress common.Address, _owner common.Address, _stablecoins []common.Address, _factoryAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.Initialize(&_StakeManager.TransactOpts, _lsdToken, _poolAddress, _owner, _stablecoins, _factoryAddress)
}

// NewEra is a paid mutator transaction binding the contract method 0x7b207727.
//
// Solidity: function newEra() returns()
func (_StakeManager *StakeManagerTransactor) NewEra(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "newEra")
}

// NewEra is a paid mutator transaction binding the contract method 0x7b207727.
//
// Solidity: function newEra() returns()
func (_StakeManager *StakeManagerSession) NewEra() (*types.Transaction, error) {
	return _StakeManager.Contract.NewEra(&_StakeManager.TransactOpts)
}

// NewEra is a paid mutator transaction binding the contract method 0x7b207727.
//
// Solidity: function newEra() returns()
func (_StakeManager *StakeManagerTransactorSession) NewEra() (*types.Transaction, error) {
	return _StakeManager.Contract.NewEra(&_StakeManager.TransactOpts)
}

// RmStablecoin is a paid mutator transaction binding the contract method 0x356772b3.
//
// Solidity: function rmStablecoin(address _stablecoin) returns()
func (_StakeManager *StakeManagerTransactor) RmStablecoin(opts *bind.TransactOpts, _stablecoin common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "rmStablecoin", _stablecoin)
}

// RmStablecoin is a paid mutator transaction binding the contract method 0x356772b3.
//
// Solidity: function rmStablecoin(address _stablecoin) returns()
func (_StakeManager *StakeManagerSession) RmStablecoin(_stablecoin common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.RmStablecoin(&_StakeManager.TransactOpts, _stablecoin)
}

// RmStablecoin is a paid mutator transaction binding the contract method 0x356772b3.
//
// Solidity: function rmStablecoin(address _stablecoin) returns()
func (_StakeManager *StakeManagerTransactorSession) RmStablecoin(_stablecoin common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.RmStablecoin(&_StakeManager.TransactOpts, _stablecoin)
}

// SetEraParams is a paid mutator transaction binding the contract method 0x8b9686a9.
//
// Solidity: function setEraParams(uint256 _eraSeconds, uint256 _eraOffset) returns()
func (_StakeManager *StakeManagerTransactor) SetEraParams(opts *bind.TransactOpts, _eraSeconds *big.Int, _eraOffset *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "setEraParams", _eraSeconds, _eraOffset)
}

// SetEraParams is a paid mutator transaction binding the contract method 0x8b9686a9.
//
// Solidity: function setEraParams(uint256 _eraSeconds, uint256 _eraOffset) returns()
func (_StakeManager *StakeManagerSession) SetEraParams(_eraSeconds *big.Int, _eraOffset *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SetEraParams(&_StakeManager.TransactOpts, _eraSeconds, _eraOffset)
}

// SetEraParams is a paid mutator transaction binding the contract method 0x8b9686a9.
//
// Solidity: function setEraParams(uint256 _eraSeconds, uint256 _eraOffset) returns()
func (_StakeManager *StakeManagerTransactorSession) SetEraParams(_eraSeconds *big.Int, _eraOffset *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SetEraParams(&_StakeManager.TransactOpts, _eraSeconds, _eraOffset)
}

// SetFactoryFeeCommission is a paid mutator transaction binding the contract method 0xc681229d.
//
// Solidity: function setFactoryFeeCommission(uint256 _factoryFeeCommission) returns()
func (_StakeManager *StakeManagerTransactor) SetFactoryFeeCommission(opts *bind.TransactOpts, _factoryFeeCommission *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "setFactoryFeeCommission", _factoryFeeCommission)
}

// SetFactoryFeeCommission is a paid mutator transaction binding the contract method 0xc681229d.
//
// Solidity: function setFactoryFeeCommission(uint256 _factoryFeeCommission) returns()
func (_StakeManager *StakeManagerSession) SetFactoryFeeCommission(_factoryFeeCommission *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SetFactoryFeeCommission(&_StakeManager.TransactOpts, _factoryFeeCommission)
}

// SetFactoryFeeCommission is a paid mutator transaction binding the contract method 0xc681229d.
//
// Solidity: function setFactoryFeeCommission(uint256 _factoryFeeCommission) returns()
func (_StakeManager *StakeManagerTransactorSession) SetFactoryFeeCommission(_factoryFeeCommission *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SetFactoryFeeCommission(&_StakeManager.TransactOpts, _factoryFeeCommission)
}

// SetMinStakeAmount is a paid mutator transaction binding the contract method 0xeb4af045.
//
// Solidity: function setMinStakeAmount(uint256 _minStakeAmount) returns()
func (_StakeManager *StakeManagerTransactor) SetMinStakeAmount(opts *bind.TransactOpts, _minStakeAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "setMinStakeAmount", _minStakeAmount)
}

// SetMinStakeAmount is a paid mutator transaction binding the contract method 0xeb4af045.
//
// Solidity: function setMinStakeAmount(uint256 _minStakeAmount) returns()
func (_StakeManager *StakeManagerSession) SetMinStakeAmount(_minStakeAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SetMinStakeAmount(&_StakeManager.TransactOpts, _minStakeAmount)
}

// SetMinStakeAmount is a paid mutator transaction binding the contract method 0xeb4af045.
//
// Solidity: function setMinStakeAmount(uint256 _minStakeAmount) returns()
func (_StakeManager *StakeManagerTransactorSession) SetMinStakeAmount(_minStakeAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SetMinStakeAmount(&_StakeManager.TransactOpts, _minStakeAmount)
}

// SetProtocolFeeCommission is a paid mutator transaction binding the contract method 0xaf737e60.
//
// Solidity: function setProtocolFeeCommission(uint256 _protocolFeeCommission) returns()
func (_StakeManager *StakeManagerTransactor) SetProtocolFeeCommission(opts *bind.TransactOpts, _protocolFeeCommission *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "setProtocolFeeCommission", _protocolFeeCommission)
}

// SetProtocolFeeCommission is a paid mutator transaction binding the contract method 0xaf737e60.
//
// Solidity: function setProtocolFeeCommission(uint256 _protocolFeeCommission) returns()
func (_StakeManager *StakeManagerSession) SetProtocolFeeCommission(_protocolFeeCommission *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SetProtocolFeeCommission(&_StakeManager.TransactOpts, _protocolFeeCommission)
}

// SetProtocolFeeCommission is a paid mutator transaction binding the contract method 0xaf737e60.
//
// Solidity: function setProtocolFeeCommission(uint256 _protocolFeeCommission) returns()
func (_StakeManager *StakeManagerTransactorSession) SetProtocolFeeCommission(_protocolFeeCommission *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SetProtocolFeeCommission(&_StakeManager.TransactOpts, _protocolFeeCommission)
}

// SetRateChangeLimit is a paid mutator transaction binding the contract method 0x19826e71.
//
// Solidity: function setRateChangeLimit(uint256 _rateChangeLimit) returns()
func (_StakeManager *StakeManagerTransactor) SetRateChangeLimit(opts *bind.TransactOpts, _rateChangeLimit *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "setRateChangeLimit", _rateChangeLimit)
}

// SetRateChangeLimit is a paid mutator transaction binding the contract method 0x19826e71.
//
// Solidity: function setRateChangeLimit(uint256 _rateChangeLimit) returns()
func (_StakeManager *StakeManagerSession) SetRateChangeLimit(_rateChangeLimit *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SetRateChangeLimit(&_StakeManager.TransactOpts, _rateChangeLimit)
}

// SetRateChangeLimit is a paid mutator transaction binding the contract method 0x19826e71.
//
// Solidity: function setRateChangeLimit(uint256 _rateChangeLimit) returns()
func (_StakeManager *StakeManagerTransactorSession) SetRateChangeLimit(_rateChangeLimit *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SetRateChangeLimit(&_StakeManager.TransactOpts, _rateChangeLimit)
}

// SetUnbondingDuration is a paid mutator transaction binding the contract method 0x40decaad.
//
// Solidity: function setUnbondingDuration(uint256 _unbondingDuration) returns()
func (_StakeManager *StakeManagerTransactor) SetUnbondingDuration(opts *bind.TransactOpts, _unbondingDuration *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "setUnbondingDuration", _unbondingDuration)
}

// SetUnbondingDuration is a paid mutator transaction binding the contract method 0x40decaad.
//
// Solidity: function setUnbondingDuration(uint256 _unbondingDuration) returns()
func (_StakeManager *StakeManagerSession) SetUnbondingDuration(_unbondingDuration *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SetUnbondingDuration(&_StakeManager.TransactOpts, _unbondingDuration)
}

// SetUnbondingDuration is a paid mutator transaction binding the contract method 0x40decaad.
//
// Solidity: function setUnbondingDuration(uint256 _unbondingDuration) returns()
func (_StakeManager *StakeManagerTransactorSession) SetUnbondingDuration(_unbondingDuration *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SetUnbondingDuration(&_StakeManager.TransactOpts, _unbondingDuration)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address _stablecoin, uint256 _stakeAmount) returns()
func (_StakeManager *StakeManagerTransactor) Stake(opts *bind.TransactOpts, _stablecoin common.Address, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "stake", _stablecoin, _stakeAmount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address _stablecoin, uint256 _stakeAmount) returns()
func (_StakeManager *StakeManagerSession) Stake(_stablecoin common.Address, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Stake(&_StakeManager.TransactOpts, _stablecoin, _stakeAmount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address _stablecoin, uint256 _stakeAmount) returns()
func (_StakeManager *StakeManagerTransactorSession) Stake(_stablecoin common.Address, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Stake(&_StakeManager.TransactOpts, _stablecoin, _stakeAmount)
}

// StakeWithPool is a paid mutator transaction binding the contract method 0x8ce39e52.
//
// Solidity: function stakeWithPool(address _poolAddress, address _stablecoin, uint256 _stakeAmount) returns()
func (_StakeManager *StakeManagerTransactor) StakeWithPool(opts *bind.TransactOpts, _poolAddress common.Address, _stablecoin common.Address, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "stakeWithPool", _poolAddress, _stablecoin, _stakeAmount)
}

// StakeWithPool is a paid mutator transaction binding the contract method 0x8ce39e52.
//
// Solidity: function stakeWithPool(address _poolAddress, address _stablecoin, uint256 _stakeAmount) returns()
func (_StakeManager *StakeManagerSession) StakeWithPool(_poolAddress common.Address, _stablecoin common.Address, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.StakeWithPool(&_StakeManager.TransactOpts, _poolAddress, _stablecoin, _stakeAmount)
}

// StakeWithPool is a paid mutator transaction binding the contract method 0x8ce39e52.
//
// Solidity: function stakeWithPool(address _poolAddress, address _stablecoin, uint256 _stakeAmount) returns()
func (_StakeManager *StakeManagerTransactorSession) StakeWithPool(_poolAddress common.Address, _stablecoin common.Address, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.StakeWithPool(&_StakeManager.TransactOpts, _poolAddress, _stablecoin, _stakeAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_StakeManager *StakeManagerTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_StakeManager *StakeManagerSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.TransferOwnership(&_StakeManager.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_StakeManager *StakeManagerTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.TransferOwnership(&_StakeManager.TransactOpts, _newOwner)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _stablecoin, uint256 _lsdTokenAmount) returns()
func (_StakeManager *StakeManagerTransactor) Unstake(opts *bind.TransactOpts, _stablecoin common.Address, _lsdTokenAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "unstake", _stablecoin, _lsdTokenAmount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _stablecoin, uint256 _lsdTokenAmount) returns()
func (_StakeManager *StakeManagerSession) Unstake(_stablecoin common.Address, _lsdTokenAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Unstake(&_StakeManager.TransactOpts, _stablecoin, _lsdTokenAmount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _stablecoin, uint256 _lsdTokenAmount) returns()
func (_StakeManager *StakeManagerTransactorSession) Unstake(_stablecoin common.Address, _lsdTokenAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Unstake(&_StakeManager.TransactOpts, _stablecoin, _lsdTokenAmount)
}

// UnstakeWithPool is a paid mutator transaction binding the contract method 0x3e49716a.
//
// Solidity: function unstakeWithPool(address _poolAddress, address _stablecoin, uint256 _lsdTokenAmount) returns()
func (_StakeManager *StakeManagerTransactor) UnstakeWithPool(opts *bind.TransactOpts, _poolAddress common.Address, _stablecoin common.Address, _lsdTokenAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "unstakeWithPool", _poolAddress, _stablecoin, _lsdTokenAmount)
}

// UnstakeWithPool is a paid mutator transaction binding the contract method 0x3e49716a.
//
// Solidity: function unstakeWithPool(address _poolAddress, address _stablecoin, uint256 _lsdTokenAmount) returns()
func (_StakeManager *StakeManagerSession) UnstakeWithPool(_poolAddress common.Address, _stablecoin common.Address, _lsdTokenAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UnstakeWithPool(&_StakeManager.TransactOpts, _poolAddress, _stablecoin, _lsdTokenAmount)
}

// UnstakeWithPool is a paid mutator transaction binding the contract method 0x3e49716a.
//
// Solidity: function unstakeWithPool(address _poolAddress, address _stablecoin, uint256 _lsdTokenAmount) returns()
func (_StakeManager *StakeManagerTransactorSession) UnstakeWithPool(_poolAddress common.Address, _stablecoin common.Address, _lsdTokenAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UnstakeWithPool(&_StakeManager.TransactOpts, _poolAddress, _stablecoin, _lsdTokenAmount)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakeManager *StakeManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakeManager *StakeManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakeManager.Contract.UpgradeToAndCall(&_StakeManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakeManager *StakeManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakeManager.Contract.UpgradeToAndCall(&_StakeManager.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_StakeManager *StakeManagerTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_StakeManager *StakeManagerSession) Withdraw() (*types.Transaction, error) {
	return _StakeManager.Contract.Withdraw(&_StakeManager.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_StakeManager *StakeManagerTransactorSession) Withdraw() (*types.Transaction, error) {
	return _StakeManager.Contract.Withdraw(&_StakeManager.TransactOpts)
}

// WithdrawProtocolFee is a paid mutator transaction binding the contract method 0x668fb6dc.
//
// Solidity: function withdrawProtocolFee(address _to) returns()
func (_StakeManager *StakeManagerTransactor) WithdrawProtocolFee(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "withdrawProtocolFee", _to)
}

// WithdrawProtocolFee is a paid mutator transaction binding the contract method 0x668fb6dc.
//
// Solidity: function withdrawProtocolFee(address _to) returns()
func (_StakeManager *StakeManagerSession) WithdrawProtocolFee(_to common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawProtocolFee(&_StakeManager.TransactOpts, _to)
}

// WithdrawProtocolFee is a paid mutator transaction binding the contract method 0x668fb6dc.
//
// Solidity: function withdrawProtocolFee(address _to) returns()
func (_StakeManager *StakeManagerTransactorSession) WithdrawProtocolFee(_to common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawProtocolFee(&_StakeManager.TransactOpts, _to)
}

// WithdrawWithPool is a paid mutator transaction binding the contract method 0xf737abac.
//
// Solidity: function withdrawWithPool(address _poolAddress) returns()
func (_StakeManager *StakeManagerTransactor) WithdrawWithPool(opts *bind.TransactOpts, _poolAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "withdrawWithPool", _poolAddress)
}

// WithdrawWithPool is a paid mutator transaction binding the contract method 0xf737abac.
//
// Solidity: function withdrawWithPool(address _poolAddress) returns()
func (_StakeManager *StakeManagerSession) WithdrawWithPool(_poolAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawWithPool(&_StakeManager.TransactOpts, _poolAddress)
}

// WithdrawWithPool is a paid mutator transaction binding the contract method 0xf737abac.
//
// Solidity: function withdrawWithPool(address _poolAddress) returns()
func (_StakeManager *StakeManagerTransactorSession) WithdrawWithPool(_poolAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawWithPool(&_StakeManager.TransactOpts, _poolAddress)
}

// StakeManagerExecuteNewEraIterator is returned from FilterExecuteNewEra and is used to iterate over the raw logs and unpacked data for ExecuteNewEra events raised by the StakeManager contract.
type StakeManagerExecuteNewEraIterator struct {
	Event *StakeManagerExecuteNewEra // Event containing the contract specifics and raw log

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
func (it *StakeManagerExecuteNewEraIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerExecuteNewEra)
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
		it.Event = new(StakeManagerExecuteNewEra)
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
func (it *StakeManagerExecuteNewEraIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerExecuteNewEraIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerExecuteNewEra represents a ExecuteNewEra event raised by the StakeManager contract.
type StakeManagerExecuteNewEra struct {
	Era  *big.Int
	Rate *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterExecuteNewEra is a free log retrieval operation binding the contract event 0x02105621fc31aa3ac04a9845beacd54c700e2ab23ff8acdd755dfd878ae61f02.
//
// Solidity: event ExecuteNewEra(uint256 indexed era, uint256 rate)
func (_StakeManager *StakeManagerFilterer) FilterExecuteNewEra(opts *bind.FilterOpts, era []*big.Int) (*StakeManagerExecuteNewEraIterator, error) {

	var eraRule []interface{}
	for _, eraItem := range era {
		eraRule = append(eraRule, eraItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "ExecuteNewEra", eraRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerExecuteNewEraIterator{contract: _StakeManager.contract, event: "ExecuteNewEra", logs: logs, sub: sub}, nil
}

// WatchExecuteNewEra is a free log subscription operation binding the contract event 0x02105621fc31aa3ac04a9845beacd54c700e2ab23ff8acdd755dfd878ae61f02.
//
// Solidity: event ExecuteNewEra(uint256 indexed era, uint256 rate)
func (_StakeManager *StakeManagerFilterer) WatchExecuteNewEra(opts *bind.WatchOpts, sink chan<- *StakeManagerExecuteNewEra, era []*big.Int) (event.Subscription, error) {

	var eraRule []interface{}
	for _, eraItem := range era {
		eraRule = append(eraRule, eraItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "ExecuteNewEra", eraRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerExecuteNewEra)
				if err := _StakeManager.contract.UnpackLog(event, "ExecuteNewEra", log); err != nil {
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

// ParseExecuteNewEra is a log parse operation binding the contract event 0x02105621fc31aa3ac04a9845beacd54c700e2ab23ff8acdd755dfd878ae61f02.
//
// Solidity: event ExecuteNewEra(uint256 indexed era, uint256 rate)
func (_StakeManager *StakeManagerFilterer) ParseExecuteNewEra(log types.Log) (*StakeManagerExecuteNewEra, error) {
	event := new(StakeManagerExecuteNewEra)
	if err := _StakeManager.contract.UnpackLog(event, "ExecuteNewEra", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StakeManager contract.
type StakeManagerInitializedIterator struct {
	Event *StakeManagerInitialized // Event containing the contract specifics and raw log

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
func (it *StakeManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerInitialized)
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
		it.Event = new(StakeManagerInitialized)
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
func (it *StakeManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerInitialized represents a Initialized event raised by the StakeManager contract.
type StakeManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StakeManager *StakeManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakeManagerInitializedIterator, error) {

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakeManagerInitializedIterator{contract: _StakeManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StakeManager *StakeManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakeManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerInitialized)
				if err := _StakeManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StakeManager *StakeManagerFilterer) ParseInitialized(log types.Log) (*StakeManagerInitialized, error) {
	event := new(StakeManagerInitialized)
	if err := _StakeManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerNewRewardIterator is returned from FilterNewReward and is used to iterate over the raw logs and unpacked data for NewReward events raised by the StakeManager contract.
type StakeManagerNewRewardIterator struct {
	Event *StakeManagerNewReward // Event containing the contract specifics and raw log

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
func (it *StakeManagerNewRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerNewReward)
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
		it.Event = new(StakeManagerNewReward)
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
func (it *StakeManagerNewRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerNewRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerNewReward represents a NewReward event raised by the StakeManager contract.
type StakeManagerNewReward struct {
	PoolAddress common.Address
	NewReward   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewReward is a free log retrieval operation binding the contract event 0x01f132e60506239e8c7aeeb029ac2aa3a506bd76fda96ffa1f02e2b193399905.
//
// Solidity: event NewReward(address poolAddress, uint256 newReward)
func (_StakeManager *StakeManagerFilterer) FilterNewReward(opts *bind.FilterOpts) (*StakeManagerNewRewardIterator, error) {

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "NewReward")
	if err != nil {
		return nil, err
	}
	return &StakeManagerNewRewardIterator{contract: _StakeManager.contract, event: "NewReward", logs: logs, sub: sub}, nil
}

// WatchNewReward is a free log subscription operation binding the contract event 0x01f132e60506239e8c7aeeb029ac2aa3a506bd76fda96ffa1f02e2b193399905.
//
// Solidity: event NewReward(address poolAddress, uint256 newReward)
func (_StakeManager *StakeManagerFilterer) WatchNewReward(opts *bind.WatchOpts, sink chan<- *StakeManagerNewReward) (event.Subscription, error) {

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "NewReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerNewReward)
				if err := _StakeManager.contract.UnpackLog(event, "NewReward", log); err != nil {
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

// ParseNewReward is a log parse operation binding the contract event 0x01f132e60506239e8c7aeeb029ac2aa3a506bd76fda96ffa1f02e2b193399905.
//
// Solidity: event NewReward(address poolAddress, uint256 newReward)
func (_StakeManager *StakeManagerFilterer) ParseNewReward(log types.Log) (*StakeManagerNewReward, error) {
	event := new(StakeManagerNewReward)
	if err := _StakeManager.contract.UnpackLog(event, "NewReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakeManager contract.
type StakeManagerOwnershipTransferredIterator struct {
	Event *StakeManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakeManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerOwnershipTransferred)
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
		it.Event = new(StakeManagerOwnershipTransferred)
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
func (it *StakeManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerOwnershipTransferred represents a OwnershipTransferred event raised by the StakeManager contract.
type StakeManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakeManager *StakeManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakeManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerOwnershipTransferredIterator{contract: _StakeManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakeManager *StakeManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakeManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerOwnershipTransferred)
				if err := _StakeManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakeManager *StakeManagerFilterer) ParseOwnershipTransferred(log types.Log) (*StakeManagerOwnershipTransferred, error) {
	event := new(StakeManagerOwnershipTransferred)
	if err := _StakeManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerSetUnbondingDurationIterator is returned from FilterSetUnbondingDuration and is used to iterate over the raw logs and unpacked data for SetUnbondingDuration events raised by the StakeManager contract.
type StakeManagerSetUnbondingDurationIterator struct {
	Event *StakeManagerSetUnbondingDuration // Event containing the contract specifics and raw log

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
func (it *StakeManagerSetUnbondingDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerSetUnbondingDuration)
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
		it.Event = new(StakeManagerSetUnbondingDuration)
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
func (it *StakeManagerSetUnbondingDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerSetUnbondingDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerSetUnbondingDuration represents a SetUnbondingDuration event raised by the StakeManager contract.
type StakeManagerSetUnbondingDuration struct {
	UnbondingDuration *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetUnbondingDuration is a free log retrieval operation binding the contract event 0x279066b062766bf26597e98ef1d6fb6ec39502061f95f271089f727c875414d0.
//
// Solidity: event SetUnbondingDuration(uint256 unbondingDuration)
func (_StakeManager *StakeManagerFilterer) FilterSetUnbondingDuration(opts *bind.FilterOpts) (*StakeManagerSetUnbondingDurationIterator, error) {

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "SetUnbondingDuration")
	if err != nil {
		return nil, err
	}
	return &StakeManagerSetUnbondingDurationIterator{contract: _StakeManager.contract, event: "SetUnbondingDuration", logs: logs, sub: sub}, nil
}

// WatchSetUnbondingDuration is a free log subscription operation binding the contract event 0x279066b062766bf26597e98ef1d6fb6ec39502061f95f271089f727c875414d0.
//
// Solidity: event SetUnbondingDuration(uint256 unbondingDuration)
func (_StakeManager *StakeManagerFilterer) WatchSetUnbondingDuration(opts *bind.WatchOpts, sink chan<- *StakeManagerSetUnbondingDuration) (event.Subscription, error) {

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "SetUnbondingDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerSetUnbondingDuration)
				if err := _StakeManager.contract.UnpackLog(event, "SetUnbondingDuration", log); err != nil {
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

// ParseSetUnbondingDuration is a log parse operation binding the contract event 0x279066b062766bf26597e98ef1d6fb6ec39502061f95f271089f727c875414d0.
//
// Solidity: event SetUnbondingDuration(uint256 unbondingDuration)
func (_StakeManager *StakeManagerFilterer) ParseSetUnbondingDuration(log types.Log) (*StakeManagerSetUnbondingDuration, error) {
	event := new(StakeManagerSetUnbondingDuration)
	if err := _StakeManager.contract.UnpackLog(event, "SetUnbondingDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerStakeIterator is returned from FilterStake and is used to iterate over the raw logs and unpacked data for Stake events raised by the StakeManager contract.
type StakeManagerStakeIterator struct {
	Event *StakeManagerStake // Event containing the contract specifics and raw log

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
func (it *StakeManagerStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerStake)
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
		it.Event = new(StakeManagerStake)
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
func (it *StakeManagerStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerStake represents a Stake event raised by the StakeManager contract.
type StakeManagerStake struct {
	Staker         common.Address
	PoolAddress    common.Address
	Stablecoin     common.Address
	TokenAmount    *big.Int
	LsdTokenAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStake is a free log retrieval operation binding the contract event 0x592d8fba044f9498232ce6c14dc7eb3a60f723f4d6731918aea4f6ae0c081044.
//
// Solidity: event Stake(address staker, address poolAddress, address stablecoin, uint256 tokenAmount, uint256 lsdTokenAmount)
func (_StakeManager *StakeManagerFilterer) FilterStake(opts *bind.FilterOpts) (*StakeManagerStakeIterator, error) {

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "Stake")
	if err != nil {
		return nil, err
	}
	return &StakeManagerStakeIterator{contract: _StakeManager.contract, event: "Stake", logs: logs, sub: sub}, nil
}

// WatchStake is a free log subscription operation binding the contract event 0x592d8fba044f9498232ce6c14dc7eb3a60f723f4d6731918aea4f6ae0c081044.
//
// Solidity: event Stake(address staker, address poolAddress, address stablecoin, uint256 tokenAmount, uint256 lsdTokenAmount)
func (_StakeManager *StakeManagerFilterer) WatchStake(opts *bind.WatchOpts, sink chan<- *StakeManagerStake) (event.Subscription, error) {

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "Stake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerStake)
				if err := _StakeManager.contract.UnpackLog(event, "Stake", log); err != nil {
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

// ParseStake is a log parse operation binding the contract event 0x592d8fba044f9498232ce6c14dc7eb3a60f723f4d6731918aea4f6ae0c081044.
//
// Solidity: event Stake(address staker, address poolAddress, address stablecoin, uint256 tokenAmount, uint256 lsdTokenAmount)
func (_StakeManager *StakeManagerFilterer) ParseStake(log types.Log) (*StakeManagerStake, error) {
	event := new(StakeManagerStake)
	if err := _StakeManager.contract.UnpackLog(event, "Stake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerUnstakeIterator is returned from FilterUnstake and is used to iterate over the raw logs and unpacked data for Unstake events raised by the StakeManager contract.
type StakeManagerUnstakeIterator struct {
	Event *StakeManagerUnstake // Event containing the contract specifics and raw log

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
func (it *StakeManagerUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerUnstake)
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
		it.Event = new(StakeManagerUnstake)
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
func (it *StakeManagerUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerUnstake represents a Unstake event raised by the StakeManager contract.
type StakeManagerUnstake struct {
	Staker         common.Address
	PoolAddress    common.Address
	Stablecoin     common.Address
	TokenAmount    *big.Int
	LsdTokenAmount *big.Int
	UnstakeIndex   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUnstake is a free log retrieval operation binding the contract event 0x2c904b13f3b9bc1fa77b4e439e2deb5507c11d7a25e042a5bf6374837582f046.
//
// Solidity: event Unstake(address staker, address poolAddress, address stablecoin, uint256 tokenAmount, uint256 lsdTokenAmount, uint256 unstakeIndex)
func (_StakeManager *StakeManagerFilterer) FilterUnstake(opts *bind.FilterOpts) (*StakeManagerUnstakeIterator, error) {

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "Unstake")
	if err != nil {
		return nil, err
	}
	return &StakeManagerUnstakeIterator{contract: _StakeManager.contract, event: "Unstake", logs: logs, sub: sub}, nil
}

// WatchUnstake is a free log subscription operation binding the contract event 0x2c904b13f3b9bc1fa77b4e439e2deb5507c11d7a25e042a5bf6374837582f046.
//
// Solidity: event Unstake(address staker, address poolAddress, address stablecoin, uint256 tokenAmount, uint256 lsdTokenAmount, uint256 unstakeIndex)
func (_StakeManager *StakeManagerFilterer) WatchUnstake(opts *bind.WatchOpts, sink chan<- *StakeManagerUnstake) (event.Subscription, error) {

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "Unstake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerUnstake)
				if err := _StakeManager.contract.UnpackLog(event, "Unstake", log); err != nil {
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

// ParseUnstake is a log parse operation binding the contract event 0x2c904b13f3b9bc1fa77b4e439e2deb5507c11d7a25e042a5bf6374837582f046.
//
// Solidity: event Unstake(address staker, address poolAddress, address stablecoin, uint256 tokenAmount, uint256 lsdTokenAmount, uint256 unstakeIndex)
func (_StakeManager *StakeManagerFilterer) ParseUnstake(log types.Log) (*StakeManagerUnstake, error) {
	event := new(StakeManagerUnstake)
	if err := _StakeManager.contract.UnpackLog(event, "Unstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the StakeManager contract.
type StakeManagerUpgradedIterator struct {
	Event *StakeManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *StakeManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerUpgraded)
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
		it.Event = new(StakeManagerUpgraded)
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
func (it *StakeManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerUpgraded represents a Upgraded event raised by the StakeManager contract.
type StakeManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StakeManager *StakeManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StakeManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerUpgradedIterator{contract: _StakeManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StakeManager *StakeManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StakeManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerUpgraded)
				if err := _StakeManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_StakeManager *StakeManagerFilterer) ParseUpgraded(log types.Log) (*StakeManagerUpgraded, error) {
	event := new(StakeManagerUpgraded)
	if err := _StakeManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the StakeManager contract.
type StakeManagerWithdrawIterator struct {
	Event *StakeManagerWithdraw // Event containing the contract specifics and raw log

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
func (it *StakeManagerWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerWithdraw)
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
		it.Event = new(StakeManagerWithdraw)
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
func (it *StakeManagerWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerWithdraw represents a Withdraw event raised by the StakeManager contract.
type StakeManagerWithdraw struct {
	Staker           common.Address
	PoolAddress      common.Address
	UnstakeIndexList []*big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xcdbb3b8ca00299a5a4a191f420fc23a9d08297641ae04dd0cd15f6c989d53972.
//
// Solidity: event Withdraw(address staker, address poolAddress, int256[] unstakeIndexList)
func (_StakeManager *StakeManagerFilterer) FilterWithdraw(opts *bind.FilterOpts) (*StakeManagerWithdrawIterator, error) {

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &StakeManagerWithdrawIterator{contract: _StakeManager.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xcdbb3b8ca00299a5a4a191f420fc23a9d08297641ae04dd0cd15f6c989d53972.
//
// Solidity: event Withdraw(address staker, address poolAddress, int256[] unstakeIndexList)
func (_StakeManager *StakeManagerFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *StakeManagerWithdraw) (event.Subscription, error) {

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerWithdraw)
				if err := _StakeManager.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xcdbb3b8ca00299a5a4a191f420fc23a9d08297641ae04dd0cd15f6c989d53972.
//
// Solidity: event Withdraw(address staker, address poolAddress, int256[] unstakeIndexList)
func (_StakeManager *StakeManagerFilterer) ParseWithdraw(log types.Log) (*StakeManagerWithdraw, error) {
	event := new(StakeManagerWithdraw)
	if err := _StakeManager.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
