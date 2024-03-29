package chunk

import (
	"testing"

	"github.com/weiser/clox/value"
)

func TestWriteChunk(t *testing.T) {
	var chnk Chunk
	InitChunk(&chnk)
	WriteChunk(&chnk, OP_RETURN, 0)
	if chnk.Code[0] != OP_RETURN {
		t.Errorf("expected OP_RETURN, got %v", chnk.Code[0])
	}
	if chnk.Lines[0] != 0 {
		t.Errorf("expected 0, got %v", chnk.Lines[0])
	}
}

func TestAddConstant(t *testing.T) {
	var chnk Chunk
	InitChunk(&chnk)
	idx := uint8(AddConstant(&chnk, value.NumberVal(1234.0)))
	WriteChunk(&chnk, OP_CONSTANT, 0)
	WriteChunk(&chnk, idx, 0)
	if value.AsNumber(chnk.Constants[idx]) != 1234.0 {
		t.Errorf("expected 1234.0, got %v", chnk.Constants[idx])
	}
	if chnk.Code[0] != OP_CONSTANT {
		t.Errorf("expected OP_CONSTANT, got %v", chnk.Code[0])
	}
	if chnk.Code[1] != 0 {
		t.Errorf("expected 1, got %v", chnk.Code[1])
	}
}
