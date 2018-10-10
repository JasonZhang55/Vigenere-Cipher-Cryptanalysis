package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"math"
)

func main() {
	path_text := os.Args[1]
	cipherText := ReadFile(path_text)

	ICs := KeyLengthGuess(cipherText)
	fmt.Println("Most probable key length is:", ICs)
}

func ReadFile(path string) string {
	content, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Print(err)
    }
    return string(content)
}

func IndexOfCoincidence(msg string) float64 {
	nums := 0.0 
	sums := 0.0
	// Initialize numbers of each letter
	values := make([]float64, 26)
	for i := 0; i < 26; i++ {
		values[i] = 0
	}

	// Total numbers of each letter
	for _, val := range msg {
		values[val - 'A']++
		nums++
	}

	// Sums of the frequnce of each letter
	for i := 0; i < 26; i++ {
		sums += values[i] * (values[i] - 1)
	}

	res := sums / (nums * (nums - 1))

	return float64(res)
}

func KeyLengthGuess(msg string) int{
	sMsg := Sanitize(msg)
	msgLen := len(sMsg)
	standard := 0.0667

	for i := 1; i <= 20; i++ {
		count := 0
		IC := 0.0
		for start := 0; count < msgLen; start++ {
			current := start
			tmp := ""
			for {
				tmp += string(sMsg[current])
				current = (current + i)
				count++
				if current >= msgLen {
					break
				}
			}
			IC += IndexOfCoincidence(tmp)
		}
		diff := math.Abs(standard - IC / float64(i))
		if diff < 0.006 {
			return i
		}
	}
	return 0
}

func Sanitize(in string) string {
	out := []rune{}
	for _, val := range in {
		if 65 <= val && val <= 90 {
			out = append(out, val)		
		} else if 97 <= val && val <= 122 {
			out = append(out, val - 32)
		}
	}

	return string(out)
}