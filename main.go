package main

import (
	"fmt"

	"github.com/rohith-prakash/diff/test"
)

func main() {
	ans, err := test.Tester("./testcases/test1.txt", "./testcases/test2.txt")
	if err == nil {
		fmt.Println(ans)
	} else {
		fmt.Println(err)
	}

}
