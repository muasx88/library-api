package validation

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/muasx88/library-api/internal/response"
)

func ValidateStruct(s interface{}) (bool, []*response.IError) {
	var errors []*response.IError

	validate := validator.New()

	// English is the fallback locale
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	translator, _ := uni.GetTranslator("en")

	// Register English translations
	en_translations.RegisterDefaultTranslations(validate, translator)

	err := validate.Struct(s)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			var el response.IError
			el.Field = strings.ToLower(e.Field())
			el.Tag = e.Tag()
			el.Message = e.Translate(translator)
			errors = append(errors, &el)
		}
		return false, errors
	}
	return true, nil
}

func isLengthValid(password string) bool {
	return len(password) >= 8
}

func containsUppercase(password string) bool {
	uppercase := regexp.MustCompile(`[A-Z]`)
	return uppercase.MatchString(password)
}

func containsLowercase(password string) bool {
	lowercase := regexp.MustCompile(`[a-z]`)
	return lowercase.MatchString(password)
}

func containsNumber(password string) bool {
	number := regexp.MustCompile(`[0-9]`)
	return number.MatchString(password)
}

func containsSpecial(password string) bool {
	special := regexp.MustCompile(`[!@#\$%\^&\*]`)
	return special.MatchString(password)
}

func ValidatePassword(password string, field string) error {

	if !isLengthValid(password) {
		return fmt.Errorf("%s must be at least 8 characters long", field)
	}

	if !containsUppercase(password) {
		return fmt.Errorf("%s must contain at least one uppercase letter", field)
	}

	if !containsLowercase(password) {
		return fmt.Errorf("%s must contain at least one lowercase letter", field)
	}

	if !containsNumber(password) {
		return fmt.Errorf("%s must contain at least one number", field)
	}

	if containsSpecial(password) {
		return fmt.Errorf("%s must not contain special character", field)
	}

	return nil
}

func ValidateISBN(isbn string, field string) error {
	isbn = removeDashesAndSpaces(isbn)

	// Check if the ISBN follows the correct format
	match, err := regexp.MatchString(`^\d{10}|\d{13}$`, isbn)
	if err != nil {
		return fmt.Errorf("%s invalid", field)
	}

	if match == false {
		return fmt.Errorf("%s invalid", field)
	}

	return nil
}

func removeDashesAndSpaces(s string) string {
	re := regexp.MustCompile(`[-\s]`)
	return re.ReplaceAllString(s, "")
}

func ValidateNoWhiteSpace(fl validator.FieldLevel) bool {
	// Get the field value
	value := fl.Field().String()

	// Remove leading and trailing spaces
	value = strings.TrimSpace(value)

	// Check if the value contains only spaces
	return value != ""
}
