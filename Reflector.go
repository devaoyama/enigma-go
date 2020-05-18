package enigmaMachine

import "log"

type Reflector struct {
	reflectorMap map[string]string
}

func NewReflector(mapAlphabet []string) *Reflector {
	var reflectorMap = make(map[string]string, len(mapAlphabet))
	for i, v := range mapAlphabet {
		reflectorMap[ALPHABET[i]] = v
	}

	reflector := &Reflector{
		reflectorMap: reflectorMap,
	}

	return reflector
}

func (Reflector *Reflector) reflect(indexNum int) int {
	char := Reflector.reflectorMap[ALPHABET[indexNum]]
	charIndexNum, err := getAlphabetIndex(char)
	if err != nil {
		log.Fatalln(err)
	}
	return charIndexNum
}
