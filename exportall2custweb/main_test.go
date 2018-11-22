package main

import (
	"fmt"
	"testing"
)

const (
	testinfile string = `C:\Users\Laurent\Golang\src\github.com\lpuig\scopeleccustomerweb\test\exportAll_test.csv`
)

func TestNewRecordSetFromFile(t *testing.T) {
	rs, err := NewRecordSetFromFile(testinfile)
	if err != nil {
		t.Fatal("NewRecordSetFromFile returns:", err)
	}

	for _, r := range rs.GetHeader().GetKeys() {
		fmt.Printf("'%s'\n", r)
	}
}
