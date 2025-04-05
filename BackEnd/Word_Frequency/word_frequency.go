package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var word string
	m := make(map[string]uint)
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter a string: \n")
	word, _ = input.ReadString('\n')
	if(len(word) != 1){
		for k:=0;k<len(word);k++{
			upper:=strings.ToUpper(string(word[k]))
			_, found := m[upper]
			if(found){
				m[upper]+=1
			}else if word[k] >= 65 && word[k]<= 190{
				m[upper]=1
			}
		}
	}else{
		main()
	}
	fmt.Println(m)

}
