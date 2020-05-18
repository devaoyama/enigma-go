package enigmaMachine

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

var ALPHABET = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func getAlphabetIndex(c string) (int, error) {
	for i, v := range ALPHABET {
		if c == v {
			return i, nil
		}
	}
	return 0, errors.New("おかしな値です")
}

type Machine struct {
	plugBoard *PlugBoard
	rotors    []*Rotor
	reflector *Reflector
}

func newMachine(plugBoard *PlugBoard, rotors []*Rotor, reflector *Reflector) *Machine {
	machine := &Machine{
		plugBoard,
		rotors,
		reflector,
	}
	return machine
}

func InitMachine(rotorsAmount int) *Machine {
	plugBoard := NewPlugBoard(getRandomAlphabet())
	var rotors []*Rotor
	for i := 0; i < rotorsAmount; i++ {
		rand.Seed(time.Now().UnixNano())
		rotors = append(rotors, NewRotor(getRandomAlphabet(), rand.Intn(5)))
	}
	reflector := NewReflector(getReflectAlphabet())
	machine := newMachine(plugBoard, rotors, reflector)

	return machine
}

func (machine *Machine) Encrypt(text string) string {
	var encryptedText string
	for _, v := range text {
		encryptedText += machine.through(string(v))
	}
	return encryptedText
}

func (machine *Machine) Decrypt(text string) string {
	for _, rotor := range machine.rotors {
		rotor.reset()
	}
	var encryptedText string
	for _, v := range text {
		encryptedText += machine.through(string(v))
	}
	for _, rotor := range machine.rotors {
		rotor.reset()
	}
	return encryptedText
}

func (machine *Machine) through(char string) string {
	char = strings.ToUpper(char)
	indexNum, err := getAlphabetIndex(char)
	if err != nil {
		return char
	}
	indexNum = machine.plugBoard.Forward(indexNum)

	for _, rotor := range machine.rotors {
		indexNum = rotor.plugBoard.Forward(indexNum)
	}
	indexNum = machine.reflector.reflect(indexNum)

	length := len(machine.rotors)

	reverseRotors := make([]*Rotor, len(machine.rotors))
	for _, rotor := range machine.rotors {
		reverseRotors[length-1] = rotor
		length--
	}
	for _, rotor := range reverseRotors {
		indexNum = rotor.plugBoard.Backward(indexNum)
	}

	indexNum = machine.plugBoard.Backward(indexNum)

	char = ALPHABET[indexNum]

	for _, rotor := range reverseRotors {
		if rotor.offset == 0 {
			continue
		}
		if rotor.rotate()%len(ALPHABET) != 0 {
			break
		}
	}

	return char
}
