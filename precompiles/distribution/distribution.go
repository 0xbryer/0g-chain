package distribution

import (
	"strings"

	precompiles_common "github.com/0glabs/0g-chain/precompiles/common"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	distributionkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

const (
	PrecompileAddress = "0x0000000000000000000000000000000000001003"

	// tx
	DistributionFunctionSetWithdrawAddress          = "setWithdrawAddress"
	DistributionFunctionWithdrawDelegatorRewards    = "withdrawDelegatorRewards"
	DistributionFunctionWithdrawValidatorCommission = "withdrawValidatorCommission"
	// query
	DistributionFunctionParams                      = "params"
	DistributionFunctionValidatorDistributionInfo   = "validatorDistributionInfo"
	DistributionFunctionValidatorOutstandingRewards = "validatorOutstandingRewards"
	DistributionFunctionValidatorCommission         = "validatorCommission"
	DistributionFunctionValidatorSlashes            = "validatorSlashes"
	DistributionFunctionDelegationRewards           = "delegationRewards"
	DistributionFunctionDelegationTotalRewards      = "delegationTotalRewards"
	DistributionFunctionDelegatorValidators         = "delegatorValidators"
	DistributionFunctionDelegatorWithdrawAddress    = "delegatorWithdrawAddress"
)

var _ vm.PrecompiledContract = &DistributionPrecompile{}
var _ precompiles_common.PrecompileCommon = &DistributionPrecompile{}

type DistributionPrecompile struct {
	abi                abi.ABI
	distributionKeeper distributionkeeper.Keeper
}

func NewDistributionPrecompile(distributionKeeper distributionkeeper.Keeper) (*DistributionPrecompile, error) {
	abi, err := abi.JSON(strings.NewReader(DistributionABI))
	if err != nil {
		return nil, err
	}
	return &DistributionPrecompile{
		abi:                abi,
		distributionKeeper: distributionKeeper,
	}, nil
}

// Address implements vm.PrecompiledContract.
func (d *DistributionPrecompile) Address() common.Address {
	return common.HexToAddress(PrecompileAddress)
}

// RequiredGas implements vm.PrecompiledContract.
func (d *DistributionPrecompile) RequiredGas(input []byte) uint64 {
	return 0
}

func (d *DistributionPrecompile) Abi() *abi.ABI {
	return &d.abi
}

func (d *DistributionPrecompile) IsTx(method string) bool {
	switch method {
	case DistributionFunctionSetWithdrawAddress,
		DistributionFunctionWithdrawDelegatorRewards,
		DistributionFunctionWithdrawValidatorCommission:
		return true
	default:
		return false
	}
}

func (d *DistributionPrecompile) KVGasConfig() storetypes.GasConfig {
	return storetypes.KVGasConfig()
}

// Run implements vm.PrecompiledContract.
func (d *DistributionPrecompile) Run(evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	ctx, stateDB, method, initialGas, args, err := precompiles_common.InitializePrecompileCall(d, evm, contract, readonly)
	if err != nil {
		return nil, err
	}

	var bz []byte
	switch method.Name {
	// queries
	case DistributionFunctionParams:
		bz, err = d.Params(ctx, evm, method, args)
	case DistributionFunctionValidatorDistributionInfo:
		bz, err = d.ValidatorDistributionInfo(ctx, evm, method, args)
	case DistributionFunctionValidatorOutstandingRewards:
		bz, err = d.ValidatorOutstankingRewards(ctx, evm, method, args)
	case DistributionFunctionValidatorCommission:
		bz, err = d.ValidatorCommission(ctx, evm, method, args)
	case DistributionFunctionValidatorSlashes:
		bz, err = d.ValidatorSlashes(ctx, evm, method, args)
	case DistributionFunctionDelegationRewards:
		bz, err = d.DelegationRewards(ctx, evm, method, args)
	case DistributionFunctionDelegationTotalRewards:
		bz, err = d.DelegationTotalRewards(ctx, evm, method, args)
	case DistributionFunctionDelegatorValidators:
		bz, err = d.DelegatorValidators(ctx, evm, method, args)
	case DistributionFunctionDelegatorWithdrawAddress:
		bz, err = d.DelegatorWithdrawAddress(ctx, evm, method, args)
	// txs
	case DistributionFunctionSetWithdrawAddress:
		bz, err = d.SetWithdrawAddress(ctx, evm, stateDB, contract, method, args)
	case DistributionFunctionWithdrawDelegatorRewards:
		bz, err = d.WithdrawDelegatorReward(ctx, evm, stateDB, contract, method, args)
	case DistributionFunctionWithdrawValidatorCommission:
		bz, err = d.WithdrawValidatorCommission(ctx, evm, stateDB, contract, method, args)
	}

	if err != nil {
		return nil, err
	}

	cost := ctx.GasMeter().GasConsumed() - initialGas

	if !contract.UseGas(cost) {
		return nil, vm.ErrOutOfGas
	}
	return bz, nil
}
