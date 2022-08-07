package middleware

import (
	"leexsh/file_sys/core/helper"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

// midwart sign
func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("UnAuthorize"))
			return
		}
		uc, err := helper.AnalyzeToke(auth)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		r.Header.Set("UserId", string(rune(uc.Id)))
		r.Header.Set("UserIdentity", uc.Identity)
		r.Header.Set("UserName", uc.Name)
		next(w, r)
	}
}
