package processor

import "testing"

// compare int array (deep)
func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestProcessShortString(t *testing.T) {
	actual := Process("12345")

	if len(actual) != 0 {
		t.Fatalf("Expected empty array but got %v", actual)
	}
}

func TestProcessLongString(t *testing.T) {
	actual := Process("12345678912345")

	if len(actual) != 0 {
		t.Fatalf("Expected empty array but got %v", actual)
	}
}

func TestProcessAllSingleDigits(t *testing.T) {
	actual := Process("7654321")
	expected := []int{7, 6, 5, 4, 3, 2, 1}

	if !IntArrayEquals(actual, expected) {
		t.Fatalf("Expected %v array but got %v", expected, actual)
	}
}

func TestProcessAllDoubleDigits(t *testing.T) {
	actual := Process("49385328564754")
	expected := []int{49, 38, 53, 28, 56, 47, 54}

	if !IntArrayEquals(actual, expected) {
		t.Fatalf("Expected %v array but got %v", expected, actual)
	}
}

func TestProcessDupDigits(t *testing.T) {
	actual := Process("49385328565454")

	if len(actual) != 0 {
		t.Fatalf("Expected empty array but got %v", actual)
	}
}

func TestProcessAllDupDigits(t *testing.T) {
	actual := Process("111111111")

	if len(actual) != 0 {
		t.Fatalf("Expected empty array but got %v", actual)
	}
}
