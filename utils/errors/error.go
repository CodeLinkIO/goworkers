package errors

import "fmt"

func ConvertPanicToError(r interface{}) error {
	switch r.(type) {
	case error:
		err, _ := r.(error)
		return err
	case string:
		return fmt.Errorf(r.(string))
	default:
		return fmt.Errorf("An error")
	}
}
