<!--
order: 4
-->

# Events

The stimulus module emits the following events:

## BeginBlocker

| Type                               | Attribute Key | Attribute Value |
| ---------------------------------- | ------------- | --------------- |
| cosmichorizon.qwoyn.game.EventSwap | sender        | {sender}        |
| cosmichorizon.qwoyn.game.EventSwap | in_amount     | {in_amount}     |
| cosmichorizon.qwoyn.game.EventSwap | out_amount    | {out_amount}    |
| []coin_spent                       | spender       | {spender}       |
| []coin_spent                       | amount        | {amount}        |
| []coin_received                    | receiver      | {receiver}      |
| []coin_received                    | amount        | {amount}        |
| []transfer                         | recipient     | {recipient}     |
| []transfer                         | sender        | {sender}        |
| []transfer                         | amount        | {amount}        |

## Msg's

### MsgDepositIntoOutpostFunding

| Type                                                        | Attribute Key | Attribute Value                                            |
| ----------------------------------------------------------- | ------------- | ---------------------------------------------------------- |
| message                                                     | action        | /cosmichorizon.qwoyn.stimulus.MsgDepositIntoOutpostFunding |
| cosmichorizon.qwoyn.stimulus.EventDepositIntoOutpostFunding | sender        | {sender}                                                   |
| cosmichorizon.qwoyn.stimulus.EventDepositIntoOutpostFunding | amount        | {amount}                                                   |

### MsgWithdrawFromOutpostFunding

| Type                                                         | Attribute Key | Attribute Value                                         |
| ------------------------------------------------------------ | ------------- | ------------------------------------------------------- |
| message                                                      | action        | /cosmichorizon.qwoyn.game.MsgWithdrawFromOutpostFunding |
| cosmichorizon.qwoyn.stimulus.EventWithdrawFromOutpostFunding | sender        | {sender}                                                |
| cosmichorizon.qwoyn.stimulus.EventWithdrawFromOutpostFunding | amount        | {amount}                                                |
