package authentication

import (
	"context"
	"net/http"
	"strings"
)

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
			ctx := request.Context()
			authentication := strings.Split(request.Header.Get("Authentication"), " ")

			if len(authentication) == 2 {
				token := authentication[2]

				if ok, claims := VerifyToken(token); ok {
					ctx = context.WithValue(ctx, "user_id", claims["user_id"])
				}
			}

			next.ServeHTTP(rw, request.WithContext(ctx))
		})
	}
}

func GetUserId(ctx context.Context) *int64 {
	userId := ctx.Value("user_id")

	if userId != nil {
		return userId.(*int64)
	}

	return nil
}
