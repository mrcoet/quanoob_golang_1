package work

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func contains(s [16]string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func DropRock(allawance int, startWith string) string {
	hexaValue := [16]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(hexaValue), func(i, j int) { hexaValue[i], hexaValue[j] = hexaValue[j], hexaValue[i] })
	var checkSum [16]string
	for i, value := range hexaValue {
		checkSum[i] = strings.Repeat(value, allawance)
	}
	randomHex := startWith
	for {
		rand.Seed(time.Now().UnixNano())
		randomHex += hexaValue[rand.Intn(len(hexaValue))]
		lenTo := len(randomHex) - 4
		if lenTo >= 0 {
			if contains(checkSum, randomHex[lenTo:]) {
				randomHex = randomHex[:len(randomHex)-1]
			}
		}
		if len(randomHex) >= 64 {
			break
		}
	}
	fmt.Println(randomHex)
	return randomHex
}

func DropRockTestOne() string {
	hexaValues := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(hexaValues), func(i, j int) { hexaValues[i], hexaValues[j] = hexaValues[j], hexaValues[i] })
	return strings.Join(hexaValues, "")
}
