package types

import (
	fmt "fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// parameter keys
var (
	KeyOwner            = []byte("Owner")
	KeyDepositDenom     = []byte("DepositDenom")
	KeyStakingInflation = []byte("StakingInflation")
	KeyUnstakingTime    = []byte("UnstakingTime")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(owner, depositDenom string, stakingInflation uint64, unstakingTime time.Duration) Params {
	return Params{
		Owner:            owner,
		DepositDenom:     depositDenom,
		StakingInflation: stakingInflation,
		UnstakingTime:    unstakingTime,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams("coho1x0fha27pejg5ajg8vnrqm33ck8tq6raafkwa9v", "stake", 1, time.Second*30)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyOwner, &p.Owner, validateOwner),
		paramtypes.NewParamSetPair(KeyDepositDenom, &p.DepositDenom, validateDenom),
		paramtypes.NewParamSetPair(KeyStakingInflation, &p.StakingInflation, validateStakingInflation),
		paramtypes.NewParamSetPair(KeyUnstakingTime, &p.UnstakingTime, validateUnstakingTime),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	err := validateOwner(p.Owner)
	if err != nil {
		return err
	}

	err = validateDenom(p.DepositDenom)
	if err != nil {
		return err
	}

	err = validateStakingInflation(p.StakingInflation)
	if err != nil {
		return err
	}

	err = validateUnstakingTime(p.UnstakingTime)
	if err != nil {
		return err
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateOwner(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if _, err := sdk.AccAddressFromBech32(v); err != nil {
		return fmt.Errorf("invalid owner address")
	}
	return nil
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

func validateStakingInflation(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateUnstakingTime(i interface{}) error {
	_, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
