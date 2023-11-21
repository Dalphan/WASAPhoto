package utils

func ValidateUsername(username string) error {
	if len(username) == 0 {
		return ErrUsernameMissing
	}
	if len(username) < 3 || len(username) > 16 {
		return ErrUsernameNotValid
	}
	return nil
}
