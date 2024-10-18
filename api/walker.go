package api

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

var fileCount int = 1

type BasicWalker struct{}

func visitDirectory(path string, d fs.DirEntry, err error) error {
	i,_ := d.Info() //todo handle error
	fmt.Printf("visited: %v\n", i.Name())
	fileCount++
	
	//naive approach, have another function that updates an init var to count totals / keep track of dir, names etc

	return nil
}

func (w BasicWalker) Walk(path string) (int64, []int64, []string) {
	filepath.WalkDir(path, visitDirectory)
	fmt.Printf("fileCount %v\n", fileCount)
	return 0, nil, nil
}
