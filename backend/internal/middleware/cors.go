package middleware

import (
	"net/http"

	"github.com/inRiza/popout/internal/config"
	"github.com/inRiza/popout/internal/pkg/httputil"
)

func CORS(cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			// in production, restrict actual domains
			if cfg.IsProd() {
				// TODO
				Origins := map[string]bool {
					"https://": true,
				}

				if _, ok := Origins[origin]; ok {
					w.Header().Set("Access-Control-Allow-Origin", origin)
				} else {
					httputil.Error(w, http.StatusForbidden, "Forbidden: invalid origin CORS issue")
					return
				}
			} else {
				w.Header().Set("Access-Control-Allow-Origin", "*")
			}

			// THIS ARE ALL THE ALLOWED HEADERS FOR THE CORS
			// allowed http methods
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			// allowed headers
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			// allowed credentials
			w.Header().Set("Access-Control-Allow-Credentials", "true") // could access JWT via cookie

			// this will handle the options request without running the next handler resulting in 404 err
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}