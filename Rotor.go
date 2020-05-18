package enigmaMachine

type Rotor struct {
	offset    int
	rotations int
	plugBoard *PlugBoard
}

func NewRotor(mapAlphabet []string, offset int) *Rotor {
	plugBoard := NewPlugBoard(mapAlphabet)
	rotor := &Rotor{
		offset:    offset,
		rotations: 0,
		plugBoard: plugBoard,
	}
	return rotor
}

func (Rotor *Rotor) rotate() int {
	var alphabet []string
	for _, v := range Rotor.plugBoard.alphabet[Rotor.offset:] {
		alphabet = append(alphabet, v)
	}
	for _, v := range Rotor.plugBoard.alphabet[:Rotor.offset] {
		alphabet = append(alphabet, v)
	}
	Rotor.plugBoard.alphabet = alphabet

	Rotor.rotations += Rotor.offset
	return Rotor.rotations
}

func (Rotor *Rotor) reset() {
	Rotor.rotations = 0
	Rotor.plugBoard.alphabet = ALPHABET
}
