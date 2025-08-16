package main

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
// func ReverseSeq(n int) []int {
// 	var tempSlice []int
// 	for i := 1; i <= n; i++ {
// 		tempSlice = append(tempSlice, i)
// 	}
// 	slices.Reverse(tempSlice)

// 	// sort.Sort(sort.Reverse(sort.IntSlice(mySlice))
// 	return tempSlice
// }

//! Grasshopper - Check for factor

// func CheckForFactor(base int, factor int) bool {
//   if  base % factor == 0 {
// 	fmt.Println(base % factor)
//     return true
//   }

//   return false
// }

//! Litres
// func Litres(time float64) int {
//   //The fun begins here.
//   return int(time * 0.5)
// }

//! SquareSum
// func SquareSum(numbers []int) int {
//     // your code here
// 	var total int
// 	var total2 int
// 	for _,value := range numbers {
// 		total += value * 2
// 		total2 += value * value
// 	}
//   return total2
// }

//! makes negaitve

// func MakeNegative(x int) int {
// 	if  x < 0 {
// 		return x
// 	}

// 	return -x
// }

//! sortting slice

func SortNumbers(numbers []int) []int {
  if numbers == nil {
    return nil
  }
  fmt.println("numbers before sort:", numbers)
//   var angka = numbers.Sort()
  return numbers
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
	// ReverseSeq(5)
	// fmt.Println(CheckForFactor(10, 2))
	// fmt.Println(CheckForFactor(63, 7))

	// fmt.Println(Litres(1.4))
	// fmt.Println(SquareSum([]int{1,2}))


	// fmt.Println(MakeNegative(9))
	// fmt.Println(MakeNegative(-606))
	// fmt.Println(MakeNegative(42))
	// fmt.Println(MakeNegative(0))

	SortNumbers([]int{1, 2, 10, 50, 5})
}
