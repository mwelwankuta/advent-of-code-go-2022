package question

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Question struct {
	File    string      // Name of the file with input data
	PartOne interface{} // Part one solution
	PartTwo interface{} // Part two solution
}

// LoadInput(fileName string)
// stores in memory the data from a given file
// e.g LoadInput("input.txt")
// the function searches the current directory for a file and panics when non is found
func (q *Question) LoadInput(dayPath, fileName string) {
	filePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	file, err := ioutil.ReadFile(fmt.Sprintf("%v/../internal/%v/%v", filePath, dayPath, fileName))
	if err != nil {
		panic(err)
	}

	q.File = string(file)
}
