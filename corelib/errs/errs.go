package errs

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

type Err struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
	Errors  string `json:"errors,omitempty"`
}

func (e *Err) Error() string {
	return e.Message
}

func New(code Code, err error) Err {
	return Err{
		Code:    code,
		Message: viper.GetString(strconv.Itoa(int(code))),
		Errors:  err.Error(),
	}
}

func NewBind(err error) Err {
	return Err{
		Code:    Bind,
		Message: viper.GetString(fmt.Sprintf("error.%d", Bind)),
		Errors:  err.Error(),
	}
}

func NewValidate(err error) Err {
	return Err{
		Code:    Validate,
		Message: viper.GetString(fmt.Sprintf("error.%d", Validate)),
		Errors:  err.Error(),
	}
}

func NewUnknown(err error) Err {
	return Err{
		Code:    Unknown,
		Message: viper.GetString(fmt.Sprintf("error.%d", Unknown)),
		Errors:  err.Error(),
	}
}
