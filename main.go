package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func getLetter(arr []string, char rune, line int) string { 
	asciiValue := (int(char) - 32) *9 + line // change char to its ascii value (-32) & *9 since we have 8 char lines & 1 extra line
			for count, _ := range arr { // go through all the letters 
				if count == asciiValue {
					count ++
					return arr[count]
					}
				}
	return ""
}

func main() {
	input := os.Args[1]
	banner := "standard" // You can change this to "shadow" or "thinkertoy" as needed

	if len(os.Args) ==  2 {
		bannerData, err := ioutil.ReadFile(fmt.Sprintf("banners/%s.txt", banner))
		if err != nil {
			fmt.Println("Error reading banner file:", err)
			os.Exit(1)
		}

		bannerLines := strings.Split(string(bannerData), "\n") // split the lines of the banner  
		cInput := strings.ReplaceAll(input, "\\n", "\n") // replace \n in input to new line 
		inputList := strings.Split(cInput, "\n") 
		for _, element := range inputList { 
			if element != "" {
				for i := 0; i < 8; i++ { // go through all 8 lines of each char
					for _, letter := range element {  
						fmt.Print(getLetter(bannerLines,letter, i))
					}
					fmt.Println("")
				}
			} else {
				fmt.Println("")
			}
			
		}
		

	} else {
		fmt.Println("Please provide input text.")
		os.Exit(1)
	}

	
}