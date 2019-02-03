//Matrix Multiplication 
package main

import(
"fmt"
"strconv"
"os"
)
//Populates 2 slices to be multiplied, then creates a new slice to generate the output matrix
func main(){
	var token=""
	var array [][]int
	var array2 [][]int
	var row []int
	var numRows=0
	var numCols=0
	var maxCols=0
	var firstRow =true
	var numRows2=0
	var numCols2=0
	var maxCols2=0
	fmt.Scan(&token)
	for token!="*"{
		if token!=","{
			i, err:=strconv.Atoi(token)
		if err!=nil{ //handle error
			fmt.Println("Please enter numbers for matrix entries")
			os.Exit(2)
		}
		row=append(row,i) //create rows of the array
		if firstRow==true{
			maxCols++
		}
		numCols++
		} else if token==","{
			firstRow=false
			if firstRow==false && maxCols!=numCols{
				fmt.Println("Invalid Matrix Entry")
				os.Exit(2)
			}
			numRows++
			numCols=0
	array=append(array,row) //adds new row to the array
	row=nil
}
fmt.Scan(&token)
}
if token=="*"{
	array=append(array,row) //adds new row to the array
	row=nil
	numRows++
	numCols=0
}
//second array========================================
firstRow=true
fmt.Scan(&token)
for token!="#"{
	if token!=","{
		i, err:=strconv.Atoi(token)
		if err!=nil{ //handle error
			fmt.Println("Please enter a number for matrix entries")
			os.Exit(2)
		}
		if firstRow==true{
			maxCols2++
		}
		numCols2++
		row=append(row,i) //create rows of the array
		} else if token==","{
			firstRow=false
			if firstRow==false && maxCols2!=numCols2{
				fmt.Println("Invalid Matrix Entry")
				os.Exit(2)
			}
	array2=append(array2,row) //adds new row to the array
	numRows2++
	numCols2=0
	row=nil
		}
	fmt.Scan(&token)
}
if token=="#"{
	array2=append(array2,row) //adds new row to the array
}
numRows2++
numCols2=0
if maxCols!=numRows2{
	fmt.Println("Cannot multiply, columns in Matrix 1 must equal rows in Matrix 2")
	os.Exit(2)
}
	//fmt.Println(array)
	//fmt.Println(array2)
var output [][]int
var outputRow []int
var outputIndex =0
for i:=0; i < numRows; i++{ 
	for j := 0; j < maxCols2; j++{ 
		for k := 0; k < maxCols; k++{ 
			outputIndex+=array[i][k]*array2[k][j]
		}
		outputRow=append(outputRow,outputIndex) 
		outputIndex=0
	} 
	output=append(output, outputRow)
	outputRow=nil
}

fmt.Println(output)
}

