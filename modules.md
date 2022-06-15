# nft module

- Multiple types of items ownable should be managed via this module, `Ships` and `Planets` etc.
- Transferable items should be managed as an nft.

# In-game items / funds manager module

## module accounts

- `active_use` module account - game manager can use any items and funds on this module account. Everything is managed on off-chain for tokens and nfts put on `active_use` vault.
- `paused_use` module account - items to be stored when the user paused the game - items and funds on this account can not be transferred to `active_use` without end user's signature and end user also can not withdraw this to his end wallet .
- `starports` accounts: these accounts will get coho rewards, and it's switched to sagans

## functions

- deposit item
- deposit funds
- withdraw item (items are changed during the game execution)
- claim new item earned during the game
- withdraw funds

All of the withdraw functionalities should have signature from game manager account.
There could be multiple game manager accounts for security.

## thoughts

- In game funds can be staked for staking rewards?

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

## Cosmos SDK Groups Module

# token swap module

- COHO / Sagan
- Any other tokens that can be issued on game

This should be done with third parties like Osmosis?

# game fee manager module

- Fee collection
- Fee distribution

# minter module

Inflation to

- validators
- funding in-game Starports
- on-chain community pool

# airdrop module

- Rather than giving full airdrop as balance it would be good to make it claimable to make sure it's given to right people who claim

# lockup module

- Probably utilize vesting module?
- Or build custom lockup module?

# EVMOS smart contract

# Add-ons

- betting
- custom incentivization mechanism
- fee random airdrop
- summoning nft (a recipe to mint a random nft)
- nft staking
