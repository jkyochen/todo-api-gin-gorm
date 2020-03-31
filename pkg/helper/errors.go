package helper

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var (
	ErrBadRequest     = InnError{Code: 101, Message: "Bad Input Request"}
	ErrNotLogin       = InnError{Code: 102, Message: "User Not Login"}
	ErrNotAuth        = InnError{Code: 103, Message: "User Can't Auth"}
	ErrUserExist      = InnError{Code: 104, Message: "User Exist"}
	ErrNotExist       = InnError{Code: 105, Message: "Todo Not Exist"}
	ErrInternalServer = InnError{Code: 500, Message: "Internal Server Error"}
)

// InnError is an error implementation that includes a time and message.
type InnError struct {
	Code    int
	Message string
}

func (e InnError) Error() string {
	return fmt.Sprintf("Error Code: %d, Error Message: %s", e.Code, e.Message)
}

// IsInnError is check error type
func IsInnError(err error) bool {
	_, ok := err.(InnError)
	return ok
}

func ErrorJSON(c *gin.Context, code int, err InnError) {
	log.Error().Err(err).Msg("json error")

	c.AbortWithStatusJSON(
		code,
		gin.H{
			"code":  err.Code,
			"error": err.Error(),
		},
	)
}
