package pkgerr

import "errors"

var (
	ErrConfigIsNil = errors.New("config is nil")
)
