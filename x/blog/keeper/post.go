package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/example/blog/x/blog/types"
	"strconv"
)

func (k Keeper) GetPostCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostCountKey))
	byteKey := types.KeyPrefix(types.PostCountKey)
	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		panic("cannot decode count")
	}

	return count
}

func (k Keeper) SetPostCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostCountKey))
	byteKey := types.KeyPrefix(types.PostCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) CreatePost(ctx sdk.Context, msg types.MsgCreatePost) {
	count := k.GetPostCount(ctx)
	var post = types.Post{
		Creator: msg.Creator,
		Id:      strconv.FormatInt(count, 10),
		Title:   msg.Title,
		Body:    msg.Body,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	key := types.KeyPrefix(types.PostKey + post.Id)
	value := k.cdc.MustMarshalBinaryBare(&post)
	store.Set(key, value)

	k.SetPostCount(ctx, count+1)
}

func (k Keeper) GetPost(ctx sdk.Context, key string) types.Post {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	var post types.Post
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.PostKey + key)), &post)
	return post
}

func (k Keeper) HasPost(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	return store.Has(types.KeyPrefix(types.PostKey + id))
}

func (k Keeper) GetPostOwner(ctx sdk.Context, key string) string {
	return k.GetPost(ctx, key).Creator
}

func (k Keeper) GetAllPost(ctx sdk.Context) (msgs []types.Post) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.PostKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Post
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
