package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var (
	KeyDepositToken           = []byte("DepositToken")
	KeyAllocationToken        = []byte("AllocationToken")
	KeyVestingDuration        = []byte("VestingDuration")
	KeyDepositEndTime         = []byte("DepositEndTime")
	KeyInitLiquidityPrice     = []byte("InitLiquidityPrice")
	KeyLiquidityBootstrapping = []byte("LiquidityBootstrapping")
	KeyLiquidityBootstrapped  = []byte("LiquidityBootstrapped")
	KeyIcsAccount             = []byte("IcsAccount")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	depositToken string,
	allocationToken string,
	vestingDuraiton uint64,
	depositEndTime uint64,
	initLiquidityPrice sdk.Dec,
	liquidityBootstrapping bool,
	liquidityBootstrapped bool,
	icsConnectionId string,
) Params {
	return Params{
		DepositToken:           depositToken,
		AllocationToken:        allocationToken,
		VestingDuration:        vestingDuraiton,
		DepositEndTime:         depositEndTime,
		InitLiquidityPrice:     initLiquidityPrice,
		LiquidityBootstrapping: liquidityBootstrapping,
		LiquidityBootstrapped:  liquidityBootstrapped,
		IcsConnectionId:        icsConnectionId,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams("ibc/C053D637CCA2A2BA030E2C5EE1B28A16F71CCB0E45E8BE52766DC1B241B77878", "uqwoyn", 86400*360, 1679578470, sdk.OneDec(), false, false, "")
	// return NewParams("stake", "uqwoyn", 86400*360, 1679578470, sdk.OneDec(), false, false, "")
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyDepositToken, &p.DepositToken, validateDenom),
		paramtypes.NewParamSetPair(KeyAllocationToken, &p.AllocationToken, validateDenom),
		paramtypes.NewParamSetPair(KeyVestingDuration, &p.VestingDuration, validateVestingDuration),
		paramtypes.NewParamSetPair(KeyDepositEndTime, &p.DepositEndTime, validateDepositEndTime),
		paramtypes.NewParamSetPair(KeyInitLiquidityPrice, &p.InitLiquidityPrice, validateInitialLiquidityPrice),
		paramtypes.NewParamSetPair(KeyLiquidityBootstrapping, &p.LiquidityBootstrapping, validateBool),
		paramtypes.NewParamSetPair(KeyLiquidityBootstrapped, &p.LiquidityBootstrapped, validateBool),
		paramtypes.NewParamSetPair(KeyIcsAccount, &p.IcsConnectionId, validateIcsConnectionId),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateDenom(p.DepositToken); err != nil {
		return err
	}

	if err := validateDenom(p.AllocationToken); err != nil {
		return err
	}

	if err := validateVestingDuration(p.VestingDuration); err != nil {
		return err
	}

	if err := validateDepositEndTime(p.DepositEndTime); err != nil {
		return err
	}

	if err := validateInitialLiquidityPrice(p.InitLiquidityPrice); err != nil {
		return err
	}

	if err := validateBool(p.LiquidityBootstrapping); err != nil {
		return err
	}

	if err := validateBool(p.LiquidityBootstrapped); err != nil {
		return err
	}

	if err := validateIcsConnectionId(p.IcsConnectionId); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if err := sdk.ValidateDenom(v); err != nil {
		return err
	}

	return nil
}

func validateVestingDuration(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("vesting duration should be positive")
	}
	return nil
}

func validateDepositEndTime(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("deposit end time should be positive")
	}
	return nil
}

func validateInitialLiquidityPrice(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.IsPositive() {
		return fmt.Errorf("initial liquidity price should be positive")
	}

	return nil
}

func validateBool(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateIcsConnectionId(i interface{}) error {
	_, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
