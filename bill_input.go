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
		items: map[string]float64{},
		tips: 0,
		
	}
	return b
}
// formate bill (recievinf functions)
func (b *bill)formate()string{
	fs :="Bill Breakdown\n"
	var total float64
	// list buill
	for k,v:=range b.items{
		fs+=fmt.Sprintf("%-25v.....%v XAF\n",k+": ", v)
		total+=v
	}
	fs+=fmt.Sprintf("%-25v.....%v XAF\n","Tip: ",b.tips)
	fs+=fmt.Sprintf("%-25v.....%v XAF\n","Total: ",total+b.tips)
	return fs

}
//update tip
func (b *bill) UpdateTip(tips float64){
	b.tips=tips
}
//add a new item to the bill
func (b *bill) Additem(name string, price float64){
	b.items[name]=price
}