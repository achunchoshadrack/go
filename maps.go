package main
import ("fmt"

)
func main(){
	// maps takes keys and assign values
	// in a map the values must be of the same type and the keys of the same type
	menu:=map [string]float64{
		"Rice": 500,
		"Chicken": 1000,
		"katikati": 1500,
	}
	fmt.Println(menu)
	fmt.Println(menu["Rice"])
	//looping in a map
	for k,v:=range menu{
		fmt.Println(k," ",v)
	}
	//maps with ints key type
	phonebook:=map [int]string{
     683519095: "DESH-CS",
	 686638711: "MAHKO",
	 653943272: "SHADRACK",
	 687654390: "PAS-CS",
	}
	fmt.Println("Phone Book")
	for k, v:=range phonebook{
		fmt.Println(k,":  ", v)
	}
	//updating a map
	phonebook[687654390]="ACHUNCHO"
	fmt.Println(phonebook)
}