package helpers

import (
	"errors"
	"gin_app/dto"
	"regexp"
	"strings"
	"unicode"
)

func ValidateSignupData(user *dto.UserSignUpRequestDTO) error {
    if user == nil {
        return errors.New("user input is nil")
    }

    user.FirstName = strings.TrimSpace(user.FirstName)
    user.LastName = strings.TrimSpace(user.LastName)
    user.DateOfBirth = strings.TrimSpace(user.DateOfBirth)
    user.Gender = strings.TrimSpace(user.Gender)
    user.Email = strings.TrimSpace(user.Email)
    user.Password = strings.TrimSpace(user.Password)
    user.CountryCode = strings.TrimSpace(user.CountryCode)

    if user.FirstName == "" || user.LastName == "" || user.DateOfBirth == "" ||
        user.Gender == "" || user.Email == "" || user.Password == "" || user.CountryCode == "" {
        return errors.New("all fields must be non-empty")
    }

    nameRegex := regexp.MustCompile(`^[A-Za-z ]+$`)
    if !nameRegex.MatchString(user.FirstName) {
        return errors.New("first name must contain only alphabets")
    }
    if !nameRegex.MatchString(user.LastName) {
        return errors.New("last name must contain only alphabets")
    }

	if !nameRegex.MatchString(user.Gender) {
		return errors.New("gender can either be Male Or Female")
	}

	if(user.Gender != "Male" && user.Gender != "Female" && user.Gender != "male" && user.Gender != "female") {
		return errors.New("gender can either be Male Or Female")
	}
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
    if !emailRegex.MatchString(user.Email) {
        return errors.New("invalid email format")
    }


    if len(user.Password) < 6 {
        return errors.New("password must be at least 6 characters long")
    }

    var hasUpper, hasLower, hasDigit bool
    for _, char := range user.Password {
        switch {
        case unicode.IsUpper(char):
            hasUpper = true
        case unicode.IsLower(char):
            hasLower = true
        case unicode.IsDigit(char):
            hasDigit = true
        }
    }

    if !hasUpper || !hasLower || !hasDigit {
        return errors.New("password must contain at least 1 uppercase, 1 lowercase letter, and 1 digit")
    }

    return nil
}
