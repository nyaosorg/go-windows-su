// +build !windows

package su

import (
	"errors"
)

func shellExecute(action, path, param, directory string) (int, error) {
	return 0, errors.New("not supported")
}
