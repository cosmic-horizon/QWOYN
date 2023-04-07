#!/bin/sh

rm -rf ~/.qwoynd

echo "weather leader certain hard busy blouse click patient balcony return elephant hire mule gather danger curious visual boy estate army marine cinnamon snake flight" > mnemonic.txt;

# Build genesis
qwoynd init --chain-id=qwoyn-1 test
qwoynd keys add validator --keyring-backend="test"
qwoynd keys add user1 --recover --keyring-backend=test < mnemonic.txt;

qwoynd add-genesis-account $(qwoynd keys show validator -a --keyring-backend="test") 1000000000000uqwoyn,1000000000000ucoho,1000000000000stake
qwoynd add-genesis-account $(qwoynd keys show user1 -a --keyring-backend="test") 1000000000000uqwoyn,1000000000000ucoho,1000000000000stake
qwoynd gentx validator 100000000stake --keyring-backend="test" --chain-id=qwoyn-1
qwoynd collect-gentxs

# sed -i 's/stake/uqwoyn/g' $HOME/.qwoynd/config/genesis.json

# Start node
qwoynd start --pruning=nothing --minimum-gas-prices="0stake"
