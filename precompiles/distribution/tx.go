package distribution

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	distributionkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/ethermint/x/evm/statedb"
)

func (d *DistributionPrecompile) SetWithdrawAddress(
	ctx sdk.Context,
	evm *vm.EVM,
	stateDB *statedb.StateDB,
	contract *vm.Contract,
	method *abi.Method,
	args []interface{},
) ([]byte, error) {
	msg, err := NewMsgSetWithdrawAddress(args, contract.CallerAddress)
	if err != nil {
		return nil, err
	}
	// validation
	// execute
	_, err = distributionkeeper.NewMsgServerImpl(d.distributionKeeper).SetWithdrawAddress(ctx, msg)
	if err != nil {
		return nil, err
	}
	// emit events
	return method.Outputs.Pack()
}

func (d *DistributionPrecompile) WithdrawDelegatorReward(
	ctx sdk.Context,
	evm *vm.EVM,
	stateDB *statedb.StateDB,
	contract *vm.Contract,
	method *abi.Method,
	args []interface{},
) ([]byte, error) {
	msg, err := NewMsgWithdrawDelegatorReward(args, contract.CallerAddress)
	if err != nil {
		return nil, err
	}
	// validation
	// execute
	_, err = distributionkeeper.NewMsgServerImpl(d.distributionKeeper).WithdrawDelegatorReward(ctx, msg)
	if err != nil {
		return nil, err
	}
	// emit events
	return method.Outputs.Pack()
}

func (d *DistributionPrecompile) WithdrawValidatorCommission(
	ctx sdk.Context,
	evm *vm.EVM,
	stateDB *statedb.StateDB,
	contract *vm.Contract,
	method *abi.Method,
	args []interface{},
) ([]byte, error) {
	msg, err := NewMsgWithdrawValidatorCommission(args, contract.CallerAddress)
	if err != nil {
		return nil, err
	}
	// validation
	// execute
	_, err = distributionkeeper.NewMsgServerImpl(d.distributionKeeper).WithdrawValidatorCommission(ctx, msg)
	if err != nil {
		return nil, err
	}
	// emit events
	return method.Outputs.Pack()
}
