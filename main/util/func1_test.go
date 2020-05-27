package util

import (
	"math"
	"testing"
)

func TestFunc1(t *testing.T) {
	res := int(math.Abs(-10))
	if res != 10{
		t.Fatalf("Abs(-10) = %d, 期望值 = %d", res, 10)
	}
}

func BenchmarkFunc1(b *testing.B){
	for i:= 0; i < b.N; i++{
		math.Cos(0.5)

	}
}