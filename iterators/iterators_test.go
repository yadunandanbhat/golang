package iterators

import "testing"

func TestIterators(t *testing.T) {
	t.Run("repeat 'a' for default defined loops", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("got %q expected %q", repeated, expected)
		}
	})
}

func BenchmarkIterators(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("x", 10)
	}
}