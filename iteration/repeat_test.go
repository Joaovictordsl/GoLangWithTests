package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat(5, "a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(5, "a")
	}
}

func ExampleRepeat() {
	repeat := Repeat(20, "b")
	fmt.Println(repeat)
	//Output: bbbbbbbbbbbbbbbbbbbb
}
