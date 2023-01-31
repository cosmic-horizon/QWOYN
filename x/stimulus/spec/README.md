# Stimulus module

It's what will be used to supply the starports with $COHO.

Requirements for the module include.

1. Shave off some of the inflationary rewards. This means that there needs to be a mechanism that takes a portion of the QWOYN inflationary rewards and purchases COHO with those rewards via the LP Daily.
2. These newly purchased COHO will then go into the outposts funding pool for purchasing Assets from players.
3. Also, when players purchase items from the outposts the module needs to take a portion of that COHO and deposit it back into the outposts funding pool. Another percentage goes to Qwoyn Studios and another portion will be burned.

As far as the COHO token is concerned. Do you remember how we are minting that token? My idea is that it is only acquired through 2 mechanisms. 1) Selling commodities using the outposts and 2) interest in the a) planetary banks and 3) universal bank. The interest should be based on the amount of tokens a player has deposited in each type of bank. So lets say, for arguments sake, a player deposits 1000 COHO in the universal bank then they will start earning a small percentage of COHO tokens, similar to staking rewards or interest in a bank account.
So we need a minting mechanism for COHO based on the banking system. We will also need an initial supply of COHO tokens which a portion will be deposited into the ouposts funding pool and a portion will be put into the LP.

Questions

- How to configure each starport on stimulus (is it needed?)
- Should add epoch module for daily swap? (Or is it okay to do per block?)
- Minter rewards are all given to fee collector, configure minter module to distribute part of inflation to stimulus module?
- Outposts funding pool should be controlled by game module owner account?
- When game players purchase items, outposts funding pool will need to be increased by game module owner account? or individual account? (Probably game module owner account)
- Should add bank mechanism (Planetary banks, universal bank) for users to run bank on the game? (Game module owner deposit? or Individual account deposit?)
- Game module owner should have permission to mint COHO tokens for selling commodities?
- How to determine the interest rate based on total deposit amount on the bank?
- Initial portion of COHO tokens should be minted into outpost funding pool at genesis, remaining put on LP, others given to COHO token manager

Shave off some of the inflationary rewards. This means that there needs to be a mechanism that takes a portion of the QWOYN inflationary rewards and purchases COHO with those rewards via the LP Daily.
