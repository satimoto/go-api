package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input graph.CreateUserInput) (*db.User, error) {
	auth, err := r.AuthenticationResolver.Repository.GetAuthenticationByCode(ctx, input.Code)

	if err != nil {
		dbUtil.LogOnError("API016", "Authentication not found", err)
		log.Printf("API016: Code=%v", input.Code)
		return nil, gqlerror.Errorf("Authentication not found")
	}

	if !auth.Signature.Valid {
		log.Printf("API017: Authentication not yet verified")
		log.Printf("API017: Challenge=%v", auth.Challenge)
		return nil, gqlerror.Errorf("Authentication not yet verified")
	}

	var circuitUserId *int64
	ipAddress := middleware.GetIPAddress(ctx)

	if len(*ipAddress) > 0 {
		if referral, err := r.ReferralRepository.GetReferralByIpAddress(ctx, *ipAddress); err == nil {
			circuitUserId = &referral.UserID
		}
	}

	referralCode := r.generateReferralCode(ctx)

	user, err := r.UserRepository.CreateUser(ctx, db.CreateUserParams{
		CommissionPercent: dbUtil.GetEnvFloat64("DEFAULT_COMMISSION_PERCENT", 7),
		DeviceToken:       input.DeviceToken,
		LinkingPubkey:     auth.LinkingPubkey.String,
		Pubkey:            input.Pubkey,
		ReferralCode:      dbUtil.SqlNullString(referralCode),
		CircuitUserID:     dbUtil.SqlNullInt64(circuitUserId),
	})

	if err != nil {
		dbUtil.LogOnError("API018", "User already exists", err)
		return nil, gqlerror.Errorf("User already exists")
	}

	_, err = r.TokenResolver.CreateToken(ctx, user.ID)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input graph.UpdateUserInput) (*db.User, error) {
	if user := middleware.GetUser(ctx, r.UserRepository); user != nil {
		updateUserParams := param.NewUpdateUserParams(*user)
		updateUserParams.DeviceToken = input.DeviceToken

		updatedUser, err := r.UserRepository.UpdateUser(ctx, updateUserParams)

		if err != nil {
			dbUtil.LogOnError("API020", "Error updating user", err)
			return nil, gqlerror.Errorf("Error updating user")
		}

		return &updatedUser, nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

func (r *userResolver) ReferralCode(ctx context.Context, obj *db.User) (*string, error) {
	return util.NullString(obj.ReferralCode)
}

func (r *mutationResolver) generateReferralCode(ctx context.Context) string {
	for {
		referralCode := util.RandomString(8)

		if _, err := r.UserRepository.GetUserByReferralCode(ctx, dbUtil.SqlNullString(referralCode)); err != nil {
			return referralCode
		}
	}
}

// User returns graph.UserResolver implementation.
func (r *Resolver) User() graph.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }

