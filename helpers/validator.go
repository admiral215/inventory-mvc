package helpers

import (
	"fmt"
	"github.com/beego/beego/v2/core/validation"
)

func ValidateDto(obj interface{}) any {
	valid := validation.Validation{}
	ok, err := valid.Valid(obj)

	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	if !ok {
		return valid.Errors
	}

	return nil
}
