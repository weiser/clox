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

func InitVM(isDebug bool) {
	_isDebug = isDebug
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

func InterpretSource(source string, isDebug bool) InterpreterResult {
	var _chunk chunk.Chunk
	chunk.InitChunk(&_chunk)

	if !compiler.Compile(source, &_chunk, _isDebug) {
		return INTERPRET_COMPILE_ERROR
	}

	_vm.Chunk = &_chunk
	_vm.Ip = _vm.Chunk.Code
	_vm.ip_idx = 0
	_vm.StackTop = 0
	return run()

}

func Interpret(chnk *chunk.Chunk) InterpreterResult {
	_vm.Chunk = chnk
	_vm.Ip = _vm.Chunk.Code
	_vm.ip_idx = 0
	_vm.StackTop = 0
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

func Peek(distance int) value.Value {
	return _vm.Stack[_vm.StackTop-distance-1]
}

func runtimeError(msg string, a ...any) {
	fmt.Println(msg, a)
	instruction := _vm.Ip[_vm.ip_idx]
	_line := _vm.Chunk.Lines[instruction]
	fmt.Printf("[line %v] in script\n", _line)
	resetStack()
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
	if !value.IsNumber(Peek(0)) || !value.IsNumber(Peek(1)) {
		runtimeError("Operands must be numbers")

	} else {
		b := value.AsNumber(Pop())
		a := value.AsNumber(Pop())
		switch op {
		case ">":
			v = value.BoolVal(a > b)
		case "<":
			v = value.BoolVal(a < b)
		case "+":
			v = value.NumberVal(a + b)
		case "-":
			v = value.NumberVal(a - b)
		case "*":
			v = value.NumberVal(a * b)
		case "/":
			v = value.NumberVal(a / b)
		}
		Push(v)
	}
}

func run() InterpreterResult {
	// BEWARE: IP refers to the initial instruction and we might need to change
	// what that points to based on what instruction we are looking at,
	// e.g. an OP_CONSTANT might need to go to the next IP to get the index for
	// the constant to load
	for {
		// if you want to renove debugging, set `_isDebug = false` in InitVM()
		if _isDebug {
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
			if !value.IsNumber(Peek(0)) {
				runtimeError("Operand must be a number")
				return INTERPRET_RUNTIME_ERROR
			}
			Push(value.NumberVal(-value.AsNumber(Pop())))
		case chunk.OP_GREATER:
			Binary_OP(">")
		case chunk.OP_LESS:
			Binary_OP("<")
		case chunk.OP_ADD:
			Binary_OP("+")
		case chunk.OP_SUBTRACT:
			Binary_OP("-")
		case chunk.OP_MULTIPLY:
			Binary_OP("*")
		case chunk.OP_DIVIDE:
			Binary_OP("/")
		case chunk.OP_NIL:
			Push(value.BoolNil())
		case chunk.OP_FALSE:
			Push(value.BoolVal(false))
		case chunk.OP_EQUAL:
			b := Pop()
			a := Pop()
			Push(value.BoolVal(value.ValuesEqual(a, b)))
		case chunk.OP_TRUE:
			Push(value.BoolVal(true))
		default:
			return INTERPRET_COMPILE_ERROR
		}
	}
}

func ReadConstant() value.Value {
	return _vm.Chunk.Constants[ReadByte()]
}

// start on 15.3, pg 277
