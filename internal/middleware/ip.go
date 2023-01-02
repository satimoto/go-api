package middleware

import (
	"context"
	"net/http"

	"github.com/satimoto/go-api/internal/util"
)

func IpContext() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
			ctx := request.Context()
			ipAddress := util.GetIPAddress(request)

			if len(ipAddress) > 0 {
				ctx = context.WithValue(ctx, "ip_address", &ipAddress)
			}

			next.ServeHTTP(rw, request.WithContext(ctx))
		})
	}
}

func GetIPAddress(ctx context.Context) *string {
	ctxIpAddress := ctx.Value("ip_address")

	if ctxIpAddress != nil {
		return ctxIpAddress.(*string)
	}

	return nil
}