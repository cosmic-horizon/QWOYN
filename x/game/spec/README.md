<!--
order: 0
title: Game Module Overview
parent:
  title: "game"
-->

# `game`

## Abstract

This module provides interface for game tokens and nfts deposit, withdrawal and token liquidity/swap operations.
Goal of the module is to manage in game items and tokens securely on-chain.

## Contents

1. **[State](01_state.md)**
   - [Game tokens](01_state.md#game-tokens)
   - [Liquidity](01_state.md#liquidity)
   - [Game nfts](01_state.md#params)
   - [Params](01_state.md#params)
2. **[Messages](02_messages.md)**
   - [MsgTransferModuleOwnership](02_messages.md#msgtransfermoduleownership)
   - [MsgWhitelistNftContracts](02_messages.md#msgwhitelistnftcontracts)
   - [MsgRemoveWhitelistedNftContracts](02_messages.md#msgremovewhitelistednftcontracts)
   - [MsgDepositNft](02_messages.md#msgdepositnft)
   - [MsgWithdrawUpdatedNft](02_messages.md#msgwithdrawupdatednft)
   - [MsgDepositToken](02_messages.md#msgdeposittoken)
   - [MsgWithdrawToken](02_messages.md#msgwithdrawtoken)
   - [MsgStakeInGameToken](02_messages.md#msgstakeingametoken)
   - [MsgBeginUnstakeInGameToken](02_messages.md#msgbeginunstakeingametoken)
   - [MsgClaimInGameStakingReward](02_messages.md#msgclaimingamestakingreward)
   - [MsgAddLiquidity](02_messages.md#msgaddliquidity)
   - [MsgRemoveLiquidity](02_messages.md#msgremoveliquidity)
   - [MsgSwap](02_messages.md#msgswap)
3. **[End-Block](03_end_block.md)**
   - [Unbonding queue](03_end_block.md#unbonding-queue)
4. **[Events](04_events.md)**
   - [EndBlocker](07_events.md#endblocker)
   - [Msg's](07_events.md#msg's)
5. **[Parameters](05_params.md)**
6. **[Client](06_client.md)**
