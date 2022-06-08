package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("cmd", "/C", "tasklist > system.txt").Output()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(out))
	}

	out2, err := exec.Command("cmd", "/C", "systeminfo >> system.txt").Output()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(out2))
	}

	out3, err := exec.Command("cmd", "/C", "ipconfig /all >> system.txt").Output()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(out3))
	}

	out4, err := exec.Command("cmd", "/C", "netstat -r >> system.txt").Output()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(out4))
	}

}
