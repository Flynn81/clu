package api

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func visitDirectory(path string, d fs.DirEntry, err error) error {
	i,_ := d.Info() //todo handle error
	fmt.Printf("visited: %v\n", i)
	
	return nil
}

type BasicWalker struct{}

func (w BasicWalker) Walk(path string) (int64, []int64, []string) {
	filepath.WalkDir(path, visitDirectory)
	return 0, nil, nil
}
