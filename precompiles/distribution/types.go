package distribution

import (
	"fmt"
	"math/big"

	precompiles_common "github.com/0glabs/0g-chain/precompiles/common"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/ethereum/go-ethereum/common"
)

// Coin is an auto generated low-level Go binding around an user-defined struct.
type Coin = struct {
	Denom  string   `json:"denom"`
	Amount *big.Int `json:"amount"`
}

// DecCoin is an auto generated low-level Go binding around an user-defined struct.
type DecCoin = struct {
	Denom  string   `json:"denom"`
	Amount *big.Int `json:"amount"`
}

// DelegationDelegatorReward is an auto generated low-level Go binding around an user-defined struct.
type DelegationDelegatorReward = struct {
	ValidatorAddress common.Address `json:"validatorAddress"`
	Reward           []DecCoin      `json:"reward"`
}

// PageRequest is an auto generated low-level Go binding around an user-defined struct.
type PageRequest = struct {
	Key        []byte `json:"key"`
	Offset     uint64 `json:"offset"`
	Limit      uint64 `json:"limit"`
	CountTotal bool   `json:"countTotal"`
	Reverse    bool   `json:"reverse"`
}

// PageResponse is an auto generated low-level Go binding around an user-defined struct.
type PageResponse = struct {
	NextKey []byte `json:"nextKey"`
	Total   uint64 `json:"total"`
}

// Params is an auto generated low-level Go binding around an user-defined struct.
type Params = struct {
	CommunityTax        *big.Int `json:"communityTax"`
	BaseProposerReward  *big.Int `json:"baseProposerReward"`
	BonusProposerReward *big.Int `json:"bonusProposerReward"`
	WithdrawAddrEnabled bool     `json:"withdrawAddrEnabled"`
}

// ValidatorDistributionInfoResponse is an auto generated low-level Go binding around an user-defined struct.
type ValidatorDistributionInfoResponse = struct {
	OperatorAddress common.Address `json:"operatorAddress"`
	SelfBondRewards []DecCoin      `json:"selfBondRewards"`
	Commission      []DecCoin      `json:"commission"`
}

// ValidatorSlashEvent is an auto generated low-level Go binding around an user-defined struct.
type ValidatorSlashEvent = struct {
	ValidatorPeriod uint64   `json:"validatorPeriod"`
	Fraction        *big.Int `json:"fraction"`
}

func NewQueryParamsRequest(args []interface{}) (*distributiontypes.QueryParamsRequest, error) {
	if len(args) != 0 {
		return nil, fmt.Errorf(precompiles_common.ErrInvalidNumberOfArgs, 0, len(args))
	}

	return &distributiontypes.QueryParamsRequest{}, nil
}

func convertParams(params distributiontypes.Params) Params {
	return Params{
		CommunityTax:        params.CommunityTax.BigInt(),
		BaseProposerReward:  params.BaseProposerReward.BigInt(),
		BonusProposerReward: params.BonusProposerReward.BigInt(),
		WithdrawAddrEnabled: params.WithdrawAddrEnabled,
	}
}

func NewQueryValidatorDistributionInfoRequest(args []interface{}) (*distributiontypes.QueryValidatorDistributionInfoRequest, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf(precompiles_common.ErrInvalidNumberOfArgs, 1, len(args))
	}

	validator := args[0].(common.Address)

	return &distributiontypes.QueryValidatorDistributionInfoRequest{
		ValidatorAddress: sdk.ValAddress(validator.Bytes()).String(),
	}, nil
}

func convertDecCoins(coins github_com_cosmos_cosmos_sdk_types.DecCoins) []DecCoin {
	ans := make([]DecCoin, len(coins))
	for i, coin := range coins {
		ans[i] = DecCoin{
			Denom:  coin.Denom,
			Amount: coin.Amount.TruncateInt().BigInt(),
		}
	}
	return ans
}

func convertValidatorDistributionInfoResponse(response *distributiontypes.QueryValidatorDistributionInfoResponse) (ValidatorDistributionInfoResponse, error) {
	operatorAddress, err := sdk.ValAddressFromBech32(response.OperatorAddress)
	if err != nil {
		return ValidatorDistributionInfoResponse{}, err
	}
	return ValidatorDistributionInfoResponse{
		OperatorAddress: common.BytesToAddress(operatorAddress.Bytes()),
		SelfBondRewards: convertDecCoins(response.SelfBondRewards),
		Commission:      convertDecCoins(response.Commission),
	}, nil
}

func NewQueryValidatorOutstankingRewardsRequest(args []interface{}) (*distributiontypes.QueryValidatorOutstandingRewardsRequest, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf(precompiles_common.ErrInvalidNumberOfArgs, 1, len(args))
	}

	validator := args[0].(common.Address)

	return &distributiontypes.QueryValidatorOutstandingRewardsRequest{
		ValidatorAddress: sdk.ValAddress(validator.Bytes()).String(),
	}, nil
}

func convertValidatorOutstandingRewardsResponse(response *distributiontypes.QueryValidatorOutstandingRewardsResponse) []DecCoin {
	return convertDecCoins(response.Rewards.Rewards)
}

func NewQueryValidatorCommissionRequest(args []interface{}) (*distributiontypes.QueryValidatorCommissionRequest, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf(precompiles_common.ErrInvalidNumberOfArgs, 1, len(args))
	}

	validator := args[0].(common.Address)

	return &distributiontypes.QueryValidatorCommissionRequest{
		ValidatorAddress: sdk.ValAddress(validator.Bytes()).String(),
	}, nil
}

func convertValidatorCommissionResponse(response *distributiontypes.QueryValidatorCommissionResponse) []DecCoin {
	return convertDecCoins(response.Commission.Commission)
}

func convertQueryPageRequest(pagination PageRequest) *query.PageRequest {
	return &query.PageRequest{
		Key:        pagination.Key,
		Offset:     pagination.Offset,
		Limit:      pagination.Limit,
		CountTotal: pagination.CountTotal,
		Reverse:    pagination.Reverse,
	}
}

func NewQueryValidatorSlashesRequest(args []interface{}) (*distributiontypes.QueryValidatorSlashesRequest, error) {
	if len(args) != 4 {
		return nil, fmt.Errorf(precompiles_common.ErrInvalidNumberOfArgs, 4, len(args))
	}

	validator := args[0].(common.Address)

	return &distributiontypes.QueryValidatorSlashesRequest{
		ValidatorAddress: sdk.ValAddress(validator.Bytes()).String(),
		StartingHeight:   args[1].(uint64),
		EndingHeight:     args[2].(uint64),
		Pagination:       convertQueryPageRequest(args[3].(PageRequest)),
	}, nil
}

func convertValidatorSlashEvent(slash distributiontypes.ValidatorSlashEvent) ValidatorSlashEvent {
	return ValidatorSlashEvent{
		ValidatorPeriod: slash.ValidatorPeriod,
		Fraction:        slash.Fraction.BigInt(),
	}
}

func convertPageResponse(pagination *query.PageResponse) PageResponse {
	if pagination == nil {
		return PageResponse{
			NextKey: make([]byte, 0),
			Total:   1,
		}
	}
	return PageResponse{
		NextKey: pagination.NextKey,
		Total:   pagination.Total,
	}
}

func NewQueryDelegationRewardsRequest(args []interface{}) (*distributiontypes.QueryDelegationRewardsRequest, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf(precompiles_common.ErrInvalidNumberOfArgs, 2, len(args))
	}

	delegator := args[0].(common.Address)
	validator := args[1].(common.Address)

	return &distributiontypes.QueryDelegationRewardsRequest{
		DelegatorAddress: sdk.AccAddress(delegator.Bytes()).String(),
		ValidatorAddress: sdk.ValAddress(validator.Bytes()).String(),
	}, nil
}

func convertDelegationRewardsResponse(response *distributiontypes.QueryDelegationRewardsResponse) []DecCoin {
	return convertDecCoins(response.Rewards)
}

func NewQueryDelegationTotalRewardsRequest(args []interface{}) (*distributiontypes.QueryDelegationTotalRewardsRequest, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf(precompiles_common.ErrInvalidNumberOfArgs, 1, len(args))
	}

	delegator := args[0].(common.Address)

	return &distributiontypes.QueryDelegationTotalRewardsRequest{
		DelegatorAddress: sdk.AccAddress(delegator.Bytes()).String(),
	}, nil
}

func convertDelegationDelegatorReward(reward distributiontypes.DelegationDelegatorReward) (DelegationDelegatorReward, error) {
	validatorAddress, err := sdk.ValAddressFromBech32(reward.ValidatorAddress)
	if err != nil {
		return DelegationDelegatorReward{}, err
	}
	return DelegationDelegatorReward{
		ValidatorAddress: common.BytesToAddress(validatorAddress.Bytes()),
		Reward:           convertDecCoins(reward.Reward),
	}, nil
}

func convertDelegationTotalRewardsResponse(response *distributiontypes.QueryDelegationTotalRewardsResponse) ([]DelegationDelegatorReward, []DecCoin, error) {
	rewards := make([]DelegationDelegatorReward, len(response.Rewards))
	var err error
	for i, reward := range response.Rewards {
		rewards[i], err = convertDelegationDelegatorReward(reward)
		if err != nil {
			return nil, nil, err
		}
	}
	return rewards, convertDecCoins(response.Total), nil
}

func NewQueryDelegatorValidatorsRequest(args []interface{}) (*distributiontypes.QueryDelegatorValidatorsRequest, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf(precompiles_common.ErrInvalidNumberOfArgs, 1, len(args))
	}

	delegator := args[0].(common.Address)

	return &distributiontypes.QueryDelegatorValidatorsRequest{
		DelegatorAddress: sdk.AccAddress(delegator.Bytes()).String(),
	}, nil
}

func convertDelegatorValidatorsResponse(response *distributiontypes.QueryDelegatorValidatorsResponse) ([]common.Address, error) {
	validators := make([]common.Address, len(response.Validators))
	for i, addr := range response.Validators {
		validatorAddress, err := sdk.ValAddressFromBech32(addr)
		if err != nil {
			return nil, err
		}
		validators[i] = common.BytesToAddress(validatorAddress.Bytes())
	}
	return validators, nil
}

func NewQueryDelegatorWithdrawAddressRequest(args []interface{}) (*distributiontypes.QueryDelegatorWithdrawAddressRequest, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf(precompiles_common.ErrInvalidNumberOfArgs, 1, len(args))
	}

	delegator := args[0].(common.Address)

	return &distributiontypes.QueryDelegatorWithdrawAddressRequest{
		DelegatorAddress: sdk.AccAddress(delegator.Bytes()).String(),
	}, nil
}

func convertWithdrawAddress(addr string) (common.Address, error) {
	withdrawAddress, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return common.Address{}, err
	}
	return common.Address(withdrawAddress.Bytes()), nil
}

func NewQueryCommunityPoolRequest(args []interface{}) (*distributiontypes.QueryCommunityPoolRequest, error) {
	if len(args) != 0 {
		return nil, fmt.Errorf(precompiles_common.ErrInvalidNumberOfArgs, 0, len(args))
	}

	return &distributiontypes.QueryCommunityPoolRequest{}, nil
}

func convertCommunityPoolResponse(response *distributiontypes.QueryCommunityPoolResponse) []DecCoin {
	return convertDecCoins(response.Pool)
}
