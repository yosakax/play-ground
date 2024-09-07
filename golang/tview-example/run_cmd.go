package main

import (
	"fmt"
	// "os/exec"
)

func main() {
	// out, err := exec.Command("az", "account", "show").CombinedOutput()
	// if err != nil {
	// 	// panic(err)
	// 	// fmt.Println(string())
	// 	fmt.Println("err")
	// }
	// fmt.Printf("%s\n", string(out))
	fmt.Println(int16('h') & 0xff)
	fmt.Println(int16('j') & 0xff)
	fmt.Println(int16('k') & 0xff)
}
