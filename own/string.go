package own

import (
	"math/rand"
	"strconv"
)

func GenerateCode(length int) string {
	// length is the number of code length
	code := ""
	for i := 0; i < length; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}
	return code
}
