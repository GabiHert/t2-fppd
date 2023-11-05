package commom

const MAX_RETRIES = 5

func Retry[T interface{}](fn func() (T, error)) (T, error) {
	var err error
	var t T
	for i := 0; i < MAX_RETRIES; i++ {
		t, err = fn()
		if err == nil {
			return t, nil
		}
	}

	return t, err
}
