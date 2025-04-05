package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var word string
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter a string: \n")
	alphanumeric:="1234567890qwertyuiopasdfghjklzxcvbnm"
	word, _ = input.ReadString('\n')
	word = strings.Trim(word,"\n")
	word=strings.ToLower(string(word))
	count := len(word)-1
	for i:=0;i<=len(word)/2;i++{
		if strings.Contains(alphanumeric, string(word[i])) && strings.Contains(alphanumeric, string(word[count])) {
			if word[i] == word[count]{
				// fmt.Printf("equal %v\n",string(word[count]))
				count-=1
			}else{
				fmt.Println("not a palindrome!")
				return
			}

		}else if strings.Contains(alphanumeric, string(word[i])) {
			// fmt.Printf("i %v\t %v\n",string(word[count]),count)
			count-=1
			i-=1
		}
		// fmt.Printf("the i %v\n",i)
		if count==i{
			break
		}
	}
	fmt.Println("palindrome!")
	return
}
