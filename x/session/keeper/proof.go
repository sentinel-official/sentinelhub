package keeper

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/sentinelhub/v12/x/session/types/v3"
)

// VerifySignature validates the cryptographic signature for a given session proof.
func (k *Keeper) VerifySignature(ctx sdk.Context, addr sdk.AccAddress, proof *v3.Proof, signature []byte) error {
	// Retrieve the account associated with the provided address.
	acc := k.account.GetAccount(ctx, addr)
	if acc == nil {
		return fmt.Errorf("account for address %s does not exist", addr)
	}

	// Extract the public key from the account.
	pubKey := acc.GetPubKey()
	if pubKey == nil {
		return fmt.Errorf("public key for address %s does not exist", addr)
	}

	// Convert the proof object into a byte slice (message to verify).
	message, err := proof.Marshal()
	if err != nil {
		return err
	}

	// Verify the signature against the marshaled message and public key.
	if !pubKey.VerifySignature(message, signature) {
		return errors.New("invalid signature for message")
	}

	// Signature is valid
	return nil
}
