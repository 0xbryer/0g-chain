// SPDX-License-Identifier: LGPL-3.0-only
pragma solidity >=0.8.0;

/**
 * @dev cosmos.base.v1beta1.Coin
 */
struct Coin {
    string denom;
    uint amount;
}

/**
 * @dev cosmos.base.v1beta1.DecCoin
 * 18 decimals
 */
struct DecCoin {
    string denom;
    uint amount;
}

/**
 * @dev cosmos.base.query.v1beta1.PageRequest
 */
struct PageRequest {
    bytes key;
    uint64 offset;
    uint64 limit;
    bool countTotal;
    bool reverse;
}

/**
 * @dev cosmos.base.query.v1beta1.PageResponse
 */
struct PageResponse {
    bytes nextKey;
    uint64 total;
}
