package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	key := os.Args[1]
	path_Text := os.Args[2]
	// key := ReadFile(path_key)
	cipherText := ReadFile(path_Text)

	if len(key) > 32 {
		fmt.Println("Key size should be less than 32 characters.")
		os.Exit(3)
	}

	plainText := Decrypt(cipherText, key)
	fmt.Println(plainText)


}

func ReadFile(path string) string {
	content, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Print(err)
    }
    return string(content)
}

// func WriteFile(content, path string) string {
// 	err := ioutil.WriteFile(path, content, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// }

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

func Decrypt(msg, key string) string {
	sMsg, sKey := Sanitize(msg), Sanitize(key)
	out := make([]rune, 0, len(msg))

	for index, val := range sMsg {
		out = append(out, DecodePair(val, rune(sKey[index%len(sKey)])))
	}
	return string(out)
}

func DecodePair(m, k rune) rune {
	return ((m - 'A') - (k - 'A') + 26) % 26 + 'A'
}