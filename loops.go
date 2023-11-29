package main
import ("fmt"
"sort"

)
func main(){
	//x:=0
	//for x<=5{
		//fmt.Println("the value of x is:",x)
		//x++
	//}
	// for i:=1;i<=5;i++{
	// 	fmt.Println("the value of x is:",i)
	// }
	names:=[]string{"John","Paul", "Ambe","Zeno","Bless","Che"}
	 sort.Strings(names)
	// for i:=0;i<len(names);i++{
	// 	fmt.Println("the name is:",names[i])
	// }
	// for index, value := range names{
	// 	fmt.Printf("the Name at index %v is %v\n", index,value)
	// }
	for _, value := range names{
		fmt.Printf("the Name is %v\n", value)
	}
}
