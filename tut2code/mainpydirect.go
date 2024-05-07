package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("python", "python.py", "foo", "bar")
	fmt.Println(cmd.Args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

// import sys

// for i in range(len(sys.argv)):
//    print str(i) + ": " + sys.argv[i]
