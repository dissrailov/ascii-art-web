package ascii

import (
	"fmt"
	"path/filepath"
)

func AsciiMain(input, filename string) (string, error) {
	filePath := filepath.Join("../ascii-art-web/banners", filename+".txt")
	if len(input) == 0 {
		return "", fmt.Errorf("Usage: go run . <text>")
	}
	inputText := input
	if inputText == "" {
		return "", fmt.Errorf("Input text is empty")
	}
	if err := HashCheck(filePath); err != nil {
		return "", fmt.Errorf(err.Error())
	}
	if err := AsciiCheck(inputText); err != nil {
		return "", fmt.Errorf(err.Error())
	}
	count := News(inputText)
	if count != 0 {
		for i := 0; i < count; i++ {
			fmt.Println("")
		}
		return "", nil
	}
	res := Ascii_Printer(filename, inputText)
	return res, nil
}
