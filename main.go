package main

import "fmt"

// Multiply returns the product of two non-negative decimal numbers given as
// strings, without converting them to integers. Grade-school multiplication
// into a digit array: digit i × digit j lands at positions i+j and i+j+1.
func Multiply(num1, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n, m := len(num1), len(num2)
	acc := make([]int, n+m)
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			prod := int(num1[i]-'0')*int(num2[j]-'0') + acc[i+j+1]
			acc[i+j+1] = prod % 10
			acc[i+j] += prod / 10
		}
	}
	out := make([]byte, 0, n+m)
	for k, d := range acc {
		if len(out) == 0 && d == 0 && k < len(acc)-1 {
			continue // skip the single possible leading zero
		}
		out = append(out, byte('0'+d))
	}
	return string(out)
}

func main() {
	fmt.Println(Multiply("123", "456"))
	fmt.Println(Multiply("999", "999"))
}
