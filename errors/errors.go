package errors

import "github.com/labstack/echo/v4"

const (
	NotFound        = "resource not found"
	AccessForbidden = "access forbidden"
)

type ResponseError struct {
	Errors map[string]interface{}
}

func NewError(err interface{}) ResponseError {
	e := ResponseError{}
	e.Errors = make(map[string]interface{})
	switch v := err.(type) {
	case *echo.HTTPError:
		e.Errors["body"] = v.Message
	case error:
		e.Errors["body"] = v.Error()
	default:
		e.Errors["body"] = v.(string)
	}
	return e
}

func NewErrorFromPredefined(message string) ResponseError {
	e := ResponseError{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = message
	return e
}
