package task9

import "errors"

func AcceptOnlyPositive(input int) error {
	if input < 0 {
		return errors.New("negative input not allowed")
	}
	return nil
}
