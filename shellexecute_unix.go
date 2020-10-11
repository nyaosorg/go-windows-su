// +build !windows

package su

import (
	"errors"
)

func (i Param) shellExecute() (int, error) {
	return 0, errors.New("not supported")
}
