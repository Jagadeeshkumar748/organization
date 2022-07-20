package keeper

import (
	"encoding/binary"
	"fmt"
	"org/x/org/types"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AddUser(ctx sdk.Context, user types.User) uint64 {
	// Get the current number of posts in the store
	count := k.GetUserCount(ctx)
	// Assign an ID to the post based on the number of posts in the store
	user.Userid = count
	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.UserKey))
	// Convert the post ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, user.Userid)
	// Marshal the post into bytes
	appendedValue := k.cdc.MustMarshal(&user)
	// Insert the post bytes using post ID as a key
	store.Set(byteKey, appendedValue)
	// Update the post count
	k.SetUserCount(ctx, count+1)
	return count
}

func (k Keeper) GetUser(ctx sdk.Context, userid string) (types.User, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.UserKey))

	byteKey := make([]byte, 8)
	userId, _ := strconv.ParseUint(userid, 0, 64)
	binary.BigEndian.PutUint64(byteKey, userId)

	bz := store.Get(byteKey)

	var user types.User

	if bz == nil {
		return types.User{}, fmt.Errorf("user not found")
	}

	k.cdc.Unmarshal(bz, &user)
	return user, nil
}

func (k Keeper) GetUserCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey (which is "blog") and PostCountKey (which is "Post-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.UserCountKey))
	// Convert the PostCountKey to bytes
	byteKey := []byte(types.UserCountKey)
	// Get the value of the count
	bz := store.Get(byteKey)
	// Return zero if the count value is not found (for example, it's the first post)
	if bz == nil {
		return 0
	}
	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetUserCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey (which is "blog") and PostCountKey (which is "Post-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.UserCountKey))
	// Convert the PostCountKey to bytes
	byteKey := []byte(types.UserCountKey)
	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	// Set the value of Post-count- to count
	store.Set(byteKey, bz)
}
