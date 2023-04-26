# Messages

In this section we describe the processing of the aquifer module messages and the corresponding updates to the state.

## MsgPutAllocationToken

`MsgPutAllocationToken` is a message to put the token to be purchased on the pool.

```protobuf
message MsgPutAllocationToken {
  string sender = 1;
  cosmos.base.v1beta1.Coin amount = 2 [ (gogoproto.nullable) = false ];
}
```

Steps:

1. Ensure `msg.Amount.Denom` is correct allocation denom
2. Send tokens from `msg.Sender` to module account
3. Emit `EventPutAllocationToken` event

## MsgTakeOutAllocationToken

`MsgTakeOutAllocationToken` is a message to take out remaining tokens from the module after deposit end time.

```protobuf
message MsgTakeOutAllocationToken {
  string sender = 1;
  cosmos.base.v1beta1.Coin amount = 2 [ (gogoproto.nullable) = false ];
}
```

Steps:

1. Ensure `msg.Amount.Denom` is correct allocation denom
2. Ensure `msg.Sender` is module maintainer
3. Ensure deposit end time reached
4. Send tokens from module account to `msg.Sender`
5. Emit `EventTakeOutAllocationToken` event

## MsgBuyAllocationToken

`MsgBuyAllocationToken` is a message to allocate vesting for allocation token when the account put deposit token.

```protobuf
message MsgBuyAllocationToken {
  string sender = 1;
  cosmos.base.v1beta1.Coin amount = 2 [ (gogoproto.nullable) = false ];
}
```

Steps:

1. Ensure `msg.Amount.Denom` is correct deposit denom
2. Ensure deposit end time not reached
3. Send deposit token from `msg.Sender` to module account.
4. Ensure the message is sent from `BaseAccount` type
5. Allocate vesting based on initial liquidity, discount and vesting duration
6. Emit `EventBuyAllocationToken` event

## MsgSetDepositEndTime

`MsgSetDepositEndTime` is a message to set deposit end time by maintainer.

```protobuf
message MsgSetDepositEndTime {
  string sender = 1;
  uint64 end_time = 2;
}
```

Steps:

1. Ensure `msg.Sender` is maintainer
2. Set deposit end time
3. Emit `EventSetDepositEndTime` event

## MsgInitICA

`MsgInitICA` is a message to initiate ICA account for Osmosis liquidity bootstrap.

```protobuf
message MsgInitICA {
  string sender = 1;
  string connection_id = 2;
}
```

Steps:

1. Ensure `msg.Sender` is maintainer
2. Set connection id on params
3. Call `RegisterInterchainAccount` through ICA controller keeper as aquifer module as admin

## MsgExecTransfer

`MsgExecTransfer` is a message to IBC transfer deposit tokens to Osmosis ICA account.

```protobuf
message MsgExecTransfer {
  string sender =  1;
  uint64 timeout_nano_second = 2;
  string transfer_channel_id = 3;
}
```

Steps:

1. Ensure `msg.Sender` is maintainer
2. Get ICA address to receive deposit token
3. Execute IBC transfer with `msg.TransferChannelId`, module account deposit token balance and `msg.TimeoutNanoSecond`

## MsgExecAddLiquidity

`MsgExecAddLiquidity` is the message to create a new pool on Osmosis with ICA account

```protobuf
message MsgExecAddLiquidity {
  string sender = 1;
  bytes msg = 2 [
    (gogoproto.nullable) = false,
    (gogoproto.customtype) = "github.com/cosmic-horizon/qwoyn/osmosis/balancer.MsgCreateBalancerPool"
  ];
}
```

1. Ensure `msg.Sender` is maintainer
2. Get ICA address to execute pool creation
3. Compose ICA message for create pool
4. Execute ICA transaction through `SendTx` interface of ICA controller keeper
