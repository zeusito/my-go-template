package validation

import (
	"encoding/json"
	"net/http"

	"github.com/zeusito/my-go-template/pkg/errors"

	"github.com/go-playground/validator/v10"
	"github.com/zeusito/my-go-template/pkg/router"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func ParseJSONAndValidate(r *http.Request, w http.ResponseWriter, dest *any) {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {

		router.RenderError(r.Context(), w, errors.New(errors.ErrorCodeBadRequest, "Invalid request body"))
	}

	if err := validate.Struct(dest); err != nil {
		router.RenderError(r.Context(), w, errors.New(errors.ErrorCodePreconditionFailed, err.Error()))
	}
}
