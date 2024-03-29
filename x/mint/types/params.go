package types

import (
	"errors"
	"fmt"
	"strings"

	appparams "github.com/cosmic-horizon/qwoyn/app/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	yaml "gopkg.in/yaml.v2"
)

// Parameter store keys
var (
	KeyMintDenom                 = []byte("MintDenom")
	KeyInflationRateChange       = []byte("InflationRateChange")
	KeyInflationMax              = []byte("InflationMax")
	KeyInflationMin              = []byte("InflationMin")
	KeyGoalBonded                = []byte("GoalBonded")
	KeyBlocksPerYear             = []byte("BlocksPerYear")
	KeyMaxCap                    = []byte("MaxCap")
	KeyOutpostFundingPoolPortion = []byte("OutpostFundingPool")
)

// ParamTable for minting module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(
	mintDenom string, inflationRateChange, inflationMax, inflationMin, goalBonded sdk.Dec, blocksPerYear uint64, maxCap sdk.Int, outpostFundingPoolPortion sdk.Dec,
) Params {

	return Params{
		MintDenom:                 mintDenom,
		InflationRateChange:       inflationRateChange,
		InflationMax:              inflationMax,
		InflationMin:              inflationMin,
		GoalBonded:                goalBonded,
		BlocksPerYear:             blocksPerYear,
		MaxCap:                    maxCap,
		OutpostFundingPoolPortion: outpostFundingPoolPortion,
	}
}

// default minting module parameters
func DefaultParams() Params {
	return Params{
		MintDenom:                 appparams.BondDenom,
		InflationRateChange:       sdk.NewDecWithPrec(13, 2),
		InflationMax:              sdk.NewDecWithPrec(40, 2),
		InflationMin:              sdk.NewDecWithPrec(7, 2),
		GoalBonded:                sdk.NewDecWithPrec(67, 2),
		BlocksPerYear:             4360000,
		MaxCap:                    sdk.NewInt(21_000_000_000_000), // 21,000,000 QWOYN
		OutpostFundingPoolPortion: sdk.NewDecWithPrec(50, 2),      // 50%
	}
}

// validate params
func (p Params) Validate() error {
	if err := validateMintDenom(p.MintDenom); err != nil {
		return err
	}
	if err := validateInflationRateChange(p.InflationRateChange); err != nil {
		return err
	}
	if err := validateInflationMax(p.InflationMax); err != nil {
		return err
	}
	if err := validateInflationMin(p.InflationMin); err != nil {
		return err
	}
	if err := validateGoalBonded(p.GoalBonded); err != nil {
		return err
	}
	if err := validateBlocksPerYear(p.BlocksPerYear); err != nil {
		return err
	}
	if p.InflationMax.LT(p.InflationMin) {
		return fmt.Errorf(
			"max inflation (%s) must be greater than or equal to min inflation (%s)",
			p.InflationMax, p.InflationMin,
		)
	}
	if err := validateMaxCap(p.MaxCap); err != nil {
		return err
	}

	if err := validateOutpostFundingPoolPortion(p.OutpostFundingPoolPortion); err != nil {
		return err
	}

	return nil

}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMintDenom, &p.MintDenom, validateMintDenom),
		paramtypes.NewParamSetPair(KeyInflationRateChange, &p.InflationRateChange, validateInflationRateChange),
		paramtypes.NewParamSetPair(KeyInflationMax, &p.InflationMax, validateInflationMax),
		paramtypes.NewParamSetPair(KeyInflationMin, &p.InflationMin, validateInflationMin),
		paramtypes.NewParamSetPair(KeyGoalBonded, &p.GoalBonded, validateGoalBonded),
		paramtypes.NewParamSetPair(KeyBlocksPerYear, &p.BlocksPerYear, validateBlocksPerYear),
		paramtypes.NewParamSetPair(KeyMaxCap, &p.MaxCap, validateMaxCap),
		paramtypes.NewParamSetPair(KeyOutpostFundingPoolPortion, &p.OutpostFundingPoolPortion, validateOutpostFundingPoolPortion),
	}
}

func validateMintDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if strings.TrimSpace(v) == "" {
		return errors.New("mint denom cannot be blank")
	}
	if err := sdk.ValidateDenom(v); err != nil {
		return err
	}

	return nil
}

func validateInflationRateChange(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("inflation rate change cannot be negative: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("inflation rate change too large: %s", v)
	}

	return nil
}

func validateInflationMax(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("max inflation cannot be negative: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("max inflation too large: %s", v)
	}

	return nil
}

func validateInflationMin(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("min inflation cannot be negative: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("min inflation too large: %s", v)
	}

	return nil
}

func validateGoalBonded(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("goal bonded cannot be negative: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("goal bonded too large: %s", v)
	}

	return nil
}

func validateBlocksPerYear(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("blocks per year must be positive: %d", v)
	}

	return nil
}

func validateMaxCap(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsZero() {
		return fmt.Errorf("max cap must be positive: %d", v)
	}

	return nil
}

func validateOutpostFundingPoolPortion(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("outpost funding pool portion cannot be negative: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("outpost funding pool portion too large: %s", v)
	}

	return nil
}
