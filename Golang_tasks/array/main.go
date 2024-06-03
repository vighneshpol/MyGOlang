package main

import "fmt"

func main() {
	fmt.Println("we are learning array")

	var name [5]string
	name[0] = "Vighnesh,"
	name[1] = "Poppy,"
	name[2] = "Visl,"
	name[3] = "Vl,"
	name[4] = "Vighu."

	fmt.Println("Values of array are ", name)

	var arr = []string{"abc", "efg", "sd", "sds", "asdsd", "sdsdsr", "sdseew", "sdswwq"}
	fmt.Println("values of 2nd array are:", arr)
	fmt.Println("length of 2nd array is:", len(arr))
	fmt.Println("capacity of 2nd array is:", cap(arr))
	fmt.Println("Name at the third index is", arr[2])
	fmt.Println("Name at the fourth index is", arr[3])
	fmt.Println("Name at the first index is", arr[0])
	fmt.Println("Name at the second index is", arr[1])
	fmt.Println("Name at the fifth index is", arr[4])
	fmt.Println("Name at the sixth index is", arr[5])
	fmt.Println("Name at the seventh index is", arr[6])

	var array [4]int
	fmt.Println("values of array 3 is ", array)

	arrShorthand := [5]string{"a", "b", "c", "d","e"}
	fmt.Println(arrShorthand)
	for index,value:= range arrShorthand{
		fmt.Printf("Value at index %d is %s\n",index,value)
	}

}
