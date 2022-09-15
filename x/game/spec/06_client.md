<!--
order: 6
-->

# Client

## CLI

A user can query and interact with the `game` module using the CLI.

### Query

The `query` commands allows users to query `game` module state.

```bash
cohod query game --help
```

#### params

The `params` command allows users to query values set as game module parameters.

```bash
# usage
cohod query game params [flags]

# example
cohod query game params
```

Example Output:

```bash
params:
  deposit_denom: stake
  owner: coho1x0fha27pejg5ajg8vnrqm33ck8tq6raafkwa9v
  staking_inflation: "1"
  unstaking_time: 30s
```

#### whitelisted contracts

The `whitelisted-contracts` command allows users to query whitelisted nft contracts.

Usage:

```bash
# usage
cohod query game whitelisted-contracts [flags]

# example
cohod query game whitelisted-contracts
```

Example Output:

```bash
contracts:
- coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc
```

#### all deposit balances

The `all-deposit-balances` command allow users to query all deposits.

```bash
# usage
cohod query game all-deposit-balances [flags]

# example
cohod query game all-deposit-balances
```

Example Output:

```bash
amount:
  amount: "1000000"
  denom: stake
deposits:
- address: coho1tpug77hddm558x39gzndx72w94zgfstzk95rdy
  amount: "1000000"
  reward_claim_time: "0001-01-01T00:00:00Z"
  staking: "0"
  unbonding: "0"
```

#### deposit balance

The `deposit-balance` command allow users to query a single deposit.

```bash
# usage
cohod query game deposit-balance [address] [flags]

# example
cohod query game deposit-balance coho1tpug77hddm558x39gzndx72w94zgfstzk95rdy
```

Example Output:

```bash
deposit:
  address: coho1tpug77hddm558x39gzndx72w94zgfstzk95rdy
  amount: "1000000"
  reward_claim_time: "0001-01-01T00:00:00Z"
  staking: "0"
  unbonding: "0"
```

#### all unbondings

The `all-unbondings` command allow users to query all active unbondings.

```bash
# usage
cohod query game all-unbondings [flags]
# example
cohod query game all-unbondings
```

Example Output:

```bash
unbondings:
- amount: "100000"
  completion_time: "2022-09-13T13:07:14.682492Z"
  creation_height: "178"
  id: "1"
  staker_address: coho1tpug77hddm558x39gzndx72w94zgfstzk95rdy
```

#### user unbondings

The `user-unbondings` command allow users to query all active unbondings for an address.

```bash
# usage
cohod query game user-unbondings [address] [flags]
# example
cohod query game user-unbondings $(cohod keys show -a validator --keyring-backend=test)
```

Example Output:

```bash
unbondings:
- amount: "100000"
  completion_time: "2022-09-13T13:16:18.650478Z"
  creation_height: "283"
  id: "2"
  staker_address: coho1tpug77hddm558x39gzndx72w94zgfstzk95rdy
```

#### liquidity

The `liquidity` command allow users to query total liquidity put for coho and qwoyn.

```bash
# usage
cohod query game liquidity [flags]
# example
cohod query game liquidity
```

Example Output:

```bash
liquidity:
  amounts:
  - amount: "1000000"
    denom: ucoho
  - amount: "1000000"
    denom: uqwoyn
```

#### estimate swap out

The `estimated-swap-out` command allow users to query estimated out amount on swap operation.

```bash
# bash
cohod query game estimated-swap-out [amount] [flags]
# example
cohod query game estimated-swap-out 10000ucoho
```

Example Output:

```bash
amount:
  amount: "9901"
  denom: uqwoyn
```

#### swap rate

The `swap-rate` command allow users to query spot price on current liquidity.

```bash
# usage
cohod query game swap-rate [flags]
# example
cohod query game swap-rate
```

Example Output:

```bash
rate: "1.000000000000000000"
src_denom: ucoho
tar_denom: uqwoyn
```

### Transactions

The `tx` commands allows users to interact with the `game` module.

```bash
cohod tx game --help
```

#### MsgTransferModuleOwnership

The command `transfer-module-ownership` allows the module owner to transfer the ownership to a different address.

Usage:

```bash
cohod tx game transfer-module-ownership [newOwner] [flags]
```

Example:

```bash
cohod tx game transfer-module-ownership coho1tpug77hddm558x39gzndx72w94zgfstzk95rdy \
  --from=moduleOwner
```

## MsgWhitelistNftContracts

The command `whitelist-contracts` allows the module owner to whitelist nft contracts deployed on cosmwasm.

Usage:

```bash
cohod tx game whitelist-contracts [contracts] [flags]
```

Example:

```bash
cohod tx game whitelist-contracts coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc \
  --from=moduleOwner
```

## MsgRemoveWhitelistedNftContracts

The command `remove-whitelisted-contracts` allows the module owner to remove whitelisted nft contracts.

Usage:

```bash
cohod tx game remove-whitelisted-contracts [contracts] [flags]
```

Example:

```bash
cohod tx game remove-whitelisted-contracts coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc \
  --from=moduleOwner
```

## MsgDepositNft

The command `deposit-nft` allows a user to deposit a whitelisted nft.

Usage:

```bash
cohod tx game deposit-nft [contract] [tokenId] [flags]
```

Example:

```bash
cohod tx game deposit-nft coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc 1 \
  --from=user
```

## MsgSignerWithdrawUpdatedNft

The command `sign-withdraw-updated-nft` allows the module owner to generate signature for a user to withdraw an nft with updates.

Usage:

```bash
cohod tx game sign-withdraw-updated-nft [contract] [tokenId] [execMsg] [flags]
```

Example:

```bash
cohod tx game sign-withdraw-updated-nft coho14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9snm4thc 1 '{"update_nft":{"token_id":"1","extension":{"ship_type":67,"owner":"200"}}}' \
  --from=moduleOwner
```

## MsgWithdrawUpdatedNft

The command `withdraw-updated-nft` allows a user to withdraw a deposited nft with updates by using module owner signature.

Usage:

```bash
cohod tx game withdraw-updated-nft [contract] [tokenId] [execMsg] [signature] [flags]
```

Example:

```bash
cohod tx game withdraw-updated-nft '{"update_nft":{"token_id":"1","extension":{"ship_type":67,"owner":"200"}}}' \
 42d6e9d3b62ffc9b0bc3f6a97cbc0857af1c7a7aa57549571d7bc72415a955d978a1790440ce53c8f9fbfa2ce70d967812eda6094d6f112d7e5736170e48e2a8 \
  --from=user
```

## MsgDepositToken

The command `deposit-token` allows a user to deposit game tokens by users.

Usage:

```bash
cohod tx game deposit-token deposit-token [coin] [flags]
```

Example:

```bash
cohod tx game deposit-token 1000000stake \
  --from=user
```

## MsgWithdrawToken

The command `withdraw-token` allows a user to withdraw game tokens by users.

Usage:

```bash
cohod tx game withdraw-token [coin] [flags]
```

Example:

```bash
cohod tx game withdraw-token 1000000stake \
  --from=user
```

## MsgStakeInGameToken

The command `stake-ingame-token` allows a user to stake deposited game tokens by users.

Usage:

```bash
cohod tx game stake-ingame-token [coin] [flags]
```

Example:

```bash
cohod tx game stake-ingame-token 1000000stake \
  --from=user
```

## MsgBeginUnstakeInGameToken

The command `begin-unstake-ingame-token` allows a user to begin unstake of staked game tokens by users.

Usage:

```bash
cohod tx game begin-unstake-ingame-token [coin] [flags]
```

Example:

```bash
cohod tx game begin-unstake-ingame-token 1000000stake \
  --from=user
```

## MsgClaimInGameStakingReward

The command `claim-ingame-staking-reward` allows a user to claim rewards from staked game tokens.

Usage:

```bash
cohod tx game claim-ingame-staking-reward [flags]
```

Example:

```bash
cohod tx game claim-ingame-staking-reward \
  --from=user
```

## MsgAddLiquidity

The command `add-liquidity` allows the module owner to put liquidity for game tokens and governance tokens.

Usage:

```bash
cohod tx game add-liquidity [coins] [flags]
```

Example:

```bash
cohod tx game add-liquidity 10000ucoho,10000uqwoyn \
  --from=user
```

## MsgRemoveLiquidity

The command `remove-liquidity` allows the module owner to remove liquidity.

Usage:

```bash
cohod tx game remove-liquidity [coins] [flags]
```

Example:

```bash
cohod tx game remove-liquidity 10000ucoho,10000uqwoyn \
  --from=user
```

## MsgSwap

The command `swap` allows a user to to swap on liquidity.

Usage:

```bash
cohod tx game swap [coin] [flags]
```

Example:

```bash
cohod tx game swap 10000ucoho \
  --from=user
```
