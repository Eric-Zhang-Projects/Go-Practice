Matrix Multiplication:


 ![image](https://user-images.githubusercontent.com/38845224/52182532-19bae080-27cc-11e9-9ca9-99ea19c879f5.png)
 
Input: 
For numbers in a row of a matrix, use spaces to separate each number. To make a new row, use a comma separated by spaces, and to multiply, use the asterisk character separated by spaces ( * ). After the last inputted number, use the # character to terminate the input. CURRENTLY CANNOT MULTIPLY MORE THAN 2 MATRICES AT ONCE.

Output: 
Numbers in a row will be grouped together in brackets, and each bracket represents another row.
For example, inputting the first multiplication in the picture above will output:

![image](https://user-images.githubusercontent.com/38845224/52182563-4838bb80-27cc-11e9-8129-30d67a3692ce.png)
 
For the second:
 
![image](https://user-images.githubusercontent.com/38845224/52182564-4bcc4280-27cc-11e9-9223-abe6f681b2dd.png)

Notes: 
I wanted to practice reading from the command line so I used fmt.scan to read a line of input. For each token, I checked to see if it was either a new row (comma), multiplication sign, or number to act accordingly. Using slices and append(), I created a slice of all numbers in a row (by scanning until a non numeric token), and then appending these “row slices” together to form a 2d slice representing one of the matrices. Following this process I would multiply the 2 together. In addition, the code will check if the entered input contains a non numeric character, an invalid matrix size (for example a 3 x 3 matrix containing a row of 4 numbers not 3), or if two matrices of incompatible size cannot be multiplied together. (1x2 * 3x4 cannot be multiplied because the number of columns in one must equal the number of rows in the other, in this case 2 does not equal 3) 
