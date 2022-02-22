package calc

import "github.com/pkg/errors"

func Add(numbers ...int) (error, int) {
	sum := 0
	if len(numbers) < 2 {
		return errors.New("Please provide more then one number"), sum
	}
	for _, num := range numbers {
		sum += num
	}
	return nil, sum
}
