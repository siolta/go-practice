// can use os/exec to replace a go process with another one, non-go or otherwise
package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// go requires the absolute path to the binary you want to exec
	// can usp exec.LookPath() similar to which
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	
	// Exec requires arguments in slice form
	args := []string{"ls", "-a", "-l", "-h"}

	// Exec also needs env vors to use
	env := os.Environ()

	// actual call to syscall.Exec, success ends current process and replaces it
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
