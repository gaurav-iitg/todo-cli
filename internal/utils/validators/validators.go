package validators

import "strconv"

func ValidateNumber(s string) (int, error) {
	number, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return number, nil
}
