package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteUser = "delete_user"

var _ sdk.Msg = &MsgDeleteUser{}

func NewMsgDeleteUser(creator string, userid string) *MsgDeleteUser {
	return &MsgDeleteUser{
		Creator: creator,
		Userid:  userid,
	}
}

func (msg *MsgDeleteUser) Route() string {
	return RouterKey
}

func (msg *MsgDeleteUser) Type() string {
	return TypeMsgDeleteUser
}

func (msg *MsgDeleteUser) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteUser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
