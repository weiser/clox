package scanner

import "testing"

func Test_ScanToken(t *testing.T) {
	InitScanner("fun")
	_t := ScanToken()
	if _t.Type != TOKEN_FUN {
		t.Errorf("expected TOKEN_FUN, got %v", t)
	}

}
