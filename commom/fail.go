package commom

import (
	"errors"
	"math/rand"
)

const (
	FAIL_PERCENTAGE = 50
)

func Fail() error {
	randomNumber := rand.Intn(100) + 1

	if randomNumber <= FAIL_PERCENTAGE {
		return errors.New("simulated failure")
	}

	return nil
}
