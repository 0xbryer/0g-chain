// SPDX-License-Identifier: LGPL-3.0-only
pragma solidity >=0.8.0;

import "./types.sol";

struct Params {
    uint communityTax;
    uint baseProposerReward;
    uint bonusProposerReward;
    bool withdrawAddrEnabled;
}

struct ValidatorDistributionInfoResponse {
    address operatorAddress;
    DecCoin[] selfBondRewards;
    DecCoin[] commission;
}

struct ValidatorSlashEvent {
    uint64 validatorPeriod;
    uint fraction; // 18 decimals
}

struct DelegationDelegatorReward {
    address validatorAddress;
    DecCoin[] reward;
}

interface IDistribution {
    /*=== events ===*/

    /*=== cosmos tx ===*/

    /**
     * @dev SetWithdrawAddress defines a method to change the withdraw address for a delegator (or validator self-delegation).
     * cosmos grpc: rpc SetWithdrawAddress(MsgSetWithdrawAddress) returns (MsgSetWithdrawAddressResponse);
     */
    function setWithdrawAddress(
        address delegatorAddress,
        address withdrawerAddress
    ) external returns (bool success);

    /**
     * @dev WithdrawDelegatorReward defines a method to withdraw rewards of delegator from a single validator.
     * cosmos grpc: rpc WithdrawDelegatorReward(MsgWithdrawDelegatorReward) returns (MsgWithdrawDelegatorRewardResponse);
     */
    function withdrawDelegatorRewards(
        address delegatorAddress,
        string memory validatorAddress
    ) external returns (Coin[] calldata amount);

    /**
     * @dev WithdrawValidatorCommission defines a method to withdraw the full commission to the validator address.
     * cosmos grpc: rpc WithdrawValidatorCommission(MsgWithdrawValidatorCommission) returns (MsgWithdrawValidatorCommissionResponse);
     */
    function withdrawValidatorCommission(
        address validatorAddress
    ) external returns (Coin[] calldata amount);

    /**
     * @dev FundCommunityPool defines a method to allow an account to directly fund the community pool.
     * cosmos grpc: rpc FundCommunityPool(MsgFundCommunityPool) returns (MsgFundCommunityPoolResponse);
     */
    function fundCommunityPool(
        address depositor,
        uint256 amount
    ) external returns (bool success);

    /*=== cosmos query ===*/

    /**
     * @dev Params queries params of the distribution module.
     * cosmos grpc: rpc Params(QueryParamsRequest) returns (QueryParamsResponse);
     */
    function params() external view returns (Params memory);

    /**
     * @dev ValidatorDistributionInfo queries validator commission and self-delegation rewards for validator
     * cosmos grpc: rpc ValidatorDistributionInfo(QueryValidatorDistributionInfoRequest) returns (QueryValidatorDistributionInfoResponse);
     */
    function validatorDistributionInfo(
        address validatorAddress
    )
        external
        view
        returns (ValidatorDistributionInfoResponse memory distributionInfo);

    /**
     * @dev ValidatorOutstandingRewards queries rewards of a validator address.
     * cosmos grpc: rpc ValidatorOutstandingRewards(QueryValidatorOutstandingRewardsRequest) returns (QueryValidatorOutstandingRewardsResponse);
     */
    function validatorOutstandingRewards(
        address validatorAddress
    ) external view returns (DecCoin[] memory rewards);

    /**
     * @dev ValidatorCommission queries accumulated commission for a validator.
     * cosmos grpc: rpc ValidatorCommission(QueryValidatorCommissionRequest) returns (QueryValidatorCommissionResponse);
     */
    function validatorCommission(
        string memory validatorAddress
    ) external view returns (DecCoin[] memory commission);

    /**
     * @dev ValidatorSlashes queries slash events of a validator.
     * cosmos grpc: rpc ValidatorSlashes(QueryValidatorSlashesRequest) returns (QueryValidatorSlashesResponse);
     */
    function validatorSlashes(
        address validatorAddress,
        uint64 startingHeight,
        uint64 endingHeight,
        PageRequest calldata pageRequest
    )
        external
        view
        returns (
            ValidatorSlashEvent[] memory slashes,
            PageResponse calldata pageResponse
        );

    /**
     * @dev DelegationRewards queries the total rewards accrued by a delegation.
     * cosmos grpc: rpc DelegationRewards(QueryDelegationRewardsRequest) returns (QueryDelegationRewardsResponse);
     */
    function delegationRewards(
        address delegatorAddress,
        address validatorAddress
    ) external view returns (DecCoin[] memory rewards);

    /**
     * @dev DelegationTotalRewards queries the total rewards accrued by a each validator.
     * cosmos grpc: rpc DelegationTotalRewards(QueryDelegationTotalRewardsRequest) returns (QueryDelegationTotalRewardsResponse);
     */
    function delegationTotalRewards(
        address delegatorAddress
    )
        external
        view
        returns (
            DelegationDelegatorReward[] memory rewards,
            DecCoin[] memory total
        );

    /**
     * @dev DelegatorValidators queries the validators of a delegator.
     * cosmos grpc: rpc DelegatorValidators(QueryDelegatorValidatorsRequest) returns (QueryDelegatorValidatorsResponse);
     */
    function delegatorValidators(
        address delegatorAddress
    ) external view returns (address[] memory validators);

    /**
     * @dev DelegatorWithdrawAddress queries withdraw address of a delegator.
     * cosmos grpc: rpc DelegatorWithdrawAddress(QueryDelegatorWithdrawAddressRequest) returns (QueryDelegatorWithdrawAddressResponse);
     */
    function delegatorWithdrawAddress(
        address delegatorAddress
    ) external view returns (address withdrawAddress);

    /**
     * @dev CommunityPool queries the community pool coins.
     * cosmos grpc: rpc CommunityPool(QueryCommunityPoolRequest) returns (QueryCommunityPoolResponse);
     */
    function CommunityPool() external view returns (DecCoin[] memory pool);
}
