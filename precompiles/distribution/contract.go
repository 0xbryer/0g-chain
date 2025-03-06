// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package distribution

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
)

// DistributionMetaData contains all meta data concerning the Distribution contract.
var DistributionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CommunityPool\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structDecCoin[]\",\"name\":\"pool\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"delegationRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structDecCoin[]\",\"name\":\"rewards\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"}],\"name\":\"delegationTotalRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structDecCoin[]\",\"name\":\"reward\",\"type\":\"tuple[]\"}],\"internalType\":\"structDelegationDelegatorReward[]\",\"name\":\"rewards\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structDecCoin[]\",\"name\":\"total\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"}],\"name\":\"delegatorValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"}],\"name\":\"delegatorWithdrawAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundCommunityPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"params\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"communityTax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseProposerReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bonusProposerReward\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"withdrawAddrEnabled\",\"type\":\"bool\"}],\"internalType\":\"structParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"withdrawerAddress\",\"type\":\"address\"}],\"name\":\"setWithdrawAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"}],\"name\":\"validatorCommission\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structDecCoin[]\",\"name\":\"commission\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"validatorDistributionInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structDecCoin[]\",\"name\":\"selfBondRewards\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structDecCoin[]\",\"name\":\"commission\",\"type\":\"tuple[]\"}],\"internalType\":\"structValidatorDistributionInfoResponse\",\"name\":\"distributionInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"validatorOutstandingRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structDecCoin[]\",\"name\":\"rewards\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startingHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endingHeight\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pageRequest\",\"type\":\"tuple\"}],\"name\":\"validatorSlashes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"validatorPeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"fraction\",\"type\":\"uint256\"}],\"internalType\":\"structValidatorSlashEvent[]\",\"name\":\"slashes\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"}],\"name\":\"withdrawDelegatorRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin[]\",\"name\":\"amount\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"withdrawValidatorCommission\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin[]\",\"name\":\"amount\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DistributionABI is the input ABI used to generate the binding from.
// Deprecated: Use DistributionMetaData.ABI instead.
var DistributionABI = DistributionMetaData.ABI

// Distribution is an auto generated Go binding around an Ethereum contract.
type Distribution struct {
	DistributionCaller     // Read-only binding to the contract
	DistributionTransactor // Write-only binding to the contract
	DistributionFilterer   // Log filterer for contract events
}

// DistributionCaller is an auto generated read-only Go binding around an Ethereum contract.
type DistributionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DistributionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DistributionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DistributionSession struct {
	Contract     *Distribution     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DistributionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DistributionCallerSession struct {
	Contract *DistributionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DistributionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DistributionTransactorSession struct {
	Contract     *DistributionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DistributionRaw is an auto generated low-level Go binding around an Ethereum contract.
type DistributionRaw struct {
	Contract *Distribution // Generic contract binding to access the raw methods on
}

// DistributionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DistributionCallerRaw struct {
	Contract *DistributionCaller // Generic read-only contract binding to access the raw methods on
}

// DistributionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DistributionTransactorRaw struct {
	Contract *DistributionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDistribution creates a new instance of Distribution, bound to a specific deployed contract.
func NewDistribution(address common.Address, backend bind.ContractBackend) (*Distribution, error) {
	contract, err := bindDistribution(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Distribution{DistributionCaller: DistributionCaller{contract: contract}, DistributionTransactor: DistributionTransactor{contract: contract}, DistributionFilterer: DistributionFilterer{contract: contract}}, nil
}

// NewDistributionCaller creates a new read-only instance of Distribution, bound to a specific deployed contract.
func NewDistributionCaller(address common.Address, caller bind.ContractCaller) (*DistributionCaller, error) {
	contract, err := bindDistribution(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DistributionCaller{contract: contract}, nil
}

// NewDistributionTransactor creates a new write-only instance of Distribution, bound to a specific deployed contract.
func NewDistributionTransactor(address common.Address, transactor bind.ContractTransactor) (*DistributionTransactor, error) {
	contract, err := bindDistribution(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DistributionTransactor{contract: contract}, nil
}

// NewDistributionFilterer creates a new log filterer instance of Distribution, bound to a specific deployed contract.
func NewDistributionFilterer(address common.Address, filterer bind.ContractFilterer) (*DistributionFilterer, error) {
	contract, err := bindDistribution(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DistributionFilterer{contract: contract}, nil
}

// bindDistribution binds a generic wrapper to an already deployed contract.
func bindDistribution(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DistributionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Distribution *DistributionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Distribution.Contract.DistributionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Distribution *DistributionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distribution.Contract.DistributionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Distribution *DistributionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Distribution.Contract.DistributionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Distribution *DistributionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Distribution.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Distribution *DistributionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distribution.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Distribution *DistributionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Distribution.Contract.contract.Transact(opts, method, params...)
}

// CommunityPool is a free data retrieval call binding the contract method 0x5637d7c9.
//
// Solidity: function CommunityPool() view returns((string,uint256)[] pool)
func (_Distribution *DistributionCaller) CommunityPool(opts *bind.CallOpts) ([]DecCoin, error) {
	var out []interface{}
	err := _Distribution.contract.Call(opts, &out, "CommunityPool")

	if err != nil {
		return *new([]DecCoin), err
	}

	out0 := *abi.ConvertType(out[0], new([]DecCoin)).(*[]DecCoin)

	return out0, err

}

// CommunityPool is a free data retrieval call binding the contract method 0x5637d7c9.
//
// Solidity: function CommunityPool() view returns((string,uint256)[] pool)
func (_Distribution *DistributionSession) CommunityPool() ([]DecCoin, error) {
	return _Distribution.Contract.CommunityPool(&_Distribution.CallOpts)
}

// CommunityPool is a free data retrieval call binding the contract method 0x5637d7c9.
//
// Solidity: function CommunityPool() view returns((string,uint256)[] pool)
func (_Distribution *DistributionCallerSession) CommunityPool() ([]DecCoin, error) {
	return _Distribution.Contract.CommunityPool(&_Distribution.CallOpts)
}

// DelegationRewards is a free data retrieval call binding the contract method 0xc9a21b7b.
//
// Solidity: function delegationRewards(address delegatorAddress, address validatorAddress) view returns((string,uint256)[] rewards)
func (_Distribution *DistributionCaller) DelegationRewards(opts *bind.CallOpts, delegatorAddress common.Address, validatorAddress common.Address) ([]DecCoin, error) {
	var out []interface{}
	err := _Distribution.contract.Call(opts, &out, "delegationRewards", delegatorAddress, validatorAddress)

	if err != nil {
		return *new([]DecCoin), err
	}

	out0 := *abi.ConvertType(out[0], new([]DecCoin)).(*[]DecCoin)

	return out0, err

}

// DelegationRewards is a free data retrieval call binding the contract method 0xc9a21b7b.
//
// Solidity: function delegationRewards(address delegatorAddress, address validatorAddress) view returns((string,uint256)[] rewards)
func (_Distribution *DistributionSession) DelegationRewards(delegatorAddress common.Address, validatorAddress common.Address) ([]DecCoin, error) {
	return _Distribution.Contract.DelegationRewards(&_Distribution.CallOpts, delegatorAddress, validatorAddress)
}

// DelegationRewards is a free data retrieval call binding the contract method 0xc9a21b7b.
//
// Solidity: function delegationRewards(address delegatorAddress, address validatorAddress) view returns((string,uint256)[] rewards)
func (_Distribution *DistributionCallerSession) DelegationRewards(delegatorAddress common.Address, validatorAddress common.Address) ([]DecCoin, error) {
	return _Distribution.Contract.DelegationRewards(&_Distribution.CallOpts, delegatorAddress, validatorAddress)
}

// DelegationTotalRewards is a free data retrieval call binding the contract method 0x54be1a28.
//
// Solidity: function delegationTotalRewards(address delegatorAddress) view returns((address,(string,uint256)[])[] rewards, (string,uint256)[] total)
func (_Distribution *DistributionCaller) DelegationTotalRewards(opts *bind.CallOpts, delegatorAddress common.Address) (struct {
	Rewards []DelegationDelegatorReward
	Total   []DecCoin
}, error) {
	var out []interface{}
	err := _Distribution.contract.Call(opts, &out, "delegationTotalRewards", delegatorAddress)

	outstruct := new(struct {
		Rewards []DelegationDelegatorReward
		Total   []DecCoin
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Rewards = *abi.ConvertType(out[0], new([]DelegationDelegatorReward)).(*[]DelegationDelegatorReward)
	outstruct.Total = *abi.ConvertType(out[1], new([]DecCoin)).(*[]DecCoin)

	return *outstruct, err

}

// DelegationTotalRewards is a free data retrieval call binding the contract method 0x54be1a28.
//
// Solidity: function delegationTotalRewards(address delegatorAddress) view returns((address,(string,uint256)[])[] rewards, (string,uint256)[] total)
func (_Distribution *DistributionSession) DelegationTotalRewards(delegatorAddress common.Address) (struct {
	Rewards []DelegationDelegatorReward
	Total   []DecCoin
}, error) {
	return _Distribution.Contract.DelegationTotalRewards(&_Distribution.CallOpts, delegatorAddress)
}

// DelegationTotalRewards is a free data retrieval call binding the contract method 0x54be1a28.
//
// Solidity: function delegationTotalRewards(address delegatorAddress) view returns((address,(string,uint256)[])[] rewards, (string,uint256)[] total)
func (_Distribution *DistributionCallerSession) DelegationTotalRewards(delegatorAddress common.Address) (struct {
	Rewards []DelegationDelegatorReward
	Total   []DecCoin
}, error) {
	return _Distribution.Contract.DelegationTotalRewards(&_Distribution.CallOpts, delegatorAddress)
}

// DelegatorValidators is a free data retrieval call binding the contract method 0xa66cb605.
//
// Solidity: function delegatorValidators(address delegatorAddress) view returns(address[] validators)
func (_Distribution *DistributionCaller) DelegatorValidators(opts *bind.CallOpts, delegatorAddress common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Distribution.contract.Call(opts, &out, "delegatorValidators", delegatorAddress)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// DelegatorValidators is a free data retrieval call binding the contract method 0xa66cb605.
//
// Solidity: function delegatorValidators(address delegatorAddress) view returns(address[] validators)
func (_Distribution *DistributionSession) DelegatorValidators(delegatorAddress common.Address) ([]common.Address, error) {
	return _Distribution.Contract.DelegatorValidators(&_Distribution.CallOpts, delegatorAddress)
}

// DelegatorValidators is a free data retrieval call binding the contract method 0xa66cb605.
//
// Solidity: function delegatorValidators(address delegatorAddress) view returns(address[] validators)
func (_Distribution *DistributionCallerSession) DelegatorValidators(delegatorAddress common.Address) ([]common.Address, error) {
	return _Distribution.Contract.DelegatorValidators(&_Distribution.CallOpts, delegatorAddress)
}

// DelegatorWithdrawAddress is a free data retrieval call binding the contract method 0x5431f450.
//
// Solidity: function delegatorWithdrawAddress(address delegatorAddress) view returns(address withdrawAddress)
func (_Distribution *DistributionCaller) DelegatorWithdrawAddress(opts *bind.CallOpts, delegatorAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Distribution.contract.Call(opts, &out, "delegatorWithdrawAddress", delegatorAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegatorWithdrawAddress is a free data retrieval call binding the contract method 0x5431f450.
//
// Solidity: function delegatorWithdrawAddress(address delegatorAddress) view returns(address withdrawAddress)
func (_Distribution *DistributionSession) DelegatorWithdrawAddress(delegatorAddress common.Address) (common.Address, error) {
	return _Distribution.Contract.DelegatorWithdrawAddress(&_Distribution.CallOpts, delegatorAddress)
}

// DelegatorWithdrawAddress is a free data retrieval call binding the contract method 0x5431f450.
//
// Solidity: function delegatorWithdrawAddress(address delegatorAddress) view returns(address withdrawAddress)
func (_Distribution *DistributionCallerSession) DelegatorWithdrawAddress(delegatorAddress common.Address) (common.Address, error) {
	return _Distribution.Contract.DelegatorWithdrawAddress(&_Distribution.CallOpts, delegatorAddress)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns((uint256,uint256,uint256,bool))
func (_Distribution *DistributionCaller) Params(opts *bind.CallOpts) (Params, error) {
	var out []interface{}
	err := _Distribution.contract.Call(opts, &out, "params")

	if err != nil {
		return *new(Params), err
	}

	out0 := *abi.ConvertType(out[0], new(Params)).(*Params)

	return out0, err

}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns((uint256,uint256,uint256,bool))
func (_Distribution *DistributionSession) Params() (Params, error) {
	return _Distribution.Contract.Params(&_Distribution.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns((uint256,uint256,uint256,bool))
func (_Distribution *DistributionCallerSession) Params() (Params, error) {
	return _Distribution.Contract.Params(&_Distribution.CallOpts)
}

// ValidatorCommission is a free data retrieval call binding the contract method 0x3dd40f78.
//
// Solidity: function validatorCommission(string validatorAddress) view returns((string,uint256)[] commission)
func (_Distribution *DistributionCaller) ValidatorCommission(opts *bind.CallOpts, validatorAddress string) ([]DecCoin, error) {
	var out []interface{}
	err := _Distribution.contract.Call(opts, &out, "validatorCommission", validatorAddress)

	if err != nil {
		return *new([]DecCoin), err
	}

	out0 := *abi.ConvertType(out[0], new([]DecCoin)).(*[]DecCoin)

	return out0, err

}

// ValidatorCommission is a free data retrieval call binding the contract method 0x3dd40f78.
//
// Solidity: function validatorCommission(string validatorAddress) view returns((string,uint256)[] commission)
func (_Distribution *DistributionSession) ValidatorCommission(validatorAddress string) ([]DecCoin, error) {
	return _Distribution.Contract.ValidatorCommission(&_Distribution.CallOpts, validatorAddress)
}

// ValidatorCommission is a free data retrieval call binding the contract method 0x3dd40f78.
//
// Solidity: function validatorCommission(string validatorAddress) view returns((string,uint256)[] commission)
func (_Distribution *DistributionCallerSession) ValidatorCommission(validatorAddress string) ([]DecCoin, error) {
	return _Distribution.Contract.ValidatorCommission(&_Distribution.CallOpts, validatorAddress)
}

// ValidatorDistributionInfo is a free data retrieval call binding the contract method 0x7240b338.
//
// Solidity: function validatorDistributionInfo(address validatorAddress) view returns((address,(string,uint256)[],(string,uint256)[]) distributionInfo)
func (_Distribution *DistributionCaller) ValidatorDistributionInfo(opts *bind.CallOpts, validatorAddress common.Address) (ValidatorDistributionInfoResponse, error) {
	var out []interface{}
	err := _Distribution.contract.Call(opts, &out, "validatorDistributionInfo", validatorAddress)

	if err != nil {
		return *new(ValidatorDistributionInfoResponse), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidatorDistributionInfoResponse)).(*ValidatorDistributionInfoResponse)

	return out0, err

}

// ValidatorDistributionInfo is a free data retrieval call binding the contract method 0x7240b338.
//
// Solidity: function validatorDistributionInfo(address validatorAddress) view returns((address,(string,uint256)[],(string,uint256)[]) distributionInfo)
func (_Distribution *DistributionSession) ValidatorDistributionInfo(validatorAddress common.Address) (ValidatorDistributionInfoResponse, error) {
	return _Distribution.Contract.ValidatorDistributionInfo(&_Distribution.CallOpts, validatorAddress)
}

// ValidatorDistributionInfo is a free data retrieval call binding the contract method 0x7240b338.
//
// Solidity: function validatorDistributionInfo(address validatorAddress) view returns((address,(string,uint256)[],(string,uint256)[]) distributionInfo)
func (_Distribution *DistributionCallerSession) ValidatorDistributionInfo(validatorAddress common.Address) (ValidatorDistributionInfoResponse, error) {
	return _Distribution.Contract.ValidatorDistributionInfo(&_Distribution.CallOpts, validatorAddress)
}

// ValidatorOutstandingRewards is a free data retrieval call binding the contract method 0xf28e9b39.
//
// Solidity: function validatorOutstandingRewards(address validatorAddress) view returns((string,uint256)[] rewards)
func (_Distribution *DistributionCaller) ValidatorOutstandingRewards(opts *bind.CallOpts, validatorAddress common.Address) ([]DecCoin, error) {
	var out []interface{}
	err := _Distribution.contract.Call(opts, &out, "validatorOutstandingRewards", validatorAddress)

	if err != nil {
		return *new([]DecCoin), err
	}

	out0 := *abi.ConvertType(out[0], new([]DecCoin)).(*[]DecCoin)

	return out0, err

}

// ValidatorOutstandingRewards is a free data retrieval call binding the contract method 0xf28e9b39.
//
// Solidity: function validatorOutstandingRewards(address validatorAddress) view returns((string,uint256)[] rewards)
func (_Distribution *DistributionSession) ValidatorOutstandingRewards(validatorAddress common.Address) ([]DecCoin, error) {
	return _Distribution.Contract.ValidatorOutstandingRewards(&_Distribution.CallOpts, validatorAddress)
}

// ValidatorOutstandingRewards is a free data retrieval call binding the contract method 0xf28e9b39.
//
// Solidity: function validatorOutstandingRewards(address validatorAddress) view returns((string,uint256)[] rewards)
func (_Distribution *DistributionCallerSession) ValidatorOutstandingRewards(validatorAddress common.Address) ([]DecCoin, error) {
	return _Distribution.Contract.ValidatorOutstandingRewards(&_Distribution.CallOpts, validatorAddress)
}

// ValidatorSlashes is a free data retrieval call binding the contract method 0xbd2e5693.
//
// Solidity: function validatorSlashes(address validatorAddress, uint64 startingHeight, uint64 endingHeight, (bytes,uint64,uint64,bool,bool) pageRequest) view returns((uint64,uint256)[] slashes, (bytes,uint64) pageResponse)
func (_Distribution *DistributionCaller) ValidatorSlashes(opts *bind.CallOpts, validatorAddress common.Address, startingHeight uint64, endingHeight uint64, pageRequest PageRequest) (struct {
	Slashes      []ValidatorSlashEvent
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _Distribution.contract.Call(opts, &out, "validatorSlashes", validatorAddress, startingHeight, endingHeight, pageRequest)

	outstruct := new(struct {
		Slashes      []ValidatorSlashEvent
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Slashes = *abi.ConvertType(out[0], new([]ValidatorSlashEvent)).(*[]ValidatorSlashEvent)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// ValidatorSlashes is a free data retrieval call binding the contract method 0xbd2e5693.
//
// Solidity: function validatorSlashes(address validatorAddress, uint64 startingHeight, uint64 endingHeight, (bytes,uint64,uint64,bool,bool) pageRequest) view returns((uint64,uint256)[] slashes, (bytes,uint64) pageResponse)
func (_Distribution *DistributionSession) ValidatorSlashes(validatorAddress common.Address, startingHeight uint64, endingHeight uint64, pageRequest PageRequest) (struct {
	Slashes      []ValidatorSlashEvent
	PageResponse PageResponse
}, error) {
	return _Distribution.Contract.ValidatorSlashes(&_Distribution.CallOpts, validatorAddress, startingHeight, endingHeight, pageRequest)
}

// ValidatorSlashes is a free data retrieval call binding the contract method 0xbd2e5693.
//
// Solidity: function validatorSlashes(address validatorAddress, uint64 startingHeight, uint64 endingHeight, (bytes,uint64,uint64,bool,bool) pageRequest) view returns((uint64,uint256)[] slashes, (bytes,uint64) pageResponse)
func (_Distribution *DistributionCallerSession) ValidatorSlashes(validatorAddress common.Address, startingHeight uint64, endingHeight uint64, pageRequest PageRequest) (struct {
	Slashes      []ValidatorSlashEvent
	PageResponse PageResponse
}, error) {
	return _Distribution.Contract.ValidatorSlashes(&_Distribution.CallOpts, validatorAddress, startingHeight, endingHeight, pageRequest)
}

// FundCommunityPool is a paid mutator transaction binding the contract method 0xed41d0b6.
//
// Solidity: function fundCommunityPool(address depositor, uint256 amount) returns(bool success)
func (_Distribution *DistributionTransactor) FundCommunityPool(opts *bind.TransactOpts, depositor common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Distribution.contract.Transact(opts, "fundCommunityPool", depositor, amount)
}

// FundCommunityPool is a paid mutator transaction binding the contract method 0xed41d0b6.
//
// Solidity: function fundCommunityPool(address depositor, uint256 amount) returns(bool success)
func (_Distribution *DistributionSession) FundCommunityPool(depositor common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Distribution.Contract.FundCommunityPool(&_Distribution.TransactOpts, depositor, amount)
}

// FundCommunityPool is a paid mutator transaction binding the contract method 0xed41d0b6.
//
// Solidity: function fundCommunityPool(address depositor, uint256 amount) returns(bool success)
func (_Distribution *DistributionTransactorSession) FundCommunityPool(depositor common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Distribution.Contract.FundCommunityPool(&_Distribution.TransactOpts, depositor, amount)
}

// SetWithdrawAddress is a paid mutator transaction binding the contract method 0xe59a2f99.
//
// Solidity: function setWithdrawAddress(address delegatorAddress, address withdrawerAddress) returns(bool success)
func (_Distribution *DistributionTransactor) SetWithdrawAddress(opts *bind.TransactOpts, delegatorAddress common.Address, withdrawerAddress common.Address) (*types.Transaction, error) {
	return _Distribution.contract.Transact(opts, "setWithdrawAddress", delegatorAddress, withdrawerAddress)
}

// SetWithdrawAddress is a paid mutator transaction binding the contract method 0xe59a2f99.
//
// Solidity: function setWithdrawAddress(address delegatorAddress, address withdrawerAddress) returns(bool success)
func (_Distribution *DistributionSession) SetWithdrawAddress(delegatorAddress common.Address, withdrawerAddress common.Address) (*types.Transaction, error) {
	return _Distribution.Contract.SetWithdrawAddress(&_Distribution.TransactOpts, delegatorAddress, withdrawerAddress)
}

// SetWithdrawAddress is a paid mutator transaction binding the contract method 0xe59a2f99.
//
// Solidity: function setWithdrawAddress(address delegatorAddress, address withdrawerAddress) returns(bool success)
func (_Distribution *DistributionTransactorSession) SetWithdrawAddress(delegatorAddress common.Address, withdrawerAddress common.Address) (*types.Transaction, error) {
	return _Distribution.Contract.SetWithdrawAddress(&_Distribution.TransactOpts, delegatorAddress, withdrawerAddress)
}

// WithdrawDelegatorRewards is a paid mutator transaction binding the contract method 0xb46a8d61.
//
// Solidity: function withdrawDelegatorRewards(address delegatorAddress, string validatorAddress) returns((string,uint256)[] amount)
func (_Distribution *DistributionTransactor) WithdrawDelegatorRewards(opts *bind.TransactOpts, delegatorAddress common.Address, validatorAddress string) (*types.Transaction, error) {
	return _Distribution.contract.Transact(opts, "withdrawDelegatorRewards", delegatorAddress, validatorAddress)
}

// WithdrawDelegatorRewards is a paid mutator transaction binding the contract method 0xb46a8d61.
//
// Solidity: function withdrawDelegatorRewards(address delegatorAddress, string validatorAddress) returns((string,uint256)[] amount)
func (_Distribution *DistributionSession) WithdrawDelegatorRewards(delegatorAddress common.Address, validatorAddress string) (*types.Transaction, error) {
	return _Distribution.Contract.WithdrawDelegatorRewards(&_Distribution.TransactOpts, delegatorAddress, validatorAddress)
}

// WithdrawDelegatorRewards is a paid mutator transaction binding the contract method 0xb46a8d61.
//
// Solidity: function withdrawDelegatorRewards(address delegatorAddress, string validatorAddress) returns((string,uint256)[] amount)
func (_Distribution *DistributionTransactorSession) WithdrawDelegatorRewards(delegatorAddress common.Address, validatorAddress string) (*types.Transaction, error) {
	return _Distribution.Contract.WithdrawDelegatorRewards(&_Distribution.TransactOpts, delegatorAddress, validatorAddress)
}

// WithdrawValidatorCommission is a paid mutator transaction binding the contract method 0x11fc0bfe.
//
// Solidity: function withdrawValidatorCommission(address validatorAddress) returns((string,uint256)[] amount)
func (_Distribution *DistributionTransactor) WithdrawValidatorCommission(opts *bind.TransactOpts, validatorAddress common.Address) (*types.Transaction, error) {
	return _Distribution.contract.Transact(opts, "withdrawValidatorCommission", validatorAddress)
}

// WithdrawValidatorCommission is a paid mutator transaction binding the contract method 0x11fc0bfe.
//
// Solidity: function withdrawValidatorCommission(address validatorAddress) returns((string,uint256)[] amount)
func (_Distribution *DistributionSession) WithdrawValidatorCommission(validatorAddress common.Address) (*types.Transaction, error) {
	return _Distribution.Contract.WithdrawValidatorCommission(&_Distribution.TransactOpts, validatorAddress)
}

// WithdrawValidatorCommission is a paid mutator transaction binding the contract method 0x11fc0bfe.
//
// Solidity: function withdrawValidatorCommission(address validatorAddress) returns((string,uint256)[] amount)
func (_Distribution *DistributionTransactorSession) WithdrawValidatorCommission(validatorAddress common.Address) (*types.Transaction, error) {
	return _Distribution.Contract.WithdrawValidatorCommission(&_Distribution.TransactOpts, validatorAddress)
}
