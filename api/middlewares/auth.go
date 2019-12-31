package middlewares

import (
	"github.com/knuckerr/go_rest/api/auth"
	"github.com/knuckerr/go_rest/api/responses"
	"net/http"
	"strings"
)

func AuthenticationRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer")
		if len(splitToken) == 1 {
			responses.JSON(w, http.StatusUnauthorized, map[string]string{"error": "user needs to be signed in to access this service"})
			return
		}
		reqToken = strings.TrimSpace(splitToken[1])
		var claims = &auth.Claims{}
		err := auth.Vaildtoken(reqToken, claims)
		if err != nil {
			responses.JSON(w, http.StatusUnauthorized, map[string]string{"error": "user needs to be signed in to access this service"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
