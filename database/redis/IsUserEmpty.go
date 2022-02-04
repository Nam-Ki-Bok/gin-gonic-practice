package redis

func IsUserEmpty(email string) bool {
	err := CLIENT.HGet("users", email).Err()

	if err != nil {
		return true
	} else {
		return false
	}
}
