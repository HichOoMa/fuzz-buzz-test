package handler

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func validationErrorMessage(err error) string {
	verrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error()
	}

	messages := make([]string, 0, len(verrs))
	for _, fe := range verrs {
		field := strings.ToLower(fe.Field())
		var reason string
		switch fe.Tag() {
		case "required":
			reason = "is required"
		case "gt":
			reason = fmt.Sprintf("must be greater than %s", fe.Param())
		case "lte":
			reason = fmt.Sprintf("must be less than or equal to %s", fe.Param())
		default:
			reason = fmt.Sprintf("failed validation (%s)", fe.Tag())
		}
		messages = append(messages, fmt.Sprintf("%s: %s", field, reason))
	}

	return strings.Join(messages, "; ")
}
