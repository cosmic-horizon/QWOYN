<div align="center" style="font-size:20px">
    <img alt="Issues" src="https://i.imgur.com/EZBSGwH.png" />
    <h3>
        <i>Decentralized Gaming Blockchain build using cosmos sdk</i>
    </h3>
</div>
<br />
<div align="center">
    <a href="https://github.com/cosmic-horizon/QWOYN/blob/main/LICENSE">
        <img alt="License" src="https://img.shields.io/badge/License-Apache%202.0-blue" />
    </a>
    <a href="https://github.com/cosmic-horizon/QWOYN/releases/latest">
        <img alt="Version" src="https://img.shields.io/github/tag/cosmic-horizon/QWOYN" />
    </a>
    <a href="https://pkg.go.dev/github.com/cosmic-horizon/QWOYN/v5">
        <img alt="Go Doc" src="https://pkg.go.dev/badge/github.com/cosmic-horizon/QWOYN/v5" />
    </a>
</div>
<br />
<div align="center">
    <a href="https://github.com/cosmic-horizon/QWOYN/commits/main">
        <img alt="Build Status" src="https://github.com/cosmic-horizon/QWOYN/workflows/Build/badge.svg" />
    </a>
    <a href="https://github.com/cosmic-horizon/QWOYN/commits/main">
        <img alt="Test Status" src="https://github.com/cosmic-horizon/QWOYN/workflows/Tests/badge.svg" />
    </a>
    <a href="https://github.com/cosmic-horizon/QWOYN/commits/main">
        <img alt="Sims Status" src="https://github.com/cosmic-horizon/QWOYN/workflows/Sims/badge.svg" />
    </a>
    <a href="https://github.com/cosmic-horizon/QWOYN/commits/main">
        <img alt="Lint Status" src="https://github.com/cosmic-horizon/QWOYN/workflows/Lint/badge.svg" />
    </a>
</div>
<br />
<div align="center">
    <a href="https://github.com/cosmic-horizon/QWOYN/issues">
        <img alt="Issues" src="https://img.shields.io/github/issues/regen-network/regen-ledger?color=blue" />
    </a>
    <a href="https://github.com/cosmic-horizon/QWOYN/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22">
        <img alt="Good First Issues" src="https://img.shields.io/github/issues/regen-network/regen-ledger/good%20first%20issue?color=blue" />
    </a>
    <a href="https://github.com/cosmic-horizon/QWOYN/discussions">
        <img alt="Discussions" src="https://img.shields.io/github/discussions/regen-network/regen-ledger?color=blue" />
    </a>
    <a href="https://discord.cosmic-horizon.com">
        <img alt="Discord" src="https://img.shields.io/discord/684494798358315010?color=blue" />
    </a>
</div>
<br />
<div align="center">
    <a href="https://codecov.io/gh/cosmic-horizon/QWOYN	">
        <img alt="Code Coverage" src="https://codecov.io/gh/regen-network/cosmic-horizon/branch/main/graph/badge.svg" />
    </a>
</div>
<br />

## Introduction

Regen Ledger is a blockchain application for ecological assets and data claims built on top of [Cosmos SDK](http://github.com/cosmos/cosmos-sdk) and [Tendermint Core](http://github.com/tendermint/tendermint). Leveraging these tools, Regen Ledger provides the infrastructure for a Proof-of-Stake blockchain network governed by a community dedicated to planetary regeneration.

Features specific to Regen Ledger are developed within this repository as custom modules that are then wired up to the main application. The custom modules developed within Regen Ledger follow the same architecture and pattern as modules developed within Cosmos SDK and other Cosmos SDK applications.

The core features that Regen Ledger aims to provide include the following:

- infrastructure for managing the issuance and retirement of ecosystem service credits
- a database of ecological state and change of state claims that spans both on and off-chain data sources
- mechanisms for automating the assessment of ecological state, making payments, and issuing assets

Qwoyn Blockchain is under heavy development and as result the above features are implemented to varying degrees of completeness. For more information about our approach and vision, see [Qwoyn Blockchain Specification](specs/qwoyn-blockchain.md).

## Documentation

Documentation for Qwoyn Blockchain is hosted at [docs.qwoyn.studio](https://docs.qwoyn.studio). This includes installation instructions for users and developers, information about live networks running Qwoyn Blockchain, instructions on how to interact with local and live networks, infrastructure and module-specific documentation, tutorials for users and developers, migration guides for developers, upgrade guides for validators, a complete list of available commands, and more.

## Contributing

Contributions are more than welcome and greatly appreciated. All the information you need to get started should be available in [Contributing Guidelines](./CONTRIBUTING.md). Please take the time to read through the contributing guidelines before opening an issue or pull request. The following prerequisites and commands cover the basics.

### Prerequisites

- [Git](https://git-scm.com) `>=2`
- [Make](https://www.gnu.org/software/make/) `>=4`
- [Go](https://golang.org/) `>=1.18`

### Go Tools

Install go tools:

```
make tools
```

### Git Hooks

Configure git hooks:

```
git config core.hooksPath scripts/githooks
```

### Lint and Format

Run linter in all go modules:

```
make lint
```

Run linter and attempt to fix errors in all go modules:

```
make lint-fix
```

Run formatting in all go modules:

```
make format
```

Run linter for all proto files:

```
make proto-lint
```

Run linter and attempt to fix errors for all proto files:

```
make proto-lint-fix
```

Run formatting for all proto files:

```
make proto-format
```

### Running Tests

Run all unit and integrations tests:

```
make test
```

### Manual Testing

Build the qwoynd binary:

```
make build
```

View the available commands:

```
./build/qwoyn help
```

## Related Repositories

- [cosmic-horizon/governance](https://github.com/cosmic-horizon/governance) - guidelines and long-form proposals for Qwoyn Mainnet
- [cosmic-horizon/mainnet](https://github.com/cosmic-horizon/mainnet) - additional information and historical record for Qwoyn Mainnet
- [cosmic-horizon/testnets](https://github.com/cosmic-horizon/testnets) - additional information and historical record for Qwoyn Testnets
