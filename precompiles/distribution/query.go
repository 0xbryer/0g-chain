package distribution

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	distributionkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/vm"
)

func (d *DistributionPrecompile) Params(ctx sdk.Context, _ *vm.EVM, method *abi.Method, args []interface{}) ([]byte, error) {
	req, err := NewQueryParamsRequest(args)
	if err != nil {
		return nil, err
	}
	response, err := distributionkeeper.Querier{Keeper: d.distributionKeeper}.Params(ctx, req)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(convertParams(response.Params))
}

func (d *DistributionPrecompile) ValidatorDistributionInfo(ctx sdk.Context, _ *vm.EVM, method *abi.Method, args []interface{}) ([]byte, error) {
	req, err := NewQueryValidatorDistributionInfoRequest(args)
	if err != nil {
		return nil, err
	}
	response, err := distributionkeeper.Querier{Keeper: d.distributionKeeper}.ValidatorDistributionInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	res, err := convertValidatorDistributionInfoResponse(response)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(res)
}

func (d *DistributionPrecompile) ValidatorOutstankingRewards(ctx sdk.Context, _ *vm.EVM, method *abi.Method, args []interface{}) ([]byte, error) {
	req, err := NewQueryValidatorOutstankingRewardsRequest(args)
	if err != nil {
		return nil, err
	}

	response, err := distributionkeeper.Querier{Keeper: d.distributionKeeper}.ValidatorOutstandingRewards(ctx, req)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(convertValidatorOutstandingRewardsResponse(response))
}

func (d *DistributionPrecompile) ValidatorCommission(ctx sdk.Context, _ *vm.EVM, method *abi.Method, args []interface{}) ([]byte, error) {
	req, err := NewQueryValidatorCommissionRequest(args)
	if err != nil {
		return nil, err
	}

	response, err := distributionkeeper.Querier{Keeper: d.distributionKeeper}.ValidatorCommission(ctx, req)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(convertValidatorCommissionResponse(response))
}

func (d *DistributionPrecompile) ValidatorSlashes(ctx sdk.Context, _ *vm.EVM, method *abi.Method, args []interface{}) ([]byte, error) {
	req, err := NewQueryValidatorSlashesRequest(args)
	if err != nil {
		return nil, err
	}

	response, err := distributionkeeper.Querier{Keeper: d.distributionKeeper}.ValidatorSlashes(ctx, req)
	if err != nil {
		return nil, err
	}

	slashes := make([]ValidatorSlashEvent, len(response.Slashes))
	for i, slash := range response.Slashes {
		slashes[i] = convertValidatorSlashEvent(slash)
	}
	paginationResult := convertPageResponse(response.Pagination)

	return method.Outputs.Pack(slashes, paginationResult)
}

func (d *DistributionPrecompile) DelegationRewards(ctx sdk.Context, _ *vm.EVM, method *abi.Method, args []interface{}) ([]byte, error) {
	req, err := NewQueryDelegationRewardsRequest(args)
	if err != nil {
		return nil, err
	}

	response, err := distributionkeeper.Querier{Keeper: d.distributionKeeper}.DelegationRewards(ctx, req)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(convertDelegationRewardsResponse(response))
}

func (d *DistributionPrecompile) DelegationTotalRewards(ctx sdk.Context, _ *vm.EVM, method *abi.Method, args []interface{}) ([]byte, error) {
	req, err := NewQueryDelegationTotalRewardsRequest(args)
	if err != nil {
		return nil, err
	}

	response, err := distributionkeeper.Querier{Keeper: d.distributionKeeper}.DelegationTotalRewards(ctx, req)
	if err != nil {
		return nil, err
	}

	rewards, total, err := convertDelegationTotalRewardsResponse(response)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(rewards, total)
}

func (d *DistributionPrecompile) DelegatorValidators(ctx sdk.Context, _ *vm.EVM, method *abi.Method, args []interface{}) ([]byte, error) {
	req, err := NewQueryDelegatorValidatorsRequest(args)
	if err != nil {
		return nil, err
	}

	response, err := distributionkeeper.Querier{Keeper: d.distributionKeeper}.DelegatorValidators(ctx, req)
	if err != nil {
		return nil, err
	}

	validators, err := convertDelegatorValidatorsResponse(response)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(validators)
}

func (d *DistributionPrecompile) DelegatorWithdrawAddress(ctx sdk.Context, _ *vm.EVM, method *abi.Method, args []interface{}) ([]byte, error) {
	req, err := NewQueryDelegatorWithdrawAddressRequest(args)
	if err != nil {
		return nil, err
	}

	response, err := distributionkeeper.Querier{Keeper: d.distributionKeeper}.DelegatorWithdrawAddress(ctx, req)
	if err != nil {
		return nil, err
	}

	withdrawAddress, err := convertWithdrawAddress(response.WithdrawAddress)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(withdrawAddress)
}
