package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgCreateFetterApp
// ------------------------------------------------------------------------------
var _ sdk.Msg = &MsgCreateFetterApp{}

// MsgCreateScavenge - struct for unjailing jailed validator
type MsgCreateFetterApp struct {
	Creator      sdk.AccAddress `json:"creator" yaml:"creator"`           // address of the scavenger creator
	Description  string         `json:"description" yaml:"description"`   // description of the scavenge
	SolutionHash string         `json:"solutionHash" yaml:"solutionHash"` // solution hash of the scavenge
	Reward       sdk.Coins      `json:"reward" yaml:"reward"`             // reward of the scavenger
}

// NewMsgCreateScavenge creates a new MsgCreateScavenge instance
func NewMsgCreateFetterApp(creator sdk.AccAddress, description, solutionHash string, reward sdk.Coins) MsgCreateFetterApp {
	return MsgCreateFetterApp{
		Creator:      creator,
		Description:  description,
		SolutionHash: solutionHash,
		Reward:       reward,
	}
}

// CreateScavengeConst is CreateScavenge Constant
const CreateFetterAppConst = "CreateScavenge"

// nolint
func (msg MsgCreateFetterApp) Route() string { return RouterKey }
func (msg MsgCreateFetterApp) Type() string  { return CreateFetterAppConst }
func (msg MsgCreateFetterApp) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgCreateFetterApp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgCreateFetterApp) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if msg.SolutionHash == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "solutionScavengerHash can't be empty")
	}
	return nil
}
