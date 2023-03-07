<!--
order: 2
-->

# Messages

In this section we describe the processing of the stimulus module messages and the corresponding updates to the state.

## MsgDepositIntoOutpostFunding

`MsgDepositIntoOutpostFunding` is a message to deposit more funds to outposts funding pool by any user.

```protobuf
message MsgDepositIntoOutpostFunding {
    string sender = 1;
    cosmos.base.v1beta1.Coin amount = 2 [ (gogoproto.nullable) = false ];
}
```

Steps:

1. Ensure `msg.Amount.Denom` is game token denom (COHO)
2. Transfer tokens from account to outposts funding pool
3. Emit `EventDepositIntoOutpostFunding` event

## MsgWithdrawFromOutpostFunding

`MsgWithdrawFromOutpostFunding` is a message to withdraw from outposts funding pool by game module owner.

```protobuf
message MsgWithdrawFromOutpostFunding {
    string sender = 1;
    cosmos.base.v1beta1.Coin amount = 2 [ (gogoproto.nullable) = false ];
}
```

Steps:

1. Ensure `msg.Amount.Denom` is game token denom (COHO)
2. Ensure `msg.Sender` is game module owner
3. Transfer tokens from outposts funding pool to sender
4. Emit `EventWithdrawFromOutpostFunding` event
