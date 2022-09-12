# Cosmic Horizon Network

# Installing CoHo

## Hardware Requirements

Here are the minimal hardware configs required for running a validator/sentry node

- 8GB RAM
- 4vCPUs
- 200GB Disk space

## Software Requirements

- Ubuntu 18.04 or higher
- [Go v1.17.1](https://golang.org/doc/install)

# Install CoHo, Generate Wallet and Start your Node

You have two options for installing the cohod binary. First, our team will be providing releases of cohod in our github repository, please check the [releases page](https://github.com/cosmic-horizon/coho/releases) for the latest version of cohod. Secondly, you can follow the steps below to compile coho yourself.

## Install Go version 1.17.1

```
sudo apt update
sudo apt install build-essential jq wget git -y

wget https://dl.google.com/go/go1.17.1.linux-amd64.tar.gz
tar -xvf go1.17.1.linux-amd64.tar.gz
sudo mv go /usr/local
```

Now add go to your bashrc -

```
echo "" >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
echo 'export GOBIN=$GOPATH/bin' >> ~/.bashrc
echo 'export PATH=$PATH:/usr/local/go/bin:$GOBIN' >> ~/.bashrc

# use this new bashrc configuration
source ~/.bashrc
```

## Build CoHo

These steps install the `cohod` binary.

```
# Clone the Repo
git clone https://github.com/cosmic-horizon/coho.git

# Install CoHo
cd coho
make install
```

## Testnet generation

```
cohod testnet --keyring-backend=test --chain-id="qwoyn-1" --v 4 --output-dir ./testnet --starting-ip-address 192.168.10.2
```
