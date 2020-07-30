package find_unique_number

import (
	"fmt"
	"testing"
	. "github.com/pranavraja/zen"
)

type TestCase struct {
	test []float32
	expected float32
}

func TestFindUnique(t *testing.T) {
	Desc(t, "Find Unique Simple Tests", func(it It) {
			var testCaseList = []TestCase {
				TestCase{[]float32{ 1.0, 1.0, 1.0, 2.0, 1.0, 1.0 }, float32(2)},
				TestCase{[]float32{ 0, 0, 0.55, 0, 0 }, float32(0.55)},
			}

			for _, testCase := range testCaseList {
				it(fmt.Sprintf("should find %g on test array", testCase.expected), func(expect Expect) {
					expect(FindUniq(testCase.test)).ToEqual(testCase.expected)
				})
			}
	})
}
