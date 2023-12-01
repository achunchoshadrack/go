package main
import ("bufio"
"fmt"
"os"
"strings"

)
//hepper function
func getinput(prompt string, r  *bufio.Reader)(string, error){
	fmt.Print(prompt)
	input, err:=r.ReadString('\n')
	return strings.TrimSpace(input), err

}

func CreateBill() bill{
reader:=bufio.NewReader(os.Stdin)
// fmt.Print("create a New Bill name:")
// name, _:=reader.ReadString('\n')
// name=strings.TrimSpace(name)
name, _:=getinput("Create a new Bill Name:",reader)
b:=NewBill(name)
fmt.Println("you have created a new bill --",b.name)
return b;
}
func promptopt(b bill){
	reader:=bufio.NewReader(os.Stdin)
	opt, _:=getinput("Choose Option (a-to add Item, s-Save, t-to Add a tip):",reader)
	fmt.Println(opt)
}

func main(){
	mybill:=CreateBill()
	promptopt(mybill)
}