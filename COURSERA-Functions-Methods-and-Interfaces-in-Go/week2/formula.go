package week2

import (
	"math"
)

func GenDisplaceFn(acceleration float64, initialVelocity float64, 
	initialDisplacement float64) func(float64) float64 {

		returnedFunction := func(time float64) float64{
								return (0.5 * acceleration * (math.Pow(time,2))) +
								(initialVelocity + time) +
								(initialDisplacement) 
							}

		return returnedFunction
}