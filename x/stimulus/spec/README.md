# Stimulus module

It's what will be used to supply the starports with $COHO.

Requirements for the module include.

1. Shave off some of the inflationary rewards. This means that there needs to be a mechanism that takes a portion of the QWOYN inflationary rewards and purchases COHO with those rewards via the LP Daily.
2. These newly purchased COHO will then go into the outposts funding pool for purchasing Assets from players.
3. Also, when players purchase items from the outposts the module needs to take a portion of that COHO and deposit it back into the outposts funding pool. Another percentage goes to Qwoyn Studios and another portion will be burned.

As far as the COHO token is concerned. Do you remember how we are minting that token? My idea is that it is only acquired through 2 mechanisms. 1) Selling commodities using the outposts and 2) interest in the a) planetary banks and 3) universal bank. The interest should be based on the amount of tokens a player has deposited in each type of bank. So lets say, for arguments sake, a player deposits 1000 COHO in the universal bank then they will start earning a small percentage of COHO tokens, similar to staking rewards or interest in a bank account.
So we need a minting mechanism for COHO based on the banking system. We will also need an initial supply of COHO tokens which a portion will be deposited into the ouposts funding pool and a portion will be put into the LP.

## Questions:

- How to configure each starport on stimulus (is it needed?)
- Should add epoch module for daily swap? (Or is it okay to do per block?)
- Minter rewards are all given to fee collector, configure minter module to distribute part of inflation to stimulus module?
- Outposts funding pool should be controlled by game module owner account?
- When game players purchase items, outposts funding pool will need to be increased by game module owner account? or individual account? (Probably game module owner account)
- Should add bank mechanism (Planetary banks, universal bank) for users to run bank on the game? (Game module owner deposit? or Individual account deposit?)
- Game module owner should have permission to mint COHO tokens for selling commodities?
- How to determine the interest rate based on total deposit amount on the bank?
- Initial portion of COHO tokens should be minted into outpost funding pool at genesis, remaining put on LP, others given to COHO token manager

## Answers:

- How to configure each starport on stimulus (is it needed?)

We will not need to configure the starports individually. I figure we have one big pool that the game can have access to and the game will then configure the starports and access the tokens as needed.

- Should add epoch module for daily swap? (Or is it okay to do per block?)

Per block is just fine.

- Minter rewards are all given to fee collector, configure minter module to distribute part of inflation to stimulus module?

Yes please configure the minter module to distribute part of the inflation. The module must then swap them as in the spec above.

- Outposts funding pool should be controlled by game module owner account?

yes

- When game players purchase items, outposts funding pool will need to be increased by game module owner account? or individual account? (Probably game module owner account)

The game module owner account should increase the pool each time there is a swap.

- Should add bank mechanism (Planetary banks, universal bank) for users to run bank on the game? (Game module owner deposit? or Individual account deposit?)

Planetary banks and universal banks are used by the players only. Universal banks should have a deposit limit and a smaller interest(inflation) rate than the planetary banks. The planetary banks do not have a deposit limit and have a highrer interest rate than the universal bank.

- Game module owner should have permission to mint COHO tokens for selling commodities?

Yes

- How to determine the interest rate based on total deposit amount on the bank?

The interest rate should be a fixed daily rate. Something like 2.5% for example. If the user has 1000 COHO in the bank then the after 24 hours or x amount of blocks per day the user will then have 1025 COHO.

- Initial portion of COHO tokens should be minted into outpost funding pool at genesis, remaining put on LP, others given to COHO token manager

That’s correct.

## Apply swap fee for LP swap COHO <> QWOYN

configure the fee to be paid in Qwoyn (static fee. e.g. 1 QWOYN)

- Where to send fees to?
