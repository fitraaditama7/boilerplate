package middleware

import (
	"arka/cmd/lib/customError"
	"arka/pkg/auth"
	"arka/pkg/casbin"
	"arka/pkg/response"
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

type contextKey string

func (c contextKey) String() string {
	return "myPackage context key " + string(c)
}

func RequiresAccessToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			logrus.Error(err)
			response.Error(w, err)
			return
		}
		claims, err := auth.ExtractTokenMetadata(r)
		if err != nil {
			logrus.Error(err)
			response.Error(w, err)
		}
		contextClaims := contextKey("claims")
		ctx := context.WithValue(r.Context(), contextClaims, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func RequiresAuthorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contextClaims := contextKey("claims")
		value := r.Context().Value(contextClaims).(*auth.AccessDetails)
		roleID := value.RoleID

		ok, err := casbin.CheckPolicy(&casbin.RoleData{
			Role:   roleID,
			Path:   r.URL.Path,
			Method: r.Method,
		})
		if err != nil {
			logrus.Error(err)
			response.Error(w, err)
			return
		}

		if ok {
			next.ServeHTTP(w, r)
		} else {
			logrus.Error(customError.ErrNotAuthorize.Detail)
			response.Error(w, customError.ErrNotAuthorize)
			return
		}
	}
}

func RequiresCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Request-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, PUT, DELETE")

		if r.Method == "OPTIONS" {
			response.Error(w, customError.ErrNoContent)
			return
		}

		next.ServeHTTP(w, r)
	}
}
