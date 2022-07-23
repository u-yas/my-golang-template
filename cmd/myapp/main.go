package main

import (
	"fmt"

	"github.com/u-yas/my-golang-template/configs"
	"github.com/u-yas/my-golang-template/utils/calc"
)

func main() {
	config := configs.LoadEnv()

	fmt.Println("my app env is ", config.AppEnv)

	sum := calc.Sum(1, 2)

	fmt.Println("sum is ", sum)
}
