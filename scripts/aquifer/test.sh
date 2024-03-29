#!/bin/sh

# start osmosis and qwoyn single nodes on different terminals
sh start_osmosis.sh
sh start.sh

# Restore keys to hermes relayer
echo "weather leader certain hard busy blouse click patient balcony return elephant hire mule gather danger curious visual boy estate army marine cinnamon snake flight" > ./relayer_mnemonic.txt
hermes --config ./hermes/config.toml keys delete --chain osmo-test --all
hermes --config ./hermes/config.toml keys add --chain osmo-test --mnemonic-file ./relayer_mnemonic.txt &
hermes --config ./hermes/config.toml keys delete --chain qwoyn-1 --all
hermes --config ./hermes/config.toml keys add --chain qwoyn-1 --mnemonic-file ./relayer_mnemonic.txt &

# create connection with and start (hermes 1.4.0)
# hermes --config ./hermes/config.toml create client --host-chain osmo-test --reference-chain qwoyn-1
# hermes --config ./hermes/config.toml create client --host-chain qwoyn-1 --reference-chain osmo-test
hermes --config ./hermes/config.toml create connection --a-chain osmo-test --b-chain qwoyn-1
hermes --config ./hermes/config.toml start

# create transfer channel
hermes --config ./hermes/config.toml create channel --a-chain osmo-test --a-port transfer --b-port transfer --a-connection connection-0

# check accounts to be managed on test
qwoynd keys show -a maintainer --keyring-backend=test
qwoyn1h9krsew6kpg9huzcqgmgmns0n48jx9yd5vr0n5
qwoynd keys show -a user1 --keyring-backend=test
qwoyn13tqzdukugulllnk3p5js3w7hzw8gclkeenzp6e
osmosisd keys show -a user1 --keyring-backend=test
osmo1h9krsew6kpg9huzcqgmgmns0n48jx9ydph3pz8

# check channels
qwoynd q ibc channel channels
- channel_id: channel-0
  connection_hops:
  - connection-1
  counterparty:
    channel_id: channel-0

# transfer tokens to QWOYN from Osmosis to buy QWOYN through aquifer module
osmosisd tx ibc-transfer transfer transfer channel-0 qwoyn13tqzdukugulllnk3p5js3w7hzw8gclkeenzp6e 1000000stake --fees=10000stake --chain-id=osmo-test --from=user1 --keyring-backend=test -y --broadcast-mode=block --node=http://localhost:16657
qwoynd query bank balances qwoyn13tqzdukugulllnk3p5js3w7hzw8gclkeenzp6e 
- amount: "1000000"
  denom: ibc/C053D637CCA2A2BA030E2C5EE1B28A16F71CCB0E45E8BE52766DC1B241B77878

IBC_OSMO=ibc/C053D637CCA2A2BA030E2C5EE1B28A16F71CCB0E45E8BE52766DC1B241B77878

# deposit tokens to be sold
qwoynd tx aquifer put-allocation-token 100000000uqwoyn --chain-id=qwoyn-1 --from=maintainer --keyring-backend=test -y --broadcast-mode=sync
qwoynd query tx 58EBEECD71BA897B61A02165C6B7420E9E1AA77E9CBC766816044B3FEC683E51

# set deposit end time by maintainer
qwoynd tx aquifer set-deposit-endtime $(($(date -u +%s) + 300)) --chain-id=qwoyn-1 --from=maintainer --keyring-backend=test -y --broadcast-mode=sync
# qwoynd tx aquifer buy-allocation-token 100000000stake --chain-id=qwoyn-1 --from=user1 --keyring-backend=test -y --broadcast-mode=sync

# buy allocation token by user
qwoynd tx aquifer buy-allocation-token 200000$IBC_OSMO --chain-id=qwoyn-1 --from=user1 --keyring-backend=test -y --broadcast-mode=sync

# init interchain account by maintainer
qwoynd tx aquifer init-ica connection-0 --chain-id=qwoyn-1 --from=maintainer --keyring-backend=test -y --broadcast-mode=sync --gas=1000000

qwoynd query intertx ica aquifer connection-0
interchain_account_address: osmo1q3vp0e6s3grtqap8ja25s3v0txqxufpfwq2w6ppqa08a00tfkrcsutdjr8

ICA_ACCOUNT=osmo1q3vp0e6s3grtqap8ja25s3v0txqxufpfwq2w6ppqa08a00tfkrcsutdjr8

qwoynd query aquifer params
# transfer required tokens to ICA account (it requires uosmo to pay fees for pool creation)
qwoynd tx aquifer exec-transfer channel-0 60000000000 --chain-id=qwoyn-1 --from=maintainer --keyring-backend=test -y --broadcast-mode=sync
osmosisd tx bank send user1 $ICA_ACCOUNT 1000000stake,10000000000uosmo --fees=10000stake --chain-id=osmo-test --keyring-backend=test -y --broadcast-mode=block --node=http://localhost:16657
qwoynd tx ibc-transfer transfer transfer channel-0 $ICA_ACCOUNT 100000000uqwoyn --chain-id=qwoyn-1 --from=user1 --keyring-backend=test -y --broadcast-mode=sync
osmosisd query bank balances $ICA_ACCOUNT --node=http://localhost:16657

# ensure that pool creation is allowed through IBC message on Osmosis
osmosisd query interchain-accounts host params --node=http://0.0.0.0:16657
# allow_messages:
# - /osmosis.gamm.poolmodels.balancer.v1beta1.MsgCreateBalancerPool
# host_enabled: true

# create pool on Osmosis
qwoynd tx aquifer exec-add-liquidity --pool-file="pool.json" --chain-id=qwoyn-1 --from=maintainer --keyring-backend=test -y --broadcast-mode=sync
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
# ensure gamm pool created after the operation
osmosisd query gamm pools --node=http://localhost:16657

# create osmosis pool as test without ICA
# osmosisd tx gamm create-pool --pool-file=pool1.json --from=user1 --keyring-backend=test -y --broadcast-mode=block --node=http://localhost:16657 --chain-id=osmo-test
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
