package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// var ipconfig, _ = exec.Command("ipconfig").Output()
	// var ping, _ = exec.Command("ping").Output()
	var tracert, _ = exec.Command("tracert").Output()
	// fmt.Println(string(ipconfig))
	// fmt.Println(string(ping))
	fmt.Println(string(tracert))
}
