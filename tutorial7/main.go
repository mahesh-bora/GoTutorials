package main
import("fmt")

func main(){
	var p *int32 =new(int32)
	var i int32
	fmt.Printf("The address of the pointer p is:%v\n", *p)
	fmt.Printf("The value of i is:%v\n", i)

	fmt.Println("-----------------------------------------")
	p = &i
	*p = 1
	fmt.Printf("The value p to pointer is\n:%v\n", *p)
	fmt.Printf("The value of i is:%v\n", i)

	var k int32 = 2
	i = k
	fmt.Printf("The value of i is:%v\n", i)

	var slice = []int{1,3,2}
	var sliceCopy = slice
	sliceCopy[2] = 5
	fmt.Println(slice)
	fmt.Println(sliceCopy)
}