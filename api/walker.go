package api

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

// todo change this
var fileCount int = 1

type InProgressOutput func(d fs.DirEntry) error

type BasicWalker struct {
	ipo InProgressOutput
}

// TODO: have visit directory take a param to say what should be outputtted as we are walking
// TODO: have function add data to a space and speed effecient in mem data store if possible, else write to a text file
func (w BasicWalker) VisitDirectory(path string, d fs.DirEntry, err error) error {
	i, _ := d.Info() //todo handle error
	fmt.Printf("visited: %v\n", i.Name())
	fileCount++

	w.ipo(d)

	//naive approach, have another function that updates an init var to count totals / keep track of dir, names etc

	return nil
}

func (w BasicWalker) Walk(path string) (int64, []int64, []string) {
	filepath.WalkDir(path, visitDirectory)
	fmt.Printf("fileCount %v\n", fileCount)
	return 0, nil, nil
}
