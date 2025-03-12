package resources

import (
	"API_Config/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uidStr := chi.URLParam(r, "id")
		if uidStr == "" {
			logrus.Errorf("parsing error : uidStr empty")
			customError := &models.CustomError{
				Message: "cannot parse uidStr, it is empty",
				Code:    http.StatusUnprocessableEntity,
			}
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
			return
		}

		uid, err := strconv.Atoi(uidStr)
		if err != nil {
			logrus.Errorf("parsing error: %s", err.Error())
			customError := &models.CustomError{
				Message: fmt.Sprintf("cannot parse uid (%s) to int", uidStr),
				Code:    http.StatusUnprocessableEntity,
			}
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
			return
		}

		ctx := context.WithValue(r.Context(), "resourceId", uid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
