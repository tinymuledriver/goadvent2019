package solutions

import "fmt"

func Day1() int {
	moduleMass := [...]int{87819, 115026, 134815, 137411, 67764, 99126, 73336, 66216, 81346, 94695, 76336, 148938, 100089, 67341, 101811, 83239, 58537, 146622, 140006, 95115, 87728, 51664, 93463, 127521, 62195, 135326, 104650, 121170, 142794, 125892, 112521, 81326, 110930, 125273, 70131, 52291, 116316, 50670, 82145, 89869, 55474, 146525, 67064, 118129, 74723, 111269, 128051, 131256, 145221, 71059, 137530, 94041, 92331, 134280, 133517, 59611, 113590, 96394, 64731, 53491, 83163, 56863, 51928, 126075, 92833, 106741, 94873, 97241, 105203, 147315, 108651, 67542, 111622, 83522, 125500, 149284, 70747, 78945, 125322, 141425, 111995, 66892, 131105, 86896, 87588, 140571, 116504, 76218, 146224, 127819, 59032, 102767, 137517, 126448, 141218, 102267, 78692, 96306, 56531, 80841}
	//var accumulator int = 0
	var massWithFuel int = 0
	for _, v := range moduleMass {
		//accumulator += massCalc(v)
		massWithFuel += fuelCost(v)
	}
	//fmt.Printf("Accumulator: %d\n", accumulator)
	fmt.Printf("Fuel cost running total: %d\n", massWithFuel)
	//fmt.Printf("Fuel cost calculated from accumulaor: %d\n", fuelCost(accumulator))
	// return accumulator  // 3348430 is correct for part 1
	return massWithFuel // 5019767 is correct for part 2
}
// should be able to refactor to not use fuel cost, but rename since it contains
// fuel calculation works
func fuelCost(mass int) int {
	newFuelMass := massCalc(mass)

	if newFuelMass < 6 {
		return newFuelMass
	}
	return fuelCost(newFuelMass) + newFuelMass
}
// mass calculation works
func massCalc(mass int) int {
	return (mass / 3) - 2
}

/*
Fuel itself requires fuel just like a module - take its mass, divide by three, round down, and subtract 2.
However, that fuel also requires fuel, and that fuel requires fuel, and so on. Any mass that would
require negative fuel should instead be treated as if it requires zero fuel; the remaining mass, if any,
is instead handled by wishing really hard, which has no mass and is outside the scope of this calculation.

So, for each module mass, calculate its fuel and add it to the total. Then, treat the fuel amount you
just calculated as the input mass and repeat the process, continuing until a fuel requirement is
zero or negative. For example:

- A module of mass 14 requires 2 fuel. This fuel requires no further fuel (2 divided by 3 and rounded
down is 0, which would call for a negative fuel), so the total fuel required is still just 2.

- At first, a module of mass 1969 requires 654 fuel. Then, this fuel requires 216 more fuel (654 / 3 - 2).
216 then requires 70 more fuel, which requires 21 fuel, which requires 5 fuel, which requires no further fuel.
So, the total fuel required for a module of mass 1969 is 654 + 216 + 70 + 21 + 5 = 966.

- The fuel required by a module of mass 100756 and its fuel is:
33583 + 11192 + 3728 + 1240 + 411 + 135 + 43 + 12 + 2 = 50346.

What is the sum of the fuel requirements for all of the modules on your spacecraft when also taking
into account the mass of the added fuel? (Calculate the fuel requirements for each module separately,
then add them all up at the end.)


*/
