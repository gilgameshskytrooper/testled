package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var (
	raspivid *exec.Cmd
)

func startRecord(command *exec.Cmd) {
	fmt.Println("startRecord")
	starterr := command.Start()
	if starterr != nil {
		fmt.Println("Can't start raspivid")
		fmt.Println(starterr.Error())
	}
}

func stopRecord(command *exec.Cmd) {
	fmt.Println("stopRecord")
	killerr := command.Process.Kill()
	if killerr != nil {
		fmt.Println("Can't kill raspivid")
		fmt.Println(killerr)
	}
}

func main() {
	raspivid = exec.Command("raspivid", "-o", os.Args[1]+".h264", "-t", "1000000000")
	startRecord(raspivid)
	time.Sleep(10 * time.Second)
	stopRecord(raspivid)

}
