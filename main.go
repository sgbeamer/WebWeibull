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
	ttfs := []float64{1750, 5000}
	ttfc := ttf

	// print ttf before sorting
	fmt.Println("Original failure input ", ttf)
	fmt.Println("Original suspended input ", ttfs)	
	//combine the two slices into 1
	l1 := int (len(ttfs))
	for i1 :=0 ; i1 < l1; i1++{
		ttfc = append(ttfc, ttfs[i1])
		}
	
	fmt.Println("Combined inputs unsorted ", ttfc)

	//sort in order
	sort.Float64s(ttfc)
	sort.Float64s(ttf) //needs to be sorted for Bernard ranking

	//print a blank line then results
	fmt.Println()
	fmt.Println("Sorted input ", ttfc)

	//Rank, Reverse Rank
	rankOrder := []int{}
	revRankOrder := []int{}
	ttfcLen := int(len(ttfc))
	for i2 := 0 ; i2 < ttfcLen ; i2++ {
		rankOrder = append(rankOrder, i2+1)
		revRankOrder = append(revRankOrder, ttfcLen-i2)
		}
	fmt.Println("Rank Order " , rankOrder)
	fmt.Println("Reverse Rank Order " , revRankOrder)
	
	//Calculate i's for Bernard's Median Ranking
	// calculate the ranks and reverse ranks
	// start with 3 float64 arrays  ttf (exists and is sorted), ranks of ttf, reverse ranks of ttf
	// output will be two arrays with values of ranks and reverse ranks associate with failures (in order)
	ttfLen := len(ttf)
	ttfrnk := []float64{}
	ttfrr := []float64{}
	for i := 0 ; i < ttfcLen ; i++ {
		for j:=0 ; j < ttfLen ; j++ {
		if ttf[j] == ttfc[i] {
			ttfrnk = append(ttfrnk , float64(rankOrder[i]))
			ttfrr = append(ttfrr , float64(revRankOrder[i]))
			} 
		}
	}
	fmt.Println()
	
	fmt.Println("Times to Failures " , ttf)
	fmt.Println("Ranking of ttf    " , ttfrnk)
	fmt.Println("Reverse rankning  " , ttfrr)
	
		
	//calculate the adjusted ranks
	ttfcLenf64 := float64(len(ttfc))  //THIS IS N
	ttfLenf64 := float64(ttfLen)  
	adjRank0 := (ttfcLenf64 + 1)/(ttfrr[0] + 1)
	adjRank := []float64{adjRank0}
	for i := 1 ; i < ttfLen ; i++ { //start indexing at 1 because initial value already calculated  
		z := i - 1
		num := (ttfrr[i]*adjRank[z]) + (ttfcLenf64 + 1)
		den := ttfrr[i] + 1
		adjRank0 := num/den
		fmt.Println("Numerator ", num)
		fmt.Println("Denominator " , den)
		fmt.Println("Adjusted Rank value " , i , adjRank0)
		adjRank = append(adjRank , adjRank0)
		}
	fmt.Println("Adjusted Ranks for input to Bernard " , adjRank)
	fmt.Println()
	
	//PROGRAM IS GOOD TO HERE -----------------------------------------------------------------------------------------
	
	// Bernard's Median Rank (i - 0.3)/(N + 0.4)
	// declare the variable
	bmr := []float64{} 	//float64
<<<<<<< HEAD
	b := int(len(ttf)) 	//int
	bf := float64(b) 	//float64
	a := float64(0)		//float64
	for i := 0; i < b; i++{
		a = a + 1
		bmr = append(bmr, ((a - 0.3) / (bf + 0.4)))
=======
	for i := 0; i < ttfLen; i++{
		bmr = append(bmr, ((adjRank[i] - 0.3) / (ttfcLenf64 + 0.4)))    
>>>>>>> f99d381700bc4b42eede5acad9b4d57763c09582
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
	for i := 0; i < ttfLen; i++{
		lnY = math.Log(ttf[i])
		lnX = math.Log(math.Log(1/(1-bmr[i])))
		sumXY = sumXY + (lnX * lnY)
		sumY = sumY + lnY
		sumX = sumX + lnX
		sumX2 = sumX2 + lnX*lnX
		sumY2 = sumY2 + lnY * lnY

	}
	// calculate the other stuff that doesn't need to loop
<<<<<<< HEAD
	avgY := sumY / bf
	avgX := sumX /bf
	rNum := (sumXY - (( sumX * sumY) / bf))
	rDenom := math.Sqrt((sumX2 - (sumX * sumX) / bf) * (sumY2 - (sumY * sumY) / bf))
	beta := (sumXY - ((sumX * sumY) / bf)) / (sumX2 - ((sumX * sumX) / bf))  //checks ok
	r := rNum / rDenom
	r2 := r * r
	outA := avgY - beta*avgX

=======
	avgY := sumY / ttfLenf64	
	avgX := sumX / ttfLenf64
	rNum := (sumXY - (( sumX * sumY) / ttfLenf64))
	rDenom := math.Sqrt((sumX2 - (sumX * sumX) / ttfLenf64) * (sumY2 - (sumY * sumY) / ttfLenf64))
	beta := (sumXY - ((sumX * sumY) / ttfLenf64)) / (sumX2 - ((sumX * sumX) / ttfLenf64))  //checks ok
	r := rNum / rDenom
	r2 := r * r
	outA := avgY - beta*avgX
	eta :=  math.Pow(math.E , outA)
	
>>>>>>> f99d381700bc4b42eede5acad9b4d57763c09582
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
	fmt.Println("r = " , r)
	fmt.Println("r² = " , r2)
<<<<<<< HEAD


=======
	fmt.Println("A = " , outA)
	fmt.Println("Weibull Beta (1/beta) = " , 1/beta)
	fmt.Println("Weibull Eta (e^A)     = " , eta)
	
>>>>>>> f99d381700bc4b42eede5acad9b4d57763c09582
}
