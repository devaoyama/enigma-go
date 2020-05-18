package enigmaMachine

import (
	"math/rand"
	"time"
)

func generateMapAlphabet(alphabet string) []string {
	var mapAlphabet []string
	for _, v := range alphabet {
		mapAlphabet = append(mapAlphabet, string(v))
	}
	return mapAlphabet
}

func getRandomAlphabet() []string {
	alphabet := make([]string, len(ALPHABET))
	copy(alphabet, ALPHABET)
	for i := len(alphabet) - 1; i >= 0; i-- {
		rand.Seed(time.Now().UnixNano())
		j := rand.Intn(i + 1)
		alphabet[i], alphabet[j] = alphabet[j], alphabet[i]
	}
	return alphabet
}

func getReflectAlphabet() []string {
	alphabet := make([]string, len(ALPHABET))
	copy(alphabet, ALPHABET)
	var indexes = make([]int, len(alphabet))
	for i := range alphabet {
		indexes[i] = i
	}

	length := len(alphabet) / 2
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		num1 := rand.Intn(len(indexes) / 2)
		num2 := rand.Intn(len(indexes) / 2)
		val1 := indexes[num1]
		indexes = append(indexes[:num1], indexes[num1+1:]...)
		val2 := indexes[num2]
		indexes = append(indexes[:num2], indexes[num2+1:]...)
		alphabet[val1], alphabet[val2] = alphabet[val2], alphabet[val1]
	}
	return alphabet
}
