package sample

import "github.com/ngenohkevin/pcbook/pb"

// NewKeyBoard returns a new sample of the keyboard
func NewKeyBoard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBoolean(),
	}
}
