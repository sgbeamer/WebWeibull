package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// Declare a slice of type float64 for times to failure
	// need to input the data later - using test data to code the math now

	ttf := []float64{10, 24, 12, 30, 15, 36}

	// print ttf before sorting
	fmt.Println("Original input ", ttf)

	//sort in order
	sort.Float64s(ttf)

	//print a blank line then results
	fmt.Println()
	fmt.Println("Sorted input ", ttf)

	// Bernard's Median Rank (i - 0.3)/(N + 0.4)
	// declare the variable
	bmr := []float64{} //float64
	b := int(len(ttf)) //int
	bf := float64(b)   //float64
	a := float64(0)    //float64
	for i := 0; i < b; i++ {
		a = a + 1
		bmr = append(bmr, ((a - 0.3) / (bf + 0.4)))
		//fmt.Println(bmr[i])
	}
	fmt.Println()
	fmt.Println("The ranks are ", bmr)

	// begin regression calcs
	// convert bmr to X = ln(ln(1/(1-bmr[i]))
	// convert ttf to Y = ln(ttf[i])
	bmrX := []float64{}   //float64
	ttfY := []float64{}   //float64
	multXY := []float64{} //float64
	sumXY := float64(0)   //float64
	sumY := float64(0)    //float64
	sumX := float64(0)    //float64
	lnY := float64(0)
	lnX := float64(0)
	for i := 0; i < b; i++ {
		lnY = math.Log(ttf[i])
		lnX = math.Log(math.Log(1 / (1 - bmr[i])))
		bmrX = append(bmrX, lnX)
		ttfY = append(ttfY, lnY)
		multXY = append(multXY, bmrX[i]*ttfY[i])
		sumXY = sumXY + multXY[i]
		sumY = sumY + ttfY[i]
		sumX = sumX + bmrX[i]
	}
	avgY := sumY / bf
	avgX := sumX / bf

	fmt.Println("MultXY of bmr * ttf = ", multXY)
	fmt.Println("Sums of bmr * ttf = ", sumXY)
	fmt.Println("Sums of ttf = ", sumY)
	fmt.Println("Average y(bmr) = ", avgY)
	fmt.Println("Average x(ttf) = ", avgX)

}
