package errs

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

type Err struct {
	Status  int    `json:"-"`
	Code    Code   `json:"code"`
	Message string `json:"message"`
	Errors  string `json:"errors,omitempty"`
}

func (e *Err) Error() string {
	return e.Message
}

func NewBind(err error) Err {
	return Err{
		Status:  http.StatusBadRequest,
		Code:    Bind,
		Message: viper.GetString(fmt.Sprintf("error.%d", Bind)),
		Errors:  err.Error(),
	}
}

func NewValidate(err error) Err {
	return Err{
		Status:  http.StatusUnprocessableEntity,
		Code:    Validate,
		Message: viper.GetString(fmt.Sprintf("error.%d", Validate)),
		Errors:  err.Error(),
	}
}

func NewUnknown(err error) Err {
	return Err{
		Status:  http.StatusInternalServerError,
		Code:    Unknown,
		Message: viper.GetString(fmt.Sprintf("error.%d", Unknown)),
		Errors:  err.Error(),
	}
}
