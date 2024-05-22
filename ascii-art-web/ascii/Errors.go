package ascii

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"strings"
)

func AsciiCheck(s string) error {
	for _, ch := range s {
		if ch != 10 && (ch < 32 || ch > 126) {
			return fmt.Errorf("Ascii error")
		}
	}
	return nil
}

const (
	StandardHash   = "fe6d3468cf5c74d8ec2a95b40f2e05338c37a4202f8fad692d2b64a9cf9b468a"
	SnadowHash     = "0bb09d80600eec3eb9d7793a6f859bedde2a2d83899b70bd78e961ed674b32f4"
	ThinkertoyHash = "e514ee1f92ecf4cbb569ce32a36f679c5329050803af3a831ec7dab1a0e3d5f3"
)

func HashCheck(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	currentHash := calculateSHA256(data)
	if fileName == "../ascii-art-web/banners/standard.txt" {
		if currentHash == StandardHash {
			return nil
		}
	}
	if fileName == "../ascii-art-web/banners/shadow.txt" {
		if currentHash == SnadowHash {
			return nil
		}
	}
	if fileName == "../ascii-art-web/banners/thinkertoy.txt" {
		if currentHash == ThinkertoyHash {
			return nil
		}
	}
	return nil
}

func NewlinesCheck(text string) int {
	count := 0
	for i := 0; i < len(text)-1; i++ {
		if text[i] == 92 && text[i+1] == 'n' {
			count++
		}
	}
	return count
}

func News(inputText string) int {
	clon := strings.Replace(inputText, "\\n", "", -1)
	if len(clon) > 0 {
		return 0
	}
	return NewlinesCheck(inputText)
}

func calculateSHA256(fileContent []byte) string {
	hash := sha256.New()
	hash.Write(fileContent)
	hashSum := hash.Sum(nil)
	return fmt.Sprintf("%x", hashSum)
}
