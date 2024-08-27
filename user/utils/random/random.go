package random

import (
	"crypto/rand"
	"math"
	"math/big"
)

const userIDLengths = 6

func GenerateUserID() int {
	min := int64(math.Pow(10, float64(userIDLengths)-1))
	max := int64(math.Pow(10, float64(userIDLengths))) - 1

	randomNumber, err := rand.Int(rand.Reader, big.NewInt(max-min))
	if err != nil {
		panic(err)
	}

	number := int(randomNumber.Int64()) + int(min)

	return number
}
