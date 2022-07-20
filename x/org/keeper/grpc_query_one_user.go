package keeper

import (
	"context"

	"org/x/org/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OneUser(goCtx context.Context, req *types.QueryOneUserRequest) (*types.QueryOneUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	userid := req.GetUserid()
	user, err := k.GetUser(ctx, userid)
	if err != nil {
		return nil, err
	}

	k.cdc.MustMarshal(&user)
	return &types.QueryOneUserResponse{User: &user}, nil
}
