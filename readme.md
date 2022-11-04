# Cosmic Horizon Network

# Installing `qwoynd`

## Hardware Requirements

Here are the minimal hardware configs required for running a validator/sentry node

- 16GB RAM
- 4vCPUs
- 500gb Disk space

## Software Requirements

- Ubuntu 18.04 or higher
- [Go v1.18.1](https://golang.org/doc/install) or higher

# Install `Qwoynd`, Generate Wallet and Start your Node

You have two options for installing the qwoynd binary. First, our team will be providing releases of qwoynd in our github repository, please check the [releases page](https://github.com/cosmic-horizon/qwoyn/releases) for the latest version of qwoynd. Secondly, you can follow the steps below to compile `Qwoynd` yourself.

## Install Go version 1.18.1

```
sudo apt update
sudo apt install build-essential jq wget git -y

wget https://dl.google.com/go/go1.18.1.linux-amd64.tar.gz
tar -xvf go1.18.1.linux-amd64.tar.gz
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

## Build `Qwoynd`

These steps install the `qwoynd` binary.

```
# Clone the Repo
git clone https://github.com/cosmic-horizon/QWOYN.git

# Install `Qwoynd`
cd QWOYN
git checkout v1.0.0-beta
make install

# check version (should be v1.0.0-beta)
qwoynd version
```
