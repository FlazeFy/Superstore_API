package generator

import (
	"math/rand"
	"strings"
)

const numberPack = "1234567890"
const letterPack = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateProductId(category string, subcategory string, n int) string {
	firstId := strings.ToUpper(category[0:3])
	secondId := strings.ToUpper(subcategory[0:2])

	thirdId := make([]byte, n)
	for i := range thirdId {
		thirdId[i] = numberPack[rand.Intn(len(numberPack))]
	}

	fullId := firstId + "-" + secondId + "-" + string(thirdId)

	return fullId
}

func GenerateCustomerId(name string, n int) string {
	split := strings.Split(name, " ")
	secondId := make([]byte, n)

	for i := range secondId {
		secondId[i] = numberPack[rand.Intn(len(numberPack))]
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

func GenerateOrderId(order_date string, n int, m int) string {
	split := strings.Split(order_date, "-")
	firstId := make([]byte, n)
	secondId := split[0]
	thirdId := make([]byte, m)

	for i := range firstId {
		firstId[i] = letterPack[rand.Intn(len(letterPack))]
	}

	for i := range thirdId {
		thirdId[i] = numberPack[rand.Intn(len(numberPack))]
	}

	fullId := string(firstId) + "-" + secondId + "-" + string(thirdId)

	return fullId
}
