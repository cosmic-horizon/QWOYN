#!/bin/sh

rm -rf ~/.cohod

# Build genesis file incl account for passed address
cohod init --chain-id=test test
cohod keys add validator --keyring-backend="test"
cohod add-genesis-account $(cohod keys show validator -a --keyring-backend="test") 100000000000000ucoho
cohod gentx validator 100000000ucoho --keyring-backend="test" --chain-id=test
cohod collect-gentxs

# Start coho
cohod start --pruning=nothing
