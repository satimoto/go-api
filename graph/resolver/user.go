package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/satimoto/go-api/graph"
	metrics "github.com/satimoto/go-api/internal/metric"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input graph.CreateUserInput) (*db.User, error) {
	auth, err := r.AuthenticationResolver.Repository.GetAuthenticationByCode(ctx, input.Code)

	if err != nil {
		metrics.RecordError("API016", "Authentication not found", err)
		log.Printf("API016: Code=%v", input.Code)
		return nil, gqlerror.Errorf("Authentication not found")
	}

	if !auth.Signature.Valid {
		log.Printf("API017: Authentication not yet verified")
		log.Printf("API017: Challenge=%v", auth.Challenge)
		return nil, gqlerror.Errorf("Authentication not yet verified")
	}

	var circuitUserID *int64
	ipAddress := middleware.GetIPAddress(ctx)

	if ipAddress != nil && len(*ipAddress) > 0 {
		if referral, err := r.ReferralRepository.GetReferralByIpAddress(ctx, *ipAddress); err == nil {
			circuitUserID = &referral.UserID
		}
	}

	referralCode := r.generateReferralCode(ctx)
	createUserParams := db.CreateUserParams{
		CommissionPercent: dbUtil.GetEnvFloat64("DEFAULT_COMMISSION_PERCENT", 7),
		DeviceToken:       dbUtil.SqlNullString(input.DeviceToken),
		LinkingPubkey:     auth.LinkingPubkey.String,
		Pubkey:            input.Pubkey,
		ReferralCode:      dbUtil.SqlNullString(referralCode),
		CircuitUserID:     dbUtil.SqlNullInt64(circuitUserID),
	}

	user, err := r.UserRepository.CreateUser(ctx, createUserParams)

	if err != nil {
		metrics.RecordError("API018", "User already exists", err)
		log.Printf("API018: Params=%#v", createUserParams)
		return nil, gqlerror.Errorf("User already exists")
	}

	_, err = r.TokenResolver.CreateToken(ctx, user.ID)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input graph.UpdateUserInput) (*db.User, error) {
	if user := middleware.GetUser(ctx, r.UserRepository); user != nil {
		updateUserParams := param.NewUpdateUserParams(*user)
		updateUserParams.DeviceToken = dbUtil.SqlNullString(input.DeviceToken)

		updatedUser, err := r.UserRepository.UpdateUser(ctx, updateUserParams)

		if err != nil {
			metrics.RecordError("API020", "Error updating user", err)
			log.Printf("API020: Params=%#v", updateUserParams)
			return nil, gqlerror.Errorf("Error updating user")
		}

		updatePendingNotificationByUserParams := db.UpdatePendingNotificationsByUserParams{
			DeviceToken: dbUtil.SqlNullString(input.DeviceToken),
			UserID:      user.ID,
		}

		err = r.PendingNotificationRepository.UpdatePendingNotificationsByUser(ctx, updatePendingNotificationByUserParams)

		if err != nil {
			metrics.RecordError("API027", "Error updating pending notifications", err)
			log.Printf("API027: Params=%#v", updatePendingNotificationByUserParams)
			return nil, gqlerror.Errorf("Error updating pending notifications")
		}

		return &updatedUser, nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context) (*db.User, error) {
	user := middleware.GetUser(ctx, r.UserRepository)

	if user != nil {
		return user, nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// DeviceToken is the resolver for the deviceToken field.
func (r *userResolver) DeviceToken(ctx context.Context, obj *db.User) (*string, error) {
	return util.NullString(obj.DeviceToken)
}

// ReferralCode is the resolver for the referralCode field.
func (r *userResolver) ReferralCode(ctx context.Context, obj *db.User) (*string, error) {
	return util.NullString(obj.ReferralCode)
}

// Node is the resolver for the node field.
func (r *userResolver) Node(ctx context.Context, obj *db.User) (*db.Node, error) {
	if obj.NodeID.Valid {
		if node, err := r.NodeRepository.GetNode(ctx, obj.NodeID.Int64); err == nil {
			return &node, nil
		}
	}

	return nil, nil
}

// User returns graph.UserResolver implementation.
func (r *Resolver) User() graph.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
