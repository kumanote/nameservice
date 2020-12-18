package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateWhois{}

type MsgCreateWhois struct {
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Value string `json:"value" yaml:"value"`
  Price string `json:"price" yaml:"price"`
}

func NewMsgCreateWhois(creator sdk.AccAddress, value string, price string) MsgCreateWhois {
  return MsgCreateWhois{
		Creator: creator,
    Value: value,
    Price: price,
	}
}

func (msg MsgCreateWhois) Route() string {
  return RouterKey
}

func (msg MsgCreateWhois) Type() string {
  return "CreateWhois"
}

func (msg MsgCreateWhois) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateWhois) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateWhois) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}