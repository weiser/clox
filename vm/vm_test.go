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

	ir := Interpret(&chnk)
	if ir != INTERPRET_OK {
		t.Errorf("expected INTERPRET_OK, got %v", ir)
	}

}
