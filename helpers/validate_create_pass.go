package helpers

import "regexp"

func ValidatePass(p string) bool {
	if len(p) < 8 {
		return false
	}
	match, _ := regexp.MatchString("[A-Z]", p)
	if !match {
		return false
	}

	match, _ = regexp.MatchString("[!@#$%^&*()_+{}|:\"<>?]", p)
	if !match {
		return false
	}

	match, _ = regexp.MatchString("[0-9]", p)
	if !match {
		return false
	}

	return true
}
