# nft module

- Multiple types of items ownable should be managed via this module, `Ships`, `Avatars` and `Planets`.
- Transferable items should be managed as an nft parameter. NFT parameters can be writeable
  - Cyrillium crystals (source of power for engines, shields and weapons)
  - Hardware (technological equipment)
  - Rations
- Ship parameters will be managed by this module

- List Of NFT Parameters - onchain/offchain.

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

- In game funds can be staked for staking rewards? Yes.  Sagans will be put in banks and the interest from storing your sagans in the banks will provide the staking rewards/inflation.

# nft marketplace

Users can interact with market to sell/buy nfts.

- auction
- nft listing

# nft upgrade

- character off-chain upgrade

# profile manager module

On-chain profile registration & ranking system for players

- rewards to top players weekly

# game state store module

Game state is always stored inside NFTs

Game state can be stored on-chain daily or weekly basis - could just store ipfs hash if the data is pretty big.

# game governance module

Subject would be related to game related stuff.
Sagan token holders have governance power on this module.

- game governance pool
- game pool spend proposal
- game parameters change proposal
- game parameters

# game competition (event) manager module

There could be funds allocated for competition.

## functions

- Organize competition
- Deposit on competition
- Start competition
- End competition
- Registration for participation
- Claim competition rewards

# governance module (native Cosmos SDK module)

On-chain governance with COHO token holders - would be interacting with technical subjects for chain maintenance.

# Cosmos SDK Groups Module

will be used for guild governance.

# token swap module

- COHO / Sagan
- Any other tokens that can be issued on game - not currently

This should be done with third parties like Osmosis? NO

In game trading module.  Private trading pair with liquidity provided by qwoyn studios. 

# game fee manager module

- Fee collection
- Fee distribution

# minter module

Inflation to

- validators
- purchasing sagans in private LP
- on-chain community pool

# airdrop module

- No airdrop, instead we will perform protocol to protocol token swaps

# lockup module

- Probably utilize vesting module? yes

# EVMOS smart contract

# Add-ons

- betting
- custom incentivization mechanism
- fee random airdrop
- summoning nft (a recipe to mint a random nft)
- nft staking
