package utils

import (
	"context"
	"errors"
	"net"
	"net/http"
	"strings"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e APIError) Error() string { return e.Message }

func MapDBError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return APIError{Code: http.StatusNotFound, Message: "data not found"}
	}

	var me *mysql.MySQLError
	if errors.As(err, &me) {
		switch me.Number {
		case 1062: // Duplicate entry
			if strings.Contains(me.Message, "email") {
				return APIError{Code: http.StatusConflict, Message: "email already registered"}
			}
			return APIError{Code: http.StatusConflict, Message: "duplicate entry"}
		case 1451:
			return APIError{Code: http.StatusConflict, Message: "cannot modify: constrained by related records"}
		case 1452:
			return APIError{Code: http.StatusBadRequest, Message: "invalid data: parent row not found"}
		case 1048:
			return APIError{Code: http.StatusBadRequest, Message: "required field cannot be null"}
		case 1406:
			return APIError{Code: http.StatusBadRequest, Message: "data too long for column"}
		case 1064:
			return APIError{Code: http.StatusBadRequest, Message: "bad request (sql syntax)"}
		default:
			return APIError{Code: http.StatusBadRequest, Message: me.Message}
		}
	}

	if errors.Is(err, context.DeadlineExceeded) {
		return APIError{Code: http.StatusGatewayTimeout, Message: "database timeout"}
	}
	var ne net.Error
	if errors.As(err, &ne) && ne.Timeout() {
		return APIError{Code: http.StatusGatewayTimeout, Message: "database timeout"}
	}

	return APIError{Code: http.StatusInternalServerError, Message: err.Error()}
}
