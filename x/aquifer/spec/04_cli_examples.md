# CLI examples for aquifer

```sh
qwoynd tx aquifer put-allocation-token 100000000uqwoyn --chain-id=qwoyn-1 --from=user1 --keyring-backend=test -y --broadcast-mode=block
qwoynd tx aquifer set-deposit-endtime $(($(date -u +%s) + 300)) --chain-id=qwoyn-1 --from=user1 --keyring-backend=test -y --broadcast-mode=block
qwoynd tx aquifer buy-allocation-token 100000000stake --chain-id=qwoyn-1 --from=user1 --keyring-backend=test -y --broadcast-mode=block
qwoynd tx aquifer buy-allocation-token 200000$IBC_OSMO --chain-id=qwoyn-1 --from=user1 --keyring-backend=test -y --broadcast-mode=block
qwoynd tx aquifer init-ica connection-1 --chain-id=qwoyn-1 --from=user1 --keyring-backend=test -y --broadcast-mode=block

qwoynd query intertx ica aquifer connection-1
qwoynd query aquifer params
qwoynd tx aquifer exec-transfer channel-0 60000000000 --chain-id=qwoyn-1 --from=user1 --keyring-backend=test -y --broadcast-mode=block
qwoynd tx aquifer exec-add-liquidity --pool-file="pool.json" --chain-id=qwoyn-1 --from=user1 --keyring-backend=test -y --broadcast-mode=block
osmosisd query gamm pools --node=http://localhost:16657
```

pool.json

```json
{
  "weights": "5ibc/D67FFE08041F9BD3378C0003A785C90577F4DD2AED6713C78680D663FA9CAEE2,5stake",
  "initial-deposit": "1000ibc/D67FFE08041F9BD3378C0003A785C90577F4DD2AED6713C78680D663FA9CAEE2,1000stake",
  "swap-fee": "0.01",
  "exit-fee": "0.01",
  "future-governor": "168h"
}
```
