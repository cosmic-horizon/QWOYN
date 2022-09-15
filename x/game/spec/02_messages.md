<!--
order: 2
-->

# Messages

In this section we describe the processing of the game module messages and the corresponding updates to the state.

## MsgTransferModuleOwnership

`MsgTransferModuleOwnership` is a message to transfer module ownership by current module owner.

```protobuf
message MsgTransferModuleOwnership {
    string sender = 1;
    string newOwner = 2;
}
```

Steps:

1. Ensure `msg.Sender` is module owner
2. Update module owner to `msg.NewOwner`
3. Emit `EventTransferModuleOwnership` event

## MsgWhitelistNftContracts

`MsgWhitelistNftContracts` is a message to whitelist multiple nft contracts by module owner.

```protobuf
message MsgWhitelistNftContracts {
    string sender = 1;
    repeated string contracts = 2;
}
```

Steps:

1. Ensure `msg.Sender` is module owner
2. Iterate `msg.Contracts` and execute followings
3. Ensure contract minter is set to module owner
4. Ensure contract owner is set to module owner
5. Add contract to whitelisted
6. Emit `EventNftContractAddWhitelist` event

## MsgRemoveWhitelistedNftContracts

`MsgRemoveWhitelistedNftContracts` is a message to remove multiple nft contracts from whitelist by module owner.

```protobuf
message MsgRemoveWhitelistedNftContracts {
  string sender = 1;
  repeated string contracts = 2;
}
```

Steps:

1. Ensure `msg.Sender` is module owner
2. Iterate `msg.Contracts` and execute followings
3. Remove contract from whitelisted
4. Emit `EventNftContractRemoveWhitelist` event

## MsgDepositNft

`MsgDepositNft` is a message to deposit a whitelist nft by the nft owner.

```protobuf
message MsgDepositNft {
  string sender = 1;
  string contract = 2;
  uint64 token_id = 3;
}
```

Steps:

1. Ensure that nft contract (`msg.Contract`) is whitelisted
2. Execute wasm contract `"transfer_nft"` to send nft from `msg.Sender` to module address
3. Ensure wasm contract execution does not fail
4. Emit `EventDepositNft` event

## MsgWithdrawUpdatedNft

`MsgWithdrawUpdatedNft` is a message to withdraw modified nft from the module from module owner's signature.

```protobuf
message MsgWithdrawUpdatedNft {
  string sender = 1;
  string contract = 2;
  uint64 token_id = 3;
  string exec_msg = 4;
  bytes signature = 5;
}
```

Steps:

1. Get sign bytes to check signature of module owner
2. Verify `msg.Signature` with module owner pubKey and sign bytes
3. Execute nft updates and ensure that no error happens
4. Execute wasm contract `"transfer_nft"` to send nft from module address to `msg.Sender`.
5. Ensure wasm contract execution does not fail
6. Emit `EventWithdrawNft` event

### MsgSignerWithdrawUpdatedNft

`MsgSignerWithdrawUpdatedNft` is a message used for sign bytes for `MsgWithdrawUpdatedNft`.

```protobuf
message MsgSignerWithdrawUpdatedNft {
  string sender = 1;
  string contract = 2;
  uint64 token_id = 3;
  string exec_msg = 4;
}
```

## MsgDepositToken

`MsgDepositToken` is a message to deposit in game token into the game.

```protobuf
message MsgDepositToken {
  string sender = 1;
  cosmos.base.v1beta1.Coin amount = 2 [ (gogoproto.nullable) = false ];
}
```

Steps:

1. Ensure `msg.Amount.Denom` is deposit denom (game token)
2. Deposit coins from `msg.Sender` to module account
3. Increase deposit amount of `msg.Sender` by `msg.Amount.Amount`
4. Emit `EventDepositToken` event

## MsgWithdrawToken

`MsgWithdrawToken` is a message to withdraw in game token from the game.

```protobuf
message MsgWithdrawToken {
  string sender = 1;
  cosmos.base.v1beta1.Coin amount = 2 [(gogoproto.nullable) = false];
}
```

Steps:

1. Ensure `msg.Amount.Denom` is deposit denom (game token)
2. Withdrw coins from module account to `msg.Sender`
3. Decrease deposit amount of `msg.Sender` by `msg.Amount.Amount`
4. Emit `EventWithdrawToken` event

## MsgStakeInGameToken

`MsgStakeInGameToken` is a message to stake in-game token to get more game tokens.

```protobuf
message MsgStakeInGameToken {
  string sender = 1;
  cosmos.base.v1beta1.Coin amount = 2 [ (gogoproto.nullable) = false ];
}
```

Steps:

1. Ensure `msg.Amount.Denom` is deposit denom (game token)
2. Claim in game staking reward collected so far
3. Ensure that `staking + amount` does not exceed `deposit.Amount`
4. Increase staking amount
5. Emit `EventStakeInGameToken` event

## MsgBeginUnstakeInGameToken

`MsgBeginUnstakeInGameToken` is a message to begin unstaking game token. This starts unbonding process and it's automatically unbonded after `params.unstaking_time` pass.

```protobuf
message MsgBeginUnstakeInGameToken {
  string sender = 1;
  cosmos.base.v1beta1.Coin amount = 2 [ (gogoproto.nullable) = false ];
}
```

Steps:

1. Ensure `msg.Amount.Denom` is deposit denom (game token)
2. Claim in game staking reward collected so far
3. Ensure that `unbonding + amount` does not exceed `deposit.Unbonding`
4. Increase unbonding amount
5. Generate new unbonding id
6. Put unbonding object on unbonding queue with new unbonding id
7. Emit `EventBeginUnstakeInGameToken` event

## MsgClaimInGameStakingReward

`MsgClaimInGameStakingReward` is a message to claim in game staking reward by staker.

```protobuf
message MsgClaimInGameStakingReward {
  string sender = 1;
}
```

Steps:

1. Calculate reward amount from staking amount and `StakingInflation` parameter
2. If reward amount is positive, mint coins to module and increase `deposit.Amount`
3. Set reward claim time
4. Emit `EventClaimInGameStakingReward` event

## MsgAddLiquidity

`MsgAddLiquidity` is a message to add liquidity by module owner for game token and governance token.

```protobuf
message MsgAddLiquidity {
  string sender = 1;
  repeated cosmos.base.v1beta1.Coin amounts = 2 [ (gogoproto.nullable) = false ];
}
```

Steps:

1. Ensure that transaction is executed by module owner
2. Increase liquidity amount by `msg.Amounts`
3. Ensure that only two tokens are added as liquidity
4. Deposit coins from module owner account to the module
5. Emit `EventAddLiquidity` event

## MsgRemoveLiquidity

`MsgRemoveLiquidity` is a message to add liquidity by module owner for game token and governance token.

```protobuf
message MsgRemoveLiquidity {
  string sender = 1;
  repeated cosmos.base.v1beta1.Coin amounts = 2 [ (gogoproto.nullable) = false ];
}
```

Steps:

1. Ensure that transaction is executed by module owner
2. Decrease liquidity amount by `msg.Amounts`
3. Withdraw coins from module owner account to the module
4. Emit `EventRemoveLiquidity` event

## MsgSwap

`MsgSwap` is a message to swap on liquidity between game token and governance token.

```protobuf
message MsgSwap {
  string sender = 1;
  cosmos.base.v1beta1.Coin amount = 2 [ (gogoproto.nullable) = false ];
}
```

Steps:

1. Transfer input amount to the module account
2. Calculate swap out amount from current liquidity and input amount
3. Transfer calculated amount from module account to `msg.Sender`
4. Increase liquidity by deposit amount
5. Decrease liquidity by withdraw amount
6. Emit `EventSwap` event
