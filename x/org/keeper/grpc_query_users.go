package keeper

import (
	"context"

	"org/x/org/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Users(c context.Context, req *types.QueryUsersRequest) (*types.QueryUsersResponse, error) {
	// Throw an error if request is nil
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	// Define a variable that will store a list of posts
	var users []*types.User
	// Get context with the information about the environment
	ctx := sdk.UnwrapSDKContext(c)
	// Get the key-value module store using the store key (in our case store key is "chain")
	store := ctx.KVStore(k.storeKey)
	// Get the part of the store that keeps posts (using post key, which is "Post-value-")
	UserStore := prefix.NewStore(store, []byte(types.UserKey))
	// Paginate the posts store based on PageRequest
	pageRes, err := query.Paginate(UserStore, req.Pagination, func(key []byte, value []byte) error {
		var user types.User
		if err := k.cdc.Unmarshal(value, &user); err != nil {
			return err
		}
		users = append(users, &user)
		return nil
	})
	// Throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	// Return a struct containing a list of posts and pagination info
	return &types.QueryUsersResponse{User: users, Pagination: pageRes}, nil
}
