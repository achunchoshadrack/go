package main
import ("fmt")
//for type int, float, strings, array, struct, boolean
func nameupdate(x *string){
	*x="mac"
}
func main(){
	var name="DESH"
	//nameupdate(name)
	//fmt.Println("the memory adrees for name is :",&name)
	m:=&name
	fmt.Println(name)
	// fmt.Println("the memory address of m is:", m)
	// fmt.Println("the Value at m is:", *m)
	nameupdate(m)
	fmt.Println(name)
}