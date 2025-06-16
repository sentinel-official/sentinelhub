package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v12/x/session/types/v3"
)

func (k *Keeper) VerifySignature(ctx sdk.Context, addr sdk.AccAddress, proof *v3.Proof, signature []byte) error {
	acc := k.account.GetAccount(ctx, addr)
	if acc == nil {
		return fmt.Errorf("account for address %s does not exist", addr)
	}

	pubKey := acc.GetPubKey()
	if pubKey == nil {
		return fmt.Errorf("public key for address %s does not exist", addr)
	}

	message, err := proof.Marshal()
	if err != nil {
		return err
	}

	if !pubKey.VerifySignature(message, signature) {
		return fmt.Errorf("invalid signature for message")
	}

	return nil
}
