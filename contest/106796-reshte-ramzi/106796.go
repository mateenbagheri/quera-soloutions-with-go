package main

import "fmt"

func encrypt(n, k int, inpStr string) (string, error) {
	r := []rune(inpStr)
	r = append([]rune{r[n-1]}, r[:n-1]...)
	for i := 0; i < k; i++ {

		for j := range r {
			if string(r[j]) != "z" {
				r[j] = r[j] + 1
				println()
			} else {
				r[j] = 'a'
			}
		}
	}

	return string(r), nil
}

func main() {
	var n, k int
	var input string
	fmt.Scanf("%d\n", &n)
	fmt.Scanf("%d\n", &k)
	fmt.Scanf("%s\n", &input)

	res, _ := encrypt(n, k, input)
	fmt.Println(res)
}
