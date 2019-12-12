package solutions

func Day2() int {
	intCode := [...]int{1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,9,19,23,2,23,13,27,1,27,9,31,2,31,6,35,1,5,35,39,1,10,39,43,2,43,6,47,1,10,47,51,2,6,51,55,1,5,55,59,1,59,9,63,1,13,63,67,2,6,67,71,1,5,71,75,2,6,75,79,2,79,6,83,1,13,83,87,1,9,87,91,1,9,91,95,1,5,95,99,1,5,99,103,2,13,103,107,1,6,107,111,1,9,111,115,2,6,115,119,1,13,119,123,1,123,6,127,1,127,5,131,2,10,131,135,2,135,10,139,1,13,139,143,1,10,143,147,1,2,147,151,1,6,151,0,99,2,14,0,0}

	// you want to teat this as a slice
	// each slice is n+4 inclusive
	// for each slice,
	// opcode = first position
	// opcode case switch 1 = add, 2 = multiply, 99 = halt and return ????, default = throw error
	// first, second, target location in array


	return len(intCode)
}
/*
An Intcode program is a list of integers separated by commas (like 1,0,0,3,99).
To run one, start by looking at the first integer (called position 0). Here, you will
find an opcode - either 1, 2, or 99. The opcode indicates what to do; for example,
99 means that the program is finished and should immediately halt. Encountering an unknown
opcode means something went wrong.

Opcode 1 adds together numbers read from two positions and stores the result in a third
position. The three integers immediately after the opcode tell you these three positions -
the first two indicate the positions from which you should read the input values, and the
third indicates the position at which the output should be stored.

For example, if your Intcode computer encounters 1,10,20,30, it should read the values at
positions 10 and 20, add those values, and then overwrite the value at position 30 with their sum.

Opcode 2 works exactly like opcode 1, except it multiplies the two inputs instead of adding
them. Again, the three integers after the opcode indicate where the inputs and outputs are, not their values.

Once you're done processing an opcode, move to the next one by stepping forward 4 positions.
*/