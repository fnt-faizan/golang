package errors

import "fmt"

func Div(a, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("dividing by zero")
	}
	return a / b, nil
}

// func main() {
// 	fmt.Println("Error handling")
// 	ans, err := Div(10, 0)
// 	fmt.Println(ans, err)
// }
