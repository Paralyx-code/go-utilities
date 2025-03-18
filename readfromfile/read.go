package readfromfile

import (
	"errors"
	"os"
	"strconv"
)

func ReadFromFileFloat(filePath string) (float64, error) {
	data, err := os.ReadFile(filePath)

	if err != nil {
		return 1000, errors.New("file not found")
	}

	valueString := string(data)
	value, err := strconv.ParseFloat(valueString, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}

	return value, nil
}

func ReadFromFileString(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)

	if err != nil {
		return "error", errors.New("file not found")
	}

	valueString := string(data)
	//value, err := strconv.ParseFloat(valueString, 64)

	//if err != nil {
	//	return "error", errors.New("failed to parse stored value")
	//}

	return valueString, nil
}
