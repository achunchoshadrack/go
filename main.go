package main
import "fmt"
func main(){
	// bits and memories
	var numOne int8
    var name1 string
    var ages[3]int=[3]int{12,14,16}
	fmt.Println(ages,len(ages))
	fmt.Println("enter your age")
	fmt.Scanln(&numOne)
	fmt.Println("enter your name")
	fmt.Scanln(&name1)
	fmt.Println("my name is ", name1,"and my age is", numOne)
	//printf string formating
	fmt.Printf("my name is %v and my age is %v \n", name1, numOne)
	fmt.Printf("my name is %q and my age is %q \n", name1, numOne)
	fmt.Printf("my name is of type %T\n", name1)
	//sprintf formated string
	var str= fmt.Sprintf("your name is %v and my age is %v \n", name1, numOne)
	fmt.Println("the saved string is", str)

}