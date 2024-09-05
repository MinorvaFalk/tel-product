package exception

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"tel/product/internal/model"
	"tel/product/pkg/logger"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	InternalServerError = "Internal Server Error"
	NotFoundError       = "Not Found Error"
)

func EchoErrorHandler(err error, c echo.Context) {
	logger.Zap().Error("api", zap.Error(err))

	response := model.HTTPResponse{
		Status:  http.StatusInternalServerError,
		Message: InternalServerError,
	}

	var exError Error
	if errors.As(err, &exError) {
		response.Status = exError.Status
		response.Message = exError.Message

		c.JSON(exError.Status, response)
		return
	}

	if echoError, ok := err.(*echo.HTTPError); ok {
		response.Status = echoError.Code
		response.Message = fmt.Sprintf("%v", echoError.Message)

		c.JSON(exError.Status, response)
		return
	}

	c.JSON(http.StatusInternalServerError, response)
}

type Error struct {
	Source  string `json:"source"`
	Status  int    `json:"status"`
	Message string `json:"message"`

	// Store original error
	Err error
}

func (e Error) Error() string {
	return e.Err.Error()
}

func NewDBQueryError(e error, msg ...string) error {
	err := Error{
		Source:  "repository",
		Status:  http.StatusInternalServerError,
		Message: InternalServerError,
		Err:     e,
	}

	if errors.Is(e, gorm.ErrRecordNotFound) {
		err.Status = http.StatusNotFound
		err.Message = NotFoundError

		return err
	}

	if len(msg) > 0 {
		err.Message = err.Message + ";" + strings.Join(msg, ";")
	}

	return err
}

func NewNotFoundError(source string, e error) error {
	return Error{
		Source:  source,
		Status:  http.StatusNotFound,
		Message: NotFoundError,
		Err:     e,
	}
}

func NewValidatonError(msg string, err error) error {
	return Error{
		Source:  "validation",
		Status:  http.StatusBadRequest,
		Message: msg,
		Err:     err,
	}
}
