package compute

import (
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
)

func TestFactorial(t *testing.T) {
	if factorial(2) != 2 {
		t.Fatalf("failed test")
	}
	if factorial(3) != 6 {
		t.Fatalf("failed test")
	}
	if factorial(4) != 24 {
		t.Fatalf("failed test")
	}
	if factorial(5) != 120 {
		t.Fatalf("failed test")
	}
	if factorial(10) != 3628800 {
		t.Fatalf("failed test")
	}
	if factorial(13) != 6227020800 {
		t.Fatalf("failed test")
	}
	if factorial(17) != 355687428096000 {
		t.Fatalf("failed test")
	}
	if factorial(20) != 2432902008176640000 {
		t.Fatalf("failed test")
	}
}

func TestCombination(t *testing.T) {
	var act int
	
	act = combination(2, 2)
	if act != 1 {
		t.Fatalf("failed test %v", act)
	}
	
	act = combination(3, 2)
	if act != 3 {
		t.Fatalf("failed test %v", act)
	}
	
	act = combination(4, 2)
	if act != 6 {
		t.Fatalf("failed test %v", act)
	}
}
