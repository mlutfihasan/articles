package helper

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func RegisterValidation(validate *validator.Validate) {
	validate.RegisterValidation("status", validateStatus)
	validate.RegisterValidation("time", validateTime)
	validate.RegisterValidation("type", validateType)
}

func validateStatus(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	const pattern = "(INPUT|CONFIRM|CONFIRM HRD|REJECT|REJECT HRD|CANCEL LEAVE)"
	regex := regexp.MustCompile(pattern)
	result := regex.MatchString(value)
	return result
}

func validateTime(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	const pattern = "(MORNING-DAYLIGHT|DAYLIGHT-AFTERNOON)"
	regex := regexp.MustCompile(pattern)
	result := regex.MatchString(value)
	return result
}

func validateType(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	const pattern = "(0.5|1.0|2.0|3.0|10.0|30.0)"
	regex := regexp.MustCompile(pattern)
	result := regex.MatchString(value)
	return result
}
