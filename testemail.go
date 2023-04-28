package goemail

import (
	"fmt"
	"regexp"
)

func IsEmail(s string) (string, error) {
	r, _ := regexp.Compile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
	if r.MatchString(s) {
		return "Valid Email", nil
	} else {
		return "", fmt.Errorf("Invalid email address: %s", s)
	}

}
