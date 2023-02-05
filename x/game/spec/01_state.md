<!--
order: 1
-->

# State

## Game tokens

### Deposit

`Deposit` is used for tracking deposited game tokens by each user. This object tracks deposit amount, staking amount, total unbonding amount by the user.

```protobuf
message Deposit {
  string address = 1;
  string amount = 2 [
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int",
    (gogoproto.nullable) = false
  ];
  string staking = 3 [
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int",
    (gogoproto.nullable) = false
  ];
  string unbonding = 4 [
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int",
    (gogoproto.nullable) = false
  ];
  google.protobuf.Timestamp reward_claim_time = 5 [ (gogoproto.nullable) = false, (gogoproto.stdtime) = true ];
}
```

- Deposit: `0x02 | format(address) -> Deposit`

### Unbonding

`Unbonding` is used for tracking unbonding entries that has unbonding period based on module parameter.

```protobuf
message Unbonding {
  uint64 id = 1;
  string staker_address = 2;
  int64 creation_height = 3;
  google.protobuf.Timestamp completion_time = 4
      [(gogoproto.nullable) = false, (gogoproto.stdtime) = true];
  string amount = 5 [
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int",
    (gogoproto.nullable)   = false
  ];
}
```

- Unbonding: `0x03 | format(unbonding_id) -> Unbonding`
- Unbonding by user: `0x04 | format(address) | format(unbonding_id) -> format(unbonding_id)`
- Unbonding by time: `0x05 | format(unbonding_completion_time) | format(unbonding_id) -> format(unbonding_id)`
- Last unbonding index: `0x06 | format(unbonding_id) -> format(last_unbonding_id)`

## Liquidity

`Liquidity` is available on game module to provide token swap between game token and native token.
Liquidity can only be added or removed by module owner.

```protobuf
message Liquidity {
  repeated cosmos.base.v1beta1.Coin amounts = 1 [ (gogoproto.nullable) = false ];
}
```

- Liquidity: `0x07 -> Liquidity`

## Game nfts

### Whitelisted contracts

Game item nft contracts are all registered by module owner before any deposit/withdraw operation.
Whitelisted contract's owner should be set to `game` module account to be registered.

- Whitelisted contract: `0x01 | contract_addr -> contract_addr`

## Params

Params is a module-wide configuration structure that stores system parameters that is governed by `gov` module.
Params contains `owner`(module owner), `deposit_denom` (game token), `staking_inflation` (in game staking %), `unstaking_time` (unstaking period for in game staking).

```protobuf
message Params {
  string owner = 1;
  string deposit_denom = 2;
  uint64 staking_inflation = 3; // percentage
  google.protobuf.Duration unstaking_time = 4 [(gogoproto.nullable) = false, (gogoproto.stdduration) = true];
}
```

- Params: `Paramsspace("game") -> Params`
