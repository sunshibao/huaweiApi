package validator

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"unicode/utf8"

	"github.com/pkg/errors"
)

type ValidateFunc func() error

type ValidateWrapper struct {
	items []ValidateFunc
}

func NewWrapper(fns ...ValidateFunc) *ValidateWrapper {
	return &ValidateWrapper{
		items: fns,
	}
}

func (vw *ValidateWrapper) AddValidateFunc(functions ...ValidateFunc) {
	for _, function := range functions {
		vw.items = append(vw.items, function)
	}
}

func (vw *ValidateWrapper) Validate() error {
	for _, v := range vw.items {
		if err := v(); err != nil {
			return err
		}
	}
	return nil
}

func ValidLength(str, keyName string, minimum, maximum int) error {
	length := utf8.RuneCountInString(str)
	if maximum > 0 && length > maximum {
		return fmt.Errorf("%q '%s' is too long", keyName, str)
	}
	if minimum > 0 && length < minimum {
		return fmt.Errorf("%q '%s' is too short", keyName, str)
	}
	return nil
}

func ValidateString(str, keyName string, minimum, maximum int) ValidateFunc {
	return func() error {
		return ValidLength(str, keyName, minimum, maximum)
	}
}

func ValidateStringPointer(str *string, keyName string, minimum, maximum int) ValidateFunc {
	return func() error {
		if str == nil {
			return nil
		}
		return ValidateString(*str, keyName, minimum, maximum)()
	}
}

func ValidateSameString(str1, keyName1, str2, keyName2 string) ValidateFunc {
	return func() error {
		if str1 == str2 {
			return nil
		}
		return fmt.Errorf("%q & %q must equal", keyName1, keyName2)
	}
}

func ValidateRegexp(re *regexp.Regexp, str, keyName string) ValidateFunc {
	return func() error {
		if !re.MatchString(str) {
			return fmt.Errorf("%q illegal", keyName)
		}
		return nil
	}
}

func ValidateEmail(email string) bool {
	b, err := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$", email)
	return b && (err == nil)
}

func ValidateMobile(mobile string) bool {
	// Regular Expression By https://github.com/VincentSit/ChinaMobilePhoneNumberRegex/blob/master/README-CN.md
	re := regexp.MustCompile(`^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[01356789]\d{2}|4(?:0\d|1[0-2]|9\d))|9[189]\d{2}|6[567]\d{2}|4(?:[14]0\d{3}|[68]\d{4}|[579]\d{2}))\d{6}$`)
	return re.MatchString(mobile)
}

func ValidateStringOptions(str string, keyName string, options []string) ValidateFunc {

	return func() error {
		for _, option := range options {
			if option == str {
				return nil
			}
		}
		return fmt.Errorf("%s not in specify options", keyName)
	}
}

func ValidateStringArrayOptions(strList []string, keyName string, options []string) ValidateFunc {

	return func() error {

		if len(strList) == 0 {
			return fmt.Errorf("%s is empty", keyName)
		}

		for _, str := range strList {

			if err := ValidateStringOptions(str, keyName, options)(); err != nil {
				return err
			}
		}

		return nil
	}
}

func ValidateIntRange(value int, keyName string, minimum, maximum int) ValidateFunc {

	return func() error {

		if minimum <= value && value <= maximum {
			return nil
		}

		return fmt.Errorf("%s out of range: minimum: %d, maximum: %d", keyName, minimum, maximum)
	}
}

func ValidateUint64Range(value uint64, keyName string, minimum, maximum uint64) ValidateFunc {

	return func() error {

		if minimum <= value && value <= maximum {
			return nil
		}

		return fmt.Errorf("%s out of range: minimum: %d, maximum: %d", keyName, minimum, maximum)
	}
}

func ValidateUint32Range(value uint32, keyName string, minimum, maximum uint32) ValidateFunc {

	return func() error {

		if minimum <= value && value <= maximum {
			return nil
		}

		return fmt.Errorf("%s out of range: minimum: %d, maximum: %d", keyName, minimum, maximum)
	}
}

func ValidateUint8Range(value uint8, keyName string, minimum, maximum uint8) ValidateFunc {

	return func() error {

		if minimum <= value && value <= maximum {
			return nil
		}

		return fmt.Errorf("%s out of range: minimum: %d, maximum: %d", keyName, minimum, maximum)
	}
}

func ValidateIP(value string, keyName string) ValidateFunc {

	return func() error {

		ip := net.ParseIP(value)
		if ip == nil {
			return fmt.Errorf("%s is invalid ip", keyName)
		}

		return nil
	}
}

func ValidateUint64PositiveInteger(value uint64, keyName string) ValidateFunc {

	return func() error {

		if value > 0 {
			return nil
		}

		return fmt.Errorf("%s must be a positive integer", keyName)
	}
}

// 检验字符串内容值是否正整数
func ValidateStringWithUint64PositiveInteger(value string, keyName string) ValidateFunc {

	return func() error {

		if len(value) == 0 {
			return fmt.Errorf("%s must be a non-empty positive integer string value", keyName)
		}

		number, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("%s must be a positive integer string value", keyName))
		}

		return ValidateUint64PositiveInteger(number, keyName)()
	}
}
