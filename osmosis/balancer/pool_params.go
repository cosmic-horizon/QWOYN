package balancer

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewPoolParams(swapFee, exitFee sdk.Dec, params *SmoothWeightChangeParams) PoolParams {
	return PoolParams{
		SwapFee:                  swapFee,
		ExitFee:                  exitFee,
		SmoothWeightChangeParams: params,
	}
}

func (params PoolParams) Validate(poolWeights []PoolAsset) error {

	return nil
}

func (params PoolParams) GetPoolSwapFee() sdk.Dec {
	return params.SwapFee
}

func (params PoolParams) GetPoolExitFee() sdk.Dec {
	return params.ExitFee
}
