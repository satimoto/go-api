package user

import "github.com/satimoto/go-datastore/db"

func NewUpdateUserParams(user db.User) db.UpdateUserParams {
	return db.UpdateUserParams{
		ID:          user.ID,
		DeviceToken: user.DeviceToken,
		NodeID:      user.NodeID,
		TokenID:     user.TokenID,
	}
}
