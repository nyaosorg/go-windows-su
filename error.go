package su

import (
	"errors"
	"syscall"
)

const (
	_ERROR_ELEVATION_REQUIRED syscall.Errno = 740
)

// IsElevationRequired reports whether err is the error object "os/exec".Command.Run() returns when the elevation is required.
func IsElevationRequired(err error) bool {
	return errors.Is(err, _ERROR_ELEVATION_REQUIRED)
}
