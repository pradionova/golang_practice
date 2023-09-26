package main
import "fmt"
func main8() {
	res := SumOfSlice([]int {1,2,3})
	fmt.Println(res)
}

func IntersectionOfSlices(slice1, slice2 []int) []int {
	length := len(slice1)
	intersection := make([] int, 0)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if slice1[i] != slice2[j] {
				break
			}
			intersection = append(intersection, slice1[i])
		}
	}
	return intersection
}
