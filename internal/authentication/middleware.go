package authentication

import (
	"context"
	"net/http"
	"strings"
)

func AuthorizationContext() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
			ctx := request.Context()
			authorization := strings.Split(request.Header.Get("Authorization"), " ")

			if len(authorization) == 2 {
				token := authorization[1]

				if ok, claims := VerifyToken(token); ok {
					userId := int64(claims["user_id"].(float64))
					ctx = context.WithValue(ctx, "user_id", &userId)
				}
			}

			next.ServeHTTP(rw, request.WithContext(ctx))
		})
	}
}

func GetUserId(ctx context.Context) *int64 {
	ctxUserId := ctx.Value("user_id")

	if ctxUserId != nil {
		return ctxUserId.(*int64)
	}

	return nil
}
