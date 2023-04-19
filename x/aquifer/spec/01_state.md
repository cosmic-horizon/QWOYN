# State

## Params

```protobuf
// Params defines the parameters for the module.
message Params {
  option (gogoproto.goproto_stringer) = false;

  string maintainer = 1;
  string deposit_token = 2; // axlUSDC.osmo
  string allocation_token = 3; // qwoyn
  uint64 vesting_duration = 4; // 1 year in seconds
  uint64 deposit_end_time = 5;
  string init_liquidity_price = 6 [
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable) = false
  ]; // price without considering decimal
  string discount = 7 [
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable) = false
  ]; // discount from initial liquidity price
  bool liquidity_bootstrapping = 8; // broadcasted ICA tx to bootstrap Osmosis LP
  bool liquidity_bootstrapped = 9; // set as true once Osmosis liquidity created
  string ica_connection_id = 10;
}
```

Aquifer module's state is managed as params it is governed by governance.

- `maintainer` is the maintainer of module to manage deposit end time, remaining QWOYN tokens after deposit period, sending signal to setup ICA, transfer of funds to Osmosis, and send ICA message for liquidity bootstrap.
- `deposit_token` is the token to be raised through the discounted price offering
- `allocation_token` is the token that is allocated in vesting
- `vesting_duration` is the vesting duration of token
- `deposit_end_time` is the time where the deposit period ends
- `init_liquidity_price` is the price for initial liquidity providing on Osmosis
- `discount` is the discount from initial liquidity price for vesting
- `liquidity_bootstrapping` is true when osmosis LP bootstrap ICA transaction is broadcasted and not received finalization message
- `liquidity_bootstrapped` is true when osmosis LP bootstrap ICA transaction has received acknowledgement packet
- `ica_connection_id` is connection id with Osmosis for ICA account used for liquidity bootstrap
