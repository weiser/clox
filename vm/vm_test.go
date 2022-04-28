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
func TestExecuteOP_BINARY(t *testing.T) {
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
