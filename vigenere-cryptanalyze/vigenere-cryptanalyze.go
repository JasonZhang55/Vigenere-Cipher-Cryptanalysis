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

	key := GetKey(cipherText, ICs)
	fmt.Println("Key:", key)
}

func ReadFile(path string) string {
	content, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Print(err)
    }
    return string(content)
}


func GetKey(msg string, length int) []string{
	key := []string{}
	alphaRate := []float64{0.08167,0.01492,0.02782,0.04253,0.12705,0.02228,0.02015,0.06094,0.06996,0.00153,0.00772,0.04025,0.02406,0.06749,0.07507,0.01929,0.0009,0.05987,0.06327,0.09056,0.02758,0.00978,0.02360,0.0015,0.01974,0.00074}
    matrix := TextToMatrix(msg,length)

    for i := 0; i < length; i++ {
    	cipherSlice := []string{}
    	for _, row := range matrix {
    		cipherSlice = append(cipherSlice, string(row[i]))
    	}
    	chsFrequency := CountFrequency(cipherSlice)
    	// fmt.Println(cipherSlice)
    	powChs := []float64{}
    	for j := 0; j < 26; j++ {
    		sums := 0.0
    		for k := 0; k < 26; k++ {
    			sums += alphaRate[k] * chsFrequency[k]
    		}
    		powChs = append(powChs, sums)
    		chsFrequency = append(chsFrequency[1:] , chsFrequency[:1]...)
    	}
    	Abs := 100.0
    	ch := ""
    	for j := 0; j < len(powChs); j++ {
    		if math.Abs(powChs[j] - 0.065546) < Abs {
    			Abs = math.Abs(powChs[j] - 0.065546)
    			ch = string(j + 65)
    			// fmt.Println(ch)
    		}
    	}
    	key = append(key, ch)
    }

    return key

}

func TextToMatrix(in string, length int) [][]string{
	textMatrix := [][]string{}
	row := []string{}
	index := 0
	for _, ch := range in {
		row = append(row, string(ch))
		index += 1
		if index % length == 0 {
			textMatrix = append(textMatrix, row)
			row = []string{}
		}
	}

	return textMatrix
}

func CountFrequency(list []string) []float64 {
	res := []float64{}
	alphabet := []string{}
	for i := 65; i < 91; i++ {
		alphabet = append(alphabet, string(rune(i)))
	}
	for _, c := range alphabet {
		count := 0.0
		for _, ch := range list {
			if c == ch {
				count ++
			}
		} 
		res = append(res, count/float64(len(list)))
	}
	return res
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