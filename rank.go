package rank

import (
	"fmt"
	"sort"
	//"math"
)

func main() {
	// Declare a slice of type float64 for times to failure
	// need to input the data later - using test data to code the math now

	ttf := []float64{7000, 1500, 4300, 2250, 4000}
	ttfs := []float64{5000, 1750}
	ttfComb := []float64{}
	//combine the 2 slices for ranking
	c := int(len(ttf)) // prepare to append ttf
	d := int(len(ttfs)) // same for ttfs
	for j:= 0; j < c; j++ {
		ttfComb = append(ttfComb, ttf[j])
		}
	for k := 0; k < d; k++ {
		ttfComb = append(ttfComb, ttfs[k])
		}




	// print ttf before sorting
	fmt.Println("Original input failure ", ttf)
	fmt.Println("Original input suspensions ", ttfs)
	fmt.Println("Combined inputs unsorted ", ttfComb)


	//sort in order
	sort.Float64s(ttfComb)
	//set up for Median Ranking  ONLY RANKING ttf slice
	rank := []float64{}
	rrank := []float64{}

	//print a blank line then results
	fmt.Println()
	fmt.Println("Sorted input ", ttfComb)

	// Bernard's Median Rank (i - 0.3)/(N + 0.4)
	// declare the variable
	bmr := []float64{} 	//float64
	b := int(len(ttfComb)) 	//int
	bf := float64(b) 	//float64
	a := float64(0)		//float64
	for i := 0; i < b; i++{
		a = a + 1
		bmr = append(bmr, ((a - 0.3) / (bf + 0.4)))
		//fmt.Println(bmr[i])
	}
	fmt.Println()
	fmt.Println("The ranks are ", bmr)

}
