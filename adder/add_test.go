package task2

import (
	"fmt"
	"testing"
)

func Test_add(t *testing.T) {
	got := Add(2, 2)
	want := 4
	if got != want {
		t.Errorf("func add not work with got %d , want %d", got, want)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func Test_Repeat(t *testing.T) {
	got := Repeat("y", 10)
	want := "yyyyyyyyyy"
	if got != want {
		t.Errorf("func not work, got = %q, want = %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	ret := Repeat("a", 10)
	fmt.Println(ret)
	// Output: aaaaaaaaaaa
}
