package main
import ("fmt"

)

func main(){
	//var n string
	//fmt.Printf("Customer's Name:")
	//fmt.Scanln(&n)
	mybill:=NewBill("Shadrack's Bill")
	mybill.UpdateTip(500)
	mybill.Additem("Fried Rice", 1500)
	fmt.Println(mybill.formate())
}