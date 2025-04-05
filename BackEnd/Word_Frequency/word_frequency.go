package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var word string
	m := make(map[string]uint)
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter a string\n")
	word, _ = input.ReadString('\n')
	if(len(word) != 1){
		for k:=0;k<len(word);k++{
			_, found := m[string(word[k])]
			if(found){
				m[string(word[k])]+=1
			}else{
				m[string(word[k])]=1
			}
		}
	}else{
		fmt.Print("Please enter a string\n")
		main()
	}
	fmt.Println(m)

}
