package main
import ("fmt"
"math"

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
// functions that return a value
func circleArea(r float64) float64{
	return math.Pi * r * r
}

func main(){
	//greetings("Mac")
    // goodby("Mac")
	funcnamecy([]string{"Paul","Mac","John"}, greetings)
    funcnamecy([]string{"Paul","Mac","John"}, goodby)
	//calling returning value functions
	a1:=circleArea(10.5)
	a2:=circleArea(5.5)
	fmt.Printf("Area of circle 1 is: %.3f\n Area of circle 2 is: %.3f",a1,a2)
}