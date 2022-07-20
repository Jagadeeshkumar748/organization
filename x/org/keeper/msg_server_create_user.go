package keeper

import (
	"context"

	"org/x/org/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateUser(goCtx context.Context, msg *types.MsgCreateUser) (*types.MsgCreateUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	user := types.User{
		Creator: msg.Creator,
		Name:    msg.Name,
		Email:   msg.Email,
	}

	userid := k.AddUser(ctx, user)

	return &types.MsgCreateUserResponse{Userid: userid}, nil
}
