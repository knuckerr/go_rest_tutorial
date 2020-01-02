package middlewares

import (
	"context"
	"github.com/go-chi/chi"
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
			responses.JSON(w, http.StatusUnauthorized, map[string]string{"error": "header must have Authorization Bearer {token}"})
			return
		}
		reqToken = strings.TrimSpace(splitToken[1])
		var claims = &auth.Claims{}
		err := auth.Vaildtoken(reqToken, claims)
		if err != nil {
			responses.JSON(w, http.StatusUnauthorized, map[string]string{"error": "user needs to be signed in to access this service"})
			return
		}
		ctw := context.WithValue(r.Context(), "claim", claims)
		next.ServeHTTP(w, r.WithContext(ctw))
	})
}

func OwnerId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := r.Context().Value("claim")
		if data == nil {
			responses.JSON(w, http.StatusUnauthorized, map[string]string{"error": "need to be authorized first"})
			return
		}
		claim, ok := data.(*auth.Claims)
		if !ok {
			responses.JSON(w, http.StatusUnauthorized, map[string]string{"error": "context need to be claim struct"})
			return
		}
		item_id := chi.URLParam(r, "id")
		if claim.Role != "admin" {
			if item_id != claim.Id {
				responses.JSON(w, http.StatusUnauthorized, map[string]string{"error": "you are not the owner of this id"})
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
