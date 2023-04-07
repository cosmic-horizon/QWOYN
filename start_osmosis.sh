#!/bin/sh

# osmosis version v13.1.0
rm -rf $HOME/.osmosisd/

cd $HOME

chain_id=osmo-test

echo "weather leader certain hard busy blouse click patient balcony return elephant hire mule gather danger curious visual boy estate army marine cinnamon snake flight" > mnemonic.txt;

osmosisd init --chain-id=$chain_id osmo-test --home=$HOME/.osmosisd
osmosisd keys add validator --keyring-backend=test --home=$HOME/.osmosisd
osmosisd keys add user1 --recover --keyring-backend=test < mnemonic.txt;
osmosisd add-genesis-account $(osmosisd keys show validator -a --keyring-backend=test --home=$HOME/.osmosisd) 100000000000stake,100000000000uosmo --home=$HOME/.osmosisd
osmosisd add-genesis-account $(osmosisd keys show user1 -a --keyring-backend=test --home=$HOME/.osmosisd) 10000000stake,100000000000uosmo --home=$HOME/.osmosisd
osmosisd gentx validator 500000000stake --keyring-backend=test --home=$HOME/.osmosisd --chain-id=$chain_id
osmosisd collect-gentxs --home=$HOME/.osmosisd

sed -i -e 's#"localhost:6060"#"localhost:6061"#g' $HOME/.osmosisd/config/config.toml
sed -i -e 's#"tcp://0.0.0.0:26656"#"tcp://0.0.0.0:16656"#g' $HOME/.osmosisd/config/config.toml
sed -i -e 's#"tcp://127.0.0.1:26657"#"tcp://127.0.0.1:16657"#g' $HOME/.osmosisd/config/config.toml
sed -i -e 's#"0.0.0.0:9091"#"0.0.0.0:9093"#g' $HOME/.osmosisd/config/app.toml
sed -i -e 's#"0.0.0.0:9090"#"0.0.0.0:9092"#g' $HOME/.osmosisd/config/app.toml
sed -i -e 's#"0.0.0.0:8080"#"0.0.0.0:8081"#g' $HOME/.osmosisd/config/app.toml
sed -i -e 's#"0.0.0.0:1317"#"0.0.0.0:1316"#g' $HOME/.osmosisd/config/app.toml
sed -i -r "325s/.*/          \"allow_messages\": \[\"\/osmosis.gamm.poolmodels.balancer.v1beta1.MsgCreateBalancerPool\"\]/"  $HOME/.osmosisd/config/genesis.json
 
osmosisd start --home=$HOME/.osmosisd


