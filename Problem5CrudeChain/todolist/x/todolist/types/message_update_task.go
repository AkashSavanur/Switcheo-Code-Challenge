package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateTask{}

func NewMsgUpdateTask(creator string, id uint64, title string, description string, completed bool) *MsgUpdateTask {
	return &MsgUpdateTask{
		Creator:     creator,
		Id:          id,
		Title:       title,
		Description: description,
		Completed:   completed,
	}
}

func (msg *MsgUpdateTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
