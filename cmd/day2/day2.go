package day2

var instructionLength int = 4

var originalIntCodes = []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 13, 1, 19, 1, 9, 19, 23, 2, 23, 13, 27, 1, 27, 9, 31, 2, 31, 6, 35, 1, 5, 35, 39, 1, 10, 39, 43, 2, 43, 6, 47, 1, 10, 47, 51, 2, 6, 51, 55, 1, 5, 55, 59, 1, 59, 9, 63, 1, 13, 63, 67, 2, 6, 67, 71, 1, 5, 71, 75, 2, 6, 75, 79, 2, 79, 6, 83, 1, 13, 83, 87, 1, 9, 87, 91, 1, 9, 91, 95, 1, 5, 95, 99, 1, 5, 99, 103, 2, 13, 103, 107, 1, 6, 107, 111, 1, 9, 111, 115, 2, 6, 115, 119, 1, 13, 119, 123, 1, 123, 6, 127, 1, 127, 5, 131, 2, 10, 131, 135, 2, 135, 10, 139, 1, 13, 139, 143, 1, 10, 143, 147, 1, 2, 147, 151, 1, 6, 151, 0, 99, 2, 14, 0, 0}
var intCodes []int

func Day2(target int) int {
	nounMin := 1
	nounMax := 100
	verbMin := 1
	verbMax := 100
	ans := 0
	currentNoun := nounMin
	currentVerb := verbMin

	for i := nounMin; i < nounMax; i++ {
		for j := verbMin; j < verbMax; j++ {
			ans = runPair(i, j)

			if ans == target {
				currentNoun = i
				currentVerb = j
				break
			}
		}
		if ans == target {
			ans = add(multiply(100, currentNoun), currentVerb)
			break
		}
	}
	return ans
}

func runPair(noun, verb int) int {
	intCodes = make([]int, len(originalIntCodes))
	copy(intCodes, originalIntCodes)

	intCodes[1] = noun
	intCodes[2] = verb

	for instructionPointer := 0; instructionPointer < len(intCodes); instructionPointer += instructionLength {
		if (len(intCodes) - instructionPointer) < instructionLength {
			continue
		}
		op := operation{opcode: intCodes[instructionPointer], noun: intCodes[instructionPointer+1], verb: intCodes[instructionPointer+2], param3: intCodes[instructionPointer+3]}
		if op.opcode != 99 {
			doMath(op)
		}
		if op.opcode == 99 {
			return intCodes[0]
		}
	}
	return intCodes[0]
}

func doMath(op operation) {
	switch op.opcode {
	case 1:
		intCodes[op.param3] = add(intCodes[op.noun], intCodes[op.verb])
	case 2:
		intCodes[op.param3] = multiply(intCodes[op.noun], intCodes[op.verb])
	}
}

func add(first, second int) int {
	return first + second
}

func multiply(first, second int) int {
	return first * second
}

type operation struct {
	opcode int
	noun   int
	verb   int
	param3 int
}
