package main

import (
	"slices"
)

// func HowMuchILoveYou(i int) string {
// 	kelopak := []int{}
// 	for j := 1; j <= i; j++ {
// 		if j%1 == 0 {
// 			fmt.Println("i love you")
// 		}
// 		if j%2 == 0 {
// 			fmt.Println("a little")
// 		}
// 		kelopak = append(kelopak, j)
// 	}
// 	fmt.Println((kelopak))
// 	return ""
// }

// ! count sheep
// func countSheep(num int) string {
// 	if num == 0 {
// 		return ""
// 	}
// 	str := ""
// 	// Your code here!
// 	for i := 1; i <= num; i++ {
// 		str += fmt.Sprintf("%d sheep... ", i)
// 	}
// 	return str
// }

// func Between(a, b int) []int {
// 	// your code here
// 	arrStart := []int{}

// 	for i := a; i <= b; i++ {
// 		arrStart = append(arrStart, i)
// 	}

// 	return arrStart
// }

// func FindMultiples(integer, limit int) []int {
// 	// Your code here!
// 	arrStart := []int{}
// 	for i := limit;
// 	return nil
//   }

// reverse value
func ReverseSeq(n int) []int {
	var tempSlice []int
	for i := 1; i <= n; i++ {
		tempSlice = append(tempSlice, i)
	}
	slices.Reverse(tempSlice)

	// sort.Sort(sort.Reverse(sort.IntSlice(mySlice))
	return tempSlice
}

func main() {
	// print(countSheep(3))
	// Between(1, 4)

	// for i := 1; i <= 10; i++ {
	// 	if i == 8 || i == 6 {
	// 		fmt.Printf("angkot ke %v sedang lembur \n", i)
	// 		continue
	// 	} else if i >= 6 {
	// 		fmt.Printf("angkot ke %v sedang tidak beroperasi \n", i)
	// 		continue
	// 	}
	// 	fmt.Printf("angkot ke %v sedang beroperasi \n", i)
	// }

	// mahasiswa := []string{"mhs1", "mhs2", "mhs3"}
	// for index, value := range mahasiswa {
	// 	fmt.Printf("mahasiwa ke %v adalah %v \n", index+1, value)
	// }

	// for {
	// 	fmt.Println("mahasiwa")
	// }
	ReverseSeq(5)

}
