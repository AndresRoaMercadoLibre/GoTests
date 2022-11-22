package tests

import (
	"TDD/model/mock"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := mock.Repeat("a")
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mock.Repeat("a")
	}
}
