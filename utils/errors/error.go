package errors

import "fmt"

// ConvertPanicToError grabs panic in recover()
// and then returns an error based on different types that it catches
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
