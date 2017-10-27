package brute

import "testing"

func TestNextByte(t *testing.T) {
	t.Run("should return 0 for z", func(t *testing.T) {
		actual := NextByte('z')
		expected := byte('0')
		if actual != expected {
			t.Fatalf("Expected %v but got %v", expected, actual)
		}
	})

	t.Run("should return a for 9", func(t *testing.T) {
		actual := NextByte('9')
		expected := byte('a')
		if actual != expected {
			t.Fatalf("Expected %v but got %v", expected, actual)
		}
	})

	t.Run("should return b for a", func(t *testing.T) {
		actual := NextByte('a')
		expected := byte('b')
		if actual != expected {
			t.Fatalf("Expected %v but got %v", expected, actual)
		}
	})
}
