package enigmaMachine

import (
	"errors"
	"log"
)

type PlugBoard struct {
	alphabet    []string
	forwardMap  map[string]string
	backwardMap map[string]string
}

func NewPlugBoard(mapAlphabet []string) *PlugBoard {
	plugBoard := &PlugBoard{
		alphabet: ALPHABET,
	}
	plugBoard.mapping(mapAlphabet)
	return plugBoard
}

func (PlugBoard *PlugBoard) getAlphabetIndex(c string) (int, error) {
	for i, v := range PlugBoard.alphabet {
		if c == v {
			return i, nil
		}
	}
	return 0, errors.New("おかしな値です")
}

func (PlugBoard *PlugBoard) mapping(mapAlphabet []string) {
	var forwardMap = make(map[string]string, len(mapAlphabet))
	for i, v := range PlugBoard.alphabet {
		forwardMap[v] = mapAlphabet[i]
	}
	var backwardMap = make(map[string]string, len(mapAlphabet))
	for i, v := range forwardMap {
		backwardMap[v] = i
	}
	PlugBoard.forwardMap = forwardMap
	PlugBoard.backwardMap = backwardMap
}

func (PlugBoard PlugBoard) Forward(indexNum int) int {
	char := PlugBoard.alphabet[indexNum]
	char = PlugBoard.forwardMap[char]
	charIndexNum, err := PlugBoard.getAlphabetIndex(char)
	if err != nil {
		log.Fatalln(err)
	}
	return charIndexNum
}

func (PlugBoard PlugBoard) Backward(indexNum int) int {
	char := PlugBoard.alphabet[indexNum]
	char = PlugBoard.backwardMap[char]
	charIndexNum, err := PlugBoard.getAlphabetIndex(char)
	if err != nil {
		log.Fatalln(err)
	}
	return charIndexNum
}
