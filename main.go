package main
import "fmt"
func main(){
	// bits and memories
	var numOne int8
    var name1 string
	//arrays
    var ages[3]int=[3]int{12,14,16}
	var ages1=[3]int{22,14,26}
	unames:=[3]string{"mac","paul","max" }
	//slices (use arrays hood)
	var Names=[]int{10,20,30}
	fmt.Println(Names,len(Names))
	// append a slice
	Names=append(Names, 89)
	fmt.Println(Names,len(Names))
	//input output
	fmt.Println(ages,len(ages))
	fmt.Println(ages1,len(ages1))
	fmt.Println(unames,len(unames))
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