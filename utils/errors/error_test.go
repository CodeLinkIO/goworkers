package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertPanicToError(t *testing.T) {
	err := ConvertPanicToError("Error 1")
	assert.Error(t, err, "Should covert string to error")
	assert.Equal(t, "Error 1", err.Error(), "The string message should remain")

	err = ConvertPanicToError(fmt.Errorf("Error 2"))
	assert.Error(t, err, "Should return error as it is")
	assert.Equal(t, "Error 2", ConvertPanicToError(err).Error(), "Error message should remain")

	err = ConvertPanicToError(1)
	assert.Error(t, err, "Should convert other type to error")
}
