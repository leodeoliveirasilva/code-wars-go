package find_unique_number

import "testing"

func TestFindUnique(t *testing.T) {
		response := FindUniq([]float32{ 1.0, 1.0, 1.0, 2.0, 1.0, 1.0 })
		expected := float32(2)
		if response != expected {
			t.Error()
		} else {
			t.Logf("FIND UNIQUE TEST PASSED")
		}
}