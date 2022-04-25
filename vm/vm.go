package vm

import (
	"fmt"

	"github.com/weiser/clox/chunk"
	"github.com/weiser/clox/value"
)

type VM struct {
	Chunk  *chunk.Chunk
	Ip     []uint8
	ip_idx int
}
type InterpreterResult int

const (
	INTERPRET_OK InterpreterResult = iota
	INTERPRET_COMPILE_ERROR
	INTERPRET_RUNTIME_ERROR
)

var _vm *VM

func InitVM() {
	if _vm == nil {
		_vm = &VM{}
	}
}

func FreeVM() {

}

func Interpret(chnk *chunk.Chunk) InterpreterResult {
	_vm.Chunk = chnk
	_vm.Ip = _vm.Chunk.Code
	_vm.ip_idx = 0
	return run()
}

// ReadByte reads the next byte of the instruction pointer,
// incrementing where the next byte is read from for next time this is called
func ReadByte() uint8 {
	_vm.ip_idx = _vm.ip_idx + 1
	return _vm.Ip[_vm.ip_idx-1]
}

func run() InterpreterResult {
	// BEWARE: IP refers to the initial instruction and we might need to change
	// what that points to based on what instruction we are looking at,
	// e.g. an OP_CONSTANT might need to go to the next IP to get the index for
	// the constant to load
	for {
		instruction := ReadByte()
		switch instruction {
		case chunk.OP_CONSTANT:
			constant := ReadConstant()
			value.PrintValue(constant)
			fmt.Println()
			// this will change, but we put it here only to have something to play with for now
			return INTERPRET_OK
		case chunk.OP_RETURN:
			return INTERPRET_OK
		default:
			return INTERPRET_COMPILE_ERROR
		}
	}
}

func ReadConstant() value.Value {
	return _vm.Chunk.Constants[ReadByte()]
}

// todoon pg 269
