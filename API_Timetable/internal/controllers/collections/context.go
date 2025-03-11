package collections

import (
	"API_Timetable/internal/models"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uid := chi.URLParam(r, "id")
		if uid == "" {
			logrus.Errorf("parsing error : uid empty")
			customError := &models.CustomError{
				Message: "cannot parse uid, it is empty",
				Code:    http.StatusUnprocessableEntity,
			}
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
			return
		}

		ctx := context.WithValue(r.Context(), "collectionId", uid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
