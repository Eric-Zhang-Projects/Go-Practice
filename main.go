//Sudoku Solver
//Based on backtracking algorithm here: https://medium.com/@ssaurel/build-a-sudoku-solver-in-java-part-1-c308bd511481

package main

import(
"fmt"
)
	var i=0
	var j =0
	
	func isInRow(board [9][9]int, row int, number int) bool{
		for i = 0; i < 9; i++{
			if board[row][i] == number{
				return true
			}
		}
		
		return false
	}
	
	// we check if a possible number is already in a column
	func isInCol(board [9][9]int, col int, number int) bool{
		for i = 0; i < 9; i++{
			if board[i][col] == number{
				return true
			}
		}
		
		return false
	}
	
	// we check if a possible number is in its 3x3 box
	func isInBox(board [9][9]int, row int, col int, number int) bool{
		var r = row - row % 3
		var c = col - col % 3
		
		for i = r; i < r + 3; i++{
			for j = c; j < c + 3; j++{
				if board[i][j] == number{
					return true
				}
			}
		}
		
		return false
	}
	
	// combined method to check if a number possible to a row,col position is ok
	func isOk(board [9][9]int, row int, col int, number int) bool{
		return !isInRow(board, row, number)  &&  !isInCol(board, col, number)  &&  !isInBox(board, row, col, number);
	}
	
	// Solve method. We will use a recursive BackTracking algorithm.
       func solve(board [9][9]int) bool{
       	var row=0
       	var col =0
       	var number=0
        for row = 0; row < 9; row++ {
         for col = 0; col < 9; col++ {
          // we search an empty cell
          if (board[row][col] == 0) {
            // we try possible numbers

            for number = 1; number <= 9; number++ {
              if isOk(board, row, col, number) {
                // number ok. it respects sudoku constraints
                board[row][col] = number
               // fmt.Println(row, col, board[row][col])
                if solve(board) { // we start backtracking recursively
                  return true
                } else { // if not a solution, we empty the cell and we continue
                  board[row][col] = 0
                }
             }
            }
            //fmt.Println("Unsolvable");
            return false // we return false if the starting board cannot be solved
           }
          }
         }
         fmt.Println("Sudoku Grid solution:")
         display(board)
         return true // sudoku solved
	}
	

	//display function for printing the board
	func display(board [9][9]int){
		for i= 0; i < 9; i++{
			for j= 0; j < 9; j++{
				fmt.Print(" ", board[i][j])
			}
		
			fmt.Println()
		}
		
		fmt.Println()
	}
	
	//main function creates an initial board state in a 2d array, with 0s denoting an "empty space"
	func main() {
	var solveMe= [9][9]int {
	{0,0,0,2,6,0,7,0,1},
	{6,8,0,0,7,0,0,9,0},
	{1,9,0,0,0,4,5,0,0},
	{8,2,0,1,0,0,0,4,0},
	{0,0,4,6,0,2,9,0,0},
	{0,5,0,0,0,3,0,2,8},
	{0,0,9,3,0,0,0,7,4},
	{0,4,0,0,5,0,0,3,6},
	{7,0,3,0,1,8,0,0,0}}
	var board [9][9]int;
	for i = 0; i < 9; i++{
		for j = 0; j < 9; j++{
			board[i][j] = solveMe[i][j]
		}
	}
		fmt.Println("Starting sudoku grid to solve")
		display(board)
		
		// we try resolution
		if solve(board)!=true{
			fmt.Println("Unsolvable board")
	}
		}




