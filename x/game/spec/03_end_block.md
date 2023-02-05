<!--
order: 3
-->

# End-Block

Each abci end block call, unbonding queue is updated.

## Unbonding queue

In game staking has unbonding period and when a user begins unbonding of in game staking token,
`Unbonding` object is put on the queue and when the time pass, endblock automatically checks the
list of matured unbondings and deducts the staking amount of the user.

### Steps

1. Iterate matured unbondings and for each unbonding
2. Decrease `deposit.Staking` amount
3. Decrease `deposit.Unbonding` amount
4. Delete unbonding object
5. Emit `EventCompleteUnstakeInGameToken` event
