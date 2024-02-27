package response

import (
	"encoding/json"
	"net/http"
	"time"

	"nasdaqvfs/pkg/errors"

	"github.com/gin-gonic/gin"
)

const (
	CodeOK = 0
)

const (
	MessageOK = "Success"
)

type Response struct {
	Code       int    `json:"code"`
	Message    string `json:"message,omitempty"`
	Result     any    `json:"result,omitempty"`
	StatusCode int    `json:"-"`
}
type MapResponse map[error]Response

func WithOK(c *gin.Context, data any) {
	WithCode(c, http.StatusOK, data)
}

func WithNoContent(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}

func WithCode(c *gin.Context, code int, data any) {
	c.JSON(code, Response{
		Code:    CodeOK,
		Message: MessageOK,
		Result:  data,
	})
}

func WithError(c *gin.Context, err error) {
	c.JSON(errors.HTTPErrorResponse(err))
}

func WithErrorAndMap(c *gin.Context, err error, mapErr MapResponse) {
	if res, ok := mapErr[err]; ok {
		status := res.StatusCode
		if status <= 0 {
			status = http.StatusBadRequest
		}

		message := res.Message
		if message == "" {
			message = err.Error()
		}
		c.JSON(status, errors.NewError(res.Code, message, err))
		return
	}
	c.JSON(errors.HTTPErrorResponse(err))
}

// Date Response
func marshalTime(dt time.Time, format string) ([]byte, error) {
	t := dt.Format(format)
	dstr, err := json.Marshal(t)
	if err != nil {
		return []byte{}, err
	}

	return dstr, nil
}

// DateResponse is a custom time.Time type that marshals to and from JSON in YYYY-MM-DD format
type DateResponse time.Time

// MarshalJSON implements the json.Marshaler interface.
func (d DateResponse) MarshalJSON() ([]byte, error) {
	return marshalTime(time.Time(d), time.DateOnly)
}

// DateTimeResponse is a custom time.Time type that marshals to and from JSON in YYYY-MM-DD HH:MM:SS format
type DateTimeResponse time.Time

// MarshalJSON implements the json.Marshaler interface.
func (d DateTimeResponse) MarshalJSON() ([]byte, error) {
	return marshalTime(time.Time(d), time.DateTime)
}
