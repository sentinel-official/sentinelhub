package utils

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetProportionOfCoin(coin sdk.Coin, share sdkmath.LegacyDec) sdk.Coin {
	decAmount := sdkmath.LegacyNewDecFromInt(coin.Amount)
	amount := decAmount.Mul(share).Ceil().TruncateInt()

	return sdk.Coin{
		Denom:  coin.Denom,
		Amount: amount,
	}
}
