package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/appleboy/go-fcm"
	"github.com/google/uuid"
	"github.com/satimoto/go-api/graph"
	metrics "github.com/satimoto/go-api/internal/metric"
	"github.com/satimoto/go-api/internal/middleware"
	"github.com/satimoto/go-api/internal/notification"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input graph.CreateUserInput) (*db.User, error) {
	backgroundCtx := context.Background()
	auth, err := r.AuthenticationResolver.Repository.GetAuthenticationByCode(backgroundCtx, input.Code)

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

	var circuitUserID, nodeId *int64
	ipAddress := middleware.GetIPAddress(ctx)

	if ipAddress != nil && len(*ipAddress) > 0 {
		if referral, err := r.ReferralRepository.GetReferralByIpAddress(backgroundCtx, *ipAddress); err == nil {
			circuitUserID = &referral.UserID
		}
	}

	isLsp := dbUtil.DefaultBool(input.Lsp, true)

	if nodes, err := r.NodeRepository.ListActiveNodes(backgroundCtx, isLsp); err == nil && len(nodes) > 0 {
		for _, n := range nodes {
			nodeId = &n.ID
			break
		}
	}

	referralCode := r.generateReferralCode(backgroundCtx)
	createUserParams := db.CreateUserParams{
		CommissionPercent: dbUtil.GetEnvFloat64("DEFAULT_COMMISSION_PERCENT", 7),
		DeviceToken:       dbUtil.SqlNullString(input.DeviceToken),
		LinkingPubkey:     auth.LinkingPubkey.String,
		Pubkey:            input.Pubkey,
		ReferralCode:      dbUtil.SqlNullString(referralCode),
		CircuitUserID:     dbUtil.SqlNullInt64(circuitUserID),
		NodeID:            dbUtil.SqlNullInt64(nodeId),
	}

	user, err := r.UserRepository.CreateUser(backgroundCtx, createUserParams)

	if err != nil {
		metrics.RecordError("API018", "User already exists", err)
		log.Printf("API018: Params=%#v", createUserParams)
		return nil, gqlerror.Errorf("User already exists")
	}

	_, err = r.TokenResolver.CreateToken(backgroundCtx, user.ID)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// PingUser is the resolver for the pingUser field.
func (r *mutationResolver) PingUser(ctx context.Context, id int64) (*graph.ResultOk, error) {
	backgroundCtx := context.Background()

	if user := middleware.GetCtxUser(ctx, r.UserRepository); user != nil && user.IsAdmin {
		if toUser, err := r.UserRepository.GetUser(backgroundCtx, id); err == nil && toUser.DeviceToken.Valid {
			ping, err := uuid.NewUUID()

			if err != nil {
				metrics.RecordError("API057", "Error generating UUID", err)
				log.Printf("API057: UserID=%v", id)
				return nil, gqlerror.Errorf("Error generating UUID")
			}

			pingStr := ping.String()
			dataPingMessage := notification.CreateDataPingNotificationDto(pingStr)

			message := &fcm.Message{
				To:               toUser.DeviceToken.String,
				ContentAvailable: true,
				Priority:         "high",
				Data:             dataPingMessage,
			}

			r.NotificationService.SendNotification(message)

			log.Printf("User %v ping sent: %v", toUser.ID, pingStr)

			return &graph.ResultOk{Ok: true}, nil
		}

		return &graph.ResultOk{Ok: false}, nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// PongUser is the resolver for the pongUser field.
func (r *mutationResolver) PongUser(ctx context.Context, input graph.PongUserInput) (*graph.ResultOk, error) {
	if userID := middleware.GetUserID(ctx); userID != nil {
		log.Printf("User %v pong received: %v", *userID, input.Pong)

		return &graph.ResultOk{Ok: true}, nil
	}

	return nil, gqlerror.Errorf("Not authenticated")
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input graph.UpdateUserInput) (*db.User, error) {
	backgroundCtx := context.Background()

	if user := middleware.GetCtxUser(ctx, r.UserRepository); user != nil {
		updateUserParams := param.NewUpdateUserParams(*user)
		updateUserParams.DeviceToken = dbUtil.SqlNullString(input.DeviceToken)
		updateUserParams.Name = dbUtil.SqlNullString(input.Name)
		updateUserParams.Address = dbUtil.SqlNullString(input.Address)
		updateUserParams.PostalCode = dbUtil.SqlNullString(input.PostalCode)
		updateUserParams.City = dbUtil.SqlNullString(input.City)
		updateUserParams.BatteryCapacity = dbUtil.SqlNullFloat64(input.BatteryCapacity)
		updateUserParams.BatteryPowerAc = dbUtil.SqlNullFloat64(input.BatteryPowerAc)
		updateUserParams.BatteryPowerDc = dbUtil.SqlNullFloat64(input.BatteryPowerDc)

		updatedUser, err := r.UserRepository.UpdateUser(backgroundCtx, updateUserParams)

		if err != nil {
			metrics.RecordError("API020", "Error updating user", err)
			log.Printf("API020: Params=%#v", updateUserParams)
			return nil, gqlerror.Errorf("Error updating user")
		}

		updatePendingNotificationByUserParams := db.UpdatePendingNotificationsByUserParams{
			DeviceToken: dbUtil.SqlNullString(input.DeviceToken),
			UserID:      user.ID,
		}

		err = r.PendingNotificationRepository.UpdatePendingNotificationsByUser(backgroundCtx, updatePendingNotificationByUserParams)

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
	user := middleware.GetCtxUser(ctx, r.UserRepository)

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
