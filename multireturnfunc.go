package main
import ("fmt"
"strings"
)
//fuctions that return multiple values 
func getinitials(n string)(string, string){
	s:=strings.ToUpper(n)
	names:=strings.Split(s, " ")
	var initials[]string
	for _,v:=range names{
		initials=append(initials,v[:1])
	}
	if len(initials)>1{
		return initials[0],initials[1]

	}
	return initials[0],"_"

}

func main (){
	fname,sname:=getinitials("Achuncho Shadrack")
	fmt.Println(fname,sname)
	fname1,sname1:=getinitials("Shadrack Mahko")
	fmt.Println(fname1,sname1)
	fname2,sname2:=getinitials("Shadrack")
	fmt.Println(fname2,sname2)

}