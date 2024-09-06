package router

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/zeusito/my-go-template/pkg/errors"

	"github.com/go-chi/chi/v5/middleware"
)

// RenderJSON is a helper function to write a JSON response
func RenderJSON(ctx context.Context, w http.ResponseWriter, httpStatusCode int, payload any) {
	// Headers
	w.Header().Set(middleware.RequestIDHeader, middleware.GetReqID(ctx))
	w.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(httpStatusCode)
	_, _ = w.Write(js)
}

// RenderError Renders an error with some sane defaults.
func RenderError(ctx context.Context, w http.ResponseWriter, err errors.CustomError) {
	var httpStatusCode int

	switch err.GetCode() {
	case errors.ErrorCodeBadRequest:
		httpStatusCode = http.StatusBadRequest
	case errors.ErrorCodeNotFound:
		httpStatusCode = http.StatusBadRequest
	case errors.ErrorCodeForbidden:
		httpStatusCode = http.StatusForbidden
	case errors.ErrorCodePreconditionFailed:
		httpStatusCode = http.StatusBadRequest
	case errors.ErrorCodeInternalError:
		httpStatusCode = http.StatusInternalServerError
	default:
		httpStatusCode = http.StatusInternalServerError
	}

	RenderJSON(ctx, w, httpStatusCode, err)
}
