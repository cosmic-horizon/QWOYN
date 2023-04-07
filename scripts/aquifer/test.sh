#!/bin/sh

hermes -c ./hermes/config.toml keys restore qwoyn-1 -m "weather leader certain hard busy blouse click patient balcony return elephant hire mule gather danger curious visual boy estate army marine cinnamon snake flight"
hermes -c ./hermes/config.toml keys restore osmo-test -m "weather leader certain hard busy blouse click patient balcony return elephant hire mule gather danger curious visual boy estate army marine cinnamon snake flight"

hermes -c ./hermes/config.toml create connection osmo-test qwoyn-1
hermes -c ./hermes/config.toml start
hermes -c ./hermes/config.toml create channel osmo-test qwoyn-1 --port-a transfer --port-b transfer

qwoynd keys show -a user1 --keyring-backend=test
qwoyn1h9krsew6kpg9huzcqgmgmns0n48jx9yd5vr0n5
osmosisd keys show -a user1 --keyring-backend=test
osmo1h9krsew6kpg9huzcqgmgmns0n48jx9ydph3pz8

qwoynd q ibc channel channels
- channel_id: channel-0
  connection_hops:
  - connection-1
  counterparty:
    channel_id: channel-0

osmosisd tx ibc-transfer transfer transfer channel-0 qwoyn1h9krsew6kpg9huzcqgmgmns0n48jx9yd5vr0n5 1000000stake --chain-id=osmo-test --from=user1 --keyring-backend=test -y --broadcast-mode=block --node=http://localhost:16657
qwoynd query bank balances qwoyn1h9krsew6kpg9huzcqgmgmns0n48jx9yd5vr0n5 
- amount: "1000000"
  denom: ibc/C053D637CCA2A2BA030E2C5EE1B28A16F71CCB0E45E8BE52766DC1B241B77878

IBC_OSMO=ibc/C053D637CCA2A2BA030E2C5EE1B28A16F71CCB0E45E8BE52766DC1B241B77878
qwoynd tx aquifer put-allocation-token 100000000uqwoyn --chain-id=qwoyn-1 --from=user1 --keyring-backend=test -y --broadcast-mode=block
qwoynd tx aquifer set-deposit-endtime $(($(date -u +%s) + 300)) --chain-id=qwoyn-1 --from=user1 --keyring-backend=test -y --broadcast-mode=block
qwoynd tx aquifer buy-allocation-token 100000000stake --chain-id=qwoyn-1 --from=user1 --keyring-backend=test -y --broadcast-mode=block
qwoynd tx aquifer buy-allocation-token 200000$IBC_OSMO --chain-id=qwoyn-1 --from=user1 --keyring-backend=test -y --broadcast-mode=block
qwoynd tx aquifer init-ica connection-1 --chain-id=qwoyn-1 --from=user1 --keyring-backend=test -y --broadcast-mode=block

qwoynd query intertx ica aquifer connection-1
interchain_account_address: osmo12pfj79vt84pmxrwjc6sh4pawq8ltz028xprne80rgqzvxtj2tryq6zke9y

qwoynd query aquifer params
qwoynd tx aquifer exec-transfer channel-0 60000000000 --chain-id=qwoyn-1 --from=user1 --keyring-backend=test -y --broadcast-mode=block
osmosisd tx bank send user1 osmo12pfj79vt84pmxrwjc6sh4pawq8ltz028xprne80rgqzvxtj2tryq6zke9y 1000000stake,10000000000uosmo --chain-id=osmo-test --keyring-backend=test -y --broadcast-mode=block --node=http://localhost:16657
qwoynd tx ibc-transfer transfer transfer channel-0 osmo12pfj79vt84pmxrwjc6sh4pawq8ltz028xprne80rgqzvxtj2tryq6zke9y 100000000uqwoyn --chain-id=qwoyn-1 --from=user1 --keyring-backend=test -y --broadcast-mode=block
osmosisd query bank balances osmo12pfj79vt84pmxrwjc6sh4pawq8ltz028xprne80rgqzvxtj2tryq6zke9y --node=http://localhost:16657

qwoynd tx aquifer exec-add-liquidity --pool-file="pool.json" --chain-id=qwoyn-1 --from=user1 --keyring-backend=test -y --broadcast-mode=block
osmosisd query gamm pools --node=http://localhost:16657

osmosisd tx gamm create-pool --pool-file=pool1.json --from=user1 --keyring-backend=test -y --broadcast-mode=block --node=http://localhost:16657 --chain-id=osmo-test

osmosisd query interchain-accounts host params --node=http://0.0.0.0:16657
# allow_messages:
# - /osmosis.gamm.poolmodels.balancer.v1beta1.MsgCreateBalancerPool
# host_enabled: true

pool.json
```
{
    "weights": "5ibc/D67FFE08041F9BD3378C0003A785C90577F4DD2AED6713C78680D663FA9CAEE2,5stake",
    "initial-deposit": "1000ibc/D67FFE08041F9BD3378C0003A785C90577F4DD2AED6713C78680D663FA9CAEE2,1000stake",
    "swap-fee": "0.01",
    "exit-fee": "0.01",
    "future-governor": "168h"
}
```
pool1.json
```
{
    "weights": "5stake,5uosmo",
    "initial-deposit": "1000stake,1000uosmo",
    "swap-fee": "0.01",
    "exit-fee": "0.01",
    "future-governor": "168h"
}
```
