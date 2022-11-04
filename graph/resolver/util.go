package resolver

import (
	"context"

	"github.com/satimoto/go-api/internal/util"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
)

func (r *mutationResolver) generateReferralCode(ctx context.Context) string {
	for {
		referralCode := util.RandomString(8)

		if _, err := r.UserRepository.GetUserByReferralCode(ctx, dbUtil.SqlNullString(referralCode)); err != nil {
			return referralCode
		}
	}
}
