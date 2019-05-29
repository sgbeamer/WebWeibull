package main

import (
	"fmt"
	"sort"
	"math"
)

func main() {
	// Declare a slice of type float64 for times to failure
	// need to input the data later - using test data to code the math now

	ttf := []float64{7000, 1500, 4300, 2250, 4000}

	// print ttf before sorting
	fmt.Println("Original input ", ttf)

	//sort in order
	sort.Float64s(ttf)

	//print a blank line then results
	fmt.Println()
	fmt.Println("Sorted input ", ttf)

	// Bernard's Median Rank (i - 0.3)/(N + 0.4)
	// declare the variable
	bmr := []float64{} 	//float64
	b := int(len(ttf)) 	//int
	bf := float64(b) 	//float64
	a := float64(0)		//float64
	for i := 0; i < b; i++{
		a = a + 1
		bmr = append(bmr, ((a - 0.3) / (bf + 0.4)))
		//fmt.Println(bmr[i])
	}
	fmt.Println()
	fmt.Println("The ranks are ", bmr)

	// begin regression calcs
	// convert bmr to lnX = ln(ln(1/(1-bmr[i]))
	// convert ttf to lnY = ln(ttf[i])

	sumXY := float64(0)		//float64
	sumY := float64(0)		//float64
	sumX := float64(0)		//float64
	sumX2 := float64(0)		//float64
	sumY2 := float64(0)
	lnY := float64(0)
	lnX := float64(0)

	//loop through all the calcs
	for i := 0; i < b; i++{
		lnY = math.Log(ttf[i])
		lnX = math.Log(math.Log(1/(1-bmr[i])))
		sumXY = sumXY + (lnX * lnY)
		sumY = sumY + lnY
		sumX = sumX + lnX
		sumX2 = sumX2 + lnX*lnX
		sumY2 = sumY2 + lnY * lnY

	}
	// calculate the other stuff that doesn't need to loop
	avgY := sumY / bf
	avgX := sumX /bf
	rNum := (sumXY - (( sumX * sumY) / bf))
	rDenom := math.Sqrt((sumX2 - (sumX * sumX) / bf) * (sumY2 - (sumY * sumY) / bf))
	beta := (sumXY - ((sumX * sumY) / bf)) / (sumX2 - ((sumX * sumX) / bf))  //checks ok
	r := rNum / rDenom
	r2 := r * r
	outA := avgY - beta*avgX

	// print results for testing
	fmt.Println("Sums of x * y = ", sumXY)
	fmt.Println("Sums of y = ", sumY)
	fmt.Println("Sum of Y^2 = ", sumY2)
	fmt.Println("Average y = ", avgY)
	fmt.Println("Sum of x = ", sumX)
	fmt.Println("Sum of x^2 = ", sumX2)
	fmt.Println("Average x = ", avgX)
	fmt.Println("Factor A = ", outA)
	fmt.Println("calc beta = ", beta)
	fmt.Println("Weibull Beta (1/beta) = " , 1/beta)
	fmt.Println("r = " , r)
	fmt.Println("r² = " , r2)


}
