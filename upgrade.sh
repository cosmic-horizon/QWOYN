#!/bin/sh

qwoynd51 tx gov submit-legacy-proposal software-upgrade "v5.2.0" \
--upgrade-height=12 \
--title="Upgrade to v5.2.0" --description="Upgrade to v5.2.0" --no-validate \
--from=validator --keyring-backend=test \
--chain-id=qwoyn-1 --yes -b sync --deposit="100000000stake"

qwoynd51 tx gov vote 1 yes --from validator --chain-id qwoyn-1 \
-b sync -y --keyring-backend test

qwoynd51 query gov proposals

qwoynd52 start --pruning=nothing --minimum-gas-prices="0stake"

qwoynd52 query bank total