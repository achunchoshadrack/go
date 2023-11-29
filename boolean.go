package main
import ("fmt"

)
func main(){
	age:=45
	fmt.Println(age<50)
	fmt.Println(age>=50)
	fmt.Println(age==45)
	fmt.Println(age!=45)
    // if statement with the breake and continue Keyword
	names:=[]string{"John","Paul", "Ambe","Zeno","Bless","Che"}
	for index, value:=range names{
		if index==1{
			fmt.Println("continuing at position",index)
			continue//at this point the program breaks the loop and starts back at fresh
		}
		if index>2{
			fmt.Println("breaking at position",index)
			break// the loop automatically terminates heare when index>2
		}
		fmt.Printf("the value at position %v is %v\n",index,value)
	}
}