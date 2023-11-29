package main
import ("fmt"

)
//basic functions with one arguments
func greetings(n string){
	fmt.Printf("Good Morning %v\n",n)
}
func goodby(n string){
	fmt.Printf("Goodby %v \n",n)
}
//complex functions with multipel arguments
func funcnamecy(n []string, f func(string)){
	for _,v:=range n{
		f(v)
	}
}
func main(){
	//greetings("Mac")
    // goodby("Mac")
	funcnamecy([]string{"Paul","Mac","John"}, greetings)
    funcnamecy([]string{"Paul","Mac","John"}, goodby)
}