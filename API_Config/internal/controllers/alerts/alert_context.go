package alerts

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := chi.URLParam(r, "id")
		ctx := context.WithValue(r.Context(), "Email", email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
