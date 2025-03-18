package writetofile

import (
	"errors"
	"fmt"
	"os"
)

func WriteToFileFloat(value float64, filePath string) error {
	valueString := fmt.Sprint(value)
	err := os.WriteFile(filePath, []byte(valueString), 0644)
	if err != nil {
		return errors.New("file not found")
	}
	return nil
}

func WriteToFileString(value string, filePath string) error {
	err := os.WriteFile(filePath, []byte(value), 0644)
	if err != nil {
		return errors.New("file not found")
	}
	return nil
}
