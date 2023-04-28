package goemail

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("hello")
	if err == nil {
		t.Error("Expected error for invalid email address")
	}
	_, err = IsEmail("hello@22323.com")
	if err == nil {
		t.Error("hello@22323.com is an email")
	}
	_, err = IsEmail("hello@ew")
	if err != nil {
		t.Error("hello@ew is an email")
	}
}
