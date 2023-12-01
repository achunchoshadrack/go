package main
import ("fmt")
//for type int, float, strings, array, struct, boolean
func nameupdate(x string)string{
	x="mac"
	return x;
}
//for type map, slices, functions
func menuupdate(y map[string]float64){
	y["Fried Rice"]=1500.0
}
func main(){
	//for type int, float, strings, array, struct, boolean
	var y string
	var name="Desh"
	y=nameupdate(name)
 	fmt.Println(y)
	//for type map, slices, functions
	menu:=map[string]float64{
		"Achu": 1000,
		"katikati": 700,
	}
	menuupdate(menu)
	fmt.Println(menu)
}