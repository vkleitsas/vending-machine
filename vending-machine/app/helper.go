package app

import "golang.org/x/crypto/bcrypt"

func containsInt(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func coinsBreakdown(amount int) map[int]int {
	breakdown := make(map[int]int)

	breakdown[100] = amount / 100
	amount = amount % 100
	breakdown[50] = amount / 50
	amount = amount % 50
	breakdown[20] = amount / 20
	amount = amount % 20
	breakdown[10] = amount / 10
	amount = amount % 10
	breakdown[5] = amount / 5

	return breakdown
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
