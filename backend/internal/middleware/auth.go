package middleware

import (
	"net/http"
	"context"
	"strings"

	"github.com/inRiza/popout/internal/model"
	"github.com/inRiza/popout/internal/pkg/httputil"
)

type contextKey string

const userContextKey = contextKey("user")

type SessionValidator interface {
	ValidateSession(ctx context.Context, token string) (*model.User, error)
}

func ValidateSession(svc SessionValidator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := extractToken(r)

			if token == "" {
				httputil.Error(w, http.StatusUnauthorized, "Unauthorized: missing token")
				return
			}

			user, err := svc.ValidateSession(r.Context(), token)
			if err != nil {
				httputil.Error(w, http.StatusUnauthorized, "Unauthorized: invalid or expired token")
				return
			}

			ctx := context.WithValue(r.Context(), userContextKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func userFromContext(ctx context.Context) (*model.User, bool) {
	user, ok := ctx.Value(userContextKey).(*model.User)
	if !ok {
		return nil, false
	}

	return user, true
}

func extractToken(r *http.Request) string {
	bearer := r.Header.Get("Authorization")
	if strings.HasPrefix(bearer, "Bearer ") {
		return strings.TrimPrefix(bearer, "Bearer ")
	}

	if cookie, err := r.Cookie("session_token"); err == nil {
		return cookie.Value
	}

	return ""
}