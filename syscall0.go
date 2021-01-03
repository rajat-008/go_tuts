package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	// fmt.Printf("%d,%d", os.Getpid, os.Getppid)
	os.Chdir("..")
	cmd := exec.Command("ls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Fprintf(os.Stdout, "hello")
	time.Sleep(1000000000)
	// fmt.Printf("%d,%d", os.Getpid, os.Getppid)
}
