package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IntStdin() int {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("正の整数を入力してください: ")
		scanner.Scan()
		input := scanner.Text()

		num, err := strconv.Atoi(input)
		if err != nil || num < 2 {
			fmt.Println("2以上の数値を入力してください。")
			continue
		}

		return num
	}
}

func factorize(n int) []int {
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			return append(append([]int{}, i), factorize(n/i)...)
		}
	}
	return nil
}

func main() {
	number := IntStdin()
	fmt.Println(factorize(number))
}
