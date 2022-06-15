# nft module

- Multiple types of items ownable should be managed via this module, `Ships`, `Avatars` and `Planets`.
- Transferable items should be managed as an nft parameter. NFT parameters can be writeable
    - Cyrillium crystals (source of power for engines, shields and weapons)
    - Hardware (technological equipment)
    - Rations
- Ship parameters will be managed by this module

# In-game items manager module

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

## functions

- deposit item
- deposit funds
- withdraw item (items are changed during the game execution)
- claim new item earned during the game
- withdraw funds

All of the withdraw functionalities should have signature from game manager account.
There could be multiple game manager for security.

# nft marketplace

Users can interact with market to sell/buy nfts.

- auction
- nft listing

# profile manager module

On-chain profile registration & ranking system for players

# game state store module

Game state can be stored on-chain daily or weekly basis - could just store ipfs hash if the data is pretty big.

# game governance module

Subject would be related to game related stuff.
Sagan token holders have governance power on this module.

- game governance pool
- game pool spend proposal
- game parameters change proposal
- game parameters

# game competition manager module

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

# token swap module

- COHO / Sagan
- Any other tokens that can be issued on game

This should be done with third parties like Osmosis?

# game fee manager module

- Fee collection
- Fee distribution
