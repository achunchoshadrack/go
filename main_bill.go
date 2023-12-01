package main
import ("bufio"
"fmt"
"os"
"strings"
"strconv"

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
	switch opt {
	case "a":
		name, _:=getinput("Iterms Name: ",reader)
		price, _:=getinput("Iterms price: ",reader)
		//parsing floats
		p, err:=strconv.ParseFloat(price, 64)
		if err!=nil {
			fmt.Println("the Price must be a number")
			promptopt(b)
		}
		b.Additem(name,p)
		fmt.Printf("Iterm Added%v........%vXAF\n",name,price)
		ch, _:=getinput("Do you want another transaction (y- for yes, n- for no):",reader)
		switch ch{
		case "y":
			promptopt(b)
		case "n":
			fmt.Println("Please save the Information to proceed further")
			promptopt(b)
		default:
			fmt.Println("that was not a valid option please choose a valid option.......")
			promptopt(b)
		}
	case "t":
		tip, _:=getinput("Enter Tip:",reader)
		t, err:=strconv.ParseFloat(tip, 64)
		if err!=nil {
			fmt.Println("the Tip must be a number")
			promptopt(b)
		}
		b.UpdateTip(t)
	    fmt.Printf("Tip Aded%v..........%vXAF",":",tip)
		ch, _:=getinput("Do you want another transaction (y- for yes, n- for no):",reader)
		switch ch{
		case "y":
			promptopt(b)
		case "n":
			fmt.Println("Please save the Information to proceed further")
			promptopt(b)
		default:
			fmt.Println("that was not a valid option please choose a valid option.......")
			promptopt(b)
		}
	case "s":
	    fmt.Println("you choose to save the bill",b)
	default:
		fmt.Println("that was not a valid option please choose a valid option.......")
		promptopt(b)
	}
}

func main(){
	mybill:=CreateBill()
	promptopt(mybill)
}