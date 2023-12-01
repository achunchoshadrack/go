package main
import ("fmt"
)
//structure decleration
type bill struct{
	name string
	items map[string]float64
	tips float64
}
func NewBill(name string)bill{
	b:=bill{
		name: name,
		items: map[string]float64{"Rice": 1000, "Achu": 1500, "Fufu":2000},
		tips: 0,
	}
	return b
}
// formate bill (recievinf functions)
func (b bill)formate()string{
	fs :="Bill Breakdown\n"
	var total float64
	// list buill
	for k,v:=range b.items{
		fs+=fmt.Sprintf("%-25v.....%v XAF\n",k+": ", v)
		total+=v
	}
	fs+=fmt.Sprintf("%-25v.....%v XAF","Total: ",total)
	return fs

}