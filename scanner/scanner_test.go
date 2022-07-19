package scanner

import "testing"

func Test_ScanToken(t *testing.T) {
	InitScanner("fun functy")
	_t := ScanToken()
	if _t.Type != TOKEN_FUN {
		t.Errorf("expected TOKEN_FUN, got %v", _t)
	}

	_t = ScanToken()
	if _t.Type != TOKEN_IDENTIFIER {
		t.Errorf("expected TOKEN_IDENTIFIER, got %v", _t)
	}
}

func Test_ScanTokenNumber(t *testing.T) {
	InitScanner("1")
	_t := ScanToken()
	if _t.Type != TOKEN_NUMBER {
		t.Errorf("expected TOKEN_NUMBER, got %v", _t)
	}
}
