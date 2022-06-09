# nft module

- Multiple types of items ownable should be managed via this module, `Ships` and `Planets` etc.
- Transferable items should be managed as an nft.

# In-game items manager module

## module accounts

- `active_use` module account - game manager can use any items and funds on this module account. Everything is managed on off-chain for tokens and nfts put on `active_use` vault.
- `paused_use` module account - items to be stored when the user paused the game - items and funds on this account can not be transferred to `active_use` without end user's signature and end user also can not withdraw this to his end wallet .

## functions

- deposit item
- deposit funds
- withdraw item (items are changed during the game execution)
- claim new item earned during the game
- withdraw funds

All of the withdraw functionalities should have signature from game manager account.
There could be multiple game manager accounts for security.

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
