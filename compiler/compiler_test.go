package compiler

import (
	"testing"

	"github.com/weiser/clox/chunk"
)

// todo 17.7 on pg 323
// TODO build tests to verify compiler works.  E.g. should be abke to debug chunks on 17.7

func TestCompile(t *testing.T) {

	var _chunk chunk.Chunk
	chunk.InitChunk(&_chunk)
	res := Compile("1 + 1", &_chunk, false)
	if !res {
		t.Errorf("Expected no errors, but got errors")
	}

}

func TestCompile2(t *testing.T) {

	var _chunk chunk.Chunk
	chunk.InitChunk(&_chunk)
	res := Compile("1 / 1", &_chunk, false)
	if !res {
		t.Errorf("Expected no errors, but got errors")
	}

}
