# Messages

In this section we describe the processing of the intertx module messages and the corresponding updates to the state.

## MsgRegisterAccount

`MsgRegisterAccount` is a message to register an interchain account.

```protobuf
message MsgRegisterAccount {
  // owner is the address of the interchain account owner.
  string owner = 1;

  // connection_id is the connection id string (i.e. channel-5).
  string connection_id = 2;

  // version is the application version string. For example, this could be an
  // ICS27 encoded metadata type or an ICS29 encoded metadata type with a nested
  // application version.
  string version = 3;
}
```

This registers an interchain controller account where the owner is `msg.Owner` and the host is one connected `msg.ConnectionId`.

## MsgSubmitTx

`MsgSubmitTx` is a message to broadcast ICA transaction.

```protobuf
message MsgSubmitTx {
  // owner is the owner address of the interchain account.
  string owner = 1;

  // connection_id is the id of the connection.
  string connection_id = 2;

  // msg is the bytes of the transaction msg to send.
  google.protobuf.Any msg = 3;
}
```

Steps:

1. Get portID from `msg.Owner`
2. Get channelId from `portID` and `msg.ConnectiondId`
3. Serialize `msg.Msg` into ICA packet data
4. Broadcast packet through ica controller keeper
