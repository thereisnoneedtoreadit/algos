package main

import "fmt"

/*
Given a stream of integers and a window size, calculate the moving average of all integers in the sliding window.

Implement the MovingAverage class:

MovingAverage(int size) Initializes the object with the size of the window size.
double next(int val) Returns the moving average of the last size values of the stream.

Example 1:

Input
["MovingAverage", "next", "next", "next", "next"]
[[3], [1], [10], [3], [5]]
Output
[null, 1.0, 5.5, 4.66667, 6.0]

Explanation
MovingAverage movingAverage = new MovingAverage(3);
movingAverage.next(1); // return 1.0 = 1 / 1
movingAverage.next(10); // return 5.5 = (1 + 10) / 2
movingAverage.next(3); // return 4.66667 = (1 + 10 + 3) / 3
movingAverage.next(5); // return 6.0 = (10 + 3 + 5) / 3

Constraints:

1 <= size <= 1000
-105 <= val <= 105
At most 104 calls will be made to next.
*/
func main() {
	movingAverage := Constructor(3)
	fmt.Println(movingAverage.Next(1))  // return 1.0 = 1 / 1
	fmt.Println(movingAverage.Next(10)) // return 5.5 = (1 + 10) / 2
	fmt.Println(movingAverage.Next(3))  // return 4.66667 = (1 + 10 + 3) / 3
	fmt.Println(movingAverage.Next(5))  // return 6.0 = (10 + 3 + 5) / 3
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size:    size,
		numbers: make([]int, 0, size),
		sum:     0,
	}
}

func (average *MovingAverage) Next(val int) float64 {
	if len(average.numbers) == average.size {
		average.sum -= average.numbers[0]
		average.numbers = average.numbers[1:]
	}

	average.sum += val
	average.numbers = append(average.numbers, val)

	return float64(average.sum) / float64(len(average.numbers))
}

type MovingAverage struct {
	size    int
	numbers []int
	sum     int
}
