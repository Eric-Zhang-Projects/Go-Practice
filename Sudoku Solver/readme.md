Sudoku Solver adadpted from Java: https://medium.com/@ssaurel/build-a-sudoku-solver-in-java-part-1-c308bd511481

Input: To manipulate the starting board, you will have to manually change it in the main function, as this program does not take any inputs from the command line.
The biggest difference between the two is that in the original Java code, they create a sudoku object and manipulate it. However, as Go does not have objects, after trying to emulate the object using a struct and having difficulties, I decided to simply create a 2d array and pass it throughout the various functions in the code. In addition, I originally attempted to change the code to accept an input text file containing a board to parse and solve, but was having a reoccurring issue of using the os.Open() function, with the file sometimes opening and other times failing to be found in my directory despite providing an absolute path. I decided against using it and just simply went with the original method of initializing the starting board in the main function itself. 