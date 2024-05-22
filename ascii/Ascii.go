package ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Ascii_Printer(fileName, text string) string {
	fileName = "../ascii-art-web/banners/" + fileName + ".txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var contentLines []string

	// Read all lines from the file and store them in contentLines
	for scanner.Scan() {
		contentLines = append(contentLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return ""
	}

	final := ""
	var result string
	var charCodes []int
	inputText := strings.Replace(text, "\\n", "\r\n", -1)
	lines := strings.Split(inputText, "\n")

	for _, line := range lines {
		charCodes = GetCharCodes(line)
		for i := 0; i < 8; i++ {
			if len(charCodes) == 0 {
				break
			}
			for j := 0; j < len(charCodes); j++ {
				if charCodes[j] == 10 {
					charCodes = charCodes[1:]
					continue
				}
				count := 0 + 9*(charCodes[j]-32)
				if count+i >= 0 && count+i < len(contentLines) {
					result += contentLines[count+i]
				}
			}
			if i < 7 {
				result += "\n"
			}
		}
	}

	lines2 := strings.Split(result, "\n")
	emptyLineCount := 0
	for _, line := range lines2 {
		if line == "" {
			emptyLineCount++
		}
	}

	if emptyLineCount == len(lines)-1 {
		final += result
	} else {
		final += result + "\n"
	}

	return final
}

func GetCharCodes(text string) []int {
	var charCodes []int
	for _, char := range text {
		charCodes = append(charCodes, int(char))
	}
	return charCodes
}
