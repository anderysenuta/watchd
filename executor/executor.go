package executor

import (
	"log"
	"os"
	"os/exec"
)

func Execute(channel chan string, command string, path string) {
	var cmd *exec.Cmd

	for {
		select {
		case <-channel:
			// stop active process
			if cmd != nil && cmd.Process != nil {
				_ = cmd.Process.Kill()
				cmd.Wait()
			}

			// start new process
			if command != "" {
				cmd = exec.Command("sh", "-c", command)
			} else {
				cmd = exec.Command(path)
			}

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Start(); err != nil {
				log.Printf("Failed to start process: %s", err)
			}
		}
	}
}
