package vm

import (
	"testing"

	"github.com/weiser/clox/chunk"
)

func TestExecuteInstructions(t *testing.T) {
	InitVM()
	var chnk chunk.Chunk
	chunk.InitChunk(&chnk)
	idx := uint8(chunk.AddConstant(&chnk, 1234.0))
	chunk.WriteChunk(&chnk, chunk.OP_CONSTANT, 0)
	chunk.WriteChunk(&chnk, idx, 0)

	// this just causes `run()` to exit
	chunk.WriteChunk(&chnk, chunk.OP_RETURN, 0)

	ir := Interpret(&chnk)
	if ir != INTERPRET_OK {
		t.Errorf("expected INTERPRET_OK, got %v", ir)
	}

}

func TestExecuteOP_NEGATE(t *testing.T) {
	InitVM()
	var chnk chunk.Chunk
	chunk.InitChunk(&chnk)
	idx := uint8(chunk.AddConstant(&chnk, 1234.0))
	chunk.WriteChunk(&chnk, chunk.OP_CONSTANT, 0)
	chunk.WriteChunk(&chnk, idx, 0)
	chunk.WriteChunk(&chnk, chunk.OP_NEGATE, 0)

	// this just causes `run()` to exit
	chunk.WriteChunk(&chnk, chunk.OP_RETURN, 0)

	ir := Interpret(&chnk)
	v := Pop()
	if ir != INTERPRET_OK {
		t.Errorf("expected INTERPRET_OK, got %v", ir)
	}
	if v != -1234.0 {
		t.Errorf("expected -1234.0, got %v", v)
	}

}

// TODO: DEBUG WHY THIS TEST FAILS
func TestExecuteOP_BINARY_ADD(t *testing.T) {
	InitVM()
	var chnk chunk.Chunk
	chunk.InitChunk(&chnk)
	idx := uint8(chunk.AddConstant(&chnk, 2))
	chunk.WriteChunk(&chnk, chunk.OP_CONSTANT, 0)
	chunk.WriteChunk(&chnk, idx, 0)
	idx = uint8(chunk.AddConstant(&chnk, 1))
	chunk.WriteChunk(&chnk, chunk.OP_CONSTANT, 0)
	chunk.WriteChunk(&chnk, idx, 0)

	chunk.WriteChunk(&chnk, chunk.OP_ADD, 0)

	// this just causes `run()` to exit
	chunk.WriteChunk(&chnk, chunk.OP_RETURN, 0)

	ir := Interpret(&chnk)
	v := Pop()
	if ir != INTERPRET_OK {
		t.Errorf("expected INTERPRET_OK, got %v", ir)
	}
	if v != 3 {
		t.Errorf("expected 3.0, got %v", v)
	}

}

func TestExecuteOP_BINARY_SUB(t *testing.T) {
	InitVM()
	var chnk chunk.Chunk
	chunk.InitChunk(&chnk)
	idx := uint8(chunk.AddConstant(&chnk, 2))
	chunk.WriteChunk(&chnk, chunk.OP_CONSTANT, 0)
	chunk.WriteChunk(&chnk, idx, 0)
	idx = uint8(chunk.AddConstant(&chnk, 1))
	chunk.WriteChunk(&chnk, chunk.OP_CONSTANT, 0)
	chunk.WriteChunk(&chnk, idx, 0)

	chunk.WriteChunk(&chnk, chunk.OP_SUBTRACT, 0)

	// this just causes `run()` to exit
	chunk.WriteChunk(&chnk, chunk.OP_RETURN, 0)

	ir := Interpret(&chnk)
	v := Pop()
	if ir != INTERPRET_OK {
		t.Errorf("expected INTERPRET_OK, got %v", ir)
	}
	if v != 1 {
		t.Errorf("expected 1.0, got %v", v)
	}
}

func TestExecuteOP_BINARY_MULTIPLY(t *testing.T) {
	InitVM()
	var chnk chunk.Chunk
	chunk.InitChunk(&chnk)
	idx := uint8(chunk.AddConstant(&chnk, 2))
	chunk.WriteChunk(&chnk, chunk.OP_CONSTANT, 0)
	chunk.WriteChunk(&chnk, idx, 0)
	idx = uint8(chunk.AddConstant(&chnk, 1))
	chunk.WriteChunk(&chnk, chunk.OP_CONSTANT, 0)
	chunk.WriteChunk(&chnk, idx, 0)

	chunk.WriteChunk(&chnk, chunk.OP_MULTIPLY, 0)

	// this just causes `run()` to exit
	chunk.WriteChunk(&chnk, chunk.OP_RETURN, 0)

	ir := Interpret(&chnk)
	v := Pop()
	if ir != INTERPRET_OK {
		t.Errorf("expected INTERPRET_OK, got %v", ir)
	}
	if v != 2 {
		t.Errorf("expected 2.0, got %v", v)
	}

}

func TestExecuteOP_BINARY_DIVIDE(t *testing.T) {
	InitVM()
	var chnk chunk.Chunk
	chunk.InitChunk(&chnk)
	idx := uint8(chunk.AddConstant(&chnk, 2))
	chunk.WriteChunk(&chnk, chunk.OP_CONSTANT, 0)
	chunk.WriteChunk(&chnk, idx, 0)
	idx = uint8(chunk.AddConstant(&chnk, 1))
	chunk.WriteChunk(&chnk, chunk.OP_CONSTANT, 0)
	chunk.WriteChunk(&chnk, idx, 0)

	chunk.WriteChunk(&chnk, chunk.OP_DIVIDE, 0)

	// this just causes `run()` to exit
	chunk.WriteChunk(&chnk, chunk.OP_RETURN, 0)

	ir := Interpret(&chnk)
	v := Pop()
	if ir != INTERPRET_OK {
		t.Errorf("expected INTERPRET_OK, got %v", ir)
	}
	if v != 2 {
		t.Errorf("expected 2.0, got %v", v)
	}

}

// Test that (2+1)*(3+4) = 21
func TestExecuteOP_BINARY_ADD_MULTIPLY(t *testing.T) {
	InitVM()
	var chnk chunk.Chunk
	chunk.InitChunk(&chnk)
	idx := uint8(chunk.AddConstant(&chnk, 2))
	chunk.WriteChunk(&chnk, chunk.OP_CONSTANT, 0)
	chunk.WriteChunk(&chnk, idx, 0)
	idx = uint8(chunk.AddConstant(&chnk, 1))
	chunk.WriteChunk(&chnk, chunk.OP_CONSTANT, 0)
	chunk.WriteChunk(&chnk, idx, 0)

	chunk.WriteChunk(&chnk, chunk.OP_ADD, 0)

	idx = uint8(chunk.AddConstant(&chnk, 3))
	chunk.WriteChunk(&chnk, chunk.OP_CONSTANT, 0)
	chunk.WriteChunk(&chnk, idx, 0)
	idx = uint8(chunk.AddConstant(&chnk, 4))
	chunk.WriteChunk(&chnk, chunk.OP_CONSTANT, 0)
	chunk.WriteChunk(&chnk, idx, 0)

	chunk.WriteChunk(&chnk, chunk.OP_ADD, 0)

	chunk.WriteChunk(&chnk, chunk.OP_MULTIPLY, 0)

	// this just causes `run()` to exit
	chunk.WriteChunk(&chnk, chunk.OP_RETURN, 0)

	ir := Interpret(&chnk)
	v := Pop()
	if ir != INTERPRET_OK {
		t.Errorf("expected INTERPRET_OK, got %v", ir)
	}
	if v != 21 {
		t.Errorf("expected 21.0, got %v", v)
	}

}
