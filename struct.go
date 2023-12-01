package main
import ("fmt"

)

func main(){
	var n string
	fmt.Printf("Customer's Name:")
	fmt.Scanln(&n)
	mybill:=NewBill(n)
	fmt.Println(mybill.formate())
}