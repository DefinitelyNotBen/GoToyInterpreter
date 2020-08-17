# GoToyInterpreter

A toy interpreter that reads a txt file line by line to compute and interpret the statements as a toy programming language. 

Each line can contain 3 operands at most, separated by spaces. You can assign values to variables, perform basic mathematical functions (addition, subtraction, division, and multiplication). 

The program will only work in integers

If a line only has one element then that element is outputted, if it's an assigned variable then the value is outputted

Example:
________
A = 10

B = 5

C = A + B

C
________

In this example, 15 is outputted

To run the program, put the filepath for the txt file after go run main.go in the command line
