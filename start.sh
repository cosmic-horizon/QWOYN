#!/bin/sh

rm -rf ~/.qwoynd

echo "weather leader certain hard busy blouse click patient balcony return elephant hire mule gather danger curious visual boy estate army marine cinnamon snake flight" > mnemonic.txt;
echo "never chuckle bird almost jacket veteran weekend original rare habit point scorpion place gadget net train more plug upon pear renew mule material dynamic" > mnemonic2.txt;

# Build genesis
qwoynd init --chain-id=qwoyn-1 test
qwoynd keys add validator --keyring-backend="test"
qwoynd keys add maintainer --recover --keyring-backend=test < mnemonic.txt;
qwoynd keys add user1 --recover --keyring-backend=test < mnemonic2.txt;

qwoynd genesis add-genesis-account $(qwoynd keys show validator -a --keyring-backend="test") 1000000000000uqwoyn,1000000000000ucoho,1000000000000stake
qwoynd genesis add-genesis-account $(qwoynd keys show maintainer -a --keyring-backend="test") 1000000000000uqwoyn,1000000000000ucoho,1000000000000stake
qwoynd genesis add-genesis-account $(qwoynd keys show user1 -a --keyring-backend="test") 1000000000000uqwoyn,1000000000000ucoho,1000000000000stake
qwoynd genesis gentx validator 100000000stake --keyring-backend="test" --chain-id=qwoyn-1
qwoynd genesis collect-gentxs

# sed -i 's/stake/uqwoyn/g' $HOME/.qwoynd/config/genesis.json

# Start node
qwoynd start --pruning=nothing --minimum-gas-prices="0stake"
