package generator

import (
	"math/rand"
	"strings"
)

const letterBytes = "1234567890"

func GenerateProductId(category string, subcategory string, n int) string {
	firstId := strings.ToUpper(category[0:3])
	secondId := strings.ToUpper(subcategory[0:2])

	thirdId := make([]byte, n)
	for i := range thirdId {
		thirdId[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	fullId := firstId + "-" + secondId + "-" + string(thirdId)

	return fullId
}

func GenerateCustomerId(name string, n int) string {
	split := strings.Split(name, " ")
	secondId := make([]byte, n)

	for i := range secondId {
		secondId[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	if len(split) == 1 {
		firstId := strings.ToUpper(name[len(name)-1:])
		fullId := firstId + "-" + string(secondId)

		return fullId
	}

	firstId := strings.ToUpper(split[0][0:1]) + strings.ToUpper(split[1][0:1])
	fullId := firstId + "-" + string(secondId)

	return fullId
}
