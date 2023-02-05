# In-game items / funds manager module

## module accounts

- `active_use` module account
  - Permanantly linked to Avatar
  - Enabled and Disabled at specific locations ie. Cryo chamber, Shipyard, Bank
  - Avatar can use any items and funds on this module account.
  - Everything is managed off-chain for tokens and nfts put in `active_use` vault.
  - Planets only exist in `active_use` account
- `sales_use` module account
  - items for sale will be stored in this account
  - user cannot cancel sale
  - creates nft(deed) tied to items for sale
  - deed can then be transferred to end wallet
- `starports` accounts: these accounts will get coho rewards, and it's switched to sagans
- `retirement_account` interchain account which retires credits on behalf of the player

## functions

- deposit item
- deposit funds
- withdraw item (items are changed during the game execution)
- claim new item earned during the game
- withdraw funds

All of the withdraw functionalities should have signature from game manager account.
There could be multiple game manager for security.

## thoughts

- In game funds can be staked for staking rewards? Yes. Sagans will be put in banks and the interest from storing your sagans in the banks will provide the staking rewards/inflation.
