package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

func KillProcess(namePrefix string) error {
	processes, err := process.Processes()
	if err != nil {
		return err
	}
	for _, p := range processes {
		n, err := p.Name()
		if err != nil {
			return err
		}
		if strings.HasPrefix(n, namePrefix) {
			if p.Pid != int32(os.Getpid()) {
				log.Printf("%v found, kill it", n)
				if err = p.Kill(); err != nil {
					return err
				}
			}
		}
	}
	return fmt.Errorf("process not found")
}
