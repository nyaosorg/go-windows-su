// +build !windows

package su

import (
	"errors"
)

func isElevated() (bool, error) {
	return false, errors.New("not supported")
}
