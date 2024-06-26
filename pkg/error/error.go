package pkgerr

import (
	"errors"
	"github.com/lib/pq"
)

func ErrorCode(err error) string {
	var pgErr *pq.Error

	if err == nil {
		return ""
	}

	if errors.As(err, &pgErr) {
		return string(pgErr.Code)
	}

	return err.Error()
}
