<!--
order: 4
-->

# Events

The game module emits the following events:

## EndBlocker

| Type                                                     | Attribute Key   | Attribute Value |
| -------------------------------------------------------- | --------------- | --------------- |
| cosmichorizon.qwoyn.game.EventCompleteUnstakeInGameToken | user            | {user}          |
| cosmichorizon.qwoyn.game.EventCompleteUnstakeInGameToken | amount          | {amount}        |
| cosmichorizon.qwoyn.game.EventCompleteUnstakeInGameToken | completion_time | {timestamp}     |
| cosmichorizon.qwoyn.game.EventCompleteUnstakeInGameToken | unbonding_id    | {unbonding_id}  |

## Msg's

### MsgTransferModuleOwnership

| Type                                                  | Attribute Key | Attribute Value                                      |
| ----------------------------------------------------- | ------------- | ---------------------------------------------------- |
| message                                               | action        | /cosmichorizon.qwoyn.game.MsgTransferModuleOwnership |
| cosmichorizon.qwoyn.game.EventTransferModuleOwnership | origin_owner  | {origin_owner}                                       |
| cosmichorizon.qwoyn.game.EventTransferModuleOwnership | new_owner     | {new_owner}                                          |

### MsgWhitelistNftContracts

| Type                                                  | Attribute Key | Attribute Value                                    |
| ----------------------------------------------------- | ------------- | -------------------------------------------------- |
| message                                               | action        | /cosmichorizon.qwoyn.game.MsgWhitelistNftContracts |
| cosmichorizon.qwoyn.game.EventNftContractAddWhitelist | contract      | {contract_addr}                                    |

### MsgRemoveWhitelistedNftContracts

| Type                                                     | Attribute Key | Attribute Value                                            |
| -------------------------------------------------------- | ------------- | ---------------------------------------------------------- |
| message                                                  | action        | /cosmichorizon.qwoyn.game.MsgRemoveWhitelistedNftContracts |
| cosmichorizon.qwoyn.game.EventNftContractRemoveWhitelist | contract      | {contract_addr}                                            |

### MsgDepositNft

| Type                                     | Attribute Key      | Attribute Value                         |
| ---------------------------------------- | ------------------ | --------------------------------------- |
| message                                  | action             | /cosmichorizon.qwoyn.game.MsgDepositNft |
| cosmichorizon.qwoyn.game.EventDepositNft | contract           | {contract_addr}                         |
| cosmichorizon.qwoyn.game.EventDepositNft | token_id           | {token_id}                              |
| cosmichorizon.qwoyn.game.EventDepositNft | owner              | {owner}                                 |
| execute                                  | \_contract_address | {contract_addr}                         |
| wasm                                     | token_id           | {token_id}                              |
| wasm                                     | recipient          | {recipient}                             |
| wasm                                     | sender             | {sender}                                |
| wasm                                     | action             | transfer_nft                            |
| wasm                                     | \_contract_address | contract_addr                           |

### MsgWithdrawUpdatedNft

| Type                                      | Attribute Key      | Attribute Value                                 |
| ----------------------------------------- | ------------------ | ----------------------------------------------- |
| message                                   | action             | /cosmichorizon.qwoyn.game.MsgWithdrawUpdatedNft |
| cosmichorizon.qwoyn.game.EventWithdrawNft | sender             | {sender}                                        |
| cosmichorizon.qwoyn.game.EventWithdrawNft | contract           | {contract}                                      |
| cosmichorizon.qwoyn.game.EventWithdrawNft | token_id           | {token_id}                                      |
| cosmichorizon.qwoyn.game.EventWithdrawNft | exec_msg           | {exec_msg}                                      |
| execute                                   | \_contract_address | {contract_addr}                                 |
| wasm                                      | token_id           | {token_id}                                      |
| wasm                                      | recipient          | {recipient}                                     |
| wasm                                      | sender             | {sender}                                        |
| wasm                                      | action             | transfer_nft                                    |
| wasm                                      | \_contract_address | {contract_addr}                                 |
| wasm                                      | token_id           | {token_id}                                      |
| wasm                                      | owner              | {owner}                                         |
| wasm                                      | minter             | {minter}                                        |
| wasm                                      | action             | mint                                            |
| wasm                                      | \_contract_address | {contract_addr}                                 |

### MsgDepositToken

| Type                                       | Attribute Key | Attribute Value                           |
| ------------------------------------------ | ------------- | ----------------------------------------- |
| message                                    | action        | /cosmichorizon.qwoyn.game.MsgDepositToken |
| message                                    | sender        | {sender}                                  |
| cosmichorizon.qwoyn.game.EventDepositToken | sender        | {sender}                                  |
| cosmichorizon.qwoyn.game.EventDepositToken | amount        | {amount}                                  |
| coin_spent                                 | spender       | {spender}                                 |
| coin_spent                                 | amount        | {amount}                                  |
| coin_received                              | receiver      | {receiver}                                |
| coin_received                              | amount        | {amount}                                  |
| transfer                                   | recipient     | {recipient}                               |
| transfer                                   | sender        | {sender}                                  |
| transfer                                   | amount        | {amount}                                  |

### MsgWithdrawToken

| Type                                        | Attribute Key | Attribute Value                           |
| ------------------------------------------- | ------------- | ----------------------------------------- |
| message                                     | action        | /cosmichorizon.qwoyn.game.MsgDepositToken |
| message                                     | sender        | {sender}                                  |
| cosmichorizon.qwoyn.game.EventWithdrawToken | sender        | {sender}                                  |
| cosmichorizon.qwoyn.game.EventWithdrawToken | amount        | {amount}                                  |
| coin_spent                                  | spender       | {spender}                                 |
| coin_spent                                  | amount        | {amount}                                  |
| coin_received                               | receiver      | {receiver}                                |
| coin_received                               | amount        | {amount}                                  |
| transfer                                    | recipient     | {recipient}                               |
| transfer                                    | sender        | {sender}                                  |
| transfer                                    | amount        | {amount}                                  |

### MsgStakeInGameToken

| Type                                                   | Attribute Key     | Attribute Value                               |
| ------------------------------------------------------ | ----------------- | --------------------------------------------- |
| message                                                | action            | /cosmichorizon.qwoyn.game.MsgStakeInGameToken |
| cosmichorizon.qwoyn.game.EventClaimInGameStakingReward | sender            | {sender}                                      |
| cosmichorizon.qwoyn.game.EventClaimInGameStakingReward | amount            | {amount}                                      |
| cosmichorizon.qwoyn.game.EventClaimInGameStakingReward | reward_claim_time | {reward_claim_time}                           |
| cosmichorizon.qwoyn.game.EventStakeInGameToken         | sender            | {sender}                                      |
| cosmichorizon.qwoyn.game.EventStakeInGameToken         | amount            | {amount}                                      |

### MsgBeginUnstakeInGameToken

| Type                                                   | Attribute Key     | Attribute Value                                      |
| ------------------------------------------------------ | ----------------- | ---------------------------------------------------- |
| message                                                | action            | /cosmichorizon.qwoyn.game.MsgBeginUnstakeInGameToken |
| cosmichorizon.qwoyn.game.EventClaimInGameStakingReward | sender            | {sender}                                             |
| cosmichorizon.qwoyn.game.EventClaimInGameStakingReward | amount            | {amount}                                             |
| cosmichorizon.qwoyn.game.EventClaimInGameStakingReward | reward_claim_time | {reward_claim_time}                                  |
| cosmichorizon.qwoyn.game.EventBeginUnstakeInGameToken  | sender            | {sender}                                             |
| cosmichorizon.qwoyn.game.EventBeginUnstakeInGameToken  | amount            | {amount}                                             |
| cosmichorizon.qwoyn.game.EventBeginUnstakeInGameToken  | completion_time   | {completion_time}                                    |

### MsgClaimInGameStakingReward

| Type                                                   | Attribute Key     | Attribute Value                                       |
| ------------------------------------------------------ | ----------------- | ----------------------------------------------------- |
| message                                                | action            | /cosmichorizon.qwoyn.game.MsgClaimInGameStakingReward |
| cosmichorizon.qwoyn.game.EventClaimInGameStakingReward | sender            | {sender}                                              |
| cosmichorizon.qwoyn.game.EventClaimInGameStakingReward | amount            | {amount}                                              |
| cosmichorizon.qwoyn.game.EventClaimInGameStakingReward | reward_claim_time | {reward_claim_time}                                   |

### MsgAddLiquidity

| Type                                       | Attribute Key | Attribute Value                           |
| ------------------------------------------ | ------------- | ----------------------------------------- |
| message                                    | action        | /cosmichorizon.qwoyn.game.MsgAddLiquidity |
| message                                    | sender        | {sender}                                  |
| cosmichorizon.qwoyn.game.EventAddLiquidity | sender        | {sender}                                  |
| cosmichorizon.qwoyn.game.EventAddLiquidity | amount        | {amount}                                  |
| coin_spent                                 | spender       | {spender}                                 |
| coin_spent                                 | amount        | {amount}                                  |
| coin_received                              | receiver      | {receiver}                                |
| coin_received                              | amount        | {amount}                                  |
| transfer                                   | recipient     | {recipient}                               |
| transfer                                   | sender        | {sender}                                  |
| transfer                                   | amount        | {amount}                                  |

### MsgRemoveLiquidity

| Type                                          | Attribute Key | Attribute Value                              |
| --------------------------------------------- | ------------- | -------------------------------------------- |
| message                                       | action        | /cosmichorizon.qwoyn.game.MsgRemoveLiquidity |
| message                                       | sender        | {sender}                                     |
| cosmichorizon.qwoyn.game.EventRemoveLiquidity | sender        | {sender}                                     |
| cosmichorizon.qwoyn.game.EventRemoveLiquidity | amount        | {amount}                                     |
| coin_spent                                    | spender       | {spender}                                    |
| coin_spent                                    | amount        | {amount}                                     |
| coin_received                                 | receiver      | {receiver}                                   |
| coin_received                                 | amount        | {amount}                                     |
| transfer                                      | recipient     | {recipient}                                  |
| transfer                                      | sender        | {sender}                                     |
| transfer                                      | amount        | {amount}                                     |

### MsgSwap

| Type                               | Attribute Key | Attribute Value                   |
| ---------------------------------- | ------------- | --------------------------------- |
| message                            | action        | /cosmichorizon.qwoyn.game.MsgSwap |
| message                            | sender        | {sender}                          |
| cosmichorizon.qwoyn.game.EventSwap | sender        | {sender}                          |
| cosmichorizon.qwoyn.game.EventSwap | in_amount     | {in_amount}                       |
| cosmichorizon.qwoyn.game.EventSwap | out_amount    | {out_amount}                      |
| []coin_spent                       | spender       | {spender}                         |
| []coin_spent                       | amount        | {amount}                          |
| []coin_received                    | receiver      | {receiver}                        |
| []coin_received                    | amount        | {amount}                          |
| []transfer                         | recipient     | {recipient}                       |
| []transfer                         | sender        | {sender}                          |
| []transfer                         | amount        | {amount}                          |
