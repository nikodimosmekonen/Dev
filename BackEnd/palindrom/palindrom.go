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
	word, _ = input.ReadString('\n')
	word=strings.ToUpper(string(word))
	count := len(word)-1
	for i:=0;i<=len(word)/2;i++{
		if word[i] >= 65 && word[i]<= 190{
			if(word[i]==word[count] ){
				count-=1
			}else if word[i]!=word[count]{
				fmt.Println("not a palindrome!")
				return
			}
		}
	}
}
