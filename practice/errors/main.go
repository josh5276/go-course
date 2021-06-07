// Errors Practice: In this example, we will assign two errors to a variable, typing them as independent types,
//     then we will test the result of validWage against our errors.
// Before you begin:
//   1. Change to the practice directory
//       - cd practice/errors
//   2. Complete the `TO DO` sections of the code
//   3. Check answers in answer/main.go

package main

import (
	"errors"
	"fmt"
)

var (
	// TODO: 1. Create two custom errors, 1 when the wage is too low and 1 when the wage is too high
)

func main() {
	// Wages is a pre-defined set of values to pass into the validWage function. Feel free
	// to modify these values to test with different numbers.
	wages := []float32{
		15.25,
		1200,
		7.25,
		5.25,
	}

	for _, wage := range wages {
		if err := validWage(wage); err != nil {
			switch err {

			case /* Error1 */:
				// TODO: 4. Catch your first error and print a message stating why we threw the error.

			case /* Error1 */:
				// TODO: 5. Catch your second error and print a message stating why we threw the error.
			}
			continue
		}
		// TODO: 6. Print a message if the wage is valid
	}

}

func validWage(f float32) error {
	if f < 7.25 {
		// TODO: 2. Return your custom error for a wage too low.
		return /* Error1 */
	}
	if f > 1000 {
		// TODO: 3. Return your custom error for a wage too high.
		return /* Error1 */
	}
	return nil
}

