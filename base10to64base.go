// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func decimalTobase64(n int) []byte {
	rep := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+-"
	res := []byte{}
	for n != 0 {
		i := n % 64
		res = append(res, rep[i])
		n /= 64
	}
	return res
}

func main() {
	fmt.Println(string(decimalTobase64(100)))
}
