package vm

import (
	"fmt"

	"github.com/weiser/clox/chunk"
	"github.com/weiser/clox/compiler"
	"github.com/weiser/clox/debug"
	"github.com/weiser/clox/value"
)

type VM struct {
	Chunk    *chunk.Chunk
	Ip       []uint8
	ip_idx   int
	Stack    [1024]value.Value
	StackTop int
}

type InterpreterResult int

const (
	INTERPRET_OK InterpreterResult = iota
	INTERPRET_COMPILE_ERROR
	INTERPRET_RUNTIME_ERROR
)

var _vm *VM
var _isDebug bool

func InitVM() {
	_isDebug = true
	//if _vm == nil {
	_vm = &VM{}
	resetStack()
	//}
}

func resetStack() {
	_vm.StackTop = 0
}

func FreeVM() {

}

func InterpretSource(source string) InterpreterResult {
	compiler.Compile(source)
	return INTERPRET_OK
}

func Interpret(chnk *chunk.Chunk) InterpreterResult {
	_vm.Chunk = chnk
	_vm.Ip = _vm.Chunk.Code
	_vm.ip_idx = 0
	return run()
}

func Push(v value.Value) {
	_vm.Stack[_vm.StackTop] = v
	_vm.StackTop += 1
}

func Pop() value.Value {
	_vm.StackTop -= 1
	return _vm.Stack[_vm.StackTop]
}

// ReadByte reads the next byte of the instruction pointer,
// incrementing where the next byte is read from for next time this is called
func ReadByte() uint8 {
	v := _vm.Ip[_vm.ip_idx]
	_vm.ip_idx = _vm.ip_idx + 1
	return v
}

func Binary_OP(op string) {
	var v value.Value
	b := Pop()
	a := Pop()
	switch op {
	case "+":
		v = a + b
	case "-":
		v = a - b
	case "*":
		v = a * b
	case "/":
		v = a / b
	}
	Push(v)
}

func run() InterpreterResult {
	// BEWARE: IP refers to the initial instruction and we might need to change
	// what that points to based on what instruction we are looking at,
	// e.g. an OP_CONSTANT might need to go to the next IP to get the index for
	// the constant to load
	for {
		// if you want to renove debugging, set `_isDebug = false` in InitVM()
		if _isDebug == true {
			fmt.Print("        ")
			for i := 0; i < _vm.StackTop; i += 1 {
				fmt.Print("[ ")
				value.PrintValue(_vm.Stack[i])
				fmt.Println(" ]")
			}
			debug.DisassembleInstruction(_vm.Chunk, int(_vm.ip_idx))
		}

		instruction := ReadByte()
		switch instruction {
		case chunk.OP_CONSTANT:
			constant := ReadConstant()
			Push(constant)
		case chunk.OP_RETURN:
			//value.PrintValue(Pop())
			//fmt.Println()
			return INTERPRET_OK
		case chunk.OP_NEGATE:
			Push(-Pop())
		case chunk.OP_ADD:
			Binary_OP("+")
		case chunk.OP_SUBTRACT:
			Binary_OP("-")
		case chunk.OP_MULTIPLY:
			Binary_OP("*")
		case chunk.OP_DIVIDE:
			Binary_OP("/")
		default:
			return INTERPRET_COMPILE_ERROR
		}
	}
}

func ReadConstant() value.Value {
	return _vm.Chunk.Constants[ReadByte()]
}
