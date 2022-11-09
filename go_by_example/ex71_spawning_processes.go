// sometimes we want to spawn non go based processes in go
package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	// start with a simple command that takes no args or input
	dateCmd := exec.Command("date")

	// .Output is a helper that handles, running a command and waiting for it to finish
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// more complicated case, here we pipe data to the process on stdin and collect results from stdout
	grepCmd := exec.Command("grep", "hello")

	// explicitly grap input/output pipes, start process, write to it, read to it, and wait to exit
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodby grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	// error checks can be done as usual w/ if err != nil
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// when spawning commands we need to privode explicit command and argument array
	// if you want to spawn a full command with only a string use bash -c
	lsCmd := exec.Command("bash", "-c", "ls -alh")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -alh")
	fmt.Println(string(lsOut))
}
