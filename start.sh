#!/bin/sh

rm -rf ~/.qwoynd

echo "silly slab oxygen reflect hawk wasp peace omit carbon pause turkey organ relax sing youth since fence increase record thing trial alien render begin" > validator.txt
echo "weather leader certain hard busy blouse click patient balcony return elephant hire mule gather danger curious visual boy estate army marine cinnamon snake flight" > mnemonic.txt;
echo "never chuckle bird almost jacket veteran weekend original rare habit point scorpion place gadget net train more plug upon pear renew mule material dynamic" > mnemonic2.txt;

# Build genesis
qwoynd init --chain-id=qwoyn-1 test
qwoynd keys add validator --keyring-backend="test" < validator.txt;
qwoynd keys add maintainer --recover --keyring-backend=test < mnemonic.txt;
qwoynd keys add user1 --recover --keyring-backend=test < mnemonic2.txt;

VALIDATOR=$(qwoynd keys show validator -a --keyring-backend="test")
MAINTAINER=$(qwoynd keys show maintainer -a --keyring-backend="test")
USER1=$(qwoynd keys show user1 -a --keyring-backend="test")
# VALIDATOR=qwoyn1hzqg4r2e789930hs88wqle25ef94xajuqay93r
# MAINTAINER=qwoyn1h9krsew6kpg9huzcqgmgmns0n48jx9yd5vr0n5
# USER1=qwoyn13tqzdukugulllnk3p5js3w7hzw8gclkeenzp6e
qwoynd genesis add-genesis-account $VALIDATOR 1000000000000uqwoyn,1000000000000ucoho,1000000000000stake
qwoynd genesis add-genesis-account $MAINTAINER 1000000000000uqwoyn,1000000000000ucoho,1000000000000stake
qwoynd genesis add-genesis-account $USER1 1000000000000uqwoyn,1000000000000ucoho,1000000000000stake
qwoynd genesis gentx validator 100000000stake --keyring-backend="test" --chain-id=qwoyn-1
qwoynd genesis collect-gentxs

# sed -i 's/stake/uqwoyn/g' $HOME/.qwoynd/config/genesis.json

# Start node
qwoynd start --pruning=nothing --minimum-gas-prices="0stake"
