package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetWhois{}

type MsgSetWhois struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Value string `json:"value" yaml:"value"`
  Price string `json:"price" yaml:"price"`
}

func NewMsgSetWhois(creator sdk.AccAddress, id string, value string, price string) MsgSetWhois {
  return MsgSetWhois{
    ID: id,
		Creator: creator,
    Value: value,
    Price: price,
	}
}

func (msg MsgSetWhois) Route() string {
  return RouterKey
}

func (msg MsgSetWhois) Type() string {
  return "SetWhois"
}

func (msg MsgSetWhois) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetWhois) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetWhois) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}