package chunk

import "testing"

func TestWriteChunk(t *testing.T) {
	var chnk Chunk
	InitChunk(&chnk)
	WriteChunk(&chnk, OP_RETURN)
	if chnk.Code[0] != OP_RETURN {
		t.Errorf("expected OP_RETURN, got %v", chnk.Code[0])
	}
}
