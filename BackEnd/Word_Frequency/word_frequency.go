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
	fmt.Print("Please enter a string: \n")
	word, _ = input.ReadString('\n')
	if(len(word) != 1){
		for k:=0;k<len(word);k++{
			_, found := m[string(word[k])]
			if(found){
				m[string(word[k])]+=1
			}else if word[k] >= 65 && word[k]<= 122{
				if word[k] > 96 || word[k]< 91 {
					m[string(word[k])]=1
				}
			}
		}
	}else{
		main()
	}
	fmt.Println(m)

}
