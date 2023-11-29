package main
import ("fmt"
"strings"
"sort"
)
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
	//slice range
	range1:=Names[0:4]
	range2:=Names[2:]
	range3:=Names[:3]
	fmt.Println(range1,range2,range3)
	//appending slice ranges
	range1=append(range1, 80)
	fmt.Println(range1)
	//strings standard librarr
	greetings:="Hello dear Golang"
	fmt.Println(strings.Contains(greetings, "Golang"))
	fmt.Println(strings.Contains(greetings, "shadrack"))
	//replacing a word with another word using the strings std lib
	fmt.Println(strings.ReplaceAll(greetings, "dear", "all"))
	//to Upper case a String
	fmt.Println(strings.ToUpper(greetings))
	// getting the position of a string
    fmt.Println(strings.Index(greetings, "an"))
	
    //sort parkage
	//sorting an array or slice in accending order
    heights:=[]int{2,3,1,5,0,10}
	sort.Ints(heights)
	sort.Ints(Names)
	fmt.Println(Names)
	fmt.Println(heights)
	// searching for an index of the member of an array or a slice
	index:=sort.SearchInts(heights, 5)
	fmt.Println(index)

    //sorting an array or slice in accending order of a string
	boys:=[]string{"john","Paul", "Ambe","Zeno","Bless","Che"}
	sort.Strings(boys)
	fmt.Println(boys)
	//Seraching for a string on a slice or an array
	fmt.Println(sort.SearchStrings(boys, "Che"))
    
	//input output
	fmt.Println(ages,len(ages))
	fmt.Println(ages1,len(ages1))
	fmt.Println(unames,len(unames))
	fmt.Println("enter your age")
	fmt.Scanln(&numOne)
	fmt.Println("enter your name")
	fmt.Scanln(&name1)
	//fmt.Println("my name is ", name1,"and my age is", numOne)
	//printf string formating
	//fmt.Printf("my name is %v and my age is %v \n", name1, numOne)
	//fmt.Printf("my name is %q and my age is %q \n", name1, numOne)
	//fmt.Printf("my name is of type %T\n", name1)
	//sprintf formated string
	//var str= fmt.Sprintf("your name is %v and my age is %v \n", name1, numOne)
	//fmt.Println("the saved string is", str)

}