#!/bin/bash

exit_with_error()
{
  echo "$1" 1>&2
  exit 1
}

KeyringBackend=test
Chain=test-1
GAS_FEE="0uqwoyn"

# Get the options
while getopts ":kc:" option; do
   case $option in
      k) # Keyring backend
         KeyringBackend=$OPTARG;;
      c) # Enter a name
         Chain=$OPTARG;;
     \?) # Invalid option
         echo "Error: Invalid option"
         exit 1
   esac
done

# Make sure the path is set correctly
export PATH=~/go/bin:$PATH

echo "qwoynd Version: `qwoynd version`"

qwoynd keys add validator --keyring-backend ${KeyringBackend} || exit_with_error "Error: Validator add failed"
qwoynd keys add delegator --keyring-backend ${KeyringBackend} || exit_with_error "Error: Delegator add failed"
qwoynd init node --chain-id ${Chain} || exit_with_error "Error: Could not init node"

# Change the staking token to uqwoyn and change voting period to 1 minute
# Note: sed works differently on different platforms
echo "Updating your staking token to uqwoyn in the genesis file..."
OS=`uname`
if [[ $OS == "Linux"* ]]; then
    echo "Your OS is a Linux variant..."
    sed -i "s/stake/uqwoyn/g" ~/.qwoynd/config/genesis.json || exit_with_error "Error: Could not update staking token"
    sed -i '/minimum-gas-prices =/c\minimum-gas-prices = "'"$GAS_FEE"'"' ~/.qwoynd/config/app.toml || exit_with_error "Error: Could not update min gas fee"
    sed -i 's/"voting_period": ".*"/"voting_period": "60s"/' ~/.qwoynd/config/genesis.json || exit_with_error "Error: Could not update voting period"
elif [[ $OS == "Darwin"* ]]; then
    echo "Your OS is Mac OS/darwin..."
    sed -i "" "s/stake/uqwoyn/g" ~/.qwoynd/config/genesis.json || exit_with_error "Error: Could not update staking token"
else
    # Dunno
    echo "Your OS is not supported"
    exit 1
fi

echo "Adding validator to genesis.json..."
qwoynd genesis add-genesis-account validator 5000000000uqwoyn --keyring-backend ${KeyringBackend} || exit_with_error "Error: Could not add validator to genesis"

echo "Adding delegator to genesis.json..."
qwoynd genesis add-genesis-account delegator 2000000000uqwoyn --keyring-backend ${KeyringBackend} || exit_with_error "Error: Could not add delegator to genesis"

echo "Adding qwoyn1qn8mzcel2svf6s47hltuxtkl8e7w07tjaaleuj to genesis.json..."
qwoynd genesis add-genesis-account qwoyn1qn8mzcel2svf6s47hltuxtkl8e7w07tjaaleuj 1000000uqwoyn --keyring-backend ${KeyringBackend} || exit_with_error "Error: Could not add delegator to genesis"

echo "Creating genesis transaction..."
qwoynd genesis gentx validator 1000000000uqwoyn --chain-id ${Chain} --keyring-backend ${KeyringBackend} || exit_with_error "Error: Genesis transaction failed"

echo "Adding genesis transaction to genesis.json..."
qwoynd genesis collect-gentxs || exit_with_error "Error: Could not add transaction to genesis"

echo "If there were no errors above, you can now type 'qwoynd start' to start your node"
