package main

import (
	"fmt"
	"testing"
)

func testPrint(t *testing.T) {
	// t.SkipNow()
	res := Print1to20()
	fmt.Println("hey")
	if res != 210 {
		t.Errorf("Wrong result of Print1to20")
	}
}

func testPrint2(t *testing.T) {
	res := Print1to20()
	res++
	if res != 211 {
		t.Errorf("Test Print2 failed")
	}
}

func TestAll(t *testing.T) {
	t.Run("TestPrint", testPrint)
	t.Run("TestPrint2", testPrint2)
}

func TestMain(m *testing.M) {
	fmt.Println("Tests begins...")
	m.Run()
}

func aaa(n int) int {
	for n > 0 {
		n--
	}
	return n
}

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		aaa(i)
	}
}
