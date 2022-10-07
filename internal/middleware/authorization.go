package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/satimoto/go-api/internal/authentication"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/user"
	"github.com/satimoto/go-datastore/pkg/util"
)

func AuthorizationContext() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
			ctx := request.Context()
			authorization := strings.Split(request.Header.Get("Authorization"), " ")

			if len(authorization) == 2 {
				token := authorization[1]

				if ok, claims := authentication.VerifyToken(token); ok {
					userID := int64(claims["user_id"].(float64))
					ctx = context.WithValue(ctx, "user_id", &userID)
				}
			}

			next.ServeHTTP(rw, request.WithContext(ctx))
		})
	}
}

func GetUserID(ctx context.Context) *int64 {
	ctxUserID := ctx.Value("user_id")

	if ctxUserID != nil {
		return ctxUserID.(*int64)
	}

	return nil
}

func GetUser(ctx context.Context, r user.UserRepository) *db.User {
	operationCtx := graphql.GetOperationContext(ctx)
	ctxUser := operationCtx.Variables["user"]

	if ctxUser != nil {
		return ctxUser.(*db.User)
	}

	userID := GetUserID(ctx)

	if userID != nil {
		user, err := r.GetUser(ctx, *userID)

		if err != nil {
			util.LogOnError("API019", "Error retrieving user", err)
			log.Printf("API019: UserID=%v", userID)
			return nil
		}

		operationCtx.Variables["user"] = &user
		return &user
	}

	return nil
}
