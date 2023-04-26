# Events

The aquifer module emits the following events:

### MsgPutAllocationToken

| Type                                                | Attribute Key | Attribute Value                                    |
| --------------------------------------------------- | ------------- | -------------------------------------------------- |
| message                                             | action        | /cosmichorizon.qwoyn.aquifer.MsgPutAllocationToken |
| cosmichorizon.qwoyn.aquifer.EventPutAllocationToken | sender        | {sender}                                           |
| cosmichorizon.qwoyn.aquifer.EventPutAllocationToken | amount        | {amount}                                           |

## MsgTakeOutAllocationToken

| Type                                                    | Attribute Key | Attribute Value                                        |
| ------------------------------------------------------- | ------------- | ------------------------------------------------------ |
| message                                                 | action        | /cosmichorizon.qwoyn.aquifer.MsgTakeOutAllocationToken |
| cosmichorizon.qwoyn.aquifer.EventTakeOutAllocationToken | sender        | {sender}                                               |
| cosmichorizon.qwoyn.aquifer.EventTakeOutAllocationToken | amount        | {amount}                                               |

## MsgBuyAllocationToken

| Type                                                | Attribute Key | Attribute Value                                    |
| --------------------------------------------------- | ------------- | -------------------------------------------------- |
| message                                             | action        | /cosmichorizon.qwoyn.aquifer.MsgBuyAllocationToken |
| cosmichorizon.qwoyn.aquifer.EventBuyAllocationToken | sender        | {sender}                                           |
| cosmichorizon.qwoyn.aquifer.EventBuyAllocationToken | amount        | {amount}                                           |

## MsgSetDepositEndTime

| Type                                               | Attribute Key | Attribute Value                                   |
| -------------------------------------------------- | ------------- | ------------------------------------------------- |
| message                                            | action        | /cosmichorizon.qwoyn.aquifer.MsgSetDepositEndTime |
| cosmichorizon.qwoyn.aquifer.EventSetDepositEndTime | time          | {time}                                            |

## MsgInitICA

| Type    | Attribute Key | Attribute Value                         |
| ------- | ------------- | --------------------------------------- |
| message | action        | /cosmichorizon.qwoyn.aquifer.MsgInitICA |

## MsgExecTransfer

| Type    | Attribute Key | Attribute Value                              |
| ------- | ------------- | -------------------------------------------- |
| message | action        | /cosmichorizon.qwoyn.aquifer.MsgExecTransfer |

## MsgExecAddLiquidity

| Type    | Attribute Key | Attribute Value                                  |
| ------- | ------------- | ------------------------------------------------ |
| message | action        | /cosmichorizon.qwoyn.aquifer.MsgExecAddLiquidity |
