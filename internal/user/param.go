package user

import "github.com/satimoto/go-datastore/db"

func NewUpdateUserParams(user db.User) db.UpdateUserParams {
	return db.UpdateUserParams{
		ID:                user.ID,
		CommissionPercent: user.CommissionPercent,
		DeviceToken:       user.DeviceToken,
		LinkingPubkey:     user.LinkingPubkey,
		NodeID:            user.NodeID,
		Pubkey:            user.Pubkey,
		IsRestricted:      user.IsRestricted,
		ReferrerID:        user.ReferrerID,
	}
}
