package main

// Testing for even or odd num for calculation
func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "even" {
		t.Error("expected: even, actual: %s, result")
	}
}