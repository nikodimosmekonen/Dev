package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type student struct{
	name string;
	subject map[string]float64;
}
func (s student) cal_grade()  float64{
	total := 0.0;
	for _,v := range s.subject {
		total+= v;
	}
	var average float64= total / float64(len(s.subject));

	return average;
}

func main(){

	// var a student;
	// a.name = "chala"
	// a.subject = make(map[string]float64)
	// a.subject["eng"]= 43.5
	// a.subject["amharic"]=42.5

	// fmt.Printf("%v got an average of %v",a.name,a.cal_grade())
	var stud student;
	fmt.Print("Enter name: ");

	input:= bufio.NewReader(os.Stdin)
	stud.name,_ = input.ReadString('\n')

	stud.name= strings.TrimSpace(stud.name);
	stud.subject = make(map[string]float64)


	
	var num int;
	fmt.Print("Enter the number of subjects: ");
	fmt.Scan(&num);

	var sub_name string ="";
	var sub_grade int= 0;

	for i:= 1; i< num+1;i++{
		
		fmt.Printf("Enter the name of subject %v and the Score recived \nName\tScore: ", i);
		fmt.Scan(&sub_name,&sub_grade)
		for sub_grade > 100 || sub_grade < 0{
			fmt.Printf("Score of %v must be between 0-100 \nEnter Score:\t",sub_name);
			fmt.Scan(&sub_grade);
		}
		stud.subject[sub_name]=float64(sub_grade)
	} 
	
	fmt.Printf("%v \nSubject \t Score\n",stud.name)

	for k,v := range stud.subject{
		fmt.Printf("%v \t %v\n",k,v)
	}
	fmt.Printf("Average of: \t %v",stud.cal_grade())
}
