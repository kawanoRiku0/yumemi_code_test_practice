package inits

import (
	"encoding/csv"
	"fmt"
	"os"
)

func InitCSVFileWriter(fName string) (*csv.Writer, error) {
	f, err := os.Create(fName)
	if err != nil {
		return nil, fmt.Errorf("creating file faile : %s \n", err.Error())
	}

	w := csv.NewWriter(f)
	return w, nil
}
