<!--
order: 4
-->

# Events

The game module emits the following events:

## EndBlocker

| Type                                                    | Attribute Key   | Attribute Value |
| ------------------------------------------------------- | --------------- | --------------- |
| cosmichorizon.coho.game.EventCompleteUnstakeInGameToken | user            | {user}          |
| cosmichorizon.coho.game.EventCompleteUnstakeInGameToken | amount          | {amount}        |
| cosmichorizon.coho.game.EventCompleteUnstakeInGameToken | completion_time | {timestamp}     |
| cosmichorizon.coho.game.EventCompleteUnstakeInGameToken | unbonding_id    | {unbonding_id}  |

## Msg's

### MsgTransferModuleOwnership

| Type                                                 | Attribute Key | Attribute Value                                     |
| ---------------------------------------------------- | ------------- | --------------------------------------------------- |
| message                                              | action        | /cosmichorizon.coho.game.MsgTransferModuleOwnership |
| cosmichorizon.coho.game.EventTransferModuleOwnership | origin_owner  | {origin_owner}                                      |
| cosmichorizon.coho.game.EventTransferModuleOwnership | new_owner     | {new_owner}                                         |

### MsgWhitelistNftContracts

| Type                                                 | Attribute Key | Attribute Value                                   |
| ---------------------------------------------------- | ------------- | ------------------------------------------------- |
| message                                              | action        | /cosmichorizon.coho.game.MsgWhitelistNftContracts |
| cosmichorizon.coho.game.EventNftContractAddWhitelist | contract      | {contract_addr}                                   |

### MsgRemoveWhitelistedNftContracts

| Type                                                    | Attribute Key | Attribute Value                                           |
| ------------------------------------------------------- | ------------- | --------------------------------------------------------- |
| message                                                 | action        | /cosmichorizon.coho.game.MsgRemoveWhitelistedNftContracts |
| cosmichorizon.coho.game.EventNftContractRemoveWhitelist | contract      | {contract_addr}                                           |

### MsgDepositNft

| Type                                    | Attribute Key      | Attribute Value                        |
| --------------------------------------- | ------------------ | -------------------------------------- |
| message                                 | action             | /cosmichorizon.coho.game.MsgDepositNft |
| cosmichorizon.coho.game.EventDepositNft | contract           | {contract_addr}                        |
| cosmichorizon.coho.game.EventDepositNft | token_id           | {token_id}                             |
| cosmichorizon.coho.game.EventDepositNft | owner              | {owner}                                |
| execute                                 | \_contract_address | {contract_addr}                        |
| wasm                                    | token_id           | {token_id}                             |
| wasm                                    | recipient          | {recipient}                            |
| wasm                                    | sender             | {sender}                               |
| wasm                                    | action             | transfer_nft                           |
| wasm                                    | \_contract_address | contract_addr                          |

### MsgWithdrawUpdatedNft

| Type                                     | Attribute Key      | Attribute Value                                |
| ---------------------------------------- | ------------------ | ---------------------------------------------- |
| message                                  | action             | /cosmichorizon.coho.game.MsgWithdrawUpdatedNft |
| cosmichorizon.coho.game.EventWithdrawNft | sender             | {sender}                                       |
| cosmichorizon.coho.game.EventWithdrawNft | contract           | {contract}                                     |
| cosmichorizon.coho.game.EventWithdrawNft | token_id           | {token_id}                                     |
| cosmichorizon.coho.game.EventWithdrawNft | exec_msg           | {exec_msg}                                     |
| execute                                  | \_contract_address | {contract_addr}                                |
| wasm                                     | token_id           | {token_id}                                     |
| wasm                                     | recipient          | {recipient}                                    |
| wasm                                     | sender             | {sender}                                       |
| wasm                                     | action             | transfer_nft                                   |
| wasm                                     | \_contract_address | {contract_addr}                                |
| wasm                                     | token_id           | {token_id}                                     |
| wasm                                     | owner              | {owner}                                        |
| wasm                                     | minter             | {minter}                                       |
| wasm                                     | action             | mint                                           |
| wasm                                     | \_contract_address | {contract_addr}                                |

### MsgDepositToken

| Type                                      | Attribute Key | Attribute Value                          |
| ----------------------------------------- | ------------- | ---------------------------------------- |
| message                                   | action        | /cosmichorizon.coho.game.MsgDepositToken |
| message                                   | sender        | {sender}                                 |
| cosmichorizon.coho.game.EventDepositToken | sender        | {sender}                                 |
| cosmichorizon.coho.game.EventDepositToken | amount        | {amount}                                 |
| coin_spent                                | spender       | {spender}                                |
| coin_spent                                | amount        | {amount}                                 |
| coin_received                             | receiver      | {receiver}                               |
| coin_received                             | amount        | {amount}                                 |
| transfer                                  | recipient     | {recipient}                              |
| transfer                                  | sender        | {sender}                                 |
| transfer                                  | amount        | {amount}                                 |

### MsgWithdrawToken

| Type                                       | Attribute Key | Attribute Value                          |
| ------------------------------------------ | ------------- | ---------------------------------------- |
| message                                    | action        | /cosmichorizon.coho.game.MsgDepositToken |
| message                                    | sender        | {sender}                                 |
| cosmichorizon.coho.game.EventWithdrawToken | sender        | {sender}                                 |
| cosmichorizon.coho.game.EventWithdrawToken | amount        | {amount}                                 |
| coin_spent                                 | spender       | {spender}                                |
| coin_spent                                 | amount        | {amount}                                 |
| coin_received                              | receiver      | {receiver}                               |
| coin_received                              | amount        | {amount}                                 |
| transfer                                   | recipient     | {recipient}                              |
| transfer                                   | sender        | {sender}                                 |
| transfer                                   | amount        | {amount}                                 |

### MsgStakeInGameToken

| Type                                                  | Attribute Key     | Attribute Value                              |
| ----------------------------------------------------- | ----------------- | -------------------------------------------- |
| message                                               | action            | /cosmichorizon.coho.game.MsgStakeInGameToken |
| cosmichorizon.coho.game.EventClaimInGameStakingReward | sender            | {sender}                                     |
| cosmichorizon.coho.game.EventClaimInGameStakingReward | amount            | {amount}                                     |
| cosmichorizon.coho.game.EventClaimInGameStakingReward | reward_claim_time | {reward_claim_time}                          |
| cosmichorizon.coho.game.EventStakeInGameToken         | sender            | {sender}                                     |
| cosmichorizon.coho.game.EventStakeInGameToken         | amount            | {amount}                                     |

### MsgBeginUnstakeInGameToken

| Type                                                  | Attribute Key     | Attribute Value                                     |
| ----------------------------------------------------- | ----------------- | --------------------------------------------------- |
| message                                               | action            | /cosmichorizon.coho.game.MsgBeginUnstakeInGameToken |
| cosmichorizon.coho.game.EventClaimInGameStakingReward | sender            | {sender}                                            |
| cosmichorizon.coho.game.EventClaimInGameStakingReward | amount            | {amount}                                            |
| cosmichorizon.coho.game.EventClaimInGameStakingReward | reward_claim_time | {reward_claim_time}                                 |
| cosmichorizon.coho.game.EventBeginUnstakeInGameToken  | sender            | {sender}                                            |
| cosmichorizon.coho.game.EventBeginUnstakeInGameToken  | amount            | {amount}                                            |
| cosmichorizon.coho.game.EventBeginUnstakeInGameToken  | completion_time   | {completion_time}                                   |

### MsgClaimInGameStakingReward

| Type                                                  | Attribute Key     | Attribute Value                                      |
| ----------------------------------------------------- | ----------------- | ---------------------------------------------------- |
| message                                               | action            | /cosmichorizon.coho.game.MsgClaimInGameStakingReward |
| cosmichorizon.coho.game.EventClaimInGameStakingReward | sender            | {sender}                                             |
| cosmichorizon.coho.game.EventClaimInGameStakingReward | amount            | {amount}                                             |
| cosmichorizon.coho.game.EventClaimInGameStakingReward | reward_claim_time | {reward_claim_time}                                  |

### MsgAddLiquidity

| Type                                      | Attribute Key | Attribute Value                          |
| ----------------------------------------- | ------------- | ---------------------------------------- |
| message                                   | action        | /cosmichorizon.coho.game.MsgAddLiquidity |
| message                                   | sender        | {sender}                                 |
| cosmichorizon.coho.game.EventAddLiquidity | sender        | {sender}                                 |
| cosmichorizon.coho.game.EventAddLiquidity | amount        | {amount}                                 |
| coin_spent                                | spender       | {spender}                                |
| coin_spent                                | amount        | {amount}                                 |
| coin_received                             | receiver      | {receiver}                               |
| coin_received                             | amount        | {amount}                                 |
| transfer                                  | recipient     | {recipient}                              |
| transfer                                  | sender        | {sender}                                 |
| transfer                                  | amount        | {amount}                                 |

### MsgRemoveLiquidity

| Type                                         | Attribute Key | Attribute Value                             |
| -------------------------------------------- | ------------- | ------------------------------------------- |
| message                                      | action        | /cosmichorizon.coho.game.MsgRemoveLiquidity |
| message                                      | sender        | {sender}                                    |
| cosmichorizon.coho.game.EventRemoveLiquidity | sender        | {sender}                                    |
| cosmichorizon.coho.game.EventRemoveLiquidity | amount        | {amount}                                    |
| coin_spent                                   | spender       | {spender}                                   |
| coin_spent                                   | amount        | {amount}                                    |
| coin_received                                | receiver      | {receiver}                                  |
| coin_received                                | amount        | {amount}                                    |
| transfer                                     | recipient     | {recipient}                                 |
| transfer                                     | sender        | {sender}                                    |
| transfer                                     | amount        | {amount}                                    |

### MsgSwap

| Type                              | Attribute Key | Attribute Value                  |
| --------------------------------- | ------------- | -------------------------------- |
| message                           | action        | /cosmichorizon.coho.game.MsgSwap |
| message                           | sender        | {sender}                         |
| cosmichorizon.coho.game.EventSwap | sender        | {sender}                         |
| cosmichorizon.coho.game.EventSwap | in_amount     | {in_amount}                      |
| cosmichorizon.coho.game.EventSwap | out_amount    | {out_amount}                     |
| []coin_spent                      | spender       | {spender}                        |
| []coin_spent                      | amount        | {amount}                         |
| []coin_received                   | receiver      | {receiver}                       |
| []coin_received                   | amount        | {amount}                         |
| []transfer                        | recipient     | {recipient}                      |
| []transfer                        | sender        | {sender}                         |
| []transfer                        | amount        | {amount}                         |
