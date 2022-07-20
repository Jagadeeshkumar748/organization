package keeper

import (
	"context"

	"org/x/org/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteUser(goCtx context.Context, msg *types.MsgDeleteUser) (*types.MsgDeleteUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	userid := msg.GetUserid()
	_, err := k.GetUser(ctx, userid)
	if err != nil {
		return nil, err
	}

	k.Delete(ctx, userid)

	return &types.MsgDeleteUserResponse{}, nil
}
