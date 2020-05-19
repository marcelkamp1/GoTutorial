package main

import (
	"fmt"
)

func main() {
	fmt.Println(expressionMatter(1, 10, 3))

}

func expressionMatter(a int, b int, c int) int {
	// your code here

	sum := a + b + c
	multiSum := a*b + c
	multiply := a * b * c
	sumPly := a + b*c
	sumPlyBracket := a * (b + c)
	bracketSumPly := (a + b) * c

	max := []int{sum, multiSum, multiply, sumPly, sumPlyBracket, bracketSumPly}
	fmt.Println(max)
	result := 0

	for _, elem := range max {
		if elem > result {
			result = elem
		}
	}

	return result
}

/*
Given three integers a ,b ,c, return the largest number obtained
after inserting the following operators and brackets: +, *, ()
In other words , try every combination of a,b,c with [*+()] ,
and return the Maximum Obtained

The numbers are always positive.
The numbers are in the range (1  ≤  a, b, c  ≤  10).
You can use the same operation more than once.
It's not necessary to place all the signs and brackets.
Repetition in numbers may occur .
You cannot swap the operands. For instance,
in the given example you cannot get expression (1 + 3) * 2 = 8.

Input >> Output Examples:
expressionsMatter(1,2,3)  ==>  return 9
Explanation:
After placing signs and brackets,
the Maximum value obtained from the expression (1+2) * 3 = 9.
expressionsMatter(1,1,1)  ==>  return 3
Explanation:
After placing signs,
the Maximum value obtained from the expression is 1 + 1 + 1 = 3.
expressionsMatter(9,1,1)  ==>  return 18
Explanation:
After placing signs and brackets,
the Maximum value obtained from the expression is 9 * (1+1) = 18.
*/
