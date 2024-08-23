package api

import (
	//"io/fs"
	"testing"
	//"testing/fstest"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// func TestVisitDirectory(t *testing.T) {
// 	// Create a mock filesystem with a single file
// 	mockFS := fstest.MapFS{
// 		"test.txt": {Data: []byte("Hello, World!")},
// 	}

// 	// Use the mock filesystem
// 	data, err := fs.ReadFile(mockFS, "test.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	println(string(data)) // Output: Hello, World!
// }
