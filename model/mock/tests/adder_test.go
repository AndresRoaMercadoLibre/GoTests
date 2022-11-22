package tests

import (
	"TDD/model/mock"
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := mock.Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but go '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := mock.Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
